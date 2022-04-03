import { createRouter, createWebHistory } from 'vue-router';
import AdminPage from '@/pages/dashboard/AdminPage';

const routes = [
  {
    path: '/',
    name: 'Admin',
    component: AdminPage,
    meta: {
      title: 'AT Robotics',
    },
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach((to) => {
  document.title = to.meta.title;
});

export default router;
