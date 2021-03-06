// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cluster.proto

package openpitrix

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/mwitkow/go-proto-validators"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Cluster struct {
	Id               *string                     `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Name             *string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description      *string                     `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	AppId            *string                     `protobuf:"bytes,4,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	AppVersion       *string                     `protobuf:"bytes,5,opt,name=app_version,json=appVersion" json:"app_version,omitempty"`
	Status           *string                     `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	TransitionStatus *string                     `protobuf:"bytes,7,opt,name=transition_status,json=transitionStatus" json:"transition_status,omitempty"`
	Created          *google_protobuf3.Timestamp `protobuf:"bytes,8,opt,name=created" json:"created,omitempty"`
	LastModified     *google_protobuf3.Timestamp `protobuf:"bytes,9,opt,name=last_modified,json=lastModified" json:"last_modified,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Cluster) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Cluster) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Cluster) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *Cluster) GetAppId() string {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return ""
}

func (m *Cluster) GetAppVersion() string {
	if m != nil && m.AppVersion != nil {
		return *m.AppVersion
	}
	return ""
}

func (m *Cluster) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *Cluster) GetTransitionStatus() string {
	if m != nil && m.TransitionStatus != nil {
		return *m.TransitionStatus
	}
	return ""
}

func (m *Cluster) GetCreated() *google_protobuf3.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Cluster) GetLastModified() *google_protobuf3.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

type Clusters struct {
	Items            []*Cluster `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Clusters) Reset()                    { *m = Clusters{} }
func (m *Clusters) String() string            { return proto.CompactTextString(m) }
func (*Clusters) ProtoMessage()               {}
func (*Clusters) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Clusters) GetItems() []*Cluster {
	if m != nil {
		return m.Items
	}
	return nil
}

type ClusterNode struct {
	Id               *string                     `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	InstanceId       *string                     `protobuf:"bytes,2,req,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	Name             *string                     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Description      *string                     `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	ClusterId        *string                     `protobuf:"bytes,5,opt,name=cluster_id,json=clusterId" json:"cluster_id,omitempty"`
	PrivateIp        *string                     `protobuf:"bytes,6,opt,name=private_ip,json=privateIp" json:"private_ip,omitempty"`
	Status           *string                     `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
	TransitionStatus *string                     `protobuf:"bytes,8,opt,name=transition_status,json=transitionStatus" json:"transition_status,omitempty"`
	Created          *google_protobuf3.Timestamp `protobuf:"bytes,9,opt,name=created" json:"created,omitempty"`
	LastModified     *google_protobuf3.Timestamp `protobuf:"bytes,10,opt,name=last_modified,json=lastModified" json:"last_modified,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *ClusterNode) Reset()                    { *m = ClusterNode{} }
func (m *ClusterNode) String() string            { return proto.CompactTextString(m) }
func (*ClusterNode) ProtoMessage()               {}
func (*ClusterNode) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *ClusterNode) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *ClusterNode) GetInstanceId() string {
	if m != nil && m.InstanceId != nil {
		return *m.InstanceId
	}
	return ""
}

func (m *ClusterNode) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ClusterNode) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *ClusterNode) GetClusterId() string {
	if m != nil && m.ClusterId != nil {
		return *m.ClusterId
	}
	return ""
}

func (m *ClusterNode) GetPrivateIp() string {
	if m != nil && m.PrivateIp != nil {
		return *m.PrivateIp
	}
	return ""
}

func (m *ClusterNode) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *ClusterNode) GetTransitionStatus() string {
	if m != nil && m.TransitionStatus != nil {
		return *m.TransitionStatus
	}
	return ""
}

func (m *ClusterNode) GetCreated() *google_protobuf3.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *ClusterNode) GetLastModified() *google_protobuf3.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

