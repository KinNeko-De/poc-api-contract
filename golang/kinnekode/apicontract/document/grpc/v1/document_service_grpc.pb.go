// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0
// source: kinnekode/apicontract/document/grpc/v1/document_service.proto

package v1

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

// DocumentServiceClient is the client API for DocumentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentServiceClient interface {
	UploadDocument(ctx context.Context, opts ...grpc.CallOption) (DocumentService_UploadDocumentClient, error)
	DownloadDocument(ctx context.Context, in *DownloadDocumentRequest, opts ...grpc.CallOption) (DocumentService_DownloadDocumentClient, error)
}

type documentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentServiceClient(cc grpc.ClientConnInterface) DocumentServiceClient {
	return &documentServiceClient{cc}
}

func (c *documentServiceClient) UploadDocument(ctx context.Context, opts ...grpc.CallOption) (DocumentService_UploadDocumentClient, error) {
	stream, err := c.cc.NewStream(ctx, &DocumentService_ServiceDesc.Streams[0], "/kinnekode.apicontract.document.grpc.v1.DocumentService/UploadDocument", opts...)
	if err != nil {
		return nil, err
	}
	x := &documentServiceUploadDocumentClient{stream}
	return x, nil
}

type DocumentService_UploadDocumentClient interface {
	Send(*UploadDocumentRequest) error
	CloseAndRecv() (*UploadDocumentResponse, error)
	grpc.ClientStream
}

type documentServiceUploadDocumentClient struct {
	grpc.ClientStream
}

func (x *documentServiceUploadDocumentClient) Send(m *UploadDocumentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *documentServiceUploadDocumentClient) CloseAndRecv() (*UploadDocumentResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadDocumentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *documentServiceClient) DownloadDocument(ctx context.Context, in *DownloadDocumentRequest, opts ...grpc.CallOption) (DocumentService_DownloadDocumentClient, error) {
	stream, err := c.cc.NewStream(ctx, &DocumentService_ServiceDesc.Streams[1], "/kinnekode.apicontract.document.grpc.v1.DocumentService/DownloadDocument", opts...)
	if err != nil {
		return nil, err
	}
	x := &documentServiceDownloadDocumentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DocumentService_DownloadDocumentClient interface {
	Recv() (*DownloadDocumentResponse, error)
	grpc.ClientStream
}

type documentServiceDownloadDocumentClient struct {
	grpc.ClientStream
}

func (x *documentServiceDownloadDocumentClient) Recv() (*DownloadDocumentResponse, error) {
	m := new(DownloadDocumentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DocumentServiceServer is the server API for DocumentService service.
// All implementations must embed UnimplementedDocumentServiceServer
// for forward compatibility
type DocumentServiceServer interface {
	UploadDocument(DocumentService_UploadDocumentServer) error
	DownloadDocument(*DownloadDocumentRequest, DocumentService_DownloadDocumentServer) error
	mustEmbedUnimplementedDocumentServiceServer()
}

// UnimplementedDocumentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDocumentServiceServer struct {
}

func (UnimplementedDocumentServiceServer) UploadDocument(DocumentService_UploadDocumentServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadDocument not implemented")
}
func (UnimplementedDocumentServiceServer) DownloadDocument(*DownloadDocumentRequest, DocumentService_DownloadDocumentServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadDocument not implemented")
}
func (UnimplementedDocumentServiceServer) mustEmbedUnimplementedDocumentServiceServer() {}

// UnsafeDocumentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentServiceServer will
// result in compilation errors.
type UnsafeDocumentServiceServer interface {
	mustEmbedUnimplementedDocumentServiceServer()
}

func RegisterDocumentServiceServer(s grpc.ServiceRegistrar, srv DocumentServiceServer) {
	s.RegisterService(&DocumentService_ServiceDesc, srv)
}

func _DocumentService_UploadDocument_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DocumentServiceServer).UploadDocument(&documentServiceUploadDocumentServer{stream})
}

type DocumentService_UploadDocumentServer interface {
	SendAndClose(*UploadDocumentResponse) error
	Recv() (*UploadDocumentRequest, error)
	grpc.ServerStream
}

type documentServiceUploadDocumentServer struct {
	grpc.ServerStream
}

func (x *documentServiceUploadDocumentServer) SendAndClose(m *UploadDocumentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *documentServiceUploadDocumentServer) Recv() (*UploadDocumentRequest, error) {
	m := new(UploadDocumentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DocumentService_DownloadDocument_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadDocumentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DocumentServiceServer).DownloadDocument(m, &documentServiceDownloadDocumentServer{stream})
}

type DocumentService_DownloadDocumentServer interface {
	Send(*DownloadDocumentResponse) error
	grpc.ServerStream
}

type documentServiceDownloadDocumentServer struct {
	grpc.ServerStream
}

func (x *documentServiceDownloadDocumentServer) Send(m *DownloadDocumentResponse) error {
	return x.ServerStream.SendMsg(m)
}

// DocumentService_ServiceDesc is the grpc.ServiceDesc for DocumentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kinnekode.apicontract.document.grpc.v1.DocumentService",
	HandlerType: (*DocumentServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadDocument",
			Handler:       _DocumentService_UploadDocument_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadDocument",
			Handler:       _DocumentService_DownloadDocument_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kinnekode/apicontract/document/grpc/v1/document_service.proto",
}
