// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pkg/serverpb/admin.proto

package serverpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DKVReplicationClient is the client API for DKVReplication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DKVReplicationClient interface {
	// GetChanges retrieves all changes from a given change number.
	GetChanges(ctx context.Context, in *GetChangesRequest, opts ...grpc.CallOption) (*GetChangesResponse, error)
	// AddReplica registers a new replica with the current master.
	AddReplica(ctx context.Context, in *Replica, opts ...grpc.CallOption) (*Status, error)
	// RemoveReplica deregisters given replica from the current master.
	RemoveReplica(ctx context.Context, in *Replica, opts ...grpc.CallOption) (*Status, error)
	// GetReplicas retrieves all the replicas of the current master.
	GetReplicas(ctx context.Context, in *GetReplicasRequest, opts ...grpc.CallOption) (*GetReplicasResponse, error)
}

type dKVReplicationClient struct {
	cc grpc.ClientConnInterface
}

func NewDKVReplicationClient(cc grpc.ClientConnInterface) DKVReplicationClient {
	return &dKVReplicationClient{cc}
}

func (c *dKVReplicationClient) GetChanges(ctx context.Context, in *GetChangesRequest, opts ...grpc.CallOption) (*GetChangesResponse, error) {
	out := new(GetChangesResponse)
	err := c.cc.Invoke(ctx, "/dkv.serverpb.DKVReplication/GetChanges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dKVReplicationClient) AddReplica(ctx context.Context, in *Replica, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/dkv.serverpb.DKVReplication/AddReplica", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dKVReplicationClient) RemoveReplica(ctx context.Context, in *Replica, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/dkv.serverpb.DKVReplication/RemoveReplica", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dKVReplicationClient) GetReplicas(ctx context.Context, in *GetReplicasRequest, opts ...grpc.CallOption) (*GetReplicasResponse, error) {
	out := new(GetReplicasResponse)
	err := c.cc.Invoke(ctx, "/dkv.serverpb.DKVReplication/GetReplicas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DKVReplicationServer is the server API for DKVReplication service.
// All implementations should embed UnimplementedDKVReplicationServer
// for forward compatibility
type DKVReplicationServer interface {
	// GetChanges retrieves all changes from a given change number.
	GetChanges(context.Context, *GetChangesRequest) (*GetChangesResponse, error)
	// AddReplica registers a new replica with the current master.
	AddReplica(context.Context, *Replica) (*Status, error)
	// RemoveReplica deregisters given replica from the current master.
	RemoveReplica(context.Context, *Replica) (*Status, error)
	// GetReplicas retrieves all the replicas of the current master.
	GetReplicas(context.Context, *GetReplicasRequest) (*GetReplicasResponse, error)
}

// UnimplementedDKVReplicationServer should be embedded to have forward compatible implementations.
type UnimplementedDKVReplicationServer struct {
}

func (UnimplementedDKVReplicationServer) GetChanges(context.Context, *GetChangesRequest) (*GetChangesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChanges not implemented")
}
func (UnimplementedDKVReplicationServer) AddReplica(context.Context, *Replica) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReplica not implemented")
}
func (UnimplementedDKVReplicationServer) RemoveReplica(context.Context, *Replica) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveReplica not implemented")
}
func (UnimplementedDKVReplicationServer) GetReplicas(context.Context, *GetReplicasRequest) (*GetReplicasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReplicas not implemented")
}

// UnsafeDKVReplicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DKVReplicationServer will
// result in compilation errors.
type UnsafeDKVReplicationServer interface {
	mustEmbedUnimplementedDKVReplicationServer()
}

func RegisterDKVReplicationServer(s grpc.ServiceRegistrar, srv DKVReplicationServer) {
	s.RegisterService(&DKVReplication_ServiceDesc, srv)
}

func _DKVReplication_GetChanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChangesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DKVReplicationServer).GetChanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkv.serverpb.DKVReplication/GetChanges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DKVReplicationServer).GetChanges(ctx, req.(*GetChangesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DKVReplication_AddReplica_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Replica)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DKVReplicationServer).AddReplica(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkv.serverpb.DKVReplication/AddReplica",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DKVReplicationServer).AddReplica(ctx, req.(*Replica))
	}
	return interceptor(ctx, in, info, handler)
}

func _DKVReplication_RemoveReplica_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Replica)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DKVReplicationServer).RemoveReplica(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkv.serverpb.DKVReplication/RemoveReplica",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DKVReplicationServer).RemoveReplica(ctx, req.(*Replica))
	}
	return interceptor(ctx, in, info, handler)
}

