// Code generated by protoc-gen-go.
// source: facesearch.proto
// DO NOT EDIT!

package dg_model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AttrConds int32

const (
	AttrConds_Attr_Unknown  AttrConds = 0
	AttrConds_Attr_Name     AttrConds = 1
	AttrConds_Attr_IdNo     AttrConds = 2
	AttrConds_Attr_GenderId AttrConds = 3
	AttrConds_Attr_AgeId    AttrConds = 4
	AttrConds_Attr_NationId AttrConds = 5
	AttrConds_Attr_Address  AttrConds = 6
	AttrConds_Attr_AgeRange AttrConds = 7
	AttrConds_Attr_IdType   AttrConds = 8
)

var AttrConds_name = map[int32]string{
	0: "Attr_Unknown",
	1: "Attr_Name",
	2: "Attr_IdNo",
	3: "Attr_GenderId",
	4: "Attr_AgeId",
	5: "Attr_NationId",
	6: "Attr_Address",
	7: "Attr_AgeRange",
	8: "Attr_IdType",
}
var AttrConds_value = map[string]int32{
	"Attr_Unknown":  0,
	"Attr_Name":     1,
	"Attr_IdNo":     2,
	"Attr_GenderId": 3,
	"Attr_AgeId":    4,
	"Attr_NationId": 5,
	"Attr_Address":  6,
	"Attr_AgeRange": 7,
	"Attr_IdType":   8,
}

func (x AttrConds) String() string {
	return proto.EnumName(AttrConds_name, int32(x))
}
func (AttrConds) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

type TimeSpacialRangeConds struct {
	SensorIds  []string          `protobuf:"bytes,1,rep,name=SensorIds,json=sensorIds" json:"SensorIds"`
	TimeRanges []*QueryTimeRange `protobuf:"bytes,2,rep,name=TimeRanges,json=timeRanges" json:"TimeRanges"`
}

func (m *TimeSpacialRangeConds) Reset()                    { *m = TimeSpacialRangeConds{} }
func (m *TimeSpacialRangeConds) String() string            { return proto.CompactTextString(m) }
func (*TimeSpacialRangeConds) ProtoMessage()               {}
func (*TimeSpacialRangeConds) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *TimeSpacialRangeConds) GetTimeRanges() []*QueryTimeRange {
	if m != nil {
		return m.TimeRanges
	}
	return nil
}

type BaseConds struct {
	QueryId           string            `protobuf:"bytes,1,opt,name=QueryId,json=queryId" json:"QueryId"`
	Columns           []string          `protobuf:"bytes,6,rep,name=Columns,json=columns" json:"Columns"`
	GroupByDimensions GroupByDimensions `protobuf:"varint,7,opt,name=GroupByDimensions,json=groupByDimensions,enum=dg.model.GroupByDimensions" json:"GroupByDimensions"`
	Offset            int32             `protobuf:"varint,8,opt,name=Offset,json=offset" json:"Offset"`
	Limit             int32             `protobuf:"varint,9,opt,name=Limit,json=limit" json:"Limit"`
	SortBy            string            `protobuf:"bytes,10,opt,name=SortBy,json=sortBy" json:"SortBy"`
	SortOrderAsc      bool              `protobuf:"varint,11,opt,name=SortOrderAsc,json=sortOrderAsc" json:"SortOrderAsc"`
	OutputFmt         DataFmtType       `protobuf:"varint,12,opt,name=OutputFmt,json=outputFmt,enum=dg.model.DataFmtType" json:"OutputFmt"`
}

func (m *BaseConds) Reset()                    { *m = BaseConds{} }
func (m *BaseConds) String() string            { return proto.CompactTextString(m) }
func (*BaseConds) ProtoMessage()               {}
func (*BaseConds) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

type CapturedConds struct {
	Ids                   []string               `protobuf:"bytes,1,rep,name=Ids,json=ids" json:"Ids"`
	TimeSpacialRangeConds *TimeSpacialRangeConds `protobuf:"bytes,2,opt,name=TimeSpacialRangeConds,json=timeSpacialRangeConds" json:"TimeSpacialRangeConds"`
	Conditions            map[string]string      `protobuf:"bytes,3,rep,name=Conditions,json=conditions" json:"Conditions" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CapturedConds) Reset()                    { *m = CapturedConds{} }
func (m *CapturedConds) String() string            { return proto.CompactTextString(m) }
func (*CapturedConds) ProtoMessage()               {}
func (*CapturedConds) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *CapturedConds) GetTimeSpacialRangeConds() *TimeSpacialRangeConds {
	if m != nil {
		return m.TimeSpacialRangeConds
	}
	return nil
}

func (m *CapturedConds) GetConditions() map[string]string {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type RegisteredConds struct {
	Ids        []string          `protobuf:"bytes,1,rep,name=Ids,json=ids" json:"Ids"`
	Repos      []string          `protobuf:"bytes,2,rep,name=Repos,json=repos" json:"Repos"`
	Conditions map[string]string `protobuf:"bytes,3,rep,name=Conditions,json=conditions" json:"Conditions" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *RegisteredConds) Reset()                    { *m = RegisteredConds{} }
