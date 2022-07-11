// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.1
// source: permission.entity.proto

package rbac_pb

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

type PermissionListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: form:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty" form:"name"`
	// @gotags: form:"page_size"
	PageSize int64 `protobuf:"varint,2,opt,name=PageSize,proto3" json:"PageSize,omitempty" form:"page_size"`
	// @gotags: form:"page"
	Page int64 `protobuf:"varint,3,opt,name=Page,proto3" json:"Page,omitempty" form:"page"`
	// @gotags: form:"-"
	DomainID int64 `protobuf:"varint,4,opt,name=DomainID,proto3" json:"DomainID,omitempty" form:"-"`
	// @gotags: form:"status"
	Status int64 `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty" form:"status"`
}

func (x *PermissionListReq) Reset() {
	*x = PermissionListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionListReq) ProtoMessage() {}

func (x *PermissionListReq) ProtoReflect() protoreflect.Message {
	mi := &file_permission_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionListReq.ProtoReflect.Descriptor instead.
func (*PermissionListReq) Descriptor() ([]byte, []int) {
	return file_permission_entity_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PermissionListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PermissionListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PermissionListReq) GetDomainID() int64 {
	if x != nil {
		return x.DomainID
	}
	return 0
}

func (x *PermissionListReq) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type PermissionCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"name" binding:"required"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" binding:"required"`
	// @gotags: json:"desc" binding:"required"
	Desc string `protobuf:"bytes,2,opt,name=Desc,proto3" json:"desc" binding:"required"`
	// @gotags: json:"domain_id"
	DomainID int64 `protobuf:"varint,3,opt,name=DomainID,proto3" json:"domain_id"`
	// @gotags: json:"status"
	Status int64 `protobuf:"varint,4,opt,name=Status,proto3" json:"status"`
}

func (x *PermissionCreateReq) Reset() {
	*x = PermissionCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_entity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionCreateReq) ProtoMessage() {}

func (x *PermissionCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_permission_entity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionCreateReq.ProtoReflect.Descriptor instead.
func (*PermissionCreateReq) Descriptor() ([]byte, []int) {
	return file_permission_entity_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionCreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PermissionCreateReq) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *PermissionCreateReq) GetDomainID() int64 {
	if x != nil {
		return x.DomainID
	}
	return 0
}

func (x *PermissionCreateReq) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type PermissionUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id" binding:"required"
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"id" binding:"required"`
	// @gotags: json:"name" binding:"required"
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"name" binding:"required"`
	// @gotags: json:"desc" binding:"required"
	Desc string `protobuf:"bytes,3,opt,name=Desc,proto3" json:"desc" binding:"required"`
	// @gotags: json:"domain_id"
	DomainID int64 `protobuf:"varint,4,opt,name=DomainID,proto3" json:"domain_id"`
	// @gotags: json:"status"
	Status int64 `protobuf:"varint,5,opt,name=Status,proto3" json:"status"`
}

func (x *PermissionUpdateReq) Reset() {
	*x = PermissionUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_entity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionUpdateReq) ProtoMessage() {}

func (x *PermissionUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_permission_entity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionUpdateReq.ProtoReflect.Descriptor instead.
func (*PermissionUpdateReq) Descriptor() ([]byte, []int) {
	return file_permission_entity_proto_rawDescGZIP(), []int{2}
}

func (x *PermissionUpdateReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PermissionUpdateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PermissionUpdateReq) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *PermissionUpdateReq) GetDomainID() int64 {
	if x != nil {
		return x.DomainID
	}
	return 0
}

