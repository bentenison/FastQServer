<template>
  <!-- temporary and bottom for mobile devices -->
  <v-navigation-drawer :absolute="$vuetify.breakpoint.mdAndUp" style="overflow: hidden; position:fixed" left app class="" v-model="drawerOpen">
    <template v-slot:prepend>
      <v-list-item two-line>
        <v-list-item-avatar>
          <img src="../../assets/rsz_al_2.png">
        </v-list-item-avatar>

        <v-list-item-content>
          <v-list-item-title>{{ $store.state.Auth.user.company }}</v-list-item-title>
          <v-list-item-subtitle>Logged In</v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>
    </template>
    <v-divider></v-divider>
    <v-list>
      <v-list nav>
        <v-list-item-group v-model="selectedItem" color="primary">
          <v-list-item v-for="item in items" :to="item.to" :key="item.to">
            <v-list-item-icon>
              <v-icon v-text="item.action"></v-icon>
            </v-list-item-icon>

            <v-list-item-content>
              <v-list-item-title class="text-h6 grey--text"> {{ item.title }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list-item-group>
      </v-list>
      <!-- <v-list-item v-model="selectedItem"
        color="primary" link>
        <v-list-item-icon>
          <v-icon>mdi-home</v-icon>
        </v-list-item-icon>

        <v-list-item-title>Home</v-list-item-title>
      </v-list-item> -->

      <v-list-group :value="false" v-for="route in drawer" :key="route.icon" :prepend-icon="route.icon">
        <template v-slot:activator>
          <v-list-item-title class="text-h6 grey--text">{{ route.section }}</v-list-item-title>
        </template>



        <v-list-item v-for="section in route.routes" :key="section.to" link :to="section.to">
          <v-list-item-title v-text="section.title"></v-list-item-title>

          <v-list-item-icon>
            <v-icon v-text="section.action"></v-icon>
          </v-list-item-icon>
        </v-list-item>
        <!-- </v-list-group> -->


      </v-list-group>
    </v-list>
  </v-navigation-drawer>
</template>
  
<script>
// Utilities
import EventBus from '@/plugins/eventBus';
import {
  mapState,
} from 'vuex'
export default {
  name: 'DashboardCoreDrawer',
  data() {
    return {
      drawerOpen: true,
      selectedItem: null,
      items: [
        {
          action: 'mdi-view-dashboard',
          title: 'Dashboard',
          to: '/dashboard'
        },
        {
          action: 'mdi-chart-arc',
          title: 'Reports',
          to: '/reports'
        },
        // {
        //   action: 'mdi-content-save',
        //   title: 'Ads',
        //   to: '/video-ads'
        // },
        // {
        //     action: 'fa-chart-line',
        //     title: 'Dashboard',
        //     to: '/dashboard'
        // },
        // {
        //     action: 'fa-user',
        //     title: 'Attractions',

        // },
      ],
      drawer: [
        {
          icon: "mdi-google-ads",
          section: "Ads",
          routes: [{
            action: 'mdi-video-box',
            title: 'Videos',
            to: '/video-ads'
          },
          {
            action: 'mdi-image-box',
            title: 'Images',
            to: '/carousel'
          },
          {
            action: 'mdi-book-clock',
            title: 'Video Schedular',
            to: '/video-schedular'
          },
          {
            action: 'mdi-bullhorn',
            title: 'Announcement',
            to: '/announcement'
          }],
        },
        {
          icon: "mdi-cogs",
          section: "Configuration",
          routes: [{
            action: 'mdi-domain',
            title: 'Company',
            to: '/company'
          },
          {
            action: 'mdi-volume-high',
            title: 'Audio',
            to: '/audio'
          },
          {
            action: 'mdi-ticket',
            title: 'Ticket',
            to: '/ticketConf'
          },
          {
            action: 'mdi-television',
            title: 'Display',
            to: '/display'
          }
          ],
        },
        {
          icon: "mdi-alpha-m-circle",
          section: "Manage",
          routes: [
            //   {
            //   action: 'mdi-source-branch',
            //   title: 'Branches',
            //   to: '/manage/branches'
            // },
            {
              action: 'mdi-account-group',
              title: 'Users',
              to: '/manage/users'
            },
            {
              action: 'mdi-account-wrench-outline',
              title: 'Services',
              to: '/manage/services'
            },
            // {
            //   action: 'mdi-account-wrench-outline',
            //   title: 'Assign Services',
            //   to: '/assign/services'
            // },
            {
              action: 'mdi-content-save',
              title: 'Counters',
              to: '/manage/counters'
            }
          ],
        },
      ],
    }
  },
  methods: {
    handleFabClick() {
      this.drawerOpen = !this.drawerOpen
    }
  },
  created() {
    console.log("EventBus",this.$eventBus);
    this.$eventBus.$on('fabClicked', this.handleFabClick)
  }
}
</script>
  
