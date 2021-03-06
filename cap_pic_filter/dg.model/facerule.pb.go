// Code generated by protoc-gen-go.
// source: facerule.proto
// DO NOT EDIT!

package dg_model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FaceRule struct {
	RuleId    string          `protobuf:"bytes,1,opt,name=RuleId,json=ruleId" json:"RuleId"`
	Timestamp int64           `protobuf:"varint,2,opt,name=Timestamp,json=timestamp" json:"Timestamp"`
	Sensors   []*SensorInRule `protobuf:"bytes,3,rep,name=Sensors,json=sensors" json:"Sensors"`
	FaceRepos []*RepoInRule   `protobuf:"bytes,4,rep,name=FaceRepos,json=faceRepos" json:"FaceRepos"`
	Times     *TimeInRule     `protobuf:"bytes,5,opt,name=Times,json=times" json:"Times"`
	Comment   string          `protobuf:"bytes,6,opt,name=Comment,json=comment" json:"Comment"`
	RepoCount int32           `protobuf:"varint,7,opt,name=RepoCount,json=repoCount" json:"RepoCount"`
	Status    TaskStatus      `protobuf:"varint,99,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
}

func (m *FaceRule) Reset()                    { *m = FaceRule{} }
func (m *FaceRule) String() string            { return proto.CompactTextString(m) }
func (*FaceRule) ProtoMessage()               {}
func (*FaceRule) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *FaceRule) GetSensors() []*SensorInRule {
	if m != nil {
		return m.Sensors
	}
	return nil
}

func (m *FaceRule) GetFaceRepos() []*RepoInRule {
	if m != nil {
		return m.FaceRepos
	}
	return nil
}

func (m *FaceRule) GetTimes() *TimeInRule {
	if m != nil {
		return m.Times
	}
	return nil
}

type Point struct {
	X int32 `protobuf:"varint,1,opt,name=X,json=x" json:"X"`
	Y int32 `protobuf:"varint,2,opt,name=Y,json=y" json:"Y"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

type Roi struct {
	Points []*Point `protobuf:"bytes,1,rep,name=Points,json=points" json:"Points"`
}

func (m *Roi) Reset()                    { *m = Roi{} }
func (m *Roi) String() string            { return proto.CompactTextString(m) }
func (*Roi) ProtoMessage()               {}
func (*Roi) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *Roi) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

type SensorInRule struct {
	SensorId     string         `protobuf:"bytes,1,opt,name=SensorId,json=sensorId" json:"SensorId"`
	SensorName   string         `protobuf:"bytes,2,opt,name=SensorName,json=sensorName" json:"SensorName"`
	Rois         []*Roi         `protobuf:"bytes,3,rep,name=Rois,json=rois" json:"Rois"`
	Switcher     SwitcherStatus `protobuf:"varint,97,opt,name=Switcher,json=switcher,enum=dg.model.SwitcherStatus" json:"Switcher"`
	Status       TaskStatus     `protobuf:"varint,98,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
	RuleSensorId string         `protobuf:"bytes,99,opt,name=RuleSensorId,json=ruleSensorId" json:"RuleSensorId"`
}

func (m *SensorInRule) Reset()                    { *m = SensorInRule{} }
func (m *SensorInRule) String() string            { return proto.CompactTextString(m) }
func (*SensorInRule) ProtoMessage()               {}
func (*SensorInRule) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *SensorInRule) GetRois() []*Roi {
	if m != nil {
		return m.Rois
	}
	return nil
}

type RepoInRule struct {
	RepoId     string     `protobuf:"bytes,1,opt,name=RepoId,json=repoId" json:"RepoId"`
	RepoName   string     `protobuf:"bytes,2,opt,name=RepoName,json=repoName" json:"RepoName"`
	Confidence float32    `protobuf:"fixed32,3,opt,name=Confidence,json=confidence" json:"Confidence"`
	CountLimit int32      `protobuf:"varint,4,opt,name=CountLimit,json=countLimit" json:"CountLimit"`
	Status     TaskStatus `protobuf:"varint,98,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
	RuleRepoId string     `protobuf:"bytes,99,opt,name=RuleRepoId,json=ruleRepoId" json:"RuleRepoId"`
}

func (m *RepoInRule) Reset()                    { *m = RepoInRule{} }
func (m *RepoInRule) String() string            { return proto.CompactTextString(m) }
func (*RepoInRule) ProtoMessage()               {}
func (*RepoInRule) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

