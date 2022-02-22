// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/controllerbus/controller/exec/exec.proto

package controller_exec

import (
	fmt "fmt"
	math "math"

	controller "github.com/aperturerobotics/controllerbus/controller"
	proto1 "github.com/aperturerobotics/controllerbus/controller/configset/proto"
	proto "github.com/golang/protobuf/proto"
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

// ControllerStatus holds basic status for a controller.
type ControllerStatus int32

const (
	// ControllerStatus_UNKNOWN is unrecognized.
	ControllerStatus_ControllerStatus_UNKNOWN ControllerStatus = 0
	// ControllerStatus_CONFIGURING indicates the controller is configuring.
	ControllerStatus_ControllerStatus_CONFIGURING ControllerStatus = 1
	// ControllerStatus_RUNNING indicates the controller is running.
	ControllerStatus_ControllerStatus_RUNNING ControllerStatus = 2
	// ControllerStatus_ERROR indicates the controller is terminated with an error.
	ControllerStatus_ControllerStatus_ERROR ControllerStatus = 3
)

var ControllerStatus_name = map[int32]string{
	0: "ControllerStatus_UNKNOWN",
	1: "ControllerStatus_CONFIGURING",
	2: "ControllerStatus_RUNNING",
	3: "ControllerStatus_ERROR",
}

var ControllerStatus_value = map[string]int32{
	"ControllerStatus_UNKNOWN":     0,
	"ControllerStatus_CONFIGURING": 1,
	"ControllerStatus_RUNNING":     2,
	"ControllerStatus_ERROR":       3,
}

func (x ControllerStatus) String() string {
	return proto.EnumName(ControllerStatus_name, int32(x))
}

func (ControllerStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_979e75fb0350f746, []int{0}
}

// ExecControllerRequest is a protobuf request to execute a controller.
type ExecControllerRequest struct {
	// ConfigSet is the controller config set to execute.
	ConfigSet *proto1.ConfigSet `protobuf:"bytes,1,opt,name=config_set,json=configSet,proto3" json:"config_set,omitempty"`
	// ConfigSetYaml is optionally the YAML form of config_set to parse.
	// Merged with config_set.
	ConfigSetYaml string `protobuf:"bytes,2,opt,name=config_set_yaml,json=configSetYaml,proto3" json:"config_set_yaml,omitempty"`
	// ConfigSetYamlOverwrite sets if the yaml portion overwrites the proto portion.
	ConfigSetYamlOverwrite bool     `protobuf:"varint,3,opt,name=config_set_yaml_overwrite,json=configSetYamlOverwrite,proto3" json:"config_set_yaml_overwrite,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *ExecControllerRequest) Reset()         { *m = ExecControllerRequest{} }
func (m *ExecControllerRequest) String() string { return proto.CompactTextString(m) }
func (*ExecControllerRequest) ProtoMessage()    {}
func (*ExecControllerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_979e75fb0350f746, []int{0}
}

func (m *ExecControllerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecControllerRequest.Unmarshal(m, b)
}
func (m *ExecControllerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecControllerRequest.Marshal(b, m, deterministic)
}
func (m *ExecControllerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecControllerRequest.Merge(m, src)
}
func (m *ExecControllerRequest) XXX_Size() int {
	return xxx_messageInfo_ExecControllerRequest.Size(m)
}
func (m *ExecControllerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecControllerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecControllerRequest proto.InternalMessageInfo

func (m *ExecControllerRequest) GetConfigSet() *proto1.ConfigSet {
	if m != nil {
		return m.ConfigSet
	}
	return nil
}

func (m *ExecControllerRequest) GetConfigSetYaml() string {
	if m != nil {
		return m.ConfigSetYaml
	}
	return ""
}

func (m *ExecControllerRequest) GetConfigSetYamlOverwrite() bool {
	if m != nil {
		return m.ConfigSetYamlOverwrite
	}
	return false
}

// ExecControllerResponse is a protobuf response stream.
type ExecControllerResponse struct {
	// Id is the configset identifier for this status report.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Status is the controller execution status.
	Status ControllerStatus `protobuf:"varint,2,opt,name=status,proto3,enum=controller.exec.ControllerStatus" json:"status,omitempty"`
	// ControllerInfo may contain the running controller info.
	ControllerInfo *controller.Info `protobuf:"bytes,3,opt,name=controller_info,json=controllerInfo,proto3" json:"controller_info,omitempty"`
	// ErrorInfo may contain the error information.
	ErrorInfo            string   `protobuf:"bytes,4,opt,name=error_info,json=errorInfo,proto3" json:"error_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecControllerResponse) Reset()         { *m = ExecControllerResponse{} }
func (m *ExecControllerResponse) String() string { return proto.CompactTextString(m) }
func (*ExecControllerResponse) ProtoMessage()    {}
func (*ExecControllerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_979e75fb0350f746, []int{1}
}

func (m *ExecControllerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecControllerResponse.Unmarshal(m, b)
}
func (m *ExecControllerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecControllerResponse.Marshal(b, m, deterministic)
}
func (m *ExecControllerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecControllerResponse.Merge(m, src)
}
func (m *ExecControllerResponse) XXX_Size() int {
	return xxx_messageInfo_ExecControllerResponse.Size(m)
}
func (m *ExecControllerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecControllerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecControllerResponse proto.InternalMessageInfo

func (m *ExecControllerResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExecControllerResponse) GetStatus() ControllerStatus {
	if m != nil {
		return m.Status
	}
	return ControllerStatus_ControllerStatus_UNKNOWN
}

func (m *ExecControllerResponse) GetControllerInfo() *controller.Info {
	if m != nil {
		return m.ControllerInfo
	}
	return nil
}

func (m *ExecControllerResponse) GetErrorInfo() string {
	if m != nil {
		return m.ErrorInfo
	}
	return ""
}

func init() {
	proto.RegisterEnum("controller.exec.ControllerStatus", ControllerStatus_name, ControllerStatus_value)
	proto.RegisterType((*ExecControllerRequest)(nil), "controller.exec.ExecControllerRequest")
	proto.RegisterType((*ExecControllerResponse)(nil), "controller.exec.ExecControllerResponse")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/controllerbus/controller/exec/exec.proto", fileDescriptor_979e75fb0350f746)
}

var fileDescriptor_979e75fb0350f746 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x4f, 0x8b, 0xda, 0x40,
	0x18, 0xc6, 0x3b, 0xb1, 0x48, 0xf3, 0x4a, 0x35, 0x0c, 0x54, 0xd2, 0x60, 0x21, 0xf5, 0x50, 0xa4,
	0x87, 0x04, 0xec, 0x29, 0x67, 0x6b, 0x45, 0x0a, 0x09, 0x8c, 0x95, 0xb2, 0xa7, 0x90, 0xc4, 0xd1,
	0x0d, 0xc4, 0x8c, 0x3b, 0x33, 0xd9, 0x75, 0xbf, 0xc3, 0x7e, 0x9d, 0x3d, 0xee, 0x77, 0x5b, 0x32,
	0xfe, 0x49, 0x36, 0xee, 0xc9, 0x4b, 0xc8, 0xfb, 0x3e, 0xcf, 0xf3, 0x9b, 0x07, 0x5e, 0xf8, 0xbd,
	0x49, 0xe5, 0x6d, 0x11, 0x3b, 0x09, 0xdb, 0xba, 0xd1, 0x8e, 0x72, 0x59, 0x70, 0xca, 0x59, 0xcc,
	0x64, 0x9a, 0x08, 0x37, 0x61, 0xb9, 0xe4, 0x2c, 0xcb, 0x28, 0x8f, 0x8b, 0xfa, 0xe4, 0xd2, 0x3d,
	0x4d, 0xd4, 0xc7, 0xd9, 0x71, 0x26, 0x19, 0xee, 0x55, 0x9a, 0x53, 0xae, 0xad, 0xe9, 0x55, 0xd8,
	0x1a, 0x45, 0x71, 0xad, 0x7f, 0xd7, 0x62, 0xd6, 0xe9, 0x46, 0x50, 0xe9, 0x2a, 0x4a, 0x35, 0x1f,
	0xa8, 0xc3, 0x67, 0x04, 0x5f, 0xa6, 0x7b, 0x9a, 0x4c, 0xce, 0x11, 0x42, 0xef, 0x0a, 0x2a, 0x24,
	0xf6, 0x00, 0x0e, 0xe6, 0x50, 0x50, 0x69, 0x22, 0x1b, 0x8d, 0x3a, 0x63, 0xcb, 0x69, 0xe4, 0x9d,
	0x89, 0x9a, 0x17, 0x54, 0x12, 0x3d, 0x39, 0xfd, 0xe2, 0x1f, 0xd0, 0xab, 0xa2, 0xe1, 0x63, 0xb4,
	0xcd, 0x4c, 0xcd, 0x46, 0x23, 0x9d, 0x7c, 0x3e, 0x7b, 0x6e, 0xa2, 0x6d, 0x86, 0x3d, 0xf8, 0xda,
	0xf0, 0x85, 0xec, 0x9e, 0xf2, 0x07, 0x9e, 0x4a, 0x6a, 0xb6, 0x6c, 0x34, 0xfa, 0x44, 0xfa, 0x6f,
	0x12, 0xc1, 0x49, 0x1d, 0xbe, 0x20, 0xe8, 0x37, 0x7b, 0x8b, 0x1d, 0xcb, 0x05, 0xc5, 0x5d, 0xd0,
	0xd2, 0x95, 0x2a, 0xac, 0x13, 0x2d, 0x5d, 0x61, 0x0f, 0xda, 0x42, 0x46, 0xb2, 0x10, 0xaa, 0x44,
	0x77, 0xfc, 0xdd, 0x69, 0x5c, 0xc8, 0xa9, 0x20, 0x0b, 0x65, 0x24, 0xc7, 0x00, 0xf6, 0xa0, 0x76,
	0xcd, 0x30, 0xcd, 0xd7, 0x4c, 0xd5, 0xea, 0x8c, 0x8d, 0x3a, 0x63, 0x9e, 0xaf, 0x19, 0xe9, 0x56,
	0x8b, 0x72, 0xc6, 0xdf, 0x00, 0x28, 0xe7, 0xec, 0x98, 0xfa, 0xa8, 0xda, 0xe8, 0x6a, 0x53, 0xca,
	0x3f, 0x9f, 0x10, 0x18, 0xcd, 0x67, 0xf1, 0x00, 0xcc, 0xe6, 0x2e, 0x5c, 0xfa, 0x7f, 0xfd, 0xe0,
	0xbf, 0x6f, 0x7c, 0xc0, 0x36, 0x0c, 0x2e, 0xd4, 0x49, 0xe0, 0xff, 0x99, 0xcf, 0x96, 0x64, 0xee,
	0xcf, 0x0c, 0xf4, 0x6e, 0x9e, 0x2c, 0x7d, 0xbf, 0x54, 0x35, 0x6c, 0x41, 0xff, 0x42, 0x9d, 0x12,
	0x12, 0x10, 0xa3, 0x15, 0xb7, 0xd5, 0x35, 0x7f, 0xbd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xf3,
	0x14, 0x34, 0x03, 0x03, 0x00, 0x00,
}
