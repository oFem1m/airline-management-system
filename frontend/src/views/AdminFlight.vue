<template>
    <div>
        <Header />
        <div class="container mt-4">
            <router-link
                v-if="flight"
                :to="{ name: 'AdminRoute', params: { id: flight.route_id } }"
                class="back-link"
            >
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
                            <h5 class="card-title">{{ member.first_name }} {{ member.last_name }}</h5>
                            <p class="card-text">
                                Email: {{ member.email }}<br />
                                Телефон: {{ member.phone }}<br />
                                Дата найма: {{ formatDate(member.hire_date) }}<br />
                                Зарплата: {{ member.salary }}<br />
                                Должность: {{ getRoleName(member.role_id) }}
                            </p>
                            <button class="btn btn-danger btn-sm float-end" @click="deleteCrewMember(member.id)">×</button>
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
                    <div class="card" style="cursor: pointer" @click="openEditTicketModal(ticket)">
                        <div class="card-body">
                            <h5 class="card-title">Место: {{ ticket.seat_number }}</h5>
                            <p class="card-text">
                                Цена: {{ ticket.price }}<br />
                                Пассажир: {{ getPassengerName(ticket.passenger_id) }}
                            </p>
                            <button class="btn btn-danger btn-sm float-end" @click.stop="removeTicket(ticket.id)">×</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <CrewModal
            ref="crewModal"
            :flightId="flightId"
            :initialCrew="crew"
            @updateCrew="fetchCrew"
        />
        <TicketModal
            ref="ticketModal"
            :flightId="flightId"
            :initialTicket="selectedTicket"
            @createTicket="handleAddTicket"
            @updateTicket="handleUpdateTicket"
        />
        <FlightModal
            v-if="flight"
            ref="flightModal"
            :initialFlight="flight"
            :routeId="flight.route_id"
            @updateFlight="handleUpdateFlight"
        />
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import Header from '@/components/Header.vue'
import CrewModal from '@/components/CrewModal.vue'
import TicketModal from '@/components/TicketModal.vue'
import FlightModal from '@/components/FlightModal.vue'
import flightApi from '@/API/flightApi'
import crewApi from '@/API/crewApi'
import ticketApi from '@/API/ticketApi'
import passengerApi from '@/API/passengerApi'
import routeApi from '@/API/routeApi'
import aircraftApi from '@/API/aircraftApi'
import airportApi from '@/API/airportApi'
import roleApi from '@/API/roleApi.js'

export default {
    name: 'AdminFlight',
    components: { Header, CrewModal, TicketModal, FlightModal },
    setup() {
        const route = useRoute()
        const flightId = parseInt(route.params.id, 10)

        const flight = ref(null)
        const crew = ref([])
        const tickets = ref([])
        const passengers = ref([])
        const aircraftInfo = ref('')
        const routeInfo = ref('')
        const selectedTicket = ref(null)
        const roles = ref([])

        // Существующая логика
        const fetchFlight = () => {
            flightApi.getFlight(flightId)
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

        const fetchRoles = () => {
            roleApi.getRoles()
                .then(res => { roles.value = res.data })
                .catch(err => console.error('Ошибка получения ролей', err))
        }

        const fetchCrew = () => {
            crewApi.getCrewByFlight(flightId)
                .then(res => { crew.value = res.data || [] })
                .catch(console.error)
        }

        const fetchTickets = () => {
            ticketApi.getTicketsByFlight(flightId)
                .then(res => { tickets.value = res.data || [] })
                .catch(err => console.error('Ошибка получения билетов', err))
        }

        const fetchPassengers = () => {
            passengerApi.getPassengers()
                .then(res => { passengers.value = res.data })
                .catch(err => console.error('Ошибка получения пассажиров', err))
        }

        const deleteCrewMember = (employeeId) => {
            crewApi.removeCrewMember(flightId, employeeId)
                .then(fetchCrew)
                .catch(console.error)
        }

        const removeTicket = (ticketId) => {
            ticketApi.deleteTicket(ticketId)
                .then(fetchTickets)
                .catch(err => console.error('Ошибка удаления билета', err))
        }

        const openCrewModal = () => crewModal.value.open()
        const openTicketModal = () => {
            selectedTicket.value = null
            ticketModal.value.open()
        }
        const openEditTicketModal = (ticket) => {
            selectedTicket.value = { ...ticket }
            ticketModal.value.open()
        }
        const editFlight = () => flightModal.value.open()

        const handleAddTicket = (data) => {
            ticketApi.addTicket(flightId, data)
                .then(fetchTickets)
                .catch(err => console.error('Ошибка при создании билета', err))
        }
        const handleUpdateTicket = (data) => {
            ticketApi.updateTicket(data.id, data)
                .then(fetchTickets)
                .catch(err => console.error('Ошибка при обновлении билета', err))
        }

        const handleUpdateFlight = (updatedFlight) => {
            flightApi.updateFlight(updatedFlight.id, updatedFlight)
                .then(fetchFlight)
                .catch(err => console.error('Ошибка обновления рейса', err))
        }

        const getRoleName = (roleId) => {
            const r = roles.value.find(x => x.id === roleId)
            return r ? r.name : 'Не задана'
        }

        const getPassengerName = (id) => {
            const p = passengers.value.find(x => x.id === id)
            return p ? `${p.first_name} ${p.last_name}` : '–'
        }

        const formatDate = (isoStr) => {
            if (!isoStr) return ''
            return new Date(isoStr).toLocaleString()
        }

        const crewModal = ref(null)
        const ticketModal = ref(null)
        const flightModal = ref(null)

        onMounted(() => {
            fetchFlight()
            fetchRoles()
            fetchCrew()
            fetchTickets()
            fetchPassengers()
        })

        return {
            flightId,
            flight,
            crew,
            tickets,
            passengers,
            aircraftInfo,
            routeInfo,
            selectedTicket,
            roles,
            fetchCrew,
            openCrewModal,
            openTicketModal,
            openEditTicketModal,
            editFlight,
            deleteCrewMember,
            removeTicket,
            handleAddTicket,
            handleUpdateTicket,
            handleUpdateFlight,
            getRoleName,
            getPassengerName,
            formatDate,
            crewModal,
            ticketModal,
            flightModal,
        }
    },
}
</script>

<style scoped></style>
