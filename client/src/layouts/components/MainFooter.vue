<template>
  <div class="container-fluid">
    <div class="container">
      <div class="main-footer row">
        <FooterInformation class="col-5" />
        <FooterWebsite class="col-3" />
        <FooterContact class="col-4" />
      </div>
    </div>
  </div>
</template>

<script>
import FooterInformation from './main-footer/FooterInformation.vue';
import FooterWebsite from './main-footer/FooterWebsite.vue';
import FooterContact from './main-footer/FooterContact.vue';

export default {
  components: { FooterInformation, FooterWebsite, FooterContact },
};
</script>

<style lang="scss" scoped>
.container-fluid {
  margin-top: auto;
  background-color: $color-secondary;
}

.main-footer {
  height: fit-content;
  padding-top: 28px;
  padding-bottom: 20px;
}
</style>
