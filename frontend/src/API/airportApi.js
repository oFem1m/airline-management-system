import apiClient from './apiClient'

export default {
    getAirports() {
        return apiClient.get('/airports')
    },
    createAirport(airport) {
        return apiClient.post('/airport', airport)
    },
    updateAirport(id, airport) {
        return apiClient.put(`/airport/${id}`, airport)
    },
    deleteAirport(id) {
        return apiClient.delete(`/airport/${id}`)
    },
}
