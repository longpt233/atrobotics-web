import { createStore } from 'vuex';
import homeModule from './modules/home.module';

export default createStore({
  modules: {
    home: homeModule,
  },
});
