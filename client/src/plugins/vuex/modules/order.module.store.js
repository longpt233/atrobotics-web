import orderApiService from '@/services/order.api.service';

const orderModule = {
  namespaced: true,
  state: {
    orderList: [],
  },
  getters: {
    orderList: (state) => state.orderList,
  },
  mutations: {
    SET_ORDER_LIST(state, orderList) {
      state.orderList = orderList;
    },
  },
  actions: {
    async getOrderList({ commit }) {
      try {
        const response = await orderApiService.getOrderList();
        commit('SET_ORDER_LIST', response?.data?.data || '');
      } catch (error) {
        commit('SET_ORDER_LIST', '');
      }
    },
  },
};

export default orderModule;
