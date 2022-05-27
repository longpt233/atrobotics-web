<template>
  <div class="delivery-address-box">
    <div class="title">{{ $t('checkout.deliveryAddress.title') }}</div>
    <div class="content">
      <div class="no-address">
        <div class="no-address__title">
          {{ $t('checkout.deliveryAddress.content.noAddress.title') }}
        </div>
        <div class="button-add-new-address d-flex justify-content-center">
          <el-button type="warning">{{
            $t('checkout.deliveryAddress.content.noAddress.addNewAddress')
          }}</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import globalMixin from '@/common/mixins';

export default {
  mixins: [globalMixin],
};
</script>

<style lang="scss" scoped>
.delivery-address-box {
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
.no-address {
  &__title {
    font-size: 18px;
    margin-bottom: 14px;
  }
  .button-add-new-address {
    width: 100%;
    button {
      padding: 18px 15px;
      font-size: 18px;
      border-radius: 5px;
    }
  }
}
</style>
