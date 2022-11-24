package excel

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs/container"
	utilRS "gitlab.kenda.com.tw/kenda/commons/v2/util/rs"
)

func productTypeRow(product rs.ProductType) []interface{} {
	row := make([]interface{}, 2)
	row[0] = "Product Type"
	if product != rs.ProductType_PRODUCT_TYPE_UNSPECIFIED {
		row[1] = rs.ProductType_name[int32(product)]
	}
	return row
}

func factoriesRow(fids []rs.FactoryID) []interface{} {
	row := make([]interface{}, len(fids)+1)
	row[0] = "Factories"
	ids := row[1:]
	for i, id := range fids {
		if id != rs.FactoryID_FACTORY_UNSPECIFIED {
			ids[i] = rs.FactoryID_name[int32(id)]
		}
	}
	return row
}

func stageStatusRow(stage rs.StageStatus) []interface{} {
	row := make([]interface{}, 2)
	row[0] = "Stage"
	if stage != rs.StageStatus_STAGE_STATUS_UNSPECIFIED {
		row[1] = rs.StageStatus_name[int32(stage)]
	}
	return row
}

const errorWriteRecipeIndex = "failed to write recipe index: %w"

func writeRecipeIndexSheet(f *excelize.File, productType rs.ProductType, r *rs.Recipe) (err error) {
	if f == nil {
		return fmt.Errorf(errorWriteRecipeIndex, errorNilFile)
	}
	if r == nil {
		return fmt.Errorf(errorWriteRecipeIndex, errorNilRecipe)
	}
	f.SetSheetName(sheetDefault, sheetRecipeIndex) // Replace the default "Sheet1" created by NewFile
	w, err := f.NewStreamWriter(sheetRecipeIndex)
	if err != nil {
		return fmt.Errorf(errorWriteRecipeIndex, err)
	}

	if err := w.SetRow("A1", []interface{}{"Recipe ID", r.GetId()}); err != nil {
		return fmt.Errorf(errorWriteRecipeIndex, err)
	}
	if err := w.SetRow("A2", []interface{}{"Product ID", r.GetProductId()}); err != nil {
		return fmt.Errorf(errorWriteRecipeIndex, err)
	}
	if err := w.SetRow("A3", productTypeRow(productType)); err != nil {
		return fmt.Errorf(errorWriteRecipeIndex, err)
	}
	if err := w.SetRow("A4", factoriesRow(r.GetFactory())); err != nil {
		return fmt.Errorf(errorWriteRecipeIndex, err)
	}
	if err := w.SetRow("A5", []interface{}{"Major Version", r.GetVersion().GetMajor()}); err != nil {
		return fmt.Errorf(errorWriteRecipeIndex, err)
	}
	if err := w.SetRow("A6", []interface{}{"Minor Version", r.GetVersion().GetMinor()}); err != nil {
		return fmt.Errorf(errorWriteRecipeIndex, err)
	}
	if err := w.SetRow("A7", []interface{}{"Parent Spec ID", r.GetParentSpecId()}); err != nil {
		return fmt.Errorf(errorWriteRecipeIndex, err)
	}
	if err := w.SetRow("A8", stageStatusRow(r.GetVersion().GetStageStatus())); err != nil {
		return fmt.Errorf(errorWriteRecipeIndex, err)
	}
	if err := setDropDownList(f, sheetRecipeIndex, "B8", keysStage); err != nil {
		return fmt.Errorf(errorWriteRecipeIndex, err)
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf(errorWriteRecipeIndex, err)
	}
	return nil
}

const errorFailedToGetSheetName = "failed to get sheet name: process type=%s, name=%s "

func processSheetName(processType rs.ProcessType, name string) (string, error) {
	if processType == rs.ProcessType_PROCESS_TYPE_UNSPECIFIED {
		return "", fmt.Errorf(errorFailedToGetSheetName, processType, name)
	}
	if name == "" {
		return "", fmt.Errorf(errorFailedToGetSheetName, processType, name)
	}
	return fmt.Sprintf("%s-%s", processType.String()[0:1], name), nil
}

