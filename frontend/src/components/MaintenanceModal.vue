<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">
                        {{ isEditMode ? 'Редактировать обслуживание' : 'Создать обслуживание' }}
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
                            <label for="maintenance_date" class="form-label">
                                Дата обслуживания
                            </label>
                            <input
                                type="datetime-local"
                                id="maintenance_date"
                                v-model="maintenance.maintenance_date"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="description" class="form-label">Описание</label>
                            <textarea
                                id="description"
                                v-model="maintenance.description"
                                class="form-control"
                                rows="3"
                                required
                            ></textarea>
                        </div>
                        <div class="mb-3">
                            <label for="performed_by" class="form-label">
                                Сотрудник, выполнивший обслуживание
                            </label>
                            <select
                                id="performed_by"
                                v-model.number="maintenance.performed_by"
                                class="form-control"
                                required
                            >
                                <option disabled value="">Выберите сотрудника</option>
                                <option v-for="emp in employees" :key="emp.id" :value="emp.id">
                                    {{ emp.first_name }} {{ emp.last_name }} (ID: {{ emp.id }})
                                </option>
                            </select>
                        </div>
                        <div class="mb-3">
                            <label for="next_maintenance_date" class="form-label">
                                Дата следующего обслуживания
                            </label>
                            <input
                                type="datetime-local"
                                id="next_maintenance_date"
                                v-model="maintenance.next_maintenance_date"
                                class="form-control"
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
import { ref, computed, defineComponent, watch, toRefs } from 'vue'
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
    name: 'MaintenanceModal',
    props: {
        initialMaintenance: {
            type: Object,
            default: null,
        },
        aircraftId: {
            type: Number,
            required: true,
        },
        employees: {
            type: Array,
            default: () => [],
        },
    },
    emits: ['createMaintenance', 'updateMaintenance'],
    setup(props, { emit }) {
        const { employees } = toRefs(props)

        const maintenance = ref({
            id: null,
            aircraft_id: null,
            maintenance_date: '',
            description: '',
            performed_by: '',
            next_maintenance_date: '',
        })

        const resetForm = () => {
            maintenance.value = {
                id: null,
                aircraft_id: props.aircraftId,
                maintenance_date: '',
                description: '',
                performed_by: '',
                next_maintenance_date: '',
            }
        }

        const isEditMode = computed(() => maintenance.value.id !== null)

        watch(
            () => props.initialMaintenance,
            (newVal) => {
                if (newVal) {
                    maintenance.value = {
                        ...newVal,
                        aircraft_id: props.aircraftId,
                        maintenance_date: fromISO(newVal.maintenance_date),
                        next_maintenance_date: fromISO(newVal.next_maintenance_date),
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
            const preparedMaintenance = {
                ...maintenance.value,
                maintenance_date: toISO(maintenance.value.maintenance_date),
                next_maintenance_date: maintenance.value.next_maintenance_date
                    ? toISO(maintenance.value.next_maintenance_date)
                    : null,
            }

            if (isEditMode.value) {
                emit('updateMaintenance', preparedMaintenance)
            } else {
                preparedMaintenance.aircraft_id = props.aircraftId
                emit('createMaintenance', preparedMaintenance)
            }
            resetForm()
            close()
        }

        return {
            modalElement,
            maintenance,
            isEditMode,
            open,
            close,
            submitForm,
          // eslint-disable-next-line vue/no-dupe-keys
            employees,
        }
    },
})
</script>

<style scoped></style>
