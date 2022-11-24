package excel

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
	utilRS "gitlab.kenda.com.tw/kenda/commons/v2/util/rs"
	utilRecipe "gitlab.kenda.com.tw/kenda/commons/v2/util/rs/recipe"
)

var (
	keysLevel       = []string{"A", "B"}
	keysBool        = []string{"TRUE", "FALSE"}
	keysStage       = fromEnumsToKeys(rs.StageStatus_name)
	keysProcessType = fromEnumsToKeys(rs.ProcessType_name)
	keysToolRole    = fromEnumsToKeys(rs.ToolRole_name)
)

// fromEnumNamesToKeys return enum names or values as slice of string,
// skip name with "UNSPECIFIED" suffix as default
func fromEnumsToKeys(names map[int32]string, skips ...int32) []string {
	var keys []int
	for key, name := range names {
		if strings.HasSuffix(name, "UNSPECIFIED") {
			continue
		}

		keys = append(keys, int(key))
		for _, skip := range skips {
			if key == skip {
				keys = keys[:len(keys)-1]
			}
		}
	}
	sort.Ints(keys)

	out := make([]string, len(keys))
	for i, key := range keys {
		out[i] = names[int32(key)]
	}
	return out
}

func keysContainerTypes(product rs.ProductType) (types []string) {
	ts := utilRecipe.ProductContainerTypes[product]
	if ts == nil {
		return types
	}
	for _, t := range ts {
		types = append(types, t.String())
	}
	return types
}

func keysMaterialTypes(product rs.ProductType) (types []string) {
	ts := utilRS.ProductMaterialTypes[product]
	if ts == nil {
		return types
	}
	for _, t := range ts {
		types = append(types, strconv.Itoa(int(t))) // use enum value as keys
	}
	return types
}

const errorSetDropDownList = "failed to set dropdown list: %w"

func setDropDownList(file *excelize.File, sheet string, axis string, keys []string) (err error) {
	if file == nil {
		return fmt.Errorf(errorSetDropDownList, errorNilFile)
	}
	if file.GetSheetIndex(sheet) == -1 {
		return fmt.Errorf(errorSetDropDownList, errors.New("sheet not exist"))
	}
	if len(axis) == 0 {
		return fmt.Errorf(errorSetDropDownList, errors.New("missing axis"))
	}
	if len(keys) == 0 {
		return fmt.Errorf(errorSetDropDownList, errors.New("missing keys"))
	}

	dv := excelize.NewDataValidation(false)
	dv.Sqref = axis
	if err := dv.SetDropList(keys); err != nil {
		return fmt.Errorf(errorSetDropDownList, err)
	}
	if err := file.AddDataValidation(sheet, dv); err != nil {
		return fmt.Errorf(errorSetDropDownList, err)
	}
	return nil
}
