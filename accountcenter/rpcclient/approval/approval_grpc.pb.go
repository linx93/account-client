// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: approval.proto

package approval

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ApprovalService_CreateApprovalConfig_FullMethodName = "/approval.ApprovalService/CreateApprovalConfig"
	ApprovalService_GetApprovalConfig_FullMethodName    = "/approval.ApprovalService/GetApprovalConfig"
	ApprovalService_GetApprovalConfigs_FullMethodName   = "/approval.ApprovalService/GetApprovalConfigs"
	ApprovalService_CreateApproval_FullMethodName       = "/approval.ApprovalService/CreateApproval"
	ApprovalService_ApproveApproval_FullMethodName      = "/approval.ApprovalService/ApproveApproval"
	ApprovalService_DenyApproval_FullMethodName         = "/approval.ApprovalService/DenyApproval"
	ApprovalService_GetMyApproval_FullMethodName        = "/approval.ApprovalService/GetMyApproval"
	ApprovalService_PageMyApproval_FullMethodName       = "/approval.ApprovalService/PageMyApproval"
	ApprovalService_GetApprovalSequences_FullMethodName = "/approval.ApprovalService/GetApprovalSequences"
)

// ApprovalServiceClient is the client API for ApprovalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApprovalServiceClient interface {
	// CreateApprovalConfig
	// 创建审批配置
	CreateApprovalConfig(ctx context.Context, in *CreateConfigRequest, opts ...grpc.CallOption) (*CreateConfigResponse, error)
	// GetConfig
	// 获取审批配置
	GetApprovalConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
	// GetConfigs
	// 获取应用租户的所有审批配置
	GetApprovalConfigs(ctx context.Context, in *GetConfigsRequest, opts ...grpc.CallOption) (*GetConfigsResponse, error)
	// CreateApproval
	// 创建审批实例
	CreateApproval(ctx context.Context, in *CreateApprovalRequest, opts ...grpc.CallOption) (*CreateApprovalResponse, error)
	// ApproveApproval
	// 审批通过
	ApproveApproval(ctx context.Context, in *ApproveRequest, opts ...grpc.CallOption) (*ApproveResponse, error)
	// DenyApproval
	// 审批不通过
	DenyApproval(ctx context.Context, in *DenyRequest, opts ...grpc.CallOption) (*DenyResponse, error)
	// GetMyApproval
	// 获取我的待审批
	GetMyApproval(ctx context.Context, in *GetMyApprovalRequest, opts ...grpc.CallOption) (*GetMyApprovalResponse, error)
	// PageMyApproval
	// 分页获取我的审批记录
	PageMyApproval(ctx context.Context, in *PageMyApprovalRequest, opts ...grpc.CallOption) (*PageMyApprovalResponse, error)
	// GetApprovalSequences
	// 获取审核实例的审核序列
	GetApprovalSequences(ctx context.Context, in *GetSequencesRequest, opts ...grpc.CallOption) (*GetSequencesResponse, error)
}

type approvalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApprovalServiceClient(cc grpc.ClientConnInterface) ApprovalServiceClient {
	return &approvalServiceClient{cc}
}

