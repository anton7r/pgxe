package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/anton7r/pgxe/escape"
)

func PrepReflect(ifc interface{}) reflect.Value {
	//TODO: move these two lines to an method above the method tree since these are
	rv := reflect.ValueOf(ifc)
	return rv.Elem()
}

//TODO: add more support
//gets the field from structs and interfaces
//the interface parameter must be a pointer
func GetNamedField(e reflect.Value, name string) (string, error) {
	kind := e.Kind()

	if kind == reflect.Struct {

		value := e.FieldByNameFunc(func(fieldName string) bool {
			return strings.EqualFold(fieldName, name)
		})

		if value.IsValid() {
			return convertValueString(value.Interface())
		}

		return "", errors.New("field '" + name + "' not found")
	}

	return "", errors.New("Only structs are currently supported by the 'ifc' parameter, but the type was " + kind.String())

}

func convertValueString(ifc interface{}) (string, error) {
	switch t := ifc.(type) {
	case string:
		return escape.QuoteString(ifc.(string)), nil
	case bool:
		if t {
			return "TRUE", nil
		}
		return "FALSE", nil
	case int:
		return strconv.Itoa(t), nil
	case int8:
		return strconv.Itoa(int(t)), nil
	case int16:
		return strconv.Itoa(int(t)), nil
	case int32:
		return strconv.Itoa(int(t)), nil
	case int64:
		return strconv.FormatInt(t, 10), nil

	case float32:
		return fmt.Sprintf("%f", t), nil
	case float64:
		return fmt.Sprintf("%f", t), nil

	case uint:
		return strconv.Itoa(int(t)), nil
	case uint8:
		return strconv.Itoa(int(t)), nil
	case uint16:
		return strconv.Itoa(int(t)), nil
	case uint32:
		return strconv.Itoa(int(t)), nil
	case uint64:
		return strconv.FormatUint(t, 10), nil
	}

	return "", errors.New("unsupported type")
}
