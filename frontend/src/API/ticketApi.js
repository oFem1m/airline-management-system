import apiClient from './apiClient'

export default {
    // Получение билетов для конкретного рейса
    getTicketsByFlight(flightId) {
        return apiClient.get(`/tickets/flight/${flightId}`)
    },

    // Создание нового билета
    addTicket(flightId, ticket) {
        return apiClient.post('/ticket', {
            flight_id: flightId,
            seat_number: ticket.seat_number,
            price: ticket.price,
        })
    },

    // Удаление билета
    deleteTicket(ticketId) {
        return apiClient.delete(`/ticket/${ticketId}`)
    },
}
