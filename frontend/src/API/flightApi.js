import apiClient from './apiClient'

export default {
    // Получение всех рейсов для конкретного маршрута по его ID
    getFlightsByRoute(routeId) {
        return apiClient.get(`/flights/route/${routeId}`)
    },
    // Получение рейса по ID
    getFlight(id) {
        return apiClient.get(`/flight/${id}`)
    },
    // Создание нового рейса
    createFlight(flight) {
        return apiClient.post('/flight', flight)
    },
    // Обновление рейса
    updateFlight(id, flight) {
        return apiClient.put(`/flight/${id}`, flight)
    },
    // Удаление рейса
    deleteFlight(id) {
        return apiClient.delete(`/flight/${id}`)
    },
}
