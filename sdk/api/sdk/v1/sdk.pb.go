// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: api/sdk/v1/sdk.proto

package v1

import (
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

type CommonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	State string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *CommonReply) Reset() {
	*x = CommonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sdk_v1_sdk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonReply) ProtoMessage() {}

func (x *CommonReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sdk_v1_sdk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonReply.ProtoReflect.Descriptor instead.
func (*CommonReply) Descriptor() ([]byte, []int) {
	return file_api_sdk_v1_sdk_proto_rawDescGZIP(), []int{0}
}

func (x *CommonReply) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *CommonReply) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type InitSdkReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string          `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	AppId   uint32          `protobuf:"varint,2,opt,name=appId,proto3" json:"appId,omitempty"`
	Data    *InitSdkRegData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Sign    string          `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *InitSdkReq) Reset() {
	*x = InitSdkReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sdk_v1_sdk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitSdkReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitSdkReq) ProtoMessage() {}

func (x *InitSdkReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_sdk_v1_sdk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitSdkReq.ProtoReflect.Descriptor instead.
func (*InitSdkReq) Descriptor() ([]byte, []int) {
	return file_api_sdk_v1_sdk_proto_rawDescGZIP(), []int{1}
}

func (x *InitSdkReq) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *InitSdkReq) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *InitSdkReq) GetData() *InitSdkRegData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *InitSdkReq) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

type InitSdkRegData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Udid    string `protobuf:"bytes,1,opt,name=udid,proto3" json:"udid,omitempty"`
	Channel uint32 `protobuf:"varint,2,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *InitSdkRegData) Reset() {
	*x = InitSdkRegData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sdk_v1_sdk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitSdkRegData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitSdkRegData) ProtoMessage() {}

func (x *InitSdkRegData) ProtoReflect() protoreflect.Message {
	mi := &file_api_sdk_v1_sdk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitSdkRegData.ProtoReflect.Descriptor instead.
func (*InitSdkRegData) Descriptor() ([]byte, []int) {
	return file_api_sdk_v1_sdk_proto_rawDescGZIP(), []int{2}
}

func (x *InitSdkRegData) GetUdid() string {
	if x != nil {
		return x.Udid
	}
	return ""
}

func (x *InitSdkRegData) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

type InitSdkReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"` //  string gameName = 3; // 暂不返回，因为要单独拿数据
}

func (x *InitSdkReply) Reset() {
	*x = InitSdkReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sdk_v1_sdk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitSdkReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitSdkReply) ProtoMessage() {}

func (x *InitSdkReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sdk_v1_sdk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitSdkReply.ProtoReflect.Descriptor instead.
func (*InitSdkReply) Descriptor() ([]byte, []int) {
	return file_api_sdk_v1_sdk_proto_rawDescGZIP(), []int{3}
}

func (x *InitSdkReply) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *InitSdkReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type RegReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string      `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	AppId   uint32      `protobuf:"varint,2,opt,name=appId,proto3" json:"appId,omitempty"`
	Data    *RegReqData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Sign    string      `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *RegReq) Reset() {
	*x = RegReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sdk_v1_sdk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegReq) ProtoMessage() {}

func (x *RegReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_sdk_v1_sdk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegReq.ProtoReflect.Descriptor instead.
func (*RegReq) Descriptor() ([]byte, []int) {
	return file_api_sdk_v1_sdk_proto_rawDescGZIP(), []int{4}
}

func (x *RegReq) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *RegReq) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *RegReq) GetData() *RegReqData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RegReq) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

type RegReqData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Passwd   string `protobuf:"bytes,2,opt,name=passwd,proto3" json:"passwd,omitempty"`
	Udid     string `protobuf:"bytes,3,opt,name=udid,proto3" json:"udid,omitempty"`
	Channel  uint32 `protobuf:"varint,4,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *RegReqData) Reset() {
	*x = RegReqData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sdk_v1_sdk_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegReqData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegReqData) ProtoMessage() {}

func (x *RegReqData) ProtoReflect() protoreflect.Message {
	mi := &file_api_sdk_v1_sdk_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegReqData.ProtoReflect.Descriptor instead.
func (*RegReqData) Descriptor() ([]byte, []int) {
	return file_api_sdk_v1_sdk_proto_rawDescGZIP(), []int{5}
}

func (x *RegReqData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegReqData) GetPasswd() string {
	if x != nil {
		return x.Passwd
	}
	return ""
}

func (x *RegReqData) GetUdid() string {
	if x != nil {
		return x.Udid
	}
	return ""
}

func (x *RegReqData) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

type RegReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *RegReplyData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RegReply) Reset() {
	*x = RegReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sdk_v1_sdk_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegReply) ProtoMessage() {}

func (x *RegReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sdk_v1_sdk_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegReply.ProtoReflect.Descriptor instead.
func (*RegReply) Descriptor() ([]byte, []int) {
	return file_api_sdk_v1_sdk_proto_rawDescGZIP(), []int{6}
}

func (x *RegReply) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RegReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RegReply) GetData() *RegReplyData {
	if x != nil {
		return x.Data
	}
	return nil
}

type RegReplyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *RegReplyData) Reset() {
	*x = RegReplyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sdk_v1_sdk_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegReplyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegReplyData) ProtoMessage() {}

func (x *RegReplyData) ProtoReflect() protoreflect.Message {
	mi := &file_api_sdk_v1_sdk_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegReplyData.ProtoReflect.Descriptor instead.
func (*RegReplyData) Descriptor() ([]byte, []int) {
	return file_api_sdk_v1_sdk_proto_rawDescGZIP(), []int{7}
}

func (x *RegReplyData) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string     `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	AppId   uint32     `protobuf:"varint,2,opt,name=appId,proto3" json:"appId,omitempty"`
	Data    *LoginData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Sign    string     `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sdk_v1_sdk_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_sdk_v1_sdk_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_api_sdk_v1_sdk_proto_rawDescGZIP(), []int{8}
}

func (x *LoginReq) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *LoginReq) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *LoginReq) GetData() *LoginData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LoginReq) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

type LoginData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Udid    string `protobuf:"bytes,1,opt,name=udid,proto3" json:"udid,omitempty"`
	Channel uint32 `protobuf:"varint,2,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *LoginData) Reset() {
	*x = LoginData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sdk_v1_sdk_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginData) ProtoMessage() {}