type TimeInRule struct {
	StartDate int64 `protobuf:"varint,1,opt,name=StartDate,json=startDate" json:"StartDate"`
	EndDate   int64 `protobuf:"varint,2,opt,name=EndDate,json=endDate" json:"EndDate"`
	StartTime int64 `protobuf:"varint,3,opt,name=StartTime,json=startTime" json:"StartTime"`
	EndTime   int64 `protobuf:"varint,4,opt,name=EndTime,json=endTime" json:"EndTime"`
	IsLong    bool  `protobuf:"varint,5,opt,name=IsLong,json=isLong" json:"IsLong"`
}

func (m *TimeInRule) Reset()                    { *m = TimeInRule{} }
func (m *TimeInRule) String() string            { return proto.CompactTextString(m) }
func (*TimeInRule) ProtoMessage()               {}
func (*TimeInRule) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

type FaceRuleSummary struct {
	Comment   string        `protobuf:"bytes,1,opt,name=Comment,json=comment" json:"Comment"`
	Timestamp int64         `protobuf:"varint,2,opt,name=Timestamp,json=timestamp" json:"Timestamp"`
	RuleId    string        `protobuf:"bytes,3,opt,name=RuleId,json=ruleId" json:"RuleId"`
	Sensor    *SensorInRule `protobuf:"bytes,4,opt,name=Sensor,json=sensor" json:"Sensor"`
	Repos     []*RepoInRule `protobuf:"bytes,5,rep,name=Repos,json=repos" json:"Repos"`
	Times     *TimeInRule   `protobuf:"bytes,6,opt,name=Times,json=times" json:"Times"`
	Status    TaskStatus    `protobuf:"varint,7,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
}

func (m *FaceRuleSummary) Reset()                    { *m = FaceRuleSummary{} }
func (m *FaceRuleSummary) String() string            { return proto.CompactTextString(m) }
func (*FaceRuleSummary) ProtoMessage()               {}
func (*FaceRuleSummary) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

func (m *FaceRuleSummary) GetSensor() *SensorInRule {
	if m != nil {
		return m.Sensor
	}
	return nil
}

func (m *FaceRuleSummary) GetRepos() []*RepoInRule {
	if m != nil {
		return m.Repos
	}
	return nil
}

func (m *FaceRuleSummary) GetTimes() *TimeInRule {
	if m != nil {
		return m.Times
	}
	return nil
}

type FaceRuleResult struct {
	AllSize      int32       `protobuf:"varint,1,opt,name=AllSize,json=allSize" json:"AllSize"`
	ReturnedSize int32       `protobuf:"varint,2,opt,name=ReturnedSize,json=returnedSize" json:"ReturnedSize"`
	FaceRules    []*FaceRule `protobuf:"bytes,3,rep,name=FaceRules,json=faceRules" json:"FaceRules"`
}

func (m *FaceRuleResult) Reset()                    { *m = FaceRuleResult{} }
func (m *FaceRuleResult) String() string            { return proto.CompactTextString(m) }
func (*FaceRuleResult) ProtoMessage()               {}
func (*FaceRuleResult) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{7} }

func (m *FaceRuleResult) GetFaceRules() []*FaceRule {
	if m != nil {
		return m.FaceRules
	}
	return nil
}

type FaceEvent struct {
}

func (m *FaceEvent) Reset()                    { *m = FaceEvent{} }
func (m *FaceEvent) String() string            { return proto.CompactTextString(m) }
func (*FaceEvent) ProtoMessage()               {}
func (*FaceEvent) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{8} }

func init() {
	proto.RegisterType((*FaceRule)(nil), "dg.model.FaceRule")
	proto.RegisterType((*Point)(nil), "dg.model.Point")
	proto.RegisterType((*Roi)(nil), "dg.model.Roi")
	proto.RegisterType((*SensorInRule)(nil), "dg.model.SensorInRule")
	proto.RegisterType((*RepoInRule)(nil), "dg.model.RepoInRule")
	proto.RegisterType((*TimeInRule)(nil), "dg.model.TimeInRule")
	proto.RegisterType((*FaceRuleSummary)(nil), "dg.model.FaceRuleSummary")
	proto.RegisterType((*FaceRuleResult)(nil), "dg.model.FaceRuleResult")
	proto.RegisterType((*FaceEvent)(nil), "dg.model.FaceEvent")
}

func init() { proto.RegisterFile("facerule.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0x14, 0x31,
	0x14, 0x66, 0x76, 0x77, 0xfe, 0x4e, 0xd7, 0x16, 0x82, 0x94, 0xa1, 0x88, 0xd4, 0xf1, 0xc2, 0x22,
	0xb2, 0x48, 0xf5, 0x05, 0xa4, 0x56, 0x28, 0x14, 0x91, 0xac, 0x17, 0x7a, 0x39, 0x9d, 0xcd, 0xb6,
	0x83, 0x3b, 0x93, 0x25, 0xc9, 0xaa, 0xf5, 0xca, 0x87, 0x10, 0x7c, 0x15, 0x9f, 0xc2, 0xb7, 0xf1,
	0xde, 0x93, 0x93, 0xc9, 0xcc, 0x2c, 0x62, 0xab, 0x57, 0xe1, 0x7c, 0x5f, 0xce, 0x49, 0xbe, 0xf3,
	0x07, 0xbb, 0xcb, 0xa2, 0x14, 0x6a, 0xb3, 0x12, 0xb3, 0xb5, 0x92, 0x46, 0xb2, 0x64, 0x71, 0x39,
	0xab, 0xe5, 0x42, 0xac, 0x0e, 0x98, 0x16, 0x85, 0x2a, 0xaf, 0x4a, 0x59, 0xd7, 0xb2, 0x71, 0x6c,
	0xfe, 0x63, 0x04, 0xc9, 0x2b, 0x74, 0xe0, 0xe8, 0xc0, 0xf6, 0x21, 0xb2, 0xe7, 0xd9, 0x22, 0x0b,
	0x0e, 0x83, 0xa3, 0x94, 0x47, 0x8a, 0x2c, 0x76, 0x0f, 0xd2, 0xb7, 0x55, 0x2d, 0xb4, 0x29, 0xea,
	0x75, 0x36, 0x42, 0x6a, 0xcc, 0x53, 0xe3, 0x01, 0xf6, 0x14, 0xe2, 0xb9, 0x68, 0xb4, 0x54, 0x3a,
	0x1b, 0x1f, 0x8e, 0x8f, 0x76, 0x8e, 0xf7, 0x67, 0xfe, 0xc9, 0x99, 0x23, 0xce, 0x1a, 0x1b, 0x96,
	0xc7, 0xda, 0x5d, 0x63, 0xc7, 0x90, 0xd2, 0x9b, 0x62, 0x2d, 0x75, 0x36, 0x21, 0x9f, 0xbb, 0xbd,
	0x8f, 0x85, 0x5b, 0x8f, 0x74, 0xe9, 0xaf, 0xb1, 0xc7, 0x10, 0xd2, 0x1f, 0xb2, 0x10, 0xdf, 0xdf,
	0xba, 0x6f, 0xe1, 0xf6, 0x7e, 0x48, 0xbf, 0x62, 0x19, 0xc4, 0x27, 0x28, 0x52, 0x34, 0x26, 0x8b,
	0x48, 0x48, 0x5c, 0x3a, 0xd3, 0x2a, 0xb1, 0xe1, 0x4e, 0xe4, 0x06, 0xb9, 0x18, 0xb9, 0x90, 0xa7,
	0xca, 0x03, 0xec, 0x09, 0x44, 0x73, 0x53, 0x98, 0x8d, 0xce, 0x4a, 0xa4, 0x76, 0xb7, 0x1e, 0x29,
	0xf4, 0x07, 0xc7, 0xf1, 0x48, 0xd3, 0x99, 0x3f, 0x84, 0xf0, 0x8d, 0xac, 0xd0, 0x6d, 0x0a, 0xc1,
	0x3b, 0xca, 0x58, 0xc8, 0x83, 0xcf, 0xd6, 0x7a, 0x4f, 0x49, 0x42, 0xeb, 0x3a, 0x9f, 0xc1, 0x98,
	0xcb, 0x8a, 0x3d, 0x82, 0x88, 0xee, 0x6a, 0xbc, 0x67, 0xe5, 0xee, 0xf5, 0x91, 0x09, 0xe7, 0xd1,
	0x9a, 0xe8, 0xfc, 0x57, 0x00, 0xd3, 0x61, 0xd2, 0xd8, 0x01, 0x24, 0xad, 0xed, 0xab, 0x92, 0xe8,
	0xd6, 0x66, 0xf7, 0x01, 0x1c, 0xf7, 0xba, 0xa8, 0x05, 0xbd, 0x99, 0x72, 0xd0, 0x1d, 0xc2, 0x1e,
	0xc0, 0x04, 0x1f, 0xf7, 0x65, 0xb9, 0x33, 0x48, 0xb1, 0xac, 0xf8, 0x44, 0x21, 0xc5, 0x9e, 0x63,
	0xf8, 0x4f, 0x95, 0x29, 0xaf, 0x84, 0xca, 0x0a, 0x12, 0x9d, 0x0d, 0xaa, 0xd7, 0x32, 0xad, 0xf0,
	0x44, 0xb7, 0xf6, 0x20, 0x51, 0x17, 0xb7, 0x27, 0x8a, 0xe5, 0x30, 0xb5, 0x52, 0x3a, 0x19, 0x25,
	0x7d, 0x74, 0xaa, 0x06, 0x58, 0xfe, 0x33, 0x00, 0xe8, 0x0b, 0x4f, 0x9d, 0x68, 0xad, 0xbe, 0x13,
	0xc9, 0xb2, 0xd9, 0xb0, 0xf8, 0x40, 0x6f, 0xa2, 0x5a, 0xdb, 0x66, 0xe3, 0x44, 0x36, 0xcb, 0x6a,
	0x21, 0x9a, 0x52, 0xa0, 0xe6, 0xe0, 0x68, 0xc4, 0xa1, 0xec, 0x10, 0xc7, 0x63, 0x99, 0xcf, 0xab,
	0xba, 0x32, 0xd8, 0x76, 0xb6, 0x42, 0xc8, 0x7b, 0xe4, 0x3f, 0x45, 0x61, 0x34, 0x6a, 0x39, 0xf7,
	0x4b, 0x27, 0x09, 0x54, 0x87, 0xe4, 0xdf, 0x50, 0x50, 0xdf, 0x99, 0xb6, 0xf1, 0x30, 0x80, 0x32,
	0x2f, 0x0b, 0x23, 0x48, 0x13, 0x8e, 0x90, 0xf6, 0x80, 0x6d, 0xd8, 0xd3, 0x66, 0x41, 0x9c, 0x1b,
	0xaf, 0x58, 0x38, 0xb3, 0xf3, 0xb3, 0xa1, 0x48, 0x93, 0xf7, 0xb3, 0x40, 0xeb, 0x47, 0xdc, 0xa4,
	0xf3, 0x23, 0x06, 0x13, 0x78, 0xa6, 0xcf, 0x65, 0x73, 0x49, 0xf3, 0x92, 0xf0, 0xa8, 0x22, 0x2b,
	0xff, 0x3e, 0x82, 0x3d, 0x3f, 0xef, 0xf3, 0x4d, 0x5d, 0x17, 0xea, 0x7a, 0x38, 0x2e, 0xc1, 0x1f,
	0xe3, 0x72, 0xc3, 0xe0, 0xf7, 0xeb, 0x62, 0xbc, 0xb5, 0x2e, 0x66, 0x98, 0x48, 0xaa, 0x2b, 0x7d,
	0xea, 0xef, 0xfb, 0x20, 0x72, 0xad, 0x6a, 0x47, 0xdb, 0xad, 0x82, 0xf0, 0x86, 0x55, 0x10, 0xaa,
	0xed, 0x35, 0x10, 0xdd, 0xbe, 0x06, 0xfa, 0x82, 0xc6, 0xff, 0x30, 0xce, 0x5f, 0x03, 0xd8, 0xf5,
	0x99, 0xe1, 0x42, 0x6f, 0x56, 0xc6, 0x26, 0xe6, 0xc5, 0x6a, 0x35, 0xaf, 0xbe, 0x88, 0x76, 0xbc,
	0xe3, 0xc2, 0x99, 0xd4, 0xd2, 0xc2, 0x6c, 0x54, 0x23, 0x16, 0x44, 0xbb, 0x79, 0x9f, 0xaa, 0x01,
	0x86, 0x7b, 0x31, 0xf5, 0xf1, 0xfc, 0x08, 0xb2, 0xfe, 0x07, 0xdd, 0x53, 0x6e, 0xc7, 0xd9, 0x4b,
	0xf9, 0x8e, 0xf3, 0x38, 0xfd, 0x88, 0xb9, 0xbf, 0x88, 0x68, 0x41, 0x3f, 0xfb, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0x69, 0xdb, 0x61, 0x4c, 0xd0, 0x05, 0x00, 0x00,
}
