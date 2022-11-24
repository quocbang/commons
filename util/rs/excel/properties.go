package excel

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
	utilRs "gitlab.kenda.com.tw/kenda/commons/v2/util/rs"
)

func constraintToString(constraint *rs.Param_Constraint) (string, bool) {
	if constraint != nil {
		return decimal.NewFromFloat(float64(constraint.GetValue())).Truncate(3).String(), true
	}
	return "", false
}

func stringToConstraint(value string) (*rs.Param_Constraint, error) {
	constraint := &rs.Param_Constraint{}
	if value != "" {
		num, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to convert string to float")
		}
		constraint.Value = float32(num)
	}
	return constraint, nil
}

func propertyToString(property *rs.Property) (string, error) {
	param := property.GetParam()
	if param == nil {
		return "", nil
	}
	if enum := param.GetEnum(); enum != nil {
		if value := enum.GetValue(); value != 0 {
			return utilRs.PropertySettingValueToName(property.GetName(), value)
		}
		return "", nil // empty strings for zero/unspecified values
	}

	// Not enum -> low/mid/high values
	return paramValuesToString(param), nil
}

func constraintToLabelledString(label string, constraint *rs.Param_Constraint) (string, bool) {
	if s, ok := constraintToString(constraint); ok {
		return fmt.Sprintf("%s:%s", label, s), true
	}
	return "", false
}

func paramValuesToString(param *rs.Param) string {
	ss := make([]string, 0, 3)
	if v, exists := constraintToLabelledString("low", param.GetMin()); exists {
		ss = append(ss, v)
	}
	if v, exists := constraintToLabelledString("mid", param.GetCentral()); exists {
		ss = append(ss, v)
	}
	if v, exists := constraintToLabelledString("high", param.GetMax()); exists {
		ss = append(ss, v)
	}
	return strings.Join(ss, ";")
}

func stringToParamValues(low, mid, high, unit string) (param *rs.Param, err error) {
	param = &rs.Param{}
	var min, central, max *rs.Param_Constraint
	if min, err = stringToConstraint(low); err != nil {
		return
	}
	if central, err = stringToConstraint(mid); err != nil {
		return
	}
	if max, err = stringToConstraint(high); err != nil {
		return
	}
	switch {
	case min.Value != 0 && central.Value == 0 && max.Value == 0:
		param = &rs.Param{
			Type: rs.Param_MIN,
			Unit: unit,
			Min:  min,
		}
	case min.Value == 0 && central.Value != 0 && max.Value == 0:
		param = &rs.Param{
			Type:    rs.Param_VALUE,
			Unit:    unit,
			Central: central,
		}
	case min.Value == 0 && central.Value == 0 && max.Value != 0:
		param = &rs.Param{
			Type: rs.Param_MAX,
			Unit: unit,
			Max:  max,
		}
	default:
		if min.Value == 0 {
			min = nil
		}
		if central.Value == 0 {
			central = nil
		}
		if max.Value == 0 {
			max = nil
		}
		param = &rs.Param{
			Type:    rs.Param_RANGE,
			Unit:    unit,
			Min:     min,
			Central: central,
			Max:     max,
		}
	}
	return
}
