<template>
  <div class="cart-list-box">
    <div class="cart-list-header d-flex justify-content-between">
      <div class="header product">{{ $t('checkout.cartList.header.product') }}</div>
      <div class="header price">{{ $t('checkout.cartList.header.price') }}</div>
      <div class="header quantity">{{ $t('checkout.cartList.header.quantity') }}</div>
      <div class="header total">{{ $t('checkout.cartList.header.total') }}</div>
    </div>
    <div class="cart-list-content d-flex flex-column">
      <CartItem v-for="item in itemList" :key="item.id" :item="item" />
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import CartItem from './CartItem.vue';

export default {
  components: { CartItem },
  computed: {
    ...mapGetters('checkout', ['itemList']),
  },
};
</script>

<style lang="scss" scoped>
.cart-list-box {
  width: 100%;
  padding: 0 30px;
}
.cart-list-header {
  margin-bottom: 8px;
  font-size: 20px;
  font-weight: 700;
  color: $color-blue;
  border-bottom: 2px solid $color-light-grey-2;
  .header {
    font-size: 20px;
    text-align: center;
  }
  .product {
    flex-basis: 350px;
  }
  .price {
    flex-basis: 130px;
  }
  .quantity {
    flex-basis: 90px;
  }
  .total {
    flex-basis: 130px;
  }
}
.cart-list-content {
  gap: 10px;
}
</style>
