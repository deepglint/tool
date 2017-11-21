package output

import (
	"bytes"
	"sync"

	"github.com/deepglint/tool/hopper/config"
	"github.com/deepglint/tool/hopper/utils"
	"github.com/golang/glog"
	"github.com/jlaffaye/ftp"
)

type FTPOutput struct {
	config        *config.FTPConfig
	isLive        bool
	ftpThreadPool chan int
	sendMessageWg sync.WaitGroup
	ftpConnect    *ftp.ServerConn
}

func NewFTPOutput(ftpOutputConfig *config.FTPConfig) (*FTPOutput, error) {
	ftpOutput := new(FTPOutput)
	ftpOutput.config = ftpOutputConfig
	return ftpOutput, nil
}

func (this *FTPOutput) Start(inputChan chan *[]byte) error {

	err := this.reConnectFTP()
	if err != nil {
		return err
	}

	this.ftpThreadPool = make(chan int, this.config.MaxThread)
	for threadCount := 0; threadCount < this.config.MaxThread; threadCount++ {
		this.ftpThreadPool <- 1
	}
	this.isLive = true
	glog.Info("[OUT_FTP] ftp output Start")

	go func(inputChan chan *[]byte) {
		defer func() {
			if this.ftpConnect != nil {
				this.ftpConnect.Logout()
				this.ftpConnect.Quit()
			}
			this.isLive = false
			glog.Info("ftp output closed")
		}()

		err := this.ftpConnect.MakeDir(this.config.Path)
		if err != nil {
			glog.Errorln(err)
		}

		for data := range inputChan {
			<-this.ftpThreadPool
			this.sendMessageWg.Add(1)
			go func(data *[]byte) {
				defer func() {
					this.ftpThreadPool <- 1
					this.sendMessageWg.Done()
				}()
				messageTmpPath := this.config.Path + utils.NewV1().String() + ".tmp"
				err := this.ftpConnect.Stor(messageTmpPath, bytes.NewBuffer(*data))
				if err != nil {
					glog.Warning("write data to ftp err:", err)
					glog.Warning("start reconnect ftp")
					err = this.reConnectFTP()
					if err != nil {
						glog.Warning("reconnect ftp err:", err)
						return
					}
					glog.Warning("reconnect ftp success")
					return
				}

				messagePath := this.config.Path + utils.NewV1().String()
				err = this.ftpConnect.Rename(messageTmpPath, messagePath)
				if err != nil {
					glog.Warning("rename data in ftp err:", err)
					glog.Warning("start reconnect ftp")
					err = this.reConnectFTP()
					if err != nil {
						glog.Warning("reconnect ftp err:", err)
						return
					}
					glog.Warning("reconnect ftp success")
					return
				}

			}(data)

		}
		this.sendMessageWg.Wait()
	}(inputChan)
	return nil
}

func (this *FTPOutput) GetLiveState() bool {
	return this.isLive
}

//re-connect ftp
func (this *FTPOutput) reConnectFTP() error {

	if this.ftpConnect != nil {
		this.ftpConnect.Logout()
		this.ftpConnect.Quit()
	}

	var err error
	//connect ftp
	this.ftpConnect, err = ftp.Connect(this.config.Address)
	if err != nil {
		glog.Warningln(err)
		return err
	}

	//login in
	err = this.ftpConnect.Login(this.config.Username, this.config.Password)
	if err != nil {
		glog.Warningln(err)
		return err
	}
	return nil
}
