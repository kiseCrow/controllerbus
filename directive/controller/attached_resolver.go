package controller

import (
	"context"
	"runtime/debug"
	"sync"

	"github.com/aperturerobotics/controllerbus/directive"
	"github.com/pkg/errors"
)

// attachedResolver tracks a resolver and its values.
// it also manages starting and stopping the resolver.
type attachedResolver struct {
	diMtx sync.Mutex
	di    *DirectiveInstance

	valsMtx sync.Mutex
	vals    []uint32

	wakeMtx         sync.Mutex
	res             directive.Resolver
	resolutionCtxCh chan context.Context
	wakeResolution  func()
}

// newAttachedResolver constructs a new attachedResolver.
func newAttachedResolver(di *DirectiveInstance, res directive.Resolver) *attachedResolver {
	return &attachedResolver{
		di:              di,
		res:             res,
		resolutionCtxCh: make(chan context.Context, 1),
	}
}

// pushHandlerContext pushes the next handler context.
func (r *attachedResolver) pushHandlerContext(ctx context.Context) {
PushLoop:
	for {
		select {
		case r.resolutionCtxCh <- ctx:
			r.wakeMtx.Lock()
			if r.wakeResolution != nil {
				go r.wakeResolution()
				r.wakeResolution = nil
			}
			r.wakeMtx.Unlock()
			break PushLoop
		default:
		}

		select {
		case <-r.resolutionCtxCh:
		default:
		}
	}
}

// handlerCtx is canceled when the handler is removed or di canceled.
func (r *attachedResolver) execResolver(handlerCtx context.Context) {
	errCh := make(chan error, 1)
	rctx := <-r.resolutionCtxCh

	// the running resolver count is incremented before execResolver is called.
	// rctx is the context for all currently running resolvers, may be canceled
	// resolutionCtxCh will emit the next rctx context when the previous is canceled
	// handlerCtx is the context for this specific handler and may be canceled separately from rctx

	// Resolve() needs to exit if either rctx or handlerCtx is canceled.
	// Create a new sub-context and cancel it when execOnce returns.

	for {
		nctx, nctxCancel := context.WithCancel(handlerCtx)
		go func(ctx context.Context) {
			// if skipCtx then rctx was already canceled
			var skipCtx bool
			select {
			case <-ctx.Done():
				skipCtx = true
			default:
			}
			if !skipCtx {
				errCh <- func() (eerr error) {
					defer func() {
						if ferr := recover(); ferr != nil && eerr == nil {
							eerr = errors.Errorf("%v", ferr)
							r.di.le.
								WithError(eerr).
								Errorf("resolver panic with stack:\n%s", debug.Stack())
						}
					}()
					return r.res.Resolve(ctx, r)
				}()
			} else {
				errCh <- nil
			}
			r.di.decrementRunningResolvers()
		}(nctx)

		// TODO: find a way to exit this goroutine here to lower goroutine count
		var gerr bool
		select {
		case err := <-errCh:
			nctxCancel()
			gerr = true
			if err != nil && err != context.Canceled {
				// TODO handle resolver fatal error
				go r.di.le.
					WithError(err).
					Warn("resolver exited with error")
				return
			}
		case <-handlerCtx.Done():
			nctxCancel()
			return
		case <-rctx.Done():
			nctxCancel()
		}

		// ensure we wait for resolve to return
		if !gerr {
			select {
			case <-errCh:
			case <-handlerCtx.Done():
				return
			}
		}

		// wait for signal to restart resolver
		select {
		case <-handlerCtx.Done():
			return
		case rctx = <-r.resolutionCtxCh:
			r.di.incrementRunningResolvers()
		default:
			// allow the routine to exit to avoid goroutine pollution
			r.wakeMtx.Lock()
			r.wakeResolution = func() {
				r.di.incrementRunningResolvers()
				r.execResolver(handlerCtx)
			}
			r.wakeMtx.Unlock()
			return
		}
	}
}

// Constructed: when resolver is returned from a handler.
//
// When DI destroyed:

// AddValue adds a value to the result, returning success and an ID. If
// AddValue returns false, value was rejected. A rejected value should be
// released immediately. If the value limit is reached, the value may not be
// accepted. The value may be accepted, immediately before the resolver is
// canceled (limit reached). It is always safe to call RemoveValue with the
// ID at any time, even if the resolver is cancelled.
func (r *attachedResolver) AddValue(val directive.Value) (uint32, bool) {
	r.diMtx.Lock()
	di := r.di
	r.diMtx.Unlock()

	if di == nil {
		return 0, false
	}

	r.valsMtx.Lock()
	id, accepted := di.emitValue(val)
	if !accepted {
		return 0, accepted
	}

	r.vals = append(r.vals, id)
	r.valsMtx.Unlock()
	return id, accepted
}

// RemoveValue removes a value from the result, returning found.
// It is safe to call this function even if the resolver is canceled.
func (r *attachedResolver) RemoveValue(id uint32) (val directive.Value, found bool) {
	r.diMtx.Lock()
	di := r.di
	r.diMtx.Unlock()

	if di == nil {
		return nil, false
	}

	r.valsMtx.Lock()
	for i, valID := range r.vals {
		if valID == id {
			found = true
			r.vals[i] = r.vals[len(r.vals)-1]
			r.vals[len(r.vals)-1] = 0
			r.vals = r.vals[:len(r.vals)-1]
			break
		}
	}
	r.valsMtx.Unlock()

	if !found {
		return
	}

	return di.purgeEmittedValue(id)
}

// _ is a type assertion.
var _ directive.ResolverHandler = ((*attachedResolver)(nil))
