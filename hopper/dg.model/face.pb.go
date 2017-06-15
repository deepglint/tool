// Code generated by protoc-gen-go.
// source: face.proto
// DO NOT EDIT!

package dg_model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type IdType int32

const (
	IdType_Id_Type_Unknown   IdType = 0
	IdType_Id_Type_IdNo      IdType = 1
	IdType_Id_Type_Passport  IdType = 2
	IdType_Id_Type_Officer   IdType = 3
	IdType_Id_Type_Sergeant  IdType = 4
	IdType_Id_Type_Temporary IdType = 5
)

var IdType_name = map[int32]string{
	0: "Id_Type_Unknown",
	1: "Id_Type_IdNo",
	2: "Id_Type_Passport",
	3: "Id_Type_Officer",
	4: "Id_Type_Sergeant",
	5: "Id_Type_Temporary",
}
var IdType_value = map[string]int32{
	"Id_Type_Unknown":   0,
	"Id_Type_IdNo":      1,
	"Id_Type_Passport":  2,
	"Id_Type_Officer":   3,
	"Id_Type_Sergeant":  4,
	"Id_Type_Temporary": 5,
}

func (x IdType) String() string {
	return proto.EnumName(IdType_name, int32(x))
}
func (IdType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type CivilAttr struct {
	Name             string     `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name"`
	IdType           IdType     `protobuf:"varint,2,opt,name=IdType,json=idType,enum=dg.model.IdType" json:"IdType"`
	IdNo             string     `protobuf:"bytes,3,opt,name=IdNo,json=idNo" json:"IdNo"`
	GenderId         int32      `protobuf:"varint,4,opt,name=GenderId,json=genderId" json:"GenderId"`
	GenderConfidence float32    `protobuf:"fixed32,5,opt,name=GenderConfidence,json=genderConfidence" json:"GenderConfidence"`
	AgeId            int32      `protobuf:"varint,6,opt,name=AgeId,json=ageId" json:"AgeId"`
	AgeConfidence    float32    `protobuf:"fixed32,7,opt,name=AgeConfidence,json=ageConfidence" json:"AgeConfidence"`
	NationId         int32      `protobuf:"varint,8,opt,name=NationId,json=nationId" json:"NationId"`
	NationConfidence float32    `protobuf:"fixed32,9,opt,name=NationConfidence,json=nationConfidence" json:"NationConfidence"`
	GlassId          int32      `protobuf:"varint,10,opt,name=GlassId,json=glassId" json:"GlassId"`
	GlassConfidence  float32    `protobuf:"fixed32,11,opt,name=GlassConfidence,json=glassConfidence" json:"GlassConfidence"`
	Addr             string     `protobuf:"bytes,12,opt,name=Addr,json=addr" json:"Addr"`
	Birthday         string     `protobuf:"bytes,13,opt,name=Birthday,json=birthday" json:"Birthday"`
	Comment          string     `protobuf:"bytes,14,opt,name=Comment,json=comment" json:"Comment"`
	Status           TaskStatus `protobuf:"varint,98,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
	AttrId           string     `protobuf:"bytes,99,opt,name=AttrId,json=attrId" json:"AttrId"`
}

func (m *CivilAttr) Reset()                    { *m = CivilAttr{} }
func (m *CivilAttr) String() string            { return proto.CompactTextString(m) }
func (*CivilAttr) ProtoMessage()               {}
func (*CivilAttr) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type ImageResult struct {
	ImageUri          string     `protobuf:"bytes,1,opt,name=ImageUri,json=imageUri" json:"ImageUri"`
	ThumbnailImageUri string     `protobuf:"bytes,2,opt,name=ThumbnailImageUri,json=thumbnailImageUri" json:"ThumbnailImageUri"`
	CutboardImageUri  string     `protobuf:"bytes,3,opt,name=CutboardImageUri,json=cutboardImageUri" json:"CutboardImageUri"`
	CutboardX         int32      `protobuf:"varint,4,opt,name=CutboardX,json=cutboardX" json:"CutboardX"`
	CutboardY         int32      `protobuf:"varint,5,opt,name=CutboardY,json=cutboardY" json:"CutboardY"`
	CutboardWidth     int32      `protobuf:"varint,6,opt,name=CutboardWidth,json=cutboardWidth" json:"CutboardWidth"`
	CutboardHeight    int32      `protobuf:"varint,7,opt,name=CutboardHeight,json=cutboardHeight" json:"CutboardHeight"`
	CutboardResWidth  int32      `protobuf:"varint,8,opt,name=CutboardResWidth,json=cutboardResWidth" json:"CutboardResWidth"`
	CutboardResHeight int32      `protobuf:"varint,9,opt,name=CutboardResHeight,json=cutboardResHeight" json:"CutboardResHeight"`
	BinData           string     `protobuf:"bytes,16,opt,name=BinData,json=binData" json:"BinData"`
	Feature           string     `protobuf:"bytes,17,opt,name=Feature,json=feature" json:"Feature"`
	IsRanked          int32      `protobuf:"varint,18,opt,name=IsRanked,json=isRanked" json:"IsRanked"`
	Status            TaskStatus `protobuf:"varint,98,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
	ImageId           string     `protobuf:"bytes,99,opt,name=ImageId,json=imageId" json:"ImageId"`
}

func (m *ImageResult) Reset()                    { *m = ImageResult{} }
func (m *ImageResult) String() string            { return proto.CompactTextString(m) }
func (*ImageResult) ProtoMessage()               {}
func (*ImageResult) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

type CapturedFace struct {
	FaceId      string       `protobuf:"bytes,1,opt,name=FaceId,json=faceId" json:"FaceId"`
	FaceReId    string       `protobuf:"bytes,2,opt,name=FaceReId,json=faceReId" json:"FaceReId"`
	Timestamp   int64        `protobuf:"varint,3,opt,name=Timestamp,json=timestamp" json:"Timestamp"`
	SensorId    string       `protobuf:"bytes,4,opt,name=SensorId,json=sensorId" json:"SensorId"`
	SensorName  string       `protobuf:"bytes,5,opt,name=SensorName,json=sensorName" json:"SensorName"`
	Confidence  float32      `protobuf:"fixed32,6,opt,name=Confidence,json=confidence" json:"Confidence"`
	IsWarned    int32        `protobuf:"varint,7,opt,name=IsWarned,json=isWarned" json:"IsWarned"`
	CivilAttr   *CivilAttr   `protobuf:"bytes,8,opt,name=CivilAttr,json=civilAttr" json:"CivilAttr"`
	ImageResult *ImageResult `protobuf:"bytes,9,opt,name=ImageResult,json=imageResult" json:"ImageResult"`
	Status      TaskStatus   `protobuf:"varint,99,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
}

func (m *CapturedFace) Reset()                    { *m = CapturedFace{} }
func (m *CapturedFace) String() string            { return proto.CompactTextString(m) }
func (*CapturedFace) ProtoMessage()               {}
func (*CapturedFace) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *CapturedFace) GetCivilAttr() *CivilAttr {
	if m != nil {
		return m.CivilAttr
	}
	return nil
}

func (m *CapturedFace) GetImageResult() *ImageResult {
	if m != nil {
		return m.ImageResult
	}
	return nil
}

type CapturedFaceResult struct {
	AllSize       int32           `protobuf:"varint,1,opt,name=AllSize,json=allSize" json:"AllSize"`
	ReturnedSize  int32           `protobuf:"varint,2,opt,name=ReturnedSize,json=returnedSize" json:"ReturnedSize"`
	CapturedFaces []*CapturedFace `protobuf:"bytes,3,rep,name=CapturedFaces,json=capturedFaces" json:"CapturedFaces"`
}

func (m *CapturedFaceResult) Reset()                    { *m = CapturedFaceResult{} }
func (m *CapturedFaceResult) String() string            { return proto.CompactTextString(m) }
func (*CapturedFaceResult) ProtoMessage()               {}
func (*CapturedFaceResult) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *CapturedFaceResult) GetCapturedFaces() []*CapturedFace {
	if m != nil {
		return m.CapturedFaces
	}
	return nil
}

type FaceRepo struct {
	RepoId       string     `protobuf:"bytes,1,opt,name=RepoId,json=repoId" json:"RepoId"`
	RepoName     string     `protobuf:"bytes,2,opt,name=RepoName,json=repoName" json:"RepoName"`
	Timestamp    int64      `protobuf:"varint,3,opt,name=Timestamp,json=timestamp" json:"Timestamp"`
	Comment      string     `protobuf:"bytes,4,opt,name=Comment,json=comment" json:"Comment"`
	PicCount     int32      `protobuf:"varint,5,opt,name=PicCount,json=picCount" json:"PicCount"`
	NameListAttr int32      `protobuf:"varint,6,opt,name=NameListAttr,json=nameListAttr" json:"NameListAttr"`
	Status       TaskStatus `protobuf:"varint,99,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
}

func (m *FaceRepo) Reset()                    { *m = FaceRepo{} }
func (m *FaceRepo) String() string            { return proto.CompactTextString(m) }
func (*FaceRepo) ProtoMessage()               {}
func (*FaceRepo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

type FaceRepoResult struct {
	AllSize      int32       `protobuf:"varint,1,opt,name=AllSize,json=allSize" json:"AllSize"`
	ReturnedSize int32       `protobuf:"varint,2,opt,name=ReturnedSize,json=returnedSize" json:"ReturnedSize"`
	FaceRepos    []*FaceRepo `protobuf:"bytes,3,rep,name=FaceRepos,json=faceRepos" json:"FaceRepos"`
}

func (m *FaceRepoResult) Reset()                    { *m = FaceRepoResult{} }
func (m *FaceRepoResult) String() string            { return proto.CompactTextString(m) }
func (*FaceRepoResult) ProtoMessage()               {}
func (*FaceRepoResult) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *FaceRepoResult) GetFaceRepos() []*FaceRepo {
	if m != nil {
		return m.FaceRepos
	}
	return nil
}

type RegisteredFace struct {
	CivilId      string         `protobuf:"bytes,1,opt,name=CivilId,json=civilId" json:"CivilId"`
	Timestamp    int64          `protobuf:"varint,2,opt,name=Timestamp,json=timestamp" json:"Timestamp"`
	FaceRepoId   string         `protobuf:"bytes,3,opt,name=FaceRepoId,json=faceRepoId" json:"FaceRepoId"`
	FaceRepoName string         `protobuf:"bytes,4,opt,name=FaceRepoName,json=faceRepoName" json:"FaceRepoName"`
	Confidence   float32        `protobuf:"fixed32,5,opt,name=Confidence,json=confidence" json:"Confidence"`
	NameListAttr int32          `protobuf:"varint,6,opt,name=NameListAttr,json=nameListAttr" json:"NameListAttr"`
	CivilAttr    *CivilAttr     `protobuf:"bytes,8,opt,name=CivilAttr,json=civilAttr" json:"CivilAttr"`
	ImageResults []*ImageResult `protobuf:"bytes,9,rep,name=ImageResults,json=imageResults" json:"ImageResults"`
	Status       TaskStatus     `protobuf:"varint,99,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
}

func (m *RegisteredFace) Reset()                    { *m = RegisteredFace{} }
func (m *RegisteredFace) String() string            { return proto.CompactTextString(m) }
func (*RegisteredFace) ProtoMessage()               {}
func (*RegisteredFace) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *RegisteredFace) GetCivilAttr() *CivilAttr {
	if m != nil {
		return m.CivilAttr
	}
	return nil
}

func (m *RegisteredFace) GetImageResults() []*ImageResult {
	if m != nil {
		return m.ImageResults
	}
	return nil
}

type RegisteredFaceResult struct {
	AllSize         int32             `protobuf:"varint,1,opt,name=AllSize,json=allSize" json:"AllSize"`
	ReturnedSize    int32             `protobuf:"varint,2,opt,name=ReturnedSize,json=returnedSize" json:"ReturnedSize"`
	RegisteredFaces []*RegisteredFace `protobuf:"bytes,3,rep,name=RegisteredFaces,json=registeredFaces" json:"RegisteredFaces"`
}

func (m *RegisteredFaceResult) Reset()                    { *m = RegisteredFaceResult{} }
func (m *RegisteredFaceResult) String() string            { return proto.CompactTextString(m) }
func (*RegisteredFaceResult) ProtoMessage()               {}
func (*RegisteredFaceResult) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *RegisteredFaceResult) GetRegisteredFaces() []*RegisteredFace {
	if m != nil {
		return m.RegisteredFaces
	}
	return nil
}

type WarnedFace struct {
	EventId         string       `protobuf:"bytes,1,opt,name=EventId,json=eventId" json:"EventId"`
	EventReId       string       `protobuf:"bytes,2,opt,name=EventReId,json=eventReId" json:"EventReId"`
	Timestamp       int64        `protobuf:"varint,3,opt,name=Timestamp,json=timestamp" json:"Timestamp"`
	SensorId        string       `protobuf:"bytes,4,opt,name=SensorId,json=sensorId" json:"SensorId"`
	SensorName      string       `protobuf:"bytes,5,opt,name=SensorName,json=sensorName" json:"SensorName"`
	FaceRepoId      string       `protobuf:"bytes,8,opt,name=FaceRepoId,json=faceRepoId" json:"FaceRepoId"`
	FaceRepoName    string       `protobuf:"bytes,9,opt,name=FaceRepoName,json=faceRepoName" json:"FaceRepoName"`
	NameListAttr    int32        `protobuf:"varint,10,opt,name=NameListAttr,json=nameListAttr" json:"NameListAttr"`
	FaceId          string       `protobuf:"bytes,16,opt,name=FaceId,json=faceId" json:"FaceId"`
	FaceReId        string       `protobuf:"bytes,17,opt,name=FaceReId,json=faceReId" json:"FaceReId"`
	CapturedImage   *ImageResult `protobuf:"bytes,18,opt,name=CapturedImage,json=capturedImage" json:"CapturedImage"`
	CivilId         string       `protobuf:"bytes,32,opt,name=CivilId,json=civilId" json:"CivilId"`
	CivilAttr       *CivilAttr   `protobuf:"bytes,33,opt,name=CivilAttr,json=civilAttr" json:"CivilAttr"`
	RegisteredImage *ImageResult `protobuf:"bytes,34,opt,name=RegisteredImage,json=registeredImage" json:"RegisteredImage"`
	RuleId          string       `protobuf:"bytes,64,opt,name=RuleId,json=ruleId" json:"RuleId"`
	Rule            string       `protobuf:"bytes,65,opt,name=Rule,json=rule" json:"Rule"`
	Confidence      float32      `protobuf:"fixed32,80,opt,name=Confidence,json=confidence" json:"Confidence"`
	IsChecked       int32        `protobuf:"varint,81,opt,name=IsChecked,json=isChecked" json:"IsChecked"`
	Comment         string       `protobuf:"bytes,82,opt,name=Comment,json=comment" json:"Comment"`
	UserIds         string       `protobuf:"bytes,83,opt,name=UserIds,json=userIds" json:"UserIds"`
	Status          TaskStatus   `protobuf:"varint,99,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
}

func (m *WarnedFace) Reset()                    { *m = WarnedFace{} }
func (m *WarnedFace) String() string            { return proto.CompactTextString(m) }
func (*WarnedFace) ProtoMessage()               {}
func (*WarnedFace) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *WarnedFace) GetCapturedImage() *ImageResult {
	if m != nil {
		return m.CapturedImage
	}
	return nil
}

