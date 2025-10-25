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
	connStr := "user=u password=p dbname=d"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer db.Close()
	queries := sqlc.New(db)
	server := logic.NewServer(queries)

	index := "./pages"
	fileServer := http.FileServer(http.Dir(index))
	http.Handle("/", fileServer)
	http.Handle("/tasks", logic.LoggingMiddleware(http.HandlerFunc(server.ElementHandler)))

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
