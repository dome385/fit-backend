package main

import "net/http"

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	übungen, err := app.DB.AllÜbungen()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, übungen)
}
