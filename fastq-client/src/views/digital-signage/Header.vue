<template>
  <v-app-bar elevation="4" app height="70px">
    <div class="d-flex align-items-center justify-content-center p-3">
      <v-img
        :src="
          ipAddress !== 'localhost'
            ? `http://${ipAddress}:8090/uploaded/logo.png`
            : `http://localhost:8080/server/uploaded/logo.png`
        "
        max-height="60"
        max-width="100"
        contain
      ></v-img>
      <h2 class="m-0">
        {{ $store.state.Auth.user.company.toUpperCase() }}
      </h2>
      <v-chip class="ma-2" color="primary" label text-color="white">
        <v-icon left> mdi-label </v-icon>
        <strong>HEAD OFFICE</strong>
      </v-chip>
    </div>
    <v-spacer></v-spacer>
    <div
      class="d-md-flex d-sm-none d-xs-none align-items-center justify-content-center flex-wrap p-3"
    >
      <v-icon> mdi-calendar </v-icon>
      <v-chip
        class="ma-2 d-flex align-items-center justify-content-center"
        label
        text-color="grey darken-2"
      >
        <h4 class="font-weight-regular">{{ date }}</h4>
      </v-chip>
      <v-icon> mdi-clock </v-icon>
      <v-chip class="ma-2" label text-color="grey darken-2">
        <h4 class="font-weight-regular">{{ time }}</h4>
      </v-chip>
    </div>
    <div class="d-flex align-items-center justify-content-center p-3">
      <v-btn fab small color="primary" class="mx-2">
        <v-icon color="white">mdi-cloud</v-icon>
      </v-btn>
    </div>
  </v-app-bar>
</template>

<script>
import moment from "moment";

export default {
  data() {
    return {
      m: moment(),
      ipAddress: null,
      week: ["SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"],
      timerID: null,
      time: "",
      date: "",
    };
  },
  methods: {
    updateTime() {
      var cd = new Date();
      this.time =
        this.zeroPadding(cd.getHours(), 2) +
        ":" +
        this.zeroPadding(cd.getMinutes(), 2) +
        ":" +
        this.zeroPadding(cd.getSeconds(), 2);
      this.date =
        this.zeroPadding(cd.getFullYear(), 4) +
        "-" +
        this.zeroPadding(cd.getMonth() + 1, 2) +
        "-" +
        this.zeroPadding(cd.getDate(), 2) +
        " " +
        this.week[cd.getDay()];
    },
    zeroPadding(num, digit) {
      var zero = "";
      for (var i = 0; i < digit; i++) {
        zero += "0";
      }
      return (zero + num).slice(-digit);
    },
  },
  created() {
    const url = new URL(window.location.origin);

    // Get the IP address from the URL
    this.ipAddress = url.hostname;
    console.log("url", this.ipAddress);
    this.timerID = setInterval(this.updateTime, 1000);
    this.updateTime();
  },
};
</script>

<style lang="scss" scoped></style>

