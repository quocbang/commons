package decimal

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"

	mesCommons "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/mes/v2/commons"
)

// FromString returns Decimal representation for s.
func FromString(s string) (*mesCommons.Decimal, error) {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return nil, err
	}
	return FromDecimal(d)
}

// MustFromString returns Decimal representation for s, it panics if any error.
func MustFromString(s string) *mesCommons.Decimal {
	dec, err := FromString(s)
	if err != nil {
		panic(err)
	}
	return dec
}

// FromFloat returns Decimal representation for f.
func FromFloat(f float64) (*mesCommons.Decimal, error) {
	return FromDecimal(decimal.NewFromFloat(f))
}

// MustFromFloat returns Decimal representation for f, it panics if any error.
func MustFromFloat(f float64) *mesCommons.Decimal {
	dec, err := FromFloat(f)
	if err != nil {
		panic(err)
	}
	return dec
}

// ToFloat returns the nearest float64 value for d*10^shift and a bool indicating
// whether f represents d*10^shift exactly.
func ToFloat(d *mesCommons.Decimal, shift int32) (f float64, ok bool) {
	return decimal.New(d.Value, d.Exp+shift).Float64()
}

// FromInt8 returns Decimal representation for i.
func FromInt8(i int8) *mesCommons.Decimal {
	return &mesCommons.Decimal{Value: int64(i)}
}

// ToInt8 returns the int8 value for d*10^shift and a bool indicating
// whether the conversion succeeds. See ToInt64 for failure conditions.
func ToInt8(d *mesCommons.Decimal, shift int32) (i int8, ok bool) {
	i64, ok := ToInt64(d, shift)
	if !ok || i64 < math.MinInt8 || i64 > math.MaxInt8 {
		return 0, false
	}
	return int8(i64), true
}

// FromInt16 returns Decimal representation for i.
func FromInt16(i int16) *mesCommons.Decimal {
	return &mesCommons.Decimal{Value: int64(i)}
}

// ToInt16 returns the int16 value for d*10^shift and a bool indicating
// whether the conversion succeeds. See ToInt64 for failure conditions.
func ToInt16(d *mesCommons.Decimal, shift int32) (i int16, ok bool) {
	i64, ok := ToInt64(d, shift)
	if !ok || i64 < math.MinInt16 || i64 > math.MaxInt16 {
		return 0, false
	}
	return int16(i64), true
}

// FromInt32 returns Decimal representation for i.
func FromInt32(i int32) *mesCommons.Decimal {
	return &mesCommons.Decimal{Value: int64(i)}
}

// ToInt32 returns the int32 value for d*10^shift and a bool indicating
// whether the conversion succeeds. See ToInt64 for failure conditions.
func ToInt32(d *mesCommons.Decimal, shift int32) (i int32, ok bool) {
	i64, ok := ToInt64(d, shift)
	if !ok || i64 < math.MinInt32 || i64 > math.MaxInt32 {
		return 0, false
	}
	return int32(i64), true
}

// FromInt64 returns Decimal representation for i.
func FromInt64(i int64) *mesCommons.Decimal {
	return &mesCommons.Decimal{Value: i}
}

// ToInt64 returns the int64 value for d*10^shift and a bool indicating
// whether the conversion succeeds.
//
// The conversion fails when
// - d.Exp + shift < 0 (even if d.Value ends with zeros), or
// - d*10^shift overflows
func ToInt64(d *mesCommons.Decimal, shift int32) (i int64, ok bool) {
	d.Exp += shift
	if d.Exp < 0 {
		return 0, false
	}
	if d.Exp == 0 {
		return d.Value, true
	}
	bigValue := decimal.New(d.Value, d.Exp).BigInt()
	if bigValue.IsInt64() {
		return bigValue.Int64(), true
	}
	return 0, false
}

