<template>
    <div>
        <Header />
        <div class="container mt-4">
            <h1>Панель управления</h1>

            <div class="mt-4">
                <div class="d-flex justify-content-between align-items-center">
                    <h2>Самолеты</h2>
                    <button class="btn btn-primary" @click="openAircraftModal">
                        Добавить самолет
                    </button>
                </div>
                <div class="mt-2">
                    <input type="checkbox" v-model="showAircrafts" id="toggleAircrafts" />
                    <label for="toggleAircrafts" class="ms-2">Показать список самолетов</label>
                </div>
                <div v-if="showAircrafts" class="mt-3">
                    <div class="row">
                        <div v-for="aircraft in aircrafts" :key="aircraft.id" class="col-md-4 mb-3">
                            <div
                                class="card"
                                style="cursor: pointer"
                                @click="goToAircraft(aircraft.id)"
                            >
                                <div class="card-body">
                                    <h5 class="card-title">{{ aircraft.tail_number }}</h5>
                                    <p class="card-text">
                                        Модель: {{ aircraft.model }}<br />
                                        Вместимость: {{ aircraft.capacity }}<br />
                                        Год выпуска: {{ aircraft.manufacture_year }}
                                    </p>
                                    <button
                                        class="btn btn-danger btn-sm float-end"
                                        @click.stop="deleteAircraft(aircraft.id)"
                                    >
                                        ×
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <hr class="my-4" />

            <div class="mt-4">
                <div class="d-flex justify-content-between align-items-center">
                    <h2>Аэропорты</h2>
                    <button class="btn btn-primary" @click="openCreateAirportModal">
                        Добавить аэропорт
                    </button>
                </div>
                <div class="mt-2">
                    <input type="checkbox" v-model="showAirports" id="toggleAirports" />
                    <label for="toggleAirports" class="ms-2">Показать список аэропортов</label>
                </div>
                <div v-if="showAirports" class="mt-3">
                    <div class="row">
                        <div v-for="airport in airports" :key="airport.id" class="col-md-4 mb-3">
                            <div
                                class="card"
                                @click="openEditAirportModal(airport)"
                                style="cursor: pointer"
                            >
                                <div class="card-body">
                                    <h5 class="card-title">{{ airport.name }}</h5>
                                    <p class="card-text">
                                        Код: {{ airport.code }}<br />
                                        Город: {{ airport.city }}<br />
                                        Страна: {{ airport.country }}<br />
                                        Часовой пояс: {{ airport.timezone }}
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
            </div>

            <CreateAircraftModal ref="createAircraftModal" @createAircraft="handleCreateAircraft" />
            <AirportModal
                ref="airportModal"
                :initialAirport="selectedAirport"
                @createAirport="handleCreateAirport"
                @updateAirport="handleUpdateAirport"
            />
        </div>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import Header from '@/components/Header.vue'
import CreateAircraftModal from '@/components/CreateAircraftModal.vue'
import AirportModal from '@/components/AirportModal.vue'
import aircraftApi from '@/API/aircraftApi'
import airportApi from '@/API/airportApi'
import { useRouter } from 'vue-router'

export default {
    name: 'AdminPanel',
    components: {
        Header,
        CreateAircraftModal,
        AirportModal,
    },
    setup() {
        const router = useRouter()
        const showAircrafts = ref(false)
        const aircrafts = ref([])

        const goToAircraft = (id) => {
            router.push(`/admin/aircraft/${id}`)
        }

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
        const openAircraftModal = () => {
            createAircraftModal.value.open()
        }

        const showAirports = ref(false)
        const airports = ref([])
        const selectedAirport = ref(null)

        const fetchAirports = () => {
            airportApi
                .getAirports()
                .then((response) => {
                    airports.value = response.data
                })
                .catch((error) => {
                    console.error('Ошибка получения аэропортов', error)
                })
        }

        const handleCreateAirport = (newAirportData) => {
            airportApi
                .createAirport(newAirportData)
                .then(() => {
                    fetchAirports()
                })
                .catch((error) => {
                    console.error('Ошибка создания аэропорта', error)
                })
        }

        const handleUpdateAirport = (airportData) => {
            airportApi
                .updateAirport(airportData.id, airportData)
                .then(() => {
                    fetchAirports()
                })
                .catch((error) => {
                    console.error('Ошибка обновления аэропорта', error)
                })
        }

        const deleteAirport = (id) => {
            airportApi
                .deleteAirport(id)
                .then(() => {
                    fetchAirports()
                })
                .catch((error) => {
                    console.error('Ошибка удаления аэропорта', error)
                })
        }

        const airportModal = ref(null)

        const openCreateAirportModal = () => {
            selectedAirport.value = null
            airportModal.value.open()
        }

        const openEditAirportModal = (airport) => {
            selectedAirport.value = { ...airport }
            airportModal.value.open()
        }

        onMounted(() => {
            fetchAircrafts()
            fetchAirports()
        })

        return {
            // Самолеты
            showAircrafts,
            aircrafts,
            openAircraftModal,
            deleteAircraft,
            handleCreateAircraft,
            createAircraftModal,
            // Аэропорты
            showAirports,
            airports,
            openCreateAirportModal,
            openEditAirportModal,
            deleteAirport,
            handleCreateAirport,
            handleUpdateAirport,
            airportModal,
            selectedAirport,
            goToAircraft,
        }
    },
}
</script>

<style scoped></style>
