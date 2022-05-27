<template>
  <div
    class="product-list-header-contanier container d-flex justify-content-between align-items-center"
  >
    <ProductListLabel />
    <ProductListFilter />
  </div>
</template>

<script>
import ProductListLabel from './ProductListLabel.vue';
import ProductListFilter from './ProductListFilter.vue';

export default {
  components: { ProductListLabel, ProductListFilter },
};
</script>

<style lang="scss" scoped>
.product-list-header-contanier {
  margin-top: 100px;
  margin-bottom: 20px;
}
</style>
