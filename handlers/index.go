package handlers

import (
	"net/http"
	"todolist/handlers/middlewares"
)

var IndexHandler = middlewares.Auth(true, http.HandlerFunc(indexHandler))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	T.ExecuteTemplate(w, "index.html", nil)
}