func _DKVReplication_GetReplicas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReplicasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DKVReplicationServer).GetReplicas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkv.serverpb.DKVReplication/GetReplicas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DKVReplicationServer).GetReplicas(ctx, req.(*GetReplicasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DKVReplication_ServiceDesc is the grpc.ServiceDesc for DKVReplication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DKVReplication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dkv.serverpb.DKVReplication",
	HandlerType: (*DKVReplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChanges",
			Handler:    _DKVReplication_GetChanges_Handler,
		},
		{
			MethodName: "AddReplica",
			Handler:    _DKVReplication_AddReplica_Handler,
		},
		{
			MethodName: "RemoveReplica",
			Handler:    _DKVReplication_RemoveReplica_Handler,
		},
		{
			MethodName: "GetReplicas",
			Handler:    _DKVReplication_GetReplicas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/serverpb/admin.proto",
}

// DKVBackupRestoreClient is the client API for DKVBackupRestore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DKVBackupRestoreClient interface {
	// Backup backs up the entire keyspace into the given filesystem location.
	Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*Status, error)
	// Restore restores the entire keyspace from an existing backup at the
	// given filesystem location.
	Restore(ctx context.Context, in *RestoreRequest, opts ...grpc.CallOption) (*Status, error)
}

type dKVBackupRestoreClient struct {
	cc grpc.ClientConnInterface
}

func NewDKVBackupRestoreClient(cc grpc.ClientConnInterface) DKVBackupRestoreClient {
	return &dKVBackupRestoreClient{cc}
}

