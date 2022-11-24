package rs

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
)

func TestPropertySettingNameToValue(t *testing.T) {
	// bad cases
	{ // empty parameter
		_, err := PropertySettingNameToValue("", "")
		assert.Equal(t, strings.HasPrefix(err.Error(), "bad property name="), true)
	}
	{ // empty value
		_, err := PropertySettingNameToValue(rs.PropertyName_RAM_POSITION.String(), "")
		assert.Equal(t, strings.HasPrefix(err.Error(), "property name=RAM_POSITION has bad property value="), true)
	}
	{ // bad value
		_, err := PropertySettingNameToValue(rs.PropertyName_RAM_POSITION.String(), "FLOAT")
		assert.Equal(t, strings.HasPrefix(err.Error(), "property name=RAM_POSITION has bad property value="), true)
	}

	// good cases
	{
		value, err := PropertySettingNameToValue(rs.PropertyName_RAM_POSITION.String(), "UP")
		assert.NoError(t, err)
		assert.Equal(t, int32(1), value)
	}
	{
		value, err := PropertySettingNameToValue(rs.PropertyName_RAM_POSITION.String(), "DOWN")
		assert.NoError(t, err)
		assert.Equal(t, int32(2), value)
	}
}

func TestPropertySettingValueToName(t *testing.T) {
	// bad cases
	{ // empty parameter
		_, err := PropertySettingValueToName("", 0)
		assert.Equal(t, strings.HasPrefix(err.Error(), "bad property name="), true)
	}
	{ // bad value
		_, err := PropertySettingValueToName(rs.PropertyName_RAM_POSITION.String(), 10)
		assert.Equal(t, strings.HasPrefix(err.Error(), "property name=RAM_POSITION has bad property value="), true)
	}
	// good cases
	{
		value, err := PropertySettingValueToName(rs.PropertyName_DROPPING_LOGIC.String(), 1)
		assert.NoError(t, err)
		assert.Equal(t, rs.DroppingLogic_TIME_ONLY.String(), value)
	}
}
