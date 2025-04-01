package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gorilla/mux"
)

type MaintenanceHandler struct {
	repo *repository.MaintenanceRepository
}

func NewMaintenanceHandler(repo *repository.MaintenanceRepository) *MaintenanceHandler {
	return &MaintenanceHandler{repo: repo}
}

// CreateMaintenance обрабатывает POST-запрос для создания нового обслуживания.
func (h *MaintenanceHandler) CreateMaintenance(w http.ResponseWriter, r *http.Request) {
	var m models.Maintenance
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}

// UpdateMaintenance обрабатывает PUT-запрос для обновления данных обслуживания по ID.
func (h *MaintenanceHandler) UpdateMaintenance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID обслуживания", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID обслуживания", http.StatusBadRequest)
		return
	}
	var m models.Maintenance
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	m.ID = id
	if err := h.repo.Update(&m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}

// GetMaintenance обрабатывает GET-запрос для получения обслуживания по ID.
func (h *MaintenanceHandler) GetMaintenance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID обслуживания", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID обслуживания", http.StatusBadRequest)
		return
	}
	m, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if m == nil {
		http.Error(w, "Обслуживание не найдено", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}

// GetAllMaintenance обрабатывает GET-запрос для получения списка всех обслуживаний.
func (h *MaintenanceHandler) GetAllMaintenance(w http.ResponseWriter, r *http.Request) {
	maintenances, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(maintenances)
}

// GetMaintenanceByAircraft обрабатывает GET-запрос для получения обслуживаний для конкретного самолёта.
func (h *MaintenanceHandler) GetMaintenanceByAircraft(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aircraftIDStr := vars["aircraftId"]
	if aircraftIDStr == "" {
		http.Error(w, "Не указан ID самолёта", http.StatusBadRequest)
		return
	}
	aircraftID, err := strconv.Atoi(aircraftIDStr)
	if err != nil {
		http.Error(w, "Некорректный ID самолёта", http.StatusBadRequest)
		return
	}
	maintenances, err := h.repo.GetByAircraft(aircraftID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(maintenances)
}

// GetMaintenanceByEmployee обрабатывает GET-запрос для получения обслуживаний, проведенных конкретным сотрудником.
func (h *MaintenanceHandler) GetMaintenanceByEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeIDStr := vars["employeeId"]
	if employeeIDStr == "" {
		http.Error(w, "Не указан ID сотрудника", http.StatusBadRequest)
		return
	}
	employeeID, err := strconv.Atoi(employeeIDStr)
	if err != nil {
		http.Error(w, "Некорректный ID сотрудника", http.StatusBadRequest)
		return
	}
	maintenances, err := h.repo.GetByEmployee(employeeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(maintenances)
}

// DeleteMaintenance обрабатывает DELETE-запрос для удаления обслуживания по ID.
func (h *MaintenanceHandler) DeleteMaintenance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID обслуживания", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID обслуживания", http.StatusBadRequest)
		return
	}
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
