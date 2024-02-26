package handlers

import (
	"net/http"
	"todolist/db"
	"todolist/handlers/middlewares"
	"todolist/util/session"
)

var IndexHandler = middlewares.Auth(true, http.HandlerFunc(indexHandler))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// get user's todolists
	user, ok := r.Context().Value("user").(*db.User)
	if !ok {
		ssn, err := session.Store.Get(r, session.SID)
		if err != nil {
			delete(ssn.Values, session.USER_ID)
		}
		http.Redirect(w, r, "/auth", http.StatusFound)
		return
	}

	q := db.New(db.Conn)
	lists, err := q.GetLists(r.Context(), user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	createdAtMap := make(map[string]string)
	for _, list := range lists {
		createdAtMap[list.ID] = list.CreatedAt.Time.Format("2006-Jan-02, Mon")
	}

	T.ExecuteTemplate(w, "index.html", map[string]any{"Lists": lists, "User": user, "CreatedAtMap": createdAtMap})
}
