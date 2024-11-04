package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/misshanya/go-todo-api/internal/services"
)

type TodoHandler struct {
	todoService *services.TodoService
}

func NewTodoHandler(todoService *services.TodoService) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.todoService.CreateTodo(r.Context(), req.Title, req.Content)
	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	fmt.Println(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := h.todoService.GetTodoByID(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Title   *string `json:"title"`
		Content *string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.todoService.UpdateTodo(r.Context(), int32(id), req.Title, req.Content)
	if err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TodoHandler) ListTodosByUpdatedAt(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.ListTodosByUpdatedAt(r.Context())
	if err != nil {
		http.Error(w, "Some problems", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	fmt.Println(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.todoService.DeleteTodo(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
