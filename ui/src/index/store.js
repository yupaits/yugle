import Vue from 'vue'
import Vuex from 'vuex'
import consts from '../consts'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    key: '',
    settingsKeys: [consts.settingsMenu[0].key]
  },
  getters: {
    key(state) {
      return state.key;
    },
    settingsKeys(state) {
      return state.settingsKeys;
    }
  },
  mutations: {
    setKey(state, key) {
      state.key = key;
    },
    setSettingsKeys(state, keys) {
      state.settingsKeys = keys;
    }
  },
  actions: {
    setKey({commit}, key) {
      commit('setKey', key);
    },
    setSettingsKeys({commit}, keys) {
      commit('setSettingsKeys', keys);
    }
  }
})
