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
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs/container"
)

func normalTools() []*rs.Tool {
	return []*rs.Tool{
		{
			Role:      rs.ToolRole_MOLD,
			Id:        "A13",
			CheckFlag: true,
		},
		{
			Role:      rs.ToolRole_PAPER,
			Id:        "U79700N004",
			CheckFlag: false,
		},
	}
}

func normalCentralValueParam() *rs.Param {
	return &rs.Param{
		Type:    rs.Param_VALUE,
		Unit:    "kg",
		Central: &rs.Param_Constraint{Value: 2.00},
	}
}

func normalMinValueParam() *rs.Param {
	return &rs.Param{
		Type: rs.Param_MIN,
		Unit: "kg",
		Min:  &rs.Param_Constraint{Value: 1.00},
	}
}

func normalMaxValueParam() *rs.Param {
	return &rs.Param{
		Type: rs.Param_MAX,
		Unit: "kg",
		Max:  &rs.Param_Constraint{Value: 3.00},
	}
}

func normalAllValueParam() *rs.Param {
	return &rs.Param{
		Type:    rs.Param_RANGE,
		Unit:    "kg",
		Central: &rs.Param_Constraint{Value: 2.00},
		Min:     &rs.Param_Constraint{Value: 1.00},
		Max:     &rs.Param_Constraint{Value: 3.00},
	}
}

func normalEnumParam() *rs.Param {
	return &rs.Param{
		Type: rs.Param_ENUM,
		Unit: "",
		Enum: &rs.Param_Enum{Value: 1},
	}
}

func normalMaterials() []*rs.Material {
	return []*rs.Material{
		{
			Type:          "20",
			Level:         "A",
			Name:          "20003",
			ContainerType: 0,
			Param:         normalCentralValueParam(),
			ReqRecipeId:   "",
		},
		{
			Type:          "20",
			Level:         "A",
			Name:          "20004",
			ContainerType: 0,
			Param:         normalMinValueParam(),
			ReqRecipeId:   "",
		},
		{
			Type:          "202",
			Level:         "",
			Name:          "U79700N004-1-CMP-1",
			ContainerType: 0,
			Param:         normalMaxValueParam(),
			ReqRecipeId:   "U79700N004-1-CMP1",
		},
		{
			Type:          "60",
			Level:         "A",
			Name:          "60005",
			ContainerType: 0,
			Param:         normalAllValueParam(),
			ReqRecipeId:   "",
		},
	}
}

func normalProperties() []*rs.Property {
	return []*rs.Property{
		{
			Name:        "property-central-value",
			Param:       normalCentralValueParam(),
			Description: "",
		},
		{
			Name:        "property-min-value",
			Param:       normalMinValueParam(),
			Description: "",
		},
		{
			Name:        "property-max-value",
			Param:       normalMaxValueParam(),
			Description: "",
		},
		{
			Name:        "property-all-value",
			Param:       normalAllValueParam(),
			Description: "",
		},
		{
			Name:        "RAM_POSITION",
			Param:       normalEnumParam(),
			Description: "",
		},
	}
}

func normalSteps() []*rs.Process_Config_Step {
	return []*rs.Process_Config_Step{
		{
			Materials:    normalMaterials(),
			Controls:     normalProperties(),
			Measurements: normalProperties(),
		},
	}
}

func normalConfigs() []*rs.Process_Config {
	return []*rs.Process_Config{
		{
			Stations: []string{"station-1", "station-2", "station-3"},
			Tools:    normalTools(),
			Steps:    normalSteps(),
		},
	}
}

func normalProcesses() []*rs.Process {
	return []*rs.Process{
		{
			Name:    "79700-1",
			Type:    rs.ProcessType_PROCESS,
			Configs: normalConfigs(),
		},
		{
			Name:    "79700-1",
			Type:    rs.ProcessType_INSPECTION,
			Configs: normalConfigs(),
		},
	}
}

func normalOptionalFlows() []*rs.Recipe_OptionalFlow {
	return []*rs.Recipe_OptionalFlow{
		{
			Name:             "optional-flow-1",
			AfterProcessName: "process-1",
			AfterProcessType: rs.ProcessType_PROCESS,
			MaxRepetitions:   10,
			Processes:        normalProcesses(),
		},
	}
}

