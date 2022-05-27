<template>
  <div class="product-list-filter d-flex align-items-center">
    <div class="brand">
      <span class="brand__title">{{ $t('productList.filter.brand.title') }}</span>
      <el-select
        class="brand__select"
        v-model="brandOptionSelected"
        :placeholder="$t('productList.filter.brand.placeholder')"
        @change="onSelectBrandOption"
        clearable
      >
        <el-option
          v-for="brand in brandList"
          :key="brand"
          :label="brand"
          :value="brand"
        />
      </el-select>
    </div>
    <div class="sort">
      <span class="sort__title">{{ $t('productList.filter.sort.title') }}</span>
      <el-select
        class="sort__select"
        v-model="sortOptionSelected"
        :placeholder="$t('productList.filter.sort.placeholder')"
        @change="onSelectSortOption"
        clearable
      >
        <el-option
          v-for="index in parseInt($t('productList.filter.sort.optionsLength'))"
          :key="index"
          :label="$t(`productList.filter.sort.options[${index - 1}].label`)"
          :value="$t(`productList.filter.sort.options[${index - 1}].value`)"
        />
      </el-select>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

import { RESET_OFFSET_PRODUCT } from '../constants';

export default {
  watchQuery: true,
  computed: {
    ...mapGetters('productList', ['brandList']),
    ...mapGetters('productList', { filterStore: 'filter' }),
  },
  data() {
    return {
      brandOptionSelected: '',
      sortOptionSelected: '',
    };
  },
  methods: {
    ...mapActions('productList', ['getProductList']),
    onSelectSortOption() {
      const filter = this.filterStore || {};
      filter['sort-by'] = this.sortOptionSelected;
      filter.offset = RESET_OFFSET_PRODUCT;

      this.getProductList(filter);
    },
  },
};
</script>

<style lang="scss" scoped>
.product-list-filter {
  gap: 15px;
}
.brand {
  width: fit-content;
  &__select {
    width: 125px;
    margin-left: 10px;
  }
}
.sort {
  width: fit-content;
  &__select {
    width: 145px;
    margin-left: 10px;
  }
}
.search-box {
  width: 180px;
}
</style>
