import { createRouter, createWebHistory } from 'vue-router';
import ShowConfession from '@/components/ShowConfession.vue';
import CreateConfession from '@/components/CreateConfession.vue';


const routes = [
  { path: '/', name: 'ShowConfession', component: ShowConfession },
  { path: '/new', name: 'CreateConfession', component: CreateConfession },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