func normalRecipe() *rs.Recipe {
	return &rs.Recipe{
		Id:          "U79700N004",
		ProductType: "201",
		ProductId:   "79700-9",
		Factory:     []rs.FactoryID{rs.FactoryID_KY, rs.FactoryID_KU},
		Version: &rs.Version{
			Major:       "004",
			Minor:       "",
			StageStatus: rs.StageStatus_PRODUCTION,
		},
		ParentSpecId:  "2-79700-3-2",
		Processes:     normalProcesses(),
		OptionalFlows: normalOptionalFlows(),
	}
}

var styleEqual = func(t *testing.T, file *excelize.File, sheet string, expected [][]int) {
	rows, err := file.GetRows(sheet)
	assert.NoError(t, err)
	assert.Equal(t, len(expected), len(rows))

	actual := make([][]int, len(rows))
	for rowIdx, row := range rows {
		assert.Equal(t, len(row), len(expected[rowIdx]))
		for cellIdx := range row {
			cellName, err := excelize.CoordinatesToCellName(cellIdx+1, rowIdx+1)
			assert.NoError(t, err)
			style, err := file.GetCellStyle(sheet, cellName)
			assert.NoError(t, err)
			actual[rowIdx] = append(actual[rowIdx], style)
		}
	}
	assert.Equal(t, expected, actual)
}

func Test_productTypeRow(t *testing.T) {
	{
		expectedRow := []interface{}{"Product Type", rs.ProductType_RUBBER.String()}

		actualRow := productTypeRow(rs.ProductType_RUBBER)
		assert.Equal(t, expectedRow, actualRow)
	}
	{ // unspecified
		expectedRow := []interface{}{"Product Type", nil}

		actualRow := productTypeRow(rs.ProductType_PRODUCT_TYPE_UNSPECIFIED)
		assert.Equal(t, expectedRow, actualRow)
	}
}

func Test_factoriesRow(t *testing.T) {
	{
		expectedFactories := []interface{}{"Factories", rs.FactoryID_KY.String(), rs.FactoryID_KU.String()}

		actualFactories := factoriesRow(normalRecipe().GetFactory())
		assert.Equal(t, expectedFactories, actualFactories)
	}
	{ // unspecified
		expectedFactories := []interface{}{"Factories", nil}

		actualFactories := factoriesRow([]rs.FactoryID{rs.FactoryID_FACTORY_UNSPECIFIED})
		assert.Equal(t, expectedFactories, actualFactories)
	}
}

func Test_stageStatusRow(t *testing.T) {
	{
		expectedRow := []interface{}{"Stage", rs.StageStatus_PRODUCTION.String()}
		actualRow := stageStatusRow(normalRecipe().GetVersion().GetStageStatus())
		assert.Equal(t, expectedRow, actualRow)
	}
	{ // unspecified
		expectedRow := []interface{}{"Stage", nil}

		actualRow := stageStatusRow(rs.StageStatus_STAGE_STATUS_UNSPECIFIED)
		assert.Equal(t, expectedRow, actualRow)
	}
}

func Test_writeRecipeIndexSheet(t *testing.T) {
	expectedRows := func() [][]string {
		return [][]string{
			{"Recipe ID", "U79700N004"},
			{"Product ID", "79700-9"},
			{"Product Type", rs.ProductType_RUBBER.String()},
			{"Factories", rs.FactoryID_KY.String(), rs.FactoryID_KU.String()},
			{"Major Version", "004"},
			{"Minor Version", ""},
			{"Parent Spec ID", "2-79700-3-2"},
			{"Stage", rs.StageStatus_PRODUCTION.String()},
		}
	}

	{ // normal recipe index
		file := excelize.NewFile()
		assert.NoError(t, writeRecipeIndexSheet(file, rs.ProductType_RUBBER, normalRecipe()))
		actualRows, err := file.GetRows(sheetRecipeIndex)
		assert.NoError(t, err)
		assert.Equal(t, expectedRows(), actualRows)
	}
	{ // product type unspecified

		recipe := normalRecipe()
		recipe.ProductType = ""

		file := excelize.NewFile()
		assert.NoError(t, writeRecipeIndexSheet(file, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, recipe))
		actualRows, err := file.GetRows(sheetRecipeIndex)
		assert.NoError(t, err)

		rows := expectedRows()
		rows[2][1] = "" // set product type to empty
		assert.Equal(t, rows, actualRows)
	}
	{ // factory unspecified
		recipe := normalRecipe()
		recipe.Factory = []rs.FactoryID{rs.FactoryID_FACTORY_UNSPECIFIED}

		file := excelize.NewFile()
		assert.NoError(t, writeRecipeIndexSheet(file, rs.ProductType_RUBBER, recipe))
		actualRows, err := file.GetRows(sheetRecipeIndex)
		assert.NoError(t, err)

		rows := expectedRows()
		rows[3] = []string{"Factories", ""} // set factories to empty
		assert.Equal(t, rows, actualRows)
	}
	{ // stage status unspecified
		recipe := normalRecipe()
		recipe.Version.StageStatus = rs.StageStatus_STAGE_STATUS_UNSPECIFIED

		file := excelize.NewFile()
		assert.NoError(t, writeRecipeIndexSheet(file, rs.ProductType_RUBBER, recipe))
		actualRows, err := file.GetRows(sheetRecipeIndex)
		assert.NoError(t, err)

		rows := expectedRows()
		rows[7][1] = "" // set stage to empty
		assert.Equal(t, rows, actualRows)
	}
	{ // nil file
		assert.Equal(t, errorNilFile, errors.Unwrap(writeRecipeIndexSheet(nil, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, nil)))
	}
	{ // nil recipe
		assert.Equal(t, errorNilRecipe, errors.Unwrap(writeRecipeIndexSheet(excelize.NewFile(), rs.ProductType_RUBBER, nil)))
	}
}

