// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: controller/storage/credential/store/v1/credential.proto

// Package store provides protobufs for storing types in the credential package.

package store

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public_id is a surrogate key suitable for use in a public API.
	// @inject_tag: `gorm:"primary_key"`
	PublicId string `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
	// The scope_id of the owning scope and must be set.
	// @inject_tag: `gorm:"not_null"`
	ScopeId string `protobuf:"bytes,2,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty" gorm:"not_null"`
}

func (x *Store) Reset() {
	*x = Store{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_credential_store_v1_credential_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store) ProtoMessage() {}

func (x *Store) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_credential_store_v1_credential_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store.ProtoReflect.Descriptor instead.
func (*Store) Descriptor() ([]byte, []int) {
	return file_controller_storage_credential_store_v1_credential_proto_rawDescGZIP(), []int{0}
}

func (x *Store) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *Store) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

type Library struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public_id is a surrogate key suitable for use in a public API.
	// @inject_tag: `gorm:"primary_key"`
	PublicId string `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
	// The store_id of the owning store and must be set.
	// @inject_tag: `gorm:"not_null"`
	StoreId string `protobuf:"bytes,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty" gorm:"not_null"`
	// credential_type is optional. If set, it indicates the type of
	// credential the library returns.
	// @inject_tag: `gorm:"default:null"`
	CredentialType string `protobuf:"bytes,3,opt,name=credential_type,json=credentialType,proto3" json:"credential_type,omitempty" gorm:"default:null"`
}

func (x *Library) Reset() {
	*x = Library{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_credential_store_v1_credential_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Library) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Library) ProtoMessage() {}

func (x *Library) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_credential_store_v1_credential_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Library.ProtoReflect.Descriptor instead.
func (*Library) Descriptor() ([]byte, []int) {
	return file_controller_storage_credential_store_v1_credential_proto_rawDescGZIP(), []int{1}
}

func (x *Library) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *Library) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *Library) GetCredentialType() string {
	if x != nil {
		return x.CredentialType
	}
	return ""
}

type Credential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public_id is a surrogate key suitable for use in a public API.
	// @inject_tag: `gorm:"primary_key"`
	PublicId string `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
}

func (x *Credential) Reset() {
	*x = Credential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_credential_store_v1_credential_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credential) ProtoMessage() {}

func (x *Credential) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_credential_store_v1_credential_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credential.ProtoReflect.Descriptor instead.
func (*Credential) Descriptor() ([]byte, []int) {
	return file_controller_storage_credential_store_v1_credential_proto_rawDescGZIP(), []int{2}
}

func (x *Credential) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

type Static struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public_id is a surrogate key suitable for use in a public API.
	// @inject_tag: `gorm:"primary_key"`
	PublicId string `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
	// The store_id of the owning store and must be set.
	// @inject_tag: `gorm:"not_null"`
	StoreId string `protobuf:"bytes,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty" gorm:"not_null"`
}

func (x *Static) Reset() {
	*x = Static{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_credential_store_v1_credential_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Static) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Static) ProtoMessage() {}

func (x *Static) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_credential_store_v1_credential_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Static.ProtoReflect.Descriptor instead.
func (*Static) Descriptor() ([]byte, []int) {
	return file_controller_storage_credential_store_v1_credential_proto_rawDescGZIP(), []int{3}
}

func (x *Static) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *Static) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

type Dynamic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public_id is a surrogate key suitable for use in a public API.
	// @inject_tag: `gorm:"primary_key"`
	PublicId string `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
	// The library_id of the owning library and must be set.
	// @inject_tag: `gorm:"not_null"`
	LibraryId string `protobuf:"bytes,2,opt,name=library_id,json=libraryId,proto3" json:"library_id,omitempty" gorm:"not_null"`
}

func (x *Dynamic) Reset() {
	*x = Dynamic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_credential_store_v1_credential_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dynamic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dynamic) ProtoMessage() {}

func (x *Dynamic) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_credential_store_v1_credential_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dynamic.ProtoReflect.Descriptor instead.
func (*Dynamic) Descriptor() ([]byte, []int) {
	return file_controller_storage_credential_store_v1_credential_proto_rawDescGZIP(), []int{4}
}

func (x *Dynamic) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *Dynamic) GetLibraryId() string {
	if x != nil {
		return x.LibraryId
	}
	return ""
}

var File_controller_storage_credential_store_v1_credential_proto protoreflect.FileDescriptor

var file_controller_storage_credential_store_v1_credential_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x22, 0x3f, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x49, 0x64, 0x22, 0x6a, 0x0a, 0x07, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x22, 0x29,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x07, 0x44,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x49, 0x64, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3b, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_storage_credential_store_v1_credential_proto_rawDescOnce sync.Once
	file_controller_storage_credential_store_v1_credential_proto_rawDescData = file_controller_storage_credential_store_v1_credential_proto_rawDesc
)

func file_controller_storage_credential_store_v1_credential_proto_rawDescGZIP() []byte {
	file_controller_storage_credential_store_v1_credential_proto_rawDescOnce.Do(func() {
		file_controller_storage_credential_store_v1_credential_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_storage_credential_store_v1_credential_proto_rawDescData)
	})
	return file_controller_storage_credential_store_v1_credential_proto_rawDescData
}

var file_controller_storage_credential_store_v1_credential_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_controller_storage_credential_store_v1_credential_proto_goTypes = []interface{}{
	(*Store)(nil),      // 0: controller.storage.credential.store.v1.Store
	(*Library)(nil),    // 1: controller.storage.credential.store.v1.Library
	(*Credential)(nil), // 2: controller.storage.credential.store.v1.Credential
	(*Static)(nil),     // 3: controller.storage.credential.store.v1.Static
	(*Dynamic)(nil),    // 4: controller.storage.credential.store.v1.Dynamic
}
var file_controller_storage_credential_store_v1_credential_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_controller_storage_credential_store_v1_credential_proto_init() }
func file_controller_storage_credential_store_v1_credential_proto_init() {
	if File_controller_storage_credential_store_v1_credential_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_storage_credential_store_v1_credential_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store); i {
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
		file_controller_storage_credential_store_v1_credential_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Library); i {
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
		file_controller_storage_credential_store_v1_credential_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credential); i {
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
		file_controller_storage_credential_store_v1_credential_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Static); i {
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
		file_controller_storage_credential_store_v1_credential_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dynamic); i {
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
			RawDescriptor: file_controller_storage_credential_store_v1_credential_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_credential_store_v1_credential_proto_goTypes,
		DependencyIndexes: file_controller_storage_credential_store_v1_credential_proto_depIdxs,
		MessageInfos:      file_controller_storage_credential_store_v1_credential_proto_msgTypes,
	}.Build()
	File_controller_storage_credential_store_v1_credential_proto = out.File
	file_controller_storage_credential_store_v1_credential_proto_rawDesc = nil
	file_controller_storage_credential_store_v1_credential_proto_goTypes = nil
	file_controller_storage_credential_store_v1_credential_proto_depIdxs = nil
}
