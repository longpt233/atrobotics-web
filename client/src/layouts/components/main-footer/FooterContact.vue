<template>
  <div class="footer-contact d-flex flex-column align-items-end">
    <div class="footer-contact__title">
      {{ $t('app.footer.contact.title') }}
    </div>
    <div class="footer-contact__subtitle">
      {{ $t('app.footer.contact.subtitle') }}
    </div>
    <div class="footer-contact__icons d-flex align-items-center">
      <a :href="zaloUrl">
        <img src="@/assets/images/icon-zalo.svg" width="32" height="32" />
      </a>
      <a :href="fanpageUrl">
        <img src="@/assets/images/icon-facebook.svg" width="40" height="40" />
      </a>
    </div>
  </div>
</template>

<script>
export default {
  setup() {
    const fanpageUrl = process.env.VUE_APP_FANEPAGE_URL;
    const zaloUrl = process.env.VUE_APP_ZALO_URL;
    return { fanpageUrl, zaloUrl };
  },
};
</script>

<style lang="scss" scoped>
.footer-contact {
  gap: 10px;
  font-size: 14px;
  &__title {
    color: $color-primary;
    font-size: 20px;
    font-weight: 700;
  }
  &__icons {
    gap: 4px;
  }
}
</style>
