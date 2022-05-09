// Code generated by protoc-gen-go. DO NOT EDIT.
// source: path.entity.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PathCreateReq struct {
	// @gotags: json:"domain_id" binding:"required"
	DomainID int64 `protobuf:"varint,1,opt,name=DomainID,proto3" json:"domain_id" binding:"required"`
	// @gotags: json:"component" binding:"required"
	Component string `protobuf:"bytes,2,opt,name=Component,proto3" json:"component" binding:"required"`
	// @gotags: json:"redirect"
	Redirect string `protobuf:"bytes,3,opt,name=Redirect,proto3" json:"redirect"`
	// @gotags: json:"parent_id"
	ParentID int64 `protobuf:"varint,4,opt,name=ParentID,proto3" json:"parent_id"`
	// @gotags: json:"path" binding:"required"
	Path string `protobuf:"bytes,5,opt,name=Path,proto3" json:"path" binding:"required"`
	// @gotags: json:"name" binding:"required"
	Name string `protobuf:"bytes,6,opt,name=Name,proto3" json:"name" binding:"required"`
	// @gotags: json:"path_type" binding:"required"
	PathType int64 `protobuf:"varint,8,opt,name=PathType,proto3" json:"path_type" binding:"required"`
	// @gotags: json:"meta"
	Meta map[string]string `protobuf:"bytes,9,rep,name=Meta,proto3" json:"meta" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// @gotags: json:"api_path" binding:"required"
	ApiPath              string   `protobuf:"bytes,10,opt,name=ApiPath,proto3" json:"api_path" binding:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PathCreateReq) Reset()         { *m = PathCreateReq{} }
func (m *PathCreateReq) String() string { return proto.CompactTextString(m) }
func (*PathCreateReq) ProtoMessage()    {}
func (*PathCreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_338b6f1ce64417a0, []int{0}
}

func (m *PathCreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathCreateReq.Unmarshal(m, b)
}
func (m *PathCreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathCreateReq.Marshal(b, m, deterministic)
}
func (m *PathCreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathCreateReq.Merge(m, src)
}
func (m *PathCreateReq) XXX_Size() int {
	return xxx_messageInfo_PathCreateReq.Size(m)
}
func (m *PathCreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PathCreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_PathCreateReq proto.InternalMessageInfo

func (m *PathCreateReq) GetDomainID() int64 {
	if m != nil {
		return m.DomainID
	}
	return 0
}

func (m *PathCreateReq) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *PathCreateReq) GetRedirect() string {
	if m != nil {
		return m.Redirect
	}
	return ""
}

func (m *PathCreateReq) GetParentID() int64 {
	if m != nil {
		return m.ParentID
	}
	return 0
}

func (m *PathCreateReq) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PathCreateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PathCreateReq) GetPathType() int64 {
	if m != nil {
		return m.PathType
	}
	return 0
}

func (m *PathCreateReq) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PathCreateReq) GetApiPath() string {
	if m != nil {
		return m.ApiPath
	}
	return ""
}

type PathUpdateReq struct {
	// @gotags: json:"domain_id" binding:"required"
	DomainID int64 `protobuf:"varint,1,opt,name=DomainID,proto3" json:"domain_id" binding:"required"`
	// @gotags: json:"component" binding:"required_if=PathType 1"
	Component string `protobuf:"bytes,2,opt,name=Component,proto3" json:"component" binding:"required_if=PathType 1"`
	// @gotags: json:"redirect"
	Redirect string `protobuf:"bytes,3,opt,name=Redirect,proto3" json:"redirect"`
	// @gotags: json:"parent_id"
	ParentID int64 `protobuf:"varint,4,opt,name=ParentID,proto3" json:"parent_id"`
	// @gotags: json:"path" binding:"required"
	Path string `protobuf:"bytes,5,opt,name=Path,proto3" json:"path" binding:"required"`
	// @gotags: json:"id" binding:"required"
	ID int64 `protobuf:"varint,6,opt,name=ID,proto3" json:"id" binding:"required"`
	// @gotags: json:"name" binding:"required"
	Name string `protobuf:"bytes,7,opt,name=Name,proto3" json:"name" binding:"required"`
	// @gotags: json:"path_type" binding:"required"
	PathType int64 `protobuf:"varint,8,opt,name=PathType,proto3" json:"path_type" binding:"required"`
	// @gotags: json:"meta"
	Meta map[string]string `protobuf:"bytes,9,rep,name=Meta,proto3" json:"meta" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// @gotags: json:"api_path" binding:"required"
	ApiPath              string   `protobuf:"bytes,10,opt,name=ApiPath,proto3" json:"api_path" binding:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PathUpdateReq) Reset()         { *m = PathUpdateReq{} }
