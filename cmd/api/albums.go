package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// createAlbumHandler POST request to /v1/albums endpoint
func (app *application) createAlbumHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new album")
}

// showAlbumHandler GET request to /v1/albums/:id endpoint
func (app *application) showAlbumHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 32)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of album %d\n", id)
}