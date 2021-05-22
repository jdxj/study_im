// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: chat.proto

package chat

import (
	gate "github.com/jdxj/study_im/proto/gate"
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

type Status int32

const (
	Status_MessageStored Status = 0
	Status_InternalError Status = 1
	Status_NotLoggedIn   Status = 2
	Status_IllegalID     Status = 3
	Status_MsgConfirmed  Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "MessageStored",
		1: "InternalError",
		2: "NotLoggedIn",
		3: "IllegalID",
		4: "MsgConfirmed",
	}
	Status_value = map[string]int32{
		"MessageStored": 0,
		"InternalError": 1,
		"NotLoggedIn":   2,
		"IllegalID":     3,
		"MsgConfirmed":  4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_chat_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_chat_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

// Content 目前使用最简单的字符串, 以后可以扩展
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// Options 表示不向客户端响应
type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

// 6步协议, 只使用4个:
// 1. msg:R
type C2CMsgR struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity *gate.Identity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	From     uint32         `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	To       uint32         `protobuf:"varint,3,opt,name=to,proto3" json:"to,omitempty"`
	Msg      *Message       `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	MsgId    int64          `protobuf:"varint,5,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
}

func (x *C2CMsgR) Reset() {
	*x = C2CMsgR{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2CMsgR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2CMsgR) ProtoMessage() {}

func (x *C2CMsgR) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2CMsgR.ProtoReflect.Descriptor instead.
func (*C2CMsgR) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *C2CMsgR) GetIdentity() *gate.Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *C2CMsgR) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *C2CMsgR) GetTo() uint32 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *C2CMsgR) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *C2CMsgR) GetMsgId() int64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

// 2. msg:A
type C2CMsgA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  Status `protobuf:"varint,1,opt,name=code,proto3,enum=chat.Status" json:"code,omitempty"`
	MsgId int64  `protobuf:"varint,2,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
}

func (x *C2CMsgA) Reset() {
	*x = C2CMsgA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2CMsgA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2CMsgA) ProtoMessage() {}

func (x *C2CMsgA) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2CMsgA.ProtoReflect.Descriptor instead.
func (*C2CMsgA) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *C2CMsgA) GetCode() Status {
	if x != nil {
		return x.Code
	}
	return Status_MessageStored
}

func (x *C2CMsgA) GetMsgId() int64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

// 6. ack:N
type C2CAckN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  uint32 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To    uint32 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	MsgId int64  `protobuf:"varint,3,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
}

func (x *C2CAckN) Reset() {
	*x = C2CAckN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2CAckN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2CAckN) ProtoMessage() {}

func (x *C2CAckN) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2CAckN.ProtoReflect.Descriptor instead.
func (*C2CAckN) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *C2CAckN) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *C2CAckN) GetTo() uint32 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *C2CAckN) GetMsgId() int64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

// 3. msg:N
type C2CMsgN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  uint32   `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	Msg   *Message `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	MsgId int64    `protobuf:"varint,3,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
}

func (x *C2CMsgN) Reset() {
	*x = C2CMsgN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2CMsgN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2CMsgN) ProtoMessage() {}

func (x *C2CMsgN) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2CMsgN.ProtoReflect.Descriptor instead.
func (*C2CMsgN) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{5}
}

func (x *C2CMsgN) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *C2CMsgN) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *C2CMsgN) GetMsgId() int64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

// 4. ack:R
type C2CAckR struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity *gate.Identity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	From     uint32         `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	To       uint32         `protobuf:"varint,3,opt,name=to,proto3" json:"to,omitempty"`
	MsgId    int64          `protobuf:"varint,4,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
}

func (x *C2CAckR) Reset() {
	*x = C2CAckR{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2CAckR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2CAckR) ProtoMessage() {}

func (x *C2CAckR) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2CAckR.ProtoReflect.Descriptor instead.
func (*C2CAckR) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{6}
}

func (x *C2CAckR) GetIdentity() *gate.Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *C2CAckR) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *C2CAckR) GetTo() uint32 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *C2CAckR) GetMsgId() int64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

// 5. ack:A
type C2CAckA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  Status `protobuf:"varint,1,opt,name=code,proto3,enum=chat.Status" json:"code,omitempty"`
	MsgId int64  `protobuf:"varint,2,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
}

func (x *C2CAckA) Reset() {
	*x = C2CAckA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2CAckA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2CAckA) ProtoMessage() {}

func (x *C2CAckA) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2CAckA.ProtoReflect.Descriptor instead.
func (*C2CAckA) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{7}
}

func (x *C2CAckA) GetCode() Status {
	if x != nil {
		return x.Code
	}
	return Status_MessageStored
}

func (x *C2CAckA) GetMsgId() int64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

type C2GSendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  uint32   `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	Group uint32   `protobuf:"varint,2,opt,name=group,proto3" json:"group,omitempty"`
	Msg   *Message `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *C2GSendRequest) Reset() {
	*x = C2GSendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2GSendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2GSendRequest) ProtoMessage() {}

func (x *C2GSendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2GSendRequest.ProtoReflect.Descriptor instead.
func (*C2GSendRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{8}
}

func (x *C2GSendRequest) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *C2GSendRequest) GetGroup() uint32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *C2GSendRequest) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

type C2GSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId uint64 `protobuf:"varint,1,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
}

