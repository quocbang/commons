package metadata

import (
	// "errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SetTimeZone(t *testing.T) {
	{ // KV
		m := &Metadata{}
		name := "Asia/Ho_Chi_Minh"
		assert.NoError(t, m.SetTimeZone(name))
		assert.Equal(t, name, m.TimeZone)
	}
	{ // KA (Ohio)
		m := &Metadata{}
		name := "America/New_York"
		assert.NoError(t, m.SetTimeZone(name))
		assert.Equal(t, name, m.TimeZone)
	}

	// bad cases
	{ // empty location name
		m := &Metadata{}
		assert.Error(t, m.SetTimeZone(""))
	}
	{ // bad location name
		m := &Metadata{}
		assert.Error(t, m.SetTimeZone("/"))
	}
	{ // bad location name
		m := &Metadata{}
		assert.Error(t, m.SetTimeZone("/Taipei"))
	}
	{ // bad location name
		m := &Metadata{}
		assert.Error(t, m.SetTimeZone("Asia/"))
	}
}
