// Code generated by protoc-gen-go.
// source: donor.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Donor model
type Donor struct {
	Addr          string               `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Id            string               `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Name          string               `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Balance       uint64               `protobuf:"varint,4,opt,name=balance" json:"balance,omitempty"`
	RsaPublicKey  string               `protobuf:"bytes,5,opt,name=rsaPublicKey" json:"rsaPublicKey,omitempty"`
	Contributions []*DonorContribution `protobuf:"bytes,6,rep,name=contributions" json:"contributions,omitempty"`
	Trackings     []*DonorTracking     `protobuf:"bytes,7,rep,name=trackings" json:"trackings,omitempty"`
}

func (m *Donor) Reset()                    { *m = Donor{} }
func (m *Donor) String() string            { return proto.CompactTextString(m) }
func (*Donor) ProtoMessage()               {}
func (*Donor) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Donor) GetContributions() []*DonorContribution {
	if m != nil {
		return m.Contributions
	}
	return nil
}

func (m *Donor) GetTrackings() []*DonorTracking {
	if m != nil {
		return m.Trackings
	}
	return nil
}

// Donor Contribution model
type DonorContribution struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Treaty    string `protobuf:"bytes,2,opt,name=treaty" json:"treaty,omitempty"`
	Amount    uint64 `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *DonorContribution) Reset()                    { *m = DonorContribution{} }
func (m *DonorContribution) String() string            { return proto.CompactTextString(m) }
func (*DonorContribution) ProtoMessage()               {}
func (*DonorContribution) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

// Donor Tracking model
type DonorTracking struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Amount    uint64 `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *DonorTracking) Reset()                    { *m = DonorTracking{} }
func (m *DonorTracking) String() string            { return proto.CompactTextString(m) }
func (*DonorTracking) ProtoMessage()               {}
func (*DonorTracking) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func init() {
	proto.RegisterType((*Donor)(nil), "protos.Donor")
	proto.RegisterType((*DonorContribution)(nil), "protos.DonorContribution")
	proto.RegisterType((*DonorTracking)(nil), "protos.DonorTracking")
}

func init() { proto.RegisterFile("donor.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x51, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x55, 0x3e, 0x9a, 0x2a, 0x57, 0x8a, 0x84, 0x25, 0x90, 0x91, 0x18, 0xa2, 0x4c, 0x99, 0x3a,
	0xd0, 0x1f, 0xc0, 0x00, 0x1b, 0x0b, 0xb2, 0x58, 0x18, 0x2f, 0xb1, 0x85, 0x2c, 0x1a, 0xbb, 0xb2,
	0x2f, 0x43, 0x7f, 0x37, 0x7f, 0x00, 0xc5, 0x31, 0xb4, 0xa1, 0xea, 0xe4, 0xbb, 0xe7, 0x77, 0xcf,
	0xef, 0x9e, 0x61, 0x25, 0xad, 0xb1, 0x6e, 0xb3, 0x77, 0x96, 0x2c, 0x2b, 0xc2, 0xe1, 0xeb, 0xef,
	0x04, 0x16, 0x2f, 0x23, 0xce, 0x18, 0xe4, 0x28, 0xa5, 0xe3, 0x49, 0x95, 0x34, 0xa5, 0x08, 0x35,
	0xbb, 0x86, 0x54, 0x4b, 0x9e, 0x06, 0x24, 0xd5, 0x72, 0xe4, 0x18, 0xec, 0x15, 0xcf, 0x26, 0xce,
	0x58, 0x33, 0x0e, 0xcb, 0x16, 0x77, 0x68, 0x3a, 0xc5, 0xf3, 0x2a, 0x69, 0x72, 0xf1, 0xdb, 0xb2,
	0x1a, 0xae, 0x9c, 0xc7, 0xb7, 0xa1, 0xdd, 0xe9, 0xee, 0x55, 0x1d, 0xf8, 0x22, 0x4c, 0xcd, 0x30,
	0xf6, 0x04, 0xeb, 0xce, 0x1a, 0x72, 0xba, 0x1d, 0x48, 0x5b, 0xe3, 0x79, 0x51, 0x65, 0xcd, 0xea,
	0xf1, 0x7e, 0xb2, 0xe9, 0x37, 0xc1, 0xdb, 0xf3, 0x09, 0x43, 0xcc, 0xf9, 0x6c, 0x0b, 0x25, 0x39,
	0xec, 0xbe, 0xb4, 0xf9, 0xf4, 0x7c, 0x19, 0x86, 0x6f, 0x67, 0xc3, 0xef, 0xf1, 0x56, 0x1c, 0x79,
	0xf5, 0x00, 0x37, 0x67, 0xc2, 0x7f, 0xcb, 0x25, 0x27, 0xcb, 0xdd, 0x41, 0x41, 0x4e, 0x21, 0x1d,
	0x62, 0x08, 0xb1, 0x1b, 0x71, 0xec, 0xed, 0x60, 0x28, 0x44, 0x91, 0x8b, 0xd8, 0xb1, 0x07, 0x28,
	0x49, 0xf7, 0xca, 0x13, 0xf6, 0xfb, 0x10, 0x47, 0x26, 0x8e, 0x40, 0xfd, 0x01, 0xeb, 0x99, 0xa5,
	0x4b, 0x4f, 0x46, 0xe9, 0xf4, 0xb2, 0x74, 0xf6, 0x4f, 0xba, 0x9d, 0xfe, 0x73, 0xfb, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x4e, 0x47, 0x6e, 0xb6, 0xe5, 0x01, 0x00, 0x00,
}
