package main

import (
	"feedbox/internals/storage"
	"os"

	"fmt"

	"github.com/joho/godotenv"
	// "fmt"
	// "log"
	// "net/http"
	// "repo/internal/storage"
	// "repo/internal/transport"
)

func main() {

	if err := godotenv.Load("config/.env"); err != nil {
		panic(err)
	}

	// Инициализация хранилища
	db, err := storage.NewStoragePG(fmt.Sprintf("user=%s dbname=%s password=%s port=%s host=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST")))
	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	//////////// test for the projects table ////////////

	// var test postgres.Project = postgres.Project{
	// 	Title:       "hello",
	// 	Description: "hellohellohellohellohellohello",
	// }

	// for i := 1; i < 11; i++ {
	// 	test.Title = fmt.Sprintf("Проект №%d", i)
	// 	if err := test.Create(db.GetDB()); err != nil {
	// 		panic(err)
	// 	}
	// }

	// log.Println("Data added to projects table")

	/////////////////////////////////////////////////////

	// Запуск сервера
	// transport.NewServer(db)
	// fmt.Println("Starting server on :8080")
	// if err := http.ListenAndServe(":8080", server.Router()); err != nil {
	// 	log.Panic("Server failed:", err)
	// }

}
