import checkoutApiService from '@/services/checkout.api.service';

const checkoutModule = {
  namespaced: true,
  state: {
    itemList: [],
    totalAmount: 0,
  },
  getters: {
    itemList: (state) => state.itemList,
    totalAmount: (state) => state.totalAmount,
  },
  mutations: {
    SET_ITEM_LIST(state, itemList) {
      state.itemList = itemList;
    },
    UPDATE_QUANTITY(state, item) {
      if (item.quantity > 0) {
        // update quantity
        state.itemList.forEach((checkoutItem) => {
          if (checkoutItem.id === item.id && checkoutItem.color === item.color) {
            checkoutItem.quantity = item.quantity;
          }
        });
      } else {
        // delete item
        state.itemList = state.itemList.filter((checkoutItem) => {
          return checkoutItem.id !== item.id && checkoutItem.color !== item.color;
        });
      }
    },
    DELETE_ITEM(state, item) {
      state.itemList = state.itemList.filter((checkoutItem) => {
        return checkoutItem.id !== item.id;
      });
    },
    SET_TOTAL_AMOUNT(state, amount) {
      state.totalAmount += amount;
    },
  },
  actions: {
    async getItemList({ commit }) {
      try {
        const response = await checkoutApiService.getItemList();
        commit('SET_ITEM_LIST', response?.data?.data || []);
      } catch {
        commit('SET_ITEM_LIST', []);
      }
    },
    async updateQuantity({ commit }, item) {
      try {
        const response = await checkoutApiService.addItem(item);
        commit('UPDATE_QUANTITY', response?.data?.data || []);
      } catch {
        commit('UPDATE_QUANTITY', []);
      }
    },
    async deleteItem({ commit }, id) {
      try {
        const response = await checkoutApiService.deleteItem(id);
        commit('DELETE_ITEM', response?.data?.data || []);
      } catch {
        commit('DELETE_ITEM', []);
      }
    },
    setTotalAmount({ commit }, amount) {
      commit('SET_TOTAL_AMOUNT', amount);
    },
  },
};

export default checkoutModule;
