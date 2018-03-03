// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

/*
Package go_micro_srv_user is a generated protocol buffer package.

It is generated from these files:
	proto/user/user.proto

It has these top-level messages:
	User
	Request
	Response
	Token
	Error
*/
package go_micro_srv_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Company  string `protobuf:"bytes,3,opt,name=company" json:"company,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Response struct {
	User   *User    `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Users  []*User  `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token  string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Valie  bool     `protobuf:"varint,2,opt,name=valie" json:"valie,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValie() bool {
	if m != nil {
		return m.Valie
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code        int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.user.User")
	proto.RegisterType((*Request)(nil), "go.micro.srv.user.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.user.Response")
	proto.RegisterType((*Token)(nil), "go.micro.srv.user.Token")
	proto.RegisterType((*Error)(nil), "go.micro.srv.user.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
	Auth(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}

func (h *UserService) Auth(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x8b, 0xdb, 0x30,
	0x14, 0x4c, 0xfc, 0x95, 0xe4, 0x85, 0x16, 0xfa, 0x68, 0xa9, 0x48, 0x2f, 0xc1, 0xa7, 0x42, 0xa9,
	0x5b, 0xd2, 0x63, 0x09, 0x34, 0x84, 0x90, 0xbb, 0xfb, 0x71, 0x57, 0xed, 0x47, 0x2b, 0x6a, 0x5b,
	0xae, 0xa4, 0xb8, 0xf4, 0x9f, 0xec, 0x3f, 0xdb, 0xbf, 0xb3, 0xe8, 0x39, 0x59, 0x16, 0xd6, 0xc9,
	0x42, 0x2e, 0xb6, 0x66, 0xde, 0x58, 0x1a, 0xcd, 0x60, 0x78, 0xd5, 0x1a, 0xed, 0xf4, 0x87, 0x83,
	0x25, 0xc3, 0x8f, 0x8c, 0x31, 0xbe, 0xf8, 0xa5, 0xb3, 0x5a, 0x15, 0x46, 0x67, 0xd6, 0x74, 0x99,
	0x1f, 0xa4, 0x1d, 0x44, 0xdf, 0x2d, 0x19, 0x7c, 0x0e, 0x81, 0x2a, 0xc5, 0x78, 0x39, 0x7e, 0x3b,
	0xcb, 0x03, 0x55, 0x22, 0x42, 0xd4, 0xc8, 0x9a, 0x44, 0xc0, 0x0c, 0xaf, 0x51, 0xc0, 0xa4, 0xd0,
	0x75, 0x2b, 0x9b, 0xff, 0x22, 0x64, 0xfa, 0x04, 0xf1, 0x25, 0xc4, 0x54, 0x4b, 0x55, 0x89, 0x88,
	0xf9, 0x1e, 0xe0, 0x02, 0xa6, 0xad, 0xb4, 0xf6, 0x9f, 0x36, 0xa5, 0x88, 0x79, 0x70, 0x8f, 0xd3,
	0x19, 0x4c, 0x72, 0xfa, 0x7b, 0x20, 0xeb, 0xd2, 0x9b, 0x31, 0x4c, 0x73, 0xb2, 0xad, 0x6e, 0x2c,
	0xe1, 0x3b, 0x88, 0xbc, 0x2f, 0x76, 0x32, 0x5f, 0xbd, 0xce, 0x1e, 0x39, 0xce, 0xbc, 0xdd, 0x9c,
	0x45, 0xf8, 0x1e, 0x62, 0xff, 0xb6, 0x22, 0x58, 0x86, 0x97, 0xd4, 0xbd, 0x0a, 0x3f, 0x42, 0x42,
	0xc6, 0x68, 0x63, 0x45, 0xc8, 0x7a, 0x31, 0xa0, 0xdf, 0x79, 0x41, 0x7e, 0xd4, 0xa5, 0x04, 0xf1,
	0x37, 0xfd, 0x87, 0x1a, 0x7f, 0x41, 0xe7, 0x17, 0xc7, 0x84, 0x7a, 0xe0, 0xd9, 0x4e, 0x56, 0xaa,
	0x4f, 0x69, 0x9a, 0xf7, 0xe0, 0x8a, 0x63, 0xd6, 0x10, 0x33, 0xe1, 0x53, 0x2f, 0x74, 0x49, 0x7c,
	0x4a, 0x9c, 0xf3, 0x1a, 0x97, 0x30, 0x2f, 0xc9, 0x16, 0x46, 0xb5, 0x4e, 0xe9, 0xe6, 0x58, 0xc8,
	0x43, 0x6a, 0x75, 0x1b, 0xc0, 0xdc, 0xdf, 0xf3, 0x2b, 0x99, 0x4e, 0x15, 0x84, 0x5f, 0x20, 0xd9,
	0x1a, 0x92, 0x8e, 0xf0, 0x5c, 0x22, 0x8b, 0x37, 0x03, 0x83, 0x53, 0x07, 0xe9, 0x08, 0xd7, 0x10,
	0xee, 0xc9, 0x5d, 0xfd, 0xf9, 0x16, 0x92, 0x3d, 0xb9, 0x4d, 0x55, 0xe1, 0x62, 0x50, 0xc8, 0xbd,
	0x3f, 0xb5, 0xc9, 0x67, 0x88, 0x36, 0x07, 0xf7, 0xfb, 0xbc, 0x89, 0xa1, 0x5c, 0xb9, 0xad, 0x74,
	0x84, 0x3b, 0x78, 0xf6, 0x43, 0x56, 0xaa, 0x94, 0x8e, 0xfa, 0x02, 0xcf, 0x8a, 0x2f, 0x6d, 0xf3,
	0x33, 0xe1, 0xff, 0xe6, 0xd3, 0x5d, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x60, 0x75, 0xe6, 0x50,
	0x03, 0x00, 0x00,
}