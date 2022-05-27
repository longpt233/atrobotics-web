<template>
  <div class="product-item-box d-flex">
    <router-link :to="`product/${productId}`">
      <img src="imageUrl" />
    </router-link>
    <div class="item-info">
      <router-link class="router-link" :to="`product/${productId}`">
        <div class="title">{{ title }}</div>
      </router-link>
      <div class="color">{{ `${$t('common.productItem.color')}: ${color}` }}</div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    imageUrl: String,
    productId: String,
    title: String,
    color: String,
  },
};
</script>

<style lang="scss" scoped>
.product-item-box {
  gap: 12px;
  width: 100%;
  img {
    width: 100px;
    height: 80px;
    margin-bottom: 6px;
    border-radius: 2px;
  }
  .item-info {
    margin: 10px 0;
    .title {
      width: 100%;
      font-size: 14px;
      font-weight: 500;
      color: $color-black-1;
      @include limit-by-n-line(2);
      &:hover {
        color: $color-primary;
      }
    }
    .color {
      font-size: 12px;
      color: $color-grey;
    }
  }
}
</style>
