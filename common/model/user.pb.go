// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: model/user.proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	AccessToken string `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *UserFormat) Reset() {
	*x = UserFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFormat) ProtoMessage() {}

func (x *UserFormat) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFormat.ProtoReflect.Descriptor instead.
func (*UserFormat) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserFormat) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserFormat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserFormat) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserFormat) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type UserInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserInput) Reset() {
	*x = UserInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInput) ProtoMessage() {}

func (x *UserInput) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInput.ProtoReflect.Descriptor instead.
func (*UserInput) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInput) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserLogin) Reset() {
	*x = UserLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLogin) ProtoMessage() {}

func (x *UserLogin) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLogin.ProtoReflect.Descriptor instead.
func (*UserLogin) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserLogin) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserLogin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_model_user_proto protoreflect.FileDescriptor

var file_model_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x5c, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x69, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x51, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3d, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x32, 0x6a, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x00,
	0x12, 0x2e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x11, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_model_user_proto_rawDescOnce sync.Once
	file_model_user_proto_rawDescData = file_model_user_proto_rawDesc
)

func file_model_user_proto_rawDescGZIP() []byte {
	file_model_user_proto_rawDescOnce.Do(func() {
		file_model_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_user_proto_rawDescData)
	})
	return file_model_user_proto_rawDescData
}

var file_model_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_model_user_proto_goTypes = []interface{}{
	(*User)(nil),       // 0: model.User
	(*UserFormat)(nil), // 1: model.UserFormat
	(*UserInput)(nil),  // 2: model.UserInput
	(*UserLogin)(nil),  // 3: model.UserLogin
}
var file_model_user_proto_depIdxs = []int32{
	2, // 0: model.Users.Register:input_type -> model.UserInput
	3, // 1: model.Users.Login:input_type -> model.UserLogin
	1, // 2: model.Users.Register:output_type -> model.UserFormat
	1, // 3: model.Users.Login:output_type -> model.UserFormat
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_model_user_proto_init() }
func file_model_user_proto_init() {
	if File_model_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_model_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFormat); i {
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
		file_model_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInput); i {
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
		file_model_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLogin); i {
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
			RawDescriptor: file_model_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_model_user_proto_goTypes,
		DependencyIndexes: file_model_user_proto_depIdxs,
		MessageInfos:      file_model_user_proto_msgTypes,
	}.Build()
	File_model_user_proto = out.File
	file_model_user_proto_rawDesc = nil
	file_model_user_proto_goTypes = nil
	file_model_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	Register(ctx context.Context, in *UserInput, opts ...grpc.CallOption) (*UserFormat, error)
	Login(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*UserFormat, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Register(ctx context.Context, in *UserInput, opts ...grpc.CallOption) (*UserFormat, error) {
	out := new(UserFormat)
	err := c.cc.Invoke(ctx, "/model.Users/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Login(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*UserFormat, error) {
	out := new(UserFormat)
	err := c.cc.Invoke(ctx, "/model.Users/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	Register(context.Context, *UserInput) (*UserFormat, error)
	Login(context.Context, *UserLogin) (*UserFormat, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) Register(context.Context, *UserInput) (*UserFormat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedUsersServer) Login(context.Context, *UserLogin) (*UserFormat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Users/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Register(ctx, req.(*UserInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Users/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Login(ctx, req.(*UserLogin))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Users_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Users_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model/user.proto",
}