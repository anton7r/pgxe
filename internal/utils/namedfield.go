package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/anton7r/pgxe/internal/escape"
)

func PrepReflect(ifc interface{}) reflect.Value {
	return reflect.ValueOf(ifc).Elem()
}

//the map has lowercase fieldname as
func GetFields(e reflect.Value) *map[string]string {
	a := map[string]string{}

	for i := 0; i < e.NumField(); i++ {
		f := e.Type().Field(i).Name
		fl := strings.ToLower(f)
		a[fl] = f
	}

	//fmt.Println(a)

	return &a
}
//Deprecated:
//TODO: add more support for different types
//gets the field from structs and interfaces
//the interface parameter must be a pointer
//
//This algorithm is quite slow as it's perfomance is determined by how many fields the struct has (calls * fields)
//This one should be used only when there is only one value that needs to be retireved from the struct
//The function below this function fixes this performance issue
func GetNamedField(e reflect.Value, name string) (string, error) {
	kind := e.Kind()

	//switch is on avg a couple of ns faster than if else here but the performance difference is not really that large so that would make a meaningful impact on the total performance
	//but as we optimize it more the difference starts to make somewhat of a meaningful impact
	switch kind {
	case reflect.Struct:
		value := e.FieldByNameFunc(func(fieldName string) bool {
			return strings.EqualFold(fieldName, name)
		})

		if value.IsValid() {
			return convertValueString(value.Interface())
		}

		return "", errors.New("field '" + name + "' not found")

	default:
		return "", errors.New("only structs are currently supported by the 'ifc' parameter, but the type was " + kind.String())
	}
}

func GetNamedFieldV2(e reflect.Value, fields *map[string]string, name string) (string, error) {
	kind := e.Kind()

	name = strings.ToLower(name)

	//switch is on avg a couple of ns faster than if else here but the performance difference is not really that large so that would make a meaningful impact on the total performance
	//but as we optimize it more the difference starts to make somewhat of a meaningful impact
	switch kind {
	case reflect.Struct: 
		fName := (*fields)[name]

		if fName == "" {
			return "", errors.New("field '" + name + "' not found")
		}

		value := e.FieldByName(fName)

		if value.IsValid() {
			return convertValueString(value.Interface())
		}

		return "", errors.New("the provided value was not valid")

	default:
		return "", errors.New("only structs are currently supported by the 'ifc' parameter, but the type was " + kind.String())
	}
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
