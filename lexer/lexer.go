//The job of this package is to parse the queries and replace placeholder values with the real values
package lexer

import (
	"bytes"
	"github.com/anton7r/pgxe/utils"
	"unicode/utf8"
)

func appendPart(b *bytes.Buffer, str string, isParameter bool, arg interface{}) error {
	if !isParameter {
		_, err := b.WriteString(str)
		if err != nil {
			return err
		}
	} else {
		str, err := utils.GetNamedField(arg, str)
		if err != nil {
			return err
		}
		_, err2 := b.WriteString(str)
		if err2 != nil {
			return err2
		}
	}
	return nil
}

//Compile is used for parsing and compiling sql queries
func Compile(sql string, arg interface{}) (string, error) {

	buf := &bytes.Buffer{}

	start := 0
	pos := 0
	onParameter := false

	for {
		r, width := utf8.DecodeRuneInString(sql[pos:])

		if r == utf8.RuneError {
			err := appendPart(buf, sql[start:pos], onParameter, arg)
			if err != nil {
				return "", err
			}

			break
		} else if r == ':' {
			err := appendPart(buf, sql[start:pos], onParameter, arg)
			if err != nil {
				return "", err
			}

			onParameter = true
			start = pos + width
		} else if r >= 'z' || r <= '1' {
			if onParameter {
				err := appendPart(buf, sql[start:pos], onParameter, arg)
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
