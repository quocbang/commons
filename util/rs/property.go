package rs

import (
	"fmt"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
)

// PropertySettingNameToValue get property setting value by specified property and value.(string)
//
// Examples:
//
// 	   PropertySettingNameToValue(rs.PropertyName_RAM_POSITION.String(), "UP")          // output: 1, nil
//
func PropertySettingNameToValue(property string, value string) (int32, error) {
	values, ok := propertySettingValueEnum[property]
	if !ok {
		return 0, fmt.Errorf("bad property name=%s", property)
	}
	val, ok := values[value]
	if !ok {
		return 0, fmt.Errorf("property name=%s has bad property value=%v", property, value)
	}
	return val, nil
}

// PropertySettingValueToName get property setting name by specified property and setting value.(int32)
//
// Examples:
//
// 	   PropertySettingValueToName(rs.PropertyName_DROPPING_LOGIC.String(), 1) // output: MIXING_TIME_ONLY, nil
//
func PropertySettingValueToName(property string, value int32) (string, error) {
	names, ok := PropertySettingNameEnum[property]
	if !ok {
		return "", fmt.Errorf("bad property name=%s", property)
	}
	name, ok := names[value]
	if !ok {
		return "", fmt.Errorf("property name=%s has bad property value=%v", property, value)
	}
	return name, nil
}

// propertySettingValueEnum is property and setting values mapping list
var propertySettingValueEnum = map[string]map[string]int32{
	rs.PropertyName_RAM_POSITION.String():                                        rs.RamPosition_value,
	rs.PropertyName_FEED_GATE_STATUS.String():                                    rs.GateStatus_value,
	rs.PropertyName_ACTION_TYPE.String():                                         rs.ActionType_value,
	rs.PropertyName_DROPPING_LOGIC.String():                                      rs.DroppingLogic_value,
	rs.PropertyName_FIRST_LEFT_COLOR_LINE.String():                               rs.ColorCode_value,
	rs.PropertyName_SECOND_LEFT_COLOR_LINE.String():                              rs.ColorCode_value,
	rs.PropertyName_THIRD_LEFT_COLOR_LINE.String():                               rs.ColorCode_value,
	rs.PropertyName_CENTER_COLOR_LINE.String():                                   rs.ColorCode_value,
	rs.PropertyName_FIRST_RIGHT_COLOR_LINE.String():                              rs.ColorCode_value,
	rs.PropertyName_SECOND_RIGHT_COLOR_LINE.String():                             rs.ColorCode_value,
	rs.PropertyName_THIRD_RIGHT_COLOR_LINE.String():                              rs.ColorCode_value,
	rs.PropertyName_FOURTH_RIGHT_COLOR_LINE.String():                             rs.ColorCode_value,
	rs.PropertyName_DISTANCE_BETWEEN_CENTER_AND_FIRST_LEFT_COLOR_LINE.String():   rs.ColorCode_value,
	rs.PropertyName_DISTANCE_BETWEEN_CENTER_AND_SECOND_LEFT_COLOR_LINE_.String(): rs.ColorCode_value,
	rs.PropertyName_DISTANCE_BETWEEN_CENTER_AND_THIRD_LEFT_COLOR_LINE.String():   rs.ColorCode_value,
	rs.PropertyName_DISTANCE_BETWEEN_CENTER_AND_FIRST_RIGHT_COLOR_LINE.String():  rs.ColorCode_value,
	rs.PropertyName_DISTANCE_BETWEEN_CENTER_AND_SECOND_RIGHT_COLOR_LINE.String(): rs.ColorCode_value,
	rs.PropertyName_DISTANCE_BETWEEN_CENTER_AND_THIRD_RIGHT_COLOR_LINE.String():  rs.ColorCode_value,
	rs.PropertyName_DISTANCE_BETWEEN_CENTER_AND_FOURTH_RIGHT_COLOR_LINE.String(): rs.ColorCode_value,
	rs.PropertyName_BUILDING_METHOD.String():                                     rs.BuildingMethod_value,
	rs.PropertyName_FIRST_BUILDING_METHOD.String():                               rs.FirstBuildingMethod_value,
	rs.PropertyName_FIRST_CAP_PLY_WINDING_STRUCTURE.String():                     rs.CapPlyStructure_value,
	rs.PropertyName_SECOND_CAP_PLY_WINDING_STRUCTURE.String():                    rs.CapPlyStructure_value,
	rs.PropertyName_EDGE_GUM_POSITION.String():                                   rs.EdgeGumPosition_value,
	rs.PropertyName_DUST_COLLECTOR_SETTING.String():                              rs.DustCollectorSetting_value,
}

