package excel

import (
	"errors"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/stretchr/testify/assert"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
)

func Test_parseSpecIndex(t *testing.T) {
	expectedSpec := *normalSpec()
	expectedSpec.Materials = nil
	expectedSpec.Properties = nil
	expectedSpec.Tools = nil

	rows := [][]string{
		{"Spec ID", "2-79700-3-2"},
		{"Product ID", "79700-9"},
		{"Product Type", rs.ProductType_RUBBER.String()},
		{"Factories", rs.FactoryID_KY.String(), rs.FactoryID_KU.String(), rs.FactoryID_KC.String()},
		{"Major Version", "3"},
		{"Minor Version", "2"},
		{"Stage", rs.StageStatus_PRODUCTION.String()},
	}
	{ // good case
		actualSpec := &rs.Spec{}
		pt, err := parseSpecIndex(rows, actualSpec)
		assert.NoError(t, err)
		assert.Equal(t, &expectedSpec, actualSpec)
		assert.Equal(t, rs.ProductType_RUBBER, pt)
	}
	// bad cases
	{ // nil row
		pt, err := parseSpecIndex(nil, &rs.Spec{})
		assert.Equal(t, errorNilRow, err)
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // empty row
		pt, err := parseSpecIndex([][]string{}, &rs.Spec{})
		assert.Equal(t, errorNilRow, err)
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // bad product type
		rows := [][]string{
			{"Product Type", "BAD"},
		}
		pt, err := parseSpecIndex(rows, &rs.Spec{})
		assert.Equal(t, errors.New("bad product type=BAD"), err)
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // empty product type
		rows := [][]string{
			{"Product Type", ""},
		}
		pt, err := parseSpecIndex(rows, &rs.Spec{})
		assert.Equal(t, errors.New("bad product type="), err)
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // product type unspecified
		rows := [][]string{
			{"Product Type", "UNSPECIFIED"},
		}
		pt, err := parseSpecIndex(rows, &rs.Spec{})
		assert.Equal(t, errors.New("bad product type=UNSPECIFIED"), err)
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // bad factory
		rows := [][]string{
			{"Factories", "KM"},
		}
		pt, err := parseSpecIndex(rows, &rs.Spec{})
		assert.Equal(t, errors.New("bad factory=KM"), err)
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // empty factory
		rows := [][]string{
			{"Factories", ""},
		}
		pt, err := parseSpecIndex(rows, &rs.Spec{})
		assert.Equal(t, errors.New("bad factory="), err)
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // factory unspecified
		rows := [][]string{
			{"Factories", "UNSPECIFIED"},
		}
		pt, err := parseSpecIndex(rows, &rs.Spec{})
		assert.Equal(t, errors.New("bad factory=UNSPECIFIED"), err)
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // bad stageStatus
		rows := [][]string{
			{"Stage", "BAD_STAGE"},
		}
		pt, err := parseSpecIndex(rows, &rs.Spec{})
		assert.Equal(t, errors.New("bad stage=BAD_STAGE"), err)
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // empty stageStatus
		rows := [][]string{
			{"Stage", ""},
		}
		pt, err := parseSpecIndex(rows, &rs.Spec{})
		assert.Equal(t, errors.New("bad stage="), err)
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
	{ // stage status unspecified
		rows := [][]string{
			{"Stage", "UNSPECIFIED"},
		}
		pt, err := parseSpecIndex(rows, &rs.Spec{})
		assert.Equal(t, errors.New("bad stage=UNSPECIFIED"), err)
		assert.Equal(t, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, pt)
	}
}

func Test_parseProperties(t *testing.T) {
	rows := map[int][]string{
		3: {"property-max-value", "", "", "3", "", "kg"},
		0: columnHeaderProperties,
		4: {"property-all-value", "1", "2", "3", "", "kg"},
		5: {"RAM_POSITION", "", "", "", "UP"},
		2: {"property-min-value", "1", "", "", "", "kg"},
		1: {"property-central-value", "", "2", "", "", "kg"},
	}
	{ // good case
		properties, err := parseProperties(rows)
		assert.NoError(t, err)
		assert.Equal(t, normalProperties(), properties)
	}
	// nil support
	{
		properties, err := parseProperties(nil)
		assert.Empty(t, err)
		assert.Empty(t, properties)
	}
	{
		properties, err := parseProperties(map[int][]string{})
		assert.Empty(t, properties)
		assert.Empty(t, err)
	}
	// bad cases
	{
		rows := map[int][]string{
			1: {"property-central-value", "", "2", "", "", "", ""},
			0: columnHeaderProperties,
		}
		properties, err := parseProperties(rows)
		assert.Equal(t, errors.New("bad section format, row index=2"), errors.Unwrap(err))
		assert.Empty(t, properties)
	}
	{ // if enum has no value
		rows := map[int][]string{
			0: columnHeaderProperties,
			1: {"ENUM_TEST", "", "", "", ""},
		}
		properties, err := parseProperties(rows)
		assert.Equal(t, errors.New("missing enum value, row index=2"), errors.Unwrap(err))
		assert.Empty(t, properties)
	}
	{
		rows := map[int][]string{
			1: {"property-central-value", "", "", "", "", ""},
			0: columnHeaderProperties,
		}
		properties, err := parseProperties(rows)
		assert.Equal(t, errors.New("missing property value, row index=2"), errors.Unwrap(err))
		assert.Empty(t, properties)
	}
	{
		rows := map[int][]string{
			1: {"property-central-value", "", "2", "", "", ""},
			0: columnHeaderProperties,
		}
		properties, err := parseProperties(rows)
		assert.Equal(t, errors.New("missing item unit, row index=2"), errors.Unwrap(err))
		assert.Empty(t, properties)
	}
	{ // bad property name
		rows := map[int][]string{
			1: {"BAD_NAME", "", "", "", "UP"},
			0: columnHeaderProperties,
		}
		properties, err := parseProperties(rows)
		assert.Equal(t, errors.New("bad property name=BAD_NAME"), errors.Unwrap(err))
		assert.Empty(t, properties)
	}
	{ // bad property value
		rows := map[int][]string{
			0: columnHeaderProperties,
			1: {"RAM_POSITION", "", "", "", "BAD_VALUE"},
		}
		properties, err := parseProperties(rows)
		assert.Equal(t, errors.New("property name=RAM_POSITION has bad property value=BAD_VALUE"), errors.Unwrap(err))
		assert.Empty(t, properties)
	}
}

func Test_parseSpecConfigSheet(t *testing.T) {
	rows := [][]string{
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
		{"RAM_POSITION", "", "", "", "UP"},
	}
	{ // good case
		actualSpec := &rs.Spec{}
		assert.NoError(t, parseSpecConfigSheet(rows, rs.ProductType_RUBBER, actualSpec))
		materials := actualSpec.GetMaterials()
		tools := actualSpec.GetTools()
		properties := actualSpec.GetProperties()

		assert.Equal(t, []interface{}{normalMaterials(), normalTools(), normalProperties()}, []interface{}{materials, tools, properties})
	}
	// bad cases
	{ // nil rows
		assert.Equal(t, errorNilRow, parseSpecConfigSheet(nil, rs.ProductType_RUBBER, &rs.Spec{}))
	}
	{ // empty rows
		assert.Equal(t, errorNilRow, parseSpecConfigSheet([][]string{}, rs.ProductType_RUBBER, &rs.Spec{}))
	}
	{ // product type unspecified
		assert.Equal(t, errorProductTypeUnspecified, parseSpecConfigSheet(rows, rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, &rs.Spec{}))
	}
	{
		rows := [][]string{
			columnHeaderSpecMaterials,
			columnHeaderProperties,
			{"property-central-value", "", "2", "", "", "kg"},
		}
		err := parseSpecConfigSheet(rows, rs.ProductType_RUBBER, &rs.Spec{})
		assert.Equal(t, errors.New("missing materials"), err)
	}
	{
		rows := [][]string{
			columnHeaderTools,
			{"A13", "BAD", "TRUE"},
		}
		err := parseSpecConfigSheet(rows, rs.ProductType_RUBBER, &rs.Spec{})
		assert.Equal(t, errors.New("bad tool role:BAD"), err)
	}
	{
		rows := [][]string{
			columnHeaderTools,
			{"", rs.ToolRole_MOLD.String(), "TRUE"},
		}
		err := parseSpecConfigSheet(rows, rs.ProductType_RUBBER, &rs.Spec{})
		assert.Equal(t, errors.New("missing tool ID: row index=2"), err)
	}
	{
		rows := [][]string{
			columnHeaderTools,
			{"A13", "MOLD", "TRUE", "BAD"},
		}
		err := parseSpecConfigSheet(rows, rs.ProductType_RUBBER, &rs.Spec{})
		assert.Equal(t, errors.New("bad section format, row index=2"), err)
	}
	{
		rows := [][]string{
			columnHeaderSpecMaterials,
			{"20004", "A", "20", "", "", "1", "", "", "kg", "xxx"},
		}
		err := parseSpecConfigSheet(rows, rs.ProductType_RUBBER, &rs.Spec{})
		assert.Equal(t, errors.New("bad section format, row index=2"), err)
	}
	{
		rows := [][]string{
			columnHeaderProperties,
			{"ENUM_TEST", "", "", "", ""},
		}
		err := parseSpecConfigSheet(rows, rs.ProductType_RUBBER, &rs.Spec{})
		assert.Equal(t, errors.New("missing enum value, row index=2"), errors.Unwrap(err))
	}
	{ // bad property name
		rows := [][]string{
			columnHeaderSpecMaterials,
			{"20003", "A", "20", "", "", "", "2", "", "kg"},
			columnHeaderProperties,
			{"BAD_NAME", "", "", "", "UP"},
		}
		err := parseSpecConfigSheet(rows, rs.ProductType_RUBBER, &rs.Spec{})
		assert.Equal(t, errors.New("bad property name=BAD_NAME"), errors.Unwrap(err))
	}
	{ // bad property value
		rows := [][]string{
			columnHeaderSpecMaterials,
			{"20003", "A", "20", "", "", "", "2", "", "kg"},
			columnHeaderProperties,
			{"RAM_POSITION", "", "", "", "BAD_VALUE"},
		}
		err := parseSpecConfigSheet(rows, rs.ProductType_RUBBER, &rs.Spec{})
		assert.Equal(t, errors.New("property name=RAM_POSITION has bad property value=BAD_VALUE"), errors.Unwrap(err))
	}
}

func Test_ParseSpec(t *testing.T) {
	{
		_, err := ToSpec(nil)
		assert.Equal(t, errorNilFile, errors.Unwrap(err))
	}
	{
		_, err := ToSpec(excelize.NewFile())
		assert.Equal(t, excelize.ErrSheetNotExist{SheetName: sheetSpecIndex}, errors.Unwrap(err))
	}
	{
		f := excelize.NewFile()
		f.NewSheet(sheetSpecIndex)
		_, err := ToSpec(f)
		assert.Equal(t, excelize.ErrSheetNotExist{SheetName: sheetSpecConfig}, errors.Unwrap(err))
	}
}