func (m *PathUpdateReq) String() string { return proto.CompactTextString(m) }
func (*PathUpdateReq) ProtoMessage()    {}
func (*PathUpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_338b6f1ce64417a0, []int{1}
}

func (m *PathUpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathUpdateReq.Unmarshal(m, b)
}
func (m *PathUpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathUpdateReq.Marshal(b, m, deterministic)
}
func (m *PathUpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathUpdateReq.Merge(m, src)
}
func (m *PathUpdateReq) XXX_Size() int {
	return xxx_messageInfo_PathUpdateReq.Size(m)
}
func (m *PathUpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PathUpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_PathUpdateReq proto.InternalMessageInfo

func (m *PathUpdateReq) GetDomainID() int64 {
	if m != nil {
		return m.DomainID
	}
	return 0
}

func (m *PathUpdateReq) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *PathUpdateReq) GetRedirect() string {
	if m != nil {
		return m.Redirect
	}
	return ""
}

func (m *PathUpdateReq) GetParentID() int64 {
	if m != nil {
		return m.ParentID
	}
	return 0
}

func (m *PathUpdateReq) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PathUpdateReq) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *PathUpdateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PathUpdateReq) GetPathType() int64 {
	if m != nil {
		return m.PathType
	}
	return 0
}

func (m *PathUpdateReq) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PathUpdateReq) GetApiPath() string {
	if m != nil {
		return m.ApiPath
	}
	return ""
}

