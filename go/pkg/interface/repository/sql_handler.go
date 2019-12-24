package repository

type SQLHandler interface {
	QueryRow(query string, args ...interface{}) Row
}

type Row interface {
	Scan(...interface{}) error
}
