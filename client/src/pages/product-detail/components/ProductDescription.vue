<template>
  <div class="product-description d-flex flex-column align-items-center">
    <div class="product-description__title">
      {{ $t('productDetail.description.title') }}
    </div>
    <div class="product-description__content">
      {{ productDetail?.longDesc }}
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  computed: {
    ...mapGetters('productDetail', ['productDetail']),
  },
};
</script>

<style lang="scss" scoped>
.product-description {
  gap: 5px;
  &__title {
    font-size: 24px;
    font-weight: 500;
    color: $color-blue;
  }
  &__content {
    width: 1100px;
  }
}
</style>
