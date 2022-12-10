// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: chat.proto

package pb

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

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	Chat(ctx context.Context, opts ...grpc.CallOption) (Chat_ChatClient, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Chat(ctx context.Context, opts ...grpc.CallOption) (Chat_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[0], "/pb.Chat/chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatChatClient{stream}
	return x, nil
}

type Chat_ChatClient interface {
	Send(*MessageReq) error
	Recv() (*MessageRes, error)
	grpc.ClientStream
}

type chatChatClient struct {
	grpc.ClientStream
}

func (x *chatChatClient) Send(m *MessageReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatChatClient) Recv() (*MessageRes, error) {
	m := new(MessageRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	Chat(Chat_ChatServer) error
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) Chat(Chat_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).Chat(&chatChatServer{stream})
}

type Chat_ChatServer interface {
	Send(*MessageRes) error
	Recv() (*MessageReq, error)
	grpc.ServerStream
}

type chatChatServer struct {
	grpc.ServerStream
}

func (x *chatChatServer) Send(m *MessageRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatChatServer) Recv() (*MessageReq, error) {
	m := new(MessageReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "chat",
			Handler:       _Chat_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}