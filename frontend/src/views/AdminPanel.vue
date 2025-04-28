<template>
    <div>
        <Header />
        <div class="container mt-4">
            <h1>Панель управления</h1>

            <div class="mt-4">
                <div class="d-flex justify-content-between align-items-center">
                    <h2>Самолёты</h2>
                    <router-link to="/admin/aircrafts" class="btn btn-secondary">
                        Управление самолётами
                    </router-link>
                </div>
            </div>

            <hr class="my-4" />

            <div class="mt-4">
                <div class="d-flex justify-content-between align-items-center">
                    <h2>Аэропорты</h2>
                    <router-link to="/admin/airports" class="btn btn-secondary">
                        Управление аэропортами
                    </router-link>
                </div>
            </div>

            <hr class="my-4" />

            <div class="mt-4">
                <div class="d-flex justify-content-between align-items-center">
                    <h2>Персонал</h2>
                    <router-link to="/admin/staff" class="btn btn-secondary">
                        Управление персоналом
                    </router-link>
                </div>
            </div>

            <hr class="my-4" />

            <div class="mt-4">
                <div class="d-flex justify-content-between align-items-center">
                    <h2>Маршруты</h2>
                    <router-link to="/admin/routes" class="btn btn-secondary">
                        Управление маршрутами
                    </router-link>
                </div>
            </div>

            <hr class="my-4" />

            <div class="d-flex justify-content-between align-items-center">
                <h2>Пассажиры</h2>
                <router-link to="/admin/passengers" class="btn btn-secondary">
                    Управление пассажирами
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import Header from '@/components/Header.vue'

export default {
    name: 'AdminPanel',
    components: { Header },
}
</script>

<style scoped></style>
