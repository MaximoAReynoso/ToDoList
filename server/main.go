package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	sqlc "example.com/server/db/sqlc" // generado por sqlc
	"example.com/server/logic"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=user password=pass dbname=database host=db port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer db.Close()
	queries := sqlc.New(db)
	server := logic.NewServer(queries)

	http.Handle("/", logic.LoggingMiddleware(http.HandlerFunc(server.GetTasks)))
	http.Handle("/tasks", logic.LoggingMiddleware(http.HandlerFunc(server.ElementHandler)))
	http.Handle("/tasks/", logic.LoggingMiddleware(http.HandlerFunc(server.ElementHandler)))

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
