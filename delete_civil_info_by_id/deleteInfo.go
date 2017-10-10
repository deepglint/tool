package main

import (
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/json"
	"io"
	"bufio"
	"bytes"
)

type Repo struct {
	RepoId       string
	RepoName     string
	Timestamp    int64
	Comment      string
	PicCount     int64
	NameListAttr int64
	Status       int64
}

type DataResp struct {
	AllSize      int64
	ReturnedSize int64
	FaceRepos    []Repo
}

type ThorResp struct {
	Data DataResp
}

type CivilAttr struct {
	Name             string     `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name"`
	IdType           int64     `protobuf:"varint,2,opt,name=IdType,json=idType,enum=dg.model.IdType" json:"IdType"`
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
	Status           int64 		`protobuf:"varint,98,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
	AttrId           string     `protobuf:"bytes,99,opt,name=AttrId,json=attrId" json:"AttrId"`
}

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
	Status            int64 	 `protobuf:"varint,98,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
	ImageId           string     `protobuf:"bytes,99,opt,name=ImageId,json=imageId" json:"ImageId"`
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
	Status       int64     `protobuf:"varint,99,opt,name=Status,json=status,enum=dg.model.TaskStatus" json:"Status"`
}

type RegisteredFaceData struct {
	AllSize      int64
	ReturnedSize int64
	RegisteredFaces []RegisteredFace
}

type RegisteredFaceResult struct {
	Code 		int64
    Msg 		string
    Redirect 	string
    Data		RegisteredFaceData
}

type BaseConds struct {
	QueryId           string            `protobuf:"bytes,1,opt,name=QueryId,json=queryId" json:"QueryId"`
	Columns           []string          `protobuf:"bytes,6,rep,name=Columns,json=columns" json:"Columns"`
	Offset            int32             `protobuf:"varint,8,opt,name=Offset,json=offset" json:"Offset"`
	Limit             int32             `protobuf:"varint,9,opt,name=Limit,json=limit" json:"Limit"`
	SortBy            string            `protobuf:"bytes,10,opt,name=SortBy,json=sortBy" json:"SortBy"`
	SortOrderAsc      bool              `protobuf:"varint,11,opt,name=SortOrderAsc,json=sortOrderAsc" json:"SortOrderAsc"`
}

type RegisteredConds struct {
	Ids        []string          `protobuf:"bytes,1,rep,name=Ids,json=ids" json:"Ids"`
	Repos      []string          `protobuf:"bytes,2,rep,name=Repos,json=repos" json:"Repos"`
	Conditions map[string]string `protobuf:"bytes,3,rep,name=Conditions,json=conditions" json:"Conditions" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

type ImageConds struct {
	Images     []*ImageResult `protobuf:"bytes,1,rep,name=Images,json=images" json:"Images"`
	Confidence float32        `protobuf:"fixed32,2,opt,name=Confidence,json=confidence" json:"Confidence"`
	CountLimit int32          `protobuf:"varint,3,opt,name=CountLimit,json=countLimit" json:"CountLimit"`
}

type AuxiliaryConds struct {
	IsWarned  int32 `protobuf:"varint,1,opt,name=IsWarned,json=isWarned" json:"IsWarned"`
	IsChecked int32 `protobuf:"varint,2,opt,name=IsChecked,json=isChecked" json:"IsChecked"`
	IsRanked  int32 `protobuf:"varint,3,opt,name=IsRanked,json=isRanked" json:"IsRanked"`
	Status    int32 `protobuf:"varint,4,opt,name=Status,json=status" json:"Status"`
}

type RegisteredFaceQuery struct {
	BaseConds      BaseConds       `protobuf:"bytes,1,opt,name=BaseConds,json=baseConds" json:"BaseConds"`
	AttrConds      RegisteredConds `protobuf:"bytes,2,opt,name=AttrConds,json=attrConds" json:"AttrConds"`
	ImageConds     ImageConds      `protobuf:"bytes,3,opt,name=ImageConds,json=imageConds" json:"ImageConds"`
	AuxiliaryConds AuxiliaryConds  `protobuf:"bytes,4,opt,name=AuxiliaryConds,json=auxiliaryConds" json:"AuxiliaryConds"`
}



func main(){
	if len(os.Args) != 4 {
		fmt.Println("格式: ./deleteInfo ip filename reponame")
		os.Exit(0)
	}

	ip := os.Args[1]
	filename := os.Args[2]
	repo_name := os.Args[3]

	resp, err := http.Get("http://" + ip + ":9876/api/face/repos")
	if err != nil {
		fmt.Println("Get repos error, exit", err)
		os.Exit(0)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Get repos error, exit", err)
		os.Exit(0)
	}
	var repos ThorResp
	err = json.Unmarshal(body, &repos)
	if err != nil {
		fmt.Println("Json unmarshal error, exit", err)
		os.Exit(0)
	}

	var repo_id []string
	for _, repo := range repos.Data.FaceRepos {
		if repo.RepoName == repo_name {
			repo_id = append(repo_id, repo.RepoId)
			fmt.Println("repo_id: ", repo_id, repo_name)
			break
		}
	}

	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("cannot open input file, exit", err)
		os.Exit(0)
	}

	r := bufio.NewReader(f)
	num := 0
	for {
		buf, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		buf = strings.Split(buf, "\n")[0]
		conds := make(map[string]string)
		conds["2"] = buf
		q := RegisteredFaceQuery{
			BaseConds :BaseConds{
				Offset:0,
				Limit:100,
			},
			AttrConds: RegisteredConds{
				Repos: repo_id,
				Conditions: conds,
			},
		}
		q_body, err := json.Marshal(q)
		// fmt.Printf("%v", q)
		if err != nil {
			fmt.Println("json marshal error", err)
			os.Exit(0)
		}
		resp, err = http.Post("http://" + ip + ":9876/api/face/civils", "application/json;charset=utf-8", bytes.NewBuffer(q_body))
		if err != nil {
			fmt.Println("get civil id error", err)
			os.Exit(0)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Get repos error, exit", err)
			os.Exit(0)
		}
		// fmt.Println(string(body))
		var registered_civil_result RegisteredFaceResult
		err = json.Unmarshal(body, &registered_civil_result)
		if err != nil {
			fmt.Println("Unmarshal civil result error", err)
			os.Exit(0)
		}
		flag := false
		for _, face := range registered_civil_result.Data.RegisteredFaces {
			civil_attr_id := face.CivilAttr.AttrId
			fmt.Println("civil_attr_id", civil_attr_id)
			req, err := http.NewRequest("DELETE", "http://" + ip + ":9876/api/face/civilattr?id=" + civil_attr_id, nil)
			if err != nil {
				fmt.Println(err)
				continue
			}
			_, err = http.DefaultClient.Do(req)
			if err != nil {
				fmt.Println(err)
			}
			flag = true
		}
		if flag {
			fmt.Printf("No.%d deleted, id: %s\n", num, buf)
		} else {
			fmt.Printf("No.%d not existed, id: %s\n", num, buf)
		}
		num++
	}
	fmt.Println("Process done")
}
