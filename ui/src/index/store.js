import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    key: ''
  },
  getters: {
    key(state) {
      return state.key;
    }
  },
  mutations: {
    setKey(state, key) {
      state.key = key;
    }
  },
  actions: {
    setKey({commit}, key) {
      commit('setKey', key);
    }
  }
})
