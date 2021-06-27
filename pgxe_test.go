package pgxe_test

import (
	"context"
	"log"
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

//TODO
func TestSelect(t *testing.T) {

}

//TODO
func TestGet(t *testing.T) {

}

//TODO
func TestQuery(t *testing.T) {

}

//TODO
func TestQueryRow(t *testing.T) {

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
