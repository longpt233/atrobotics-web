<template>
  <div class="footer-website d-flex flex-column align-items-start">
    <div class="footer-website__title">
      {{ $t('app.footer.website.title') }}
    </div>
    <div class="footer-website__website" v-for="website in websites" :key="website.id">
      <a class="router-link" :href="website.link">
        {{ website.title }}
      </a>
    </div>
  </div>
</template>

<script>
import { websites } from '../../constants';
export default {
  setup() {
    return { websites };
  },
  components: {},
};
</script>

<style lang="scss" scoped>
.footer-website {
  gap: 10px;
  font-size: 14px;
  &__title {
    color: $color-primary;
    font-size: 20px;
    font-weight: 700;
  }
}
</style>
