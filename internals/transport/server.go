package transport

import (
	"feedbox/internals/storage"
	"fmt"
	"net/http"
)

func NewServer(db *storage.StoragePG) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/forms", getForms)

	/*
		mux.HandleFunc("/api/forms", postForms)
		mux.HandleFunc("/api/forms/{id}", getFormById)
		mux.HandleFunc("/api/forms/{id}/schema", getSchemaByForm)
		mux.HandleFunc("/api/forms/{id}/submit", postSchemaData)
		mux.HandleFunc("/api/forms/{id}/responses", getStatus)
		mux.HandleFunc("/api/responses/{response_id}", updateStatus)
	*/

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
