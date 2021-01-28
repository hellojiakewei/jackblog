import { createStore } from 'vuex'
import createPersistedState from 'vuex-persistedstate'
import { loginApi } from '@/api/user'
import { Iuserinfo } from '@/model/user'
export default createStore({
  state: {
    userinfo: {},
    token:''
  },
  mutations: {
    SET_USER: (state, data) => {
      state.userinfo = data
    }
  },
  actions: {
    login({ commit }, params:Iuserinfo) {
      return new Promise((resolve, reject) => {
        loginApi(params).then(res => {
          commit('SET_USER', res.data)
          resolve(res)
        }).catch(error => {
          reject(error)
        })
      })
    }

  },
  modules: {

  },
  plugins: [createPersistedState()],
})

// const modules = (modulesFiles => {
//   return modulesFiles.keys().reduce((modules, modulePath) => {
//       const moduleName = modulePath.replace(/^\.\/(.*)\.\w+$/, '$1');
//       const value = modulesFiles(modulePath);
//       modules[moduleName] = value.default;
//       return modules;
//   }, {});
// })(require.context('./module', false, /\.js$/));