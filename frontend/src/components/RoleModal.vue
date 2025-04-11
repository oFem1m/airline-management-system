<template>
    <div class="modal fade" ref="modalElement" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">
                        {{ isEditMode ? 'Редактировать роль' : 'Добавить роль' }}
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
                                v-model="role.name"
                                class="form-control"
                                required
                            />
                        </div>
                        <div class="mb-3">
                            <label for="description" class="form-label">Описание</label>
                            <textarea
                                id="description"
                                v-model="role.description"
                                class="form-control"
                                rows="3"
                                required
                            ></textarea>
                        </div>
                        <button type="submit" class="btn btn-primary">
                            {{ isEditMode ? 'Изменить' : 'Добавить' }}
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
    name: 'RoleModal',
    props: {
        initialRole: {
            type: Object,
            default: null,
        },
    },
    emits: ['createRole', 'updateRole'],
    setup(props, { emit }) {
        const role = ref({
            id: null,
            name: '',
            description: '',
        })

        const resetForm = () => {
            role.value = {
                id: null,
                name: '',
                description: '',
            }
        }

        const isEditMode = computed(() => role.value.id !== null)

        watch(
            () => props.initialRole,
            (newVal) => {
                if (newVal) {
                    role.value = { ...newVal }
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
                emit('updateRole', { ...role.value })
            } else {
                emit('createRole', { ...role.value })
            }
            resetForm()
            close()
        }

        return {
            modalElement,
            role,
            isEditMode,
            open,
            close,
            submitForm,
        }
    },
})
</script>

<style scoped></style>
