import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import AdminPanel from '@/views/AdminPanel.vue'
import AdminAircraft from '@/views/AdminAircraft.vue'
import AdminStaff from '@/views/AdminStaff.vue'
import AdminRoutes from '@/views/AdminRoutes.vue'
import AdminRoute from '@/views/AdminRoute.vue'

const routes = [
    { path: '/', name: 'Home', component: Home },
    { path: '/admin', name: 'AdminPanel', component: AdminPanel },
    { path: '/admin/aircraft/:id', name: 'AdminAircraft', component: AdminAircraft, props: true },
    { path: '/admin/staff', name: 'AdminStaff', component: AdminStaff },
    { path: '/admin/routes', name: 'AdminRoutes', component: AdminRoutes },
    { path: '/admin/route/:id', name: 'AdminRoute', component: AdminRoute, props: true  },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