func (c *approvalServiceClient) CreateApprovalConfig(ctx context.Context, in *CreateConfigRequest, opts ...grpc.CallOption) (*CreateConfigResponse, error) {
	out := new(CreateConfigResponse)
	err := c.cc.Invoke(ctx, ApprovalService_CreateApprovalConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvalServiceClient) GetApprovalConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := c.cc.Invoke(ctx, ApprovalService_GetApprovalConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvalServiceClient) GetApprovalConfigs(ctx context.Context, in *GetConfigsRequest, opts ...grpc.CallOption) (*GetConfigsResponse, error) {
	out := new(GetConfigsResponse)
	err := c.cc.Invoke(ctx, ApprovalService_GetApprovalConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvalServiceClient) CreateApproval(ctx context.Context, in *CreateApprovalRequest, opts ...grpc.CallOption) (*CreateApprovalResponse, error) {
	out := new(CreateApprovalResponse)
	err := c.cc.Invoke(ctx, ApprovalService_CreateApproval_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvalServiceClient) ApproveApproval(ctx context.Context, in *ApproveRequest, opts ...grpc.CallOption) (*ApproveResponse, error) {
	out := new(ApproveResponse)
	err := c.cc.Invoke(ctx, ApprovalService_ApproveApproval_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvalServiceClient) DenyApproval(ctx context.Context, in *DenyRequest, opts ...grpc.CallOption) (*DenyResponse, error) {
	out := new(DenyResponse)
	err := c.cc.Invoke(ctx, ApprovalService_DenyApproval_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvalServiceClient) GetMyApproval(ctx context.Context, in *GetMyApprovalRequest, opts ...grpc.CallOption) (*GetMyApprovalResponse, error) {
	out := new(GetMyApprovalResponse)
	err := c.cc.Invoke(ctx, ApprovalService_GetMyApproval_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvalServiceClient) PageMyApproval(ctx context.Context, in *PageMyApprovalRequest, opts ...grpc.CallOption) (*PageMyApprovalResponse, error) {
	out := new(PageMyApprovalResponse)
	err := c.cc.Invoke(ctx, ApprovalService_PageMyApproval_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvalServiceClient) GetApprovalSequences(ctx context.Context, in *GetSequencesRequest, opts ...grpc.CallOption) (*GetSequencesResponse, error) {
	out := new(GetSequencesResponse)
	err := c.cc.Invoke(ctx, ApprovalService_GetApprovalSequences_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApprovalServiceServer is the server API for ApprovalService service.
// All implementations must embed UnimplementedApprovalServiceServer
// for forward compatibility
type ApprovalServiceServer interface {
	// CreateApprovalConfig
	// 创建审批配置
	CreateApprovalConfig(context.Context, *CreateConfigRequest) (*CreateConfigResponse, error)
	// GetConfig
	// 获取审批配置
	GetApprovalConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
	// GetConfigs
	// 获取应用租户的所有审批配置
	GetApprovalConfigs(context.Context, *GetConfigsRequest) (*GetConfigsResponse, error)
	// CreateApproval
	// 创建审批实例
	CreateApproval(context.Context, *CreateApprovalRequest) (*CreateApprovalResponse, error)
	// ApproveApproval
	// 审批通过
	ApproveApproval(context.Context, *ApproveRequest) (*ApproveResponse, error)
	// DenyApproval
	// 审批不通过
	DenyApproval(context.Context, *DenyRequest) (*DenyResponse, error)
	// GetMyApproval
	// 获取我的待审批
	GetMyApproval(context.Context, *GetMyApprovalRequest) (*GetMyApprovalResponse, error)
	// PageMyApproval
	// 分页获取我的审批记录
	PageMyApproval(context.Context, *PageMyApprovalRequest) (*PageMyApprovalResponse, error)
	// GetApprovalSequences
	// 获取审核实例的审核序列
	GetApprovalSequences(context.Context, *GetSequencesRequest) (*GetSequencesResponse, error)
	mustEmbedUnimplementedApprovalServiceServer()
}

// UnimplementedApprovalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApprovalServiceServer struct {
}

func (UnimplementedApprovalServiceServer) CreateApprovalConfig(context.Context, *CreateConfigRequest) (*CreateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApprovalConfig not implemented")
}
func (UnimplementedApprovalServiceServer) GetApprovalConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApprovalConfig not implemented")
}
func (UnimplementedApprovalServiceServer) GetApprovalConfigs(context.Context, *GetConfigsRequest) (*GetConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApprovalConfigs not implemented")
}
func (UnimplementedApprovalServiceServer) CreateApproval(context.Context, *CreateApprovalRequest) (*CreateApprovalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApproval not implemented")
}
func (UnimplementedApprovalServiceServer) ApproveApproval(context.Context, *ApproveRequest) (*ApproveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveApproval not implemented")
}
func (UnimplementedApprovalServiceServer) DenyApproval(context.Context, *DenyRequest) (*DenyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenyApproval not implemented")
}
func (UnimplementedApprovalServiceServer) GetMyApproval(context.Context, *GetMyApprovalRequest) (*GetMyApprovalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyApproval not implemented")
}
func (UnimplementedApprovalServiceServer) PageMyApproval(context.Context, *PageMyApprovalRequest) (*PageMyApprovalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageMyApproval not implemented")
}
func (UnimplementedApprovalServiceServer) GetApprovalSequences(context.Context, *GetSequencesRequest) (*GetSequencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApprovalSequences not implemented")
}
func (UnimplementedApprovalServiceServer) mustEmbedUnimplementedApprovalServiceServer() {}

// UnsafeApprovalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApprovalServiceServer will
// result in compilation errors.
type UnsafeApprovalServiceServer interface {
	mustEmbedUnimplementedApprovalServiceServer()
}

func RegisterApprovalServiceServer(s grpc.ServiceRegistrar, srv ApprovalServiceServer) {
	s.RegisterService(&ApprovalService_ServiceDesc, srv)
}

func _ApprovalService_CreateApprovalConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalServiceServer).CreateApprovalConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApprovalService_CreateApprovalConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalServiceServer).CreateApprovalConfig(ctx, req.(*CreateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprovalService_GetApprovalConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalServiceServer).GetApprovalConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApprovalService_GetApprovalConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalServiceServer).GetApprovalConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprovalService_GetApprovalConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalServiceServer).GetApprovalConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApprovalService_GetApprovalConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalServiceServer).GetApprovalConfigs(ctx, req.(*GetConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprovalService_CreateApproval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApprovalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalServiceServer).CreateApproval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApprovalService_CreateApproval_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalServiceServer).CreateApproval(ctx, req.(*CreateApprovalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprovalService_ApproveApproval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalServiceServer).ApproveApproval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApprovalService_ApproveApproval_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalServiceServer).ApproveApproval(ctx, req.(*ApproveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprovalService_DenyApproval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DenyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalServiceServer).DenyApproval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApprovalService_DenyApproval_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalServiceServer).DenyApproval(ctx, req.(*DenyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprovalService_GetMyApproval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyApprovalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalServiceServer).GetMyApproval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApprovalService_GetMyApproval_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalServiceServer).GetMyApproval(ctx, req.(*GetMyApprovalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprovalService_PageMyApproval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageMyApprovalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalServiceServer).PageMyApproval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApprovalService_PageMyApproval_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalServiceServer).PageMyApproval(ctx, req.(*PageMyApprovalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprovalService_GetApprovalSequences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSequencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalServiceServer).GetApprovalSequences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApprovalService_GetApprovalSequences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalServiceServer).GetApprovalSequences(ctx, req.(*GetSequencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApprovalService_ServiceDesc is the grpc.ServiceDesc for ApprovalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApprovalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "approval.ApprovalService",
	HandlerType: (*ApprovalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApprovalConfig",
			Handler:    _ApprovalService_CreateApprovalConfig_Handler,
		},
		{
			MethodName: "GetApprovalConfig",
			Handler:    _ApprovalService_GetApprovalConfig_Handler,
		},
		{
			MethodName: "GetApprovalConfigs",
			Handler:    _ApprovalService_GetApprovalConfigs_Handler,
		},
		{
			MethodName: "CreateApproval",
			Handler:    _ApprovalService_CreateApproval_Handler,
		},
		{
			MethodName: "ApproveApproval",
			Handler:    _ApprovalService_ApproveApproval_Handler,
		},
		{
			MethodName: "DenyApproval",
			Handler:    _ApprovalService_DenyApproval_Handler,
		},
		{
			MethodName: "GetMyApproval",
			Handler:    _ApprovalService_GetMyApproval_Handler,
		},
		{
			MethodName: "PageMyApproval",
			Handler:    _ApprovalService_PageMyApproval_Handler,
		},
		{
			MethodName: "GetApprovalSequences",
			Handler:    _ApprovalService_GetApprovalSequences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "approval.proto",
}