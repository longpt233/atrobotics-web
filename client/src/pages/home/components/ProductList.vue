<template>
  <div class="container overflow-hidden">
    <div v-for="index in 2" :key="index" class="row g-5">
      <div v-for="index in 3" :key="index" class="col-4 p-3">
        <img src="@/assets/images/image-product.png" />
      </div>
    </div>
  </div>
</template>

<script>
import { productList } from '@/services/home.api.service';

console.log(productList);

export default {};
</script>

<style lang="scss" scoped>
img {
  width: 100%;
  height: auto;
  cursor: pointer;
}
</style>
