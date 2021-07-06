package pgxe

import (
	"github.com/anton7r/pgx-scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

type Rows struct {
	rows pgx.Rows
}

func (r *Rows) Scan(dst interface{}) {
	pgxscan.ScanAll(dst, r.rows)
}

func (r *Rows) ScanOne(dst interface{}) {
	pgxscan.ScanOne(dst, r.rows)
}

func (r *Rows) X() {
	
}