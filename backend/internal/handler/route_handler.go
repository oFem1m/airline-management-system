package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gorilla/mux"
)

// RouteHandler хранит методы для обработки HTTP-запросов к маршрутам.
type RouteHandler struct {
	repo *repository.RouteRepository
}

// NewRouteHandler возвращает новый обработчик для маршрутов.
func NewRouteHandler(repo *repository.RouteRepository) *RouteHandler {
	return &RouteHandler{repo: repo}
}

// CreateRoute обрабатывает POST-запрос для создания маршрута.
func (h *RouteHandler) CreateRoute(w http.ResponseWriter, r *http.Request) {
	var route models.Route
	if err := json.NewDecoder(r.Body).Decode(&route); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&route); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(route)
}

// DeleteRoute обрабатывает DELETE-запрос для удаления маршрута по ID.
func (h *RouteHandler) DeleteRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID маршрута", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID маршрута", http.StatusBadRequest)
		return
	}
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetAllRoutes обрабатывает GET-запрос для получения списка всех маршрутов.
func (h *RouteHandler) GetAllRoutes(w http.ResponseWriter, r *http.Request) {
	routes, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routes)
}

// GetRoute обрабатывает GET-запрос для получения маршрута по его ID.
func (h *RouteHandler) GetRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID маршрута", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID маршрута", http.StatusBadRequest)
		return
	}
	route, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if route == nil {
		http.Error(w, "Маршрут не найден", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(route)
}

// GetRoutesByAirport обрабатывает GET-запрос для получения маршрутов, связанных с конкретным аэропортом.
func (h *RouteHandler) GetRoutesByAirport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airportIDStr := vars["airportId"]
	if airportIDStr == "" {
		http.Error(w, "Не указан ID аэропорта", http.StatusBadRequest)
		return
	}
	airportID, err := strconv.Atoi(airportIDStr)
	if err != nil {
		http.Error(w, "Некорректный ID аэропорта", http.StatusBadRequest)
		return
	}
	routes, err := h.repo.GetRoutesByAirport(airportID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routes)
}
