package excel

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	utilRecipe "gitlab.kenda.com.tw/kenda/commons/v2/util/rs/recipe"

	"github.com/360EntSecGroup-Skylar/excelize/v2"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs/container"
	utilRS "gitlab.kenda.com.tw/kenda/commons/v2/util/rs"
)

func sheetPrefixToProcessType(prefix string) rs.ProcessType {
	switch prefix {
	case "P":
		return rs.ProcessType_PROCESS
	case "I":
		return rs.ProcessType_INSPECTION
	case "M":
		return rs.ProcessType_MEASUREMENT
	default:
		return rs.ProcessType_PROCESS_TYPE_UNSPECIFIED
	}
}

const errorParseRecipeProcessSheet = "failed to parse process sheet: %w"

func parseRecipeProcessList(rows [][]string, r *rs.Recipe) error {
	if len(rows) == 0 {
		return fmt.Errorf(errorParseRecipeProcessSheet, errorNilRow)
	}

	var standardStart, optionalStart bool
	var standardSection [][]string
	var optionalSection [][]string
	var optionalSections [][][]string
	{ // 0. classify rows into two sections: standard (standardSection) , optional (optionalSections)

		for key, row := range rows {
			rowIdx := key + 1
			// 0.1 standard
			if reflect.DeepEqual(row, columnHeaderFlowProcesses) && !optionalStart { // no more standard section after optional started
				standardStart = true
				optionalStart = false
				continue
			}
			if standardStart {
				if len(row) == 0 {
					standardStart = false
					continue
				}
				if len(row) != 4 {
					return fmt.Errorf("bad standard process section format: row index=%d", rowIdx)
				}
				standardSection = append(standardSection, row)
				continue
			}
			// 0.2 optional
			if reflect.DeepEqual(row, columnHeaderOptionalFlowInfo) {
				standardStart = false
				optionalStart = true
				if optionalSection != nil { // append optional flow when another optional flow start
					optionalSections = append(optionalSections, optionalSection)
					optionalSection = nil
				}
				continue
			}
			if optionalStart {
				if len(row) == 0 {
					optionalStart = false
					continue
				}
				if len(row) != 3 && len(row) != 4 {
					return fmt.Errorf("bad standard optional process section format: row index=%d", rowIdx)
				}
				if reflect.DeepEqual(row, columnHeaderFlowProcesses) { // skip header
					continue
				}

				optionalSection = append(optionalSection, row)

				if len(rows)-1 == key { // append last optional flow if this is last row
					if optionalSection != nil {
						optionalSections = append(optionalSections, optionalSection)
						optionalSection = nil
					}
				}

				continue
			}
		}
	}

	{ // 1. make standard process
		for _, row := range standardSection {
			r.Processes = append(r.Processes, &rs.Process{
				Name:    row[0],
				Type:    rs.ProcessType(rs.ProcessType_value[row[1]]),
				Configs: []*rs.Process_Config{},
			})
		}
	}

	{ // 2. make optional process
		for _, optionalSection := range optionalSections {
			var optionalFlow rs.Recipe_OptionalFlow
			for key, row := range optionalSection {
				if key == 0 { // first row -> optional process info
					if row[1][1:2] != "-" {
						return fmt.Errorf("bad after process=%s", row[1])
					}
					processType := sheetPrefixToProcessType(strings.Split(row[1], "-")[0])
					if processType == rs.ProcessType_PROCESS_TYPE_UNSPECIFIED {
						return fmt.Errorf("bad after process type=%s", row[1])
					}

					maxRepetitions, err := strconv.Atoi(row[2])
					if err != nil {
						return fmt.Errorf("bad max repetitions=%s", row[2])
					}
					optionalFlow = rs.Recipe_OptionalFlow{
						Name:             row[0],
						AfterProcessName: row[1][2:],
						AfterProcessType: processType,
						MaxRepetitions:   int32(maxRepetitions),
					}
				} else { // other rows -> process name, type
					// TODO: get product ID of each process from excel
					optionalFlow.Processes = append(optionalFlow.Processes, &rs.Process{
						Name:    row[0],
						Type:    rs.ProcessType(rs.ProcessType_value[row[1]]),
						Configs: []*rs.Process_Config{},
					})
				}
			}
			r.OptionalFlows = append(r.OptionalFlows, &optionalFlow) // every each optionalSection == optional flow
		}
	}
	return nil
}

