import Vue from "vue";
import Vuex from "vuex";
// import auth from "./modules/auth";
import modules from "./modules/modules"
import createPersistedState from "vuex-persistedstate";

Vue.use(Vuex);

export default new Vuex.Store({
  modules,
  plugins: [createPersistedState()]
});