func (m *RegisteredConds) String() string            { return proto.CompactTextString(m) }
func (*RegisteredConds) ProtoMessage()               {}
func (*RegisteredConds) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *RegisteredConds) GetConditions() map[string]string {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type WarnedConds struct {
	Ids                   []string               `protobuf:"bytes,1,rep,name=Ids,json=ids" json:"Ids"`
	Repos                 []string               `protobuf:"bytes,2,rep,name=Repos,json=repos" json:"Repos"`
	TimeSpacialRangeConds *TimeSpacialRangeConds `protobuf:"bytes,3,opt,name=TimeSpacialRangeConds,json=timeSpacialRangeConds" json:"TimeSpacialRangeConds"`
	Conditions            map[string]string      `protobuf:"bytes,4,rep,name=Conditions,json=conditions" json:"Conditions" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *WarnedConds) Reset()                    { *m = WarnedConds{} }
func (m *WarnedConds) String() string            { return proto.CompactTextString(m) }
func (*WarnedConds) ProtoMessage()               {}
func (*WarnedConds) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *WarnedConds) GetTimeSpacialRangeConds() *TimeSpacialRangeConds {
	if m != nil {
		return m.TimeSpacialRangeConds
	}
	return nil
}

func (m *WarnedConds) GetConditions() map[string]string {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type AuxiliaryConds struct {
	IsWarned  int32 `protobuf:"varint,1,opt,name=IsWarned,json=isWarned" json:"IsWarned"`
	IsChecked int32 `protobuf:"varint,2,opt,name=IsChecked,json=isChecked" json:"IsChecked"`
	IsRanked  int32 `protobuf:"varint,3,opt,name=IsRanked,json=isRanked" json:"IsRanked"`
	Status    int32 `protobuf:"varint,4,opt,name=Status,json=status" json:"Status"`
}

func (m *AuxiliaryConds) Reset()                    { *m = AuxiliaryConds{} }
func (m *AuxiliaryConds) String() string            { return proto.CompactTextString(m) }
func (*AuxiliaryConds) ProtoMessage()               {}
func (*AuxiliaryConds) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

type ImageConds struct {
	Images     []*ImageResult `protobuf:"bytes,1,rep,name=Images,json=images" json:"Images"`
	Confidence float32        `protobuf:"fixed32,2,opt,name=Confidence,json=confidence" json:"Confidence"`
	CountLimit int32          `protobuf:"varint,3,opt,name=CountLimit,json=countLimit" json:"CountLimit"`
}

func (m *ImageConds) Reset()                    { *m = ImageConds{} }
func (m *ImageConds) String() string            { return proto.CompactTextString(m) }
func (*ImageConds) ProtoMessage()               {}
func (*ImageConds) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *ImageConds) GetImages() []*ImageResult {
	if m != nil {
		return m.Images
	}
	return nil
}

type CapturedFaceQuery struct {
	BaseConds      *BaseConds      `protobuf:"bytes,1,opt,name=BaseConds,json=baseConds" json:"BaseConds"`
	AttrConds      *CapturedConds  `protobuf:"bytes,2,opt,name=AttrConds,json=attrConds" json:"AttrConds"`
	ImageConds     *ImageConds     `protobuf:"bytes,3,opt,name=ImageConds,json=imageConds" json:"ImageConds"`
	AuxiliaryConds *AuxiliaryConds `protobuf:"bytes,4,opt,name=AuxiliaryConds,json=auxiliaryConds" json:"AuxiliaryConds"`
}

func (m *CapturedFaceQuery) Reset()                    { *m = CapturedFaceQuery{} }
func (m *CapturedFaceQuery) String() string            { return proto.CompactTextString(m) }
func (*CapturedFaceQuery) ProtoMessage()               {}
func (*CapturedFaceQuery) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

func (m *CapturedFaceQuery) GetBaseConds() *BaseConds {
	if m != nil {
		return m.BaseConds
	}
	return nil
}

func (m *CapturedFaceQuery) GetAttrConds() *CapturedConds {
	if m != nil {
		return m.AttrConds
	}
	return nil
}

func (m *CapturedFaceQuery) GetImageConds() *ImageConds {
	if m != nil {
		return m.ImageConds
	}
	return nil
}

func (m *CapturedFaceQuery) GetAuxiliaryConds() *AuxiliaryConds {
	if m != nil {
		return m.AuxiliaryConds
	}
	return nil
}

type RegisteredFaceQuery struct {
	BaseConds      *BaseConds       `protobuf:"bytes,1,opt,name=BaseConds,json=baseConds" json:"BaseConds"`
	AttrConds      *RegisteredConds `protobuf:"bytes,2,opt,name=AttrConds,json=attrConds" json:"AttrConds"`
	ImageConds     *ImageConds      `protobuf:"bytes,3,opt,name=ImageConds,json=imageConds" json:"ImageConds"`
	AuxiliaryConds *AuxiliaryConds  `protobuf:"bytes,4,opt,name=AuxiliaryConds,json=auxiliaryConds" json:"AuxiliaryConds"`
}

func (m *RegisteredFaceQuery) Reset()                    { *m = RegisteredFaceQuery{} }
func (m *RegisteredFaceQuery) String() string            { return proto.CompactTextString(m) }
func (*RegisteredFaceQuery) ProtoMessage()               {}
func (*RegisteredFaceQuery) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

func (m *RegisteredFaceQuery) GetBaseConds() *BaseConds {
	if m != nil {
		return m.BaseConds
	}
	return nil
}

func (m *RegisteredFaceQuery) GetAttrConds() *RegisteredConds {
	if m != nil {
		return m.AttrConds
	}
	return nil
}

func (m *RegisteredFaceQuery) GetImageConds() *ImageConds {
	if m != nil {
		return m.ImageConds
	}
	return nil
}

func (m *RegisteredFaceQuery) GetAuxiliaryConds() *AuxiliaryConds {
	if m != nil {
		return m.AuxiliaryConds
	}
	return nil
}

type FaceRuleQuery struct {
	BaseConds      *BaseConds      `protobuf:"bytes,1,opt,name=BaseConds,json=baseConds" json:"BaseConds"`
	SensorName     string          `protobuf:"bytes,2,opt,name=SensorName,json=sensorName" json:"SensorName"`
	AuxiliaryConds *AuxiliaryConds `protobuf:"bytes,3,opt,name=AuxiliaryConds,json=auxiliaryConds" json:"AuxiliaryConds"`
}

func (m *FaceRuleQuery) Reset()                    { *m = FaceRuleQuery{} }
func (m *FaceRuleQuery) String() string            { return proto.CompactTextString(m) }
func (*FaceRuleQuery) ProtoMessage()               {}
func (*FaceRuleQuery) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{9} }

func (m *FaceRuleQuery) GetBaseConds() *BaseConds {
	if m != nil {
		return m.BaseConds
	}
	return nil
}

func (m *FaceRuleQuery) GetAuxiliaryConds() *AuxiliaryConds {
	if m != nil {
		return m.AuxiliaryConds
	}
	return nil
}

type FaceEventQuery struct {
	BaseConds      *BaseConds      `protobuf:"bytes,1,opt,name=BaseConds,json=baseConds" json:"BaseConds"`
	AttrConds      *WarnedConds    `protobuf:"bytes,2,opt,name=AttrConds,json=attrConds" json:"AttrConds"`
	AuxiliaryConds *AuxiliaryConds `protobuf:"bytes,3,opt,name=AuxiliaryConds,json=auxiliaryConds" json:"AuxiliaryConds"`
}

func (m *FaceEventQuery) Reset()                    { *m = FaceEventQuery{} }
func (m *FaceEventQuery) String() string            { return proto.CompactTextString(m) }
func (*FaceEventQuery) ProtoMessage()               {}
func (*FaceEventQuery) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{10} }

func (m *FaceEventQuery) GetBaseConds() *BaseConds {
	if m != nil {
		return m.BaseConds
	}
	return nil
}

func (m *FaceEventQuery) GetAttrConds() *WarnedConds {
	if m != nil {
		return m.AttrConds
	}
	return nil
}

func (m *FaceEventQuery) GetAuxiliaryConds() *AuxiliaryConds {
	if m != nil {
		return m.AuxiliaryConds
	}
	return nil
}

func init() {
	proto.RegisterType((*TimeSpacialRangeConds)(nil), "dg.model.TimeSpacialRangeConds")
	proto.RegisterType((*BaseConds)(nil), "dg.model.BaseConds")
	proto.RegisterType((*CapturedConds)(nil), "dg.model.CapturedConds")
	proto.RegisterType((*RegisteredConds)(nil), "dg.model.RegisteredConds")
	proto.RegisterType((*WarnedConds)(nil), "dg.model.WarnedConds")
	proto.RegisterType((*AuxiliaryConds)(nil), "dg.model.AuxiliaryConds")
	proto.RegisterType((*ImageConds)(nil), "dg.model.ImageConds")
	proto.RegisterType((*CapturedFaceQuery)(nil), "dg.model.CapturedFaceQuery")
	proto.RegisterType((*RegisteredFaceQuery)(nil), "dg.model.RegisteredFaceQuery")
	proto.RegisterType((*FaceRuleQuery)(nil), "dg.model.FaceRuleQuery")
	proto.RegisterType((*FaceEventQuery)(nil), "dg.model.FaceEventQuery")
	proto.RegisterEnum("dg.model.AttrConds", AttrConds_name, AttrConds_value)
}

func init() { proto.RegisterFile("facesearch.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 863 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x56, 0xdd, 0x8a, 0xe4, 0x44,
	0x14, 0x36, 0xc9, 0x76, 0xba, 0x73, 0xba, 0xa7, 0x27, 0x53, 0xbb, 0xa3, 0x71, 0x14, 0x95, 0x80,
	0xf8, 0x03, 0x36, 0x38, 0xa3, 0xb8, 0x08, 0x82, 0x3d, 0xb3, 0xb3, 0x43, 0x40, 0x66, 0xb0, 0x66,
	0x17, 0x2f, 0x97, 0xda, 0xa4, 0xa6, 0x37, 0x4c, 0x92, 0x6a, 0x53, 0x95, 0xd5, 0x46, 0x10, 0x9f,
	0xc1, 0x67, 0x10, 0xc4, 0x57, 0xf0, 0xd6, 0x67, 0xf2, 0xc2, 0x3b, 0xeb, 0xa7, 0xf3, 0xd3, 0x3d,
	0xdd, 0xc2, 0xee, 0xb0, 0xe0, 0x5d, 0xbe, 0xf3, 0x53, 0xe7, 0x3b, 0xe7, 0xd4, 0x39, 0x15, 0xf0,
	0xaf, 0x48, 0x4c, 0x39, 0x25, 0x65, 0xfc, 0x6c, 0x32, 0x2f, 0x99, 0x60, 0x68, 0x90, 0xcc, 0x26,
	0x39, 0x4b, 0x68, 0x76, 0x80, 0x8c, 0x3c, 0x66, 0x79, 0xce, 0x0a, 0xa3, 0x3d, 0x00, 0x65, 0xbf,
	0xfc, 0x1e, 0x75, 0x35, 0x21, 0x83, 0xfd, 0x47, 0x69, 0x4e, 0x2f, 0xe7, 0x24, 0x4e, 0x49, 0x86,
	0x49, 0x31, 0xa3, 0x27, 0xac, 0x48, 0x38, 0x7a, 0x1b, 0xbc, 0x4b, 0x5a, 0x70, 0x56, 0x46, 0x09,
	0x0f, 0xac, 0xf7, 0x9c, 0x0f, 0x3d, 0xec, 0xf1, 0x5a, 0x80, 0xee, 0x03, 0x28, 0x37, 0x6d, 0xcf,
	0x03, 0x5b, 0xaa, 0x87, 0x87, 0xc1, 0xa4, 0xe6, 0x30, 0xf9, 0xb6, 0xa2, 0xe5, 0xa2, 0x31, 0xc0,
	0x20, 0x1a, 0xdb, 0xf0, 0x0f, 0x1b, 0xbc, 0x63, 0xc2, 0x97, 0x51, 0x02, 0xe8, 0x6b, 0xdb, 0x28,
	0x91, 0x31, 0x2c, 0x19, 0xa3, 0xff, 0xbd, 0x81, 0x4a, 0x73, 0xc2, 0xb2, 0x2a, 0x2f, 0x78, 0xe0,
	0xea, 0xe8, 0xfd, 0xd8, 0x40, 0x14, 0xc1, 0xde, 0x59, 0xc9, 0xaa, 0xf9, 0xf1, 0xe2, 0x81, 0x3c,
	0xb6, 0xe0, 0x29, 0x93, 0x36, 0x7d, 0xe9, 0x3d, 0x3e, 0x7c, 0xab, 0xa5, 0x70, 0xc3, 0x04, 0xef,
	0xcd, 0xd6, 0x45, 0xe8, 0x75, 0x70, 0x2f, 0xae, 0xae, 0x38, 0x15, 0xc1, 0x40, 0xfa, 0xf7, 0xb0,
	0xcb, 0x34, 0x42, 0xf7, 0xa0, 0xf7, 0x4d, 0x9a, 0xa7, 0x22, 0xf0, 0xb4, 0xb8, 0x97, 0x29, 0xa0,
	0xac, 0x2f, 0x59, 0x29, 0x8e, 0x17, 0x01, 0x68, 0xae, 0x2e, 0xd7, 0x08, 0x85, 0x30, 0x52, 0xf2,
	0x8b, 0x32, 0xa1, 0xe5, 0x94, 0xc7, 0xc1, 0x50, 0x6a, 0x07, 0x78, 0xc4, 0x3b, 0x32, 0x74, 0x04,
	0xde, 0x45, 0x25, 0xe6, 0x95, 0x78, 0x98, 0x8b, 0x60, 0xa4, 0xc9, 0xee, 0xb7, 0x64, 0x1f, 0x10,
	0x41, 0xa4, 0xe2, 0xd1, 0x62, 0x4e, 0xb1, 0xc7, 0x6a, 0xbb, 0xf0, 0x17, 0x1b, 0x76, 0x4e, 0xc8,
	0x5c, 0x54, 0x25, 0x4d, 0x4c, 0xbd, 0x7c, 0x70, 0xda, 0x7e, 0x38, 0xa9, 0x94, 0x3c, 0xde, 0xd2,
	0x40, 0xd9, 0x14, 0x4b, 0x36, 0xe5, 0xdd, 0x36, 0xc8, 0x46, 0x33, 0xbc, 0x2f, 0x36, 0xb6, 0xff,
	0x0c, 0x40, 0x7d, 0xa4, 0x42, 0x57, 0xd7, 0xd1, 0x0d, 0xfe, 0xa0, 0x3d, 0x6b, 0x85, 0xd5, 0xa4,
	0xb5, 0x3c, 0x2d, 0x44, 0xb9, 0xc0, 0x10, 0x37, 0x82, 0x83, 0xaf, 0x60, 0x77, 0x4d, 0xad, 0x92,
	0xb8, 0xa6, 0x8b, 0x65, 0xc3, 0xd5, 0xa7, 0xaa, 0xf7, 0x73, 0x92, 0x55, 0x54, 0x93, 0xf6, 0xb0,
	0x01, 0x5f, 0xda, 0xf7, 0xad, 0xf0, 0x2f, 0x0b, 0x76, 0x31, 0x9d, 0xa5, 0x5c, 0xd0, 0xff, 0x28,
	0x82, 0xf4, 0xc7, 0x74, 0xce, 0xcc, 0x4d, 0x94, 0xfe, 0xa5, 0x02, 0xf2, 0xa2, 0xdc, 0xcc, 0xe1,
	0xa3, 0x36, 0x87, 0xb5, 0x63, 0x5f, 0x65, 0x16, 0xbf, 0xda, 0x30, 0xfc, 0x8e, 0x94, 0xc5, 0x8b,
	0x66, 0xb0, 0xb5, 0xb9, 0xce, 0xad, 0x9a, 0x7b, 0xba, 0x52, 0x98, 0x3b, 0xba, 0x30, 0xef, 0xb7,
	0x67, 0x75, 0x98, 0xbe, 0xca, 0xa2, 0xfc, 0x0c, 0xe3, 0x69, 0xf5, 0x63, 0x9a, 0xa5, 0xa4, 0x5c,
	0x18, 0x5e, 0x07, 0x30, 0x88, 0xb8, 0x89, 0xae, 0x8f, 0xe8, 0xe1, 0x41, 0xba, 0xc4, 0x6a, 0x1f,
	0x45, 0xfc, 0xe4, 0x19, 0x8d, 0xaf, 0xa5, 0xd2, 0xd6, 0x4a, 0x2f, 0xad, 0x05, 0xc6, 0x53, 0x66,
	0xa8, 0x94, 0x4e, 0xed, 0x69, 0xb0, 0x1e, 0x5b, 0x41, 0x44, 0xa5, 0x32, 0xd5, 0x43, 0xce, 0x35,
	0x0a, 0x7f, 0x02, 0x88, 0x72, 0x52, 0xd7, 0xe4, 0x13, 0x70, 0x35, 0x32, 0x5d, 0x19, 0x76, 0xa7,
	0x53, 0xcb, 0x31, 0xe5, 0x55, 0x26, 0xb0, 0x9b, 0x6a, 0x23, 0xf4, 0x8e, 0x2e, 0xe1, 0x55, 0x9a,
	0xd0, 0x22, 0x36, 0xb9, 0xd9, 0xba, 0x36, 0x4b, 0x89, 0xd1, 0x57, 0x85, 0x30, 0x6b, 0xc4, 0x50,
	0x92, 0xfa, 0x5a, 0x12, 0xfe, 0x6d, 0xc1, 0x5e, 0x3d, 0x44, 0x0f, 0xe5, 0x72, 0xd6, 0x5b, 0x10,
	0x7d, 0xda, 0xd9, 0x8d, 0xba, 0x02, 0xc3, 0xc3, 0xbb, 0x2d, 0x8f, 0x46, 0x85, 0xbd, 0xa7, 0xcd,
	0x06, 0xfd, 0x1c, 0xbc, 0xa9, 0x10, 0x65, 0x77, 0xe6, 0xdf, 0xd8, 0x32, 0xa7, 0xd8, 0x23, 0xb5,
	0x25, 0xfa, 0xac, 0x9b, 0xfc, 0xf2, 0x3a, 0xdd, 0x5b, 0x4b, 0xd9, 0x38, 0x41, 0xda, 0x16, 0xe9,
	0xeb, 0xf5, 0x96, 0xe9, 0x92, 0xae, 0xac, 0xfe, 0x55, 0x3d, 0x1e, 0x93, 0x15, 0x1c, 0xfe, 0x63,
	0xc1, 0xdd, 0x76, 0xf0, 0x6e, 0x95, 0xf9, 0x17, 0x37, 0x33, 0x7f, 0x73, 0xeb, 0x74, 0xff, 0x1f,
	0x72, 0xff, 0xcd, 0x82, 0x1d, 0x95, 0x31, 0xae, 0xb2, 0x97, 0xcf, 0x5a, 0x5e, 0x2c, 0xf3, 0x2e,
	0x9f, 0x93, 0xbc, 0x1e, 0x2a, 0xe0, 0x8d, 0x64, 0x03, 0x4d, 0xe7, 0x05, 0x69, 0xfe, 0x69, 0xc1,
	0x58, 0xd1, 0x3c, 0x7d, 0x4e, 0x0b, 0xf1, 0xd2, 0x3c, 0x8f, 0x6e, 0x76, 0x67, 0x7f, 0xe3, 0x8a,
	0xe9, 0x76, 0xe6, 0xd6, 0xe4, 0x3f, 0xfe, 0xdd, 0xea, 0xc4, 0x95, 0xeb, 0x68, 0xa4, 0xc0, 0x93,
	0xc7, 0xc5, 0x75, 0xc1, 0x7e, 0x28, 0xfc, 0xd7, 0xd0, 0x8e, 0x51, 0x3f, 0x51, 0xb5, 0xf2, 0xad,
	0x06, 0x46, 0xc9, 0x39, 0xf3, 0x6d, 0xb4, 0x07, 0x3b, 0x1a, 0x9e, 0xd1, 0x42, 0xbe, 0xdb, 0x51,
	0xe2, 0x3b, 0x68, 0x0c, 0xa0, 0x45, 0xd3, 0x19, 0x95, 0xf8, 0x4e, 0x63, 0x72, 0x4e, 0xd4, 0xda,
	0x93, 0xa2, 0x5e, 0x13, 0x65, 0x9a, 0x24, 0x25, 0xe5, 0xdc, 0x77, 0x1b, 0x23, 0xe9, 0xa4, 0xd7,
	0xae, 0xdf, 0x47, 0xbb, 0x30, 0x5c, 0x46, 0x52, 0xaf, 0xbc, 0x3f, 0x78, 0xea, 0xea, 0x1f, 0xb0,
	0xa3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x84, 0xed, 0x06, 0xcc, 0x09, 0x00, 0x00,
}