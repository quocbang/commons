package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewBuilder(t *testing.T) {
	{ // nil string
		sb := NewBuilder().String()
		assert.Equal(t, "", sb)
	}
	{ // normal
		sb := NewBuilder("hey, ", "it's ", "me!").String()
		assert.Equal(t, "hey, it's me!", sb)
	}
}

func TestBuilder_WriteStrings(t *testing.T) {
	{ // zero length
		b := NewBuilder()
		len, err := b.WriteStrings()
		assert.NoError(t, err)
		assert.Equal(t, 0, len)
	}
	{ // string length
		b := NewBuilder("hey, ", "it's ", "me!")
		len, err := b.WriteStrings("hey, ", "it's ", "me!")
		assert.NoError(t, err)
		assert.Equal(t, 13, len)
	}
}
