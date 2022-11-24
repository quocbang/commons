package excel

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const (
	styleHighlight = `{"fill":{"type":"pattern","color":["#ffd700"],"pattern":1}}`
	styleWarning   = `{"font":{"color":"#ff0000"}}`
	styleHyperlink = `{"font":{"color":"#0000ff","underline":"single"}}`
)

type styles struct {
	highlight int
	hyperlink int
	warning   int
}

func setDefaultStyles(f *excelize.File) (*styles, error) {
	if f == nil {
		return nil, errorNilFile
	}
	highlight, err := f.NewStyle(styleHighlight)
	if err != nil {
		return nil, fmt.Errorf("failed to set highlight style: %w", err)
	}
	hyperlink, err := f.NewStyle(styleHyperlink)
	if err != nil {
		return nil, fmt.Errorf("failed to set hyperlink style: %w", err)
	}
	warning, err := f.NewStyle(styleWarning)
	if err != nil {
		return nil, fmt.Errorf("failed to set warning style: %w", err)
	}
	return &styles{highlight: highlight, hyperlink: hyperlink, warning: warning}, nil
}

func cellsWithStyle(style int, values ...string) []interface{} {
	cells := make([]interface{}, len(values))
	for i, value := range values {
		cells[i] = excelize.Cell{StyleID: style, Value: value}
	}
	return cells
}
