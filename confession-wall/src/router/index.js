import { createRouter, createWebHistory } from 'vue-router'
import Bar from '../views/Bar.vue'
import Main from '../views/user/Main.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/a',
      name: 'bar',
      component: Bar,
      children: [
        { path: '/main', component: Main}
      ]
    },
  ]
})

export default router
