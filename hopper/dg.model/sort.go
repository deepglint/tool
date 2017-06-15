package dg_model

const (
	SORTBY_TIMESTAMP        = "ts"
	SORTBY_CONFIDENCE       = "confidence"
	SORTBY_COLOR            = "color_id"
	SORTBY_COLOR_CONFIDENCE = "color_confidence"
	SORTBY_BRAND            = "brand_id"
	SORTBY_BRAND_CONFIDENCE = "model_type_confidence"
)

type TimestampSortedRetVehicles []*RetVehicle
type ColorSortedRetVehicles []*RetVehicle
type BrandSortedRetVehicles []*RetVehicle
type TimestampSortedPlates []*VehicleTimeSpacial

func (p TimestampSortedRetVehicles) Len() int {
	return len(p)
}

func (p TimestampSortedRetVehicles) Less(i, j int) bool {
	return p[i].Metadata.Timestamp > p[j].Metadata.Timestamp
}

func (p TimestampSortedRetVehicles) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//
func (p ColorSortedRetVehicles) Len() int {
	return len(p)
}

func (p ColorSortedRetVehicles) Less(i, j int) bool {
	return p[i].RecVehicle.Color.ColorId < p[j].RecVehicle.Color.ColorId
}

func (p ColorSortedRetVehicles) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//
func (p BrandSortedRetVehicles) Len() int {
	return len(p)
}

func (p BrandSortedRetVehicles) Less(i, j int) bool {
	return p[i].RecVehicle.ModelType.BrandId < p[j].RecVehicle.ModelType.BrandId
}

func (p BrandSortedRetVehicles) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//
func (v TimestampSortedPlates) Len() int {
	return len(v)
}

func (v TimestampSortedPlates) Less(i, j int) bool {
	return v[i].Timestamp < v[j].Timestamp
}

func (v TimestampSortedPlates) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

/*---------------*/
type TimestampSortedRegisteredCivil []*RegisteredFace

func (this TimestampSortedRegisteredCivil) Len() int {
	return len(this)
}

func (this TimestampSortedRegisteredCivil) Less(i, j int) bool {
	return this[i].Timestamp < this[j].Timestamp
}

func (this TimestampSortedRegisteredCivil) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

/*---------------*/
type ConfidenceSortedRegisteredCivil []*RegisteredFace

func (this ConfidenceSortedRegisteredCivil) Len() int {
	return len(this)
}

func (this ConfidenceSortedRegisteredCivil) Less(i, j int) bool {
	return this[i].Confidence < this[j].Confidence
}

func (this ConfidenceSortedRegisteredCivil) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

/*---------------*/
type TimestampSortedCapturedFace []*CapturedFace

func (this TimestampSortedCapturedFace) Len() int {
	return len(this)
}

func (this TimestampSortedCapturedFace) Less(i, j int) bool {
	return this[i].Timestamp < this[j].Timestamp
}

func (this TimestampSortedCapturedFace) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

/*---------------*/
type ConfidenceSortedCapturedFace []*CapturedFace

func (this ConfidenceSortedCapturedFace) Len() int {
	return len(this)
}

func (this ConfidenceSortedCapturedFace) Less(i, j int) bool {
	return this[i].Confidence < this[j].Confidence
}

func (this ConfidenceSortedCapturedFace) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
