// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: base.proto

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

type SESSION int32

const (
	SESSION_UID       SESSION = 0
	SESSION_DOMAIN_ID SESSION = 1
	SESSION_ROLE_IDS  SESSION = 2
)

// Enum value maps for SESSION.
var (
	SESSION_name = map[int32]string{
		0: "UID",
		1: "DOMAIN_ID",
		2: "ROLE_IDS",
	}
	SESSION_value = map[string]int32{
		"UID":       0,
		"DOMAIN_ID": 1,
		"ROLE_IDS":  2,
	}
)

func (x SESSION) Enum() *SESSION {
	p := new(SESSION)
	*p = x
	return p
}

func (x SESSION) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SESSION) Descriptor() protoreflect.EnumDescriptor {
	return file_base_proto_enumTypes[0].Descriptor()
}

func (SESSION) Type() protoreflect.EnumType {
	return &file_base_proto_enumTypes[0]
}

func (x SESSION) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SESSION.Descriptor instead.
func (SESSION) EnumDescriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

type DefaultRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"msg"
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	// @inject_tag: json:"data"
	Data string `protobuf:"bytes,3,opt,name=Data,proto3" json:"data"`
}

func (x *DefaultRes) Reset() {
	*x = DefaultRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultRes) ProtoMessage() {}

func (x *DefaultRes) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultRes.ProtoReflect.Descriptor instead.
func (*DefaultRes) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

func (x *DefaultRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DefaultRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *DefaultRes) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type EmptyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReq) Reset() {
	*x = EmptyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReq.ProtoReflect.Descriptor instead.
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

// 主键请求
type DefaultPkReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Pk:
	//	*DefaultPkReq_ID
	//	*DefaultPkReq_Code
	Pk isDefaultPkReq_Pk `protobuf_oneof:"pk"`
	// @gotags: json:"domain_id" binding:"required"
	DomainID int64 `protobuf:"varint,3,opt,name=DomainID,proto3" json:"domain_id" binding:"required"`
}

func (x *DefaultPkReq) Reset() {
	*x = DefaultPkReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultPkReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultPkReq) ProtoMessage() {}

func (x *DefaultPkReq) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultPkReq.ProtoReflect.Descriptor instead.
func (*DefaultPkReq) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{2}
}

func (m *DefaultPkReq) GetPk() isDefaultPkReq_Pk {
	if m != nil {
		return m.Pk
	}
	return nil
}

func (x *DefaultPkReq) GetID() int64 {
	if x, ok := x.GetPk().(*DefaultPkReq_ID); ok {
		return x.ID
	}
	return 0
}

func (x *DefaultPkReq) GetCode() string {
	if x, ok := x.GetPk().(*DefaultPkReq_Code); ok {
		return x.Code
	}
	return ""
}

func (x *DefaultPkReq) GetDomainID() int64 {
	if x != nil {
		return x.DomainID
	}
	return 0
}

type isDefaultPkReq_Pk interface {
	isDefaultPkReq_Pk()
}

type DefaultPkReq_ID struct {
	// @gotags: json:"id"
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3,oneof" json:"id"`
}

type DefaultPkReq_Code struct {
	// @gotags: json:"code"
	Code string `protobuf:"bytes,2,opt,name=Code,proto3,oneof" json:"code"`
}

func (*DefaultPkReq_ID) isDefaultPkReq_Pk() {}

func (*DefaultPkReq_Code) isDefaultPkReq_Pk() {}

// 主键请求
type DomainIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainID int64 `protobuf:"varint,3,opt,name=DomainID,proto3" json:"DomainID,omitempty"`
}

func (x *DomainIDReq) Reset() {
	*x = DomainIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainIDReq) ProtoMessage() {}

func (x *DomainIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainIDReq.ProtoReflect.Descriptor instead.
func (*DomainIDReq) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{3}
}

func (x *DomainIDReq) GetDomainID() int64 {
	if x != nil {
		return x.DomainID
	}
	return 0
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID      int64   `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
	DomainID int64   `protobuf:"varint,2,opt,name=DomainID,proto3" json:"DomainID,omitempty"`
	RoleIDs  []int64 `protobuf:"varint,3,rep,packed,name=RoleIDs,proto3" json:"RoleIDs,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{4}
}

func (x *Session) GetUID() int64 {
	if x != nil {
		return x.UID
	}
	return 0
}

func (x *Session) GetDomainID() int64 {
	if x != nil {
		return x.DomainID
	}
	return 0
}

func (x *Session) GetRoleIDs() []int64 {
	if x != nil {
		return x.RoleIDs
	}
	return nil
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x62,
	0x61, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0a, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x22, 0x58, 0x0a, 0x0c, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x42,
	0x04, 0x0a, 0x02, 0x70, 0x6b, 0x22, 0x29, 0x0a, 0x0b, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44,
	0x22, 0x51, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x55,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x6f, 0x6c,
	0x65, 0x49, 0x44, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x44, 0x73, 0x2a, 0x2f, 0x0a, 0x07, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x12, 0x07,
	0x0a, 0x03, 0x55, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x4f, 0x4d, 0x41, 0x49,
	0x4e, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x49,
	0x44, 0x53, 0x10, 0x02, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2e, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x5f,
	0x70, 0x62, 0x3b, 0x72, 0x62, 0x61, 0x63, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_base_proto_goTypes = []interface{}{
	(SESSION)(0),         // 0: rbacProto.SESSION
	(*DefaultRes)(nil),   // 1: rbacProto.DefaultRes
	(*EmptyReq)(nil),     // 2: rbacProto.EmptyReq
	(*DefaultPkReq)(nil), // 3: rbacProto.DefaultPkReq
	(*DomainIDReq)(nil),  // 4: rbacProto.DomainIDReq
	(*Session)(nil),      // 5: rbacProto.Session
}
var file_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultRes); i {
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
		file_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReq); i {
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
		file_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultPkReq); i {
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
		file_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainIDReq); i {
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
		file_base_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
	file_base_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*DefaultPkReq_ID)(nil),
		(*DefaultPkReq_Code)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		EnumInfos:         file_base_proto_enumTypes,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}