func checkProcessSheetExists(f *excelize.File, processes []*rs.Process) (names []string, err error) {
	for _, process := range processes {
		name, err := processSheetName(process.GetType(), process.GetName())
		if err != nil {
			return nil, fmt.Errorf("failed to check process sheet exists: %w", err)
		}
		if f.GetSheetIndex(name) == -1 {
			return nil, fmt.Errorf("sheet=%s not exists", name)
		}
		names = append(names, name)
	}
	return names, nil
}

const errorParseRecipeSheetIndex = "failed to parse recipe index sheet: %w"

func parseRecipeIndex(rows [][]string, r *rs.Recipe) (productType rs.ProductType, err error) {
	if len(rows) == 0 {
		err = fmt.Errorf(errorParseRecipeSheetIndex, errorNilRow)
		return
	}

	r.Version = &rs.Version{}
	for _, row := range rows {
		if len(row) <= 1 { // row has no columns or has title only
			continue
		}

		title, values := row[0], row[1:]
		switch title {
		case "Recipe ID":
			r.Id = values[0]
		case "Product ID":
			r.ProductId = values[0]
		case "Product Type":
			if pt, ok := rs.ProductType_value[strings.ToUpper(values[0])]; !ok || pt == int32(rs.ProductType_PRODUCT_TYPE_UNSPECIFIED) {
				err = fmt.Errorf(errorParseRecipeSheetIndex, fmt.Errorf("bad product type=%s", values[0]))
				return
			} else {
				productType = rs.ProductType(pt)
				r.ProductType = strconv.Itoa(int(pt))
			}
		case "Factories":
			for _, value := range values {
				if factory, ok := rs.FactoryID_value[strings.ToUpper(value)]; !ok || factory == int32(rs.FactoryID_FACTORY_UNSPECIFIED) {
					err = fmt.Errorf(errorParseRecipeSheetIndex, fmt.Errorf("bad factory=%s", value))
					return
				} else {
					r.Factory = append(r.Factory, rs.FactoryID(factory))
				}
			}
		case "Major Version":
			r.Version.Major = values[0]
		case "Minor Version":
			r.Version.Minor = values[0]
		case "Parent Spec ID":
			r.ParentSpecId = values[0]
		case "Stage":
			if stage, ok := rs.StageStatus_value[strings.ToUpper(values[0])]; !ok || stage == int32(rs.StageStatus_STAGE_STATUS_UNSPECIFIED) {
				err = fmt.Errorf(errorParseRecipeSheetIndex, fmt.Errorf("bad stage=%s", values[0]))
				return
			} else {
				r.Version.StageStatus = rs.StageStatus(stage)
			}
		}
	}
	return
}

