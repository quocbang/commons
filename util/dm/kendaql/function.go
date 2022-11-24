package kendaql

import (
	"errors"
	"reflect"
	"time"

	pbkql "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/kendaql"
)

// Func is the function.
type Func struct {
	f   *pbkql.Function
	err error
}

// Count is a COUNT function.
func Count(columnName string) *Func {
	return &Func{
		f: &pbkql.Function{
			Name: "COUNT",
			Parameters: []*pbkql.Function_Param{
				&pbkql.Function_Param{Value: &pbkql.Function_Param_ColumnValue{ColumnValue: &pbkql.Column{
					Name: columnName,
				}}},
			},
			ExpectedType: pbkql.DataType_INT,
		},
	}
}

// Max is a MAX function.
func Max(columnName string) *Func {
	return &Func{
		f: &pbkql.Function{
			Name: "MAX",
			Parameters: []*pbkql.Function_Param{
				&pbkql.Function_Param{Value: &pbkql.Function_Param_ColumnValue{ColumnValue: &pbkql.Column{
					Name: columnName,
				}}},
			},
		},
	}
}

func (f *Func) isCol() (*pbkql.Col, error) {
	if f.err != nil {
		return nil, f.err
	}
	return &pbkql.Col{Col: &pbkql.Col_Function{
		Function: f.f,
	}}, nil
}

// ForTable specifies what table contains this column.
func (f *Func) ForTable(name string) *Func {
	params := f.f.GetParameters()
	if len(params) == 0 {
		f.err = errors.New("kendaql: the length of the parameter is 0")
		return f
	}

	col := params[0].GetColumnValue()
	if col == nil {
		f.err = errors.New("kendaql: the first of the parameter is not column type")
		return f
	}

	col.TableAlias = name
	return f
}

// WithExpectedType expects what data type will return.
func (f *Func) WithExpectedType(typ pbkql.DataType) *Func {
	f.f.ExpectedType = typ
	return f
}

// Bools sets the values.
func (f *Func) Bools(values ...bool) *ColumnWithValue {
	return f.interfaces(reflect.ValueOf(values))
}

// Strings sets the values.
func (f *Func) Strings(values ...string) *ColumnWithValue {
	return f.interfaces(reflect.ValueOf(values))
}

// MultiBytes sets the values.
func (f *Func) MultiBytes(values ...[]byte) *ColumnWithValue {
	return f.interfaces(reflect.ValueOf(values))
}

// Ints sets the values.
func (f *Func) Ints(values ...int) *ColumnWithValue {
	return f.interfaces(reflect.ValueOf(values))
}

// Int8s sets the values.
func (f *Func) Int8s(values ...int8) *ColumnWithValue {
	return f.interfaces(reflect.ValueOf(values))
}

// Int16s sets the values.
func (f *Func) Int16s(values ...int16) *ColumnWithValue {
	return f.interfaces(reflect.ValueOf(values))
}

// Int32s sets the values.
func (f *Func) Int32s(values ...int32) *ColumnWithValue {
	return f.interfaces(reflect.ValueOf(values))
}

// Int64s sets the values.
func (f *Func) Int64s(values ...int64) *ColumnWithValue {
	return f.interfaces(reflect.ValueOf(values))
}

// Float32s sets the values.
func (f *Func) Float32s(values ...float32) *ColumnWithValue {
	return f.interfaces(reflect.ValueOf(values))
}

// Float64s sets the values.
func (f *Func) Float64s(values ...float64) *ColumnWithValue {
	return f.interfaces(reflect.ValueOf(values))
}

// Times sets the values.
func (f *Func) Times(values ...time.Time) *ColumnWithValue {
	return f.interfaces(reflect.ValueOf(values))
}

func (f *Func) interfaces(rv reflect.Value) *ColumnWithValue {
	vals, err := toBytesSlice(rv)
	if err != nil {
		return &ColumnWithValue{err: err}
	}

	col, err := f.isCol()
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
