<template>
    <div>
        <Header />

        <div class="container mt-4">
            <router-link to="/admin/aircrafts" class="back-link">
                <span class="back-arrow">&#8592;</span> Назад
            </router-link>
            <h1>Информация о самолёте</h1>
            <div v-if="aircraft" class="mt-3">
                <p><strong>Бортовой номер:</strong> {{ aircraft.tail_number }}</p>
                <p><strong>Модель:</strong> {{ aircraft.model }}</p>
                <p><strong>Вместимость:</strong> {{ aircraft.capacity }}</p>
                <p><strong>Год выпуска:</strong> {{ aircraft.manufacture_year }}</p>
            </div>

            <hr class="my-4" />

            <div class="d-flex justify-content-between align-items-center">
                <h2>Обслуживания</h2>
                <button class="btn btn-primary" @click="openCreateMaintenanceModal">
                    Добавить обслуживание
                </button>
            </div>

            <div class="row mt-3">
                <div
                    v-for="maintenance in maintenances"
                    :key="maintenance.id"
                    class="col-md-4 mb-3"
                >
                    <div
                        class="card"
                        style="cursor: pointer"
                        @click="openEditMaintenanceModal(maintenance)"
                    >
                        <div class="card-body">
                            <h5 class="card-title">
                                Дата обслуживания: {{ maintenance.maintenance_date }}
                            </h5>
                            <p class="card-text">
                                {{ maintenance.description }}<br />
                                <strong>Следующее обслуживание:</strong>
                                {{ maintenance.next_maintenance_date }}<br />
                                <strong>Выполнил:</strong>
                                {{ getEmployeeName(maintenance.performed_by) }}
                            </p>
                            <button
                                class="btn btn-danger btn-sm float-end"
                                @click.stop="deleteMaintenance(maintenance.id)"
                            >
                                ×
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            <hr class="my-4" />

            <div class="d-flex justify-content-between align-items-center">
                <h2>Рейсы самолёта</h2>
            </div>
            <div v-if="flights.length" class="row mt-3">
                <div
                    v-for="flight in flights"
                    :key="flight.id"
                    class="col-md-4 mb-3"
                    style="cursor: pointer"
                    @click="goToFlight(flight.id)"
                >
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">Рейс {{ flight.flight_number }}</h5>
                            <p class="card-text">
                                Вылет: {{ flight.departure_time }}<br />
                                Прилет: {{ flight.arrival_time }}<br />
                                Статус: {{ flight.status }}
                            </p>
                            <button
                                class="btn btn-danger btn-sm float-end"
                                @click.stop="deleteFlight(flight.id)"
                            >
                                ×
                            </button>
                        </div>
                    </div>
                </div>
            </div>
            <p v-else class="text-muted">Нет запланированных рейсов.</p>
        </div>

        <MaintenanceModal
            ref="maintenanceModal"
            :initialMaintenance="selectedMaintenance"
            :aircraftId="aircraftId"
            :employees="employees"
            @createMaintenance="handleCreateMaintenance"
            @updateMaintenance="handleUpdateMaintenance"
        />
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Header from '@/components/Header.vue'
import aircraftApi from '@/API/aircraftApi'
import maintenanceApi from '@/API/maintenanceApi'
import employeeApi from '@/API/employeeApi'
import flightApi from '@/API/flightApi'
import MaintenanceModal from '@/components/MaintenanceModal.vue'

export default {
    name: 'AdminAircraft',
    components: {
        Header,
        MaintenanceModal,
    },
    setup() {
        const route = useRoute()
        const router = useRouter()
        const aircraftId = parseInt(route.params.id, 10)

        const aircraft = ref(null)
        const maintenances = ref([])
        const employees = ref([])
        const flights = ref([])
        const selectedMaintenance = ref(null)

        const fetchAircraft = () => {
            aircraftApi
                .getAircraft(aircraftId)
                .then((response) => {
                    aircraft.value = response.data
                })
                .catch((error) => {
                    alert(error.response.data)
                })
        }

        const fetchMaintenances = () => {
            maintenanceApi
                .getMaintenancesByAircraft(aircraftId)
                .then((response) => {
                    maintenances.value = response.data
                })
                .catch((error) => {
                    alert(error.response.data)
                })
        }

        const fetchEmployees = () => {
            employeeApi
                .getEmployees()
                .then((response) => {
                    employees.value = response.data
                })
                .catch((error) => {
                    alert(error.response.data)
                })
        }

        const fetchFlights = () => {
            flightApi
                .getFlightsByAircraft(aircraftId)
                .then(response => {
                    if(response.data) {
                        flights.value = response.data
                    }
                })
                .catch(error => {
                    alert(error.response.data)
                })
        }

        const deleteMaintenance = maintenanceId => {
            maintenanceApi
                .deleteMaintenance(maintenanceId)
                .then(() => {
                    fetchMaintenances()
                })
                .catch(error => {
                    alert(error.response.data)
                })
        }

        const openCreateMaintenanceModal = () => {
            selectedMaintenance.value = null
            maintenanceModal.value.open()
        }

        const openEditMaintenanceModal = (maintenance) => {
            selectedMaintenance.value = { ...maintenance }
            maintenanceModal.value.open()
        }

        const handleCreateMaintenance = (newMaintenance) => {
            maintenanceApi
                .createMaintenance(newMaintenance)
                .then(() => {
                    fetchMaintenances()
                })
                .catch((error) => {
                    alert(error.response.data)
                })
        }

        const handleUpdateMaintenance = (maintenanceData) => {
            maintenanceApi
                .updateMaintenance(maintenanceData.id, maintenanceData)
                .then(() => {
                    fetchMaintenances()
                })
                .catch((error) => {
                    alert(error.response.data)
                })
        }

        const deleteFlight = (flightId) => {
            flightApi
                .deleteFlight(flightId)
                .then(() => {
                    fetchFlights()
                })
                .catch((error) => {
                    alert(error.response.data)
                })
        }

        const goToFlight = (flightId) => {
            router.push({ name: 'AdminFlight', params: { id: flightId } })
        }

        const getEmployeeName = (employeeId) => {
            const emp = employees.value.find(e => e.id === employeeId)
            return emp ? `${emp.first_name} ${emp.last_name} (ID: ${emp.id})` : 'Не задан'
        }

        const maintenanceModal = ref(null)

        onMounted(() => {
            fetchAircraft()
            fetchMaintenances()
            fetchEmployees()
            fetchFlights()
        })

        return {
            aircraft,
            maintenances,
            employees,
            flights,
            selectedMaintenance,
            maintenanceModal,
            aircraftId,
            deleteMaintenance,
            openCreateMaintenanceModal,
            openEditMaintenanceModal,
            handleCreateMaintenance,
            handleUpdateMaintenance,
            deleteFlight,
            goToFlight,
            getEmployeeName,
        }
    },
}
</script>

<style scoped></style>
