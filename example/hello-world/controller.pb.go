// Code generated by protoc-gen-go-lite. DO NOT EDIT.
// protoc-gen-go-lite version: v0.6.3
// source: github.com/aperturerobotics/controllerbus/example/hello-world/controller.proto

package main

import (
	fmt "fmt"
	io "io"
	strconv "strconv"
	strings "strings"

	protobuf_go_lite "github.com/aperturerobotics/protobuf-go-lite"
	json "github.com/aperturerobotics/protobuf-go-lite/json"
)

// ToyControllerConfig contains configuration for the toy controller.
type ToyControllerConfig struct {
	unknownFields []byte
	// Name is the name to say hello to.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ToyControllerConfig) Reset() {
	*x = ToyControllerConfig{}
}

func (*ToyControllerConfig) ProtoMessage() {}

func (x *ToyControllerConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *ToyControllerConfig) CloneVT() *ToyControllerConfig {
	if m == nil {
		return (*ToyControllerConfig)(nil)
	}
	r := new(ToyControllerConfig)
	r.Name = m.Name
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ToyControllerConfig) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (this *ToyControllerConfig) EqualVT(that *ToyControllerConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ToyControllerConfig) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*ToyControllerConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}

// MarshalProtoJSON marshals the ToyControllerConfig message to JSON.
func (x *ToyControllerConfig) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Name != "" || s.HasField("name") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name")
		s.WriteString(x.Name)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ToyControllerConfig to JSON.
func (x *ToyControllerConfig) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the ToyControllerConfig message from JSON.
func (x *ToyControllerConfig) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "name":
			s.AddField("name")
			x.Name = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the ToyControllerConfig from JSON.
func (x *ToyControllerConfig) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

func (m *ToyControllerConfig) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ToyControllerConfig) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ToyControllerConfig) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ToyControllerConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (x *ToyControllerConfig) MarshalProtoText() string {
	var sb strings.Builder
	sb.WriteString("ToyControllerConfig { ")
	if x.Name != "" {
		sb.WriteString(" name: ")
		sb.WriteString(strconv.Quote(x.Name))
	}
	sb.WriteString("}")
	return sb.String()
}
func (x *ToyControllerConfig) String() string {
	return x.MarshalProtoText()
}
func (m *ToyControllerConfig) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ToyControllerConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ToyControllerConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
