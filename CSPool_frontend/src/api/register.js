import {ref} from "vue";
import request from '@/utils/request';
import {ElMessage} from "element-plus"; // 引入预配置的axios实例

export const registerData = ref({
    username:"",
    password:"",
    re_password:"",
    invitation:""
})

export const registerRules = {
    username: [
        { required: true, message: 'username can\'t be empty', trigger: 'blur' },
    ],
    password: [
        { required: true, message: 'password can\'t be empty', trigger: 'blur' },
        { min: 4, max: 16, message: 'password must be 4 to 16 non-blank characters in length', trigger: 'blur' }
    ],
    re_password: [
        { required: true, message: 'password can\'t be empty', trigger: 'blur' },
        { min: 4, max: 16, message: 'password must be 4 to 16 non-blank characters in length', trigger: 'blur' }
    ]
}

export const registerService = async () => {
    try {
        const response = await request.post('/user/register', registerData.value);
        // 注意这里使用 response.data.code 和 response.data.msg 访问响应数据
        if(response.data.code === 1000){
            ElMessage.success(response.data.msg ? response.data.msg : 'register success');
        } else {
            ElMessage.error(response.data.msg ? response.data.msg : 'register failed');
        }
    } catch (error) {
        console.error(error);
    }
};

export const clearRegisterData = ()=>{
    registerData.value={
        username:'',
        password:'',
        rePassword:'',
        invitation:''
    }
}