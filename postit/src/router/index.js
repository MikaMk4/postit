import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [
  {
    path: '/:pathMatch(.*)*',
    name: 'not-found',
    component: () => import('../views/NotFoundView.vue')
  },
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/board/:bid/post/:pid',
    name: 'post',
    component: () => import('../views/PostView.vue')
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/LoginView.vue')
  },
  {
    path: '/settings',
    name: 'settings',
    component: () => import('../views/SettingsView.vue')
  },
  {
    path: '/boards',
    name: 'boards',
    component: () => import('../views/BoardListView.vue')
  },
  {
    path: '/b/:id',
    name: 'board',
    component: () => import('../views/BoardView.vue')
  },
  {
    path: '/b/:id/create',
    name: 'create-post',
    component: () => import('../views/PostCreateView.vue')
  },
  {
    path: '/create-board',
    name: 'create-board',
    component: () => import('../views/BoardCreateView.vue')
  },
  {
    path: '/b/:id/edit',
    name: 'edit-board',
    component: () => import('../views/BoardEditView.vue')
  },
  {
    path: '/u/:id',
    name: 'user',
    component: () => import('../views/ProfileView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
