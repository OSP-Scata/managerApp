package helpers

import (
	"database/sql"
	"fmt"
)

func DBInit() (*sql.DB, error) {
	var (
		err error
		db  *sql.DB
	)
	connStr := "user=postgres password=password port=5432 dbname=assessment_manager sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to PostgreSQL.")
	return db, err
}
