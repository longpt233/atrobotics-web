<template>
  <div class="main-layout d-flex flex-column">
    <MainHeader />
    <router-view v-slot="{ Component }" :key="new Date()">
      <component :is="Component" />
    </router-view>
    <MainFooter />
  </div>
  <ZaloChat />
  <MessengerChat />
</template>

<script>
import MainHeader from './components/MainHeader.vue';
import MainFooter from './components/MainFooter.vue';
import { MessengerChat } from '@/plugins/social/messenger';
import { ZaloChat } from '@/plugins/social/zalo';

export default {
  components: { MainHeader, MainFooter, MessengerChat, ZaloChat },
};
</script>

<style lang="scss" scoped>
.main-layout {
  height: 100vh;
}
</style>
