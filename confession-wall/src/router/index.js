import { createRouter, createWebHistory } from 'vue-router'
import Bar from '../views/Bar.vue'
import LoginReg from '../views/Account/LoginReg.vue'
import Main from '../views/user/Main.vue'
import Confession from '../views/user/Confession/Confession.vue'
import NewConfession from '../views/user/Confession/NewConfession.vue'
import Blacklist from '../views/user/Blacklist/Blacklist.vue'
import Profile from '../views/user/Profile/Profile.vue'
import EditProfile from '../views/user/Profile/EditProfile.vue'
import NewProfile from '../views/user/Profile/NewProfile.vue'
import ProfileC from ',,/views/user/Profile/ProfileC.vue'
import Setting from '../views/user/Setting/Setting.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/bar',
      name: 'bar',
      component: Bar,
      children: [
        { path: '/main', name: 'main', component: Main},
        { path: '/confession', name: 'confession', component: Confession},
        { path: '/newconfession', name: 'newconfession', component: NewConfession},
        { path: '/blacklist', name: 'blacklist', component: Blacklist},
        { path: '/profile', name: 'profile', component: Profile},
        { path: '/profilec', name: 'profilec', component: ProfileC},
        { path: '/editprofile', name: 'editprofile', component: EditProfile},
        { path: '/setting', name: 'setting', component: Setting},
      ]
    },
    {
      path: '/newprofile',
      name: 'newprofile',
      component: NewProfile
    },
    {
      path: '/',
      name: 'loginreg',
      component: LoginReg
    },
  ]
})

export default router
