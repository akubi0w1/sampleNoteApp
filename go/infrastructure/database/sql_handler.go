package database

import (
	"database/sql"
	"fmt"
	"log"
	"note-app/config"
	"note-app/interface/datastore"

	_ "github.com/go-sql-driver/mysql"
)

// SQLHandler sqlの操作
type SQLHandler struct {
	DB *sql.DB
}

// NewSQLHandler sqlハンドラの作成
func NewSQLHandler() datastore.SQLHandler {
	conf := config.LoadDBConfig()
	conn, err := sql.Open(conf.Driver,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			conf.User,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Database,
		))
	if err != nil {
		log.Fatal(err)
	}
	handler := new(SQLHandler)
	handler.DB = conn
	return handler
}

// Execute impl: sql.Exec
func (sh *SQLHandler) Execute(query string, args ...interface{}) (datastore.Result, error) {
	result, err := sh.DB.Exec(query, args...)
	if err != nil {
		return &SQLResult{}, err
	}
	return &SQLResult{Result: result}, nil
}

// Query impl: sql.Query
func (sh *SQLHandler) Query(query string, args ...interface{}) (datastore.Rows, error) {
	rows, err := sh.DB.Query(query, args...)
	if err != nil {
		return &SQLRows{}, err
	}
	return &SQLRows{Rows: rows}, nil
}

// QueryRow impl: sql.QueryRow
func (sh *SQLHandler) QueryRow(query string, args ...interface{}) datastore.Row {
	row := sh.DB.QueryRow(query, args...)
	return &SQLRow{Row: row}
}

// SQLResult result
type SQLResult struct {
	Result sql.Result
}

// LastInsertId impl: result.lastinsertid
func (r *SQLResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

// RowsAffected impl: result.rowsaffected
func (r *SQLResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

// SQLRows sql.rows
type SQLRows struct {
	Rows *sql.Rows
}

// Scan impl: rows.scan
func (r *SQLRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

// Next impl: rows.next
func (r *SQLRows) Next() bool {
	return r.Rows.Next()
}

// Close impl: rows.close
func (r *SQLRows) Close() error {
	return r.Rows.Close()
}

// SQLRow sql.row
type SQLRow struct {
	Row *sql.Row
}

// Scan impl: row.scan
func (r *SQLRow) Scan(dest ...interface{}) error {
	return r.Row.Scan(dest...)
}
