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
func Prepare(sql string) *PreparedQuery {
	return &PreparedQuery{
		lexer.Prepare(sql),
	}
}

func (p *PreparedQuery) str(args ...interface{}) (string, error) {
	return lexer.Finalize(p.parts, args...)
}

func (p *PreparedQuery) Select(db *DB, target interface{}, args ...interface{}) error {
	str, err := p.str(args...)
	if err != nil {
		return err
	}

	return db.Select(target, str)
}

func (p *PreparedQuery) Get(db *DB, target interface{}, args ...interface{}) error {
	str, err := p.str(args...)
	if err != nil {
		return err
	}

	return db.Get(target, str)
}

func (p *PreparedQuery) Query(db *DB, args ...interface{}) (pgx.Rows, error) {
	str, err := p.str(args...)
	if err != nil {
		return nil, err
	}

	return db.Query(str)
}

func (p *PreparedQuery) QueryRow(db *DB, args ...interface{}) (pgx.Row, error) {
	str, err := p.str(args...)
	if err != nil {
		return nil, err
	}

	return db.QueryRow(str), nil
}

func (p *PreparedQuery) Exec(db *DB, args ...interface{}) (pgconn.CommandTag, error) {
	str, err := p.str(args...)
	if err != nil {
		return nil, err
	}

	return db.Exec(str)
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

func (p *PreparedNamedQuery) str(arg interface{}) (string, error) {
	return lexer.FinalizeNamed(p.parts, arg)
}

func (p *PreparedNamedQuery) Select(db *DB, target interface{}, arg interface{}) error {
	str, err := p.str(arg)
	if err != nil {
		return err
	}

	return db.Select(target, str)
}

func (p *PreparedNamedQuery) Get(db *DB, target interface{}, arg interface{}) error {
	str, err := p.str(arg)
	if err != nil {
		return err
	}

	return db.Get(target, str)
}

func (p *PreparedNamedQuery) Exec(db *DB, arg interface{}) (pgconn.CommandTag, error) {
	str, err := p.str(arg)
	if err != nil {
		return nil, err
	}

	return db.Exec(str)
}

func (p *PreparedNamedQuery) Query(db *DB, arg interface{}) (pgx.Rows, error) {
	str, err := p.str(arg)
	if err != nil {
		return nil, err
	}

	return db.Query(str)
}

func (p *PreparedNamedQuery) QueryRow(db *DB, arg interface{}) (pgx.Row, error) {
	str, err := p.str(arg)
	if err != nil {
		return nil, err
	}

	return db.QueryRow(str), nil
}
