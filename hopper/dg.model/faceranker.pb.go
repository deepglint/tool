// Code generated by protoc-gen-go.
// source: faceranker.proto
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

type RankRepoSizeRequest struct {
}

func (m *RankRepoSizeRequest) Reset()                    { *m = RankRepoSizeRequest{} }
func (m *RankRepoSizeRequest) String() string            { return proto.CompactTextString(m) }
func (*RankRepoSizeRequest) ProtoMessage()               {}
func (*RankRepoSizeRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type RankRepoSizeResponse struct {
	Context  *RankResponseContext `protobuf:"bytes,1,opt,name=Context,json=context" json:"Context"`
	Size     int32                `protobuf:"varint,2,opt,name=Size,json=size" json:"Size"`
	Capacity int32                `protobuf:"varint,3,opt,name=Capacity,json=capacity" json:"Capacity"`
}

func (m *RankRepoSizeResponse) Reset()                    { *m = RankRepoSizeResponse{} }
func (m *RankRepoSizeResponse) String() string            { return proto.CompactTextString(m) }
func (*RankRepoSizeResponse) ProtoMessage()               {}
func (*RankRepoSizeResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *RankRepoSizeResponse) GetContext() *RankResponseContext {
	if m != nil {
		return m.Context
	}
	return nil
}

type RankRequestContext struct {
	SessionId string            `protobuf:"bytes,1,opt,name=SessionId,json=sessionId" json:"SessionId"`
	UserName  string            `protobuf:"bytes,2,opt,name=UserName,json=userName" json:"UserName"`
	Token     string            `protobuf:"bytes,3,opt,name=Token,json=token" json:"Token"`
	Type      RankType          `protobuf:"varint,4,opt,name=Type,json=type,enum=dg.model.RankType" json:"Type"`
	Params    map[string]string `protobuf:"bytes,5,rep,name=Params,json=params" json:"Params" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *RankRequestContext) Reset()                    { *m = RankRequestContext{} }
func (m *RankRequestContext) String() string            { return proto.CompactTextString(m) }
func (*RankRequestContext) ProtoMessage()               {}
func (*RankRequestContext) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *RankRequestContext) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type RankResponseContext struct {
	SessionId  string           `protobuf:"bytes,1,opt,name=SessionId,json=sessionId" json:"SessionId"`
	Status     string           `protobuf:"bytes,2,opt,name=Status,json=status" json:"Status"`
	Message    string           `protobuf:"bytes,3,opt,name=Message,json=message" json:"Message"`
	RequestTs  *Time            `protobuf:"bytes,4,opt,name=RequestTs,json=requestTs" json:"RequestTs"`
	ResponseTs *Time            `protobuf:"bytes,5,opt,name=ResponseTs,json=responseTs" json:"ResponseTs"`
	DebugTs    map[string]*Time `protobuf:"bytes,6,rep,name=DebugTs,json=debugTs" json:"DebugTs" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *RankResponseContext) Reset()                    { *m = RankResponseContext{} }
func (m *RankResponseContext) String() string            { return proto.CompactTextString(m) }
func (*RankResponseContext) ProtoMessage()               {}
func (*RankResponseContext) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *RankResponseContext) GetRequestTs() *Time {
	if m != nil {
		return m.RequestTs
	}
	return nil
}

func (m *RankResponseContext) GetResponseTs() *Time {
	if m != nil {
		return m.ResponseTs
	}
	return nil
}

func (m *RankResponseContext) GetDebugTs() map[string]*Time {
	if m != nil {
		return m.DebugTs
	}
	return nil
}

type RankFeatureRequest struct {
	Context       *RankRequestContext `protobuf:"bytes,1,opt,name=Context,json=context" json:"Context"`
	Feature       *FeatureVector      `protobuf:"bytes,2,opt,name=Feature,json=feature" json:"Feature"`
	MaxCandidates int32               `protobuf:"varint,3,opt,name=MaxCandidates,json=maxCandidates" json:"MaxCandidates"`
}

func (m *RankFeatureRequest) Reset()                    { *m = RankFeatureRequest{} }
func (m *RankFeatureRequest) String() string            { return proto.CompactTextString(m) }
func (*RankFeatureRequest) ProtoMessage()               {}
func (*RankFeatureRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *RankFeatureRequest) GetContext() *RankRequestContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *RankFeatureRequest) GetFeature() *FeatureVector {
	if m != nil {
		return m.Feature
	}
	return nil
}

