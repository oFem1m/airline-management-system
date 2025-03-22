package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gorilla/mux"
)

type AircraftHandler struct {
	repo *repository.AircraftRepository
}

// NewAircraftHandler возвращает новый обработчик для самолётов
func NewAircraftHandler(repo *repository.AircraftRepository) *AircraftHandler {
	return &AircraftHandler{repo: repo}
}

// CreateAircraft обрабатывает POST-запрос на создание нового самолёта
func (h *AircraftHandler) CreateAircraft(w http.ResponseWriter, r *http.Request) {
	var aircraft models.Aircraft

	if err := json.NewDecoder(r.Body).Decode(&aircraft); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&aircraft); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(aircraft)
	if err != nil {
		return
	}
}

// DeleteAircraft обрабатывает DELETE-запрос на удаление самолёта по ID
func (h *AircraftHandler) DeleteAircraft(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID самолёта", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID самолёта", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