type PathListReq struct {
	// @gotags: form:"page_size"
	PageSize int64 `protobuf:"varint,1,opt,name=PageSize,proto3" json:"PageSize,omitempty" form:"page_size"`
	// @gotags: form:"page"
	Page int64 `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty" form:"page"`
	// @gotags: form:"name"
	Name string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty" form:"name"`
	// @gotags: form:"path"
	Path string `protobuf:"bytes,4,opt,name=Path,proto3" json:"Path,omitempty" form:"path"`
	// @gotags: form:"path_type"
	PathType int64 `protobuf:"varint,5,opt,name=PathType,proto3" json:"PathType,omitempty" form:"path_type"`
	// @gotags: form:"-"
	DomainID             int64    `protobuf:"varint,6,opt,name=DomainID,proto3" json:"DomainID,omitempty" form:"-"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PathListReq) Reset()         { *m = PathListReq{} }
func (m *PathListReq) String() string { return proto.CompactTextString(m) }
func (*PathListReq) ProtoMessage()    {}
func (*PathListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_338b6f1ce64417a0, []int{2}
}

func (m *PathListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathListReq.Unmarshal(m, b)
}
func (m *PathListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathListReq.Marshal(b, m, deterministic)
}
func (m *PathListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathListReq.Merge(m, src)
}
func (m *PathListReq) XXX_Size() int {
	return xxx_messageInfo_PathListReq.Size(m)
}
func (m *PathListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PathListReq.DiscardUnknown(m)
}

var xxx_messageInfo_PathListReq proto.InternalMessageInfo

func (m *PathListReq) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *PathListReq) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *PathListReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PathListReq) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PathListReq) GetPathType() int64 {
	if m != nil {
		return m.PathType
	}
	return 0
}

func (m *PathListReq) GetDomainID() int64 {
	if m != nil {
		return m.DomainID
	}
	return 0
}

func init() {
	proto.RegisterType((*PathCreateReq)(nil), "proto.PathCreateReq")
	proto.RegisterMapType((map[string]string)(nil), "proto.PathCreateReq.MetaEntry")
	proto.RegisterType((*PathUpdateReq)(nil), "proto.PathUpdateReq")
	proto.RegisterMapType((map[string]string)(nil), "proto.PathUpdateReq.MetaEntry")
	proto.RegisterType((*PathListReq)(nil), "proto.PathListReq")
}

func init() { proto.RegisterFile("path.entity.proto", fileDescriptor_338b6f1ce64417a0) }

var fileDescriptor_338b6f1ce64417a0 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x49, 0x26, 0x69, 0x9b, 0x53, 0xee, 0xe5, 0xde, 0xc1, 0xc5, 0x50, 0x44, 0x4a, 0x57,
	0x5d, 0x45, 0xa8, 0x0b, 0x45, 0x57, 0xda, 0xb8, 0x28, 0xa8, 0x94, 0x51, 0x37, 0xee, 0xa6, 0xf6,
	0xd0, 0x06, 0x6d, 0x32, 0xc6, 0xa3, 0x10, 0xdf, 0xc5, 0x37, 0xf2, 0x11, 0x7c, 0x18, 0x99, 0x99,
	0x26, 0xa6, 0x6e, 0x05, 0x71, 0x95, 0xf3, 0x9f, 0x39, 0xf3, 0x4f, 0xfe, 0xef, 0xc0, 0x7f, 0xad,
	0x68, 0x19, 0x63, 0x46, 0x29, 0x95, 0xb1, 0x2e, 0x72, 0xca, 0x79, 0x68, 0x3f, 0x83, 0x37, 0x1f,
	0xfe, 0x4c, 0x15, 0x2d, 0xc7, 0x05, 0x2a, 0x42, 0x89, 0x0f, 0xbc, 0x07, 0x9d, 0x24, 0x5f, 0xa9,
	0x34, 0x9b, 0x24, 0xc2, 0xeb, 0x7b, 0x43, 0x26, 0x6b, 0xcd, 0xb7, 0x21, 0x1a, 0xe7, 0x2b, 0x9d,
	0x67, 0x98, 0x91, 0xf0, 0xfb, 0xde, 0x30, 0x92, 0x9f, 0x0d, 0x73, 0x53, 0xe2, 0x3c, 0x2d, 0xf0,
	0x96, 0x04, 0xb3, 0x87, 0xb5, 0x36, 0x67, 0x53, 0x55, 0x60, 0x46, 0x93, 0x44, 0x04, 0xce, 0xb5,
	0xd2, 0x9c, 0x43, 0x60, 0x7e, 0x41, 0x84, 0xf6, 0x8e, 0xad, 0x4d, 0xef, 0x42, 0xad, 0x50, 0xb4,
	0x5c, 0xcf, 0xd4, 0xce, 0x83, 0x96, 0x57, 0xa5, 0x46, 0xd1, 0xa9, 0x3c, 0x9c, 0xe6, 0x23, 0x08,
	0xce, 0x91, 0x94, 0x88, 0xfa, 0x6c, 0xd8, 0x1d, 0xed, 0xb8, 0x90, 0xf1, 0x46, 0xb2, 0xd8, 0x0c,
	0x9c, 0x66, 0x54, 0x94, 0xd2, 0xce, 0x72, 0x01, 0xed, 0x63, 0x9d, 0xda, 0xa7, 0xc1, 0x3e, 0x53,
	0xc9, 0xde, 0x3e, 0x44, 0xf5, 0x30, 0xff, 0x07, 0xec, 0x0e, 0x4b, 0xcb, 0x22, 0x92, 0xa6, 0xe4,
	0x5b, 0x10, 0x3e, 0xab, 0xfb, 0x27, 0x5c, 0x23, 0x70, 0xe2, 0xd0, 0x3f, 0xf0, 0x06, 0xef, 0x6b,
	0x9c, 0xd7, 0x7a, 0xfe, 0xab, 0x70, 0xfe, 0x05, 0x7f, 0x92, 0x58, 0x98, 0x4c, 0xfa, 0x6e, 0xc6,
	0xe2, 0x6d, 0x7f, 0x1b, 0x6f, 0x9d, 0xf4, 0x27, 0xf0, 0xbe, 0x7a, 0xd0, 0x35, 0x0e, 0x67, 0xe9,
	0x23, 0xad, 0xe1, 0x4e, 0xd5, 0x02, 0x2f, 0xd3, 0x17, 0xac, 0xe0, 0x56, 0xda, 0x61, 0x58, 0x38,
	0x13, 0x26, 0x6d, 0x5d, 0xc7, 0x66, 0x8d, 0xd8, 0x15, 0xae, 0xa0, 0x81, 0xab, 0x89, 0x22, 0xfc,
	0x82, 0xa2, 0xb9, 0xd0, 0xd6, 0xe6, 0x42, 0x4f, 0xe0, 0xa6, 0x13, 0xc7, 0xbb, 0x7a, 0x76, 0xa4,
	0x67, 0xb3, 0x96, 0x65, 0xb4, 0xf7, 0x11, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x60, 0x3a, 0xd2, 0x7c,
	0x03, 0x00, 0x00,
}
