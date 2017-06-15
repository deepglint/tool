package configs

import "github.com/deepglint/filter/dg.model"

type FilterCenterConfig struct {
	TransformConfig *MatrixTransformConfig
	Quality         *QualityScore
	Path            string
}

type QualityScore struct {
	AlignScore float32
	FaceAngle  int
	BlurScore  float32
	MinHeights int
	MinWidth   int
}

type MatrixTransformConfig struct {
	RecogniseURL        string
	MaxThread           int
	TimeOut             int
	Etcd                EtcdConfig
	GetFeatureContext   dg_model.WitnessRequestContext
	GetQualityContext   dg_model.WitnessRequestContext
	QualityThreshold    float32
	IsQulityFilterOpen  bool
	IsEdgeDetectionOpen bool
	FaceDetectThreshold float32
	Weed                WeedConfig
}

type EtcdConfig struct {
	Addr      []string
	Interval  uint16
	FilterDir string
	Timeout   int
}

type WeedConfig struct {
	Host        string
	Replication string
	Collection  string
	TTLStr      string
}

type Img struct {
	Filename string
	Uri      string
	Bindata  string
}
