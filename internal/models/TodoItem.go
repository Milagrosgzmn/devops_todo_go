package models

import (
	"errors"
	"strings"

	"github.com/Milagrosgzmn/devops_todo_go.git/internal/constants"
)

type TodoItem struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	State       string  `json:"state"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at,omitempty"`
}

// Validate valida los campos del TodoItem
func (t *TodoItem) Validate() error {
	// Validar título
	if strings.TrimSpace(t.Title) == "" {
		return errors.New("el título es requerido")
	}

	// Validar estado
	if strings.TrimSpace(t.State) == "" {
		return errors.New("el estado es requerido")
	}

	if !constants.IsValidState(t.State) {
		return errors.New(constants.EstadoInvalido)
	}

	return nil
}
