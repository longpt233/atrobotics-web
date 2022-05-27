import { createStore } from 'vuex';
import homeModule from './modules/home.module.store';
import productListModule from './modules/productList.module.store';
import productDetailModule from './modules/productDetail.module.store';
import authModule from './modules/auth.module.store';
import userModule from './modules/user.module.store';
import checkoutModule from './modules/checkout.module.store';
// import orderModule from './modules/orderModule.module.store';

export default createStore({
  modules: {
    home: homeModule,
    productList: productListModule,
    productDetail: productDetailModule,
    auth: authModule,
    user: userModule,
    checkout: checkoutModule,
    // order: orderModule,
  },
});
