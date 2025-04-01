package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gorilla/mux"
)

// BookingHandler хранит методы для обработки HTTP-запросов к бронированиям.
type BookingHandler struct {
	repo *repository.BookingRepository
}

// NewBookingHandler возвращает новый обработчик для бронирований
func NewBookingHandler(repo *repository.BookingRepository) *BookingHandler {
	return &BookingHandler{repo: repo}
}

// CreateBooking обрабатывает POST-запрос для создания нового бронирования.
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

// UpdateBooking обрабатывает PUT-запрос для обновления данных бронирования по ID.
func (h *BookingHandler) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID бронирования", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID бронирования", http.StatusBadRequest)
		return
	}
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	booking.ID = id
	if err := h.repo.Update(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

// GetBooking обрабатывает GET-запрос для получения бронирования по ID.
func (h *BookingHandler) GetBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID бронирования", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID бронирования", http.StatusBadRequest)
		return
	}
	booking, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if booking == nil {
		http.Error(w, "Бронирование не найдено", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

// GetAllBookings обрабатывает GET-запрос для получения всех бронирований.
func (h *BookingHandler) GetAllBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

// GetBookingsByPassenger обрабатывает GET-запрос для получения бронирований конкретного пассажира.
func (h *BookingHandler) GetBookingsByPassenger(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	passengerIDStr := vars["passengerId"]
	if passengerIDStr == "" {
		http.Error(w, "Не указан ID пассажира", http.StatusBadRequest)
		return
	}
	passengerID, err := strconv.Atoi(passengerIDStr)
	if err != nil {
		http.Error(w, "Некорректный ID пассажира", http.StatusBadRequest)
		return
	}
	bookings, err := h.repo.GetByPassenger(passengerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

// DeleteBooking обрабатывает DELETE-запрос для удаления бронирования по ID.
func (h *BookingHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID бронирования", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID бронирования", http.StatusBadRequest)
		return
	}
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
