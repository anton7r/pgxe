package lexer

import (
	"unicode/utf8"
)

type Part struct {
	Str     string
	OnParam bool
}

type PartNamed struct {
	Str     string
	OnParam bool
}

//PrepareNamed is used for parsing and preparing sql queries
func PrepareNamed(sql string) (*[]PartNamed, error) {
	sql = Format(sql)

	sl := []PartNamed{} //slice

	start := 0
	pos := 0
	onParameter := false

	for {
		r, width := utf8.DecodeRuneInString(sql[pos:])

		if r == utf8.RuneError {
			sl = append(sl, PartNamed{sql[start:pos], onParameter})

			break
		} else if r == delim {
			sl = append(sl, PartNamed{sql[start:pos], onParameter})

			onParameter = true
			start = pos + width
		} else if r > 'z' || r < '1' {
			if onParameter {
				sl = append(sl, PartNamed{sql[start:pos], onParameter})

				onParameter = false
				start = pos
			}
		}

		pos += width
	}

	return &sl, nil
}

//Prepare is used for parsing and preparing sql queries
func Prepare(sql string) *[]Part {
	sql = Format(sql)

	sl := []Part{} // slice

	start := 0
	pos := 0
	onParameter := false

	for {
		r, width := utf8.DecodeRuneInString(sql[pos:])

		if r == utf8.RuneError {
			sl = append(sl, Part{sql[start:pos], onParameter})

			break
		} else if r == '$' {
			sl = append(sl, Part{sql[start:pos], onParameter})

			onParameter = true
			start = pos + width
		} else if r > 'z' || r < '1' {
			if onParameter {
				sl = append(sl, Part{sql[start:pos], onParameter})

				onParameter = false
				start = pos
			}
		}

		pos += width
	}

	return &sl
}
