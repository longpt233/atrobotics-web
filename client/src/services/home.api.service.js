import axiosConfig from "../config/axiosConfig"
class HomeApiService {
  /**
   * API get list product with limit ... and offset ...
   * @returns response of API
   */
  async getProductList(limit, offset) {
    let response;
    try {
      response = await axiosConfig.get(
        `/user/products?limit=${limit}&offset=${offset}`,
      );
    } catch (error) {
      return;
    }
    return response;
  }

  /**
   * Get user profile API
   * @param {String} token 
   * @returns response of API
   */
  async getUserProfile(token){
    let response;
    try{
      response = axiosConfig.get(
        "/user/auth/info",
        {
          headers:{
            "Authorization": `Bearer ${token}`
          }
        }
      );
      console.log("get user info token: ", token)
    }catch(error){
      return;
    }
    return response;
  }
}

const homeApiService = new HomeApiService();

export default homeApiService;