func (x *PermissionUpdateReq) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type PermissionDelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id" binding:"required"
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"id" binding:"required"`
	// @gotags: json:"domain_id" binding:"required"
	DomainID int64 `protobuf:"varint,2,opt,name=DomainID,proto3" json:"domain_id" binding:"required"`
}

func (x *PermissionDelReq) Reset() {
	*x = PermissionDelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_entity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionDelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionDelReq) ProtoMessage() {}

func (x *PermissionDelReq) ProtoReflect() protoreflect.Message {
	mi := &file_permission_entity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionDelReq.ProtoReflect.Descriptor instead.
func (*PermissionDelReq) Descriptor() ([]byte, []int) {
	return file_permission_entity_proto_rawDescGZIP(), []int{3}
}

func (x *PermissionDelReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PermissionDelReq) GetDomainID() int64 {
	if x != nil {
		return x.DomainID
	}
	return 0
}

type PermissionPathListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: form:"permission_id"
	PermissionID int64 `protobuf:"varint,1,opt,name=PermissionID,proto3" json:"PermissionID,omitempty" form:"permission_id"`
	// @gotags: form:"domain_id"
	DomainID int64 `protobuf:"varint,4,opt,name=DomainID,proto3" json:"DomainID,omitempty" form:"domain_id"`
}

func (x *PermissionPathListReq) Reset() {
	*x = PermissionPathListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_entity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionPathListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionPathListReq) ProtoMessage() {}

func (x *PermissionPathListReq) ProtoReflect() protoreflect.Message {
	mi := &file_permission_entity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionPathListReq.ProtoReflect.Descriptor instead.
func (*PermissionPathListReq) Descriptor() ([]byte, []int) {
	return file_permission_entity_proto_rawDescGZIP(), []int{4}
}

func (x *PermissionPathListReq) GetPermissionID() int64 {
	if x != nil {
		return x.PermissionID
	}
	return 0
}

func (x *PermissionPathListReq) GetDomainID() int64 {
	if x != nil {
		return x.DomainID
	}
	return 0
}

type PermissionPathSetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"permission_id"
	PermissionID int64 `protobuf:"varint,1,opt,name=PermissionID,proto3" json:"permission_id"`
	// @gotags: json:"path_ids" binding:"required"
	PathIDs []int64 `protobuf:"varint,2,rep,packed,name=PathIDs,proto3" json:"path_ids" binding:"required"`
	// @gotags: json:"domain_id"
	DomainID int64 `protobuf:"varint,3,opt,name=DomainID,proto3" json:"domain_id"`
}

func (x *PermissionPathSetReq) Reset() {
	*x = PermissionPathSetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_entity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionPathSetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionPathSetReq) ProtoMessage() {}

func (x *PermissionPathSetReq) ProtoReflect() protoreflect.Message {
	mi := &file_permission_entity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionPathSetReq.ProtoReflect.Descriptor instead.
func (*PermissionPathSetReq) Descriptor() ([]byte, []int) {
	return file_permission_entity_proto_rawDescGZIP(), []int{5}
}

func (x *PermissionPathSetReq) GetPermissionID() int64 {
	if x != nil {
		return x.PermissionID
	}
	return 0
}

func (x *PermissionPathSetReq) GetPathIDs() []int64 {
	if x != nil {
		return x.PathIDs
	}
	return nil
}

func (x *PermissionPathSetReq) GetDomainID() int64 {
	if x != nil {
		return x.DomainID
	}
	return 0
}

type GetDomainPolicyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainID int64 `protobuf:"varint,3,opt,name=DomainID,proto3" json:"DomainID,omitempty"`
}

func (x *GetDomainPolicyReq) Reset() {
	*x = GetDomainPolicyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_entity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDomainPolicyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDomainPolicyReq) ProtoMessage() {}

func (x *GetDomainPolicyReq) ProtoReflect() protoreflect.Message {
	mi := &file_permission_entity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDomainPolicyReq.ProtoReflect.Descriptor instead.
func (*GetDomainPolicyReq) Descriptor() ([]byte, []int) {
	return file_permission_entity_proto_rawDescGZIP(), []int{6}
}

func (x *GetDomainPolicyReq) GetDomainID() int64 {
	if x != nil {
		return x.DomainID
	}
	return 0
}

var File_permission_entity_proto protoreflect.FileDescriptor

var file_permission_entity_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x62, 0x61, 0x63, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x71, 0x0a, 0x13, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x65, 0x73,
	0x63, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3e, 0x0a, 0x10, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x22, 0x57, 0x0a, 0x15, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x49, 0x44, 0x22, 0x70, 0x0a, 0x14, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x18,
	0x0a, 0x07, 0x50, 0x61, 0x74, 0x68, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x07, 0x50, 0x61, 0x74, 0x68, 0x49, 0x44, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x49, 0x44, 0x22, 0x30, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2e, 0x2f, 0x72, 0x62, 0x61,
	0x63, 0x5f, 0x70, 0x62, 0x3b, 0x72, 0x62, 0x61, 0x63, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_permission_entity_proto_rawDescOnce sync.Once
	file_permission_entity_proto_rawDescData = file_permission_entity_proto_rawDesc
)

func file_permission_entity_proto_rawDescGZIP() []byte {
	file_permission_entity_proto_rawDescOnce.Do(func() {
		file_permission_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_permission_entity_proto_rawDescData)
	})
	return file_permission_entity_proto_rawDescData
}

var file_permission_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_permission_entity_proto_goTypes = []interface{}{
	(*PermissionListReq)(nil),     // 0: rbacProto.PermissionListReq
	(*PermissionCreateReq)(nil),   // 1: rbacProto.PermissionCreateReq
	(*PermissionUpdateReq)(nil),   // 2: rbacProto.PermissionUpdateReq
	(*PermissionDelReq)(nil),      // 3: rbacProto.PermissionDelReq
	(*PermissionPathListReq)(nil), // 4: rbacProto.PermissionPathListReq
	(*PermissionPathSetReq)(nil),  // 5: rbacProto.PermissionPathSetReq
	(*GetDomainPolicyReq)(nil),    // 6: rbacProto.GetDomainPolicyReq
}
var file_permission_entity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_permission_entity_proto_init() }
func file_permission_entity_proto_init() {
	if File_permission_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_permission_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionListReq); i {
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
		file_permission_entity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionCreateReq); i {
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
		file_permission_entity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionUpdateReq); i {
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
		file_permission_entity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionDelReq); i {
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
		file_permission_entity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionPathListReq); i {
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
		file_permission_entity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionPathSetReq); i {
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
		file_permission_entity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDomainPolicyReq); i {
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
			RawDescriptor: file_permission_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_permission_entity_proto_goTypes,
		DependencyIndexes: file_permission_entity_proto_depIdxs,
		MessageInfos:      file_permission_entity_proto_msgTypes,
	}.Build()
	File_permission_entity_proto = out.File
	file_permission_entity_proto_rawDesc = nil
	file_permission_entity_proto_goTypes = nil
	file_permission_entity_proto_depIdxs = nil
}
