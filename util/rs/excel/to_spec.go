package excel

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
	utilRS "gitlab.kenda.com.tw/kenda/commons/v2/util/rs"
)

func parseSpecIndex(rows [][]string, spec *rs.Spec) (productType rs.ProductType, err error) {
	if len(rows) == 0 {
		err = errorNilRow
		return
	}

	spec.Version = &rs.Version{}

	for _, row := range rows {
		if len(row) <= 1 { // row has no columns or has title only
			continue
		}

		title, values := row[0], row[1:]
		switch title {
		case "Spec ID":
			spec.Id = values[0]
		case "Product ID":
			spec.ProductId = values[0]
		case "Product Type":
			if pt, ok := rs.ProductType_value[strings.ToUpper(values[0])]; !ok || pt == int32(rs.ProductType_PRODUCT_TYPE_UNSPECIFIED) {
				err = fmt.Errorf("bad product type=%s", values[0])
				return
			} else {
				productType = rs.ProductType(pt)
				spec.ProductType = strconv.Itoa(int(pt))
			}
		case "Factories":
			for _, value := range values {
				if factory, ok := rs.FactoryID_value[strings.ToUpper(value)]; !ok || factory == int32(rs.FactoryID_FACTORY_UNSPECIFIED) {
					err = fmt.Errorf("bad factory=%s", value)
					return
				} else {
					spec.Factory = append(spec.Factory, rs.FactoryID(factory))
				}
			}
		case "Major Version":
			spec.Version.Major = values[0]
		case "Minor Version":
			spec.Version.Minor = values[0]
		case "Stage":
			if stage, ok := rs.StageStatus_value[strings.ToUpper(values[0])]; !ok || stage == int32(rs.StageStatus_STAGE_STATUS_UNSPECIFIED) {
				err = fmt.Errorf("bad stage=%s", values[0])
				return
			} else {
				spec.Version.StageStatus = rs.StageStatus(stage)
			}
		}
	}
	return
}

const errorParseProperties = "failed to parse properties: %w"

func parseProperties(rows map[int][]string) (properties []*rs.Property, err error) {
	var keys []int
	for k := range rows {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, key := range keys {
		rowIdx := key + 1
		row := rows[key]
		if reflect.DeepEqual(row, columnHeaderProperties) { // skip title
			continue
		}
		if len(row) == 0 { // skip empty row
			continue
		}

		switch len(row) {
		default:
			return nil, fmt.Errorf(errorParseProperties, fmt.Errorf("bad section format, row index=%d", rowIdx))
		case 5: // enum property
			if row[4] == "" {
				return nil, fmt.Errorf(errorParseProperties, fmt.Errorf("missing enum value, row index=%d", rowIdx))
			}
			enum, err := utilRS.PropertySettingNameToValue(row[0], row[4])
			if err != nil {
				return nil, fmt.Errorf(errorParseProperties, err)
			}
			properties = append(properties, &rs.Property{
				Name: row[0],
				Param: &rs.Param{
					Type: rs.Param_ENUM,
					Enum: &rs.Param_Enum{Value: enum},
				},
			})
		case 6: // value property
			if row[1] == "" && row[2] == "" && row[3] == "" {
				return nil, fmt.Errorf(errorParseProperties, fmt.Errorf("missing property value, row index=%d", rowIdx))
			}
			if row[5] == "" {
				return nil, fmt.Errorf(errorParseProperties, fmt.Errorf("missing item unit, row index=%d", rowIdx))
			}
			param, err := stringToParamValues(row[1], row[2], row[3], row[5])
			if err != nil {
				return nil, fmt.Errorf(errorParseProperties, err)
			}
			properties = append(properties, &rs.Property{
				Name:  row[0],
				Param: param,
			})
		}
	}
	return
}

func parseSpecConfigSheet(rows [][]string, productType rs.ProductType, spec *rs.Spec) (err error) {
	if len(rows) == 0 {
		return errorNilRow
	}

	if productType == rs.ProductType_PRODUCT_TYPE_UNSPECIFIED {
		return errorProductTypeUnspecified
	}

	type configSection map[string]map[int][]string // section name -> row index -> row

	var name string
	section := configSection{}
	for key, row := range rows {
		rowIdx := key + 1
		if len(row) == 0 { // skip empty row
			continue
		}
		if reflect.DeepEqual(row, columnHeaderTools) {
			name = columnHeaderTools[0]
			section[name] = map[int][]string{}
			continue
		}
		if reflect.DeepEqual(row, columnHeaderSpecMaterials) {
			name = columnHeaderSpecMaterials[0]
			section[name] = map[int][]string{}
			continue
		}
		if reflect.DeepEqual(row, columnHeaderProperties) {
			name = columnHeaderProperties[0]
			section[name] = map[int][]string{}
			continue
		}

		switch name {
		case columnHeaderTools[0]:
			if len(row) != 3 {
				return fmt.Errorf("bad section format, row index=%d", rowIdx)
			}
		case columnHeaderSpecMaterials[0]:
			if len(row) != 9 {
				return fmt.Errorf("bad section format, row index=%d", rowIdx)
			}
			row = append([]string{"1"}, row...) // spec only have one step
		case columnHeaderProperties[0]:
			if len(row) != 6 && // value property
				len(row) != 5 { // enum property
				return fmt.Errorf("bad section format, row index=%d", rowIdx)
			}
		default: // continue if no section name
			continue
		}
		section[name][key] = row
	}

	if t, ok := section[columnHeaderTools[0]]; ok {
		tools, err := parseTools(t)
		if err != nil {
			return err
		}
		spec.Tools = tools
	}

	if m, ok := section[columnHeaderSpecMaterials[0]]; ok {
		materials, err := parseMaterials(productType, m)
		if err != nil {
			return err
		}
		if len(materials) == 0 { // spec must have materials
			return errors.New("missing materials")
		}
		spec.Materials = materials[1]
	}

	if p, ok := section[columnHeaderProperties[0]]; ok {
		properties, err := parseProperties(p)
		if err != nil {
			return err
		}
		spec.Properties = properties
	}

	return nil
}

const errorFailedToParseSpecExcel = "failed to parse spec excel: err=%w"

// ToSpec parse excel file to spec
func ToSpec(file *excelize.File) (spec *rs.Spec, err error) {
	if file == nil {
		return nil, fmt.Errorf(errorFailedToParseSpecExcel, errorNilFile)
	}

	spec = &rs.Spec{}

	// 0. get rows of necessary sheets
	indexRows, err := file.GetRows(sheetSpecIndex)
	if err != nil {
		return nil, fmt.Errorf(errorFailedToParseSpecExcel, err)
	}
	configRows, err := file.GetRows(sheetSpecConfig)
	if err != nil {
		return nil, fmt.Errorf(errorFailedToParseSpecExcel, err)
	}

	// 1. parse necessary sheets
	// 1.1 index
	var productType rs.ProductType
	if productType, err = parseSpecIndex(indexRows, spec); err != nil {
		return nil, fmt.Errorf("failed to parse sheet Spec: err=%w", err)
	}

	// 1.2 config
	if err = parseSpecConfigSheet(configRows, productType, spec); err != nil {
		return nil, fmt.Errorf("failed to parse sheet Config: err=%w", err)
	}

	return
}
