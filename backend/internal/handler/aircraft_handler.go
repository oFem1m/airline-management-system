package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gorilla/mux"
)

// AircraftHandler хранит методы для обработки HTTP-запросов к самолетам.
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

// DeleteAircraft обрабатывает POST-запрос на удаление самолёта по ID
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

// GetAllAircrafts обрабатывает GET-запрос и возвращает список всех самолётов
func (h *AircraftHandler) GetAllAircrafts(w http.ResponseWriter, r *http.Request) {
	aircrafts, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(aircrafts); err != nil {
		http.Error(w, "Ошибка при сериализации ответа", http.StatusInternalServerError)
		return
	}
}

// GetAircraft обрабатывает GET-запрос для получения одного самолёта по ID.
func (h *AircraftHandler) GetAircraft(w http.ResponseWriter, r *http.Request) {
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

	aircraft, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if aircraft == nil {
		http.Error(w, "Самолёт не найден", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aircraft)
}
