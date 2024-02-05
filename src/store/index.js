import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    role: "",
    stream:""
  },
  mutations: {
    setRole(state, role) {
      state.role = role;
    },
    setStream(state, stream) {
      state.stream = stream;
    },
  },
  actions: {},
  modules: {},
});
