// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: testing/event/v1/event.proto

package event

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestAuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Auth Method in the system that should be used for authentication.
	// @inject_tag: `class:"public"`
	AuthMethodId string `protobuf:"bytes,1,opt,name=auth_method_id,proto3" json:"auth_method_id,omitempty" class:"public"`
	// This can be "cookie" or "token". If not provided, "token" will be used. "cookie" activates a split-cookie method where the token is split partially between http-only and regular cookies in order
	// to keep it safe from rogue JS in the browser. Deprecated, use "type" instead.
	// @inject_tag: `class:"public"`
	//
	// Deprecated: Do not use.
	TokenType string `protobuf:"bytes,2,opt,name=token_type,proto3" json:"token_type,omitempty" class:"public"`
	// This can be "cookie" or "token". If not provided, "token" will be used. "cookie" activates a split-cookie method where the token is split partially between http-only and regular cookies in order
	// to keep it safe from rogue JS in the browser.
	// @inject_tag: `class:"public"`
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" class:"public"`
	// Attributes are passed to the Auth Method; the valid keys and values depend on the type of Auth Method as well as the command.
	Attributes *structpb.Struct `protobuf:"bytes,4,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// The command to perform.
	// @inject_tag: `class:"public"`
	Command string `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty" class:"public"`
}

func (x *TestAuthenticateRequest) Reset() {
	*x = TestAuthenticateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testing_event_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestAuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestAuthenticateRequest) ProtoMessage() {}

func (x *TestAuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testing_event_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestAuthenticateRequest.ProtoReflect.Descriptor instead.
func (*TestAuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_testing_event_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *TestAuthenticateRequest) GetAuthMethodId() string {
	if x != nil {
		return x.AuthMethodId
	}
	return ""
}

// Deprecated: Do not use.
func (x *TestAuthenticateRequest) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *TestAuthenticateRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TestAuthenticateRequest) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *TestAuthenticateRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type TestAuthenticateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Valid keys and values depend on the type of Auth Method as well as the command.
	Attributes *structpb.Struct `protobuf:"bytes,4,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// The command that was performed.
	// @inject_tag: `class:"public"`
	Command string `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty" class:"public"`
}

func (x *TestAuthenticateResponse) Reset() {
	*x = TestAuthenticateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testing_event_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestAuthenticateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestAuthenticateResponse) ProtoMessage() {}

func (x *TestAuthenticateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testing_event_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestAuthenticateResponse.ProtoReflect.Descriptor instead.
func (*TestAuthenticateResponse) Descriptor() ([]byte, []int) {
	return file_testing_event_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *TestAuthenticateResponse) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *TestAuthenticateResponse) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

var File_testing_event_v1_event_proto protoreflect.FileDescriptor

var file_testing_event_v1_event_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a,
	0x17, 0x54, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64,
	0x12, 0x22, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x18,
	0x54, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4a, 0x04, 0x08, 0x01, 0x10,
	0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x32, 0x89, 0x02, 0x0a, 0x15, 0x54, 0x65,
	0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xef, 0x01, 0x0a, 0x10, 0x54, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x83, 0x01, 0x92, 0x41, 0x47, 0x12, 0x45, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x6e,
	0x20, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x33, 0x22, 0x2e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x2f, 0x7b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testing_event_v1_event_proto_rawDescOnce sync.Once
	file_testing_event_v1_event_proto_rawDescData = file_testing_event_v1_event_proto_rawDesc
)

func file_testing_event_v1_event_proto_rawDescGZIP() []byte {
	file_testing_event_v1_event_proto_rawDescOnce.Do(func() {
		file_testing_event_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_testing_event_v1_event_proto_rawDescData)
	})
	return file_testing_event_v1_event_proto_rawDescData
}

var file_testing_event_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_testing_event_v1_event_proto_goTypes = []interface{}{
	(*TestAuthenticateRequest)(nil),  // 0: testing.event.v1.TestAuthenticateRequest
	(*TestAuthenticateResponse)(nil), // 1: testing.event.v1.TestAuthenticateResponse
	(*structpb.Struct)(nil),          // 2: google.protobuf.Struct
}
var file_testing_event_v1_event_proto_depIdxs = []int32{
	2, // 0: testing.event.v1.TestAuthenticateRequest.attributes:type_name -> google.protobuf.Struct
	2, // 1: testing.event.v1.TestAuthenticateResponse.attributes:type_name -> google.protobuf.Struct
	0, // 2: testing.event.v1.TestAuthMethodService.TestAuthenticate:input_type -> testing.event.v1.TestAuthenticateRequest
	1, // 3: testing.event.v1.TestAuthMethodService.TestAuthenticate:output_type -> testing.event.v1.TestAuthenticateResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_testing_event_v1_event_proto_init() }
func file_testing_event_v1_event_proto_init() {
	if File_testing_event_v1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testing_event_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestAuthenticateRequest); i {
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
		file_testing_event_v1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestAuthenticateResponse); i {
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
			RawDescriptor: file_testing_event_v1_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_testing_event_v1_event_proto_goTypes,
		DependencyIndexes: file_testing_event_v1_event_proto_depIdxs,
		MessageInfos:      file_testing_event_v1_event_proto_msgTypes,
	}.Build()
	File_testing_event_v1_event_proto = out.File
	file_testing_event_v1_event_proto_rawDesc = nil
	file_testing_event_v1_event_proto_goTypes = nil
	file_testing_event_v1_event_proto_depIdxs = nil
}