func stationsRow(style int, stations []string) []interface{} {
	row := make([]interface{}, len(stations)+1)
	row[0] = excelize.Cell{StyleID: style, Value: "STATIONS"}
	ss := row[1:]
	for i, station := range stations {
		ss[i] = excelize.Cell{StyleID: style, Value: station}
	}
	return row
}

func upperBool(b bool) string {
	if b {
		return "TRUE"
	}
	return "FALSE"
}

const errorWriteToolsSection = "failed to write tools section: %w"

func writeToolsSection(w *excelize.StreamWriter, row *int, style int, tools []*rs.Tool) (err error) {
	if w == nil {
		return fmt.Errorf(errorWriteToolsSection, errorNilStreamWriter)
	}
	if row == nil {
		return fmt.Errorf(errorWriteToolsSection, errorNilRow)
	}

	if err := w.SetRow(fmt.Sprintf("A%d", *row), cellsWithStyle(style, columnHeaderTools...)); err != nil {
		return fmt.Errorf(errorWriteToolsSection, err)
	}
	*row++
	for _, tool := range tools {
		if err = w.SetRow(fmt.Sprintf("A%d", *row), []interface{}{tool.GetId(), tool.GetRole(), upperBool(tool.GetCheckFlag())}); err != nil {
			return fmt.Errorf(errorWriteToolsSection, err)
		}
		if err = setDropDownList(w.File, w.Sheet, fmt.Sprintf("C%d", *row), keysToolRole); err != nil {
			return fmt.Errorf(errorWriteToolsSection, err)
		}
		if err = setDropDownList(w.File, w.Sheet, fmt.Sprintf("D%d", *row), keysBool); err != nil {
			return fmt.Errorf(errorWriteToolsSection, err)
		}
		*row++
	}
	*row++ // section divider
	return nil
}

func containerToString(ct container.Type) string {
	if ct == container.Type_CONTAINER_TYPE_UNSPECIFIED {
		return ""
	}
	return ct.String()
}

const errorWriteMaterialsSection = "failed to write materials section: %w"

func writeMaterialsSection(w *excelize.StreamWriter, row *int, style int, productType rs.ProductType, steps []*rs.Process_Config_Step) (err error) {
	if w == nil {
		return fmt.Errorf(errorWriteMaterialsSection, errorNilStreamWriter)
	}
	if row == nil {
		return fmt.Errorf(errorWriteMaterialsSection, errorNilRow)
	}
	if productType == rs.ProductType_PRODUCT_TYPE_UNSPECIFIED {
		return fmt.Errorf(errorWriteMaterialsSection, errorProductTypeUnspecified)
	}
	if err := w.SetRow(fmt.Sprintf("A%d", *row), cellsWithStyle(style, columnHeaderRecipeMaterials...)); err != nil {
		return fmt.Errorf(errorWriteMaterialsSection, err)
	}
	*row++

	for stepIndex, step := range steps {
		if step.GetMaterials() == nil {
			continue
		}
		for _, material := range step.GetMaterials() {
			param := material.GetParam()
			low, _ := constraintToString(param.GetMin())
			mid, _ := constraintToString(param.GetCentral())
			high, _ := constraintToString(param.GetMax())
			if err := w.SetRow(
				fmt.Sprintf("A%d", *row),
				[]interface{}{stepIndex + 1,
					material.GetName(), material.GetLevel(), material.GetType(), material.GetReqRecipeId(),
					containerToString(material.GetContainerType()),
					low, mid, high, material.GetParam().GetUnit(),
				},
			); err != nil {
				return fmt.Errorf(errorWriteMaterialsSection, err)
			}

			if err := setDropDownList(w.File, w.Sheet, fmt.Sprintf("C%d", *row), keysLevel); err != nil {
				return fmt.Errorf(errorWriteMaterialsSection, err)
			}
			if keys := keysMaterialTypes(productType); len(keys) > 0 {
				if err := setDropDownList(w.File, w.Sheet, fmt.Sprintf("D%d", *row), keys); err != nil {
					return fmt.Errorf(errorWriteMaterialsSection, err)
				}
			}
			if keys := keysContainerTypes(productType); len(keys) > 0 {
				if err := setDropDownList(w.File, w.Sheet, fmt.Sprintf("F%d", *row), keys); err != nil {
					return fmt.Errorf(errorWriteMaterialsSection, err)
				}
			}
			*row++
		}
	}
	*row++ // section divider
	return nil
}

