<template>
  <div class="container h-100">
    <div class="row h-100">
      <div class="row col-md-12 h-100">
        <v-card class="col-md-12 h-auto" flat>
          <EditTable
            :data="tableData"
            :headers="headers"
            :name="'counterservice'"
          />

          <h5 class="h4 grey--text mt-5">Assign Services to Counter</h5>
          <div class="row mt-2 d-flex align-items-center gap-3 ml-2">
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
            <v-col cols="4" sm="4" class="m-0 p-0">
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
            </v-col>
            <v-col cols="12" sm="12" class="">
              <!-- <p class="h6 grey--text">Services</p> -->
              <v-btn
                class="mr-3"
                color="primary"
                @click.prevent="assignServices"
                ><v-icon left>mdi-content-save</v-icon>Save</v-btn
              >
              <v-btn color="warning"
                ><v-icon left>mdi-restore</v-icon>Reset</v-btn
              >
            </v-col>
            <!-- </div> -->
          </div>
          <!-- <v-btn color="primary" class=""><v-icon left>mdi-content-save-all</v-icon>Save config</v-btn> -->
        </v-card>
      </div>
    </div>
  </div>
</template>

<script>
import DataTables from "../tables/dataTables.vue";
import EditTable from "../tables/editTable.vue";
import axios from "axios";
export default {
  data() {
    return {
      tableData: [],
      services: [],
      counters: [],
      selectedServices: null,
      selectedCounter: null,
      headers: [
        {
          text: "CounterName",
          align: "start",
          sortable: false,
          value: "counterName",
        },
        { text: "CounterNumber", value: "counterNumber" },
        { text: "Services", value: "services" },
        { text: "Actions", value: "actions", sortable: false },
      ],
    };
  },
  components: { DataTables, EditTable },
  methods: {
    addBranch() {
      // this.branch.services = this.services
      console.log("services:::::", this.branch);
      this.branch.companyCode = this.$store.state.Auth.user.company_code;
      axios
        .post("/branch/addbranch", this.branch)
        .then((res) => {
          console.log("res:::::", res);
          this.$toast.success("branch added successfully.");
          this.branch = {};
        })
        .catch((err) => {
          console.log(err.response);
          this.$toast.error("error occured while adding branch!!!");
        });
    },

    getServices() {
      // console.log("store>>>>>>>>>>>>>>>>>>>",);
      return new Promise((resolve, reject) => {
        axios
          .get(
            `/service/getAllservice/${this.$store.getters.getUser.company_code}`
          )
          .then((res) => {
            // console.log(">>>>>>>>>>>>>>>>>", res);

            res.data.message.forEach((element) => {
              let data = {};
              data.Code = element.Code;
              data.Name = element.Name;
              data.Prefix = element.Prefix;
              data.NumberStarts = element.NumberStarts;
              data.NumberEnds = element.NumberEnds;
              data.Hide = element.Hide;
              data.id = element.ID;
              data.Service_Arabic = element.Workflow;
              this.services.push(data);
            });
            resolve();
            this.$toast.success("services fetched successfully.");
            // this.connection.send("Hello World !!!")
          })
          .catch((err) => {
            console.log(err.response);
            // this.$toast.error("error occured while getting service!!!")
            reject(err);
          });
      });
    },
    getAllCounters() {
      this.counters = [];
      return new Promise((resolve, reject) => {
        axios
          .get(
            `/counter/getAllcounter/${this.$store.getters.getUser.company_code}`
          )
          .then((res) => {
            // console.log(">>>>>>>>>>>>>>>>>", res);
            res.data.message.forEach((element) => {
              let data = {};
              data.counterNumber = element.CounterNumber;
              data.counterName = element.CounterName;
              data.branchCode = element.BranchCode;
              data.branchName = element.BranchName;
              data.date = element.CreatedAt;
              data.id = element.ID;
              this.counters.push(data);
            });
            resolve();
            this.$toast.success("counters fetched successfully.");
            // this.connection.send("Hello World !!!")
          })
          .catch((err) => {
            console.log(err.response);
            this.$toast.error("error occured while getting counters!!!");
            reject(err);
          });
      });
    },
    getAllAssignedServices() {
      axios
        .get(`/service/getAllassignedServices`)
        .then((res) => {
          //   console.log(">>>>>>>>>>>>>>>>>", res);
          res.data.forEach((element) => {
            let data = {};
            // console.log("element", element);
            this.counters.forEach((counter) => {
              //   console.log("counter", counter);
              if (counter.id === element.CounterID) {
                data.counterNumber = counter.counterNumber;
                data.counterName = counter.counterName;
                data.services = element.ServiceID;
                data.id = element.CounterServiceID
              }
            });
            // console.log("DAaya >>>>>>>>>>>>", data);
            this.tableData.push(data);
          });

          //   this.$toast.success("counters fetched successfully.");
          // this.connection.send("Hello World !!!")
        })
        .catch((err) => {
          console.log(err.response);
          //   this.$toast.error("error occured while getting counters!!!");
        });
    },
    getassignedServiceforCounter() {
      axios
        .get(
          `/service/assignService/${this.$store.getters.getUser.company_code}`
        )
        .then((res) => {
          //   console.log(">>>>>>>>>>>>>>>>>", res);
          this.$toast.success("counter services fetched successfully.");
          // this.connection.send("Hello World !!!")
        })
        .catch((err) => {
          console.log(err.response);
          this.$toast.error("error occured while getting counter services!!!");
        });
    },
    assignServices() {
      let svcs = [];
      let counterSvc = {};
      this.selectedServices.forEach((svc) => {
        svcs.push(svc.Name);
        svcs.push(svc.Service_Arabic);
      });
      counterSvc.CounterID = this.selectedCounter.id;
      counterSvc.ServiceID = svcs.toString();
      console.log("object,", counterSvc);
      // return;
      //   console.log("assigned services", counterSvc);
      axios
        .post("/service/addassignedServices", counterSvc)
        .then((res) => {
          console.log("res:::::", res);
          this.$toast.success("service assigned successfully.");
          //   this.branch = {};
          this.selectedCounter = null;
          this.selectedServices = null;
          this.getAllAssignedServices()
        })
        .catch((err) => {
          console.log(err.response);
          this.$toast.error("error occured while assigning service!!!");
        });

      //   console.log("assigned counter", this.selectedCounter);
    },
    updateassignedServices() {
      axios
        .post("/service/updateassignedServices", this.branch)
        .then((res) => {
          //   console.log("res:::::", res);
          this.$toast.success("services updated successfully.");
          this.branch = {};
        })
        .catch((err) => {
          console.log(err.response);
          this.$toast.error("error occured while updating counter services!!!");
        });
    },
    reset() {
      this.branch = {};
      this.service = {};
    },
  },
  mounted() {
    // this.getBranches();
    this.getServices().then((res) => {
      this.getAllCounters().then((res) => {
        this.getAllAssignedServices();
      });
    });
  },
};
</script>

<style lang="scss" scoped></style>