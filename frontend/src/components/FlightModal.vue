<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">
                        {{ isEditMode ? 'Редактирование рейса' : 'Создание рейса' }}
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
                            <label class="form-label">Номер рейса</label>
                            <input
                                type="text"
                                v-model="flight.flight_number"
                                class="form-control"
                                required
                            />
                        </div>

                        <div class="mb-3">
                            <label class="form-label">Самолет</label>
                            <select v-model="flight.aircraft_id" class="form-control" required>
                                <option disabled value="">Выберите самолет</option>
                                <option
                                    v-for="aircraft in aircrafts"
                                    :key="aircraft.id"
                                    :value="aircraft.id"
                                >
                                    {{ aircraft.tail_number }} - {{ aircraft.model }}
                                </option>
                            </select>
                        </div>

                        <div class="mb-3">
                            <label class="form-label">Время вылета</label>
                            <input
                                type="datetime-local"
                                v-model="flight.departure_time"
                                class="form-control"
                                required
                            />
                        </div>

                        <div class="mb-3">
                            <label class="form-label">Время прибытия</label>
                            <input
                                type="datetime-local"
                                v-model="flight.arrival_time"
                                class="form-control"
                                required
                            />
                        </div>

                        <div class="mb-3">
                            <label class="form-label">Статус</label>
                            <select v-model="flight.status" class="form-control" required>
                                <option value="scheduled">Запланировано</option>
                                <option value="delayed">Задержано</option>
                                <option value="completed">Выполнено</option>
                            </select>
                        </div>

                        <button type="submit" class="btn btn-primary">
                            {{ isEditMode ? 'Сохранить' : 'Создать' }}
                        </button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { ref, computed, onMounted, defineComponent } from 'vue'
import { Modal } from 'bootstrap'
import aircraftApi from '@/API/aircraftApi'
import { toISO } from '../../utils/time.js'

export default defineComponent({
    name: 'FlightModal',
    props: {
        initialFlight: {
            type: Object,
            default: () => ({}), // Default to an empty object
        },
        routeId: {
            type: Number,
            required: true, // Make sure routeId is passed
        },
    },
    emits: ['createFlight', 'updateFlight'],
    setup(props, { emit }) {
        const aircrafts = ref([])
        const flight = ref({
            id: null,
            flight_number: '',
            aircraft_id: '',
            departure_time: '',
            arrival_time: '',
            status: 'scheduled',
        })

        const isEditMode = computed(() => flight.value.id !== null)

        const resetForm = () => {
            flight.value = {
                id: null,
                flight_number: '',
                aircraft_id: '',
                departure_time: '',
                arrival_time: '',
                status: 'scheduled',
            }
        }

        const submitForm = () => {
            const preparedFlight = {
                ...flight.value,
                departure_time: toISO(flight.value.departure_time),
                arrival_time: toISO(flight.value.arrival_time),
                route_id: props.routeId,
            }
            if (isEditMode.value) {
                emit('updateFlight', { ...preparedFlight })
            } else {
                emit('createFlight', { ...preparedFlight })
            }
            resetForm()
            close()
        }

        const modalElement = ref(null)
        let modalInstance = null

        const open = () => {
            if (!modalInstance) {
                modalInstance = new Modal(modalElement.value)
            }
            modalInstance.show()
        }
        const close = () => modalInstance && modalInstance.hide()

        const fetchAircrafts = () => {
            aircraftApi
                .getAircrafts()
                .then((res) => (aircrafts.value = res.data))
                .catch((err) => console.error('Ошибка получения самолётов', err))
        }

        onMounted(() => {
            if (props.initialFlight && Object.keys(props.initialFlight).length) {
                flight.value = { ...props.initialFlight }
            }
            fetchAircrafts()
        })

        return {
            aircrafts,
            flight,
            isEditMode,
            modalElement,
            open,
            close,
            submitForm,
        }
    },
})
</script>

<style scoped></style>
