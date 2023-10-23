package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, _ := sql.Open("sqlite3", "test.db")
	defer func() { _ = db.Close() }()
	if db == nil {
		log.Fatal("connecting db fail")
	}
	_, _ = db.Exec("create table helloworld(name text);")
	result, err := db.Exec("insert into helloworld(`name`) values (?)", "wzt")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println("insert ", affected, " rows success")
	}
	row := db.QueryRow("select name from helloworld limit 1")
	var name string
	if err := row.Scan(&name); err == nil {
		log.Println("data from db is ", name)
	}
}
