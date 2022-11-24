package kendaql

import (
	"fmt"
	dmKql "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/kendaql"
)

//Parse value
func Parse(data []byte, typ dmKql.DataType) (i interface{}, err error) {
	if len(data) == 0 {
		if typ == dmKql.DataType_STRING {
			return "", nil
		}
		return nil, nil
	}
	switch typ {
	case dmKql.DataType_BOOL:
		i, err = ValueToBool(data)
	case dmKql.DataType_STRING:
		i, err = ValueToString(data)
	case dmKql.DataType_BYTES:
		i, err = ValueToBytes(data)
	case dmKql.DataType_INT:
		i, err = ValueToInt(data)
	case dmKql.DataType_INT8:
		i, err = ValueToInt8(data)
	case dmKql.DataType_INT16:
		i, err = ValueToInt16(data)
	case dmKql.DataType_INT32:
		i, err = ValueToInt32(data)
	case dmKql.DataType_INT64:
		i, err = ValueToInt64(data)
	case dmKql.DataType_FLOAT32:
		i, err = ValueToFloat32(data)
	case dmKql.DataType_FLOAT64:
		i, err = ValueToFloat64(data)
	case dmKql.DataType_TIME:
		i, err = ValueToTime(data)
	default:
		err = fmt.Errorf("not supported type")
	}
	return
}
