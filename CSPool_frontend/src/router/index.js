import { createRouter, createWebHashHistory } from 'vue-router'
import RegisterView from '../views/Register.vue';
import HomeView from '../views/Home.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomeView,
  },
  {
    path: '/register', // 定义路由路径
    name: 'register', // 定义路由名称
    component: RegisterView, // 指定路由组件
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
