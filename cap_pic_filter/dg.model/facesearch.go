package dg_model

// A temporary solution
import (
	"bytes"
	"strconv"
	"strings"
)

func (this *BaseConds) ParseMetaSearch() (*MetaSearch, *MetaCount) {
	condition := &LogicalCondition{
		Conditions: []*Condition{
			&Condition{Key: "status", Value: TaskStatus_Task_Status_Deleted, Kind: OP_NOT_EQUAL},
		},
		Kind: LOGICAL_AND,
	}
	metaCount := &MetaCount{
		Conditions: condition,
	}
	metaSearch := &MetaSearch{
		Columns:    ParseColumns(this.Columns),
		Conditions: condition,
		Group:      ParseGroupSet(this.GroupByDimensions),
		Order:      ParseOrderSet(this.SortBy, this.SortOrderAsc),
		Page:       ParsePagingSet(this.Offset, this.Limit),
	}
	return metaSearch, metaCount
}

func (this *FaceRuleQuery) Valid() bool {
	if this.GetBaseConds() == nil {
		return false
	}
	return true
}

func (this *FaceRuleQuery) ParseMetaSearch() (*MetaSearch, *MetaCount) {
	metaSearch := new(MetaSearch)
	metaCount := new(MetaCount)
	if this.GetBaseConds() != nil {
		metaSearch.Columns = ParseColumns(this.GetBaseConds().Columns)
		metaSearch.Group = ParseGroupSet(this.GetBaseConds().GroupByDimensions)
		metaSearch.Order = ParseOrderSet(this.GetBaseConds().SortBy, this.GetBaseConds().SortOrderAsc)
		metaSearch.Page = ParsePagingSet(this.GetBaseConds().Offset, this.GetBaseConds().Limit)
	}
	// NOTE: use rule_sensor_id instead of sensor_id, in case of 'one sensor multi-rules' condition
	partition := &PartitionSet{Key: "sensor_id", Count: 1}
	metaSearch.Partition = partition
	metaCount.Group = &GroupSet{Key: "sensor_id"}

	// default not deleted rule
	conds := make([]*Condition, 0)
	conds = append(conds, &Condition{Key: "status", Value: TaskStatus_Task_Status_Deleted, Kind: OP_NOT_EQUAL})
	if this.GetAuxiliaryConds() != nil {
		conds = append(conds, this.GetAuxiliaryConds().ParseConditions()...)
	}
	if this.SensorName != "" {
		conds = append(conds, &Condition{Key: "sensor_name", Value: fuzzy(this.SensorName), Kind: OP_LIKE})
	}

	logicalCondition := &LogicalCondition{Conditions: conds, Kind: LOGICAL_AND}
	metaSearch.Conditions = logicalCondition
	metaCount.Conditions = logicalCondition

	return metaSearch, metaCount
}

/*
	Columns    []*Column
	Conditions *LogicalCondition
	Group      *GroupSet
	Order      *OrderSet
	Have       *HavingSet
	Page       *PagingSet
*/
func (this *FaceEventQuery) Valid() bool {
	if this.GetBaseConds() == nil || this.GetAttrConds() == nil || !this.GetAttrConds().Valid() {
		return false
	}
	return true
}

func (this *FaceEventQuery) ParseMetaSearch() (*MetaSearch, *MetaCount) {
	metaSearch := new(MetaSearch)
	metaCount := new(MetaCount)
	if this.GetBaseConds() != nil {
		metaSearch.Columns = ParseColumns(this.GetBaseConds().Columns)
		metaSearch.Group = ParseGroupSet(this.GetBaseConds().GroupByDimensions)
		metaSearch.Order = ParseOrderSet(this.GetBaseConds().SortBy, this.GetBaseConds().SortOrderAsc)
		metaSearch.Page = ParsePagingSet(this.GetBaseConds().Offset, this.GetBaseConds().Limit)
	}
	partition := &PartitionSet{Key: "event_reid", Count: 1}
	metaSearch.Partition = partition
	metaCount.Group = &GroupSet{Key: "event_reid"}

	conds := make([]*Condition, 0)
	conds = append(conds, &Condition{Key: "status", Value: TaskStatus_Task_Status_Deleted, Kind: OP_NOT_EQUAL})
	if this.GetAuxiliaryConds() != nil {
		conds = append(conds, this.GetAuxiliaryConds().ParseConditions()...)
	}
	logicalCondition := this.GetAttrConds().ParseConditions()
	if logicalCondition.Conditions != nil {
		logicalCondition.Conditions = append(logicalCondition.Conditions, conds...)
	} else {
		logicalCondition.Conditions = conds
	}

	metaSearch.Conditions = logicalCondition
	metaCount.Conditions = logicalCondition

	return metaSearch, metaCount
}

