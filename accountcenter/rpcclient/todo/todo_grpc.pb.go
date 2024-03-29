// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: todo.proto

package todo

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
	TodoService_CreateTodo_FullMethodName = "/todo.TodoService/CreateTodo"
	TodoService_GetTodos_FullMethodName   = "/todo.TodoService/GetTodos"
	TodoService_PageTodos_FullMethodName  = "/todo.TodoService/PageTodos"
	TodoService_PushRemark_FullMethodName = "/todo.TodoService/PushRemark"
	TodoService_GetRemarks_FullMethodName = "/todo.TodoService/GetRemarks"
)

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error)
	GetTodos(ctx context.Context, in *GetTodosRequest, opts ...grpc.CallOption) (*GetTodosResponse, error)
	PageTodos(ctx context.Context, in *PageTodosRequest, opts ...grpc.CallOption) (*PageTodosResponse, error)
	PushRemark(ctx context.Context, in *PushRemarkRequest, opts ...grpc.CallOption) (*PushRemarkResponse, error)
	GetRemarks(ctx context.Context, in *GetRemarksRequest, opts ...grpc.CallOption) (*GetRemarksResponse, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error) {
	out := new(CreateTodoResponse)
	err := c.cc.Invoke(ctx, TodoService_CreateTodo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodos(ctx context.Context, in *GetTodosRequest, opts ...grpc.CallOption) (*GetTodosResponse, error) {
	out := new(GetTodosResponse)
	err := c.cc.Invoke(ctx, TodoService_GetTodos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) PageTodos(ctx context.Context, in *PageTodosRequest, opts ...grpc.CallOption) (*PageTodosResponse, error) {
	out := new(PageTodosResponse)
	err := c.cc.Invoke(ctx, TodoService_PageTodos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) PushRemark(ctx context.Context, in *PushRemarkRequest, opts ...grpc.CallOption) (*PushRemarkResponse, error) {
	out := new(PushRemarkResponse)
	err := c.cc.Invoke(ctx, TodoService_PushRemark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetRemarks(ctx context.Context, in *GetRemarksRequest, opts ...grpc.CallOption) (*GetRemarksResponse, error) {
	out := new(GetRemarksResponse)
	err := c.cc.Invoke(ctx, TodoService_GetRemarks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
	GetTodos(context.Context, *GetTodosRequest) (*GetTodosResponse, error)
	PageTodos(context.Context, *PageTodosRequest) (*PageTodosResponse, error)
	PushRemark(context.Context, *PushRemarkRequest) (*PushRemarkResponse, error)
	GetRemarks(context.Context, *GetRemarksRequest) (*GetRemarksResponse, error)
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodoServiceServer) GetTodos(context.Context, *GetTodosRequest) (*GetTodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodos not implemented")
}
func (UnimplementedTodoServiceServer) PageTodos(context.Context, *PageTodosRequest) (*PageTodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageTodos not implemented")
}
func (UnimplementedTodoServiceServer) PushRemark(context.Context, *PushRemarkRequest) (*PushRemarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushRemark not implemented")
}
func (UnimplementedTodoServiceServer) GetRemarks(context.Context, *GetRemarksRequest) (*GetRemarksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRemarks not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_CreateTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_GetTodos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetTodos(ctx, req.(*GetTodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_PageTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageTodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).PageTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_PageTodos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).PageTodos(ctx, req.(*PageTodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_PushRemark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRemarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).PushRemark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_PushRemark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).PushRemark(ctx, req.(*PushRemarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetRemarks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRemarksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetRemarks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_GetRemarks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetRemarks(ctx, req.(*GetRemarksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "GetTodos",
			Handler:    _TodoService_GetTodos_Handler,
		},
		{
			MethodName: "PageTodos",
			Handler:    _TodoService_PageTodos_Handler,
		},
		{
			MethodName: "PushRemark",
			Handler:    _TodoService_PushRemark_Handler,
		},
		{
			MethodName: "GetRemarks",
			Handler:    _TodoService_GetRemarks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
