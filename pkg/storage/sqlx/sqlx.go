package sqlx

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitSqlx() (err error) {
	DB, err = sqlx.Open("mysql", "root:root@tcp(192.168.1.6:3306)/test?charset=utf8")
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("db.Ping() error: %w", err)
	}

	return nil
}

func Close() error {
	if err := DB.Close(); err != nil {
		return fmt.Errorf("close db error: %w", err)
	}
	return nil
}
