package ormazing

import (
	"github.com/RichardKyire/ormazing/log"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestEngine(t *testing.T) {
	engine, err := NewEngine("sqlite3", "ormazing.db")
	if err != nil {
		log.Error("init ormazing err")
	}
	session := engine.NewSession()
	session.Clear()
	engine.Close()
}

func TestCreateTable(t *testing.T) {
	engine, err := NewEngine("sqlite3", "ormazing.db")
	if err != nil {
		log.Error("init ormazing err")
	}
	defer engine.Close()
	session := engine.NewSession()
	session.FillSql("create table user(name text)")
	result, err := session.Execute()
	row, _ := result.RowsAffected()
	log.Info("execute sql success,affected row is ", row)
}

func TestInsert(t *testing.T) {
	engine, err := NewEngine("sqlite3", "ormazing.db")
	if err != nil {
		log.Error("init ormazing err")
	}
	defer engine.Close()
	session := engine.NewSession()
	session.FillSql("insert into user(name) values(?)", "wzt", "lb")
	result, err := session.Execute()
	row, _ := result.RowsAffected()
	log.Info("execute sql success,affected row is ", row)
}

func TestQueryCount(t *testing.T) {
	engine, err := NewEngine("sqlite3", "ormazing.db")
	if err != nil {
		log.Error("init ormazing err")
	}
	defer engine.Close()
	session := engine.NewSession()
	row := session.FillSql("select count(1) from user").QueryRow()
	var count int
	if err := row.Scan(&count); err != nil {
		log.Error("query row error:", err)
	}
	log.Info("query row success, count is ", count)
}

func TestQuery(t *testing.T) {
	engine, err := NewEngine("sqlite3", "ormazing.db")
	if err != nil {
		log.Error("init ormazing err")
	}
	defer engine.Close()
	session := engine.NewSession()
	rows := session.FillSql("select `name` from user").Query()
	var (
		name string
	)
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Error(err)
		}
		log.Info("the data from user table is ", name)
	}

	err = rows.Err()
	if err != nil {
		log.Error(err)
	}
}
