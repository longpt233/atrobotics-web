import { createRouter, createWebHistory } from 'vue-router';

import MainLayout from '@/layouts/MainLayout.vue';
import HomePage from '@/pages/home/HomePage.vue';

const routes = [
  {
    path: '/',
    component: MainLayout,
    children: [
      {
        path: '',
        name: 'HomePage',
        component: HomePage,
      },
    ],
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
