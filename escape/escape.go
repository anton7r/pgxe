package escape

import (
	"encoding/hex"
	"strings"
)

//QuoteString and QuoteBytes are from jackc/pgx/internal/sanitize package

func QuoteString(str string) string {
	return "'" + strings.ReplaceAll(str, "'", "''") + "'"
}

func QuoteBytes(buf []byte) string {
	return `'\x` + hex.EncodeToString(buf) + "'"
}
