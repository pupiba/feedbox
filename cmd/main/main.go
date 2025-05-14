package main

import (
	"feedbox/internals/transport"
	// "fmt"
	// "log"
	// "net/http"
	// "repo/internal/storage"
	// "repo/internal/transport"
)

func main() {
	// Инициализация хранилища
	// store := storage.NewPostgresStorage()

	// server := transport.NewServer(store)

	// Запуск сервера
	transport.NewServer()
	// fmt.Println("Starting server on :8080")
	// if err := http.ListenAndServe(":8080", server.Router()); err != nil {
	// 	log.Panic("Server failed:", err)
	// }
}
