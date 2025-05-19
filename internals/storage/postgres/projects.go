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

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(project *Project) error {
	return r.db.QueryRow(
		`INSERT INTO projects (title, description)
        VALUES ($1, $2)
        RETURNING id, created_at`,
		project.Title, project.Description,
	).Scan(&project.ID, &project.CreatedAt)
}

func (r *ProjectRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM projects WHERE id = $1`, id)
	return err
}