func Test_processSheetName(t *testing.T) {
	// normal process type
	{
		name, err := processSheetName(rs.ProcessType_PROCESS, "79700-1")
		assert.NoError(t, err)
		assert.Equal(t, "P-79700-1", name)
	}
	{
		name, err := processSheetName(rs.ProcessType_INSPECTION, "79700-1")
		assert.NoError(t, err)
		assert.Equal(t, "I-79700-1", name)
	}
	{
		name, err := processSheetName(rs.ProcessType_MEASUREMENT, "79700-1")
		assert.NoError(t, err)
		assert.Equal(t, "M-79700-1", name)
	}
	// unspecified
	{
		_, err := processSheetName(rs.ProcessType_PROCESS_TYPE_UNSPECIFIED, "79700-1")
		assert.Equal(t, errors.New("failed to get sheet name: process type=PROCESS_TYPE_UNSPECIFIED, name=79700-1 "), err)
	}
	{ // empty name
		_, err := processSheetName(rs.ProcessType_PROCESS, "")
		assert.Equal(t, errors.New("failed to get sheet name: process type=PROCESS, name= "), err)
	}
}

func Test_stationsRow(t *testing.T) {
	{
		expectedRows := []interface{}{
			excelize.Cell{StyleID: 0, Value: "STATIONS"},
			excelize.Cell{StyleID: 0, Value: "station-1"},
			excelize.Cell{StyleID: 0, Value: "station-2"},
			excelize.Cell{StyleID: 0, Value: "station-3"},
		}
		assert.Equal(t, expectedRows, stationsRow(0, normalRecipe().GetProcesses()[0].GetConfigs()[0].GetStations()))
	}
	{ // nil stations
		assert.Equal(t, []interface{}{excelize.Cell{StyleID: 0, Value: "STATIONS"}}, stationsRow(0, nil))
	}
	{ // empty stations
		assert.Equal(t, []interface{}{excelize.Cell{StyleID: 0, Value: "STATIONS"}}, stationsRow(0, []string{}))
	}
}

func Test_upperBool(t *testing.T) {
	{
		assert.Equal(t, "TRUE", upperBool(true))
	}
	{
		assert.Equal(t, "FALSE", upperBool(false))
	}
}

