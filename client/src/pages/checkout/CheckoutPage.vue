<template>
  <div class="container d-flex justify-content-between">
    <CartList />
    <div class="checkout-summary d-flex flex-column">
      <DeliveryAddress />
      <CartTotals />
    </div>
  </div>
</template>

<script>
import CartList from './components/CartList.vue';
import CartTotals from './components/CartTotals.vue';
import DeliveryAddress from './components/DeliveryAddress.vue';
import { mapActions } from 'vuex';

export default {
  components: { CartList, DeliveryAddress, CartTotals },
  created() {
    this.loadData();
  },
  methods: {
    ...mapActions('checkout', ['getItemList']),
    async loadData() {
      await this.getItemList();
    },
  },
};
</script>

<style lang="scss" scoped>
.container {
  gap: 50px;
  margin-top: 110px;
  margin-bottom: 60px;
}
.checkout-summary {
  gap: 20px;
}
</style>
