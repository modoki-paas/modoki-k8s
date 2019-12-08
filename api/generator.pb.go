// Code generated by protoc-gen-go. DO NOT EDIT.
// source: generator.proto

/*
Package modoki is a generated protocol buffer package.

It is generated from these files:
	generator.proto
	service.proto

It has these top-level messages:
	KubernetesConfig
	YAML
	OperateRequest
	OperateResponse
	MetricsRequest
	MetricsResponse
	AppSpec
	AppStatus
	AppCreateRequest
	AppCreateResponse
*/
package modoki

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type OperateKind int32

const (
	OperateKind_Apply  OperateKind = 0
	OperateKind_Delete OperateKind = 1
)

var OperateKind_name = map[int32]string{
	0: "Apply",
	1: "Delete",
}
var OperateKind_value = map[string]int32{
	"Apply":  0,
	"Delete": 1,
}

func (x OperateKind) String() string {
	return proto.EnumName(OperateKind_name, int32(x))
}
func (OperateKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type KubernetesConfig struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *KubernetesConfig) Reset()                    { *m = KubernetesConfig{} }
func (m *KubernetesConfig) String() string            { return proto.CompactTextString(m) }
func (*KubernetesConfig) ProtoMessage()               {}
func (*KubernetesConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KubernetesConfig) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type YAML struct {
	Config string `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
}

func (m *YAML) Reset()                    { *m = YAML{} }
func (m *YAML) String() string            { return proto.CompactTextString(m) }
func (*YAML) ProtoMessage()               {}
func (*YAML) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *YAML) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

type OperateRequest struct {
	Id         string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Kind       OperateKind       `protobuf:"varint,2,opt,name=kind,enum=modoki.OperateKind" json:"kind,omitempty"`
	Performer  int32             `protobuf:"varint,3,opt,name=performer" json:"performer,omitempty"`
	Spec       *AppSpec          `protobuf:"bytes,4,opt,name=spec" json:"spec,omitempty"`
	DeleteYaml *YAML             `protobuf:"bytes,5,opt,name=delete_yaml,json=deleteYaml" json:"delete_yaml,omitempty"`
	ApplyYaml  *YAML             `protobuf:"bytes,6,opt,name=apply_yaml,json=applyYaml" json:"apply_yaml,omitempty"`
	K8SConfig  *KubernetesConfig `protobuf:"bytes,7,opt,name=k8s_config,json=k8sConfig" json:"k8s_config,omitempty"`
}

func (m *OperateRequest) Reset()                    { *m = OperateRequest{} }
func (m *OperateRequest) String() string            { return proto.CompactTextString(m) }
func (*OperateRequest) ProtoMessage()               {}
func (*OperateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OperateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OperateRequest) GetKind() OperateKind {
	if m != nil {
		return m.Kind
	}
	return OperateKind_Apply
}

func (m *OperateRequest) GetPerformer() int32 {
	if m != nil {
		return m.Performer
	}
	return 0
}

func (m *OperateRequest) GetSpec() *AppSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *OperateRequest) GetDeleteYaml() *YAML {
	if m != nil {
		return m.DeleteYaml
	}
	return nil
}

func (m *OperateRequest) GetApplyYaml() *YAML {
	if m != nil {
		return m.ApplyYaml
	}
	return nil
}

func (m *OperateRequest) GetK8SConfig() *KubernetesConfig {
	if m != nil {
		return m.K8SConfig
	}
	return nil
}

type OperateResponse struct {
	DeleteYaml *YAML `protobuf:"bytes,1,opt,name=delete_yaml,json=deleteYaml" json:"delete_yaml,omitempty"`
	ApplyYaml  *YAML `protobuf:"bytes,2,opt,name=apply_yaml,json=applyYaml" json:"apply_yaml,omitempty"`
}

func (m *OperateResponse) Reset()                    { *m = OperateResponse{} }
func (m *OperateResponse) String() string            { return proto.CompactTextString(m) }
func (*OperateResponse) ProtoMessage()               {}
func (*OperateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OperateResponse) GetDeleteYaml() *YAML {
	if m != nil {
		return m.DeleteYaml
	}
	return nil
}

func (m *OperateResponse) GetApplyYaml() *YAML {
	if m != nil {
		return m.ApplyYaml
	}
	return nil
}

type MetricsRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *MetricsRequest) Reset()                    { *m = MetricsRequest{} }
func (m *MetricsRequest) String() string            { return proto.CompactTextString(m) }
func (*MetricsRequest) ProtoMessage()               {}
func (*MetricsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MetricsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MetricsResponse struct {
	Status *AppStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *MetricsResponse) Reset()                    { *m = MetricsResponse{} }
func (m *MetricsResponse) String() string            { return proto.CompactTextString(m) }
func (*MetricsResponse) ProtoMessage()               {}
func (*MetricsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MetricsResponse) GetStatus() *AppStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*KubernetesConfig)(nil), "modoki.KubernetesConfig")
	proto.RegisterType((*YAML)(nil), "modoki.YAML")
	proto.RegisterType((*OperateRequest)(nil), "modoki.OperateRequest")
	proto.RegisterType((*OperateResponse)(nil), "modoki.OperateResponse")
	proto.RegisterType((*MetricsRequest)(nil), "modoki.MetricsRequest")
	proto.RegisterType((*MetricsResponse)(nil), "modoki.MetricsResponse")
	proto.RegisterEnum("modoki.OperateKind", OperateKind_name, OperateKind_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Generator service

type GeneratorClient interface {
	Operate(ctx context.Context, in *OperateRequest, opts ...grpc.CallOption) (*OperateResponse, error)
	Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error)
}

type generatorClient struct {
	cc *grpc.ClientConn
}

func NewGeneratorClient(cc *grpc.ClientConn) GeneratorClient {
	return &generatorClient{cc}
}

func (c *generatorClient) Operate(ctx context.Context, in *OperateRequest, opts ...grpc.CallOption) (*OperateResponse, error) {
	out := new(OperateResponse)
	err := grpc.Invoke(ctx, "/modoki.Generator/Operate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorClient) Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := grpc.Invoke(ctx, "/modoki.Generator/Metrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Generator service

type GeneratorServer interface {
	Operate(context.Context, *OperateRequest) (*OperateResponse, error)
	Metrics(context.Context, *MetricsRequest) (*MetricsResponse, error)
}

func RegisterGeneratorServer(s *grpc.Server, srv GeneratorServer) {
	s.RegisterService(&_Generator_serviceDesc, srv)
}

func _Generator_Operate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).Operate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.Generator/Operate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).Operate(ctx, req.(*OperateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Generator_Metrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).Metrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.Generator/Metrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).Metrics(ctx, req.(*MetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Generator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "modoki.Generator",
	HandlerType: (*GeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Operate",
			Handler:    _Generator_Operate_Handler,
		},
		{
			MethodName: "Metrics",
			Handler:    _Generator_Metrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generator.proto",
}

func init() { proto.RegisterFile("generator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x8b, 0xd4, 0x30,
	0x14, 0xc0, 0x4d, 0x9d, 0xe9, 0xd8, 0x37, 0xda, 0x8e, 0x11, 0xd6, 0xb2, 0x88, 0x94, 0x2a, 0x58,
	0x15, 0x07, 0x19, 0x0f, 0x2e, 0x8b, 0x97, 0x51, 0xc1, 0xc3, 0xba, 0x08, 0xf1, 0xb4, 0x5e, 0x96,
	0x6e, 0xfb, 0x76, 0x09, 0xfd, 0x93, 0x98, 0x64, 0x84, 0x3d, 0xfb, 0x49, 0xfc, 0xa6, 0xd2, 0x34,
	0xdd, 0xae, 0xf5, 0x0f, 0x78, 0x6b, 0xdf, 0xfb, 0xbd, 0xe4, 0xf7, 0x5e, 0x12, 0x88, 0x2e, 0xb0,
	0x45, 0x95, 0x1b, 0xa1, 0xd6, 0x52, 0x09, 0x23, 0xa8, 0xdf, 0x88, 0x52, 0x54, 0x7c, 0xff, 0x8e,
	0x46, 0xf5, 0x8d, 0x17, 0xd8, 0x87, 0xd3, 0x97, 0xb0, 0x3a, 0xda, 0x9d, 0xa1, 0x6a, 0xd1, 0xa0,
	0x7e, 0x27, 0xda, 0x73, 0x7e, 0x41, 0x1f, 0x40, 0xd0, 0xe6, 0x0d, 0x6a, 0x99, 0x17, 0x18, 0x93,
	0x84, 0x64, 0x01, 0x1b, 0x03, 0xe9, 0x43, 0x98, 0x9d, 0x6c, 0x8f, 0x3f, 0xd2, 0x3d, 0xf0, 0x0b,
	0xcb, 0x3b, 0xc4, 0xfd, 0xa5, 0x3f, 0x3c, 0x08, 0x3f, 0xc9, 0x6e, 0x6f, 0x64, 0xf8, 0x75, 0x87,
	0xda, 0xd0, 0x10, 0x3c, 0x5e, 0x3a, 0xcc, 0xe3, 0x25, 0x7d, 0x02, 0xb3, 0x8a, 0xb7, 0x65, 0xec,
	0x25, 0x24, 0x0b, 0x37, 0xf7, 0xd6, 0xbd, 0xda, 0xda, 0x55, 0x1d, 0xf1, 0xb6, 0x64, 0x16, 0xe8,
	0x4c, 0x24, 0xaa, 0x73, 0xa1, 0x1a, 0x54, 0xf1, 0xcd, 0x84, 0x64, 0x73, 0x36, 0x06, 0xe8, 0x23,
	0x98, 0x69, 0x89, 0x45, 0x3c, 0x4b, 0x48, 0xb6, 0xdc, 0x44, 0xc3, 0x32, 0x5b, 0x29, 0x3f, 0x4b,
	0x2c, 0x98, 0x4d, 0xd2, 0x17, 0xb0, 0x2c, 0xb1, 0x46, 0x83, 0xa7, 0x97, 0x79, 0x53, 0xc7, 0x73,
	0xcb, 0xde, 0x1e, 0xd8, 0xae, 0x13, 0x06, 0x3d, 0x70, 0x92, 0x37, 0x35, 0x7d, 0x0e, 0x90, 0x4b,
	0x59, 0x5f, 0xf6, 0xb4, 0xff, 0x07, 0x3a, 0xb0, 0x79, 0x0b, 0xbf, 0x06, 0xa8, 0x0e, 0xf4, 0xa9,
	0x1b, 0xc3, 0xc2, 0xc2, 0xf1, 0x00, 0x4f, 0xc7, 0xca, 0x82, 0xea, 0xc0, 0x7d, 0xa6, 0x0d, 0x44,
	0x57, 0x23, 0xd2, 0x52, 0xb4, 0x1a, 0xa7, 0x9e, 0xe4, 0xbf, 0x3c, 0xbd, 0x7f, 0x7a, 0xa6, 0x09,
	0x84, 0xc7, 0x68, 0x14, 0x2f, 0xf4, 0x5f, 0x4e, 0x24, 0x7d, 0x03, 0xd1, 0x15, 0xe1, 0x84, 0x9e,
	0x82, 0xaf, 0x4d, 0x6e, 0x76, 0xda, 0xb9, 0xdc, 0xbd, 0x3e, 0x5f, 0x9b, 0x60, 0x0e, 0x78, 0xf6,
	0x18, 0x96, 0xd7, 0xce, 0x8e, 0x06, 0x30, 0xdf, 0x76, 0x7b, 0xaf, 0x6e, 0x50, 0x00, 0xff, 0xbd,
	0x95, 0x5e, 0x91, 0xcd, 0x77, 0x02, 0xc1, 0x87, 0xe1, 0x56, 0xd2, 0x43, 0x58, 0xb8, 0x1a, 0xba,
	0x37, 0xb9, 0x00, 0x4e, 0x72, 0xff, 0xfe, 0x6f, 0x71, 0xa7, 0x76, 0x08, 0x0b, 0x67, 0x3b, 0xd6,
	0xfe, 0xda, 0xe0, 0x58, 0x3b, 0x69, 0xeb, 0xed, 0xad, 0x2f, 0xee, 0x25, 0x9c, 0xf9, 0xf6, 0x05,
	0xbc, 0xfa, 0x19, 0x00, 0x00, 0xff, 0xff, 0x27, 0x1b, 0x4a, 0x86, 0x2b, 0x03, 0x00, 0x00,
}
