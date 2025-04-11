<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">
                        {{ isEditMode ? 'Редактировать сотрудника' : 'Добавить сотрудника' }}
                    </h5>
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
                            <label for="first_name" class="form-label">Имя</label>
                            <input
                                type="text"
                                id="first_name"
                                v-model="employee.first_name"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="last_name" class="form-label">Фамилия</label>
                            <input
                                type="text"
                                id="last_name"
                                v-model="employee.last_name"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="role_id" class="form-label">ID роли</label>
                            <input
                                type="number"
                                id="role_id"
                                v-model="employee.role_id"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="hire_date" class="form-label">Дата найма</label>
                            <input
                                type="datetime-local"
                                id="hire_date"
                                v-model="employee.hire_date"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="salary" class="form-label">Зарплата</label>
                            <input
                                type="number"
                                step="0.01"
                                id="salary"
                                v-model="employee.salary"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="email" class="form-label">Email</label>
                            <input
                                type="email"
                                id="email"
                                v-model="employee.email"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="phone" class="form-label">Телефон</label>
                            <input
                                type="text"
                                id="phone"
                                v-model="employee.phone"
                                class="form-control"
                                required
                            />
                        </div>
                        <button type="submit" class="btn btn-primary">
                            {{ isEditMode ? 'Изменить' : 'Добавить' }}
                        </button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { ref, computed, defineComponent, watch } from 'vue'
import { Modal } from 'bootstrap'

function fromISO(isoStr) {
    if (!isoStr) return ''
    const date = new Date(isoStr)
    const year = date.getFullYear().toString().padStart(4, '0')
    const month = (date.getMonth() + 1).toString().padStart(2, '0')
    const day = date.getDate().toString().padStart(2, '0')
    const hours = date.getHours().toString().padStart(2, '0')
    const minutes = date.getMinutes().toString().padStart(2, '0')
    return `${year}-${month}-${day}T${hours}:${minutes}`
}

function toISO(localDateTimeStr) {
    if (!localDateTimeStr) return null
    const date = new Date(localDateTimeStr)
    return date.toISOString()
}

export default defineComponent({
    name: 'EmployeeModal',
    props: {
        initialEmployee: {
            type: Object,
            default: null,
        },
    },
    emits: ['createEmployee', 'updateEmployee'],
    setup(props, { emit }) {
        const employee = ref({
            id: null,
            first_name: '',
            last_name: '',
            role_id: '',
            hire_date: '',
            salary: '',
            email: '',
            phone: '',
        })

        const resetForm = () => {
            employee.value = {
                id: null,
                first_name: '',
                last_name: '',
                role_id: '',
                hire_date: '',
                salary: '',
                email: '',
                phone: '',
            }
        }

        const isEditMode = computed(() => employee.value.id !== null)

        watch(
            () => props.initialEmployee,
            (newVal) => {
                if (newVal) {
                    employee.value = {
                        ...newVal,
                        hire_date: fromISO(newVal.hire_date),
                    }
                } else {
                    resetForm()
                }
            },
            { immediate: true },
        )

        const modalElement = ref(null)
        let modalInstance = null

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
            const preparedEmployee = {
                ...employee.value,
                hire_date: toISO(employee.value.hire_date),
            }
            if (isEditMode.value) {
                emit('updateEmployee', preparedEmployee)
            } else {
                emit('createEmployee', preparedEmployee)
            }
            resetForm()
            close()
        }

        return {
            modalElement,
            employee,
            isEditMode,
            open,
            close,
            submitForm,
        }
    },
})
</script>

<style scoped></style>