func parseTools(rows map[int][]string) ([]*rs.Tool, error) {
	var tools []*rs.Tool

	var keys []int
	for k := range rows {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, key := range keys {
		rowIdx := key + 1
		row := rows[key]
		if reflect.DeepEqual(row, columnHeaderTools) { // skip title
			continue
		}
		if len(row) == 0 {
			continue // skip empty row
		}
		if len(row) != 3 {
			return nil, fmt.Errorf("bad tool section format: row index=%d", rowIdx)
		}

		var tool rs.Tool
		if len(row[0]) == 0 {
			return nil, fmt.Errorf("missing tool ID: row index=%d", rowIdx)
		} else {
			tool.Id = row[0]
		}
		if len(row[1]) == 0 {
			return nil, fmt.Errorf("missing tool role: row index=%d", rowIdx)
		} else {
			toolRole, ok := rs.ToolRole_value[strings.ToUpper(row[1])]
			if !ok {
				return nil, fmt.Errorf("bad tool role:%s", row[1])
			}
			tool.Role = rs.ToolRole(toolRole)
		}
		if len(row[2]) == 0 {
			return nil, fmt.Errorf("missing tool check: row index=%d", rowIdx)
		} else {
			check, err := strconv.ParseBool(row[2])
			if err != nil {
				return nil, fmt.Errorf("bad tool check:%s, row index=%d", row[3], rowIdx)
			} else {
				tool.CheckFlag = check
			}
		}
		tools = append(tools, &tool)
	}
	return tools, nil
}

const errorParseMaterials = "failed to parse materials: %w"

func parseMaterials(productType rs.ProductType, rows map[int][]string) (map[int][]*rs.Material, error) {
	materials := make(map[int][]*rs.Material)
	var currentStep int

	var keys []int
	for k := range rows {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, key := range keys {
		rowIdx := key + 1
		row := rows[key]
		if reflect.DeepEqual(row, columnHeaderRecipeMaterials) { // skip title
			continue
		}
		if len(row) == 0 { // skip empty row
			continue
		}
		if len(row) < 10 {
			return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("bad section format: row index=%d", rowIdx))
		}

		step, err := strconv.Atoi(row[0])
		if err != nil || step <= 0 {
			return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("bad step=%s, row index=%d", row[0], rowIdx))
		}

		if (productType == rs.ProductType_RUBBER ||
			productType == rs.ProductType_ACCELERATOR_INGREDIENT ||
			productType == rs.ProductType_COMPOUNDING_INGREDIENT) && (len(row[1]) == 5) {
			if row[1][0:2] != row[3] { // material name should have type prefix
				return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("material name=%s and type=%s mismatch", row[1], row[3]))
			}
		}
		if row[3] == "202" && !strings.Contains(strings.ToUpper(row[1]), "CMP") ||
			row[3] != "202" && strings.Contains(strings.ToUpper(row[1]), "CMP") { // name of type 202 should contains "CMP"
			return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("material name=%s and type=%s mismatch", row[1], row[3]))
		}
		if row[3] == "203" && !strings.Contains(strings.ToUpper(row[1]), "ACL") || // name of type 203 should contains "ACL"
			row[3] != "203" && strings.Contains(strings.ToUpper(row[1]), "ACL") {
			return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("material name=%s and type=%s mismatch", row[1], row[3]))
		}

		if strings.ToUpper(row[2]) != "A" && strings.ToUpper(row[2]) != "B" && row[2] != "" {
			return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("bad material level=%s, row index=%d", row[2], rowIdx))
		}

		if len(row[1]) == 5 && productType == rs.ProductType_RUBBER && // raw material in mixing must have level
			row[2] != "A" && row[2] != "B" {
			return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("bad material=%s level=%s: row index=%d", row[1], row[2], rowIdx))
		}

		if t, err := strconv.Atoi(row[3]); err != nil {
			return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("bad material main type: err=%w, row index=%d", err, rowIdx))
		} else {
			var ok bool
			types := utilRS.ProductMaterialTypes[productType]
			for _, v := range types {
				if v == rs.ProductType(t) {
					ok = true
				}
			}
			if !ok {
				return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("bad material main type: product=%s, type=%d row index=%d", productType, t, rowIdx))
			}
		}

		// TODO: check recipe ID exists ?

		var ct container.Type
		if len(row[5]) > 0 {
			if containerTypes, ok := utilRecipe.ProductContainerTypes[productType]; !ok {
				return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("missing product type=%s mapping list, row index=%d", productType, rowIdx))
			} else {
				var ok bool
				for _, t := range containerTypes {
					if t.String() == row[5] {
						ok = true
						ct = t
					}
				}
				if !ok {
					return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("bad container type=%s of product type=%s mapping list, row index=%d", row[5], productType, rowIdx))
				}
			}
		}

		if len(row[6]) == 0 && len(row[7]) == 0 && len(row[8]) == 0 {
			return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("bad material quantity, low=%s mid=%s high=%s, row index=%d", row[6], row[7], row[8], rowIdx))
		}

		if len(row[6]) > 0 {
			if _, err := strconv.ParseFloat(row[6], 32); err != nil {
				return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("bad low material quantity=%s, err=%w row index=%d", row[6], err, rowIdx))
			}
		}
		if len(row[7]) > 0 {
			if _, err := strconv.ParseFloat(row[7], 32); err != nil {
				return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("bad mid material quantity=%s, err=%w row index=%d", row[6], err, rowIdx))
			}
		}
		if len(row[8]) > 0 {
			if _, err := strconv.ParseFloat(row[8], 32); err != nil {
				return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("bad high material quantity=%s, err=%w row index=%d", row[6], err, rowIdx))
			}
		}
		if len(row[9]) == 0 {
			return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("missing material unit, row index=%d", rowIdx))
		}
		param, err := stringToParamValues(row[6], row[7], row[8], row[9])
		if err != nil {
			return nil, fmt.Errorf(errorParseMaterials, fmt.Errorf("failed to parse param: err=%w", err))
		}

		material := &rs.Material{
			Name:          row[1],
			Level:         strings.ToUpper(row[2]),
			Type:          row[3],
			ReqRecipeId:   row[4],
			ContainerType: ct,
			Param:         param,
		}

		if currentStep == 0 || // init
			currentStep > 0 && currentStep != step { // new step start
			currentStep = step
		}
		materials[currentStep] = append(materials[currentStep], material)
	}
	return materials, nil
}

