import { createRouter, createWebHistory } from 'vue-router'
import LoginRegisterView from '../views/LoginRegister.vue';
import HomeView from '../views/Home.vue';
import UploadPostView from '../views/PostUpload.vue';
import TimeListView from '../views/PostShow.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomeView,
    children: [
      { path: '/post/upload', component: UploadPostView },
      { path: '/post/list', component: TimeListView },
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
