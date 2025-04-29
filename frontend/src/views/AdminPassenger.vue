<template>
    <div>
        <Header />

        <div class="container mt-4">
            <router-link to="/admin/passengers" class="back-link">
                <span class="back-arrow">&#8592;</span> Назад
            </router-link>
            <h1>Информация о пассажире</h1>

            <div v-if="passenger" class="mt-3">
                <p><strong>Имя:</strong> {{ passenger.first_name }}</p>
                <p><strong>Фамилия:</strong> {{ passenger.last_name }}</p>
                <p><strong>Email:</strong> {{ passenger.email }}</p>
                <p><strong>Телефон:</strong> {{ passenger.phone }}</p>
                <p><strong>Паспорт:</strong> {{ passenger.passport_number }}</p>
                <button class="btn btn-secondary" @click="openEditModal">
                    Редактировать пассажира
                </button>
            </div>

            <hr class="my-4" />

            <div class="d-flex justify-content-between align-items-center">
                <h2>Рейсы пассажира</h2>
            </div>
            <div v-if="flights.length" class="row mt-3">
                <div
                    v-for="flight in flights"
                    :key="flight.id"
                    class="col-md-4 mb-3"
                    style="cursor: pointer"
                    @click="goToFlight(flight.id)"
                >
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">Рейс {{ flight.flight_number }}</h5>
                            <p class="card-text">
                                Вылет: {{ flight.departure_time }}<br />
                                Прилет: {{ flight.arrival_time }}<br />
                                Статус: {{ flight.status }}
                            </p>
                        </div>
                    </div>
                </div>
            </div>
            <p v-else class="text-muted">Нет рейсов.</p>

            <hr class="my-4" />

            <h2>Билеты пассажира</h2>
            <div v-if="tickets.length" class="row mt-3">
                <div
                    v-for="ticket in tickets"
                    :key="ticket.id"
                    class="col-md-4 mb-3"
                >
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">Место: {{ ticket.seat_number }}</h5>
                            <p class="card-text">
                                Цена: {{ ticket.price }}<br />
                                Рейс: {{ ticket.flight_id }}
                            </p>
                            <button
                                class="btn btn-danger btn-sm float-end"
                                @click="deleteTicket(ticket.id)"
                            >
                                ×
                            </button>
                        </div>
                    </div>
                </div>
            </div>
            <p v-else class="text-muted">Нет билетов.</p>
        </div>

        <PassengerModal
            ref="editModal"
            :initialPassenger="passenger"
            @updatePassenger="handleUpdate"
        />
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Header from '@/components/Header.vue'
import PassengerModal from '@/components/PassengerModal.vue'
import passengerApi from '@/API/passengerApi'
import ticketApi from '@/API/ticketApi'
import flightApi from '@/API/flightApi'

export default {
    name: 'AdminPassenger',
    components: { Header, PassengerModal },
    setup() {
        const route = useRoute()
        const router = useRouter()
        const passengerId = parseInt(route.params.id, 10)

        const passenger = ref(null)
        const tickets = ref([])
        const flights = ref([])
        const editModal = ref(null)

        const fetchPassenger = () => {
            passengerApi
                .getPassenger(passengerId)
                .then(res => {
                    passenger.value = res.data
                })
                .catch(err => console.error('Ошибка получения пассажира', err))
        }

        const fetchTickets = () => {
            ticketApi
                .getTicketsByPassenger(passengerId)
                .then(res => {
                    tickets.value = res.data
                    fetchFlights()
                })
                .catch(err => console.error('Ошибка получения билетов', err))
        }

        const fetchFlights = () => {
            const ids = [...new Set(tickets.value.map(t => t.flight_id))]
            if (!ids.length) {
                flights.value = []
                return
            }
            Promise.all(ids.map(id => flightApi.getFlight(id)))
                .then(responses => {
                    flights.value = responses.map(r => r.data)
                })
                .catch(err => console.error('Ошибка получения рейсов', err))
        }

        const goToFlight = (id) => {
            router.push({ name: 'AdminFlight', params: { id } })
        }

        const deleteTicket = (id) => {
            ticketApi
                .deleteTicket(id)
                .then(fetchTickets)
                .catch(err => console.error('Ошибка удаления билета', err))
        }

        const openEditModal = () => {
            editModal.value.open()
        }

        const handleUpdate = (upd) => {
            passengerApi
                .updatePassenger(upd.id, upd)
                .then(fetchPassenger)
                .catch(err => console.error('Ошибка обновления пассажира', err))
        }

        onMounted(() => {
            fetchPassenger()
            fetchTickets()
        })

        return {
            passenger,
            tickets,
            flights,
            editModal,
            goToFlight,
            deleteTicket,
            openEditModal,
            handleUpdate
        }
    }
}
</script>

<style scoped></style>
