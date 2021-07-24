package pgxe_test

import (
	"testing"

	"github.com/anton7r/pgxe"
)

func TestPrepareSelect(t *testing.T) {
	prep := pgxe.Prepare("SELECT * FROM employees WHERE ID <= $1")

	emps := []Employee{}

	err := prep.Select(db, &emps, 3)
	if err != nil {
		t.Fatal("Errored in preparation: " + err.Error())
	}

	if len(emps) != 3 {
		t.Fatal("Expected length 3 but got " + string(rune(len(emps))))
	}
}

func TestPrepareGet(t *testing.T) {
	prep := pgxe.Prepare("SELECT * FROM employees WHERE ID = $1 LIMIT 1")

	emp := Employee{}

	err := prep.Get(db, &emp, 1)
	if err != nil {
		t.Fatal("Errored during retrieval: " + err.Error())
	}

	if emp.Name != "ABC" {
		t.Error("Name wasnt as expected and got: " + emp.Name)
	}

	if emp.Surname != "123" {
		t.Error("Surname wasnt as expected and got: " + emp.Surname)
	}

	if emp.PaymentAddress != "FooBar" {
		t.Error("PaymentAddress wasnt as expected and got: " + emp.PaymentAddress)
	}

	if emp.ID != 1 {
		t.Errorf("ID wasnt as expected and got: %o", emp.ID)
	}

}

func TestPrepareQueryRow(t *testing.T) {
	prep := pgxe.Prepare("SELECT * FROM employees WHERE ID = $1 LIMIT 1")

	emp := Employee{}

	row, err := prep.QueryRow(db, 1)
	if err != nil {
		t.Fatal("Errored during retrieval: " + err.Error())
	}

	err = row.Scan(&emp.Name, &emp.Surname, &emp.PaymentAddress, &emp.ID)
	if err != nil {
		t.Fatal("Errored during scan: " + err.Error())
	}

	if emp.Name != "ABC" {
		t.Error("Name wasnt as expected and got: " + emp.Name)
	}

	if emp.Surname != "123" {
		t.Error("Surname wasnt as expected and got: " + emp.Surname)
	}

	if emp.PaymentAddress != "FooBar" {
		t.Error("PaymentAddress wasnt as expected and got: " + emp.PaymentAddress)
	}

	if emp.ID != 1 {
		t.Errorf("ID wasnt as expected and got: %o", emp.ID)
	}

}

func TestPrepareNamedQuery(t *testing.T) {
	prep, err := pgxe.PrepareNamed("SELECT * FROM employees WHERE ID <= :Id")

	if err != nil {
		t.Fatal("Errored in preparation: " + err.Error())
	}

	rows, er2 := prep.Query(db, &tNamed{Id: 3})
	if er2 != nil {
		t.Error("Test errored: " + er2.Error())
	}

	for rows.Next() {
		emp := Employee{}
		rows.Scan(&emp.Name, &emp.Surname, &emp.PaymentAddress, &emp.ID)

		if err != nil {
			t.Error("Test errored:" + err.Error())
			t.FailNow()
		}

		if emp.Name == "" {
			t.Error("Name is empty")
		}

		if emp.Surname == "" {
			t.Error("Surname is empty")
		}

		if emp.PaymentAddress == "" {
			t.Error("PaymentAddress is empty")
		}

		if emp.ID == 0 {
			t.Errorf("ID was 0")
		}
	}
}
