// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/controllerbus/bus/api/api.proto

package bus_api

import (
	context "context"
	fmt "fmt"
	math "math"

	controller "github.com/aperturerobotics/controllerbus/controller"
	exec "github.com/aperturerobotics/controllerbus/controller/exec"
	directive "github.com/aperturerobotics/controllerbus/directive"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Config are configuration arguments.
type Config struct {
	// EnableExecController enables the exec controller API.
	EnableExecController bool     `protobuf:"varint,1,opt,name=enable_exec_controller,json=enableExecController,proto3" json:"enable_exec_controller,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fbbe4e64990f1, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetEnableExecController() bool {
	if m != nil {
		return m.EnableExecController
	}
	return false
}

// GetBusInfoRequest is the request type for GetBusInfo.
type GetBusInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBusInfoRequest) Reset()         { *m = GetBusInfoRequest{} }
func (m *GetBusInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetBusInfoRequest) ProtoMessage()    {}
func (*GetBusInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fbbe4e64990f1, []int{1}
}

func (m *GetBusInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBusInfoRequest.Unmarshal(m, b)
}
func (m *GetBusInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBusInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetBusInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBusInfoRequest.Merge(m, src)
}
func (m *GetBusInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetBusInfoRequest.Size(m)
}
func (m *GetBusInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBusInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBusInfoRequest proto.InternalMessageInfo

// GetBusInfoResponse is the response type for GetBusInfo.
type GetBusInfoResponse struct {
	// RunningControllers is the list of running controllers.
	RunningControllers []*controller.Info `protobuf:"bytes,1,rep,name=running_controllers,json=runningControllers,proto3" json:"running_controllers,omitempty"`
	// RunningDirectives is the list of running directives.
	RunningDirectives    []*directive.DirectiveState `protobuf:"bytes,2,rep,name=running_directives,json=runningDirectives,proto3" json:"running_directives,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *GetBusInfoResponse) Reset()         { *m = GetBusInfoResponse{} }
func (m *GetBusInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetBusInfoResponse) ProtoMessage()    {}
func (*GetBusInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fbbe4e64990f1, []int{2}
}

func (m *GetBusInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBusInfoResponse.Unmarshal(m, b)
}
func (m *GetBusInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBusInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetBusInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBusInfoResponse.Merge(m, src)
}
func (m *GetBusInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetBusInfoResponse.Size(m)
}
func (m *GetBusInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBusInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBusInfoResponse proto.InternalMessageInfo

func (m *GetBusInfoResponse) GetRunningControllers() []*controller.Info {
	if m != nil {
		return m.RunningControllers
	}
	return nil
}

func (m *GetBusInfoResponse) GetRunningDirectives() []*directive.DirectiveState {
	if m != nil {
		return m.RunningDirectives
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "bus.api.Config")
	proto.RegisterType((*GetBusInfoRequest)(nil), "bus.api.GetBusInfoRequest")
	proto.RegisterType((*GetBusInfoResponse)(nil), "bus.api.GetBusInfoResponse")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/controllerbus/bus/api/api.proto", fileDescriptor_cb9fbbe4e64990f1)
}

var fileDescriptor_cb9fbbe4e64990f1 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0xc7, 0x8d, 0x42, 0x95, 0x08, 0x62, 0xd3, 0x22, 0x75, 0xbd, 0x94, 0x3d, 0x68, 0x4f, 0xa9,
	0x54, 0x6f, 0x82, 0x60, 0x3f, 0xa8, 0x5e, 0xdb, 0x07, 0x28, 0x9b, 0x38, 0xad, 0x81, 0x9a, 0xac,
	0x99, 0xa4, 0xf4, 0x75, 0x7c, 0x19, 0x9f, 0x4b, 0xba, 0x1f, 0xcd, 0x16, 0x05, 0x8b, 0x87, 0x59,
	0x66, 0xe7, 0xe3, 0x37, 0xf9, 0xcf, 0xd0, 0x87, 0x85, 0x72, 0x6f, 0x5e, 0x70, 0x69, 0xde, 0xbb,
	0x49, 0x0a, 0xd6, 0x79, 0x0b, 0xd6, 0x08, 0xe3, 0x94, 0xc4, 0xae, 0x34, 0xda, 0x59, 0xb3, 0x5c,
	0x82, 0x15, 0x1e, 0xbb, 0x1b, 0x4b, 0x52, 0xb5, 0x31, 0x9e, 0x5a, 0xe3, 0x0c, 0x3b, 0x16, 0x1e,
	0x79, 0x92, 0xaa, 0x68, 0xb4, 0x3f, 0x25, 0xfc, 0x55, 0xdc, 0x9c, 0x17, 0x0d, 0xff, 0x85, 0x81,
	0x35, 0xc8, 0xec, 0x53, 0x50, 0x06, 0xfb, 0x53, 0x5e, 0x95, 0x05, 0xe9, 0xd4, 0x0a, 0x82, 0x97,
	0x43, 0xe2, 0x47, 0x5a, 0x1b, 0x18, 0x3d, 0x57, 0x0b, 0x76, 0x4f, 0x2f, 0x40, 0x27, 0x62, 0x09,
	0xb3, 0xcd, 0x8c, 0x59, 0x00, 0xb4, 0x48, 0x9b, 0x74, 0x4e, 0x26, 0xcd, 0x3c, 0x3b, 0x5a, 0x83,
	0x1c, 0x6c, 0x73, 0x71, 0x83, 0xd6, 0xc7, 0xe0, 0xfa, 0x1e, 0x5f, 0xf4, 0xdc, 0x4c, 0xe0, 0xc3,
	0x03, 0xba, 0xf8, 0x93, 0x50, 0x56, 0x8d, 0x62, 0x6a, 0x34, 0x02, 0x7b, 0xa2, 0x0d, 0xeb, 0xb5,
	0x56, 0x7a, 0x51, 0xa1, 0x63, 0x8b, 0xb4, 0x8f, 0x3a, 0xa7, 0xbd, 0x73, 0x5e, 0x59, 0x53, 0xd6,
	0xc6, 0x8a, 0xe2, 0x30, 0x0d, 0xd9, 0x33, 0x2d, 0xa3, 0xb3, 0xad, 0x12, 0x6c, 0x1d, 0x66, 0x84,
	0x4b, 0x1e, 0xc4, 0x0d, 0x4b, 0x6f, 0xea, 0x12, 0x07, 0x93, 0x7a, 0xd1, 0xb4, 0x0d, 0x63, 0xef,
	0x8b, 0xd0, 0x66, 0x20, 0xf7, 0x3d, 0x4e, 0xc1, 0xae, 0x94, 0x04, 0x36, 0xa6, 0x34, 0xbc, 0x9d,
	0x45, 0xbc, 0xb8, 0x3d, 0xff, 0x21, 0x33, 0xba, 0xfa, 0x35, 0x97, 0x8b, 0x8d, 0x0f, 0x18, 0xd0,
	0xb3, 0xdd, 0x65, 0xb1, 0xeb, 0xaa, 0xc6, 0xec, 0x92, 0xbb, 0x05, 0x25, 0xf8, 0xe6, 0xcf, 0xba,
	0x72, 0xc8, 0x2d, 0x11, 0xb5, 0xec, 0x90, 0x77, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xfa,
	0xbf, 0x49, 0xe2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ControllerBusServiceClient is the client API for ControllerBusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControllerBusServiceClient interface {
	// GetBusInfo requests information about the controller bus.
	GetBusInfo(ctx context.Context, in *GetBusInfoRequest, opts ...grpc.CallOption) (*GetBusInfoResponse, error)
	// ExecController executes a controller configuration on the bus.
	ExecController(ctx context.Context, in *exec.ExecControllerRequest, opts ...grpc.CallOption) (ControllerBusService_ExecControllerClient, error)
}

type controllerBusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerBusServiceClient(cc grpc.ClientConnInterface) ControllerBusServiceClient {
	return &controllerBusServiceClient{cc}
}

func (c *controllerBusServiceClient) GetBusInfo(ctx context.Context, in *GetBusInfoRequest, opts ...grpc.CallOption) (*GetBusInfoResponse, error) {
	out := new(GetBusInfoResponse)
	err := c.cc.Invoke(ctx, "/bus.api.ControllerBusService/GetBusInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerBusServiceClient) ExecController(ctx context.Context, in *exec.ExecControllerRequest, opts ...grpc.CallOption) (ControllerBusService_ExecControllerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ControllerBusService_serviceDesc.Streams[0], "/bus.api.ControllerBusService/ExecController", opts...)
	if err != nil {
		return nil, err
	}
	x := &controllerBusServiceExecControllerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ControllerBusService_ExecControllerClient interface {
	Recv() (*exec.ExecControllerResponse, error)
	grpc.ClientStream
}

type controllerBusServiceExecControllerClient struct {
	grpc.ClientStream
}

func (x *controllerBusServiceExecControllerClient) Recv() (*exec.ExecControllerResponse, error) {
	m := new(exec.ExecControllerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ControllerBusServiceServer is the server API for ControllerBusService service.
type ControllerBusServiceServer interface {
	// GetBusInfo requests information about the controller bus.
	GetBusInfo(context.Context, *GetBusInfoRequest) (*GetBusInfoResponse, error)
	// ExecController executes a controller configuration on the bus.
	ExecController(*exec.ExecControllerRequest, ControllerBusService_ExecControllerServer) error
}

// UnimplementedControllerBusServiceServer can be embedded to have forward compatible implementations.
type UnimplementedControllerBusServiceServer struct {
}

func (*UnimplementedControllerBusServiceServer) GetBusInfo(ctx context.Context, req *GetBusInfoRequest) (*GetBusInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusInfo not implemented")
}
func (*UnimplementedControllerBusServiceServer) ExecController(req *exec.ExecControllerRequest, srv ControllerBusService_ExecControllerServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecController not implemented")
}

func RegisterControllerBusServiceServer(s *grpc.Server, srv ControllerBusServiceServer) {
	s.RegisterService(&_ControllerBusService_serviceDesc, srv)
}

func _ControllerBusService_GetBusInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerBusServiceServer).GetBusInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bus.api.ControllerBusService/GetBusInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerBusServiceServer).GetBusInfo(ctx, req.(*GetBusInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerBusService_ExecController_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(exec.ExecControllerRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControllerBusServiceServer).ExecController(m, &controllerBusServiceExecControllerServer{stream})
}

type ControllerBusService_ExecControllerServer interface {
	Send(*exec.ExecControllerResponse) error
	grpc.ServerStream
}

type controllerBusServiceExecControllerServer struct {
	grpc.ServerStream
}

func (x *controllerBusServiceExecControllerServer) Send(m *exec.ExecControllerResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ControllerBusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bus.api.ControllerBusService",
	HandlerType: (*ControllerBusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBusInfo",
			Handler:    _ControllerBusService_GetBusInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecController",
			Handler:       _ControllerBusService_ExecController_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/aperturerobotics/controllerbus/bus/api/api.proto",
}
