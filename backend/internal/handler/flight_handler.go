package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gorilla/mux"
)

// FlightHandler хранит методы для обработки HTTP-запросов к рейсам.
type FlightHandler struct {
	repo *repository.FlightRepository
}

func NewFlightHandler(repo *repository.FlightRepository) *FlightHandler {
	return &FlightHandler{repo: repo}
}

// CreateFlight обрабатывает POST-запрос для создания нового рейса.
func (h *FlightHandler) CreateFlight(w http.ResponseWriter, r *http.Request) {
	var flight models.Flight
	if err := json.NewDecoder(r.Body).Decode(&flight); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&flight); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flight)
}

// DeleteFlight обрабатывает DELETE-запрос для удаления рейса по ID.
func (h *FlightHandler) DeleteFlight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID рейса", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID рейса", http.StatusBadRequest)
		return
	}
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetAllFlights обрабатывает GET-запрос для получения всех рейсов.
func (h *FlightHandler) GetAllFlights(w http.ResponseWriter, r *http.Request) {
	flights, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}

// GetFlight обрабатывает GET-запрос для получения рейса по ID.
func (h *FlightHandler) GetFlight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID рейса", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID рейса", http.StatusBadRequest)
		return
	}
	flight, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if flight == nil {
		http.Error(w, "Рейс не найден", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flight)
}

// GetFlightsByRoute обрабатывает GET-запрос для получения всех рейсов по route_id.
func (h *FlightHandler) GetFlightsByRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	routeIdStr := vars["routeId"]
	if routeIdStr == "" {
		http.Error(w, "Не указан ID маршрута", http.StatusBadRequest)
		return
	}
	routeId, err := strconv.Atoi(routeIdStr)
	if err != nil {
		http.Error(w, "Некорректный ID маршрута", http.StatusBadRequest)
		return
	}
	flights, err := h.repo.GetByRoute(routeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}

// GetFlightsByAircraft обрабатывает GET-запрос для получения всех рейсов для конкретного самолёта.
func (h *FlightHandler) GetFlightsByAircraft(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aircraftIdStr := vars["aircraftId"]
	if aircraftIdStr == "" {
		http.Error(w, "Не указан ID самолёта", http.StatusBadRequest)
		return
	}
	aircraftId, err := strconv.Atoi(aircraftIdStr)
	if err != nil {
		http.Error(w, "Некорректный ID самолёта", http.StatusBadRequest)
		return
	}
	flights, err := h.repo.GetByAircraft(aircraftId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}

// GetFlightsByAirport обрабатывает GET-запрос для получения всех рейсов для конкретного аэропорта.
func (h *FlightHandler) GetFlightsByAirport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airportIdStr := vars["airportId"]
	if airportIdStr == "" {
		http.Error(w, "Не указан ID аэропорта", http.StatusBadRequest)
		return
	}
	airportId, err := strconv.Atoi(airportIdStr)
	if err != nil {
		http.Error(w, "Некорректный ID аэропорта", http.StatusBadRequest)
		return
	}
	flights, err := h.repo.GetByAirport(airportId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}
