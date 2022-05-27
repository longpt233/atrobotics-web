import productApiService from '@/services/product.api.service';

const productDetailModule = {
  namespaced: true,
  state: {
    productDetail: null,
  },
  getters: {
    productDetail: (state) => state.productDetail,
  },
  mutations: {
    SET_PRODUCT_DETAIL(state, productDetail) {
      state.productDetail = productDetail;
    },
  },
  actions: {
    async getProductDetail({ commit }, productId) {
      try {
        const response = await productApiService.getProductDetail(productId);
        commit('SET_PRODUCT_DETAIL', response?.data?.data || null);
      } catch {
        commit('SET_PRODUCT_DETAIL', null);
      }
    },
  },
};

export default productDetailModule;
