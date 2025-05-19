package storage

import (
	"encoding/json"
	"feedbox/internals/storage/postgres"
	"time"
)

type ProjectRepository interface {
	Create(project *postgres.Project) error
	Delete(id int) error
	// Другие методы...
}

type FormRepository interface {
	Create(form *postgres.Form) error
	GetByID(id int) (*postgres.Form, error)
	// Другие методы...
}

type FeedbackRepository interface {
	// Методы для работы с feedback
}

type Form struct {
	ID          int             `json:"id"`
	ProjectID   int             `json:"project_id" validate:"required"`
	Title       string          `json:"title" validate:"required,max=64"`
	Description string          `json:"description" validate:"required"`
	Schema      json.RawMessage `json:"schema" validate:"required"` // JSONB
	CreatedAt   time.Time       `json:"created_at"`
}

type FeedbackStatus string

type Feedback struct {
	ID        int             `json:"id"`
	FormID    int             `json:"form_id" validate:"required"`
	Data      json.RawMessage `json:"data" validate:"required"` // JSONB
	Status    FeedbackStatus  `json:"status"`
	CreatedAt time.Time       `json:"created_at"`
}
