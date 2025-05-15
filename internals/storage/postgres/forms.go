package postgres

import (
	"encoding/json"
	"time"
)

type Form struct {
	ID          int             `json:"id"`
	ProjectID   int             `json:"project_id" validate:"required"`
	Title       string          `json:"title" validate:"required,max=64"`
	Description string          `json:"description" validate:"required"`
	Schema      json.RawMessage `json:"schema" validate:"required"` // JSONB
	CreatedAt   time.Time       `json:"created_at"`
}
