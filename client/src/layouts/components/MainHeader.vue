<template>
  <div class="main-header-box">
    <HeadingBar />
    <NavigationBar />
  </div>
</template>

<script>
import NavigationBar from './main-header/NavigationBar.vue';
import HeadingBar from './main-header/HeadingBar.vue';

export default {
  components: { NavigationBar, HeadingBar },
};
</script>

<style lang="scss" scoped>
.main-header-box {
  position: fixed;
  top: 0;
  width: 100vw;
  z-index: map-get($zIndex, 'menu');
}
</style>
