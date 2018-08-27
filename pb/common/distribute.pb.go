// Code generated by protoc-gen-go.
// source: distribute.proto
// DO NOT EDIT!

package Carrier

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Distribute_Protocol int32

const (
	Distribute_UNKNOWN  Distribute_Protocol = 0
	Distribute_LOGRT    Distribute_Protocol = 1
	Distribute_LOCATION Distribute_Protocol = 2
	Distribute_REQP     Distribute_Protocol = 3
	Distribute_TEARCHK  Distribute_Protocol = 4
	Distribute_MESSAGE  Distribute_Protocol = 5
	Distribute_CROSS    Distribute_Protocol = 6
	Distribute_FRESET   Distribute_Protocol = 7
	Distribute_LOWPSET  Distribute_Protocol = 8
	Distribute_CTL      Distribute_Protocol = 9
	Distribute_SRVSET   Distribute_Protocol = 10
	Distribute_UPGRADE  Distribute_Protocol = 11
	Distribute_ACK      Distribute_Protocol = 12
)

var Distribute_Protocol_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "LOGRT",
	2:  "LOCATION",
	3:  "REQP",
	4:  "TEARCHK",
	5:  "MESSAGE",
	6:  "CROSS",
	7:  "FRESET",
	8:  "LOWPSET",
	9:  "CTL",
	10: "SRVSET",
	11: "UPGRADE",
	12: "ACK",
}
var Distribute_Protocol_value = map[string]int32{
	"UNKNOWN":  0,
	"LOGRT":    1,
	"LOCATION": 2,
	"REQP":     3,
	"TEARCHK":  4,
	"MESSAGE":  5,
	"CROSS":    6,
	"FRESET":   7,
	"LOWPSET":  8,
	"CTL":      9,
	"SRVSET":   10,
	"UPGRADE":  11,
	"ACK":      12,
}

func (x Distribute_Protocol) String() string {
	return proto.EnumName(Distribute_Protocol_name, int32(x))
}
func (Distribute_Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{0, 0} }

type Distribute struct {
	Header   *Header             `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Protocol Distribute_Protocol `protobuf:"varint,2,opt,name=protocol,enum=Carrier.Distribute_Protocol" json:"protocol,omitempty"`
	Logrt    *DistributeLogRt    `protobuf:"bytes,3,opt,name=logrt" json:"logrt,omitempty"`
	Location *DistributeLocation `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
	Reqp     *DistributeReqp     `protobuf:"bytes,5,opt,name=reqp" json:"reqp,omitempty"`
	Tearchk  *DistributeTearchk  `protobuf:"bytes,6,opt,name=tearchk" json:"tearchk,omitempty"`
	Message  *DistributeMessage  `protobuf:"bytes,7,opt,name=message" json:"message,omitempty"`
	Cross    *DistributeCross    `protobuf:"bytes,8,opt,name=cross" json:"cross,omitempty"`
	Freset   *DistributeFreset   `protobuf:"bytes,9,opt,name=freset" json:"freset,omitempty"`
	Lowpset  *DistributeLowpset  `protobuf:"bytes,10,opt,name=lowpset" json:"lowpset,omitempty"`
	Ctl      *DistributeCtl      `protobuf:"bytes,11,opt,name=ctl" json:"ctl,omitempty"`
	Srvset   *DistributeSrvset   `protobuf:"bytes,12,opt,name=srvset" json:"srvset,omitempty"`
	Upgrade  *DistributeUpgrade  `protobuf:"bytes,13,opt,name=upgrade" json:"upgrade,omitempty"`
	Ack      *DistributeAck      `protobuf:"bytes,14,opt,name=ack" json:"ack,omitempty"`
}

func (m *Distribute) Reset()                    { *m = Distribute{} }
func (m *Distribute) String() string            { return proto.CompactTextString(m) }
func (*Distribute) ProtoMessage()               {}
func (*Distribute) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *Distribute) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Distribute) GetProtocol() Distribute_Protocol {
	if m != nil {
		return m.Protocol
	}
	return Distribute_UNKNOWN
}

func (m *Distribute) GetLogrt() *DistributeLogRt {
	if m != nil {
		return m.Logrt
	}
	return nil
}

