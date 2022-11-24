package excel

import (
	"errors"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/stretchr/testify/assert"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
)

func Test_sheetPrefixToProcessType(t *testing.T) {
	{
		assert.Equal(t, rs.ProcessType_PROCESS, sheetPrefixToProcessType("P"))
	}
	{
		assert.Equal(t, rs.ProcessType_INSPECTION, sheetPrefixToProcessType("I"))
	}
	{
		assert.Equal(t, rs.ProcessType_MEASUREMENT, sheetPrefixToProcessType("M"))
	}
	{ // wrong name
		assert.Equal(t, rs.ProcessType_PROCESS_TYPE_UNSPECIFIED, sheetPrefixToProcessType("R"))
	}
	{ // empty name
		assert.Equal(t, rs.ProcessType_PROCESS_TYPE_UNSPECIFIED, sheetPrefixToProcessType(""))
	}
}

func Test_parseRecipeProcessList(t *testing.T) {
	{
		rows := [][]string{
			{"STANDARD PROCESSES"},
			columnHeaderFlowProcesses,
			{"79700-1", "PROCESS", "P-79700-1", "79700-1"},
			{"79700-1", "INSPECTION", "I-79700-1", "79700-1"},
			nil,
			{"OPTIONAL PROCESSES"},
			columnHeaderOptionalFlowInfo,
			{"optional-flow-1", "P-process-1", "10"},
			columnHeaderFlowProcesses,
			{"79700-1", "PROCESS", "P-79700-1", "79700-1"},
			{"79700-1", "INSPECTION", "I-79700-1", "79700-1"},
		}

		actualRecipe := rs.Recipe{}
		err := parseRecipeProcessList(rows, &actualRecipe)
		assert.NoError(t, err)
		actualProcesses := actualRecipe.GetProcesses()
		actualOptional := actualRecipe.GetOptionalFlows()

		processes := normalProcesses()
		processes[0].Configs = []*rs.Process_Config{}
		processes[1].Configs = []*rs.Process_Config{}

		optionalFlows := normalOptionalFlows()
		optionalFlows[0].Processes = processes

		assert.Equal(t, processes, actualProcesses)
		assert.Equal(t, optionalFlows, actualOptional)
	}
	// bad cases
	{ // nil row
		err := parseRecipeProcessList(nil, &rs.Recipe{})
		assert.Equal(t, errorNilRow, errors.Unwrap(err))
	}
	{ // empty row
		err := parseRecipeProcessList([][]string{}, &rs.Recipe{})
		assert.Equal(t, errorNilRow, errors.Unwrap(err))
	}
	{ // bad standard process section
		rows := [][]string{
			{"STANDARD PROCESSES"},
			columnHeaderFlowProcesses,
			{"79700-1", "PROCESS", "P-79700-1", "79700-1", "5"},
		}
		err := parseRecipeProcessList(rows, &rs.Recipe{})
		assert.Equal(t, errors.New("bad standard process section format: row index=3"), err)
	}
	{ // bad standard optional process section
		rows := [][]string{
			{"OPTIONAL PROCESSES"},
			columnHeaderOptionalFlowInfo,
			{"optional-flow-1", "P-process-1", "10", "4", "5"},
		}
		err := parseRecipeProcessList(rows, &rs.Recipe{})
		assert.Equal(t, errors.New("bad standard optional process section format: row index=3"), err)
	}
	{ // bad after process
		rows := [][]string{
			{"OPTIONAL PROCESSES"},
			columnHeaderOptionalFlowInfo,
			{"optional-flow-1", "P&process&1", "10"},
		}
		err := parseRecipeProcessList(rows, &rs.Recipe{})
		assert.Equal(t, errors.New("bad after process=P&process&1"), err)
	}
	{ // bad after process type
		rows := [][]string{
			{"OPTIONAL PROCESSES"},
			columnHeaderOptionalFlowInfo,
			{"optional-flow-1", "J-process-1", "10"},
		}
		err := parseRecipeProcessList(rows, &rs.Recipe{})
		assert.Equal(t, errors.New("bad after process type=J-process-1"), err)
	}
}

