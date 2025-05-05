import apiClient from './apiClient'

export default {
    // Получение списка сотрудников, назначенных на рейс
    getCrewByFlight(flightId) {
        return apiClient.get(`/flightcrew/flight/${flightId}`)
    },

    // Назначение сотрудника на рейс
    addCrewMember(flightId, crewMember) {
        return apiClient.post('/flightcrew', {
            flight_id: flightId,
            employee_id: crewMember.employeeId,
        })
    },

    // Удаление сотрудника с рейса
    removeCrewMember(flightId, employeeId) {
        return apiClient.delete(`/flightcrew`, {
            params: {
                flight_id: flightId,
                employee_id: employeeId,
            },
        })
    },
}
