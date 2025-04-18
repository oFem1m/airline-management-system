<template>
    <div>
        <Header />
        <div class="container mt-4">
            <router-link to="/admin/routes" class="back-link">
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
                <button class="btn btn-secondary" @click="editFlight">Редактировать рейс</button>
            </div>

            <hr class="my-4" />

            <div class="d-flex justify-content-between align-items-center">
                <h2>Экипаж</h2>
                <button class="btn btn-primary" @click="openCrewModal">Назначить сотрудника</button>
            </div>
            <div class="row mt-3">
                <div v-for="member in crew" :key="member.employee_id" class="col-md-4 mb-3">
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">
                                {{ member.first_name }} {{ member.last_name }}
                            </h5>
                            <p class="card-text">{{ member.role_name }}</p>
                            <button
                                class="btn btn-danger btn-sm float-end"
                                @click="removeCrewMember(member.employee_id)"
                            >
                                ×
                            </button>
                        </div>
                    </div>
                </div>
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
                            <p class="card-text">Цена: {{ ticket.price }}</p>
                            <button
                                class="btn btn-danger btn-sm float-end"
                                @click="removeTicket(ticket.id)"
                            >
                                ×
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            <CrewModal ref="crewModal" :flightId="flightId" @addCrewMember="handleAddCrewMember" />
            <TicketModal ref="ticketModal" :flightId="flightId" @addTicket="handleAddTicket" />
        </div>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import Header from '@/components/Header.vue'
import CrewModal from '@/components/CrewModal.vue'
import TicketModal from '@/components/TicketModal.vue'
import flightApi from '@/API/flightApi'
import crewApi from '@/API/crewApi'
import ticketApi from '@/API/ticketApi'
import routeApi from '@/API/routeApi'
import aircraftApi from '@/API/aircraftApi'
import airportApi from '@/API/airportApi'

export default {
    name: 'AdminFlight',
    components: { Header, CrewModal, TicketModal },
    setup() {
        const route = useRoute()
        const flightId = parseInt(route.params.id, 10)

        const flight = ref(null)
        const crew = ref([])
        const tickets = ref([])
        const aircraftInfo = ref('')
        const routeInfo = ref('')

        function fetchFlight() {
            flightApi
                .getFlight(flightId)
                .then((res) => {
                    flight.value = res.data
                    return Promise.all([
                        aircraftApi.getAircraft(res.data.aircraft_id),
                        routeApi.getRoute(res.data.route_id),
                    ])
                })
                .then(([airRes, rtRes]) => {
                    aircraftInfo.value = `${airRes.data.tail_number} – ${airRes.data.model}`
                    const r = rtRes.data
                    return Promise.all([
                        airportApi.getAirport(r.departure_airport_id),
                        airportApi.getAirport(r.arrival_airport_id),
                    ])
                })
                .then(([depRes, arrRes]) => {
                    routeInfo.value = `${depRes.data.city} (${depRes.data.code}) – ${arrRes.data.city} (${arrRes.data.code})`
                })
                .catch(console.error)
        }

        function fetchCrew() {
            crewApi
                .getCrewByFlight(flightId)
                .then((res) => {
                    crew.value = res.data
                })
                .catch(console.error)
        }

        function fetchTickets() {
            ticketApi
                .getTicketsByFlight(flightId)
                .then((res) => {
                    tickets.value = res.data
                })
                .catch(console.error)
        }

        function openCrewModal() {
            crewModal.value.open()
        }

        function openTicketModal() {
            ticketModal.value.open()
        }

        function handleAddCrewMember() {
            fetchCrew()
        }

        function handleAddTicket() {
            fetchTickets()
        }

        function removeCrewMember(employeeId) {
            crewApi.removeCrewMember(flightId, employeeId).then(fetchCrew).catch(console.error)
        }

        function removeTicket(ticketId) {
            ticketApi.deleteTicket(ticketId).then(fetchTickets).catch(console.error)
        }

        const crewModal = ref(null)
        const ticketModal = ref(null)

        onMounted(() => {
            fetchFlight()
            fetchCrew()
            fetchTickets()
        })

        return {
            flightId,
            flight,
            crew,
            tickets,
            aircraftInfo,
            routeInfo,
            openCrewModal,
            openTicketModal,
            handleAddCrewMember,
            handleAddTicket,
            removeCrewMember,
            removeTicket,
            crewModal,
            ticketModal,
        }
    },
}
</script>

<style scoped></style>
