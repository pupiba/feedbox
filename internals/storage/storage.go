package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Хранилище
type StoragePG struct {
	db *sql.DB
}

// интерфейс хранилища
type Storage interface {
	// Получение списка всех форм
	GetAllForms() ([]Form, error)
	// Создание новой формы
	CreateForm(form Form) (Form, error)
	// Получение одной формы по ID
	GetFormByID(id string) (Form, error)
	// Получение схемы формы по ID формы
	GetFormSchema(formID string) (map[string]interface{}, error)
	// Отправка данных пользователем (создание ответа на форму)
	SubmitFormResponse(formID string, response Feedback) (Feedback, error)
	// Получение всех ответов по ID формы
	GetFormResponses(formID string) ([]Feedback, error)
	// Обновление статуса ответа
	UpdateResponseStatus(responseID string, status string) (Feedback, error)
	// Очистка БД
	ClearDB() error
	// Закрытие соединения с БД (если нужно)
	Close() error
}

// Конструктор хранилища
func NewStoragePG(dsn string) (*StoragePG, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	tables := []string{
		`CREATE TABLE IF NOT EXISTS projects (
			id SERIAL PRIMARY KEY,
			title VARCHAR(64) NOT NULL,
			description TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,

		`CREATE TABLE IF NOT EXISTS forms (
			id SERIAL PRIMARY KEY,
			project_id INTEGER NOT NULL,
			title VARCHAR(64) NOT NULL,
			description TEXT NOT NULL,
			schema JSONB NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
		)`,

		`CREATE TABLE IF NOT EXISTS feedback (
			id SERIAL PRIMARY KEY,
			form_id INTEGER NOT NULL,
			data JSONB NOT NULL,
			status VARCHAR(16) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (form_id) REFERENCES forms(id) ON DELETE CASCADE
		)`,
	}

	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			return nil, fmt.Errorf("failed to create table: %v", err)
		}
	}

	log.Println("Database connection established")
	return &StoragePG{db: db}, nil
}

func (d *StoragePG) ClearDB() error {
	if _, err := d.db.Exec(`DROP TABLE IF EXISTS projects, forms, feedback CASCADE;`); err != nil {
		return fmt.Errorf("failed to clear database: %v", err)
	}
	log.Println("The database tables have been deleted")
}

func (d *StoragePG) Close() error {
	return d.db.Close()
}

func (d *StoragePG) GetAllForms() ([]Form, error) {

}

func (d *StoragePG) CreateForm(form Form) (Form, error) {

}

func (d *StoragePG) GetFormByID(id string) (Form, error) {

}

func (d *StoragePG) GetFormSchema(formID string) (map[string]interface{}, error) {

}

func (d *StoragePG) SubmitFormResponse(formID string, response Feedback) (Feedback, error) {

}

func (d *StoragePG) GetFormResponses(formID string) ([]Feedback, error) {

}

func (d *StoragePG) UpdateResponseStatus(responseID string, status string) (Feedback, error) {

}
