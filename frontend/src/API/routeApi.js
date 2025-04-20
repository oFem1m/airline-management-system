import apiClient from './apiClient'

export default {
    // Получить все маршруты
    getRoutes() {
        return apiClient.get('/routes')
    },
    // Создать маршрут
    createRoute(route) {
        return apiClient.post('/route', route)
    },
    // Обновить маршрут
    updateRoute(id, route) {
        return apiClient.put(`/route/${id}`, route)
    },
    // Удалить маршрут
    deleteRoute(id) {
        return apiClient.delete(`/route/${id}`)
    },
    // Получить маршруты, проходящие через конкретный аэропорт
    getRoutesByAirport(airportId) {
        return apiClient.get(`/routes/airport/${airportId}`)
    },
    // Получить конкретный маршрут по ID
    getRoute(id) {
        return apiClient.get(`/route/${id}`)
    },
}
