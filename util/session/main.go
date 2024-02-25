package session

import (
	"context"
	"os"
	"todolist/db"
	"todolist/util"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte(util.Default(os.Getenv("SESSION_SECRET"), "secret-value")))

func GetUserFromSession(session *sessions.Session) *db.User {
	if userId, ok := session.Values["userId"].(string); !ok {
		return nil
	} else {
		// fetch user from db
		q := db.New(db.Conn)
		user, err := q.GetUser(context.Background(), userId)
		if err != nil {
			println(err.Error())
			return nil
		}
		return &user
	}
}