type ClusterNodes struct {
	Items            []*ClusterNode `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *ClusterNodes) Reset()                    { *m = ClusterNodes{} }
func (m *ClusterNodes) String() string            { return proto.CompactTextString(m) }
func (*ClusterNodes) ProtoMessage()               {}
func (*ClusterNodes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *ClusterNodes) GetItems() []*ClusterNode {
	if m != nil {
		return m.Items
	}
	return nil
}

type ClusterId struct {
	Id               *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClusterId) Reset()                    { *m = ClusterId{} }
func (m *ClusterId) String() string            { return proto.CompactTextString(m) }
func (*ClusterId) ProtoMessage()               {}
func (*ClusterId) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *ClusterId) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type ClusterIds struct {
	Ids              *string `protobuf:"bytes,1,req,name=ids" json:"ids,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClusterIds) Reset()                    { *m = ClusterIds{} }
func (m *ClusterIds) String() string            { return proto.CompactTextString(m) }
func (*ClusterIds) ProtoMessage()               {}
func (*ClusterIds) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *ClusterIds) GetIds() string {
	if m != nil && m.Ids != nil {
		return *m.Ids
	}
	return ""
}

type ClusterListRequest struct {
	PageSize         *int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,def=10" json:"page_size,omitempty"`
	PageNumber       *int32 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,def=1" json:"page_number,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ClusterListRequest) Reset()                    { *m = ClusterListRequest{} }
