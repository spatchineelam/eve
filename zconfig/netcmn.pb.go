// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netcmn.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NetworkType int32

const (
	NetworkType_NETWORKTYPENOOP NetworkType = 0
	NetworkType_V4              NetworkType = 4
	NetworkType_V6              NetworkType = 6
	NetworkType_L2              NetworkType = 2
	NetworkType_LISP            NetworkType = 10
)

var NetworkType_name = map[int32]string{
	0:  "NETWORKTYPENOOP",
	4:  "V4",
	6:  "V6",
	2:  "L2",
	10: "LISP",
}
var NetworkType_value = map[string]int32{
	"NETWORKTYPENOOP": 0,
	"V4":              4,
	"V6":              6,
	"L2":              2,
	"LISP":            10,
}

func (x NetworkType) String() string {
	return proto.EnumName(NetworkType_name, int32(x))
}
func (NetworkType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

// import "fw.proto";
type IpRange struct {
	Start string `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End   string `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
}

func (m *IpRange) Reset()                    { *m = IpRange{} }
func (m *IpRange) String() string            { return proto.CompactTextString(m) }
func (*IpRange) ProtoMessage()               {}
func (*IpRange) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *IpRange) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *IpRange) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

type Ipv4Spec struct {
	Dhcp    bool     `protobuf:"varint,2,opt,name=dhcp" json:"dhcp,omitempty"`
	Subnet  string   `protobuf:"bytes,3,opt,name=subnet" json:"subnet,omitempty"`
	Mask    string   `protobuf:"bytes,4,opt,name=mask" json:"mask,omitempty"`
	Gateway string   `protobuf:"bytes,5,opt,name=gateway" json:"gateway,omitempty"`
	Domain  string   `protobuf:"bytes,6,opt,name=domain" json:"domain,omitempty"`
	Ntp     string   `protobuf:"bytes,7,opt,name=ntp" json:"ntp,omitempty"`
	Dns     []string `protobuf:"bytes,8,rep,name=dns" json:"dns,omitempty"`
	// for IPAM management when dhcp is turned on.
	// If none provided, system will default pool.
	IpRange string `protobuf:"bytes,9,opt,name=ipRange" json:"ipRange,omitempty"`
}

func (m *Ipv4Spec) Reset()                    { *m = Ipv4Spec{} }
func (m *Ipv4Spec) String() string            { return proto.CompactTextString(m) }
func (*Ipv4Spec) ProtoMessage()               {}
func (*Ipv4Spec) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *Ipv4Spec) GetDhcp() bool {
	if m != nil {
		return m.Dhcp
	}
	return false
}

func (m *Ipv4Spec) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

func (m *Ipv4Spec) GetMask() string {
	if m != nil {
		return m.Mask
	}
	return ""
}

func (m *Ipv4Spec) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *Ipv4Spec) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Ipv4Spec) GetNtp() string {
	if m != nil {
		return m.Ntp
	}
	return ""
}

func (m *Ipv4Spec) GetDns() []string {
	if m != nil {
		return m.Dns
	}
	return nil
}

func (m *Ipv4Spec) GetIpRange() string {
	if m != nil {
		return m.IpRange
	}
	return ""
}

type Ipv6Spec struct {
	Dhcp    bool     `protobuf:"varint,2,opt,name=dhcp" json:"dhcp,omitempty"`
	Subnet  string   `protobuf:"bytes,3,opt,name=subnet" json:"subnet,omitempty"`
	Mask    string   `protobuf:"bytes,4,opt,name=mask" json:"mask,omitempty"`
	Gateway string   `protobuf:"bytes,5,opt,name=gateway" json:"gateway,omitempty"`
	Domain  string   `protobuf:"bytes,6,opt,name=domain" json:"domain,omitempty"`
	Ntp     string   `protobuf:"bytes,7,opt,name=ntp" json:"ntp,omitempty"`
	Dns     []string `protobuf:"bytes,8,rep,name=dns" json:"dns,omitempty"`
	// for IPAM management when dhcp is turned on.
	// If none provided, system will default pool.
	IpRange string `protobuf:"bytes,9,opt,name=ipRange" json:"ipRange,omitempty"`
}

