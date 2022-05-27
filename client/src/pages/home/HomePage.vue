<template>
  <MainBanner />
  <ProductsByCategoryList />
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

import MainBanner from './components/MainBanner.vue';
import ProductsByCategoryList from './components/ProductsByCategoryList.vue';

export default {
  components: { MainBanner, ProductsByCategoryList },
  computed: {
    ...mapGetters('home', ['categoryList']),
  },
  created() {
    this.loadData();
  },
  methods: {
    ...mapActions('home', [
      'getBannerList',
      'getCategoryList',
      'getProductListByCategory',
    ]),
    async loadData() {
      this.getBannerList();
    },
  },
};
</script>

<style lang="scss" scoped></style>
