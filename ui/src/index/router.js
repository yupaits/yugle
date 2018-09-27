import Vue from 'vue'
import Router from 'vue-router'
import 'vue-cookies'
import consts from '../consts'
import store from './store'

import Index from './views/Index'
import Dashboard from './views/Dashboard'

import Settings from './views/Settings'
import UserManage from './views/settings/UserManage'
import RoleManage from './views/settings/RoleManage'
import PermissionManage from './views/settings/PermissionManage'
import ModifyPassword from './views/settings/ModifyPassword'

import Profile from './views/Profile'
import Picture from './views/Picture'
import BingPicture from './views/picture/BingPicture'
import ShotOnOnePlus from './views/picture/ShotOnOnePlus'
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
          path: 'profile',
          component: Profile
        },
        {
          path: 'settings',
          component: Settings,
          redirect: '/settings/user',
          children: [
            {
              path: 'user',
              component: UserManage
            },
            {
              path: 'role',
              component: RoleManage
            },
            {
              path: 'permission',
              component: PermissionManage
            },
            {
              path: 'password',
              component: ModifyPassword
            },
          ]
        },
        {
          path: 'picture',
          component: Picture,
          children: [
            {
              path: 'bing',
              component: BingPicture,
              meta: {key: 'bing-picture'}
            },
            {
              path: 'shot_on_oneplus',
              component: ShotOnOnePlus,
              meta: {key: 'shot-on-oneplus'}
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
  if (to.path.startsWith('/settings')) {
    store.dispatch('setSettingsKeys', [to.path.split('/')[2]]);
  }
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