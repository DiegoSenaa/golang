package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringCon := "root:12345@tcp(127.0.0.1:3306)/devbook?charset=utf8&parseTime=True&loc=Local"
	db, er := sql.Open("mysql", stringCon)

	if er != nil {
		log.Fatal(er)
	}
	defer db.Close()

	if er = db.Ping(); er != nil {
		log.Fatal(er)
	}

	fmt.Println("Conexão aberta com o banco")

	rows, er := db.Query("select * from usuarios")

	if er != nil {
		log.Fatal(er)
	}

	defer rows.Close()

	fmt.Println(rows)
}
