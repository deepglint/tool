package input

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/deepglint/tool/cap_pic_filter/configs"
	"github.com/golang/glog"
)

type FileInput struct {
	path          string
	stopChan      chan bool
	resultChan    chan *configs.Img
	sendMessageWg sync.WaitGroup
}

func NewFileInput(path string) *FileInput {
	return &FileInput{
		path:       path,
		stopChan:   make(chan bool),
		resultChan: make(chan *configs.Img, 1),
	}
}

func (this *FileInput) Start() {
	defer close(this.resultChan)
	// for {
	select {
	case <-this.stopChan:
		this.sendMessageWg.Wait()
		close(this.resultChan)
		glog.Errorln("Close file input")
		return
	default:
		_ = filepath.Walk(this.path, func(datapath string, f os.FileInfo, err error) error {
			glog.Infoln(datapath)
			if f.IsDir() {
				return nil
			}
			file, file_err := os.OpenFile(datapath, os.O_RDONLY, 0644)
			if file_err != nil {
				glog.Errorln(file_err)
				return file_err
			}
			data, read_err := ioutil.ReadAll(file)
			if read_err != nil {
				glog.Errorln(read_err)
				return read_err
			}
			img := &configs.Img{}
			if len(data) > 4 && string(data[0:4]) == "http" {
				img.Uri = string(data)
			} else {
				img.Bindata = string(data)
			}
			img.Filename = datapath
			this.resultChan <- img
			return nil
		})
	}
	// }
}

func (this *FileInput) Stop() {
	close(this.stopChan)
}

func (this *FileInput) GetResultChan() chan *configs.Img {
	return this.resultChan
}
