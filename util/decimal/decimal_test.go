package decimal

import (
	"math"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	mesCommons "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/mes/v2/commons"
)

func TestFromString(t *testing.T) {
	assert := assert.New(t)

	// valid number
	{
		dec, err := FromString("123.456")
		assert.NoError(err)
		assert.Equal(&mesCommons.Decimal{
			Value: 123456,
			Exp:   -3,
		}, dec)
	}
	// bad string
	{
		_, err := FromString("123.456a")
		assert.Error(err)
	}
}

func TestMustFromString(t *testing.T) {
	assert := assert.New(t)

	// valid number
	{
		dec := MustFromString("123.456")
		assert.Equal(&mesCommons.Decimal{
			Value: 123456,
			Exp:   -3,
		}, dec)
	}
	// bad string
	{
		assert.Panics(func() {
			MustFromString("123.456a")
		})
	}
}

func TestFromFloat(t *testing.T) {
	{
		dec, err := FromFloat(999999999.999999)
		assert.NoError(t, err)
		assert.Equal(t, &mesCommons.Decimal{Value: 999999999999999, Exp: -6}, dec)
	}
	{
		dec, err := FromFloat(0)
		assert.NoError(t, err)
		assert.Equal(t, &mesCommons.Decimal{}, dec)
	}
	{
		dec, err := FromFloat(-999999999.999999)
		assert.NoError(t, err)
		assert.Equal(t, &mesCommons.Decimal{Value: -999999999999999, Exp: -6}, dec)
	}
	{
		dec, err := FromFloat(.12345678)
		assert.NoError(t, err)
		assert.Equal(t, &mesCommons.Decimal{Value: 12345678, Exp: -8}, dec)
	}
	{
		dec, err := FromFloat(.00000001)
		assert.NoError(t, err)
		assert.Equal(t, &mesCommons.Decimal{Value: 1, Exp: -8}, dec)
	}
	{
		dec, err := FromFloat(12345678)
		assert.NoError(t, err)
		assert.Equal(t, &mesCommons.Decimal{Value: 12345678, Exp: 0}, dec)
	}
}

func TestMustFromFloat(t *testing.T) {
	assert := assert.New(t)

	// normal
	{
		assert.NotPanics(func() {
			MustFromFloat(math.MaxFloat64)
		})
	}
}

func TestToFloat(t *testing.T) {
	{
		f, exact := ToFloat(&mesCommons.Decimal{Value: 999999999999999, Exp: -6}, 0)
		if assert.False(t, exact) {
			assert.Equal(t, float64(999999999.999999), f)
		}
	}
	{
		f, exact := ToFloat(&mesCommons.Decimal{Value: -999999999999999, Exp: -6}, 0)
		if assert.False(t, exact) {
			assert.Equal(t, float64(-999999999.999999), f)
		}
	}
	{
		f, exact := ToFloat(&mesCommons.Decimal{Value: 999999999999999, Exp: 2}, -2)
		if assert.True(t, exact) {
			assert.Equal(t, float64(999999999999999), f)
		}
	}
	{
		f, exact := ToFloat(&mesCommons.Decimal{Value: -999999999999999, Exp: 2}, -2)
		if assert.True(t, exact) {
			assert.Equal(t, float64(-999999999999999), f)
		}
	}
	{
		f, exact := ToFloat(&mesCommons.Decimal{Value: 999999999999999, Exp: -4}, -2)
		if assert.False(t, exact) {
			assert.Equal(t, float64(999999999.999999), f)
		}
	}
	{
		f, exact := ToFloat(&mesCommons.Decimal{Value: -999999999999999, Exp: -4}, -2)
		if assert.False(t, exact) {
			assert.Equal(t, float64(-999999999.999999), f)
		}
	}
	{
		f, exact := ToFloat(&mesCommons.Decimal{Value: 0, Exp: 0}, 0)
		if assert.True(t, exact) {
			assert.Equal(t, float64(0), f)
		}
	}
	{
		f, exact := ToFloat(&mesCommons.Decimal{Value: 0, Exp: 0}, -6)
		if assert.True(t, exact) {
			assert.Equal(t, float64(0), f)
		}
	}
	{
		f, exact := ToFloat(&mesCommons.Decimal{Value: 1234567}, -4)
		if assert.False(t, exact) {
			assert.Equal(t, 123.4567, f)
		}
	}
	{
		f, exact := ToFloat(&mesCommons.Decimal{Value: 1}, -4)
		if assert.False(t, exact) {
			assert.Equal(t, .0001, f)
		}
	}
	{
		f, exact := ToFloat(&mesCommons.Decimal{Value: 9, Exp: 5}, -2)
		if assert.True(t, exact) {
			assert.Equal(t, float64(9000), f)
		}
	}
	{
		f, exact := ToFloat(&mesCommons.Decimal{Value: 1234, Exp: -2}, -2)
		if assert.False(t, exact) {
			assert.Equal(t, .1234, f)
		}
	}
	{
		f, exact := ToFloat(&mesCommons.Decimal{Value: 1, Exp: 3}, -4)
		if assert.False(t, exact) {
			assert.Equal(t, float64(0.1), f)
		}
	}
	{
		f, exact := ToFloat(&mesCommons.Decimal{Value: 25, Exp: 0}, -2)
		if assert.True(t, exact) {
			assert.Equal(t, float64(0.25), f)
		}
	}
}

