import apiClient from './apiClient'

export default {
    // Получение списка всех бронирований
    getBookings() {
        return apiClient.get('/bookings')
    },
    // Создание нового бронирования
    createBooking(booking) {
        return apiClient.post('/booking', {
            passenger_id: booking.passenger_id,
            booking_date: booking.booking_date,
            status: booking.status,
        })
    },
    // Получение бронирования по ID
    getBooking(id) {
        return apiClient.get(`/booking/${id}`)
    },
    // Обновление бронирования по ID
    updateBooking(id, booking) {
        return apiClient.put(`/booking/${id}`, {
            passenger_id: booking.passenger_id,
            booking_date: booking.booking_date,
            status: booking.status,
        })
    },
    // Удаление бронирования по ID
    deleteBooking(id) {
        return apiClient.delete(`/booking/${id}`)
    },
}
