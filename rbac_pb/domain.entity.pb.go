// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.1
// source: domain.entity.proto

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

type DomainCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"name" binding:"required"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" binding:"required"`
	// @gotags: json:"domain" binding:"required"
	Domain string `protobuf:"bytes,2,opt,name=Domain,proto3" json:"domain" binding:"required"`
	// @gotags: json:"status" binding:"required"
	Status int32 `protobuf:"varint,3,opt,name=Status,proto3" json:"status" binding:"required"`
}

func (x *DomainCreateReq) Reset() {
	*x = DomainCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainCreateReq) ProtoMessage() {}

func (x *DomainCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_domain_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainCreateReq.ProtoReflect.Descriptor instead.
func (*DomainCreateReq) Descriptor() ([]byte, []int) {
	return file_domain_entity_proto_rawDescGZIP(), []int{0}
}

func (x *DomainCreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DomainCreateReq) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *DomainCreateReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DomainListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: form:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty" form:"name"`
	// @gotags: form:"domain"
	Domain string `protobuf:"bytes,2,opt,name=Domain,proto3" json:"Domain,omitempty" form:"domain"`
	// @gotags: form:"page_size"
	PageSize int64 `protobuf:"varint,3,opt,name=PageSize,proto3" json:"PageSize,omitempty" form:"page_size"`
	// @gotags: form:"page"
	Page int64 `protobuf:"varint,4,opt,name=Page,proto3" json:"Page,omitempty" form:"page"`
	// @gotags: form:"status"
	Status int64 `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty" form:"status"`
}

func (x *DomainListReq) Reset() {
	*x = DomainListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_entity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainListReq) ProtoMessage() {}

func (x *DomainListReq) ProtoReflect() protoreflect.Message {
	mi := &file_domain_entity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainListReq.ProtoReflect.Descriptor instead.
func (*DomainListReq) Descriptor() ([]byte, []int) {
	return file_domain_entity_proto_rawDescGZIP(), []int{1}
}

func (x *DomainListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DomainListReq) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *DomainListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *DomainListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *DomainListReq) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DomainDelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: form:"name" json:"name" binding:"required"
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"name" form:"name" binding:"required"`
}

func (x *DomainDelReq) Reset() {
	*x = DomainDelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_entity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainDelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainDelReq) ProtoMessage() {}

func (x *DomainDelReq) ProtoReflect() protoreflect.Message {
	mi := &file_domain_entity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainDelReq.ProtoReflect.Descriptor instead.
func (*DomainDelReq) Descriptor() ([]byte, []int) {
	return file_domain_entity_proto_rawDescGZIP(), []int{2}
}

func (x *DomainDelReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DomainUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"  binding:"required"
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"id" binding:"required"`
	// @gotags: json:"name" binding:"required"
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"name" binding:"required"`
	// @gotags: json:"domain" binding:"required"
	Domain string `protobuf:"bytes,3,opt,name=Domain,proto3" json:"domain" binding:"required"`
	// @gotags: json:"status" binding:"required"
	Status int32 `protobuf:"varint,4,opt,name=Status,proto3" json:"status" binding:"required"`
}

func (x *DomainUpdateReq) Reset() {
	*x = DomainUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_entity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainUpdateReq) ProtoMessage() {}

func (x *DomainUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_domain_entity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainUpdateReq.ProtoReflect.Descriptor instead.
func (*DomainUpdateReq) Descriptor() ([]byte, []int) {
	return file_domain_entity_proto_rawDescGZIP(), []int{3}
}

func (x *DomainUpdateReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DomainUpdateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DomainUpdateReq) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *DomainUpdateReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_domain_entity_proto protoreflect.FileDescriptor

var file_domain_entity_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x62, 0x61, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x55, 0x0a, 0x0f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x0d, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1e, 0x0a,
	0x0c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x65, 0x0a,
	0x0f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2e, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x5f,
	0x70, 0x62, 0x3b, 0x72, 0x62, 0x61, 0x63, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_domain_entity_proto_rawDescOnce sync.Once
	file_domain_entity_proto_rawDescData = file_domain_entity_proto_rawDesc
)

func file_domain_entity_proto_rawDescGZIP() []byte {
	file_domain_entity_proto_rawDescOnce.Do(func() {
		file_domain_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_entity_proto_rawDescData)
	})
	return file_domain_entity_proto_rawDescData
}

var file_domain_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_domain_entity_proto_goTypes = []interface{}{
	(*DomainCreateReq)(nil), // 0: rbacProto.DomainCreateReq
	(*DomainListReq)(nil),   // 1: rbacProto.DomainListReq
	(*DomainDelReq)(nil),    // 2: rbacProto.DomainDelReq
	(*DomainUpdateReq)(nil), // 3: rbacProto.DomainUpdateReq
}
var file_domain_entity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_domain_entity_proto_init() }
func file_domain_entity_proto_init() {
	if File_domain_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainCreateReq); i {
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
		file_domain_entity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainListReq); i {
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
		file_domain_entity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainDelReq); i {
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
		file_domain_entity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainUpdateReq); i {
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
			RawDescriptor: file_domain_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_entity_proto_goTypes,
		DependencyIndexes: file_domain_entity_proto_depIdxs,
		MessageInfos:      file_domain_entity_proto_msgTypes,
	}.Build()
	File_domain_entity_proto = out.File
	file_domain_entity_proto_rawDesc = nil
	file_domain_entity_proto_goTypes = nil
	file_domain_entity_proto_depIdxs = nil
}