const errorWriteControlsSection = "failed to write controls section: %w"

func writeControlsSection(w *excelize.StreamWriter, row *int, style int, steps []*rs.Process_Config_Step) (err error) {
	if w == nil {
		return fmt.Errorf(errorWriteControlsSection, errorNilStreamWriter)
	}
	if row == nil {
		return fmt.Errorf(errorWriteControlsSection, errorNilRow)
	}
	if len(steps) == 0 {
		return fmt.Errorf(errorWriteControlsSection, errorEmptySteps)
	}

	names := []string{"CONTROL"}
	units := []string{"UNIT"}
	dvs := map[string][]string{}
	for _, property := range steps[0].GetControls() { // first step has all control items of process
		names = append(names, property.GetName())
		units = append(units, property.GetParam().GetUnit())
		if dv := fromEnumsToKeys(utilRS.PropertySettingNameEnum[property.GetName()]); dv != nil {
			dvs[property.GetName()] = dv
		}
	}
	if err = w.SetRow(fmt.Sprintf("A%d", *row), cellsWithStyle(style, names...)); err != nil {
		return fmt.Errorf(errorWriteControlsSection, err)
	}
	*row++
	if err = w.SetRow(fmt.Sprintf("A%d", *row), cellsWithStyle(style, units...)); err != nil {
		return fmt.Errorf(errorWriteControlsSection, err)
	}
	*row++

	for stepIndex, step := range steps {
		if step.GetControls() == nil {
			continue
		}
		values := []interface{}{stepIndex + 1}
		for _, name := range names[1:] { // skip row title
			var exists bool
			for _, property := range step.GetControls() {
				if name != property.GetName() {
					continue
				}
				if value, err := propertyToString(property); err != nil {
					return fmt.Errorf(errorWriteControlsSection, err)
				} else {
					values = append(values, value)
					exists = true

					// set dropdown list if enum control
					if dv, ok := dvs[name]; ok && len(dv) > 0 {
						cellName, err := excelize.CoordinatesToCellName(len(values), *row) // col and row count from 1
						if err != nil {
							return fmt.Errorf(errorWriteControlsSection, err)
						}
						if err = setDropDownList(w.File, w.Sheet, cellName, dv); err != nil {
							return fmt.Errorf(errorWriteControlsSection, err)
						}
					}
				}
			}
			if !exists {
				values = append(values, "")
			}
		}
		if err := w.SetRow(fmt.Sprintf("A%d", *row), values); err != nil {
			return fmt.Errorf(errorWriteControlsSection, err)
		}
		*row++
	}
	*row++ // section divider
	return nil
}

const errorWriteMeasurementsSection = "failed to write measurements section: %w"

func writeMeasurementsSection(w *excelize.StreamWriter, row *int, style int, steps []*rs.Process_Config_Step) (err error) {
	if w == nil {
		return fmt.Errorf(errorWriteMeasurementsSection, errorNilStreamWriter)
	}
	if row == nil {
		return fmt.Errorf(errorWriteMeasurementsSection, errorNilRow)
	}
	if err := w.SetRow(fmt.Sprintf("A%d", *row), cellsWithStyle(style, columnHeaderMeasurements...)); err != nil {
		return fmt.Errorf(errorWriteMeasurementsSection, err)
	}
	*row++

	for stepIndex, step := range steps {
		if step.GetMeasurements() == nil {
			continue
		}
		for _, measurement := range step.GetMeasurements() {
			param := measurement.GetParam()
			low, _ := constraintToString(param.GetMin())
			mid, _ := constraintToString(param.GetCentral())
			high, _ := constraintToString(param.GetMax())
			if err := w.SetRow(fmt.Sprintf("A%d", *row),
				[]interface{}{
					stepIndex + 1,
					measurement.GetName(),
					low,
					mid,
					high,
					measurement.GetParam().GetUnit()},
			); err != nil {
				return fmt.Errorf(errorWriteMeasurementsSection, err)
			}
			*row++
		}
	}
	*row++ // section divider
	return nil
}

