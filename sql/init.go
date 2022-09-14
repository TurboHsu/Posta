package sql

import (
	s "database/sql"

	"github.com/TurboHsu/Posta/config"
	"github.com/TurboHsu/Posta/util"
	_ "github.com/mattn/go-sqlite3"
)

var DB *s.DB

func Init() {
	var err error
	switch config.C.SQL.Engine {
	case "SQLite":
		DB, err = s.Open("sqlite3", config.C.SQL.SQLite.File)
		util.HandleErr(err)
	case "MySQL":
		//In Progress
	default:
		util.LogFatal("Database not specified.")
	}
}
