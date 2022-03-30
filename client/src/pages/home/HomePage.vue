<template>
  <MainBanner />
  <ProductList />
  <router-view />
</template>

<script>
// @ is an alias to /src

import MainBanner from './components/MainBanner.vue';
import ProductList from './components/ProductList.vue';

export default {
  name: 'HomePage',
  components: { MainBanner, ProductList },
};
</script>

<style lang="scss" scoped></style>
