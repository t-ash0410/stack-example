// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: ticketmgr/v1/service.proto

package ticketmgrv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateTicketRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RequestedBy   string                 `protobuf:"bytes,1,opt,name=requested_by,json=requestedBy,proto3" json:"requested_by,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Deadline      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=deadline,proto3" json:"deadline,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTicketRequest) Reset() {
	*x = CreateTicketRequest{}
	mi := &file_ticketmgr_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTicketRequest) ProtoMessage() {}

func (x *CreateTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticketmgr_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTicketRequest.ProtoReflect.Descriptor instead.
func (*CreateTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticketmgr_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTicketRequest) GetRequestedBy() string {
	if x != nil {
		return x.RequestedBy
	}
	return ""
}

func (x *CreateTicketRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTicketRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTicketRequest) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

type CreateTicketResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TicketId      string                 `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTicketResponse) Reset() {
	*x = CreateTicketResponse{}
	mi := &file_ticketmgr_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTicketResponse) ProtoMessage() {}

func (x *CreateTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticketmgr_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTicketResponse.ProtoReflect.Descriptor instead.
func (*CreateTicketResponse) Descriptor() ([]byte, []int) {
	return file_ticketmgr_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTicketResponse) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

type UpdateTicketRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TicketId      string                 `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	RequestedBy   string                 `protobuf:"bytes,2,opt,name=requested_by,json=requestedBy,proto3" json:"requested_by,omitempty"`
	Title         *string                `protobuf:"bytes,3,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Description   *string                `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Deadline      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTicketRequest) Reset() {
	*x = UpdateTicketRequest{}
	mi := &file_ticketmgr_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTicketRequest) ProtoMessage() {}

func (x *UpdateTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticketmgr_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTicketRequest.ProtoReflect.Descriptor instead.
func (*UpdateTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticketmgr_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTicketRequest) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

func (x *UpdateTicketRequest) GetRequestedBy() string {
	if x != nil {
		return x.RequestedBy
	}
	return ""
}

func (x *UpdateTicketRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *UpdateTicketRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateTicketRequest) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

type UpdateTicketResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTicketResponse) Reset() {
	*x = UpdateTicketResponse{}
	mi := &file_ticketmgr_v1_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTicketResponse) ProtoMessage() {}

func (x *UpdateTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticketmgr_v1_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTicketResponse.ProtoReflect.Descriptor instead.
func (*UpdateTicketResponse) Descriptor() ([]byte, []int) {
	return file_ticketmgr_v1_service_proto_rawDescGZIP(), []int{3}
}

type DeleteTicketRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TicketId      string                 `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTicketRequest) Reset() {
	*x = DeleteTicketRequest{}
	mi := &file_ticketmgr_v1_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTicketRequest) ProtoMessage() {}

func (x *DeleteTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticketmgr_v1_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTicketRequest.ProtoReflect.Descriptor instead.
func (*DeleteTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticketmgr_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTicketRequest) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

type DeleteTicketResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTicketResponse) Reset() {
	*x = DeleteTicketResponse{}
	mi := &file_ticketmgr_v1_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTicketResponse) ProtoMessage() {}

func (x *DeleteTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticketmgr_v1_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTicketResponse.ProtoReflect.Descriptor instead.
func (*DeleteTicketResponse) Descriptor() ([]byte, []int) {
	return file_ticketmgr_v1_service_proto_rawDescGZIP(), []int{5}
}

var File_ticketmgr_v1_service_proto protoreflect.FileDescriptor

var file_ticketmgr_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36,
	0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x33, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x22, 0xe9, 0x01, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x32, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9d, 0x02, 0x0a, 0x10,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4d, 0x67, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x57, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x21, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x67, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x57, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x21, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x67,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb5, 0x01, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x76, 0x31,
	0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x2d, 0x61,
	0x73, 0x68, 0x30, 0x34, 0x31, 0x30, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2d, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x6d,
	0x67, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x6d, 0x67, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x6d, 0x67, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x67, 0x72, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticketmgr_v1_service_proto_rawDescOnce sync.Once
	file_ticketmgr_v1_service_proto_rawDescData = file_ticketmgr_v1_service_proto_rawDesc
)

func file_ticketmgr_v1_service_proto_rawDescGZIP() []byte {
	file_ticketmgr_v1_service_proto_rawDescOnce.Do(func() {
		file_ticketmgr_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticketmgr_v1_service_proto_rawDescData)
	})
	return file_ticketmgr_v1_service_proto_rawDescData
}

var file_ticketmgr_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ticketmgr_v1_service_proto_goTypes = []any{
	(*CreateTicketRequest)(nil),   // 0: ticketmgr.v1.CreateTicketRequest
	(*CreateTicketResponse)(nil),  // 1: ticketmgr.v1.CreateTicketResponse
	(*UpdateTicketRequest)(nil),   // 2: ticketmgr.v1.UpdateTicketRequest
	(*UpdateTicketResponse)(nil),  // 3: ticketmgr.v1.UpdateTicketResponse
	(*DeleteTicketRequest)(nil),   // 4: ticketmgr.v1.DeleteTicketRequest
	(*DeleteTicketResponse)(nil),  // 5: ticketmgr.v1.DeleteTicketResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_ticketmgr_v1_service_proto_depIdxs = []int32{
	6, // 0: ticketmgr.v1.CreateTicketRequest.deadline:type_name -> google.protobuf.Timestamp
	6, // 1: ticketmgr.v1.UpdateTicketRequest.deadline:type_name -> google.protobuf.Timestamp
	0, // 2: ticketmgr.v1.TicketMgrService.CreateTicket:input_type -> ticketmgr.v1.CreateTicketRequest
	2, // 3: ticketmgr.v1.TicketMgrService.UpdateTicket:input_type -> ticketmgr.v1.UpdateTicketRequest
	4, // 4: ticketmgr.v1.TicketMgrService.DeleteTicket:input_type -> ticketmgr.v1.DeleteTicketRequest
	1, // 5: ticketmgr.v1.TicketMgrService.CreateTicket:output_type -> ticketmgr.v1.CreateTicketResponse
	3, // 6: ticketmgr.v1.TicketMgrService.UpdateTicket:output_type -> ticketmgr.v1.UpdateTicketResponse
	5, // 7: ticketmgr.v1.TicketMgrService.DeleteTicket:output_type -> ticketmgr.v1.DeleteTicketResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ticketmgr_v1_service_proto_init() }
func file_ticketmgr_v1_service_proto_init() {
	if File_ticketmgr_v1_service_proto != nil {
		return
	}
	file_ticketmgr_v1_service_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ticketmgr_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticketmgr_v1_service_proto_goTypes,
		DependencyIndexes: file_ticketmgr_v1_service_proto_depIdxs,
		MessageInfos:      file_ticketmgr_v1_service_proto_msgTypes,
	}.Build()
	File_ticketmgr_v1_service_proto = out.File
	file_ticketmgr_v1_service_proto_rawDesc = nil
	file_ticketmgr_v1_service_proto_goTypes = nil
	file_ticketmgr_v1_service_proto_depIdxs = nil
}