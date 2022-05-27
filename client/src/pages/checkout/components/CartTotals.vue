<template>
  <div class="cart-totals-box">
    <div class="title">{{ $t('checkout.cartTotals.title') }}</div>
    <div class="content">
      <div class="totals d-flex justify-content-between">
        <div class="totals__title">
          {{ $t('checkout.cartTotals.content.totals') }}
        </div>
        <div class="totals__value">{{ parsePrice(totalAmount) }}</div>
      </div>
      <div class="additional-info">{{ $t('checkout.cartTotals.additionalInfo') }}</div>
      <div class="button-checkout d-flex justify-content-center">
        <el-button type="primary">
          {{ $t('checkout.cartTotals.submit') }}
        </el-button>
      </div>
    </div>
  </div>
</template>

<script>
import globalMixin from '@/common/mixins';
import { mapGetters } from 'vuex';

export default {
  mixins: [globalMixin],
  computed: {
    ...mapGetters('checkout', ['totalAmount']),
  },
};
</script>

<style lang="scss" scoped>
.cart-totals-box {
  flex-shrink: 0;
  width: 400px;
}
.title {
  margin-bottom: 12px;
  text-align: center;
  font-size: 20px;
  font-weight: 700;
  color: $color-blue;
}
.content {
  width: 100%;
  height: fit-content;
  padding: 30px;
  background-color: $color-light-grey-1;
}
.totals {
  margin-bottom: 16px;
  font-size: 20px;
  font-weight: 500;
  color: $color-dark-blue-1;
  border-bottom: 2px solid $color-light-grey-3;
}
.additional-info {
  margin-bottom: 8px;
  font-style: italic;
  color: $color-dark-grey;
}
.button-checkout {
  width: 100%;
  button {
    padding: 18px 15px;
    font-size: 18px;
    border-radius: 5px;
  }
}
</style>
