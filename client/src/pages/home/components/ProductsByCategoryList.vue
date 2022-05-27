<template>
  <div class="product-category-list">
    <ProductsByCategoryItem category="Máy hút bụi" />
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import ProductsByCategoryItem from './ProductsByCategoryItem.vue';

export default {
  components: { ProductsByCategoryItem },
  computed: {
    ...mapGetters('home', ['productListByCategory']),
  },
};
</script>

<style lang="scss" scoped></style>
