package center

import (
	"errors"
	"sync"

	"github.com/deepglint/tool/cap_pic_filter/configs"
	"github.com/deepglint/tool/cap_pic_filter/input"
	"github.com/deepglint/tool/cap_pic_filter/transform"
	"github.com/golang/glog"
)

type FilterCenter struct {
	matrixTransform *transform.MatrixTransform
	fileinput       *input.FileInput
	islive          bool
	mutex           sync.Mutex
}

func NewFilterCenter(config *configs.FilterCenterConfig) (*FilterCenter, error) {

	center := &FilterCenter{
		fileinput: input.NewFileInput(config.Path),
	}
	mt, err := transform.NewMatrixTransform(config)
	if err != nil {
		glog.Errorln(err)
		return nil, err
	}
	center.matrixTransform = mt
	return center, nil

}

func (this *FilterCenter) Start() (err error) {

	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.islive {
		return errors.New("filter server is open")
	}
	glog.Errorln("filter server start")
	this.islive = true

	go this.fileinput.Start()
	resultChan := this.fileinput.GetResultChan()

	this.matrixTransform.StartTransform(resultChan)

	return nil
}

func (this *FilterCenter) Close() (err error) {

	this.mutex.Lock()
	defer this.mutex.Unlock()
	if !this.islive {
		return errors.New("isd server has stoped")
	}

	return nil
}
