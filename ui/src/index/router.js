import Vue from 'vue'
import Router from 'vue-router'
import 'vue-cookies'
import consts from '../consts'
import store from './store'

import Index from './views/Index'
import Dashboard from './views/Dashboard'
import Picture from './views/Picture'
import BingPicture from './views/picture/BingPicture'
import Video from './views/Video'
import Movie from './views/Movie'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/',
      component: Index,
      redirect: '/dashboard',
      children: [
        {
          path: 'dashboard',
          component: Dashboard
        },
        {
          path: 'picture',
          component: Picture,
          children: [
            {
              path: 'bing',
              component: BingPicture,
              meta: {key: 'bing-picture'}
            }
          ]
        },
        {
          path: 'video',
          component: Video,
          meta: {key: 'video'}
        },
        {
          path: 'movie',
          component: Movie,
          meta: {key: 'movie'}
        }
      ]
    },
  ]
})

router.beforeEach((to, from, next) => {
  const key = to.meta.key;
  store.dispatch('setKey', key || '');
  const accessToken = window.$cookies.get(consts.token.cookieName);
  const tokenValid = accessToken !== undefined && accessToken !== null && accessToken.trim() !== '';
  if (tokenValid) {
    next();
  } else {
    window.location.href = '/login';
  }
  // next();
})

export default router