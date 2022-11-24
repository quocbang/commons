package mes

import "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/mes/code"

func MaterialStatusMapping(materialStatus code.MaterialStatus) []code.MaterialStatus {
	return map[code.MaterialStatus][]code.MaterialStatus{
		code.MaterialStatus_AVAILABLE: []code.MaterialStatus{
			code.MaterialStatus_AVAILABLE,
			code.MaterialStatus_HOLD,
			code.MaterialStatus_INSPECTION,
			code.MaterialStatus_MOUNTED,
			code.MaterialStatus_UNAVAILABLE,
		},
		code.MaterialStatus_HOLD: []code.MaterialStatus{
			code.MaterialStatus_AVAILABLE,
			code.MaterialStatus_UNAVAILABLE,
		},
	}[materialStatus]
}