func Test_checkProcessSheetExists(t *testing.T) {
	f := excelize.NewFile()
	{
		f.NewSheet("P-79700-1")
		f.NewSheet("I-79700-1")
		actualNames, err := checkProcessSheetExists(f, normalProcesses())
		assert.NoError(t, err)
		assert.Equal(t, []string{"P-79700-1", "I-79700-1"}, actualNames)
	}
	// bad cases
	{ // process type unspecified
		processes := []*rs.Process{
			{
				Name: "79700-1",
				Type: rs.ProcessType_PROCESS_TYPE_UNSPECIFIED,
			},
		}
		_, err := checkProcessSheetExists(f, processes)
		assert.Equal(t, errors.New("failed to get sheet name: process type=PROCESS_TYPE_UNSPECIFIED, name=79700-1 "), errors.Unwrap(err))
	}
	{ // sheet not exists
		processes := []*rs.Process{
			{
				Name: "79700-1",
				Type: rs.ProcessType_INSPECTION,
			},
		}
		_, err := checkProcessSheetExists(excelize.NewFile(), processes)
		assert.Equal(t, errors.New("sheet=I-79700-1 not exists"), err)
	}
	{ // nil file
		name, err := checkProcessSheetExists(nil, nil)
		assert.Empty(t, name)
		assert.Empty(t, err)
	}
	{ // nil processes
		name, err := checkProcessSheetExists(excelize.NewFile(), nil)
		assert.Empty(t, name)
		assert.Empty(t, err)
	}
	{ // empty processes
		name, err := checkProcessSheetExists(excelize.NewFile(), []*rs.Process{})
		assert.Empty(t, name)
		assert.Empty(t, err)
	}
}

func Test_parseRecipeIndex(t *testing.T) {
	expectedRecipe := *normalRecipe()
	expectedRecipe.Processes = nil
	expectedRecipe.OptionalFlows = nil

	rows := [][]string{
		{"Recipe ID", "U79700N004"},
		{"Product ID", "79700-9"},
		{"Product Type", rs.ProductType_RUBBER.String()},
		{"Factories", rs.FactoryID_KY.String(), rs.FactoryID_KU.String()},
		{"Major Version", "004"},
		{"Minor Version", ""},
		{"Parent Spec ID", "2-79700-3-2"},
		{"Stage", rs.StageStatus_PRODUCTION.String()},
	}
	{ // good case
		actualRecipe := &rs.Recipe{}
		pt, err := parseRecipeIndex(rows, actualRecipe)
		assert.NoError(t, err)
		assert.Equal(t, rs.ProductType_RUBBER, pt)
		assert.Equal(t, &expectedRecipe, actualRecipe)
	}
	// bad cases
	{ // nil row
		pt, err := parseRecipeIndex(nil, &rs.Recipe{})
		assert.Equal(t, errorNilRow, errors.Unwrap(err))
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // empty row
		pt, err := parseRecipeIndex([][]string{}, &rs.Recipe{})
		assert.Equal(t, errorNilRow, errors.Unwrap(err))
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // product type unspecified
		rows := [][]string{
			{"Product Type", rs.ProductType_PRODUCT_TYPE_UNSPECIFIED.String()},
		}
		pt, err := parseRecipeIndex(rows, &rs.Recipe{})
		assert.Equal(t, errors.New("bad product type=PRODUCT_TYPE_UNSPECIFIED"), errors.Unwrap(err))
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // bad product type
		rows := [][]string{
			{"Product Type", "BAD_PRODUCT_TYPE"},
		}
		pt, err := parseRecipeIndex(rows, &rs.Recipe{})
		assert.Equal(t, errors.New("bad product type=BAD_PRODUCT_TYPE"), errors.Unwrap(err))
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // empty product type
		rows := [][]string{
			{"Product Type", ""},
		}
		pt, err := parseRecipeIndex(rows, &rs.Recipe{})
		assert.Equal(t, errors.New("bad product type="), errors.Unwrap(err))
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // factory unspecified
		rows := [][]string{
			{"Factories", rs.FactoryID_FACTORY_UNSPECIFIED.String()},
		}
		pt, err := parseRecipeIndex(rows, &rs.Recipe{})
		assert.Equal(t, errors.New("bad factory=FACTORY_UNSPECIFIED"), errors.Unwrap(err))
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // bad factory
		rows := [][]string{
			{"Factories", "KM"},
		}
		actualRecipe := &rs.Recipe{}
		pt, err := parseRecipeIndex(rows, actualRecipe)
		assert.Equal(t, errors.Unwrap(err), errors.New("bad factory=KM"))
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // empty factory
		rows := [][]string{
			{"Factories", ""},
		}
		actualRecipe := &rs.Recipe{}
		pt, err := parseRecipeIndex(rows, actualRecipe)
		assert.Equal(t, errors.Unwrap(err), errors.New("bad factory="))
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // stageStatus unspcified
		rows := [][]string{
			{"Stage", rs.StageStatus_STAGE_STATUS_UNSPECIFIED.String()},
		}
		pt, err := parseRecipeIndex(rows, &rs.Recipe{})
		assert.Equal(t, errors.New("bad stage=STAGE_STATUS_UNSPECIFIED"), errors.Unwrap(err))
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // bad stageStatus
		rows := [][]string{
			{"Stage", "TEST"},
		}
		actualRecipe := &rs.Recipe{}
		pt, err := parseRecipeIndex(rows, actualRecipe)
		assert.Equal(t, errors.Unwrap(err), errors.New("bad stage=TEST"))
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // empty stageStatus
		rows := [][]string{
			{"Stage", ""},
		}
		actualRecipe := &rs.Recipe{}
		pt, err := parseRecipeIndex(rows, actualRecipe)
		assert.Equal(t, errors.Unwrap(err), errors.New("bad stage="))
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}

}

