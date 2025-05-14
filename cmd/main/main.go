package main

import (
	"log"
	"net/http"
	"repo/internal/storage"
	"repo/internal/transport"
)

func main() {
	// Инициализация хранилища
	store := storage.NewPostgresStorage()

	// Создание HTTP сервера
	server := http.NewServer(store)

	// Запуск сервера
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", server.Router()); err != nil {
		log.Fatal("Server failed:", err)
	}
}
