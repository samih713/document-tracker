package db

import (
	"database/sql"
	_ "github.com/ncruces/go-sqlite3/driver"
)

const (
	CONNECTION = "file"
	DRIVERNAME = "sqlite3"
)

type DataBase struct {
}

func Open(path string) (*sql.DB, error) {
	dataSourceName := CONNECTION + ":" + path
	database, err := sql.Open(DRIVERNAME, dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := database.Ping(); err != nil {
		return nil, err
	}

	_, err = database.Exec(`PRAGMA foreign_keys = ON;`) // enforce foreign key checking
	if err != nil {
		return nil, err
	}

	return database, err
}

func Init(database *sql.DB) error {
	_, err := database.Exec(schemaSQL)
	return err
}
