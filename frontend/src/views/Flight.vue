<template>
    <div>
        <Header />
        <div class="container mt-4">
            <router-link to="/" class="back-link">
                <span class="back-arrow">&#8592;</span> Назад
            </router-link>

            <h1>Информация о рейсе</h1>
            <div v-if="flight">
                <p><strong>Номер рейса:</strong> {{ flight.flight_number }}</p>
                <p><strong>Самолёт:</strong> {{ aircraftInfo }}</p>
                <p><strong>Маршрут:</strong> {{ routeInfo }}</p>
                <p><strong>Время вылета:</strong> {{ flight.departure_time }}</p>
                <p><strong>Время прибытия:</strong> {{ flight.arrival_time }}</p>
                <p><strong>Статус:</strong> {{ flight.status }}</p>
                <button class="btn btn-secondary" @click="openBookingModal">Купить билеты</button>
            </div>

            <hr class="my-4" />

            <BookingModal
                ref="bookingModal"
                :initialBooking="null"
                @createBooking="handleCreateBooking"
            />
        </div>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Header from '@/components/Header.vue'
import flightApi from '@/API/flightApi'
import routeApi from '@/API/routeApi'
import aircraftApi from '@/API/aircraftApi'
import airportApi from '@/API/airportApi'
import bookingApi from '@/API/bookingApi'
import BookingModal from '@/components/BookingModal.vue'

export default {
    name: 'FlightDetail',
    components: { Header, BookingModal },
    setup() {
        const route = useRoute()
        const router = useRouter()
        const flightId = parseInt(route.params.id, 10)

        const flight = ref(null)
        const aircraftInfo = ref('')
        const routeInfo = ref('')

        const bookingModal = ref(null)

        const fetchFlight = () => {
            flightApi.getFlight(flightId)
                .then((res) => {
                    flight.value = res.data
                    return Promise.all([
                        aircraftApi.getAircraft(res.data.aircraft_id),
                        routeApi.getRoute(res.data.route_id)
                    ])
                })
                .then(([airRes, rtRes]) => {
                    aircraftInfo.value = `${airRes.data.tail_number} – ${airRes.data.model}`
                    const r = rtRes.data
                    return Promise.all([
                        airportApi.getAirport(r.departure_airport_id),
                        airportApi.getAirport(r.arrival_airport_id)
                    ])
                })
                .then(([depRes, arrRes]) => {
                    routeInfo.value = `${depRes.data.city} (${depRes.data.code}) – ${arrRes.data.city} (${arrRes.data.code})`
                })
                .catch(console.error)
        }

        const formatDate = (isoStr) => {
            if (!isoStr) return ''
            return new Date(isoStr).toLocaleString()
        }

        function openBookingModal() {
            bookingModal.value.open()
        }

        function handleCreateBooking(newBooking) {
            const payload = {
                ...newBooking,
                flight_id: flightId
            }
            bookingApi.createBooking(payload)
                .then(res => {
                    router.push({ name: 'BookingDetails', params: { id: res.data.id } })
                })
                .catch(err => {
                    console.error('Ошибка при создании бронирования', err)
                })
        }

        onMounted(fetchFlight)

        return {
            flightId,
            flight,
            aircraftInfo,
            routeInfo,
            formatDate,
            bookingModal,
            openBookingModal,
            handleCreateBooking,
        }
    }
}
</script>

<style scoped></style>
