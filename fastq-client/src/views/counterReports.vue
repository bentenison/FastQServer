<template>
  <div class="container">
    <h3 class="grey--text text-center">Active Counter Reports</h3>
    <v-row class="d-flex align-center justify-center">
      <v-col cols="8" md="8" lg="8">
        <tiime-pickers />
      </v-col>
      <v-col cols="3" md="3" lg="3">
        <v-select
          v-model="selectedUser"
          :items="users"
          label="Select User"
          item-text="Firstname"
          item-value="ID"
          return-object
          hide-details
        ></v-select>
      </v-col>
      <!-- <v-col cols="3" md="3" lg="3">
        <v-select
          label="Select"
          item-text="Name"
          item-value="ID"
          return-object
          single-line
          hide-details
        ></v-select>
      </v-col> -->
      <v-col cols="3" md="3" lg="3">
        <v-btn large text color="primary" class="mt-2 text-md-button">
          <v-icon left>mdi-export</v-icon>
          Submit
        </v-btn>
      </v-col>
    </v-row>
    <v-row class="d-flex align-center justify-center">
      <v-col cols="12" md="12" lg="12" class="d-flex justify-content-between">
        <h2 class="grey--text">Tickets Served By User</h2>
        <v-btn
          large
          text
          color="success"
          class="mt-1 text-md-button"
          @click="exportToExcel(userData, 'tickets_served')"
        >
          <v-icon left>mdi-export</v-icon>
          Export Excel
        </v-btn>
      </v-col>
      <!-- <v-col cols="12" md="12" lg="12">
                <DataTables :headers="tableheaders" :items="userData" />
            </v-col> -->
      <v-col cols="12" md="12" lg="12" class="d-flex justify-content-between">
        <h2 class="grey--text">Active Time Of User Per Day</h2>
        <v-btn
          large
          text
          color="success"
          class="mt-1 text-md-button"
          @click="exportToExcel(activeData, 'active_time_user_per_day')"
        >
          <v-icon left>mdi-export</v-icon>
          Export Excel
        </v-btn>
      </v-col>
      <!-- <v-col cols="12" md="12" lg="12">
                <DataTables :headers="activeheaders" :items="activeData" />
            </v-col> -->
    </v-row>
  </div>
</template>

<script>
import axios from "axios";
import TiimePickers from "../components/TiimePickers.vue";
export default {
  components: { TiimePickers },
  data() {
    return {
      tableData: [],
      users: null,
      selectedUser: null,
    };
  },
  methods: {
    getAllUser() {
      let payload = {
        company_code: this.$store.state.Auth.user.company_code,
        branch_code: "",
      };
      //   this.tableData = [];
      axios
        .post(`/user/getAlluser`, payload)
        .then((res) => {
          console.log("object", res);
          this.users = res.data.message;
          // this.reset()
        })
        .catch((err) => {
          console.log(err.response);
          this.$toast.error("error occured while getting all user!!!");
        });
    },
    timeSelected(d) {
      console.log("selecfted times", d);
    },
  },
  created() {
    this.$eventBus.$on("time-selected", this.timeSelected);
    this.getAllUser();
  },
  mounted() {},
};
</script>

<style lang="scss" scoped>
</style>