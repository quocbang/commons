package rs

import (
	"fmt"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
)

//ToErpStageName turns recipe/spec stage to erp stage name
func ToErpStageName(stage rs.StageStatus) (string, error) {
	v, ok := toErpStageName[stage]
	if !ok {
		return "", fmt.Errorf("bad stage=%v", stage)
	}
	return v, nil
}

var toErpStageName = map[rs.StageStatus]string{
	rs.StageStatus_TEMPORARY:    "R",
	rs.StageStatus_PRODUCTION:   "N",
	rs.StageStatus_EXPERIMENTAL: "P",
	rs.StageStatus_TRIAL_RUN:    "T",
}

//ToErpStageValue turns recipe/spec stage to erp stage value
func ToErpStageValue(stage rs.StageStatus) (int, error) {
	v, ok := toErpStageValue[stage]
	if !ok {
		return 0, fmt.Errorf("bad stage=%v", stage)
	}
	return v, nil
}

var toErpStageValue = map[rs.StageStatus]int{
	rs.StageStatus_TEMPORARY:    0,
	rs.StageStatus_PRODUCTION:   1,
	rs.StageStatus_EXPERIMENTAL: 3,
	rs.StageStatus_TRIAL_RUN:    6,
}
