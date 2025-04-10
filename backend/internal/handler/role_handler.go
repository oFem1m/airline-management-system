package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gorilla/mux"
)

// RoleHandler хранит методы для обработки HTTP-запросов к ролям.
type RoleHandler struct {
	repo *repository.RoleRepository
}

// NewRoleHandler возвращает новый обработчик для ролей.
func NewRoleHandler(repo *repository.RoleRepository) *RoleHandler {
	return &RoleHandler{repo: repo}
}

// CreateRole создает новую роль.
func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&role); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

// UpdateRole обновляет данные роли по ID.
func (h *RoleHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID роли", http.StatusBadRequest)
		return
	}
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	role.ID = id
	if err := h.repo.Update(&role); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

// GetAllRoles возвращает список всех ролей.
func (h *RoleHandler) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}

// GetRole возвращает роль по ID.
func (h *RoleHandler) GetRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID роли", http.StatusBadRequest)
		return
	}
	role, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if role == nil {
		http.Error(w, "Роль не найдена", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

// DeleteRole удаляет роль по ID.
func (h *RoleHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID роли", http.StatusBadRequest)
		return
	}
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
