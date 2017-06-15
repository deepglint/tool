package dg_model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	OP_EQUAL         int = 0
	OP_NOT_EQUAL         = 1
	OP_GREATER           = 2
	OP_LESS              = 3
	OP_GREATER_EQUAL     = 4
	OP_LESS_EQUAL        = 5
	OP_IN                = 6
	OP_NOT_IN            = 7
	OP_LIKE              = 8
	OP_ANY               = 9
	OP_ALL               = 10
	OP_AND               = 11
	OP_OR                = 12
	OP_NOT               = 13

	// TODO: append
)

const (
	OP_AGGRERATE_COUNT    int = 1
	OP_AGGRERATE_MAX          = 2
	OP_AGGRERATE_MIN          = 3
	OP_AGGRERATE_SUM          = 4
	OP_AGGRERATE_AVG          = 5
	OP_AGGRERATE_DISTINCT     = 6
)

const (
	LOGICAL_AND int = 0
	LOGICAL_OR      = 1
	LOGICAL_XOR     = 2
	LOGICAL_NOT     = 3
)

const (
	EXECUTE_OK             int = 0
	ERROR_CODE_TIMEOUT     int = 2015
	ERROR_EMPTY_TABLE_NAME int = 2016
	ERROR_EMPTY_CQL_PARSED int = 2017
	ERROR_UNKONEN          int = 2018
)

const (
	ElementType_DEFAULT int = 0
	ElementType_OLD         = 1
	ElementType_NEW         = 2
)

type Validation interface {
	Valid() bool
}

// meta value for Insert/Update
type Field struct {
	Key   string
	Value interface{}
}

func (this *Field) Valid() bool {
	if this.Key == "" {
		return false
	}
	return true
}

// a kind of Field/Value , string without ''
type Expression struct {
	Elements []Element
}

type Element struct {
	Type   int
	Value  string
	Parent string //table name ; need by postgres - old value
}

/* condition used by searching */
// meta value for Update/Search/Delete
type Condition struct {
	Key   string
	Value interface{}
	Kind  int
}

// TODO: yiqun
func (this *Condition) Valid() bool {
	if this.Key == "" || this.Kind < 0 || ConvertValueToString(this.Value) == "" {
		return false
	}

	return true
}

/* multi conditions */
type LogicalCondition struct {
	LogicalConditions []*LogicalCondition
	Conditions        []*Condition
	Kind              int
}

func (this *LogicalCondition) Valid() bool {
	for _, logicalCondition := range this.LogicalConditions {
		if !logicalCondition.Valid() {
			return false
		}
	}

	for _, condition := range this.Conditions {
		if !condition.Valid() {
			return false
		}
	}

	return true
}

/* metadata of column */
type Column struct {
	Key  string
	Kind int
}

func (this *Column) Valid() bool {
	if this.Key == "" {
		return false
	}
	return true
}

/* metadata to insert */
type MetaInsert struct {
	Table  string
	Fields []*Field
}

func (this *MetaInsert) Valid() bool {
	if this.Table == "" {
		return false
	}

	for _, field := range this.Fields {
		if !field.Valid() {
			return false
		}
	}
	return true
}

type BatchMetaInsert struct {
	Table       string
	MultiFields [][]*Field
}

func (this *BatchMetaInsert) Valid() bool {
	if this.Table == "" {
		return false
	}
	if len(this.MultiFields) == 0 {
		return false
	}
	for _, fields := range this.MultiFields {
		if len(fields) == 0 {
			return false
		}
		for _, field := range fields {
			if !field.Valid() {
				return false
			}
		}
	}
	return true
}

type MetaInsertOrUpdate struct {
	Table        string
	Fields       []*Field
	UpdateFields []*Field
	UniqueColumn []string // postgtes upsert need CONFLICT constraint name or column
}

func (this *MetaInsertOrUpdate) Valid() bool {
	if this.Table == "" {
		return false
	}
	for _, field := range this.Fields {
		if !field.Valid() {
			return false
		}
	}
	for _, field := range this.UpdateFields {
		if !field.Valid() {
			return false
		}
	}

	if len(this.UniqueColumn) == 0 {
		return false
	}

	return true
}

