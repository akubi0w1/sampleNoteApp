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
