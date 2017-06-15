package transform

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"math"
	"os"
	"sync"
	"time"

	"github.com/deepglint/filter/configs"
	"github.com/deepglint/filter/dg.model"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

type MatrixTransform struct {
	config             *configs.MatrixTransformConfig
	quality            *configs.QualityScore
	resultChanSlice    []chan []*dg_model.GenericObj
	getImageInfoChanWg sync.WaitGroup
	getFeatureWg       sync.WaitGroup
	sendOutputWg       sync.WaitGroup
	// getFeaturemutex    sync.Mutex
	// getFeatureThreadCount int
	getFeatureThreadChan chan int
	conn                 *grpc.ClientConn
	client               dg_model.WitnessServiceClient
	isLive               bool
}

func NewMatrixTransform(config *configs.FilterCenterConfig) (*MatrixTransform, error) {

	l := new(MatrixTransform)
	l.config = config.TransformConfig
	l.quality = config.Quality
	if l.config.MaxThread <= 0 {
		l.config.MaxThread = 1
	}
	return l, nil
}

func (this *MatrixTransform) StartTransform(imgChanSlice chan *configs.Img) (err error) {

	this.conn, err = grpc.Dial(this.config.RecogniseURL, grpc.WithInsecure())
	if err != nil {
		glog.Errorln("grpc connect err:", err.Error())
		return err
	}
	glog.Errorln("%v", this.conn)
	this.client = dg_model.NewWitnessServiceClient(this.conn)

	this.isLive = true
	//initilize getFeatureThreadChan
	this.getFeatureThreadChan = make(chan int, this.config.MaxThread)
	for threadCount := 0; threadCount < this.config.MaxThread; threadCount++ {
		this.getFeatureThreadChan <- 1
	}

	go func(imgChanSlice chan *configs.Img) {
		stopChan := make(chan bool)
		defer func() {
			this.conn.Close()
			this.client = nil
			this.isLive = false
			stopChan <- true
			glog.Infoln("MatrixTransform ResultChan Close")
		}()

		for img := range imgChanSlice {
			glog.Infoln("transform: ", img.Filename)
			this.getImageInfoChanWg.Add(1)
			go func(img *configs.Img) {
				defer this.getImageInfoChanWg.Done()
				this.RecognizeHandler(img)
				this.getFeatureWg.Wait()
			}(img)
		}

		this.getImageInfoChanWg.Wait()

		//wait all send message to output thread close

	}(imgChanSlice)

	return nil
}

func (this *MatrixTransform) RecognizeHandler(img *configs.Img) error {
	//require getFeatureThread
	this.getFeatureWg.Add(1)
	<-this.getFeatureThreadChan

	// glog.Infoln("transform begin to handle message ,message time:", libraFMessage.WideImageInfo.TimeStamp, " now time:", time.Now().UnixNano()/1000000)
	go func(img *configs.Img) {
		defer func() {
			//release getFeatureThread
			this.getFeatureThreadChan <- 1
			this.getFeatureWg.Done()
		}()

		// glog.Infoln("begin handle tartget  message ,message time:", libraFImageInfo.TimeStamp, " now time:", time.Now().UnixNano()/1000000)

		//set http post  ctx of matrix to check whether the pic has face
		ctx := context.Background()
		//set time out
		context.WithTimeout(ctx, time.Duration(this.config.TimeOut)*time.Second)
		//create witness

		binData := base64.StdEncoding.EncodeToString([]byte(img.Bindata))
		uri := img.Uri
		if binData == "" && uri == "" {
			return
		}
		witnessImage := dg_model.WitnessImage{
			Data: &dg_model.Image{
				BinData: binData,
				URI:     uri,
			},
			WitnessMetaData: &dg_model.SrcMetadata{
				//SensorId: imageInfo.SensorID,
				SensorIdStr: "test",
			},
		}

		//face exist check
		req := &dg_model.WitnessRequest{
			Context: &this.config.GetQualityContext,
			Image:   &witnessImage,
		}

		//send post request
		resp, err := this.client.Recognize(ctx, req)
		if err != nil {
			glog.Infoln("Matrix image quality decect failed: ", err)
			return
		}

		//check response
		if resp == nil || resp.GetResult() == nil ||
			resp.GetResult().GetImage() == nil || len(resp.GetResult().GetPedestrian()) != 1 {
			return
		}

		// var confidence float32 = 1

		recPedestrian := resp.GetResult().GetPedestrian()[0]

		//check face
		if recPedestrian.GetFace() == nil {
			glog.Infoln("recPedestrian.GetFace()==nil,image uri:")
			return
		}

		//质量检测
		file, err := os.OpenFile("result.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			glog.Errorln(err)
			return
		}
		defer file.Close()

		face := recPedestrian.GetFace()
		if face.GetImg() == nil || face.GetImg().GetCutboard() == nil || face.GetPose() == nil ||
			face.GetPose().Angles == nil || len(face.GetPose().Angles) < 2 {

			glog.Errorln(img.Filename, ": NULL Img/Cutboard/Pose/Angle")
			return
		}
		if face.AlignScore < this.quality.AlignScore || face.BlurScore < this.quality.BlurScore ||
			math.Abs(float64(face.GetPose().Angles[0])) > 15 || math.Abs(float64(face.GetPose().Angles[1])) > 15 ||
			face.GetImg().GetCutboard().Height < 60 || face.GetImg().GetCutboard().Width < 60 {

			rst := fmt.Sprintln("Fail ", img.Filename, face.AlignScore, face.BlurScore, face.GetPose().Angles[0],
				face.GetPose().Angles[1], face.GetImg().GetCutboard().Height, face.GetImg().GetCutboard().Width)

			fmt.Print(rst)
			io.WriteString(file, rst)
		} else {
			// os.Rename(img.Filename, path.Dir(img.Filename)+"/Pass"+path.Base(img.Filename))
			rst := fmt.Sprintln("Pass ", img.Filename, face.AlignScore, face.BlurScore, face.GetPose().Angles[0],
				face.GetPose().Angles[1], face.GetImg().GetCutboard().Height, face.GetImg().GetCutboard().Width)

			fmt.Print(rst)
			io.WriteString(file, rst)
		}
		return
	}(img)
	return nil
}

