package database

import (
	"app/pkg/interface/repository"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type sqlHandler struct {
	DB *sql.DB
}

func NewSQLHandler() repository.SQLHandler {
	conn, err := sql.Open("mysql", "root:password@tcp(localhost:3307)/note_app")
	if err != nil {
		log.Fatal(err)
	}

	var sh sqlHandler
	sh.DB = conn
	return &sh
}

func (sh *sqlHandler) Execute(query string, args ...interface{}) (repository.Result, error) {
	result, err := sh.DB.Exec(query, args...)
	if err != nil {
		return &sqlResult{}, err
	}
	return &sqlResult{
		Result: result,
	}, nil
}

type sqlResult struct {
	Result sql.Result
}

func (r *sqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r *sqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

func (sh *sqlHandler) QueryRow(query string, args ...interface{}) repository.Row {
	row := sh.DB.QueryRow(query, args...)
	return &sqlRow{
		Row: row,
	}
}

type sqlRow struct {
	Row *sql.Row
}

func (r *sqlRow) Scan(dest ...interface{}) error {
	return r.Row.Scan(dest...)
}

func (sh *sqlHandler) Query(query string, args ...interface{}) (repository.Rows, error) {
	rows, err := sh.DB.Query(query, args...)
	if err != nil {
		return &sqlRows{}, err
	}
	return &sqlRows{
		Rows: rows,
	}, nil
}

type sqlRows struct {
	Rows *sql.Rows
}

func (r *sqlRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r *sqlRows) Close() error {
	return r.Rows.Close()
}

func (r *sqlRows) Next() bool {
	return r.Rows.Next()
}
