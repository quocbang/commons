package excel

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
	utilRS "gitlab.kenda.com.tw/kenda/commons/v2/util/rs"
)

const errorWriteSpecIndex = "failed to write spec index: %w"

func writeSpecIndexSheet(f *excelize.File, productType rs.ProductType, r *rs.Spec) (err error) {
	if f == nil {
		return fmt.Errorf(errorWriteSpecIndex, errorNilFile)
	}
	if r == nil {
		return fmt.Errorf(errorWriteSpecIndex, errorNilSpec)
	}
	f.SetSheetName(sheetDefault, sheetSpecIndex) // Replace the default "Sheet1" created by NewFile
	w, err := f.NewStreamWriter(sheetSpecIndex)
	if err != nil {
		return fmt.Errorf(errorWriteSpecIndex, err)
	}

	if err := w.SetRow("A1", []interface{}{"Spec ID", r.GetId()}); err != nil {
		return fmt.Errorf(errorWriteSpecIndex, err)
	}
	if err := w.SetRow("A2", []interface{}{"Product ID", r.GetProductId()}); err != nil {
		return fmt.Errorf(errorWriteSpecIndex, err)
	}
	if err := w.SetRow("A3", productTypeRow(productType)); err != nil {
		return fmt.Errorf(errorWriteSpecIndex, err)
	}
	if err := w.SetRow("A4", factoriesRow(r.GetFactory())); err != nil {
		return fmt.Errorf(errorWriteSpecIndex, err)
	}
	if err := w.SetRow("A5", []interface{}{"Major Version", r.GetVersion().GetMajor()}); err != nil {
		return fmt.Errorf(errorWriteSpecIndex, err)
	}
	if err := w.SetRow("A6", []interface{}{"Minor Version", r.GetVersion().GetMinor()}); err != nil {
		return fmt.Errorf(errorWriteSpecIndex, err)
	}
	if err := w.SetRow("A7", stageStatusRow(r.GetVersion().GetStageStatus())); err != nil {
		return fmt.Errorf(errorWriteSpecIndex, err)
	}
	if err := setDropDownList(f, sheetSpecIndex, "B7", keysStage); err != nil {
		return fmt.Errorf(errorWriteSpecIndex, err)
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf(errorWriteSpecIndex, err)
	}
	return nil
}

func writeSpecMaterialsSection(w *excelize.StreamWriter, row *int, style int, productType rs.ProductType, materials []*rs.Material) (err error) {
	if w == nil {
		return fmt.Errorf(errorWriteMaterialsSection, errorNilStreamWriter)
	}
	if row == nil {
		return fmt.Errorf(errorWriteMaterialsSection, errorNilRow)
	}
	if productType == rs.ProductType_PRODUCT_TYPE_UNSPECIFIED {
		return fmt.Errorf(errorWriteMaterialsSection, errorProductTypeUnspecified)
	}

	if err := w.SetRow(fmt.Sprintf("A%d", *row), cellsWithStyle(style, columnHeaderSpecMaterials...)); err != nil {
		return fmt.Errorf(errorWriteMaterialsSection, err)
	}
	*row++

	for _, material := range materials {
		param := material.GetParam()
		low, _ := constraintToString(param.GetMin())
		mid, _ := constraintToString(param.GetCentral())
		high, _ := constraintToString(param.GetMax())
		if err := w.SetRow(
			fmt.Sprintf("A%d", *row),
			[]interface{}{
				material.GetName(), material.GetLevel(), material.GetType(), material.GetReqRecipeId(),
				containerToString(material.GetContainerType()),
				low, mid, high, material.GetParam().GetUnit(),
			}); err != nil {
			return fmt.Errorf(errorWriteMaterialsSection, err)
		}
		if err := setDropDownList(w.File, w.Sheet, fmt.Sprintf("B%d", *row), keysLevel); err != nil {
			return fmt.Errorf(errorWriteMaterialsSection, err)
		}
		if keys := keysMaterialTypes(productType); len(keys) > 0 {
			if err := setDropDownList(w.File, w.Sheet, fmt.Sprintf("C%d", *row), keys); err != nil {
				return fmt.Errorf(errorWriteMaterialsSection, err)
			}
		}
		if keys := keysContainerTypes(productType); len(keys) > 0 {
			if err := setDropDownList(w.File, w.Sheet, fmt.Sprintf("E%d", *row), keys); err != nil {
				return fmt.Errorf(errorWriteMaterialsSection, err)
			}
		}
		*row++
	}
	*row++ // section divider
	return nil
}

const errorWriteSpecPropertiesSection = "failed to write spec properties section: %w"

