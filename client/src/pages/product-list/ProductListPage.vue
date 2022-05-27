<template>
  <ProductListHeader />
  <ProductList />
</template>

<script>
import ProductListHeader from './components/ProductListHeader.vue';
import ProductList from './components/ProductList.vue';
import { mapActions } from 'vuex';

import { DEFAULT_LIMIT_PRODUCT, RESET_OFFSET_PRODUCT } from './constants';

export default {
  components: { ProductListHeader, ProductList },
  created() {
    this.loadData();
  },
  methods: {
    ...mapActions('productList', ['getProductList', 'getBrandList']),
    loadData() {
      const filter = this.$route.query;
      filter.limit = DEFAULT_LIMIT_PRODUCT;
      filter.offset = RESET_OFFSET_PRODUCT;
      this.getProductList(filter);

      this.getBrandList();
    },
  },
};
</script>

<style lang="scss" scoped></style>
