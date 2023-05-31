
const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/Index/Index.vue') }
    ]
  },
  {
    path: '/index',
    name: 'index',
    component: () => import('layouts/Layout.vue'),
    children: [
      { path: '', name: 'index_child', component: () => import('pages/Index/Index2.vue') }
    ]
  },
  {
    path: '/detail',
    component: () => import('layouts/Layout.vue'),
    children: [
      { path: '', component: () => import('pages/detail.vue') }
    ]
  },
  {
    path: '/MyHistory',
    component: () => import('layouts/Layout.vue'),
    children: [
      { path: '', component: () => import('pages/MyHistory.vue') }
    ]
  },
  {
    path: '/login',
    component: () => import('pages/login.vue')
  },
  {
    path: '/register',
    component: () => import('pages/register.vue')
  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes
