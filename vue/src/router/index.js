import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Home from '../views/Home.vue'
import Students from '../views/Students.vue'
import Laboratories from '../views/Laboratories.vue'
import UserManagement from '../views/UserManagement.vue'
import CourseManagement from '../views/CourseManagement.vue'

const routes = [
    { path: '/', component: Login },
    { path: '/home', component: Home, children: [
            { path: 'user', component: UserManagement },
            { path: 'students', component: Students },
            { path: 'laboratories', component: Laboratories },
            { path: 'courses', component: CourseManagement }
        ]}
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router
