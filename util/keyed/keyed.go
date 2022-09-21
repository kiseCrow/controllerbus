package keyed

import (
	"context"
	"sort"
	"sync"
)

// Routine is a function called as a goroutine.
// If nil is returned, exits cleanly permanently.
// If an error is returned, can be restarted later.
type Routine func(ctx context.Context) error

// Constructor returns a function to start for the given key.
// If nil is returned, skips starting that routine.
type Constructor func(key string) Routine

// Keyed manages a set of goroutines with associated Keys.
type Keyed struct {
	// ctorCb is the constructor callback
	ctorCb Constructor

	// mtx guards below fields
	mtx sync.Mutex
	// ctx is the current root context
	ctx context.Context
	// routines is the set of running routines
	routines map[string]*runningRoutine
}

// NewKeyed constructs a new Keyed execution manager.
// Note: routines won't start until SetContext is called.
func NewKeyed(ctorCb Constructor) *Keyed {
	if ctorCb == nil {
		ctorCb = func(key string) Routine {
			return nil
		}
	}
	return &Keyed{
		ctorCb:   ctorCb,
		routines: make(map[string]*runningRoutine, 1),
	}
}

// SetContext updates the root context, restarting all running routines.
// if restart is true, all errored routines also restart
func (k *Keyed) SetContext(ctx context.Context, restart bool) {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	sameCtx := k.ctx == ctx
	if sameCtx && !restart {
		return
	}

	k.ctx = ctx
	for _, rr := range k.routines {
		if sameCtx && rr.err == nil {
			continue
		}
		if rr.err == nil || restart {
			if rr.ctxCancel != nil {
				rr.ctxCancel()
				rr.ctx, rr.ctxCancel = nil, nil
			}
			rr.start(ctx)
		}
	}
}

// GetKeys returns the list of keys registered with the Keyed instance.
// Note: this is an instantaneous snapshot.
func (k *Keyed) GetKeys() []string {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	keys := make([]string, len(k.routines))
	for k := range k.routines {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// SyncKeys synchronizes the list of running routines with the given list.
// If restart=true, restarts any failed routines in the list.
func (k *Keyed) SyncKeys(keys []string, restart bool) {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	if k.ctx != nil {
		select {
		case <-k.ctx.Done():
			k.ctx = nil
		default:
		}
	}

	routines := make(map[string]*runningRoutine, len(keys))
	for _, key := range keys {
		v := routines[key]
		if v != nil {
			continue
		}
		v, existed := k.routines[key]
		if !existed {
			v = newRunningRoutine(k, k.ctorCb(key))
			k.routines[key] = v
		}
		routines[key] = v
		if !existed || restart {
			if k.ctx != nil {
				v.start(k.ctx)
			}
		}
	}
	for key, rr := range k.routines {
		if _, ok := routines[key]; ok {
			continue
		}
		if rr.ctxCancel != nil {
			rr.ctxCancel()
		}
		delete(k.routines, key)
	}
}

// GetRoutine returns the routine for the given key.
// Note: this is an instantaneous snapshot.
func (k *Keyed) GetRoutine(key string) Routine {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	v, existed := k.routines[key]
	if !existed {
		return nil
	}

	return v.routine
}

// CondResetRoutine checks the condition function, if it returns true, closes
// the routine, constructs a new one, and replaces it (hard reset).
func (k *Keyed) CondResetRoutine(key string, cond func(Routine) bool) (existed bool, reset bool) {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	if k.ctx != nil {
		select {
		case <-k.ctx.Done():
			k.ctx = nil
		default:
		}
	}

	v, existed := k.routines[key]
	if !existed {
		return false, false
	}
	if cond != nil && !cond(v.routine) {
		return true, false
	}

	if v.ctxCancel != nil {
		v.ctxCancel()
	}
	v = newRunningRoutine(k, k.ctorCb(key))
	k.routines[key] = v
	if k.ctx != nil {
		v.start(k.ctx)
	}

	return true, true
}
