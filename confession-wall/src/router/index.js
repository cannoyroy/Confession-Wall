import { createRouter, createWebHistory } from 'vue-router'
import Bar from '../views/Bar.vue'
import LoginReg from '../views/Account/LoginReg.vue'
import Main from '../views/user/Main.vue'
import Confession from '../views/user/Confession/Confession.vue'
// import NewConfession from '../views/user/Confession/NewConfession.vue'
import Blacklist from '../views/user/Blacklist/Blacklist.vue'
import Profile from '../views/user/Profile/Profile.vue'
import EditProfile from '../views/user/Profile/EditProfile.vue'
import NewProfile from '../views/user/Profile/NewProfile.vue'
import ProfileC from '../views/user/Profile/ProfileC.vue'
import Setting from '../views/user/Setting/Setting.vue'

import BarAdmin from '@/views/Admin/BarAdmin.vue'
import getReport from '@/views/Admin/getReport.vue'
import historyReport from '@/views/Admin/historyReport.vue'
import handleReport from '@/views/Admin/handleReport.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/bar',
      name: 'bar',
      component: Bar,
      meta: { requiresAuth: true },
      children: [
        { path: '/main', name: 'main', component: Main,meta: { requiresAuth: true }},
        { path: '/confession', name: 'confession', component: Confession,meta: { requiresAuth: true }},
        // { path: '/newconfession', name: 'newconfession', component: NewConfession,meta: { requiresAuth: true }},
        { path: '/blacklist', name: 'blacklist', component: Blacklist,meta: { requiresAuth: true }},
        { path: '/profile', name: 'profile', component: Profile,meta: { requiresAuth: true }},
        { path: '/profilec', name: 'profilec', component: ProfileC,meta: { requiresAuth: true }},
        { path: '/editprofile', name: 'editprofile', component: EditProfile,meta: { requiresAuth: true }},
        { path: '/setting', name: 'setting', component: Setting,meta: { requiresAuth: true }},
      ]
    },
    {
      path: '/admin',
      component: BarAdmin,
      meta: { requiresAuth: true },
      children:[
        {path: 'reports/handle',  name: 'adhandle', component: handleReport,meta: { requiresAuth: true }},
        {path: 'reports',  name: 'adget', component: getReport,meta: { requiresAuth: true }},
        {path: 'reports/history',  name: 'adhistory', component: historyReport,meta: { requiresAuth: true }},
      ]
    },
    {
      path: '/newprofile',
      name: 'newprofile',
      component: NewProfile,
      meta: { requiresAuth: true }
    },
    {
      path: '/',
      name: 'loginreg',
      component: LoginReg
    },
  ]
})


// 在 router/index.ts 中
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token');
  const expiresAt = localStorage.getItem('expiresAt');
  const isAuthenticated = token !== '0' && (Date.now() < expiresAt); // 判断 token 是否存在且未过期
  console.log(token, expiresAt, Date.now() < expiresAt, isAuthenticated);
  // 清除 token 如果目标路由是登录页面
  if (to.name === 'loginreg') {
    localStorage.removeItem('token'); // 清除 token
    localStorage.removeItem('expiresAt'); // 清除过期时间
  }

  // 如果目标路由需要认证且用户未认证
  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ path: '/' }); // 重定向到登录页面
  } else {
    next(); // 继续访问目标路由
  }
});


export default router
