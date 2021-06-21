package lexer_test

import (
	"testing"

	"github.com/anton7r/pgxe/internal/lexer"
)

type TestStruct struct {
	Name  string
	Email string
}

func TestCompile(t *testing.T) {
	str, err := lexer.Compile("SELECT * FROM users WHERE name = :Name AND email = :Email", &TestStruct{Name: "Foo", Email: "foo@barbar"})
	if err != nil {
		t.Error("Failed: ", err.Error())
	}

	expected := "SELECT * FROM users WHERE name = 'Foo' AND email = 'foo@barbar'"

	if str != expected {
		t.Error("Expected \"" + expected + "\" but was \"" + str + "\"")
	}
}

func TestCompileWrong(t *testing.T) {
	_, err := lexer.Compile("SELECT * FROM users WHERE name = :Namd AND email = :EmaiI", &TestStruct{Name: "Foo", Email: "foo@barbar"})
	if err == nil {
		t.Error("Did not error")
	}
}

func BenchmarkCompile(b *testing.B) {
	ts := &TestStruct{Name: "Foo", Email: "foo@barbar"}
	for i := 0; i < b.N; i++ {
		lexer.Compile("SELECT * FROM users WHERE name = :name AND email = :email", ts)
	}
}

func BenchmarkCompile_2(b *testing.B) {
	ts := &TestStruct{Name: "Foo", Email: "foo@barbar"}
	for i := 0; i < b.N; i++ {
		lexer.Compile("SELECT * FROM users WHERE name = :name AND email = :email SORT BY name DESC", ts)
	}
}
