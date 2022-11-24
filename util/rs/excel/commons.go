package excel

import "errors"

const (
	sheetDefault = "Sheet1"

	// for recipe
	sheetRecipeIndex  = "Recipe"
	sheetProcessIndex = "Process"

	// for spec
	sheetSpecIndex  = "Spec"
	sheetSpecConfig = "Config"
)

var (
	columnHeaderTools           = []string{"TOOL", "ROLE", "MUST_CHECK"}
	columnHeaderRecipeMaterials = []string{"MATERIAL", "NAME", "LEVEL", "MAIN TYPE", "REQ_RECIPE_ID", "CONTAINER_TYPE", "LOW", "MID", "HIGH", "UNIT"}
	columnHeaderMeasurements    = []string{"MEASUREMENT", "NAME", "LOW", "MID", "HIGH", "UNIT"}

	columnHeaderFlowProcesses    = []string{"NAME", "TYPE", "SHEET", "PRODUCT NAME"}
	columnHeaderOptionalFlowInfo = []string{"FLOW NAME", "AFTER PROCESS", "MAX REPETITIONS"}

	// for spec
	columnHeaderSpecMaterials = []string{"MATERIAL", "LEVEL", "MAIN TYPE", "REQ_RECIPE_ID", "CONTAINER_TYPE", "LOW", "MID", "HIGH", "UNIT"}
	columnHeaderProperties    = []string{"PROPERTY", "LOW", "MID", "HIGH", "ENUM", "UNIT"}
)

var (
	errorNilFile                = errors.New("nil file")
	errorNilRow                 = errors.New("nil row")
	errorNilStreamWriter        = errors.New("nil stream writer")
	errorNilStyles              = errors.New("nil styles")
	errorNilRecipe              = errors.New("nil recipe")
	errorNilSpec                = errors.New("nil spec")
	errorNilProcess             = errors.New("nil process")
	errorEmptySteps             = errors.New("empty steps")
	errorEmptySheetName         = errors.New("empty sheet name")
	errorProductTypeUnspecified = errors.New("product type unspecified")
)
