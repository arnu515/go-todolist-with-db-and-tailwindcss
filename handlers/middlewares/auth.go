package middlewares

import (
	"context"
	"net/http"
	"todolist/util/session"
)

func Auth(required bool, handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the session
		ssn, err := session.Store.Get(r, session.SID)
		if err != nil {
			http.Redirect(w, r, "/auth", http.StatusFound)
			return
		}

		// Get user from cookie
		user := session.GetUserFromSession(ssn)
		if user == nil {
			http.Redirect(w, r, "/auth", http.StatusFound)
			return
		}
		req := r.WithContext(context.WithValue(r.Context(), "user", user))

		handler.ServeHTTP(w, req)
	}
}
