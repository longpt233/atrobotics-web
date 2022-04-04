<template>
  <MainBanner />
  <ProductList />
  <p>{{ productList }}</p>
  <router-view />
</template>

<script>
import MainBanner from './components/MainBanner.vue';
import ProductList from './components/ProductList.vue';
// import store from '@/store/index';
import { mapActions, mapGetters } from 'vuex';

export default {
  name: 'HomePage',
  components: { MainBanner, ProductList },
  computed: {
    ...mapGetters('home', ['productList','userInfo']),
    ...mapGetters('auth', ['token', 'registerUser']),
  },
  methods: {
    ...mapActions('home', ['getProductList','getUserProfile']),
    ...mapActions('auth', ['login', 'register']),
    loadData() {
      this.getProductList({limit: 3, offset: 1});
      this.login({ email: 'tuannha@gmail.com', password: 'tuannguvl123' });
      // const testRegisterForm = {
      //   email: 'tester03@gmail.com',
      //   password: 'tester123',
      //   firstname: 'Test3',
      //   lastname: 'Tester3',
      //   phone: '09345343452',
      //   address: 'Ha Noi',
      // };
      // this.register(testRegisterForm);
      this.getUserProfile("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDkwNDY5NjIsImlhdCI6MTY0OTAzNjE2MiwidXNlcklEIjoiOTA4YzhlMmMtMjQwMi00NDJkLWFiN2QtNDUxZmZhMWEwNWI1In0.pzKmBjAH_LPRT79IReoiJxkAmpkSPQcIyx3LQ8L37ng");
    },
  },
  created() {
    this.loadData();
  },
};
</script>

<style lang="scss" scoped></style>
