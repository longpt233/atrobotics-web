<template>
  <div class="container container">
    <div class="product-list d-grid justify-content-between">
      <ProductItem
        v-for="productItem in productList.data"
        :key="productItem.id"
        :product-item="productItem"
      />
    </div>
    <div class="pagination d-flex justify-content-center">
      <el-pagination
        background
        layout="prev, pager, next"
        hide-on-single-page
        :page-size="DEFAULT_LIMIT_PRODUCT"
        :total="productList.total"
        :current-page="currentPage"
        @current-change="updateCurrentPage"
      />
    </div>
  </div>
</template>

<script>
import ProductItem from './ProductItem.vue';
import { mapGetters, mapActions } from 'vuex';

import { DEFAULT_LIMIT_PRODUCT, RESET_OFFSET_PRODUCT } from '../constants';

export default {
  setup() {
    return { DEFAULT_LIMIT_PRODUCT };
  },
  components: { ProductItem },
  computed: {
    ...mapGetters('productList', ['productList']),
    ...mapGetters('productList', { filterStore: 'filter' }),
  },
  data() {
    return {
      currentPage: 1,
    };
  },
  watch: {
    'filterStore.sort-by': {
      handler() {
        this.resetCurrentPage();
      },
      deep: true,
    },
    'filterStore.brand': {
      handler() {
        this.resetCurrentPage();
      },
      deep: true,
    },
  },
  methods: {
    ...mapActions('productList', ['getProductList']),
    updateCurrentPage(newPage) {
      this.currentPage = newPage;
      const filter = this.filterStore || {};
      filter.offset = this.currentPage;

      this.getProductList(filter);
    },
    resetCurrentPage() {
      this.currentPage = RESET_OFFSET_PRODUCT;
    },
  },
};
</script>

<style lang="scss" scoped>
.container {
  margin-bottom: 40px;
}
.product-list {
  row-gap: 40px;
  grid-template-columns: repeat(auto-fill, 270px);
  margin-bottom: 40px;
}
.pagination {
  text-align: center;
}
</style>
