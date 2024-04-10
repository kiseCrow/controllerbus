// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0-devel
// 	protoc        v5.26.1
// source: github.com/aperturerobotics/controllerbus/bus/api/api.proto

package bus_api

import (
	reflect "reflect"
	sync "sync"

	controller "github.com/aperturerobotics/controllerbus/controller"
	exec "github.com/aperturerobotics/controllerbus/controller/exec"
	directive "github.com/aperturerobotics/controllerbus/directive"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Config are configuration arguments.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// EnableExecController enables the exec controller API.
	EnableExecController bool `protobuf:"varint,1,opt,name=enable_exec_controller,json=enableExecController,proto3" json:"enable_exec_controller,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetEnableExecController() bool {
	if x != nil {
		return x.EnableExecController
	}
	return false
}

// GetBusInfoRequest is the request type for GetBusInfo.
type GetBusInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBusInfoRequest) Reset() {
	*x = GetBusInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBusInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusInfoRequest) ProtoMessage() {}

func (x *GetBusInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusInfoRequest.ProtoReflect.Descriptor instead.
func (*GetBusInfoRequest) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_rawDescGZIP(), []int{1}
}

// GetBusInfoResponse is the response type for GetBusInfo.
type GetBusInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RunningControllers is the list of running controllers.
	RunningControllers []*controller.Info `protobuf:"bytes,1,rep,name=running_controllers,json=runningControllers,proto3" json:"running_controllers,omitempty"`
	// RunningDirectives is the list of running directives.
	RunningDirectives []*directive.DirectiveState `protobuf:"bytes,2,rep,name=running_directives,json=runningDirectives,proto3" json:"running_directives,omitempty"`
}

func (x *GetBusInfoResponse) Reset() {
	*x = GetBusInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBusInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusInfoResponse) ProtoMessage() {}

func (x *GetBusInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusInfoResponse.ProtoReflect.Descriptor instead.
func (*GetBusInfoResponse) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetBusInfoResponse) GetRunningControllers() []*controller.Info {
	if x != nil {
		return x.RunningControllers
	}
	return nil
}

func (x *GetBusInfoResponse) GetRunningDirectives() []*directive.DirectiveState {
	if x != nil {
		return x.RunningDirectives
	}
	return nil
}

var File_github_com_aperturerobotics_controllerbus_bus_api_api_proto protoreflect.FileDescriptor

var file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x75, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62,
	0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x69, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x75,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x44, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x75, 0x73, 0x2f, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x34, 0x0a, 0x16, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x78, 0x65,
	0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x14, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa1, 0x01,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x13, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x12, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x48, 0x0a, 0x12, 0x72, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x11,
	0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x32, 0xc6, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x42, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x42, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x62, 0x75, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x2e,
	0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_rawDescOnce sync.Once
	file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_rawDescData = file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_rawDesc
)

func file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_rawDescGZIP() []byte {
	file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_rawDescOnce.Do(func() {
		file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_rawDescData)
	})
	return file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_rawDescData
}

var file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_goTypes = []interface{}{
	(*Config)(nil),                      // 0: bus.api.Config
	(*GetBusInfoRequest)(nil),           // 1: bus.api.GetBusInfoRequest
	(*GetBusInfoResponse)(nil),          // 2: bus.api.GetBusInfoResponse
	(*controller.Info)(nil),             // 3: controller.Info
	(*directive.DirectiveState)(nil),    // 4: directive.DirectiveState
	(*exec.ExecControllerRequest)(nil),  // 5: controller.exec.ExecControllerRequest
	(*exec.ExecControllerResponse)(nil), // 6: controller.exec.ExecControllerResponse
}
var file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_depIdxs = []int32{
	3, // 0: bus.api.GetBusInfoResponse.running_controllers:type_name -> controller.Info
	4, // 1: bus.api.GetBusInfoResponse.running_directives:type_name -> directive.DirectiveState
	1, // 2: bus.api.ControllerBusService.GetBusInfo:input_type -> bus.api.GetBusInfoRequest
	5, // 3: bus.api.ControllerBusService.ExecController:input_type -> controller.exec.ExecControllerRequest
	2, // 4: bus.api.ControllerBusService.GetBusInfo:output_type -> bus.api.GetBusInfoResponse
	6, // 5: bus.api.ControllerBusService.ExecController:output_type -> controller.exec.ExecControllerResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_init() }
func file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_init() {
	if File_github_com_aperturerobotics_controllerbus_bus_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBusInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBusInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_goTypes,
		DependencyIndexes: file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_depIdxs,
		MessageInfos:      file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_msgTypes,
	}.Build()
	File_github_com_aperturerobotics_controllerbus_bus_api_api_proto = out.File
	file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_rawDesc = nil
	file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_goTypes = nil
	file_github_com_aperturerobotics_controllerbus_bus_api_api_proto_depIdxs = nil
}
