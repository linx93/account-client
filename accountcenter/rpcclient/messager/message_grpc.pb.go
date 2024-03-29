// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: message.proto

package messager

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
	MessageService_CreateTemplate_FullMethodName   = "/messager.MessageService/CreateTemplate"
	MessageService_GetTemplate_FullMethodName      = "/messager.MessageService/GetTemplate"
	MessageService_GetTemplates_FullMethodName     = "/messager.MessageService/GetTemplates"
	MessageService_CreateBinding_FullMethodName    = "/messager.MessageService/CreateBinding"
	MessageService_SetReceivers_FullMethodName     = "/messager.MessageService/SetReceivers"
	MessageService_SendMessage_FullMethodName      = "/messager.MessageService/SendMessage"
	MessageService_PageMessage_FullMethodName      = "/messager.MessageService/PageMessage"
	MessageService_MakeMessagesRead_FullMethodName = "/messager.MessageService/MakeMessagesRead"
)

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error)
	GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*GetTemplateResponse, error)
	GetTemplates(ctx context.Context, in *GetTemplatesRequest, opts ...grpc.CallOption) (*GetTemplatesResponse, error)
	CreateBinding(ctx context.Context, in *CreateBindingRequest, opts ...grpc.CallOption) (*CreateBindingResponse, error)
	SetReceivers(ctx context.Context, in *SetReceiversRequest, opts ...grpc.CallOption) (*SetReceiversResponse, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	PageMessage(ctx context.Context, in *PageMessageRequest, opts ...grpc.CallOption) (*PageMessageResponse, error)
	MakeMessagesRead(ctx context.Context, in *MakeMessagesReadRequest, opts ...grpc.CallOption) (*MakeMessagesReadResponse, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error) {
	out := new(CreateTemplateResponse)
	err := c.cc.Invoke(ctx, MessageService_CreateTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*GetTemplateResponse, error) {
	out := new(GetTemplateResponse)
	err := c.cc.Invoke(ctx, MessageService_GetTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetTemplates(ctx context.Context, in *GetTemplatesRequest, opts ...grpc.CallOption) (*GetTemplatesResponse, error) {
	out := new(GetTemplatesResponse)
	err := c.cc.Invoke(ctx, MessageService_GetTemplates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) CreateBinding(ctx context.Context, in *CreateBindingRequest, opts ...grpc.CallOption) (*CreateBindingResponse, error) {
	out := new(CreateBindingResponse)
	err := c.cc.Invoke(ctx, MessageService_CreateBinding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SetReceivers(ctx context.Context, in *SetReceiversRequest, opts ...grpc.CallOption) (*SetReceiversResponse, error) {
	out := new(SetReceiversResponse)
	err := c.cc.Invoke(ctx, MessageService_SetReceivers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, MessageService_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) PageMessage(ctx context.Context, in *PageMessageRequest, opts ...grpc.CallOption) (*PageMessageResponse, error) {
	out := new(PageMessageResponse)
	err := c.cc.Invoke(ctx, MessageService_PageMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) MakeMessagesRead(ctx context.Context, in *MakeMessagesReadRequest, opts ...grpc.CallOption) (*MakeMessagesReadResponse, error) {
	out := new(MakeMessagesReadResponse)
	err := c.cc.Invoke(ctx, MessageService_MakeMessagesRead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error)
	GetTemplate(context.Context, *GetTemplateRequest) (*GetTemplateResponse, error)
	GetTemplates(context.Context, *GetTemplatesRequest) (*GetTemplatesResponse, error)
	CreateBinding(context.Context, *CreateBindingRequest) (*CreateBindingResponse, error)
	SetReceivers(context.Context, *SetReceiversRequest) (*SetReceiversResponse, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	PageMessage(context.Context, *PageMessageRequest) (*PageMessageResponse, error)
	MakeMessagesRead(context.Context, *MakeMessagesReadRequest) (*MakeMessagesReadResponse, error)
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}
func (UnimplementedMessageServiceServer) GetTemplate(context.Context, *GetTemplateRequest) (*GetTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplate not implemented")
}
func (UnimplementedMessageServiceServer) GetTemplates(context.Context, *GetTemplatesRequest) (*GetTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplates not implemented")
}
func (UnimplementedMessageServiceServer) CreateBinding(context.Context, *CreateBindingRequest) (*CreateBindingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBinding not implemented")
}
func (UnimplementedMessageServiceServer) SetReceivers(context.Context, *SetReceiversRequest) (*SetReceiversResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetReceivers not implemented")
}
func (UnimplementedMessageServiceServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessageServiceServer) PageMessage(context.Context, *PageMessageRequest) (*PageMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageMessage not implemented")
}
func (UnimplementedMessageServiceServer) MakeMessagesRead(context.Context, *MakeMessagesReadRequest) (*MakeMessagesReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeMessagesRead not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_CreateTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).CreateTemplate(ctx, req.(*CreateTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_GetTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetTemplate(ctx, req.(*GetTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_GetTemplates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetTemplates(ctx, req.(*GetTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_CreateBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).CreateBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_CreateBinding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).CreateBinding(ctx, req.(*CreateBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SetReceivers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReceiversRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SetReceivers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_SetReceivers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SetReceivers(ctx, req.(*SetReceiversRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_PageMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).PageMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_PageMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).PageMessage(ctx, req.(*PageMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_MakeMessagesRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeMessagesReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).MakeMessagesRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_MakeMessagesRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).MakeMessagesRead(ctx, req.(*MakeMessagesReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messager.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTemplate",
			Handler:    _MessageService_CreateTemplate_Handler,
		},
		{
			MethodName: "GetTemplate",
			Handler:    _MessageService_GetTemplate_Handler,
		},
		{
			MethodName: "GetTemplates",
			Handler:    _MessageService_GetTemplates_Handler,
		},
		{
			MethodName: "CreateBinding",
			Handler:    _MessageService_CreateBinding_Handler,
		},
		{
			MethodName: "SetReceivers",
			Handler:    _MessageService_SetReceivers_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _MessageService_SendMessage_Handler,
		},
		{
			MethodName: "PageMessage",
			Handler:    _MessageService_PageMessage_Handler,
		},
		{
			MethodName: "MakeMessagesRead",
			Handler:    _MessageService_MakeMessagesRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
