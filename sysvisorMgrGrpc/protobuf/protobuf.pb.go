// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegisterReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{0}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RegisterResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResp) Reset()         { *m = RegisterResp{} }
func (m *RegisterResp) String() string { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()    {}
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{1}
}

func (m *RegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResp.Unmarshal(m, b)
}
func (m *RegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResp.Marshal(b, m, deterministic)
}
func (m *RegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResp.Merge(m, src)
}
func (m *RegisterResp) XXX_Size() int {
	return xxx_messageInfo_RegisterResp.Size(m)
}
func (m *RegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResp proto.InternalMessageInfo

type UnregisterReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterReq) Reset()         { *m = UnregisterReq{} }
func (m *UnregisterReq) String() string { return proto.CompactTextString(m) }
func (*UnregisterReq) ProtoMessage()    {}
func (*UnregisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{2}
}

func (m *UnregisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterReq.Unmarshal(m, b)
}
func (m *UnregisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterReq.Marshal(b, m, deterministic)
}
func (m *UnregisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterReq.Merge(m, src)
}
func (m *UnregisterReq) XXX_Size() int {
	return xxx_messageInfo_UnregisterReq.Size(m)
}
func (m *UnregisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterReq proto.InternalMessageInfo

func (m *UnregisterReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UnregisterResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterResp) Reset()         { *m = UnregisterResp{} }
func (m *UnregisterResp) String() string { return proto.CompactTextString(m) }
func (*UnregisterResp) ProtoMessage()    {}
func (*UnregisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{3}
}

func (m *UnregisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterResp.Unmarshal(m, b)
}
func (m *UnregisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterResp.Marshal(b, m, deterministic)
}
func (m *UnregisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterResp.Merge(m, src)
}
func (m *UnregisterResp) XXX_Size() int {
	return xxx_messageInfo_UnregisterResp.Size(m)
}
func (m *UnregisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterResp proto.InternalMessageInfo

type SubidAllocReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Size                 uint64   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubidAllocReq) Reset()         { *m = SubidAllocReq{} }
func (m *SubidAllocReq) String() string { return proto.CompactTextString(m) }
func (*SubidAllocReq) ProtoMessage()    {}
func (*SubidAllocReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{4}
}

func (m *SubidAllocReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubidAllocReq.Unmarshal(m, b)
}
func (m *SubidAllocReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubidAllocReq.Marshal(b, m, deterministic)
}
func (m *SubidAllocReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubidAllocReq.Merge(m, src)
}
func (m *SubidAllocReq) XXX_Size() int {
	return xxx_messageInfo_SubidAllocReq.Size(m)
}
func (m *SubidAllocReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SubidAllocReq.DiscardUnknown(m)
}

var xxx_messageInfo_SubidAllocReq proto.InternalMessageInfo

func (m *SubidAllocReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SubidAllocReq) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type SubidAllocResp struct {
	Uid                  uint32   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid                  uint32   `protobuf:"varint,2,opt,name=gid,proto3" json:"gid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubidAllocResp) Reset()         { *m = SubidAllocResp{} }
func (m *SubidAllocResp) String() string { return proto.CompactTextString(m) }
func (*SubidAllocResp) ProtoMessage()    {}
func (*SubidAllocResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{5}
}

func (m *SubidAllocResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubidAllocResp.Unmarshal(m, b)
}
func (m *SubidAllocResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubidAllocResp.Marshal(b, m, deterministic)
}
func (m *SubidAllocResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubidAllocResp.Merge(m, src)
}
func (m *SubidAllocResp) XXX_Size() int {
	return xxx_messageInfo_SubidAllocResp.Size(m)
}
func (m *SubidAllocResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SubidAllocResp.DiscardUnknown(m)
}

var xxx_messageInfo_SubidAllocResp proto.InternalMessageInfo

func (m *SubidAllocResp) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *SubidAllocResp) GetGid() uint32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

type Mount struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Dest                 string   `protobuf:"bytes,2,opt,name=dest,proto3" json:"dest,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Opt                  []string `protobuf:"bytes,4,rep,name=opt,proto3" json:"opt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mount) Reset()         { *m = Mount{} }
func (m *Mount) String() string { return proto.CompactTextString(m) }
func (*Mount) ProtoMessage()    {}
func (*Mount) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{6}
}

func (m *Mount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mount.Unmarshal(m, b)
}
func (m *Mount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mount.Marshal(b, m, deterministic)
}
func (m *Mount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mount.Merge(m, src)
}
func (m *Mount) XXX_Size() int {
	return xxx_messageInfo_Mount.Size(m)
}
func (m *Mount) XXX_DiscardUnknown() {
	xxx_messageInfo_Mount.DiscardUnknown(m)
}

