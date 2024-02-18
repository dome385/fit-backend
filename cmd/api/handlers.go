package main

import (
	"fit-backend/internal/models"
	"net/http"
)

func (app *application) AllÜbungen(w http.ResponseWriter, r *http.Request) {
	übungen, err := app.DB.AllÜbungen()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, übungen)
}

func (app *application) InsertÜbung(w http.ResponseWriter, r *http.Request) {
	var übung models.Übung

	err := app.readJSON(w, r, &übung)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_, err = app.DB.AddÜbung(übung)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "übung geupdated",
	}

	app.writeJSON(w, http.StatusAccepted, resp)

}
