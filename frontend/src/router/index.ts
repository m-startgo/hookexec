import { createRouter } from 'vue-router';
import { createWebHistory } from 'vue-router';

const routes = [
  {
    path: '/',
    name: 'HomePage',
    component: () => import('@src/views/HomePage.vue'),
  },
  {
    path: '/about',
    name: 'AboutPage',
    component: () => import('@src/views/AboutPage.vue'),
  },
  {
    path: '/component',
    name: 'ComponentPage',
    component: () => import('@src/views/ComponentPage.vue'),
  },
  {
    path: '/test',
    name: 'TestPage',
    // route level code-splitting
    // this generates a separate chunk (About.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import('@src/views/TestPage.vue'),
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@src/views/NotFound.vue'),
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
