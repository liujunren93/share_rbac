// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: path.entity.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PathCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"-" binding:"required"
	DomainID int64 `protobuf:"varint,1,opt,name=DomainID,proto3" json:"-" binding:"required"`
	// @gotags: json:"component" binding:"required_if=PathType 1"
	Component string `protobuf:"bytes,2,opt,name=Component,proto3" json:"component" binding:"required_if=PathType 1"`
	// @gotags: json:"redirect"
	Redirect string `protobuf:"bytes,3,opt,name=Redirect,proto3" json:"redirect"`
	// @gotags: json:"parent_id"
	ParentID int64 `protobuf:"varint,4,opt,name=ParentID,proto3" json:"parent_id"`
	// @gotags: json:"path"
	Path string `protobuf:"bytes,5,opt,name=Path,proto3" json:"path"`
	// @gotags: json:"name" binding:"required"
	Name string `protobuf:"bytes,6,opt,name=Name,proto3" json:"name" binding:"required"`
	// @gotags: json:"path_type" binding:"required"
	PathType int64 `protobuf:"varint,8,opt,name=PathType,proto3" json:"path_type" binding:"required"`
	// @gotags: json:"meta"
	Meta map[string]string `protobuf:"bytes,9,rep,name=Meta,proto3" json:"meta" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// @gotags: json:"api_path" binding:"required"
	ApiPath string `protobuf:"bytes,10,opt,name=ApiPath,proto3" json:"api_path" binding:"required"`
	// @gotags: json:"method"
	Method string `protobuf:"bytes,11,opt,name=Method,proto3" json:"method"`
	// @gotags: json:"action"
	Action string `protobuf:"bytes,12,opt,name=Action,proto3" json:"action"`
	// @gotags: json:"key"
	Key string `protobuf:"bytes,13,opt,name=Key,proto3" json:"key"`
}

func (x *PathCreateReq) Reset() {
	*x = PathCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_path_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathCreateReq) ProtoMessage() {}

func (x *PathCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_path_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathCreateReq.ProtoReflect.Descriptor instead.
func (*PathCreateReq) Descriptor() ([]byte, []int) {
	return file_path_entity_proto_rawDescGZIP(), []int{0}
}

func (x *PathCreateReq) GetDomainID() int64 {
	if x != nil {
		return x.DomainID
	}
	return 0
}

func (x *PathCreateReq) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *PathCreateReq) GetRedirect() string {
	if x != nil {
		return x.Redirect
	}
	return ""
}

func (x *PathCreateReq) GetParentID() int64 {
	if x != nil {
		return x.ParentID
	}
	return 0
}

func (x *PathCreateReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PathCreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PathCreateReq) GetPathType() int64 {
	if x != nil {
		return x.PathType
	}
	return 0
}

func (x *PathCreateReq) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *PathCreateReq) GetApiPath() string {
	if x != nil {
		return x.ApiPath
	}
	return ""
}

func (x *PathCreateReq) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *PathCreateReq) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *PathCreateReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type PathUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"-" binding:"required"
	DomainID int64 `protobuf:"varint,1,opt,name=DomainID,proto3" json:"-" binding:"required"`
	// @gotags: json:"component" binding:"required_if=PathType 1"
	Component string `protobuf:"bytes,2,opt,name=Component,proto3" json:"component" binding:"required_if=PathType 1"`
	// @gotags: json:"redirect"
	Redirect string `protobuf:"bytes,3,opt,name=Redirect,proto3" json:"redirect"`
	// @gotags: json:"parent_id"
	ParentID int64 `protobuf:"varint,4,opt,name=ParentID,proto3" json:"parent_id"`
	// @gotags: json:"path"
	Path string `protobuf:"bytes,5,opt,name=Path,proto3" json:"path"`
	// @gotags: json:"id" binding:"required"
	ID int64 `protobuf:"varint,6,opt,name=ID,proto3" json:"id" binding:"required"`
	// @gotags: json:"name" binding:"required"
	Name string `protobuf:"bytes,7,opt,name=Name,proto3" json:"name" binding:"required"`
	// @gotags: json:"path_type" binding:"required"
	PathType int64 `protobuf:"varint,8,opt,name=PathType,proto3" json:"path_type" binding:"required"`
	// @gotags: json:"meta"
	Meta map[string]string `protobuf:"bytes,9,rep,name=Meta,proto3" json:"meta" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// @gotags: json:"api_path" binding:"required"
	ApiPath string `protobuf:"bytes,10,opt,name=ApiPath,proto3" json:"api_path" binding:"required"`
	// @gotags: json:"method"
	Method string `protobuf:"bytes,11,opt,name=Method,proto3" json:"method"`
	// @gotags: json:"action"
	Action string `protobuf:"bytes,12,opt,name=Action,proto3" json:"action"`
	// @gotags: json:"key"
	Key string `protobuf:"bytes,13,opt,name=Key,proto3" json:"key"`
}

