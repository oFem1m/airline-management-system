<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">Назначение экипажа</h5>
                    <button type="button" class="btn-close" @click="close"></button>
                </div>
                <div class="modal-body">
                    <label class="form-label">Выберите сотрудников</label>
                    <div v-for="emp in employees" :key="emp.id" class="form-check">
                        <input
                            class="form-check-input"
                            type="checkbox"
                            :id="'emp-' + emp.id"
                            :value="emp.id"
                            v-model="selected"
                        />
                        <label class="form-check-label" :for="'emp-' + emp.id">
                            {{ emp.first_name }} {{ emp.last_name }} (ID: {{ emp.id }})
                        </label>
                    </div>
                </div>
                <div class="modal-footer">
                    <button class="btn btn-secondary" @click="close">Отмена</button>
                    <button class="btn btn-primary" @click="submitForm">Сохранить</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { ref, onMounted, watch, defineComponent } from 'vue'
import { Modal } from 'bootstrap'
import employeeApi from '@/API/employeeApi'

export default defineComponent({
    name: 'CrewModal',
    props: {
        flightId: { type: Number, required: true },
        initialCrew: { type: Array, default: () => [] },
    },
    emits: ['updateCrew'],
    setup(props, { emit }) {
        const employees = ref([])
        const selected = ref([])
        const modalElement = ref(null)
        let modalInstance = null

        const fetchEmployees = () =>
            employeeApi.getEmployees().then((r) => (employees.value = r.data))

        watch(
            () => props.initialCrew,
            (crew) => {
                selected.value = crew.map((m) => m.employee_id)
            },
            { immediate: true },
        )

        const open = () => {
            if (!modalInstance) modalInstance = new Modal(modalElement.value)
            modalInstance.show()
        }
        const close = () => modalInstance && modalInstance.hide()

        const submitForm = () => {
            emit('updateCrew', selected.value)
            close()
        }

        onMounted(fetchEmployees)

        return { employees, selected, modalElement, open, close, submitForm }
    },
})
</script>

<style scoped></style>
