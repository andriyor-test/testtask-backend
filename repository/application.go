package repository

import (
	"github.com/andriyor/testtask-backend/driver"
	"github.com/andriyor/testtask-backend/models"
	_ "github.com/lib/pq"
	"time"
)

var Schema = `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE application (
	id     uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text,
    key text,
    created timestamp
);
`

func GetApplications() []models.Application {
	var people []models.Application
	driver.DBConn.Select(&people, "SELECT * FROM  application")
	return people
}

func CreateApplication(application models.Application) models.Application {
	tx := driver.DBConn.MustBegin()
	var people models.Application
	err := tx.QueryRowx("INSERT INTO application (name, key, created) VALUES ($1, $2, $3) RETURNING *",
		application.Name, application.Key, time.Now()).StructScan(&people)

	if err != nil {
		panic(err.Error())
	}

	tx.Commit()
	return people
}
