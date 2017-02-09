// Code generated by protoc-gen-go.
// source: cebcoin.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// account model digitalAssets
type Account struct {
	Addr         string               `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Balance      uint64               `protobuf:"varint,2,opt,name=balance" json:"balance,omitempty"`
	RsaPublicKey string               `protobuf:"bytes,3,opt,name=rsaPublicKey" json:"rsaPublicKey,omitempty"`
	Txouts       map[string]*TX_TXOUT `protobuf:"bytes,4,rep,name=txouts" json:"txouts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Account) GetTxouts() map[string]*TX_TXOUT {
	if m != nil {
		return m.Txouts
	}
	return nil
}

// utxo tx
type TX struct {
	Version   uint64      `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Timestamp int64       `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Txin      []*TX_TXIN  `protobuf:"bytes,3,rep,name=txin" json:"txin,omitempty"`
	Txout     []*TX_TXOUT `protobuf:"bytes,4,rep,name=txout" json:"txout,omitempty"`
	InputData string      `protobuf:"bytes,5,opt,name=InputData" json:"InputData,omitempty"`
	Founder   string      `protobuf:"bytes,6,opt,name=founder" json:"founder,omitempty"`
}

func (m *TX) Reset()                    { *m = TX{} }
func (m *TX) String() string            { return proto.CompactTextString(m) }
func (*TX) ProtoMessage()               {}
func (*TX) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *TX) GetTxin() []*TX_TXIN {
	if m != nil {
		return m.Txin
	}
	return nil
}

func (m *TX) GetTxout() []*TX_TXOUT {
	if m != nil {
		return m.Txout
	}
	return nil
}

// txin not specified who has this txin, because creator can use their txout only, txin must be creator's previous txout
type TX_TXIN struct {
	Idx          uint32 `protobuf:"varint,1,opt,name=idx" json:"idx,omitempty"`
	SourceTxHash string `protobuf:"bytes,2,opt,name=sourceTxHash" json:"sourceTxHash,omitempty"`
	// indicate owner's addr
	Addr string `protobuf:"bytes,3,opt,name=addr" json:"addr,omitempty"`
}

func (m *TX_TXIN) Reset()                    { *m = TX_TXIN{} }
func (m *TX_TXIN) String() string            { return proto.CompactTextString(m) }
func (*TX_TXIN) ProtoMessage()               {}
func (*TX_TXIN) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

type TX_TXOUT struct {
	Value uint64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Addr  string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
	Attr  string `protobuf:"bytes,3,opt,name=attr" json:"attr,omitempty"`
	Sign  string `protobuf:"bytes,4,opt,name=sign" json:"sign,omitempty"`
}

func (m *TX_TXOUT) Reset()                    { *m = TX_TXOUT{} }
func (m *TX_TXOUT) String() string            { return proto.CompactTextString(m) }
func (*TX_TXOUT) ProtoMessage()               {}
func (*TX_TXOUT) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 1} }

func init() {
	proto.RegisterType((*Account)(nil), "protos.Account")
	proto.RegisterType((*TX)(nil), "protos.TX")
	proto.RegisterType((*TX_TXIN)(nil), "protos.TX.TXIN")
	proto.RegisterType((*TX_TXOUT)(nil), "protos.TX.TXOUT")
}

