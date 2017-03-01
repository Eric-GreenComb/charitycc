// Code generated by protoc-gen-go.
// source: bargain.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	bargain.proto
	cebcoin.proto
	channel.proto
	donor.proto
	foundation.proto
	process.proto
	smartContract.proto
	test.proto

It has these top-level messages:
	Bargain
	AttachFile
	BargainTrack
	BargainHistory
	Account
	TX
	Channel
	ChannelFeeTrack
	Donor
	DonorContribution
	DonorTrack
	Foundation
	FoundationFeeTrack
	ProcessDonored
	ProcessDrawed
	SmartContract
	SmartContractExt
	SmartContractTrack
	SmartContractHistory
	TestArray
	TestMap
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Bargain model
type Bargain struct {
	Addr        string        `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Id          string        `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Name        string        `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Detail      string        `protobuf:"bytes,4,opt,name=detail" json:"detail,omitempty"`
	StartTime   string        `protobuf:"bytes,5,opt,name=startTime" json:"startTime,omitempty"`
	EndTime     string        `protobuf:"bytes,6,opt,name=endTime" json:"endTime,omitempty"`
	PartyA      string        `protobuf:"bytes,7,opt,name=partyA" json:"partyA,omitempty"`
	PartyB      string        `protobuf:"bytes,8,opt,name=partyB" json:"partyB,omitempty"`
	DepositBank string        `protobuf:"bytes,9,opt,name=depositBank" json:"depositBank,omitempty"`
	BankAccount string        `protobuf:"bytes,10,opt,name=bankAccount" json:"bankAccount,omitempty"`
	AttachFiles []*AttachFile `protobuf:"bytes,11,rep,name=attachFiles" json:"attachFiles,omitempty"`
	Status      int64         `protobuf:"varint,12,opt,name=status" json:"status,omitempty"`
	Remark      string        `protobuf:"bytes,13,opt,name=remark" json:"remark,omitempty"`
}

func (m *Bargain) Reset()                    { *m = Bargain{} }
func (m *Bargain) String() string            { return proto.CompactTextString(m) }
func (*Bargain) ProtoMessage()               {}
func (*Bargain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Bargain) GetAttachFiles() []*AttachFile {
	if m != nil {
		return m.AttachFiles
	}
	return nil
}

type AttachFile struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Desc string `protobuf:"bytes,2,opt,name=desc" json:"desc,omitempty"`
	Hash string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
}

func (m *AttachFile) Reset()                    { *m = AttachFile{} }
func (m *AttachFile) String() string            { return proto.CompactTextString(m) }
func (*AttachFile) ProtoMessage()               {}
func (*AttachFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BargainTrack struct {
	Addr  string            `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Trans []*BargainHistory `protobuf:"bytes,2,rep,name=trans" json:"trans,omitempty"`
}

func (m *BargainTrack) Reset()                    { *m = BargainTrack{} }
func (m *BargainTrack) String() string            { return proto.CompactTextString(m) }
func (*BargainTrack) ProtoMessage()               {}
func (*BargainTrack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BargainTrack) GetTrans() []*BargainHistory {
	if m != nil {
		return m.Trans
	}
	return nil
}

