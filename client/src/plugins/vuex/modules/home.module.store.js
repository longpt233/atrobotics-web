import bannerApiService from '@/services/banner.api.service';
import productApiService from '@/services/product.api.service';

const homeModule = {
  namespaced: true,
  state: {
    bannerList: [],
    categoryList: [],
    productListByCategory: [],
  },
  getters: {
    bannerList: (state) => state.bannerList,
    categoryList: (state) => state.categoryList,
    productListByCategory: (state) => state.productListByCategory,
  },
  mutations: {
    SET_BANNER_LIST(state, bannerList) {
      state.bannerList = bannerList;
    },
    SET_CATEGORY_LIST(state, categoryList) {
      state.categoryList = categoryList;
    },
    SET_PRODUCT_LIST_BY_CATEGORY(state, productListByCategory) {
      state.productListByCategory.push(productListByCategory);
    },
  },
  actions: {
    async getBannerList({ commit }) {
      try {
        const response = await bannerApiService.getBannerList();
        commit('SET_BANNER_LIST', response?.data?.data || []);
      } catch {
        commit('SET_BANNER_LIST', []);
      }
    },
    async getCategoryList({ commit }) {
      try {
        const response = await productApiService.getCategoryList();
        commit('SET_CATEGORY_LIST', response?.data?.data || []);
      } catch {
        commit('SET_CATEGORY_LIST', []);
      }
    },
    async getProductListByCategory({ commit }, category) {
      try {
        const response = await productApiService.getProductListByCategory(category.id);
        let productListByCategory = response?.data?.data;
        productListByCategory.category = category.name;
        commit('SET_PRODUCT_LIST_BY_CATEGORY', productListByCategory || []);
      } catch {
        commit('SET_PRODUCT_LIST_BY_CATEGORY', []);
      }
    },
  },
};

export default homeModule;
