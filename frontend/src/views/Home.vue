<template>
    <div>
        <Header />
        <div class="container mt-4">
            <h1>Главная страница</h1>
            <p>Добро пожаловать в систему управления авиакомпанией!</p>
        </div>
    </div>
</template>

<script>
import Header from '@/components/Header.vue'

export default {
    name: 'Home',
    components: {
        Header,
    },
}
</script>

<style scoped></style>
