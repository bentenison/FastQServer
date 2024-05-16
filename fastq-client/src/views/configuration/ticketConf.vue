<template>
  <div class="container h-100">
    <div class="row h-100">
      <div class="row col-md-12 h-100">
        <v-card class="col-md-6 h-auto" flat>
          <div class="d-flex align-center">
            <v-checkbox
              v-model="ticketconf.image_header"
              input-value="true"
              label="Has Image Header"
            ></v-checkbox>
            <!-- <p class="h6 grey--text">Has Image Header</p> -->
          </div>
          <div class="d-flex flex-column">
            <p class="h6 grey--text">
              No Show Expiration <span class="caption">(seconds)</span>
            </p>
            <v-text-field
              type="number"
              v-model="ticketconf.no_show_exp"
              label="label"
              id="id"
            ></v-text-field>
          </div>
          <div class="d-flex">
            <!-- <p class="h6 grey--text"></p> -->
            <v-checkbox
              label="Has Text Header"
              value
              :input-value="false"
              v-model="ticketconf.text_header"
            ></v-checkbox>
          </div>
          <div class="d-flex flex-column">
            <!-- <p class="h6 grey--text"></p> -->
            <v-checkbox
              label="Print Duplicate?"
              :input-value="false"
              v-model="ticketconf.print_duplicate"
            ></v-checkbox>
          </div>
          <div class="d-flex flex-column">
            <p class="h6 grey--text">Date Format:</p>
            <v-select
              :items="items"
              v-model="ticketconf.date_format"
              class="m-0"
              label="date format"
            ></v-select>
          </div>
          <div class="d-flex flex-column align-items-start">
            <!-- <p class="h6 grey--text"></p> -->
            <v-checkbox
              label="Print Estimated Waiting Time?"
              :input-value="false"
              v-model="ticketconf.estimated_time"
            ></v-checkbox>
          </div>

          <v-btn
            color="primary"
            class=""
            v-if="createOne"
            @click="addTicketConfig"
            ><v-icon left>mdi-content-save-all</v-icon>Save config</v-btn
          >
          <v-btn
            color="primary"
            class=""
            v-if="!createOne"
            @click="updateTicketConf"
            ><v-icon left>mdi-update</v-icon>Update config</v-btn
          >
        </v-card>
        <div class="col-md-5 p-3">
          <h5 class="h6 grey--text">Image Header</h5>
          <v-file-input
            label="Select a file to upload"
            filled
            prepend-icon="mdi-image-edit"
            @change="handleFileChange"
            show-size
            accept="'image/jpeg,image/png'"
          ></v-file-input>
          <p class="caption">
            Maximum file size is 1 MB , only jpg/png files allowed Maximum
            dimension for width is 200
          </p>
          <v-btn
            color="primary"
            elevation="2"
            outlined
            :loading="false"
            @click="uploadVideo"
            ><v-icon left> mdi-upload </v-icon>Upload</v-btn
          >
          <div class="d-flex flex-column align-items-start mt-5">
            <!-- <p class="h6 grey--text">Print Queue Position?</p> -->
            <v-checkbox
              v-model="ticketconf.print_position"
              :input-value="false"
              label="Print Queue Position?"
            ></v-checkbox>
          </div>
          <div class="d-flex flex-column align-items-start">
            <!-- <p class="h6 grey--text"></p> -->
            <v-checkbox
              label="Hide Date"
              :input-value="false"
              v-model="ticketconf.hide_time"
            ></v-checkbox>
          </div>
          <!-- <div class="d-flex flex-column p-0">
            <p class="h6 grey--text">
              Print Header
              <span class="caption">(maximum 26 characters allowed)</span>
            </p>
            <v-text-field
              class="m-0 p-0"
              name="header text"
              v-model="ticketconf.header_text"
              label="header"
              id="id"
            ></v-text-field>
          </div> -->
          <!-- <div class="d-flex flex-column">
                        <p class="h6 grey--text">Ticket Number of Digits:</p>
                        <v-select :items="items" v-model="value" class="m-0 p-0" label="country"></v-select>
                    </div> -->
        </div>
        <div class="col-md-12 p-3">
          <h5 class="h6 grey--text">Printer Configurations</h5>
          <div class="d-flex flex-column align-items-start mt-5">
            <!-- <p class="h6 grey--text">Print Queue Position?</p> -->
            <v-checkbox
              v-model="printerconf.isUsbPrinter"
              :input-value="false"
              label="USB Printer"
            ></v-checkbox>
          </div>
          <div class="d-flex flex-column align-items-start">
            <!-- <p class="h6 grey--text"></p> -->
            <v-checkbox
              label="Bluetooth Printer"
              :input-value="false"
              v-model="printerconf.isBTPrinter"
            ></v-checkbox>
          </div>
          <div class="d-flex flex-column p-0">
            <p class="h6 grey--text">
              PrinterName
              <span class="caption">(maximum 26 characters allowed)</span>
            </p>
            <v-text-field
              class="m-0 p-0"
              name="header text"
              v-model="printerconf.printerName"
              label="header"
              id="id"
            ></v-text-field>
          </div>

          <v-btn
            color="primary"
            elevation="2"
            outlined
            :loading="false"
            @click="AddPrinterConf"
            ><v-icon left> mdi-printer-3d </v-icon>Add Printer</v-btn
          >
          <!-- <div class="d-flex flex-column">
                        <p class="h6 grey--text">Ticket Number of Digits:</p>
                        <v-select :items="items" v-model="value" class="m-0 p-0" label="country"></v-select>
                    </div> -->
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
import * as types from "../../store/types";
import moment from "moment";
export default {
  data() {
    return {
      items: ["YYYY-MM-DD", "DD/MM/YYYY", "MMMM D, YYYY", "ddd, MMM D YYYY"],
      panel: "",
      readonly: false,
      ticketconf: {
        image_header: false,
        image_url: "",
        no_show_exp: "",
        text_header: false,
        header_text: "",
        hide_time: false,
        print_duplicate: false,
        date_format: "",
        print_position: false,
        estimated_time: false,
        // pop_up_delay: "",
        created_at: "",
        updated_at: "",
        updated_by: "",
        id: "",
        company_code: "",
        company_name: "",
      },
      printerconf: {},
      selectedFile: null,
      videoModel: null,
      createOne: false,
    };
  },
  components: { DataTables },
  methods: {
    playSound() {
      // console.log("object:::::");
      // this.$refs.
      // this.$refs.dingAudio.play()
    },
    handleFileChange(event) {
      console.log("object,event", event);
      this.videoModel = {
        file_name: event.name,
        type: event.type,
        size: event.size,
        modified_at: moment(event.lastModified).format("YYYY-MM-DD HH:mm:ss"),
        // created_at:,
        updated_by: this.$store.state.Auth.user.id,
        // branch_code:,
        company_code: this.$store.state.Auth.user.company_code,
        company_name: this.$store.state.Auth.user.company,
        // branch_name:,
      };
      this.ticketconf.image_url = event.name;
      this.selectedFile = event;
      // console.log("object", this.selectedFile);
    },
    uploadVideo() {
      const formData = new FormData();
      // this.selectedFile.name = "logo.png"
      this.selectedFile = new File([this.selectedFile], "logo.png", {
        type: this.selectedFile.type,
      });
      // let newFile = new File([this.selectedFile], "logo.png")
      formData.append("video", this.selectedFile);

      axios
        .post("/upload", formData)
        .then((response) => {
          // this.AddVideo();
          // this.GetAllVideos()
          // console.log(data);
          this.selectedFile = {};
          alert("Image uploaded successfully!");
        })
        .catch((error) => {
          console.error("Error uploading image:", error);
          alert("Error uploading image. Please try again.");
        });
    },
    AddPrinterConf() {
      // console.log(this.printerconf, JSON.stringify(this.printerconf));
      this.ticketconf.header_text = JSON.stringify(this.printerconf);
      this.updateTicketConf();
    },
    AddVideo() {
      return new Promise((resolve, reject) => {
        axios
          .post(`/config/setvideo`, this.videoModel)
          .then((res) => {
            // this.clients = res.data
            console.log("res >>>>>>>>>>:::::", res);
            this.$toast.success("Added new image!!!");
            resolve(res.data);
            // this.$store.commit("SET_USER", res.data)
            // this.$store.commit(types.MUTATE_LOADER_OFF)
          })
          .catch((err) => {
            console.log(err.response);
            // reject(err.response)
            reject(err);
            // this.$store.commit(types.MUTATE_LOADER_OFF)
            this.$toast.error("error occured while adding image file!!!");
          });
      });
    },
    addTicketConfig() {
      // console.log("object", this.audioconf);
      return new Promise((resolve, reject) => {
        // this.id = this.$route.query.company
        // if (this.id) {
        this.ticketconf.company_code = this.$store.state.Auth.user.company_code;
        this.ticketconf.company_name = this.$store.state.Auth.user.company;
        this.ticketconf.updated_by = this.$store.state.Auth.user.id;
        this.ticketconf.no_show_exp = parseInt(this.ticketconf.no_show_exp);
        // this.parseInt(this.announcement.speed)
        // this.ticketconf.repeat_call = parseInt(this.audioconf.repeat_call)
        this.$store.commit(types.MUTATE_LOADER_ON);
        axios
          .post(`/config/addticketconf`, this.ticketconf)
          .then((res) => {
            console.log("res :::::>>>>>>>>>>", res);
            this.$toast.success("Added ticket configuration successfully!!!");
            resolve();
            // this.audioConf = res.data
            // this.$store.commit("SET_USER", res.data)
            this.$store.commit(types.MUTATE_LOADER_OFF);
          })
          .catch((err) => {
            console.log(err.response.data.error.includes("sql:"));
            // if (err.response.data.error.includes("sql:")) {
            //     this.createOne = true
            // }
            this.$store.commit(types.MUTATE_LOADER_OFF);
            reject(err.response.data);
            this.$toast.error("error occured while adding ticketconf!!!");
          });
        // }
      });
    },
    getTicketConf() {
      return new Promise((resolve, reject) => {
        // this.id = this.$route.query.company
        // if (this.id) {
        this.$store.commit(types.MUTATE_LOADER_ON);
        axios
          .get(
            `/config/getticketconf/${this.$store.state.Auth.user.company_code}`
          )
          .then((res) => {
            console.log("res :::::", res);
            this.ticketconf = res.data;
            this.printerconf = JSON.parse(this.ticketconf.header_text)
            // this.$store.commit("SET_USER", res.data)
            this.$store.commit(types.MUTATE_LOADER_OFF);
          })
          .catch((err) => {
            // console.log(err.response.data.error.includes("sql:"));
            this.$store.commit(types.MUTATE_LOADER_OFF);
            if (err.response.data.error.includes("sql:")) {
              this.createOne = true;
            }
            // reject(err.response.data)
            this.$toast.error("error occured while getting ticket config!!!");
          });
        // }
      });
    },
    updateTicketConf() {
      return new Promise((resolve, reject) => {
        // this.id = this.$route.query.company
        // if (this.id) {
        this.$store.commit(types.MUTATE_LOADER_ON);
        axios
          .post(`/config/updateticketconf`, this.ticketconf)
          .then((res) => {
            console.log("res :::::", res);
            // this.$store.commit("SET_USER", res.data)
            this.$store.commit(types.MUTATE_LOADER_OFF);
            this.$toast.success("Updated ticket config!!!");
          })
          .catch((err) => {
            console.log(err.response);
            reject(err.response);
            this.$store.commit(types.MUTATE_LOADER_OFF);
            this.$toast.error("error occured while adding announcement!!!");
          });
        // }
      });
    },
    // updateAudioConf() {
    //     return new Promise((resolve, reject) => {
    //         // this.id = this.$route.query.company
    //         // if (this.id) {
    //         this.$store.commit(types.MUTATE_LOADER_ON);
    //         axios.post(`/config/updateannounce`).then(res => {
    //             console.log("res :::::", res);
    //             // this.$store.commit("SET_USER", res.data)
    //             this.$store.commit(types.MUTATE_LOADER_OFF)
    //         }).catch(err => {
    //             console.log(err.response);
    //             reject(err.response)
    //             this.$store.commit(types.MUTATE_LOADER_OFF)
    //             this.$toast.error("error occured while adding announcement!!!")
    //         })
    //         // }
    //     })
    // },
  },
  created() {
    this.getTicketConf();
    // this.GetAllVideos()
  },
};
</script>

<style lang="scss" scoped></style>