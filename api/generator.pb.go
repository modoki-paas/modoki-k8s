// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: generator.proto

package modoki

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type OperateKind int32

const (
	OperateKind_Apply  OperateKind = 0
	OperateKind_Delete OperateKind = 1
)

// Enum value maps for OperateKind.
var (
	OperateKind_name = map[int32]string{
		0: "Apply",
		1: "Delete",
	}
	OperateKind_value = map[string]int32{
		"Apply":  0,
		"Delete": 1,
	}
)

func (x OperateKind) Enum() *OperateKind {
	p := new(OperateKind)
	*p = x
	return p
}

func (x OperateKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperateKind) Descriptor() protoreflect.EnumDescriptor {
	return file_generator_proto_enumTypes[0].Descriptor()
}

func (OperateKind) Type() protoreflect.EnumType {
	return &file_generator_proto_enumTypes[0]
}

func (x OperateKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperateKind.Descriptor instead.
func (OperateKind) EnumDescriptor() ([]byte, []int) {
	return file_generator_proto_rawDescGZIP(), []int{0}
}

type KubernetesConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *KubernetesConfig) Reset() {
	*x = KubernetesConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesConfig) ProtoMessage() {}

func (x *KubernetesConfig) ProtoReflect() protoreflect.Message {
	mi := &file_generator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesConfig.ProtoReflect.Descriptor instead.
func (*KubernetesConfig) Descriptor() ([]byte, []int) {
	return file_generator_proto_rawDescGZIP(), []int{0}
}

func (x *KubernetesConfig) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type YAML struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config string `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *YAML) Reset() {
	*x = YAML{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YAML) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YAML) ProtoMessage() {}

func (x *YAML) ProtoReflect() protoreflect.Message {
	mi := &file_generator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YAML.ProtoReflect.Descriptor instead.
func (*YAML) Descriptor() ([]byte, []int) {
	return file_generator_proto_rawDescGZIP(), []int{1}
}

func (x *YAML) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

type OperateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Domain    string            `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Kind      OperateKind       `protobuf:"varint,3,opt,name=kind,proto3,enum=modoki.OperateKind" json:"kind,omitempty"`
	Spec      *AppSpec          `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	Status    *AppStatus        `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Yaml      *YAML             `protobuf:"bytes,6,opt,name=yaml,proto3" json:"yaml,omitempty"`
	K8SConfig *KubernetesConfig `protobuf:"bytes,7,opt,name=k8s_config,json=k8sConfig,proto3" json:"k8s_config,omitempty"`
}

func (x *OperateRequest) Reset() {
	*x = OperateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateRequest) ProtoMessage() {}

func (x *OperateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_generator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateRequest.ProtoReflect.Descriptor instead.
func (*OperateRequest) Descriptor() ([]byte, []int) {
	return file_generator_proto_rawDescGZIP(), []int{2}
}

func (x *OperateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OperateRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *OperateRequest) GetKind() OperateKind {
	if x != nil {
		return x.Kind
	}
	return OperateKind_Apply
}

func (x *OperateRequest) GetSpec() *AppSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *OperateRequest) GetStatus() *AppStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *OperateRequest) GetYaml() *YAML {
	if x != nil {
		return x.Yaml
	}
	return nil
}

func (x *OperateRequest) GetK8SConfig() *KubernetesConfig {
	if x != nil {
		return x.K8SConfig
	}
	return nil
}

type OperateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yaml   *YAML      `protobuf:"bytes,1,opt,name=yaml,proto3" json:"yaml,omitempty"`
	Status *AppStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OperateResponse) Reset() {
	*x = OperateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateResponse) ProtoMessage() {}

