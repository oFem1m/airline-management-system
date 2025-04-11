<template>
    <div>
        <Header />
        <div class="container mt-4">
            <h1>Управление персоналом</h1>

            <div class="mt-4">
                <div class="d-flex justify-content-between align-items-center">
                    <h2>Сотрудники</h2>
                    <button class="btn btn-primary" @click="openCreateEmployeeModal">
                        Добавить сотрудника
                    </button>
                </div>
                <div class="row mt-3">
                    <div v-for="emp in employees" :key="emp.id" class="col-md-4 mb-3">
                        <div
                            class="card"
                            style="cursor: pointer"
                            @click="openEditEmployeeModal(emp)"
                        >
                            <div class="card-body">
                                <h5 class="card-title">{{ emp.first_name }} {{ emp.last_name }}</h5>
                                <p class="card-text">
                                    Email: {{ emp.email }}<br />
                                    Телефон: {{ emp.phone }}<br />
                                    Дата найма: {{ formatDate(emp.hire_date) }}<br />
                                    Зарплата: {{ emp.salary }}<br />
                                    Роль (ID): {{ emp.role_id }}
                                </p>
                                <button
                                    class="btn btn-danger btn-sm float-end"
                                    @click.stop="deleteEmployee(emp.id)"
                                >
                                    ×
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <hr class="my-4" />

            <div class="mt-4">
                <div class="d-flex justify-content-between align-items-center">
                    <h2>Роли</h2>
                    <button class="btn btn-primary" @click="openCreateRoleModal">
                        Добавить роль
                    </button>
                </div>
                <div class="row mt-3">
                    <div v-for="role in roles" :key="role.id" class="col-md-4 mb-3">
                        <div class="card" style="cursor: pointer" @click="openEditRoleModal(role)">
                            <div class="card-body">
                                <h5 class="card-title">{{ role.name }}</h5>
                                <p class="card-text">
                                    {{ role.description }}
                                </p>
                                <button
                                    class="btn btn-danger btn-sm float-end"
                                    @click.stop="deleteRole(role.id)"
                                >
                                    ×
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <EmployeeModal
            ref="employeeModal"
            :initialEmployee="selectedEmployee"
            @createEmployee="handleCreateEmployee"
            @updateEmployee="handleUpdateEmployee"
        />
        <RoleModal
            ref="roleModal"
            :initialRole="selectedRole"
            @createRole="handleCreateRole"
            @updateRole="handleUpdateRole"
        />
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import Header from '@/components/Header.vue'
import EmployeeModal from '@/components/EmployeeModal.vue'
import RoleModal from '@/components/RoleModal.vue'
import employeeApi from '@/API/employeeApi'
import roleApi from '@/API/roleApi'

export default {
    name: 'AdminStaff',
    components: {
        Header,
        EmployeeModal,
        RoleModal,
    },
    setup() {
        const employees = ref([])
        const roles = ref([])

        const selectedEmployee = ref(null)
        const selectedRole = ref(null)

        const employeeModal = ref(null)
        const roleModal = ref(null)

        const fetchEmployees = () => {
            employeeApi
                .getEmployees()
                .then((response) => {
                    employees.value = response.data
                })
                .catch((error) => {
                    console.error('Ошибка получения сотрудников', error)
                })
        }

        const fetchRoles = () => {
            roleApi
                .getRoles()
                .then((response) => {
                    roles.value = response.data
                })
                .catch((error) => {
                    console.error('Ошибка получения ролей', error)
                })
        }

        const deleteEmployee = (id) => {
            employeeApi
                .deleteEmployee(id)
                .then(() => {
                    fetchEmployees()
                })
                .catch((error) => {
                    console.error('Ошибка при удалении сотрудника', error)
                })
        }

        const deleteRole = (id) => {
            roleApi
                .deleteRole(id)
                .then(() => {
                    fetchRoles()
                })
                .catch((error) => {
                    console.error('Ошибка при удалении роли', error)
                })
        }

        const openCreateEmployeeModal = () => {
            selectedEmployee.value = null
            employeeModal.value.open()
        }

        const openEditEmployeeModal = (emp) => {
            selectedEmployee.value = { ...emp }
            employeeModal.value.open()
        }

        const handleCreateEmployee = (newEmployee) => {
            employeeApi
                .createEmployee(newEmployee)
                .then(() => {
                    fetchEmployees()
                })
                .catch((error) => {
                    console.error('Ошибка создания сотрудника', error)
                })
        }

        const handleUpdateEmployee = (employeeData) => {
            employeeApi
                .updateEmployee(employeeData.id, employeeData)
                .then(() => {
                    fetchEmployees()
                })
                .catch((error) => {
                    console.error('Ошибка обновления сотрудника', error)
                })
        }

        const openCreateRoleModal = () => {
            selectedRole.value = null
            roleModal.value.open()
        }

        const openEditRoleModal = (rl) => {
            selectedRole.value = { ...rl }
            roleModal.value.open()
        }

        const handleCreateRole = (newRole) => {
            roleApi
                .createRole(newRole)
                .then(() => {
                    fetchRoles()
                })
                .catch((error) => {
                    console.error('Ошибка создания роли', error)
                })
        }

        const handleUpdateRole = (roleData) => {
            roleApi
                .updateRole(roleData.id, roleData)
                .then(() => {
                    fetchRoles()
                })
                .catch((error) => {
                    console.error('Ошибка обновления роли', error)
                })
        }

        const formatDate = (isoStr) => {
            if (!isoStr) return ''
            const date = new Date(isoStr)
            return date.toLocaleString()
        }

        onMounted(() => {
            fetchEmployees()
            fetchRoles()
        })

        return {
            employees,
            roles,
            selectedEmployee,
            selectedRole,
            openCreateEmployeeModal,
            openEditEmployeeModal,
            handleCreateEmployee,
            handleUpdateEmployee,
            deleteEmployee,
            openCreateRoleModal,
            openEditRoleModal,
            handleCreateRole,
            handleUpdateRole,
            deleteRole,
            employeeModal,
            roleModal,
            formatDate,
        }
    },
}
</script>

<style scoped></style>
