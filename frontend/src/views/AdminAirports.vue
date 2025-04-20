<template>
    <div>
        <Header />

        <div class="container mt-4">
            <router-link to="/admin" class="back-link">
                <span class="back-arrow">&#8592;</span> Назад
            </router-link>

            <h1>Управление аэропортами</h1>

            <div class="mt-4">
                <div class="d-flex justify-content-between align-items-center">
                    <h2>Аэропорты</h2>
                    <button class="btn btn-primary" @click="openCreateModal">
                        Добавить аэропорт
                    </button>
                </div>

                <div class="row mt-3">
                    <div v-for="airport in airports" :key="airport.id" class="col-md-4 mb-3">
                        <div class="card" style="cursor: pointer" @click="goToAirport(airport.id)">
                            <div class="card-body">
                                <h5 class="card-title">{{ airport.name }} ({{ airport.code }})</h5>
                                <p class="card-text">
                                    Город: {{ airport.city }}<br />
                                    Страна: {{ airport.country }}
                                </p>
                                <button
                                    class="btn btn-danger btn-sm float-end"
                                    @click.stop="deleteAirport(airport.id)"
                                >
                                    ×
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <AirportModal ref="modal" :initialAirport="null" @createAirport="fetchAirports" />
        </div>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import Header from '@/components/Header.vue'
import AirportModal from '@/components/AirportModal.vue'
import airportApi from '@/API/airportApi'

export default {
    name: 'AdminAirports',
    components: { Header, AirportModal },
    setup() {
        const router = useRouter()
        const airports = ref([])
        const modal = ref(null)

        const fetchAirports = () => {
            airportApi
                .getAirports()
                .then((res) => (airports.value = res.data))
                .catch((err) => console.error('Ошибка получения аэропортов', err))
        }

        const goToAirport = (airportId) => {
            router.push({ name: 'AdminAirport', params: { id: airportId } })
        }

        const deleteAirport = (id) => {
            airportApi
                .deleteAirport(id)
                .then(fetchAirports)
                .catch((err) => console.error('Ошибка удаления аэропорта', err))
        }

        const openCreateModal = () => {
            modal.value.open()
        }

        onMounted(() => {
            fetchAirports()
        })

        return {
            airports,
            modal,
            goToAirport,
            deleteAirport,
            openCreateModal,
            fetchAirports,
        }
    },
}
</script>

<style scoped></style>