func Test_writeToolsSection(t *testing.T) {
	// good cases
	{ // normal case
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		styles, err := setDefaultStyles(w.File)
		assert.NoError(t, err)
		row := 1
		assert.NoError(t, writeToolsSection(w, &row, styles.warning, normalTools()))
		assert.NoError(t, w.Flush())

		expectedRows := [][]string{
			columnHeaderTools,
			{"A13", rs.ToolRole_MOLD.String(), "TRUE"},
			{"U79700N004", rs.ToolRole_PAPER.String(), "FALSE"},
		}
		actualRows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, expectedRows, actualRows)

		expectedRowStyle := [][]int{
			{styles.warning, styles.warning, styles.warning},
			{0, 0, 0},
			{0, 0, 0},
		}
		styleEqual(t, w.File, sheetDefault, expectedRowStyle)
	}
	{ // nil tools support
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		row := 1
		assert.NoError(t, writeToolsSection(w, &row, 0, nil))
		assert.NoError(t, w.Flush())

		rows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, [][]string{columnHeaderTools}, rows)
	}
	{ // empty tools support
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		row := 1
		assert.NoError(t, writeToolsSection(w, &row, 0, []*rs.Tool{}))
		assert.NoError(t, w.Flush())

		rows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, [][]string{columnHeaderTools}, rows)
	}

	// bad cases
	{ // nil stream writer
		assert.Equal(t, errorNilStreamWriter, errors.Unwrap(writeToolsSection(nil, nil, 0, nil)))
	}
	{ // nil row
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, errorNilRow, errors.Unwrap(writeToolsSection(w, nil, 0, nil)))
		assert.NoError(t, w.Flush())
	}
}

func Test_containerToString(t *testing.T) {
	{ // unspecified
		assert.Equal(t, "", containerToString(container.Type_CONTAINER_TYPE_UNSPECIFIED))
	}
	{ // normal case
		assert.Equal(t, container.Type__233T.String(), containerToString(container.Type__233T))
	}
}

func Test_writeMaterialsSection(t *testing.T) {
	// good cases
	{ // normal case
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		styles, err := setDefaultStyles(w.File)
		assert.NoError(t, err)
		row := 1
		assert.NoError(t, writeMaterialsSection(w, &row, styles.warning, rs.ProductType_RUBBER, normalSteps()))
		assert.NoError(t, w.Flush())

		expectedRows := [][]string{
			columnHeaderRecipeMaterials,
			{"1", "20003", "A", "20", "", "", "", "2", "", "kg"},
			{"1", "20004", "A", "20", "", "", "1", "", "", "kg"},
			{"1", "U79700N004-1-CMP-1", "", "202", "U79700N004-1-CMP1", "", "", "", "3", "kg"},
			{"1", "60005", "A", "60", "", "", "1", "2", "3", "kg"},
		}
		actualRows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, expectedRows, actualRows)

		expectedStyle := [][]int{
			{styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
		styleEqual(t, w.File, sheetDefault, expectedStyle)
	}
	{ // nil steps support
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		row := 1
		assert.NoError(t, writeMaterialsSection(w, &row, 0, rs.ProductType_RUBBER, nil))
		assert.NoError(t, w.Flush())

		rows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, [][]string{columnHeaderRecipeMaterials}, rows)
	}
	{ // empty steps support
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		row := 1
		assert.NoError(t, writeMaterialsSection(w, &row, 0, rs.ProductType_RUBBER, []*rs.Process_Config_Step{}))
		assert.NoError(t, w.Flush())

		rows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, [][]string{columnHeaderRecipeMaterials}, rows)
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

func Test_writeControlsSection(t *testing.T) {
	// good cases
	{ // normal case
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		styles, err := setDefaultStyles(w.File)
		assert.NoError(t, err)
		row := 1
		assert.NoError(t, writeControlsSection(w, &row, styles.warning, normalSteps()))
		assert.NoError(t, w.Flush())

		expectedRows := [][]string{
			{"CONTROL", "property-central-value", "property-min-value", "property-max-value", "property-all-value", rs.PropertyName_RAM_POSITION.String()},
			{"UNIT", "kg", "kg", "kg", "kg", ""},
			{"1", "mid:2", "low:1", "high:3", "low:1;mid:2;high:3", rs.RamPosition_UP.String()},
		}
		actualRows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, expectedRows, actualRows)

		expectedRowStyle := [][]int{
			{styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning},
			{styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning},
			{0, 0, 0, 0, 0, 0},
		}

		styleEqual(t, w.File, sheetDefault, expectedRowStyle)
	}

	// bad cases
	{ // nil stream writer
		assert.Equal(t, errorNilStreamWriter, errors.Unwrap(writeControlsSection(nil, nil, 0, nil)))
	}
	{ // nil row
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, errorNilRow, errors.Unwrap(writeControlsSection(w, nil, 0, nil)))
		assert.NoError(t, w.Flush())
	}
	{ // nil steps
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)
		row := 1
		assert.Equal(t, errorEmptySteps, errors.Unwrap(writeControlsSection(w, &row, 0, nil)))
		assert.NoError(t, w.Flush())
	}
	{ // empty steps
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)
		row := 1
		assert.Equal(t, errorEmptySteps, errors.Unwrap(writeControlsSection(w, &row, 0, []*rs.Process_Config_Step{})))
		assert.NoError(t, w.Flush())
	}
	{ // bad enum value
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		row := 1
		config := normalConfigs()
		config[0].Steps[0].Controls[4].Param.Enum.Value = -1
		assert.Equal(t, errors.New("property name=RAM_POSITION has bad property value=-1"), errors.Unwrap(writeControlsSection(w, &row, 0, config[0].Steps)))
		assert.NoError(t, w.Flush())
	}
}

