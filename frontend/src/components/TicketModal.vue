<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">
                        {{ isEditMode ? 'Редактировать билет' : 'Добавить билет' }}
                    </h5>
                    <button type="button" class="btn-close" aria-label="Close"
                            @click="close"></button>
                </div>
                <div class="modal-body">
                    <form @submit.prevent="submitForm">
                        <div class="mb-3">
                            <label for="seatNumber" class="form-label">Место</label>
                            <input
                                type="text"
                                id="seatNumber"
                                v-model="ticket.seat_number"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="price" class="form-label">Цена</label>
                            <input
                                type="number"
                                id="price"
                                v-model.number="ticket.price"
                                class="form-control"
                                required
                            />
                        </div>
                        <button type="submit" class="btn btn-primary">
                            {{ isEditMode ? 'Сохранить' : 'Добавить' }}
                        </button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { ref, computed, watch } from 'vue'
import { Modal } from 'bootstrap'

export default {
    name: 'TicketModal',
    props: {
        flightId: {
            type: Number,
            required: true
        },
        initialTicket: {
            type: Object,
            default: null
        }
    },
    emits: ['createTicket', 'updateTicket'],
    setup(props, { emit }) {
        const ticket = ref({
            id: null,
            flight_id: props.flightId,
            seat_number: '',
            price: 0
        })
        const modalElement = ref(null)
        let modalInstance = null

        const isEditMode = computed(() => ticket.value.id !== null)

        watch(
            () => props.initialTicket,
            (newVal) => {
                if (newVal && newVal.id != null) {
                    ticket.value = {
                        id: newVal.id,
                        flight_id: props.flightId,
                        seat_number: newVal.seat_number,
                        price: newVal.price
                    }
                } else {
                    ticket.value = {
                        id: null,
                        flight_id: props.flightId,
                        seat_number: '',
                        price: 0
                    }
                }
            },
            { immediate: true }
        )

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
            if (isEditMode.value) {
                emit('updateTicket', { ...ticket.value })
            } else {
                emit('createTicket', { ...ticket.value })
            }
            close()
        }

        return {
            ticket,
            modalElement,
            open,
            close,
            submitForm,
            isEditMode
        }
    }
}
</script>

<style scoped></style>
