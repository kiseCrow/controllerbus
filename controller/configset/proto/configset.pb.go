// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/controllerbus/controller/configset/proto/configset.proto

package configset_proto

import (
	fmt "fmt"
	math "math"

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

// ConfigSet contains a configuration set.
type ConfigSet struct {
	// Configurations contains the controller configurations.
	Configurations       map[string]*ControllerConfig `protobuf:"bytes,1,rep,name=configurations,proto3" json:"configurations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ConfigSet) Reset()         { *m = ConfigSet{} }
func (m *ConfigSet) String() string { return proto.CompactTextString(m) }
func (*ConfigSet) ProtoMessage()    {}
func (*ConfigSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab230962ea467981, []int{0}
}

func (m *ConfigSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSet.Unmarshal(m, b)
}
func (m *ConfigSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSet.Marshal(b, m, deterministic)
}
func (m *ConfigSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSet.Merge(m, src)
}
func (m *ConfigSet) XXX_Size() int {
	return xxx_messageInfo_ConfigSet.Size(m)
}
func (m *ConfigSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSet.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSet proto.InternalMessageInfo

func (m *ConfigSet) GetConfigurations() map[string]*ControllerConfig {
	if m != nil {
		return m.Configurations
	}
	return nil
}

// ControllerConfig contains a controller configuration.
type ControllerConfig struct {
	// Id is the config ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Revision is the revision number of the configuration.
	Revision uint64 `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`
	// Config is the configuration object in protobuf form.
	Config               []byte   `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ControllerConfig) Reset()         { *m = ControllerConfig{} }
func (m *ControllerConfig) String() string { return proto.CompactTextString(m) }
func (*ControllerConfig) ProtoMessage()    {}
func (*ControllerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab230962ea467981, []int{1}
}

func (m *ControllerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerConfig.Unmarshal(m, b)
}
func (m *ControllerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerConfig.Marshal(b, m, deterministic)
}
func (m *ControllerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerConfig.Merge(m, src)
}
func (m *ControllerConfig) XXX_Size() int {
	return xxx_messageInfo_ControllerConfig.Size(m)
}
func (m *ControllerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerConfig proto.InternalMessageInfo

func (m *ControllerConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ControllerConfig) GetRevision() uint64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *ControllerConfig) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigSet)(nil), "configset.proto.ConfigSet")
	proto.RegisterMapType((map[string]*ControllerConfig)(nil), "configset.proto.ConfigSet.ConfigurationsEntry")
	proto.RegisterType((*ControllerConfig)(nil), "configset.proto.ControllerConfig")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/controllerbus/controller/configset/proto/configset.proto", fileDescriptor_ab230962ea467981)
}

var fileDescriptor_ab230962ea467981 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xd9, 0x44, 0x8b, 0x9d, 0x4a, 0x2d, 0x2b, 0x48, 0xe8, 0x29, 0xf6, 0x94, 0xd3, 0x06,
	0xea, 0x41, 0xf1, 0x2a, 0xbe, 0xc0, 0x2a, 0xbd, 0xe7, 0xcf, 0x5a, 0x07, 0x63, 0xa6, 0x4c, 0x66,
	0x0b, 0x7d, 0x44, 0xdf, 0x4a, 0x9a, 0xc4, 0x5a, 0x42, 0x6e, 0xf3, 0x9b, 0xfd, 0xbe, 0xf9, 0xb1,
	0xf0, 0xbe, 0x45, 0xf9, 0xf4, 0xb9, 0x29, 0xe8, 0x3b, 0xcd, 0x76, 0x8e, 0xc5, 0xb3, 0x63, 0xca,
	0x49, 0xb0, 0x68, 0xd2, 0x82, 0x6a, 0x61, 0xaa, 0x2a, 0xc7, 0xb9, 0x3f, 0xa7, 0xe3, 0xf8, 0x81,
	0xdb, 0xc6, 0x49, 0xba, 0x63, 0x12, 0xfa, 0x67, 0xd3, 0xb2, 0xbe, 0x19, 0x2c, 0x56, 0x3f, 0x0a,
	0xa6, 0x2f, 0xed, 0xee, 0xcd, 0x89, 0xde, 0xc0, 0xbc, 0x0b, 0x78, 0xce, 0x04, 0xa9, 0x6e, 0x22,
	0x15, 0x87, 0xc9, 0x6c, 0x6d, 0xcc, 0xa0, 0x67, 0x4e, 0x9d, 0x7e, 0xfa, 0x2b, 0xbc, 0xd6, 0xc2,
	0x07, 0x3b, 0xb8, 0xb2, 0x2c, 0xe1, 0x76, 0x24, 0xa6, 0x17, 0x10, 0x7e, 0xb9, 0x43, 0xa4, 0x62,
	0x95, 0x4c, 0xed, 0x71, 0xd4, 0x8f, 0x70, 0xb9, 0xcf, 0x2a, 0xef, 0xa2, 0x20, 0x56, 0xc9, 0x6c,
	0x7d, 0x3f, 0xe6, 0xed, 0xff, 0xda, 0x1d, 0xb4, 0x5d, 0xfe, 0x39, 0x78, 0x52, 0xab, 0x0d, 0x2c,
	0x86, 0xcf, 0x7a, 0x0e, 0x01, 0x96, 0xbd, 0x21, 0xc0, 0x52, 0x2f, 0xe1, 0x8a, 0xdd, 0x1e, 0x1b,
	0xa4, 0xba, 0x75, 0x5c, 0xd8, 0x13, 0xeb, 0x3b, 0x98, 0x74, 0xba, 0x28, 0x8c, 0x55, 0x72, 0x6d,
	0x7b, 0xca, 0x27, 0xad, 0xfa, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x46, 0xb8, 0x7c, 0xe6, 0x93,
	0x01, 0x00, 0x00,
}
