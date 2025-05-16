package transport

import (
	"fmt"
	"net/http"
)

// GET /api/forms
func getForms(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")

	if key != "" {
		fmt.Fprintln(w, "key: ", key)
	} else {
		fmt.Fprintln(w, "No Quary parm")
	}
}

// POST /api/forms
func postForms(w http.ResponseWriter, r *http.Request) {}

// GET /api/forms/{id}
func getFormById(w http.ResponseWriter, r *http.Request) {}

// GET /api/forms/{id}/schema
func getSchemaByForm(w http.ResponseWriter, r *http.Request) {}

// POST /api/forms/{id}/submit
func postSchemaData(w http.ResponseWriter, r *http.Request) {}

// GET /api/forms/{id}/responses
func getStatus(w http.ResponseWriter, r *http.Request) {}

// PATCH /api/responses/{response_id}
func updateStatus(w http.ResponseWriter, r *http.Request) {}
