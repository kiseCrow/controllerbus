// Code generated by protoc-gen-srpc. DO NOT EDIT.
// protoc-gen-srpc version: v0.32.9
// source: github.com/aperturerobotics/controllerbus/bus/api/api.proto

package bus_api

import (
	context "context"

	exec "github.com/aperturerobotics/controllerbus/controller/exec"
	srpc "github.com/aperturerobotics/starpc/srpc"
)

type SRPCControllerBusServiceClient interface {
	SRPCClient() srpc.Client

	GetBusInfo(ctx context.Context, in *GetBusInfoRequest) (*GetBusInfoResponse, error)
	ExecController(ctx context.Context, in *exec.ExecControllerRequest) (SRPCControllerBusService_ExecControllerClient, error)
}

type srpcControllerBusServiceClient struct {
	cc        srpc.Client
	serviceID string
}

func NewSRPCControllerBusServiceClient(cc srpc.Client) SRPCControllerBusServiceClient {
	return &srpcControllerBusServiceClient{cc: cc, serviceID: SRPCControllerBusServiceServiceID}
}

func NewSRPCControllerBusServiceClientWithServiceID(cc srpc.Client, serviceID string) SRPCControllerBusServiceClient {
	if serviceID == "" {
		serviceID = SRPCControllerBusServiceServiceID
	}
	return &srpcControllerBusServiceClient{cc: cc, serviceID: serviceID}
}

func (c *srpcControllerBusServiceClient) SRPCClient() srpc.Client { return c.cc }

func (c *srpcControllerBusServiceClient) GetBusInfo(ctx context.Context, in *GetBusInfoRequest) (*GetBusInfoResponse, error) {
	out := new(GetBusInfoResponse)
	err := c.cc.ExecCall(ctx, c.serviceID, "GetBusInfo", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srpcControllerBusServiceClient) ExecController(ctx context.Context, in *exec.ExecControllerRequest) (SRPCControllerBusService_ExecControllerClient, error) {
	stream, err := c.cc.NewStream(ctx, c.serviceID, "ExecController", in)
	if err != nil {
		return nil, err
	}
	strm := &srpcControllerBusService_ExecControllerClient{stream}
	if err := strm.CloseSend(); err != nil {
		return nil, err
	}
	return strm, nil
}

type SRPCControllerBusService_ExecControllerClient interface {
	srpc.Stream
	Recv() (*exec.ExecControllerResponse, error)
	RecvTo(*exec.ExecControllerResponse) error
}

type srpcControllerBusService_ExecControllerClient struct {
	srpc.Stream
}

func (x *srpcControllerBusService_ExecControllerClient) Recv() (*exec.ExecControllerResponse, error) {
	m := new(exec.ExecControllerResponse)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcControllerBusService_ExecControllerClient) RecvTo(m *exec.ExecControllerResponse) error {
	return x.MsgRecv(m)
}

type SRPCControllerBusServiceServer interface {
	GetBusInfo(context.Context, *GetBusInfoRequest) (*GetBusInfoResponse, error)
	ExecController(*exec.ExecControllerRequest, SRPCControllerBusService_ExecControllerStream) error
}

type SRPCControllerBusServiceUnimplementedServer struct{}

func (s *SRPCControllerBusServiceUnimplementedServer) GetBusInfo(context.Context, *GetBusInfoRequest) (*GetBusInfoResponse, error) {
	return nil, srpc.ErrUnimplemented
}

func (s *SRPCControllerBusServiceUnimplementedServer) ExecController(*exec.ExecControllerRequest, SRPCControllerBusService_ExecControllerStream) error {
	return srpc.ErrUnimplemented
}

const SRPCControllerBusServiceServiceID = "bus.api.ControllerBusService"

type SRPCControllerBusServiceHandler struct {
	serviceID string
	impl      SRPCControllerBusServiceServer
}

// NewSRPCControllerBusServiceHandler constructs a new RPC handler.
// serviceID: if empty, uses default: bus.api.ControllerBusService
func NewSRPCControllerBusServiceHandler(impl SRPCControllerBusServiceServer, serviceID string) srpc.Handler {
	if serviceID == "" {
		serviceID = SRPCControllerBusServiceServiceID
	}
	return &SRPCControllerBusServiceHandler{impl: impl, serviceID: serviceID}
}

// SRPCRegisterControllerBusService registers the implementation with the mux.
// Uses the default serviceID: bus.api.ControllerBusService
func SRPCRegisterControllerBusService(mux srpc.Mux, impl SRPCControllerBusServiceServer) error {
	return mux.Register(NewSRPCControllerBusServiceHandler(impl, ""))
}

func (d *SRPCControllerBusServiceHandler) GetServiceID() string { return d.serviceID }

func (SRPCControllerBusServiceHandler) GetMethodIDs() []string {
	return []string{
		"GetBusInfo",
		"ExecController",
	}
}

func (d *SRPCControllerBusServiceHandler) InvokeMethod(
	serviceID, methodID string,
	strm srpc.Stream,
) (bool, error) {
	if serviceID != "" && serviceID != d.GetServiceID() {
		return false, nil
	}

	switch methodID {
	case "GetBusInfo":
		return true, d.InvokeMethod_GetBusInfo(d.impl, strm)
	case "ExecController":
		return true, d.InvokeMethod_ExecController(d.impl, strm)
	default:
		return false, nil
	}
}

func (SRPCControllerBusServiceHandler) InvokeMethod_GetBusInfo(impl SRPCControllerBusServiceServer, strm srpc.Stream) error {
	req := new(GetBusInfoRequest)
	if err := strm.MsgRecv(req); err != nil {
		return err
	}
	out, err := impl.GetBusInfo(strm.Context(), req)
	if err != nil {
		return err
	}
	return strm.MsgSend(out)
}

func (SRPCControllerBusServiceHandler) InvokeMethod_ExecController(impl SRPCControllerBusServiceServer, strm srpc.Stream) error {
	req := new(exec.ExecControllerRequest)
	if err := strm.MsgRecv(req); err != nil {
		return err
	}
	serverStrm := &srpcControllerBusService_ExecControllerStream{strm}
	return impl.ExecController(req, serverStrm)
}

type SRPCControllerBusService_GetBusInfoStream interface {
	srpc.Stream
}

type srpcControllerBusService_GetBusInfoStream struct {
	srpc.Stream
}

type SRPCControllerBusService_ExecControllerStream interface {
	srpc.Stream
	Send(*exec.ExecControllerResponse) error
	SendAndClose(*exec.ExecControllerResponse) error
}

type srpcControllerBusService_ExecControllerStream struct {
	srpc.Stream
}

func (x *srpcControllerBusService_ExecControllerStream) Send(m *exec.ExecControllerResponse) error {
	return x.MsgSend(m)
}

func (x *srpcControllerBusService_ExecControllerStream) SendAndClose(m *exec.ExecControllerResponse) error {
	if m != nil {
		if err := x.MsgSend(m); err != nil {
			return err
		}
	}
	return x.CloseSend()
}
