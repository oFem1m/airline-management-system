<template>
    <div>
        <Header />
        <div class="container mt-4">
            <a href="#" class="back-link" @click.prevent="goBack">
                <span class="back-arrow">&#8592;</span> Назад
            </a>
            <h1>Информация о бронировании</h1>
            <div v-if="booking" class="mt-3">
                <p><strong>Пассажир:</strong> {{ getPassengerName(booking.passenger_id) }}</p>
                <p><strong>Дата бронирования:</strong> {{ booking.booking_date }}</p>
                <p><strong>Статус:</strong> {{ booking.status }}</p>
            </div>

            <hr class="my-4" />

            <div class="d-flex justify-content-between align-items-center">
                <h2>Билеты</h2>
                <button class="btn btn-primary" @click="openTicketModal">Добавить билет</button>
            </div>
            <div class="row mt-3">
                <div v-for="ticket in tickets" :key="ticket.id" class="col-md-4 mb-3">
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">Место: {{ ticket.seat_number }}</h5>
                            <p class="card-text">
                                Цена: {{ ticket.price }}<br />
                                Пассажир: {{ getPassengerName(ticket.passenger_id) }}<br />
                                Рейс: {{ getFlightNumber(ticket.flight_id) }}
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <TicketModal
            ref="ticketModal"
            :bookingId="bookingId"
            :flightId="flightId"
            :price=5000
            :initialTicket="selectedTicket"
            @createTicket="handleAddTicket"
        />
    </div>
</template>

<script>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Header from '@/components/Header.vue'
import TicketModal from '@/components/TicketModal.vue'
import bookingApi from '@/API/bookingApi'
import ticketApi from '@/API/ticketApi'
import passengerApi from '@/API/passengerApi'
import flightApi from '@/API/flightApi'

export default {
    name: 'AdminBooking',
    components: { Header, TicketModal },
    setup() {
        const route = useRoute()
        const router = useRouter()
        const bookingId = parseInt(route.params.id, 10)
        const flightId = parseInt(route.params.flight_id, 10)

        const booking = ref(null)
        const tickets = ref([])
        const passengers = ref([])
        const flights = ref([])
        const selectedTicket = ref(null)

        const fetchBooking = () => {
            bookingApi.getBooking(bookingId)
                .then(res => { booking.value = res.data })
                .catch(err => console.error('Ошибка получения бронирования', err))
        }

        const fetchTickets = () => {
            ticketApi.getTicketsByBooking(bookingId)
                .then(res => { tickets.value = res.data })
                .catch(err => console.error('Ошибка получения билетов', err))
        }

        const fetchPassengers = () => {
            passengerApi.getPassengers()
                .then(res => { passengers.value = res.data })
                .catch(err => console.error('Ошибка получения пассажиров', err))
        }

        const fetchFlights = () => {
            flightApi.getFlights()
                .then(res => { flights.value = res.data })
                .catch(err => console.error('Ошибка получения рейсов', err))
        }

        const getPassengerName = id => {
            const p = passengers.value.find(x => x.id === id)
            return p ? `${p.first_name} ${p.last_name}` : '–'
        }

        const getFlightNumber = id => {
            const f = flights.value.find(x => x.id === id)
            return f ? f.flight_number : '–'
        }

        const passengerName = computed(() => {
            if (!booking.value) return ''
            const p = passengers.value.find(x => x.id === booking.value.passenger_id)
            return p ? `${p.first_name} ${p.last_name}` : '–'
        })

        const openTicketModal = () => {
            selectedTicket.value = null
            ticketModal.value.open()
        }

        const handleAddTicket = data => {
            ticketApi.addTicketForBooking(bookingId, data)
                .then(fetchTickets)
                .catch(err => console.error('Ошибка добавления билета', err))
        }

        const ticketModal = ref(null)

        function goBack() {
            router.back()
        }

        onMounted(() => {
            fetchPassengers()
            fetchFlights()
            fetchBooking()
            fetchTickets()
        })

        return {
            bookingId,
            flightId,
            booking,
            tickets,
            selectedTicket,
            ticketModal,
            getPassengerName,
            getFlightNumber,
            passengerName,
            openTicketModal,
            handleAddTicket,
            goBack,
        }
    },
}
</script>

<style scoped></style>
