<template>
  <div class="container h-100">
    <div class="row h-100">
      <div class="row col-md-12 h-100">
        <v-card class="col-md-7 h-auto" flat>
          <div class="d-flex flex-column">
            <p class="h6 grey--text">Company Name</p>
            <p class="h4 grey--text">{{ $store.state.Auth.user.company }}</p>
          </div>
          <div class="d-flex flex-column">
            <p class="h6 grey--text">Contact</p>
            <p class="h4 grey--text">{{ $store.state.Auth.user.email }}</p>
          </div>
          <div class="d-flex flex-column">
            <p class="h6 grey--text">Company Code</p>
            <p class="h4 grey--text">
              {{ $store.state.Auth.user.company_code }}
            </p>
          </div>
          <div class="d-flex flex-column">
            <p class="h6 grey--text">Country:</p>
            <!-- <v-select :items="items" v-model="$store.state.Auth.user.country" label="country"></v-select> -->
            <p class="h4 grey--text">{{ $store.state.Auth.user.country }}</p>
          </div>
          <div class="d-flex flex-column">
            <p class="h6 grey--text">Time Zone:</p>
            <!-- <v-select :items="items" v-model="$store.state.Auth.user.timezone" label="timezone"></v-select> -->
            <p class="h4 grey--text">{{ $store.state.Auth.user.timezone }}</p>
          </div>
          <div class="d-flex flex-column">
            <p class="h6 grey--text">Ticket Print URL:</p>
            <!-- <v-select :items="items" v-model="$store.state.Auth.user.timezone" label="timezone"></v-select> -->
            <p class="h4 grey--text">{{ url }}</p>
          </div>

          <v-btn color="primary" class=""
            ><v-icon left>mdi-content-save-all</v-icon>Save config</v-btn
          >
        </v-card>
        <div class="col-md-5 p-3">
          <h5 class="h6 grey--text">Change IP Address</h5>
          <v-text-field
            name="IP"
            label="IP Address"
            id="id"
            class="w-80"
            v-model="ipAdd"
          ></v-text-field>
          <v-btn
            color="primary"
            elevation="2"
            outlined
            :loading="false"
            @click="UpdateIPAdd"
            ><v-icon left> mdi-upload </v-icon>Save Changes</v-btn
          >
        </div>
        <div class="row col-md-12">
          <v-expansion-panels
            v-model="panel"
            :readonly="readonly"
            multiple
            flat
            style="z-index: 10"
            v-if="false"
          >
            <v-expansion-panel>
              <v-expansion-panel-header class="grey--text">
                <p class="text-center h5">AUDIT TRAIL</p>
              </v-expansion-panel-header>
              <v-expansion-panel-content>
                <DataTables />
              </v-expansion-panel-content>
            </v-expansion-panel>
          </v-expansion-panels>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import DataTables from "../tables/dataTables.vue";

export default {
  data() {
    return {
      panel: "",
      readonly: false,
      items: ["India", "America", "Europe", "Japan"],
      ServerDetails: null,
      ipAdd: null,
    };
  },
  components: { DataTables },
  computed: {
    url() {
      if (this.ServerDetails) {
        return (
          "https://" +
          this.ServerDetails.server_ip +
          ":443/fastqclient/#/ticket?company=" +
          this.$store.state.Auth.user.company_code
        );
      }
    },
  },
  methods: {
    getServerByCompany() {
      return new Promise((resolve, reject) => {
        axios
          .get(`/config/getServer/${this.$store.state.Auth.user.company_code}`)
          .then((res) => {
            // this.allvideos = res.data
            // console.log("res >>>>>>>>>>:::::", res);
            // this.tableData = []
            this.ServerDetails = res.data;
            resolve(res.data);
            // this.$store.commit("SET_USER", res.data)
            // this.$store.commit(types.MUTATE_LOADER_OFF)
          })
          .catch((err) => {
            console.log(err.response);
            // reject(err.response)
            reject(err);
            // this.$store.commit(types.MUTATE_LOADER_OFF)
            this.$toast.error("error occured while getting server details!!!");
          });
      });
    },
    UpdateIPAdd() {
      return new Promise((resolve, reject) => {
        axios
          .post(`/config/updateIp`, {
            server_ip: this.ipAdd,
            id: this.$store.state.Auth.user.company_code,
          })
          .then((res) => {
            // this.allvideos = res.data
            // console.log("res >>>>>>>>>>:::::", res);
            // this.tableData = []
            this.$toast.success("Server IP Updated!!");
            window.location.reload()
            this.ipAdd = null;
            // this.ServerDetails = res.data;
            resolve(res.data);
            // this.$store.commit("SET_USER", res.data)
            // this.$store.commit(types.MUTATE_LOADER_OFF)
          })
          .catch((err) => {
            console.log(err.response);
            // reject(err.response)
            reject(err);
            // this.$store.commit(types.MUTATE_LOADER_OFF)
            this.$toast.error("error occured while updating server details!!!");
          });
      });
    },
  },
  mounted() {
    this.getServerByCompany();
  },
};
</script>

<style lang="scss" scoped></style>