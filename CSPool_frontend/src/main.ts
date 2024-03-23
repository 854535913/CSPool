import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import {createPinia} from 'pinia'
import { createPersistedState } from 'pinia-persistedstate-plugin'


const app = createApp(App);
const pinia = createPinia();
const persist = createPersistedState();
app.use(store)
    .use(router)
    .use(ElementPlus)
    .use(pinia)
    .mount('#app');
pinia.use(persist)