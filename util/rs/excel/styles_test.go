package excel

import (
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/stretchr/testify/assert"
)

func Test_cellsWithStyle(t *testing.T) {
	{
		expectedCells := []interface{}{
			excelize.Cell{StyleID: 5, Value: "TOOL"},
			excelize.Cell{StyleID: 5, Value: "ID"},
			excelize.Cell{StyleID: 5, Value: "ROLE"},
			excelize.Cell{StyleID: 5, Value: "CHECK"},
		}
		actualCells := cellsWithStyle(5, []string{"TOOL", "ID", "ROLE", "CHECK"}...)
		assert.Equal(t, expectedCells, actualCells)
	}
	{
		cells := cellsWithStyle(0, []string{}...)
		assert.Empty(t, cells)
	}
}

func Test_setDefaultStyles(t *testing.T) {
	{ // nil file
		_, err := setDefaultStyles(nil)
		assert.EqualError(t, err, errorNilFile.Error())
	}
	{
		file := excelize.NewFile()
		style, err := setDefaultStyles(file)
		assert.NoError(t, err)
		assert.Equal(t, &styles{highlight: 1, hyperlink: 2, warning: 3}, style)

	}
}
