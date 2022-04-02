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
    ...mapActions('home', ['getProductList']),
    ...mapGetters('home', ['productList']),
  },
  methods: {
    loadData() {
      this.getProductList;
      console.log(this.productList);
    },
  },
  created() {
    this.loadData();
  },
};
</script>

<style lang="scss" scoped></style>
