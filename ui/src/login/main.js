import Vue from 'vue'
import App from './App.vue'
import router from './router'
import VueCookies from 'vue-cookies'
import consts from "../consts"
import request from '../request'

import 'bootstrap/dist/css/bootstrap.min.css'

Vue.config.productionTip = false

Vue.use(VueCookies)

Vue.prototype.auth = request.auth
Vue.prototype.consts = consts

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
