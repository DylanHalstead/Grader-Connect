import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LogInView from '../views/LogInView.vue'
import AssignmentView from '../views/AssignmentView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        title: 'Home'
      }
    },
    {
      path: '/login',
      name: 'login',
      component: LogInView,
      meta: {
        title: 'Login'
      }
    },
    {
      path: '/assignments',
      name: 'assignments',
      component: AssignmentView,
      meta: {
        title: 'Assignments'
      }
    }
  ]
})

export default router
