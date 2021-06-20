package pgxe

import (
	"github.com/anton7r/pgxe/internal/lexer"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type PreparedQuery struct {
	parts *[]lexer.Part
}

//Prepares a traditional SQL-query with $1, $2, $3 placeholders
//NOTE: This does not use the pgx way to prepare queries under the hood
func Prepare(sql string) (*PreparedQuery, error) {
	parts, err := lexer.Prepare(sql)

	if err != nil {
		return nil, err
	}

	return &PreparedQuery{
		parts,
	}, nil
}

func (p *PreparedQuery) toStr(args ...interface{}) string {
	//TODO:
	return ""
}

func (p *PreparedQuery) Select(db DB, target *interface{}, args ...interface{}) error {
	return db.Select(target, p.toStr(args))
}

func (p *PreparedQuery) Get(db DB, target *interface{}, args ...interface{}) error {
	return db.Get(target, p.toStr(args))
}

func (p *PreparedQuery) Query(db DB, args ...interface{}) (pgx.Rows, error) {
	return db.Query(p.toStr(args))
}

func (p *PreparedQuery) QueryRow(db DB, args ...interface{}) pgx.Row {
	return db.QueryRow(p.toStr(args))
}

func (p *PreparedQuery) Exec(db DB, args ...interface{}) (pgconn.CommandTag, error) {
	return db.Exec(p.toStr(args))
}

type PreparedNamedQuery struct {
	parts *[]lexer.PartNamed
}

//Prepares a named SQL-query with :FieldName placeholders
//NOTE: This does not use the pgx way to prepare queries under the hood
func PrepareNamed(sql string) (*PreparedNamedQuery, error) {
	parts, err := lexer.PrepareNamed(sql)

	if err != nil {
		return nil, err
	}

	return &PreparedNamedQuery{
		parts,
	}, nil
}

func (p *PreparedNamedQuery) toStr(arg *interface{}) string {
	//TODO:
	return ""
}

func (p *PreparedNamedQuery) Select(db DB, target *interface{}, arg *interface{}) error {
	return db.Select(target, p.toStr(arg))
}

func (p *PreparedNamedQuery) Get(db DB, target *interface{}, arg *interface{}) error {
	return db.Get(target, p.toStr(arg))
}

func (p *PreparedNamedQuery) Exec(db DB, arg *interface{}) (pgconn.CommandTag, error) {
	return db.Exec(p.toStr(arg))
}

func (p *PreparedNamedQuery) Query(db DB, arg *interface{}) (pgx.Rows, error) {
	return db.Query(p.toStr(arg))
}

func (p *PreparedNamedQuery) QueryRow(db DB, arg *interface{}) pgx.Row {
	return db.QueryRow(p.toStr(arg))
}
