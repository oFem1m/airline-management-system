import apiClient from './apiClient'

export default {
    // Получение билетов для конкретного рейса
    getTicketsByFlight(flightId) {
        return apiClient.get(`/tickets/flight/${flightId}`)
    },

    // Получение билетов для конкретного пассажира
    getTicketsByPassenger(passengerId) {
        return apiClient.get(`/tickets/passenger/${passengerId}`)
    },

    // Получение билетов для конкретного бронирования
    getTicketsByBooking(bookingId) {
        return apiClient.get(`/tickets/booking/${bookingId}`)
    },

    // Создание нового билета (включая необязательный booking_id)
    addTicket(flightId, ticket) {
        return apiClient.post('/ticket', {
            flight_id: flightId,
            passenger_id: ticket.passenger_id,
            booking_id: ticket.booking_id || null,
            seat_number: ticket.seat_number,
            price: ticket.price
        })
    },

    // Создание нового билета в контексте бронирования
    addTicketForBooking(bookingId, ticket) {
        return apiClient.post('/ticket', {
            flight_id: ticket.flight_id,
            passenger_id: ticket.passenger_id,
            booking_id: bookingId,
            seat_number: ticket.seat_number,
            price: ticket.price
        })
    },

    // Обновление существующего билета (включая необязательный booking_id)
    updateTicket(id, ticket) {
        return apiClient.put(`/ticket/${id}`, {
            flight_id: ticket.flight_id,
            passenger_id: ticket.passenger_id,
            booking_id: ticket.booking_id || null,
            seat_number: ticket.seat_number,
            price: ticket.price
        })
    },

    // Обновление билета в контексте бронирования
    updateTicketForBooking(ticketId, ticket) {
        return apiClient.put(`/ticket/${ticketId}`, {
            flight_id: ticket.flight_id,
            passenger_id: ticket.passenger_id,
            booking_id: ticket.booking_id,
            seat_number: ticket.seat_number,
            price: ticket.price
        })
    },

    // Удаление билета
    deleteTicket(ticketId) {
        return apiClient.delete(`/ticket/${ticketId}`)
    }
}
