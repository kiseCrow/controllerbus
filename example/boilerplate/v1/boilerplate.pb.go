// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.19.3
// source: github.com/aperturerobotics/controllerbus/example/boilerplate/v1/boilerplate.proto

package boilerplate_v1

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

// Boilerplate implements the boilerplate directive.
type Boilerplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// MessageText is the message to print with the boilerplate.
	// This is an example field.
	// The keyword "message" prevents us from using that as the field name.
	MessageText string `protobuf:"bytes,1,opt,name=message_text,json=messageText,proto3" json:"message_text,omitempty"`
}

func (x *Boilerplate) Reset() {
	*x = Boilerplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Boilerplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Boilerplate) ProtoMessage() {}

func (x *Boilerplate) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Boilerplate.ProtoReflect.Descriptor instead.
func (*Boilerplate) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_rawDescGZIP(), []int{0}
}

func (x *Boilerplate) GetMessageText() string {
	if x != nil {
		return x.MessageText
	}
	return ""
}

// BoilerplateResult implements the boilerplate directive result.
type BoilerplateResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PrintedLen is the final length of the printed message.
	PrintedLen uint32 `protobuf:"varint,1,opt,name=printed_len,json=printedLen,proto3" json:"printed_len,omitempty"`
}

func (x *BoilerplateResult) Reset() {
	*x = BoilerplateResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoilerplateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoilerplateResult) ProtoMessage() {}

func (x *BoilerplateResult) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoilerplateResult.ProtoReflect.Descriptor instead.
func (*BoilerplateResult) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_rawDescGZIP(), []int{1}
}

func (x *BoilerplateResult) GetPrintedLen() uint32 {
	if x != nil {
		return x.PrintedLen
	}
	return 0
}

var File_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto protoreflect.FileDescriptor

var file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_rawDesc = []byte{
	0x0a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x75, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x22, 0x30, 0x0a, 0x0b, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x65, 0x78, 0x74, 0x22, 0x34, 0x0a, 0x11, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x4c, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_rawDescOnce sync.Once
	file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_rawDescData = file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_rawDesc
)

func file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_rawDescGZIP() []byte {
	file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_rawDescOnce.Do(func() {
		file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_rawDescData)
	})
	return file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_rawDescData
}

var file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_goTypes = []interface{}{
	(*Boilerplate)(nil),       // 0: boilerplate.v1.Boilerplate
	(*BoilerplateResult)(nil), // 1: boilerplate.v1.BoilerplateResult
}
var file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_init()
}
func file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_init() {
	if File_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Boilerplate); i {
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
		file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoilerplateResult); i {
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
			RawDescriptor: file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_goTypes,
		DependencyIndexes: file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_depIdxs,
		MessageInfos:      file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_msgTypes,
	}.Build()
	File_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto = out.File
	file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_rawDesc = nil
	file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_goTypes = nil
	file_github_com_aperturerobotics_controllerbus_example_boilerplate_v1_boilerplate_proto_depIdxs = nil
}
