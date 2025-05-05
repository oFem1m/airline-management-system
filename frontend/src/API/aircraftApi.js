import apiClient from './apiClient'

export default {
    getAircrafts() {
        return apiClient.get('/aircrafts')
    },
    getAircraft(id) {
        return apiClient.get(`/aircraft/${id}`)
    },
    createAircraft(aircraft) {
        return apiClient.post('/aircraft', aircraft)
    },
    deleteAircraft(id) {
        return apiClient.delete(`/aircraft/${id}`)
    },
}