func Test_parseTools(t *testing.T) {
	{
		rows := map[int][]string{
			1: {"A13", "MOLD", "TRUE"},
			0: columnHeaderTools,
			2: {"U79700N004", "PAPER", "FALSE"},
		}
		tools, err := parseTools(rows)
		assert.NoError(t, err)

		assert.Equal(t, normalTools(), tools)
	}
	// nil support
	{
		tools, err := parseTools(nil)
		assert.Empty(t, tools)
		assert.Empty(t, err)
	}
	{
		tools, err := parseTools(map[int][]string{})
		assert.Empty(t, tools)
		assert.Empty(t, err)
	}
	// bad cases
	{
		rows := map[int][]string{
			0: {"NO", "A13", "MOLD", "TRUE"},
		}
		_, err := parseTools(rows)
		assert.Equal(t, errors.New("bad tool section format: row index=1"), err)
	}
	{
		rows := map[int][]string{
			0: columnHeaderTools,
			1: {"", "MOLD", "TRUE"},
		}
		_, err := parseTools(rows)
		assert.Equal(t, errors.New("missing tool ID: row index=2"), err)
	}
	{
		rows := map[int][]string{
			0: columnHeaderTools,
			1: {"A13", "", "TRUE"},
		}
		_, err := parseTools(rows)
		assert.Equal(t, errors.New("missing tool role: row index=2"), err)
	}
	{
		rows := map[int][]string{
			0: columnHeaderTools,
			1: {"A13", "TEST", "TRUE"},
		}
		_, err := parseTools(rows)
		assert.Equal(t, errors.New("bad tool role:TEST"), err)
	}
	{
		rows := map[int][]string{
			0: columnHeaderTools,
			1: {"A13", "MOLD", ""},
		}
		_, err := parseTools(rows)
		assert.Equal(t, errors.New("missing tool check: row index=2"), err)
	}
}

