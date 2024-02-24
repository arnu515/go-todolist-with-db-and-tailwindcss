package handlers

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	println(T.DefinedTemplates())
	T.ExecuteTemplate(w, "index.html", nil)
}