func init() { proto.RegisterFile("cebcoin.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6b, 0xea, 0x40,
	0x14, 0x25, 0x1f, 0x2a, 0x8e, 0x4f, 0x9e, 0x0c, 0x6f, 0x31, 0xf8, 0xde, 0x42, 0x7c, 0x50, 0x5c,
	0x65, 0xa1, 0x9b, 0xd2, 0x5d, 0xa1, 0x85, 0x8a, 0xd0, 0xca, 0x90, 0x82, 0x5d, 0x4e, 0x92, 0x69,
	0x1b, 0xaa, 0x33, 0x92, 0xb9, 0x91, 0xf8, 0x63, 0xfb, 0x17, 0xfa, 0x1b, 0xca, 0xdc, 0x49, 0x8c,
	0x42, 0x57, 0xb9, 0xf7, 0x9c, 0xcc, 0x3d, 0xe7, 0x7e, 0x90, 0x61, 0x2a, 0x93, 0x54, 0xe7, 0x2a,
	0xda, 0x17, 0x1a, 0x34, 0xed, 0xe2, 0xc7, 0x4c, 0x3f, 0x3d, 0xd2, 0xbb, 0x4d, 0x53, 0x5d, 0x2a,
	0xa0, 0x94, 0x84, 0x22, 0xcb, 0x0a, 0xe6, 0x4d, 0xbc, 0x59, 0x9f, 0x63, 0x4c, 0x19, 0xe9, 0x25,
	0x62, 0x2b, 0x54, 0x2a, 0x99, 0x3f, 0xf1, 0x66, 0x21, 0x6f, 0x52, 0x3a, 0x25, 0xbf, 0x0a, 0x23,
	0xd6, 0x65, 0xb2, 0xcd, 0xd3, 0x95, 0x3c, 0xb2, 0x00, 0x5f, 0x5d, 0x60, 0x74, 0x41, 0xba, 0x50,
	0xe9, 0x12, 0x0c, 0x0b, 0x27, 0xc1, 0x6c, 0x30, 0xff, 0xeb, 0xd4, 0x4d, 0x54, 0x4b, 0x46, 0x31,
	0xb2, 0xf7, 0x0a, 0x8a, 0x23, 0xaf, 0x7f, 0x1d, 0xaf, 0xc8, 0xe0, 0x0c, 0xa6, 0x23, 0x12, 0x7c,
	0xc8, 0x63, 0x6d, 0xca, 0x86, 0xf4, 0x8a, 0x74, 0x0e, 0x62, 0x5b, 0x3a, 0x47, 0x83, 0xf9, 0xa8,
	0x29, 0x1a, 0x6f, 0xa2, 0x78, 0xf3, 0xf4, 0x1c, 0x73, 0x47, 0xdf, 0xf8, 0xd7, 0xde, 0xf4, 0xcb,
	0x27, 0x7e, 0xbc, 0xb1, 0x6d, 0x1c, 0x64, 0x61, 0x72, 0xad, 0xb0, 0x50, 0xc8, 0x9b, 0x94, 0xfe,
	0x23, 0x7d, 0xc8, 0x77, 0xd2, 0x80, 0xd8, 0xed, 0xb1, 0x60, 0xc0, 0x5b, 0x80, 0xfe, 0x27, 0x21,
	0x54, 0xb9, 0x62, 0x01, 0xda, 0xff, 0x7d, 0xa1, 0xb4, 0x7c, 0xe4, 0x48, 0x5a, 0x3f, 0x68, 0xbd,
	0x6e, 0xf2, 0x07, 0x3f, 0x48, 0x5b, 0xa9, 0xa5, 0xda, 0x97, 0x70, 0x27, 0x40, 0xb0, 0x0e, 0xf6,
	0xd3, 0x02, 0xd6, 0xe2, 0xab, 0x2e, 0x55, 0x26, 0x0b, 0xd6, 0x45, 0xae, 0x49, 0xc7, 0x6b, 0x12,
	0x5a, 0x35, 0x3b, 0x89, 0x3c, 0xab, 0xb0, 0x81, 0x21, 0xb7, 0xa1, 0xdd, 0x81, 0xd1, 0x65, 0x91,
	0xca, 0xb8, 0x7a, 0x10, 0xe6, 0x1d, 0xfd, 0xf7, 0xf9, 0x05, 0x76, 0xda, 0x6a, 0xd0, 0x6e, 0x75,
	0xfc, 0x42, 0x3a, 0xe8, 0x8c, 0xfe, 0x69, 0x46, 0xe9, 0xa6, 0xe2, 0x92, 0xd3, 0x13, 0xff, 0xec,
	0x10, 0x2c, 0x06, 0xd0, 0x96, 0x01, 0x40, 0xcc, 0xe4, 0x6f, 0x8a, 0x85, 0x0e, 0xb3, 0x71, 0xe2,
	0x0e, 0x6b, 0xf1, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x75, 0x22, 0x26, 0x70, 0x02, 0x00, 0x00,
}
