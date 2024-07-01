<template>
  <div class="container h-100" :key="tableData.length">
    <div class="row h-100">
      <div class="col-md-12 h-100">
        <v-card class="row h-100">
          <DataTables :headers="headers" :items="tableData" :name="'videos'" />

          <v-card class="h-70" flat>
            <div class="col-md-6 p-3">
              <h5 class="text--secondary">Upload a New Video</h5>
              <v-file-input
                label="Select a file to upload"
                filled
                prepend-icon="mdi-video"
                v-model="selectedFile"
                @change="handleFileChange"
                :loader-height="3"
              ></v-file-input>
              <p class="caption">
                Maximum file size is 100 MB , only mp4/MP4 files allowed
              </p>
              <v-btn
                color="primary"
                elevation="2"
                outlined
                :loading="false"
                @click="uploadVideo"
                ><v-icon left> mdi-upload </v-icon>Upload</v-btn
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
          </v-card>
        </v-card>
      </div>
      <!-- <div class="col-md-12 h100">
          
        </div> -->
    </div>
  </div>
</template>

<script>
import axios from "axios";
import DataTables from "../tables/dataTables.vue";
import DraggableTable from "../tables/draggableTable.vue";
import moment from "moment";

export default {
  components: { DataTables, DraggableTable },
  data() {
    return {
      selectedFile: null,
      panel: "",
      readonly: false,
      videoModel: {},
      allvideos: [],
      tableData: [],
      headers: [
        {
          text: "FileName",
          align: "start",
          sortable: true,
          value: "file_name",
        },
        { text: "Type", value: "type" },
        { text: "Size", value: "size" },
        { text: "Date Created", value: "created_at" },
        // { text: 'User Type', value: 'uType' },
        // { text: 'Suspended?', value: 'isSuspended' },
        { text: "Modified At", value: "modified_at" },
        { text: "Actions", value: "actions", sortable: false },
      ],
    };
  },
  methods: {
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
      this.selectedFile = event;
      // console.log("object", this.selectedFile);
    },
    uploadVideo() {
      const formData = new FormData();
      formData.append("video", this.selectedFile);

      axios
        .post("/upload", formData)
        .then((response) => {
          this.AddVideo();
          // console.log(data);
          this.selectedFile = {};
          // alert("Video uploaded successfully!");
          this.GetAllVideos();
        })
        .catch((error) => {
          console.error("Error uploading video:", error);
          alert("Error uploading video. Please try again.");
        });
    },
    GetAllVideos() {
      return new Promise((resolve, reject) => {
        axios
          .get(
            `/config/getallvideo/${this.$store.state.Auth.user.company_code}`
          )
          .then((res) => {
            this.allvideos = res.data;
            console.log("res >>>>>>>>>>:::::", res);
            this.tableData = [];
            if (res.data) {
              res.data.forEach((element) => {
                let data = {};
                data.file_name = element.file_name;
                data.type = element.type;
                data.size = element.size;
                data.modified_at = moment(element.modified_at).format(
                  "MMMM DD, YYYY HH:mm:ss"
                );
                data.created_at = moment(element.created_at).format(
                  "MMMM DD, YYYY HH:mm:ss"
                );
                data.id = element.id;
                // data.Hide = element.Hide
                // data.id = element.ID
                this.tableData.push(data);
              });
              resolve(res.data);
            }
            // this.$store.commit("SET_USER", res.data)
            // this.$store.commit(types.MUTATE_LOADER_OFF)
          })
          .catch((err) => {
            console.log(err.response);
            // reject(err.response)
            reject(err);
            // this.$store.commit(types.MUTATE_LOADER_OFF)
            this.$toast.error(
              "error occured while getting connected clients!!!"
            );
          });
      });
    },
    AddVideo() {
      return new Promise((resolve, reject) => {
        axios
          .post(`/config/setvideo`, this.videoModel)
          .then((res) => {
            // this.clients = res.data
            console.log("res >>>>>>>>>>:::::", res);
            this.$toast.success("Added new video!!!");
            resolve(res.data);
            // this.$store.commit("SET_USER", res.data)
            // this.$store.commit(types.MUTATE_LOADER_OFF)
          })
          .catch((err) => {
            console.log(err.response);
            // reject(err.response)
            reject(err);
            // this.$store.commit(types.MUTATE_LOADER_OFF)
            this.$toast.error("error occured while adding video config!!!");
          });
      });
    },
  },
  created() {
    this.GetAllVideos();
  },
};
</script>

<style lang="scss" scoped></style>