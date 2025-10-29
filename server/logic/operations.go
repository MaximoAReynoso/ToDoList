package logic

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	sqlc "example.com/server/db/sqlc" // generado por sqlc
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
	tasks, _ := s.queries.ListTasks(context.Background())

	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(tasks)
}

// GET /tasks/{id} - Obtener elemento específico
func (s *Server) GetTask(w http.ResponseWriter, r *http.Request, id int) {
	task, err := s.queries.GetTask(context.Background(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	}

	w.Header().Set("Content-Type", "application/json")

	mu.Lock()
	defer mu.Unlock()

	json.NewEncoder(w).Encode(task)
}

// POST /tasks - Crear nuevo elemento
func (s *Server) CreateTasks(w http.ResponseWriter, r *http.Request) {
	var newTask sqlc.CreateTaskParams

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ValidateTask(newTask.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbTask, err := s.queries.CreateTask(context.Background(), sqlc.CreateTaskParams{
		Title:       newTask.Title,
		Description: newTask.Description,
		Completed:   newTask.Completed,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dbTask)

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
	err := s.queries.DeleteTask(context.Background(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	}

	mu.Lock()
	defer mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
