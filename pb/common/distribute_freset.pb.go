// Code generated by protoc-gen-go.
// source: distribute_freset.proto
// DO NOT EDIT!

package Carrier

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DistributeFreset struct {
	// imei
	Imei string `protobuf:"bytes,1,opt,name=imei" json:"imei,omitempty"`
	// 序列号
	Serial string `protobuf:"bytes,2,opt,name=serial" json:"serial,omitempty"`
	// 单位 : 分
	Frequency string `protobuf:"bytes,3,opt,name=frequency" json:"frequency,omitempty"`
}

func (m *DistributeFreset) Reset()                    { *m = DistributeFreset{} }
func (m *DistributeFreset) String() string            { return proto.CompactTextString(m) }
func (*DistributeFreset) ProtoMessage()               {}
func (*DistributeFreset) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *DistributeFreset) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *DistributeFreset) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func (m *DistributeFreset) GetFrequency() string {
	if m != nil {
		return m.Frequency
	}
	return ""
}

func init() {
	proto.RegisterType((*DistributeFreset)(nil), "Carrier.DistributeFreset")
}

func init() { proto.RegisterFile("distribute_freset.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xc9, 0x2c, 0x2e,
	0x29, 0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x8d, 0x4f, 0x2b, 0x4a, 0x2d, 0x4e, 0x2d, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x77, 0x4e, 0x2c, 0x2a, 0xca, 0x4c, 0x2d, 0x52, 0x72, 0xe6, 0x12,
	0x70, 0x81, 0xab, 0x71, 0x03, 0x2b, 0x11, 0xe2, 0xe1, 0x62, 0xc9, 0xcc, 0x4d, 0xcd, 0x94, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x14, 0xe2, 0xe3, 0x62, 0x2b, 0x4e, 0x2d, 0xca, 0x4c, 0xcc, 0x91, 0x60,
	0x02, 0xf3, 0x05, 0xb9, 0x38, 0xd3, 0x8a, 0x52, 0x0b, 0x4b, 0x53, 0xf3, 0x92, 0x2b, 0x25, 0x98,
	0x41, 0x42, 0x49, 0x6c, 0x60, 0x43, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x55, 0x60,
	0x8f, 0x6f, 0x00, 0x00, 0x00,
}
