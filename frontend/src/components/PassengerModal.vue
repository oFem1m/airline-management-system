<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">
                        {{ isEditMode ? 'Редактировать пассажира' : 'Добавить пассажира' }}
                    </h5>
                    <button type="button" class="btn-close" @click="close"></button>
                </div>
                <div class="modal-body">
                    <form @submit.prevent="submitForm">
                        <div class="mb-3">
                            <label for="first_name" class="form-label">Имя</label>
                            <input
                                id="first_name"
                                v-model="passenger.first_name"
                                type="text"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="last_name" class="form-label">Фамилия</label>
                            <input
                                id="last_name"
                                v-model="passenger.last_name"
                                type="text"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="email" class="form-label">Email</label>
                            <input
                                id="email"
                                v-model="passenger.email"
                                type="email"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="phone" class="form-label">Телефон</label>
                            <input
                                id="phone"
                                v-model="passenger.phone"
                                type="text"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="passport_number" class="form-label">Номер паспорта</label>
                            <input
                                id="passport_number"
                                v-model="passenger.passport_number"
                                type="text"
                                class="form-control"
                                required
                            />
                        </div>
                        <button type="submit" class="btn btn-primary">
                            {{ isEditMode ? 'Сохранить' : 'Создать' }}
                        </button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { ref, computed, watch, defineComponent } from 'vue'
import { Modal } from 'bootstrap'

export default defineComponent({
    name: 'PassengerModal',
    props: {
        initialPassenger: {
            type: Object,
            default: null,
        },
    },
    emits: ['createPassenger', 'updatePassenger'],
    setup(props, { emit }) {
        const passenger = ref({
            id: null,
            first_name: '',
            last_name: '',
            email: '',
            phone: '',
            passport_number: '',
        })

        const isEditMode = computed(() => passenger.value.id !== null)

        const resetForm = () => {
            passenger.value = {
                id: null,
                first_name: '',
                last_name: '',
                email: '',
                phone: '',
                passport_number: '',
            }
        }

        watch(
            () => props.initialPassenger,
            (val) => {
                if (val) {
                    passenger.value = { ...val }
                } else {
                    resetForm()
                }
            },
            { immediate: true }
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
            if (modalInstance) modalInstance.hide()
        }

        const submitForm = () => {
            const payload = { ...passenger.value }
            if (isEditMode.value) {
                emit('updatePassenger', payload)
            } else {
                emit('createPassenger', payload)
            }
            close()
            resetForm()
        }

        return { passenger, isEditMode, modalElement, open, close, submitForm }
    },
})
</script>

<style scoped></style>
