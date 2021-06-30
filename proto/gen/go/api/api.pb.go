// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: api/api.proto

package api

import (
	simple_interface "github.com/fregie/simple/proto/gen/go/simple-interface"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Code int32

const (
	Code_OK            Code = 0 // 请求成功
	Code_InternalError Code = 10005
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:     "OK",
		10005: "InternalError",
	}
	Code_value = map[string]int32{
		"OK":            0,
		"InternalError": 10005,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_api_api_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_api_api_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

type CreateSessionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proto      string                      `protobuf:"bytes,1,opt,name=Proto,proto3" json:"Proto,omitempty"`
	ConfigType simple_interface.ConfigType `protobuf:"varint,2,opt,name=ConfigType,proto3,enum=interface.ConfigType" json:"ConfigType,omitempty"`
	Opt        *simple_interface.Option    `protobuf:"bytes,3,opt,name=Opt,proto3" json:"Opt,omitempty"`
	CustomOpt  string                      `protobuf:"bytes,4,opt,name=customOpt,proto3" json:"customOpt,omitempty"`
}

func (x *CreateSessionReq) Reset() {
	*x = CreateSessionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionReq) ProtoMessage() {}

func (x *CreateSessionReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionReq.ProtoReflect.Descriptor instead.
func (*CreateSessionReq) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSessionReq) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *CreateSessionReq) GetConfigType() simple_interface.ConfigType {
	if x != nil {
		return x.ConfigType
	}
	return simple_interface.ConfigType_JSON
}

func (x *CreateSessionReq) GetOpt() *simple_interface.Option {
	if x != nil {
		return x.Opt
	}
	return nil
}

func (x *CreateSessionReq) GetCustomOpt() string {
	if x != nil {
		return x.CustomOpt
	}
	return ""
}

type CreateSessionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   Code   `protobuf:"varint,1,opt,name=Code,proto3,enum=simple.Code" json:"Code,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Config string `protobuf:"bytes,3,opt,name=Config,proto3" json:"Config,omitempty"`
}

func (x *CreateSessionRsp) Reset() {
	*x = CreateSessionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRsp) ProtoMessage() {}

func (x *CreateSessionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRsp.ProtoReflect.Descriptor instead.
func (*CreateSessionRsp) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSessionRsp) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

func (x *CreateSessionRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CreateSessionRsp) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string                      `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Proto      string                      `protobuf:"bytes,2,opt,name=Proto,proto3" json:"Proto,omitempty"`
	ConfigType simple_interface.ConfigType `protobuf:"varint,3,opt,name=ConfigType,proto3,enum=interface.ConfigType" json:"ConfigType,omitempty"`
	Config     string                      `protobuf:"bytes,4,opt,name=Config,proto3" json:"Config,omitempty"`
	Opt        *simple_interface.Option    `protobuf:"bytes,5,opt,name=Opt,proto3" json:"Opt,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[2]
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
	return file_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *Session) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Session) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *Session) GetConfigType() simple_interface.ConfigType {
	if x != nil {
		return x.ConfigType
	}
	return simple_interface.ConfigType_JSON
}

func (x *Session) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Session) GetOpt() *simple_interface.Option {
	if x != nil {
		return x.Opt
	}
	return nil
}

type GetAllSessionsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllSessionsReq) Reset() {
	*x = GetAllSessionsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSessionsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSessionsReq) ProtoMessage() {}

func (x *GetAllSessionsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSessionsReq.ProtoReflect.Descriptor instead.
func (*GetAllSessionsReq) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{3}
}

type GetAllSessionsRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     Code       `protobuf:"varint,1,opt,name=Code,proto3,enum=simple.Code" json:"Code,omitempty"`
	Msg      string     `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Sessions []*Session `protobuf:"bytes,3,rep,name=Sessions,proto3" json:"Sessions,omitempty"`
}

func (x *GetAllSessionsRsp) Reset() {
	*x = GetAllSessionsRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSessionsRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSessionsRsp) ProtoMessage() {}

func (x *GetAllSessionsRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSessionsRsp.ProtoReflect.Descriptor instead.
func (*GetAllSessionsRsp) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllSessionsRsp) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

func (x *GetAllSessionsRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetAllSessionsRsp) GetSessions() []*Session {
	if x != nil {
		return x.Sessions
	}
	return nil
}

type DeleteSessionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteSessionReq) Reset() {
	*x = DeleteSessionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSessionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSessionReq) ProtoMessage() {}

