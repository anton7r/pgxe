package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/anton7r/pgxe/internal/escape"
)

func PrepReflect(ifc interface{}) reflect.Value {
	return reflect.ValueOf(ifc).Elem()
}

//TODO: add more support for different types
//gets the field from structs and interfaces
//the interface parameter must be a pointer
func GetNamedField(e reflect.Value, name string) (string, error) {
	kind := e.Kind()

	//switch is on avg a couple of ns faster than if else here but the performance difference is not really that large so that would make a meaningful impact on the total performance
	//but as we optimize it more the difference starts to make somewhat of a meaningful impact
	switch kind {
	case reflect.Struct:
		//value := e.FieldByNameFunc(func(fieldName string) bool {
		//	return strings.EqualFold(fieldName, name)
		//})
		value := e.FieldByName(name)

		if value.IsValid() {
			return ConvertValueString(value.Interface())
		}

		return "", errors.New("field '" + name + "' not found")

	default:
		return "", errors.New("only structs are currently supported by the 'ifc' parameter, but the type was " + kind.String())
	}
}

func ConvertValueString(ifc interface{}) (string, error) {
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
