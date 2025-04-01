package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gorilla/mux"
)

// TicketHandler хранит методы для обработки HTTP-запросов к билетам.
type TicketHandler struct {
	repo *repository.TicketRepository
}

// NewTicketHandler возвращает новый обработчик для билетов
func NewTicketHandler(repo *repository.TicketRepository) *TicketHandler {
	return &TicketHandler{repo: repo}
}

// CreateTicket обрабатывает POST-запрос для создания нового билета.
func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&ticket); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ticket)
}

// UpdateTicket обрабатывает PUT-запрос для обновления информации о билете по его ID.
func (h *TicketHandler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID билета", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID билета", http.StatusBadRequest)
		return
	}
	var ticket models.Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	ticket.ID = id
	if err := h.repo.Update(&ticket); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ticket)
}

// GetTicket обрабатывает GET-запрос для получения билета по его ID.
func (h *TicketHandler) GetTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID билета", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID билета", http.StatusBadRequest)
		return
	}
	ticket, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if ticket == nil {
		http.Error(w, "Билет не найден", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ticket)
}

// GetAllTickets обрабатывает GET-запрос для получения списка всех билетов.
func (h *TicketHandler) GetAllTickets(w http.ResponseWriter, r *http.Request) {
	tickets, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

// GetTicketsByPassenger обрабатывает GET-запрос для получения билетов для конкретного пассажира.
func (h *TicketHandler) GetTicketsByPassenger(w http.ResponseWriter, r *http.Request) {
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
	tickets, err := h.repo.GetByPassenger(passengerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

// GetTicketsByFlight обрабатывает GET-запрос для получения билетов на рейс по его ID.
func (h *TicketHandler) GetTicketsByFlight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	flightIDStr := vars["flightId"]
	if flightIDStr == "" {
		http.Error(w, "Не указан ID рейса", http.StatusBadRequest)
		return
	}
	flightID, err := strconv.Atoi(flightIDStr)
	if err != nil {
		http.Error(w, "Некорректный ID рейса", http.StatusBadRequest)
		return
	}
	tickets, err := h.repo.GetByFlight(flightID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

// GetTicketsByBooking обрабатывает GET-запрос для получения билетов в бронировании по его ID.
func (h *TicketHandler) GetTicketsByBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingIDStr := vars["bookingId"]
	if bookingIDStr == "" {
		http.Error(w, "Не указан ID бронирования", http.StatusBadRequest)
		return
	}
	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		http.Error(w, "Некорректный ID бронирования", http.StatusBadRequest)
		return
	}
	tickets, err := h.repo.GetByBooking(bookingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

// DeleteTicket обрабатывает DELETE-запрос для удаления билета по его ID.
func (h *TicketHandler) DeleteTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не указан ID билета", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID билета", http.StatusBadRequest)
		return
	}
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
