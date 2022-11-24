package excel

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/stretchr/testify/assert"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
)

func normalSpec() *rs.Spec {
	return &rs.Spec{
		Id:          "2-79700-3-2",
		ProductId:   "79700-9",
		ProductType: "201",
		Factory:     []rs.FactoryID{rs.FactoryID_KY, rs.FactoryID_KU, rs.FactoryID_KC},
		Version: &rs.Version{
			Major:       "3",
			Minor:       "2",
			StageStatus: rs.StageStatus_PRODUCTION,
		},
		Materials:  normalMaterials(),
		Tools:      normalTools(),
		Properties: normalProperties(),
	}
}

func Test_writeSpecIndexSheet(t *testing.T) {
	expectedRows := func() [][]string {
		return [][]string{
			{"Spec ID", "2-79700-3-2"},
			{"Product ID", "79700-9"},
			{"Product Type", rs.ProductType_RUBBER.String()},
			{"Factories", rs.FactoryID_KY.String(), rs.FactoryID_KU.String(), rs.FactoryID_KC.String()},
			{"Major Version", "3"},
			{"Minor Version", "2"},
			{"Stage", rs.StageStatus_PRODUCTION.String()},
		}
	}

	{ // normal spec index
		file := excelize.NewFile()
		assert.NoError(t, writeSpecIndexSheet(file, rs.ProductType_RUBBER, normalSpec()))
		actualRows, err := file.GetRows(sheetSpecIndex)
		assert.NoError(t, err)
		assert.Equal(t, expectedRows(), actualRows)
	}
	{ // product type unspecified
		spec := normalSpec()
		spec.ProductType = ""

		file := excelize.NewFile()
		assert.NoError(t, writeSpecIndexSheet(file, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, spec))
		actualRows, err := file.GetRows(sheetSpecIndex)
		assert.NoError(t, err)

		rows := expectedRows()
		rows[2][1] = "" // set product type to empty
		assert.Equal(t, rows, actualRows)
	}
	{ // factory unspecified
		spec := normalSpec()
		spec.Factory = []rs.FactoryID{rs.FactoryID_FACTORY_UNSPECIFIED}

		file := excelize.NewFile()
		assert.NoError(t, writeSpecIndexSheet(file, rs.ProductType_RUBBER, spec))
		actualRows, err := file.GetRows(sheetSpecIndex)
		assert.NoError(t, err)

		rows := expectedRows()
		rows[3] = []string{"Factories", ""} // set factories to empty
		assert.Equal(t, rows, actualRows)
	}
	{ // stage status unspecified
		spec := normalSpec()
		spec.Version.StageStatus = rs.StageStatus_STAGE_STATUS_UNSPECIFIED

		file := excelize.NewFile()
		assert.NoError(t, writeSpecIndexSheet(file, rs.ProductType_RUBBER, spec))
		actualRows, err := file.GetRows(sheetSpecIndex)
		assert.NoError(t, err)

		rows := expectedRows()
		rows[6][1] = "" // set stage to empty
		assert.Equal(t, rows, actualRows)
	}
	{ // nil file
		assert.Equal(t, errorNilFile, errors.Unwrap(writeSpecIndexSheet(nil, rs.ProductType_RUBBER, normalSpec())))
	}
	{ // nil spec
		assert.Equal(t, errorNilSpec, errors.Unwrap(writeSpecIndexSheet(excelize.NewFile(), rs.ProductType_RUBBER, nil)))
	}
}

