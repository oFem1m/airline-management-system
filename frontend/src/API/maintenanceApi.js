import apiClient from './apiClient'

export default {
    getMaintenancesByAircraft(aircraftId) {
        return apiClient.get(`/maintenances/aircraft/${aircraftId}`)
    },
    createMaintenance(maintenance) {
        return apiClient.post('/maintenance', maintenance)
    },
    updateMaintenance(id, maintenance) {
        return apiClient.put(`/maintenance/${id}`, maintenance)
    },
    deleteMaintenance(id) {
        return apiClient.delete(`/maintenance/${id}`)
    },
}
