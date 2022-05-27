<template>
  <div class="product-detail-page d-flex flex-column">
    <ProductCard v-loading="isLoading" />
    <ProductDescription />
  </div>
</template>

<script>
import { mapActions } from 'vuex';
import ProductCard from './components/ProductCard.vue';
import ProductDescription from './components/ProductDescription.vue';

export default {
  components: { ProductCard, ProductDescription },
  created() {
    this.loadData();
  },
  data() {
    return {
      isLoading: true,
    };
  },
  methods: {
    ...mapActions('productDetail', ['getProductDetail']),
    async loadData() {
      this.isLoading = true;
      const productId = this.$route.params.id;
      await this.getProductDetail(productId);
      this.isLoading = false;
    },
  },
};
</script>

<style lang="scss" scoped>
.product-detail-page {
  margin-top: 130px;
  padding-bottom: 50px;
  gap: 30px;
}
</style>
