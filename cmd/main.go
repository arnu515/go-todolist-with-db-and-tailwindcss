package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"todolist/db"
	"todolist/handlers"
)

func main() {
	mux := http.NewServeMux()

	defer db.Conn.Close(context.Background())
	handlers.RegisterHandlers(mux)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}

	log.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
