<template>
    <div>
        <Header />
        <div class="container mt-4">
            <router-link to="/admin" class="back-link">
                <span class="back-arrow">&#8592;</span> Назад
            </router-link>
            <h1>Управление маршрутами</h1>
            <div class="mt-4">
                <div class="d-flex justify-content-between align-items-center">
                    <h2>Маршруты</h2>
                    <button class="btn btn-primary" @click="openCreateRouteModal">
                        Создать маршрут
                    </button>
                </div>
                <div class="row mt-3">
                    <div v-for="route in routes" :key="route.id" class="col-md-4 mb-3">
                        <div class="card" @click="goToRoute(route.id)">
                            <div class="card-body">
                                <h5 class="card-title">
                                    Маршрут {{ getAirportLabel(route.departure_airport_id) }} —
                                    {{ getAirportLabel(route.arrival_airport_id) }}
                                </h5>
                                <p class="card-text">
                                    Расстояние: {{ route.distance }} км<br />
                                    Время: {{ route.duration_minutes }} мин
                                </p>
                                <button
                                    class="btn btn-danger btn-sm float-end"
                                    @click.stop="deleteRoute(route.id)"
                                >
                                    ×
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <CreateRouteModal ref="createRouteModal" @createRoute="handleCreateRoute" />
        </div>
    </div>
</template>

<script>
import { ref, onMounted, defineComponent } from 'vue'
import { useRouter } from 'vue-router'
import Header from '@/components/Header.vue'
import CreateRouteModal from '@/components/RouteModal.vue'
import routeApi from '@/API/routeApi'
import airportApi from '@/API/airportApi'

export default defineComponent({
    name: 'AdminRoutes',
    components: { Header, CreateRouteModal },
    setup() {
        const routes = ref([])
        const airports = ref([])
        const router = useRouter()

        const fetchRoutes = () => {
            routeApi
                .getRoutes()
                .then((res) => (routes.value = res.data))
                .catch((err) => console.error('Ошибка получения маршрутов', err))
        }
        const fetchAirports = () => {
            airportApi
                .getAirports()
                .then((res) => (airports.value = res.data))
                .catch((err) => console.error('Ошибка получения аэропортов', err))
        }

        const deleteRoute = (id) => {
            routeApi
                .deleteRoute(id)
                .then(fetchRoutes)
                .catch((err) => console.error('Ошибка удаления маршрута', err))
        }

        const openCreateRouteModal = () => createRouteModal.value.open()
        const handleCreateRoute = (newRoute) => {
            routeApi
                .createRoute(newRoute)
                .then(() => {
                    fetchRoutes()
                })
                .catch((err) => console.error('Ошибка создания маршрута', err))
        }

        const getAirportLabel = (id) => {
            const a = airports.value.find((x) => x.id === id)
            return a ? `${a.city} (${a.code})` : id
        }

        const goToRoute = (routeId) => {
            router.push({ name: 'AdminRoute', params: { id: routeId } })
        }

        const createRouteModal = ref(null)

        onMounted(() => {
            fetchAirports()
            fetchRoutes()
        })

        return {
            routes,
            airports,
            createRouteModal,
            openCreateRouteModal,
            handleCreateRoute,
            deleteRoute,
            getAirportLabel,
            goToRoute,
        }
    },
})
</script>
