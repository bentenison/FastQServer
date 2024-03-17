<template>
  <v-app-bar id="app-bar" fixed app color="#F5F5F5" flat height="60">
    <v-btn
      class="mr-3"
      elevation="1"
      fab
      small
      @click="$eventBus.$emit('fabClicked')"
    >
      <v-icon v-if="false"> mdi-view-quilt </v-icon>

      <v-icon v-else> mdi-dots-vertical </v-icon>
    </v-btn>

    <v-toolbar-title
      class="hidden-sm-and-down font-weight-light"
      v-text="$route.name"
    />

    <v-spacer />

    <v-text-field color="secondary" hide-details style="max-width: 165px">
      <template v-if="$vuetify.breakpoint.mdAndUp" v-slot:append-outer>
        <v-btn class="mt-n2" elevation="1" fab small>
          <v-icon>mdi-magnify</v-icon>
        </v-btn>
      </template>
    </v-text-field>

    <div class="mx-3" />

    <!-- <v-btn class="ml-2" min-width="0" text to="/">
      <v-icon>mdi-view-dashboard</v-icon>
    </v-btn> -->

    <!-- <v-menu
        bottom
        left
        offset-y
        origin="top right"
        transition="scale-transition"
      >
        <template v-slot:activator="{ attrs, on }">
          <v-btn
            class="ml-2"
            min-width="0"
            text
            v-bind="attrs"
            v-on="on"
          >
            <v-badge
              color="red"
              overlap
              bordered
            >
              <template v-slot:badge>
                <span>5</span>
              </template>
  
              <v-icon>mdi-bell</v-icon>
            </v-badge>
          </v-btn>
        </template>
  
        <v-list
          :tile="false"
          nav
        >
          <div>
            <app-bar-item
              v-for="(n, i) in notifications"
              :key="`item-${i}`"
            >
              <v-list-item-title v-text="n" />
            </app-bar-item>
          </div>
        </v-list>
      </v-menu> -->

    <div class="text-center">
      <v-menu
        v-model="menu"
        :close-on-content-click="true"
        :nudge-width="200"
        offset-x
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn class="ml-2" min-width="0" text v-bind="attrs" v-on="on">
            <v-icon>mdi-account</v-icon>
          </v-btn>
        </template>
        <v-card>
          <v-list>
            <v-list-item>
              <v-list-item-avatar>
                <v-avatar color="red" class="text-center">
                  <span class="white--text text-h5 text-center">{{
                    initials
                  }}</span>
                </v-avatar>
              </v-list-item-avatar>

              <v-list-item-content>
                <v-list-item-title
                  >{{ $store.state.Auth.user.firstname }}
                  {{ $store.state.Auth.user.lastname }}</v-list-item-title
                >
                <v-list-item-subtitle>{{
                  $store.state.Auth.user.company.toUpperCase()
                }}</v-list-item-subtitle>
              </v-list-item-content>

              <v-list-item-action>
                <v-btn :class="fav ? 'red--text' : ''" icon @click="fav = !fav">
                  <v-icon>mdi-heart</v-icon>
                </v-btn>
              </v-list-item-action>
            </v-list-item>
          </v-list>

          <v-divider></v-divider>

          <v-list>
            <v-list-item>
              <v-list-item-title
                ><v-btn
                  block
                  text
                  color="error"
                  class="mt-1 text-md-button"
                  @click="$router.push('/')"
                >
                  <v-icon left>mdi-logout</v-icon>
                  logout
                </v-btn></v-list-item-title
              >
            </v-list-item>

            <v-list-item>
              <v-list-item-title
                ><v-btn
                  block
                  text
                  color="primary"
                  class="mt-1 text-md-button"
                  @click="$router.push('/company')"
                >
                  <v-icon left>mdi-card-account-details</v-icon>
                  profile
                </v-btn></v-list-item-title
              >
            </v-list-item>
          </v-list>

          <v-card-actions>
            <v-spacer></v-spacer>

            <v-btn text @click="menu = false"> Cancel </v-btn>
            <v-btn color="primary" text @click="menu = false"> Save </v-btn>
          </v-card-actions>
        </v-card>
      </v-menu>
    </div>
  </v-app-bar>
</template>

<script>
// import EventBus from '@/plugins/eventBus';
export default {
  data() {
    return {
      menu: false,
      fav: false,
    };
  },
  methods: {},
  computed: {
    initials() {
      const firstInitial = this.$store.state.Auth.user.firstname
        ? this.$store.state.Auth.user.firstname[0]
        : "";
      const lastInitial = this.$store.state.Auth.user.lastname
        ? this.$store.state.Auth.user.lastname[0]
        : "";
      return `${firstInitial}${lastInitial}`.toUpperCase();
    },
  },
};
</script>

<style lang="sass" scoped>
</style>