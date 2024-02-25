package handlers

import "net/http"

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/{$}", IndexHandler)
	mux.HandleFunc("/auth/{$}", AuthHandler)

	mux.Handle("/", http.FileServer(http.Dir("./static")))
}
