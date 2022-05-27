import authApiService from '@/services/auth.api.service';
import axiosService from '@/plugins/axios';

const authModule = {
  namespaced: true,
  state: {
    token: '',
  },
  getters: {
    token: (state) => state.token,
  },
  mutations: {
    SET_TOKEN(state, token) {
      state.token = token;
      localStorage.setItem('token', token);
      axiosService.defaults.headers.common['Authorization'] = `Bearer ${token}`;
    },
    REMOVE_TOKEN(state) {
      state.token = '';
      localStorage.setItem('token', '');
      axiosService.defaults.headers.common['Authorization'] = '';
    },
  },
  actions: {
    async login({ commit, dispatch }, loginForm) {
      try {
        const response = await authApiService.login(loginForm);
        commit('SET_TOKEN', response?.data?.data || '');
        await dispatch('user/getUserInfo', null, { root: true });
      } catch (error) {
        commit('SET_TOKEN', '');
        return error.response?.data?.data;
      }
    },
    async register(_context, registerForm) {
      try {
        await authApiService.register(registerForm);
      } catch (error) {
        return error.response?.data?.message;
      }
    },
    async autoLogin({ commit, dispatch }, token) {
      commit('SET_TOKEN', token);
      await dispatch('user/getUserInfo', null, { root: true });
    },
    async logout({ commit, dispatch }) {
      commit('REMOVE_TOKEN');
      await dispatch('user/removeUserInfo', null, { root: true });
    },
  },
};

export default authModule;
