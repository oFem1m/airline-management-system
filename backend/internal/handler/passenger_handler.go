package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gorilla/mux"
)

// PassengerHandler хранит методы для обработки HTTP-запросов к пассажирам.
type PassengerHandler struct {
	repo *repository.PassengerRepository
}

// NewPassengerHandler возвращает новый обработчик для пассажиров
func NewPassengerHandler(repo *repository.PassengerRepository) *PassengerHandler {
	return &PassengerHandler{repo: repo}
}

// CreatePassenger обрабатывает POST-запрос для создания нового пассажира.
func (h *PassengerHandler) CreatePassenger(w http.ResponseWriter, r *http.Request) {
	var p models.Passenger
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// GetAllPassengers обрабатывает GET-запрос для получения списка всех пассажиров.
func (h *PassengerHandler) GetAllPassengers(w http.ResponseWriter, r *http.Request) {
	passengers, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(passengers)
}

// GetPassenger обрабатывает GET-запрос для получения пассажира по ID.
func (h *PassengerHandler) GetPassenger(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID пассажира", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID пассажира", http.StatusBadRequest)
		return
	}
	passenger, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if passenger == nil {
		http.Error(w, "Пассажир не найден", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(passenger)
}

// UpdatePassenger обрабатывает PUT-запрос для обновления данных пассажира по ID.
func (h *PassengerHandler) UpdatePassenger(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID пассажира", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID пассажира", http.StatusBadRequest)
		return
	}
	var p models.Passenger
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	p.ID = id
	if err := h.repo.Update(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// DeletePassenger обрабатывает DELETE-запрос для удаления пассажира по ID.
func (h *PassengerHandler) DeletePassenger(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID пассажира", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID пассажира", http.StatusBadRequest)
		return
	}
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
