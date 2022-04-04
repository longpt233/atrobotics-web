import axios from "axios";

const instance = axios.create({
    baseURL:'http://atroboticsvn.com/api/v1'
});

export default instance;