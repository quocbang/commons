package excel

import (
	"errors"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/stretchr/testify/assert"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
)

func Test_fromEnumNamesToKeys(t *testing.T) {
	// good cases
	// enum names to keys
	{ // with skips
		names := rs.ProcessType_name
		skips := []int32{int32(rs.ProcessType_PROCESS_TYPE_UNSPECIFIED)}
		count := len(names) - len(skips)
		assert.Equal(t, len(fromEnumsToKeys(names, skips...)), count)
	}
	{ // without skips
		names := rs.ProcessType_name
		count := len(names) - 1
		assert.Equal(t, len(fromEnumsToKeys(names)), count)
	}
	// enum values to keys
	{ // with skips
		names := rs.ProcessType_name
		skips := []int32{int32(rs.ProcessType_PROCESS_TYPE_UNSPECIFIED)}
		count := len(names) - len(skips)
		assert.Equal(t, len(fromEnumsToKeys(names, skips...)), count)
	}
	{ // without skips
		names := rs.ProcessType_name
		count := len(names) - 1
		assert.Equal(t, len(fromEnumsToKeys(names)), count)
	}

	{ // nil names with no skips
		assert.Equal(t, []string{}, fromEnumsToKeys(nil))
	}
	{ // nil names with skips
		assert.Equal(t, []string{}, fromEnumsToKeys(nil, 1))
	}
}

func Test_setDropDownList(t *testing.T) {
	file := excelize.NewFile()

	// bad cases
	{ // nil file
		assert.Equal(t, errorNilFile, errors.Unwrap(setDropDownList(nil, "", "", []string{})))
	}
	{ // sheet not exist
		assert.Equal(t, errors.New("sheet not exist"), errors.Unwrap(setDropDownList(file, "", "", []string{})))
	}
	{ // missing SQRef
		assert.Equal(t, errors.New("missing axis"), errors.Unwrap(setDropDownList(file, sheetDefault, "", []string{})))
	}
	{ // missing keys
		assert.Equal(t, errors.New("missing keys"), errors.Unwrap(setDropDownList(file, sheetDefault, "A1", []string{})))
	}

	{ // normal case
		err := setDropDownList(file, sheetDefault, "A1", []string{"apple", "pineapple", "grapple", "dapple"})
		assert.NoError(t, err)
	}
}
