<template>
  <div class="product-list-label-box">
    <div class="title">{{ keyword || $t('productList.label.title') }}</div>
    <div class="result-count">
      {{ `${productList.total} ${$t('productList.label.resultCount')}` }}
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  data() {
    return {
      keyword: this.$route.query.q,
    };
  },
  computed: {
    ...mapGetters('productList', ['productList']),
  },
};
</script>

<style lang="scss" scoped>
.title {
  font-size: 24px;
  font-weight: 500;
  color: $color-blue;
}
</style>
