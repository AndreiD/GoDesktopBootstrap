import Vue from 'vue'
import Router from 'vue-router'
import Home from '../src/components/Home'
import Settings from '../src/components/Settings'

Vue.use(Router)

export const router = new Router({
  mode: 'hash',
  base: process.env.BASE_URL,
  routes: [{
    path: '/',
    name: 'home',
    component: Home,
  },
  {
    path: '/settings',
    name: 'settings',
    component: Settings,
  },
  ]
})

export default router
