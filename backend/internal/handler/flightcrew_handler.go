package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gorilla/mux"
)

// FlightCrewHandler хранит методы для работы с экипажами (FlightCrew).
type FlightCrewHandler struct {
	repo *repository.FlightCrewRepository
}

// NewFlightCrewHandler возвращает новый обработчик для экипажей.
func NewFlightCrewHandler(repo *repository.FlightCrewRepository) *FlightCrewHandler {
	return &FlightCrewHandler{repo: repo}
}

// AssignEmployee назначает сотрудника на рейс.
func (h *FlightCrewHandler) AssignEmployee(w http.ResponseWriter, r *http.Request) {
	var fc models.FlightCrew
	if err := json.NewDecoder(r.Body).Decode(&fc); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	if err := h.repo.AssignEmployeeToFlight(&fc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fc)
}

// RemoveEmployee удаляет сотрудника с рейса.
// Используется удаление по query-параметрам: flight_id и employee_id.
func (h *FlightCrewHandler) RemoveEmployee(w http.ResponseWriter, r *http.Request) {
	flightIDStr := r.URL.Query().Get("flight_id")
	employeeIDStr := r.URL.Query().Get("employee_id")
	if flightIDStr == "" || employeeIDStr == "" {
		http.Error(w, "Необходимо указать flight_id и employee_id", http.StatusBadRequest)
		return
	}
	flightID, err := strconv.Atoi(flightIDStr)
	if err != nil {
		http.Error(w, "Некорректный flight_id", http.StatusBadRequest)
		return
	}
	employeeID, err := strconv.Atoi(employeeIDStr)
	if err != nil {
		http.Error(w, "Некорректный employee_id", http.StatusBadRequest)
		return
	}
	if err := h.repo.RemoveEmployeeFromFlight(flightID, employeeID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetEmployeesByFlight возвращает список сотрудников, назначенных на рейс.
func (h *FlightCrewHandler) GetEmployeesByFlight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	flightIDStr := vars["flight_id"]
	flightID, err := strconv.Atoi(flightIDStr)
	if err != nil {
		http.Error(w, "Некорректный flight_id", http.StatusBadRequest)
		return
	}
	employees, err := h.repo.GetEmployeesByFlight(flightID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

// GetFlightsByEmployee возвращает список рейсов, на которых закреплён сотрудник.
func (h *FlightCrewHandler) GetFlightsByEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeIDStr := vars["employee_id"]
	employeeID, err := strconv.Atoi(employeeIDStr)
	if err != nil {
		http.Error(w, "Некорректный employee_id", http.StatusBadRequest)
		return
	}
	flights, err := h.repo.GetFlightsByEmployee(employeeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}
