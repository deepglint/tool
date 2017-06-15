package dg_model

import (
	"time"
)

func (q *QueryTimeSpacialRange) IsValid() bool {
	if len(q.GetTimeRanges()) == 0 || len(q.SensorIds) == 0 {
		return false
	}

	// check start/end time
	for i := 0; i < len(q.GetTimeRanges()); i++ {
		if q.GetTimeRanges()[i].TimestampStart > q.GetTimeRanges()[i].TimestampEnd {
			q.GetTimeRanges()[i].TimestampStart, q.GetTimeRanges()[i].TimestampEnd = q.GetTimeRanges()[i].TimestampEnd, q.GetTimeRanges()[i].TimestampStart
		}
	}

	return true
}

func (q *QueryTimeSpacialRange) MaxTimestamp() int64 {
	max := time.Now().Unix() * 1000
	for _, timeRange := range q.GetTimeRanges() {
		if timeRange.TimestampEnd > max {
			max = timeRange.TimestampEnd
		}
	}
	return max
}

func (v *VehicleQuery) IsValid() bool {
	if v == nil || v.GetBaseQuery() == nil || v.GetBaseQuery().QueryId == "" || len(v.GetBaseQuery().GetQueryTimeSpacialRanges()) == 0 {
		return false
	}
	return true
}

func (v *BicycleQuery) IsValid() bool {
	if v == nil || v.GetBaseQuery() == nil || v.GetBaseQuery().QueryId == "" || len(v.GetBaseQuery().GetQueryTimeSpacialRanges()) == 0 {
		return false
	}
	return true
}

func (v *TricycleQuery) IsValid() bool {
	if v == nil || v.GetBaseQuery() == nil || v.GetBaseQuery().QueryId == "" || len(v.GetBaseQuery().GetQueryTimeSpacialRanges()) == 0 {
		return false
	}
	return true
}

func (p *PedestrianQuery) IsValid() bool {
	if p == nil || p.GetBaseQuery() == nil || p.GetBaseQuery().QueryId == "" || len(p.GetBaseQuery().GetQueryTimeSpacialRanges()) == 0 {
		return false
	}
	return true
}

func (this *AlarmRule) IsValid() bool {
	if this.GetAlarmContent() == nil || !this.GetAlarmContent().GetQueryTimeSpacialRange().IsValid() {
		return false
	}

	return true
}

func (this *SensorInfo) IsValid() bool {
	if this.RepoId == 0 || (this.SensorId == 0 && this.Name == "") {
		return false
	}

	return true
}
