<template>
    <div>
        <Header />
        <div class="container mt-4">
            <h1>Панель управления</h1>

            <div class="mt-4">
                <div class="d-flex justify-content-between align-items-center">
                    <h2>Самолеты</h2>
                    <button class="btn btn-primary" @click="openModal">Создать самолет</button>
                </div>
                <div class="mt-2">
                    <input type="checkbox" v-model="showAircrafts" id="toggleAircrafts" />
                    <label for="toggleAircrafts" class="ms-2">Показать список самолетов</label>
                </div>
                <div v-if="showAircrafts" class="mt-3">
                    <div class="row">
                        <div v-for="aircraft in aircrafts" :key="aircraft.id" class="col-md-4 mb-3">
                            <div class="card">
                                <div class="card-body">
                                    <h5 class="card-title">{{ aircraft.tail_number }}</h5>
                                    <p class="card-text">
                                        Модель: {{ aircraft.model }}<br />
                                        Вместимость: {{ aircraft.capacity }}<br />
                                        Год выпуска: {{ aircraft.manufacture_year }}
                                    </p>
                                    <button
                                        class="btn btn-danger btn-sm float-end"
                                        @click="deleteAircraft(aircraft.id)"
                                    >
                                        ×
                                    </button>
                                </div>
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
import Header from '@/components/Header.vue'
import CreateAircraftModal from '@/components/CreateAircraftModal.vue'
import aircraftApi from '@/API/aircraftApi'

export default {
    name: 'AdminPanel',
    components: {
        Header,
        CreateAircraftModal,
    },
    setup() {
        const showAircrafts = ref(false)
        const aircrafts = ref([])

        const fetchAircrafts = () => {
            aircraftApi
                .getAircrafts()
                .then((response) => {
                    aircrafts.value = response.data
                })
                .catch((error) => {
                    console.error('Ошибка получения самолетов', error)
                })
        }

        const handleCreateAircraft = (newAircraftData) => {
            aircraftApi
                .createAircraft(newAircraftData)
                .then(() => {
                    fetchAircrafts()
                })
                .catch((error) => {
                    console.error('Ошибка создания самолета', error)
                })
        }

        const deleteAircraft = (id) => {
            aircraftApi
                .deleteAircraft(id)
                .then(() => {
                    fetchAircrafts()
                })
                .catch((error) => {
                    console.error('Ошибка удаления самолета', error)
                })
        }

        const createAircraftModal = ref(null)

        const openModal = () => {
            createAircraftModal.value.open()
        }

        onMounted(() => {
            fetchAircrafts()
        })

        return {
            showAircrafts,
            aircrafts,
            openModal,
            deleteAircraft,
            handleCreateAircraft,
            createAircraftModal,
        }
    },
}
</script>

<style scoped>

</style>
