<template>
  <div
    id="fb-customer-chat"
    class="fb-customerchat"
    :page-id="id"
    atribution="biz_inbox"
  ></div>
</template>

<script>
export default {
  setup() {
    const id = process.env.VUE_APP_ID_MESSENGER;
    return { id };
  },
};
</script>

<style lang="scss" scoped></style>

<style lang="scss">
@use './style.scss';
</style>
