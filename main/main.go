package main

import (
	"github.com/RichardKyire/ormazing"
	"github.com/RichardKyire/ormazing/log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := ormazing.NewEngine("sqlite3", "ormazing.db")
	defer engine.Close()
	session := engine.NewSession()
	session.FillSql("drop table if exists user;").Execute()
	session.FillSql("create table user(name text);").Execute()
	result, _ := session.FillSql("insert into user(`name`) values(?)", "wzt").Execute()
	count, _ := result.RowsAffected()
	log.Info("insert data success,%d affected count ", count)
}
