import homeApiService from '@/services/home.api.service';

const homeModule = {
  namespaced: true,
  state: {
    productList: [],
    userInfo: {}
  },
  getters: {
    productList: (state) => state.productList,
    userInfo: (state) => state.userInfo,
  },
  mutations: {
    SET_PRODUCT_LIST(state, productList) {
      state.productList = productList;
    },
    SET_USER_PROFILE(state, userInfo){
      state.userInfo = userInfo;
    },
  },
  actions: {
    async getProductList({ commit },{limit, offset}) {
      const response = await homeApiService.getProductList(limit, offset);
      if (response) {
        commit('SET_PRODUCT_LIST', response?.data?.data || []);
        console.log(response?.data?.data);
      } else {
        commit('SET_PRODUCT_LIST', []);
      }
    },

    async getUserProfile({commit}, token){
      const response = await homeApiService.getUserProfile(token);
      if(response){
        commit('SET_USER_PROFILE',response?.data?.data || {});
        console.log("User profile: ",response?.data?.data);
      }else {
        commit('SET_USER_PROFILE', {})
      }
    }
  },
};
export default homeModule;
