import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Antd from 'ant-design-vue'
import api from '../api'
import consts from '../consts'

import 'ant-design-vue/dist/antd.min.css'
import '../style/index.css'
import '../style/icon.css'

Vue.config.productionTip = false

Vue.use(Antd)

Vue.prototype.$consts = consts
Vue.prototype.$api = api

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
