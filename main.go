package main

import (
	"github.com/andriyor/testtask-backend/driver"
	"github.com/andriyor/testtask-backend/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	driver.Connect()
	//driver.DBConn.MustExec(repository.Schema)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handler.GetApplications)
	r.Post("/", handler.CreateApplication)
	http.ListenAndServe(":8123", r)
}