func (m *WarnedFace) GetCivilAttr() *CivilAttr {
	if m != nil {
		return m.CivilAttr
	}
	return nil
}

func (m *WarnedFace) GetRegisteredImage() *ImageResult {
	if m != nil {
		return m.RegisteredImage
	}
	return nil
}

type WarnedFaceResult struct {
	AllSize      int32         `protobuf:"varint,1,opt,name=AllSize,json=allSize" json:"AllSize"`
	ReturnedSize int32         `protobuf:"varint,2,opt,name=ReturnedSize,json=returnedSize" json:"ReturnedSize"`
	WarnedFaces  []*WarnedFace `protobuf:"bytes,3,rep,name=WarnedFaces,json=warnedFaces" json:"WarnedFaces"`
	ReId         string        `protobuf:"bytes,4,opt,name=ReId,json=reId" json:"ReId"`
	IsInform     int32         `protobuf:"varint,5,opt,name=IsInform,json=isInform" json:"IsInform"`
	FaceId       string        `protobuf:"bytes,6,opt,name=FaceId,json=faceId" json:"FaceId"`
}

func (m *WarnedFaceResult) Reset()                    { *m = WarnedFaceResult{} }
func (m *WarnedFaceResult) String() string            { return proto.CompactTextString(m) }
func (*WarnedFaceResult) ProtoMessage()               {}
func (*WarnedFaceResult) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *WarnedFaceResult) GetWarnedFaces() []*WarnedFace {
	if m != nil {
		return m.WarnedFaces
	}
	return nil
}

