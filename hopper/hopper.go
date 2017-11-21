package main

import (
	"flag"

	"github.com/deepglint/tool/hopper/center"
	"github.com/deepglint/tool/hopper/config"

	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/golang/glog"
)

func main() {

	flag.Parse()
	data, err := ioutil.ReadFile("config/config.json")

	if err != nil {
		glog.Errorln("read config err:", err)
		return
	}

	centerConfig := new(config.Config)

	err = json.Unmarshal(data, &centerConfig)
	if err != nil {
		glog.Errorln("parse config err:", err)
		return
	}
	glog.Infoln("%v", centerConfig)

	center, err := center.NewCenter(centerConfig)
	if err != nil {
		glog.Errorln("parse config err:", err)
		// fmt.Println(err)
		return
	}

	if centerConfig.HopperOut {
		go func() {
			for {
				err := center.StartHopperOut()
				if err != nil {
					glog.Errorln("start hopper out err:", err, "restart center")
					center.CloseHopperOut()
					time.Sleep(time.Duration(1) * time.Second)
					continue
				}
				break
			}
		}()
	}
	if centerConfig.HopperIn {
		go func() {
			for {
				glog.Infoln("start")
				err := center.StartHopperIn()
				if err != nil {
					glog.Errorln("start hopper in err:", err, "restart center")
					center.CloseHopperIn()
					time.Sleep(time.Duration(1) * time.Second)
					continue
				}
				break
			}
		}()
	}

	wait := make(chan int)
	<-wait
}