func Test_writeMeasurementsSection(t *testing.T) {
	// good cases
	{ // normal case
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		styles, err := setDefaultStyles(w.File)
		assert.NoError(t, err)
		row := 1
		assert.NoError(t, writeMeasurementsSection(w, &row, styles.warning, normalSteps()))
		assert.NoError(t, w.Flush())

		expectedRows := [][]string{
			columnHeaderMeasurements,
			{"1", "property-central-value", "", "2", "", "kg"},
			{"1", "property-min-value", "1", "", "", "kg"},
			{"1", "property-max-value", "", "", "3", "kg"},
			{"1", "property-all-value", "1", "2", "3", "kg"},
			{"1", "RAM_POSITION", "", "", "", ""},
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
	{ // nil steps support
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		row := 1
		assert.NoError(t, writeMeasurementsSection(w, &row, 0, nil))
		assert.NoError(t, w.Flush())

		rows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, [][]string{columnHeaderMeasurements}, rows)
	}
	{ // empty steps support
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)

		row := 1
		assert.NoError(t, writeMeasurementsSection(w, &row, 0, []*rs.Process_Config_Step{}))
		assert.NoError(t, w.Flush())

		rows, err := w.File.GetRows(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, [][]string{columnHeaderMeasurements}, rows)
	}

	// bad cases
	{ // nil stream writer
		assert.Equal(t, errorNilStreamWriter, errors.Unwrap(writeMeasurementsSection(nil, nil, 0, nil)))
	}
	{ // nil row
		w, err := excelize.NewFile().NewStreamWriter(sheetDefault)
		assert.NoError(t, err)
		assert.Equal(t, errorNilRow, errors.Unwrap(writeMeasurementsSection(w, nil, 0, nil)))
		assert.NoError(t, w.Flush())
	}
}

