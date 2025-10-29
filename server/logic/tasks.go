package logic

import "errors"

type Task struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func ValidateTask(t string) error {
	if t == "" {
		return errors.New("formato de task invalido")
	}
	return nil
}
