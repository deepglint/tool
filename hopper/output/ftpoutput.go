package output

import (
	"bytes"
	"sync"

	"github.com/deepglint/tool/hopper/config"
	"github.com/deepglint/tool/hopper/utils"

	"github.com/deepglint/glog"
	"github.com/ftp"
)

type FTPOutput struct {
	config        *config.FTPConfig
	isLive        bool
	ftpThreadPool chan int
	sendMessageWg sync.WaitGroup
}

func NewFTPOutput(ftpOutputConfig *config.FTPConfig) (*FTPOutput, error) {
	ftpOutput := new(FTPOutput)
	ftpOutput.config = ftpOutputConfig
	return ftpOutput, nil
}

func (this *FTPOutput) Start(inputChan chan *[]byte) error {

	ftpCon, err := ftp.Dial(this.config.Address)
	if err != nil {
		return err
	}

	err = ftpCon.Login(this.config.Username, this.config.Password)
	if err != nil {
		return err
	}
	this.ftpThreadPool = make(chan int, this.config.MaxThread)
	for threadCount := 0; threadCount < this.config.MaxThread; threadCount++ {
		this.ftpThreadPool <- 1
	}
	this.isLive = true
	glog.Info("[OUT_FTP] ftp output Start")

	go func(ftpServerConn *ftp.ServerConn, inputChan chan *[]byte) {
		defer func() {
			this.isLive = false
			glog.Info("ftp output closed")
		}()
		for data := range inputChan {
			<-this.ftpThreadPool
			this.sendMessageWg.Add(1)
			go func(data *[]byte) {
				defer func() {
					this.ftpThreadPool <- 1
					this.sendMessageWg.Done()
				}()
				messagePath := this.config.Path + utils.NewV1().String()
				err := ftpServerConn.Stor(messagePath, bytes.NewBuffer(*data))
				if err != nil {
					glog.Warning("write data to ftp err:", err)
				}
			}(data)

		}
		this.sendMessageWg.Wait()
	}(ftpCon, inputChan)
	return nil
}

func (this *FTPOutput) GetLiveState() bool {
	return this.isLive
}