func Test_writeSpecMaterialsSection(t *testing.T) {
	// good cases
	{ // normal case
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		styles, err := setDefaultStyles(w.File)
		assert.NoError(t, err)
		row := 1
		assert.NoError(t, writeSpecMaterialsSection(w, &row, styles.warning, rs.ProductType_RUBBER, normalMaterials()))
		assert.NoError(t, w.Flush())

		expectedRows := [][]string{
			columnHeaderSpecMaterials,
			{"20003", "A", "20", "", "", "", "2", "", "kg"},
			{"20004", "A", "20", "", "", "1", "", "", "kg"},
			{"U79700N004-1-CMP-1", "", "202", "U79700N004-1-CMP1", "", "", "", "3", "kg"},
			{"60005", "A", "60", "", "", "1", "2", "3", "kg"},
		}
		actualRows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, expectedRows, actualRows)

		expectedStyle := [][]int{
			{styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
		styleEqual(t, w.File, sheetDefault, expectedStyle)
	}
	{ // nil material support
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		row := 1
		assert.NoError(t, writeSpecMaterialsSection(w, &row, 0, rs.ProductType_RUBBER, nil))
		assert.NoError(t, w.Flush())

		rows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, [][]string{columnHeaderSpecMaterials}, rows)
	}
	{ // empty material support
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		row := 1
		assert.NoError(t, writeSpecMaterialsSection(w, &row, 0, rs.ProductType_RUBBER, []*rs.Material{}))
		assert.NoError(t, w.Flush())

		rows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, [][]string{columnHeaderSpecMaterials}, rows)
	}

	// bad cases
	{ // nil stream writer
		assert.Equal(t, errorNilStreamWriter, errors.Unwrap(writeMaterialsSection(nil, nil, 0, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, nil)))
	}
	{ // nil row
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, errorNilRow, errors.Unwrap(writeMaterialsSection(w, nil, 0, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, nil)))
		assert.NoError(t, w.Flush())
	}
	{ // product type unspecified
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)
		row := 1
		assert.Equal(t, errorProductTypeUnspecified, errors.Unwrap(writeMaterialsSection(w, &row, 0, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, normalSteps())))
		assert.NoError(t, w.Flush())
	}
}

func Test_writeSpecPropertiesSection(t *testing.T) {
	// good cases
	{ // normal case
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		styles, err := setDefaultStyles(w.File)
		assert.NoError(t, err)
		row := 1
		assert.NoError(t, writeSpecPropertiesSection(w, &row, styles.warning, rs.ProductType_RUBBER, normalProperties()))
		assert.NoError(t, w.Flush())

		expectedRows := [][]string{
			columnHeaderProperties,
			{"property-central-value", "", "2", "", "", "kg"},
			{"property-min-value", "1", "", "", "", "kg"},
			{"property-max-value", "", "", "3", "", "kg"},
			{"property-all-value", "1", "2", "3", "", "kg"},
			{"RAM_POSITION", "", "", "", "UP", ""},
		}
		actualRows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, expectedRows, actualRows)

		expectedRowStyle := [][]int{
			{styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		}
		styleEqual(t, w.File, sheetDefault, expectedRowStyle)
	}
	{ // nil properties support
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		row := 1
		assert.NoError(t, writeSpecPropertiesSection(w, &row, 0, rs.ProductType_RUBBER, nil))
		assert.NoError(t, w.Flush())

		rows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, [][]string{columnHeaderProperties}, rows)
	}
	{ // empty properties support
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		row := 1
		assert.NoError(t, writeSpecPropertiesSection(w, &row, 0, rs.ProductType_RUBBER, []*rs.Property{}))
		assert.NoError(t, w.Flush())

		rows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, [][]string{columnHeaderProperties}, rows)
	}

	// bad cases
	{ // nil stream writer
		assert.Equal(t, errorNilStreamWriter, errors.Unwrap(writeSpecPropertiesSection(nil, nil, 0, rs.ProductType_RUBBER, nil)))
	}
	{ // nil row
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, errorNilRow, errors.Unwrap(writeSpecPropertiesSection(w, nil, 0, rs.ProductType_RUBBER, nil)))
		assert.NoError(t, w.Flush())
	}
	{ // product type unspecified
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)
		row := 1
		assert.Equal(t, errorProductTypeUnspecified, errors.Unwrap(writeSpecPropertiesSection(w, &row, 0, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, nil)))
		assert.NoError(t, w.Flush())
	}
	{ // bad enum value
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		row := 1
		spec := normalSpec()
		spec.Properties[4].Param.Enum.Value = 4
		assert.Equal(t, errors.New("property name=RAM_POSITION has bad property value=4"), errors.Unwrap(writeSpecPropertiesSection(w, &row, 0, rs.ProductType_RUBBER, spec.GetProperties())))
		assert.NoError(t, w.Flush())
	}
	{ // enum value unspecified
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		row := 1
		properties := normalProperties()
		properties[4].Param.Enum.Value = 0
		assert.NoError(t, writeSpecPropertiesSection(w, &row, 0, rs.ProductType_RUBBER, properties))
		assert.NoError(t, w.Flush())

		rows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, rows[5], []string{"RAM_POSITION", "", "", "", "", ""})
	}
}

