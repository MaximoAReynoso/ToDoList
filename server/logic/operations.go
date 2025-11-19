package logic

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	sqlc "example.com/server/db/sqlc" // generado por sqlc
	"example.com/server/views"
)

var mu sync.Mutex

// Server contiene la instancia de sqlc
type Server struct {
	queries *sqlc.Queries
}

// Constructor
func NewServer(queries *sqlc.Queries) *Server {
	return &Server{queries: queries}
}

// GET /tasks - Listar todos los elementos
func (s *Server) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, _ := s.queries.ListTasks(r.Context())

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.Index(tasks).Render(r.Context(), w)
}

// GET /tasks/{id} - Obtener elemento específico
func (s *Server) GetTask(w http.ResponseWriter, r *http.Request, id int) {
	task, err := s.queries.GetTask(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.Index([]sqlc.Task{task}).Render(r.Context(), w)
}

// POST /tasks - Crear nuevo elemento
func (s *Server) CreateTasks(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	description := r.FormValue("description")
	completedText := r.FormValue("completed")
	completed := completedText == "on"

	err = ValidateTask(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = s.queries.CreateTask(r.Context(), sqlc.CreateTaskParams{
		Title:       title,
		Description: description,
		Completed:   completed,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// PUT /tasks/{id} - Actualizar elemento
func (s *Server) UpdateTask(w http.ResponseWriter, r *http.Request, id int) {
	var updatedTask sqlc.UpdateTaskParams
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedTask.ID = int32(id)

	mu.Lock()
	defer mu.Unlock()

	err2 := s.queries.UpdateTask(context.Background(), updatedTask)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)

}

// DELETE /tasks/{id} - Eliminar elemento
func (s *Server) DeleteTask(w http.ResponseWriter, r *http.Request, id int) {
	err := s.queries.DeleteTask(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	}

	w.WriteHeader(http.StatusNoContent)
}