// PropertySettingNameEnum is property and setting names mapping list
var PropertySettingNameEnum = map[string]map[int32]string{
	rs.PropertyName_RAM_POSITION.String():                                        rs.RamPosition_name,
	rs.PropertyName_FEED_GATE_STATUS.String():                                    rs.GateStatus_name,
	rs.PropertyName_ACTION_TYPE.String():                                         rs.ActionType_name,
	rs.PropertyName_DROPPING_LOGIC.String():                                      rs.DroppingLogic_name,
	rs.PropertyName_FIRST_LEFT_COLOR_LINE.String():                               rs.ColorCode_name,
	rs.PropertyName_SECOND_LEFT_COLOR_LINE.String():                              rs.ColorCode_name,
	rs.PropertyName_THIRD_LEFT_COLOR_LINE.String():                               rs.ColorCode_name,
	rs.PropertyName_CENTER_COLOR_LINE.String():                                   rs.ColorCode_name,
	rs.PropertyName_FIRST_RIGHT_COLOR_LINE.String():                              rs.ColorCode_name,
	rs.PropertyName_SECOND_RIGHT_COLOR_LINE.String():                             rs.ColorCode_name,
	rs.PropertyName_THIRD_RIGHT_COLOR_LINE.String():                              rs.ColorCode_name,
	rs.PropertyName_FOURTH_RIGHT_COLOR_LINE.String():                             rs.ColorCode_name,
	rs.PropertyName_DISTANCE_BETWEEN_CENTER_AND_FIRST_LEFT_COLOR_LINE.String():   rs.ColorCode_name,
	rs.PropertyName_DISTANCE_BETWEEN_CENTER_AND_SECOND_LEFT_COLOR_LINE_.String(): rs.ColorCode_name,
	rs.PropertyName_DISTANCE_BETWEEN_CENTER_AND_THIRD_LEFT_COLOR_LINE.String():   rs.ColorCode_name,
	rs.PropertyName_DISTANCE_BETWEEN_CENTER_AND_FIRST_RIGHT_COLOR_LINE.String():  rs.ColorCode_name,
	rs.PropertyName_DISTANCE_BETWEEN_CENTER_AND_SECOND_RIGHT_COLOR_LINE.String(): rs.ColorCode_name,
	rs.PropertyName_DISTANCE_BETWEEN_CENTER_AND_THIRD_RIGHT_COLOR_LINE.String():  rs.ColorCode_name,
	rs.PropertyName_DISTANCE_BETWEEN_CENTER_AND_FOURTH_RIGHT_COLOR_LINE.String(): rs.ColorCode_name,
	rs.PropertyName_BUILDING_METHOD.String():                                     rs.BuildingMethod_name,
	rs.PropertyName_FIRST_BUILDING_METHOD.String():                               rs.FirstBuildingMethod_name,
	rs.PropertyName_FIRST_CAP_PLY_WINDING_STRUCTURE.String():                     rs.CapPlyStructure_name,
	rs.PropertyName_SECOND_CAP_PLY_WINDING_STRUCTURE.String():                    rs.CapPlyStructure_name,
	rs.PropertyName_EDGE_GUM_POSITION.String():                                   rs.EdgeGumPosition_name,
	rs.PropertyName_DUST_COLLECTOR_SETTING.String():                              rs.DustCollectorSetting_name,
}
