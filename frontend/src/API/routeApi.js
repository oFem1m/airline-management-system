import apiClient from './apiClient'

export default {
    getRoutes() {
        return apiClient.get('/routes')
    },
    createRoute(route) {
        return apiClient.post('/route', route)
    },
    deleteRoute(id) {
        return apiClient.delete(`/route/${id}`)
    },
}