const errorWriteProcessSheet = "failed to write process sheet: %w"

func writeProcessSheet(f *excelize.File, processSheet string, styles *styles, productType rs.ProductType, process *rs.Process) (err error) {
	if f == nil {
		return fmt.Errorf(errorWriteProcessSheet, errorNilFile)
	}
	if len(processSheet) == 0 {
		return fmt.Errorf(errorWriteProcessSheet, errorEmptySheetName)
	}
	if styles == nil {
		return fmt.Errorf(errorWriteProcessSheet, errorNilStyles)
	}
	if process == nil {
		return fmt.Errorf(errorWriteProcessSheet, errorNilProcess)
	}
	if productType == rs.ProductType_PRODUCT_TYPE_UNSPECIFIED {
		return fmt.Errorf(errorWriteProcessSheet, errorProductTypeUnspecified)
	}
	if f.GetSheetIndex(processSheet) != -1 { // sheet name conflict is acceptable, optional process could use inspection process as same as standard process.
		return nil
	}

	f.NewSheet(processSheet)
	processWriter, err := f.NewStreamWriter(processSheet)
	if err != nil {
		return fmt.Errorf(errorWriteProcessSheet, err)
	}

	row := 1
	for _, cfg := range process.Configs { // each config has its own stations
		if err := processWriter.SetRow(fmt.Sprintf("A%d", row), stationsRow(styles.highlight, cfg.GetStations())); err != nil {
			return fmt.Errorf(errorWriteProcessSheet, err)
		}
		row++

		if err := writeToolsSection(processWriter, &row, styles.warning, cfg.GetTools()); err != nil {
			return fmt.Errorf(errorWriteProcessSheet, err)
		}
		if err := writeMaterialsSection(processWriter, &row, styles.warning, productType, cfg.GetSteps()); err != nil {
			return fmt.Errorf(errorWriteProcessSheet, err)
		}
		if err = writeControlsSection(processWriter, &row, styles.warning, cfg.GetSteps()); err != nil {
			return fmt.Errorf(errorWriteProcessSheet, err)
		}
		if err := writeMeasurementsSection(processWriter, &row, styles.warning, cfg.GetSteps()); err != nil {
			return fmt.Errorf(errorWriteProcessSheet, err)
		}
	}
	if err := processWriter.Flush(); err != nil {
		return fmt.Errorf(errorWriteProcessSheet, err)
	}
	return nil
}

const errorWriteFlowProcesses = "failed to write flow processes: %w"

func writeFlowProcesses(f *excelize.File, w *excelize.StreamWriter, styles *styles, row *int, productType rs.ProductType, processes []*rs.Process) (err error) {
	if f == nil {
		return fmt.Errorf(errorWriteFlowProcesses, errorNilFile)
	}
	if w == nil {
		return fmt.Errorf(errorWriteFlowProcesses, errorNilStreamWriter)
	}
	if styles == nil {
		return fmt.Errorf(errorWriteFlowProcesses, errorNilStyles)
	}
	if row == nil {
		return fmt.Errorf(errorWriteFlowProcesses, errorNilRow)
	}
	if productType == rs.ProductType_PRODUCT_TYPE_UNSPECIFIED {
		return fmt.Errorf(errorWriteFlowProcesses, errorProductTypeUnspecified)
	}

	if err := w.SetRow(fmt.Sprintf("A%d", *row), cellsWithStyle(styles.warning, columnHeaderFlowProcesses...)); err != nil {
		return fmt.Errorf(errorWriteFlowProcesses, err)
	}
	*row++
	for _, process := range processes {
		if process == nil {
			return fmt.Errorf(errorWriteFlowProcesses, errorNilProcess)
		}

		sheetName, err := processSheetName(process.GetType(), process.GetName())
		if err != nil {
			return fmt.Errorf(errorWriteFlowProcesses, err)
		}

		processSheetLink := excelize.Cell{StyleID: styles.hyperlink, Value: sheetName}
		// TODO: 目前無product name時以process name匯出，待recipe補入product name欄位後修改寫法
		if err := w.SetRow(fmt.Sprintf("A%d", *row), []interface{}{
			process.GetName(),
			process.GetType(),
			processSheetLink,
			sheetName[2:]}); err != nil {
			return fmt.Errorf(errorWriteFlowProcesses, err)
		}
		if err := setDropDownList(f, sheetProcessIndex, fmt.Sprintf("B%d", *row), keysProcessType); err != nil {
			return fmt.Errorf(errorWriteFlowProcesses, err)
		}
		// create link from cell to the target process cell.
		if err := w.File.SetCellHyperLink(sheetProcessIndex, fmt.Sprintf("C%d", *row), "'"+sheetName+"'"+"!A1", "Location"); err != nil {
			return fmt.Errorf(errorWriteFlowProcesses, err)
		}

		if err := writeProcessSheet(f, sheetName, styles, productType, process); err != nil {
			return fmt.Errorf(errorWriteFlowProcesses, err)
		}
		*row++
	}
	return nil
}