func parseStepsControls(items, units []string, stepsValues map[int][]string) (map[int][]*rs.Property, error) {
	stepsProperties := map[int][]*rs.Property{}
	for step, values := range stepsValues {
		var properties []*rs.Property
		for key := 0; key < len(values); key++ { // count of values could less than items
			if values[key] == "" {
				continue // skip item without value
			}
			if key+0 > len(values) {
				break
			}
			if value, err := utilRS.PropertySettingNameToValue(items[key], values[key]); err == nil { // enum value
				properties = append(properties, &rs.Property{
					Name: items[key],
					Param: &rs.Param{
						Type: rs.Param_ENUM,
						Unit: units[key],
						Enum: &rs.Param_Enum{Value: value},
					},
				})
			} else { // float values
				var low, mid, high string
				vs := strings.Split(values[key], ";")
				for _, value := range vs {
					switch {
					case strings.HasPrefix(value, "low:"):
						low = value[4:]
					case strings.HasPrefix(value, "mid:"):
						mid = value[4:]
					case strings.HasPrefix(value, "high:"):
						high = value[5:]
					default:
						return nil, fmt.Errorf("bad value format: step=%d, item=%s, value=%s", step, items[key], values[key])
					}
				}
				if len(units[key]) == 0 {
					return nil, fmt.Errorf("failed to property param, step=%d, item=%s, value=%s, err=%w", step, items[key], values[key], errors.New("missing unit"))
				}
				param, err := stringToParamValues(low, mid, high, units[key])
				if err != nil {
					return nil, fmt.Errorf("failed to property param, step=%d, item=%s, value=%s, err=%w", step, items[key], values[key], err)
				}
				properties = append(properties, &rs.Property{
					Name:  items[key],
					Param: param,
				})
			}
		}
		stepsProperties[step] = properties
	}
	return stepsProperties, nil
}

const errorParseControls = "failed to parse controls: %w"

func parseControls(rows map[int][]string) (map[int][]*rs.Property, error) {
	var items, units []string
	values := map[int][]string{}
	var currentStep int

	var keys []int
	for k := range rows {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, key := range keys {
		rowIdx := key + 1
		row := rows[key]
		if row[0] == "CONTROL" {
			items = append(items, row[1:]...) // get all control items from title
			continue
		}
		if row[0] == "UNIT" {
			units = append(units, row[1:]...) // get all control items from sub title
			continue
		}

		if len(row) < 2 {
			return nil, fmt.Errorf(errorParseControls, fmt.Errorf("bad section format, row index=%d", rowIdx))
		}
		step, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, fmt.Errorf(errorParseControls, fmt.Errorf("bad step=%s", row[0]))
		}

		if currentStep == 0 || // init
			currentStep > 0 && currentStep != step { // new step start
			currentStep = step
		}
		values[currentStep] = append(values[currentStep], row[1:]...)
	}

	if len(values) == 0 && len(items) > 0 && len(units) > 0 {
		return nil, fmt.Errorf(errorParseControls, errors.New("missing values"))
	}
	controls, err := parseStepsControls(items, units, values)
	if err != nil {
		return nil, fmt.Errorf(errorParseControls, err)
	}
	return controls, nil
}

const errorParseMeasurements = "failed to parse measurements: %w"

