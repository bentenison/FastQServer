<template>
  <v-app-bar
    id="app-bar"
    fixed
    app
    color="#F5F5F5"
    flat
    :height="!hide ? 75 : 50"
  >
    <v-btn class="mr-3" elevation="1" fab small>
      <!-- <v-icon v-if="value">
        mdi-view-quilt
      </v-icon> -->

      <v-icon> mdi-dots-vertical </v-icon>
    </v-btn>

    <v-toolbar-title
      class="hidden-sm-and-down font-weight-light"
      v-text="$route.name"
    />

    <v-img
      class="mx-2"
      src="../../../../assets/sec-icon.svg"
      max-height="45"
      max-width="100"
      contain
      v-if="!hide"
    ></v-img>
    <!-- <v-toolbar-title class="hidden-sm-and-down" v-text="'Branch Name'" /> -->
    <v-spacer />

    <div class="mx-3" />

    <!-- <v-btn class="ml-2" min-width="0" text color="primary">
      <v-icon color="primary">mdi-cog</v-icon>setting
    </v-btn> -->
    <v-chip class="ma-2" color="primary" label text-color="white" v-if="!hide">
      <v-icon left> mdi-label </v-icon>
      <strong>Head Office - {{ $store.state.Auth.user.company }}</strong>
    </v-chip>

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

    <v-btn
      class="ml-2"
      min-width="0"
      text
      color="primary"
      v-if="!hide && $route.name === 'Teller'"
      @click="$eventBus.$emit('float-mode')"
    >
      <v-icon color="primary">mdi-window-minimize</v-icon> Float Mode
    </v-btn>
    <v-btn
      class="ml-2"
      min-width="0"
      text
      color="primary"
      v-if="hide && $route.name === 'Teller'"
      @click="$eventBus.$emit('float-mode-off')"
    >
      <v-icon color="primary">mdi-window-maximize</v-icon> exit float mode
    </v-btn>
    <v-btn class="ml-2" min-width="0" text color="error" @click="killapp">
      <v-icon color="error">mdi-close</v-icon> Close
    </v-btn>
    <v-btn
      class="ml-2"
      min-width="0"
      text
      color="error"
      v-if="!hide && $store.state.Auth.counterUser"
      @click="logout"
    >
      <v-icon color="error">mdi-logout</v-icon> logout
    </v-btn>
  </v-app-bar>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      value: true,
      hide: false,
    };
  },
  methods: {
    killapp() {
      this.logout();
      window.close();
      // sessionStorage.clear();
      // this.$router.push("/teller/");
    },
    logout() {
      console.log("object user logged out");
      axios
        .get(
          `/license/updatecounteruser/${this.$store.state.Auth.activeCounter.ID}`
        )
        .then((res) => {
          this.$toast.success("logged out!!!");
          localStorage.clear();
          sessionStorage.clear();
          this.$router.push("/teller/");
        })
        .catch((err) => {
          this.$toast.error("error logging out");
          console.error(err);
        });
    },
  },
  beforeDestroy() {
    // this.$eventBus.$off('float-mode')
    // this.$eventBus.$off('float-mode-off')
  },
  created() {
    this.$eventBus.$on("float-mode", () => {
      this.hide = true;
    });
    this.$eventBus.$on("float-mode-off", () => {
      this.hide = false;
    });
  },
};
</script>

<style lang="sass" scoped>
</style>