func Test_parseMaterials(t *testing.T) {
	{ // good case
		rows := map[int][]string{
			2: {"1", "20004", "A", "20", "", "", "1", "", "", "kg"},
			0: columnHeaderRecipeMaterials,
			1: {"1", "20003", "A", "20", "", "", "", "2", "", "kg"},
			4: {"1", "60005", "A", "60", "", "", "1", "2", "3", "kg"},
			3: {"1", "U79700N004-1-CMP-1", "", "202", "U79700N004-1-CMP1", "", "", "", "3", "kg"},
		}

		expected := map[int][]*rs.Material{
			1: normalMaterials(),
		}

		actual, err := parseMaterials(rs.ProductType_RUBBER, rows)
		assert.NoError(t, err)

		assert.Equal(t, expected, actual)
	}
	// nil support
	{
		materials, err := parseMaterials(rs.ProductType_RUBBER, nil)
		assert.Empty(t, materials)
		assert.Empty(t, err)
	}
	{
		materials, err := parseMaterials(rs.ProductType_RUBBER, map[int][]string{})
		assert.Empty(t, materials)
		assert.Empty(t, err)
	}
	// bad cases
	{
		rows := map[int][]string{
			0: columnHeaderRecipeMaterials,
			1: {"1", "20003", "C", "20", "", "", "", "2", ""},
		}
		_, err := parseMaterials(rs.ProductType_RUBBER, rows)
		assert.Equal(t, errors.New("bad section format: row index=2"), errors.Unwrap(err))
	}
	{
		rows := map[int][]string{
			0: columnHeaderRecipeMaterials,
			1: {"1", "20003", "C", "20", "", "", "", "2", "", "PHR"},
		}
		_, err := parseMaterials(rs.ProductType_RUBBER, rows)
		assert.Equal(t, errors.New("bad material level=C, row index=2"), errors.Unwrap(err))
	}
	{
		rows := map[int][]string{
			0: columnHeaderRecipeMaterials,
			1: {"1", "79700-1", "A", "202", "", "", "", "2", "", "PHR"},
		}
		_, err := parseMaterials(rs.ProductType_RUBBER, rows)
		assert.Equal(t, errors.New("material name=79700-1 and type=202 mismatch"), errors.Unwrap(err))
	}
	{
		rows := map[int][]string{
			0: columnHeaderRecipeMaterials,
			1: {"1", "20003", "A", "30", "", "", "", "2", "", "PHR"},
		}
		_, err := parseMaterials(rs.ProductType_RUBBER, rows)
		assert.Equal(t, errors.New("material name=20003 and type=30 mismatch"), errors.Unwrap(err))
	}
	{
		rows := map[int][]string{
			0: columnHeaderRecipeMaterials,
			1: {"1", "U79700N004-1-CMP1-1", "", "30", "", "", "", "2", "", "PHR"},
		}
		_, err := parseMaterials(rs.ProductType_RUBBER, rows)
		assert.Equal(t, errors.New("material name=U79700N004-1-CMP1-1 and type=30 mismatch"), errors.Unwrap(err))
	}
	{
		rows := map[int][]string{
			0: columnHeaderRecipeMaterials,
			1: {"1", "U79700N004-1-ACL1-1", "", "30", "", "", "", "2", "", "PHR"},
		}
		_, err := parseMaterials(rs.ProductType_RUBBER, rows)
		assert.Equal(t, errors.New("material name=U79700N004-1-ACL1-1 and type=30 mismatch"), errors.Unwrap(err))
	}
	{
		rows := map[int][]string{
			0: columnHeaderRecipeMaterials,
			1: {"1", "U79700N004-1-ACL1-1", "", "203", "", "99", "", "2", "", "PHR"},
		}
		_, err := parseMaterials(rs.ProductType_RUBBER, rows)
		assert.Equal(t, errors.New("bad container type=99 of product type=RUBBER mapping list, row index=2"), errors.Unwrap(err))
	}
	{
		rows := map[int][]string{
			0: columnHeaderRecipeMaterials,
			1: {"1", "20003", "A", "20", "", "", "", "", "", "PHR"},
		}
		_, err := parseMaterials(rs.ProductType_RUBBER, rows)
		assert.Equal(t, errors.New("bad material quantity, low= mid= high=, row index=2"), errors.Unwrap(err))
	}
	{
		rows := map[int][]string{
			0: columnHeaderRecipeMaterials,
			1: {"1", "20003", "A", "20", "", "", "", "", "3", ""},
		}
		_, err := parseMaterials(rs.ProductType_RUBBER, rows)
		assert.Equal(t, errors.New("missing material unit, row index=2"), errors.Unwrap(err))
	}
}

