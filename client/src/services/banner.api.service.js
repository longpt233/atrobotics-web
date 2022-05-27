import axiosService from '@/plugins/axios';

class BannerApiService {
  async getBannerList() {
    return await axiosService.get('/user/banners/top-3-newest');
  }
}

const bannerApiService = new BannerApiService();
export default bannerApiService;
