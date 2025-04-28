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
                            <h5 class="card-title">
                                {{ member.first_name }} {{ member.last_name }}
                            </h5>
                            <p class="card-text">
                                Email: {{ member.email }}<br />
                                Телефон: {{ member.phone }}<br />
                                Дата найма: {{ formatDate(member.hire_date) }}<br />
                                Зарплата: {{ member.salary }}<br />
                                Должность: {{ getRoleName(member.role_id) }}
                            </p>
                            <button
                                class="btn btn-danger btn-sm float-end"
                                @click="deleteCrewMember(member.id)"
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
        </div>

        <CrewModal
            ref="crewModal"
            :flightId="flightId"
            :initialCrew="crew"
            @updateCrew="fetchCrew"
        />
        <TicketModal ref="ticketModal" :flightId="flightId" @addTicket="handleAddTicket" />
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
        const roles = ref([])
        const tickets = ref([])
        const aircraftInfo = ref('')
        const routeInfo = ref('')

        const fetchFlight = () => {
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

        const fetchRoles = () => {
            roleApi
                .getRoles()
                .then((response) => {
                    roles.value = response.data
                })
                .catch((error) => {
                    console.error('Ошибка получения ролей', error)
                })
        }

        const fetchCrew = () => {
            crewApi
                .getCrewByFlight(flightId)
                .then((res) => {
                    crew.value = res.data || []
                })
                .catch(console.error)
        }

        const fetchTickets = () => {
            ticketApi
                .getTicketsByFlight(flightId)
                .then((res) => {
                    tickets.value = res.data || []
                })
                .catch(console.error)
        }

        const deleteCrewMember = (employeeId) => {
            crewApi.removeCrewMember(flightId, employeeId).then(fetchCrew).catch(console.error)
        }

        const removeTicket = (ticketId) => {
            ticketApi.deleteTicket(ticketId).then(fetchTickets).catch(console.error)
        }

        const openCrewModal = () => crewModal.value.open()
        const openTicketModal = () => ticketModal.value.open()
        const editFlight = () => flightModal.value.open()

        const handleAddTicket = () => fetchTickets()

        const handleUpdateFlight = (updatedFlight) => {
            flightApi
                .updateFlight(updatedFlight.id, updatedFlight)
                .then(fetchFlight)
                .catch((err) => console.error('Ошибка обновления рейса', err))
        }

        const crewModal = ref(null)
        const ticketModal = ref(null)
        const flightModal = ref(null)

        const formatDate = (isoStr) => {
            if (!isoStr) return ''
            const date = new Date(isoStr)
            return date.toLocaleString()
        }

        const getRoleName = (roleId) => {
            const role = roles.value.find((r) => r.id === roleId)
            return role ? role.name : 'Не задана'
        }

        onMounted(() => {
            fetchFlight()
            fetchRoles()
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
            editFlight,
            fetchRoles,
            fetchCrew,
            handleAddTicket,
            deleteCrewMember,
            removeTicket,
            handleUpdateFlight,
            crewModal,
            ticketModal,
            flightModal,
            formatDate,
            getRoleName,
        }
    },
}
</script>

<style scoped></style>
