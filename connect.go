package pgxe

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Connection struct {
	User     string // pgx defaults this to your OS username
	Password string
	Database string // the name of the database
	Host     string // defaults to localhost
	Port     string // defaults to 5432

	SslMode string // defaults to disable

	PoolMaxConns          int // default 4
	PoolMinConns          int // default 0
	PoolMaxConnLifetime   int // default time.Hour
	PoolMaxConnIdleTime   int // default time.Minute * 30
	PoolHealthCheckPeriod int // default time.Minute

	Logger pgx.Logger
}

//Connect connects to the database
//NOTE: It does not support all the features that are on connection string yet Use is currently recommended to be used instead
func Connect(conn Connection) (*DB, error) {
	if conn.SslMode == "" {
		conn.SslMode = "disable"
	}

	if conn.Port == "" {
		conn.Port = "5432"
	}

	if conn.Host == "" {
		conn.Host = "localhost"
	}

	var connStr string = "user=" + conn.User +
		" password=" + conn.Password +
		" dbname=" + conn.Database +
		" host=" + conn.Host +
		" port=" + conn.Port +
		" sslmode=" + conn.SslMode

	if conn.PoolMaxConns != 0 {
		connStr += " pool_max_conns=" + strconv.Itoa(conn.PoolMaxConns)
	}

	if conn.PoolMinConns != 0 {
		connStr += " pool_min_conns=" + strconv.Itoa(conn.PoolMinConns)
	}

	if conn.PoolMaxConnLifetime != 0 {
		connStr += " pool_max_conn_lifetime=" + strconv.Itoa(conn.PoolMaxConnLifetime)
	}

	if conn.PoolMaxConnIdleTime != 0 {
		connStr += " pool_max_conn_idle_time=" + strconv.Itoa(conn.PoolMaxConnIdleTime)
	}

	if conn.PoolHealthCheckPeriod != 0 {
		connStr += " pool_health_check_period=" + strconv.Itoa(conn.PoolHealthCheckPeriod)
	}

	config, err := pgxpool.ParseConfig(connStr) // we need to parse it because otherwise it could not use pgx default values and could be more prone to bugs
	if err != nil {
		return nil, err
	}

	if conn.Logger != nil {
		config.ConnConfig.Logger = conn.Logger
	}

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
