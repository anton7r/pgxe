package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"github.com/anton7r/pgxe/escape"
)

//TODO: add more support
//gets the field from structs and interfaces
//the interface parameter must be a pointer
func GetNamedField(ifc interface{}, name string) (string, error) {
	name = strings.ToLower(name)
	
	//TODO: move these two lines to an method above the method tree since these are
	rv := reflect.ValueOf(ifc)
	e := rv.Elem()
	kind := e.Kind()

	if kind == reflect.Struct {

		value := e.FieldByNameFunc(func(fieldName string) bool {
			return strings.ToLower(fieldName) == name
		})

		if value.IsValid() {
			return convertValueString(value.Interface())
		}

		return "", errors.New("field '" + name + "' not found")
	}

	return "", errors.New("Only structs are currently supported by the 'ifc' parameter, but the type was " + kind.String())

}

func convertValueString(ifc interface{}) (string, error) {
	switch ifc.(type) {
	case string:
		return escape.QuoteString(ifc.(string)), nil
	case bool:
		if ifc.(bool) {
			return "TRUE", nil
		}
		return "FALSE", nil
	case int:
		return strconv.Itoa(ifc.(int)), nil
	case int8:
		return strconv.Itoa(int(ifc.(int8))), nil
	case int16:
		return strconv.Itoa(int(ifc.(int16))), nil
	case int32:
		return strconv.Itoa(int(ifc.(int32))), nil
	case int64:
		return strconv.FormatInt(ifc.(int64), 10), nil

	case float32:
		return fmt.Sprintf("%f", ifc.(float32)), nil
	case float64:
		return fmt.Sprintf("%f", ifc.(float64)), nil

	case uint:
		return strconv.Itoa(int(ifc.(uint))), nil
	case uint8:
		return strconv.Itoa(int(ifc.(uint8))), nil
	case uint16:
		return strconv.Itoa(int(ifc.(uint16))), nil
	case uint32:
		return strconv.Itoa(int(ifc.(uint32))), nil
	case uint64:
		return strconv.FormatUint(ifc.(uint64), 10), nil
	}

	return "", errors.New("unsupported type")
}
