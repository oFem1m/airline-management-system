import apiClient from './apiClient'

export default {
    getRoutes() {
        return apiClient.get('/routes')
    },
    createRoute(route) {
        return apiClient.post('/route', route)
    },
    updateRoute(id, route) {
        return apiClient.put(`/route/${id}`, route)
    },
    deleteRoute(id) {
        return apiClient.delete(`/route/${id}`)
    },
}
