import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/index.vue'),
    meta: { guest: true },
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/register/index.vue'),
    meta: { guest: true },
  },
  {
    path: '/',
    component: () => import('@/views/layout/index.vue'),
    children: [
      {
        path: '',
        name: 'Home',
        component: () => import('@/views/home/index.vue'),
      },
      {
        path: 'create',
        name: 'CreateArticle',
        component: () => import('@/views/article/create.vue'),
        meta: { auth: true },
      },
      {
        path: 'article/:id',
        name: 'ArticleDetail',
        component: () => import('@/views/article/detail.vue'),
      },
      {
        path: 'friends',
        name: 'Friends',
        component: () => import('@/views/friends/index.vue'),
        meta: { auth: true },
      },
      {
        path: 'admin',
        name: 'Admin',
        component: () => import('@/views/admin/index.vue'),
        meta: { auth: true, admin: true },
      },
    ],
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/',
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, _from, next) => {
  const userStore = useUserStore()

  if (to.meta.guest && userStore.isLoggedIn) {
    next('/')
    return
  }

  if (to.meta.auth && !userStore.isLoggedIn) {
    next('/login')
    return
  }

  if (to.meta.admin && !userStore.isAdmin) {
    next('/')
    return
  }

  next()
})

export default router