func (c *dKVBackupRestoreClient) Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/dkv.serverpb.DKVBackupRestore/Backup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dKVBackupRestoreClient) Restore(ctx context.Context, in *RestoreRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/dkv.serverpb.DKVBackupRestore/Restore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DKVBackupRestoreServer is the server API for DKVBackupRestore service.
// All implementations should embed UnimplementedDKVBackupRestoreServer
// for forward compatibility
type DKVBackupRestoreServer interface {
	// Backup backs up the entire keyspace into the given filesystem location.
	Backup(context.Context, *BackupRequest) (*Status, error)
	// Restore restores the entire keyspace from an existing backup at the
	// given filesystem location.
	Restore(context.Context, *RestoreRequest) (*Status, error)
}

// UnimplementedDKVBackupRestoreServer should be embedded to have forward compatible implementations.
type UnimplementedDKVBackupRestoreServer struct {
}

func (UnimplementedDKVBackupRestoreServer) Backup(context.Context, *BackupRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Backup not implemented")
}
func (UnimplementedDKVBackupRestoreServer) Restore(context.Context, *RestoreRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

// UnsafeDKVBackupRestoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DKVBackupRestoreServer will
// result in compilation errors.
type UnsafeDKVBackupRestoreServer interface {
	mustEmbedUnimplementedDKVBackupRestoreServer()
}

func RegisterDKVBackupRestoreServer(s grpc.ServiceRegistrar, srv DKVBackupRestoreServer) {
	s.RegisterService(&DKVBackupRestore_ServiceDesc, srv)
}

func _DKVBackupRestore_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DKVBackupRestoreServer).Backup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkv.serverpb.DKVBackupRestore/Backup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DKVBackupRestoreServer).Backup(ctx, req.(*BackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DKVBackupRestore_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DKVBackupRestoreServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkv.serverpb.DKVBackupRestore/Restore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DKVBackupRestoreServer).Restore(ctx, req.(*RestoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DKVBackupRestore_ServiceDesc is the grpc.ServiceDesc for DKVBackupRestore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DKVBackupRestore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dkv.serverpb.DKVBackupRestore",
	HandlerType: (*DKVBackupRestoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Backup",
			Handler:    _DKVBackupRestore_Backup_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _DKVBackupRestore_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/serverpb/admin.proto",
}

// DKVClusterClient is the client API for DKVCluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DKVClusterClient interface {
	// AddNode adds the given DKV node to the cluster that the
	// current node is a member of.
	AddNode(ctx context.Context, in *AddNodeRequest, opts ...grpc.CallOption) (*Status, error)
	// RemoveNode removes the given DKV node from the cluster that
	// the current node is a member of.
	RemoveNode(ctx context.Context, in *RemoveNodeRequest, opts ...grpc.CallOption) (*Status, error)
	// ListNodes retrieves the current set of DKV nodes from
	// the Nexus cluster.
	ListNodes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListNodesResponse, error)
}

type dKVClusterClient struct {
	cc grpc.ClientConnInterface
}

func NewDKVClusterClient(cc grpc.ClientConnInterface) DKVClusterClient {
	return &dKVClusterClient{cc}
}

func (c *dKVClusterClient) AddNode(ctx context.Context, in *AddNodeRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/dkv.serverpb.DKVCluster/AddNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dKVClusterClient) RemoveNode(ctx context.Context, in *RemoveNodeRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/dkv.serverpb.DKVCluster/RemoveNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dKVClusterClient) ListNodes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListNodesResponse, error) {
	out := new(ListNodesResponse)
	err := c.cc.Invoke(ctx, "/dkv.serverpb.DKVCluster/ListNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DKVClusterServer is the server API for DKVCluster service.
// All implementations should embed UnimplementedDKVClusterServer
// for forward compatibility
type DKVClusterServer interface {
	// AddNode adds the given DKV node to the cluster that the
	// current node is a member of.
	AddNode(context.Context, *AddNodeRequest) (*Status, error)
	// RemoveNode removes the given DKV node from the cluster that
	// the current node is a member of.
	RemoveNode(context.Context, *RemoveNodeRequest) (*Status, error)
	// ListNodes retrieves the current set of DKV nodes from
	// the Nexus cluster.
	ListNodes(context.Context, *emptypb.Empty) (*ListNodesResponse, error)
}

// UnimplementedDKVClusterServer should be embedded to have forward compatible implementations.
type UnimplementedDKVClusterServer struct {
}

func (UnimplementedDKVClusterServer) AddNode(context.Context, *AddNodeRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNode not implemented")
}
func (UnimplementedDKVClusterServer) RemoveNode(context.Context, *RemoveNodeRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveNode not implemented")
}
func (UnimplementedDKVClusterServer) ListNodes(context.Context, *emptypb.Empty) (*ListNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodes not implemented")
}

// UnsafeDKVClusterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DKVClusterServer will
// result in compilation errors.
type UnsafeDKVClusterServer interface {
	mustEmbedUnimplementedDKVClusterServer()
}

func RegisterDKVClusterServer(s grpc.ServiceRegistrar, srv DKVClusterServer) {
	s.RegisterService(&DKVCluster_ServiceDesc, srv)
}

func _DKVCluster_AddNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DKVClusterServer).AddNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkv.serverpb.DKVCluster/AddNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DKVClusterServer).AddNode(ctx, req.(*AddNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DKVCluster_RemoveNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DKVClusterServer).RemoveNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkv.serverpb.DKVCluster/RemoveNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DKVClusterServer).RemoveNode(ctx, req.(*RemoveNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DKVCluster_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DKVClusterServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkv.serverpb.DKVCluster/ListNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DKVClusterServer).ListNodes(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DKVCluster_ServiceDesc is the grpc.ServiceDesc for DKVCluster service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DKVCluster_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dkv.serverpb.DKVCluster",
	HandlerType: (*DKVClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNode",
			Handler:    _DKVCluster_AddNode_Handler,
		},
		{
			MethodName: "RemoveNode",
			Handler:    _DKVCluster_RemoveNode_Handler,
		},
		{
			MethodName: "ListNodes",
			Handler:    _DKVCluster_ListNodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/serverpb/admin.proto",
}

// DKVDiscoveryClient is the client API for DKVDiscovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DKVDiscoveryClient interface {
	// Update status of the given database and vBucket
	UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*Status, error)
	// Retrieve all active nodes in cluster with their status after filtering based on request params
	GetClusterInfo(ctx context.Context, in *GetClusterInfoRequest, opts ...grpc.CallOption) (*GetClusterInfoResponse, error)
}

type dKVDiscoveryClient struct {
	cc grpc.ClientConnInterface
}

func NewDKVDiscoveryClient(cc grpc.ClientConnInterface) DKVDiscoveryClient {
	return &dKVDiscoveryClient{cc}
}

