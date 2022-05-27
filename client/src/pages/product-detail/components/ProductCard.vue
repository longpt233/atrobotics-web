<template>
  <div class="product-card">
    <el-card :body-style="{ padding: '12px' }">
      <div class="product-card-content d-flex">
        <ProductImages />
        <ProductInfo />
      </div>
    </el-card>
  </div>
</template>

<script>
import ProductImages from './ProductImages.vue';
import ProductInfo from './ProductInfo.vue';

export default {
  components: { ProductImages, ProductInfo },
};
</script>

<style lang="scss" scoped>
.product-card {
  width: 1100px;
  display: block;
  margin: 0 auto;
}

.product-card-content {
  gap: 40px;
}
</style>
