// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.29
// source: github.com/aperturerobotics/controllerbus/bus/api/api.proto

package bus_api

import (
	context "context"
	errors "errors"

	exec "github.com/aperturerobotics/controllerbus/controller/exec"
	drpc1 "github.com/planetscale/vtprotobuf/codec/drpc"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_github_com_aperturerobotics_controllerbus_bus_api_api_proto struct{}

func (drpcEncoding_File_github_com_aperturerobotics_controllerbus_bus_api_api_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return drpc1.Marshal(msg)
}

func (drpcEncoding_File_github_com_aperturerobotics_controllerbus_bus_api_api_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return drpc1.Unmarshal(buf, msg)
}

type DRPCControllerBusServiceClient interface {
	DRPCConn() drpc.Conn

	GetBusInfo(ctx context.Context, in *GetBusInfoRequest) (*GetBusInfoResponse, error)
	ExecController(ctx context.Context, in *exec.ExecControllerRequest) (DRPCControllerBusService_ExecControllerClient, error)
}

type drpcControllerBusServiceClient struct {
	cc drpc.Conn
}

func NewDRPCControllerBusServiceClient(cc drpc.Conn) DRPCControllerBusServiceClient {
	return &drpcControllerBusServiceClient{cc}
}

func (c *drpcControllerBusServiceClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcControllerBusServiceClient) GetBusInfo(ctx context.Context, in *GetBusInfoRequest) (*GetBusInfoResponse, error) {
	out := new(GetBusInfoResponse)
	err := c.cc.Invoke(ctx, "/bus.api.ControllerBusService/GetBusInfo", drpcEncoding_File_github_com_aperturerobotics_controllerbus_bus_api_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcControllerBusServiceClient) ExecController(ctx context.Context, in *exec.ExecControllerRequest) (DRPCControllerBusService_ExecControllerClient, error) {
	stream, err := c.cc.NewStream(ctx, "/bus.api.ControllerBusService/ExecController", drpcEncoding_File_github_com_aperturerobotics_controllerbus_bus_api_api_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcControllerBusService_ExecControllerClient{stream}
	if err := x.MsgSend(in, drpcEncoding_File_github_com_aperturerobotics_controllerbus_bus_api_api_proto{}); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCControllerBusService_ExecControllerClient interface {
	drpc.Stream
	Recv() (*exec.ExecControllerResponse, error)
}

type drpcControllerBusService_ExecControllerClient struct {
	drpc.Stream
}

func (x *drpcControllerBusService_ExecControllerClient) Recv() (*exec.ExecControllerResponse, error) {
	m := new(exec.ExecControllerResponse)
	if err := x.MsgRecv(m, drpcEncoding_File_github_com_aperturerobotics_controllerbus_bus_api_api_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcControllerBusService_ExecControllerClient) RecvMsg(m *exec.ExecControllerResponse) error {
	return x.MsgRecv(m, drpcEncoding_File_github_com_aperturerobotics_controllerbus_bus_api_api_proto{})
}

type DRPCControllerBusServiceServer interface {
	GetBusInfo(context.Context, *GetBusInfoRequest) (*GetBusInfoResponse, error)
	ExecController(*exec.ExecControllerRequest, DRPCControllerBusService_ExecControllerStream) error
}

type DRPCControllerBusServiceUnimplementedServer struct{}

func (s *DRPCControllerBusServiceUnimplementedServer) GetBusInfo(context.Context, *GetBusInfoRequest) (*GetBusInfoResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCControllerBusServiceUnimplementedServer) ExecController(*exec.ExecControllerRequest, DRPCControllerBusService_ExecControllerStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCControllerBusServiceDescription struct{}

func (DRPCControllerBusServiceDescription) NumMethods() int { return 2 }

func (DRPCControllerBusServiceDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/bus.api.ControllerBusService/GetBusInfo", drpcEncoding_File_github_com_aperturerobotics_controllerbus_bus_api_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCControllerBusServiceServer).
					GetBusInfo(
						ctx,
						in1.(*GetBusInfoRequest),
					)
			}, DRPCControllerBusServiceServer.GetBusInfo, true
	case 1:
		return "/bus.api.ControllerBusService/ExecController", drpcEncoding_File_github_com_aperturerobotics_controllerbus_bus_api_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCControllerBusServiceServer).
					ExecController(
						in1.(*exec.ExecControllerRequest),
						&drpcControllerBusService_ExecControllerStream{in2.(drpc.Stream)},
					)
			}, DRPCControllerBusServiceServer.ExecController, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterControllerBusService(mux drpc.Mux, impl DRPCControllerBusServiceServer) error {
	return mux.Register(impl, DRPCControllerBusServiceDescription{})
}

type DRPCControllerBusService_GetBusInfoStream interface {
	drpc.Stream
	SendAndClose(*GetBusInfoResponse) error
}

type drpcControllerBusService_GetBusInfoStream struct {
	drpc.Stream
}

func (x *drpcControllerBusService_GetBusInfoStream) SendAndClose(m *GetBusInfoResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_github_com_aperturerobotics_controllerbus_bus_api_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCControllerBusService_ExecControllerStream interface {
	drpc.Stream
	Send(*exec.ExecControllerResponse) error
}

type drpcControllerBusService_ExecControllerStream struct {
	drpc.Stream
}

func (x *drpcControllerBusService_ExecControllerStream) Send(m *exec.ExecControllerResponse) error {
	return x.MsgSend(m, drpcEncoding_File_github_com_aperturerobotics_controllerbus_bus_api_api_proto{})
}
