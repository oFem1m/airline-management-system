<template>
    <div>
        <Header />

        <div class="container mt-4">
            <router-link to="/admin" class="back-link">
                <span class="back-arrow">&#8592;</span> Назад
            </router-link>
            <h1>Управление пассажирами</h1>

            <div class="d-flex justify-content-between align-items-center my-4">
                <h2>Пассажиры</h2>
                <button class="btn btn-primary" @click="openCreateModal">
                    Добавить пассажира
                </button>
            </div>

            <div class="row">
                <div
                    v-for="p in passengers"
                    :key="p.id"
                    class="col-md-4 mb-3"
                    style="cursor: pointer"
                    @click="goToPassenger(p.id)"
                >
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">{{ p.first_name }} {{ p.last_name }}</h5>
                            <p class="card-text">
                                Email: {{ p.email }}<br />
                                Телефон: {{ p.phone }}<br />
                                Паспорт: {{ p.passport_number }}
                            </p>
                            <button
                                class="btn btn-danger btn-sm float-end"
                                @click.stop="deletePassenger(p.id)"
                            >
                                ×
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <PassengerModal
            ref="passengerModal"
            @createPassenger="handleCreate"
        />
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import Header from '@/components/Header.vue'
import PassengerModal from '@/components/PassengerModal.vue'
import passengerApi from '@/API/passengerApi'

export default {
    name: 'AdminPassengers',
    components: { Header, PassengerModal },
    setup() {
        const router = useRouter()
        const passengers = ref([])
        const passengerModal = ref(null)

        const fetchPassengers = () => {
            passengerApi
                .getPassengers()
                .then((res) => { passengers.value = res.data })
                .catch((err) => alert(err.response.data))
        }

        const goToPassenger = (id) => {
            router.push({ name: 'AdminPassenger', params: { id } })
        }

        const openCreateModal = () => {
            passengerModal.value.open()
        }

        const handleCreate = (newP) => {
            passengerApi
                .addPassenger(newP)
                .then(fetchPassengers)
                .catch((err) => alert(err.response.data))
        }

        const deletePassenger = (id) => {
            passengerApi
                .deletePassenger(id)
                .then(fetchPassengers)
                .catch((err) => alert(err.response.data))
        }

        onMounted(fetchPassengers)

        return {
            passengers,
            passengerModal,
            goToPassenger,
            openCreateModal,
            handleCreate,
            deletePassenger,
        }
    },
}
</script>

<style scoped></style>
