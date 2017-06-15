package dg_model

type ISDFilter struct {
	//sensor sn-id
	SensorID string
	//serial id
	SerialID string
	//whether sensor is alive
	IsAlive bool
	//sensor's filter  threshold
	FilterThreshold float32
}