func (x *LoginData) ProtoReflect() protoreflect.Message {
	mi := &file_api_sdk_v1_sdk_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginData.ProtoReflect.Descriptor instead.
func (*LoginData) Descriptor() ([]byte, []int) {
	return file_api_sdk_v1_sdk_proto_rawDescGZIP(), []int{9}
}

func (x *LoginData) GetUdid() string {
	if x != nil {
		return x.Udid
	}
	return ""
}

func (x *LoginData) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sdk_v1_sdk_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sdk_v1_sdk_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_api_sdk_v1_sdk_proto_rawDescGZIP(), []int{10}
}

func (x *LoginReply) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LoginReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type CheckEnterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string     `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	AppId   uint32     `protobuf:"varint,2,opt,name=appId,proto3" json:"appId,omitempty"`
	Data    *LoginData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Sign    string     `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *CheckEnterReq) Reset() {
	*x = CheckEnterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sdk_v1_sdk_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckEnterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEnterReq) ProtoMessage() {}

func (x *CheckEnterReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_sdk_v1_sdk_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEnterReq.ProtoReflect.Descriptor instead.
func (*CheckEnterReq) Descriptor() ([]byte, []int) {
	return file_api_sdk_v1_sdk_proto_rawDescGZIP(), []int{11}
}

func (x *CheckEnterReq) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *CheckEnterReq) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *CheckEnterReq) GetData() *LoginData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CheckEnterReq) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

