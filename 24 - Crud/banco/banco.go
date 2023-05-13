package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	stringCon := "root:12345@tcp(127.0.0.1:3306)/devbook?charset=utf8&parseTime=True&loc=Local"
	db, er := sql.Open("mysql", stringCon)

	if er != nil {
		return nil, er
	}

	if er = db.Ping(); er != nil {
		return nil, er
	}

	return db, nil
}
