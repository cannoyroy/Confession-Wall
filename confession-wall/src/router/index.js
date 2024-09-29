import { createRouter, createWebHistory } from 'vue-router'
import Bar from '../views/Bar.vue'
import LoginReg from '../views/Account/LoginReg.vue'
import Main from '../views/user/Main.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'bar',
      component: Bar,
      children: [
        { path: '/main', name: 'main', component: Main}
      ]
    },
    {
      path: '/loginreg',
      name: 'loginreg',
      component: LoginReg
    },
  ]
})

export default router
