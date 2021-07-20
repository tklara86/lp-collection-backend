package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/albums", app.createAlbumHandler)
	router.HandlerFunc(http.MethodGet, "/v1/albums/:id", app.showAlbumHandler)

	return router
}