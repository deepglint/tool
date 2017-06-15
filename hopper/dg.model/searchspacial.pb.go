// Code generated by protoc-gen-go.
// source: searchspacial.proto
// DO NOT EDIT!

package dg_model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TaskType int32

const (
	TaskType_Task_Type_Unknown TaskType = 0
	TaskType_Task_Type_Image   TaskType = 1
	TaskType_Task_Type_Video   TaskType = 2
	TaskType_Task_Type_Local   TaskType = 3
)

var TaskType_name = map[int32]string{
	0: "Task_Type_Unknown",
	1: "Task_Type_Image",
	2: "Task_Type_Video",
	3: "Task_Type_Local",
}
var TaskType_value = map[string]int32{
	"Task_Type_Unknown": 0,
	"Task_Type_Image":   1,
	"Task_Type_Video":   2,
	"Task_Type_Local":   3,
}

func (x TaskType) String() string {
	return proto.EnumName(TaskType_name, int32(x))
}
func (TaskType) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

type RepoInfo struct {
	RepoId       int32      `protobuf:"varint,1,opt,name=RepoId,json=repoId" json:"RepoId"`
	Name         string     `protobuf:"bytes,2,opt,name=Name,json=name" json:"Name"`
	Status       TaskStatus `protobuf:"varint,3,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
	IsBound      bool       `protobuf:"varint,4,opt,name=IsBound,json=isBound" json:"IsBound"`
	UniqueRepoId string     `protobuf:"bytes,5,opt,name=UniqueRepoId,json=uniqueRepoId" json:"UniqueRepoId"`
	UserId       int32      `protobuf:"varint,16,opt,name=UserId,json=userId" json:"UserId"`
}

func (m *RepoInfo) Reset()                    { *m = RepoInfo{} }
func (m *RepoInfo) String() string            { return proto.CompactTextString(m) }
func (*RepoInfo) ProtoMessage()               {}
func (*RepoInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

type SensorInfo struct {
	Id             int32      `protobuf:"varint,1,opt,name=Id,json=id" json:"Id"`
	RepoId         int32      `protobuf:"varint,4,opt,name=RepoId,json=repoId" json:"RepoId"`
	SensorId       int32      `protobuf:"varint,5,opt,name=SensorId,json=sensorId" json:"SensorId"`
	Name           string     `protobuf:"bytes,6,opt,name=Name,json=name" json:"Name"`
	Longitude      float32    `protobuf:"fixed32,7,opt,name=Longitude,json=longitude" json:"Longitude"`
	Latitude       float32    `protobuf:"fixed32,8,opt,name=Latitude,json=latitude" json:"Latitude"`
	MinTimestamp   int64      `protobuf:"varint,9,opt,name=MinTimestamp,json=minTimestamp" json:"MinTimestamp"`
	MaxTimestamp   int64      `protobuf:"varint,10,opt,name=MaxTimestamp,json=maxTimestamp" json:"MaxTimestamp"`
	Status         TaskStatus `protobuf:"varint,11,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
	IsBound        bool       `protobuf:"varint,12,opt,name=IsBound,json=isBound" json:"IsBound"`
	UniqueRepoId   string     `protobuf:"bytes,13,opt,name=UniqueRepoId,json=uniqueRepoId" json:"UniqueRepoId"`
	UniqueSensorId string     `protobuf:"bytes,14,opt,name=UniqueSensorId,json=uniqueSensorId" json:"UniqueSensorId"`
	PicCount       int32      `protobuf:"varint,16,opt,name=PicCount,json=picCount" json:"PicCount"`
}

