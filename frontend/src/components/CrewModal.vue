<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">Назначить сотрудника</h5>
                    <button
                        type="button"
                        class="btn-close"
                        aria-label="Close"
                        @click="close"
                    ></button>
                </div>
                <div class="modal-body">
                    <form @submit.prevent="submitForm">
                        <div class="mb-3">
                            <label for="employeeId" class="form-label">Выберите сотрудника</label>
                            <select
                                id="employeeId"
                                v-model="selectedEmployee"
                                class="form-control"
                                required
                            >
                                <option v-for="emp in employees" :key="emp.id" :value="emp.id">
                                    {{ emp.first_name }} {{ emp.last_name }}
                                </option>
                            </select>
                        </div>
                        <button type="submit" class="btn btn-primary">Назначить</button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { Modal } from 'bootstrap'
import employeeApi from '@/API/employeeApi'

export default {
    name: 'CrewModal',
    props: {
        flightId: {
            type: Number,
            required: true,
        },
    },
    emits: ['addCrewMember'],
    setup(props, { emit }) {
        const employees = ref([])
        const selectedEmployee = ref(null)
        const modalElement = ref(null)
        let modalInstance = null

        const fetchEmployees = () => {
            employeeApi
                .getEmployees()
                .then((response) => {
                    employees.value = response.data
                })
                .catch(console.error)
        }

        const open = () => {
            if (!modalInstance) {
                modalInstance = new Modal(modalElement.value)
            }
            modalInstance.show()
        }

        const close = () => {
            if (modalInstance) {
                modalInstance.hide()
            }
        }

        const submitForm = () => {
            emit('addCrewMember', { flightId: props.flightId, employeeId: selectedEmployee.value })
            close()
        }

        onMounted(() => {
            fetchEmployees()
        })

        return {
            employees,
            selectedEmployee,
            open,
            close,
            submitForm,
        }
    },
}
</script>

<style scoped></style>
