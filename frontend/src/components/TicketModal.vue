<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">Добавить билет</h5>
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
                                v-model="ticket.price"
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
import { ref } from 'vue'
import { Modal } from 'bootstrap'

export default {
    name: 'TicketModal',
    props: {
        flightId: {
            type: Number,
            required: true,
        },
    },
    emits: ['addTicket'],
    setup(props, { emit }) {
        const ticket = ref({
            seat_number: '',
            price: 0,
        })
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
            emit('addTicket', { flightId: props.flightId, ...ticket.value })
            close()
        }

        return {
            ticket,
            open,
            close,
            submitForm,
        }
    },
}
</script>

<style scoped></style>