func (x *OperateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_generator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateResponse.ProtoReflect.Descriptor instead.
func (*OperateResponse) Descriptor() ([]byte, []int) {
	return file_generator_proto_rawDescGZIP(), []int{3}
}

func (x *OperateResponse) GetYaml() *YAML {
	if x != nil {
		return x.Yaml
	}
	return nil
}

func (x *OperateResponse) GetStatus() *AppStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type MetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    *AppStatus        `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	K8SConfig *KubernetesConfig `protobuf:"bytes,2,opt,name=k8s_config,json=k8sConfig,proto3" json:"k8s_config,omitempty"`
}

func (x *MetricsRequest) Reset() {
	*x = MetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsRequest) ProtoMessage() {}

func (x *MetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_generator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsRequest.ProtoReflect.Descriptor instead.
func (*MetricsRequest) Descriptor() ([]byte, []int) {
	return file_generator_proto_rawDescGZIP(), []int{4}
}

func (x *MetricsRequest) GetStatus() *AppStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *MetricsRequest) GetK8SConfig() *KubernetesConfig {
	if x != nil {
		return x.K8SConfig
	}
	return nil
}

type MetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *AppStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MetricsResponse) Reset() {
	*x = MetricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsResponse) ProtoMessage() {}

func (x *MetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_generator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsResponse.ProtoReflect.Descriptor instead.
func (*MetricsResponse) Descriptor() ([]byte, []int) {
	return file_generator_proto_rawDescGZIP(), []int{5}
}

func (x *MetricsResponse) GetStatus() *AppStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_generator_proto protoreflect.FileDescriptor

var file_generator_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x6f, 0x6b, 0x69, 0x1a, 0x09, 0x61, 0x70, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x10, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x1e, 0x0a, 0x04, 0x59, 0x41, 0x4d, 0x4c, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x8c, 0x02, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x27, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x6f, 0x6b, 0x69, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x6f, 0x6b,
	0x69, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12,
	0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x6f, 0x6b, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x79, 0x61,
	0x6d, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x6f, 0x6b,
	0x69, 0x2e, 0x59, 0x41, 0x4d, 0x4c, 0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x12, 0x37, 0x0a, 0x0a,
	0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x6f, 0x6b, 0x69, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x6b, 0x38, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x5e, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x79, 0x61, 0x6d, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x6f, 0x6b, 0x69, 0x2e,
	0x59, 0x41, 0x4d, 0x4c, 0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64,
	0x6f, 0x6b, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x74, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x6f, 0x6b, 0x69,
	0x2e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x6f, 0x6b, 0x69, 0x2e,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x09, 0x6b, 0x38, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x3c, 0x0a, 0x0f, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x6d, 0x6f, 0x64, 0x6f, 0x6b, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x24, 0x0a, 0x0b, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x01, 0x32,
	0x83, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x3a, 0x0a,
	0x07, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x6f, 0x6b,
	0x69, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x6f, 0x6b, 0x69, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x12, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x6f, 0x6b, 0x69, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d,
	0x6f, 0x64, 0x6f, 0x6b, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x6f, 0x6b, 0x69, 0x2d, 0x70, 0x61, 0x61, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x6f, 0x6b, 0x69, 0x2d, 0x6b, 0x38, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x6d,
	0x6f, 0x64, 0x6f, 0x6b, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_generator_proto_rawDescOnce sync.Once
	file_generator_proto_rawDescData = file_generator_proto_rawDesc
)

func file_generator_proto_rawDescGZIP() []byte {
	file_generator_proto_rawDescOnce.Do(func() {
		file_generator_proto_rawDescData = protoimpl.X.CompressGZIP(file_generator_proto_rawDescData)
	})
	return file_generator_proto_rawDescData
}

var file_generator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_generator_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_generator_proto_goTypes = []interface{}{
	(OperateKind)(0),         // 0: modoki.OperateKind
	(*KubernetesConfig)(nil), // 1: modoki.KubernetesConfig
	(*YAML)(nil),             // 2: modoki.YAML
	(*OperateRequest)(nil),   // 3: modoki.OperateRequest
	(*OperateResponse)(nil),  // 4: modoki.OperateResponse
	(*MetricsRequest)(nil),   // 5: modoki.MetricsRequest
	(*MetricsResponse)(nil),  // 6: modoki.MetricsResponse
	(*AppSpec)(nil),          // 7: modoki.AppSpec
	(*AppStatus)(nil),        // 8: modoki.AppStatus
}
var file_generator_proto_depIdxs = []int32{
	0,  // 0: modoki.OperateRequest.kind:type_name -> modoki.OperateKind
	7,  // 1: modoki.OperateRequest.spec:type_name -> modoki.AppSpec
	8,  // 2: modoki.OperateRequest.status:type_name -> modoki.AppStatus
	2,  // 3: modoki.OperateRequest.yaml:type_name -> modoki.YAML
	1,  // 4: modoki.OperateRequest.k8s_config:type_name -> modoki.KubernetesConfig
	2,  // 5: modoki.OperateResponse.yaml:type_name -> modoki.YAML
	8,  // 6: modoki.OperateResponse.status:type_name -> modoki.AppStatus
	8,  // 7: modoki.MetricsRequest.status:type_name -> modoki.AppStatus
	1,  // 8: modoki.MetricsRequest.k8s_config:type_name -> modoki.KubernetesConfig
	8,  // 9: modoki.MetricsResponse.status:type_name -> modoki.AppStatus
	3,  // 10: modoki.Generator.Operate:input_type -> modoki.OperateRequest
	5,  // 11: modoki.Generator.Metrics:input_type -> modoki.MetricsRequest
	4,  // 12: modoki.Generator.Operate:output_type -> modoki.OperateResponse
	6,  // 13: modoki.Generator.Metrics:output_type -> modoki.MetricsResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_generator_proto_init() }
func file_generator_proto_init() {
	if File_generator_proto != nil {
		return
	}
	file_app_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_generator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesConfig); i {
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
		file_generator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YAML); i {
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
		file_generator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateRequest); i {
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
		file_generator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateResponse); i {
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
		file_generator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsRequest); i {
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
		file_generator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsResponse); i {
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
			RawDescriptor: file_generator_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_generator_proto_goTypes,
		DependencyIndexes: file_generator_proto_depIdxs,
		EnumInfos:         file_generator_proto_enumTypes,
		MessageInfos:      file_generator_proto_msgTypes,
	}.Build()
	File_generator_proto = out.File
	file_generator_proto_rawDesc = nil
	file_generator_proto_goTypes = nil
	file_generator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GeneratorClient is the client API for Generator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeneratorClient interface {
	Operate(ctx context.Context, in *OperateRequest, opts ...grpc.CallOption) (*OperateResponse, error)
	Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error)
}

type generatorClient struct {
	cc grpc.ClientConnInterface
}

func NewGeneratorClient(cc grpc.ClientConnInterface) GeneratorClient {
	return &generatorClient{cc}
}

func (c *generatorClient) Operate(ctx context.Context, in *OperateRequest, opts ...grpc.CallOption) (*OperateResponse, error) {
	out := new(OperateResponse)
	err := c.cc.Invoke(ctx, "/modoki.Generator/Operate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorClient) Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, "/modoki.Generator/Metrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeneratorServer is the server API for Generator service.
type GeneratorServer interface {
	Operate(context.Context, *OperateRequest) (*OperateResponse, error)
	Metrics(context.Context, *MetricsRequest) (*MetricsResponse, error)
}

// UnimplementedGeneratorServer can be embedded to have forward compatible implementations.
type UnimplementedGeneratorServer struct {
}

func (*UnimplementedGeneratorServer) Operate(context.Context, *OperateRequest) (*OperateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Operate not implemented")
}
func (*UnimplementedGeneratorServer) Metrics(context.Context, *MetricsRequest) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Metrics not implemented")
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
