package main

import (
	"fmt"
	"net/http"
)

func main() {
	index := "./pages"
	fileServer := http.FileServer(http.Dir(index))
	http.Handle("/", fileServer)

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
