<template>
  <div
    class="zalo-chat-widget"
    :data-oaid="id"
    data-autopopup="0"
    data-width="300"
    data-height="300"
  ></div>
</template>

<script>
export default {
  setup() {
    const id = process.env.VUE_APP_ID_ZALO;
    return { id };
  },
};
</script>

<style lang="scss" scoped>
.zalo-chat-widget,
.zalo-share-button {
  bottom: 93px !important;
  right: 25px !important;
  z-index: map-get($zIndex, popup) !important;
}
</style>
