<template>
  <div class="product-info-box d-flex flex-column justify-content-evenly">
    <div>
      <div class="product-name">{{ productDetail?.name }}</div>
      <div class="product-price">
        <span>{{ parsePrice(productDetail?.price) }}</span>
        <sup>{{ $t('common.currency') }}</sup>
      </div>
      <div class="product-color d-flex align-items-center">
        <span>{{ `${$t('productDetail.color')}: ` }}</span>
        <div
          class="button-select-color"
          v-for="color in productDetail?.color"
          :key="color"
          :class="color === selectedColor ? 'selected-color' : ''"
          @click="selectedColor = color"
        >
          <div class="icon-color" :style="{ backgroundColor: color }"></div>
        </div>
      </div>
      <div class="product-short-desc">
        {{ productDetail?.shortDesc }}
      </div>
    </div>
    <div>
      <el-button
        class="button-buy-now"
        type="primary"
        @click="buyNow"
        :disabled="productDetail?.available <= 0"
      >
        <span>{{ $t('productDetail.buyNow') }}</span>
      </el-button>
      <el-button
        class="button-add-to-cart"
        type="primary"
        @click="addToCart"
        :disabled="productDetail?.available <= 0"
      >
        <span>{{ $t('productDetail.addToCart') }}</span>
        <img src="@/assets/images/icon-cart.svg" />
      </el-button>
      <div class="product-status d-flex">
        <div class="product-available">
          <span class="product-available--stocking" v-if="productDetail?.available > 0">
            {{ $t('productDetail.stocking') }}
          </span>
          <span class="product-available--out-of-stock" v-else>
            {{ $t('productDetail.outOfStock') }}
          </span>
        </div>
        <div class="product-sold">
          {{ `${$t('productDetail.sold')}: ${productDetail?.sold}` }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import globalMixin from '@/common/mixins';
import { PageName } from '@/common/constants';
import { mapGetters } from 'vuex';

export default {
  mixins: [globalMixin],
  computed: {
    ...mapGetters('productDetail', ['productDetail']),
  },
  data() {
    return {
      selectedColor: '',
    };
  },
  methods: {
    buyNow() {
      if (this.selectedColor === '') {
        return this.showWarningNotification(
          this.$t('productDetail.notification.requestChooseColor.title'),
          this.$t('productDetail.notification.requestChooseColor.message'),
        );
      }
      this.$router.push({
        name: PageName.LOGIN_PAGE,
      });
    },
    addToCart() {
      if (this.selectedColor === '') {
        return this.showWarningNotification(
          this.$t('productDetail.notification.requestChooseColor.title'),
          this.$t('productDetail.notification.requestChooseColor.message'),
        );
      }
      return this.showSuccessNotification(
        this.$t('productDetail.notification.successAddToCart.title'),
        this.$t('productDetail.notification.successAddToCart.message'),
      );
    },
  },
};
</script>

<style lang="scss" scoped>
.product-name {
  margin-bottom: 3px;
  font-size: 26px;
  font-weight: 500;
  color: $color-dark-blue;
  @include limit-by-n-line(2);
}

.product-price {
  margin-bottom: 5px;
  font-size: 18px;
  font-weight: 500;
  color: $color-red;
}

.product-color {
  gap: 10px;
  margin-bottom: 7px;
  font-size: 16px;
  .button-select-color {
    padding: 2px;
    border-radius: 45px;
  }
  .icon-color {
    width: 18px;
    height: 18px;
    border-radius: 45px;
    cursor: pointer;
  }
}
.selected-color {
  border: 2px solid $color-black;
}

.product-short-desc {
  font-size: 15px;
  margin-bottom: 15px;
  color: $color-dark-grey;
}

.button-buy-now,
.button-add-to-cart {
  padding: 18px 15px;
  margin-bottom: 12px;
  font-size: 18px;
  border-radius: 5px;
}
.button-buy-now {
  margin-right: 4px;
  @include set-background-color($color-green);
}
.button-add-to-cart {
  img {
    margin-left: 3px;
  }
}

.product-available {
  margin: 0 70px 0 20px;
  font-weight: 700;
  &--stocking {
    color: $color-green;
  }
  &--out-of-stock {
    color: $color-red;
  }
}
.product-sold {
  font-weight: 500;
}
</style>