func (m *SensorInfo) Reset()                    { *m = SensorInfo{} }
func (m *SensorInfo) String() string            { return proto.CompactTextString(m) }
func (*SensorInfo) ProtoMessage()               {}
func (*SensorInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

type DayStatistics struct {
	Day      int64 `protobuf:"varint,1,opt,name=Day,json=day" json:"Day"`
	PicCount int32 `protobuf:"varint,2,opt,name=PicCount,json=picCount" json:"PicCount"`
	RetCount int32 `protobuf:"varint,3,opt,name=RetCount,json=retCount" json:"RetCount"`
	PeoCount int32 `protobuf:"varint,4,opt,name=PeoCount,json=peoCount" json:"PeoCount"`
	BicCount int32 `protobuf:"varint,5,opt,name=BicCount,json=bicCount" json:"BicCount"`
	TriCount int32 `protobuf:"varint,6,opt,name=TriCount,json=triCount" json:"TriCount"`
}

func (m *DayStatistics) Reset()                    { *m = DayStatistics{} }
func (m *DayStatistics) String() string            { return proto.CompactTextString(m) }
func (*DayStatistics) ProtoMessage()               {}
func (*DayStatistics) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

type DayStatisticsResult struct {
	Statistics []*DayStatistics `protobuf:"bytes,1,rep,name=Statistics,json=statistics" json:"Statistics"`
}

func (m *DayStatisticsResult) Reset()                    { *m = DayStatisticsResult{} }
func (m *DayStatisticsResult) String() string            { return proto.CompactTextString(m) }
func (*DayStatisticsResult) ProtoMessage()               {}
func (*DayStatisticsResult) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{3} }

func (m *DayStatisticsResult) GetStatistics() []*DayStatistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func init() {
	proto.RegisterType((*RepoInfo)(nil), "dg.model.RepoInfo")
	proto.RegisterType((*SensorInfo)(nil), "dg.model.SensorInfo")
	proto.RegisterType((*DayStatistics)(nil), "dg.model.DayStatistics")
	proto.RegisterType((*DayStatisticsResult)(nil), "dg.model.DayStatisticsResult")
	proto.RegisterEnum("dg.model.TaskType", TaskType_name, TaskType_value)
}

func init() { proto.RegisterFile("searchspacial.proto", fileDescriptor16) }

