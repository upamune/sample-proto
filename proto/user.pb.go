// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/user.proto

It has these top-level messages:
	User
	CreateRequest
	CreateResponse
	GetRequest
	GetResponse
	GetAllRequest
	GetAllResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Sex int32

const (
	Sex_MALE   Sex = 0
	Sex_FEMALE Sex = 1
)

var Sex_name = map[int32]string{
	0: "MALE",
	1: "FEMALE",
}
var Sex_value = map[string]int32{
	"MALE":   0,
	"FEMALE": 1,
}

func (x Sex) String() string {
	return proto1.EnumName(Sex_name, int32(x))
}
func (Sex) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Status int32

const (
	Status_OK   Status = 0
	Status_FAIL Status = 1
)

var Status_name = map[int32]string{
	0: "OK",
	1: "FAIL",
}
var Status_value = map[string]int32{
	"OK":   0,
	"FAIL": 1,
}

func (x Status) String() string {
	return proto1.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type User struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Sex  Sex    `protobuf:"varint,3,opt,name=sex,enum=proto.Sex" json:"sex,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto1.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetSex() Sex {
	if m != nil {
		return m.Sex
	}
	return Sex_MALE
}

type CreateRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Sex  Sex    `protobuf:"varint,3,opt,name=sex,enum=proto.Sex" json:"sex,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *CreateRequest) GetSex() Sex {
	if m != nil {
		return m.Sex
	}
	return Sex_MALE
}

type CreateResponse struct {
	Status  Status `protobuf:"varint,1,opt,name=status,enum=proto.Status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto1.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func (m *CreateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetResponse struct {
	Status  Status `protobuf:"varint,1,opt,name=status,enum=proto.Status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	User    *User  `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func (m *GetResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetAllRequest struct {
}

