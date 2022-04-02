import homeApiService from '@/services/home.api.service';

const homeModule = {
  namespaced: true,
  state: {
    productList: [],
  },
  getters: {
    productList: (state) => state.productList,
  },
  mutations: {
    SET_PRODUCT_LIST(state, productList) {
      state.productList = productList;
    },
  },
  actions: {
    async getProductList({ commit }) {
      const response = await homeApiService.getProductList();
      if (response) {
        commit('SET_PRODUCT_LIST', response?.data?.data || []);
        console.log(response?.data?.data);
      } else {
        commit('SET_PRODUCT_LIST', []);
      }
    },
  },
};
export default homeModule;
