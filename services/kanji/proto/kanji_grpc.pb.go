// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: services/kanji/proto/kanji.proto

package kanji

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

// KanjiClient is the client API for Kanji service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KanjiClient interface {
	GetKanji(ctx context.Context, in *KanjiRequest, opts ...grpc.CallOption) (*KanjiResponse, error)
	GetKanjiByLevelRange(ctx context.Context, in *KanjiLevelRangeRequest, opts ...grpc.CallOption) (Kanji_GetKanjiByLevelRangeClient, error)
	LoadAllKanji(ctx context.Context, in *WaniKaniTokenRequest, opts ...grpc.CallOption) (*LoadKanjiResponse, error)
}

type kanjiClient struct {
	cc grpc.ClientConnInterface
}

func NewKanjiClient(cc grpc.ClientConnInterface) KanjiClient {
	return &kanjiClient{cc}
}

func (c *kanjiClient) GetKanji(ctx context.Context, in *KanjiRequest, opts ...grpc.CallOption) (*KanjiResponse, error) {
	out := new(KanjiResponse)
	err := c.cc.Invoke(ctx, "/aeriqu.kanikaki.kanji.Kanji/GetKanji", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kanjiClient) GetKanjiByLevelRange(ctx context.Context, in *KanjiLevelRangeRequest, opts ...grpc.CallOption) (Kanji_GetKanjiByLevelRangeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Kanji_ServiceDesc.Streams[0], "/aeriqu.kanikaki.kanji.Kanji/GetKanjiByLevelRange", opts...)
	if err != nil {
		return nil, err
	}
	x := &kanjiGetKanjiByLevelRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Kanji_GetKanjiByLevelRangeClient interface {
	Recv() (*KanjiResponse, error)
	grpc.ClientStream
}

type kanjiGetKanjiByLevelRangeClient struct {
	grpc.ClientStream
}

func (x *kanjiGetKanjiByLevelRangeClient) Recv() (*KanjiResponse, error) {
	m := new(KanjiResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kanjiClient) LoadAllKanji(ctx context.Context, in *WaniKaniTokenRequest, opts ...grpc.CallOption) (*LoadKanjiResponse, error) {
	out := new(LoadKanjiResponse)
	err := c.cc.Invoke(ctx, "/aeriqu.kanikaki.kanji.Kanji/LoadAllKanji", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KanjiServer is the server API for Kanji service.
// All implementations must embed UnimplementedKanjiServer
// for forward compatibility
type KanjiServer interface {
	GetKanji(context.Context, *KanjiRequest) (*KanjiResponse, error)
	GetKanjiByLevelRange(*KanjiLevelRangeRequest, Kanji_GetKanjiByLevelRangeServer) error
	LoadAllKanji(context.Context, *WaniKaniTokenRequest) (*LoadKanjiResponse, error)
	mustEmbedUnimplementedKanjiServer()
}

// UnimplementedKanjiServer must be embedded to have forward compatible implementations.
type UnimplementedKanjiServer struct {
}

func (UnimplementedKanjiServer) GetKanji(context.Context, *KanjiRequest) (*KanjiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKanji not implemented")
}
func (UnimplementedKanjiServer) GetKanjiByLevelRange(*KanjiLevelRangeRequest, Kanji_GetKanjiByLevelRangeServer) error {
	return status.Errorf(codes.Unimplemented, "method GetKanjiByLevelRange not implemented")
}
func (UnimplementedKanjiServer) LoadAllKanji(context.Context, *WaniKaniTokenRequest) (*LoadKanjiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadAllKanji not implemented")
}
func (UnimplementedKanjiServer) mustEmbedUnimplementedKanjiServer() {}

// UnsafeKanjiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KanjiServer will
// result in compilation errors.
type UnsafeKanjiServer interface {
	mustEmbedUnimplementedKanjiServer()
}

func RegisterKanjiServer(s grpc.ServiceRegistrar, srv KanjiServer) {
	s.RegisterService(&Kanji_ServiceDesc, srv)
}

func _Kanji_GetKanji_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KanjiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KanjiServer).GetKanji(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aeriqu.kanikaki.kanji.Kanji/GetKanji",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KanjiServer).GetKanji(ctx, req.(*KanjiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kanji_GetKanjiByLevelRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(KanjiLevelRangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KanjiServer).GetKanjiByLevelRange(m, &kanjiGetKanjiByLevelRangeServer{stream})
}

type Kanji_GetKanjiByLevelRangeServer interface {
	Send(*KanjiResponse) error
	grpc.ServerStream
}

type kanjiGetKanjiByLevelRangeServer struct {
	grpc.ServerStream
}

func (x *kanjiGetKanjiByLevelRangeServer) Send(m *KanjiResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Kanji_LoadAllKanji_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaniKaniTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KanjiServer).LoadAllKanji(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aeriqu.kanikaki.kanji.Kanji/LoadAllKanji",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KanjiServer).LoadAllKanji(ctx, req.(*WaniKaniTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Kanji_ServiceDesc is the grpc.ServiceDesc for Kanji service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Kanji_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aeriqu.kanikaki.kanji.Kanji",
	HandlerType: (*KanjiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKanji",
			Handler:    _Kanji_GetKanji_Handler,
		},
		{
			MethodName: "LoadAllKanji",
			Handler:    _Kanji_LoadAllKanji_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetKanjiByLevelRange",
			Handler:       _Kanji_GetKanjiByLevelRange_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services/kanji/proto/kanji.proto",
}
