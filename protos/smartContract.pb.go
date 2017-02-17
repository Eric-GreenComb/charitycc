// Code generated by protoc-gen-go.
// source: smartContract.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// SmartContract model
type SmartContract struct {
	Addr                 string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Id                   string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Name                 string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Detail               string `protobuf:"bytes,4,opt,name=detail" json:"detail,omitempty"`
	Goal                 uint64 `protobuf:"varint,5,opt,name=goal" json:"goal,omitempty"`
	PartyA               string `protobuf:"bytes,6,opt,name=partyA" json:"partyA,omitempty"`
	PartyB               string `protobuf:"bytes,7,opt,name=partyB" json:"partyB,omitempty"`
	Type                 uint64 `protobuf:"varint,8,opt,name=type" json:"type,omitempty"`
	FundAddr             string `protobuf:"bytes,9,opt,name=fundAddr" json:"fundAddr,omitempty"`
	FundName             string `protobuf:"bytes,10,opt,name=fundName" json:"fundName,omitempty"`
	FundManangerFee      uint64 `protobuf:"varint,11,opt,name=fundManangerFee" json:"fundManangerFee,omitempty"`
	ChannelAddr          string `protobuf:"bytes,12,opt,name=channelAddr" json:"channelAddr,omitempty"`
	ChannelName          string `protobuf:"bytes,13,opt,name=channelName" json:"channelName,omitempty"`
	ChannelFee           uint64 `protobuf:"varint,14,opt,name=channelFee" json:"channelFee,omitempty"`
	CreateTimestamp      int64  `protobuf:"varint,15,opt,name=createTimestamp" json:"createTimestamp,omitempty"`
	UtilTimestamp        int64  `protobuf:"varint,16,opt,name=utilTimestamp" json:"utilTimestamp,omitempty"`
	TerminationTimestamp int64  `protobuf:"varint,17,opt,name=terminationTimestamp" json:"terminationTimestamp,omitempty"`
	Foundation           string `protobuf:"bytes,18,opt,name=foundation" json:"foundation,omitempty"`
	Attachhash           string `protobuf:"bytes,19,opt,name=attachhash" json:"attachhash,omitempty"`
	Status               uint64 `protobuf:"varint,20,opt,name=status" json:"status,omitempty"`
	Remark               string `protobuf:"bytes,21,opt,name=remark" json:"remark,omitempty"`
}

func (m *SmartContract) Reset()                    { *m = SmartContract{} }
func (m *SmartContract) String() string            { return proto.CompactTextString(m) }
func (*SmartContract) ProtoMessage()               {}
func (*SmartContract) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type SmartContractExt struct {
	Addr          string         `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Balance       uint64         `protobuf:"varint,2,opt,name=balance" json:"balance,omitempty"`
	Total         uint64         `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	ValidTotal    uint64         `protobuf:"varint,4,opt,name=validTotal" json:"validTotal,omitempty"`
	DonateNumber  uint64         `protobuf:"varint,5,opt,name=donateNumber" json:"donateNumber,omitempty"`
	SmartContract *SmartContract `protobuf:"bytes,6,opt,name=smartContract" json:"smartContract,omitempty"`
}

func (m *SmartContractExt) Reset()                    { *m = SmartContractExt{} }
func (m *SmartContractExt) String() string            { return proto.CompactTextString(m) }
func (*SmartContractExt) ProtoMessage()               {}
func (*SmartContractExt) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *SmartContractExt) GetSmartContract() *SmartContract {
	if m != nil {
		return m.SmartContract
	}
	return nil
}