var fileDescriptor16 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0x69, 0x92, 0x49, 0xd3, 0x33, 0x9d, 0x10, 0x5c, 0x2e, 0xd1, 0x88, 0x05, 0xca, 0x02,
	0x8d, 0x10, 0xea, 0x62, 0x58, 0xb0, 0x1f, 0xd8, 0x54, 0x2a, 0x23, 0xe4, 0x69, 0xd9, 0x0e, 0x9e,
	0xc4, 0x14, 0x6b, 0x12, 0x3b, 0xc4, 0x8e, 0xa0, 0x2f, 0xc2, 0x5b, 0xf0, 0x02, 0x3c, 0x1d, 0x8e,
	0xed, 0xa6, 0x89, 0x84, 0x90, 0x58, 0xd5, 0xff, 0xff, 0x9d, 0xe3, 0x73, 0x71, 0x03, 0x0b, 0x49,
	0x49, 0x93, 0x7f, 0x95, 0x35, 0xc9, 0x19, 0x29, 0x97, 0x75, 0x23, 0x94, 0x40, 0x51, 0xb1, 0x5b,
	0x56, 0xa2, 0xa0, 0xe5, 0x39, 0xb2, 0x38, 0x17, 0x55, 0x25, 0xb8, 0xa5, 0xd9, 0xef, 0x09, 0x44,
	0x98, 0xd6, 0x62, 0xc5, 0xbf, 0x08, 0xf4, 0x14, 0x42, 0x73, 0x2e, 0xd2, 0xc9, 0x8b, 0xc9, 0xc5,
	0x09, 0x0e, 0x1b, 0xa3, 0x10, 0x82, 0xe0, 0x9a, 0x54, 0x34, 0xf5, 0xb4, 0x3b, 0xc3, 0x01, 0xd7,
	0x67, 0xf4, 0x1a, 0xc2, 0x1b, 0x45, 0x54, 0x2b, 0x53, 0x5f, 0xbb, 0xf1, 0xe5, 0xe3, 0xe5, 0xa1,
	0xce, 0x72, 0x43, 0xe4, 0xbd, 0x65, 0x38, 0x94, 0xe6, 0x17, 0xa5, 0x30, 0x5d, 0xc9, 0x2b, 0xd1,
	0xf2, 0x22, 0x0d, 0x74, 0x78, 0x84, 0xa7, 0xcc, 0x4a, 0x94, 0xc1, 0x7c, 0xcb, 0xd9, 0xb7, 0x96,
	0xba, 0xca, 0x27, 0xa6, 0xc6, 0xbc, 0x1d, 0x78, 0x5d, 0x5f, 0x5b, 0x49, 0x1b, 0x4d, 0x13, 0xdb,
	0x57, 0x6b, 0x54, 0xf6, 0xd3, 0x07, 0xb8, 0xa1, 0x5c, 0x8a, 0xc6, 0xb4, 0x1f, 0x83, 0xd7, 0xb7,
	0xee, 0xb1, 0x62, 0x30, 0x4e, 0x30, 0x1a, 0xe7, 0x1c, 0x22, 0x97, 0x65, 0xcb, 0x9d, 0xe0, 0x48,
	0x3a, 0xdd, 0x8f, 0x1a, 0x0e, 0x46, 0x7d, 0x0e, 0xb3, 0xb5, 0xe0, 0x3b, 0xa6, 0xda, 0x82, 0xa6,
	0x53, 0x0d, 0x3c, 0x3c, 0x2b, 0x0f, 0x46, 0x77, 0xdb, 0x9a, 0x28, 0x0b, 0x23, 0x03, 0xa3, 0xd2,
	0xe9, 0x6e, 0xb8, 0x0f, 0x8c, 0x6f, 0x58, 0x45, 0xf5, 0x1e, 0xaa, 0x3a, 0x9d, 0x69, 0xee, 0xe3,
	0x79, 0x35, 0xf0, 0x4c, 0x0c, 0xf9, 0x71, 0x8c, 0x01, 0x17, 0x33, 0xf0, 0x06, 0xcb, 0x3e, 0xfd,
	0xbf, 0x65, 0xcf, 0xff, 0xbd, 0xec, 0xb3, 0xbf, 0x2c, 0xfb, 0x25, 0xc4, 0x36, 0xa6, 0xdf, 0x51,
	0x6c, 0xa2, 0xe2, 0x76, 0xe4, 0x76, 0x73, 0x7f, 0x64, 0xf9, 0x3b, 0x7d, 0xaf, 0x72, 0xcf, 0x12,
	0xd5, 0x4e, 0x67, 0xbf, 0x26, 0x70, 0xf6, 0x9e, 0xec, 0xbb, 0xbe, 0x98, 0x54, 0x2c, 0x97, 0x28,
	0x01, 0x5f, 0x1b, 0xe6, 0x71, 0x7c, 0xec, 0x17, 0x64, 0x3f, 0xca, 0xf7, 0xc6, 0xf9, 0x1d, 0xc3,
	0x54, 0x59, 0xe6, 0x5b, 0xd6, 0x38, 0x6d, 0xf2, 0xa8, 0xb0, 0x2c, 0x70, 0x79, 0x4e, 0x77, 0xec,
	0xea, 0x70, 0xa7, 0x7b, 0xd9, 0xbb, 0xc1, 0x9d, 0x9b, 0x86, 0x59, 0x16, 0x5a, 0xa6, 0x9c, 0xce,
	0xae, 0x61, 0x31, 0x6a, 0x17, 0x53, 0xd9, 0x96, 0x0a, 0xbd, 0xd5, 0x7f, 0xaf, 0xde, 0xd3, 0xbd,
	0xfb, 0x17, 0xa7, 0x97, 0xcf, 0x8e, 0xab, 0x1f, 0xa7, 0x80, 0xec, 0xcf, 0xaf, 0x3e, 0xeb, 0x5a,
	0xfa, 0x5d, 0x36, 0xfb, 0x9a, 0xa2, 0x27, 0xf0, 0xa8, 0x3b, 0xdf, 0x76, 0xe2, 0x76, 0xcb, 0xef,
	0xb9, 0xf8, 0xce, 0x93, 0x07, 0x68, 0x01, 0x0f, 0x8f, 0xf6, 0xaa, 0x22, 0x3b, 0x9a, 0x4c, 0xc6,
	0xe6, 0x27, 0x56, 0x50, 0x91, 0x78, 0x63, 0x73, 0x2d, 0x72, 0x52, 0x26, 0xfe, 0x5d, 0x68, 0x3e,
	0xdf, 0x37, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x72, 0x3d, 0x39, 0xf3, 0x03, 0x00, 0x00,
}