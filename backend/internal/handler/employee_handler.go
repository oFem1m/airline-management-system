package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gorilla/mux"
)

// EmployeeHandler хранит методы для обработки HTTP-запросов к сотрудникам.
type EmployeeHandler struct {
	repo *repository.EmployeeRepository
}

// NewEmployeeHandler возвращает новый объект EmployeeHandler.
func NewEmployeeHandler(repo *repository.EmployeeRepository) *EmployeeHandler {
	return &EmployeeHandler{repo: repo}
}

// CreateEmployee обрабатывает POST-запрос для создания нового сотрудника.
func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

// DeleteEmployee обрабатывает DELETE-запрос для удаления сотрудника по ID.
func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID сотрудника", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID сотрудника", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAllEmployees обрабатывает GET-запрос для получения списка всех сотрудников.
func (h *EmployeeHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

// GetEmployee обрабатывает GET-запрос для получения конкретного сотрудника по ID.
func (h *EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID сотрудника", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID сотрудника", http.StatusBadRequest)
		return
	}

	emp, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if emp == nil {
		http.Error(w, "Сотрудник не найден", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

// UpdateEmployee обрабатывает PUT-запрос для изменения данных сотрудника.
func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID сотрудника", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID сотрудника", http.StatusBadRequest)
		return
	}

	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}
	// Убедимся, что ID из URL совпадает с ID в теле запроса
	emp.ID = id

	if err := h.repo.Update(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}
