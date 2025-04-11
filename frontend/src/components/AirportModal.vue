<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">
                        {{ isEditMode ? 'Редактировать аэропорт' : 'Создать аэропорт' }}
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
                            <label for="name" class="form-label">Название</label>
                            <input
                                type="text"
                                id="name"
                                v-model="airport.name"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="code" class="form-label">Код</label>
                            <input
                                type="text"
                                id="code"
                                v-model="airport.code"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="city" class="form-label">Город</label>
                            <input
                                type="text"
                                id="city"
                                v-model="airport.city"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="country" class="form-label">Страна</label>
                            <input
                                type="text"
                                id="country"
                                v-model="airport.country"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="timezone" class="form-label">Часовой пояс</label>
                            <input
                                type="text"
                                id="timezone"
                                v-model="airport.timezone"
                                class="form-control"
                                required
                            />
                        </div>
                        <button type="submit" class="btn btn-primary">
                            {{ isEditMode ? 'Изменить' : 'Создать' }}
                        </button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { ref, computed, defineComponent, watch } from 'vue'
import { Modal } from 'bootstrap'

export default defineComponent({
    name: 'AirportModal',
    props: {
        initialAirport: {
            type: Object,
            default: null,
        },
    },
    emits: ['createAirport', 'updateAirport'],
    setup(props, { emit }) {
        const airport = ref({
            id: null,
            name: '',
            code: '',
            city: '',
            country: '',
            timezone: '',
        })

        const resetForm = () => {
            airport.value = {
                id: null,
                name: '',
                code: '',
                city: '',
                country: '',
                timezone: '',
            }
        }

        const isEditMode = computed(() => airport.value.id !== null)

        watch(
            () => props.initialAirport,
            (newVal) => {
                if (newVal) {
                    airport.value = { ...newVal }
                } else {
                    resetForm()
                }
            },
            { immediate: true },
        )

        const modalElement = ref(null)
        let modalInstance = null

        const open = () => {
            if (!modalInstance) {
                modalInstance = new Modal(modalElement.value)
            }
            modalInstance.show()
        }

        const close = () => {
            if (modalInstance) {
                modalInstance.hide()
            }
        }

        const submitForm = () => {
            if (isEditMode.value) {
                emit('updateAirport', { ...airport.value })
            } else {
                emit('createAirport', { ...airport.value })
            }
            resetForm()
            close()
        }

        return {
            modalElement,
            airport,
            isEditMode,
            open,
            close,
            submitForm,
        }
    },
})
</script>

<style scoped></style>
