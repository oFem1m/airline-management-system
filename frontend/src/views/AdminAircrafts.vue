<template>
    <div>
        <Header />

        <div class="container mt-4">
            <router-link to="/admin" class="back-link">
                <span class="back-arrow">&#8592;</span> Назад
            </router-link>

            <h1>Управление самолётами</h1>

            <div class="mt-4">
                <div class="d-flex justify-content-between align-items-center">
                    <h2>Самолёты</h2>
                    <button class="btn btn-primary" @click="openCreateAircraftModal">
                        Добавить самолёт
                    </button>
                </div>

                <div class="row mt-3">
                    <div v-for="air in aircrafts" :key="air.id" class="col-md-4 mb-3">
                        <div class="card" style="cursor: pointer" @click="goToAircraft(air.id)">
                            <div class="card-body">
                                <h5 class="card-title">{{ air.tail_number }}</h5>
                                <p class="card-text">
                                    Модель: {{ air.model }}<br />
                                    Вместимость: {{ air.capacity }}<br />
                                    Год выпуска: {{ air.manufacture_year }}
                                </p>
                                <button
                                    class="btn btn-danger btn-sm float-end"
                                    @click.stop="deleteAircraft(air.id)"
                                >
                                    ×
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <CreateAircraftModal ref="createAircraftModal" @createAircraft="handleCreateAircraft" />
        </div>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

import Header from '@/components/Header.vue'
import CreateAircraftModal from '@/components/CreateAircraftModal.vue'
import aircraftApi from '@/API/aircraftApi'

export default {
    name: 'AdminAircrafts',
    components: { Header, CreateAircraftModal },
    setup() {
        const router = useRouter()
        const aircrafts = ref([])
        const createAircraftModal = ref(null)

        const fetchAircrafts = () => {
            aircraftApi
                .getAircrafts()
                .then((res) => (aircrafts.value = res.data))
                .catch((err) => console.error('Ошибка получения самолётов', err))
        }

        const goToAircraft = (id) => {
            router.push({ name: 'AdminAircraft', params: { id } })
        }

        const deleteAircraft = (id) => {
            aircraftApi
                .deleteAircraft(id)
                .then(fetchAircrafts)
                .catch((err) => console.error('Ошибка удаления самолёта', err))
        }

        const handleCreateAircraft = () => {
            fetchAircrafts()
        }

        const openCreateAircraftModal = () => {
            createAircraftModal.value.open()
        }

        onMounted(() => {
            fetchAircrafts()
        })

        return {
            aircrafts,
            createAircraftModal,
            goToAircraft,
            deleteAircraft,
            handleCreateAircraft,
            openCreateAircraftModal,
        }
    },
}
</script>

<style scoped></style>