func TestFromInt8(t *testing.T) {
	{
		dec := FromInt8(math.MaxInt8)
		assert.Equal(t, &mesCommons.Decimal{Value: math.MaxInt8}, dec)
	}
	{
		dec := FromInt8(math.MinInt8)
		assert.Equal(t, &mesCommons.Decimal{Value: math.MinInt8}, dec)
	}
	{
		dec := FromInt8(0)
		assert.Equal(t, &mesCommons.Decimal{}, dec)
	}
}

func TestToInt8(t *testing.T) {
	{
		i, ok := ToInt8(&mesCommons.Decimal{Value: math.MaxInt8}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int8(math.MaxInt8), i)
		}
	}
	{
		i, ok := ToInt8(&mesCommons.Decimal{Value: math.MinInt8}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int8(math.MinInt8), i)
		}
	}
	{
		i, ok := ToInt8(&mesCommons.Decimal{Value: 0}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int8(0), i)
		}
	}
	{
		i, ok := ToInt8(&mesCommons.Decimal{Value: 1}, 2)
		if assert.True(t, ok) {
			assert.Equal(t, int8(100), i)
		}
	}
	{
		i, ok := ToInt8(&mesCommons.Decimal{Value: 123}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int8(123), i)
		}
	}
	{
		i, ok := ToInt8(&mesCommons.Decimal{Value: -123}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int8(-123), i)
		}
	}
	{
		_, ok := ToInt8(&mesCommons.Decimal{Value: math.MaxInt8 + 1}, 0)
		assert.False(t, ok)
	}
	{
		_, ok := ToInt8(&mesCommons.Decimal{Value: math.MinInt8 - 1}, 0)
		assert.False(t, ok)
	}
	{
		_, ok := ToInt8(&mesCommons.Decimal{Value: -327, Exp: -1}, 0)
		assert.False(t, ok)
	}
}

func TestFromInt16(t *testing.T) {
	{
		dec := FromInt16(math.MaxInt16)
		assert.Equal(t, &mesCommons.Decimal{Value: math.MaxInt16}, dec)
	}
	{
		dec := FromInt16(math.MinInt16)
		assert.Equal(t, &mesCommons.Decimal{Value: math.MinInt16}, dec)
	}
	{
		dec := FromInt16(0)
		assert.Equal(t, &mesCommons.Decimal{}, dec)
	}
}

