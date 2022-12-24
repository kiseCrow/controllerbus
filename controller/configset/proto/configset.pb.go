// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.9
// source: github.com/aperturerobotics/controllerbus/controller/configset/proto/configset.proto

package configset_proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ConfigSet contains a configuration set.
type ConfigSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configurations contains the controller configurations.
	Configurations map[string]*ControllerConfig `protobuf:"bytes,1,rep,name=configurations,proto3" json:"configurations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConfigSet) Reset() {
	*x = ConfigSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSet) ProtoMessage() {}

func (x *ConfigSet) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSet.ProtoReflect.Descriptor instead.
func (*ConfigSet) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigSet) GetConfigurations() map[string]*ControllerConfig {
	if x != nil {
		return x.Configurations
	}
	return nil
}

// ControllerConfig contains a controller configuration.
type ControllerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id is the config ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Revision is the revision number of the configuration.
	Revision uint64 `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`
	// Config is the configuration object.
	// Supports: protobuf and json (must start with {).
	Config []byte `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ControllerConfig) Reset() {
	*x = ControllerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerConfig) ProtoMessage() {}

func (x *ControllerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerConfig.ProtoReflect.Descriptor instead.
func (*ControllerConfig) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_rawDescGZIP(), []int{1}
}

func (x *ControllerConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ControllerConfig) GetRevision() uint64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *ControllerConfig) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto protoreflect.FileDescriptor

var file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_rawDesc = []byte{
	0x0a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x65, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x65, 0x74, 0x12, 0x56, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x64, 0x0a,
	0x13, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x56, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_rawDescOnce sync.Once
	file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_rawDescData = file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_rawDesc
)

func file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_rawDescGZIP() []byte {
	file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_rawDescOnce.Do(func() {
		file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_rawDescData)
	})
	return file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_rawDescData
}

var file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_goTypes = []interface{}{
	(*ConfigSet)(nil),        // 0: configset.proto.ConfigSet
	(*ControllerConfig)(nil), // 1: configset.proto.ControllerConfig
	nil,                      // 2: configset.proto.ConfigSet.ConfigurationsEntry
}
var file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_depIdxs = []int32{
	2, // 0: configset.proto.ConfigSet.configurations:type_name -> configset.proto.ConfigSet.ConfigurationsEntry
	1, // 1: configset.proto.ConfigSet.ConfigurationsEntry.value:type_name -> configset.proto.ControllerConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_init()
}
func file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_init() {
	if File_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigSet); i {
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
		file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerConfig); i {
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
			RawDescriptor: file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_goTypes,
		DependencyIndexes: file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_depIdxs,
		MessageInfos:      file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_msgTypes,
	}.Build()
	File_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto = out.File
	file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_rawDesc = nil
	file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_goTypes = nil
	file_github_com_aperturerobotics_controllerbus_controller_configset_proto_configset_proto_depIdxs = nil
}
