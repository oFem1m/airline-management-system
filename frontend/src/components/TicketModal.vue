<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">
                        {{ isEditMode ? 'Редактировать билет' : 'Добавить билет' }}
                    </h5>
                    <button type="button" class="btn-close" aria-label="Close" @click="close"></button>
                </div>
                <div class="modal-body">
                    <form @submit.prevent="submitForm">
                        <div class="mb-3">
                            <label for="flight" class="form-label">Рейс</label>
                            <select
                                id="flight"
                                v-model.number="ticket.flight_id"
                                class="form-control"
                                required
                            >
                                <option disabled value="">Выберите рейс</option>
                                <option
                                    v-for="f in flights"
                                    :key="f.id"
                                    :value="f.id"
                                >
                                    {{ f.flight_number }}
                                </option>
                            </select>
                        </div>
                        <div class="mb-3">
                            <label for="passenger" class="form-label">Пассажир</label>
                            <select
                                id="passenger"
                                v-model.number="ticket.passenger_id"
                                class="form-control"
                                required
                            >
                                <option disabled value="">Выберите пассажира</option>
                                <option
                                    v-for="p in passengers"
                                    :key="p.id"
                                    :value="p.id"
                                >
                                    {{ p.first_name }} {{ p.last_name }}
                                </option>
                            </select>
                        </div>
                        <div class="mb-3">
                            <label for="seatNumber" class="form-label">Место</label>
                            <input
                                type="text"
                                id="seatNumber"
                                v-model="ticket.seat_number"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="price" class="form-label">Цена</label>
                            <input
                                type="number"
                                id="price"
                                v-model.number="ticket.price"
                                class="form-control"
                                required
                            />
                        </div>
                        <button type="submit" class="btn btn-primary">
                            {{ isEditMode ? 'Сохранить' : 'Добавить' }}
                        </button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { ref, computed, watch, onMounted } from 'vue'
import { Modal } from 'bootstrap'
import passengerApi from '@/API/passengerApi'
import flightApi from '@/API/flightApi'

export default {
    name: 'TicketModal',
    props: {
        flightId: {
            type: Number,
            default: null,
        },
        bookingId: {
            type: Number,
            default: null,
        },
        initialTicket: {
            type: Object,
            default: null,
        },
    },
    emits: ['createTicket', 'updateTicket'],
    setup(props, { emit }) {
        const ticket = ref({
            id: null,
            flight_id: props.flightId || '',
            passenger_id: '',
            booking_id: props.bookingId || null,
            seat_number: '',
            price: 0,
        })
        const passengers = ref([])
        const flights = ref([])
        const modalElement = ref(null)
        let modalInstance = null

        const isEditMode = computed(() => ticket.value.id !== null)

        const fetchPassengers = () => {
            passengerApi.getPassengers()
                .then(res => {
                    passengers.value = res.data
                })
                .catch(err => console.error('Ошибка загрузки пассажиров', err))
        }

        const fetchFlights = () => {
            flightApi.getFlights()
                .then(res => {
                    flights.value = res.data
                })
                .catch(err => console.error('Ошибка загрузки рейсов', err))
        }

        watch(
            () => props.initialTicket,
            (newVal) => {
                if (newVal && newVal.id != null) {
                    ticket.value = {
                        id: newVal.id,
                        flight_id: newVal.flight_id,
                        passenger_id: newVal.passenger_id,
                        booking_id: newVal.booking_id || null,
                        seat_number: newVal.seat_number,
                        price: newVal.price,
                    }
                } else {
                    ticket.value = {
                        id: null,
                        flight_id: props.flightId || '',
                        passenger_id: '',
                        booking_id: props.bookingId || null,
                        seat_number: '',
                        price: 0,
                    }
                }
            },
            { immediate: true }
        )

        onMounted(() => {
            fetchPassengers()
            fetchFlights()
        })

        const open = () => {
            if (!modalInstance) modalInstance = new Modal(modalElement.value)
            modalInstance.show()
        }
        const close = () => {
            if (modalInstance) modalInstance.hide()
        }

        const submitForm = () => {
            if (isEditMode.value) {
                emit('updateTicket', { ...ticket.value })
            } else {
                emit('createTicket', { ...ticket.value })
            }
            close()
        }

        return {
            ticket,
            passengers,
            flights,
            isEditMode,
            modalElement,
            open,
            close,
            submitForm,
        }
    },
}
</script>

<style scoped></style>
