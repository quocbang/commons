package kendaql

import (
	"time"

	"gitlab.kenda.com.tw/kenda/commons/v2/util/math"
)

// ValueToBool []byte to bool.
func ValueToBool(bs []byte) (bool, error) {
	var val bool
	err := math.ParseFromBytes(bs, &val)
	return val, err
}

// ValueToString []byte to string.
func ValueToString(bs []byte) (string, error) {
	return string(bs), nil
}

// ValueToBytes []byte to []byte.
func ValueToBytes(bs []byte) ([]byte, error) {
	return bs, nil
}

// ValueToInt []byte to int.
func ValueToInt(bs []byte) (int, error) {
	var val int
	err := math.ParseFromBytes(bs, &val)
	return val, err
}

// ValueToInt8 []byte to Int8.
func ValueToInt8(bs []byte) (int8, error) {
	var val int8
	err := math.ParseFromBytes(bs, &val)
	return val, err
}

// ValueToInt16 []byte to Int16.
func ValueToInt16(bs []byte) (int16, error) {
	var val int16
	err := math.ParseFromBytes(bs, &val)
	return val, err
}

// ValueToInt32 []byte to toInt32.
func ValueToInt32(bs []byte) (int32, error) {
	var val int32
	err := math.ParseFromBytes(bs, &val)
	return val, err
}

// ValueToInt64 []byte to Int64.
func ValueToInt64(bs []byte) (int64, error) {
	var val int64
	err := math.ParseFromBytes(bs, &val)
	return val, err
}

// ValueToFloat32 []byte to float32.
func ValueToFloat32(bs []byte) (float32, error) {
	var val float32
	err := math.ParseFromBytes(bs, &val)
	return val, err
}

// ValueToFloat64 []byte to float64.
func ValueToFloat64(bs []byte) (float64, error) {
	var val float64
	err := math.ParseFromBytes(bs, &val)
	return val, err
}

// ValueToTime []byte to time.
func ValueToTime(bs []byte) (time.Time, error) {
	var val time.Time
	err := val.UnmarshalBinary(bs)
	return val, err
}

// ToValue any type to []byte.
func ToValue(data interface{}) ([]byte, error) {
	switch t := data.(type) {
	case string:
		return []byte(t), nil
	case []byte:
		return t, nil
	case time.Time:
		return t.MarshalBinary()
	}
	return math.ToBytes(data)
}
