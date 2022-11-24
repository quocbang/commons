package validator

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBindJSON(t *testing.T) {

	New("validate")

	var s struct {
		Foo string `json:"foo" validate:"required"`
	}
	assert := assert.New(t)

	// testing BindJSON
	{
		// good case
		assert.NoError(BindJSON(bytes.NewReader([]byte(`{"foo": "FOO"}`)), &s))
	}
	{
		// missing request
		assert.Error(BindJSON(nil, &s))
	}

	// testing decodeJSON
	{
		// good case
		assert.NoError(decodeJSON(bytes.NewReader([]byte(`{"foo": "FOO"}`)), &s))
	}
	{
		// unknown field fooo
		assert.Error(decodeJSON(bytes.NewReader([]byte(`{"fooo": "FOO"}`)), &s))
	}

	// testing validate
	{
		// Foo on the 'required' tag
		assert.Error(decodeJSON(bytes.NewReader([]byte(`{"foo": ""}`)), &s))
	}
}
