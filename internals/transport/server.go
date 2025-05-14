package transport

import (
	"fmt"
	"net/http"
)

func NewServer() {
	http.HandleFunc("/", handler)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
