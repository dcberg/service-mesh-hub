// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service-mesh-hub/services/apiserver/api/v1/workloads.proto

package v1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	types1 "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	types "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MeshWorkload struct {
	Spec                 *types.MeshWorkloadSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Ref                  *types1.ResourceRef     `protobuf:"bytes,3,opt,name=ref,proto3" json:"ref,omitempty"`
	Labels               map[string]string       `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MeshWorkload) Reset()         { *m = MeshWorkload{} }
func (m *MeshWorkload) String() string { return proto.CompactTextString(m) }
func (*MeshWorkload) ProtoMessage()    {}
func (*MeshWorkload) Descriptor() ([]byte, []int) {
	return fileDescriptor_95215d2ce5e56efe, []int{0}
}
func (m *MeshWorkload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshWorkload.Unmarshal(m, b)
}
func (m *MeshWorkload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshWorkload.Marshal(b, m, deterministic)
}
func (m *MeshWorkload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshWorkload.Merge(m, src)
}
func (m *MeshWorkload) XXX_Size() int {
	return xxx_messageInfo_MeshWorkload.Size(m)
}
func (m *MeshWorkload) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshWorkload.DiscardUnknown(m)
}

var xxx_messageInfo_MeshWorkload proto.InternalMessageInfo

func (m *MeshWorkload) GetSpec() *types.MeshWorkloadSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *MeshWorkload) GetRef() *types1.ResourceRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *MeshWorkload) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type ListMeshWorkloadsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMeshWorkloadsRequest) Reset()         { *m = ListMeshWorkloadsRequest{} }
func (m *ListMeshWorkloadsRequest) String() string { return proto.CompactTextString(m) }
func (*ListMeshWorkloadsRequest) ProtoMessage()    {}
func (*ListMeshWorkloadsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_95215d2ce5e56efe, []int{1}
}
func (m *ListMeshWorkloadsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMeshWorkloadsRequest.Unmarshal(m, b)
}
func (m *ListMeshWorkloadsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMeshWorkloadsRequest.Marshal(b, m, deterministic)
}
func (m *ListMeshWorkloadsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMeshWorkloadsRequest.Merge(m, src)
}
func (m *ListMeshWorkloadsRequest) XXX_Size() int {
	return xxx_messageInfo_ListMeshWorkloadsRequest.Size(m)
}
func (m *ListMeshWorkloadsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMeshWorkloadsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMeshWorkloadsRequest proto.InternalMessageInfo

type ListMeshWorkloadsResponse struct {
	MeshWorkloads        []*MeshWorkload `protobuf:"bytes,1,rep,name=mesh_workloads,json=meshWorkloads,proto3" json:"mesh_workloads,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListMeshWorkloadsResponse) Reset()         { *m = ListMeshWorkloadsResponse{} }
func (m *ListMeshWorkloadsResponse) String() string { return proto.CompactTextString(m) }
func (*ListMeshWorkloadsResponse) ProtoMessage()    {}
func (*ListMeshWorkloadsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_95215d2ce5e56efe, []int{2}
}
func (m *ListMeshWorkloadsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMeshWorkloadsResponse.Unmarshal(m, b)
}
func (m *ListMeshWorkloadsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMeshWorkloadsResponse.Marshal(b, m, deterministic)
}
func (m *ListMeshWorkloadsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMeshWorkloadsResponse.Merge(m, src)
}
func (m *ListMeshWorkloadsResponse) XXX_Size() int {
	return xxx_messageInfo_ListMeshWorkloadsResponse.Size(m)
}
func (m *ListMeshWorkloadsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMeshWorkloadsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMeshWorkloadsResponse proto.InternalMessageInfo

func (m *ListMeshWorkloadsResponse) GetMeshWorkloads() []*MeshWorkload {
	if m != nil {
		return m.MeshWorkloads
	}
	return nil
}

func init() {
	proto.RegisterType((*MeshWorkload)(nil), "grpc.solo.io.MeshWorkload")
	proto.RegisterMapType((map[string]string)(nil), "grpc.solo.io.MeshWorkload.LabelsEntry")
	proto.RegisterType((*ListMeshWorkloadsRequest)(nil), "grpc.solo.io.ListMeshWorkloadsRequest")
	proto.RegisterType((*ListMeshWorkloadsResponse)(nil), "grpc.solo.io.ListMeshWorkloadsResponse")
}

func init() {
	proto.RegisterFile("service-mesh-hub/services/apiserver/api/v1/workloads.proto", fileDescriptor_95215d2ce5e56efe)
}

var fileDescriptor_95215d2ce5e56efe = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6f, 0xd4, 0x30,
	0x10, 0xc5, 0x49, 0xb2, 0x54, 0xe0, 0x2d, 0x50, 0x2c, 0x0e, 0x21, 0xa7, 0x55, 0x0e, 0x50, 0x21,
	0xd6, 0xd6, 0x86, 0x0b, 0x14, 0xa9, 0xd2, 0x82, 0xb8, 0xa0, 0x72, 0x31, 0x07, 0x24, 0x0e, 0xa0,
	0xc4, 0x9d, 0x6c, 0xac, 0xcd, 0xd6, 0xc6, 0x93, 0x04, 0x85, 0x2b, 0x5f, 0x1c, 0x39, 0x49, 0x21,
	0xdd, 0x3f, 0xa2, 0xb7, 0xc9, 0xcc, 0xfc, 0xde, 0x9b, 0x71, 0x86, 0x9c, 0x21, 0xd8, 0x46, 0x49,
	0x98, 0x6f, 0x00, 0x8b, 0x79, 0x51, 0x67, 0x7c, 0x48, 0x20, 0x4f, 0x8d, 0x72, 0x31, 0x58, 0x17,
	0xf1, 0x66, 0xc1, 0x7f, 0x6a, 0xbb, 0x2e, 0x75, 0x7a, 0x89, 0xcc, 0x58, 0x5d, 0x69, 0x7a, 0xbc,
	0xb2, 0x46, 0x32, 0xd4, 0xa5, 0x66, 0x4a, 0x47, 0x6f, 0x77, 0x94, 0x1c, 0x76, 0xa9, 0x50, 0xea,
	0x06, 0x6c, 0xcb, 0x9b, 0x45, 0x5a, 0x9a, 0x22, 0x5d, 0x70, 0x57, 0xff, 0x7e, 0xad, 0xd5, 0x4b,
	0x45, 0x2f, 0xf7, 0xc2, 0x52, 0x5b, 0xf8, 0xc7, 0x59, 0xc8, 0xfb, 0xee, 0xf8, 0xb7, 0x4f, 0x8e,
	0x3f, 0x01, 0x16, 0x5f, 0x06, 0x11, 0x7a, 0x4e, 0x26, 0x68, 0x40, 0x86, 0xde, 0xcc, 0x3b, 0x9d,
	0x26, 0x2f, 0xd8, 0x5f, 0x57, 0xf6, 0x0b, 0x4c, 0xd1, 0xda, 0xeb, 0x21, 0xd9, 0x98, 0xfa, 0x6c,
	0x40, 0x8a, 0x8e, 0xa3, 0x09, 0x09, 0x2c, 0xe4, 0x61, 0xd0, 0xe1, 0x33, 0xe6, 0x7c, 0xb7, 0x49,
	0x01, 0xa8, 0x6b, 0x2b, 0x41, 0x40, 0x2e, 0x5c, 0x33, 0x3d, 0x27, 0x47, 0x65, 0x9a, 0x41, 0x89,
	0xe1, 0x64, 0x16, 0x9c, 0x4e, 0x93, 0x67, 0x6c, 0xfc, 0x1c, 0x37, 0x9c, 0xd8, 0x45, 0xd7, 0xf8,
	0xe1, 0xaa, 0xb2, 0xad, 0x18, 0xa8, 0xe8, 0x0d, 0x99, 0x8e, 0xd2, 0xf4, 0x84, 0x04, 0x6b, 0x68,
	0xbb, 0x0d, 0xee, 0x0b, 0x17, 0xd2, 0x27, 0xe4, 0x6e, 0x93, 0x96, 0x35, 0x84, 0x7e, 0x97, 0xeb,
	0x3f, 0xce, 0xfc, 0xd7, 0xde, 0xc7, 0xc9, 0x3d, 0xff, 0x24, 0x88, 0x23, 0x12, 0x5e, 0x28, 0xac,
	0xc6, 0x46, 0x28, 0xe0, 0x47, 0x0d, 0x58, 0xc5, 0xdf, 0xc8, 0xd3, 0x3d, 0x35, 0x34, 0xfa, 0x0a,
	0x81, 0x2e, 0xc9, 0xc3, 0x1b, 0xff, 0x00, 0x43, 0xaf, 0xdb, 0x20, 0x3a, 0xbc, 0x81, 0x78, 0xb0,
	0x19, 0x4b, 0x25, 0x2d, 0x79, 0x34, 0x2e, 0x2f, 0x8d, 0xa2, 0x39, 0x79, 0xbc, 0x63, 0x49, 0xb7,
	0x1e, 0xe5, 0xd0, 0xbc, 0xd1, 0xf3, 0xff, 0xf6, 0xf5, 0xb3, 0xc7, 0x77, 0xde, 0xbd, 0xff, 0xba,
	0x5c, 0xa9, 0xaa, 0xa8, 0x33, 0x26, 0xf5, 0x86, 0x3b, 0x62, 0xae, 0x34, 0xbf, 0xcd, 0x19, 0x9b,
	0xf5, 0x6a, 0x38, 0xe5, 0xec, 0xa8, 0x3b, 0xa4, 0x57, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x61,
	0x42, 0x5d, 0xb2, 0xff, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MeshWorkloadApiClient is the client API for MeshWorkloadApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MeshWorkloadApiClient interface {
	ListMeshWorkloads(ctx context.Context, in *ListMeshWorkloadsRequest, opts ...grpc.CallOption) (*ListMeshWorkloadsResponse, error)
}

type meshWorkloadApiClient struct {
	cc *grpc.ClientConn
}

func NewMeshWorkloadApiClient(cc *grpc.ClientConn) MeshWorkloadApiClient {
	return &meshWorkloadApiClient{cc}
}

func (c *meshWorkloadApiClient) ListMeshWorkloads(ctx context.Context, in *ListMeshWorkloadsRequest, opts ...grpc.CallOption) (*ListMeshWorkloadsResponse, error) {
	out := new(ListMeshWorkloadsResponse)
	err := c.cc.Invoke(ctx, "/grpc.solo.io.MeshWorkloadApi/ListMeshWorkloads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeshWorkloadApiServer is the server API for MeshWorkloadApi service.
type MeshWorkloadApiServer interface {
	ListMeshWorkloads(context.Context, *ListMeshWorkloadsRequest) (*ListMeshWorkloadsResponse, error)
}

// UnimplementedMeshWorkloadApiServer can be embedded to have forward compatible implementations.
type UnimplementedMeshWorkloadApiServer struct {
}

func (*UnimplementedMeshWorkloadApiServer) ListMeshWorkloads(ctx context.Context, req *ListMeshWorkloadsRequest) (*ListMeshWorkloadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMeshWorkloads not implemented")
}

func RegisterMeshWorkloadApiServer(s *grpc.Server, srv MeshWorkloadApiServer) {
	s.RegisterService(&_MeshWorkloadApi_serviceDesc, srv)
}

func _MeshWorkloadApi_ListMeshWorkloads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMeshWorkloadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshWorkloadApiServer).ListMeshWorkloads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.solo.io.MeshWorkloadApi/ListMeshWorkloads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshWorkloadApiServer).ListMeshWorkloads(ctx, req.(*ListMeshWorkloadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MeshWorkloadApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.solo.io.MeshWorkloadApi",
	HandlerType: (*MeshWorkloadApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMeshWorkloads",
			Handler:    _MeshWorkloadApi_ListMeshWorkloads_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service-mesh-hub/services/apiserver/api/v1/workloads.proto",
}
