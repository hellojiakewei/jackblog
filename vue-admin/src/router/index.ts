import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Layout from '../views/Layout.vue'
import Login from '../views/Login/Login.vue'
import {useStore} from 'vuex'
const store = useStore()
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
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})
router.beforeEach((to,from,next)=>{
  // console.log(store.state.userinfo)
  console.log("to",to)
  console.log("from",from)
  console.log("next",next)
  next()
})

export default router
