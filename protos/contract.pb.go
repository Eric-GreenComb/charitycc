// Code generated by protoc-gen-go.
// source: contract.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Contract model
type Contract struct {
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
	Remark      string        `protobuf:"bytes,11,opt,name=remark" json:"remark,omitempty"`
	Status      int64         `protobuf:"varint,12,opt,name=status" json:"status,omitempty"`
	AttachFiles []*AttachFile `protobuf:"bytes,13,rep,name=attachFiles" json:"attachFiles,omitempty"`
}

func (m *Contract) Reset()                    { *m = Contract{} }
func (m *Contract) String() string            { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()               {}
func (*Contract) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Contract) GetAttachFiles() []*AttachFile {
	if m != nil {
		return m.AttachFiles
	}
	return nil
}

type AttachFile struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Hash string `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
}

func (m *AttachFile) Reset()                    { *m = AttachFile{} }
func (m *AttachFile) String() string            { return proto.CompactTextString(m) }
func (*AttachFile) ProtoMessage()               {}
func (*AttachFile) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func init() {
	proto.RegisterType((*Contract)(nil), "protos.Contract")
	proto.RegisterType((*AttachFile)(nil), "protos.AttachFile")
}

func init() { proto.RegisterFile("contract.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x91, 0xb1, 0x6e, 0x83, 0x30,
	0x14, 0x45, 0x05, 0x49, 0x93, 0xf0, 0x68, 0x33, 0x78, 0xa8, 0xde, 0xd0, 0x01, 0x65, 0x62, 0xca,
	0xd0, 0xe6, 0x07, 0xa0, 0x52, 0x3f, 0x00, 0xf5, 0x07, 0x5e, 0xb0, 0x25, 0x2c, 0x82, 0x41, 0xf6,
	0xcb, 0xd0, 0x7f, 0xee, 0x47, 0x54, 0xb6, 0x29, 0x61, 0xf2, 0xbb, 0xe7, 0x5e, 0xdd, 0xe1, 0x1a,
	0x8e, 0xed, 0x68, 0xd8, 0x52, 0xcb, 0xe7, 0xc9, 0x8e, 0x3c, 0x8a, 0x5d, 0x78, 0xdc, 0xe9, 0x37,
	0x85, 0xc3, 0xe7, 0x6c, 0x09, 0x01, 0x5b, 0x92, 0xd2, 0x62, 0x52, 0x24, 0x65, 0xd6, 0x84, 0x5b,
	0x1c, 0x21, 0xd5, 0x12, 0xd3, 0x40, 0x52, 0x2d, 0x7d, 0xc6, 0xd0, 0xa0, 0x70, 0x13, 0x33, 0xfe,
	0x16, 0xaf, 0xb0, 0x93, 0x8a, 0x49, 0xdf, 0x70, 0x1b, 0xe8, 0xac, 0xc4, 0x1b, 0x64, 0x8e, 0xc9,
	0xf2, 0xb7, 0x1e, 0x14, 0x3e, 0x05, 0xeb, 0x01, 0x04, 0xc2, 0x5e, 0x19, 0x19, 0xbc, 0x5d, 0xf0,
	0xfe, 0xa5, 0xef, 0x9b, 0xc8, 0xf2, 0x4f, 0x85, 0xfb, 0xd8, 0x17, 0xd5, 0xc2, 0x6b, 0x3c, 0xac,
	0x78, 0x2d, 0x0a, 0xc8, 0xa5, 0x9a, 0x46, 0xa7, 0xb9, 0x26, 0xd3, 0x63, 0x16, 0xcc, 0x35, 0xf2,
	0x89, 0x2b, 0x99, 0xbe, 0x6a, 0xdb, 0xf1, 0x6e, 0x18, 0x21, 0x26, 0x56, 0xc8, 0x77, 0x5b, 0x35,
	0x90, 0xed, 0x31, 0x8f, 0xdd, 0x51, 0x79, 0xee, 0x98, 0xf8, 0xee, 0xf0, 0xb9, 0x48, 0xca, 0x4d,
	0x33, 0x2b, 0x71, 0x81, 0x9c, 0x98, 0xa9, 0xed, 0xbe, 0xf4, 0x4d, 0x39, 0x7c, 0x29, 0x36, 0x65,
	0xfe, 0x2e, 0xe2, 0xba, 0xee, 0x5c, 0x2d, 0x56, 0xb3, 0x8e, 0x9d, 0x2e, 0x00, 0x0f, 0x6b, 0xd9,
	0x32, 0x59, 0x6d, 0x29, 0x60, 0xdb, 0x91, 0xeb, 0xe6, 0xc5, 0xc3, 0x7d, 0x8d, 0x9f, 0xf5, 0xf1,
	0x17, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xf5, 0x63, 0xe5, 0xc5, 0x01, 0x00, 0x00,
}
