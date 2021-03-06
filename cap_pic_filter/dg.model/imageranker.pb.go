// Code generated by protoc-gen-go.
// source: imageranker.proto
// DO NOT EDIT!

package dg_model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CutboardObject struct {
	ObjectId int64     `protobuf:"varint,1,opt,name=ObjectId,json=objectId" json:"ObjectId"`
	Type     ObjType   `protobuf:"varint,2,opt,name=Type,json=type,enum=dg.model.ObjType" json:"Type"`
	Rect     *Cutboard `protobuf:"bytes,3,opt,name=Rect,json=rect" json:"Rect"`
}

func (m *CutboardObject) Reset()                    { *m = CutboardObject{} }
func (m *CutboardObject) String() string            { return proto.CompactTextString(m) }
func (*CutboardObject) ProtoMessage()               {}
func (*CutboardObject) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *CutboardObject) GetRect() *Cutboard {
	if m != nil {
		return m.Rect
	}
	return nil
}

// ////////////// INDEX ///////////////////////
type ImageRankerIndexData struct {
	Image   *Image            `protobuf:"bytes,1,opt,name=Image,json=image" json:"Image"`
	Objects []*CutboardObject `protobuf:"bytes,2,rep,name=Objects,json=objects" json:"Objects"`
}

func (m *ImageRankerIndexData) Reset()                    { *m = ImageRankerIndexData{} }
func (m *ImageRankerIndexData) String() string            { return proto.CompactTextString(m) }
func (*ImageRankerIndexData) ProtoMessage()               {}
func (*ImageRankerIndexData) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *ImageRankerIndexData) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *ImageRankerIndexData) GetObjects() []*CutboardObject {
	if m != nil {
		return m.Objects
	}
	return nil
}

type ImageRankerIndexResult struct {
	IndexedObjects []int64 `protobuf:"varint,1,rep,packed,name=IndexedObjects,json=indexedObjects" json:"IndexedObjects"`
}

func (m *ImageRankerIndexResult) Reset()                    { *m = ImageRankerIndexResult{} }
func (m *ImageRankerIndexResult) String() string            { return proto.CompactTextString(m) }
func (*ImageRankerIndexResult) ProtoMessage()               {}
func (*ImageRankerIndexResult) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

type ImageRankerIndexRequest struct {
	Data *ImageRankerIndexData `protobuf:"bytes,1,opt,name=Data,json=data" json:"Data"`
}

func (m *ImageRankerIndexRequest) Reset()                    { *m = ImageRankerIndexRequest{} }
func (m *ImageRankerIndexRequest) String() string            { return proto.CompactTextString(m) }
func (*ImageRankerIndexRequest) ProtoMessage()               {}
func (*ImageRankerIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *ImageRankerIndexRequest) GetData() *ImageRankerIndexData {
	if m != nil {
		return m.Data
	}
	return nil
}

type ImageRankerIndexResponse struct {
	Result *ImageRankerIndexResult `protobuf:"bytes,1,opt,name=Result,json=result" json:"Result"`
}

func (m *ImageRankerIndexResponse) Reset()                    { *m = ImageRankerIndexResponse{} }
func (m *ImageRankerIndexResponse) String() string            { return proto.CompactTextString(m) }
func (*ImageRankerIndexResponse) ProtoMessage()               {}
func (*ImageRankerIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *ImageRankerIndexResponse) GetResult() *ImageRankerIndexResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type ImageRankerIndexBatchRequest struct {
	DataList []*ImageRankerIndexData `protobuf:"bytes,1,rep,name=DataList,json=dataList" json:"DataList"`
}

