<template>
    <div>
        <Header />
        <div class="container mt-4">
            <div class="card p-4 mb-4">
                <div class="row align-items-end">
                    <div class="col-md-4 position-relative">
                        <label class="form-label">Откуда</label>
                        <input
                            type="text"
                            class="form-control"
                            v-model="fromQuery"
                            @input="onFromInput"
                            placeholder="Город или аэропорт"
                        />
                        <ul
                            class="list-group position-absolute w-100"
                            v-if="showFromSuggestions"
                            style="z-index: 1000;"
                        >
                            <li
                                class="list-group-item list-group-item-action"
                                v-for="airport in filteredFromAirports"
                                :key="airport.id"
                                @click="selectFrom(airport)"
                            >
                                {{ airport.city }} ({{ airport.code }})
                            </li>
                        </ul>
                    </div>

                    <div class="col-auto text-center">
                        <button
                            class="btn btn-outline-secondary"
                            @click="swapFromTo"
                            title="Поменять местами"
                            style="margin-top: 30px;"
                        >
                            ↔
                        </button>
                    </div>

                    <div class="col-md-4 position-relative">
                        <label class="form-label">Куда</label>
                        <input
                            type="text"
                            class="form-control"
                            v-model="toQuery"
                            @input="onToInput"
                            placeholder="Город или аэропорт"
                        />
                        <ul
                            class="list-group position-absolute w-100"
                            v-if="showToSuggestions"
                            style="z-index: 1000;"
                        >
                            <li
                                class="list-group-item list-group-item-action"
                                v-for="airport in filteredToAirports"
                                :key="airport.id"
                                @click="selectTo(airport)"
                            >
                                {{ airport.city }} ({{ airport.code }})
                            </li>
                        </ul>
                    </div>

                    <div class="col-md-3">
                        <label class="form-label">Дата</label>
                        <input
                            type="date"
                            class="form-control"
                            v-model="date"
                        />
                    </div>

                    <div class="col-md-1">
                        <button
                            class="btn btn-primary w-100"
                            @click="searchFlights"
                            style="margin-top: 24px;"
                        >
                            Искать
                        </button>
                    </div>
                </div>
            </div>

            <div v-if="flights.length">
                <div
                    v-for="flight in flights"
                    :key="flight.id"
                    class="card mb-3"
                >
                    <div class="card-body">
                        <h5 class="card-title">Рейс {{ flight.flight_number }}</h5>
                        <p class="card-text">
                            Вылет: {{ formatDate(flight.departure_time) }}<br/>
                            Прилет: {{ formatDate(flight.arrival_time) }}<br/>
                            Статус: {{ flight.status }}
                        </p>
                    </div>
                </div>
            </div>
            <p v-else-if="searched" class="text-muted">Рейсы не найдены.</p>
        </div>
    </div>
</template>

<script>
import { ref, onMounted, computed } from 'vue'
import Header from '@/components/Header.vue'
import airportApi from '@/API/airportApi'
import routeApi from '@/API/routeApi'
import flightApi from '@/API/flightApi'

export default {
    name: 'Home',
    components: { Header },
    setup() {
        const airports = ref([])
        const fromQuery = ref('')
        const toQuery = ref('')
        const date = ref('')
        const fromAirport = ref(null)
        const toAirport = ref(null)
        const flights = ref([])
        const searched = ref(false)

        const showFromSuggestions = ref(false)
        const showToSuggestions = ref(false)

        onMounted(() => {
            airportApi.getAirports()
                .then(res => { airports.value = res.data })
                .catch(console.error)
        })

        const filteredFromAirports = computed(() =>
            airports.value.filter(a =>
                a.city.toLowerCase().includes(fromQuery.value.toLowerCase()) ||
                a.code.toLowerCase().includes(fromQuery.value.toLowerCase())
            ).slice(0, 5)
        )
        const filteredToAirports = computed(() =>
            airports.value.filter(a =>
                a.city.toLowerCase().includes(toQuery.value.toLowerCase()) ||
                a.code.toLowerCase().includes(toQuery.value.toLowerCase())
            ).slice(0, 5)
        )

        function onFromInput() {
            fromAirport.value = null
            showFromSuggestions.value = !!fromQuery.value
        }
        function onToInput() {
            toAirport.value = null
            showToSuggestions.value = !!toQuery.value
        }

        function selectFrom(a) {
            fromAirport.value = a
            fromQuery.value = `${a.city} (${a.code})`
            showFromSuggestions.value = false
        }
        function selectTo(a) {
            toAirport.value = a
            toQuery.value = `${a.city} (${a.code})`
            showToSuggestions.value = false
        }

        function swapFromTo() {
            const tmpA = fromAirport.value
            fromAirport.value = toAirport.value
            toAirport.value = tmpA
            const tmpQ = fromQuery.value
            fromQuery.value = toQuery.value
            toQuery.value = tmpQ
        }

        function searchFlights() {
            searched.value = true
            flights.value = []

            if (!fromAirport.value || !toAirport.value) return

            routeApi.getRoutes()
                .then(res => {
                    const validRoutes = res.data.filter(r =>
                        r.departure_airport_id === fromAirport.value.id &&
                        r.arrival_airport_id === toAirport.value.id
                    )
                    return Promise.all(validRoutes.map(r =>
                        flightApi.getFlightsByRoute(r.id)
                    ))
                })
                .then(responses => {
                    let all = []
                    responses.forEach(r => { all = all.concat(r.data) })
                    if (date.value) {
                        all = all.filter(f =>
                            f.departure_time.startsWith(date.value)
                        )
                    }
                    flights.value = all
                })
                .catch(console.error)
        }

        function formatDate(iso) {
            return new Date(iso).toLocaleString()
        }

        return {
            fromQuery, toQuery, date,
            flights, searched,
            showFromSuggestions, showToSuggestions,
            filteredFromAirports, filteredToAirports,
            onFromInput, onToInput,
            selectFrom, selectTo,
            swapFromTo, searchFlights,
            formatDate,
        }
    },
}
</script>

<style scoped>
.list-group {
    max-height: 150px;
    overflow-y: auto;
}
</style>
