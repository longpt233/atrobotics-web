import axiosService from '@/plugins/axios';

class CheckoutApiService {
  async getItemList() {
    return await axiosService.get('/user/auth/cart/list');
  }
  async addItem(item) {
    return await axiosService.post('/user/auth/cart/add', item);
  }
  async deleteItem(id) {
    return await axiosService.delete(`/user/auth/cart/${id}`);
  }
}

const checkoutApiService = new CheckoutApiService();
export default checkoutApiService;