func Test_writeProcessSheet(t *testing.T) {
	{ // type: process
		f := excelize.NewFile()
		styles, err := setDefaultStyles(f)
		assert.NoError(t, err)

		processSheet := "P-79700-1"
		expectedRows := [][]string{
			{"STATIONS", "station-1", "station-2", "station-3"},
			columnHeaderTools,
			{"A13", rs.ToolRole_MOLD.String(), "TRUE"},
			{"U79700N004", rs.ToolRole_PAPER.String(), "FALSE"},
			nil,
			columnHeaderRecipeMaterials,
			{"1", "20003", "A", "20", "", "", "", "2", "", "kg"},
			{"1", "20004", "A", "20", "", "", "1", "", "", "kg"},
			{"1", "U79700N004-1-CMP-1", "", "202", "U79700N004-1-CMP1", "", "", "", "3", "kg"},
			{"1", "60005", "A", "60", "", "", "1", "2", "3", "kg"},
			nil,
			{"CONTROL", "property-central-value", "property-min-value", "property-max-value", "property-all-value", "RAM_POSITION"},
			{"UNIT", "kg", "kg", "kg", "kg", ""},
			{"1", "mid:2", "low:1", "high:3", "low:1;mid:2;high:3", "UP"},
			nil,
			columnHeaderMeasurements,
			{"1", "property-central-value", "", "2", "", "kg"},
			{"1", "property-min-value", "1", "", "", "kg"},
			{"1", "property-max-value", "", "", "3", "kg"},
			{"1", "property-all-value", "1", "2", "3", "kg"},
			{"1", "RAM_POSITION", "", "", "", ""},
		}

		assert.NoError(t, writeProcessSheet(f, processSheet, styles, rs.ProductType_RUBBER, normalProcesses()[0]))

		actualRows, err := f.GetRows(processSheet)
		assert.NoError(t, err)
		assert.Equal(t, expectedRows, actualRows)
		expectedRowStyle := [][]int{
			{styles.highlight, styles.highlight, styles.highlight, styles.highlight},
			{styles.warning, styles.warning, styles.warning},
			{0, 0, 0},
			{0, 0, 0},
			nil,
			{styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			nil,
			{styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning},
			{styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning},
			{0, 0, 0, 0, 0, 0},
			nil,
			{styles.warning, styles.warning, styles.warning, styles.warning, styles.warning, styles.warning},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		}

		styleEqual(t, f, processSheet, expectedRowStyle)
	}

	// bad cases
	{ // nil file
		assert.Equal(t, errorNilFile, errors.Unwrap(writeProcessSheet(nil, "", nil, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, nil)))
	}
	{ // empty process sheet name
		assert.Equal(t, errorEmptySheetName, errors.Unwrap(writeProcessSheet(excelize.NewFile(), "", nil, rs.ProductType_RUBBER, normalProcesses()[0])))
	}
	{ // nil styles
		assert.Equal(t, errorNilStyles, errors.Unwrap(writeProcessSheet(excelize.NewFile(), "P-79700-1", nil, rs.ProductType_RUBBER, normalProcesses()[0])))
	}
	{ // nil process
		styles, err := setDefaultStyles(excelize.NewFile())
		assert.NoError(t, err)
		assert.Equal(t, errorNilProcess, errors.Unwrap(writeProcessSheet(excelize.NewFile(), "P-79700-1", styles, rs.ProductType_RUBBER, nil)))
	}
	{ // product type unspecified
		styles, err := setDefaultStyles(excelize.NewFile())
		assert.NoError(t, err)
		assert.Equal(t, errorProductTypeUnspecified, errors.Unwrap(writeProcessSheet(excelize.NewFile(), "P-79700-1", styles, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, normalProcesses()[0])))
	}
}

func Test_writeFlowProcesses(t *testing.T) {
	file := excelize.NewFile()
	file.NewSheet(sheetProcessIndex)
	w, err := file.NewStreamWriter(sheetProcessIndex)
	assert.NoError(t, err)
	styles, err := setDefaultStyles(file)
	assert.NoError(t, err)
	row := 1

	{
		expectedRows := [][]string{
			columnHeaderFlowProcesses,
			{"79700-1", rs.ProcessType_PROCESS.String(), "P-79700-1", "79700-1"},
			{"79700-1", rs.ProcessType_INSPECTION.String(), "I-79700-1", "79700-1"},
		}

		assert.NoError(t, writeFlowProcesses(file, w, styles, &row, rs.ProductType_RUBBER, normalProcesses()))
		assert.NoError(t, w.Flush())

		actualRows, err := file.GetRows(sheetProcessIndex)
		assert.NoError(t, err)
		assert.Equal(t, expectedRows, actualRows)

		expectedRowStyle := [][]int{
			{styles.warning, styles.warning, styles.warning, styles.warning},
			{0, 0, 2, 0},
			{0, 0, 2, 0},
		}
		styleEqual(t, w.File, sheetProcessIndex, expectedRowStyle)
	}

	// bad cases
	{ // nil file
		assert.Equal(t, errorNilFile, errors.Unwrap(writeFlowProcesses(nil, w, styles, &row, rs.ProductType_RUBBER, normalProcesses())))
		assert.NoError(t, w.Flush())
	}
	{ // nil stream writer
		assert.Equal(t, errorNilStreamWriter, errors.Unwrap(writeFlowProcesses(file, nil, styles, &row, rs.ProductType_RUBBER, normalProcesses())))
	}
	{ // nil styles
		assert.Equal(t, errorNilStyles, errors.Unwrap(writeFlowProcesses(file, w, nil, &row, rs.ProductType_RUBBER, normalProcesses())))
		assert.NoError(t, w.Flush())
	}
	{ // nil row
		assert.Equal(t, errorNilRow, errors.Unwrap(writeFlowProcesses(file, w, styles, nil, rs.ProductType_RUBBER, normalProcesses())))
		assert.NoError(t, w.Flush())
	}
	{ // product type unspecified
		assert.Equal(t, errorProductTypeUnspecified, errors.Unwrap(writeFlowProcesses(file, w, styles, &row, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, normalProcesses())))
		assert.NoError(t, w.Flush())
	}
	{ // empty process sheet name - type unspecified
		processes := normalProcesses()
		processes[0].Type = rs.ProcessType_PROCESS_TYPE_UNSPECIFIED
		err := writeFlowProcesses(file, w, styles, &row, rs.ProductType_RUBBER, processes)
		assert.Equal(t, errors.New("failed to get sheet name: process type=PROCESS_TYPE_UNSPECIFIED, name=79700-1 "), errors.Unwrap(err))
		assert.NoError(t, w.Flush())
	}
	{ // empty process sheet name - name
		processes := normalProcesses()
		processes[0].Name = ""
		err := writeFlowProcesses(file, w, styles, &row, rs.ProductType_RUBBER, processes)
		assert.Equal(t, errors.New("failed to get sheet name: process type=PROCESS, name= "), errors.Unwrap(err))
		assert.NoError(t, w.Flush())
	}
}

func Test_writeProcessIndexSheet(t *testing.T) {
	f := excelize.NewFile()
	styles, err := setDefaultStyles(f)
	assert.NoError(t, err)

	{
		expectedRows := [][]string{
			{"STANDARD PROCESSES"},
			columnHeaderFlowProcesses,
			{"79700-1", rs.ProcessType_PROCESS.String(), "P-79700-1", "79700-1"},
			{"79700-1", rs.ProcessType_INSPECTION.String(), "I-79700-1", "79700-1"},
			nil,
			{"OPTIONAL PROCESSES"},
			columnHeaderOptionalFlowInfo,
			{"optional-flow-1", "P-process-1", "10"},
			columnHeaderFlowProcesses,
			{"79700-1", rs.ProcessType_PROCESS.String(), "P-79700-1", "79700-1"},
			{"79700-1", rs.ProcessType_INSPECTION.String(), "I-79700-1", "79700-1"},
		}

		assert.NoError(t, writeProcessIndexSheet(f, styles, rs.ProductType_RUBBER, normalProcesses(), normalOptionalFlows()))

		actualRows, err := f.GetRows(sheetProcessIndex)
		assert.NoError(t, err)
		assert.Equal(t, expectedRows, actualRows)

		expectedRowStyle := [][]int{
			{styles.highlight},
			{styles.warning, styles.warning, styles.warning, styles.warning},
			{0, 0, styles.hyperlink, 0},
			{0, 0, styles.hyperlink, 0},
			nil,
			{styles.highlight},
			{styles.warning, styles.warning, styles.warning},
			{0, 0, 0},
			{styles.warning, styles.warning, styles.warning, styles.warning},
			{0, 0, styles.hyperlink, 0},
			{0, 0, styles.hyperlink, 0},
		}
		styleEqual(t, f, sheetProcessIndex, expectedRowStyle)
	}

	// bad cases
	{ // nil file
		assert.Equal(t, errorNilFile, errors.Unwrap(writeProcessIndexSheet(nil, styles, rs.ProductType_RUBBER, normalProcesses(), normalOptionalFlows())))
	}
	{ // nil styles
		assert.Equal(t, errorNilStyles, errors.Unwrap(writeProcessIndexSheet(f, nil, rs.ProductType_RUBBER, normalProcesses(), normalOptionalFlows())))
	}
	{ // product type unspecified
		assert.Equal(t, errorProductTypeUnspecified, errors.Unwrap(writeProcessIndexSheet(f, styles, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, normalProcesses(), normalOptionalFlows())))
	}
}

func Test_FromRecipe(t *testing.T) {
	// good case
	{
		readfile, err := ioutil.ReadFile("recipe.json")
		assert.NoError(t, err)

		var recipe = rs.Recipe{}
		assert.NoError(t, json.Unmarshal(readfile, &recipe))

		_, err = FromRecipe(&recipe)
		assert.NoError(t, err)
	}

	// bad cases
	{ // nil recipe
		_, err := FromRecipe(nil)
		assert.Equal(t, errorNilRecipe, err)
	}
	{ // empty product type
		recipe := normalRecipe()
		recipe.ProductType = ""
		_, err := FromRecipe(recipe)
		assert.Equal(t, strings.HasPrefix(err.Error(), "fail to convert product type"), true)
	}
	{ // empty product type
		recipe := normalRecipe()
		recipe.ProductType = "0"
		_, err := FromRecipe(recipe)
		assert.Equal(t, errors.Unwrap(err), errorProductTypeUnspecified)
	}
}
