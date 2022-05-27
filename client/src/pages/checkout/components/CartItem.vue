<template>
  <div class="cart-item-contaienr d-flex justify-content-between align-items-center">
    <ProductItem
      class="product"
      :image-url="item?.productData?.images[0]"
      :product-id="item?.productId"
      :title="item?.productData?.name"
      :color="item?.color"
    />
    <div class="price">
      <span>{{ ` ${parsePrice(item?.productData?.price)}` }}</span>
      <sup>{{ $t('common.currency') }}</sup>
    </div>
    <el-input-number v-model="quantity" :min="1" @change="onChangeQuantity" />
    <div class="total">
      <span>{{ ` ${parsePrice(total)}` }}</span>
      <sup>{{ $t('common.currency') }}</sup>
    </div>
  </div>
</template>

<script>
import ProductItem from '@/components/ProductItem.vue';
import globalMixin from '@/common/mixins';
import { mapActions } from 'vuex';

export default {
  mixins: [globalMixin],
  components: { ProductItem },
  props: {
    item: Object,
  },
  data() {
    return {
      quantity: this.item?.quantity,
    };
  },
  methods: {
    ...mapActions('checkout', ['updateQuantity']),
    async updateQuantity(quantity) {
      this.updateQuantity(quantity);
    },
    onChangeQuantity(newQuantity) {},
  },
  computed: {
    total() {
      return this.item?.productData?.price * this.quantity;
    },
  },
};
</script>

<style lang="scss" scoped>
.cart-item-contaienr {
  padding: 0 0 4px;
  border-bottom: 1px solid $color-light-grey-3;
  .product {
    flex-basis: 350px;
  }
  .price {
    flex-basis: 130px;
    text-align: center;
  }
  .quantity {
    flex-basis: 80px;
    gap: 8px;
    font-size: 15px;
    font-weight: 500;
    button {
      width: 20px;
      height: 20px;
      padding: 0;
      font-size: 20px;
      border-radius: 4px;
      &:hover,
      &:focus {
        color: $color-white;
      }
    }
    .button-inc-quantity {
      background-color: $color-green;
    }
    .button-dec-quantity {
      background-color: $color-red-1;
    }
  }
  .total {
    flex-basis: 130px;
    text-align: center;
  }
}
</style>