// FromUint8 returns the uint8 value for d*10^shift and a bool indicating
// whether the conversion succeeds. See ToUint64 for failure conditions.
func FromUint8(i uint8) *mesCommons.Decimal {
	return &mesCommons.Decimal{Value: int64(i)}
}

// ToUint8 transforms Decimal into uint8.
func ToUint8(d *mesCommons.Decimal, shift int32) (i uint8, ok bool) {
	u64, ok := ToUint64(d, shift)
	if !ok || u64 > math.MaxUint8 {
		return 0, false
	}
	return uint8(u64), true
}

// FromUint16 returns the uint16 value for d*10^shift and a bool indicating
// whether the conversion succeeds. See ToUint64 for failure conditions.
func FromUint16(i uint16) *mesCommons.Decimal {
	return &mesCommons.Decimal{Value: int64(i)}
}

// ToUint16 transforms Decimal into uint16.
func ToUint16(d *mesCommons.Decimal, shift int32) (i uint16, ok bool) {
	u64, ok := ToUint64(d, shift)
	if !ok || u64 > math.MaxUint16 {
		return 0, false
	}
	return uint16(u64), true
}

// FromUint32 returns the uint32 value for d*10^shift and a bool indicating
// whether the conversion succeeds. See ToUint64 for failure conditions.
func FromUint32(i uint32) *mesCommons.Decimal {
	return &mesCommons.Decimal{Value: int64(i)}
}

// ToUint32 transforms Decimal into uint32.
func ToUint32(d *mesCommons.Decimal, shift int32) (i uint32, ok bool) {
	u64, ok := ToUint64(d, shift)
	if !ok || u64 > math.MaxUint32 {
		return 0, false
	}
	return uint32(u64), true
}

// FromUint64 returns Decimal representation for i and a bool indicating
// whether the conversion succeeds. The conversion fails when i exceeds
// MaxInt64 and as a result cannot be held by mesCommons.Decimal.
func FromUint64(i uint64) (d *mesCommons.Decimal, ok bool) {
	if i > math.MaxInt64 {
		return nil, false
	}
	return &mesCommons.Decimal{Value: int64(i)}, true
}

// ToUint64 returns the uint64 value for d*10^shift and a bool indicating
// whether the conversion succeeds.
//
// The conversion fails when
// - d.Value < 0,
// - d.Exp + shift < 0 (even if d.Value ends with zeros), or
// - d*10^shift overflows
func ToUint64(d *mesCommons.Decimal, shift int32) (i uint64, ok bool) {
	if d.Value < 0 {
		return 0, false
	}
	d.Exp += shift
	if d.Exp < 0 {
		return 0, false
	}
	if d.Exp == 0 {
		return uint64(d.Value), true
	}
	bigValue := decimal.New(d.Value, d.Exp).BigInt()
	if bigValue.IsUint64() {
		return bigValue.Uint64(), true
	}
	return 0, false
}

// FromDecimal returns Decimal representation for d.
func FromDecimal(d decimal.Decimal) (*mesCommons.Decimal, error) {
	if c := d.Coefficient(); c.IsInt64() {
		return &mesCommons.Decimal{
			Value: d.Coefficient().Int64(), // 取到浮點數的整數, e.g: f=0.005, value=5, exp=-3
			Exp:   d.Exponent(),
		}, nil
	}
	return nil, fmt.Errorf("the coefficient can't be represented as int64 value")
}

// MustFromDecimal returns Decimal representation for d, it panics if any error.
func MustFromDecimal(d decimal.Decimal) *mesCommons.Decimal {
	dec, err := FromDecimal(d)
	if err != nil {
		panic(err)
	}
	return dec
}

// ToDecimal returns decimal.Decimal for calculation.
func ToDecimal(d *mesCommons.Decimal) decimal.Decimal {
	return decimal.New(d.Value, d.Exp)
}

// IsZero returns true if decimal value is 0, otherwise returns false
func IsZero(d *mesCommons.Decimal) bool {
	return d.Value == 0
}