func (c *dKVDiscoveryClient) UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/dkv.serverpb.DKVDiscovery/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dKVDiscoveryClient) GetClusterInfo(ctx context.Context, in *GetClusterInfoRequest, opts ...grpc.CallOption) (*GetClusterInfoResponse, error) {
	out := new(GetClusterInfoResponse)
	err := c.cc.Invoke(ctx, "/dkv.serverpb.DKVDiscovery/GetClusterInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DKVDiscoveryServer is the server API for DKVDiscovery service.
// All implementations should embed UnimplementedDKVDiscoveryServer
// for forward compatibility
type DKVDiscoveryServer interface {
	// Update status of the given database and vBucket
	UpdateStatus(context.Context, *UpdateStatusRequest) (*Status, error)
	// Retrieve all active nodes in cluster with their status after filtering based on request params
	GetClusterInfo(context.Context, *GetClusterInfoRequest) (*GetClusterInfoResponse, error)
}

// UnimplementedDKVDiscoveryServer should be embedded to have forward compatible implementations.
type UnimplementedDKVDiscoveryServer struct {
}

func (UnimplementedDKVDiscoveryServer) UpdateStatus(context.Context, *UpdateStatusRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedDKVDiscoveryServer) GetClusterInfo(context.Context, *GetClusterInfoRequest) (*GetClusterInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterInfo not implemented")
}

// UnsafeDKVDiscoveryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DKVDiscoveryServer will
// result in compilation errors.
type UnsafeDKVDiscoveryServer interface {
	mustEmbedUnimplementedDKVDiscoveryServer()
}

func RegisterDKVDiscoveryServer(s grpc.ServiceRegistrar, srv DKVDiscoveryServer) {
	s.RegisterService(&DKVDiscovery_ServiceDesc, srv)
}

func _DKVDiscovery_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DKVDiscoveryServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkv.serverpb.DKVDiscovery/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DKVDiscoveryServer).UpdateStatus(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DKVDiscovery_GetClusterInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DKVDiscoveryServer).GetClusterInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkv.serverpb.DKVDiscovery/GetClusterInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DKVDiscoveryServer).GetClusterInfo(ctx, req.(*GetClusterInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DKVDiscovery_ServiceDesc is the grpc.ServiceDesc for DKVDiscovery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DKVDiscovery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dkv.serverpb.DKVDiscovery",
	HandlerType: (*DKVDiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateStatus",
			Handler:    _DKVDiscovery_UpdateStatus_Handler,
		},
		{
			MethodName: "GetClusterInfo",
			Handler:    _DKVDiscovery_GetClusterInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/serverpb/admin.proto",
}

// DKVDiscoveryNodeClient is the client API for DKVDiscoveryNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DKVDiscoveryNodeClient interface {
	// Get status of region
	GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RegionInfo, error)
}

type dKVDiscoveryNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewDKVDiscoveryNodeClient(cc grpc.ClientConnInterface) DKVDiscoveryNodeClient {
	return &dKVDiscoveryNodeClient{cc}
}

func (c *dKVDiscoveryNodeClient) GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RegionInfo, error) {
	out := new(RegionInfo)
	err := c.cc.Invoke(ctx, "/dkv.serverpb.DKVDiscoveryNode/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DKVDiscoveryNodeServer is the server API for DKVDiscoveryNode service.
// All implementations should embed UnimplementedDKVDiscoveryNodeServer
// for forward compatibility
type DKVDiscoveryNodeServer interface {
	// Get status of region
	GetStatus(context.Context, *emptypb.Empty) (*RegionInfo, error)
}

// UnimplementedDKVDiscoveryNodeServer should be embedded to have forward compatible implementations.
type UnimplementedDKVDiscoveryNodeServer struct {
}

func (UnimplementedDKVDiscoveryNodeServer) GetStatus(context.Context, *emptypb.Empty) (*RegionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}

// UnsafeDKVDiscoveryNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DKVDiscoveryNodeServer will
// result in compilation errors.
type UnsafeDKVDiscoveryNodeServer interface {
	mustEmbedUnimplementedDKVDiscoveryNodeServer()
}

func RegisterDKVDiscoveryNodeServer(s grpc.ServiceRegistrar, srv DKVDiscoveryNodeServer) {
	s.RegisterService(&DKVDiscoveryNode_ServiceDesc, srv)
}

func _DKVDiscoveryNode_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DKVDiscoveryNodeServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkv.serverpb.DKVDiscoveryNode/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DKVDiscoveryNodeServer).GetStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DKVDiscoveryNode_ServiceDesc is the grpc.ServiceDesc for DKVDiscoveryNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DKVDiscoveryNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dkv.serverpb.DKVDiscoveryNode",
	HandlerType: (*DKVDiscoveryNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _DKVDiscoveryNode_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/serverpb/admin.proto",
}
