// Code generated by protoc-gen-go.
// source: preview.proto
// DO NOT EDIT!

package dg_model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PreviewCmdType int32

const (
	PreviewCmdType_Cmd_Type_Unknown PreviewCmdType = 0
	PreviewCmdType_Cmd_Type_Start   PreviewCmdType = 1
	PreviewCmdType_Cmd_Type_Stop    PreviewCmdType = 2
)

var PreviewCmdType_name = map[int32]string{
	0: "Cmd_Type_Unknown",
	1: "Cmd_Type_Start",
	2: "Cmd_Type_Stop",
}
var PreviewCmdType_value = map[string]int32{
	"Cmd_Type_Unknown": 0,
	"Cmd_Type_Start":   1,
	"Cmd_Type_Stop":    2,
}

func (x PreviewCmdType) String() string {
	return proto.EnumName(PreviewCmdType_name, int32(x))
}
func (PreviewCmdType) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

type PreviewKeyType int32

const (
	PreviewKeyType_Key_Type_Unknown   PreviewKeyType = 0
	PreviewKeyType_Key_Type_Default   PreviewKeyType = 1
	PreviewKeyType_Key_Type_SensorId  PreviewKeyType = 2
	PreviewKeyType_Key_Type_SessionId PreviewKeyType = 3
)

var PreviewKeyType_name = map[int32]string{
	0: "Key_Type_Unknown",
	1: "Key_Type_Default",
	2: "Key_Type_SensorId",
	3: "Key_Type_SessionId",
}
var PreviewKeyType_value = map[string]int32{
	"Key_Type_Unknown":   0,
	"Key_Type_Default":   1,
	"Key_Type_SensorId":  2,
	"Key_Type_SessionId": 3,
}

