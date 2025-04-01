package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gorilla/mux"
)

// AirportHandler хранит методы для обработки HTTP-запросов к аэропортам.
type AirportHandler struct {
	repo *repository.AirportRepository
}

// NewAirportHandler возвращает новый обработчик для аэропортов
func NewAirportHandler(repo *repository.AirportRepository) *AirportHandler {
	return &AirportHandler{repo: repo}
}

// CreateAirport обрабатывает POST-запрос на создание нового аэропорта
func (h *AirportHandler) CreateAirport(w http.ResponseWriter, r *http.Request) {
	var airport models.Airport

	if err := json.NewDecoder(r.Body).Decode(&airport); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&airport); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(airport)
	if err != nil {
		return
	}
}

// DeleteAirport обрабатывает POST-запрос на удаление аэропорта по ID
func (h *AirportHandler) DeleteAirport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID аэропорта", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID аэропорта", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAllAirports обрабатывает GET-запрос и возвращает список всех аэропортов
func (h *AirportHandler) GetAllAirports(w http.ResponseWriter, r *http.Request) {
	airports, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(airports); err != nil {
		http.Error(w, "Ошибка при сериализации ответа", http.StatusInternalServerError)
		return
	}
}

// GetAirport обрабатывает GET-запрос для получения одного аэропота по ID.
func (h *AirportHandler) GetAirport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID аэропорта", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID аэропорта", http.StatusBadRequest)
		return
	}

	airport, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if airport == nil {
		http.Error(w, "Аэропорт не найден", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(airport)
}