func writeSpecPropertiesSection(w *excelize.StreamWriter, row *int, style int, productType rs.ProductType, properties []*rs.Property) (err error) {
	if w == nil {
		return fmt.Errorf(errorWriteSpecPropertiesSection, errorNilStreamWriter)
	}
	if row == nil {
		return fmt.Errorf(errorWriteSpecPropertiesSection, errorNilRow)
	}
	if productType == rs.ProductType_PRODUCT_TYPE_UNSPECIFIED {
		return fmt.Errorf(errorWriteSpecPropertiesSection, errorProductTypeUnspecified)
	}

	if err := w.SetRow(fmt.Sprintf("A%d", *row), cellsWithStyle(style, columnHeaderProperties...)); err != nil {
		return fmt.Errorf(errorWriteSpecPropertiesSection, err)
	}
	*row++

	var enumValue string
	for _, property := range properties {
		if enum := property.GetParam().GetEnum(); enum != nil {
			if value := enum.GetValue(); value != 0 {
				enumValue, err = utilRS.PropertySettingValueToName(property.GetName(), value)
				if err != nil {
					return fmt.Errorf(errorWriteSpecPropertiesSection, err)
				}
			}
			// set dropdown list
			if err := setDropDownList(w.File, w.Sheet, fmt.Sprintf("E%d", *row), fromEnumsToKeys(utilRS.PropertySettingNameEnum[property.GetName()])); err != nil {
				return fmt.Errorf(errorWriteSpecPropertiesSection, err)
			}
		}
		param := property.GetParam()
		low, _ := constraintToString(param.GetMin())
		mid, _ := constraintToString(param.GetCentral())
		high, _ := constraintToString(param.GetMax())
		if err := w.SetRow(
			fmt.Sprintf("A%d", *row),
			[]interface{}{
				property.GetName(),
				low,
				mid,
				high,
				enumValue,
				property.GetParam().GetUnit(),
			},
		); err != nil {
			return fmt.Errorf(errorWriteSpecPropertiesSection, err)
		}
		*row++
	}
	*row++ // section divider
	return nil
}

const errorWriteSpecConfigsSheet = "failed to write spec configs sheet: %w"

func writeSpecConfigsSheet(f *excelize.File, styles *styles, productType rs.ProductType, r *rs.Spec) (err error) {
	if f == nil {
		return fmt.Errorf(errorWriteSpecConfigsSheet, errorNilFile)
	}
	if styles == nil {
		return fmt.Errorf(errorWriteSpecConfigsSheet, errorNilStyles)
	}
	if productType == rs.ProductType_PRODUCT_TYPE_UNSPECIFIED {
		return fmt.Errorf(errorWriteSpecConfigsSheet, errorProductTypeUnspecified)
	}
	if r == nil {
		return fmt.Errorf(errorWriteSpecConfigsSheet, errorNilSpec)
	}

	f.NewSheet(sheetSpecConfig)
	w, err := f.NewStreamWriter(sheetSpecConfig)
	if err != nil {
		return fmt.Errorf(errorWriteSpecConfigsSheet, err)
	}

	row := 1

	if r.GetTools() != nil {
		if err := writeToolsSection(w, &row, styles.warning, r.GetTools()); err != nil {
			return fmt.Errorf(errorWriteSpecConfigsSheet, err)
		}
	}
	if r.GetMaterials() != nil {
		if err := writeSpecMaterialsSection(w, &row, styles.warning, productType, r.GetMaterials()); err != nil {
			return fmt.Errorf(errorWriteSpecConfigsSheet, err)
		}
	}
	if r.GetProperties() != nil {
		if err := writeSpecPropertiesSection(w, &row, styles.warning, productType, r.GetProperties()); err != nil {
			return fmt.Errorf(errorWriteSpecConfigsSheet, err)
		}
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf(errorWriteSpecConfigsSheet, err)
	}
	return nil
}

// FromSpec export excel file from spec
func FromSpec(r *rs.Spec) (buf *bytes.Buffer, err error) {
	if r == nil {
		return nil, errorNilSpec
	}

	f := excelize.NewFile()

	styles, err := setDefaultStyles(f)
	if err != nil {
		return nil, err
	}

	pt, err := strconv.Atoi(r.GetProductType())
	if err != nil {
		return nil, fmt.Errorf("fail to convert product type to integer representation=%s: %w", r.GetProductType(), err)
	}
	productType := rs.ProductType(pt)

	if err = writeSpecIndexSheet(f, productType, r); err != nil {
		return nil, err
	}

	if err = writeSpecConfigsSheet(f, styles, productType, r); err != nil {
		return nil, err
	}

	return f.WriteToBuffer()
}