type BargainHistory struct {
	Amount    uint64 `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *BargainHistory) Reset()                    { *m = BargainHistory{} }
func (m *BargainHistory) String() string            { return proto.CompactTextString(m) }
func (*BargainHistory) ProtoMessage()               {}
func (*BargainHistory) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Bargain)(nil), "protos.Bargain")
	proto.RegisterType((*AttachFile)(nil), "protos.AttachFile")
	proto.RegisterType((*BargainTrack)(nil), "protos.BargainTrack")
	proto.RegisterType((*BargainHistory)(nil), "protos.BargainHistory")
}

func init() { proto.RegisterFile("bargain.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0xbd, 0x6e, 0xc2, 0x30,
	0x10, 0x56, 0x7e, 0x80, 0x72, 0x01, 0x06, 0x0f, 0xc8, 0x43, 0x87, 0x28, 0x13, 0x43, 0xc5, 0xd0,
	0xf6, 0x05, 0xc8, 0x80, 0x18, 0xab, 0x88, 0x17, 0x38, 0x62, 0xab, 0x58, 0x21, 0x4e, 0x64, 0x1f,
	0x03, 0xcf, 0xdc, 0x97, 0xa8, 0x6c, 0x07, 0x48, 0xa5, 0x4e, 0xb9, 0xef, 0x27, 0xa7, 0xbb, 0xef,
	0x0c, 0xcb, 0x13, 0x9a, 0x6f, 0x54, 0x7a, 0xdb, 0x9b, 0x8e, 0x3a, 0x36, 0xf5, 0x1f, 0x5b, 0xfc,
	0xc4, 0x30, 0x2b, 0x83, 0xc2, 0x18, 0xa4, 0x28, 0x84, 0xe1, 0x51, 0x1e, 0x6d, 0xe6, 0x95, 0xaf,
	0xd9, 0x0a, 0x62, 0x25, 0x78, 0xec, 0x99, 0x58, 0x09, 0xe7, 0xd1, 0xd8, 0x4a, 0x9e, 0x04, 0x8f,
	0xab, 0xd9, 0x1a, 0xa6, 0x42, 0x12, 0xaa, 0x0b, 0x4f, 0x3d, 0x3b, 0x20, 0xf6, 0x0a, 0x73, 0x4b,
	0x68, 0xe8, 0xa8, 0x5a, 0xc9, 0x27, 0x5e, 0x7a, 0x12, 0x8c, 0xc3, 0x4c, 0x6a, 0xe1, 0xb5, 0xa9,
	0xd7, 0xee, 0xd0, 0xf5, 0xeb, 0xd1, 0xd0, 0x6d, 0xc7, 0x67, 0xa1, 0x5f, 0x40, 0x0f, 0xbe, 0xe4,
	0x2f, 0x23, 0xbe, 0x64, 0x39, 0x64, 0x42, 0xf6, 0x9d, 0x55, 0x54, 0xa2, 0x6e, 0xf8, 0xdc, 0x8b,
	0x63, 0xca, 0x39, 0x4e, 0xa8, 0x9b, 0x5d, 0x5d, 0x77, 0x57, 0x4d, 0x1c, 0x82, 0x63, 0x44, 0xb1,
	0x4f, 0xc8, 0x90, 0x08, 0xeb, 0xf3, 0x5e, 0x5d, 0xa4, 0xe5, 0x59, 0x9e, 0x6c, 0xb2, 0x77, 0x16,
	0xc2, 0xb2, 0xdb, 0xdd, 0x43, 0xaa, 0xc6, 0x36, 0x37, 0x91, 0x25, 0xa4, 0xab, 0xe5, 0x8b, 0x3c,
	0xda, 0x24, 0xd5, 0x80, 0x1c, 0x6f, 0x64, 0x8b, 0xa6, 0xe1, 0xcb, 0x30, 0x69, 0x40, 0xc5, 0x01,
	0xe0, 0xd9, 0xea, 0x91, 0x65, 0x34, 0xca, 0x92, 0x41, 0x2a, 0xa4, 0xad, 0x87, 0xc4, 0x7d, 0xed,
	0xb8, 0x33, 0xda, 0xf3, 0x3d, 0x73, 0x57, 0x17, 0x5f, 0xb0, 0x18, 0xce, 0x76, 0x34, 0x58, 0x37,
	0xff, 0xde, 0xee, 0x0d, 0x26, 0x64, 0x50, 0x5b, 0x1e, 0xfb, 0x6d, 0xd6, 0xf7, 0x6d, 0x86, 0x1f,
	0x0f, 0xca, 0x52, 0x67, 0x6e, 0x55, 0x30, 0x15, 0x7b, 0x58, 0xfd, 0x15, 0xdc, 0x16, 0xd8, 0xfa,
	0xc0, 0x5c, 0xd7, 0xb4, 0x1a, 0x90, 0xbb, 0x2b, 0xa9, 0x56, 0x5a, 0xc2, 0xb6, 0xf7, 0x83, 0x26,
	0xd5, 0x93, 0x38, 0x85, 0x97, 0xf5, 0xf1, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x00, 0xb5, 0xc4, 0x1e,
	0x71, 0x02, 0x00, 0x00,
}
