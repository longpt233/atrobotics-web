import axiosService from '@/plugins/axios';

class UserApiService {
  async getUserInfo() {
    return await axiosService.get('/user/auth/info');
  }
}

const userApiService = new UserApiService();
export default userApiService;