func (m *Distribute) GetLocation() *DistributeLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Distribute) GetReqp() *DistributeReqp {
	if m != nil {
		return m.Reqp
	}
	return nil
}

func (m *Distribute) GetTearchk() *DistributeTearchk {
	if m != nil {
		return m.Tearchk
	}
	return nil
}

func (m *Distribute) GetMessage() *DistributeMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Distribute) GetCross() *DistributeCross {
	if m != nil {
		return m.Cross
	}
	return nil
}

func (m *Distribute) GetFreset() *DistributeFreset {
	if m != nil {
		return m.Freset
	}
	return nil
}

func (m *Distribute) GetLowpset() *DistributeLowpset {
	if m != nil {
		return m.Lowpset
	}
	return nil
}

func (m *Distribute) GetCtl() *DistributeCtl {
	if m != nil {
		return m.Ctl
	}
	return nil
}

func (m *Distribute) GetSrvset() *DistributeSrvset {
	if m != nil {
		return m.Srvset
	}
	return nil
}

func (m *Distribute) GetUpgrade() *DistributeUpgrade {
	if m != nil {
		return m.Upgrade
	}
	return nil
}

func (m *Distribute) GetAck() *DistributeAck {
	if m != nil {
		return m.Ack
	}
	return nil
}

func init() {
	proto.RegisterType((*Distribute)(nil), "Carrier.Distribute")
	proto.RegisterEnum("Carrier.Distribute_Protocol", Distribute_Protocol_name, Distribute_Protocol_value)
}

