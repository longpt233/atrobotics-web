import axios from 'axios';
class HomeApiService {
  async getProductList() {
    let response;
    try {
      response = await axios.get(
        'http://atroboticsvn.com/api/v1/user/products?limit=3&offset=1',
      );
    } catch (error) {
      return;
    }
    return response;
  }
}

const homeApiService = new HomeApiService();

export default homeApiService;
