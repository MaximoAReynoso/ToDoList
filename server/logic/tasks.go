package logic

import "errors"

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func ValidateTask(t Task) error {
	if t.Title == "" {
		return errors.New("formato de producto invalido")
	}
	return nil
}