func TestToInt16(t *testing.T) {
	{
		i, ok := ToInt16(&mesCommons.Decimal{Value: math.MaxInt16}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int16(math.MaxInt16), i)
		}
	}
	{
		i, ok := ToInt16(&mesCommons.Decimal{Value: math.MinInt16}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int16(math.MinInt16), i)
		}
	}
	{
		i, ok := ToInt16(&mesCommons.Decimal{Value: 0}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int16(0), i)
		}
	}
	{
		i, ok := ToInt16(&mesCommons.Decimal{Value: 12}, 3)
		if assert.True(t, ok) {
			assert.Equal(t, int16(12000), i)
		}
	}
	{
		i, ok := ToInt16(&mesCommons.Decimal{Value: 1234}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int16(1234), i)
		}
	}
	{
		i, ok := ToInt16(&mesCommons.Decimal{Value: -1234}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int16(-1234), i)
		}
	}
	{
		_, ok := ToInt16(&mesCommons.Decimal{Value: math.MaxInt16 + 1}, 0)
		assert.False(t, ok)
	}
	{
		_, ok := ToInt16(&mesCommons.Decimal{Value: math.MinInt16 - 1}, 0)
		assert.False(t, ok)
	}
	{
		_, ok := ToInt16(&mesCommons.Decimal{Value: -327, Exp: -1}, 0)
		assert.False(t, ok)
	}
}

func TestFromInt32(t *testing.T) {
	{
		dec := FromInt32(math.MaxInt32)
		assert.Equal(t, &mesCommons.Decimal{Value: math.MaxInt32}, dec)
	}
	{
		dec := FromInt32(math.MinInt32)
		assert.Equal(t, &mesCommons.Decimal{Value: math.MinInt32}, dec)
	}
	{
		dec := FromInt32(0)
		assert.Equal(t, &mesCommons.Decimal{}, dec)
	}
}

func TestToInt32(t *testing.T) {
	{
		i, ok := ToInt32(&mesCommons.Decimal{Value: math.MaxInt32}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int32(math.MaxInt32), i)
		}
	}
	{
		i, ok := ToInt32(&mesCommons.Decimal{Value: math.MinInt32}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int32(math.MinInt32), i)
		}
	}
	{
		i, ok := ToInt32(&mesCommons.Decimal{Value: 0}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int32(0), i)
		}
	}
	{
		i, ok := ToInt32(&mesCommons.Decimal{Value: 123}, 3)
		if assert.True(t, ok) {
			assert.Equal(t, int32(123000), i)
		}
	}
	{
		i, ok := ToInt32(&mesCommons.Decimal{Value: 1234567890}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int32(1234567890), i)
		}
	}
	{
		i, ok := ToInt32(&mesCommons.Decimal{Value: -1234567890}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int32(-1234567890), i)
		}
	}
	{
		_, ok := ToInt32(&mesCommons.Decimal{Value: math.MaxInt32 + 1}, 0)
		assert.False(t, ok)
	}
	{
		_, ok := ToInt32(&mesCommons.Decimal{Value: math.MinInt32 - 1}, 0)
		assert.False(t, ok)
	}
	{
		_, ok := ToInt32(&mesCommons.Decimal{Value: -3271234056, Exp: -1}, 0)
		assert.False(t, ok)
	}
}

func TestFromInt64(t *testing.T) {
	{
		dec := FromInt64(math.MaxInt64)
		assert.Equal(t, &mesCommons.Decimal{Value: math.MaxInt64}, dec)
	}
	{
		dec := FromInt64(math.MinInt64)
		assert.Equal(t, &mesCommons.Decimal{Value: math.MinInt64}, dec)
	}
	{
		dec := FromInt64(0)
		assert.Equal(t, &mesCommons.Decimal{}, dec)
	}
}

