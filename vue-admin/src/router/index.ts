import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Layout from '../views/Layout.vue'
import ArticleAdd from '../views/Article/ArticleAdd.vue'
import Login from '../views/Login/Login.vue'
import CategoryAdd from '../views/Category/CategoryAdd.vue'
const routes: Array<RouteRecordRaw> = [
  {
    path: '/home',
    name: 'Home',
    component: Layout,
    meta: { name: "首页" },
    children: [
      {
      path: 'articleAdd',
      name: 'ArticleAdd',
      component: () => import('@/views/Article/ArticleAdd.vue'),
      meta: { name: "添加文章" }
    },
    {
      path: 'categoryAdd',
      name: 'CategoryAdd',
      component: () => import('@/views/Category/CategoryAdd.vue'),
      meta: { name: "添加类别" }
    },
  ]
  },

  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  // {
  //   path: '/about',
  //   name: 'About',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  // }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