func (x *C2GSendResponse) Reset() {
	*x = C2GSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2GSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2GSendResponse) ProtoMessage() {}

func (x *C2GSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2GSendResponse.ProtoReflect.Descriptor instead.
func (*C2GSendResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{9}
}

func (x *C2GSendResponse) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

type C2GPushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  uint32   `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	Group uint32   `protobuf:"varint,2,opt,name=group,proto3" json:"group,omitempty"`
	Msg   *Message `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	MsgId uint64   `protobuf:"varint,4,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
}

func (x *C2GPushRequest) Reset() {
	*x = C2GPushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2GPushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2GPushRequest) ProtoMessage() {}

func (x *C2GPushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2GPushRequest.ProtoReflect.Descriptor instead.
func (*C2GPushRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{10}
}

func (x *C2GPushRequest) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *C2GPushRequest) GetGroup() uint32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *C2GPushRequest) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *C2GPushRequest) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

type C2GPushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId uint64 `protobuf:"varint,1,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
}

func (x *C2GPushResponse) Reset() {
	*x = C2GPushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2GPushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2GPushResponse) ProtoMessage() {}

func (x *C2GPushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2GPushResponse.ProtoReflect.Descriptor instead.
func (*C2GPushResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{11}
}

func (x *C2GPushResponse) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

type C2SPullMsgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   uint32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	MsgId uint64 `protobuf:"varint,2,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	Limit uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *C2SPullMsgRequest) Reset() {
	*x = C2SPullMsgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SPullMsgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SPullMsgRequest) ProtoMessage() {}

func (x *C2SPullMsgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SPullMsgRequest.ProtoReflect.Descriptor instead.
func (*C2SPullMsgRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{12}
}

func (x *C2SPullMsgRequest) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *C2SPullMsgRequest) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *C2SPullMsgRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type PullMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From     uint32   `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	Group    uint32   `protobuf:"varint,2,opt,name=group,proto3" json:"group,omitempty"`
	Msg      *Message `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	MsgId    uint64   `protobuf:"varint,4,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	SendTime int64    `protobuf:"varint,5,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
}

func (x *PullMsg) Reset() {
	*x = PullMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullMsg) ProtoMessage() {}

func (x *PullMsg) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullMsg.ProtoReflect.Descriptor instead.
func (*PullMsg) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{13}
}

func (x *PullMsg) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *PullMsg) GetGroup() uint32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *PullMsg) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *PullMsg) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *PullMsg) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

type C2SPullMsgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg []*PullMsg `protobuf:"bytes,1,rep,name=msg,proto3" json:"msg,omitempty"`
}

func (x *C2SPullMsgResponse) Reset() {
	*x = C2SPullMsgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SPullMsgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SPullMsgResponse) ProtoMessage() {}

