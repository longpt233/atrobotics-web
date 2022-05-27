<template>
  <router-link class="router-link" :to="`/product/${productItem.id}`">
    <el-card class="product-item-box" :body-style="{ padding: '10px' }">
      <div class="image-box d-flex flex-column justify-content-center">
        <img :src="productItem.images[0]" />
      </div>
      <div class="product-info">
        <div class="product-info__name">{{ productItem.name }}</div>
        <div class="product-info__price">
          <span>{{ parsePrice(productItem.price) }}</span>
          <sup>{{ $t('common.currency') }}</sup>
        </div>
      </div>
    </el-card>
  </router-link>
</template>

<script>
import globalMixin from '@/common/mixins';
export default {
  mixins: [globalMixin],
  props: {
    productItem: Object,
  },
};
</script>

<style lang="scss" scoped>
.product-item-box {
  width: 270px;
  height: 295px;
}
.image-box {
  width: 100%;
  height: 200px;
  margin-bottom: 20px;
  img {
    height: 100%;
  }
}
.product-info {
  text-align: center;
  &__name {
    font-weight: 500;
    @include limit-by-n-line(1);
  }
  &__price {
    font-weight: 500;
    color: $color-red;
  }
}
</style>
