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

            <!-- Блок маршрутов -->
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

            <!-- Блок вылетов -->
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

            <!-- Блок прилётов -->
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

            <!-- Модальное редактирование -->
            <AirportModal ref="modal" :initialAirport="airport" @updateAirport="fetchAirport" />
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
        const id = parseInt(route.params.id, 10)

        const airport = ref(null)
        const routes = ref([])
        const departures = ref([])
        const arrivals = ref([])
        const allAirports = ref([])
        const modal = ref(null)

        const fetchAirport = () => {
            airportApi
                .getAirport(id)
                .then((res) => (airport.value = res.data))
                .catch((err) => console.error(err))
        }

        const fetchRoutes = () => {
            routeApi
                .getRoutesByAirport(id)
                .then((res) => (routes.value = res.data))
                .catch((err) => console.error(err))
        }

        const fetchFlights = () => {
            flightApi
                .getFlightsByAirport(id)
                .then((res) => {
                    const flights = res.data
                    return Promise.all(
                        flights.map((f) =>
                            routeApi
                                .getRoute(f.route_id)
                                .then((r) => ({ flight: f, route: r.data })),
                        ),
                    )
                })
                .then((pairs) => {
                    const mapped = pairs.map(({ flight, route }) => ({ ...flight, route }))
                    departures.value = mapped.filter(
                        (item) => item.route.departure_airport_id === id,
                    )
                    arrivals.value = mapped.filter((item) => item.route.arrival_airport_id === id)
                })
                .catch((err) => console.error('Ошибка получения рейсов по аэропорту', err))
        }

        const fetchAllAirports = () => {
            airportApi
                .getAirports()
                .then((res) => (allAirports.value = res.data))
                .catch((err) => console.error(err))
        }
        const label = (aid) => {
            const a = allAirports.value.find((x) => x.id === aid)
            return a ? `${a.city} (${a.code})` : aid
        }

        const deleteRoute = (rid) => {
            routeApi.deleteRoute(rid).then(fetchRoutes)
        }
        const deleteFlight = (fid) => {
            flightApi.deleteFlight(fid).then(fetchFlights)
        }
        const goFlight = (fid) => {
            router.push({ name: 'AdminFlight', params: { id: fid } })
        }

        const openEdit = () => {
            modal.value.open()
        }

        onMounted(() => {
            fetchAirport()
            fetchRoutes()
            fetchFlights()
            fetchAllAirports()
        })

        return {
            airport,
            routes,
            departures,
            arrivals,
            modal,
            fetchAirport,
            deleteRoute,
            deleteFlight,
            goFlight,
            openEdit,
            label,
        }
    },
}
</script>

<style scoped></style>
