import axiosService from '@/plugins/axios';

class OrderApiService {
  async getOrderList(loginForm) {
    return await axiosService.post('/user/auth/orders', loginForm);
  }
}

const orderApiService = new OrderApiService();
export default orderApiService;
