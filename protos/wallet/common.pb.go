// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet/common.proto

package wallet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Status int32

const (
	Status_INVALID Status = 0
	Status_VALID   Status = 1
	Status_ISSUED  Status = 2
)

var Status_name = map[int32]string{
	0: "INVALID",
	1: "VALID",
	2: "ISSUED",
}
var Status_value = map[string]int32{
	"INVALID": 0,
	"VALID":   1,
	"ISSUED":  2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Metadata defines on chain metadata of decentralized identity
type Metadata struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Metadata) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Metadata) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Metadata)(nil), "wallet.Metadata")
	proto.RegisterEnum("wallet.Status", Status_name, Status_value)
}

func init() { proto.RegisterFile("wallet/common.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4f, 0xcc, 0xc9,
	0x49, 0x2d, 0xd1, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x83, 0x08, 0x2a, 0x19, 0x71, 0x71, 0xf8, 0xa6, 0x96, 0x24, 0xa6, 0x24, 0x96, 0x24, 0x0a,
	0x09, 0x71, 0xb1, 0x64, 0x24, 0x16, 0x67, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9,
	0x20, 0x31, 0x90, 0x9c, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x98, 0xad, 0xa5, 0xc3, 0xc5,
	0x16, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0x2c, 0xc4, 0xcd, 0xc5, 0xee, 0xe9, 0x17, 0xe6, 0xe8, 0xe3,
	0xe9, 0x22, 0xc0, 0x20, 0xc4, 0xc9, 0xc5, 0x0a, 0x61, 0x32, 0x0a, 0x71, 0x71, 0xb1, 0x79, 0x06,
	0x07, 0x87, 0xba, 0xba, 0x08, 0x30, 0x39, 0x19, 0x47, 0x19, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0x16, 0x55, 0x24, 0xe6, 0x25, 0x67, 0x24, 0x66, 0xe6, 0xe9,
	0x17, 0xa7, 0x64, 0xeb, 0xa6, 0xe7, 0xeb, 0x42, 0x9c, 0xa5, 0x0f, 0x76, 0x56, 0xb1, 0x3e, 0xc4,
	0x59, 0x49, 0x6c, 0x60, 0xae, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x08, 0x3d, 0x3d, 0xbc,
	0x00, 0x00, 0x00,
}
