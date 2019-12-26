package repository

type SQLHandler interface {
	Execute(query string, args ...interface{}) (Result, error)
	QueryRow(query string, args ...interface{}) Row
	Query(query string, args ...interface{}) (Rows, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
}

type Rows interface {
	Scan(...interface{}) error
	Close() error
	Next() bool
}