func Test_writeSpecConfigsSheet(t *testing.T) {
	f := excelize.NewFile()
	styles, err := setDefaultStyles(f)
	assert.NoError(t, err)

	expectedRows := [][]string{
		columnHeaderTools,
		{"A13", rs.ToolRole_MOLD.String(), "TRUE"},
		{"U79700N004", rs.ToolRole_PAPER.String(), "FALSE"},
		nil,
		columnHeaderSpecMaterials,
		{"20003", "A", "20", "", "", "", "2", "", "kg"},
		{"20004", "A", "20", "", "", "1", "", "", "kg"},
		{"U79700N004-1-CMP-1", "", "202", "U79700N004-1-CMP1", "", "", "", "3", "kg"},
		{"60005", "A", "60", "", "", "1", "2", "3", "kg"},
		nil,
		columnHeaderProperties,
		{"property-central-value", "", "2", "", "", "kg"},
		{"property-min-value", "1", "", "", "", "kg"},
		{"property-max-value", "", "", "3", "", "kg"},
		{"property-all-value", "1", "2", "3", "", "kg"},
		{"RAM_POSITION", "", "", "", "UP", ""},
	}

	assert.NoError(t, writeSpecConfigsSheet(f, styles, rs.ProductType_RUBBER, normalSpec()))
	actualRows, err := f.GetRows(sheetSpecConfig)
	assert.NoError(t, err)
	assert.Equal(t, expectedRows, actualRows)
	expectedRowStyle := [][]int{
		{styles.warning, styles.warning, styles.warning},
		{0, 0, 0},
		{0, 0, 0},
		nil,
		{styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		nil,
		{styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}

	styleEqual(t, f, sheetSpecConfig, expectedRowStyle)

	// bad cases
	{ // nil file
		assert.Equal(t, errorNilFile, errors.Unwrap(writeSpecConfigsSheet(nil, nil, rs.ProductType_RUBBER, nil)))
	}
	{ // nil styles
		assert.Equal(t, errorNilStyles, errors.Unwrap(writeSpecConfigsSheet(excelize.NewFile(), nil, rs.ProductType_RUBBER, nil)))
	}
	{ // product type unspecified
		styles, err := setDefaultStyles(excelize.NewFile())
		assert.NoError(t, err)
		assert.Equal(t, errorProductTypeUnspecified, errors.Unwrap(writeSpecConfigsSheet(excelize.NewFile(), styles, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, nil)))
	}
	{ // nil spec
		styles, err := setDefaultStyles(excelize.NewFile())
		assert.NoError(t, err)
		assert.Equal(t, errorNilSpec, errors.Unwrap(writeSpecConfigsSheet(excelize.NewFile(), styles, rs.ProductType_RUBBER, nil)))
	}
}

func Test_FromSpec(t *testing.T) {
	// good case
	{
		readFile, err := ioutil.ReadFile("spec.json")
		assert.NoError(t, err)

		var spec = rs.Spec{}
		assert.NoError(t, json.Unmarshal(readFile, &spec))

		_, err = FromSpec(&spec)
		assert.NoError(t, err)
	}

	// bad cases
	{ // nil spec
		_, err := FromSpec(nil)
		assert.Equal(t, errorNilSpec, err)
	}
	{ // empty product type
		spec := normalSpec()
		spec.ProductType = ""
		_, err := FromSpec(spec)
		assert.Equal(t, strings.HasPrefix(err.Error(), "fail to convert product type"), true)
	}
	{ // product type unspecified
		spec := normalSpec()
		spec.ProductType = "0"
		_, err := FromSpec(spec)
		assert.Equal(t, errors.Unwrap(err), errorProductTypeUnspecified)
	}
}
