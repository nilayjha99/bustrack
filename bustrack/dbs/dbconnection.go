package dbs

import (
	"bustrack/tools"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres123"
	dbname   = "bustrack"
)

var db *sql.DB
var err error

func GetDB() *sql.DB {
	if db == nil {
		connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			utils.PanicIf(err)
		}
	}

	return db
}

func CloseDB(db *sql.DB) {
	err = db.Close()
	utils.PanicIf(err)
}