func (x *C2SPullMsgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SPullMsgResponse.ProtoReflect.Descriptor instead.
func (*C2SPullMsgResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{14}
}

func (x *C2SPullMsgResponse) GetMsg() []*PullMsg {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x1a, 0x0f, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x22, 0x09, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x91, 0x01,
	0x0a, 0x07, 0x43, 0x32, 0x43, 0x4d, 0x73, 0x67, 0x52, 0x12, 0x2a, 0x0a, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1f, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x64, 0x22, 0x42, 0x0a, 0x07, 0x43, 0x32, 0x43, 0x4d, 0x73, 0x67, 0x41, 0x12, 0x20, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x15,
	0x0a, 0x06, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6d, 0x73, 0x67, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x07, 0x43, 0x32, 0x43, 0x41, 0x63, 0x6b, 0x4e,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x07, 0x43,
	0x32, 0x43, 0x4d, 0x73, 0x67, 0x4e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1f, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x6d,
	0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x73, 0x67,
	0x49, 0x64, 0x22, 0x70, 0x0a, 0x07, 0x43, 0x32, 0x43, 0x41, 0x63, 0x6b, 0x52, 0x12, 0x2a, 0x0a,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x15, 0x0a,
	0x06, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d,
	0x73, 0x67, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x07, 0x43, 0x32, 0x43, 0x41, 0x63, 0x6b, 0x41, 0x12,
	0x20, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x0e, 0x43, 0x32, 0x47, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x1f, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x28, 0x0a, 0x0f, 0x43, 0x32, 0x47, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x22,
	0x72, 0x0a, 0x0e, 0x43, 0x32, 0x47, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1f, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x15, 0x0a, 0x06,
	0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x73,
	0x67, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0f, 0x43, 0x32, 0x47, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x22, 0x52, 0x0a,
	0x11, 0x43, 0x32, 0x53, 0x50, 0x75, 0x6c, 0x6c, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x88, 0x01, 0x0a, 0x07, 0x50, 0x75, 0x6c, 0x6c, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1f, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x12,
	0x43, 0x32, 0x53, 0x50, 0x75, 0x6c, 0x6c, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x4d, 0x73, 0x67, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x2a, 0x60, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a,
	0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64,
	0x49, 0x6e, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x6c, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x49,
	0x44, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x65, 0x64, 0x10, 0x04, 0x32, 0x59, 0x0a, 0x03, 0x43, 0x32, 0x43, 0x12, 0x28, 0x0a, 0x06,
	0x43, 0x32, 0x43, 0x4d, 0x73, 0x67, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x32,
	0x43, 0x4d, 0x73, 0x67, 0x52, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x32, 0x43,
	0x4d, 0x73, 0x67, 0x41, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x06, 0x43, 0x32, 0x43, 0x41, 0x63, 0x6b,
	0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x32, 0x43, 0x41, 0x63, 0x6b, 0x52, 0x1a,
	0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_chat_proto_goTypes = []interface{}{
	(Status)(0),                // 0: chat.Status
	(*Message)(nil),            // 1: chat.Message
	(*Options)(nil),            // 2: chat.Options
	(*C2CMsgR)(nil),            // 3: chat.C2CMsgR
	(*C2CMsgA)(nil),            // 4: chat.C2CMsgA
	(*C2CAckN)(nil),            // 5: chat.C2CAckN
	(*C2CMsgN)(nil),            // 6: chat.C2CMsgN
	(*C2CAckR)(nil),            // 7: chat.C2CAckR
	(*C2CAckA)(nil),            // 8: chat.C2CAckA
	(*C2GSendRequest)(nil),     // 9: chat.C2GSendRequest
	(*C2GSendResponse)(nil),    // 10: chat.C2GSendResponse
	(*C2GPushRequest)(nil),     // 11: chat.C2GPushRequest
	(*C2GPushResponse)(nil),    // 12: chat.C2GPushResponse
	(*C2SPullMsgRequest)(nil),  // 13: chat.C2SPullMsgRequest
	(*PullMsg)(nil),            // 14: chat.PullMsg
	(*C2SPullMsgResponse)(nil), // 15: chat.C2SPullMsgResponse
	(*gate.Identity)(nil),      // 16: gate.Identity
}
var file_chat_proto_depIdxs = []int32{
	16, // 0: chat.C2CMsgR.identity:type_name -> gate.Identity
	1,  // 1: chat.C2CMsgR.msg:type_name -> chat.Message
	0,  // 2: chat.C2CMsgA.code:type_name -> chat.Status
	1,  // 3: chat.C2CMsgN.msg:type_name -> chat.Message
	16, // 4: chat.C2CAckR.identity:type_name -> gate.Identity
	0,  // 5: chat.C2CAckA.code:type_name -> chat.Status
	1,  // 6: chat.C2GSendRequest.msg:type_name -> chat.Message
	1,  // 7: chat.C2GPushRequest.msg:type_name -> chat.Message
	1,  // 8: chat.PullMsg.msg:type_name -> chat.Message
	14, // 9: chat.C2SPullMsgResponse.msg:type_name -> chat.PullMsg
	3,  // 10: chat.C2C.C2CMsg:input_type -> chat.C2CMsgR
	7,  // 11: chat.C2C.C2CAck:input_type -> chat.C2CAckR
	4,  // 12: chat.C2C.C2CMsg:output_type -> chat.C2CMsgA
	2,  // 13: chat.C2C.C2CAck:output_type -> chat.Options
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2CMsgR); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2CMsgA); i {
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
		file_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2CAckN); i {
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
		file_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2CMsgN); i {
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
		file_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2CAckR); i {
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
		file_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2CAckA); i {
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
		file_chat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2GSendRequest); i {
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
		file_chat_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2GSendResponse); i {
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
		file_chat_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2GPushRequest); i {
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
		file_chat_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2GPushResponse); i {
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
		file_chat_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2SPullMsgRequest); i {
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
		file_chat_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullMsg); i {
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
		file_chat_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2SPullMsgResponse); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		EnumInfos:         file_chat_proto_enumTypes,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
