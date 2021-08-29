package pgxe_test

import (
	"context"
	"log"
	"strconv"
	"testing"

	"github.com/anton7r/pgxe"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxe.DB

func conn() (*pgxe.DB, error) {
	return pgxe.Connect(pgxe.Connection{
		User:     "test",
		Password: "test",
		Database: "test",
		Port:     "5432",
	})
}

func init() {
	var err error
	db, err = conn()
	if err != nil {
		log.Panic("Tests errored in the initialization phase: " + err.Error())
	}
}

func TestUse(t *testing.T) {
	pool, err := pgxpool.Connect(context.TODO(), "user=test password=test dbname=test port=5432 sslmode=disable")
	if err != nil {
		t.Error("Error happened: " + err.Error())
	}

	db := pgxe.Use(pool)

	if db == nil {
		t.Error("unexpected error")
	}
}

func TestConnectClose(t *testing.T) {
	db, err := conn()
	if err != nil {
		t.Error(err.Error())
	}

	db.Close()
}

type Employee struct {
	ID             int
	Name           string
	Surname        string
	PaymentAddress string `db:"paymentAddress"`
}

func TestSelect(t *testing.T) {
	emps := &[]Employee{}

	err := db.Select(emps, "SELECT * FROM employees LIMIT 2")

	if err != nil {
		t.Error("Test errored:" + err.Error())
		t.FailNow()
	}

	if len(*emps) != 2 {
		t.Error("Test failed expected to return 2 rows but got " + strconv.Itoa(len(*emps)))
	}
}

func TestGet(t *testing.T) {
	emp := &Employee{}

	err := db.Get(emp, "SELECT * FROM employees LIMIT 1")

	if err != nil {
		t.Error("Test errored:" + err.Error())
		t.FailNow()
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

func TestQuery(t *testing.T) {
	rows, err := db.Query("SELECT * FROM employees")

	if err != nil {
		t.Error("Could't retrieve rows from db: " + err.Error())
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

func TestQueryRow(t *testing.T) {
	emp := Employee{}

	obj := db.QueryRow("SELECT * FROM employees LIMIT 1")
	err := obj.Scan(&emp.Name, &emp.Surname, &emp.PaymentAddress, &emp.ID)

	if err != nil {
		t.Error("Test errored:" + err.Error())
		t.FailNow()
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

func TestExec(t *testing.T) {
	c, err := db.Exec("INSERT INTO employees (name, surname, \"paymentAddress\") VALUES ($1, $1, $1)", "Hi")

	if err != nil {
		t.Error("Insertion to the database failed: " + err.Error())
		t.FailNow()
	}

	t.Log("got command tag: " + c.String())
}

type tNamed struct {
	Id int
}

func TestNamedSelect(t *testing.T) {
	emps := []Employee{}

	err := db.NamedSelect(&emps, "SELECT * FROM employees WHERE id <= :Id", &tNamed{Id: 3})
	if err != nil {
		t.Error("Test errored: " + err.Error())
		t.FailNow()
	}

	if len(emps) != 3 {
		t.Error("The length of the slice does not match expectations.")
	}
}

func TestNamedGet(t *testing.T) {
	emp := Employee{}

	err := db.NamedGet(&emp, "SELECT * FROM employees WHERE id = :Id LIMIT 1", &tNamed{Id: 1})

	if err != nil {
		t.Error("Test errored:" + err.Error())
		t.FailNow()
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

func TestNamedQuery(t *testing.T) {
	rows, err := db.NamedQuery("SELECT * FROM employees WHERE ID <= :Id", &tNamed{Id: 3})

	if err != nil {
		t.Error("Could't retrieve rows from db: " + err.Error())
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

func TestNamedQueryRow(t *testing.T) {
	emp := Employee{}

	row, err := db.NamedQueryRow("SELECT * FROM employees WHERE id = :Id LIMIT 1", &tNamed{Id: 1})
	if err != nil {
		t.Error("Test errored:" + err.Error())
		t.FailNow()
	}

	err = row.Scan(&emp.Name, &emp.Surname, &emp.PaymentAddress, &emp.ID)
	if err != nil {
		t.Error("Test errored 2:" + err.Error())
		t.FailNow()
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

type tNamedE struct {
	Value string
}

func TestNamedExec(t *testing.T) {
	c, err := db.NamedExec("INSERT INTO employees (name, surname, \"paymentAddress\") VALUES (:Value, :Value, :Value)", &tNamedE{Value: "Hi!"})

	if err != nil {
		t.Error("Insertion to the database failed: " + err.Error())
		t.FailNow()
	}

	t.Log("got command tag: " + c.String())
}