func (this *CapturedFaceQuery) Valid() bool {
	if this.GetBaseConds() == nil || this.GetAttrConds() == nil || !this.GetAttrConds().Valid() {
		return false
	}
	return true
}

func (this *CapturedFaceQuery) ParseMetaSearch() (*MetaSearch, *MetaCount) {
	metaSearch := &MetaSearch{}
	metaCount := &MetaCount{}
	if this.GetBaseConds() != nil {
		metaSearch.Columns = ParseColumns(this.GetBaseConds().Columns)
		metaSearch.Group = ParseGroupSet(this.GetBaseConds().GroupByDimensions)
		metaSearch.Order = ParseOrderSet(this.GetBaseConds().SortBy, this.GetBaseConds().SortOrderAsc)
		metaSearch.Page = ParsePagingSet(this.GetBaseConds().Offset, this.GetBaseConds().Limit)
	}

	conds := make([]*Condition, 0)
	conds = append(conds, &Condition{Key: "status", Value: TaskStatus_Task_Status_Deleted, Kind: OP_NOT_EQUAL})
	if this.GetAuxiliaryConds() != nil {
		conds = append(conds, this.GetAuxiliaryConds().ParseConditions()...)
	}
	logicalCondition := this.GetAttrConds().ParseConditions()
	if logicalCondition.Conditions != nil {
		logicalCondition.Conditions = append(logicalCondition.Conditions, conds...)
	} else {
		logicalCondition.Conditions = conds
	}

	metaSearch.Conditions = logicalCondition
	metaCount.Conditions = logicalCondition

	return metaSearch, metaCount
}

// --------------------------
func (this *CapturedConds) Valid() bool {
	if this.GetTimeSpacialRangeConds() == nil || !this.GetTimeSpacialRangeConds().Valid() {
		return false
	}
	return true
}

func (this *CapturedConds) ParseConditions() *LogicalCondition {
	logicalCondtion := this.GetTimeSpacialRangeConds().ParseLogicalCondition()
	if this.GetConditions() != nil {
		conds := parseConditions(this.GetConditions())
		logicalCondtion.Conditions = append(logicalCondtion.Conditions, conds...)
	}
	if len(this.Ids) > 0 {
		logicalCondtion.Conditions = append(logicalCondtion.Conditions, &Condition{Key: "face_id", Value: this.Ids, Kind: OP_IN})
	}
	return logicalCondtion
}

// --------------------------
func (this *RegisteredFaceQuery) Valid() bool {
	if this.GetBaseConds() == nil || this.GetAttrConds() == nil || !this.GetAttrConds().Valid() {
		return false
	}
	// TODO:
	return true
}

func (this *RegisteredFaceQuery) ParseMetaSearch(partition bool) (*MetaSearch, *MetaCount) {
	metaSearch := &MetaSearch{}
	metaCount := &MetaCount{}
	if this.GetBaseConds() != nil {
		metaSearch.Columns = ParseColumns(this.GetBaseConds().Columns)
		metaSearch.Group = ParseGroupSet(this.GetBaseConds().GroupByDimensions)
		metaSearch.Order = ParseOrderSet(this.GetBaseConds().SortBy, this.GetBaseConds().SortOrderAsc)
		metaSearch.Page = ParsePagingSet(this.GetBaseConds().Offset, this.GetBaseConds().Limit)
	}
	if partition {
		partition := &PartitionSet{Key: "civil_attr_id", Count: 1}
		metaSearch.Partition = partition
		metaCount.Group = &GroupSet{Key: "civil_attr_id"}
	}

	conds := make([]*Condition, 0)
	conds = append(conds, &Condition{Key: "status", Value: TaskStatus_Task_Status_Deleted, Kind: OP_NOT_EQUAL})
	if this.GetAuxiliaryConds() != nil {
		conds = append(conds, this.GetAuxiliaryConds().ParseConditions()...)
	}
	if this.GetAttrConds() != nil {
		conds = append(conds, this.GetAttrConds().ParseConditions()...)
	}
	conditions := &LogicalCondition{Conditions: conds, Kind: LOGICAL_AND}
	metaSearch.Conditions = conditions
	metaCount.Conditions = conditions

	return metaSearch, metaCount
}

