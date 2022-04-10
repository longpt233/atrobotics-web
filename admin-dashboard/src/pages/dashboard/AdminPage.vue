<template>
  <SideBar />
  <router-view />
</template>

<script>
import SideBar from './components/SideBar.vue';

export default {
  name: 'AdminPage',
  components: { SideBar },
};
</script>

<style lang="scss" scoped></style>
