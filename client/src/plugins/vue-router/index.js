import { createRouter, createWebHistory } from 'vue-router';
import { PageName } from '@/common/constants';
import MainLayout from '@/layouts/MainLayout.vue';
import HomePage from '@/pages/home/HomePage.vue';
import ProductListPage from '@/pages/product-list/ProductListPage.vue';
import LoginPage from '@/pages/auth/LoginPage.vue';
import RegisterPage from '@/pages/auth/RegisterPage.vue';
import CheckoutPage from '@/pages/checkout/CheckoutPage.vue';
import ProductDetailPage from '@/pages/product-detail/ProductDetailPage.vue';

const routes = [
  {
    path: '/',
    component: MainLayout,
    children: [
      {
        path: '/',
        name: PageName.HOME_PAGE,
        component: HomePage,
      },
      {
        path: '/home',
        redirect: '/',
      },
      {
        path: '/login',
        name: PageName.LOGIN_PAGE,
        component: LoginPage,
      },
      {
        path: '/signup',
        name: PageName.REGISTER_PAGE,
        component: RegisterPage,
      },
      {
        path: '/product',
        name: PageName.PRODUCT_LIST_PAGE,
        component: ProductListPage,
      },
      {
        path: '/product/:id',
        name: PageName.PRODUCT_DETAIL_PAGE,
        component: ProductDetailPage,
      },
      {
        path: '/checkout',
        name: PageName.CHECKOUT_PAGE,
        component: CheckoutPage,
        meta: { requiresAuth: true },
      },
      {
        path: '/checkout',
        name: PageName.CHECKOUT_PAGE,
        component: CheckoutPage,
        meta: { requiresAuth: true },
      },
      {
        path: '/order',
        name: PageName.ORDER_PAGE,
        component: CheckoutPage,
        meta: { requiresAuth: true },
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach((to, from, next) => {
  const loggedIn = localStorage.getItem('token');
  if (to.matched.some((record) => record.meta.requiresAuth) && !loggedIn) {
    next('/login');
  }
  next();
});

export default router;
