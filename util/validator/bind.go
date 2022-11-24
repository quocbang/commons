package validator

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// New create new validator and set tag name
func New(tag string) {
	validate = validator.New()
	validate.SetTagName(tag)
}

// BindJSON decode io.Reader to point struct and validate with tag name
func BindJSON(r io.Reader, obj interface{}) error {
	if r == nil {
		return fmt.Errorf("invalid request")
	}
	return decodeJSON(r, obj)
}

func decodeJSON(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return validateStruct(obj)
}

func validateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		if err := validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}
