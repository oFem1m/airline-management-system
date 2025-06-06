<template>
    <div>
        <Header />
        <div class="container mt-4">
            <router-link to="/admin" class="back-link">
                <span class="back-arrow">&#8592;</span> Назад
            </router-link>
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
                                    Должность: {{ getRoleName(emp.role_id) }}
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
                    <h2>Должности</h2>
                    <button class="btn btn-primary" @click="openCreateRoleModal">
                        Добавить должность
                    </button>
                </div>
                <div class="row mt-3">
                    <div v-for="rl in roles" :key="rl.id" class="col-md-4 mb-3">
                        <div class="card" style="cursor: pointer" @click="openEditRoleModal(rl)">
                            <div class="card-body">
                                <h5 class="card-title">{{ rl.name }}</h5>
                                <p class="card-text">
                                    {{ rl.description }}
                                </p>
                                <button
                                    class="btn btn-danger btn-sm float-end"
                                    @click.stop="deleteRole(rl.id)"
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
            :roles="roles"
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
                    alert(error.response.data)
                })
        }

        const fetchRoles = () => {
            roleApi
                .getRoles()
                .then((response) => {
                    roles.value = response.data
                })
                .catch((error) => {
                    alert(error.response.data)
                })
        }

        const deleteEmployee = (id) => {
            employeeApi
                .deleteEmployee(id)
                .then(() => {
                    fetchEmployees()
                })
                .catch((error) => {
                    alert(error.response.data)
                })
        }

        const deleteRole = (id) => {
            roleApi
                .deleteRole(id)
                .then(() => {
                    fetchRoles()
                })
                .catch((error) => {
                    alert(error.response.data)
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
                    alert(error.response.data)
                })
        }

        const handleUpdateEmployee = (employeeData) => {
            employeeApi
                .updateEmployee(employeeData.id, employeeData)
                .then(() => {
                    fetchEmployees()
                })
                .catch((error) => {
                    alert(error.response.data)
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
                    alert(error.response.data)
                })
        }

        const handleUpdateRole = (roleData) => {
            roleApi
                .updateRole(roleData.id, roleData)
                .then(() => {
                    fetchRoles()
                })
                .catch((error) => {
                    alert(error.response.data)
                })
        }

        const formatDate = (isoStr) => {
            if (!isoStr) return ''
            const date = new Date(isoStr)
            return date.toLocaleString()
        }

        const getRoleName = (roleId) => {
            const role = roles.value.find((r) => r.id === roleId)
            return role ? role.name : 'Не задана'
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
            getRoleName,
        }
    },
}
</script>

<style scoped></style>
