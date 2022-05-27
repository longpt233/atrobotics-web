import axios from 'axios';

const options = {
  baseURL: process.env.VUE_APP_API_URL,
  responseType: 'json',
};

const axiosInstance = axios.create(options);

export default axiosInstance;