func (m *ImageRankerIndexBatchRequest) Reset()                    { *m = ImageRankerIndexBatchRequest{} }
func (m *ImageRankerIndexBatchRequest) String() string            { return proto.CompactTextString(m) }
func (*ImageRankerIndexBatchRequest) ProtoMessage()               {}
func (*ImageRankerIndexBatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

func (m *ImageRankerIndexBatchRequest) GetDataList() []*ImageRankerIndexData {
	if m != nil {
		return m.DataList
	}
	return nil
}

type ImageRankerIndexBatchResponse struct {
	Results []*ImageRankerIndexResult `protobuf:"bytes,1,rep,name=Results,json=results" json:"Results"`
}

func (m *ImageRankerIndexBatchResponse) Reset()                    { *m = ImageRankerIndexBatchResponse{} }
func (m *ImageRankerIndexBatchResponse) String() string            { return proto.CompactTextString(m) }
func (*ImageRankerIndexBatchResponse) ProtoMessage()               {}
func (*ImageRankerIndexBatchResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

func (m *ImageRankerIndexBatchResponse) GetResults() []*ImageRankerIndexResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// ////////////// QUERY ///////////////////////
type ImageRankerQueryData struct {
	Type       ObjType     `protobuf:"varint,1,opt,name=Type,json=type,enum=dg.model.ObjType" json:"Type"`
	Image      *Image      `protobuf:"bytes,2,opt,name=Image,json=image" json:"Image"`
	CropArea   *Cutboard   `protobuf:"bytes,3,opt,name=CropArea,json=cropArea" json:"CropArea"`
	FocusAreas []*Cutboard `protobuf:"bytes,4,rep,name=FocusAreas,json=focusAreas" json:"FocusAreas"`
	Limit      int32       `protobuf:"varint,5,opt,name=Limit,json=limit" json:"Limit"`
}

func (m *ImageRankerQueryData) Reset()                    { *m = ImageRankerQueryData{} }
func (m *ImageRankerQueryData) String() string            { return proto.CompactTextString(m) }
func (*ImageRankerQueryData) ProtoMessage()               {}
func (*ImageRankerQueryData) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{7} }

func (m *ImageRankerQueryData) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *ImageRankerQueryData) GetCropArea() *Cutboard {
	if m != nil {
		return m.CropArea
	}
	return nil
}

func (m *ImageRankerQueryData) GetFocusAreas() []*Cutboard {
	if m != nil {
		return m.FocusAreas
	}
	return nil
}

type ImageRankerQueryResult struct {
	AllMatchCount int64     `protobuf:"varint,1,opt,name=AllMatchCount,json=allMatchCount" json:"AllMatchCount"`
	DataIds       []int64   `protobuf:"varint,2,rep,packed,name=DataIds,json=dataIds" json:"DataIds"`
	Scores        []float32 `protobuf:"fixed32,3,rep,packed,name=Scores,json=scores" json:"Scores"`
}

func (m *ImageRankerQueryResult) Reset()                    { *m = ImageRankerQueryResult{} }
func (m *ImageRankerQueryResult) String() string            { return proto.CompactTextString(m) }
func (*ImageRankerQueryResult) ProtoMessage()               {}
func (*ImageRankerQueryResult) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{8} }

type ImageRankerQueryRequest struct {
	Data *ImageRankerQueryData `protobuf:"bytes,1,opt,name=Data,json=data" json:"Data"`
}

func (m *ImageRankerQueryRequest) Reset()                    { *m = ImageRankerQueryRequest{} }
func (m *ImageRankerQueryRequest) String() string            { return proto.CompactTextString(m) }
func (*ImageRankerQueryRequest) ProtoMessage()               {}
func (*ImageRankerQueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{9} }

func (m *ImageRankerQueryRequest) GetData() *ImageRankerQueryData {
	if m != nil {
		return m.Data
	}
	return nil
}

type ImageRankerQueryResponse struct {
	Result *ImageRankerQueryResult `protobuf:"bytes,1,opt,name=Result,json=result" json:"Result"`
}

func (m *ImageRankerQueryResponse) Reset()                    { *m = ImageRankerQueryResponse{} }
func (m *ImageRankerQueryResponse) String() string            { return proto.CompactTextString(m) }
func (*ImageRankerQueryResponse) ProtoMessage()               {}
func (*ImageRankerQueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{10} }

func (m *ImageRankerQueryResponse) GetResult() *ImageRankerQueryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type ImageRankerQueryBatchRequest struct {
	DataList []*ImageRankerQueryData `protobuf:"bytes,1,rep,name=DataList,json=dataList" json:"DataList"`
}

func (m *ImageRankerQueryBatchRequest) Reset()                    { *m = ImageRankerQueryBatchRequest{} }
func (m *ImageRankerQueryBatchRequest) String() string            { return proto.CompactTextString(m) }
func (*ImageRankerQueryBatchRequest) ProtoMessage()               {}
func (*ImageRankerQueryBatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{11} }

func (m *ImageRankerQueryBatchRequest) GetDataList() []*ImageRankerQueryData {
	if m != nil {
		return m.DataList
	}
	return nil
}

type ImageRankerQueryBatchResponse struct {
	Results []*ImageRankerQueryResult `protobuf:"bytes,2,rep,name=Results,json=results" json:"Results"`
}

func (m *ImageRankerQueryBatchResponse) Reset()                    { *m = ImageRankerQueryBatchResponse{} }
func (m *ImageRankerQueryBatchResponse) String() string            { return proto.CompactTextString(m) }
func (*ImageRankerQueryBatchResponse) ProtoMessage()               {}
func (*ImageRankerQueryBatchResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{12} }

