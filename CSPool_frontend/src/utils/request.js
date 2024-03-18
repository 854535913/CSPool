import axios from 'axios';

const request = axios.create({
    baseURL: '/api', // 你的API基础URL
});
export default request;
