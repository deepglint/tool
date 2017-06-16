package main

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/deepglint/tool/cap_pic_filter/center"
	"github.com/deepglint/tool/cap_pic_filter/configs"
	"github.com/golang/glog"
)

func main() {

	config := &configs.FilterCenterConfig{}
	data, err := ioutil.ReadFile("configs/config.json")
	if err != nil {
		glog.Errorln("read filter config err:", err)
		return
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		glog.Errorln("parse filter config err:", err)
		return
	}

	glog.Errorf("conf: %v", config)

	center, err := center.NewFilterCenter(config)
	if err != nil {
		glog.Errorln("parse filter config err:", err)
		// fmt.Println(err)
		return
	}

	for {
		err := center.Start()
		if err != nil {
			glog.Errorln("start filter err:", err, "restart filter")
			center.Close()
			time.Sleep(time.Duration(1) * time.Second)
			continue
		}
		break
	}

	wait := make(chan int)
	<-wait
}
