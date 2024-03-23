import axios from 'axios';
import {useTokenStore} from '@/store/token.js'
import router from '@/router'
import {ElMessage} from "element-plus";

const baseURL = '/api';
const request = axios.create({baseURL});
export default request;

request.interceptors.request.use(
    (config)=>{
        const tokenStore = useTokenStore();
        if(tokenStore.token){
            config.headers.Authorization = `Bearer ${tokenStore.token}`;
        }
        return config;
    },
    (err)=>{
        Promise.reject(err)
    }
)
request.interceptors.response.use(
    response => {
        if(response.data.code===1000){
            return response;
        }else if (response.data.code===1019){
            ElMessage.error('请先登录')
        }else{
            ElMessage.error(response.data.msg)
        }
        return Promise.reject(response)
    },
    err => {
        ElMessage.error('server error')
        return Promise.reject(err);
    }
)
