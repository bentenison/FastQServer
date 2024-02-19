// import * as types from '../types'
import axios from "axios";
import router from "../../router";
import * as types from "../types";
// import MQL from '@/plugins/mql.js'
import { getTokenPayload, storeTokens } from "../../plugins/utils";
// export default {
//   namespaced: true,

// };
export const state = {
  token: sessionStorage.getItem("user-token") || null,
  branchToken: sessionStorage.getItem("branch-token") || null,
  refreshtoken: sessionStorage.getItem("refresh-token") || null,
  user: sessionStorage.getItem("user") || null,
  status: "",
  loader: false,
  counterUser: null,
  activeCounter: null,
  activeTicket: null,
  activeTickets: [],
};
export const getters = {
  isAuthenticated: (state) => !!state.token,
  authStatus: (state) => state.status,
  getRefreshToken: (state) => state.refreshToken,
  getIdToken: (state) => state.token,
  getUser: (state) => state.user,
};
export const mutations = {
  [types.MUTATE_AUTH_REQUEST]: (state) => {
    state.status = "loading";
  },
  ["SET_USER"]: (state, user) => {
    state.user = user;
  },
  ["ADD_ACTIVE"]: (state, ticket) => {
    state.activeTickets.push(ticket);
  },
  ["REMOVE_ACTIVE"]: (state, index) => {
    state.activeTickets.splice(index, 1);
  },
  ["SET_REFRESH_TOKEN"]: (state, refreshToken) => {
    state.refreshToken = refreshToken;
  },
  ["SET_TOKEN"]: (state, token) => {
    state.token = token;
  },
  ["SET_BRANCH_TOKEN"]: (state, token) => {
    state.branchToken = token;
  },
  ["SET_ACTIVE_COUNTER"]: (state, counter) => {
    state.activeCounter = counter;
  },
  ["SET_ACTIVE_USER"]: (state, counteruser) => {
    state.counterUser = counteruser;
  },
  ["SET_ACTIVE_TICKET"]: (state, ticket) => {
    state.activeTicket = ticket;
  },
  [types.MUTATE_LOADER_ON]: (state) => {
    state.loader = true;
  },
  [types.MUTATE_LOADER_OFF]: (state) => {
    state.loader = false;
  },

  [types.MUTATE_AUTH_SUCCESS]: (state, token) => {
    state.status = "success";
    state.token = token;
  },

  [types.MUTATE_AUTH_ERROR]: (state) => {
    state.status = "error";
  },
};

export const actions = {
  AUTH_REQUEST: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      commit(types.MUTATE_AUTH_REQUEST, JSON.stringify(payload));
      // sessionStorage.setItem("user-token", "token");
      commit(types.MUTATE_LOADER_ON);
      axios
        .post("/signup", payload)
        .then((res) => {
          // console.log(res);
          commit(types.MUTATE_LOADER_OFF);
          resolve(res);
        })
        .catch((err) => {
          // this.$toast.error("error signing up. Try again ");
          reject(err);
          commit(types.MUTATE_LOADER_OFF);
        });
    });
  },
  AUTH_LOGIN: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      commit(types.MUTATE_AUTH_REQUEST, JSON.stringify(payload));
      // sessionStorage.setItem("user-token", "token");
      commit(types.MUTATE_LOADER_ON);
      axios
        .post("/signin", payload)
        .then((res) => {
          console.log(res);
          let idToken = res.data.message.idtoken;
          let refreshToken = res.data.message.refreshtoken;
          axios.defaults.headers.common = {
            Authorization: "Bearer " + idToken,
          };
          // axios.defaults.headers.common["Authorization"] = idToken;
          commit("SET_TOKEN", idToken);
          commit("SET_REFRESH_TOKEN", refreshToken);
          let claims = getTokenPayload(idToken);
          commit("SET_USER", claims.user);
          // sessionStorage.setItem("user-token", idToken);
          // storeTokens(idToken,refreshToken)
          // sessionStorage.setItem("refresh-token", refreshToken);
          // sessionStorage.setItem("user", JSON.stringify(claims.user));
          commit(types.MUTATE_AUTH_SUCCESS, res);
          commit(types.MUTATE_LOADER_OFF);
          resolve(res);
        })
        .catch((err) => {
          // this.$toast.error("error signing in ");
          reject(err);
          commit(types.MUTATE_LOADER_OFF);
        });
    });
  },
  AUTH_BRANCH: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      // commit(types.MUTATE_AUTH_REQUEST, JSON.stringify(payload));
      // sessionStorage.setItem("user-token", "token");
      commit(types.MUTATE_LOADER_ON);
      axios
        .post("/branch-login", payload)
        .then((res) => {
          console.log(res);
          let idToken = res.data.message.idtoken;
          // let refreshToken = res.data.message.refreshtoken;
          // axios.defaults.headers.common = {
          //   Authorization: "Bearer " + idToken,
          // };
          // axios.defaults.headers.common["Authorization"] = idToken;
          commit("SET_BRANCH_TOKEN", idToken);
          // commit("SET_REFRESH_TOKEN", refreshToken);
          // let claims = getTokenPayload(idToken);
          // commit("SET_USER", claims.user);
          // sessionStorage.setItem("user-token", idToken);
          // storeTokens(idToken,refreshToken)
          // sessionStorage.setItem("refresh-token", refreshToken);
          // sessionStorage.setItem("user", JSON.stringify(claims.user));
          // commit(types.MUTATE_AUTH_SUCCESS, res);
          commit(types.MUTATE_LOADER_OFF);
          resolve(res);
        })
        .catch((err) => {
          // this.$toast.error("error signing in ");
          reject(err);
          commit(types.MUTATE_LOADER_OFF);
        });
    });
  },
  AUTH_REFRESH: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      let token = {};
      token.refreshToken = payload;
      commit(types.MUTATE_LOADER_ON);
      axios
        .post("/token/refresh", token)
        .then((res) => {
          console.log(res);
          storeTokens(res.data.tokens.idtoken, res.data.tokens.refreshtoken);
          axios.defaults.headers.common[
            "Authorization"
          ] = `Bearer ${res.data.tokens.idtoken}`;
          let idToken = res.data.tokens.idtoken;
          let refreshToken = res.data.tokens.refreshtoken;
          // axios.defaults.headers.common["Authorization"] = idToken;
          commit("SET_TOKEN", idToken);
          commit("SET_REFRESH_TOKEN", refreshToken);
          let claims = getTokenPayload(idToken);
          commit("SET_USER", claims.user);
          commit(types.MUTATE_LOADER_OFF);
          resolve(res);
        })
        .catch((err) => {
          reject(err);
          commit(types.MUTATE_LOADER_OFF);
        });
      // this.$toast.error("error signing in ");
    });
  },
  SET_CUTRRENT: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      commit("SET_ACTIVE_TICKET", payload);
      resolve()
    })
  },
  AUTH_LOGOUT: () => {
    return new Promise((resolve, reject) => {
      sessionStorage.removeItem("user-token");
      // remove the axios default header
      delete axios.defaults.headers.common["Authorization"];
      router.push({
        name: "SignIn",
      });
      resolve();
    }).catch((err) => {
      // this.$toast.error("error signing in ");
      reject(err);
      commit(types.MUTATE_LOADER_OFF);
    });
  },
};
