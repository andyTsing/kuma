// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: mesh/v1alpha1/timeout.proto

package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type Timeout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of selectors to match dataplanes that are sources of traffic.
	Sources []*Selector `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	// List of selectors to match services that are destinations of traffic.
	Destinations []*Selector   `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
	Conf         *Timeout_Conf `protobuf:"bytes,3,opt,name=conf,proto3" json:"conf,omitempty"`
}

func (x *Timeout) Reset() {
	*x = Timeout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_timeout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timeout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timeout) ProtoMessage() {}

func (x *Timeout) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_timeout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timeout.ProtoReflect.Descriptor instead.
func (*Timeout) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_timeout_proto_rawDescGZIP(), []int{0}
}

func (x *Timeout) GetSources() []*Selector {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *Timeout) GetDestinations() []*Selector {
	if x != nil {
		return x.Destinations
	}
	return nil
}

func (x *Timeout) GetConf() *Timeout_Conf {
	if x != nil {
		return x.Conf
	}
	return nil
}

type Timeout_Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ConnectTimeout defines time to establish connection
	ConnectTimeout *duration.Duration `protobuf:"bytes,1,opt,name=connect_timeout,json=connectTimeout,proto3" json:"connect_timeout,omitempty"`
	Tcp            *Timeout_Conf_Tcp  `protobuf:"bytes,2,opt,name=tcp,proto3" json:"tcp,omitempty"`
	Http           *Timeout_Conf_Http `protobuf:"bytes,3,opt,name=http,proto3" json:"http,omitempty"`
	Grpc           *Timeout_Conf_Grpc `protobuf:"bytes,4,opt,name=grpc,proto3" json:"grpc,omitempty"`
}

func (x *Timeout_Conf) Reset() {
	*x = Timeout_Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_timeout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timeout_Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timeout_Conf) ProtoMessage() {}

func (x *Timeout_Conf) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_timeout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timeout_Conf.ProtoReflect.Descriptor instead.
func (*Timeout_Conf) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_timeout_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Timeout_Conf) GetConnectTimeout() *duration.Duration {
	if x != nil {
		return x.ConnectTimeout
	}
	return nil
}

func (x *Timeout_Conf) GetTcp() *Timeout_Conf_Tcp {
	if x != nil {
		return x.Tcp
	}
	return nil
}

func (x *Timeout_Conf) GetHttp() *Timeout_Conf_Http {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Timeout_Conf) GetGrpc() *Timeout_Conf_Grpc {
	if x != nil {
		return x.Grpc
	}
	return nil
}

