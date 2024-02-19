<template>
  <v-content app class="p-0">
    <router-view width="100%" :style="$route.name === 'Teller' ? '' : 'margin-top:60px'" />


    <DashboardCoreFooter />
  </v-content>
</template>
  
<script>
// import Footer from "../core/Footer.vue"
export default {
  // components:{
  //   Footer
  // },
  name: 'DashboardCoreView',
  components: {
    DashboardCoreFooter: () => import('./Footer'),
  },
}
</script>