const errorWriteProcessIndex = "failed to write process index sheet: %w"

func writeProcessIndexSheet(f *excelize.File, styles *styles, productType rs.ProductType, processes []*rs.Process, optionalFlows []*rs.Recipe_OptionalFlow) (err error) {
	if f == nil {
		return fmt.Errorf(errorWriteProcessIndex, errorNilFile)
	}
	if styles == nil {
		return fmt.Errorf(errorWriteProcessIndex, errorNilStyles)
	}
	if productType == rs.ProductType_PRODUCT_TYPE_UNSPECIFIED {
		return fmt.Errorf(errorWriteProcessIndex, errorProductTypeUnspecified)
	}

	f.NewSheet(sheetProcessIndex)
	w, err := f.NewStreamWriter(sheetProcessIndex)
	if err != nil {
		return fmt.Errorf(errorWriteProcessIndex, err)
	}

	row := 1

	// Standard processes
	if err := w.SetRow(fmt.Sprintf("A%d", row), cellsWithStyle(styles.highlight, "STANDARD PROCESSES")); err != nil {
		return fmt.Errorf(errorWriteProcessIndex, err)
	}
	row++
	if err := writeFlowProcesses(f, w, styles, &row, productType, processes); err != nil {
		return fmt.Errorf(errorWriteProcessIndex, err)
	}
	row++ // section divider

	// Optional flows of processes
	if err := w.SetRow(fmt.Sprintf("A%d", row), cellsWithStyle(styles.highlight, "OPTIONAL PROCESSES")); err != nil {
		return fmt.Errorf(errorWriteProcessIndex, err)
	}
	row++
	for _, flow := range optionalFlows {
		sheetName, err := processSheetName(flow.GetAfterProcessType(), flow.GetAfterProcessName())
		if err != nil {
			return err
		}
		// Header for the optional flow
		if err := w.SetRow(fmt.Sprintf("A%d", row), cellsWithStyle(styles.warning, columnHeaderOptionalFlowInfo...)); err != nil {
			return fmt.Errorf(errorWriteProcessIndex, err)
		}
		row++
		if err := w.SetRow(fmt.Sprintf("A%d", row), []interface{}{
			flow.GetName(),
			sheetName,
			flow.GetMaxRepetitions()}); err != nil {
			return fmt.Errorf(errorWriteProcessIndex, err)
		}
		row++

		// Process list of the optional flow
		if err := writeFlowProcesses(f, w, styles, &row, productType, flow.GetProcesses()); err != nil {
			return fmt.Errorf(errorWriteProcessIndex, err)
		}
		row++
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf(errorWriteProcessIndex, err)
	}
	return nil
}

// FromRecipe export excel file from recipe
func FromRecipe(r *rs.Recipe) (buf *bytes.Buffer, err error) {
	if r == nil {
		return nil, errorNilRecipe
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

	if err = writeRecipeIndexSheet(f, productType, r); err != nil {
		return nil, err
	}

	if err = writeProcessIndexSheet(f, styles, productType, r.GetProcesses(), r.GetOptionalFlows()); err != nil {
		return nil, err
	}

	return f.WriteToBuffer()
}
