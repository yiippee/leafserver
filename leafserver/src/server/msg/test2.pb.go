// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test2.proto

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Hello2 struct {
	Id   int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Hello2) Reset()                    { *m = Hello2{} }
func (m *Hello2) String() string            { return proto.CompactTextString(m) }
func (*Hello2) ProtoMessage()               {}
func (*Hello2) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Hello2) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Hello2) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Hello2)(nil), "msg.Hello2")
}

func init() { proto.RegisterFile("test2.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 90 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x49, 0x2d, 0x2e,
	0x31, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0xd2, 0xe1, 0x62,
	0xf3, 0x48, 0xcd, 0xc9, 0xc9, 0x37, 0x12, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0d, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60,
	0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x93, 0xd8, 0xc0, 0x3a, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xea, 0xe2, 0xdf, 0x72, 0x48, 0x00, 0x00, 0x00,
}
