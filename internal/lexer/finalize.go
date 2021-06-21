package lexer

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/anton7r/pgxe/internal/utils"
)

//FinalizeNamed is a function that finalizes the named sql queries before they are executed
func FinalizeNamed(parts *[]PartNamed, arg interface{}) (string, error) {
	buf := &bytes.Buffer{}
	prepped := utils.PrepReflect(arg)

	for _, part := range *parts {
		if part.OnParam {
			r, err := utils.GetNamedField(prepped, part.Str)
			if err != nil {
				return "", err
			}

			_, err2 := buf.WriteString(r)
			if err2 != nil {
				return "", err2
			}

		} else {
			_, err := buf.WriteString(part.Str)
			if err != nil {
				return "", err
			}
		}
	}

	return buf.String(), nil
}

//Finalize is a function that finalizes the traditional sql queries before they are executed
func Finalize(parts *[]Part, args ...interface{}) (string, error) {
	buf := &bytes.Buffer{}
	l := len(args)

	for _, part := range *parts {
		if part.OnParam {
			i, err := strconv.ParseInt(part.Str, 10, 0)
			if err != nil {
				return "", err
			}

			if int(i) > l {
				return "", errors.New("index too large")
			}

			str, err2 := utils.ConvertValueString(args[i-1])
			if err2 != nil {
				return "", err2
			}

			buf.WriteString(str)

		} else {
			_, err := buf.WriteString(part.Str)
			if err != nil {
				return "", err
			}
		}
	}

	return buf.String(), nil
}
