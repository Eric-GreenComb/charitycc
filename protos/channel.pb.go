// Code generated by protoc-gen-go.
// source: channel.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Channel model
type Channel struct {
	Addr    string        `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Id      string        `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Name    string        `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Balance uint64        `protobuf:"varint,4,opt,name=balance" json:"balance,omitempty"`
	Fees    []*ChannelFee `protobuf:"bytes,5,rep,name=fees" json:"fees,omitempty"`
}

func (m *Channel) Reset()                    { *m = Channel{} }
func (m *Channel) String() string            { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()               {}
func (*Channel) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Channel) GetFees() []*ChannelFee {
	if m != nil {
		return m.Fees
	}
	return nil
}

// ChannelFee model
type ChannelFee struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Amount    uint64 `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *ChannelFee) Reset()                    { *m = ChannelFee{} }
func (m *ChannelFee) String() string            { return proto.CompactTextString(m) }
func (*ChannelFee) ProtoMessage()               {}
func (*ChannelFee) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*Channel)(nil), "protos.Channel")
	proto.RegisterType((*ChannelFee)(nil), "protos.ChannelFee")
}

func init() { proto.RegisterFile("channel.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0xbd, 0xae, 0x82, 0x30,
	0x14, 0xc7, 0x53, 0xe8, 0x85, 0x70, 0x6e, 0xee, 0x1d, 0xce, 0x60, 0x3a, 0x38, 0x10, 0x06, 0xc3,
	0xc4, 0xa0, 0x8f, 0x60, 0xe2, 0x03, 0x74, 0x70, 0x2f, 0xf4, 0x18, 0x49, 0x68, 0x21, 0xb4, 0x3e,
	0x81, 0x2f, 0x6e, 0x68, 0x51, 0x9c, 0xce, 0xff, 0x2b, 0xf9, 0xe5, 0xc0, 0x5f, 0x77, 0x57, 0xd6,
	0xd2, 0xd0, 0x4c, 0xf3, 0xe8, 0x47, 0xcc, 0xc2, 0x71, 0xd5, 0x93, 0x41, 0x7e, 0x8e, 0x0d, 0x22,
	0x70, 0xa5, 0xf5, 0x2c, 0x58, 0xc9, 0xea, 0x42, 0x06, 0x8d, 0xff, 0x90, 0xf4, 0x5a, 0x24, 0x21,
	0x49, 0x7a, 0xbd, 0x6c, 0xac, 0x32, 0x24, 0xd2, 0xb8, 0x59, 0x34, 0x0a, 0xc8, 0x5b, 0x35, 0x28,
	0xdb, 0x91, 0xe0, 0x25, 0xab, 0xb9, 0x7c, 0x5b, 0x3c, 0x00, 0xbf, 0x11, 0x39, 0xf1, 0x53, 0xa6,
	0xf5, 0xef, 0x11, 0x23, 0xdb, 0x35, 0x2b, 0xf0, 0x42, 0x24, 0x43, 0x5f, 0x5d, 0x01, 0xb6, 0xec,
	0xc3, 0x60, 0x5f, 0x8c, 0x1d, 0x64, 0xca, 0x8c, 0x0f, 0xeb, 0x03, 0x99, 0xcb, 0xd5, 0xe1, 0x1e,
	0x0a, 0xdf, 0x1b, 0x72, 0x5e, 0x99, 0x29, 0xd0, 0x53, 0xb9, 0x05, 0x6d, 0xfc, 0xf2, 0xf4, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x48, 0x4f, 0x66, 0x91, 0xfd, 0x00, 0x00, 0x00,
}
