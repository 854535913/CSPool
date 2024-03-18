import {ref} from "vue";
import request from '@/utils/request'; // 引入预配置的axios实例

export const registerData = ref({
    username:"",
    password:"",
    re_password:""
})

export const registerRules = {
    username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
    ],
    password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 5, max: 16, message: '长度为5~16位非空字符', trigger: 'blur' }
    ],
    re_password: [
        { required: true, message: '请再次输入密码', trigger: 'blur' },
        { min: 5, max: 16, message: '长度为5~16位非空字符', trigger: 'blur' }
    ]
}

export const registerService = async () => {
    try {
        const response = await request.post('/register', registerData.value);
        // 注意这里使用 response.data.code 和 response.data.msg 访问响应数据
        if(response.data.code === 1000){
            alert(response.data.msg ? response.data.msg : '注册成功');
        } else {
            alert(response.data.msg ? response.data.msg : '注册失败');
        }
    } catch (error) {
        console.error(error);
        alert('发生错误');
    }
};