func TestToInt64(t *testing.T) {
	{
		i, ok := ToInt64(&mesCommons.Decimal{Value: math.MaxInt64}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int64(math.MaxInt64), i)
		}
	}
	{
		i, ok := ToInt64(&mesCommons.Decimal{Value: math.MinInt64}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int64(math.MinInt64), i)
		}
	}
	{
		i, ok := ToInt64(&mesCommons.Decimal{Value: 0}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int64(0), i)
		}
	}
	{
		i, ok := ToInt64(&mesCommons.Decimal{Value: 123}, 6)
		if assert.True(t, ok) {
			assert.Equal(t, int64(123000000), i)
		}
	}
	{
		i, ok := ToInt64(&mesCommons.Decimal{Value: 1234567890123456789}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int64(1234567890123456789), i)
		}
	}
	{
		i, ok := ToInt64(&mesCommons.Decimal{Value: -1234567890123456789}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int64(-1234567890123456789), i)
		}
	}
	{
		i, ok := ToInt64(&mesCommons.Decimal{Value: 123456789, Exp: 1}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, int64(1234567890), i)
		}
	}
	{
		_, ok := ToInt64(&mesCommons.Decimal{Value: math.MaxInt64}, 1)
		assert.False(t, ok)
	}
	{
		_, ok := ToInt64(&mesCommons.Decimal{Value: math.MinInt64}, 1)
		assert.False(t, ok)
	}
	{
		_, ok := ToInt64(&mesCommons.Decimal{Value: 1234567890123456789, Exp: -1}, 0)
		assert.False(t, ok)
	}
	{
		_, ok := ToInt64(&mesCommons.Decimal{Value: -1234567890123456789, Exp: -1}, 0)
		assert.False(t, ok)
	}
}

func TestFromUint8(t *testing.T) {
	{
		dec := FromUint8(math.MaxUint8)
		assert.Equal(t, &mesCommons.Decimal{Value: math.MaxUint8}, dec)
	}
	{
		dec := FromUint8(0)
		assert.Equal(t, &mesCommons.Decimal{Value: 0}, dec)
	}
}

func TestToUint8(t *testing.T) {
	{
		i, ok := ToUint8(&mesCommons.Decimal{Value: math.MaxUint8}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, uint8(math.MaxUint8), i)
		}
	}
	{
		i, ok := ToUint8(&mesCommons.Decimal{Value: 10}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, uint8(10), i)
		}
	}
	{
		i, ok := ToUint8(&mesCommons.Decimal{}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, uint8(0), i)
		}
	}
	{
		_, ok := ToUint8(&mesCommons.Decimal{Value: -1}, 0)
		assert.False(t, ok)
	}
	{
		_, ok := ToUint8(&mesCommons.Decimal{Value: 12, Exp: -1}, 0)
		assert.False(t, ok)
	}
	{
		_, ok := ToUint8(&mesCommons.Decimal{Value: math.MaxUint8 + 1}, 0)
		assert.False(t, ok)
	}
}

func TestFromUint16(t *testing.T) {
	{
		dec := FromUint16(math.MaxUint16)
		assert.Equal(t, &mesCommons.Decimal{Value: math.MaxUint16}, dec)
	}
	{
		dec := FromUint16(0)
		assert.Equal(t, &mesCommons.Decimal{Value: 0}, dec)
	}
}

func TestToUint16(t *testing.T) {
	{
		i, ok := ToUint16(&mesCommons.Decimal{Value: math.MaxUint16}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, uint16(math.MaxUint16), i)
		}
	}
	{
		i, ok := ToUint16(&mesCommons.Decimal{Value: 10000}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, uint16(10000), i)
		}
	}
	{
		i, ok := ToUint16(&mesCommons.Decimal{}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, uint16(0), i)
		}
	}
	{
		_, ok := ToUint16(&mesCommons.Decimal{Value: -1}, 0)
		assert.False(t, ok)
	}
	{
		_, ok := ToUint16(&mesCommons.Decimal{Value: 1234, Exp: -1}, 0)
		assert.False(t, ok)
	}
	{
		_, ok := ToUint16(&mesCommons.Decimal{Value: math.MaxUint16 + 1}, 0)
		assert.False(t, ok)
	}
}

