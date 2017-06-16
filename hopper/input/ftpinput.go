package input

import (
	"bytes"
	"context"
	"errors"
	"io"
	// "io/ioutil"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/deepglint/tool/hopper/config"
	models "github.com/deepglint/tool/hopper/dg.model"
	"github.com/deepglint/tool/hopper/output"
	"github.com/deepglint/tool/hopper/utils"
	"github.com/ftp"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

type FTPInput struct {
	ip          string
	username    string
	password    string
	walkingpath string
	speedlimit  int
	conn        *ftp.ServerConn
	output      *output.KafkaWriter
	isLive      bool
	devices     map[string]string
}

type WalkFunc func(datapath string, info os.FileMode, err error) error

func NewFTPInput(config *config.Config) (*FTPInput, error) {
	return &FTPInput{
		ip:          config.FTPInputConfig.Address,
		username:    config.FTPInputConfig.Username,
		password:    config.FTPInputConfig.Password,
		walkingpath: config.FTPInputConfig.Path,
		speedlimit:  config.SpeedLimit,
		isLive:      false,
		devices:     config.Devices,
	}, nil
}

func (this *FTPInput) Setup(config *config.Config) error {

	var err error
	if o, err := output.NewKafkaWriter(config); err == nil {
		this.output = o
	} else {
		glog.Errorln(err)
		return err
	}
	if err = this.output.Setup(); err != nil {
		glog.Errorln(err)
		return err
	}

	for i := 0; i < 3; i++ {
		if this.conn, err = ftp.DialTimeout(this.ip, time.Second*time.Duration(3)); err != nil {
			continue
		} else {
			break
		}
	}
	if err != nil {
		glog.Errorln(err)
		return err
	} else {
		if err = this.conn.Login(this.username, this.password); err != nil {
			glog.Errorln(err)
			return err
		} else {
		}
	}
	return nil
}

func (this *FTPInput) Clean() error {
	defer func() {
		this.output = nil
		this.conn = nil
	}()

	e1 := this.output.Clean()
	e2 := this.conn.Quit()

	switch {
	case e1 != nil:
		return e1
	case e2 != nil:
		return e2
	default:
		return nil
	}
}

func (this *FTPInput) walk(datapath string, cnt int, walkFn WalkFunc) (error, int) {
	ss, _ := utils.GbkToUtf8([]byte(datapath))
	glog.Errorln("walking", string(ss))

	var entries []*ftp.Entry
	var err error
	if entries, err = this.conn.List(datapath); err != nil {
		return err, 0
	}
	for _, entry := range entries {
		switch entry.Type {
		case ftp.EntryTypeFolder:
			if entry.Name == "." {
			} else if entry.Name == ".." || entry.Name == "上级目录" {
			} else if entry.Name == "processed" {
			} else {
				if err, count := this.walk(datapath+"/"+entry.Name, cnt, walkFn); err != nil {
					glog.Errorln(err)
					return err, 0
				} else {
					cnt = count
				}
			}
		case ftp.EntryTypeFile:
			cnt++
			if err = walkFn(datapath+"/"+entry.Name, os.FileMode(0), nil); err != nil {
				if err.Error() != "filename format error" {
					return err, cnt
				} else {
					glog.Errorln(err)
				}
			}
			glog.Errorln(cnt, this.speedlimit, time.Now())
			if cnt >= this.speedlimit {
				return errors.New("reach speedlimit"), 0
			}
		}
	}
	return err, cnt
}

func (this *FTPInput) Start() error {
	defer this.output.Perform()
	cnt := 0

	for _ = range time.NewTicker(time.Second).C {
		err, count := this.walk(this.walkingpath, cnt, func(datapath string, f os.FileMode, err error) error {
			glog.Errorln(datapath)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			select {
			case <-ctx.Done():
				glog.Errorln("Get image timeout")
				return nil
			default:
				image, err := this.conn.Retr(datapath)
				if err != nil {
					glog.Errorln(err)
					return nil
				}
				buf := new(bytes.Buffer)
				_, err = io.Copy(buf, image)
				if err != nil {
					glog.Errorln(err)
					return nil
				}
				defer image.Close()
				// ioutil.WriteFile(filename, buf.Bytes(), 0644)

				var obj models.GenericObj
				err = proto.Unmarshal(buf.Bytes(), &obj)
				if err != nil {
					glog.Warningf("Unmarshal models.GenericObj from protobuffer error: %s\n", err.Error())
					return nil
				}
				pedestrians := &models.PedestrianObj{}
				if obj.FmtType == models.DataFmtType_JSON {
					err = json.Unmarshal(obj.BinData, pedestrians)
					if err != nil {
						glog.Warningf("Unmarshal models.PedestrianObj from json error: %s\n", err.Error())
						return nil
					}
				} else if obj.FmtType == models.DataFmtType_PROTOBUF {
					err = proto.Unmarshal(obj.BinData, pedestrians)
					if err != nil {
						glog.Warningf("Unmarshal models.PedestrianObj from protobuffer error: %s\n", err.Error())
						return nil
					}
				} else {
					glog.Warningf("Can not recognize serialization type, FmtType=%d\n", obj.FmtType)
					return nil
				}

				old_sensorid := pedestrians.GetMetadata().SensorIdStr
				filename := filepath.Base(filepath.Dir(datapath))
				// glog.Errorf("filename: %s, pedestrian: %v", filename, pedestrians)
				if id, ok := this.devices[filename]; ok {
					pedestrians.GetMetadata().SensorIdStr = id
					glog.Errorln(old_sensorid, id)
				} else {
					return nil
				}
				var body []byte
				if obj.FmtType == models.DataFmtType_JSON {
					body, err = json.Marshal(pedestrians)
					if err != nil {
						glog.Warningf("Marshal models.PedestrianObj to json error: %s\n", err.Error())
						return nil
					}
				} else if obj.FmtType == models.DataFmtType_PROTOBUF {
					body, err = proto.Marshal(pedestrians)
					if err != nil {
						glog.Warningf("Marshal models.PedestrianObj to protobuffer error: %s\n", err.Error())
						return nil
					}
				} else {
					glog.Warningf("Can not recognize serialization type, FmtType=%d\n", obj.FmtType)
					return nil
				}
				obj.BinData = body
				info, err := proto.Marshal(&obj)
				if err != nil {
					glog.Warningf("Marshal models.GenericObj to protobuffer error: %s\n", err.Error())
					return nil
				}

				if err := this.output.Write(info, 100); err != nil {
					glog.Errorln(err)
					return nil //omit write error
				}

				err = this.conn.Delete(datapath)
				if err != nil {
					glog.Errorln(err)
				}
			}
			return nil
		})
		cnt = count
		if err != nil && err.Error() == "reach speedlimit" {
			this.output.Perform()
			continue
		}
	}

	return nil
}