func (m *ClusterListRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterListRequest) ProtoMessage()               {}
func (*ClusterListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

const Default_ClusterListRequest_PageSize int32 = 10
const Default_ClusterListRequest_PageNumber int32 = 1

func (m *ClusterListRequest) GetPageSize() int32 {
	if m != nil && m.PageSize != nil {
		return *m.PageSize
	}
	return Default_ClusterListRequest_PageSize
}

func (m *ClusterListRequest) GetPageNumber() int32 {
	if m != nil && m.PageNumber != nil {
		return *m.PageNumber
	}
	return Default_ClusterListRequest_PageNumber
}

type ClusterListResponse struct {
	TotalItems       *int32     `protobuf:"varint,1,opt,name=total_items,json=totalItems" json:"total_items,omitempty"`
	TotalPages       *int32     `protobuf:"varint,2,opt,name=total_pages,json=totalPages" json:"total_pages,omitempty"`
	PageSize         *int32     `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	CurrentPage      *int32     `protobuf:"varint,4,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
	Items            []*Cluster `protobuf:"bytes,5,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *ClusterListResponse) Reset()                    { *m = ClusterListResponse{} }
func (m *ClusterListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterListResponse) ProtoMessage()               {}
func (*ClusterListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *ClusterListResponse) GetTotalItems() int32 {
	if m != nil && m.TotalItems != nil {
		return *m.TotalItems
	}
	return 0
}

func (m *ClusterListResponse) GetTotalPages() int32 {
	if m != nil && m.TotalPages != nil {
		return *m.TotalPages
	}
	return 0
}

func (m *ClusterListResponse) GetPageSize() int32 {
	if m != nil && m.PageSize != nil {
		return *m.PageSize
	}
	return 0
}

func (m *ClusterListResponse) GetCurrentPage() int32 {
	if m != nil && m.CurrentPage != nil {
		return *m.CurrentPage
	}
	return 0
}

func (m *ClusterListResponse) GetItems() []*Cluster {
	if m != nil {
		return m.Items
	}
	return nil
}

type ClusterNodeId struct {
	Id               *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClusterNodeId) Reset()                    { *m = ClusterNodeId{} }
func (m *ClusterNodeId) String() string            { return proto.CompactTextString(m) }
func (*ClusterNodeId) ProtoMessage()               {}
func (*ClusterNodeId) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *ClusterNodeId) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type ClusterNodeIds struct {
	Ids              *string `protobuf:"bytes,1,req,name=ids" json:"ids,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClusterNodeIds) Reset()                    { *m = ClusterNodeIds{} }
func (m *ClusterNodeIds) String() string            { return proto.CompactTextString(m) }
func (*ClusterNodeIds) ProtoMessage()               {}
func (*ClusterNodeIds) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *ClusterNodeIds) GetIds() string {
	if m != nil && m.Ids != nil {
		return *m.Ids
	}
	return ""
}

type ClusterNodeListRequest struct {
	PageSize         *int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,def=10" json:"page_size,omitempty"`
	PageNumber       *int32 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,def=1" json:"page_number,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ClusterNodeListRequest) Reset()                    { *m = ClusterNodeListRequest{} }
func (m *ClusterNodeListRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterNodeListRequest) ProtoMessage()               {}
func (*ClusterNodeListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

const Default_ClusterNodeListRequest_PageSize int32 = 10
const Default_ClusterNodeListRequest_PageNumber int32 = 1

func (m *ClusterNodeListRequest) GetPageSize() int32 {
	if m != nil && m.PageSize != nil {
		return *m.PageSize
	}
	return Default_ClusterNodeListRequest_PageSize
}

func (m *ClusterNodeListRequest) GetPageNumber() int32 {
	if m != nil && m.PageNumber != nil {
		return *m.PageNumber
	}
	return Default_ClusterNodeListRequest_PageNumber
}

type ClusterNodeListResponse struct {
	TotalItems       *int32         `protobuf:"varint,1,opt,name=total_items,json=totalItems" json:"total_items,omitempty"`
	TotalPages       *int32         `protobuf:"varint,2,opt,name=total_pages,json=totalPages" json:"total_pages,omitempty"`
	PageSize         *int32         `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	CurrentPage      *int32         `protobuf:"varint,4,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
	Items            []*ClusterNode `protobuf:"bytes,5,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *ClusterNodeListResponse) Reset()                    { *m = ClusterNodeListResponse{} }
func (m *ClusterNodeListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterNodeListResponse) ProtoMessage()               {}
func (*ClusterNodeListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

func (m *ClusterNodeListResponse) GetTotalItems() int32 {
	if m != nil && m.TotalItems != nil {
		return *m.TotalItems
	}
	return 0
}

func (m *ClusterNodeListResponse) GetTotalPages() int32 {
	if m != nil && m.TotalPages != nil {
		return *m.TotalPages
	}
	return 0
}

func (m *ClusterNodeListResponse) GetPageSize() int32 {
	if m != nil && m.PageSize != nil {
		return *m.PageSize
	}
	return 0
}

func (m *ClusterNodeListResponse) GetCurrentPage() int32 {
	if m != nil && m.CurrentPage != nil {
		return *m.CurrentPage
	}
	return 0
}

func (m *ClusterNodeListResponse) GetItems() []*ClusterNode {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Cluster)(nil), "openpitrix.Cluster")
	proto.RegisterType((*Clusters)(nil), "openpitrix.Clusters")
	proto.RegisterType((*ClusterNode)(nil), "openpitrix.ClusterNode")
	proto.RegisterType((*ClusterNodes)(nil), "openpitrix.ClusterNodes")
	proto.RegisterType((*ClusterId)(nil), "openpitrix.ClusterId")
	proto.RegisterType((*ClusterIds)(nil), "openpitrix.ClusterIds")
	proto.RegisterType((*ClusterListRequest)(nil), "openpitrix.ClusterListRequest")
	proto.RegisterType((*ClusterListResponse)(nil), "openpitrix.ClusterListResponse")
	proto.RegisterType((*ClusterNodeId)(nil), "openpitrix.ClusterNodeId")
	proto.RegisterType((*ClusterNodeIds)(nil), "openpitrix.ClusterNodeIds")
	proto.RegisterType((*ClusterNodeListRequest)(nil), "openpitrix.ClusterNodeListRequest")
	proto.RegisterType((*ClusterNodeListResponse)(nil), "openpitrix.ClusterNodeListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ClusterService service

type ClusterServiceClient interface {
	GetClusters(ctx context.Context, in *ClusterIds, opts ...grpc.CallOption) (*Clusters, error)
	GetClusterList(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListResponse, error)
	CreateCluster(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	UpdateCluster(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	DeleteClusters(ctx context.Context, in *ClusterIds, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	GetClusterNodes(ctx context.Context, in *ClusterNodeIds, opts ...grpc.CallOption) (*ClusterNodes, error)
	GetClusterNodeList(ctx context.Context, in *ClusterNodeListRequest, opts ...grpc.CallOption) (*ClusterNodeListResponse, error)
	CreateClusterNodes(ctx context.Context, in *ClusterNodes, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	UpdateClusterNode(ctx context.Context, in *ClusterNode, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	DeleteClusterNodes(ctx context.Context, in *ClusterNodeIds, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
}

type clusterServiceClient struct {
	cc *grpc.ClientConn
}

func NewClusterServiceClient(cc *grpc.ClientConn) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) GetClusters(ctx context.Context, in *ClusterIds, opts ...grpc.CallOption) (*Clusters, error) {
	out := new(Clusters)
	err := grpc.Invoke(ctx, "/openpitrix.ClusterService/GetClusters", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) GetClusterList(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListResponse, error) {
	out := new(ClusterListResponse)
	err := grpc.Invoke(ctx, "/openpitrix.ClusterService/GetClusterList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) CreateCluster(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.ClusterService/CreateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) UpdateCluster(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.ClusterService/UpdateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) DeleteClusters(ctx context.Context, in *ClusterIds, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.ClusterService/DeleteClusters", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) GetClusterNodes(ctx context.Context, in *ClusterNodeIds, opts ...grpc.CallOption) (*ClusterNodes, error) {
	out := new(ClusterNodes)
	err := grpc.Invoke(ctx, "/openpitrix.ClusterService/GetClusterNodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) GetClusterNodeList(ctx context.Context, in *ClusterNodeListRequest, opts ...grpc.CallOption) (*ClusterNodeListResponse, error) {
	out := new(ClusterNodeListResponse)
	err := grpc.Invoke(ctx, "/openpitrix.ClusterService/GetClusterNodeList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) CreateClusterNodes(ctx context.Context, in *ClusterNodes, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.ClusterService/CreateClusterNodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) UpdateClusterNode(ctx context.Context, in *ClusterNode, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.ClusterService/UpdateClusterNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) DeleteClusterNodes(ctx context.Context, in *ClusterNodeIds, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.ClusterService/DeleteClusterNodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClusterService service

type ClusterServiceServer interface {
	GetClusters(context.Context, *ClusterIds) (*Clusters, error)
	GetClusterList(context.Context, *ClusterListRequest) (*ClusterListResponse, error)
	CreateCluster(context.Context, *Cluster) (*google_protobuf2.Empty, error)
	UpdateCluster(context.Context, *Cluster) (*google_protobuf2.Empty, error)
	DeleteClusters(context.Context, *ClusterIds) (*google_protobuf2.Empty, error)
	GetClusterNodes(context.Context, *ClusterNodeIds) (*ClusterNodes, error)
	GetClusterNodeList(context.Context, *ClusterNodeListRequest) (*ClusterNodeListResponse, error)
	CreateClusterNodes(context.Context, *ClusterNodes) (*google_protobuf2.Empty, error)
	UpdateClusterNode(context.Context, *ClusterNode) (*google_protobuf2.Empty, error)
	DeleteClusterNodes(context.Context, *ClusterNodeIds) (*google_protobuf2.Empty, error)
}

func RegisterClusterServiceServer(s *grpc.Server, srv ClusterServiceServer) {
	s.RegisterService(&_ClusterService_serviceDesc, srv)
}

func _ClusterService_GetClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ClusterService/GetClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetClusters(ctx, req.(*ClusterIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_GetClusterList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetClusterList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ClusterService/GetClusterList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetClusterList(ctx, req.(*ClusterListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ClusterService/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).CreateCluster(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).UpdateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ClusterService/UpdateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).UpdateCluster(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_DeleteClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).DeleteClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ClusterService/DeleteClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).DeleteClusters(ctx, req.(*ClusterIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_GetClusterNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterNodeIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetClusterNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ClusterService/GetClusterNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetClusterNodes(ctx, req.(*ClusterNodeIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_GetClusterNodeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterNodeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetClusterNodeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ClusterService/GetClusterNodeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetClusterNodeList(ctx, req.(*ClusterNodeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_CreateClusterNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterNodes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).CreateClusterNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ClusterService/CreateClusterNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).CreateClusterNodes(ctx, req.(*ClusterNodes))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_UpdateClusterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).UpdateClusterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ClusterService/UpdateClusterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).UpdateClusterNode(ctx, req.(*ClusterNode))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_DeleteClusterNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterNodeIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).DeleteClusterNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ClusterService/DeleteClusterNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).DeleteClusterNodes(ctx, req.(*ClusterNodeIds))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClusters",
			Handler:    _ClusterService_GetClusters_Handler,
		},
		{
			MethodName: "GetClusterList",
			Handler:    _ClusterService_GetClusterList_Handler,
		},
		{
			MethodName: "CreateCluster",
			Handler:    _ClusterService_CreateCluster_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _ClusterService_UpdateCluster_Handler,
		},
		{
			MethodName: "DeleteClusters",
			Handler:    _ClusterService_DeleteClusters_Handler,
		},
		{
			MethodName: "GetClusterNodes",
			Handler:    _ClusterService_GetClusterNodes_Handler,
		},
		{
			MethodName: "GetClusterNodeList",
			Handler:    _ClusterService_GetClusterNodeList_Handler,
		},
		{
			MethodName: "CreateClusterNodes",
			Handler:    _ClusterService_CreateClusterNodes_Handler,
		},
		{
			MethodName: "UpdateClusterNode",
			Handler:    _ClusterService_UpdateClusterNode_Handler,
		},
		{
			MethodName: "DeleteClusterNodes",
			Handler:    _ClusterService_DeleteClusterNodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cluster.proto",
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 1074 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0x41, 0x53, 0x23, 0x45,
	0x14, 0xc7, 0x6b, 0x12, 0x02, 0xe4, 0x05, 0x70, 0xd3, 0x40, 0x08, 0x61, 0x35, 0xd9, 0xd1, 0x43,
	0xc8, 0x92, 0x0c, 0x50, 0xeb, 0x96, 0x62, 0x59, 0xba, 0xa0, 0x65, 0xa5, 0x4a, 0x91, 0x0a, 0xbb,
	0x5a, 0xae, 0x42, 0x6c, 0x66, 0x7a, 0x87, 0x5e, 0x93, 0x99, 0x76, 0xba, 0x03, 0xc8, 0x16, 0x07,
	0x3d, 0x7a, 0xd4, 0x6f, 0xe4, 0xd1, 0xab, 0x77, 0xa9, 0xa2, 0xfc, 0x20, 0x56, 0xf7, 0x74, 0x92,
	0xc9, 0x66, 0xc6, 0x88, 0x7a, 0xf0, 0x44, 0xe6, 0xbd, 0xff, 0xfb, 0xbf, 0xe9, 0xf7, 0xeb, 0x6e,
	0x06, 0xe6, 0xed, 0x4e, 0x8f, 0x0b, 0x12, 0x34, 0x58, 0xe0, 0x0b, 0x1f, 0x81, 0xcf, 0x88, 0xc7,
	0xa8, 0x08, 0xe8, 0x45, 0xe9, 0xae, 0xeb, 0xfb, 0x6e, 0x87, 0x58, 0x98, 0x51, 0x0b, 0x7b, 0x9e,
	0x2f, 0xb0, 0xa0, 0xbe, 0xc7, 0x43, 0x65, 0x69, 0x43, 0xfd, 0xb1, 0xeb, 0x2e, 0xf1, 0xea, 0xfc,
	0x1c, 0xbb, 0x2e, 0x09, 0x2c, 0x9f, 0x29, 0x45, 0x8c, 0x7a, 0x4d, 0x7b, 0xa9, 0xa7, 0x93, 0xde,
	0x33, 0x8b, 0x74, 0x99, 0xf8, 0x4e, 0x27, 0xcb, 0x2f, 0x27, 0x05, 0xed, 0x12, 0x2e, 0x70, 0x97,
	0x69, 0xc1, 0x43, 0x97, 0x8a, 0xd3, 0xde, 0x49, 0xc3, 0xf6, 0xbb, 0x56, 0xf7, 0x9c, 0x8a, 0x6f,
	0xfc, 0x73, 0xcb, 0xf5, 0xeb, 0x2a, 0x59, 0x3f, 0xc3, 0x1d, 0xea, 0x60, 0xe1, 0x07, 0xdc, 0x1a,
	0xfc, 0xd4, 0x75, 0xf9, 0xb1, 0x17, 0x31, 0x7f, 0x4f, 0xc1, 0xcc, 0x5e, 0xb8, 0x64, 0xb4, 0x0e,
	0x29, 0xea, 0x14, 0x8d, 0x4a, 0xaa, 0x9a, 0xdd, 0x5d, 0xbd, 0xb9, 0x2e, 0x2f, 0xc3, 0xe2, 0xb1,
	0xdd, 0xa9, 0x7f, 0x89, 0xeb, 0x97, 0x8f, 0xea, 0x4f, 0x37, 0xeb, 0x6f, 0x1f, 0xbd, 0x78, 0xeb,
	0xea, 0x8d, 0x56, 0x8a, 0x3a, 0x08, 0xc1, 0x94, 0x87, 0xbb, 0xa4, 0x98, 0xaa, 0x18, 0xd5, 0x6c,
	0x4b, 0xfd, 0x46, 0x15, 0xc8, 0x39, 0x84, 0xdb, 0x01, 0x55, 0xab, 0x2e, 0xa6, 0x55, 0x2a, 0x1a,
	0x42, 0xcb, 0x30, 0x8d, 0x19, 0x6b, 0x53, 0xa7, 0x38, 0xa5, 0x92, 0x19, 0xcc, 0x58, 0xd3, 0x41,
	0x65, 0xc8, 0xc9, 0xf0, 0x19, 0x09, 0xb8, 0x2c, 0xcc, 0xa8, 0x1c, 0x60, 0xc6, 0x3e, 0x0b, 0x23,
	0xa8, 0x00, 0xd3, 0x5c, 0x60, 0xd1, 0xe3, 0xc5, 0x69, 0x95, 0xd3, 0x4f, 0xe8, 0x3e, 0xe4, 0x45,
	0x80, 0x3d, 0x4e, 0xa5, 0x7b, 0x5b, 0x4b, 0x66, 0x94, 0xe4, 0xce, 0x30, 0x71, 0x18, 0x8a, 0x1f,
	0xc0, 0x8c, 0x1d, 0x10, 0x2c, 0x88, 0x53, 0x9c, 0xad, 0x18, 0xd5, 0xdc, 0x76, 0xa9, 0x11, 0xce,
	0xb9, 0xd1, 0x9f, 0x73, 0xe3, 0x71, 0x7f, 0xce, 0xad, 0xbe, 0x14, 0xbd, 0x07, 0xf3, 0x1d, 0xcc,
	0x45, 0xbb, 0xeb, 0x3b, 0xf4, 0x19, 0x25, 0x4e, 0x31, 0x3b, 0xb1, 0x76, 0x4e, 0x16, 0x7c, 0xa2,
	0xf5, 0xe6, 0x9b, 0x30, 0xab, 0xe7, 0xcb, 0xd1, 0x3a, 0x64, 0xa8, 0x20, 0x5d, 0x5e, 0x34, 0x2a,
	0xe9, 0x6a, 0x6e, 0x7b, 0xb1, 0x31, 0xdc, 0x5d, 0x0d, 0x2d, 0x6a, 0x85, 0x0a, 0xf3, 0xc7, 0x34,
	0xe4, 0x74, 0x68, 0xdf, 0x77, 0x08, 0xaa, 0x45, 0xd8, 0x94, 0x6e, 0xae, 0xcb, 0x05, 0x58, 0x3a,
	0xb6, 0x3b, 0x5e, 0x2c, 0x9c, 0x32, 0xe4, 0xa8, 0xc7, 0x05, 0xf6, 0x6c, 0x22, 0x67, 0x9d, 0x92,
	0x45, 0x2d, 0xe8, 0x87, 0x9a, 0x43, 0x7a, 0xe9, 0x64, 0x7a, 0x53, 0xe3, 0xf4, 0x5e, 0x05, 0xd0,
	0x87, 0x43, 0xba, 0x86, 0x94, 0xb2, 0x3a, 0xd2, 0x74, 0x64, 0x9a, 0x05, 0xf4, 0x0c, 0x0b, 0xd2,
	0xa6, 0x4c, 0x83, 0xca, 0xea, 0x48, 0x93, 0x45, 0x18, 0xce, 0x4c, 0x66, 0x38, 0x3b, 0x99, 0x61,
	0xf6, 0x5f, 0x30, 0x84, 0x5b, 0x32, 0x7c, 0x17, 0xe6, 0x22, 0x2c, 0x38, 0xaa, 0x8f, 0x72, 0x5c,
	0x89, 0xe1, 0x28, 0x85, 0x7d, 0x96, 0x0f, 0x21, 0xbb, 0x37, 0x18, 0xd3, 0xdf, 0x3f, 0x64, 0xe6,
	0x3e, 0xc0, 0xa0, 0x8e, 0xa3, 0xf7, 0x21, 0x4d, 0x1d, 0xae, 0x2b, 0x1b, 0x37, 0xd7, 0xe5, 0x1a,
	0x54, 0x8f, 0xab, 0x63, 0xa5, 0x1b, 0xeb, 0xb5, 0x18, 0x3b, 0x59, 0x6a, 0x7e, 0x01, 0x48, 0xfb,
	0x7d, 0x4c, 0xb9, 0x68, 0x91, 0x6f, 0x7b, 0x84, 0x0b, 0x54, 0x86, 0x2c, 0xc3, 0x2e, 0x69, 0x73,
	0x7a, 0x49, 0x8a, 0x46, 0xc5, 0xa8, 0x66, 0x76, 0x52, 0x5b, 0x9b, 0xad, 0x59, 0x19, 0x3c, 0xa4,
	0x97, 0x04, 0x99, 0x90, 0x53, 0x02, 0xaf, 0xd7, 0x3d, 0x21, 0x81, 0x3a, 0xf2, 0x99, 0x1d, 0x63,
	0xab, 0x05, 0x32, 0xba, 0xaf, 0x82, 0xe6, 0x2f, 0x06, 0x2c, 0x8e, 0x78, 0x73, 0xe6, 0x7b, 0x9c,
	0xc8, 0xad, 0x28, 0x7c, 0x81, 0x3b, 0xed, 0xfe, 0xbc, 0x8c, 0x6a, 0xa6, 0x05, 0x2a, 0xd4, 0x94,
	0x91, 0xa1, 0x40, 0x9a, 0xf1, 0xd0, 0x5c, 0x0b, 0x0e, 0x64, 0x04, 0xad, 0x45, 0x5f, 0x2f, 0xad,
	0xd2, 0xc3, 0x57, 0xbb, 0x07, 0x73, 0x76, 0x2f, 0x08, 0x88, 0x27, 0x54, 0xbd, 0xda, 0xb5, 0x99,
	0x56, 0x4e, 0xc7, 0xa4, 0xc1, 0xf0, 0xcc, 0x65, 0x26, 0x9e, 0xb9, 0x77, 0x60, 0x3e, 0x42, 0xaf,
	0xe9, 0xdc, 0xe6, 0xd0, 0x99, 0x8f, 0x61, 0x61, 0xa4, 0x98, 0xa3, 0xdd, 0x28, 0xb0, 0xcd, 0x9b,
	0xeb, 0xf2, 0x06, 0xd4, 0x24, 0x30, 0x2f, 0x8e, 0x98, 0x17, 0x8f, 0xec, 0x08, 0x0a, 0x11, 0xd7,
	0xff, 0x1c, 0xdb, 0xaf, 0x06, 0xac, 0x8c, 0xf9, 0xff, 0x6f, 0xd0, 0xd5, 0x47, 0xd1, 0x4d, 0x38,
	0x66, 0xdb, 0xdf, 0xcf, 0x0e, 0x10, 0x1c, 0x92, 0xe0, 0x8c, 0xda, 0x04, 0x7d, 0x0e, 0xb9, 0x8f,
	0x88, 0x18, 0xdc, 0xbf, 0x85, 0x18, 0x87, 0xa6, 0xc3, 0x4b, 0x4b, 0x31, 0x71, 0x6e, 0x96, 0x7e,
	0xf8, 0xed, 0x8f, 0x9f, 0x53, 0x4b, 0x08, 0x59, 0x67, 0x5b, 0x96, 0xbe, 0xe7, 0xb8, 0xf5, 0x82,
	0x3a, 0xfc, 0x0a, 0xb9, 0xb0, 0x30, 0x34, 0x96, 0x63, 0x43, 0xaf, 0xc5, 0x78, 0x44, 0x78, 0x95,
	0xca, 0x89, 0xf9, 0x70, 0xde, 0xe6, 0x92, 0x6a, 0xb7, 0x80, 0xe6, 0xa2, 0xed, 0xd0, 0x13, 0x98,
	0xdf, 0x53, 0xd7, 0x58, 0xff, 0x9f, 0x74, 0xdc, 0x06, 0x2e, 0x15, 0xc6, 0xae, 0xb2, 0x0f, 0xe5,
	0xf7, 0x84, 0xb9, 0xa2, 0x3c, 0xf3, 0xe6, 0x88, 0xe7, 0x8e, 0x51, 0x43, 0x4f, 0x61, 0xfe, 0x09,
	0x73, 0xfe, 0xa9, 0xed, 0x5d, 0x65, 0x5b, 0x30, 0xf3, 0x2f, 0x4f, 0xe6, 0x4a, 0x7a, 0x7f, 0x05,
	0x0b, 0x1f, 0x90, 0x0e, 0x19, 0x78, 0x27, 0xcf, 0x3d, 0xc9, 0x5f, 0x4f, 0xbe, 0x16, 0x37, 0xf9,
	0x53, 0x78, 0x65, 0x38, 0xf9, 0xf0, 0x3a, 0x2e, 0x25, 0x6c, 0x0c, 0xd9, 0xa2, 0x98, 0x90, 0xe3,
	0x66, 0x59, 0x35, 0x59, 0x45, 0x2b, 0x91, 0x26, 0x6d, 0x4f, 0xa6, 0x74, 0xa7, 0x0b, 0x40, 0xa3,
	0x9d, 0x14, 0x67, 0x33, 0xc1, 0x30, 0xca, 0xfa, 0xf5, 0xbf, 0xd4, 0x68, 0xde, 0xab, 0xaa, 0xff,
	0x22, 0xca, 0x8f, 0xf5, 0x47, 0x0e, 0xa0, 0x11, 0xe8, 0xe1, 0x32, 0x13, 0x97, 0x72, 0x2b, 0x4e,
	0x61, 0x0b, 0xc9, 0xc9, 0x85, 0xfc, 0xc8, 0x1e, 0x50, 0xdf, 0x19, 0x49, 0x87, 0x2c, 0xb1, 0xc7,
	0x3d, 0xd5, 0x63, 0xcd, 0x2c, 0xc4, 0x8e, 0x51, 0x6d, 0x08, 0x0a, 0x68, 0x64, 0x43, 0x4c, 0xa6,
	0x96, 0xd4, 0x4c, 0x33, 0xab, 0x25, 0x31, 0xdb, 0xfd, 0xfa, 0xa7, 0x47, 0x47, 0xe8, 0x81, 0x79,
	0x1f, 0xd0, 0xa7, 0x8c, 0x78, 0x07, 0x34, 0x20, 0xf4, 0xa2, 0x72, 0x10, 0xf8, 0xcf, 0x89, 0x2d,
	0xd0, 0xf2, 0xa9, 0x10, 0x8c, 0xef, 0x58, 0x56, 0xa4, 0x31, 0xf5, 0xb7, 0x33, 0x9b, 0x8d, 0xcd,
	0xc6, 0x56, 0xcd, 0x30, 0xb6, 0xef, 0x60, 0xc6, 0x3a, 0xd4, 0x56, 0x1f, 0xc8, 0xd6, 0x73, 0xee,
	0x7b, 0x3b, 0x63, 0x91, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x78, 0x3a, 0xae, 0x21, 0x0c,
	0x00, 0x00,
}
