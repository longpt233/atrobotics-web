import productApiService from '@/services/product.api.service';

const productListModule = {
  namespaced: true,
  state: {
    productSearchResults: [],
    productList: [],
    brandList: [],
    filter: null,
  },
  getters: {
    productSearchResults: (state) => state.productSearchResults,
    productList: (state) => state.productList,
    brandList: (state) => state.brandList,
    filter: (state) => state.filter,
  },
  mutations: {
    SET_PRODUCT_SEARCH_RESULTS(state, productSearchResults) {
      state.productSearchResults = productSearchResults;
    },
    SET_PRODUCT_LIST(state, productList) {
      state.productList = productList;
    },
    SET_BRAND_LIST(state, brandList) {
      state.brandList = brandList;
    },
    SET_FILTER(state, filter) {
      state.filter = filter;
    },
  },
  actions: {
    async searchProduct({ commit }, keyword) {
      try {
        const response = await productApiService.searchProductByKeyword(keyword);
        commit('SET_PRODUCT_SEARCH_RESULTS', response?.data?.data?.data || []);
      } catch {
        commit('SET_PRODUCT_SEARCH_RESULTS', []);
      }
    },
    async getProductList({ commit }, filter) {
      try {
        const response = await productApiService.getProductList(filter);
        commit('SET_PRODUCT_LIST', response?.data?.data || []);
        commit('SET_FILTER', filter);
      } catch {
        commit('SET_PRODUCT_LIST', []);
      }
    },
    async getBrandList({ commit }) {
      try {
        const response = await productApiService.getBrandList();
        commit('SET_BRAND_LIST', response?.data?.data || []);
      } catch {
        commit('SET_BRAND_LIST', []);
      }
    },
  },
};

export default productListModule;
