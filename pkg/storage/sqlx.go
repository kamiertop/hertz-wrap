package storage

import "github.com/jmoiron/sqlx"

func InitSqlx() error {
	var (
		driver         string
		dataSourceName string
	)

	db, err := sqlx.Open(driver, dataSourceName)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return db.Ping()
}
