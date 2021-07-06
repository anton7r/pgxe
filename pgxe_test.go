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

func init() {
	var err error
	db, err = pgxe.Connect(pgxe.Connection{
		User:     "test",
		Password: "test",
		DbName:   "test",
		DbPort:   "5432",
	})

	if err != nil {
		log.Panic("Tests errored in the initialization phase: " + err.Error())
	}
}

func TestUse(t *testing.T) {
	pool, err := pgxpool.Connect(context.TODO(), "user="+"test"+" password="+"test"+" dbname="+"test"+" port="+"5432"+" sslmode=disable")
	if err != nil {
		t.Error("Error happened: " + err.Error())
	}

	db := pgxe.Use(pool)

	if db == nil {
		t.Error("unexpected error")
	}
}

func TestConnect(t *testing.T) {
	db, err := pgxe.Connect(pgxe.Connection{
		User:     "test",
		Password: "test",
		DbName:   "test",
		DbPort:   "5432",
	})

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

//TODO
func TestQuery(t *testing.T) {

}

func TestQueryRow(t *testing.T) {
	emp := &Employee{}

	err := db.QueryRow("SELECT * FROM employees LIMIT 1").
		Scan(emp)

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

//TODO
func TestExec(t *testing.T) {

}

//TODO
func TestNamedSelect(t *testing.T) {

}

//TODO
func TestNamedGet(t *testing.T) {

}

//TODO
func TestNamedQuery(t *testing.T) {

}

//TODO
func TestNamedQueryRow(t *testing.T) {

}

//TODO
func TestNamedExec(t *testing.T) {

}

//TODO
func TestClose(t *testing.T) {
	db.Close()
}
