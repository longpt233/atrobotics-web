import userApiService from '@/services/user.api.service';

const userModule = {
  namespaced: true,
  state: {
    userInfo: null,
  },
  getters: {
    userInfo: (state) => state.userInfo,
  },
  mutations: {
    SET_USER_INFO(state, userInfo) {
      state.userInfo = userInfo;
    },
  },
  actions: {
    async getUserInfo({ commit }) {
      try {
        const response = await userApiService.getUserInfo();
        commit('SET_USER_INFO', response?.data?.data || null);
      } catch {
        commit('SET_USER_INFO', null);
      }
    },
    async removeUserInfo({ commit }) {
      commit('SET_USER_INFO', null);
    },
  },
};

export default userModule;
