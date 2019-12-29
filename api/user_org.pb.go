// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_org.proto

package modoki

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User struct {
	UserId    string                     `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Name      string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	RoleName  string                     `protobuf:"bytes,3,opt,name=role_name,json=roleName" json:"role_name,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *User) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

func (m *User) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type Organization struct {
	Id    int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	OrgId string `protobuf:"bytes,2,opt,name=org_id,json=orgId" json:"org_id,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Organization) Reset()                    { *m = Organization{} }
func (m *Organization) String() string            { return proto.CompactTextString(m) }
func (*Organization) ProtoMessage()               {}
func (*Organization) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Organization) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Organization) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *Organization) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserAddRequest struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *UserAddRequest) Reset()                    { *m = UserAddRequest{} }
func (m *UserAddRequest) String() string            { return proto.CompactTextString(m) }
func (*UserAddRequest) ProtoMessage()               {}
func (*UserAddRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *UserAddRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserAddResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *UserAddResponse) Reset()                    { *m = UserAddResponse{} }
func (m *UserAddResponse) String() string            { return proto.CompactTextString(m) }
func (*UserAddResponse) ProtoMessage()               {}
func (*UserAddResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *UserAddResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserDeleteRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *UserDeleteRequest) Reset()                    { *m = UserDeleteRequest{} }
func (m *UserDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*UserDeleteRequest) ProtoMessage()               {}
func (*UserDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *UserDeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UserDeleteResponse struct {
}

func (m *UserDeleteResponse) Reset()                    { *m = UserDeleteResponse{} }
func (m *UserDeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*UserDeleteResponse) ProtoMessage()               {}
func (*UserDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

type UserFindByIDRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *UserFindByIDRequest) Reset()                    { *m = UserFindByIDRequest{} }
func (m *UserFindByIDRequest) String() string            { return proto.CompactTextString(m) }
func (*UserFindByIDRequest) ProtoMessage()               {}
func (*UserFindByIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *UserFindByIDRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type UserFindByIDResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *UserFindByIDResponse) Reset()                    { *m = UserFindByIDResponse{} }
func (m *UserFindByIDResponse) String() string            { return proto.CompactTextString(m) }
func (*UserFindByIDResponse) ProtoMessage()               {}
func (*UserFindByIDResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *UserFindByIDResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type OrganizationAddRequest struct {
	Organization *Organization `protobuf:"bytes,1,opt,name=organization" json:"organization,omitempty"`
}

func (m *OrganizationAddRequest) Reset()                    { *m = OrganizationAddRequest{} }
func (m *OrganizationAddRequest) String() string            { return proto.CompactTextString(m) }
func (*OrganizationAddRequest) ProtoMessage()               {}
func (*OrganizationAddRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *OrganizationAddRequest) GetOrganization() *Organization {
	if m != nil {
		return m.Organization
	}
	return nil
}

type OrganizationAddResponse struct {
	Organization *Organization `protobuf:"bytes,1,opt,name=organization" json:"organization,omitempty"`
}

func (m *OrganizationAddResponse) Reset()                    { *m = OrganizationAddResponse{} }
func (m *OrganizationAddResponse) String() string            { return proto.CompactTextString(m) }
func (*OrganizationAddResponse) ProtoMessage()               {}
func (*OrganizationAddResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *OrganizationAddResponse) GetOrganization() *Organization {
	if m != nil {
		return m.Organization
	}
	return nil
}

type OrganizationDeleteRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *OrganizationDeleteRequest) Reset()                    { *m = OrganizationDeleteRequest{} }
func (m *OrganizationDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*OrganizationDeleteRequest) ProtoMessage()               {}
func (*OrganizationDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *OrganizationDeleteRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type OrganizationDeleteResponse struct {
}

func (m *OrganizationDeleteResponse) Reset()                    { *m = OrganizationDeleteResponse{} }
func (m *OrganizationDeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*OrganizationDeleteResponse) ProtoMessage()               {}
func (*OrganizationDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

type OrganizationInviteUserRequest struct {
	Target       int32 `protobuf:"varint,1,opt,name=target" json:"target,omitempty"`
	Organization int32 `protobuf:"varint,2,opt,name=organization" json:"organization,omitempty"`
}

func (m *OrganizationInviteUserRequest) Reset()                    { *m = OrganizationInviteUserRequest{} }
func (m *OrganizationInviteUserRequest) String() string            { return proto.CompactTextString(m) }
func (*OrganizationInviteUserRequest) ProtoMessage()               {}
func (*OrganizationInviteUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{12} }

func (m *OrganizationInviteUserRequest) GetTarget() int32 {
	if m != nil {
		return m.Target
	}
	return 0
}

func (m *OrganizationInviteUserRequest) GetOrganization() int32 {
	if m != nil {
		return m.Organization
	}
	return 0
}

type OrganizationInviteUserResponse struct {
}

func (m *OrganizationInviteUserResponse) Reset()                    { *m = OrganizationInviteUserResponse{} }
func (m *OrganizationInviteUserResponse) String() string            { return proto.CompactTextString(m) }
func (*OrganizationInviteUserResponse) ProtoMessage()               {}
func (*OrganizationInviteUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{13} }

type OrganizationListUserRequest struct {
	Organization int32 `protobuf:"varint,1,opt,name=organization" json:"organization,omitempty"`
}

func (m *OrganizationListUserRequest) Reset()                    { *m = OrganizationListUserRequest{} }
func (m *OrganizationListUserRequest) String() string            { return proto.CompactTextString(m) }
func (*OrganizationListUserRequest) ProtoMessage()               {}
func (*OrganizationListUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{14} }

func (m *OrganizationListUserRequest) GetOrganization() int32 {
	if m != nil {
		return m.Organization
	}
	return 0
}

type OrganizationListUserResponse struct {
	Organizations []*User `protobuf:"bytes,1,rep,name=organizations" json:"organizations,omitempty"`
}

func (m *OrganizationListUserResponse) Reset()                    { *m = OrganizationListUserResponse{} }
func (m *OrganizationListUserResponse) String() string            { return proto.CompactTextString(m) }
func (*OrganizationListUserResponse) ProtoMessage()               {}
func (*OrganizationListUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{15} }

func (m *OrganizationListUserResponse) GetOrganizations() []*User {
	if m != nil {
		return m.Organizations
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "modoki.User")
	proto.RegisterType((*Organization)(nil), "modoki.Organization")
	proto.RegisterType((*UserAddRequest)(nil), "modoki.UserAddRequest")
	proto.RegisterType((*UserAddResponse)(nil), "modoki.UserAddResponse")
	proto.RegisterType((*UserDeleteRequest)(nil), "modoki.UserDeleteRequest")
	proto.RegisterType((*UserDeleteResponse)(nil), "modoki.UserDeleteResponse")
	proto.RegisterType((*UserFindByIDRequest)(nil), "modoki.UserFindByIDRequest")
	proto.RegisterType((*UserFindByIDResponse)(nil), "modoki.UserFindByIDResponse")
	proto.RegisterType((*OrganizationAddRequest)(nil), "modoki.OrganizationAddRequest")
	proto.RegisterType((*OrganizationAddResponse)(nil), "modoki.OrganizationAddResponse")
	proto.RegisterType((*OrganizationDeleteRequest)(nil), "modoki.OrganizationDeleteRequest")
	proto.RegisterType((*OrganizationDeleteResponse)(nil), "modoki.OrganizationDeleteResponse")
	proto.RegisterType((*OrganizationInviteUserRequest)(nil), "modoki.OrganizationInviteUserRequest")
	proto.RegisterType((*OrganizationInviteUserResponse)(nil), "modoki.OrganizationInviteUserResponse")
	proto.RegisterType((*OrganizationListUserRequest)(nil), "modoki.OrganizationListUserRequest")
	proto.RegisterType((*OrganizationListUserResponse)(nil), "modoki.OrganizationListUserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserOrg service

type UserOrgClient interface {
	UserAdd(ctx context.Context, in *UserAddRequest, opts ...grpc.CallOption) (*UserAddResponse, error)
	UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error)
	UserFindByID(ctx context.Context, in *UserFindByIDRequest, opts ...grpc.CallOption) (*UserFindByIDResponse, error)
	OrganizationAdd(ctx context.Context, in *OrganizationAddRequest, opts ...grpc.CallOption) (*OrganizationAddResponse, error)
	OrganizationDelete(ctx context.Context, in *OrganizationDeleteRequest, opts ...grpc.CallOption) (*OrganizationDeleteResponse, error)
	OrganizationInviteUser(ctx context.Context, in *OrganizationInviteUserRequest, opts ...grpc.CallOption) (*OrganizationInviteUserResponse, error)
	OrganizationListUser(ctx context.Context, in *OrganizationListUserRequest, opts ...grpc.CallOption) (*OrganizationListUserResponse, error)
}

type userOrgClient struct {
	cc *grpc.ClientConn
}

func NewUserOrgClient(cc *grpc.ClientConn) UserOrgClient {
	return &userOrgClient{cc}
}

func (c *userOrgClient) UserAdd(ctx context.Context, in *UserAddRequest, opts ...grpc.CallOption) (*UserAddResponse, error) {
	out := new(UserAddResponse)
	err := grpc.Invoke(ctx, "/modoki.UserOrg/UserAdd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrgClient) UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error) {
	out := new(UserDeleteResponse)
	err := grpc.Invoke(ctx, "/modoki.UserOrg/UserDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrgClient) UserFindByID(ctx context.Context, in *UserFindByIDRequest, opts ...grpc.CallOption) (*UserFindByIDResponse, error) {
	out := new(UserFindByIDResponse)
	err := grpc.Invoke(ctx, "/modoki.UserOrg/UserFindByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrgClient) OrganizationAdd(ctx context.Context, in *OrganizationAddRequest, opts ...grpc.CallOption) (*OrganizationAddResponse, error) {
	out := new(OrganizationAddResponse)
	err := grpc.Invoke(ctx, "/modoki.UserOrg/OrganizationAdd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrgClient) OrganizationDelete(ctx context.Context, in *OrganizationDeleteRequest, opts ...grpc.CallOption) (*OrganizationDeleteResponse, error) {
	out := new(OrganizationDeleteResponse)
	err := grpc.Invoke(ctx, "/modoki.UserOrg/OrganizationDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrgClient) OrganizationInviteUser(ctx context.Context, in *OrganizationInviteUserRequest, opts ...grpc.CallOption) (*OrganizationInviteUserResponse, error) {
	out := new(OrganizationInviteUserResponse)
	err := grpc.Invoke(ctx, "/modoki.UserOrg/OrganizationInviteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrgClient) OrganizationListUser(ctx context.Context, in *OrganizationListUserRequest, opts ...grpc.CallOption) (*OrganizationListUserResponse, error) {
	out := new(OrganizationListUserResponse)
	err := grpc.Invoke(ctx, "/modoki.UserOrg/OrganizationListUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserOrg service

type UserOrgServer interface {
	UserAdd(context.Context, *UserAddRequest) (*UserAddResponse, error)
	UserDelete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error)
	UserFindByID(context.Context, *UserFindByIDRequest) (*UserFindByIDResponse, error)
	OrganizationAdd(context.Context, *OrganizationAddRequest) (*OrganizationAddResponse, error)
	OrganizationDelete(context.Context, *OrganizationDeleteRequest) (*OrganizationDeleteResponse, error)
	OrganizationInviteUser(context.Context, *OrganizationInviteUserRequest) (*OrganizationInviteUserResponse, error)
	OrganizationListUser(context.Context, *OrganizationListUserRequest) (*OrganizationListUserResponse, error)
}

func RegisterUserOrgServer(s *grpc.Server, srv UserOrgServer) {
	s.RegisterService(&_UserOrg_serviceDesc, srv)
}

func _UserOrg_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOrgServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.UserOrg/UserAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOrgServer).UserAdd(ctx, req.(*UserAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOrg_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOrgServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.UserOrg/UserDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOrgServer).UserDelete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOrg_UserFindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFindByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOrgServer).UserFindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.UserOrg/UserFindByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOrgServer).UserFindByID(ctx, req.(*UserFindByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOrg_OrganizationAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOrgServer).OrganizationAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.UserOrg/OrganizationAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOrgServer).OrganizationAdd(ctx, req.(*OrganizationAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOrg_OrganizationDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOrgServer).OrganizationDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.UserOrg/OrganizationDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOrgServer).OrganizationDelete(ctx, req.(*OrganizationDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOrg_OrganizationInviteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationInviteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOrgServer).OrganizationInviteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.UserOrg/OrganizationInviteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOrgServer).OrganizationInviteUser(ctx, req.(*OrganizationInviteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOrg_OrganizationListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOrgServer).OrganizationListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.UserOrg/OrganizationListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOrgServer).OrganizationListUser(ctx, req.(*OrganizationListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserOrg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "modoki.UserOrg",
	HandlerType: (*UserOrgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserAdd",
			Handler:    _UserOrg_UserAdd_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _UserOrg_UserDelete_Handler,
		},
		{
			MethodName: "UserFindByID",
			Handler:    _UserOrg_UserFindByID_Handler,
		},
		{
			MethodName: "OrganizationAdd",
			Handler:    _UserOrg_OrganizationAdd_Handler,
		},
		{
			MethodName: "OrganizationDelete",
			Handler:    _UserOrg_OrganizationDelete_Handler,
		},
		{
			MethodName: "OrganizationInviteUser",
			Handler:    _UserOrg_OrganizationInviteUser_Handler,
		},
		{
			MethodName: "OrganizationListUser",
			Handler:    _UserOrg_OrganizationListUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_org.proto",
}

func init() { proto.RegisterFile("user_org.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdb, 0x6e, 0xd3, 0x30,
	0x18, 0x56, 0x0f, 0xeb, 0xd6, 0xbf, 0xa5, 0x13, 0xa6, 0xb4, 0x5d, 0x5a, 0xb6, 0xe2, 0x01, 0x9a,
	0x84, 0x94, 0x49, 0xdd, 0xcd, 0xe0, 0xae, 0x63, 0x42, 0x8a, 0x84, 0x98, 0x14, 0xe0, 0x86, 0x5d,
	0x54, 0x19, 0x36, 0x51, 0x44, 0x1b, 0x17, 0xc7, 0x45, 0x82, 0x07, 0xe4, 0x01, 0x78, 0x22, 0x64,
	0xc7, 0x21, 0x4e, 0xe2, 0x6e, 0x13, 0x77, 0xb1, 0xbf, 0xc3, 0x7f, 0xc8, 0x97, 0x16, 0x7a, 0x9b,
	0x84, 0xf2, 0x05, 0xe3, 0xa1, 0xbb, 0xe6, 0x4c, 0x30, 0xd4, 0x5a, 0x31, 0xc2, 0xbe, 0x45, 0xce,
	0x51, 0xc8, 0x58, 0xb8, 0xa4, 0xa7, 0xea, 0xf6, 0x66, 0xf3, 0xf5, 0x54, 0x44, 0x2b, 0x9a, 0x88,
	0x60, 0xb5, 0x4e, 0x89, 0xf8, 0x77, 0x0d, 0x9a, 0x9f, 0x12, 0xca, 0xd1, 0x10, 0x76, 0x95, 0x47,
	0x44, 0x46, 0xb5, 0x69, 0xed, 0xa4, 0xed, 0xb7, 0xe4, 0xd1, 0x23, 0x08, 0x41, 0x33, 0x0e, 0x56,
	0x74, 0x54, 0x57, 0xb7, 0xea, 0x19, 0x8d, 0xa1, 0xcd, 0xd9, 0x92, 0x2e, 0x14, 0xd0, 0x50, 0xc0,
	0x9e, 0xbc, 0x78, 0x2f, 0xc1, 0x57, 0x00, 0x5f, 0x38, 0x0d, 0x04, 0x25, 0x8b, 0x40, 0x8c, 0x60,
	0x5a, 0x3b, 0xe9, 0xcc, 0x1c, 0x37, 0x6d, 0xc4, 0xcd, 0x1a, 0x71, 0x3f, 0x66, 0x8d, 0xf8, 0x6d,
	0xcd, 0x9e, 0x0b, 0x29, 0xdd, 0xac, 0x49, 0x26, 0xed, 0xdc, 0x2d, 0xd5, 0xec, 0xb9, 0xc0, 0x1e,
	0x74, 0xaf, 0x78, 0x18, 0xc4, 0xd1, 0xaf, 0x40, 0x44, 0x2c, 0x46, 0x3d, 0xa8, 0xeb, 0x51, 0x76,
	0xfc, 0x7a, 0x44, 0xd0, 0x63, 0x68, 0x31, 0x1e, 0xca, 0xf1, 0xd2, 0x41, 0x76, 0x18, 0x0f, 0x8d,
	0xe9, 0x1a, 0xf9, 0x74, 0x78, 0x06, 0x3d, 0xb9, 0x92, 0x39, 0x21, 0x3e, 0xfd, 0xbe, 0xa1, 0x89,
	0x40, 0x53, 0x68, 0xca, 0x6d, 0x28, 0xbb, 0xce, 0xac, 0xeb, 0xa6, 0xdb, 0x75, 0x25, 0xcb, 0x57,
	0x08, 0x3e, 0x83, 0xfd, 0x7f, 0x9a, 0x64, 0xcd, 0xe2, 0x84, 0xde, 0x43, 0x74, 0x0c, 0x0f, 0xe5,
	0xe9, 0x92, 0x2e, 0xa9, 0xa0, 0x59, 0xad, 0xbc, 0xf1, 0xb6, 0x6c, 0x1c, 0xf7, 0x01, 0x99, 0xa4,
	0xd4, 0x1c, 0xbb, 0xf0, 0x48, 0xde, 0xbe, 0x8d, 0x62, 0x72, 0xf1, 0xd3, 0xbb, 0xcc, 0xc4, 0xdb,
	0xde, 0x22, 0x3e, 0x87, 0x7e, 0x91, 0x7f, 0xef, 0x26, 0x7d, 0x18, 0x98, 0x8b, 0x35, 0xb6, 0x72,
	0x0e, 0x5d, 0x66, 0x20, 0xda, 0xa3, 0x9f, 0x79, 0x98, 0x2a, 0xbf, 0xc0, 0xc4, 0x1f, 0x60, 0x58,
	0xf1, 0xd4, 0x0d, 0xfd, 0xbf, 0xe9, 0x4b, 0x38, 0x30, 0xd1, 0x6d, 0x5b, 0x55, 0x71, 0xc0, 0x13,
	0x70, 0x6c, 0x64, 0xbd, 0xdd, 0x6b, 0x78, 0x62, 0xa2, 0x5e, 0xfc, 0x23, 0x12, 0x54, 0xed, 0x44,
	0xdb, 0x0d, 0xa0, 0x25, 0x02, 0x1e, 0x52, 0xa1, 0x2d, 0xf5, 0x09, 0xe1, 0x52, 0xf7, 0x75, 0x85,
	0x16, 0xfb, 0x9c, 0xc2, 0xe1, 0x36, 0x73, 0x5d, 0x7e, 0x0e, 0x63, 0x93, 0xf1, 0x2e, 0x4a, 0x84,
	0x59, 0x1c, 0x5b, 0x56, 0x54, 0x2e, 0xe2, 0xc3, 0xc4, 0x6e, 0xa1, 0xd7, 0x3c, 0x83, 0x07, 0x26,
	0x3f, 0x19, 0xd5, 0xa6, 0x8d, 0x4a, 0x00, 0x8a, 0x94, 0xd9, 0x9f, 0x26, 0xec, 0xca, 0xfb, 0x2b,
	0x1e, 0xa2, 0xd7, 0xe9, 0xe3, 0x9c, 0x10, 0x34, 0x30, 0x35, 0x79, 0x3c, 0x9c, 0x61, 0xe5, 0x5e,
	0xd7, 0x7e, 0x03, 0x90, 0x27, 0x1a, 0x1d, 0x98, 0xb4, 0xc2, 0x4b, 0x73, 0x1c, 0x1b, 0xa4, 0x4d,
	0x3c, 0xe8, 0x9a, 0x81, 0x46, 0x63, 0x93, 0x5b, 0xfa, 0x2c, 0x9c, 0x89, 0x1d, 0xd4, 0x56, 0x3e,
	0xec, 0x97, 0xd2, 0x88, 0x0e, 0x6d, 0x79, 0x33, 0x66, 0x3b, 0xda, 0x8a, 0x6b, 0xcf, 0x6b, 0x40,
	0xd5, 0x7c, 0xa1, 0xa7, 0x36, 0x59, 0x71, 0x66, 0x7c, 0x1b, 0x45, 0x9b, 0x87, 0xc5, 0x4f, 0x32,
	0x4f, 0x10, 0x7a, 0x6e, 0x53, 0x57, 0xe2, 0xeb, 0xbc, 0xb8, 0x8b, 0xa6, 0x0b, 0x05, 0xd0, 0xb7,
	0xa5, 0x08, 0x1d, 0xdb, 0xf4, 0xa5, 0x98, 0x3a, 0xcf, 0x6e, 0x27, 0xa5, 0x25, 0x2e, 0xf6, 0x3e,
	0xeb, 0xff, 0xaa, 0x9b, 0x96, 0xfa, 0x81, 0x3f, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x74,
	0xbc, 0xba, 0xcc, 0x06, 0x00, 0x00,
}
