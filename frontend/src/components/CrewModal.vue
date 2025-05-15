<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">Назначение экипажа</h5>
                    <button type="button" class="btn-close" @click="close"></button>
                </div>
                <div v-if="availableEmployees.length" class="modal-body">
                    <label class="form-label">Выберите сотрудников</label>
                    <div v-for="emp in availableEmployees" :key="emp.id" class="form-check">
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
                <div v-else class="modal-body">
                    <label class="form-label">Нет доступных сотрудников</label>
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
import { ref, computed, onMounted, defineComponent } from 'vue'
import { Modal } from 'bootstrap'
import employeeApi from '@/API/employeeApi'
import crewApi from '@/API/crewApi'

export default defineComponent({
    name: 'CrewModal',
    props: {
        flightId: { type: Number, required: true },
        initialCrew: { type: Array, default: () => [] },
    },
    emits: ['updateCrew'],
    setup(props, { emit }) {
        const allEmployees = ref([])
        const selected = ref([])
        const modalElement = ref(null)
        let modalInstance = null

        const initialCrewIds = computed(() =>
            props.initialCrew.map(member => member.id)
        )

        const availableEmployees = computed(() =>
            allEmployees.value.filter(emp => !initialCrewIds.value.includes(emp.id))
        )

        const fetchEmployees = () =>
            employeeApi.getEmployees().then(res => {
                allEmployees.value = res.data
            })

        const open = () => {
            selected.value = []
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
            const calls = selected.value.map(empId =>
                crewApi.addCrewMember(props.flightId, { employeeId: empId })
            )
            Promise.all(calls)
                .then(() => {
                    emit('updateCrew')
                    close()
                })
                .catch(err => alert(err.response.data))
        }

        onMounted(fetchEmployees)

        return {
            availableEmployees,
            selected,
            modalElement,
            open,
            close,
            submitForm,
        }
    },
})
</script>

<style scoped></style>