var xxx_messageInfo_Mount proto.InternalMessageInfo

func (m *Mount) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Mount) GetDest() string {
	if m != nil {
		return m.Dest
	}
	return ""
}

func (m *Mount) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Mount) GetOpt() []string {
	if m != nil {
		return m.Opt
	}
	return nil
}

type SupMountsReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Rootfs               string   `protobuf:"bytes,2,opt,name=rootfs,proto3" json:"rootfs,omitempty"`
	Uid                  uint32   `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid                  uint32   `protobuf:"varint,4,opt,name=gid,proto3" json:"gid,omitempty"`
	ShiftUids            bool     `protobuf:"varint,5,opt,name=shiftUids,proto3" json:"shiftUids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupMountsReq) Reset()         { *m = SupMountsReq{} }
func (m *SupMountsReq) String() string { return proto.CompactTextString(m) }
func (*SupMountsReq) ProtoMessage()    {}
func (*SupMountsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{7}
}

func (m *SupMountsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupMountsReq.Unmarshal(m, b)
}
func (m *SupMountsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupMountsReq.Marshal(b, m, deterministic)
}
func (m *SupMountsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupMountsReq.Merge(m, src)
}
func (m *SupMountsReq) XXX_Size() int {
	return xxx_messageInfo_SupMountsReq.Size(m)
}
func (m *SupMountsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SupMountsReq.DiscardUnknown(m)
}

var xxx_messageInfo_SupMountsReq proto.InternalMessageInfo

func (m *SupMountsReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SupMountsReq) GetRootfs() string {
	if m != nil {
		return m.Rootfs
	}
	return ""
}

func (m *SupMountsReq) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *SupMountsReq) GetGid() uint32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *SupMountsReq) GetShiftUids() bool {
	if m != nil {
		return m.ShiftUids
	}
	return false
}