func (m *ImageRankerQueryBatchResponse) GetResults() []*ImageRankerQueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*CutboardObject)(nil), "dg.model.CutboardObject")
	proto.RegisterType((*ImageRankerIndexData)(nil), "dg.model.ImageRankerIndexData")
	proto.RegisterType((*ImageRankerIndexResult)(nil), "dg.model.ImageRankerIndexResult")
	proto.RegisterType((*ImageRankerIndexRequest)(nil), "dg.model.ImageRankerIndexRequest")
	proto.RegisterType((*ImageRankerIndexResponse)(nil), "dg.model.ImageRankerIndexResponse")
	proto.RegisterType((*ImageRankerIndexBatchRequest)(nil), "dg.model.ImageRankerIndexBatchRequest")
	proto.RegisterType((*ImageRankerIndexBatchResponse)(nil), "dg.model.ImageRankerIndexBatchResponse")
	proto.RegisterType((*ImageRankerQueryData)(nil), "dg.model.ImageRankerQueryData")
	proto.RegisterType((*ImageRankerQueryResult)(nil), "dg.model.ImageRankerQueryResult")
	proto.RegisterType((*ImageRankerQueryRequest)(nil), "dg.model.ImageRankerQueryRequest")
	proto.RegisterType((*ImageRankerQueryResponse)(nil), "dg.model.ImageRankerQueryResponse")
	proto.RegisterType((*ImageRankerQueryBatchRequest)(nil), "dg.model.ImageRankerQueryBatchRequest")
	proto.RegisterType((*ImageRankerQueryBatchResponse)(nil), "dg.model.ImageRankerQueryBatchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ImageRankerService service

type ImageRankerServiceClient interface {
	Index(ctx context.Context, in *ImageRankerIndexRequest, opts ...grpc.CallOption) (*ImageRankerIndexResponse, error)
	BatchIndex(ctx context.Context, in *ImageRankerIndexBatchRequest, opts ...grpc.CallOption) (*ImageRankerIndexBatchResponse, error)
	Query(ctx context.Context, in *ImageRankerQueryRequest, opts ...grpc.CallOption) (*ImageRankerQueryResponse, error)
	BatchQuery(ctx context.Context, in *ImageRankerQueryBatchRequest, opts ...grpc.CallOption) (*ImageRankerQueryBatchResponse, error)
}

type imageRankerServiceClient struct {
	cc *grpc.ClientConn
}

func NewImageRankerServiceClient(cc *grpc.ClientConn) ImageRankerServiceClient {
	return &imageRankerServiceClient{cc}
}

func (c *imageRankerServiceClient) Index(ctx context.Context, in *ImageRankerIndexRequest, opts ...grpc.CallOption) (*ImageRankerIndexResponse, error) {
	out := new(ImageRankerIndexResponse)
	err := grpc.Invoke(ctx, "/dg.model.ImageRankerService/Index", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageRankerServiceClient) BatchIndex(ctx context.Context, in *ImageRankerIndexBatchRequest, opts ...grpc.CallOption) (*ImageRankerIndexBatchResponse, error) {
	out := new(ImageRankerIndexBatchResponse)
	err := grpc.Invoke(ctx, "/dg.model.ImageRankerService/BatchIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageRankerServiceClient) Query(ctx context.Context, in *ImageRankerQueryRequest, opts ...grpc.CallOption) (*ImageRankerQueryResponse, error) {
	out := new(ImageRankerQueryResponse)
	err := grpc.Invoke(ctx, "/dg.model.ImageRankerService/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageRankerServiceClient) BatchQuery(ctx context.Context, in *ImageRankerQueryBatchRequest, opts ...grpc.CallOption) (*ImageRankerQueryBatchResponse, error) {
	out := new(ImageRankerQueryBatchResponse)
	err := grpc.Invoke(ctx, "/dg.model.ImageRankerService/BatchQuery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ImageRankerService service

type ImageRankerServiceServer interface {
	Index(context.Context, *ImageRankerIndexRequest) (*ImageRankerIndexResponse, error)
	BatchIndex(context.Context, *ImageRankerIndexBatchRequest) (*ImageRankerIndexBatchResponse, error)
	Query(context.Context, *ImageRankerQueryRequest) (*ImageRankerQueryResponse, error)
	BatchQuery(context.Context, *ImageRankerQueryBatchRequest) (*ImageRankerQueryBatchResponse, error)
}

func RegisterImageRankerServiceServer(s *grpc.Server, srv ImageRankerServiceServer) {
	s.RegisterService(&_ImageRankerService_serviceDesc, srv)
}

func _ImageRankerService_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRankerIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageRankerServiceServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.model.ImageRankerService/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageRankerServiceServer).Index(ctx, req.(*ImageRankerIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageRankerService_BatchIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRankerIndexBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageRankerServiceServer).BatchIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.model.ImageRankerService/BatchIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageRankerServiceServer).BatchIndex(ctx, req.(*ImageRankerIndexBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageRankerService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRankerQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageRankerServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.model.ImageRankerService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageRankerServiceServer).Query(ctx, req.(*ImageRankerQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageRankerService_BatchQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRankerQueryBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageRankerServiceServer).BatchQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.model.ImageRankerService/BatchQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageRankerServiceServer).BatchQuery(ctx, req.(*ImageRankerQueryBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImageRankerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dg.model.ImageRankerService",
	HandlerType: (*ImageRankerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _ImageRankerService_Index_Handler,
		},
		{
			MethodName: "BatchIndex",
			Handler:    _ImageRankerService_BatchIndex_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _ImageRankerService_Query_Handler,
		},
		{
			MethodName: "BatchQuery",
			Handler:    _ImageRankerService_BatchQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor8,
}

func init() { proto.RegisterFile("imageranker.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x95, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xd7, 0x26, 0x69, 0xa2, 0x37, 0x56, 0x34, 0x6b, 0x1a, 0x51, 0x05, 0x68, 0x58, 0x30,
	0x76, 0xea, 0xa1, 0x5c, 0xd0, 0x4e, 0x8c, 0x22, 0xa4, 0x4a, 0x9b, 0x00, 0x6f, 0x27, 0x38, 0xa0,
	0x34, 0x31, 0x23, 0xd0, 0xd6, 0x9d, 0xed, 0x20, 0x26, 0xbe, 0x2b, 0x1f, 0x84, 0x13, 0x8e, 0xed,
	0xb4, 0x49, 0xdb, 0xa4, 0x83, 0x4b, 0x37, 0xbb, 0x7f, 0xbf, 0xf7, 0x7b, 0xce, 0x2f, 0x2a, 0xec,
	0xa7, 0xd3, 0xe8, 0x9a, 0xf2, 0x68, 0xf6, 0x9d, 0xf2, 0xfe, 0x9c, 0x33, 0xc9, 0x50, 0x90, 0x5c,
	0xf7, 0xa7, 0x2c, 0xa1, 0x93, 0xde, 0xbd, 0x98, 0x4d, 0xa7, 0x6c, 0x66, 0xf6, 0xf1, 0x2f, 0xe8,
	0x0e, 0x33, 0x39, 0x66, 0x11, 0x4f, 0xde, 0x8d, 0xbf, 0xd1, 0x58, 0xa2, 0x1e, 0x04, 0xe6, 0xbf,
	0x51, 0x12, 0xb6, 0x8e, 0x5a, 0x27, 0x0e, 0x09, 0x98, 0x5d, 0xa3, 0x67, 0xe0, 0x5e, 0xdd, 0xce,
	0x69, 0xd8, 0x56, 0xfb, 0xdd, 0xc1, 0x7e, 0xbf, 0x28, 0xda, 0x57, 0x27, 0xf2, 0x2f, 0x88, 0x2b,
	0xd5, 0x27, 0x3a, 0x06, 0x97, 0xa8, 0x03, 0xa1, 0xa3, 0x62, 0xbb, 0x03, 0xb4, 0x8c, 0x15, 0xad,
	0x88, 0xcb, 0xd5, 0xf7, 0xf8, 0x06, 0x0e, 0x46, 0x39, 0x29, 0xd1, 0xa4, 0xa3, 0x59, 0x42, 0x7f,
	0xbe, 0x89, 0x64, 0xa4, 0xda, 0x78, 0x7a, 0x5f, 0xf7, 0xdf, 0x1d, 0xdc, 0x5f, 0x16, 0x30, 0x71,
	0x4f, 0xcf, 0x87, 0x06, 0xe0, 0x1b, 0x52, 0xa1, 0x80, 0x1c, 0x15, 0x0c, 0xd7, 0x3b, 0x99, 0x00,
	0xf1, 0xcd, 0x08, 0x02, 0xbf, 0x82, 0xc3, 0xd5, 0x96, 0x84, 0x8a, 0x6c, 0x22, 0x15, 0x74, 0x57,
	0x2f, 0x69, 0x52, 0x14, 0x6d, 0xa9, 0xa2, 0x0e, 0xe9, 0xa6, 0x95, 0x5d, 0x7c, 0x01, 0x0f, 0xd6,
	0x2b, 0xdc, 0x64, 0x54, 0x48, 0x05, 0xe4, 0xe6, 0xfc, 0x16, 0xfb, 0xf1, 0x2a, 0x76, 0x75, 0x4a,
	0xe2, 0x26, 0xea, 0x13, 0x5f, 0x41, 0xb8, 0x01, 0x68, 0xce, 0x66, 0x82, 0xa2, 0x97, 0xd0, 0x31,
	0x70, 0xb6, 0xe2, 0x51, 0x7d, 0x45, 0x93, 0x23, 0x1d, 0xae, 0xff, 0xe2, 0x8f, 0xf0, 0x70, 0x35,
	0xf1, 0x3a, 0x92, 0xf1, 0xd7, 0x82, 0xf4, 0x14, 0x82, 0x9c, 0xe1, 0x3c, 0x15, 0x52, 0x8f, 0xb9,
	0x9d, 0x36, 0x48, 0x6c, 0x1e, 0x7f, 0x82, 0x47, 0x35, 0xb5, 0x2d, 0xf6, 0x29, 0xf8, 0x06, 0x47,
	0xd8, 0xda, 0xdb, 0xb9, 0x7d, 0xc3, 0x2d, 0xf0, 0xef, 0x56, 0xc5, 0x89, 0x0f, 0x19, 0xe5, 0xb7,
	0xd6, 0x09, 0xa3, 0x5e, 0xab, 0x59, 0xbd, 0x85, 0x3a, 0xed, 0x46, 0x75, 0xfa, 0x10, 0x0c, 0x39,
	0x9b, 0x9f, 0x71, 0x1a, 0x35, 0x58, 0x1a, 0xc4, 0x36, 0xa3, 0x9e, 0x2c, 0xbc, 0x65, 0x71, 0x26,
	0xf2, 0x85, 0x08, 0x5d, 0x3d, 0xd5, 0xa6, 0x13, 0xf0, 0x65, 0x91, 0x42, 0x07, 0xe0, 0x9d, 0xa7,
	0xd3, 0x54, 0x86, 0x9e, 0x6a, 0xe0, 0x11, 0x6f, 0x92, 0x2f, 0xf0, 0xbc, 0x22, 0xa0, 0x9e, 0xcf,
	0x0a, 0xf8, 0x14, 0xf6, 0xce, 0x26, 0x93, 0x8b, 0xfc, 0x2a, 0x87, 0x2c, 0x9b, 0x49, 0xfb, 0xf6,
	0xed, 0x45, 0xe5, 0x4d, 0x14, 0x82, 0x9f, 0xdf, 0xc7, 0x28, 0x31, 0xd2, 0x3b, 0xc4, 0x4f, 0xcc,
	0x12, 0x1d, 0x42, 0xe7, 0x32, 0x66, 0xea, 0x22, 0xd5, 0x44, 0xce, 0x49, 0x9b, 0x74, 0x84, 0x5e,
	0xad, 0x08, 0x6b, 0x3b, 0xde, 0x5d, 0xd8, 0xc5, 0x23, 0xd8, 0x28, 0x6c, 0x31, 0xc0, 0xbf, 0x08,
	0x5b, 0x1a, 0xba, 0x46, 0x58, 0x9d, 0xf8, 0x2f, 0x61, 0x97, 0xb4, 0x75, 0xc2, 0x96, 0x6b, 0xaf,
	0x0b, 0xdb, 0x6e, 0x10, 0xb6, 0xcc, 0x5d, 0x08, 0x3b, 0xf8, 0xd3, 0x06, 0x54, 0xca, 0x5c, 0x52,
	0xfe, 0x23, 0x8d, 0x29, 0x7a, 0xaf, 0x3c, 0xcc, 0xfd, 0x46, 0x4f, 0x9a, 0xdc, 0xd7, 0xb3, 0xf5,
	0x70, 0xe3, 0xeb, 0xa1, 0x11, 0xf1, 0x0e, 0xfa, 0x0c, 0xa0, 0xa9, 0x4d, 0xd9, 0xe3, 0xfa, 0x33,
	0xe5, 0x7b, 0xeb, 0x3d, 0xdf, 0x9a, 0x5b, 0x34, 0x50, 0xc8, 0x7a, 0xc2, 0x1a, 0xe4, 0xb2, 0x38,
	0x35, 0xc8, 0x15, 0x19, 0x4a, 0xc8, 0xa6, 0xec, 0x71, 0xfd, 0x99, 0x3b, 0x20, 0xaf, 0x3f, 0x36,
	0xbc, 0x33, 0xee, 0xe8, 0x1f, 0xb1, 0x17, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x22, 0x6c, 0xa0,
	0xe0, 0xf1, 0x06, 0x00, 0x00,
}
