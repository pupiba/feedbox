package transport

import (
	"fmt"
	"net/http"
)

func NewServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /forms", getForms)
	mux.HandleFunc("POST /api/forms", postForms)
	mux.HandleFunc("GET /api/forms/{id}", getFormById)
	mux.HandleFunc("GET /api/forms/{id}/schema", getSchemaByForm)
	mux.HandleFunc("POST /api/forms/{id}/submit", postSchemaData)
	mux.HandleFunc("GET /api/forms/{id}/responses", getStatus)
	mux.HandleFunc("PATCH /api/responses/{response_id}", updateStatus)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
