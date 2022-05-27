import axiosService from '@/plugins/axios';
import qs from 'qs';

class ProductApiService {
  async searchProductByKeyword(keyword) {
    return await axiosService.get(`/user/products?q=${keyword}`);
  }
  async getProductList(filter) {
    return await axiosService.get(`/user/products?${qs.stringify(filter)}`);
  }
  async getCategoryList() {
    return await axiosService.get('/user/categories');
  }
  async getBrandList() {
    return await axiosService.get('/user/all-brand');
  }
  async getProductListByCategory(categoryId) {
    return await axiosService.get(`/user/products/by-category?categoryId=${categoryId}`);
  }
  async getProductDetail(productId) {
    return await axiosService.get(`/user/products/${productId}`);
  }
}

const productApiService = new ProductApiService();
export default productApiService;
