<template>
  <div class="row mt-5 d-flex align-items-center justify-content-center gap-3">
    <!-- <div class="col-md-12 row d-flex align-items-center mt-5"> -->
    <v-col cols="4" sm="4" class="m-0 p-0">
      <p class="h6 grey--text">Select Services to Assign</p>
      <v-select
        v-model="selectedServices"
        :items="services"
        label="Select"
        multiple
        chips
        item-text="Name"
        item-value="ID"
        return-object
        single-line
      ></v-select>
    </v-col>
    <!-- <v-col cols="4" sm="4" class="m-0 p-0">
      <p class="h6 grey--text">Select Counter To Assign Services</p>
      <v-select
        v-model="selectedCounter"
        :items="counters"
        item-text="counterName"
        item-value="ID"
        return-object
        single-line
        label="Select Counter"
      ></v-select>
    </v-col> -->
    <v-col cols="12" sm="12" class="d-flex align-items-center justify-content-center">
      <!-- <p class="h6 grey--text">Services</p> -->
      <v-btn class="mr-3" color="primary" @click.prevent="editAssignedService"
        ><v-icon left>mdi-content-save</v-icon>Save</v-btn
      >
    </v-col>
    <!-- </div> -->
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      id: null,
      services: [],
      CounterServices: null,
      selectedServices: [],
    };
  },
  methods: {
    getassignedServiceforCounter() {
      axios
        .get(`/service/assignService/${this.id}`)
        .then((res) => {
          // console.log(">>>>>>>>>>>>>>>>>", res);
          this.CounterServices = res.data;
          this.services = this.CounterServices.ServiceID.split(",");
          this.$toast.success("counter services fetched successfully.");
          // this.connection.send("Hello World !!!")
        })
        .catch((err) => {
          console.log(err.response);
          this.$toast.error("error occured while getting counter services!!!");
        });
    },
  },
  mounted() {
    this.getassignedServiceforCounter();
  },
  created() {
    this.id = this.$route.params.id;
  },
};
</script>

<style scoped lang="scss">
</style>