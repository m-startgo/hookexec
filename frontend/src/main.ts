import 'normalize.css';
import 'simplebar-vue/dist/simplebar.min.css';

import '@src/assets/main.scss';

import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { SetAxiosConfig } from '@src/api/http';

import App from '@src/App.vue';
import router from '@src/router';

const app = createApp(App);
app.use(createPinia());
app.use(router);
app.mount('#vue-app');

SetAxiosConfig();

// 所有资源全部加载完毕之后再开始进行一些骚操作
window.onload = () => {};