func TestFromUint32(t *testing.T) {
	{
		dec := FromUint32(math.MaxUint32)
		assert.Equal(t, &mesCommons.Decimal{Value: math.MaxUint32}, dec)
	}
	{
		dec := FromUint32(0)
		assert.Equal(t, &mesCommons.Decimal{Value: 0}, dec)
	}
}

func TestToUint32(t *testing.T) {
	{
		i, ok := ToUint32(&mesCommons.Decimal{Value: math.MaxUint32}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, uint32(math.MaxUint32), i)
		}
	}
	{
		i, ok := ToUint32(&mesCommons.Decimal{Value: 429496729}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, uint32(429496729), i)
		}
	}
	{
		i, ok := ToUint32(&mesCommons.Decimal{}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, uint32(0), i)
		}
	}
	{
		_, ok := ToUint32(&mesCommons.Decimal{Value: -1}, 0)
		assert.False(t, ok)
	}
	{
		_, ok := ToUint32(&mesCommons.Decimal{Value: 4294967295, Exp: -1}, 0)
		assert.False(t, ok)
	}
	{
		_, ok := ToUint32(&mesCommons.Decimal{Value: math.MaxUint32 + 1}, 0)
		assert.False(t, ok)
	}
}

func TestFromUint64(t *testing.T) {
	{
		_, ok := FromUint64(math.MaxUint64)
		assert.False(t, ok)
	}
	{
		dec, ok := FromUint64(0)
		if assert.True(t, ok) {
			assert.Equal(t, &mesCommons.Decimal{}, dec)
		}
	}
	{
		dec, ok := FromUint64(math.MaxInt64)
		if assert.True(t, ok) {
			assert.Equal(t, &mesCommons.Decimal{Value: math.MaxInt64}, dec)
		}
	}
	{
		_, ok := FromUint64(math.MaxInt64 + 1)
		assert.False(t, ok)
	}
}

func TestToUint64(t *testing.T) {
	{
		i, ok := ToUint64(&mesCommons.Decimal{Value: math.MaxInt64}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, uint64(math.MaxInt64), i)
		}
	}
	{
		i, ok := ToUint64(&mesCommons.Decimal{Value: 42949672912345}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, uint64(42949672912345), i)
		}
	}
	{
		i, ok := ToUint64(&mesCommons.Decimal{}, 0)
		if assert.True(t, ok) {
			assert.Equal(t, uint64(0), i)
		}
	}
	{
		i, ok := ToUint64(&mesCommons.Decimal{Value: 1844674407370955161}, 1)
		if assert.True(t, ok) {
			assert.Equal(t, uint64(18446744073709551610), i)
		}
	}
	{
		_, ok := ToUint64(&mesCommons.Decimal{Value: 1844674407370955161}, 2)
		assert.False(t, ok)
	}
	{
		_, ok := ToUint64(&mesCommons.Decimal{Value: -1}, 0)
		assert.False(t, ok)
	}
	{
		_, ok := ToUint64(&mesCommons.Decimal{Value: 4294967295, Exp: -1}, 0)
		assert.False(t, ok)
	}
}

func TestFromDecimal(t *testing.T) {
	assert := assert.New(t)

	// normal
	{
		dec, err := FromDecimal(decimal.New(math.MaxInt64, 2))
		assert.NoError(err)
		assert.Equal(&mesCommons.Decimal{
			Value: math.MaxInt64,
			Exp:   2,
		}, dec)
	}
	// conversion failed
	{
		dec, err := FromDecimal(decimal.New(math.MaxInt64, 0).Add(decimal.New(1, 0)))
		assert.Error(err)
		assert.Empty(dec)
	}
}

func TestMustFromDecimal(t *testing.T) {
	assert := assert.New(t)

	// normal
	{
		assert.NotPanics(func() {
			MustFromDecimal(decimal.New(30, 0))
		})
	}
	// panic error
	{
		assert.Panics(func() {
			MustFromDecimal(decimal.New(math.MaxInt64, 0).Add(decimal.New(7, 0)))
		})
	}
}
