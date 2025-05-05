<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">
                        {{ isEditMode ? 'Редактировать бронирование' : 'Добавить бронирование' }}
                    </h5>
                    <button type="button" class="btn-close" @click="close" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <form @submit.prevent="submitForm">
                        <div class="mb-3">
                            <label for="passenger" class="form-label">Пассажир</label>
                            <select id="passenger" v-model.number="booking.passenger_id" class="form-control" required>
                                <option disabled value="">Выберите пассажира</option>
                                <option v-for="p in passengers" :key="p.id" :value="p.id">
                                    {{ p.first_name }} {{ p.last_name }}
                                </option>
                            </select>
                        </div>
                        <div class="mb-3">
                            <label for="bookingDate" class="form-label">Дата бронирования</label>
                            <input
                                type="datetime-local"
                                id="bookingDate"
                                v-model="booking.booking_date"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="status" class="form-label">Статус</label>
                            <select id="status" v-model="booking.status" class="form-control" required>
                                <option value="pending">Pending</option>
                                <option value="confirmed">Confirmed</option>
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
import { ref, computed, watch, onMounted } from 'vue'
import { Modal } from 'bootstrap'
import passengerApi from '@/API/passengerApi'
import { fromISO, toISO } from '../../utils/time.js'

export default {
    name: 'BookingModal',
    props: {
        initialBooking: {
            type: Object,
            default: null,
        },
    },
    emits: ['createBooking', 'updateBooking'],
    setup(props, { emit }) {
        const booking = ref({
            id: null,
            passenger_id: '',
            booking_date: '',
            status: 'pending',
        })
        const passengers = ref([])
        const modalElement = ref(null)
        let modalInstance = null

        function getLocalDateTime() {
            const d = new Date()
            const pad = n => String(n).padStart(2, '0')
            return `${d.getFullYear()}-${pad(d.getMonth()+1)}-${pad(d.getDate())}T${pad(d.getHours())}:${pad(d.getMinutes())}`
        }

        const isEditMode = computed(() => booking.value.id !== null)

        const resetForm = () => {
            booking.value = {
                id: null,
                passenger_id: '',
                booking_date: getLocalDateTime(),
                status: 'pending',
            }
        }

        watch(
            () => props.initialBooking,
            (newVal) => {
                if (newVal) {
                    booking.value = {
                        id: newVal.id,
                        passenger_id: newVal.passenger_id,
                        booking_date: fromISO(newVal.booking_date),
                        status: newVal.status,
                    }
                } else {
                    resetForm()
                }
            },
            { immediate: true }
        )

        const fetchPassengers = () => {
            passengerApi.getPassengers()
                .then(res => { passengers.value = res.data })
                .catch(err => console.error('Ошибка загрузки пассажиров', err))
        }

        const open = () => {
            if (!isEditMode.value) {
                resetForm()
            }
            if (!modalInstance) modalInstance = new Modal(modalElement.value)
            modalInstance.show()
        }
        const close = () => {
            if (modalInstance) modalInstance.hide()
        }

        const submitForm = () => {
            const payload = {
                ...booking.value,
                booking_date: toISO(booking.value.booking_date),
            }
            if (isEditMode.value) {
                emit('updateBooking', payload)
            } else {
                emit('createBooking', payload)
            }
            resetForm()
            close()
        }

        onMounted(fetchPassengers)

        return {
            booking,
            passengers,
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
