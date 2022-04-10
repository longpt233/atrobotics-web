import { createStore } from 'vuex';
import authModule from './modules/auth.module';
import homeModule from './modules/home.module';

export default createStore({
  modules: {
    home: homeModule,
    auth: authModule,
  },
});
