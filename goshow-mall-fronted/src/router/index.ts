import { createRouter, createWebHistory } from 'vue-router'
// @ts-ignore
import HomeView from '@/pages/HomeView.vue'
import AboutView from '@/pages/AboutView.vue'

// @ts-ignore
// @ts-ignore
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView,
    },
  ],
})

export default router
