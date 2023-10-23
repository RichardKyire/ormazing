package session

import (
	"database/sql"
	"github.com/RichardKyire/ormazing/log"
	"strings"
)

type Session struct {
	db     *sql.DB
	sql    strings.Builder
	params []interface{}
}

func New(db *sql.DB) *Session {
	return &Session{db: db}
}

func (session *Session) Clear() {
	session.sql.Reset()
	session.params = nil
}

func (session *Session) FillSql(sql string, params ...interface{}) *Session {
	session.sql.WriteString(sql)
	session.params = append(session.params, params...)
	return session

}

func (session *Session) Execute() (result sql.Result, err error) {
	defer session.Clear()
	log.Info("execute sql:", session.sql.String(), " params:", session.params)
	if result, err = session.db.Exec(session.sql.String(), session.params...); err != nil {
		log.Error("execute sql err:", err)
	}
	return
}

func (session *Session) QueryRow() (row *sql.Row) {
	defer session.Clear()
	log.Info("execute sql:", session.sql.String(), " params:", session.params)
	row = session.db.QueryRow(session.sql.String(), session.params...)
	return
}

func (session *Session) Query() (rows *sql.Rows) {
	defer session.Clear()
	log.Info("execute sql:", session.sql.String(), ", params:", session.params)
	rows, err := session.db.Query(session.sql.String(), session.params...)
	if err != nil {
		log.Error(err)
	}
	return
}
