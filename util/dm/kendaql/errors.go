package kendaql

import (
	"errors"
	"fmt"
)

// error
var (
	ErrNoParameter        = errors.New("kendaql: got the zero number of input parameter")
	ErrOddParameter       = errors.New("kendaql: got the odd number of input parameter")
	ErrSliceParameterOnly = errors.New("kendaql: only supported slice type of input parameter")
)

type callerError struct {
	caller string
	err    error
}

func (e *callerError) Error() string {
	return fmt.Sprintf("%v for '%s' function", e.err, e.caller)
}
