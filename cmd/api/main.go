package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/misshanya/go-todo-api/internal/db"
	"github.com/misshanya/go-todo-api/internal/handlers"
	"github.com/misshanya/go-todo-api/internal/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("could not connect to db")
	}
	defer conn.Close(ctx)

	queries := db.New(conn)
	todoService := services.NewTodoService(queries)
	todoHandler := handlers.NewTodoHandler(todoService)

	router.Post("/todo", todoHandler.CreateTodo)
	router.Get("/todo/{id}", todoHandler.GetTodoByID)
	router.Put("/todo/{id}", todoHandler.UpdateTodo)
	router.Get("/todo", todoHandler.ListTodosByUpdatedAt)
	router.Delete("/todo/{id}", todoHandler.DeleteTodo)

	http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), router)
}
