<template>
    <div class="container h-100">

        <div class="row h-100">
            <div class="row col-md-12 h-100">
                <v-card class="col-md-10 h-auto" flat>
                    <div class="d-flex row align-items-center">
                        <!-- <p class="h6 grey--text">Hide Date</p> -->
                        <div class="col-md-6 p-3">
                            <h5 class="text--secondary"> Upload a New Audio</h5>
                            <v-file-input label="Select a file to upload" filled prepend-icon="mdi-video"
                                @change="handleFileChange" :loader-height="3"></v-file-input>
                            <p class="caption">Maximum file size is 1mb MB , only mp3/wav files allowed</p>
                            <v-btn color="primary" elevation="2" outlined :loading="false" @click="uploadVideo"><v-icon
                                    left>
                                    mdi-upload
                                </v-icon>Upload</v-btn>
                        </div>
                        <v-checkbox v-model="audioconf.hide_date" :label="'Hide Date'"></v-checkbox>
                        <!-- <v-checkbox :label="''"></v-checkbox> -->
                        <!-- <v-radio-group v-model="ex7" row class="m-0">
                            <v-radio label="Yes" color="primary" value="true"></v-radio>
                            <v-radio label="No" color="primary" value="false"></v-radio>
                        </v-radio-group> -->
                    </div>

                    <v-row>
                        <!-- <p class="h6 grey--text py-0">Text to Speech Message</p> -->
                        <div class="col-md-6">
                            <p class="h6 grey--text py-0">Repeat Call</p>
                            <v-text-field label="Solo field" type="number" v-model="audioconf.repeat_call"
                                solo></v-text-field>
                        </div>
                        <div class="col-md-6">
                            <p class="h6 grey--text py-0">Select Bell</p>
                            <v-select label="Solo field" :items="audios" @change="playSound" v-model="audioconf.bell_name"
                                solo></v-select>
                            <audio ref="dingAudio" controls
                                :src="`http://localhost:8090/uploaded/${audioconf.bell_name}`"></audio>
                            <!-- <audio ref="dingAudio" :src="'../../../public/audio/' + audioconf.bell_name"></audio> -->
                        </div>
                    </v-row>

                    <div class="col-md-12 d-flex align-items-center justify-content-between">
                        <div class="d-flex flex-column align-items-center">
                            <!-- <p class="h6 grey--text">Starting Bell Tone</p> -->
                            <v-checkbox v-model="audioconf.starting_bell" :label="'Starting tone for ticket'"></v-checkbox>
                        </div>
                        <div class="d-flex flex-column align-items-center">
                            <!-- <p class="h6 grey--text">Ending Bell Tone</p> -->
                            <v-checkbox v-model="audioconf.ending_bell" :label="'Ending tone for ticket'"></v-checkbox>
                        </div>
                    </div>
                    <!-- <div class="row d-flex flex-wrap align-items-center justify-content-around">
                        <v-select :items="items" label="Solo field" height="10" class="col-md-6" solo>
                            <template #append>
                            <v-btn class="mb-1" outlined color="indigo"> Copy </v-btn>
                        </template>
                        </v-select> -->
                    <!-- <div class="col-md-4 d-flex align-items-center">
                            <v-btn color="primary" text class="">Insert Keyword</v-btn>
                        </div> -->
                    <!-- </div> -->

                    <v-btn color="primary" class="" v-if="createOne" @click="addAudioConfig"><v-icon
                            left>mdi-content-save-all</v-icon>Save
                        config</v-btn>
                    <v-btn color="primary" @click="updateAudioConfig" v-if="!createOne"><v-icon
                            left>mdi-update</v-icon>Update</v-btn>
                </v-card>
                <div class="col-md-5">

                </div>
                <div class="row col-md-12">
                    <v-expansion-panels v-model="panel" :readonly="readonly" multiple flat>
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
import axios from 'axios';
import DataTables from '../tables/dataTables.vue';
import * as types from "../../store/types"
export default {
    data() {
        return {
            items: ["India", "America", "Europe", "Japan"],
            panel: "",
            readonly: false,
            audioconf: {
                id: "",
                hide_date: "",
                message: "",
                tts_lang: "",
                repeat_call: "",
                branch_code: "",
                branch_name: "",
                company_code: "",
                company_id: "",
                company_name: "",
                created_at: "",
                updated_at: "",
                updated_by: "",
                bell_name: "",
                starting_bell: "",
                ending_bell: "",
            },
            videoModel: {},
            audios: [],
            selectedFile: null,
            selectedAudio: null,
            createOne: false,
            audioConf: null
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
                modified_at: moment(event.lastModified).format('YYYY-MM-DD HH:mm:ss'),
                // created_at:,
                updated_by: this.$store.state.Auth.user.id,
                // branch_code:,
                company_code: this.$store.state.Auth.user.company_code,
                company_name: this.$store.state.Auth.user.company,
                // branch_name:,
            }
            this.selectedFile = event;
            // console.log("object", this.selectedFile);
        },
        uploadVideo() {
            const formData = new FormData();
            formData.append("video", this.selectedFile);

            axios.post("/upload", formData)
                .then(response => {
                    this.AddVideo()
                    // this.GetAllVideos()
                    // console.log(data);
                    this.selectedFile = {}
                    alert("Audio uploaded successfully!");
                })
                .catch(error => {
                    console.error("Error uploading audio:", error);
                    alert("Error uploading audio. Please try again.");
                });
        },
        AddVideo() {
            return new Promise((resolve, reject) => {
                axios.post(`/config/setvideo`, this.videoModel).then(res => {
                    // this.clients = res.data
                    console.log("res >>>>>>>>>>:::::", res);
                    this.$toast.success("Added new audio!!!")
                    resolve(res.data)
                    // this.$store.commit("SET_USER", res.data)
                    // this.$store.commit(types.MUTATE_LOADER_OFF)
                }).catch(err => {
                    console.log(err.response);
                    // reject(err.response)
                    reject(err)
                    // this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while adding audio file!!!")
                })

            })

        },
        GetAllVideos() {
            return new Promise((resolve, reject) => {
                axios.get(`/config/getallvideo/${this.$store.state.Auth.user.company_code}`).then(res => {
                    // this.allvideos = res.data
                    console.log("res >>>>>>>>>>:::::", res);
                    // this.tableData = []
                    res.data.forEach(element => {
                        if (element.type === 'audio/wav') {
                            this.audios.push(element.file_name)
                        }
                    });
                    resolve(res.data)
                    // this.$store.commit("SET_USER", res.data)
                    // this.$store.commit(types.MUTATE_LOADER_OFF)
                }).catch(err => {
                    console.log(err.response);
                    // reject(err.response)
                    reject(err)
                    // this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while getting all audios!!!")
                })
            })
        },
        addAudioConfig() {
            // console.log("object", this.audioconf);
            return new Promise((resolve, reject) => {
                // this.id = this.$route.query.company
                // if (this.id) {
                this.audioconf.company_code = this.$store.state.Auth.user.company_code
                this.audioconf.company_name = this.$store.state.Auth.user.company
                this.audioconf.updated_by = this.$store.state.Auth.user.id
                // this.parseInt(this.announcement.speed)
                this.audioconf.repeat_call = parseInt(this.audioconf.repeat_call)
                this.$store.commit(types.MUTATE_LOADER_ON);
                axios.post(`/config/addaudioconf`, this.audioconf).then(res => {
                    console.log("res :::::", res);
                    this.$toast.success("Added audio configuration successfully!!!")
                    resolve()
                    // this.audioConf = res.data
                    // this.$store.commit("SET_USER", res.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                }).catch(err => {
                    console.log(err.response.data.error.includes("sql:"));
                    // if (err.response.data.error.includes("sql:")) {
                    //     this.createOne = true
                    // }
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    reject(err.response.data)
                    this.$toast.error("error occured while adding audioconf!!!")
                })
                // }
            })
        },
        updateAudioConfig() {
            // console.log("object", this.audioconf);
            return new Promise((resolve, reject) => {
                // this.id = this.$route.query.company
                // if (this.id) {
                this.audioconf.company_code = this.$store.state.Auth.user.company_code
                this.audioconf.company_name = this.$store.state.Auth.user.company
                this.audioconf.updated_by = this.$store.state.Auth.user.id
                // this.parseInt(this.announcement.speed)
                this.audioconf.repeat_call = parseInt(this.audioconf.repeat_call)
                this.$store.commit(types.MUTATE_LOADER_ON);
                axios.post(`/config/updateaudioconf`, this.audioconf).then(res => {
                    console.log("res :::::", res);
                    this.$toast.success("audio configuration updated successfully!!!")
                    resolve()
                    // this.audioConf = res.data
                    // this.$store.commit("SET_USER", res.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                }).catch(err => {
                    console.log(err.response.data.error.includes("sql:"));
                    // if (err.response.data.error.includes("sql:")) {
                    //     this.createOne = true
                    // }
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    reject(err.response.data)
                    this.$toast.error("error occured while adding audioconf!!!")
                })
                // }
            })
        },
        getAudioConf() {
            return new Promise((resolve, reject) => {
                // this.id = this.$route.query.company
                // if (this.id) {
                this.$store.commit(types.MUTATE_LOADER_ON);
                axios.get(`/config/getaudioconf/${this.$store.state.Auth.user.company_code}`).then(res => {
                    console.log("res :::::", res);
                    this.audioconf = res.data
                    // this.$store.commit("SET_USER", res.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                }).catch(err => {
                    console.log(err.response.data.error.includes("sql:"));
                    if (err.response.data.error.includes("sql:")) {
                        this.createOne = true
                    }
                    // reject(err.response.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while getting announcement!!!")
                })
                // }
            })
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
        this.getAudioConf()
        this.GetAllVideos()
    }
}
</script>

<style lang="scss" scoped></style>