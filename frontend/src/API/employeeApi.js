import apiClient from './apiClient'

export default {
    getEmployees() {
        return apiClient.get('/employees')
    },
    createEmployee(employee) {
        return apiClient.post('/employee', employee)
    },
    updateEmployee(id, employee) {
        return apiClient.put(`/employee/${id}`, employee)
    },
    deleteEmployee(id) {
        return apiClient.delete(`/employee/${id}`)
    },
}