type CheckEnterData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid      string `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *CheckEnterData) Reset() {
	*x = CheckEnterData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sdk_v1_sdk_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckEnterData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEnterData) ProtoMessage() {}

func (x *CheckEnterData) ProtoReflect() protoreflect.Message {
	mi := &file_api_sdk_v1_sdk_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEnterData.ProtoReflect.Descriptor instead.
func (*CheckEnterData) Descriptor() ([]byte, []int) {
	return file_api_sdk_v1_sdk_proto_rawDescGZIP(), []int{12}
}

func (x *CheckEnterData) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *CheckEnterData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CheckEnterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CheckEnterReply) Reset() {
	*x = CheckEnterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sdk_v1_sdk_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckEnterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEnterReply) ProtoMessage() {}

func (x *CheckEnterReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sdk_v1_sdk_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEnterReply.ProtoReflect.Descriptor instead.
func (*CheckEnterReply) Descriptor() ([]byte, []int) {
	return file_api_sdk_v1_sdk_proto_rawDescGZIP(), []int{13}
}

func (x *CheckEnterReply) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CheckEnterReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_api_sdk_v1_sdk_proto protoreflect.FileDescriptor

var file_api_sdk_v1_sdk_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x64, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0b,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x7c, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x64, 0x6b,
	0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x53,
	0x64, 0x6b, 0x52, 0x65, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x69, 0x67, 0x6e, 0x22, 0x3e, 0x0a, 0x0e, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x64, 0x6b, 0x52, 0x65,
	0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x64, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x22, 0x34, 0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x64, 0x6b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x74, 0x0a, 0x06, 0x52, 0x65, 0x67,
	0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x52, 0x65,
	0x71, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x67, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22,
	0x6e, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x64, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x64, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22,
	0x5a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x20, 0x0a, 0x0c, 0x52,
	0x65, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x75, 0x0a,
	0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x69, 0x67, 0x6e, 0x22, 0x39, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x64, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22,
	0x32, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x7a, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x67, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22,
	0x3e, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x37, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xcb, 0x02, 0x0a, 0x03, 0x53, 0x64, 0x6b,
	0x12, 0x54, 0x0a, 0x07, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x64, 0x6b, 0x12, 0x12, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x64, 0x6b, 0x52, 0x65, 0x71, 0x1a,
	0x14, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x64, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f,
	0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x69, 0x74, 0x5f,
	0x73, 0x64, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0x43, 0x0a, 0x03, 0x52, 0x65, 0x67, 0x12, 0x0e, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x3a, 0x01, 0x2a, 0x12, 0x4b, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x22, 0x11, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x5c, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x42, 0x34, 0x0a, 0x15, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x42,
	0x0a, 0x53, 0x64, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x0d, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_sdk_v1_sdk_proto_rawDescOnce sync.Once
	file_api_sdk_v1_sdk_proto_rawDescData = file_api_sdk_v1_sdk_proto_rawDesc
)

func file_api_sdk_v1_sdk_proto_rawDescGZIP() []byte {
	file_api_sdk_v1_sdk_proto_rawDescOnce.Do(func() {
		file_api_sdk_v1_sdk_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_sdk_v1_sdk_proto_rawDescData)
	})
	return file_api_sdk_v1_sdk_proto_rawDescData
}

var file_api_sdk_v1_sdk_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_sdk_v1_sdk_proto_goTypes = []interface{}{
	(*CommonReply)(nil),     // 0: sdk.v1.CommonReply
	(*InitSdkReq)(nil),      // 1: sdk.v1.InitSdkReq
	(*InitSdkRegData)(nil),  // 2: sdk.v1.InitSdkRegData
	(*InitSdkReply)(nil),    // 3: sdk.v1.InitSdkReply
	(*RegReq)(nil),          // 4: sdk.v1.RegReq
	(*RegReqData)(nil),      // 5: sdk.v1.RegReqData
	(*RegReply)(nil),        // 6: sdk.v1.RegReply
	(*RegReplyData)(nil),    // 7: sdk.v1.RegReplyData
	(*LoginReq)(nil),        // 8: sdk.v1.LoginReq
	(*LoginData)(nil),       // 9: sdk.v1.LoginData
	(*LoginReply)(nil),      // 10: sdk.v1.LoginReply
	(*CheckEnterReq)(nil),   // 11: sdk.v1.CheckEnterReq
	(*CheckEnterData)(nil),  // 12: sdk.v1.CheckEnterData
	(*CheckEnterReply)(nil), // 13: sdk.v1.CheckEnterReply
}
var file_api_sdk_v1_sdk_proto_depIdxs = []int32{
	2,  // 0: sdk.v1.InitSdkReq.data:type_name -> sdk.v1.InitSdkRegData
	5,  // 1: sdk.v1.RegReq.data:type_name -> sdk.v1.RegReqData
	7,  // 2: sdk.v1.RegReply.data:type_name -> sdk.v1.RegReplyData
	9,  // 3: sdk.v1.LoginReq.data:type_name -> sdk.v1.LoginData
	9,  // 4: sdk.v1.CheckEnterReq.data:type_name -> sdk.v1.LoginData
	1,  // 5: sdk.v1.Sdk.InitSdk:input_type -> sdk.v1.InitSdkReq
	4,  // 6: sdk.v1.Sdk.Reg:input_type -> sdk.v1.RegReq
	8,  // 7: sdk.v1.Sdk.Login:input_type -> sdk.v1.LoginReq
	11, // 8: sdk.v1.Sdk.CheckEnter:input_type -> sdk.v1.CheckEnterReq
	3,  // 9: sdk.v1.Sdk.InitSdk:output_type -> sdk.v1.InitSdkReply
	6,  // 10: sdk.v1.Sdk.Reg:output_type -> sdk.v1.RegReply
	10, // 11: sdk.v1.Sdk.Login:output_type -> sdk.v1.LoginReply
	0,  // 12: sdk.v1.Sdk.CheckEnter:output_type -> sdk.v1.CommonReply
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_sdk_v1_sdk_proto_init() }
func file_api_sdk_v1_sdk_proto_init() {
	if File_api_sdk_v1_sdk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_sdk_v1_sdk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonReply); i {
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
		file_api_sdk_v1_sdk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitSdkReq); i {
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
		file_api_sdk_v1_sdk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitSdkRegData); i {
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
		file_api_sdk_v1_sdk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitSdkReply); i {
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
		file_api_sdk_v1_sdk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegReq); i {
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
		file_api_sdk_v1_sdk_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegReqData); i {
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
		file_api_sdk_v1_sdk_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegReply); i {
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
		file_api_sdk_v1_sdk_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegReplyData); i {
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
		file_api_sdk_v1_sdk_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_api_sdk_v1_sdk_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginData); i {
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
		file_api_sdk_v1_sdk_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_api_sdk_v1_sdk_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckEnterReq); i {
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
		file_api_sdk_v1_sdk_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckEnterData); i {
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
		file_api_sdk_v1_sdk_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckEnterReply); i {
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
			RawDescriptor: file_api_sdk_v1_sdk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_sdk_v1_sdk_proto_goTypes,
		DependencyIndexes: file_api_sdk_v1_sdk_proto_depIdxs,
		MessageInfos:      file_api_sdk_v1_sdk_proto_msgTypes,
	}.Build()
	File_api_sdk_v1_sdk_proto = out.File
	file_api_sdk_v1_sdk_proto_rawDesc = nil
	file_api_sdk_v1_sdk_proto_goTypes = nil
	file_api_sdk_v1_sdk_proto_depIdxs = nil
}
