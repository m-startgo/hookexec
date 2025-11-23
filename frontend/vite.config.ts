import { fileURLToPath, URL } from 'node:url';
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import vueDevTools from 'vite-plugin-vue-devtools';
import path from 'node:path';

import legacy from '@vitejs/plugin-legacy';

const __dirname = fileURLToPath(new URL('.', import.meta.url));
console.info('Vue3项目启动,__dirname', __dirname);

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    legacy({
      targets: ['since 2020', 'not dead'],
    }),
    vueDevTools(),
  ],
  resolve: {
    alias: {
      '@src': path.resolve(__dirname, './src'),
    },
  },
  server: {
    host: '0.0.0.0',
    port: 9989,
    proxy: {
      // 将所有以 /api 开头的请求转发到后端服务
      '/api': {
        target: 'http://localhost:9900/',
        changeOrigin: true,
        secure: false,
        ws: true,
        // 如果后端不需要 /api 前缀，可以取消注释下面的 rewrite
        // rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
});