func parseMeasurements(rows map[int][]string) (map[int][]*rs.Property, error) {
	measurements := make(map[int][]*rs.Property)
	var currentStep int

	var keys []int
	for k := range rows {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, key := range keys {
		rowIdx := key + 1
		row := rows[key]
		if reflect.DeepEqual(row, columnHeaderMeasurements) { // skip title
			continue
		}
		if len(row) == 0 { // skip empty row
			continue
		}
		if len(row) != 6 {
			return nil, fmt.Errorf(errorParseMeasurements, fmt.Errorf("bad section format, row index=%d", rowIdx))
		}
		step, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, fmt.Errorf(errorParseMeasurements, fmt.Errorf("bad step=%s: err=%w", row[0], err))
		}
		if len(row[1]) == 0 {
			return nil, fmt.Errorf(errorParseMeasurements, fmt.Errorf("missing item name, row index=%d", rowIdx))
		}
		if len(row[2]) == 0 && len(row[3]) == 0 && len(row[4]) == 0 {
			return nil, fmt.Errorf(errorParseMeasurements, fmt.Errorf("missing item value, low=%s mid=%s high=%s, row index=%d", row[2], row[3], row[4], rowIdx))
		}
		if len(row[5]) == 0 {
			return nil, fmt.Errorf(errorParseMeasurements, fmt.Errorf("missing item unit, row index=%d", rowIdx))
		}
		param, err := stringToParamValues(row[2], row[3], row[4], row[5])
		if err != nil {
			return nil, fmt.Errorf(errorParseMeasurements, err)
		}
		measurement := &rs.Property{
			Name:  row[1],
			Param: param,
		}
		if currentStep == 0 || // init
			step > 0 && step != currentStep { // new measurement item start
			currentStep = step
		}
		measurements[currentStep] = append(measurements[currentStep], measurement)
	}
	return measurements, nil
}

func parseProcessSheets(file *excelize.File, productType rs.ProductType, names []string) (processes []*rs.Process, err error) {
	if productType == rs.ProductType_PRODUCT_TYPE_UNSPECIFIED {
		return nil, errorProductTypeUnspecified
	}

	processes = []*rs.Process{}
	configs := []*rs.Process_Config{}
	for key, name := range names {
		rows, err := file.GetRows(name)
		if err != nil {
			return nil, err
		}
		if configs, err = parseProcessSheet(rows, productType); err != nil {
			return nil, fmt.Errorf("failed to parse process sheet=%s, err: %w", name, err)
		}
		processes = append(processes, &rs.Process{
			Name: name[2:],
			Type: sheetPrefixToProcessType(name[0:1]),
		})
		processes[key].Configs = configs
	}
	return processes, nil
}

