import axiosService from '@/plugins/axios';

class AuthApiService {
  async login(loginForm) {
    return await axiosService.post('/user/login', loginForm);
  }
  async register(registerForm) {
    return await axiosService.post('/user/register', registerForm);
  }
}

const authApiService = new AuthApiService();
export default authApiService;
