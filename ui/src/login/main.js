import Vue from 'vue'
import App from './App.vue'
import router from './router'
import Antd from 'ant-design-vue'
import VueCookies from 'vue-cookies'
import consts from "../consts"
import request from '../request'

import 'ant-design-vue/dist/antd.min.css'
import '../style/index.css'

Vue.config.productionTip = false

Vue.use(Antd)
Vue.use(VueCookies)

Vue.prototype.$auth = request.auth
Vue.prototype.$consts = consts

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
