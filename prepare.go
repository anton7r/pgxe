package pgxe

type PreparedQuery struct {
}

func Prepare(sql string) *PreparedQuery {
	return &PreparedQuery{}
}