func Test_parseStepsControls(t *testing.T) {
	{
		rows := map[int][]string{
			1: {"mid:2", "low:1", "high:3", "low:1;mid:2;high:3", "UP"},
		}
		items := []string{"property-central-value", "property-min-value", "property-max-value", "property-all-value", "RAM_POSITION"}
		units := []string{"kg", "kg", "kg", "kg", ""}

		expected := map[int][]*rs.Property{
			1: normalProperties(),
		}
		actual, err := parseStepsControls(items, units, rows)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	}
	// nil support
	{
		controls, err := parseStepsControls(nil, nil, nil)
		assert.Empty(t, controls)
		assert.Empty(t, err)
	}
	{
		controls, err := parseStepsControls([]string{}, []string{}, map[int][]string{})
		assert.Empty(t, err)
		assert.Empty(t, controls)
	}
	// bad cases
	{
		rows := map[int][]string{
			1: {"no:1"},
		}
		items := []string{"err_test"}
		units := []string{""}
		_, err := parseStepsControls(items, units, rows)
		assert.Equal(t, errors.New("bad value format: step=1, item=err_test, value=no:1"), err)
	}
	{
		rows := map[int][]string{
			1: {"mid:2"},
		}
		items := []string{"property-central-value"}
		units := []string{""}
		_, err := parseStepsControls(items, units, rows)
		assert.Equal(t, errors.New("missing unit"), errors.Unwrap(err))
	}
	{ // bad property name
		rows := map[int][]string{
			1: {"UP"},
		}
		items := []string{"BAD_PROPERTY_NAME"}
		units := []string{}
		_, err := parseStepsControls(items, units, rows)
		assert.Equal(t, errors.New("bad value format: step=1, item=BAD_PROPERTY_NAME, value=UP"), err)
	}
	{ // bad property value
		rows := map[int][]string{
			1: {"BAD_PROPERTY_VALUE"},
		}
		items := []string{"RAM_POSITION"}
		units := []string{}
		_, err := parseStepsControls(items, units, rows)
		assert.Equal(t, errors.New("bad value format: step=1, item=RAM_POSITION, value=BAD_PROPERTY_VALUE"), err)
	}
}

