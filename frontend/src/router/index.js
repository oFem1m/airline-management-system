import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import AdminPanel from '@/views/AdminPanel.vue'
import AdminAircraft from '@/views/AdminAircraft.vue'
import AdminStaff from '@/views/AdminStaff.vue'
import AdminRoutes from '@/views/AdminRoutes.vue'
import AdminRoute from '@/views/AdminRoute.vue'
import AdminFlight from '@/views/AdminFlight.vue'
import AdminAircrafts from '@/views/AdminAircrafts.vue'
import AdminAirports from '@/views/AdminAirports.vue'
import AdminAirport from '@/views/AdminAirport.vue'
import AdminPassengers from '@/views/AdminPassengers.vue'
import AdminPassenger from '@/views/AdminPassenger.vue'
import AdminBookings from '@/views/AdminBookings.vue'
import AdminBooking from '@/views/AdminBooking.vue'
import Flight from '@/views/Flight.vue'
import Booking from '@/views/Booking.vue'

const routes = [
    { path: '/', name: 'Home', component: Home },
    { path: '/flight/:id', name: 'Flight', component: Flight, props: true },
    { path: '/flight/:flight_id/booking/:id', name: 'Booking', component: Booking, props: true },
    { path: '/admin', name: 'AdminPanel', component: AdminPanel },
    { path: '/admin/aircrafts', name: 'AdminAircrafts', component: AdminAircrafts },
    { path: '/admin/aircraft/:id', name: 'AdminAircraft', component: AdminAircraft, props: true },
    { path: '/admin/airports', name: 'AdminAirports', component: AdminAirports },
    { path: '/admin/airport/:id', name: 'AdminAirport', component: AdminAirport, props: true },
    { path: '/admin/staff', name: 'AdminStaff', component: AdminStaff },
    { path: '/admin/routes', name: 'AdminRoutes', component: AdminRoutes },
    { path: '/admin/route/:id', name: 'AdminRoute', component: AdminRoute, props: true  },
    { path: '/admin/flight/:id', name: 'AdminFlight', component: AdminFlight, props: true  },
    { path: '/admin/passengers', name: 'AdminPassengers', component: AdminPassengers },
    { path: '/admin/passenger/:id', name: 'AdminPassenger', component: AdminPassenger, props: true },
    { path: '/admin/bookings', name: 'AdminBookings', component: AdminBookings },
    { path: '/admin/booking/:id', name: 'AdminBooking', component: AdminBooking, props: true },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
