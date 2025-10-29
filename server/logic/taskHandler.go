package logic

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Manejador para /elements y /elements/{id}
func (s *Server) ElementHandler(w http.ResponseWriter, r *http.Request) {
	// Extraer ID del path
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 3 {
		switch r.Method {
		case http.MethodGet:
			s.GetTasks(w, r)
		case http.MethodPost:
			s.CreateTasks(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else {
		id, err := strconv.Atoi(parts[2])
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			s.GetTask(w, r, id)
		case http.MethodPut:
			s.UpdateTask(w, r, id)
		case http.MethodDelete:
			s.DeleteTask(w, r, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// Middleware de logging
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
