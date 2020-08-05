package handler

import (
	"encoding/json"
	"github.com/andriyor/testtask-backend/models"
	"github.com/andriyor/testtask-backend/repository"
	"io/ioutil"
	"log"
	"net/http"
)

func GetApplications(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts := repository.GetApplications()
	respondWithJSON(w, http.StatusOK, posts)
}

func CreateApplication(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var post models.Application
	err = json.Unmarshal(body, &post)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}

	newPost := repository.CreateApplication(post)
	respondWithJSON(w, http.StatusOK, newPost)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
