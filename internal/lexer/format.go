package lexer

import (
	"bytes"
	"unicode/utf8"
)

//Format is just an internal algo to format sql queries
//It basically removes unneccessary whitespace from queries
//Automatically used by prepare
//CURRENTLY DOES NOT REMOVE COMMENTS
func Format(q string) string {
	buf := &bytes.Buffer{}

	start := 0
	pos := 0
	isCharacter := false

	lastRune := 'a'
	previousRune := ' '

	for {
		r, width := utf8.DecodeRuneInString(q[pos:])
		nextPos := pos + width

		if r == utf8.RuneError {
			if isCharacter {
				buf.WriteString(" " + q[start:pos])
			}

			break
		} else if r == ' ' || r == '\n' || r == '\t' {
			if isCharacter {
				cStr := q[start:pos]
				if lastRune == 'a' || previousRune == '(' || previousRune == ')' || lastRune == '(' || lastRune == ')' || lastRune == ',' {
					buf.WriteString(cStr)
				} else {
					buf.WriteString(" " + cStr)
				}
				lastRune = previousRune
			}

			start = nextPos
			isCharacter = false
		} else {
			isCharacter = true

			previousRune = r
		}

		pos = nextPos
	}

	return buf.String()
}
