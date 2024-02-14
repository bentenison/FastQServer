<template>
    <div class="container h-100">
        <div class="row h-100">
            <div class="col-md-12 h-100">
                <v-card class="row h-100" flat>
                    <v-card class="" v-for="(day, index) in days" :key="day" flat>
                        <div class="d-flex">
                            <h3 class="grey--text p-3">{{ day }}</h3>
                            <v-spacer></v-spacer>
                            <div class="d-md-flex d-xs-none align-items-center flex-wrap p-3">
                                <v-chip class="ma-2" label color="warning" text-center text-color=""
                                    @click="addvideos(day)">
                                    <v-icon left>
                                        mdi-plus-circle
                                    </v-icon>
                                    <h5 class="font-weight-regular h6 m-auto">Add Video</h5>
                                </v-chip>
                                <!-- <v-chip class="ma-2" label color="info" text-color="white">
                                    <v-icon left>
                                        mdi-sort
                                    </v-icon>
                                    <h5 class="font-weight-regular h6 m-auto">Sort Order</h5>
                                </v-chip> -->
                            </div>
                        </div>
                        <DraggableTable :uclass="day" :headers="headers" :items="videos[day]"
                            v-if="done && videos[day].length !== 0" />
                        <p class="grey--text" v-else>No data to show. Add videos</p>
                    </v-card>
                </v-card>
            </div>

        </div>
        <v-row justify="center">
            <v-dialog v-model="dialog" fullscreen hide-overlay transition="dialog-bottom-transition">
                <v-card>
                    <v-toolbar dark color="primary">
                        <v-btn icon dark @click="dialog = false">
                            <v-icon>mdi-close</v-icon>
                        </v-btn>
                        <v-toolbar-title>Select videos to add into schedular</v-toolbar-title>
                        <v-spacer></v-spacer>
                        <v-toolbar-items>
                            <v-btn dark text @click="addSelectedvideos()">
                                Save
                            </v-btn>
                        </v-toolbar-items>
                    </v-toolbar>
                    <v-row class="d-flex align-center justify-center mt-5" fill-height>
                        <v-col cols="10">
                            <v-data-table v-model="selected" :headers="headers" :items="AllVideos" :single-select="false"
                                item-key="file_name" show-select class="elevation-1">
                                <!-- <template v-slot:top>
                                    <v-switch v-model="singleSelect" label="Single select" class="pa-3"></v-switch>
                                </template> -->
                            </v-data-table>
                        </v-col>
                    </v-row>
                </v-card>
            </v-dialog>
        </v-row>
    </div>
</template>

<script>
import axios from 'axios';
import DataTables from '../tables/dataTables.vue';
import DraggableTable from '../tables/draggableTable.vue';

export default {
    components: { DataTables, DraggableTable },
    data() {
        return {
            days: ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'],
            selected: null,
            videos: {},
            AllVideos: [],
            dialog: false,
            done: false,
            headers: [
                {
                    text: '',
                    sortable: true,
                    value: '',
                },
                {
                    text: 'Index',
                    sortable: true,
                    value: '',
                },
                {
                    text: 'FileName',
                    sortable: true,
                    value: 'file_name',
                },
                { text: 'Type', value: 'type' },
                { text: 'Size', value: 'size' },
                { text: 'Date Created', value: 'created_at' },
                // { text: 'User Type', value: 'uType' },
                // { text: 'Suspended?', value: 'isSuspended' },
                { text: 'Company Name', value: 'company_name' },
                // { text: 'Actions', value: 'actions', sortable: false },
            ],
            videoSchedular: {
                id: "",
                video_id: "",
                day: "",
                ordinality: "",
                created_at: "",
                created_by: "",
                updated_at: "",
                updated_by: "",
                branch_code: "",
                company_code: "",
            },
            selectedIndex: null,
        }
    },
    methods: {
        addvideos(index) {
            this.dialog = true
            this.selectedIndex = index
        },
        getAllSchedulars() {
            return new Promise((resolve, reject) => {
                axios.get(`/config/getallschedularconfig/${this.$store.state.Auth.user.company_code}`).then(res => {
                    // this.allvideos = res.data
                    // console.log("res >>>>>>>>>>:::::", res);
                    // this.tableData = []
                    // res.data.forEach(element => {
                    //     if (element.type === 'video/mp4') {
                    //         this.AllVideos.push(element)
                    //     }
                    // });
                    console.log("res >>>>>>>>>>:::::", res.data);

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
        GetAllVideos() {
            return new Promise((resolve, reject) => {
                axios.get(`/config/getallvideo/${this.$store.state.Auth.user.company_code}`).then(res => {
                    // this.allvideos = res.data
                    // console.log("res >>>>>>>>>>:::::", res);
                    // this.tableData = []
                    res.data.forEach(element => {
                        if (element.type === 'video/mp4') {
                            this.AllVideos.push(element)
                        }
                    });
                    console.log("res >>>>>>>>>>:::::", this.AllVideos);
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
        addSelectedvideos() {
            console.log(" selected objects", this.selected);
            console.log("selected index", this.selectedIndex);
            this.videos[this.selectedIndex].push(...this.selected)
            this.videos[this.selectedIndex].forEach((element, index) => {
                let payload = {
                    video_id: element.file_name,
                    day: this.selectedIndex,
                    ordinality: index + 1,
                    company_code: this.$store.state.Auth.user.company_code,
                }
                this.AddSchedular(payload)
            });
            this.dialog = false
            console.log("all videos", this.videos);
        },
        AddSchedular(data) {
            return new Promise((resolve, reject) => {
                axios.post(`/config/setschedularconfig`, data).then(res => {
                    // this.clients = res.data
                    console.log("res >>>>>>>>>>:::::", res);
                    this.$toast.success("added video schedular!!!")
                    resolve(res.data)
                    // this.$store.commit("SET_USER", res.data)
                    // this.$store.commit(types.MUTATE_LOADER_OFF)
                }).catch(err => {
                    console.log(err.response);
                    // reject(err.response)
                    reject(err)
                    // this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while adding video schedular config!!!")
                })

            })

        },
    },
    created() {
        this.days.forEach((element, index) => {
            this.videos[element] = []
        });
        this.GetAllVideos()
        this.getAllSchedulars().then(res => {
            res.forEach(element => {
                console.log("for eachh element", element);
                let index = this.AllVideos.findIndex((a) => a.file_name === element.video_id)
                console.log("Index", index);
                this.videos[element.day].push(this.AllVideos[index])
                console.log("videos map", this.videos);
            });
            this.done = true
        })
    }
}
</script>

<style lang="scss" scoped></style>