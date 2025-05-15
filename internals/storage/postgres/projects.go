package postgres

import (
	"database/sql"
	"time"
)

type Project struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" validate:"required,max=64"`
	Description string    `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

func (p *Project) Create(db *sql.DB) error {
	err := db.QueryRow(
		`INSERT INTO PROJECTS (title, description)
		VALUES ($1, $2)
		RETURNING id, created_at`,
		p.Title, p.Description,
	).Scan(&p.ID, &p.CreatedAt)
	return err
}
