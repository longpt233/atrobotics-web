<template>
  <div class="container-fluid">
    <div
      class="navigation-bar-content container d-flex justify-content-between align-items-center"
    >
      <div class="navigation d-flex justify-content-between align-items-center">
        <img src="@/assets/images/logo-nobg.jpg" />
        <router-link class="router-link" to="/home">
          <span> {{ $t('app.header.navigationBar.home') }} </span>
        </router-link>
        <router-link class="router-link" to="/product">
          <span> {{ $t('app.header.navigationBar.productList') }} </span>
        </router-link>
        <a :href="fanpageUrl" target="_blank">
          <span>{{ $t('app.header.navigationBar.fanpage') }}</span>
        </a>
      </div>
      <SearchBox
        :placeholder="$t('app.header.navigationBar.searchBox.placeholder')"
        @on-change-keyword="getProductSuggestions"
        @on-select-keyword="onSelectKeyword"
      />
    </div>
  </div>
</template>

<script>
import { debounce, uniqBy } from 'lodash';
import { mapActions, mapGetters } from 'vuex';

import SearchBox from '@/components/SearchBox.vue';
import { PageName } from '@/common/constants';
import { SEARCH_DEBOUNCE_TIME } from '../../constants';
export default {
  components: { SearchBox },
  setup() {
    const fanpageUrl = process.env.VUE_APP_FANEPAGE_URL;
    return { fanpageUrl };
  },
  computed: {
    ...mapGetters('productList', ['productSearchResults']),
  },
  methods: {
    ...mapActions('productList', ['searchProduct']),
    getProductSuggestions({ keyword, callback }) {
      debounce(() => {
        this.searchProduct(keyword);
        const suggestionResults = uniqBy(
          [
            { value: keyword },
            ...this.productSearchResults.map((result) => {
              return { value: result.name };
            }),
          ],
          'value',
        );
        callback(suggestionResults);
      }, SEARCH_DEBOUNCE_TIME)();
    },
    onSelectKeyword(keyword) {
      const filter = {};
      if (keyword.trim() !== '') {
        filter.q = keyword;
      }

      this.$router.push({
        name: PageName.PRODUCT_LIST_PAGE,
        query: filter,
      });
    },
  },
};
</script>

<style lang="scss" scoped>
.container-fluid {
  height: 45px;
  background-color: $color-white;
  box-shadow: 0 7px 4px -7px $color-grey;
  .navigation-bar-content {
    height: 100%;
  }
}

.navigation {
  gap: 20px;
  width: fit-content;
  img {
    width: auto;
    height: 40px;
  }
  a {
    color: $color-black;
    text-decoration: none;
  }
  span {
    font-size: 16px;
    font-weight: 500;
    color: $color-black;
    &:hover {
      color: $color-primary;
    }
  }
}
</style>
