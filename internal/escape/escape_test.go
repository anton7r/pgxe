package escape_test

import (
	"testing"

	"github.com/anton7r/pgxe/internal/escape"
)

func TestEscape(t *testing.T) {
	r := escape.QuoteString("Dan's")
	if r != "'Dan''s'" {
		t.Error("Escape failed!")
	}
}

func TestEscapeBytes(t *testing.T) {
	r := escape.QuoteBytes([]byte("Dan's"))
	if r != "'\\x44616e2773'" {
		t.Error("Escape failed! Got: " + r)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		escape.QuoteString("Dan's")
	}
}