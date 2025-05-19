package postgres

import (
	"database/sql"
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

type FormRepository struct {
	db *sql.DB
}

func NewFormRepository(db *sql.DB) *FormRepository {
	return &FormRepository{db: db}
}

func (r *FormRepository) CreateForm(form Form, db *sql.DB) error {
	query := `
		INSERT INTO forms (project_id, title, description, schema)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	return db.QueryRow(
		query,
		form.ProjectID,
		form.Title,
		form.Description,
		form.Schema, // JSONB автоматически преобразуется
	).Scan(&form.ID, &form.CreatedAt)
}

func (r *FormRepository) GetFormByID(db *sql.DB, id int) (*Form, error) {
	form := &Form{}
	query := `
		SELECT id, project_id, title, description, schema, created_at
		FROM forms
		WHERE id = $1
	`

	err := db.QueryRow(query, id).Scan(
		&form.ID,
		&form.ProjectID,
		&form.Title,
		&form.Description,
		&form.Schema, // JSONB автоматически сканируется
		&form.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return form, nil
}