func init() { proto.RegisterFile("distribute.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xef, 0x6e, 0xd3, 0x3c,
	0x18, 0xc5, 0xdf, 0xac, 0xcd, 0x9f, 0x3e, 0xed, 0x3b, 0x2c, 0x0b, 0x36, 0xaf, 0x20, 0x31, 0x0d,
	0x21, 0x86, 0x10, 0xf9, 0x30, 0xae, 0x20, 0xca, 0xb2, 0x0e, 0x35, 0x6b, 0x4a, 0x92, 0xb1, 0x8f,
	0xc8, 0x4d, 0x4d, 0x57, 0x35, 0x90, 0xcc, 0xf1, 0xe0, 0x82, 0xb8, 0x50, 0x90, 0x1d, 0xd3, 0x16,
	0xc9, 0xf9, 0x98, 0xf3, 0xfc, 0x72, 0x8e, 0xed, 0xf3, 0x00, 0x5a, 0xae, 0x1b, 0xc1, 0xd7, 0x8b,
	0x47, 0xc1, 0xfc, 0x9a, 0x57, 0xa2, 0xc2, 0x6e, 0x48, 0x39, 0x5f, 0x33, 0x3e, 0x86, 0x05, 0x6d,
	0xb4, 0x38, 0x3e, 0xda, 0x61, 0x5f, 0xca, 0x6a, 0xc5, 0x85, 0xd6, 0x4f, 0xfe, 0xd1, 0x0b, 0x2a,
	0xd6, 0xd5, 0x77, 0x3d, 0x7a, 0xb6, 0x37, 0xe2, 0xec, 0xa1, 0xd6, 0x32, 0xd9, 0x93, 0x05, 0xa3,
	0xbc, 0xb8, 0xdf, 0x18, 0x26, 0xdf, 0x58, 0xd3, 0xd0, 0x95, 0x29, 0xbd, 0xe0, 0x55, 0xd3, 0x68,
	0xfd, 0x78, 0x4f, 0xff, 0xca, 0x59, 0xc3, 0x84, 0xc1, 0xaa, 0xac, 0x7e, 0xd6, 0xbb, 0xc9, 0xd3,
	0x7d, 0x2b, 0x51, 0x1a, 0x8c, 0x1a, 0xfe, 0xc3, 0x6c, 0xf4, 0x58, 0xaf, 0x38, 0x5d, 0x32, 0x83,
	0x11, 0x2d, 0xf4, 0x1d, 0xce, 0x7e, 0xdb, 0x00, 0x97, 0xdb, 0x01, 0x7e, 0x09, 0xce, 0x3d, 0xa3,
	0x4b, 0xc6, 0x89, 0x75, 0x6a, 0x9d, 0x0f, 0x2f, 0x9e, 0xf8, 0xfa, 0x71, 0xfd, 0x6b, 0x25, 0x63,
	0x1f, 0x3c, 0xf5, 0x63, 0x51, 0x95, 0xe4, 0xe0, 0xd4, 0x3a, 0x3f, 0xbc, 0x78, 0xb1, 0x45, 0x76,
	0x3e, 0xfe, 0x5c, 0x33, 0xf8, 0x0d, 0xd8, 0xea, 0xf9, 0x49, 0x4f, 0xf9, 0x11, 0x03, 0x1c, 0x57,
	0xab, 0x54, 0xe0, 0xf7, 0xe0, 0xfd, 0xed, 0x83, 0xf4, 0x15, 0xfb, 0xdc, 0xc8, 0xb6, 0x08, 0x7e,
	0x0d, 0x7d, 0xd9, 0x11, 0xb1, 0x15, 0x7a, 0x6c, 0x40, 0x53, 0xf6, 0x50, 0xe3, 0x77, 0xe0, 0xea,
	0xce, 0x88, 0xa3, 0xc8, 0xb1, 0x81, 0xcc, 0x5b, 0x42, 0xc2, 0xba, 0x46, 0xe2, 0x76, 0xc2, 0x37,
	0x2d, 0x21, 0x2f, 0xa6, 0x9a, 0x25, 0x5e, 0xe7, 0xc5, 0x42, 0x39, 0xc7, 0x6f, 0xc1, 0x69, 0xab,
	0x26, 0x03, 0x45, 0x9e, 0x18, 0xc8, 0x2b, 0x05, 0xc8, 0x03, 0xe8, 0xf2, 0x09, 0x74, 0x1e, 0x20,
	0x6e, 0x09, 0xfc, 0x0a, 0x7a, 0x85, 0x28, 0xc9, 0x50, 0x81, 0x47, 0xa6, 0x78, 0x51, 0xca, 0xf0,
	0x76, 0x3d, 0xc8, 0xa8, 0x33, 0x3c, 0x53, 0x80, 0x0c, 0xd7, 0x0b, 0x43, 0xfe, 0xef, 0x0c, 0xbf,
	0x6d, 0x09, 0x19, 0x4e, 0x8b, 0x0d, 0x39, 0xec, 0x0c, 0x0f, 0x8a, 0xcd, 0xd9, 0x2f, 0x0b, 0xbc,
	0xed, 0x22, 0x0c, 0xc1, 0xbd, 0x9d, 0x4d, 0x67, 0xc9, 0xdd, 0x0c, 0xfd, 0x87, 0x07, 0x60, 0xc7,
	0xc9, 0x24, 0xcd, 0x91, 0x85, 0x47, 0xe0, 0xc5, 0x49, 0x18, 0xe4, 0x1f, 0x93, 0x19, 0x3a, 0xc0,
	0x1e, 0xf4, 0xd3, 0xe8, 0xd3, 0x1c, 0xf5, 0x24, 0x9f, 0x47, 0x41, 0x1a, 0x5e, 0x4f, 0x51, 0x5f,
	0x7e, 0xdc, 0x44, 0x59, 0x16, 0x4c, 0x22, 0x64, 0xcb, 0x9f, 0xc3, 0x34, 0xc9, 0x32, 0xe4, 0x60,
	0x00, 0xe7, 0x2a, 0x8d, 0xb2, 0x28, 0x47, 0xae, 0x64, 0xe2, 0xe4, 0x6e, 0x2e, 0x3f, 0x3c, 0xec,
	0x42, 0x2f, 0xcc, 0x63, 0x34, 0x90, 0x44, 0x96, 0x7e, 0x96, 0x22, 0xa8, 0x23, 0xcc, 0x27, 0x69,
	0x70, 0x19, 0xa1, 0xa1, 0x24, 0x82, 0x70, 0x8a, 0x46, 0x0b, 0x47, 0xed, 0xf3, 0x87, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x9f, 0x12, 0x3e, 0x73, 0x59, 0x04, 0x00, 0x00,
}
