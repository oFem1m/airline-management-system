<template>
    <div>
        <Header />

        <div class="container mt-4">
            <router-link to="/admin/airports" class="back-link">
                <span class="back-arrow">&#8592;</span> Назад
            </router-link>

            <h1>Информация об аэропорте</h1>
            <div v-if="airport" class="mt-3">
                <p><strong>Название:</strong> {{ airport.name }}</p>
                <p><strong>Код:</strong> {{ airport.code }}</p>
                <p><strong>Город:</strong> {{ airport.city }}</p>
                <p><strong>Страна:</strong> {{ airport.country }}</p>
                <p><strong>Часовой пояс:</strong> {{ airport.timezone }}</p>
                <button class="btn btn-secondary" @click="openEdit">Редактировать</button>
            </div>

            <hr class="my-4" />

            <h2>Маршруты</h2>
            <div class="row mt-3">
                <div v-for="r in routes" :key="r.id" class="col-md-4 mb-3">
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">
                                Маршрут {{ label(r.departure_airport_id) }} —
                                {{ label(r.arrival_airport_id) }}
                            </h5>
                            <p class="card-text">
                                Расстояние: {{ r.distance }} км<br />
                                Время: {{ r.duration_minutes }} мин
                            </p>
                            <button
                                class="btn btn-danger btn-sm float-end"
                                @click="deleteRoute(r.id)"
                            >
                                ×
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            <hr class="my-4" />

            <h2>Рейсы (вылет)</h2>
            <div class="row mt-3">
                <div
                    v-for="f in departures"
                    :key="f.id"
                    class="col-md-4 mb-3"
                    style="cursor: pointer"
                    @click="goFlight(f.id)"
                >
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">Рейс {{ f.flight_number }}</h5>
                            <p class="card-text">
                                Время вылета: {{ f.departure_time }}<br />
                                Статус: {{ f.status }}
                            </p>
                            <button
                                class="btn btn-danger btn-sm float-end"
                                @click.stop="deleteFlight(f.id)"
                            >
                                ×
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            <hr class="my-4" />

            <h2>Рейсы (прилет)</h2>
            <div class="row mt-3">
                <div
                    v-for="f in arrivals"
                    :key="f.id"
                    class="col-md-4 mb-3"
                    style="cursor: pointer"
                    @click="goFlight(f.id)"
                >
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">Рейс {{ f.flight_number }}</h5>
                            <p class="card-text">
                                Время прилёта: {{ f.arrival_time }}<br />
                                Статус: {{ f.status }}
                            </p>
                            <button
                                class="btn btn-danger btn-sm float-end"
                                @click.stop="deleteFlight(f.id)"
                            >
                                ×
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            <AirportModal ref="modal" :initialAirport="airport" @updateAirport="handleUpdateAirport" />
        </div>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import Header from '@/components/Header.vue'
import AirportModal from '@/components/AirportModal.vue'
import airportApi from '@/API/airportApi'
import routeApi from '@/API/routeApi'
import flightApi from '@/API/flightApi'

export default {
    name: 'AdminAirport',
    components: { Header, AirportModal },
    setup() {
        const router = useRouter()
        const route = useRoute()
        const airportId = parseInt(route.params.id, 10)

        const airport = ref(null)
        const allAirports = ref([])
        const routes = ref([])
        const departures = ref([])
        const arrivals = ref([])
        const modal = ref(null)

        const fetchAirport = () =>
            airportApi.getAirport(airportId).then((response) => (airport.value = response.data))

        const fetchAllAirports = () =>
            airportApi.getAirports().then((response) => (allAirports.value = response.data))

        const fetchRoutes = () =>
            routeApi
                .getRoutesByAirport(airportId)
                .then((response) => (routes.value = response.data))

        const fetchFlights = () =>
            flightApi.getFlightsByAirport(airportId).then((response) => {
                const flights = response.data || []
                departures.value = flights.filter((flight) => {
                    const route = routes.value.find((x) => x.id === flight.route_id)
                    return route && route.departure_airport_id === airportId
                })
                arrivals.value = flights.filter((flight) => {
                    const rt = routes.value.find((x) => x.id === flight.route_id)
                    return rt && rt.arrival_airport_id === airportId
                })
            })

        const label = (airportId) => {
            const airport = allAirports.value.find((x) => x.id === airportId)
            return airport ? `${airport.city} (${airport.code})` : airportId
        }

        const deleteRoute = (id) => routeApi.deleteRoute(id).then(refreshAll)
        const deleteFlight = (id) => flightApi.deleteFlight(id).then(fetchFlights)
        const goFlight = (id) => router.push({ name: 'AdminFlight', params: { id } })
        const openEdit = () => modal.value.open()

        const refreshAll = () =>
            Promise.all([fetchAirport(), fetchRoutes()]).then(fetchFlights)

        const handleUpdateAirport = (updatedAirport) =>
            airportApi
                .updateAirport(updatedAirport.id, updatedAirport)
                .then(refreshAll)
                .catch((err) => alert(err.response.data))

        onMounted(() =>
            Promise.all([fetchAirport(), fetchAllAirports(), fetchRoutes()]).then(fetchFlights)
        )

        return {
            airport,
            routes,
            departures,
            arrivals,
            modal,
            label,
            deleteRoute,
            deleteFlight,
            goFlight,
            openEdit,
            handleUpdateAirport,
        }
    },
}
</script>

<style scoped></style>
