import Vue from 'vue'
import VueRouter from 'vue-router'
import 'vue-cookies'
import consts from "../consts"

import Login from './views/Login'

Vue.use(VueRouter)

const router = new VueRouter({
  routes: [
    {
      path: '/',
      component: Login
    }
  ]
});

router.beforeEach((to, from, next) => {
  const accessToken = window.$cookies.get(consts.token.cookieName);
  const tokenValid = accessToken !== undefined && accessToken !== null && accessToken.trim() !== '';
  if (tokenValid) {
    window.location.href = '/';
  } else {
    next();
  }
})

export default router