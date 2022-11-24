package excel

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
)

func Test_constraintToString(t *testing.T) {
	{ // good case
		actualValue, bool := constraintToString(&rs.Param_Constraint{Value: 20})
		assert.Equal(t, []interface{}{"20", true}, []interface{}{actualValue, bool})
	}
	{ // nil param constraint
		actualValue, bool := constraintToString(nil)
		assert.Equal(t, []interface{}{"", false}, []interface{}{actualValue, bool})
	}
}

func Test_stringToConstraint(t *testing.T) {
	{ // good case
		actualValue, err := stringToConstraint("45")
		assert.Equal(t, []interface{}{&rs.Param_Constraint{Value: 45}, nil}, []interface{}{actualValue, err})
	}
	{ // empty value
		actualValue, err := stringToConstraint("")
		assert.Equal(t, []interface{}{&rs.Param_Constraint{Value: 0}, nil}, []interface{}{actualValue, err})
	}
}

func Test_propertyToString(t *testing.T) {
	// nil param
	{
		nilProperty := &rs.Property{}
		value, err := propertyToString(nilProperty)
		assert.NoError(t, err)
		assert.Equal(t, "", value)
	}

	// enum tests
	{ // 1. unspecified
		unspecifiedProperty := &rs.Property{
			Name: rs.PropertyName_DROPPING_LOGIC.String(),
			Param: &rs.Param{
				Enum: &rs.Param_Enum{Value: 0},
			},
		}
		value, err := propertyToString(unspecifiedProperty)
		assert.NoError(t, err)
		assert.Equal(t, "", value)
	}
	{ // 2. bad value
		badValueProperty := &rs.Property{
			Name: rs.PropertyName_ACTION_TYPE.String(),
			Param: &rs.Param{
				Enum: &rs.Param_Enum{Value: 7},
			},
		}
		_, err := propertyToString(badValueProperty)
		assert.Equal(t, errors.New("property name=ACTION_TYPE has bad property value=7"), err)
	}
	{ // 3. good case
		goodProperty := &rs.Property{
			Name: rs.PropertyName_RAM_POSITION.String(),
			Param: &rs.Param{
				Enum: &rs.Param_Enum{Value: 1},
			},
		}
		enum, err := propertyToString(goodProperty)
		assert.NoError(t, err)
		assert.Equal(t, rs.RamPosition_UP.String(), enum)
	}

	// value tests
	{ // 1. range (max + min)
		maxMinProperty := &rs.Property{
			Param: &rs.Param{
				Max: &rs.Param_Constraint{Value: 45},
				Min: &rs.Param_Constraint{Value: 38},
			},
		}
		value, _ := propertyToString(maxMinProperty)
		assert.Equal(t, "low:38;high:45", value)
	}
	{ // 2. range (max + mid + min)
		allValuesProperty := &rs.Property{
			Param: &rs.Param{
				Max:     &rs.Param_Constraint{Value: 10},
				Central: &rs.Param_Constraint{Value: 9},
				Min:     &rs.Param_Constraint{Value: 8},
			},
		}
		value, _ := propertyToString(allValuesProperty)
		assert.Equal(t, "low:8;mid:9;high:10", value)
	}
	{ // 3. mid
		midProperty := &rs.Property{Param: &rs.Param{Central: &rs.Param_Constraint{Value: 40}}}
		value, _ := propertyToString(midProperty)
		assert.Equal(t, "mid:40", value)
	}
	{ // 4. max
		maxProperty := &rs.Property{Param: &rs.Param{Max: &rs.Param_Constraint{Value: 25}}}
		value, _ := propertyToString(maxProperty)
		assert.Equal(t, "high:25", value)
	}
	{ // 5. min
		minProperty := &rs.Property{Param: &rs.Param{Min: &rs.Param_Constraint{Value: 10}}}
		value, _ := propertyToString(minProperty)
		assert.Equal(t, "low:10", value)
	}
}

func Test_constraintToLabelledString(t *testing.T) {
	{ // good case
		actual, bool := constraintToLabelledString("low", &rs.Param_Constraint{Value: 20})
		assert.Equal(t, []interface{}{"low:20", true}, []interface{}{actual, bool})
	}
	{ // nil param constraint
		actual, bool := constraintToLabelledString("", nil)
		assert.Equal(t, []interface{}{"", false}, []interface{}{actual, bool})
	}
}

func Test_paramValuesToString(t *testing.T) {
	goodParam := &rs.Param{
		Min:     &rs.Param_Constraint{Value: 20},
		Central: &rs.Param_Constraint{Value: 25},
		Max:     &rs.Param_Constraint{Value: 30},
	}
	// range test
	{
		actualValue := paramValuesToString(goodParam)
		assert.Equal(t, "low:20;mid:25;high:30", actualValue)
	}
	{
		value := *goodParam
		value.Central = nil
		actualValue := paramValuesToString(&value)
		assert.Equal(t, "low:20;high:30", actualValue)
	}
	// min test
	{
		value := *goodParam
		value.Central = nil
		value.Max = nil
		actualValue := paramValuesToString(&value)
		assert.Equal(t, "low:20", actualValue)
	}
	// central test
	{
		value := *goodParam
		value.Min = nil
		value.Max = nil
		actualValue := paramValuesToString(&value)
		assert.Equal(t, "mid:25", actualValue)
	}
	// max test
	{
		value := *goodParam
		value.Min = nil
		value.Central = nil
		actualValue := paramValuesToString(&value)
		assert.Equal(t, "high:30", actualValue)
	}
}

func Test_stringToParamValues(t *testing.T) {
	expectedParam := &rs.Param{
		Unit:    "PHR",
		Min:     &rs.Param_Constraint{Value: 1},
		Central: &rs.Param_Constraint{Value: 2},
		Max:     &rs.Param_Constraint{Value: 3},
	}
	// good case
	{ // min test
		actualParam, err := stringToParamValues("1", "", "", "PHR")
		assert.NoError(t, err)

		param := *expectedParam
		param.Type = rs.Param_MIN
		param.Central = nil
		param.Max = nil
		assert.Equal(t, &param, actualParam)
	}
	{ // central test
		actualParam, err := stringToParamValues("", "2", "", "PHR")
		assert.NoError(t, err)

		param := *expectedParam
		param.Type = rs.Param_VALUE
		param.Min = nil
		param.Max = nil
		assert.Equal(t, &param, actualParam)
	}
	{ // max test
		actualParam, err := stringToParamValues("", "", "3", "PHR")
		assert.NoError(t, err)

		param := *expectedParam
		param.Type = rs.Param_MAX
		param.Min = nil
		param.Central = nil
		assert.Equal(t, &param, actualParam)
	}
	{ // range test
		actualParam, err := stringToParamValues("1", "", "3", "PHR")
		assert.NoError(t, err)

		param := *expectedParam
		param.Type = rs.Param_RANGE
		param.Central = nil
		assert.Equal(t, &param, actualParam)
	}
}
