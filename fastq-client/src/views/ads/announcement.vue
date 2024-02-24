<template>
    <div class="container h-100">
        <div class="row h-100">
            <div class="col-md-12 h-100">
                <v-card class="row h-100" flat>
                    <div class="row align-self-start justify-content-start mt-5">
                        <div class="col-md-12">
                            <h5 class="grey--text pb-2">All Announcements</h5>
                            <!-- <DataTables :headers="headers" :items="allannouncement" /> -->
                            <v-data-table v-model="selected" :headers="headers" :items="allannouncement"
                                :single-select="true" item-key="announce_text" show-select class="elevation-1">
                                <template v-slot:foot>
                                    <v-row class="m-2">
                                        <v-col cols="12" class="p-2">
                                            <v-btn color="primary" @click="selectAnnouncement"><v-icon
                                                    left>mdi-plus</v-icon>Add to
                                                selected</v-btn>
                                        </v-col>
                                    </v-row>
                                </template>
                            </v-data-table>
                        </div>
                        <div class="col-md-12">
                            <h5 class="grey--text pb-2">Display Announcement</h5>
                            <v-textarea class="" solo name="input-7-4" v-model="announcement.announce_text"
                                label="Enter announcement text"></v-textarea>
                        </div>
                        <div class="col-md-4">
                            <h5 class="grey--text pb-2">Announcement Speed</h5>
                            <v-text-field class="" label="speed" type="number" v-model="announcement.speed" outlined
                                dense></v-text-field>
                            <v-btn color="primary" @click="addAnnouncement"><v-icon
                                    left>mdi-content-save</v-icon>Save</v-btn>
                            <!-- <v-btn color="primary" @click="updateAnnouncement" v-if="!createOne"><v-icon
                                    left>mdi-update</v-icon>Update</v-btn> -->
                        </div>
                        <div class="row col-md-12 mt-3">
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
                </v-card>
            </div>
        </div>
    </div>
</template>

<script>
import DataTables from '../tables/dataTables.vue';
import * as types from "../../store/types"
import axios from 'axios';
export default {
    components: { DataTables },
    data() {
        return {
            selected: null,
            panel: "",
            readonly: false,
            announcement: {},
            id: null,
            createOne: false,
            resAnnouncement: null,
            allannouncement: [],
            headers: [
                {
                    text: 'AnnounceText',
                    align: 'start',
                    sortable: true,
                    value: 'announce_text',
                },
                { text: 'Speed', value: 'speed' },
                // { text: 'Created At', value: 'size' },
                { text: 'Date Created', value: 'created_at' },
                // { text: 'User Type', value: 'uType' },
                // { text: 'Suspended?', value: 'isSuspended' },
                // { text: 'Modified At', value: 'modified_at' },
                // { text: 'Actions', value: 'actions', sortable: false },
            ],

        }
    },
    methods: {
        addAnnouncement() {
            return new Promise((resolve, reject) => {
                // this.id = this.$route.query.company
                // if (this.id) {
                this.announcement.speed = parseInt(this.announcement.speed)
                this.announcement.company_code = this.$store.state.Auth.user.company_code
                this.announcement.company_code = this.$store.state.Auth.user.company_code
                this.announcement.updated_by = this.$store.state.Auth.user.id
                this.$store.commit(types.MUTATE_LOADER_ON);
                axios.post(`/config/addannounce`, this.announcement).then(res => {
                    console.log("res :::::", res);
                    // this.$store.commit("SET_USER", res.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.success("announcement added successfully!!!")
                }).catch(err => {
                    console.log(err.response);
                    reject(err.response)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while adding announcement!!!")
                })
                // }
            })
        },
        selectAnnouncement() {
            return new Promise((resolve, reject) => {
                // this.id = this.$route.query.company
                // if (this.id) {
                if (!this.selected) {
                    this.$toast.error("select one announcement fisrt!!!")
                    return
                }
                this.$store.commit(types.MUTATE_LOADER_ON);
                axios.get(`/config/selecttodisplay/${this.$store.state.Auth.user.company_code}?id=${this.selected[0].id}`).then(res => {
                    console.log("res :::::", res);
                    this.$toast.success("announcement set as default successfully!!!")
                    // this.$store.commit("SET_USER", res.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                }).catch(err => {
                    console.log(err.response);
                    reject(err.response)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while adding announcement!!!")
                })
                // }
            })
        },
        updateAnnouncement() {
            return new Promise((resolve, reject) => {
                // this.id = this.$route.query.company
                // if (this.id) {
                this.$store.commit(types.MUTATE_LOADER_ON);
                axios.post(`/config/updateannounce`).then(res => {
                    console.log("res :::::", res);
                    // this.$store.commit("SET_USER", res.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                }).catch(err => {
                    console.log(err.response);
                    reject(err.response)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while adding announcement!!!")
                })
                // }
            })
        },
        getAnnouncement() {
            return new Promise((resolve, reject) => {
                // this.id = this.$route.query.company
                // if (this.id) {
                this.$store.commit(types.MUTATE_LOADER_ON);
                axios.get(`/config/getannounce/${this.$store.state.Auth.user.company_code}`).then(res => {
                    console.log("res :::::", res);
                    this.announcement = res.data
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
        getAllAnnouncement() {
            return new Promise((resolve, reject) => {
                // this.id = this.$route.query.company
                // if (this.id) {
                this.$store.commit(types.MUTATE_LOADER_ON);
                axios.get(`/config/getallannounce/${this.$store.state.Auth.user.company_code}`).then(res => {
                    console.log("res :::::", res);
                    this.allannouncement = res.data
                    // this.$store.commit("SET_USER", res.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                }).catch(err => {
                    console.log(err.response.data.error.includes("sql:"));
                    // if (err.response.data.error.includes("sql:")) {
                    //     // this.createOne = true
                    // }
                    // reject(err.response.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while getting announcement!!!")
                })
                // }
            })
        },
    },
    mounted() {
        // this.getAnnouncement()
        this.getAllAnnouncement()
    }

}
</script>

<style lang="scss" scoped></style>