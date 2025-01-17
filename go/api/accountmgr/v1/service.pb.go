// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: accountmgr/v1/service.proto

package accountmgrv1

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

type SlackSSORequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SlackUserId   string                 `protobuf:"bytes,3,opt,name=slack_user_id,json=slackUserId,proto3" json:"slack_user_id,omitempty"`
	SlackTeamId   string                 `protobuf:"bytes,4,opt,name=slack_team_id,json=slackTeamId,proto3" json:"slack_team_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SlackSSORequest) Reset() {
	*x = SlackSSORequest{}
	mi := &file_accountmgr_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SlackSSORequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlackSSORequest) ProtoMessage() {}

func (x *SlackSSORequest) ProtoReflect() protoreflect.Message {
	mi := &file_accountmgr_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlackSSORequest.ProtoReflect.Descriptor instead.
func (*SlackSSORequest) Descriptor() ([]byte, []int) {
	return file_accountmgr_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *SlackSSORequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SlackSSORequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SlackSSORequest) GetSlackUserId() string {
	if x != nil {
		return x.SlackUserId
	}
	return ""
}

func (x *SlackSSORequest) GetSlackTeamId() string {
	if x != nil {
		return x.SlackTeamId
	}
	return ""
}

type SlackSSOResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SlackSSOResponse) Reset() {
	*x = SlackSSOResponse{}
	mi := &file_accountmgr_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SlackSSOResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlackSSOResponse) ProtoMessage() {}

func (x *SlackSSOResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accountmgr_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlackSSOResponse.ProtoReflect.Descriptor instead.
func (*SlackSSOResponse) Descriptor() ([]byte, []int) {
	return file_accountmgr_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *SlackSSOResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_accountmgr_v1_service_proto protoreflect.FileDescriptor

var file_accountmgr_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x83, 0x01, 0x0a,
	0x0f, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x53, 0x53, 0x4f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x6c,
	0x61, 0x63, 0x6b, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0d, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x54, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x22, 0x2b, 0x0a, 0x10, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x53, 0x53, 0x4f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32,
	0x62, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x67, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x08, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x53, 0x53, 0x4f,
	0x12, 0x1e, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x53, 0x53, 0x4f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x53, 0x53, 0x4f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0xbc, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x2d, 0x61, 0x73, 0x68, 0x30, 0x34, 0x31, 0x30, 0x2f,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x6d, 0x67, 0x72, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x6d, 0x67, 0x72, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x6d,
	0x67, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x6d,
	0x67, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x6d,
	0x67, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x6d, 0x67, 0x72, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accountmgr_v1_service_proto_rawDescOnce sync.Once
	file_accountmgr_v1_service_proto_rawDescData = file_accountmgr_v1_service_proto_rawDesc
)

func file_accountmgr_v1_service_proto_rawDescGZIP() []byte {
	file_accountmgr_v1_service_proto_rawDescOnce.Do(func() {
		file_accountmgr_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_accountmgr_v1_service_proto_rawDescData)
	})
	return file_accountmgr_v1_service_proto_rawDescData
}

var file_accountmgr_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_accountmgr_v1_service_proto_goTypes = []any{
	(*SlackSSORequest)(nil),  // 0: accountmgr.v1.SlackSSORequest
	(*SlackSSOResponse)(nil), // 1: accountmgr.v1.SlackSSOResponse
}
var file_accountmgr_v1_service_proto_depIdxs = []int32{
	0, // 0: accountmgr.v1.AccountMgrService.SlackSSO:input_type -> accountmgr.v1.SlackSSORequest
	1, // 1: accountmgr.v1.AccountMgrService.SlackSSO:output_type -> accountmgr.v1.SlackSSOResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_accountmgr_v1_service_proto_init() }
func file_accountmgr_v1_service_proto_init() {
	if File_accountmgr_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_accountmgr_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accountmgr_v1_service_proto_goTypes,
		DependencyIndexes: file_accountmgr_v1_service_proto_depIdxs,
		MessageInfos:      file_accountmgr_v1_service_proto_msgTypes,
	}.Build()
	File_accountmgr_v1_service_proto = out.File
	file_accountmgr_v1_service_proto_rawDesc = nil
	file_accountmgr_v1_service_proto_goTypes = nil
	file_accountmgr_v1_service_proto_depIdxs = nil
}
