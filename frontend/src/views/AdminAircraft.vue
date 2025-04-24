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
                <div v-for="m in maintenances" :key="m.id" class="col-md-4 mb-3">
                    <div class="card" style="cursor: pointer" @click="openEditMaintenanceModal(m)">
                        <div class="card-body">
                            <h5 class="card-title">Дата: {{ m.maintenance_date }}</h5>
                            <p class="card-text">
                                {{ m.description }}<br />
                                <strong>Следующее:</strong> {{ m.next_maintenance_date }}<br />
                                <strong>Выполнил:</strong> {{ getEmployeeName(m.performed_by) }}
                            </p>
                            <button
                                class="btn btn-danger btn-sm float-end"
                                @click.stop="deleteMaintenance(m.id)"
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
                    v-for="f in flights"
                    :key="f.id"
                    class="col-md-4 mb-3"
                    style="cursor: pointer"
                    @click="goToFlight(f.id)"
                >
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">Рейс {{ f.flight_number }}</h5>
                            <p class="card-text">
                                Вылет: {{ f.departure_time }}<br />
                                Прилет: {{ f.arrival_time }}<br />
                                Статус: {{ f.status }}
                            </p>
                            <button
                                class="btn btn-danger btn-sm float-end"
                                @click.stop="deleteFlight(f.id)"
                            >
                                ×
                            </button>
                        </div>
                    </div>
                </div>
            </div>
            <p v-else class="text-muted">Нет запланированных рейсов.</p>

            <MaintenanceModal
                ref="maintenanceModal"
                :initialMaintenance="selectedMaintenance"
                :aircraftId="aircraftId"
                :employees="employees"
                @createMaintenance="handleCreateMaintenance"
                @updateMaintenance="handleUpdateMaintenance"
            />
        </div>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Header from '@/components/Header.vue'
import MaintenanceModal from '@/components/MaintenanceModal.vue'
import aircraftApi from '@/API/aircraftApi'
import maintenanceApi from '@/API/maintenanceApi'
import employeeApi from '@/API/employeeApi'
import flightApi from '@/API/flightApi'

export default {
    name: 'AdminAircraft',
    components: { Header, MaintenanceModal },
    setup() {
        const route = useRoute()
        const router = useRouter()
        const aircraftId = Number(route.params.id)

        const aircraft = ref(null)
        const maintenances = ref([])
        const employees = ref([])

        const selectedMaintenance = ref(null)

        const flights = ref([])

        const fetchAircraft = () =>
            aircraftApi.getAircraft(aircraftId).then((r) => (aircraft.value = r.data))

        const fetchMaintenances = () =>
            maintenanceApi
                .getMaintenancesByAircraft(aircraftId)
                .then((r) => (maintenances.value = r.data))

        const fetchEmployees = () =>
            employeeApi.getEmployees().then((r) => (employees.value = r.data))

        const fetchFlights = () =>
            flightApi.getFlightsByAircraft(aircraftId).then((r) => (flights.value = r.data))

        const deleteMaintenance = (id) =>
            maintenanceApi.deleteMaintenance(id).then(fetchMaintenances)

        const openCreateMaintenanceModal = () => {
            selectedMaintenance.value = null
            maintenanceModal.value.open()
        }
        const openEditMaintenanceModal = (m) => {
            selectedMaintenance.value = { ...m }
            maintenanceModal.value.open()
        }
        const handleCreateMaintenance = () => fetchMaintenances()
        const handleUpdateMaintenance = () => fetchMaintenances()

        const deleteFlight = (id) => flightApi.deleteFlight(id).then(fetchFlights)

        const goToFlight = (id) => router.push({ name: 'AdminFlight', params: { id } })

        const getEmployeeName = (eid) => {
            const e = employees.value.find((x) => x.id === eid)
            return e ? `${e.first_name} ${e.last_name}` : '—'
        }

        const maintenanceModal = ref(null)

        onMounted(() => {
            Promise.all([fetchAircraft(), fetchMaintenances(), fetchEmployees(), fetchFlights()])
        })

        return {
            aircraft,
            maintenances,
            employees,
            selectedMaintenance,
            flights,
            deleteMaintenance,
            openCreateMaintenanceModal,
            openEditMaintenanceModal,
            handleCreateMaintenance,
            handleUpdateMaintenance,
            deleteFlight,
            goToFlight,
            getEmployeeName,
            aircraftId,
            maintenanceModal,
        }
    },
}
</script>

<style scoped></style>
