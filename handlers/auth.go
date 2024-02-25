package handlers

import (
	"io"
	"net/http"
	"net/url"
	"todolist/db"
	"todolist/util/session"

	"github.com/oklog/ulid/v2"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		T.ExecuteTemplate(w, "auth.html", nil)
		return
	}

	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}
		data, err := url.ParseQuery(string(body))
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			T.ExecuteTemplate(w, "auth.html", map[string]string{"Error": "Invalid body. Please try again."})
			return
		}
		username, password := data.Get("username"), data.Get("password")
		msg := ""
		if len(username) < 4 && len(username) > 64 {
			msg = "The username must be between 4 to 64 characters long."
		}
		if len(password) == 0 {
			msg = "Please enter a password."
		}
		if msg != "" {
			w.WriteHeader(http.StatusBadRequest)
			T.ExecuteTemplate(w, "auth.html", map[string]string{"Username": username, "Password": password, "Error": msg})
			return
		}

		// check if user exists
		q := db.New(db.Conn)
		user, err := q.GetUserByUsername(r.Context(), username)
		if err != nil {
			if err.Error() == "no rows in result set" {
				user, err = q.CreateUser(r.Context(), db.CreateUserParams{ID: ulid.Make().String(), Username: username, Password: password})
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					T.ExecuteTemplate(w, "auth.html", map[string]string{"Username": username, "Password": password, "Error": err.Error()})
					return
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				T.ExecuteTemplate(w, "auth.html", map[string]string{"Username": username, "Password": password, "Error": err.Error()})
				return
			}
		}

		ssn, err := session.Store.Get(r, session.SID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			T.ExecuteTemplate(w, "auth.html", map[string]string{"Username": username, "Password": password, "Error": err.Error()})
			return
		}

		ssn.Values[session.USER_ID] = user.ID
		ssn.Save(r, w)

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}
