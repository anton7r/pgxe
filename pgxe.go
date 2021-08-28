package pgxe

import (
	"context"

	"github.com/anton7r/pgxe/internal/lexer"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"

	"github.com/anton7r/pgx-scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

//No one likes typing context.Background() over and over again when it is not neccessary to be used

//DB stores the internal PgxPool
type DB struct {
	Pool *pgxpool.Pool
}

//Select is a high-level function that is used to retrieve data from database into slices which has structs.
func (DB *DB) Select(target interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(context.Background(), DB.Pool, target, query, args...)
}

//Get is a high-level function that is used to retrieve data from database into a single struct
func (DB *DB) Get(target interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(context.Background(), DB.Pool, target, query, args...)
}

//Query is used to retrieve multiple rows from the database
func (DB *DB) Query(query string, args ...interface{}) (pgx.Rows, error) {
	return DB.Pool.Query(context.Background(), query, args...)
}

//QueryRow is used to retrieve a single row from the database
func (DB *DB) QueryRow(query string, args ...interface{}) pgx.Row {
	return DB.Pool.QueryRow(context.Background(), query, args...)
}

//Exec is used to run actions on database that wont return any values to the client
func (DB *DB) Exec(query string, args ...interface{}) (pgconn.CommandTag, error) {
	return DB.Pool.Exec(context.Background(), query, args...)
}

//NamedSelect is a high-level function that is used to retrieve data from database into structs
func (DB *DB) NamedSelect(target interface{}, query string, arg interface{}) error {
	query, err := lexer.Compile(query, arg)
	if err != nil {
		return err
	}

	return pgxscan.Select(context.Background(), DB.Pool, target, query)
}

//NamedGet is a high-level function that is used to retrieve data from database into structs
func (DB *DB) NamedGet(target interface{}, query string, arg interface{}) error {
	query, err := lexer.Compile(query, arg)
	if err != nil {
		return err
	}

	return pgxscan.Get(context.Background(), DB.Pool, target, query)
}

//NamedQuery simplifies the parameters and allows the use of named parameters
func (DB *DB) NamedQuery(query string, arg interface{}) (pgx.Rows, error) {
	query, err := lexer.Compile(query, arg)
	if err != nil {
		return nil, err
	}

	return DB.Pool.Query(context.Background(), query)
}

//NamedQueryRow simplifies the parameters and allows the use of named parameters
func (DB *DB) NamedQueryRow(query string, arg interface{}) (pgx.Row, error) {
	query, err := lexer.Compile(query, arg)
	if err != nil {
		return nil, err
	}

	return DB.Pool.QueryRow(context.Background(), query), nil
}

//NamedExec simplifies the parameters and allows the use of named parameters
func (DB *DB) NamedExec(query string, arg interface{}) (pgconn.CommandTag, error) {
	query, err := lexer.Compile(query, arg)
	if err != nil {
		return nil, err
	}

	return DB.Pool.Exec(context.Background(), query)
}

//Close closes the underlying pool
func (DB *DB) Close() {
	DB.Pool.Close()
}
