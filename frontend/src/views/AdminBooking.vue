<template>
    <div>
        <Header />
        <div class="container mt-4">
            <router-link to="/admin/bookings" class="back-link">
                <span class="back-arrow">&#8592;</span> Назад
            </router-link>
            <h1>Информация о бронировании</h1>
            <div v-if="booking" class="mt-3">
                <p><strong>Пассажир:</strong> {{ getPassengerName(booking.passenger_id) }}</p>
                <p><strong>Дата бронирования:</strong> {{ booking.booking_date }}</p>
                <p><strong>Статус:</strong> {{ booking.status }}</p>
                <button class="btn btn-secondary" @click="openEditBookingModal">Редактировать</button>
            </div>

            <hr class="my-4" />

            <div class="d-flex justify-content-between align-items-center">
                <h2>Билеты</h2>
                <button class="btn btn-primary" @click="openTicketModal">Добавить билет</button>
            </div>
            <div class="row mt-3">
                <div v-for="ticket in tickets" :key="ticket.id" class="col-md-4 mb-3">
                    <div class="card" style="cursor: pointer" @click="openEditTicketModal(ticket)">
                        <div class="card-body">
                            <h5 class="card-title">Место: {{ ticket.seat_number }}</h5>
                            <p class="card-text">
                                Цена: {{ ticket.price }}<br />
                                Пассажир: {{ getPassengerName(ticket.passenger_id) }}<br />
                                Рейс: {{ getFlightNumber(ticket.flight_id) }}
                            </p>
                            <button
                                class="btn btn-danger btn-sm float-end"
                                @click.stop="removeTicket(ticket.id)"
                            >
                                ×
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <BookingModal
            ref="bookingModal"
            :initialBooking="selectedBooking"
            @updateBooking="handleUpdateBooking"
        />
        <TicketModal
            ref="ticketModal"
            :bookingId="bookingId"
            :initialTicket="selectedTicket"
            @createTicket="handleAddTicket"
            @updateTicket="handleUpdateTicket"
        />
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import Header from '@/components/Header.vue'
import BookingModal from '@/components/BookingModal.vue'
import TicketModal from '@/components/TicketModal.vue'
import bookingApi from '@/API/bookingApi'
import ticketApi from '@/API/ticketApi'
import passengerApi from '@/API/passengerApi'
import flightApi from '@/API/flightApi'

export default {
    name: 'AdminBooking',
    components: { Header, BookingModal, TicketModal },
    setup() {
        const route = useRoute()
        const bookingId = parseInt(route.params.id, 10)

        const booking = ref(null)
        const tickets = ref([])
        const passengers = ref([])
        const flights = ref([])
        const selectedBooking = ref(null)
        const selectedTicket = ref(null)

        const fetchBooking = () => {
            bookingApi
                .getBooking(bookingId)
                .then((res) => {
                    booking.value = res.data
                })
                .catch((err) => console.error('Ошибка получения бронирования', err))
        }

        const fetchTickets = () => {
            ticketApi
                .getTicketsByBooking(bookingId)
                .then((res) => {
                    tickets.value = res.data
                })
                .catch((err) => console.error('Ошибка получения билетов', err))
        }

        const fetchPassengers = () => {
            passengerApi
                .getPassengers()
                .then((res) => {
                    passengers.value = res.data
                })
                .catch((err) => console.error('Ошибка получения пассажиров', err))
        }

        const fetchFlights = () => {
            flightApi
                .getFlights()
                .then((res) => {
                    flights.value = res.data
                })
                .catch((err) => console.error('Ошибка получения рейсов', err))
        }

        const getPassengerName = (id) => {
            const p = passengers.value.find((x) => x.id === id)
            return p ? `${p.first_name} ${p.last_name}` : '–'
        }

        const getFlightNumber = (id) => {
            const f = flights.value.find((x) => x.id === id)
            return f ? f.flight_number : '–'
        }

        const openEditBookingModal = () => {
            selectedBooking.value = { ...booking.value }
            bookingModal.value.open()
        }

        const handleUpdateBooking = (data) => {
            bookingApi
                .updateBooking(data.id, data)
                .then(fetchBooking)
                .catch((err) => console.error('Ошибка обновления бронирования', err))
        }

        const openTicketModal = () => {
            selectedTicket.value = null
            ticketModal.value.open()
        }

        const openEditTicketModal = (ticket) => {
            selectedTicket.value = { ...ticket }
            ticketModal.value.open()
        }

        const handleAddTicket = (data) => {
            ticketApi
                .addTicketForBooking(bookingId, data)
                .then(fetchTickets)
                .catch((err) => console.error('Ошибка добавления билета', err))
        }

        const handleUpdateTicket = (data) => {
            ticketApi
                .updateTicketForBooking(data.id, data)
                .then(fetchTickets)
                .catch((err) => console.error('Ошибка обновления билета', err))
        }

        const removeTicket = (id) => {
            ticketApi
                .deleteTicket(id)
                .then(fetchTickets)
                .catch((err) => console.error('Ошибка удаления билета', err))
        }

        const bookingModal = ref(null)
        const ticketModal = ref(null)

        onMounted(() => {
            fetchPassengers()
            fetchFlights()
            fetchBooking()
            fetchTickets()
        })

        return {
            bookingId,
            booking,
            tickets,
            selectedBooking,
            selectedTicket,
            bookingModal,
            ticketModal,
            getPassengerName,
            getFlightNumber,
            openEditBookingModal,
            handleUpdateBooking,
            openTicketModal,
            openEditTicketModal,
            handleAddTicket,
            handleUpdateTicket,
            removeTicket,
        }
    },
}
</script>

<style scoped></style>
