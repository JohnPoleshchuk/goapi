package main

import (
	"encoding/json"
	"net/http"

	"github.com/JohnPoleshchuk/goapi/internal/store"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	var data store.Post
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = app.store.Posts.Create(r.Context(), &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