type BatchMetaInsertOrUpdate struct {
	Table        string
	MultiFields  [][]*Field
	UpdateFields []*Field
	UniqueColumn []string // postgtes upsert need CONFLICT constraint name or column
}

func (this *BatchMetaInsertOrUpdate) Valid() bool {
	if this.Table == "" {
		return false
	}
	if len(this.MultiFields) == 0 {
		return false
	}
	for _, fields := range this.MultiFields {
		if len(fields) == 0 {
			return false
		}
		for _, field := range fields {
			if !field.Valid() {
				return false
			}
		}
	}
	for _, field := range this.UpdateFields {
		if !field.Valid() {
			return false
		}
	}
	return true
}

type GroupSet struct {
	Key   string
	Paral bool
}

type OrderSet struct {
	Key string
	Asc bool
}

type PagingSet struct {
	Offset int32
	Limit  int32
}

type HavingSet struct {
	Kind   int
	Column string
	Op     int
	Value  int32
}

type PartitionSet struct {
	Key   string
	Count int
}

/* metadata to search */
type MetaSearch struct {
	Table      string
	Alias      string
	Columns    []*Column
	Conditions *LogicalCondition
	Group      *GroupSet
	Order      *OrderSet
	Have       *HavingSet
	Page       *PagingSet
	Partition  *PartitionSet
}

func (this *MetaSearch) Valid() bool {
	if this.Table == "" || this.Alias == "" {
		return false
	}

	for _, column := range this.Columns {
		if !column.Valid() {
			return false
		}
	}

	if this.Conditions != nil {
		return this.Conditions.Valid()
	}
	return true
}

func (this *MetaSearch) String() string {
	body, _ := json.Marshal(this)
	return string(body)
}

type MetaUpdate struct {
	Table      string
	Alias      string
	Fields     []*Field
	Conditions []*Condition
}

func (this *MetaUpdate) Valid() bool {
	if this.Table == "" {
		return false
	}

	for _, field := range this.Fields {
		if !field.Valid() {
			return false
		}
	}

	for _, condition := range this.Conditions {
		if !condition.Valid() {
			return false
		}
	}
	return true
}

type MetaDelete struct {
	Table      string
	Alias      string
	Conditions []*Condition
}

// TODO: yiqun
func (this *MetaDelete) Valid() bool {
	if this.Table == "" {
		return false
	}

	for _, condition := range this.Conditions {
		if !condition.Valid() {
			return false
		}
	}
	return true
}

func ConvertValueToString(v interface{}) string {
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Bool:
		return fmt.Sprintf("%t", value.Bool())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%g", value.Float())
	case reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("%g", value.Complex())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", value.Uint())
	case reflect.String:
		return fmt.Sprintf("'%s'", value.String())
	case reflect.Array, reflect.Slice:
		sb := bytes.Buffer{}
		sb.WriteString("(")
		for i := 0; i < value.Len(); i++ {
			if sb.Len() > 1 {
				sb.WriteString(",")
			}
			if tmp := ConvertValueToString(value.Index(i).Interface()); tmp != "" {
				sb.WriteString(tmp)
			}
		}
		sb.WriteString(")")
		if sb.Len() == 2 {
			sb.Reset()
		}
		return sb.String()
	default:
		return ""
	}
}

type MetaCount struct {
	Table      string
	Alias      string
	Conditions *LogicalCondition
	Partition  *PartitionSet
	Group      *GroupSet
}

func (this *MetaCount) Valid() bool {
	if this.Table == "" {
		return false
	}
	return true
}

type MetaAssociateSearch struct {
	AssociateColumns    map[string]string
	BaseMetaSearch      *MetaSearch
	AssociateMetaSearch *MetaSearch
	Page                *PagingSet
}

func (this *MetaAssociateSearch) Valid() bool {
	return len(this.AssociateColumns) > 0 && this.BaseMetaSearch.Valid() && this.AssociateMetaSearch.Valid()
}

type MetaInsertGroup struct {
	RepoId             int32
	SensorId           int32
	Timestamp          int64
	PlateText          string
	MainMetaInsert     *MetaInsert
	AttachedMetaInsert []*MetaInsert
}
