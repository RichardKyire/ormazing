package ormazing

import (
	"database/sql"
	"github.com/RichardKyire/ormazing/log"
	"github.com/RichardKyire/ormazing/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver string, dataSource string) (engine *Engine, err error) {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		log.Error(err)
		return
	}

	engine = &Engine{db: db}
	log.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	err := engine.db.Close()
	if err != nil {
		log.Error("close the engine fail")
	} else {
		log.Info("close the engine success")
	}
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
