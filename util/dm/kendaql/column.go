package kendaql

import (
	"reflect"
	"time"

	pbkql "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/kendaql"
)

// Col is the column.
type Col struct {
	col *pbkql.Column
}

// Column creates a column.
func Column(name string) *Col {
	return &Col{
		col: &pbkql.Column{
			Name: name,
		},
	}
}

func (column *Col) isCol() (*pbkql.Col, error) {
	return &pbkql.Col{Col: &pbkql.Col_Column{
		Column: column.col,
	}}, nil
}

// ForTable specifies what table contains this column.
func (column *Col) ForTable(table string) *Col {
	column.col.TableAlias = table
	return column
}

// Bools sets the values.
func (column *Col) Bools(values ...bool) *ColumnWithValue {
	return column.interfaces(reflect.ValueOf(values))
}

// Strings sets the values.
func (column *Col) Strings(values ...string) *ColumnWithValue {
	return column.interfaces(reflect.ValueOf(values))
}

// MultiBytes sets the values.
func (column *Col) MultiBytes(values ...[]byte) *ColumnWithValue {
	return column.interfaces(reflect.ValueOf(values))
}

// Ints sets the values.
func (column *Col) Ints(values ...int) *ColumnWithValue {
	return column.interfaces(reflect.ValueOf(values))
}

// Int8s sets the values.
func (column *Col) Int8s(values ...int8) *ColumnWithValue {
	return column.interfaces(reflect.ValueOf(values))
}

// Int16s sets the values.
func (column *Col) Int16s(values ...int16) *ColumnWithValue {
	return column.interfaces(reflect.ValueOf(values))
}

// Int32s sets the values.
func (column *Col) Int32s(values ...int32) *ColumnWithValue {
	return column.interfaces(reflect.ValueOf(values))
}

// Int64s sets the values.
func (column *Col) Int64s(values ...int64) *ColumnWithValue {
	return column.interfaces(reflect.ValueOf(values))
}

// Float32s sets the values.
func (column *Col) Float32s(values ...float32) *ColumnWithValue {
	return column.interfaces(reflect.ValueOf(values))
}

// Float64s sets the values.
func (column *Col) Float64s(values ...float64) *ColumnWithValue {
	return column.interfaces(reflect.ValueOf(values))
}

// Times sets the values.
func (column *Col) Times(values ...time.Time) *ColumnWithValue {
	return column.interfaces(reflect.ValueOf(values))
}

func (column *Col) interfaces(rv reflect.Value) *ColumnWithValue {
	vals, err := toBytesSlice(rv)
	if err != nil {
		return &ColumnWithValue{err: err}
	}

	col, err := column.isCol()
	if err != nil {
		return &ColumnWithValue{err: err}
	}

	return &ColumnWithValue{
		mf: &pbkql.Field{
			Column: col,
			Values: vals,
		},
	}
}

// Asc determines the way to order.
func (column *Col) Asc() *pbkql.OrderBy {
	return &pbkql.OrderBy{
		Type:   pbkql.Order_ASCENDING,
		Column: column.col,
	}
}

// Desc determines the way to order.
func (column *Col) Desc() *pbkql.OrderBy {
	return &pbkql.OrderBy{
		Type:   pbkql.Order_DESCENDING,
		Column: column.col,
	}
}

// ColumnWithValue is used to do mutation operating.
type ColumnWithValue struct {
	mf  *pbkql.Field
	err error
}

// Equal Equal
func (column *ColumnWithValue) Equal() *CondField {
	if column.err != nil {
		return &CondField{e: column.err}
	}
	return &CondField{
		cf: &pbkql.ConditionField{
			Condition: pbkql.Condition_EQUAL,
			Field:     column.mf,
		},
	}
}

// NotEqual Not_equal
func (column *ColumnWithValue) NotEqual() *CondField {
	if column.err != nil {
		return &CondField{e: column.err}
	}
	return &CondField{
		cf: &pbkql.ConditionField{
			Condition: pbkql.Condition_NOT_EQUAL,
			Field:     column.mf,
		},
	}
}

