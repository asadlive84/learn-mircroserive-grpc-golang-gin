// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/pb/auth.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CheckUserRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckUserRequest) Reset()         { *m = CheckUserRequest{} }
func (m *CheckUserRequest) String() string { return proto.CompactTextString(m) }
func (*CheckUserRequest) ProtoMessage()    {}
func (*CheckUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_329ec36c14ab2ce3, []int{0}
}

func (m *CheckUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckUserRequest.Unmarshal(m, b)
}
func (m *CheckUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckUserRequest.Marshal(b, m, deterministic)
}
func (m *CheckUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckUserRequest.Merge(m, src)
}
func (m *CheckUserRequest) XXX_Size() int {
	return xxx_messageInfo_CheckUserRequest.Size(m)
}
func (m *CheckUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckUserRequest proto.InternalMessageInfo

func (m *CheckUserRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type CheckUserResponse struct {
	Status               int64            `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string           `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Message              string           `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	UserInfo             *RegisterRequest `protobuf:"bytes,4,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CheckUserResponse) Reset()         { *m = CheckUserResponse{} }
func (m *CheckUserResponse) String() string { return proto.CompactTextString(m) }
func (*CheckUserResponse) ProtoMessage()    {}
func (*CheckUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_329ec36c14ab2ce3, []int{1}
}

func (m *CheckUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckUserResponse.Unmarshal(m, b)
}
func (m *CheckUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckUserResponse.Marshal(b, m, deterministic)
}
func (m *CheckUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckUserResponse.Merge(m, src)
}
func (m *CheckUserResponse) XXX_Size() int {
	return xxx_messageInfo_CheckUserResponse.Size(m)
}
func (m *CheckUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckUserResponse proto.InternalMessageInfo

func (m *CheckUserResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CheckUserResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CheckUserResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CheckUserResponse) GetUserInfo() *RegisterRequest {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

type RegisterRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_329ec36c14ab2ce3, []int{2}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type RegisterResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_329ec36c14ab2ce3, []int{3}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RegisterResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_329ec36c14ab2ce3, []int{4}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type LoginResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_329ec36c14ab2ce3, []int{5}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *LoginResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckUserRequest)(nil), "auth.CheckUserRequest")
	proto.RegisterType((*CheckUserResponse)(nil), "auth.CheckUserResponse")
	proto.RegisterType((*RegisterRequest)(nil), "auth.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "auth.RegisterResponse")
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "auth.LoginResponse")
}

func init() {
	proto.RegisterFile("pkg/pb/auth.proto", fileDescriptor_329ec36c14ab2ce3)
}

var fileDescriptor_329ec36c14ab2ce3 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x86, 0xad, 0x7c, 0x58, 0x06, 0x8d, 0x30, 0x22, 0x36, 0x3d, 0x91, 0x9e, 0x88, 0x07, 0x48,
	0xea, 0xd1, 0xc4, 0xf8, 0x75, 0x31, 0xf1, 0x54, 0xa2, 0x89, 0x5e, 0x4c, 0xc1, 0xa1, 0x6d, 0x90,
	0x6e, 0xdd, 0xd9, 0xea, 0xcf, 0xf0, 0x1f, 0xf9, 0xdb, 0x4c, 0xb7, 0x1f, 0x22, 0x7a, 0x22, 0xde,
	0xfa, 0x4e, 0x67, 0xde, 0x7d, 0x76, 0xe7, 0x85, 0x6e, 0xb2, 0x08, 0xc6, 0xc9, 0x74, 0xec, 0xa7,
	0x2a, 0x1c, 0x25, 0x52, 0x28, 0x81, 0xf5, 0xec, 0xdb, 0x39, 0x86, 0xce, 0x55, 0x48, 0xb3, 0xc5,
	0x1d, 0x93, 0xf4, 0xe8, 0x35, 0x25, 0x56, 0xd8, 0x87, 0x66, 0xca, 0x24, 0x6f, 0xae, 0x2d, 0x63,
	0x60, 0x0c, 0x5b, 0x5e, 0xa1, 0x9c, 0x0f, 0x03, 0xba, 0x2b, 0xcd, 0x9c, 0x88, 0x98, 0x29, 0xeb,
	0x66, 0xe5, 0xab, 0x94, 0x75, 0x77, 0xcd, 0x2b, 0x14, 0xf6, 0xa0, 0x41, 0x52, 0x0a, 0x69, 0x6d,
	0x6b, 0x93, 0x5c, 0xa0, 0x05, 0x3b, 0x4b, 0x62, 0xf6, 0x03, 0xb2, 0x6a, 0xba, 0x5e, 0x4a, 0x74,
	0xa1, 0x95, 0x9d, 0xf3, 0x14, 0xc5, 0x73, 0x61, 0xd5, 0x07, 0xc6, 0xb0, 0xed, 0x1e, 0x8e, 0x34,
	0xaf, 0x47, 0x41, 0xc4, 0xaa, 0xe2, 0xf3, 0x4c, 0xcd, 0x13, 0xcf, 0x85, 0xf3, 0x00, 0xfb, 0x6b,
	0x3f, 0xf5, 0xb1, 0x4b, 0x3f, 0x7a, 0x29, 0xd8, 0x73, 0x81, 0x36, 0x98, 0x89, 0xcf, 0xfc, 0x2e,
	0xe4, 0x73, 0xc1, 0x53, 0xe9, 0x6c, 0x22, 0x09, 0x45, 0x5c, 0x02, 0xe5, 0xc2, 0x39, 0x87, 0xce,
	0xb7, 0xf5, 0x26, 0x57, 0x75, 0xee, 0x61, 0xf7, 0x56, 0x04, 0x51, 0xfc, 0xdf, 0x64, 0x13, 0xd8,
	0x2b, 0x7c, 0x37, 0xda, 0x40, 0x0f, 0x1a, 0x4a, 0x2c, 0x28, 0x2e, 0x4d, 0xb5, 0x70, 0x3f, 0x0d,
	0x68, 0x5f, 0xa4, 0x2a, 0x9c, 0x90, 0x7c, 0x8b, 0x66, 0x84, 0xa7, 0x60, 0x96, 0xd7, 0xc7, 0xbf,
	0xd7, 0x60, 0xf7, 0xd7, 0xcb, 0x39, 0x8e, 0xb3, 0x85, 0x2e, 0x34, 0x34, 0x21, 0x62, 0xde, 0xb2,
	0xfa, 0x0c, 0xf6, 0xc1, 0x8f, 0x5a, 0x35, 0x73, 0x06, 0xad, 0x2a, 0x5b, 0x58, 0x58, 0xaf, 0x27,
	0xd3, 0x3e, 0xfa, 0x55, 0x2f, 0xe7, 0x2f, 0xe1, 0xd1, 0x1c, 0x8d, 0xf3, 0x94, 0x4f, 0x9b, 0x3a,
	0xe1, 0x27, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0x78, 0x7f, 0xad, 0xf6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// rpc Validate(ValidateRequest) returns (ValidateResponse) {}
	CheckUser(ctx context.Context, in *CheckUserRequest, opts ...grpc.CallOption) (*CheckUserResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckUser(ctx context.Context, in *CheckUserRequest, opts ...grpc.CallOption) (*CheckUserResponse, error) {
	out := new(CheckUserResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/CheckUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// rpc Validate(ValidateRequest) returns (ValidateResponse) {}
	CheckUser(context.Context, *CheckUserRequest) (*CheckUserResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedAuthServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServiceServer) CheckUser(ctx context.Context, req *CheckUserRequest) (*CheckUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUser not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/CheckUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckUser(ctx, req.(*CheckUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AuthService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "CheckUser",
			Handler:    _AuthService_CheckUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/auth.proto",
}
