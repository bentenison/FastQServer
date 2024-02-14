import Vue from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
import vuetify from "./plugins/vuetify";
import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import VueApexCharts from "vue-apexcharts";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
import store from "./store";
import axios from "axios";
import vueLocalStorage from "vue-localstorage";
import loader from "vue-ui-preloader";
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";
import { getTokenPayload, storeTokens } from "./plugins/utils";


const options = {
  // You can set your default options here
  position: "top-right",
  timeout: 2000,
  closeOnClick: true,
  pauseOnHover: true,
  draggable: true,
  draggablePercent: 0.6,
  showCloseButtonOnHover: false,
  // hideProgressBar: true,
  closeButton: "button",
  // icon: true,
  rtl: false,
};
var EventBus = new Vue();
Object.defineProperties(Vue.prototype, {
  $eventBus: {
    get: function () {
      return EventBus;
    }
  }
});
Vue.use(Toast, options);
Vue.use(loader);
Vue.use(VueApexCharts);

Vue.component("apexchart", VueApexCharts);
// Make BootstrapVue available throughout your project
Vue.use(BootstrapVue);
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin);
Vue.config.productionTip = false;
// sessionStorage.setItem('user-token', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.qCJ-hcgSTLgkaT7kfI6--xRm4IEpPFQmMj3UZ94gNo4')
Vue.use(vueLocalStorage);
const isProduction = process.env.NODE_ENV === 'production'
var baseURL = isProduction ? "/" : "/server";

axios.defaults.headers.common[
  "Authorization"
] = `Bearer ${store.state.Auth.token}`;
// console.log(store.state.Auth.token);
axios.defaults.baseURL = baseURL;

// axios.interceptors.request.use(
//   (config) => {
//     console.log("::::::::::::::::", config);
// if (config.authorization !== false) {
//   const token = getCurrentAccessToken();
//   if (token) {
//     config.headers.Authorization = "Bearer " + token;
//   }
// }
//     return config;
//   },
//   (error) => {
//     console.log("::::::::::::::::", error);
//     return error;
//   }
// );
// router.beforeEach((to, from, next) => {
//   let token = store.getters.getIdToken;
//   // const currentUser = store.state.currentUser;

//   if( !token) {
//       next('/');
//   } else if(to.path == '/' && currentUser) {
//       next('/fastqmenu');
//   } else {
//       next();
//   }
// });
// router.beforeEach((to, from, next) => {
//   if (to.matched.some((record) => record.meta.guest)) {
//     if (store.getters.isAuthenticated) {
//       next("/posts");
//       return;
//     }
//     next();
//   } else {
//     next();
//   }
// });
// axios.interceptors.request.use(function (config) {
//   config.headers.common["Access-Control-Allow-Origin"] = "*";
//   if (config.url.indexOf("r/") !== -1) {
//     let token = localStorage.getItem("user-token");
//     if (token) {
//       config.headers.common["authorization"] = "Bearer " + token;
//       // config.headers.post['Content-Type'] = 'application/json'
//       // next()
//     } else {
//       // store.dispatch('userLogout')
//       router.replace("/login");
//     }
//   }
//   return config;
// });
// axios.interceptors.response.use((response)=>{
//   console.log('response',response)
// })
axios.interceptors.response.use(
  function (response) {
    // Any status code that lie within the range of 2xx cause this function to trigger
    // Do something with response data
    // console.log("response in main js interceptor",response.config)
    return response;
  },
  function (error) {
    if (error.response.status == 401) {
      // store.dispatch('userLogout')
      localStorage.clear();
      router.push("/");
    }
    return Promise.reject(error);
  }
);
new Vue({
  router,
  vuetify,
  store,
  render: (h) => h(App),
}).$mount("#app");