func Test_parseControls(t *testing.T) {
	{
		rows := map[int][]string{
			2: {"1", "mid:2", "low:1", "high:3", "low:1;mid:2;high:3", "UP"},
			0: {"CONTROL", "property-central-value", "property-min-value", "property-max-value", "property-all-value", "RAM_POSITION"},
			1: {"UNIT", "kg", "kg", "kg", "kg", ""},
		}

		expected := map[int][]*rs.Property{
			1: normalProperties(),
		}

		actual, err := parseControls(rows)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	}
	// nil support
	{
		controls, err := parseControls(nil)
		assert.Empty(t, controls)
		assert.Empty(t, err)
	}
	{
		controls, err := parseControls(map[int][]string{})
		assert.Empty(t, controls)
		assert.Empty(t, err)
	}
	// bad cases
	{
		rows := map[int][]string{
			0: {"1"},
		}
		_, err := parseControls(rows)
		assert.Equal(t, errors.New("bad section format, row index=1"), errors.Unwrap(err))
	}
	{
		rows := map[int][]string{
			0: {"HI", "property-central-value", "property-min-value", "property-max-value", "property-all-value", "RAM_POSITION"},
			1: {"UNIT", "", "", "", "kg", ""},
		}
		_, err := parseControls(rows)
		assert.Equal(t, errors.New("bad step=HI"), errors.Unwrap(err))
	}
	{
		rows := map[int][]string{
			0: {"CONTROL", "property-central-value", "property-central-value", "property-max-value", "property-all-value", "RAM_POSITION"},
			1: {"UNIT", "", "", "", "kg", ""},
		}
		_, err := parseControls(rows)
		assert.Equal(t, errors.New("missing values"), errors.Unwrap(err))
	}
	{ // bad property value
		rows := map[int][]string{
			0: {"CONTROL", "RAM_POSITION"},
			1: {"UNIT", ""},
			2: {"1", "HI"},
		}
		_, err := parseControls(rows)
		assert.Equal(t, errors.New("bad value format: step=1, item=RAM_POSITION, value=HI"), errors.Unwrap(err))
	}
	{ // bad property name
		rows := map[int][]string{
			0: {"CONTROL", "UNKNOWN"},
			1: {"UNIT", ""},
			2: {"1", "UP"},
		}
		_, err := parseControls(rows)
		assert.Equal(t, errors.New("bad value format: step=1, item=UNKNOWN, value=UP"), errors.Unwrap(err))
	}
}
func Test_parseMeasurements(t *testing.T) {
	{ // good case
		rows := map[int][]string{
			2: {"1", "property-min-value", "1", "", "", "kg"},
			0: columnHeaderMeasurements,
			3: {"1", "property-max-value", "", "", "3", "kg"},
			1: {"1", "property-central-value", "", "2", "", "kg"},
			4: {"1", "property-all-value", "1", "2", "3", "kg"},
		}
		measurements := make(map[int][]*rs.Property)
		measurements[1] = normalProperties()[0:4]

		actual, err := parseMeasurements(rows)
		assert.NoError(t, err)

		assert.Equal(t, measurements, actual)
	}
	// nil support
	{
		measurements, err := parseMeasurements(nil)
		assert.Empty(t, measurements)
		assert.Empty(t, err)
	}
	{
		measurements, err := parseMeasurements(map[int][]string{})
		assert.Empty(t, measurements)
		assert.Empty(t, err)
	}
	// bad cases
	{
		rows := map[int][]string{
			0: columnHeaderMeasurements,
			1: {"1", "property-central-value", "", "2", "", "kg", ""},
		}
		_, err := parseMeasurements(rows)
		assert.Equal(t, errors.New("bad section format, row index=2"), errors.Unwrap(err))
	}
	{
		rows := map[int][]string{
			0: columnHeaderMeasurements,
			1: {"1", "", "", "2", "", "kg"},
		}
		_, err := parseMeasurements(rows)
		assert.Equal(t, errors.New("missing item name, row index=2"), errors.Unwrap(err))
	}
	{
		rows := map[int][]string{
			0: columnHeaderMeasurements,
			1: {"1", "property-central-value", "", "", "", "kg"},
		}
		_, err := parseMeasurements(rows)
		assert.Equal(t, errors.New("missing item value, low= mid= high=, row index=2"), errors.Unwrap(err))
	}
	{
		rows := map[int][]string{
			0: columnHeaderMeasurements,
			1: {"1", "property-central-value", "", "", "3", ""},
		}
		_, err := parseMeasurements(rows)
		assert.Equal(t, errors.New("missing item unit, row index=2"), errors.Unwrap(err))
	}
}

func Test_parseProcessSheets(t *testing.T) {
	// good cases
	{ // nil support
		processes, err := parseProcessSheets(nil, rs.ProductType_RUBBER, nil)
		assert.Empty(t, processes)
		assert.NoError(t, err)
	}
	{ // empty names
		actual, err := parseProcessSheets(excelize.NewFile(), rs.ProductType_RUBBER, []string{})
		assert.Empty(t, actual)
		assert.NoError(t, err)
	}
	{ // nil names
		actual, err := parseProcessSheets(excelize.NewFile(), rs.ProductType_RUBBER, nil)
		assert.Empty(t, actual)
		assert.NoError(t, err)
	}

	{ // product type unspecified
		_, err := parseProcessSheets(excelize.NewFile(), rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, []string{})
		assert.Equal(t, errorProductTypeUnspecified, err)
	}
}

