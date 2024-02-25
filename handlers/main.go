package handlers

import "net/http"

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/{$}", IndexHandler)

	mux.Handle("/", http.FileServer(http.Dir("./static")))
}
