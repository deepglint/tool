package dg_model

import (
	"time"
)

func (this *FaceRule) Valid() bool {
	if this.GetSensors() == nil || this.GetTimes() == nil || len(this.GetSensors()) == 0 {
		return false
	}
	if this.Timestamp == 0 {
		this.Timestamp = time.Now().UnixNano() / 1000000
	}
	return true
}