func Test_parseProcessSheet(t *testing.T) {
	{
		rows := [][]string{
			{"STATIONS", "station-1", "station-2", "station-3"},
			columnHeaderTools,
			{"A13", rs.ToolRole_MOLD.String(), "TRUE"},
			{"U79700N004", rs.ToolRole_PAPER.String(), "FALSE"},
			columnHeaderRecipeMaterials,
			{"1", "20003", "A", "20", "", "", "", "2", "", "kg"},
			{"1", "20004", "A", "20", "", "", "1", "", "", "kg"},
			{"1", "U79700N004-1-CMP-1", "", "202", "U79700N004-1-CMP1", "", "", "", "3", "kg"},
			{"1", "60005", "A", "60", "", "", "1", "2", "3", "kg"},
			{"CONTROL", "property-central-value", "property-min-value", "property-max-value", "property-all-value", rs.PropertyName_RAM_POSITION.String()},
			{"UNIT", "kg", "kg", "kg", "kg", ""},
			{"1", "mid:2", "low:1", "high:3", "low:1;mid:2;high:3", rs.RamPosition_UP.String()},
			columnHeaderMeasurements,
			{"1", "property-central-value", "", "2", "", "kg"},
			{"1", "property-min-value", "1", "", "", "kg"},
			{"1", "property-max-value", "", "", "3", "kg"},
			{"1", "property-all-value", "1", "2", "3", "kg"},
		}

		expected := normalConfigs()
		expected[0].Steps[0].Measurements = expected[0].Steps[0].Measurements[:4]

		actual, err := parseProcessSheet(rows, rs.ProductType_RUBBER)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	}
	{ // nil support
		config, err := parseProcessSheet(nil, rs.ProductType_RUBBER)
		assert.Empty(t, config)
		assert.Empty(t, err)
	}
	{
		config, err := parseProcessSheet([][]string{}, rs.ProductType_RUBBER)
		assert.Empty(t, config)
		assert.Empty(t, err)
	}
	// bad cases
	{
		rows := [][]string{
			columnHeaderTools,
			{"A13", rs.ToolRole_MOLD.String(), "TRUE", ""},
		}
		_, err := parseProcessSheet(rows, rs.ProductType_RUBBER)
		assert.Equal(t, errors.New("bad TOOL section format, row index=2"), err)
	}
	{
		rows := [][]string{
			columnHeaderRecipeMaterials,
			{"1", "material-1", "A", "20", "", "", "", "2", ""},
		}
		_, err := parseProcessSheet(rows, rs.ProductType_RUBBER)
		assert.Equal(t, errors.New("bad MATERIAL section format, row index=2"), err)
	}
	{
		rows := [][]string{
			{"CONTROL"},
			{"HI"},
		}
		_, err := parseProcessSheet(rows, rs.ProductType_RUBBER)
		assert.Equal(t, errors.New("bad CONTROL section format, row index=2"), err)
	}
	{
		rows := [][]string{
			columnHeaderMeasurements,
			{"1", "property-all-value", "1", "2", "3"},
		}
		_, err := parseProcessSheet(rows, rs.ProductType_RUBBER)
		assert.Equal(t, errors.New("bad MEASUREMENT section format, row index=2"), err)
	}
}

func Test_ParseExcel(t *testing.T) {
	{
		_, err := ToRecipe(nil)
		assert.Equal(t, errorNilFile, errors.Unwrap(err))
	}
	{
		_, err := ToRecipe(excelize.NewFile())
		assert.Equal(t, excelize.ErrSheetNotExist{SheetName: sheetRecipeIndex}, errors.Unwrap(err))
	}
	{
		f := excelize.NewFile()
		f.NewSheet(sheetRecipeIndex)
		_, err := ToRecipe(f)
		assert.Equal(t, excelize.ErrSheetNotExist{SheetName: sheetProcessIndex}, errors.Unwrap(err))
	}
}