type SmartContractTrack struct {
	Addr  string                  `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Trans []*SmartContractHistory `protobuf:"bytes,2,rep,name=trans" json:"trans,omitempty"`
}

func (m *SmartContractTrack) Reset()                    { *m = SmartContractTrack{} }
func (m *SmartContractTrack) String() string            { return proto.CompactTextString(m) }
func (*SmartContractTrack) ProtoMessage()               {}
func (*SmartContractTrack) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *SmartContractTrack) GetTrans() []*SmartContractHistory {
	if m != nil {
		return m.Trans
	}
	return nil
}

type SmartContractHistory struct {
	BargainName string `protobuf:"bytes,1,opt,name=bargainName" json:"bargainName,omitempty"`
	BargainAddr string `protobuf:"bytes,2,opt,name=bargainAddr" json:"bargainAddr,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Amount      uint64 `protobuf:"varint,4,opt,name=amount" json:"amount,omitempty"`
	Timestamp   int64  `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *SmartContractHistory) Reset()                    { *m = SmartContractHistory{} }
func (m *SmartContractHistory) String() string            { return proto.CompactTextString(m) }
func (*SmartContractHistory) ProtoMessage()               {}
func (*SmartContractHistory) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func init() {
	proto.RegisterType((*SmartContract)(nil), "protos.SmartContract")
	proto.RegisterType((*SmartContractExt)(nil), "protos.SmartContractExt")
	proto.RegisterType((*SmartContractTrack)(nil), "protos.SmartContractTrack")
	proto.RegisterType((*SmartContractHistory)(nil), "protos.SmartContractHistory")
}

func init() { proto.RegisterFile("smartContract.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x54, 0xcf, 0x6e, 0xda, 0x4c,
	0x10, 0x97, 0xc1, 0x90, 0x30, 0x84, 0x24, 0xdf, 0x86, 0x44, 0xab, 0x4f, 0x51, 0x85, 0x50, 0x0f,
	0x9c, 0x72, 0xa0, 0xc7, 0x9e, 0x92, 0xaa, 0x55, 0x2f, 0xcd, 0xc1, 0xe5, 0xd8, 0xcb, 0x60, 0x6f,
	0x60, 0x15, 0x7b, 0x17, 0xad, 0xc7, 0x55, 0x79, 0x9f, 0xbe, 0x53, 0x9f, 0xa3, 0x6f, 0x50, 0xed,
	0xac, 0xc1, 0x76, 0xc4, 0x89, 0xf9, 0xfd, 0xd1, 0xcc, 0xd8, 0xfe, 0x0d, 0x70, 0x53, 0x16, 0xe8,
	0xe8, 0x93, 0x35, 0xe4, 0x30, 0xa5, 0x87, 0x9d, 0xb3, 0x64, 0xc5, 0x90, 0x7f, 0xca, 0xf9, 0xdf,
	0x18, 0x26, 0xdf, 0xdb, 0xba, 0x10, 0x10, 0x63, 0x96, 0x39, 0x19, 0xcd, 0xa2, 0xc5, 0x28, 0xe1,
	0x5a, 0x5c, 0x42, 0x4f, 0x67, 0xb2, 0xc7, 0x4c, 0x4f, 0x67, 0xde, 0x63, 0xb0, 0x50, 0xb2, 0x1f,
	0x3c, 0xbe, 0x16, 0x77, 0x30, 0xcc, 0x14, 0xa1, 0xce, 0x65, 0xcc, 0x6c, 0x8d, 0xbc, 0x77, 0x63,
	0x31, 0x97, 0x83, 0x59, 0xb4, 0x88, 0x13, 0xae, 0xbd, 0x77, 0x87, 0x8e, 0xf6, 0x8f, 0x72, 0x18,
	0xbc, 0x01, 0x1d, 0xf9, 0x27, 0x79, 0xd6, 0xe2, 0x9f, 0x7c, 0x0f, 0xda, 0xef, 0x94, 0x3c, 0x0f,
	0x3d, 0x7c, 0x2d, 0xfe, 0x87, 0xf3, 0x97, 0xca, 0x64, 0x8f, 0x7e, 0xd7, 0x11, 0xbb, 0x8f, 0xf8,
	0xa0, 0x3d, 0xfb, 0x1d, 0xa1, 0xd1, 0x3c, 0x16, 0x0b, 0xb8, 0xf2, 0xf5, 0x37, 0x34, 0x68, 0x36,
	0xca, 0x7d, 0x51, 0x4a, 0x8e, 0xb9, 0xed, 0x5b, 0x5a, 0xcc, 0x60, 0x9c, 0x6e, 0xd1, 0x18, 0x95,
	0xf3, 0x90, 0x0b, 0x6e, 0xd4, 0xa6, 0x5a, 0x0e, 0x1e, 0x35, 0xe9, 0x38, 0x78, 0xda, 0x3b, 0x80,
	0x1a, 0xfa, 0x41, 0x97, 0x3c, 0xa8, 0xc5, 0xf8, 0x6d, 0x52, 0xa7, 0x90, 0xd4, 0x4a, 0x17, 0xaa,
	0x24, 0x2c, 0x76, 0xf2, 0x6a, 0x16, 0x2d, 0xfa, 0xc9, 0x5b, 0x5a, 0xbc, 0x87, 0x49, 0x45, 0x3a,
	0x6f, 0x7c, 0xd7, 0xec, 0xeb, 0x92, 0x62, 0x09, 0x53, 0x52, 0xae, 0xd0, 0x06, 0x49, 0x5b, 0xd3,
	0x98, 0xff, 0x63, 0xf3, 0x49, 0xcd, 0xef, 0xf8, 0x62, 0x2b, 0x93, 0x31, 0x2d, 0x05, 0x3f, 0x44,
	0x8b, 0xf1, 0x3a, 0x12, 0x61, 0xba, 0xdd, 0x62, 0xb9, 0x95, 0x37, 0x41, 0x6f, 0x18, 0xff, 0xd5,
	0x4a, 0x42, 0xaa, 0x4a, 0x39, 0xe5, 0xe7, 0xab, 0x91, 0xe7, 0x9d, 0x2a, 0xd0, 0xbd, 0xca, 0xdb,
	0xf0, 0x35, 0x03, 0x9a, 0xff, 0x89, 0xe0, 0xba, 0x93, 0xb9, 0xcf, 0xbf, 0x4e, 0xc7, 0x4e, 0xc2,
	0xd9, 0x1a, 0x73, 0x34, 0xa9, 0xe2, 0xec, 0xc5, 0xc9, 0x01, 0x8a, 0x29, 0x0c, 0xc8, 0x12, 0xe6,
	0x9c, 0xc0, 0x38, 0x09, 0xc0, 0x2f, 0xfa, 0x13, 0x73, 0x9d, 0xad, 0x58, 0x8a, 0xc3, 0xcb, 0x6e,
	0x18, 0x31, 0x87, 0x8b, 0xcc, 0x1a, 0x24, 0xf5, 0x5c, 0x15, 0x6b, 0xe5, 0xea, 0x48, 0x76, 0x38,
	0xf1, 0x11, 0x26, 0x9d, 0x7b, 0xe1, 0x84, 0x8e, 0x97, 0xb7, 0xe1, 0x6e, 0xca, 0x87, 0xce, 0xe2,
	0x49, 0xd7, 0x3b, 0xff, 0x01, 0xa2, 0xa3, 0xaf, 0x1c, 0xa6, 0xaf, 0x27, 0x1f, 0x6d, 0x09, 0x03,
	0x72, 0x68, 0x4a, 0xd9, 0x9b, 0xf5, 0x17, 0xe3, 0xe5, 0xfd, 0xc9, 0xf6, 0x5f, 0x75, 0x49, 0xd6,
	0xed, 0x93, 0x60, 0x9d, 0xff, 0x8e, 0x60, 0x7a, 0x4a, 0xf7, 0x31, 0x5c, 0xa3, 0xdb, 0xa0, 0x36,
	0x1c, 0xc3, 0x30, 0xa7, 0x4d, 0xb5, 0x1c, 0x1c, 0xe5, 0x5e, 0xc7, 0xc1, 0x51, 0x3e, 0x9c, 0x58,
	0x7d, 0xd2, 0x7c, 0x62, 0x77, 0x30, 0xc4, 0xc2, 0x56, 0x86, 0xea, 0x77, 0x59, 0x23, 0x71, 0x0f,
	0x23, 0x3a, 0x26, 0x6b, 0xc0, 0xc9, 0x6a, 0x88, 0x75, 0xf8, 0x6b, 0xf9, 0xf0, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x8d, 0x6c, 0xf5, 0xba, 0x78, 0x04, 0x00, 0x00,
}