func (x PreviewKeyType) String() string {
	return proto.EnumName(PreviewKeyType_name, int32(x))
}
func (PreviewKeyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

type PreviewDataType int32

const (
	PreviewDataType_Data_Type_Unknown   PreviewDataType = 0
	PreviewDataType_Data_Type_Original  PreviewDataType = 1
	PreviewDataType_Data_Type_Rendered  PreviewDataType = 2
	PreviewDataType_Data_Type_Structure PreviewDataType = 3
	PreviewDataType_Data_Type_Cutboard  PreviewDataType = 4
)

var PreviewDataType_name = map[int32]string{
	0: "Data_Type_Unknown",
	1: "Data_Type_Original",
	2: "Data_Type_Rendered",
	3: "Data_Type_Structure",
	4: "Data_Type_Cutboard",
}
var PreviewDataType_value = map[string]int32{
	"Data_Type_Unknown":   0,
	"Data_Type_Original":  1,
	"Data_Type_Rendered":  2,
	"Data_Type_Structure": 3,
	"Data_Type_Cutboard":  4,
}

func (x PreviewDataType) String() string {
	return proto.EnumName(PreviewDataType_name, int32(x))
}
func (PreviewDataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

type NanomsgCmd struct {
	CmdType     PreviewCmdType `protobuf:"varint,1,opt,name=CmdType,json=cmdType,enum=dg.model.PreviewCmdType" json:"CmdType"`
	KeyType     PreviewKeyType `protobuf:"varint,2,opt,name=KeyType,json=keyType,enum=dg.model.PreviewKeyType" json:"KeyType"`
	Key         string         `protobuf:"bytes,3,opt,name=Key,json=key" json:"Key"`
	PubEnabled  bool           `protobuf:"varint,4,opt,name=PubEnabled,json=pubEnabled" json:"PubEnabled"`
	PubAddress  string         `protobuf:"bytes,5,opt,name=PubAddress,json=pubAddress" json:"PubAddress"`
	PubInterval int32          `protobuf:"varint,6,opt,name=PubInterval,json=pubInterval" json:"PubInterval"`
	DeadTime    int64          `protobuf:"varint,7,opt,name=DeadTime,json=deadTime" json:"DeadTime"`
}

func (m *NanomsgCmd) Reset()                    { *m = NanomsgCmd{} }
func (m *NanomsgCmd) String() string            { return proto.CompactTextString(m) }
func (*NanomsgCmd) ProtoMessage()               {}
func (*NanomsgCmd) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

type NanomsgData struct {
	Key     string `protobuf:"bytes,1,opt,name=Key,json=key" json:"Key"`
	BinData []byte `protobuf:"bytes,2,opt,name=BinData,json=binData,proto3" json:"BinData"`
}

func (m *NanomsgData) Reset()                    { *m = NanomsgData{} }
func (m *NanomsgData) String() string            { return proto.CompactTextString(m) }
func (*NanomsgData) ProtoMessage()               {}
func (*NanomsgData) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

type PreviewCmd struct {
	DataType    PreviewDataType `protobuf:"varint,1,opt,name=DataType,json=dataType,enum=dg.model.PreviewDataType" json:"DataType"`
	RepoId      int32           `protobuf:"varint,2,opt,name=RepoId,json=repoId" json:"RepoId"`
	SensorId    int32           `protobuf:"varint,3,opt,name=SensorId,json=sensorId" json:"SensorId"`
	DeadTime    int64           `protobuf:"varint,4,opt,name=DeadTime,json=deadTime" json:"DeadTime"`
	KeyType     PreviewKeyType  `protobuf:"varint,8,opt,name=KeyType,json=keyType,enum=dg.model.PreviewKeyType" json:"KeyType"`
	Key         string          `protobuf:"bytes,9,opt,name=Key,json=key" json:"Key"`
	PubAddress  string          `protobuf:"bytes,10,opt,name=PubAddress,json=pubAddress" json:"PubAddress"`
	PubInterval int32           `protobuf:"varint,11,opt,name=PubInterval,json=pubInterval" json:"PubInterval"`
}

func (m *PreviewCmd) Reset()                    { *m = PreviewCmd{} }
func (m *PreviewCmd) String() string            { return proto.CompactTextString(m) }
func (*PreviewCmd) ProtoMessage()               {}
func (*PreviewCmd) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

// deprecated
type PreviewData struct {
	Id           int64  `protobuf:"varint,1,opt,name=Id,json=id" json:"Id"`
	Timestamp    int64  `protobuf:"varint,2,opt,name=Timestamp,json=timestamp" json:"Timestamp"`
	RepoId       int32  `protobuf:"varint,3,opt,name=RepoId,json=repoId" json:"RepoId"`
	SensorId     int32  `protobuf:"varint,4,opt,name=SensorId,json=sensorId" json:"SensorId"`
	PlateText    string `protobuf:"bytes,5,opt,name=PlateText,json=plateText" json:"PlateText"`
	PlateType    string `protobuf:"bytes,6,opt,name=PlateType,json=plateType" json:"PlateType"`
	PlateColor   string `protobuf:"bytes,7,opt,name=PlateColor,json=plateColor" json:"PlateColor"`
	VehicleType  string `protobuf:"bytes,8,opt,name=VehicleType,json=vehicleType" json:"VehicleType"`
	VehicleColor string `protobuf:"bytes,9,opt,name=VehicleColor,json=vehicleColor" json:"VehicleColor"`
	VehicleImage string `protobuf:"bytes,10,opt,name=VehicleImage,json=vehicleImage" json:"VehicleImage"`
	Brand        string `protobuf:"bytes,11,opt,name=Brand,json=brand" json:"Brand"`
	SubBrand     string `protobuf:"bytes,12,opt,name=SubBrand,json=subBrand" json:"SubBrand"`
	ModelYear    string `protobuf:"bytes,13,opt,name=ModelYear,json=modelYear" json:"ModelYear"`
}

func (m *PreviewData) Reset()                    { *m = PreviewData{} }
func (m *PreviewData) String() string            { return proto.CompactTextString(m) }
func (*PreviewData) ProtoMessage()               {}
func (*PreviewData) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

type PreviewImage struct {
	SensorId string `protobuf:"bytes,1,opt,name=SensorId,json=sensorId" json:"SensorId"`
	BinData  string `protobuf:"bytes,2,opt,name=BinData,json=binData" json:"BinData"`
}

func (m *PreviewImage) Reset()                    { *m = PreviewImage{} }
func (m *PreviewImage) String() string            { return proto.CompactTextString(m) }
func (*PreviewImage) ProtoMessage()               {}
func (*PreviewImage) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{4} }

func init() {
	proto.RegisterType((*NanomsgCmd)(nil), "dg.model.NanomsgCmd")
	proto.RegisterType((*NanomsgData)(nil), "dg.model.NanomsgData")
	proto.RegisterType((*PreviewCmd)(nil), "dg.model.PreviewCmd")
	proto.RegisterType((*PreviewData)(nil), "dg.model.PreviewData")
	proto.RegisterType((*PreviewImage)(nil), "dg.model.PreviewImage")
	proto.RegisterEnum("dg.model.PreviewCmdType", PreviewCmdType_name, PreviewCmdType_value)
	proto.RegisterEnum("dg.model.PreviewKeyType", PreviewKeyType_name, PreviewKeyType_value)
	proto.RegisterEnum("dg.model.PreviewDataType", PreviewDataType_name, PreviewDataType_value)
}

func init() { proto.RegisterFile("preview.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 635 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6f, 0x13, 0x3f,
	0x10, 0xfd, 0x6d, 0x36, 0x1f, 0xbb, 0x93, 0x34, 0xbf, 0xad, 0xf9, 0x5a, 0x10, 0x42, 0x55, 0x4e,
	0xa8, 0x87, 0x1c, 0x8a, 0x38, 0x70, 0xa4, 0x0d, 0x87, 0x0a, 0x15, 0x2a, 0xb7, 0x20, 0x71, 0xaa,
	0xbc, 0xb5, 0x09, 0xab, 0xec, 0xda, 0x2b, 0xaf, 0x93, 0xd2, 0x3b, 0x1c, 0xf8, 0x1b, 0xf8, 0x67,
	0xb1, 0xbd, 0xde, 0x8f, 0xa4, 0xaa, 0x10, 0xa7, 0xcc, 0xbc, 0x37, 0x1e, 0xcf, 0x7b, 0x9e, 0x2c,
	0xec, 0x15, 0x92, 0x6d, 0x52, 0x76, 0x33, 0x2f, 0xa4, 0x50, 0x02, 0x05, 0x74, 0x39, 0xcf, 0x05,
	0x65, 0xd9, 0xec, 0x47, 0x0f, 0xe0, 0x03, 0xe1, 0x22, 0x2f, 0x97, 0x27, 0x39, 0x45, 0x47, 0x30,
	0xd2, 0x3f, 0x97, 0xb7, 0x05, 0x8b, 0xbd, 0x03, 0xef, 0xe5, 0xf4, 0x28, 0x9e, 0xd7, 0xa5, 0xf3,
	0xf3, 0xaa, 0x85, 0xe3, 0xf1, 0xe8, 0xba, 0x0a, 0xcc, 0x99, 0xf7, 0xec, 0xd6, 0x9e, 0xe9, 0xdd,
	0x73, 0xc6, 0xf1, 0x78, 0xb4, 0xaa, 0x02, 0x14, 0x81, 0xaf, 0xb1, 0xd8, 0xd7, 0xf5, 0x21, 0xf6,
	0x35, 0x8a, 0x5e, 0x00, 0x9c, 0xaf, 0x93, 0x77, 0x9c, 0x24, 0x19, 0xa3, 0x71, 0x5f, 0x13, 0x01,
	0x86, 0xa2, 0x41, 0x1c, 0xff, 0x96, 0x52, 0xc9, 0xca, 0x32, 0x1e, 0xd8, 0x83, 0x86, 0x77, 0x08,
	0x3a, 0x80, 0xb1, 0xe6, 0x4f, 0xb9, 0x62, 0x72, 0x43, 0xb2, 0x78, 0xa8, 0x0b, 0x06, 0x78, 0x5c,
	0xb4, 0x10, 0x7a, 0x06, 0xc1, 0x82, 0x11, 0x7a, 0x99, 0xe6, 0x2c, 0x1e, 0x69, 0xda, 0xc7, 0x01,
	0x75, 0xf9, 0xec, 0x0d, 0x8c, 0x9d, 0x0b, 0x0b, 0xa2, 0x48, 0x3d, 0x9e, 0xd7, 0x8e, 0x17, 0xc3,
	0xe8, 0x38, 0xe5, 0x86, 0xb4, 0x22, 0x27, 0x78, 0x94, 0x54, 0xe9, 0xec, 0xb7, 0x76, 0xb0, 0xb5,
	0x06, 0xbd, 0xd6, 0xb7, 0x68, 0xb8, 0x63, 0xe1, 0xd3, 0x3b, 0x76, 0xd4, 0x05, 0x7a, 0x00, 0x17,
	0xa1, 0xc7, 0x30, 0xc4, 0xac, 0x10, 0xa7, 0xd4, 0xb6, 0x1f, 0xe0, 0xa1, 0xb4, 0x99, 0x19, 0xfa,
	0x82, 0xf1, 0x52, 0x48, 0xcd, 0xf8, 0x96, 0x09, 0x4a, 0x97, 0x6f, 0x09, 0xea, 0x6f, 0x0b, 0xea,
	0x3e, 0x4a, 0xf0, 0x8f, 0x8f, 0x12, 0xee, 0x3e, 0x4a, 0x6d, 0x3a, 0xfc, 0xcd, 0xf4, 0xf1, 0x1d,
	0xd3, 0x67, 0x3f, 0x7d, 0x5d, 0xd2, 0xaa, 0x46, 0x53, 0xe8, 0x69, 0x25, 0x9e, 0x9d, 0xb6, 0x97,
	0x52, 0xf4, 0x1c, 0x42, 0x33, 0x6f, 0xa9, 0x48, 0x5e, 0x58, 0xe9, 0x3e, 0x0e, 0x55, 0x0d, 0x74,
	0x5c, 0xf1, 0xef, 0x75, 0xa5, 0xbf, 0xe3, 0x8a, 0xee, 0x78, 0x9e, 0x11, 0xc5, 0x2e, 0xd9, 0x77,
	0xe5, 0xf6, 0x24, 0x2c, 0x6a, 0xa0, 0x65, 0x8d, 0x33, 0xc3, 0x2e, 0x6b, 0x1c, 0x30, 0x7a, 0x4d,
	0x72, 0x22, 0x32, 0x21, 0xed, 0x92, 0x18, 0xbd, 0x0d, 0x62, 0xf4, 0x7e, 0x66, 0xdf, 0xd2, 0xeb,
	0x8c, 0x35, 0xce, 0x86, 0x78, 0xbc, 0x69, 0x21, 0x34, 0x83, 0x89, 0xab, 0xa8, 0x7a, 0x54, 0x66,
	0x4e, 0x36, 0x1d, 0xac, 0x53, 0x73, 0x9a, 0x93, 0x25, 0x73, 0xbe, 0xd6, 0x35, 0x16, 0x43, 0x0f,
	0x61, 0x70, 0x2c, 0x09, 0xa7, 0xd6, 0xd3, 0x10, 0x0f, 0x12, 0x93, 0x58, 0xdd, 0xeb, 0xa4, 0x22,
	0x26, 0x96, 0x08, 0x4a, 0x97, 0x1b, 0x65, 0x67, 0xe6, 0x79, 0xbf, 0x30, 0x22, 0xe3, 0xbd, 0x4a,
	0x59, 0x5e, 0x03, 0xb3, 0x05, 0x4c, 0xdc, 0x33, 0x54, 0xfd, 0xbb, 0x0e, 0x7a, 0xae, 0x53, 0xed,
	0xe0, 0xce, 0xae, 0x87, 0xcd, 0xae, 0x1f, 0x9e, 0xc1, 0x74, 0xfb, 0x2b, 0xa0, 0xe7, 0x8c, 0x74,
	0x78, 0x65, 0xe2, 0xab, 0x4f, 0x7c, 0xc5, 0xc5, 0x0d, 0x8f, 0xfe, 0x43, 0x08, 0xa6, 0x0d, 0x7a,
	0xa1, 0x88, 0x54, 0x91, 0x87, 0xf6, 0x61, 0xaf, 0x83, 0x89, 0x22, 0xea, 0x1d, 0xae, 0x9a, 0x76,
	0x6e, 0x17, 0x4d, 0x3b, 0x1d, 0xee, 0xb6, 0xeb, 0xa2, 0x0b, 0xf6, 0x95, 0xac, 0x33, 0xd3, 0xf0,
	0x11, 0xec, 0x37, 0x68, 0xad, 0x25, 0xea, 0xe9, 0x9d, 0x41, 0x1d, 0xb8, 0x2c, 0x53, 0xc1, 0x35,
	0xee, 0x1f, 0xfe, 0xf2, 0xe0, 0xff, 0x9d, 0xff, 0x9f, 0x69, 0x61, 0xe2, 0xdd, 0xfb, 0x74, 0x8b,
	0x16, 0xfe, 0x28, 0xd3, 0x65, 0xca, 0x49, 0xa6, 0x6f, 0xdc, 0xc2, 0x31, 0xe3, 0x94, 0x49, 0x66,
	0xae, 0x7c, 0x02, 0x0f, 0x5a, 0xfc, 0x42, 0xc9, 0xf5, 0xb5, 0x5a, 0x4b, 0x16, 0xf9, 0xdb, 0x07,
	0x4e, 0xd6, 0x2a, 0x11, 0x44, 0xd2, 0xa8, 0x9f, 0x0c, 0xed, 0x67, 0xf8, 0xd5, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x5f, 0x68, 0xc4, 0xa9, 0x97, 0x05, 0x00, 0x00,
}