func (x *PathUpdateReq) Reset() {
	*x = PathUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_path_entity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathUpdateReq) ProtoMessage() {}

func (x *PathUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_path_entity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathUpdateReq.ProtoReflect.Descriptor instead.
func (*PathUpdateReq) Descriptor() ([]byte, []int) {
	return file_path_entity_proto_rawDescGZIP(), []int{1}
}

func (x *PathUpdateReq) GetDomainID() int64 {
	if x != nil {
		return x.DomainID
	}
	return 0
}

func (x *PathUpdateReq) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *PathUpdateReq) GetRedirect() string {
	if x != nil {
		return x.Redirect
	}
	return ""
}

func (x *PathUpdateReq) GetParentID() int64 {
	if x != nil {
		return x.ParentID
	}
	return 0
}

func (x *PathUpdateReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PathUpdateReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PathUpdateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PathUpdateReq) GetPathType() int64 {
	if x != nil {
		return x.PathType
	}
	return 0
}

func (x *PathUpdateReq) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *PathUpdateReq) GetApiPath() string {
	if x != nil {
		return x.ApiPath
	}
	return ""
}

func (x *PathUpdateReq) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *PathUpdateReq) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *PathUpdateReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type PathListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
	DomainID int64 `protobuf:"varint,6,opt,name=DomainID,proto3" json:"DomainID,omitempty" form:"-"`
}

func (x *PathListReq) Reset() {
	*x = PathListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_path_entity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathListReq) ProtoMessage() {}

func (x *PathListReq) ProtoReflect() protoreflect.Message {
	mi := &file_path_entity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathListReq.ProtoReflect.Descriptor instead.
func (*PathListReq) Descriptor() ([]byte, []int) {
	return file_path_entity_proto_rawDescGZIP(), []int{2}
}

func (x *PathListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PathListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PathListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PathListReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PathListReq) GetPathType() int64 {
	if x != nil {
		return x.PathType
	}
	return 0
}

func (x *PathListReq) GetDomainID() int64 {
	if x != nil {
		return x.DomainID
	}
	return 0
}

var File_path_entity_proto protoreflect.FileDescriptor

var file_path_entity_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x61, 0x74, 0x68, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x03, 0x0a, 0x0d, 0x50,
	0x61, 0x74, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x74, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x74, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x70, 0x69, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70, 0x69, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b,
	0x65, 0x79, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9e, 0x03, 0x0a, 0x0d,
	0x50, 0x61, 0x74, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a,
	0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x74, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x74, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x70, 0x69, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70, 0x69, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4b, 0x65, 0x79, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9d, 0x01, 0x0a,
	0x0b, 0x50, 0x61, 0x74, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_path_entity_proto_rawDescOnce sync.Once
	file_path_entity_proto_rawDescData = file_path_entity_proto_rawDesc
)

func file_path_entity_proto_rawDescGZIP() []byte {
	file_path_entity_proto_rawDescOnce.Do(func() {
		file_path_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_path_entity_proto_rawDescData)
	})
	return file_path_entity_proto_rawDescData
}

var file_path_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_path_entity_proto_goTypes = []interface{}{
	(*PathCreateReq)(nil), // 0: proto.PathCreateReq
	(*PathUpdateReq)(nil), // 1: proto.PathUpdateReq
	(*PathListReq)(nil),   // 2: proto.PathListReq
	nil,                   // 3: proto.PathCreateReq.MetaEntry
	nil,                   // 4: proto.PathUpdateReq.MetaEntry
}
var file_path_entity_proto_depIdxs = []int32{
	3, // 0: proto.PathCreateReq.Meta:type_name -> proto.PathCreateReq.MetaEntry
	4, // 1: proto.PathUpdateReq.Meta:type_name -> proto.PathUpdateReq.MetaEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_path_entity_proto_init() }
func file_path_entity_proto_init() {
	if File_path_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_path_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathCreateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_path_entity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathUpdateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_path_entity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_path_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_path_entity_proto_goTypes,
		DependencyIndexes: file_path_entity_proto_depIdxs,
		MessageInfos:      file_path_entity_proto_msgTypes,
	}.Build()
	File_path_entity_proto = out.File
	file_path_entity_proto_rawDesc = nil
	file_path_entity_proto_goTypes = nil
	file_path_entity_proto_depIdxs = nil
}
