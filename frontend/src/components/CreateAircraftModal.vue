<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">Добавить самолет</h5>
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
                            <label for="tail_number" class="form-label">Бортовой номер</label>
                            <input
                                type="text"
                                id="tail_number"
                                v-model="aircraft.tail_number"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="model" class="form-label">Модель</label>
                            <input
                                type="text"
                                id="model"
                                v-model="aircraft.model"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="capacity" class="form-label">Вместимость</label>
                            <input
                                type="number"
                                id="capacity"
                                v-model="aircraft.capacity"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="manufacture_year" class="form-label">Год выпуска</label>
                            <input
                                type="number"
                                id="manufacture_year"
                                v-model="aircraft.manufacture_year"
                                class="form-control"
                                required
                            />
                        </div>
                        <button type="submit" class="btn btn-primary">Добавить</button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { ref, defineComponent } from 'vue'
import { Modal } from 'bootstrap'

export default defineComponent({
    name: 'CreateAircraftModal',
    emits: ['createAircraft'],
    setup(props, { emit }) {
        const modalElement = ref(null)
        let modalInstance = null

        const aircraft = ref({
            tail_number: '',
            model: '',
            capacity: '',
            manufacture_year: '',
        })

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

        const resetForm = () => {
            aircraft.value = {
                tail_number: '',
                model: '',
                capacity: '',
                manufacture_year: '',
            }
        }

        const submitForm = () => {
            emit('createAircraft', { ...aircraft.value })
            resetForm()
            close()
        }

        return {
            modalElement,
            aircraft,
            open,
            close,
            submitForm,
        }
    },
})
</script>

<style scoped>

</style>
