package models

import (
	"time"
)

type Application struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Key     string    `json:"key"`
	Created time.Time `json:"created"`
}
