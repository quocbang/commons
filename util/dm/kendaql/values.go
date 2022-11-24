package kendaql

import (
	"reflect"
)

func toBytesSlice(rv reflect.Value) ([][]byte, error) {
	if rv.Kind() != reflect.Slice {
		return nil, &callerError{caller: "toBytesSlice", err: ErrSliceParameterOnly}
	}

	length := rv.Len()
	vals := make([][]byte, length)
	for i := 0; i < length; i++ {
		val, err := ToValue(rv.Index(i).Interface())
		if err != nil {
			return nil, err
		}
		vals[i] = val
	}
	return vals, nil
}
