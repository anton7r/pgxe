package pgxe

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Connection struct {
	User     string
	Password string
	DbName   string
	DbPort   string
	SslMode  string // defaults to disable

	Logger pgx.Logger
}

//Connect connects to the database
//NOTE: It does not support all the features that are on connection string yet Use is currently recommended to be used instead
func Connect(conn Connection) (*DB, error) {
	if conn.SslMode == "" {
		conn.SslMode = "disable"
	}

	var connStr string = "user=" + conn.User +
		" password=" + conn.Password +
		" dbname=" + conn.DbName +
		" port=" + conn.DbPort +
		" sslmode=" + conn.SslMode

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	config.ConnConfig.Logger = conn.Logger

	db, err2 := pgxpool.ConnectConfig(context.Background(), config)
	if err2 != nil {
		return nil, err2
	}

	return &DB{Pool: db}, nil
}

//Use uses an already existing connection
//It assumes that the fed connection pool is not null
func Use(conn *pgxpool.Pool) *DB {
	return &DB{Pool: conn}
}
