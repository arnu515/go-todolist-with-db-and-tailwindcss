package handlers

import "net/http"

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/{$}", IndexHandler)
	mux.HandleFunc("/auth/{$}", AuthHandler)
	mux.HandleFunc("/list/{listId}/{$}", ListHandler)

	mux.Handle("/", http.FileServer(http.Dir("./static")))
}