// --------------------------
func (this *RegisteredConds) Valid() bool {
	// TODO:
	return true
}

func (this *RegisteredConds) ParseConditions() []*Condition {
	conds := make([]*Condition, 0)
	if len(this.Repos) == 1 {
		conds = append(conds, &Condition{Key: "repo_id", Value: this.Repos[0], Kind: OP_EQUAL})
	} else if len(this.Repos) > 1 {
		conds = append(conds, &Condition{Key: "repo_id", Value: this.Repos, Kind: OP_IN})
	}
	if len(this.Ids) > 0 {
		conds = append(conds, &Condition{Key: "civil_id", Value: this.Ids, Kind: OP_IN})
	}

	if this.GetConditions() != nil {
		conds = append(conds, parseConditions(this.GetConditions())...)
	}

	return conds
}

// --------------------------
func (this *TimeSpacialRangeConds) Valid() bool {
	if this.GetTimeRanges() == nil || len(this.GetTimeRanges()) == 0 || len(this.SensorIds) == 0 {
		return false
	}
	return true
}

func (this *TimeSpacialRangeConds) ParseLogicalCondition() *LogicalCondition {
	conds := make([]*Condition, 0)
	if len(this.SensorIds) == 1 {
		conds = append(conds, &Condition{Key: "sensor_id", Value: this.SensorIds[0], Kind: OP_EQUAL})
	} else if len(this.SensorIds) > 1 {
		conds = append(conds, &Condition{Key: "sensor_id", Value: this.SensorIds, Kind: OP_IN})
	}

	lconds := make([]*LogicalCondition, 0)
	for _, timerange := range this.GetTimeRanges() {
		lconds = append(lconds, timerange.ParseLogicalCondition())
	}
	logicalCondtion := &LogicalCondition{
		LogicalConditions: lconds,
		Kind:              LOGICAL_OR,
	}
	return &LogicalCondition{
		Conditions:        conds,
		LogicalConditions: []*LogicalCondition{logicalCondtion},
		Kind:              LOGICAL_AND,
	}
}

func (this *QueryTimeRange) ParseLogicalCondition() *LogicalCondition {
	return &LogicalCondition{
		Conditions: []*Condition{
			&Condition{Key: "ts", Value: this.TimestampStart, Kind: OP_GREATER_EQUAL},
			&Condition{Key: "ts", Value: this.TimestampEnd, Kind: OP_LESS_EQUAL},
		},
		Kind: LOGICAL_AND,
	}
}

func parseConditions(values map[string]string) []*Condition {
	conds := make([]*Condition, 0)
	for key, value := range values {
		if value == "" {
			continue
		}
		index, err := strconv.ParseInt(key, 10, 32)
		if err != nil {
			continue
		}
		switch AttrConds(int32(index)) {
		// TODO: fuzzy string
		case AttrConds_Attr_Name:
			conds = append(conds, &Condition{Key: "name", Value: fuzzy(value), Kind: OP_LIKE})
		case AttrConds_Attr_IdNo:
			conds = append(conds, &Condition{Key: "id_no", Value: fuzzy(value), Kind: OP_LIKE})
		case AttrConds_Attr_GenderId:
			conds = append(conds, &Condition{Key: "gender_id", Value: value, Kind: OP_EQUAL})
		case AttrConds_Attr_AgeId:
			conds = append(conds, &Condition{Key: "age_id", Value: value, Kind: OP_EQUAL})
		case AttrConds_Attr_NationId:
			conds = append(conds, &Condition{Key: "nation_id", Value: value, Kind: OP_EQUAL})
		case AttrConds_Attr_Address:
			conds = append(conds, &Condition{Key: "address", Value: fuzzy(value), Kind: OP_LIKE})
		case AttrConds_Attr_AgeRange:
			fields := strings.Split(value, "~")
			if len(fields) != 2 {
				break
			}
			conds = append(conds, &Condition{Key: "age_id", Value: fields[0], Kind: OP_GREATER_EQUAL})
			conds = append(conds, &Condition{Key: "age_id", Value: fields[1], Kind: OP_LESS_EQUAL})
		case AttrConds_Attr_IdType:
			conds = append(conds, &Condition{Key: "id_type", Value: value, Kind: OP_EQUAL})
		}
	}
	return conds
}

