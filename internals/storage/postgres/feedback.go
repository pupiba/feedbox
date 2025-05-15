package postgres

import (
	"time"
)

type Project struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" validate:"required,max=64"`
	Description string    `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
}
