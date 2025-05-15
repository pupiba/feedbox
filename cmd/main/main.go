package main

import (
	"feedbox/internals/storage"
	"feedbox/internals/transport"
	"fmt"
	// "fmt"
	// "log"
	// "net/http"
	// "repo/internal/storage"
	// "repo/internal/transport"
)

func main() {
	// Инициализация хранилища
	dsn := "user=pupiba dbname=feedbox password=qwerty port=5432 host=127.0.0.1"
	db, err := storage.NewStoragePG(dsn)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.InitTables(); err != nil {
		fmt.Printf("Failed to initialize tables: %v", err)
	}

	// Запуск сервера
	transport.NewServer()
	// fmt.Println("Starting server on :8080")
	// if err := http.ListenAndServe(":8080", server.Router()); err != nil {
	// 	log.Panic("Server failed:", err)
	// }
}