func (m *Ipv6Spec) Reset()                    { *m = Ipv6Spec{} }
func (m *Ipv6Spec) String() string            { return proto.CompactTextString(m) }
func (*Ipv6Spec) ProtoMessage()               {}
func (*Ipv6Spec) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *Ipv6Spec) GetDhcp() bool {
	if m != nil {
		return m.Dhcp
	}
	return false
}

func (m *Ipv6Spec) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

func (m *Ipv6Spec) GetMask() string {
	if m != nil {
		return m.Mask
	}
	return ""
}

func (m *Ipv6Spec) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *Ipv6Spec) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Ipv6Spec) GetNtp() string {
	if m != nil {
		return m.Ntp
	}
	return ""
}

func (m *Ipv6Spec) GetDns() []string {
	if m != nil {
		return m.Dns
	}
	return nil
}

func (m *Ipv6Spec) GetIpRange() string {
	if m != nil {
		return m.IpRange
	}
	return ""
}

type NameToEid struct {
	Hostname string   `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Eids     []string `protobuf:"bytes,2,rep,name=eids" json:"eids,omitempty"`
}

func (m *NameToEid) Reset()                    { *m = NameToEid{} }
func (m *NameToEid) String() string            { return proto.CompactTextString(m) }
func (*NameToEid) ProtoMessage()               {}
func (*NameToEid) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *NameToEid) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *NameToEid) GetEids() []string {
	if m != nil {
		return m.Eids
	}
	return nil
}

type EIDAllocation struct {
	Allocate            bool   `protobuf:"varint,1,opt,name=allocate" json:"allocate,omitempty"`
	Exportprivate       bool   `protobuf:"varint,2,opt,name=exportprivate" json:"exportprivate,omitempty"`
	Allocationprefix    []byte `protobuf:"bytes,3,opt,name=allocationprefix,proto3" json:"allocationprefix,omitempty"`
	Allocationprefixlen uint32 `protobuf:"varint,4,opt,name=allocationprefixlen" json:"allocationprefixlen,omitempty"`
}

func (m *EIDAllocation) Reset()                    { *m = EIDAllocation{} }
func (m *EIDAllocation) String() string            { return proto.CompactTextString(m) }
func (*EIDAllocation) ProtoMessage()               {}
func (*EIDAllocation) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *EIDAllocation) GetAllocate() bool {
	if m != nil {
		return m.Allocate
	}
	return false
}

func (m *EIDAllocation) GetExportprivate() bool {
	if m != nil {
		return m.Exportprivate
	}
	return false
}

func (m *EIDAllocation) GetAllocationprefix() []byte {
	if m != nil {
		return m.Allocationprefix
	}
	return nil
}

func (m *EIDAllocation) GetAllocationprefixlen() uint32 {
	if m != nil {
		return m.Allocationprefixlen
	}
	return 0
}

type Lispspec struct {
	Iid      uint32         `protobuf:"varint,1,opt,name=iid" json:"iid,omitempty"`
	Eidalloc *EIDAllocation `protobuf:"bytes,2,opt,name=eidalloc" json:"eidalloc,omitempty"`
	Nmtoeid  []*NameToEid   `protobuf:"bytes,3,rep,name=nmtoeid" json:"nmtoeid,omitempty"`
}

func (m *Lispspec) Reset()                    { *m = Lispspec{} }
func (m *Lispspec) String() string            { return proto.CompactTextString(m) }
func (*Lispspec) ProtoMessage()               {}
func (*Lispspec) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *Lispspec) GetIid() uint32 {
	if m != nil {
		return m.Iid
	}
	return 0
}

func (m *Lispspec) GetEidalloc() *EIDAllocation {
	if m != nil {
		return m.Eidalloc
	}
	return nil
}

func (m *Lispspec) GetNmtoeid() []*NameToEid {
	if m != nil {
		return m.Nmtoeid
	}
	return nil
}

type L2Spec struct {
}

func (m *L2Spec) Reset()                    { *m = L2Spec{} }
func (m *L2Spec) String() string            { return proto.CompactTextString(m) }
func (*L2Spec) ProtoMessage()               {}
func (*L2Spec) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func init() {
	proto.RegisterType((*IpRange)(nil), "ipRange")
	proto.RegisterType((*Ipv4Spec)(nil), "ipv4spec")
	proto.RegisterType((*Ipv6Spec)(nil), "ipv6spec")
	proto.RegisterType((*NameToEid)(nil), "NameToEid")
	proto.RegisterType((*EIDAllocation)(nil), "EIDAllocation")
	proto.RegisterType((*Lispspec)(nil), "lispspec")
	proto.RegisterType((*L2Spec)(nil), "l2spec")
	proto.RegisterEnum("NetworkType", NetworkType_name, NetworkType_value)
}

func init() { proto.RegisterFile("netcmn.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x53, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x25, 0xd9, 0x74, 0xb3, 0x9d, 0x34, 0xb0, 0x72, 0x11, 0xb2, 0x90, 0x10, 0x51, 0xd4, 0x43,
	0x95, 0xc3, 0x06, 0x42, 0xd5, 0x0b, 0x17, 0x40, 0xcd, 0xa1, 0xa2, 0x4a, 0xa2, 0x25, 0x02, 0xc1,
	0xcd, 0x59, 0x4f, 0x13, 0xab, 0xbb, 0xb6, 0xb5, 0x76, 0xd2, 0x8f, 0x3f, 0xc5, 0x9d, 0x5f, 0x87,
	0x6c, 0x27, 0x91, 0x0a, 0xfc, 0x01, 0x4e, 0x7e, 0xef, 0xcd, 0xf8, 0xed, 0xf3, 0x8c, 0x16, 0x8e,
	0x24, 0xda, 0xa2, 0x92, 0x99, 0xae, 0x95, 0x55, 0xfd, 0xb7, 0xd0, 0x16, 0x3a, 0x67, 0x72, 0x89,
	0xe4, 0x39, 0x1c, 0x18, 0xcb, 0x6a, 0x4b, 0x1b, 0xbd, 0xc6, 0xe9, 0x61, 0x1e, 0x08, 0x49, 0x21,
	0x42, 0xc9, 0x69, 0xd3, 0x6b, 0x0e, 0xf6, 0x7f, 0x35, 0x20, 0x11, 0x7a, 0x73, 0x66, 0x34, 0x16,
	0x84, 0x40, 0x8b, 0xaf, 0x0a, 0xed, 0xeb, 0x49, 0xee, 0x31, 0x79, 0x01, 0xb1, 0x59, 0x2f, 0x24,
	0x5a, 0x1a, 0xf9, 0x5b, 0x5b, 0xe6, 0x7a, 0x2b, 0x66, 0x6e, 0x68, 0xcb, 0xab, 0x1e, 0x13, 0x0a,
	0xed, 0x25, 0xb3, 0x78, 0xcb, 0xee, 0xe9, 0x81, 0x97, 0x77, 0xd4, 0xb9, 0x70, 0x55, 0x31, 0x21,
	0x69, 0x1c, 0x5c, 0x02, 0x73, 0x81, 0xa4, 0xd5, 0xb4, 0x1d, 0x02, 0x49, 0xab, 0x9d, 0xc2, 0xa5,
	0xa1, 0x49, 0x2f, 0x72, 0x0a, 0x97, 0xc6, 0xb9, 0x6e, 0x5f, 0x45, 0x0f, 0x83, 0xeb, 0x96, 0xee,
	0xc2, 0x9f, 0xff, 0x97, 0xe1, 0xdf, 0xc3, 0xe1, 0x84, 0x55, 0x38, 0x57, 0x63, 0xc1, 0xc9, 0x4b,
	0x48, 0x56, 0xca, 0x58, 0xc9, 0x2a, 0xdc, 0x6e, 0x6c, 0xcf, 0x5d, 0x58, 0x14, 0xdc, 0xd0, 0xa6,
	0x77, 0xf5, 0xb8, 0xff, 0xb3, 0x01, 0xdd, 0xf1, 0xe5, 0xc5, 0xc7, 0xb2, 0x54, 0x05, 0xb3, 0x42,
	0x49, 0xe7, 0xc0, 0x02, 0x0b, 0x0e, 0x49, 0xbe, 0xe7, 0xe4, 0x04, 0xba, 0x78, 0xa7, 0x55, 0x6d,
	0x75, 0x2d, 0x36, 0xae, 0x21, 0xcc, 0xe8, 0xb1, 0x48, 0x06, 0x90, 0xb2, 0xbd, 0x9f, 0xae, 0xf1,
	0x5a, 0xdc, 0xf9, 0xb1, 0x1d, 0xe5, 0x7f, 0xe9, 0xe4, 0x0d, 0x1c, 0xff, 0xa9, 0x95, 0x28, 0xfd,
	0x3c, 0xbb, 0xf9, 0xbf, 0x4a, 0x7d, 0x09, 0x49, 0x29, 0x8c, 0xf6, 0xab, 0x4a, 0x21, 0x12, 0x82,
	0xfb, 0x98, 0xdd, 0xdc, 0x41, 0x32, 0x80, 0x04, 0x05, 0xf7, 0xf7, 0x7c, 0xb8, 0xce, 0xe8, 0x69,
	0xf6, 0xe8, 0x7d, 0xf9, 0xbe, 0x4e, 0x4e, 0xa0, 0x2d, 0x2b, 0xab, 0x50, 0x70, 0x1a, 0xf5, 0xa2,
	0xd3, 0xce, 0x08, 0xb2, 0xfd, 0x20, 0xf3, 0x5d, 0xa9, 0x9f, 0x40, 0x5c, 0x8e, 0xdc, 0xd7, 0x06,
	0x17, 0xd0, 0x99, 0xa0, 0xbd, 0x55, 0xf5, 0xcd, 0xfc, 0x5e, 0x23, 0x39, 0x86, 0x67, 0x93, 0xf1,
	0xfc, 0xdb, 0x34, 0xff, 0x3c, 0xff, 0x3e, 0x1b, 0x4f, 0xa6, 0xd3, 0x59, 0xfa, 0x84, 0xc4, 0xd0,
	0xfc, 0x7a, 0x96, 0xb6, 0xfc, 0x79, 0x9e, 0xc6, 0xee, 0xbc, 0x1a, 0xa5, 0x4d, 0x92, 0x40, 0xeb,
	0xea, 0xf2, 0xcb, 0x2c, 0x85, 0x4f, 0x1f, 0xe0, 0x75, 0xa1, 0xaa, 0xec, 0x01, 0x39, 0x72, 0x96,
	0x15, 0xa5, 0x5a, 0xf3, 0x6c, 0x6d, 0xb0, 0xde, 0x88, 0x02, 0xc3, 0xef, 0xf7, 0xe3, 0xd5, 0x52,
	0xd8, 0xd5, 0x7a, 0x91, 0x15, 0xaa, 0x1a, 0x86, 0xbe, 0x21, 0xd3, 0x62, 0xf8, 0x50, 0x28, 0x79,
	0x2d, 0x96, 0x8b, 0xd8, 0x77, 0xbd, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x85, 0x9c, 0xb1, 0x2e,
	0xb4, 0x03, 0x00, 0x00,
}
