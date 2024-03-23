import request from "@/utils/request";
import {ElMessage} from "element-plus";
import useUserInfoStore from "@/store/userInfo";

export const getUserInfoService = async () => {
    const userInfoStore = useUserInfoStore();
    try {
        const response = await request.get('/user/info');
        if(response.data.code === 1000){
            let level =''
            if (response.data.data.Level===1){
                level ='Owner'
            }else if (response.data.data.Level===2){
                level ='Admin'
            }else{
                level ='User'
            }
            userInfoStore.setInfo({
                username: response.data.data.Username,
                level: level
            });
        } else {
            console.log(response.data.msg);
            return [];
        }
    } catch (error) {
        console.error(error);
        return [];
    }
};