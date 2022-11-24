package validator

import (
	"fmt"
	"reflect"
)

type types map[string]reflect.Type

var typeList = make(types)

// RegisterType is a function for registering structure, use it in init func.
// RegisterType is a unsafe-thread function.
func RegisterType(id string, rt reflect.Type) error {
	if rt.Kind() != reflect.Struct {
		return fmt.Errorf("passing by reflect.Type is not a struct")
	}

	if _, ok := typeList[id]; ok {
		return fmt.Errorf("id :%v is already exists", id)
	}

	typeList[id] = rt
	return nil
}

type defaultStruct struct {
	Value string `json:"value,omitempty" binding:"required"`
}

// GetConfigType get config from specific struct,
// it will return pointer *struct
func GetConfigType(id string) interface{} {
	rt, ok := typeList[id]
	if !ok {
		rt = reflect.TypeOf(defaultStruct{})
	}
	return reflect.New(rt).Interface()
}
