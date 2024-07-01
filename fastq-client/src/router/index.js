import Vue from "vue";
import VueRouter from "vue-router";
// import HomeView from "../views/HomeView.vue";
import store from "../store";
import Index from "../views/Index.vue";
import IndexT from "../views/pages/teller/IndexT.vue";
import Test from "../views/Test.vue";
Vue.use(VueRouter);

const routes = [
  {
    name: "Login",
    path: "/branch-login",
    component: () => import("@/views/pages/BranchLogin"),
    // meta: { brachLogin: true },
  },
  {
    path: "/dashboard",
    name: "home",
    component: Index,
    children: [
      {
        name: "Dashboard",
        path: "/dashboard",
        component: () => import("@/views/Dashboard"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Reports",
        path: "/reports",
        component: () => import("@/views/Reports"),
        meta: { RequiresAuth: true },
      },
      // {
      //   name: "Counter Reports",
      //   path: "/counterReports",
      //   component: () => import("@/views/counterReports"),
      //   meta: { RequiresAuth: true },
      // },
      {
        name: "Videos",
        path: "/video-ads",
        component: () => import("@/views/ads/videoAds"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Video Schedular",
        path: "/video-schedular",
        component: () => import("@/views/ads/videoSchedular"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Image Carousel",
        path: "/carousel",
        component: () => import("@/views/ads/adCarousel"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Assign Services",
        path: "/assign/services",
        component: () => import("@/views/manage/assignServices"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Announcement",
        path: "/announcement",
        component: () => import("@/views/ads/announcement"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Company Configuration",
        path: "/company",
        component: () => import("@/views/configuration/companyConf"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Display Configuration",
        path: "/display",
        component: () => import("@/views/configuration/displayConf"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Ticket Configuration",
        path: "/ticketconf",
        component: () => import("@/views/configuration/ticketConf"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Audio Configuration",
        path: "/audio",
        component: () => import("@/views/configuration/audioConf"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Manage Branches",
        path: "/manage/branches",
        component: () => import("@/views/manage/branchesManage"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Edit Branch",
        path: "/edit/branch/:id",
        component: () => import("@/views/manage/editBranch"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Manage Users",
        path: "/manage/users",
        component: () => import("@/views/manage/usersManage"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Edit User",
        path: "/edit/user/:id",
        component: () => import("@/views/manage/editUser"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Manage Counters",
        path: "/manage/counters",
        component: () => import("@/views/manage/counterManage"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Edit Counter",
        path: "/edit/counter/:id",
        component: () => import("@/views/manage/editCounter"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Manage Services",
        path: "/manage/services",
        component: () => import("@/views/manage/serviceManage"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Edit Service",
        path: "/edit/service/:id",
        component: () => import("@/views/manage/editService"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Manage Sections",
        path: "/manage/sections",
        component: () => import("@/views/manage/sectionManage"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Edit Section",
        path: "/edit/section",
        component: () => import("@/views/manage/editSection"),
        meta: { RequiresAuth: true },
      },
      {
        name: "Edit Assigned Service",
        path: "/edit/counterservice/:id",
        component: () => import("@/views/manage/editassignedServices"),
        meta: { RequiresAuth: true },
      },
      // Pages
    ],
  },
  {
    path: "/teller",
    name: "teller",
    component: IndexT,
    children: [
      {
        name: "Teller",
        path: "teller-screen",
        component: () => import("@/views/pages/teller/core/Teller"),
        meta: { RequiresAuth: true },
      },
      {
        name: "TellerLogin",
        path: "/",
        component: () => import("@/views/pages/teller/pages/TellerLogin.vue"),
        meta: { RequiresAuth: true },
      },

      // Pages
    ],
  },
  {
    path: "/about",
    name: "about",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: Test,
  },
  {
    path: "/ticket",
    name: "ticket",
    component: () =>
      import(
        /* webpackChunkName: "ticket" */ "../views/pages/ticket/TicketPrint.vue"
      ),
    meta: { RequiresAuth: true },
  },
  {
    path: "/fastqmenu",
    name: "menu",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(
        /* webpackChunkName: "menu" */ "../views/pages/ticket/FastqMenu.vue"
      ),
    meta: { RequiresAuth: true },
  },
  {
    path: "/",
    name: "SignIn",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.

    //HERE use a beforeEnter guard to route authenticated user back to menu
    // beforeEnter: async (to, from, next) => {
    //   const isLoggedIn = store.getters["isLoggedin"];

    //   if (isLoggedIn) {
    //       return next("/dashboard");
    //   }

    //   next();
    // },
    component: () =>
      import(/* webpackChunkName: "signin" */ "../views/pages/Login.vue"),
    meta: { RequiresAuth: false },
  },
  // {
  //   path: "/",
  //   name: "test",
  //   component: Test,
  // },
  {
    path: "/signup",
    name: "SignUp",
    component: () =>
      import(/* webpackChunkName: "signup" */ "../views/pages/SignUpV2.vue"),
    meta: { RequiresAuth: false },
  },
  {
    path: "/forgot-password",
    name: "Forgot Password",
    component: () =>
      import(
        /* webpackChunkName: "signup" */ "../views/pages/ForgotPassword.vue"
      ),
    meta: { RequiresAuth: false },
  },
  {
    path: "/ds",
    name: "signage",
    component: () =>
      import(
        /* webpackChunkName: "signup" */ "../views/digital-signage/SignageTemplate2.vue"
      ),
    meta: { RequiresAuth: false },
  },
  {
    path: "/dscontainer",
    name: "Signagecontainer",
    component: () =>
      import(
        /* webpackChunkName: "signup" */ "../views/digital-signage/SignageContainer.vue"
      ),
    meta: { RequiresAuth: false },
  },
  {
    path: "/ticketprintcontainer",
    name: "Ticketprintcontainer",
    component: () =>
      import(
        /* webpackChunkName: "ticketcontainer" */ "../views/pages/ticket/PrintContainer.vue"
      ),
    meta: { RequiresAuth: false },
  },
  {
    path: "/:catchAll(.*)*",
    name: "NotFound",
    component: () =>
      import(/* webpackChunkName: "not_found" */ "../views/AboutView.vue"),
  },
];

const router = new VueRouter({
  // mode: "history",
  routes,
});

export default router;