func parseProcessSheet(rows [][]string, productType rs.ProductType) (configs []*rs.Process_Config, err error) {
	if rows == nil {
		return
	}

	type configSection map[string]map[int][]string // section name -> row index -> row
	type configSections []configSection

	var name string
	section := configSection{}
	var sections configSections
	for key, row := range rows {
		rowIdx := key + 1
		if len(row) == 0 { // skip empty row
			continue
		}
		if row[0] == "STATIONS" {
			name = "STATIONS"

			if section[name] != nil { // append to sections last section when new section start
				newSection := configSection{} // copy to new map to avoid value overwrite problem
				for k, v := range section {
					newSection[k] = v
				}
				sections = append(sections, newSection)
			}

			section[name] = map[int][]string{} // new section start
			section[name][0] = row[1:]         // stations
			continue
		}
		if reflect.DeepEqual(row, columnHeaderTools) {
			name = columnHeaderTools[0]
			section[name] = map[int][]string{}
			continue
		}
		if reflect.DeepEqual(row, columnHeaderRecipeMaterials) {
			name = columnHeaderRecipeMaterials[0]
			section[name] = map[int][]string{}
			continue
		}
		if row[0] == "CONTROL" {
			name = "CONTROL"
			section[name] = map[int][]string{}
			section[name][key] = row // control items
			continue
		}
		if reflect.DeepEqual(row, columnHeaderMeasurements) {
			name = columnHeaderMeasurements[0]
			section[name] = map[int][]string{}
			continue
		}

		switch name {
		case "STATIONS":
			if len(row) < 2 {
				return nil, fmt.Errorf("bad %s section format, row index=%d", name, rowIdx)
			}
		case columnHeaderTools[0]:
			if len(row) != 3 {
				return nil, fmt.Errorf("bad %s section format, row index=%d", name[0:4], rowIdx)
			}
		case columnHeaderRecipeMaterials[0]:
			if len(row) != 10 {
				return nil, fmt.Errorf("bad %s section format, row index=%d", name, rowIdx)
			}
		case "CONTROL":
			if len(row) < 2 && row[0] != "CONTROL" && row[0] != "UNIT" {
				return nil, fmt.Errorf("bad %s section format, row index=%d", name, rowIdx)
			}
		case columnHeaderMeasurements[0]:
			if len(row) != 6 {
				return nil, fmt.Errorf("bad %s section format, row index=%d", name, rowIdx)
			}
		default: // continue if no section name
			continue
		}
		section[name][key] = row
	}
	if section["STATIONS"] != nil { // append the last section if not nil
		newSection := configSection{}
		for k, v := range section {
			newSection[k] = v
		}
		sections = append(sections, newSection)
	}

	var tools []*rs.Tool
	var materials map[int][]*rs.Material
	var controls map[int][]*rs.Property
	var measurements map[int][]*rs.Property
	for _, section := range sections {
		if t, ok := section[columnHeaderTools[0]]; ok {
			if tools, err = parseTools(t); err != nil {
				return nil, err
			}
		}
		if m, ok := section[columnHeaderRecipeMaterials[0]]; ok {
			if materials, err = parseMaterials(productType, m); err != nil {
				return nil, err
			}
		}
		if c, ok := section["CONTROL"]; ok {
			if controls, err = parseControls(c); err != nil {
				return nil, err
			}
		}
		if m, ok := section[columnHeaderMeasurements[0]]; ok {
			if measurements, err = parseMeasurements(m); err != nil {
				return nil, err
			}
		}
		config := &rs.Process_Config{
			Stations: section["STATIONS"][0], // always be zero
			Tools:    tools,
		}
		for step := 1; err == nil; step++ {
			var hasStep bool
			if _, ok := materials[step]; ok {
				hasStep = true
			}
			if _, ok := controls[step]; ok {
				hasStep = true
			}
			if _, ok := measurements[step]; ok {
				hasStep = true
			}
			if !hasStep {
				err = errors.New("EOF")
				continue
			}
			config.Steps = append(config.Steps, &rs.Process_Config_Step{
				Materials:    materials[step],
				Controls:     controls[step],
				Measurements: measurements[step],
			})
		}
		configs = append(configs, config)
	}

	return configs, nil
}

const errorFailedToParseRecipeExcel = "failed to parse recipe excel: err=%w"

// ToRecipe parse excel file to recipe
func ToRecipe(f *excelize.File) (recipe *rs.Recipe, err error) {
	if f == nil {
		return nil, fmt.Errorf(errorFailedToParseRecipeExcel, errorNilFile)
	}

	recipe = &rs.Recipe{}

	// 0. get rows of necessary sheets
	indexRows, err := f.GetRows(sheetRecipeIndex)
	if err != nil {
		return nil, fmt.Errorf(errorFailedToParseRecipeExcel, err)
	}
	processRows, err := f.GetRows(sheetProcessIndex)
	if err != nil {
		return nil, fmt.Errorf(errorFailedToParseRecipeExcel, err)
	}

	// 1. parse necessary sheets
	var productType rs.ProductType
	var processes []*rs.Process
	if productType, err = parseRecipeIndex(indexRows, recipe); err != nil {
		return nil, fmt.Errorf("failed to parse sheet Recipe: err=%w", err)
	}
	if err = parseRecipeProcessList(processRows, recipe); err != nil {
		return nil, fmt.Errorf("failed to parse sheet Process: err=%w", err)
	}

	// 2 check and parse process sheets (standard & optional)
	if names, err := checkProcessSheetExists(f, recipe.GetProcesses()); err != nil {
		return nil, err
	} else {
		if processes, err = parseProcessSheets(f, productType, names); err != nil {
			return nil, err
		}
		recipe.Processes = processes
	}
	for _, flow := range recipe.GetOptionalFlows() {
		if names, err := checkProcessSheetExists(f, flow.GetProcesses()); err != nil {
			return nil, err
		} else {
			if processes, err = parseProcessSheets(f, productType, names); err != nil {
				return nil, err
			}
			flow.Processes = processes
		}
	}
	return recipe, nil
}
