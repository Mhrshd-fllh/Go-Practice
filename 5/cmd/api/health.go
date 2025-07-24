package main

import (
	"net/http"
	"project/internal/store"
)

func (app *application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))

	app.store.Posts.Create(r.Context())

}
