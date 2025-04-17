<template>
    <div>
        <Header />
        <div class="container mt-4">
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
                        <div class="card" style="cursor: pointer">
                            <div class="card-body">
                                <h5 class="card-title">
                                    Маршрут {{ route.departure_airport_id }} -
                                    {{ route.arrival_airport_id }}
                                </h5>
                                <p class="card-text">
                                    Расстояние: {{ route.distance }} км <br />
                                    Время в пути: {{ route.duration_minutes }} минут
                                </p>
                                <button
                                    class="btn btn-danger btn-sm float-end"
                                    @click="deleteRoute(route.id)"
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
import { ref, onMounted } from 'vue'
import Header from '@/components/Header.vue'
import CreateRouteModal from '@/components/CreateRouteModal.vue'
import routeApi from '@/API/routeApi'

export default {
    name: 'AdminRoutes',
    components: {
        Header,
        CreateRouteModal,
    },
    setup() {
        const routes = ref([])
        const fetchRoutes = () => {
            routeApi
                .getRoutes()
                .then((response) => {
                    routes.value = response.data
                })
                .catch((error) => {
                    console.error('Ошибка получения маршрутов', error)
                })
        }

        const deleteRoute = (id) => {
            routeApi
                .deleteRoute(id)
                .then(() => {
                    fetchRoutes()
                })
                .catch((error) => {
                    console.error('Ошибка удаления маршрута', error)
                })
        }

        const openCreateRouteModal = () => {
            createRouteModal.value.open()
        }

        const handleCreateRoute = (newRoute) => {
            routeApi
                .createRoute(newRoute)
                .then(() => {
                    fetchRoutes()
                })
                .catch((error) => {
                    console.error('Ошибка создания маршрута', error)
                })
        }

        const createRouteModal = ref(null)

        onMounted(() => {
            fetchRoutes()
        })

        return {
            routes,
            openCreateRouteModal,
            deleteRoute,
            handleCreateRoute,
            createRouteModal,
        }
    },
}
</script>
