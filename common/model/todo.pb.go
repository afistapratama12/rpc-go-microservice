// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: model/todo.proto

package model

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Category    string `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	UserId      int32  `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_model_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_model_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Todo) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Todo) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type TodoInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Category    string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *TodoInput) Reset() {
	*x = TodoInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoInput) ProtoMessage() {}

func (x *TodoInput) ProtoReflect() protoreflect.Message {
	mi := &file_model_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoInput.ProtoReflect.Descriptor instead.
func (*TodoInput) Descriptor() ([]byte, []int) {
	return file_model_todo_proto_rawDescGZIP(), []int{1}
}

func (x *TodoInput) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TodoInput) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TodoInput) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type TodoCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string     `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Todo   *TodoInput `protobuf:"bytes,2,opt,name=Todo,proto3" json:"Todo,omitempty"`
}

func (x *TodoCreate) Reset() {
	*x = TodoCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoCreate) ProtoMessage() {}

func (x *TodoCreate) ProtoReflect() protoreflect.Message {
	mi := &file_model_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoCreate.ProtoReflect.Descriptor instead.
func (*TodoCreate) Descriptor() ([]byte, []int) {
	return file_model_todo_proto_rawDescGZIP(), []int{2}
}

func (x *TodoCreate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TodoCreate) GetTodo() *TodoInput {
	if x != nil {
		return x.Todo
	}
	return nil
}

type TodoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Todo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *TodoList) Reset() {
	*x = TodoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoList) ProtoMessage() {}

func (x *TodoList) ProtoReflect() protoreflect.Message {
	mi := &file_model_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoList.ProtoReflect.Descriptor instead.
func (*TodoList) Descriptor() ([]byte, []int) {
	return file_model_todo_proto_rawDescGZIP(), []int{3}
}

func (x *TodoList) GetList() []*Todo {
	if x != nil {
		return x.List
	}
	return nil
}

type Deleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Deleted) Reset() {
	*x = Deleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deleted) ProtoMessage() {}

func (x *Deleted) ProtoReflect() protoreflect.Message {
	mi := &file_model_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deleted.ProtoReflect.Descriptor instead.
func (*Deleted) Descriptor() ([]byte, []int) {
	return file_model_todo_proto_rawDescGZIP(), []int{4}
}

func (x *Deleted) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TodoId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TodoId) Reset() {
	*x = TodoId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoId) ProtoMessage() {}

func (x *TodoId) ProtoReflect() protoreflect.Message {
	mi := &file_model_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoId.ProtoReflect.Descriptor instead.
func (*TodoId) Descriptor() ([]byte, []int) {
	return file_model_todo_proto_rawDescGZIP(), []int{5}
}

func (x *TodoId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TodoUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Todo *TodoInput `protobuf:"bytes,2,opt,name=Todo,proto3" json:"Todo,omitempty"`
}

func (x *TodoUpdate) Reset() {
	*x = TodoUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoUpdate) ProtoMessage() {}

func (x *TodoUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_model_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoUpdate.ProtoReflect.Descriptor instead.
func (*TodoUpdate) Descriptor() ([]byte, []int) {
	return file_model_todo_proto_rawDescGZIP(), []int{6}
}

func (x *TodoUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TodoUpdate) GetTodo() *TodoInput {
	if x != nil {
		return x.Todo
	}
	return nil
}

type TodoUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *TodoUserId) Reset() {
	*x = TodoUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_todo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoUserId) ProtoMessage() {}

func (x *TodoUserId) ProtoReflect() protoreflect.Message {
	mi := &file_model_todo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoUserId.ProtoReflect.Descriptor instead.
func (*TodoUserId) Descriptor() ([]byte, []int) {
	return file_model_todo_proto_rawDescGZIP(), []int{7}
}

func (x *TodoUserId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_model_todo_proto protoreflect.FileDescriptor

var file_model_todo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x83, 0x01, 0x0a, 0x04, 0x54, 0x6f,
	0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x5f, 0x0a, 0x09, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x22, 0x4b, 0x0a, 0x0a, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x22, 0x2b, 0x0a,
	0x08, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x23, 0x0a, 0x07, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x18, 0x0a, 0x06, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x0a, 0x54, 0x6f, 0x64,
	0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x22, 0x25, 0x0a,
	0x0a, 0x54, 0x6f, 0x64, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x32, 0xfa, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x30,
	0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x11, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0f, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x12, 0x2b, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x11, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x0b,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x22, 0x00, 0x12, 0x2b, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0d, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a,
	0x0b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x22, 0x00, 0x12, 0x31,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x64, 0x1a,
	0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22,
	0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_todo_proto_rawDescOnce sync.Once
	file_model_todo_proto_rawDescData = file_model_todo_proto_rawDesc
)

func file_model_todo_proto_rawDescGZIP() []byte {
	file_model_todo_proto_rawDescOnce.Do(func() {
		file_model_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_todo_proto_rawDescData)
	})
	return file_model_todo_proto_rawDescData
}

var file_model_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_model_todo_proto_goTypes = []interface{}{
	(*Todo)(nil),       // 0: model.Todo
	(*TodoInput)(nil),  // 1: model.TodoInput
	(*TodoCreate)(nil), // 2: model.TodoCreate
	(*TodoList)(nil),   // 3: model.TodoList
	(*Deleted)(nil),    // 4: model.Deleted
	(*TodoId)(nil),     // 5: model.TodoId
	(*TodoUpdate)(nil), // 6: model.TodoUpdate
	(*TodoUserId)(nil), // 7: model.TodoUserId
}
var file_model_todo_proto_depIdxs = []int32{
	1, // 0: model.TodoCreate.Todo:type_name -> model.TodoInput
	0, // 1: model.TodoList.list:type_name -> model.Todo
	1, // 2: model.TodoUpdate.Todo:type_name -> model.TodoInput
	7, // 3: model.Todos.ListTodo:input_type -> model.TodoUserId
	2, // 4: model.Todos.AddTodo:input_type -> model.TodoCreate
	5, // 5: model.Todos.GetTodoById:input_type -> model.TodoId
	6, // 6: model.Todos.UpdateTodoById:input_type -> model.TodoUpdate
	5, // 7: model.Todos.DeleteTodoById:input_type -> model.TodoId
	3, // 8: model.Todos.ListTodo:output_type -> model.TodoList
	0, // 9: model.Todos.AddTodo:output_type -> model.Todo
	0, // 10: model.Todos.GetTodoById:output_type -> model.Todo
	0, // 11: model.Todos.UpdateTodoById:output_type -> model.Todo
	4, // 12: model.Todos.DeleteTodoById:output_type -> model.Deleted
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_model_todo_proto_init() }
func file_model_todo_proto_init() {
	if File_model_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_model_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoInput); i {
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
		file_model_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoCreate); i {
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
		file_model_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoList); i {
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
		file_model_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deleted); i {
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
		file_model_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoId); i {
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
		file_model_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoUpdate); i {
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
		file_model_todo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoUserId); i {
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
			RawDescriptor: file_model_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_model_todo_proto_goTypes,
		DependencyIndexes: file_model_todo_proto_depIdxs,
		MessageInfos:      file_model_todo_proto_msgTypes,
	}.Build()
	File_model_todo_proto = out.File
	file_model_todo_proto_rawDesc = nil
	file_model_todo_proto_goTypes = nil
	file_model_todo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TodosClient is the client API for Todos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodosClient interface {
	ListTodo(ctx context.Context, in *TodoUserId, opts ...grpc.CallOption) (*TodoList, error)
	AddTodo(ctx context.Context, in *TodoCreate, opts ...grpc.CallOption) (*Todo, error)
	GetTodoById(ctx context.Context, in *TodoId, opts ...grpc.CallOption) (*Todo, error)
	UpdateTodoById(ctx context.Context, in *TodoUpdate, opts ...grpc.CallOption) (*Todo, error)
	DeleteTodoById(ctx context.Context, in *TodoId, opts ...grpc.CallOption) (*Deleted, error)
}

type todosClient struct {
	cc grpc.ClientConnInterface
}

func NewTodosClient(cc grpc.ClientConnInterface) TodosClient {
	return &todosClient{cc}
}

func (c *todosClient) ListTodo(ctx context.Context, in *TodoUserId, opts ...grpc.CallOption) (*TodoList, error) {
	out := new(TodoList)
	err := c.cc.Invoke(ctx, "/model.Todos/ListTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) AddTodo(ctx context.Context, in *TodoCreate, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/model.Todos/AddTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) GetTodoById(ctx context.Context, in *TodoId, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/model.Todos/GetTodoById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) UpdateTodoById(ctx context.Context, in *TodoUpdate, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/model.Todos/UpdateTodoById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) DeleteTodoById(ctx context.Context, in *TodoId, opts ...grpc.CallOption) (*Deleted, error) {
	out := new(Deleted)
	err := c.cc.Invoke(ctx, "/model.Todos/DeleteTodoById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodosServer is the server API for Todos service.
type TodosServer interface {
	ListTodo(context.Context, *TodoUserId) (*TodoList, error)
	AddTodo(context.Context, *TodoCreate) (*Todo, error)
	GetTodoById(context.Context, *TodoId) (*Todo, error)
	UpdateTodoById(context.Context, *TodoUpdate) (*Todo, error)
	DeleteTodoById(context.Context, *TodoId) (*Deleted, error)
}

// UnimplementedTodosServer can be embedded to have forward compatible implementations.
type UnimplementedTodosServer struct {
}

func (*UnimplementedTodosServer) ListTodo(context.Context, *TodoUserId) (*TodoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodo not implemented")
}
func (*UnimplementedTodosServer) AddTodo(context.Context, *TodoCreate) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTodo not implemented")
}
func (*UnimplementedTodosServer) GetTodoById(context.Context, *TodoId) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoById not implemented")
}
func (*UnimplementedTodosServer) UpdateTodoById(context.Context, *TodoUpdate) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodoById not implemented")
}
func (*UnimplementedTodosServer) DeleteTodoById(context.Context, *TodoId) (*Deleted, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodoById not implemented")
}

func RegisterTodosServer(s *grpc.Server, srv TodosServer) {
	s.RegisterService(&_Todos_serviceDesc, srv)
}

func _Todos_ListTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).ListTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/ListTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).ListTodo(ctx, req.(*TodoUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_AddTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).AddTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/AddTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).AddTodo(ctx, req.(*TodoCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_GetTodoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).GetTodoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/GetTodoById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).GetTodoById(ctx, req.(*TodoId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_UpdateTodoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).UpdateTodoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/UpdateTodoById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).UpdateTodoById(ctx, req.(*TodoUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_DeleteTodoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).DeleteTodoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/DeleteTodoById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).DeleteTodoById(ctx, req.(*TodoId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Todos_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Todos",
	HandlerType: (*TodosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTodo",
			Handler:    _Todos_ListTodo_Handler,
		},
		{
			MethodName: "AddTodo",
			Handler:    _Todos_AddTodo_Handler,
		},
		{
			MethodName: "GetTodoById",
			Handler:    _Todos_GetTodoById_Handler,
		},
		{
			MethodName: "UpdateTodoById",
			Handler:    _Todos_UpdateTodoById_Handler,
		},
		{
			MethodName: "DeleteTodoById",
			Handler:    _Todos_DeleteTodoById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model/todo.proto",
}