// Tcp defines timeouts that are applied when the protocol is TCP
type Timeout_Conf_Tcp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IdleTimeout is defined as the period in which there are no bytes sent or received on either the upstream or downstream connection
	IdleTimeout *duration.Duration `protobuf:"bytes,1,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
}

func (x *Timeout_Conf_Tcp) Reset() {
	*x = Timeout_Conf_Tcp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_timeout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timeout_Conf_Tcp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timeout_Conf_Tcp) ProtoMessage() {}

func (x *Timeout_Conf_Tcp) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_timeout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timeout_Conf_Tcp.ProtoReflect.Descriptor instead.
func (*Timeout_Conf_Tcp) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_timeout_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Timeout_Conf_Tcp) GetIdleTimeout() *duration.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

// Http defines timeouts that are applied when the protocol is HTTP
type Timeout_Conf_Http struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RequestTimeout is a span between the point at which the entire downstream request (i.e. end-of-stream) has been processed and when the upstream response has been completely processed
	RequestTimeout *duration.Duration `protobuf:"bytes,1,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	// IdleTimeout is the time at which a downstream or upstream connection will be terminated if there are no active streams
	IdleTimeout *duration.Duration `protobuf:"bytes,2,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
}

func (x *Timeout_Conf_Http) Reset() {
	*x = Timeout_Conf_Http{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_timeout_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timeout_Conf_Http) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timeout_Conf_Http) ProtoMessage() {}

func (x *Timeout_Conf_Http) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_timeout_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timeout_Conf_Http.ProtoReflect.Descriptor instead.
func (*Timeout_Conf_Http) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_timeout_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *Timeout_Conf_Http) GetRequestTimeout() *duration.Duration {
	if x != nil {
		return x.RequestTimeout
	}
	return nil
}

func (x *Timeout_Conf_Http) GetIdleTimeout() *duration.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

// Grpc defines timeouts that are applied when the protocol is GRPC
type Timeout_Conf_Grpc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// StreamIdleTimeout is the amount of time that the connection manager will allow a stream to exist with no upstream or downstream activity
	StreamIdleTimeout *duration.Duration `protobuf:"bytes,1,opt,name=stream_idle_timeout,json=streamIdleTimeout,proto3" json:"stream_idle_timeout,omitempty"`
	// MaxStreamDuration is the maximum time that a stream’s lifetime will span
	MaxStreamDuration *duration.Duration `protobuf:"bytes,2,opt,name=max_stream_duration,json=maxStreamDuration,proto3" json:"max_stream_duration,omitempty"`
}

func (x *Timeout_Conf_Grpc) Reset() {
	*x = Timeout_Conf_Grpc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_timeout_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timeout_Conf_Grpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timeout_Conf_Grpc) ProtoMessage() {}

func (x *Timeout_Conf_Grpc) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_timeout_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timeout_Conf_Grpc.ProtoReflect.Descriptor instead.
func (*Timeout_Conf_Grpc) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_timeout_proto_rawDescGZIP(), []int{0, 0, 2}
}

func (x *Timeout_Conf_Grpc) GetStreamIdleTimeout() *duration.Duration {
	if x != nil {
		return x.StreamIdleTimeout
	}
	return nil
}

func (x *Timeout_Conf_Grpc) GetMaxStreamDuration() *duration.Duration {
	if x != nil {
		return x.MaxStreamDuration
	}
	return nil
}

var File_mesh_v1alpha1_timeout_proto protoreflect.FileDescriptor

var file_mesh_v1alpha1_timeout_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6b,
	0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x1a, 0x1c, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa3, 0x06, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b,
	0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x75, 0x6d, 0x61,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0xe7, 0x04, 0x0a, 0x04,
	0x43, 0x6f, 0x6e, 0x66, 0x12, 0x42, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x03, 0x74, 0x63, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x54, 0x63, 0x70, 0x52, 0x03, 0x74, 0x63, 0x70,
	0x12, 0x39, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x39, 0x0a, 0x04, 0x67,
	0x72, 0x70, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6b, 0x75, 0x6d, 0x61,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x47, 0x72, 0x70, 0x63,
	0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x43, 0x0a, 0x03, 0x54, 0x63, 0x70, 0x12, 0x3c, 0x0a,
	0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x69, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x88, 0x01, 0x0a, 0x04,
	0x48, 0x74, 0x74, 0x70, 0x12, 0x42, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x9c, 0x01, 0x0a, 0x04, 0x47, 0x72, 0x70, 0x63, 0x12,
	0x49, 0x0a, 0x13, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x49, 0x0a, 0x13, 0x6d, 0x61,
	0x78, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x68, 0x71, 0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mesh_v1alpha1_timeout_proto_rawDescOnce sync.Once
	file_mesh_v1alpha1_timeout_proto_rawDescData = file_mesh_v1alpha1_timeout_proto_rawDesc
)

func file_mesh_v1alpha1_timeout_proto_rawDescGZIP() []byte {
	file_mesh_v1alpha1_timeout_proto_rawDescOnce.Do(func() {
		file_mesh_v1alpha1_timeout_proto_rawDescData = protoimpl.X.CompressGZIP(file_mesh_v1alpha1_timeout_proto_rawDescData)
	})
	return file_mesh_v1alpha1_timeout_proto_rawDescData
}

var file_mesh_v1alpha1_timeout_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mesh_v1alpha1_timeout_proto_goTypes = []interface{}{
	(*Timeout)(nil),           // 0: kuma.mesh.v1alpha1.Timeout
	(*Timeout_Conf)(nil),      // 1: kuma.mesh.v1alpha1.Timeout.Conf
	(*Timeout_Conf_Tcp)(nil),  // 2: kuma.mesh.v1alpha1.Timeout.Conf.Tcp
	(*Timeout_Conf_Http)(nil), // 3: kuma.mesh.v1alpha1.Timeout.Conf.Http
	(*Timeout_Conf_Grpc)(nil), // 4: kuma.mesh.v1alpha1.Timeout.Conf.Grpc
	(*Selector)(nil),          // 5: kuma.mesh.v1alpha1.Selector
	(*duration.Duration)(nil), // 6: google.protobuf.Duration
}
var file_mesh_v1alpha1_timeout_proto_depIdxs = []int32{
	5,  // 0: kuma.mesh.v1alpha1.Timeout.sources:type_name -> kuma.mesh.v1alpha1.Selector
	5,  // 1: kuma.mesh.v1alpha1.Timeout.destinations:type_name -> kuma.mesh.v1alpha1.Selector
	1,  // 2: kuma.mesh.v1alpha1.Timeout.conf:type_name -> kuma.mesh.v1alpha1.Timeout.Conf
	6,  // 3: kuma.mesh.v1alpha1.Timeout.Conf.connect_timeout:type_name -> google.protobuf.Duration
	2,  // 4: kuma.mesh.v1alpha1.Timeout.Conf.tcp:type_name -> kuma.mesh.v1alpha1.Timeout.Conf.Tcp
	3,  // 5: kuma.mesh.v1alpha1.Timeout.Conf.http:type_name -> kuma.mesh.v1alpha1.Timeout.Conf.Http
	4,  // 6: kuma.mesh.v1alpha1.Timeout.Conf.grpc:type_name -> kuma.mesh.v1alpha1.Timeout.Conf.Grpc
	6,  // 7: kuma.mesh.v1alpha1.Timeout.Conf.Tcp.idle_timeout:type_name -> google.protobuf.Duration
	6,  // 8: kuma.mesh.v1alpha1.Timeout.Conf.Http.request_timeout:type_name -> google.protobuf.Duration
	6,  // 9: kuma.mesh.v1alpha1.Timeout.Conf.Http.idle_timeout:type_name -> google.protobuf.Duration
	6,  // 10: kuma.mesh.v1alpha1.Timeout.Conf.Grpc.stream_idle_timeout:type_name -> google.protobuf.Duration
	6,  // 11: kuma.mesh.v1alpha1.Timeout.Conf.Grpc.max_stream_duration:type_name -> google.protobuf.Duration
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_mesh_v1alpha1_timeout_proto_init() }
func file_mesh_v1alpha1_timeout_proto_init() {
	if File_mesh_v1alpha1_timeout_proto != nil {
		return
	}
	file_mesh_v1alpha1_selector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mesh_v1alpha1_timeout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timeout); i {
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
		file_mesh_v1alpha1_timeout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timeout_Conf); i {
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
		file_mesh_v1alpha1_timeout_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timeout_Conf_Tcp); i {
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
		file_mesh_v1alpha1_timeout_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timeout_Conf_Http); i {
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
		file_mesh_v1alpha1_timeout_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timeout_Conf_Grpc); i {
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
			RawDescriptor: file_mesh_v1alpha1_timeout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mesh_v1alpha1_timeout_proto_goTypes,
		DependencyIndexes: file_mesh_v1alpha1_timeout_proto_depIdxs,
		MessageInfos:      file_mesh_v1alpha1_timeout_proto_msgTypes,
	}.Build()
	File_mesh_v1alpha1_timeout_proto = out.File
	file_mesh_v1alpha1_timeout_proto_rawDesc = nil
	file_mesh_v1alpha1_timeout_proto_goTypes = nil
	file_mesh_v1alpha1_timeout_proto_depIdxs = nil
}