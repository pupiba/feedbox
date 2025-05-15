package postgres

import (
	"encoding/json"
	"time"
)

type FeedbackStatus string

const (
	StatusPending  FeedbackStatus = "pending"
	StatusReviewed FeedbackStatus = "reviewed"
	StatusArchived FeedbackStatus = "archived"
)

type Feedback struct {
	ID        int             `json:"id"`
	FormID    int             `json:"form_id" validate:"required"`
	Data      json.RawMessage `json:"data" validate:"required"` // JSONB
	Status    FeedbackStatus  `json:"status"`
	CreatedAt time.Time       `json:"created_at"`
}
