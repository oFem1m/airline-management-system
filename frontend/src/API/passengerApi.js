import apiClient from './apiClient'

export default {
    // Получение списка пассажиров
    getPassengers() {
        return apiClient.get('/passengers')
    },

    // Получение пассажира по ID
    getPassenger(id) {
        return apiClient.get(`/passenger/${id}`)
    },

    // Добавление нового пассажира
    addPassenger(passenger) {
        return apiClient.post('/passenger', {
            first_name: passenger.first_name,
            last_name: passenger.last_name,
            email: passenger.email,
            phone: passenger.phone,
            passport_number: passenger.passport_number,
        })
    },

    // Обновление информации о пассажире
    updatePassenger(id, passenger) {
        return apiClient.put(`/passenger/${id}`, {
            first_name: passenger.first_name,
            last_name: passenger.last_name,
            email: passenger.email,
            phone: passenger.phone,
            passport_number: passenger.passport_number,
        })
    },

    // Удаление пассажира
    deletePassenger(id) {
        return apiClient.delete(`/passenger/${id}`)
    },
}
