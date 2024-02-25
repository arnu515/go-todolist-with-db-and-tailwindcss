package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"

	"todolist/util"
)

var Conn = util.Must(pgx.Connect(context.Background(), util.Default(os.Getenv("DATABASE_URL"), "postgres://postgres:postgres@localhost:5432/postgres")))
