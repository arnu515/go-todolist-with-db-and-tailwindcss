package handlers

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	T.ExecuteTemplate(w, "index.html", nil)
}
