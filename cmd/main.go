package main

import (
	"log"
	"net/http"

	"todoapp/internal/db"
	"todoapp/internal/todo"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env

	err := godotenv.Load()
	if err != nil {
		log.Println(" no .env file found")
	}

	err = db.Connect()
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}
	defer db.Close()

	repo := todo.NewRepository(db.Pool)
	service := todo.NewService(repo)
	handler := todo.NewHandler(service)

	mux := mux.NewRouter()
	mux.HandleFunc("/todos", handler.GetTodos).Methods("GET")
	mux.HandleFunc("/todos", handler.CreateTodo).Methods("POST")
	mux.HandleFunc("/todos/id={id}", handler.UpdateTodo).Methods("PUT")
	mux.HandleFunc("/todos/id={id}", handler.DeleteTodo).Methods("DELETE")
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