// Like Like
func (column *ColumnWithValue) Like() *CondField {
	if column.err != nil {
		return &CondField{e: column.err}
	}
	return &CondField{
		cf: &pbkql.ConditionField{
			Condition: pbkql.Condition_LIKE,
			Field:     column.mf,
		},
	}
}

// NotLike Not_like
func (column *ColumnWithValue) NotLike() *CondField {
	if column.err != nil {
		return &CondField{e: column.err}
	}
	return &CondField{
		cf: &pbkql.ConditionField{
			Condition: pbkql.Condition_NOT_LIKE,
			Field:     column.mf,
		},
	}
}

// Greater Greater
func (column *ColumnWithValue) Greater() *CondField {
	if column.err != nil {
		return &CondField{e: column.err}
	}
	return &CondField{
		cf: &pbkql.ConditionField{
			Condition: pbkql.Condition_GREATER,
			Field:     column.mf,
		},
	}
}

// GreaterEqual Greater_equal
func (column *ColumnWithValue) GreaterEqual() *CondField {
	if column.err != nil {
		return &CondField{e: column.err}
	}
	return &CondField{
		cf: &pbkql.ConditionField{
			Condition: pbkql.Condition_GREATER_EQUAL,
			Field:     column.mf,
		},
	}
}

// Less Less
func (column *ColumnWithValue) Less() *CondField {
	if column.err != nil {
		return &CondField{e: column.err}
	}
	return &CondField{
		cf: &pbkql.ConditionField{
			Condition: pbkql.Condition_LESS,
			Field:     column.mf,
		},
	}
}

// LessEqual Less_equal
func (column *ColumnWithValue) LessEqual() *CondField {
	if column.err != nil {
		return &CondField{e: column.err}
	}
	return &CondField{
		cf: &pbkql.ConditionField{
			Condition: pbkql.Condition_LESS_EQUAL,
			Field:     column.mf,
		},
	}
}

// Between Between
func (column *ColumnWithValue) Between() *CondField {
	if column.err != nil {
		return &CondField{e: column.err}
	}
	return &CondField{
		cf: &pbkql.ConditionField{
			Condition: pbkql.Condition_BETWEEN,
			Field:     column.mf,
		},
	}
}

// NotBetween Not_between
func (column *ColumnWithValue) NotBetween() *CondField {
	if column.err != nil {
		return &CondField{e: column.err}
	}
	return &CondField{
		cf: &pbkql.ConditionField{
			Condition: pbkql.Condition_NOT_BETWEEN,
			Field:     column.mf,
		},
	}
}

// In In
func (column *ColumnWithValue) In() *CondField {
	if column.err != nil {
		return &CondField{e: column.err}
	}
	return &CondField{
		cf: &pbkql.ConditionField{
			Condition: pbkql.Condition_IN,
			Field:     column.mf,
		},
	}
}

// NotIn Not_in
func (column *ColumnWithValue) NotIn() *CondField {
	if column.err != nil {
		return &CondField{e: column.err}
	}
	return &CondField{
		cf: &pbkql.ConditionField{
			Condition: pbkql.Condition_NOT_IN,
			Field:     column.mf,
		},
	}
}

type isLogicalGroup interface {
	logicalGroupNode() *pbkql.LogicalGroupNode
	logicalGroup() *pbkql.LogicalGroup
	err() error
}

// CondField is kandaql condition field.
type CondField struct {
	cf *pbkql.ConditionField
	e  error
}

func (cf *CondField) logicalGroupNode() *pbkql.LogicalGroupNode {
	return &pbkql.LogicalGroupNode{
		Field: &pbkql.LogicalGroupNode_Condition{Condition: cf.cf},
	}
}

func (cf *CondField) logicalGroup() *pbkql.LogicalGroup {
	return And(cf).lg
}

func (cf *CondField) err() error {
	return cf.e
}
