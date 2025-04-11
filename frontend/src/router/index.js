import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import AdminPanel from '@/views/AdminPanel.vue'

const routes = [
    { path: '/', name: 'Home', component: Home },
    { path: '/admin', name: 'AdminPanel', component: AdminPanel },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
