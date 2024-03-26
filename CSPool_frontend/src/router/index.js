import { createRouter, createWebHistory } from 'vue-router'
import LoginRegisterView from '../views/LoginRegister.vue';
import LayoutView from '../views/Layout.vue';
import UploadPostView from '../views/PostUpload.vue';
import TimeListView from '../views/PostShow.vue';
import ReviewView from '../views/PostReview.vue';
import HomeView from '../views/Homepage.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: LayoutView,
    redirect:'/home',
    children: [
      { path: '/home', component: HomeView },
      { path: '/post/upload', component: UploadPostView },
      { path: '/post/list', component: TimeListView },
      { path: '/post/review', component: ReviewView },
    ]
  },
  {
    path: '/user/login', // 定义路由路径
    component: LoginRegisterView, // 指定路由组件
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes:routes
})

export default router
