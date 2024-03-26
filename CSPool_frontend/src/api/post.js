import {ref} from "vue";
import request from "@/utils/request";
import {useTokenStore} from '@/store/token.js'
import {ElMessage} from "element-plus";
export const postData = ref({
    title: '',
    description: '',
    link: '',
    tag: ''
})
export const uploadService = async () => {
    const tokenStore = useTokenStore();
    try {
        const response = await request.post('/post/upload', postData.value);
        if(response.data.code === 1000){
            ElMessage.success(response.data.msg ? response.data.msg : 'post success');
        } else {
            ElMessage.error(response.data.msg ? response.data.msg : 'post failed');
        }
    } catch (error) {
        console.error(error);
    }
};
export const tableData = ref({
    postID: '',
    postLike: '',
    postTime: '',
    title: '',
    description: '',
    link: '',
    tag: '',
    authorName: '',
})
export const getPostByTimeService = async () => {
    try {
        const response = await request.get('/post/time');
        if(response.data.code === 1000){
            console.log(response.data.msg);
            return response.data.data.map(post => ({
                postID: post.PostID,
                postLike: post.PostLike,
                postTime: post.PostTime,
                title: post.title,
                description: post.description,
                link: post.link,
                tag: post.tag,
                authorName: post.AuthorName,
            }));
        } else {
            console.log(response.data.msg);
            return [];
        }
    } catch (error) {
        console.error(error);
        return [];
    }
};
export const getPostByLikeService = async () => {
    try {
        const response = await request.get('/post/like');
        if(response.data.code === 1000){
            console.log(response.data.msg);
            return response.data.data.map(post => ({
                postID: post.PostID,
                postLike: post.PostLike,
                postTime: post.PostTime,
                title: post.title,
                description: post.description,
                link: post.link,
                tag: post.tag,
                authorName: post.AuthorName,
            }));
        } else {
            console.log(response.data.msg);
            return [];
        }
    } catch (error) {
        console.error(error);
        return [];
    }
};
export const getPostByUnderreviewService = async () => {
    try {
        const response = await request.get('/post/review');
        if(response.data.code === 1000){
            console.log(response.data.msg);
            return response.data.data.map(post => ({
                postID: post.PostID,
                postTime: post.PostTime,
                title: post.title,
                description: post.description,
                link: post.link,
                tag: post.tag,
                authorName: post.AuthorName,
            }));
        } else {
            console.log(response.data.msg);
            return [];
        }
    } catch (error) {
        console.error(error);
        return [];
    }
};
export const actionPostService = async (vid, action) => {
    try {
        const url = `/post/${vid}/${action}`;
        const response = await request.post(url);
        console.log(response.data.msg);
    } catch (error) {
        console.error(error);
    }
    window.location.reload();
};
export const publishPostService = async (vid) => {
    try {
        const url = `/post/publish/${vid}`;
        const response = await request.post(url);
        console.log(response.data.msg);
    } catch (error) {
        console.error(error);
    }
    window.location.reload();
};