func (m *GetAllRequest) Reset()                    { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()               {}
func (*GetAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type GetAllResponse struct {
	Status  Status  `protobuf:"varint,1,opt,name=status,enum=proto.Status" json:"status,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Users   []*User `protobuf:"bytes,3,rep,name=users" json:"users,omitempty"`
}

func (m *GetAllResponse) Reset()                    { *m = GetAllResponse{} }
func (m *GetAllResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetAllResponse) ProtoMessage()               {}
func (*GetAllResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetAllResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func (m *GetAllResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetAllResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UpdateRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Sex  Sex    `protobuf:"varint,3,opt,name=sex,enum=proto.Sex" json:"sex,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto1.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRequest) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UpdateRequest) GetSex() Sex {
	if m != nil {
		return m.Sex
	}
	return Sex_MALE
}

type UpdateResponse struct {
	Status  Status `protobuf:"varint,1,opt,name=status,enum=proto.Status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto1.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func (m *UpdateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteResponse struct {
	Status  Status `protobuf:"varint,1,opt,name=status,enum=proto.Status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto1.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeleteResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func (m *DeleteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto1.RegisterType((*User)(nil), "proto.User")
	proto1.RegisterType((*CreateRequest)(nil), "proto.CreateRequest")
	proto1.RegisterType((*CreateResponse)(nil), "proto.CreateResponse")
	proto1.RegisterType((*GetRequest)(nil), "proto.GetRequest")
	proto1.RegisterType((*GetResponse)(nil), "proto.GetResponse")
	proto1.RegisterType((*GetAllRequest)(nil), "proto.GetAllRequest")
	proto1.RegisterType((*GetAllResponse)(nil), "proto.GetAllResponse")
	proto1.RegisterType((*UpdateRequest)(nil), "proto.UpdateRequest")
	proto1.RegisterType((*UpdateResponse)(nil), "proto.UpdateResponse")
	proto1.RegisterType((*DeleteRequest)(nil), "proto.DeleteRequest")
	proto1.RegisterType((*DeleteResponse)(nil), "proto.DeleteResponse")
	proto1.RegisterEnum("proto.Sex", Sex_name, Sex_value)
	proto1.RegisterEnum("proto.Status", Status_name, Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	// CREATE
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// READ
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	// UPDATE
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// DELETE
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/proto.UserService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/proto.UserService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := grpc.Invoke(ctx, "/proto.UserService/GetAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/proto.UserService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/proto.UserService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	// CREATE
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// READ
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	// UPDATE
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// DELETE
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UserService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}

func init() { proto1.RegisterFile("proto/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xed, 0x4a, 0xe3, 0x40,
	0x14, 0x6d, 0x9a, 0x74, 0x76, 0x7b, 0x43, 0xb2, 0xd9, 0xcb, 0x2e, 0x84, 0xee, 0x82, 0x31, 0x22,
	0x94, 0x22, 0x15, 0x22, 0x79, 0x80, 0xa2, 0x6d, 0x51, 0x2b, 0x62, 0x42, 0x1f, 0x20, 0xea, 0x45,
	0x84, 0xb6, 0xa9, 0x99, 0xa9, 0xf4, 0x4d, 0x7c, 0x5d, 0x99, 0xc9, 0x87, 0x26, 0x3f, 0xfa, 0x2b,
	0xbf, 0x7a, 0x7b, 0xbf, 0xce, 0x99, 0x7b, 0x4e, 0xc0, 0xd9, 0x66, 0xa9, 0x48, 0xcf, 0x77, 0x9c,
	0xb2, 0xb1, 0x0a, 0xb1, 0xa7, 0x7e, 0xfc, 0x1b, 0x30, 0x96, 0x9c, 0x32, 0x44, 0x30, 0x36, 0xc9,
	0x9a, 0x5c, 0xcd, 0xd3, 0x86, 0xfd, 0x48, 0xc5, 0xe8, 0x80, 0x9e, 0xbc, 0x90, 0xdb, 0xf5, 0xb4,
	0x61, 0x2f, 0x92, 0x21, 0xfe, 0x07, 0x9d, 0xd3, 0xde, 0xd5, 0x3d, 0x6d, 0x68, 0x07, 0x90, 0x6f,
	0x1a, 0xc7, 0xb4, 0x8f, 0x64, 0xda, 0x8f, 0xc1, 0xba, 0xcc, 0x28, 0x11, 0x14, 0xd1, 0xdb, 0x8e,
	0xb8, 0x68, 0x65, 0xe9, 0x03, 0xd8, 0xe5, 0x52, 0xbe, 0x4d, 0x37, 0x9c, 0xf0, 0x14, 0x18, 0x17,
	0x89, 0xd8, 0x71, 0xb5, 0xd7, 0x0e, 0xac, 0x72, 0x44, 0x25, 0xa3, 0xa2, 0x88, 0x2e, 0xfc, 0x58,
	0x13, 0xe7, 0x25, 0x58, 0x3f, 0x2a, 0xff, 0xfa, 0x1e, 0xc0, 0x9c, 0xc4, 0x01, 0x92, 0x7e, 0x0a,
	0xa6, 0xea, 0x68, 0x09, 0x11, 0x8f, 0xc0, 0x90, 0xa7, 0x57, 0x6f, 0x34, 0x03, 0xb3, 0x18, 0x97,
	0x87, 0x8f, 0x54, 0xc1, 0xff, 0x05, 0xd6, 0x9c, 0xc4, 0x64, 0xb5, 0x2a, 0x58, 0xf9, 0x02, 0xec,
	0x32, 0xd1, 0x16, 0x89, 0x63, 0xe8, 0x49, 0x2c, 0xee, 0xea, 0x9e, 0xde, 0x64, 0x91, 0x57, 0xa4,
	0x82, 0xcb, 0xed, 0x73, 0xfb, 0x0a, 0x96, 0x4b, 0xdb, 0x52, 0xf0, 0x04, 0xac, 0x2b, 0x5a, 0xd1,
	0x41, 0x9e, 0x12, 0xb7, 0x6c, 0x6a, 0x09, 0x77, 0xf4, 0x0f, 0xf4, 0x98, 0xf6, 0xf8, 0x13, 0x8c,
	0xbb, 0xc9, 0x62, 0xea, 0x74, 0x10, 0x80, 0xcd, 0xa6, 0x2a, 0xd6, 0x46, 0x03, 0x60, 0xf9, 0x22,
	0x64, 0xd0, 0xbd, 0xbf, 0x75, 0x3a, 0xb2, 0x6f, 0x36, 0xb9, 0x5e, 0x38, 0x5a, 0xf0, 0xd1, 0x05,
	0x53, 0x1e, 0x3a, 0xa6, 0xec, 0xfd, 0xf5, 0x89, 0x30, 0x04, 0x96, 0xbb, 0x1a, 0xff, 0x14, 0x1c,
	0x6a, 0x5f, 0xce, 0xe0, 0x6f, 0x23, 0x5b, 0x3c, 0xe0, 0x0c, 0xf4, 0x39, 0x09, 0xfc, 0x5d, 0x54,
	0xbf, 0x5c, 0x3c, 0xc0, 0xef, 0xa9, 0xa2, 0x3b, 0x04, 0x96, 0x7b, 0xa8, 0x02, 0xa9, 0x79, 0xac,
	0x02, 0x69, 0x18, 0x2d, 0x04, 0x96, 0xeb, 0x55, 0x8d, 0xd5, 0x3c, 0x51, 0x8d, 0x35, 0x44, 0x0d,
	0x81, 0xe5, 0xe7, 0xae, 0xc6, 0x6a, 0x12, 0x55, 0x63, 0x75, 0x4d, 0x1e, 0x99, 0xca, 0x5e, 0x7c,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xd8, 0x02, 0x76, 0xa2, 0x04, 0x00, 0x00,
}