const (
	blurMax float64 = 30.0
	blurMin float64 = 0

	angleMax float64 = 30
	angleMin float64 = 0

	heightMax float64 = 150
	heightMin float64 = 20
)

func (this *MatrixTransform) GetFaceQualityScore(face *dg_model.RecFace) float32 {
	// fmt.Println("alignScore:", face.GetAlignScore())
	// fmt.Println("blurScore:", face.GetBlurScore())
	// fmt.Println("angle[0]:", face.GetPose().GetAngles()[0])
	// fmt.Println("angle[1]:", face.GetPose().GetAngles()[1])
	// fmt.Println("height:", face.GetImg().GetCutboard().GetHeight())
	var blurWeight float64 = 0
	var angleWeight float64 = 1.0
	var heightWeight float64 = 1.0

	var maxWeight float64 = blurWeight + angleWeight + heightWeight
	if face.AlignScore < 0.2 {
		angleWeight *= (float64)(face.AlignScore / 0.2)
	}

	if face.GetImg() == nil ||
		face.GetImg().GetCutboard() == nil ||
		face.GetPose() == nil ||
		face.GetPose().Angles == nil ||
		len(face.GetPose().Angles) < 2 {
		return -1
	}

	if !IsFace(face.GetImg().GetCutboard().Confidence, face.AlignScore, this.config.FaceDetectThreshold) {
		return -1
	}

	// float h = f.img().cutboard().height();
	h := float64(face.GetImg().GetCutboard().Height)

	if math.Abs(float64(face.BlurScore)) > blurMax ||
		math.Abs(float64(face.GetPose().Angles[0])) > angleMax ||
		math.Abs(float64(face.GetPose().Angles[1])) > angleMax ||
		h < heightMin {
		return 0
	}

	if h > heightMax {
		h = heightMax
	}
	h = (h - heightMin) / (heightMax - heightMin)

	confidence := ((1-float64(face.BlurScore)/blurMax)*blurWeight +
		(1-math.Abs(float64(face.GetPose().Angles[0]))/angleMax)*0.5*angleWeight +
		(1-math.Abs(float64(face.GetPose().Angles[1]))/angleMax)*0.5*angleWeight +
		h*heightWeight) / maxWeight

	return float32(confidence)
}

func IsFace(coutboardConfidence float32, alignScore float32, faceDetectThreshold float32) bool {

	if alignScore >= 0.5 {
		return true
	}

	if coutboardConfidence < 0.7 {
		return false
	}
	fuseScore := coutboardConfidence + alignScore*5
	if fuseScore < 1.4 && alignScore < 0.0001 {
		return false
	}
	if fuseScore > faceDetectThreshold {
		return true
	}
	return false

}
