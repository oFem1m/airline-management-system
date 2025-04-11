import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import AdminPanel from '@/views/AdminPanel.vue'
import AdminAircraft from '@/views/AdminAircraft.vue'
import AdminStaff from '@/views/AdminStaff.vue'

const routes = [
    { path: '/', name: 'Home', component: Home },
    { path: '/admin', name: 'AdminPanel', component: AdminPanel },
    { path: '/admin/aircraft/:id', name: 'AdminAircraft', component: AdminAircraft, props: true },
    { path: '/admin/staff', name: 'AdminStaff', component: AdminStaff },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
