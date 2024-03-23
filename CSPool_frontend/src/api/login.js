import {ref} from "vue";
import request from '@/utils/request'; // 引入预配置的axios实例
import {useTokenStore} from '@/store/token.js'
import {ElMessage} from "element-plus";
import router from "@/router";
import useUserInfoStore from "@/store/userInfo";
export const loginData = ref({
    username:"",
    password:""
})
export const loginRules = {
    username: [
        { required: true, message: 'username can\'t be empty', trigger: 'blur' },
    ],
    password: [
        { required: true, message: 'password can\'t be empty', trigger: 'blur' }
    ]
}
export const loginService = async () => {
    const tokenStore = useTokenStore();
    try {
        const response = await request.post('/user/login', loginData.value);
        if(response.data.code === 1000){
            ElMessage.success(response.data.msg ? response.data.msg : 'Login success');
            tokenStore.setToken(response.data.data);
            router.push('/')
        } else {
            ElMessage.error(response.data.msg ? response.data.msg : 'Login failed');
        }
    } catch (error) {
        console.error(error);
    }
};
export const clearLoginData = ()=>{
    loginData.value={
        username:'',
        password:''
    }
}

export const logoutService = () => {
    const tokenStore = useTokenStore();
    const userInfoStore = useUserInfoStore();
    tokenStore.removeToken();
    userInfoStore.removeInfo()
    router.push('/user/login')
};