func init() {
	proto.RegisterType((*CivilAttr)(nil), "dg.model.CivilAttr")
	proto.RegisterType((*ImageResult)(nil), "dg.model.ImageResult")
	proto.RegisterType((*CapturedFace)(nil), "dg.model.CapturedFace")
	proto.RegisterType((*CapturedFaceResult)(nil), "dg.model.CapturedFaceResult")
	proto.RegisterType((*FaceRepo)(nil), "dg.model.FaceRepo")
	proto.RegisterType((*FaceRepoResult)(nil), "dg.model.FaceRepoResult")
	proto.RegisterType((*RegisteredFace)(nil), "dg.model.RegisteredFace")
	proto.RegisterType((*RegisteredFaceResult)(nil), "dg.model.RegisteredFaceResult")
	proto.RegisterType((*WarnedFace)(nil), "dg.model.WarnedFace")
	proto.RegisterType((*WarnedFaceResult)(nil), "dg.model.WarnedFaceResult")
	proto.RegisterEnum("dg.model.IdType", IdType_name, IdType_value)
}

func init() { proto.RegisterFile("face.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 1217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x57, 0xcf, 0x8f, 0xdb, 0x44,
	0x14, 0x26, 0xbf, 0x9c, 0x64, 0x92, 0xcd, 0x3a, 0xd3, 0x6d, 0x65, 0x55, 0x08, 0x95, 0x08, 0xa1,
	0xaa, 0x42, 0x2b, 0x58, 0x24, 0x10, 0x02, 0x09, 0x76, 0x03, 0x2d, 0x91, 0xd0, 0xb2, 0x78, 0xb3,
	0x6a, 0x7b, 0xaa, 0x26, 0xf6, 0x38, 0x19, 0x6d, 0x6c, 0x47, 0xb6, 0xd3, 0x6a, 0x39, 0x71, 0xe3,
	0xca, 0x91, 0x3f, 0x0b, 0xfe, 0x07, 0x8e, 0x9c, 0xb8, 0x23, 0x78, 0x6f, 0xc6, 0x63, 0x8f, 0xdd,
	0x2e, 0xdb, 0x45, 0xdb, 0x93, 0xfd, 0xbe, 0x37, 0xf3, 0xe6, 0xf9, 0x7d, 0xdf, 0x7b, 0xb6, 0x09,
	0x09, 0x98, 0xc7, 0xf7, 0x37, 0x49, 0x9c, 0xc5, 0xb4, 0xe7, 0x2f, 0xf7, 0xc3, 0xd8, 0xe7, 0xeb,
	0xbb, 0x34, 0xe5, 0x2c, 0xf1, 0x56, 0x5e, 0x1c, 0x86, 0x71, 0xa4, 0xbc, 0x93, 0xbf, 0x5b, 0xa4,
	0x3f, 0x15, 0xcf, 0xc5, 0xfa, 0x30, 0xcb, 0x12, 0x4a, 0x49, 0xfb, 0x98, 0x85, 0xdc, 0x69, 0xdc,
	0x6b, 0xdc, 0xef, 0xbb, 0xed, 0x08, 0xee, 0xe9, 0x7d, 0x62, 0xcd, 0xfc, 0xf9, 0xc5, 0x86, 0x3b,
	0x4d, 0x40, 0x47, 0x07, 0xf6, 0xbe, 0x0e, 0xb8, 0xaf, 0x70, 0xd7, 0x12, 0xf2, 0x8a, 0xbb, 0x67,
	0xfe, 0x71, 0xec, 0xb4, 0xd4, 0x6e, 0x01, 0xf7, 0xf4, 0x2e, 0xe9, 0x3d, 0xe2, 0x91, 0xcf, 0x93,
	0x99, 0xef, 0xb4, 0x01, 0xef, 0xb8, 0xbd, 0x65, 0x6e, 0xd3, 0x07, 0xc4, 0x56, 0xbe, 0x69, 0x1c,
	0x05, 0xc2, 0xe7, 0x91, 0xc7, 0x9d, 0x0e, 0xac, 0x69, 0xba, 0xf6, 0xb2, 0x86, 0xd3, 0x3d, 0xd2,
	0x39, 0x5c, 0x72, 0x08, 0x62, 0xc9, 0x20, 0x1d, 0x86, 0x06, 0x7d, 0x8f, 0xec, 0x00, 0x6a, 0x6c,
	0xef, 0xca, 0xed, 0x3b, 0xcc, 0x04, 0x31, 0x87, 0x63, 0x96, 0x89, 0x38, 0x82, 0xed, 0x3d, 0x95,
	0x43, 0x94, 0xdb, 0x98, 0x83, 0xf2, 0x19, 0x41, 0xfa, 0x2a, 0x87, 0xa8, 0x86, 0x53, 0x87, 0x74,
	0x1f, 0xad, 0x59, 0x9a, 0x42, 0x18, 0x22, 0xc3, 0x74, 0x97, 0xca, 0x84, 0x1a, 0xed, 0x4a, 0x8f,
	0x11, 0x64, 0x20, 0x83, 0xec, 0x2e, 0xab, 0x30, 0xd6, 0xe8, 0xd0, 0xf7, 0x13, 0x67, 0xa8, 0x6a,
	0xc4, 0xe0, 0x1e, 0xf3, 0x3b, 0x12, 0x49, 0xb6, 0xf2, 0xd9, 0x85, 0xb3, 0x23, 0xf1, 0xde, 0x22,
	0xb7, 0xf1, 0xcc, 0x29, 0xf0, 0xc5, 0xa3, 0xcc, 0x19, 0x49, 0x57, 0xd7, 0x53, 0x26, 0xfd, 0x80,
	0x58, 0xa7, 0x19, 0xcb, 0xb6, 0xa9, 0xb3, 0x90, 0xbc, 0xec, 0x95, 0xbc, 0xcc, 0x59, 0x7a, 0xae,
	0x7c, 0xae, 0x95, 0xca, 0x2b, 0xbd, 0x43, 0x2c, 0x64, 0x18, 0x52, 0xf7, 0x64, 0x18, 0x8b, 0x49,
	0x6b, 0xf2, 0x4f, 0x8b, 0x0c, 0x66, 0x21, 0xd4, 0xcb, 0xe5, 0xe9, 0x76, 0x9d, 0x61, 0x2e, 0xd2,
	0x3c, 0x4b, 0x44, 0xae, 0x82, 0x9e, 0xc8, 0x6d, 0x38, 0x71, 0x3c, 0x5f, 0x6d, 0xc3, 0x45, 0xc4,
	0xc4, 0xba, 0x58, 0xd4, 0x94, 0x8b, 0xc6, 0x59, 0xdd, 0x81, 0x95, 0x9d, 0x6e, 0xb3, 0x45, 0xcc,
	0x12, 0xbf, 0x58, 0xac, 0x94, 0x61, 0x7b, 0x35, 0x9c, 0xbe, 0x0d, 0x22, 0xcc, 0xb1, 0x27, 0xb9,
	0x4c, 0xfa, 0x7a, 0xd1, 0x13, 0xd3, 0xfb, 0x54, 0x0a, 0xc4, 0xf0, 0x3e, 0x45, 0x0d, 0x68, 0xef,
	0x63, 0xe1, 0x67, 0xab, 0x5c, 0x21, 0x3b, 0x9e, 0x09, 0xd2, 0xf7, 0xc9, 0x48, 0xaf, 0xfa, 0x96,
	0x8b, 0xe5, 0x2a, 0x93, 0x52, 0xe9, 0xb8, 0x23, 0xaf, 0x82, 0x9a, 0x59, 0x43, 0x45, 0x54, 0x40,
	0xa5, 0x99, 0x22, 0x6b, 0x8d, 0x63, 0x3d, 0x8c, 0xb5, 0x79, 0xd8, 0xbe, 0x5c, 0x3c, 0xf6, 0xea,
	0x0e, 0x64, 0xf2, 0x48, 0x44, 0x5f, 0xb3, 0x8c, 0x39, 0xb6, 0x62, 0x72, 0xa1, 0x4c, 0xf4, 0x3c,
	0xe4, 0xc0, 0x52, 0xc2, 0x9d, 0xb1, 0xf2, 0x04, 0xca, 0x94, 0x6c, 0xa4, 0x2e, 0x8b, 0xce, 0xb9,
	0xef, 0x50, 0xa5, 0x5c, 0x91, 0xdb, 0xd7, 0xe4, 0x1f, 0xce, 0x90, 0xd5, 0x2e, 0x04, 0xd0, 0x15,
	0xca, 0x9c, 0xfc, 0xd5, 0x24, 0xc3, 0x29, 0xdb, 0xe0, 0x79, 0xfe, 0x43, 0x18, 0x1b, 0x28, 0x15,
	0xbc, 0xc2, 0x4a, 0x25, 0x00, 0x2b, 0x90, 0x16, 0x26, 0x83, 0xb8, 0x8b, 0x1e, 0xc5, 0x7a, 0x2f,
	0xc8, 0x6d, 0xa4, 0x68, 0x2e, 0x42, 0x0e, 0x87, 0x85, 0x1b, 0xc9, 0x72, 0xcb, 0xed, 0x67, 0x1a,
	0xc0, 0x9d, 0xa7, 0x3c, 0x4a, 0x63, 0x3d, 0x04, 0x60, 0x67, 0x9a, 0xdb, 0xf4, 0x1d, 0x42, 0x94,
	0x4f, 0x0e, 0x9e, 0x8e, 0xf4, 0x92, 0xb4, 0x40, 0xd0, 0x6f, 0x74, 0x95, 0x25, 0xbb, 0x8a, 0x78,
	0x95, 0xe6, 0x9e, 0xa5, 0x8f, 0x59, 0x12, 0x41, 0x89, 0xba, 0xba, 0x44, 0xca, 0xa6, 0x1f, 0x19,
	0xb3, 0x4d, 0xb2, 0x38, 0x38, 0xb8, 0x55, 0x56, 0xa9, 0x70, 0x81, 0x9a, 0x8a, 0x09, 0xf8, 0x69,
	0xa5, 0x1d, 0x24, 0x9b, 0x83, 0x83, 0xdb, 0xc6, 0xc8, 0x2b, 0x9d, 0xee, 0x40, 0x18, 0x8d, 0x53,
	0xd2, 0xe1, 0x5d, 0x4d, 0xc7, 0xe4, 0x97, 0x06, 0xa1, 0x66, 0xd1, 0xf3, 0x20, 0xc0, 0xd2, 0xe1,
	0x7a, 0x7d, 0x2a, 0x7e, 0x54, 0x23, 0x18, 0x26, 0x0c, 0x53, 0x26, 0x9d, 0x90, 0xa1, 0xcb, 0x61,
	0x39, 0x3c, 0x96, 0x74, 0x37, 0xa5, 0x7b, 0x98, 0x18, 0x18, 0xfd, 0x02, 0x3a, 0xc1, 0x88, 0x99,
	0x02, 0x11, 0x2d, 0xc8, 0xfe, 0x8e, 0xf1, 0xc8, 0xe6, 0x91, 0x3b, 0x9e, 0xb9, 0x78, 0xf2, 0x47,
	0x43, 0xf3, 0xbb, 0x89, 0x51, 0x03, 0x78, 0x2d, 0x35, 0x90, 0x48, 0x0b, 0xab, 0x8d, 0xb8, 0xe4,
	0x2a, 0xd7, 0x40, 0x92, 0xdb, 0x57, 0x68, 0xc0, 0x18, 0x64, 0xed, 0xea, 0x20, 0x83, 0x98, 0x27,
	0xc2, 0x9b, 0xc6, 0x5b, 0x70, 0xa9, 0xee, 0xee, 0x6d, 0x72, 0x1b, 0x1f, 0x1b, 0x63, 0x7f, 0x27,
	0xd2, 0x4c, 0x92, 0xa8, 0x7a, 0x7b, 0x18, 0x19, 0xd8, 0x35, 0x2b, 0xff, 0x53, 0x83, 0x8c, 0xf4,
	0x63, 0xde, 0x48, 0xd5, 0x3f, 0x24, 0x7d, 0x1d, 0x4f, 0x57, 0x9c, 0x96, 0x19, 0x14, 0x47, 0xf5,
	0x03, 0xbd, 0x68, 0xf2, 0x67, 0x93, 0x8c, 0x5c, 0xbe, 0x84, 0xfc, 0xb9, 0xee, 0x39, 0xac, 0x0e,
	0x6a, 0xb0, 0x28, 0x78, 0xd7, 0x53, 0x66, 0xb5, 0xaa, 0xcd, 0x7a, 0x55, 0xa1, 0x3b, 0xf4, 0x09,
	0xb0, 0x55, 0x8d, 0x57, 0x12, 0x14, 0x08, 0x3e, 0x80, 0xf6, 0x4b, 0xce, 0x54, 0xe9, 0x87, 0x81,
	0x81, 0xd5, 0x3a, 0xac, 0xf3, 0x52, 0x87, 0xbd, 0x0e, 0x07, 0xff, 0xa3, 0xd3, 0x3e, 0x23, 0x43,
	0xa3, 0x99, 0x52, 0x68, 0xb5, 0xd6, 0xe5, 0xad, 0x36, 0x34, 0x5a, 0x2d, 0xbd, 0x26, 0xe3, 0xbf,
	0x36, 0xc8, 0x5e, 0xb5, 0xdc, 0x37, 0xc2, 0xfb, 0x11, 0xd9, 0xad, 0x46, 0xd5, 0xec, 0x3b, 0x65,
	0x36, 0xb5, 0x63, 0x77, 0x93, 0xea, 0x86, 0xc9, 0xef, 0x1d, 0x42, 0xd4, 0xac, 0xd2, 0x2a, 0xf8,
	0xe6, 0x39, 0xb4, 0x44, 0xa9, 0x02, 0xae, 0x4c, 0x54, 0x81, 0xf4, 0x18, 0xc3, 0xb7, 0xcf, 0x35,
	0xf0, 0x66, 0xa7, 0xaf, 0xa1, 0xaf, 0xde, 0x95, 0xfa, 0xea, 0xbf, 0x42, 0x5f, 0x75, 0xfd, 0x90,
	0x57, 0xe8, 0xa7, 0x7c, 0xe7, 0xd8, 0x97, 0xbe, 0x73, 0xc6, 0xb5, 0x77, 0xce, 0xe7, 0xe5, 0xb8,
	0x93, 0x52, 0x91, 0x6f, 0xc8, 0x4b, 0x15, 0x54, 0x4c, 0x3b, 0x09, 0x9a, 0x0d, 0x77, 0xaf, 0xda,
	0x70, 0x15, 0x29, 0xbf, 0xfb, 0x5a, 0x52, 0xfe, 0xd2, 0x94, 0x82, 0xca, 0x65, 0xf2, 0x5f, 0xb9,
	0x18, 0x3a, 0x50, 0xd9, 0xe0, 0xb8, 0xdd, 0xae, 0xf1, 0x21, 0xbf, 0xca, 0xc7, 0xad, 0xb4, 0xf0,
	0x6b, 0x11, 0x71, 0xe7, 0x50, 0x7d, 0x2d, 0x22, 0x5a, 0x6b, 0xd7, 0x93, 0x97, 0xda, 0x15, 0xc4,
	0x30, 0x4b, 0xa7, 0x2b, 0xee, 0xe1, 0x47, 0xc3, 0x0f, 0xea, 0x6b, 0x49, 0x68, 0xc0, 0x1c, 0xc3,
	0x6e, 0x75, 0x0c, 0x83, 0xe7, 0x2c, 0xc5, 0xef, 0xf2, 0xd4, 0x39, 0x55, 0x9e, 0xad, 0x32, 0xaf,
	0xd9, 0x6e, 0xbf, 0x35, 0x88, 0x5d, 0x6a, 0xfa, 0x46, 0x5a, 0xed, 0x13, 0x32, 0x28, 0x23, 0xea,
	0x36, 0x33, 0xb2, 0x30, 0x8e, 0x1b, 0xbc, 0x28, 0x17, 0xca, 0xf2, 0xf1, 0x42, 0xf5, 0xed, 0x24,
	0x57, 0xd4, 0x2c, 0x9d, 0x45, 0x41, 0x9c, 0x84, 0xfa, 0x6d, 0x23, 0x72, 0xdb, 0x50, 0xa1, 0x65,
	0xaa, 0xf0, 0xc1, 0xcf, 0x0d, 0xfd, 0x0f, 0x44, 0x6f, 0x91, 0xdd, 0x99, 0xff, 0x0c, 0x6f, 0x9f,
	0x9d, 0x45, 0xe7, 0x51, 0xfc, 0x22, 0xb2, 0xdf, 0xa2, 0x36, 0x8c, 0xb2, 0x1c, 0xc4, 0x1f, 0x20,
	0xbb, 0x01, 0xbf, 0x2b, 0xb6, 0x46, 0x4e, 0xe0, 0x07, 0x60, 0x13, 0x27, 0x99, 0xdd, 0x34, 0x37,
	0x7f, 0x1f, 0x04, 0xc2, 0xe3, 0x89, 0xdd, 0x32, 0x97, 0x9e, 0xf2, 0x64, 0xc9, 0x59, 0x94, 0xd9,
	0x6d, 0x7a, 0x9b, 0x8c, 0x35, 0x3a, 0xe7, 0x21, 0xec, 0x67, 0xc9, 0x85, 0xdd, 0x59, 0x58, 0xf2,
	0xaf, 0xed, 0xe3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x73, 0xe4, 0x35, 0xe3, 0xe1, 0x0d, 0x00,
	0x00,
}
