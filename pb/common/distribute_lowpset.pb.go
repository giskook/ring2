// Code generated by protoc-gen-go.
// source: distribute_lowpset.proto
// DO NOT EDIT!

package Carrier

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DistributeLowpset struct {
	// imei
	Imei string `protobuf:"bytes,1,opt,name=imei" json:"imei,omitempty"`
	// 序列号
	Serial string `protobuf:"bytes,2,opt,name=serial" json:"serial,omitempty"`
	// 单位1% 例如 30 表示 30%
	Lowpset string `protobuf:"bytes,3,opt,name=lowpset" json:"lowpset,omitempty"`
}

func (m *DistributeLowpset) Reset()                    { *m = DistributeLowpset{} }
func (m *DistributeLowpset) String() string            { return proto.CompactTextString(m) }
func (*DistributeLowpset) ProtoMessage()               {}
func (*DistributeLowpset) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *DistributeLowpset) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *DistributeLowpset) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func (m *DistributeLowpset) GetLowpset() string {
	if m != nil {
		return m.Lowpset
	}
	return ""
}

func init() {
	proto.RegisterType((*DistributeLowpset)(nil), "Carrier.DistributeLowpset")
}

func init() { proto.RegisterFile("distribute_lowpset.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xc9, 0x2c, 0x2e,
	0x29, 0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x8d, 0xcf, 0xc9, 0x2f, 0x2f, 0x28, 0x4e, 0x2d, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x77, 0x4e, 0x2c, 0x2a, 0xca, 0x4c, 0x2d, 0x52, 0x72, 0xe2,
	0x12, 0x74, 0x81, 0x2b, 0xf2, 0x81, 0xa8, 0x11, 0xe2, 0xe1, 0x62, 0xc9, 0xcc, 0x4d, 0xcd, 0x94,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x14, 0xe2, 0xe3, 0x62, 0x2b, 0x4e, 0x2d, 0xca, 0x4c, 0xcc, 0x91,
	0x60, 0x02, 0xf3, 0xf9, 0xb9, 0xd8, 0xa1, 0x86, 0x49, 0x30, 0x83, 0x04, 0x92, 0xd8, 0xc0, 0x66,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x05, 0xf0, 0x10, 0x6f, 0x00, 0x00, 0x00,
}
