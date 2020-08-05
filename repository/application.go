package repository

import (
	"fmt"
	"github.com/andriyor/testtask-backend/driver"
	"github.com/andriyor/testtask-backend/models"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

var Schema = `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE application (
	id     uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text,
    key text,
    created timestamp
);
`

func GetApplications() []models.Application {
	var people []models.Application
	query := fmt.Sprintf("SELECT id, PGP_SYM_DECRYPT(key::bytea, '%s') as key , name, created FROM  application",
		os.Getenv("ENCRYPT_PASSWORD"))
	driver.DBConn.Select(&people, query)
	return people
}

func CreateApplication(application models.Application) models.Application {
	tx := driver.DBConn.MustBegin()
	var people models.Application
	query := fmt.Sprintf("INSERT INTO application (name, key, created) VALUES ($1, PGP_SYM_ENCRYPT($2,'AES_KEY'), $3)"+
		" RETURNING id, PGP_SYM_DECRYPT(key::bytea, '%s') as key , name, created", os.Getenv("ENCRYPT_PASSWORD"))
	err := tx.QueryRowx(query, application.Name, application.Key, time.Now()).StructScan(&people)

	if err != nil {
		log.Fatalln(err)
	}

	_ = tx.Commit()
	return people
}
