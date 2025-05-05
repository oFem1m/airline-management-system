import apiClient from './apiClient'

export default {
    getRoles() {
        return apiClient.get('/roles')
    },
    createRole(role) {
        return apiClient.post('/role', role)
    },
    updateRole(id, role) {
        return apiClient.put(`/role/${id}`, role)
    },
    deleteRole(id) {
        return apiClient.delete(`/role/${id}`)
    },
}
