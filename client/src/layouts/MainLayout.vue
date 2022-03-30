<template>
  <MainHeader />
  <NavigationBar />
  <router-view />
</template>

<script>
import MainHeader from './components/MainHeader.vue';
import NavigationBar from './components/NavigationBar.vue';

export default {
  components: {
    MainHeader,
    NavigationBar,
  },
};
</script>

<style lang="scss" scoped></style>