type SupMountsReqResp struct {
	Mounts               []*Mount `protobuf:"bytes,1,rep,name=mounts,proto3" json:"mounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupMountsReqResp) Reset()         { *m = SupMountsReqResp{} }
func (m *SupMountsReqResp) String() string { return proto.CompactTextString(m) }
func (*SupMountsReqResp) ProtoMessage()    {}
func (*SupMountsReqResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{8}
}

func (m *SupMountsReqResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupMountsReqResp.Unmarshal(m, b)
}
func (m *SupMountsReqResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupMountsReqResp.Marshal(b, m, deterministic)
}
func (m *SupMountsReqResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupMountsReqResp.Merge(m, src)
}
func (m *SupMountsReqResp) XXX_Size() int {
	return xxx_messageInfo_SupMountsReqResp.Size(m)
}
func (m *SupMountsReqResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SupMountsReqResp.DiscardUnknown(m)
}

var xxx_messageInfo_SupMountsReqResp proto.InternalMessageInfo

func (m *SupMountsReqResp) GetMounts() []*Mount {
	if m != nil {
		return m.Mounts
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterReq)(nil), "protobuf.RegisterReq")
	proto.RegisterType((*RegisterResp)(nil), "protobuf.RegisterResp")
	proto.RegisterType((*UnregisterReq)(nil), "protobuf.UnregisterReq")
	proto.RegisterType((*UnregisterResp)(nil), "protobuf.UnregisterResp")
	proto.RegisterType((*SubidAllocReq)(nil), "protobuf.SubidAllocReq")
	proto.RegisterType((*SubidAllocResp)(nil), "protobuf.SubidAllocResp")
	proto.RegisterType((*Mount)(nil), "protobuf.Mount")
	proto.RegisterType((*SupMountsReq)(nil), "protobuf.SupMountsReq")
	proto.RegisterType((*SupMountsReqResp)(nil), "protobuf.SupMountsReqResp")
}

func init() { proto.RegisterFile("protobuf.proto", fileDescriptor_c77a803fcbc0c059) }

var fileDescriptor_c77a803fcbc0c059 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x3d, 0x6f, 0xdb, 0x30,
	0x10, 0x8d, 0x3e, 0x22, 0x58, 0x97, 0x48, 0x35, 0x08, 0x54, 0x51, 0x85, 0x16, 0x15, 0xb8, 0x44,
	0x93, 0x86, 0xa4, 0x5b, 0xa6, 0xa4, 0x5d, 0x03, 0x14, 0x34, 0x32, 0x74, 0x8c, 0x2d, 0x5a, 0x26,
	0xe0, 0x8a, 0x34, 0x49, 0x15, 0x75, 0x7f, 0x4e, 0x7f, 0x69, 0x41, 0x5a, 0x5f, 0x46, 0x85, 0x6c,
	0x8f, 0xef, 0xee, 0xbd, 0x3b, 0xdd, 0x9d, 0x20, 0x16, 0x92, 0x6b, 0xbe, 0x6e, 0xb7, 0xa5, 0x05,
	0x68, 0xd1, 0xbf, 0xf1, 0x27, 0xb8, 0x22, 0xb4, 0x66, 0x4a, 0x53, 0x49, 0xe8, 0x01, 0xc5, 0xe0,
	0xb2, 0x2a, 0x75, 0x72, 0xa7, 0x08, 0x89, 0xcb, 0x2a, 0x1c, 0xc3, 0xf5, 0x18, 0x56, 0x02, 0x7f,
	0x86, 0xe8, 0xa5, 0x91, 0x6f, 0x08, 0x96, 0x10, 0x4f, 0x13, 0x94, 0xc0, 0xf7, 0x10, 0xad, 0xda,
	0x35, 0xab, 0x1e, 0xf7, 0x7b, 0xbe, 0x99, 0x91, 0x20, 0x04, 0xbe, 0x62, 0x7f, 0x68, 0xea, 0xe6,
	0x4e, 0xe1, 0x13, 0x8b, 0xf1, 0x17, 0x88, 0xa7, 0x22, 0x25, 0xd0, 0x12, 0xbc, 0xb6, 0x93, 0x45,
	0xc4, 0x40, 0xc3, 0xd4, 0xac, 0xb2, 0xb2, 0x88, 0x18, 0x88, 0x7f, 0xc0, 0xe5, 0x33, 0x6f, 0x1b,
	0x8d, 0x12, 0x08, 0x14, 0x6f, 0xe5, 0x86, 0x76, 0x65, 0xba, 0x97, 0x29, 0x55, 0x51, 0xa5, 0xad,
	0x26, 0x24, 0x16, 0x1b, 0x4e, 0x1f, 0x05, 0x4d, 0xbd, 0x13, 0x67, 0xb0, 0xb1, 0xe6, 0x42, 0xa7,
	0x7e, 0xee, 0x15, 0x21, 0x31, 0x10, 0xff, 0x86, 0xeb, 0x55, 0x2b, 0xac, 0xbb, 0x9a, 0xfb, 0x88,
	0x04, 0x02, 0xc9, 0xb9, 0xde, 0xaa, 0xce, 0xbb, 0x7b, 0xf5, 0x6d, 0x7b, 0xff, 0xb5, 0xed, 0x0f,
	0x6d, 0xa3, 0x8f, 0x10, 0xaa, 0x1d, 0xdb, 0xea, 0x17, 0x56, 0xa9, 0xf4, 0x32, 0x77, 0x8a, 0x05,
	0x19, 0x09, 0xfc, 0x00, 0xcb, 0x69, 0x65, 0x3b, 0x8c, 0x5b, 0x08, 0x7e, 0x5a, 0x22, 0x75, 0x72,
	0xaf, 0xb8, 0xba, 0x7b, 0x57, 0x0e, 0x0b, 0xb6, 0x89, 0xa4, 0x0b, 0xdf, 0xfd, 0x75, 0xe1, 0x46,
	0x1d, 0xd5, 0x2f, 0xa6, 0xb8, 0x7c, 0xae, 0xe5, 0x4a, 0xbf, 0x6a, 0xfa, 0x75, 0xf7, 0xda, 0x34,
	0x74, 0x8f, 0x1e, 0x60, 0xd1, 0xef, 0x16, 0xbd, 0x1f, 0x0d, 0x26, 0xe7, 0x90, 0x25, 0x73, 0xb4,
	0x12, 0xf8, 0x02, 0x3d, 0x02, 0x8c, 0x7b, 0x46, 0x37, 0x63, 0xde, 0xd9, 0x79, 0x64, 0xe9, 0x7c,
	0xa0, 0xb7, 0x18, 0x77, 0x3c, 0xb5, 0x38, 0x3b, 0x97, 0xa9, 0xc5, 0xf9, 0x49, 0xe0, 0x0b, 0xf4,
	0xcd, 0x9c, 0xe7, 0x61, 0x18, 0x0f, 0x4a, 0xa6, 0xb9, 0xe3, 0xcc, 0xb2, 0x6c, 0x9e, 0x3f, 0xb9,
	0x3c, 0xdd, 0xc2, 0x07, 0xc6, 0xcb, 0x5a, 0x8a, 0x4d, 0xd9, 0xcf, 0x6a, 0xc8, 0x7f, 0x1a, 0x7e,
	0x95, 0xef, 0xce, 0x3a, 0xb0, 0xf8, 0xfe, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0xfb, 0xed,
	0xa8, 0x4f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SysvisorMgrStateChannelClient is the client API for SysvisorMgrStateChannel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SysvisorMgrStateChannelClient interface {
	// Container registration
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	// Container Unregistration
	Unregister(ctx context.Context, in *UnregisterReq, opts ...grpc.CallOption) (*UnregisterResp, error)
	// Subuid(gid) allocation request
	SubidAlloc(ctx context.Context, in *SubidAllocReq, opts ...grpc.CallOption) (*SubidAllocResp, error)
	// Supplementary mount request
	ReqSupMounts(ctx context.Context, in *SupMountsReq, opts ...grpc.CallOption) (*SupMountsReqResp, error)
}

type sysvisorMgrStateChannelClient struct {
	cc *grpc.ClientConn
}

func NewSysvisorMgrStateChannelClient(cc *grpc.ClientConn) SysvisorMgrStateChannelClient {
	return &sysvisorMgrStateChannelClient{cc}
}

func (c *sysvisorMgrStateChannelClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, "/protobuf.sysvisorMgrStateChannel/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysvisorMgrStateChannelClient) Unregister(ctx context.Context, in *UnregisterReq, opts ...grpc.CallOption) (*UnregisterResp, error) {
	out := new(UnregisterResp)
	err := c.cc.Invoke(ctx, "/protobuf.sysvisorMgrStateChannel/Unregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysvisorMgrStateChannelClient) SubidAlloc(ctx context.Context, in *SubidAllocReq, opts ...grpc.CallOption) (*SubidAllocResp, error) {
	out := new(SubidAllocResp)
	err := c.cc.Invoke(ctx, "/protobuf.sysvisorMgrStateChannel/SubidAlloc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysvisorMgrStateChannelClient) ReqSupMounts(ctx context.Context, in *SupMountsReq, opts ...grpc.CallOption) (*SupMountsReqResp, error) {
	out := new(SupMountsReqResp)
	err := c.cc.Invoke(ctx, "/protobuf.sysvisorMgrStateChannel/ReqSupMounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysvisorMgrStateChannelServer is the server API for SysvisorMgrStateChannel service.
type SysvisorMgrStateChannelServer interface {
	// Container registration
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	// Container Unregistration
	Unregister(context.Context, *UnregisterReq) (*UnregisterResp, error)
	// Subuid(gid) allocation request
	SubidAlloc(context.Context, *SubidAllocReq) (*SubidAllocResp, error)
	// Supplementary mount request
	ReqSupMounts(context.Context, *SupMountsReq) (*SupMountsReqResp, error)
}

func RegisterSysvisorMgrStateChannelServer(s *grpc.Server, srv SysvisorMgrStateChannelServer) {
	s.RegisterService(&_SysvisorMgrStateChannel_serviceDesc, srv)
}

func _SysvisorMgrStateChannel_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysvisorMgrStateChannelServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.sysvisorMgrStateChannel/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysvisorMgrStateChannelServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysvisorMgrStateChannel_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysvisorMgrStateChannelServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.sysvisorMgrStateChannel/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysvisorMgrStateChannelServer).Unregister(ctx, req.(*UnregisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysvisorMgrStateChannel_SubidAlloc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubidAllocReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysvisorMgrStateChannelServer).SubidAlloc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.sysvisorMgrStateChannel/SubidAlloc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysvisorMgrStateChannelServer).SubidAlloc(ctx, req.(*SubidAllocReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysvisorMgrStateChannel_ReqSupMounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupMountsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysvisorMgrStateChannelServer).ReqSupMounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.sysvisorMgrStateChannel/ReqSupMounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysvisorMgrStateChannelServer).ReqSupMounts(ctx, req.(*SupMountsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SysvisorMgrStateChannel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.sysvisorMgrStateChannel",
	HandlerType: (*SysvisorMgrStateChannelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _SysvisorMgrStateChannel_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _SysvisorMgrStateChannel_Unregister_Handler,
		},
		{
			MethodName: "SubidAlloc",
			Handler:    _SysvisorMgrStateChannel_SubidAlloc_Handler,
		},
		{
			MethodName: "ReqSupMounts",
			Handler:    _SysvisorMgrStateChannel_ReqSupMounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf.proto",
}
