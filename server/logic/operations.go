package logic

import (
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
	completed := completedText == "true"

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

	listadoCompleto, err := s.queries.ListTasks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	views.Listado(listadoCompleto).Render(r.Context(), w)
}

// PUT /tasks/{id} - Actualizar elemento
func (s *Server) UpdateTask(w http.ResponseWriter, r *http.Request, id int) {
	updatedTask, err := s.queries.GetTask(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTask.Completed = !updatedTask.Completed
	err = s.queries.UpdateTask(r.Context(), sqlc.UpdateTaskParams{
		ID:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Completed:   updatedTask.Completed,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotModified)
		return
	}

	views.Boton(updatedTask).Render(r.Context(), w)
}

// DELETE /tasks/{id} - Eliminar elemento
func (s *Server) DeleteTask(w http.ResponseWriter, r *http.Request, id int) {
	err := s.queries.DeleteTask(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	}

	w.WriteHeader(http.StatusOK)
}
