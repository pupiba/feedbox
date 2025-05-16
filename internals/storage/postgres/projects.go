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
		`INSERT INTO projects (title, description)
		VALUES ($1, $2)
		RETURNING id, created_at`,
		p.Title, p.Description,
	).Scan(&p.ID, &p.CreatedAt)
	return err
}

func DeleteProject(db *sql.DB, id int) error {
	_, err := db.Exec(
		`DELETE FROM projects WHERE id = $1`,
		id,
	)
	return err
}
