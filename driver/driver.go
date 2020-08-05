package driver

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var DBConn *sqlx.DB

func Connect() *sqlx.DB {
	db, err := sqlx.Connect("postgres", os.Getenv("PG_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	DBConn = db
	return db
}
