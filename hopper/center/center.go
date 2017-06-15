package center

import (
	"github.com/deepglint/hopper/config"
	"github.com/deepglint/hopper/input"
	"github.com/deepglint/hopper/output"

	"errors"
	"sync"
	"time"

	"github.com/golang/glog"
)

type Center struct {
	kafkaInput *input.KafkaInput
	ftpoutput  *output.FTPOutput
	ftpinput   *input.FTPInput
	islive     bool
	mutex      sync.Mutex
}

func NewCenter(centerConfig *config.Config) (*Center, error) {

	center := new(Center)

	if centerConfig.HopperOut {
		// init kafkaInput
		kafkaInput, err := input.NewKafkaInput(centerConfig.KafkaInputConfig)
		if err != nil {
			return nil, err
		}

		// init ftpOutput
		ftpOutput, err := output.NewFTPOutput(centerConfig.FTPOutputConfig)
		if err != nil {
			return nil, err
		}

		center.kafkaInput = kafkaInput
		center.ftpoutput = ftpOutput

	}

	if centerConfig.HopperIn {
		ftpinput, err := input.NewFTPInput(centerConfig)
		if err != nil {
			return nil, err
		}
		err = ftpinput.Setup(centerConfig)
		if err != nil {
			return nil, err
		}
		center.ftpinput = ftpinput
	}
	//init center

	return center, nil
}

func (this *Center) StartHopperOut() error {

	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.islive {
		return errors.New("server is open")
	}
	glog.Infof("server start")
	this.islive = true

	err := this.kafkaInput.Start()
	if err != nil {
		return err
	}

	err = this.ftpoutput.Start(this.kafkaInput.GetMessageInfoChan())
	if err != nil {
		return err
	}
	return nil
}

func (this *Center) CloseHopperOut() error {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if !this.islive {
		return errors.New("server has stoped")
	}

	this.kafkaInput.Close()
	for {
		if this.kafkaInput.GetLiveState() || this.ftpoutput.GetLiveState() {
			glog.Info("wait server stoped")
			<-time.After(time.Second)
			continue
		}
		break
	}

	this.islive = false
	return nil
}

func (this *Center) StartHopperIn() error {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	err := this.ftpinput.Start()
	if err != nil {
		return err
	}
	return nil
}

func (this *Center) CloseHopperIn() error {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	err := this.ftpinput.Clean()
	if err != nil {
		return err
	}
	return nil
}
