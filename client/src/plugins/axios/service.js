export default class ApiService {
  constructor(axiosInstance, baseUrl) {
    this.axiosInstance = axiosInstance;
    this.baseUrl = baseUrl;
  }
}
