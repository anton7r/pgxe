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
	Pool  *pgxpool.Pool
	delim string
}

type Connection struct {
	User     string
	Password string
	DbName   string
	DbPort   string
}

type ExtraConfig struct {
	Delim string // this defaults to
}

type Rows struct {
}

type Row struct {
}

func getDelimOrDefault(e ExtraConfig) string {
	if e.Delim != "" {
		return e.Delim
	} else {
		return ":"
	}
}

//Connect connects to the database
//NOTE: It does not support all the features that are on connection string yet Use is currently recommended to be used instead
func Connect(conn Connection, eConf ...ExtraConfig) (*DB, error) {
	var connStr string = "user=" + conn.User +
		" password=" + conn.Password +
		" dbname=" + conn.DbName +
		" port=" + conn.DbPort +
		" sslmode=disable"

	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	ec := eConf[0]

	return &DB{delim: getDelimOrDefault(ec), Pool: db}, nil
}

//Use uses an already existing connection
//It assumes that the fed connection pool is not null
func Use(conn *pgxpool.Pool, eConf ...ExtraConfig) *DB {
	ec := eConf[0]

	return &DB{delim: getDelimOrDefault(ec), Pool: conn}
}

//Select is a high-level function that is used to retrieve data from database into structs
func (DB *DB) Select(target *interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(context.Background(), DB.Pool, target, query, args...)
}

//Select is a high-level function that is used to retrieve data from database into structs
func (DB *DB) Get(target *interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(context.Background(), DB.Pool, target, query, args...)
}

//Query simplifies the parameters
func (DB *DB) Query(query string, args ...interface{}) (pgx.Rows, error) {
	return DB.Pool.Query(context.Background(), query, args...)
}

//QueryRow simplifies the parameters
func (DB *DB) QueryRow(query string, args ...interface{}) pgx.Row {
	return DB.Pool.QueryRow(context.Background(), query, args...)
}

//Exec simplifies the parameters
func (DB *DB) Exec(query string, args ...interface{}) (pgconn.CommandTag, error) {
	return DB.Pool.Exec(context.Background(), query, args...)
}

//NamedSelect is a high-level function that is used to retrieve data from database into structs
func (DB *DB) NamedSelect(target *interface{}, query string, arg *interface{}) error {
	query, err := lexer.Compile(query, arg)
	if err != nil {
		return err
	}

	return pgxscan.Select(context.Background(), DB.Pool, target, query)
}

//NamedGet is a high-level function that is used to retrieve data from database into structs
func (DB *DB) NamedGet(target *interface{}, query string, arg *interface{}) error {
	query, err := lexer.Compile(query, arg)
	if err != nil {
		return err
	}

	return pgxscan.Get(context.Background(), DB.Pool, target, query)
}

//NamedQuery simplifies the parameters and allows the use of named parameters
func (DB *DB) NamedQuery(query string, arg *interface{}) (pgx.Rows, error) {
	query, err := lexer.Compile(query, arg)
	if err != nil {
		return nil, err
	}

	return DB.Pool.Query(context.Background(), query)
}

//NamedQueryRow simplifies the parameters and allows the use of named parameters
func (DB *DB) NamedQueryRow(query string, arg *interface{}) (pgx.Row, error) {
	query, err := lexer.Compile(query, arg)
	if err != nil {
		return nil, err
	}

	return DB.Pool.QueryRow(context.Background(), query), nil
}

//NamedExec simplifies the parameters and allows the use of named parameters
func (DB *DB) NamedExec(query string, arg *interface{}) (pgconn.CommandTag, error) {
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
