package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterType(t *testing.T) {

	id := "configsID"
	type testingStruct struct{}

	assert := assert.New(t)

	// registering duplicate id
	{
		st := struct{}{}
		assert.NoError(RegisterType(id, reflect.TypeOf(st)))
		assert.Error(RegisterType(id, reflect.TypeOf(st)))
		resetList()
	}

	// available type
	{
		st := testingStruct{}
		assert.NoError(RegisterType(id, reflect.TypeOf(st)))
		resetList()
	}

	// NOT available type
	{
		// predeclared type
		{
			for _, v := range []interface{}{
				"", int8(0), int16(0), int32(0), int64(0), int(0),
				true, float32(0), float64(0),
			} {
				assert.Errorf(RegisterType(id, reflect.TypeOf(v)), "type %T", v)
				resetList()
			}
		}
		// pointer type
		{
			st := &testingStruct{}
			assert.Error(RegisterType(id, reflect.TypeOf(st)))
			resetList()
		}
	}
}

func TestGetConfigType(t *testing.T) {

	type testingStruct struct {
		Value string
	}
	id := "configsID"
	st := testingStruct{}

	// get the structure with a registry id
	{
		if err := RegisterType(id, reflect.TypeOf(st)); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, &testingStruct{}, GetConfigType(id))
		resetList()
	}

	// get the structure without a registry id
	{
		assert.Equal(t, &defaultStruct{}, GetConfigType("withoutRegistryID"))
	}

}

func resetList() {
	typeList = make(types)
}