type RankItem struct {
	Id         string            `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id"`
	Score      float32           `protobuf:"fixed32,2,opt,name=Score,json=score" json:"Score"`
	Name       string            `protobuf:"bytes,3,opt,name=Name,json=name" json:"Name"`
	URI        string            `protobuf:"bytes,4,opt,name=URI,json=uRI" json:"URI"`
	Data       string            `protobuf:"bytes,5,opt,name=Data,json=data" json:"Data"`
	Attributes map[string]string `protobuf:"bytes,6,rep,name=Attributes,json=attributes" json:"Attributes" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *RankItem) Reset()                    { *m = RankItem{} }
func (m *RankItem) String() string            { return proto.CompactTextString(m) }
func (*RankItem) ProtoMessage()               {}
func (*RankItem) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *RankItem) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type SearchRequest struct {
	Context *RankRequestContext `protobuf:"bytes,1,opt,name=Context,json=context" json:"Context"`
	Col     string              `protobuf:"bytes,2,opt,name=Col,json=col" json:"Col"`
	Key     string              `protobuf:"bytes,3,opt,name=Key,json=key" json:"Key"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *SearchRequest) GetContext() *RankRequestContext {
	if m != nil {
		return m.Context
	}
	return nil
}

type SearchResponse struct {
	Context *RankResponseContext `protobuf:"bytes,1,opt,name=Context,json=context" json:"Context"`
	Results []*RankItem          `protobuf:"bytes,2,rep,name=Results,json=results" json:"Results"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *SearchResponse) GetContext() *RankResponseContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *SearchResponse) GetResults() []*RankItem {
	if m != nil {
		return m.Results
	}
	return nil
}

type RankFeatureResponse struct {
	Context    *RankResponseContext `protobuf:"bytes,1,opt,name=Context,json=context" json:"Context"`
	Candidates []*RankItem          `protobuf:"bytes,2,rep,name=Candidates,json=candidates" json:"Candidates"`
}

func (m *RankFeatureResponse) Reset()                    { *m = RankFeatureResponse{} }
func (m *RankFeatureResponse) String() string            { return proto.CompactTextString(m) }
func (*RankFeatureResponse) ProtoMessage()               {}
func (*RankFeatureResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *RankFeatureResponse) GetContext() *RankResponseContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *RankFeatureResponse) GetCandidates() []*RankItem {
	if m != nil {
		return m.Candidates
	}
	return nil
}

type RankImageRequest struct {
	Context         *RankRequestContext `protobuf:"bytes,1,opt,name=Context,json=context" json:"Context"`
	Image           *Image              `protobuf:"bytes,2,opt,name=Image,json=image" json:"Image"`
	InterestedAreas []*Cutboard         `protobuf:"bytes,3,rep,name=InterestedAreas,json=interestedAreas" json:"InterestedAreas"`
	Candidates      []*FeatureVector    `protobuf:"bytes,4,rep,name=Candidates,json=candidates" json:"Candidates"`
	MaxCandidates   int32               `protobuf:"varint,5,opt,name=MaxCandidates,json=maxCandidates" json:"MaxCandidates"`
}

func (m *RankImageRequest) Reset()                    { *m = RankImageRequest{} }
func (m *RankImageRequest) String() string            { return proto.CompactTextString(m) }
func (*RankImageRequest) ProtoMessage()               {}
func (*RankImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *RankImageRequest) GetContext() *RankRequestContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *RankImageRequest) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *RankImageRequest) GetInterestedAreas() []*Cutboard {
	if m != nil {
		return m.InterestedAreas
	}
	return nil
}

func (m *RankImageRequest) GetCandidates() []*FeatureVector {
	if m != nil {
		return m.Candidates
	}
	return nil
}

type RankImageResponse struct {
	Context    *RankResponseContext `protobuf:"bytes,1,opt,name=Context,json=context" json:"Context"`
	Candidates []*RankItem          `protobuf:"bytes,2,rep,name=Candidates,json=candidates" json:"Candidates"`
}

func (m *RankImageResponse) Reset()                    { *m = RankImageResponse{} }
func (m *RankImageResponse) String() string            { return proto.CompactTextString(m) }
func (*RankImageResponse) ProtoMessage()               {}
func (*RankImageResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

func (m *RankImageResponse) GetContext() *RankResponseContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *RankImageResponse) GetCandidates() []*RankItem {
	if m != nil {
		return m.Candidates
	}
	return nil
}

type GetImageContentRequest struct {
	URI string `protobuf:"bytes,1,opt,name=URI,json=uRI" json:"URI"`
}

func (m *GetImageContentRequest) Reset()                    { *m = GetImageContentRequest{} }
func (m *GetImageContentRequest) String() string            { return proto.CompactTextString(m) }
func (*GetImageContentRequest) ProtoMessage()               {}
func (*GetImageContentRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{11} }

type GetImageContentResponse struct {
	Data string `protobuf:"bytes,1,opt,name=Data,json=data" json:"Data"`
}

func (m *GetImageContentResponse) Reset()                    { *m = GetImageContentResponse{} }
func (m *GetImageContentResponse) String() string            { return proto.CompactTextString(m) }
func (*GetImageContentResponse) ProtoMessage()               {}
func (*GetImageContentResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{12} }

// Deprecated
type FeatureRankingRequest struct {
	ReqId           int64            `protobuf:"varint,1,opt,name=ReqId,json=reqId" json:"ReqId"`
	Type            RecognizeType    `protobuf:"varint,2,opt,name=Type,json=type,enum=dg.model.RecognizeType" json:"Type"`
	Image           *Image           `protobuf:"bytes,3,opt,name=Image,json=image" json:"Image"`
	InterestedAreas []*Cutboard      `protobuf:"bytes,4,rep,name=InterestedAreas,json=interestedAreas" json:"InterestedAreas"`
	DisabledAreas   []*Cutboard      `protobuf:"bytes,5,rep,name=DisabledAreas,json=disabledAreas" json:"DisabledAreas"`
	Candidates      []*FeatureVector `protobuf:"bytes,6,rep,name=Candidates,json=candidates" json:"Candidates"`
	Limit           int32            `protobuf:"varint,7,opt,name=Limit,json=limit" json:"Limit"`
}

func (m *FeatureRankingRequest) Reset()                    { *m = FeatureRankingRequest{} }
func (m *FeatureRankingRequest) String() string            { return proto.CompactTextString(m) }
func (*FeatureRankingRequest) ProtoMessage()               {}
func (*FeatureRankingRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{13} }

func (m *FeatureRankingRequest) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *FeatureRankingRequest) GetInterestedAreas() []*Cutboard {
	if m != nil {
		return m.InterestedAreas
	}
	return nil
}

func (m *FeatureRankingRequest) GetDisabledAreas() []*Cutboard {
	if m != nil {
		return m.DisabledAreas
	}
	return nil
}

func (m *FeatureRankingRequest) GetCandidates() []*FeatureVector {
	if m != nil {
		return m.Candidates
	}
	return nil
}

// Deprecated
type FeatureRankingResponse struct {
	ReqId  int64     `protobuf:"varint,1,opt,name=ReqId,json=reqId" json:"ReqId"`
	Ids    []int64   `protobuf:"varint,2,rep,packed,name=Ids,json=ids" json:"Ids"`
	Scores []float32 `protobuf:"fixed32,3,rep,packed,name=Scores,json=scores" json:"Scores"`
}

func (m *FeatureRankingResponse) Reset()                    { *m = FeatureRankingResponse{} }
func (m *FeatureRankingResponse) String() string            { return proto.CompactTextString(m) }
func (*FeatureRankingResponse) ProtoMessage()               {}
func (*FeatureRankingResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{14} }

type InfoAndFeature struct {
	Info    *RankItem      `protobuf:"bytes,1,opt,name=Info,json=info" json:"Info"`
	Feature *FeatureVector `protobuf:"bytes,2,opt,name=Feature,json=feature" json:"Feature"`
}

func (m *InfoAndFeature) Reset()                    { *m = InfoAndFeature{} }
func (m *InfoAndFeature) String() string            { return proto.CompactTextString(m) }
func (*InfoAndFeature) ProtoMessage()               {}
func (*InfoAndFeature) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{15} }

func (m *InfoAndFeature) GetInfo() *RankItem {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *InfoAndFeature) GetFeature() *FeatureVector {
	if m != nil {
		return m.Feature
	}
	return nil
}

type AddFeaturesRequest struct {
	Context  *RankRequestContext `protobuf:"bytes,1,opt,name=Context,json=context" json:"Context"`
	Features []*InfoAndFeature   `protobuf:"bytes,2,rep,name=Features,json=features" json:"Features"`
}

func (m *AddFeaturesRequest) Reset()                    { *m = AddFeaturesRequest{} }
func (m *AddFeaturesRequest) String() string            { return proto.CompactTextString(m) }
func (*AddFeaturesRequest) ProtoMessage()               {}
func (*AddFeaturesRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{16} }

func (m *AddFeaturesRequest) GetContext() *RankRequestContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *AddFeaturesRequest) GetFeatures() []*InfoAndFeature {
	if m != nil {
		return m.Features
	}
	return nil
}

type AddFeaturesResponse struct {
	Context *RankResponseContext `protobuf:"bytes,1,opt,name=Context,json=context" json:"Context"`
}

func (m *AddFeaturesResponse) Reset()                    { *m = AddFeaturesResponse{} }
func (m *AddFeaturesResponse) String() string            { return proto.CompactTextString(m) }
func (*AddFeaturesResponse) ProtoMessage()               {}
func (*AddFeaturesResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{17} }

func (m *AddFeaturesResponse) GetContext() *RankResponseContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func init() {
	proto.RegisterType((*RankRepoSizeRequest)(nil), "dg.model.RankRepoSizeRequest")
	proto.RegisterType((*RankRepoSizeResponse)(nil), "dg.model.RankRepoSizeResponse")
	proto.RegisterType((*RankRequestContext)(nil), "dg.model.RankRequestContext")
	proto.RegisterType((*RankResponseContext)(nil), "dg.model.RankResponseContext")
	proto.RegisterType((*RankFeatureRequest)(nil), "dg.model.RankFeatureRequest")
	proto.RegisterType((*RankItem)(nil), "dg.model.RankItem")
	proto.RegisterType((*SearchRequest)(nil), "dg.model.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "dg.model.SearchResponse")
	proto.RegisterType((*RankFeatureResponse)(nil), "dg.model.RankFeatureResponse")
	proto.RegisterType((*RankImageRequest)(nil), "dg.model.RankImageRequest")
	proto.RegisterType((*RankImageResponse)(nil), "dg.model.RankImageResponse")
	proto.RegisterType((*GetImageContentRequest)(nil), "dg.model.GetImageContentRequest")
	proto.RegisterType((*GetImageContentResponse)(nil), "dg.model.GetImageContentResponse")
	proto.RegisterType((*FeatureRankingRequest)(nil), "dg.model.FeatureRankingRequest")
	proto.RegisterType((*FeatureRankingResponse)(nil), "dg.model.FeatureRankingResponse")
	proto.RegisterType((*InfoAndFeature)(nil), "dg.model.InfoAndFeature")
	proto.RegisterType((*AddFeaturesRequest)(nil), "dg.model.AddFeaturesRequest")
	proto.RegisterType((*AddFeaturesResponse)(nil), "dg.model.AddFeaturesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for SimilarityService service

type SimilarityServiceClient interface {
	RankImage(ctx context.Context, in *RankImageRequest, opts ...grpc.CallOption) (*RankImageResponse, error)
	RankFeature(ctx context.Context, in *RankFeatureRequest, opts ...grpc.CallOption) (*RankFeatureResponse, error)
	AddFeatures(ctx context.Context, in *AddFeaturesRequest, opts ...grpc.CallOption) (*AddFeaturesResponse, error)
	GetImageContent(ctx context.Context, in *GetImageContentRequest, opts ...grpc.CallOption) (*GetImageContentResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	RankRepoSize(ctx context.Context, in *RankRepoSizeRequest, opts ...grpc.CallOption) (*RankRepoSizeResponse, error)
	// Deprecated
	GetRankedVector(ctx context.Context, in *FeatureRankingRequest, opts ...grpc.CallOption) (*FeatureRankingResponse, error)
}

type similarityServiceClient struct {
	cc *grpc.ClientConn
}

func NewSimilarityServiceClient(cc *grpc.ClientConn) SimilarityServiceClient {
	return &similarityServiceClient{cc}
}

func (c *similarityServiceClient) RankImage(ctx context.Context, in *RankImageRequest, opts ...grpc.CallOption) (*RankImageResponse, error) {
	out := new(RankImageResponse)
	err := grpc.Invoke(ctx, "/dg.model.SimilarityService/RankImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *similarityServiceClient) RankFeature(ctx context.Context, in *RankFeatureRequest, opts ...grpc.CallOption) (*RankFeatureResponse, error) {
	out := new(RankFeatureResponse)
	err := grpc.Invoke(ctx, "/dg.model.SimilarityService/RankFeature", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *similarityServiceClient) AddFeatures(ctx context.Context, in *AddFeaturesRequest, opts ...grpc.CallOption) (*AddFeaturesResponse, error) {
	out := new(AddFeaturesResponse)
	err := grpc.Invoke(ctx, "/dg.model.SimilarityService/AddFeatures", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *similarityServiceClient) GetImageContent(ctx context.Context, in *GetImageContentRequest, opts ...grpc.CallOption) (*GetImageContentResponse, error) {
	out := new(GetImageContentResponse)
	err := grpc.Invoke(ctx, "/dg.model.SimilarityService/GetImageContent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *similarityServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/dg.model.SimilarityService/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *similarityServiceClient) RankRepoSize(ctx context.Context, in *RankRepoSizeRequest, opts ...grpc.CallOption) (*RankRepoSizeResponse, error) {
	out := new(RankRepoSizeResponse)
	err := grpc.Invoke(ctx, "/dg.model.SimilarityService/RankRepoSize", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *similarityServiceClient) GetRankedVector(ctx context.Context, in *FeatureRankingRequest, opts ...grpc.CallOption) (*FeatureRankingResponse, error) {
	out := new(FeatureRankingResponse)
	err := grpc.Invoke(ctx, "/dg.model.SimilarityService/GetRankedVector", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SimilarityService service

type SimilarityServiceServer interface {
	RankImage(context.Context, *RankImageRequest) (*RankImageResponse, error)
	RankFeature(context.Context, *RankFeatureRequest) (*RankFeatureResponse, error)
	AddFeatures(context.Context, *AddFeaturesRequest) (*AddFeaturesResponse, error)
	GetImageContent(context.Context, *GetImageContentRequest) (*GetImageContentResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	RankRepoSize(context.Context, *RankRepoSizeRequest) (*RankRepoSizeResponse, error)
	// Deprecated
	GetRankedVector(context.Context, *FeatureRankingRequest) (*FeatureRankingResponse, error)
}

func RegisterSimilarityServiceServer(s *grpc.Server, srv SimilarityServiceServer) {
	s.RegisterService(&_SimilarityService_serviceDesc, srv)
}

func _SimilarityService_RankImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RankImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimilarityServiceServer).RankImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.model.SimilarityService/RankImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimilarityServiceServer).RankImage(ctx, req.(*RankImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimilarityService_RankFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RankFeatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimilarityServiceServer).RankFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.model.SimilarityService/RankFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimilarityServiceServer).RankFeature(ctx, req.(*RankFeatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimilarityService_AddFeatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFeaturesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimilarityServiceServer).AddFeatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.model.SimilarityService/AddFeatures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimilarityServiceServer).AddFeatures(ctx, req.(*AddFeaturesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimilarityService_GetImageContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimilarityServiceServer).GetImageContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.model.SimilarityService/GetImageContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimilarityServiceServer).GetImageContent(ctx, req.(*GetImageContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimilarityService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimilarityServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.model.SimilarityService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimilarityServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimilarityService_RankRepoSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RankRepoSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimilarityServiceServer).RankRepoSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.model.SimilarityService/RankRepoSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimilarityServiceServer).RankRepoSize(ctx, req.(*RankRepoSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimilarityService_GetRankedVector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeatureRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimilarityServiceServer).GetRankedVector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.model.SimilarityService/GetRankedVector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimilarityServiceServer).GetRankedVector(ctx, req.(*FeatureRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SimilarityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dg.model.SimilarityService",
	HandlerType: (*SimilarityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RankImage",
			Handler:    _SimilarityService_RankImage_Handler,
		},
		{
			MethodName: "RankFeature",
			Handler:    _SimilarityService_RankFeature_Handler,
		},
		{
			MethodName: "AddFeatures",
			Handler:    _SimilarityService_AddFeatures_Handler,
		},
		{
			MethodName: "GetImageContent",
			Handler:    _SimilarityService_GetImageContent_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _SimilarityService_Search_Handler,
		},
		{
			MethodName: "RankRepoSize",
			Handler:    _SimilarityService_RankRepoSize_Handler,
		},
		{
			MethodName: "GetRankedVector",
			Handler:    _SimilarityService_GetRankedVector_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor5,
}

func init() { proto.RegisterFile("faceranker.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 1114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x57, 0x5d, 0x92, 0xdb, 0x44,
	0x10, 0xc6, 0xb2, 0x64, 0xd9, 0xbd, 0xbb, 0xf6, 0x66, 0x92, 0x6c, 0x5c, 0x22, 0x81, 0x45, 0x05,
	0x54, 0x2a, 0x04, 0x57, 0x61, 0x28, 0x08, 0x14, 0xa9, 0x62, 0xf1, 0x12, 0xca, 0x90, 0x04, 0x6a,
	0xec, 0xa4, 0xf2, 0x3a, 0x96, 0x66, 0x8d, 0x6a, 0x2d, 0xc9, 0x48, 0xe3, 0xb0, 0xe6, 0x05, 0x8a,
	0x37, 0x5e, 0xb8, 0x05, 0x87, 0xe0, 0x06, 0x5c, 0x80, 0x63, 0x70, 0x07, 0xe6, 0x4f, 0xf6, 0xc8,
	0x3f, 0x4b, 0x82, 0x1f, 0x78, 0xb1, 0xd5, 0xd3, 0x3d, 0xdd, 0x3d, 0x5f, 0x7f, 0xdd, 0x23, 0xc1,
	0xe1, 0x19, 0x09, 0x68, 0x46, 0x92, 0x73, 0x9a, 0x75, 0xa6, 0x59, 0xca, 0x52, 0x54, 0x0f, 0xc7,
	0x9d, 0x38, 0x0d, 0xe9, 0xc4, 0xdb, 0x0f, 0xd2, 0x38, 0x4e, 0x13, 0xb5, 0xee, 0xed, 0xc7, 0x84,
	0x65, 0xd1, 0x85, 0x92, 0xfc, 0xeb, 0x70, 0x15, 0xf3, 0x5d, 0x98, 0x4e, 0xd3, 0x41, 0xf4, 0x23,
	0xc5, 0xf4, 0xfb, 0x19, 0xcd, 0x99, 0xff, 0x13, 0x5c, 0x2b, 0x2f, 0xe7, 0xd3, 0x34, 0xc9, 0x29,
	0xfa, 0x08, 0xdc, 0x5e, 0x9a, 0x30, 0x7a, 0xc1, 0xda, 0x95, 0xe3, 0xca, 0xed, 0xbd, 0xee, 0xad,
	0x4e, 0x11, 0xa6, 0xa3, 0x36, 0x28, 0x43, 0x6d, 0x84, 0xdd, 0x40, 0x3d, 0x20, 0x04, 0xb6, 0x70,
	0xd4, 0xb6, 0xf8, 0x2e, 0x07, 0xdb, 0x39, 0x7f, 0x46, 0x1e, 0xd4, 0x7b, 0x64, 0x4a, 0x82, 0x88,
	0xcd, 0xdb, 0x55, 0xb9, 0x5e, 0x0f, 0xb4, 0xec, 0xff, 0x6a, 0x01, 0x52, 0x0e, 0x65, 0x42, 0xda,
	0x1f, 0xba, 0x09, 0x8d, 0x01, 0xcd, 0xf3, 0x28, 0x4d, 0xfa, 0xa1, 0xcc, 0xa0, 0x81, 0x1b, 0x79,
	0xb1, 0x20, 0x1c, 0x3e, 0xc9, 0x69, 0xf6, 0x98, 0xc4, 0x2a, 0x50, 0x03, 0xd7, 0x67, 0x5a, 0x46,
	0xd7, 0xc0, 0x19, 0xa6, 0xe7, 0x34, 0x91, 0x91, 0x1a, 0xd8, 0x61, 0x42, 0x40, 0x6f, 0x83, 0x3d,
	0x9c, 0x4f, 0x69, 0xdb, 0xe6, 0x8b, 0xcd, 0x2e, 0x2a, 0x1f, 0x46, 0x68, 0xb0, 0xcd, 0xf8, 0x2f,
	0xfa, 0x0c, 0x6a, 0xdf, 0x92, 0x8c, 0xc4, 0x79, 0xdb, 0x39, 0xae, 0xf2, 0x63, 0xdf, 0x5e, 0x3d,
	0xb6, 0x99, 0x65, 0x47, 0x99, 0x7e, 0x91, 0xb0, 0x6c, 0x8e, 0x6b, 0x53, 0x29, 0x78, 0x1f, 0xc3,
	0x9e, 0xb1, 0x8c, 0x0e, 0xa1, 0x7a, 0x4e, 0xe7, 0xfa, 0x08, 0xe2, 0x51, 0x24, 0xf8, 0x9c, 0x4c,
	0x66, 0x45, 0xe6, 0x4a, 0xf8, 0xc4, 0xba, 0x57, 0xf1, 0xff, 0xb2, 0x8a, 0x22, 0x95, 0xc0, 0xfd,
	0x17, 0x30, 0x8e, 0xa0, 0x36, 0x60, 0x84, 0xcd, 0x72, 0xed, 0xb0, 0x96, 0x4b, 0x09, 0xb5, 0xc1,
	0x7d, 0xc4, 0x8d, 0xc8, 0x98, 0x6a, 0x28, 0xdc, 0x58, 0x89, 0xe8, 0x2e, 0x34, 0xf4, 0x41, 0x86,
	0xb9, 0x44, 0x64, 0xaf, 0xdb, 0x5c, 0x9e, 0x73, 0x18, 0xc5, 0x14, 0x37, 0xb2, 0xc2, 0x00, 0x75,
	0x00, 0x8a, 0x84, 0x86, 0x02, 0x96, 0x4d, 0xe6, 0x90, 0x2d, 0x2c, 0xd0, 0x29, 0xb8, 0xa7, 0x74,
	0x34, 0x1b, 0x73, 0xe3, 0x9a, 0xc4, 0xf0, 0xce, 0xa5, 0xd4, 0xe9, 0x68, 0x63, 0x85, 0xa2, 0x1b,
	0x2a, 0xc9, 0xfb, 0x0a, 0xf6, 0x4d, 0xc5, 0x06, 0x1c, 0xdf, 0x34, 0x71, 0x5c, 0x4f, 0xc9, 0xc0,
	0xf5, 0xf7, 0x8a, 0xe2, 0xd8, 0x03, 0xca, 0x81, 0xc9, 0x0a, 0xee, 0xa3, 0x0f, 0x57, 0x39, 0x7e,
	0xf3, 0xb2, 0x62, 0x2f, 0x29, 0xfe, 0x1e, 0xb8, 0xda, 0x93, 0x0e, 0x7d, 0x63, 0xb9, 0x4f, 0x2b,
	0x9e, 0xd2, 0x80, 0xa5, 0x19, 0x76, 0xcf, 0x94, 0xc8, 0x73, 0x3d, 0x78, 0x44, 0x2e, 0x7a, 0x24,
	0x09, 0xa3, 0x90, 0x30, 0x9a, 0xeb, 0x36, 0x38, 0x88, 0xcd, 0x45, 0xff, 0xef, 0x0a, 0xd4, 0x45,
	0xe0, 0x3e, 0xa3, 0x31, 0x6a, 0x82, 0xb5, 0xa8, 0xb6, 0x15, 0x85, 0x82, 0x36, 0x83, 0x20, 0xd5,
	0x31, 0x2d, 0xec, 0xe4, 0x42, 0x10, 0xed, 0x26, 0xbb, 0x40, 0x55, 0xd8, 0x4e, 0x44, 0x07, 0x70,
	0xa8, 0x9e, 0xe0, 0xbe, 0x2c, 0x2c, 0x87, 0x6a, 0x86, 0xfb, 0xc2, 0xea, 0x94, 0x30, 0x22, 0x8b,
	0xc7, 0xad, 0x78, 0x34, 0x82, 0x3e, 0x07, 0x38, 0x61, 0x7c, 0x40, 0x8c, 0x66, 0x22, 0x1f, 0x55,
	0x29, 0xbf, 0x0c, 0x80, 0xc8, 0xa3, 0xb3, 0x34, 0x52, 0x15, 0x02, 0xb2, 0x58, 0xf0, 0xee, 0x43,
	0x6b, 0x45, 0xfd, 0x52, 0x7c, 0x3f, 0x87, 0x83, 0x01, 0x25, 0x59, 0xf0, 0xdd, 0xae, 0x15, 0xe1,
	0x41, 0x7b, 0xe9, 0x44, 0x07, 0xa8, 0x06, 0xe9, 0x44, 0xac, 0x7c, 0x4d, 0xe7, 0x1a, 0x16, 0x91,
	0x86, 0xff, 0x03, 0x34, 0x8b, 0x60, 0xbb, 0xce, 0xb8, 0xbb, 0xe0, 0x72, 0xdd, 0x6c, 0xc2, 0x44,
	0xcb, 0x09, 0xdc, 0xd0, 0x3a, 0x6e, 0xd8, 0xcd, 0x94, 0x89, 0xff, 0x4b, 0x45, 0x75, 0xf5, 0x82,
	0x7d, 0xbb, 0x86, 0xef, 0x02, 0x18, 0x4c, 0xda, 0x9e, 0x01, 0x04, 0x4b, 0x6a, 0xfd, 0x66, 0xc1,
	0xa1, 0x54, 0xc4, 0x7c, 0x00, 0xec, 0x0a, 0xf7, 0x5b, 0xe0, 0x48, 0x3f, 0x9a, 0xfe, 0xad, 0xe5,
	0x2e, 0xe5, 0xde, 0x89, 0xc4, 0x1f, 0xfa, 0x14, 0x5a, 0x7d, 0xbe, 0x81, 0xe3, 0xc0, 0x68, 0x78,
	0x92, 0x51, 0x22, 0x68, 0xbf, 0x92, 0x6c, 0x6f, 0xc6, 0x46, 0x29, 0xc9, 0x42, 0xdc, 0x8a, 0xca,
	0xa6, 0x1c, 0x1e, 0xf3, 0x94, 0xb6, 0xdc, 0xb8, 0xb5, 0xd1, 0x8c, 0xa3, 0xae, 0xf7, 0x9a, 0xb3,
	0xa9, 0xd7, 0x7e, 0xae, 0xc0, 0x15, 0x03, 0x90, 0xff, 0xa3, 0x26, 0x77, 0xe0, 0xe8, 0x4b, 0xca,
	0x64, 0x02, 0xd2, 0x5f, 0xc2, 0x8a, 0xc2, 0xe8, 0x0e, 0xae, 0x2c, 0x3a, 0xd8, 0x7f, 0x17, 0x6e,
	0xac, 0xd9, 0xea, 0x9c, 0x8b, 0xe6, 0xae, 0x2c, 0x9b, 0xdb, 0xff, 0xd3, 0x82, 0xeb, 0x05, 0xdf,
	0x78, 0xe8, 0x28, 0x19, 0x17, 0xae, 0x79, 0x37, 0xf2, 0x47, 0x3d, 0x59, 0xaa, 0xd8, 0xc9, 0x84,
	0x80, 0xde, 0xd1, 0xd7, 0xa3, 0x25, 0xaf, 0x47, 0x03, 0x66, 0x4c, 0x83, 0x74, 0x9c, 0xf0, 0x4b,
	0xdc, 0xb8, 0x23, 0x17, 0xe5, 0xaf, 0xbe, 0x6c, 0xf9, 0xed, 0x17, 0x2f, 0xff, 0x3d, 0x38, 0x38,
	0x8d, 0x72, 0x32, 0x9a, 0x14, 0x7b, 0x9d, 0xad, 0x7b, 0x0f, 0x42, 0xd3, 0x70, 0x85, 0x38, 0xb5,
	0x17, 0x27, 0x0e, 0x87, 0xe6, 0x61, 0x14, 0x47, 0xac, 0xed, 0x4a, 0xc2, 0x38, 0x13, 0x21, 0xf8,
	0xcf, 0xe0, 0x68, 0x15, 0x49, 0x0d, 0xfc, 0x66, 0x28, 0x79, 0xed, 0xfa, 0xa1, 0xa2, 0x40, 0x15,
	0x57, 0xa3, 0x30, 0x97, 0x17, 0xb4, 0x18, 0xd6, 0x8a, 0xfe, 0x16, 0xbf, 0xa0, 0xa5, 0xc4, 0xc7,
	0x5f, 0xb3, 0x9f, 0x9c, 0xa5, 0x27, 0x49, 0xa8, 0x03, 0x88, 0xb7, 0x14, 0xb1, 0xa2, 0xb9, 0xb7,
	0x89, 0x3f, 0x76, 0xc4, 0xf5, 0xff, 0xe1, 0x06, 0x12, 0x53, 0x08, 0x9d, 0x84, 0x45, 0xa4, 0x7c,
	0xd7, 0x11, 0xf0, 0x01, 0xd4, 0x0b, 0x57, 0x9a, 0xed, 0x6d, 0x83, 0x06, 0xa5, 0x53, 0xe1, 0xba,
	0xce, 0x21, 0xf7, 0x1f, 0xc3, 0xd5, 0x52, 0x0e, 0x3b, 0x76, 0x5d, 0xf7, 0x0f, 0x1b, 0xae, 0x0c,
	0x78, 0x95, 0x26, 0x24, 0xe3, 0xef, 0x92, 0x03, 0x9a, 0x3d, 0x8f, 0x02, 0x8a, 0x1e, 0xf0, 0xd7,
	0x9b, 0xa2, 0xb3, 0x91, 0xb7, 0x02, 0xa2, 0x31, 0xff, 0xbc, 0x57, 0x37, 0xea, 0x54, 0x2c, 0xff,
	0x15, 0xf4, 0x10, 0xf6, 0x8c, 0xb9, 0x8d, 0x56, 0x90, 0x29, 0xbf, 0x4c, 0x78, 0xb7, 0xb6, 0x68,
	0x4d, 0x6f, 0xc6, 0xd9, 0x4d, 0x6f, 0xeb, 0x65, 0x31, 0xbd, 0x6d, 0x00, 0x8c, 0x7b, 0x7b, 0x06,
	0xad, 0x95, 0x79, 0x80, 0x8e, 0x97, 0x7b, 0x36, 0x8f, 0x15, 0xef, 0x8d, 0x4b, 0x2c, 0x16, 0x9e,
	0xef, 0x73, 0xb6, 0xca, 0x7b, 0x12, 0x19, 0xa4, 0x2a, 0x5d, 0xd3, 0x5e, 0x7b, 0x5d, 0xb1, 0xd8,
	0xfe, 0x0d, 0xec, 0x9b, 0x1f, 0x14, 0x68, 0xad, 0x94, 0xa5, 0xef, 0x0f, 0xef, 0xb5, 0x6d, 0xea,
	0x85, 0xc3, 0xa7, 0xf2, 0xa4, 0x42, 0x49, 0x43, 0x45, 0x6a, 0xf4, 0xfa, 0x1a, 0xdb, 0xcb, 0x43,
	0xce, 0x3b, 0xde, 0x6e, 0x50, 0xf8, 0x1d, 0xd5, 0xe4, 0x77, 0xd1, 0xfb, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x54, 0xc2, 0x98, 0xd0, 0x51, 0x0d, 0x00, 0x00,
}
