package lexer_test

import (
	"testing"

	"github.com/anton7r/pgxe/internal/lexer"
)

type testingStruct struct {
	Name string
}

func TestNamedComplete(t *testing.T) {
	x, err := lexer.PrepareNamed("SELECT * FROM users WHERE name = :Name")
	if err != nil {
		t.Error("Error while testing:" + err.Error())
	}

	d := &testingStruct{Name: "TEST"}

	sql, err2 := lexer.FinalizeNamed(x, d)
	if err2 != nil {
		t.Error("Error while testing:" + err2.Error())
	}

	exp := "SELECT * FROM users WHERE name = 'TEST'"

	if sql != exp {
		t.Error("Expected: " + exp + "\nBut instead got:" + sql)
	}
}

func TestNamedComplete2(t *testing.T) {
	x, err := lexer.PrepareNamed("SELECT * FROM users WHERE name = :Name ORDER BY ID DESC")
	if err != nil {
		t.Error("Error while testing:" + err.Error())
	}

	d := &testingStruct{Name: "TEST"}

	sql, err2 := lexer.FinalizeNamed(x, d)
	if err2 != nil {
		t.Error("Error while testing:" + err2.Error())
	}

	exp := "SELECT * FROM users WHERE name = 'TEST' ORDER BY ID DESC"

	if sql != exp {
		t.Error("Expected: " + exp + "\nBut instead got:" + sql)
	}
}

func TestNamedIncorrectFieldName(t *testing.T) {
	x, err := lexer.PrepareNamed("SELECT * FROM users WHERE name = :Namd")
	if err != nil {
		t.Error("Error while testing:" + err.Error())
	}

	d := &testingStruct{Name: "TEST"}

	_, err2 := lexer.FinalizeNamed(x, d)
	if err2 == nil {
		t.Error("Did not error!")
	}
}

func TestComplete(t *testing.T) {
	x, err := lexer.Prepare("SELECT * FROM users WHERE name = $1")
	if err != nil {
		t.Error("Error while testing:" + err.Error())
	}

	sql, err2 := lexer.Finalize(x, "TEST")
	if err2 != nil {
		t.Errorf("Error while testing:" + err2.Error() + "\n %+v", x)
	}

	exp := "SELECT * FROM users WHERE name = 'TEST'"

	if sql != exp {
		t.Error("Expected: " + exp + "\nBut instead got:" + sql)
	}
}

func TestComplete2(t *testing.T) {
	x, err := lexer.Prepare("SELECT * FROM users WHERE name = $1 ORDER BY ID DESC")
	if err != nil {
		t.Error("Error while testing:" + err.Error())
	}

	sql, err2 := lexer.Finalize(x, "TEST")
	if err2 != nil {
		t.Errorf("Error while testing:" + err2.Error() + "\n %+v", x)
	}

	exp := "SELECT * FROM users WHERE name = 'TEST' ORDER BY ID DESC"

	if sql != exp {
		t.Error("Expected: " + exp + "\nBut instead got:" + sql)
	}
}

func TestIncorrectIndex(t *testing.T) {
	x, err := lexer.Prepare("SELECT * FROM users WHERE name = $2")
	if err != nil {
		t.Error("Error while testing:" + err.Error())
	}

	_, err2 := lexer.Finalize(x, "TEST")
	if err2 == nil {
		t.Errorf("Did not error!")
	}
}

func TestIncorrectIndex2(t *testing.T) {
	x, err := lexer.Prepare("SELECT * FROM users WHERE name = $g")
	if err != nil {
		t.Error("Error while testing:" + err.Error())
	}

	_, err2 := lexer.Finalize(x, "TEST")
	if err2 == nil {
		t.Errorf("Did not error!")
	}
}
// func TestPrepareNamed(t *testing.T) {

// }
