<template>
    <div>
        <Header />
        <div class="container mt-4">
            <h1>Информация о маршруте</h1>

            <div v-if="routeData && airports.length" class="mt-3">
                <p>
                    <strong>Аэропорт отправления:</strong>
                    {{ getAirportLabel(routeData.departure_airport_id) }}
                </p>
                <p>
                    <strong>Аэропорт назначения:</strong>
                    {{ getAirportLabel(routeData.arrival_airport_id) }}
                </p>
                <p><strong>Расстояние:</strong> {{ routeData.distance }} км</p>
                <p><strong>Длительность:</strong> {{ routeData.duration_minutes }} минут</p>
                <button class="btn btn-secondary" @click="openEditRouteModal">Редактировать</button>
            </div>

            <hr class="my-4" />

            <h2>Рейсы для этого маршрута</h2>
            <div v-if="flights.length" class="row mt-3">
                <div v-for="flight in flights" :key="flight.id" class="col-md-4 mb-3">
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">Рейс {{ flight.flight_number }}</h5>
                            <p class="card-text">
                                Время вылета: {{ flight.departure_time }}<br />
                                Время прибытия: {{ flight.arrival_time }}<br />
                                Статус: {{ flight.status }}
                            </p>
                        </div>
                    </div>
                </div>
            </div>

            <RouteModal
                ref="routeModal"
                :initialRoute="routeData"
                @updateRoute="handleUpdateRoute"
            />
        </div>
    </div>
</template>

<script>
import { ref, onMounted, defineComponent } from 'vue'
import { useRoute } from 'vue-router'
import Header from '@/components/Header.vue'
import RouteModal from '@/components/RouteModal.vue'
import routeApi from '@/API/routeApi'
import flightApi from '@/API/flightApi'
import airportApi from '@/API/airportApi'

export default defineComponent({
    name: 'AdminRoute',
    components: {
        Header,
        RouteModal,
    },
    setup() {
        const route = useRoute()
        const routeId = parseInt(route.params.id, 10)

        const routeData = ref(null)
        const flights = ref([])
        const airports = ref([])

        const fetchAirports = () => {
            return airportApi
                .getAirports()
                .then((res) => (airports.value = res.data))
                .catch((err) => console.error('Ошибка получения аэропортов', err))
        }

        const fetchRoute = () => {
            return routeApi
                .getRoutes()
                .then((res) => {
                    routeData.value = res.data.find((r) => r.id === routeId) || null
                })
                .catch((err) => console.error('Ошибка получения маршрута', err))
        }

        const fetchFlights = () => {
            flightApi
                .getFlightsByRoute(routeId)
                .then((res) => (flights.value = res.data))
                .catch((err) => console.error('Ошибка получения рейсов', err))
        }

        const getAirportLabel = (id) => {
            const a = airports.value.find((x) => x.id === id)
            return a ? `${a.city} (${a.code})` : id
        }

        const routeModal = ref(null)
        const openEditRouteModal = () => {
            routeModal.value.open()
        }

        const handleUpdateRoute = (updatedRoute) => {
            routeApi
                .updateRoute(updatedRoute.id, updatedRoute)
                .then(fetchRoute)
                .catch((err) => console.error('Ошибка обновления маршрута', err))
        }

        onMounted(async () => {
            await fetchAirports()
            await fetchRoute()
            fetchFlights()
        })

        return {
            routeData,
            flights,
            airports,
            getAirportLabel,
            openEditRouteModal,
            handleUpdateRoute,
            routeModal,
        }
    },
})
</script>

<style scoped></style>
