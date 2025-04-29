<template>
    <div>
        <Header />

        <div class="container mt-4">
            <router-link to="/admin" class="back-link">
                <span class="back-arrow">&#8592;</span> Назад
            </router-link>
            <h1>Управление бронированиями</h1>

            <div class="d-flex justify-content-between align-items-center my-4">
                <h2>Бронирования</h2>
                <button class="btn btn-primary" @click="openCreateModal">
                    Добавить бронирование
                </button>
            </div>

            <div class="row">
                <div v-for="b in bookings" :key="b.id" class="col-md-4 mb-3">
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">№ {{ b.id }}</h5>
                            <p class="card-text">
                                Пассажир: {{ getPassengerName(b.passenger_id) }}<br />
                                Дата бронирования: {{ b.booking_date }}<br />
                                Статус: {{ b.status }}
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <BookingModal
            ref="bookingModal"
            @createBooking="handleCreate"
        />
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import Header from '@/components/Header.vue'
import BookingModal from '@/components/BookingModal.vue'
import bookingApi from '@/API/bookingApi'
import passengerApi from '@/API/passengerApi'

export default {
    name: 'AdminBookings',
    components: { Header, BookingModal },
    setup() {
        const bookings = ref([])
        const passengers = ref([])
        const bookingModal = ref(null)

        const fetchBookings = () => {
            bookingApi.getBookings()
                .then(res => {
                    bookings.value = res.data
                })
                .catch(err => console.error('Ошибка получения бронирований', err))
        }

        const fetchPassengers = () => {
            passengerApi.getPassengers()
                .then(res => {
                    passengers.value = res.data
                })
                .catch(err => console.error('Ошибка получения пассажиров', err))
        }

        const getPassengerName = id => {
            const p = passengers.value.find(x => x.id === id)
            return p ? `${p.first_name} ${p.last_name} (ID: ${p.id})` : id
        }

        const openCreateModal = () => {
            bookingModal.value.open()
        }

        const handleCreate = newBooking => {
            bookingApi.createBooking(newBooking)
                .then(() => {
                    fetchBookings()
                })
                .catch(err => console.error('Ошибка создания бронирования', err))
        }

        onMounted(() => {
            fetchPassengers()
            fetchBookings()
        })

        return {
            bookings,
            bookingModal,
            openCreateModal,
            handleCreate,
            getPassengerName,
        }
    },
}
</script>

<style scoped></style>
