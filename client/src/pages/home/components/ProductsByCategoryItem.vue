<template>
  <div class="products-by-category-item d-flex flex-column align-items-center">
    <div class="products-by-category-item__title">
      {{ category }}
    </div>
    {{ productInfo }}
  </div>
</template>

<script>
export default {
  props: {
    category: String,
    productInfo: Object,
  },
};
</script>

<style lang="scss" scoped>
.products-by-category-item {
  &__title {
    font-size: 28px;
    font-weight: 500;
    color: $color-blue;
  }
}
</style>
