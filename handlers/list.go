// /:listId

package handlers

import (
	"net/http"
	"todolist/db"
	"todolist/handlers/middlewares"
	"todolist/util/session"
)

var ListHandler = middlewares.Auth(true, http.HandlerFunc(listHandler))

func listHandler(w http.ResponseWriter, r *http.Request) {
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
	list, err := q.GetList(r.Context(), db.GetListParams{ID: r.PathValue("listId"), UserID: user.ID})
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	todos, err := q.GetTodos(r.Context(), db.GetTodosParams{UserID: user.ID, ListID: list.ID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	T.ExecuteTemplate(w, "list.html", map[string]any{"List": list, "User": user, "Todos": todos})
}
