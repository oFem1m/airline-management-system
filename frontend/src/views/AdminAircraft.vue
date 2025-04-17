<template>
    <div>
        <Header />

        <div class="container mt-4">
            <router-link to="/admin" class="back-link">
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
import { useRoute } from 'vue-router'
import Header from '@/components/Header.vue'
import aircraftApi from '@/API/aircraftApi'
import maintenanceApi from '@/API/maintenanceApi'
import employeeApi from '@/API/employeeApi'
import MaintenanceModal from '@/components/MaintenanceModal.vue'

export default {
    name: 'AdminAircraft',
    components: {
        // eslint-disable-next-line vue/no-reserved-component-names
        Header,
        MaintenanceModal,
    },
    setup() {
        const route = useRoute()
        const aircraftId = parseInt(route.params.id, 10)

        const aircraft = ref(null)
        const maintenances = ref([])
        const employees = ref([])

        const selectedMaintenance = ref(null)

        const fetchAircraft = () => {
            aircraftApi
                .getAircraft(aircraftId)
                .then((response) => {
                    aircraft.value = response.data
                })
                .catch((error) => {
                    console.error('Ошибка при получении самолёта', error)
                })
        }

        const fetchMaintenances = () => {
            maintenanceApi
                .getMaintenancesByAircraft(aircraftId)
                .then((response) => {
                    maintenances.value = response.data
                })
                .catch((error) => {
                    console.error('Ошибка при получении обслуживаний', error)
                })
        }

        const fetchEmployees = () => {
            employeeApi
                .getEmployees()
                .then((response) => {
                    employees.value = response.data
                })
                .catch((error) => {
                    console.error('Ошибка получения сотрудников', error)
                })
        }

        const deleteMaintenance = (maintenanceId) => {
            maintenanceApi
                .deleteMaintenance(maintenanceId)
                .then(() => {
                    fetchMaintenances()
                })
                .catch((error) => {
                    console.error('Ошибка при удалении обслуживания', error)
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
                    console.error('Ошибка при создании обслуживания', error)
                })
        }

        const handleUpdateMaintenance = (maintenanceData) => {
            maintenanceApi
                .updateMaintenance(maintenanceData.id, maintenanceData)
                .then(() => {
                    fetchMaintenances()
                })
                .catch((error) => {
                    console.error('Ошибка при обновлении обслуживания', error)
                })
        }

        const getEmployeeName = (employeeId) => {
            const emp = employees.value.find((e) => e.id === employeeId)
            return emp ? `${emp.first_name} ${emp.last_name} (ID: ${emp.id})` : 'Не задан'
        }

        const maintenanceModal = ref(null)

        onMounted(() => {
            fetchAircraft()
            fetchMaintenances()
            fetchEmployees()
        })

        return {
            aircraft,
            maintenances,
            deleteMaintenance,
            openCreateMaintenanceModal,
            openEditMaintenanceModal,
            handleCreateMaintenance,
            handleUpdateMaintenance,
            selectedMaintenance,
            maintenanceModal,
            aircraftId,
            employees,
            getEmployeeName,
        }
    },
}
</script>

<style scoped></style>
