<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">
                        {{ isEditMode ? 'Редактировать маршрут' : 'Создать маршрут' }}
                    </h5>
                    <button
                        type="button"
                        class="btn-close"
                        aria-label="Close"
                        @click="close"
                    ></button>
                </div>
                <div class="modal-body">
                    <form @submit.prevent="submitForm">
                        <div class="mb-3">
                            <label class="form-label">Аэропорт отправления</label>
                            <select
                                v-model="route.departure_airport_id"
                                class="form-control"
                                required
                            >
                                <option disabled value="">Выберите аэропорт</option>
                                <option
                                    v-for="airport in airports"
                                    :key="airport.id"
                                    :value="airport.id"
                                >
                                    {{ airport.city }} ({{ airport.code }})
                                </option>
                            </select>
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Аэропорт назначения</label>
                            <select
                                v-model="route.arrival_airport_id"
                                class="form-control"
                                required
                            >
                                <option disabled value="">Выберите аэропорт</option>
                                <option
                                    v-for="airport in airports"
                                    :key="airport.id"
                                    :value="airport.id"
                                >
                                    {{ airport.city }} ({{ airport.code }})
                                </option>
                            </select>
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Расстояние (км)</label>
                            <input
                                type="number"
                                v-model="route.distance"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Длительность (минут)</label>
                            <input
                                type="number"
                                v-model="route.duration_minutes"
                                class="form-control"
                                required
                            />
                        </div>
                        <button type="submit" class="btn btn-primary">
                            {{ isEditMode ? 'Сохранить изменения' : 'Создать' }}
                        </button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { ref, computed, onMounted, watch, defineComponent } from 'vue'
import { Modal } from 'bootstrap'
import airportApi from '@/API/airportApi'

export default defineComponent({
    name: 'RouteModal',
    props: {
        initialRoute: {
            type: Object,
            default: null,
        },
    },
    emits: ['createRoute', 'updateRoute'],
    setup(props, { emit }) {
        const airports = ref([])
        const route = ref({
            id: null,
            departure_airport_id: '',
            arrival_airport_id: '',
            distance: '',
            duration_minutes: '',
        })

        const isEditMode = computed(() => route.value.id !== null)

        const resetForm = () => {
            route.value = {
                id: null,
                departure_airport_id: '',
                arrival_airport_id: '',
                distance: '',
                duration_minutes: '',
            }
        }

        const submitForm = () => {
            if (isEditMode.value) {
                emit('updateRoute', { ...route.value })
            } else {
                emit('createRoute', { ...route.value })
            }
            resetForm()
            close()
        }

        const modalElement = ref(null)
        let modalInstance = null

        const open = () => {
            if (!modalInstance) {
                modalInstance = new Modal(modalElement.value)
            }
            modalInstance.show()
        }
        const close = () => modalInstance && modalInstance.hide()

        const fetchAirports = () => {
            airportApi
                .getAirports()
                .then((res) => {
                    airports.value = res.data
                })
                .catch((err) => {
                    console.error('Ошибка получения аэропортов', err)
                })
        }

        watch(
            () => props.initialRoute,
            (newVal) => {
                if (newVal) {
                    route.value = {
                        id: newVal.id,
                        departure_airport_id: newVal.departure_airport_id,
                        arrival_airport_id: newVal.arrival_airport_id,
                        distance: newVal.distance,
                        duration_minutes: newVal.duration_minutes,
                    }
                } else {
                    resetForm()
                }
            },
            { immediate: true },
        )

        onMounted(fetchAirports)

        return {
            airports,
            route,
            isEditMode,
            modalElement,
            open,
            close,
            submitForm,
        }
    },
})
</script>

<style scoped></style>
