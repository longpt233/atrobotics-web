import axiosConfig from '../config/axiosConfig';
class AuthApiService {
  /**
   * Login API for user and admin
   * @param {String} email
   * @param {String} password
   * @returns response of API
   */
  async login(email, password) {
    let response;
    try {
      response = axiosConfig.post('/user/login', {
        email: email,
        password: password,
      });
    } catch (error) {
      console.log(error);
      return;
    }
    return response;
  }
  
  /**
   * Register API for user
   * @param {Object} registerForm 
   * @returns response of API
   */
  async register(registerForm) {
    let response;
    try {
      response = axiosConfig.post('/user/register', registerForm);
    } catch (error) {
      console.log(error);
      return;
    }
    return response;
  }
}

const authApiService = new AuthApiService();
export default authApiService;
