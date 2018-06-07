package util

import (
	"database/sql"
)

// OpenMySQLConnection devolve uma conexão válida com um banco MySQL.
func OpenMySQLConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