func ParseConds(values map[string]string) map[string]string {
	rets := make(map[string]string)
	for k, v := range values {
		if v == "" {
			continue
		}
		index, err := strconv.ParseInt(k, 10, 32)
		if err != nil {
			continue
		}
		switch AttrConds(int32(index)) {
		case AttrConds_Attr_Name:
			rets["name"] = v
		case AttrConds_Attr_IdNo:
			rets["id_no"] = v
		case AttrConds_Attr_GenderId:
			rets["gender_id"] = v
		case AttrConds_Attr_AgeId:
			rets["age_id"] = v
		case AttrConds_Attr_NationId:
			rets["nation_id"] = v
		case AttrConds_Attr_Address:
			rets["address"] = v
		}
	}
	return rets
}

func fuzzy(value string) string {
	buf := new(bytes.Buffer)
	buf.WriteString("%")
	buf.WriteString(value)
	buf.WriteString("%")
	return buf.String()
}

func (this *WarnedConds) Valid() bool {
	if this.GetTimeSpacialRangeConds() == nil || !this.GetTimeSpacialRangeConds().Valid() {
		return false
	}
	return true
}

func (this *WarnedConds) ParseConditions() *LogicalCondition {
	logicalCondtion := this.GetTimeSpacialRangeConds().ParseLogicalCondition()

	conds := make([]*Condition, 0)
	if len(this.Repos) == 1 {
		conds = append(conds, &Condition{Key: "repo_id", Value: this.Repos[0], Kind: OP_EQUAL})
	} else if len(this.Repos) > 1 {
		conds = append(conds, &Condition{Key: "repo_id", Value: this.Repos, Kind: OP_IN})
	}
	if this.GetConditions() != nil {
		conds = append(conds, parseConditions(this.GetConditions())...)
	}
	if len(conds) > 0 {
		logicalCondtion.Conditions = append(logicalCondtion.Conditions, conds...)
	}
	return logicalCondtion
}

func (this *AuxiliaryConds) ParseConditions() []*Condition {
	conds := make([]*Condition, 0)
	if this.IsWarned != 0 {
		conds = append(conds, &Condition{Key: "is_warned", Value: this.IsWarned, Kind: OP_EQUAL})
	}

	if this.IsChecked != 0 {
		conds = append(conds, &Condition{Key: "is_checked", Value: this.IsChecked, Kind: OP_EQUAL})
	}

	if this.IsRanked != 0 {
		conds = append(conds, &Condition{Key: "is_ranked", Value: this.IsRanked, Kind: OP_EQUAL})
	}

	if this.Status != 0 {
		conds = append(conds, &Condition{Key: "status", Value: this.Status, Kind: OP_EQUAL})
	}

	return conds
}

// --------------------------
func ParseConditons(conds map[string]string) []*Condition {

	return nil
}

func ParsePagingSet(offset, limit int32) *PagingSet {
	if limit != 0 {
		return &PagingSet{
			Offset: offset,
			Limit:  limit,
		}
	}
	return nil
}

func ParseOrderSet(sortby string, asc bool) *OrderSet {
	// if sortby == "" {
	sortby = "ts" // TODO: use ts temporarily
	// }

	return &OrderSet{
		Key: sortby,
		Asc: asc,
	}
}

func ParseGroupSet(group GroupByDimensions) *GroupSet {
	// TODO:
	return nil
}

func ParseColumns(columns []string) []*Column {
	rets := make([]*Column, 0)
	if columns == nil || len(columns) == 0 {
		rets = append(rets, &Column{Key: "*"})
		return rets
	}

	for _, column := range columns {
		rets = append(rets, &Column{Key: column})
	}

	return rets
}