func (x *DeleteSessionReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSessionReq.ProtoReflect.Descriptor instead.
func (*DeleteSessionReq) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSessionReq) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeleteSessionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Code   `protobuf:"varint,1,opt,name=Code,proto3,enum=simple.Code" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *DeleteSessionRsp) Reset() {
	*x = DeleteSessionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSessionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSessionRsp) ProtoMessage() {}

func (x *DeleteSessionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSessionRsp.ProtoReflect.Descriptor instead.
func (*DeleteSessionRsp) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSessionRsp) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

func (x *DeleteSessionRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetProtosReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProtosReq) Reset() {
	*x = GetProtosReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProtosReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProtosReq) ProtoMessage() {}

func (x *GetProtosReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProtosReq.ProtoReflect.Descriptor instead.
func (*GetProtosReq) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{7}
}

type GetProtosRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   Code     `protobuf:"varint,1,opt,name=Code,proto3,enum=simple.Code" json:"Code,omitempty"`
	Msg    string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Protos []string `protobuf:"bytes,3,rep,name=Protos,proto3" json:"Protos,omitempty"`
}

func (x *GetProtosRsp) Reset() {
	*x = GetProtosRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProtosRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProtosRsp) ProtoMessage() {}

func (x *GetProtosRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProtosRsp.ProtoReflect.Descriptor instead.
func (*GetProtosRsp) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{8}
}

func (x *GetProtosRsp) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

func (x *GetProtosRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetProtosRsp) GetProtos() []string {
	if x != nil {
		return x.Protos
	}
	return nil
}

var File_api_api_proto protoreflect.FileDescriptor

var file_api_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x03, 0x4f, 0x70, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x4f, 0x70, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x22, 0x5e, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70,
	0x12, 0x20, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xa3, 0x01, 0x0a,
	0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35,
	0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a,
	0x03, 0x4f, 0x70, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x4f,
	0x70, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x22, 0x74, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67,
	0x12, 0x2b, 0x0a, 0x08, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x22, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x22, 0x46, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x22, 0x5a, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2a, 0x22, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a,
	0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x95, 0x4e, 0x32, 0xeb, 0x02, 0x0a, 0x09, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x41, 0x50, 0x49, 0x12, 0x5b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x18, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x16, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x19, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x73, 0x70, 0x22, 0x13, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x58, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x14, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x52, 0x73, 0x70, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x72, 0x65, 0x67, 0x69, 0x65, 0x2f, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_api_proto_rawDescOnce sync.Once
	file_api_api_proto_rawDescData = file_api_api_proto_rawDesc
)

func file_api_api_proto_rawDescGZIP() []byte {
	file_api_api_proto_rawDescOnce.Do(func() {
		file_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_proto_rawDescData)
	})
	return file_api_api_proto_rawDescData
}

var file_api_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_api_proto_goTypes = []interface{}{
	(Code)(0),                        // 0: simple.Code
	(*CreateSessionReq)(nil),         // 1: simple.CreateSessionReq
	(*CreateSessionRsp)(nil),         // 2: simple.CreateSessionRsp
	(*Session)(nil),                  // 3: simple.Session
	(*GetAllSessionsReq)(nil),        // 4: simple.GetAllSessionsReq
	(*GetAllSessionsRsp)(nil),        // 5: simple.GetAllSessionsRsp
	(*DeleteSessionReq)(nil),         // 6: simple.DeleteSessionReq
	(*DeleteSessionRsp)(nil),         // 7: simple.DeleteSessionRsp
	(*GetProtosReq)(nil),             // 8: simple.GetProtosReq
	(*GetProtosRsp)(nil),             // 9: simple.GetProtosRsp
	(simple_interface.ConfigType)(0), // 10: interface.ConfigType
	(*simple_interface.Option)(nil),  // 11: interface.Option
}
var file_api_api_proto_depIdxs = []int32{
	10, // 0: simple.CreateSessionReq.ConfigType:type_name -> interface.ConfigType
	11, // 1: simple.CreateSessionReq.Opt:type_name -> interface.Option
	0,  // 2: simple.CreateSessionRsp.Code:type_name -> simple.Code
	10, // 3: simple.Session.ConfigType:type_name -> interface.ConfigType
	11, // 4: simple.Session.Opt:type_name -> interface.Option
	0,  // 5: simple.GetAllSessionsRsp.Code:type_name -> simple.Code
	3,  // 6: simple.GetAllSessionsRsp.Sessions:type_name -> simple.Session
	0,  // 7: simple.DeleteSessionRsp.Code:type_name -> simple.Code
	0,  // 8: simple.GetProtosRsp.Code:type_name -> simple.Code
	1,  // 9: simple.SimpleAPI.CreateSession:input_type -> simple.CreateSessionReq
	4,  // 10: simple.SimpleAPI.GetAllSessions:input_type -> simple.GetAllSessionsReq
	6,  // 11: simple.SimpleAPI.DeleteSession:input_type -> simple.DeleteSessionReq
	8,  // 12: simple.SimpleAPI.GetProtos:input_type -> simple.GetProtosReq
	2,  // 13: simple.SimpleAPI.CreateSession:output_type -> simple.CreateSessionRsp
	5,  // 14: simple.SimpleAPI.GetAllSessions:output_type -> simple.GetAllSessionsRsp
	7,  // 15: simple.SimpleAPI.DeleteSession:output_type -> simple.DeleteSessionRsp
	9,  // 16: simple.SimpleAPI.GetProtos:output_type -> simple.GetProtosRsp
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_api_proto_init() }
func file_api_api_proto_init() {
	if File_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionReq); i {
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
		file_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionRsp); i {
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
		file_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSessionsReq); i {
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
		file_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSessionsRsp); i {
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
		file_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSessionReq); i {
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
		file_api_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSessionRsp); i {
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
		file_api_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProtosReq); i {
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
		file_api_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProtosRsp); i {
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
			RawDescriptor: file_api_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		EnumInfos:         file_api_api_proto_enumTypes,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_rawDesc = nil
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}
