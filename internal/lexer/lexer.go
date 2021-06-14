//The job of this package is to parse the queries and replace placeholder values with the real values
package lexer

import (
	"bytes"
	"reflect"
	"unicode/utf8"

	"github.com/anton7r/pgxe/internal/utils"
)

func appendPart(b *bytes.Buffer, str string, isParameter bool, arg reflect.Value) error {
	if isParameter {
		var err error
		str, err = utils.GetNamedField(arg, str)
		if err != nil {
			return err
		}
	}

	_, err2 := b.WriteString(str)
	if err2 != nil {
		return err2
	}

	return nil
}

const delim = ':'

//Compile is used for parsing and compiling sql queries
func Compile(sql string, arg interface{}) (string, error) {

	buf := &bytes.Buffer{}

	start := 0
	pos := 0
	onParameter := false

	prepped := utils.PrepReflect(arg)

	for {
		r, width := utf8.DecodeRuneInString(sql[pos:])

		if r == utf8.RuneError {
			err := appendPart(buf, sql[start:pos], onParameter, prepped)
			if err != nil {
				return "", err
			}

			break
		} else if r == delim {
			err := appendPart(buf, sql[start:pos], onParameter, prepped)
			if err != nil {
				return "", err
			}

			onParameter = true
			start = pos + width
		} else if r >= 'z' || r <= '1' {
			if onParameter {
				err := appendPart(buf, sql[start:pos], onParameter, prepped)
				if err != nil {
					return "", err
				}

				onParameter = false
				start = pos
			}
		}

		pos += width
	}

	return buf.String(), nil
}
