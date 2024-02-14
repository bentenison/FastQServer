<template>
    <div class="container h-100">
        <div class="row h-100">
            <div class="row col-md-12 h-100">
                <v-card class="col-md-12 h-auto" flat>
                    <!-- <div class="row d-flex">
                        <v-col cols="12" sm="4">
                            <v-text-field value="" label="Service Code" solo></v-text-field>
                        </v-col>

                        <v-col cols="12" sm="4">
                            <v-text-field value="" label="Service Name" solo></v-text-field>
                        </v-col>
                        <v-col class="pl-0 col-md-4">
                            <v-btn large text color="primary" class="mt-1 text-md-button"> <v-icon left>mdi-magnify</v-icon>
                                Search
                            </v-btn>
                        </v-col>
                    </div>
                    <v-divider></v-divider> -->
                    <EditTable :data="tableData" :headers="headers" :name="'service'" />
                    <v-col class="pl-0 col-md-2">
                        <v-btn large text color="success" class="mt-1 text-md-button"> <v-icon left>mdi-export</v-icon>
                            Export Excel
                        </v-btn>
                    </v-col>
                    <div class="row">
                        <div class="col-md-12 mt-5">
                            <h5 class="h4 grey--text mt-3">Add a New Service</h5>
                            <div class="col-md-12 d-flex flex-wrap">
                                <v-col cols="12" sm="5" class="m-0 p-0">
                                    <p class="h6 grey--text">Name</p>
                                    <v-text-field label="name" class="m-0 p-0" placeholder="" v-model="service.name"
                                        outlined dense></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="3" class="">
                                    <v-checkbox v-model="service.hide" label="Hide This Transaction" color="primary"
                                        class="ml-md-5" :value="true" hide-details></v-checkbox>
                                </v-col>
                                <v-col cols="12" sm="3" class="">
                                    <v-checkbox v-model="service.show_display" label="Show in Ticket Display"
                                        color="primary" class="" :value="true" hide-details></v-checkbox>
                                </v-col>
                            </div>
                            <div class="col-md-12 d-flex flex-wrap">
                                <div class="col-md-12 d-flex flex-wrap justify-content-between">
                                    <v-col cols="12" sm="2" class="m-0 p-0">
                                        <p class="h6 grey--text">Code</p>
                                        <v-text-field label="" class="m-0 p-0" v-model="service.code" placeholder=""
                                            outlined dense></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="2" class="m-0  p-0">
                                        <p class="h6 grey--text">Prefix</p>
                                        <v-text-field label="" class="m-0 p-0" placeholder="" v-model="service.prefix"
                                            outlined dense></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="3" class="m-0  p-0">
                                        <p class="h6 grey--text">Number Starts At</p>
                                        <v-text-field label="" class="m-0 p-0" placeholder=""
                                            v-model="service.number_starts" outlined dense></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="3" class="m-0  p-0">
                                        <p class="h6 grey--text">Number Ends At</p>
                                        <v-text-field label="" class="m-0 p-0" placeholder="" v-model="service.number_ends"
                                            outlined dense></v-text-field>
                                    </v-col>
                                </div>
                                <v-col cols="12" sm="5" class="m-0 mr-md-3 p-0">
                                    <p class="h6 grey--text">Description</p>
                                    <v-textarea label="" rows="3" class="m-0 p-0" placeholder="" outlined
                                        v-model="service.description" dense></v-textarea>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 ml-md-3 p-0">
                                    <p class="h6 grey--text">Show Message</p>
                                    <v-textarea label="" rows="3" class="m-0 p-0" placeholder="" outlined
                                        v-model="service.show_message" dense></v-textarea>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 mr-md-3 p-0">
                                    <p class="h6 grey--text">Start Time</p>
                                    <v-menu ref="menu" :close-on-content-click="false" top attach :nudge-right="40"
                                        :return-value.sync="service.start_time" transition="scale-transition" offset-y
                                        max-width="290px" min-width="290px">
                                        <template v-slot:activator="{ on, attrs }">
                                            <v-text-field v-model="service.start_time" label="Select start time"
                                                prepend-icon="mdi-clock-time-four-outline" readonly v-bind="attrs"
                                                v-on="on"></v-text-field>
                                        </template>
                                        <v-time-picker v-model="service.start_time" full-width
                                            @click:minute="$refs.menu.save(service.start_time)"></v-time-picker>
                                    </v-menu>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 ml-md-3 p-0">
                                    <p class="h6 grey--text">End Time</p>
                                    <v-menu ref="menu2" :close-on-content-click="false" top attach :nudge-right="40"
                                        :return-value.sync="service.end_time" transition="scale-transition" offset-y
                                        max-width="290px" min-width="290px">
                                        <template v-slot:activator="{ on, attrs }">
                                            <v-text-field v-model="service.end_time" label="Select end time"
                                                prepend-icon="mdi-clock-time-four-outline" readonly v-bind="attrs"
                                                v-on="on"></v-text-field>
                                        </template>
                                        <v-time-picker v-model="service.end_time" full-width
                                            @click:minute="$refs.menu2.save(service.end_time)"></v-time-picker>
                                    </v-menu>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 p-0">
                                    <p class="h6 grey--text">Default Processing Time</p>
                                    <v-text-field label="" class="m-0 p-0" placeholder="" outlined dense
                                        v-model="service.default_time"></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="12" class="">
                                    <!-- <p class="h6 grey--text">Services</p> -->
                                    <v-btn class="mr-3" color="primary" @click.prevent="addservice"><v-icon
                                            left>mdi-content-save</v-icon>Save</v-btn>
                                    <v-btn color="warning" @click.prevent="clear"><v-icon
                                            left>mdi-restore</v-icon>Reset</v-btn>
                                </v-col>

                            </div>

                        </div>
                    </div>
                    <!-- <v-btn color="primary" class=""><v-icon left>mdi-content-save-all</v-icon>Save config</v-btn> -->
                </v-card>

                <div class="row col-md-12">
                    <v-expansion-panels v-model="panel" multiple flat>
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
import EditTable from '../tables/editTable.vue';
import moment from 'moment';

export default {
    data() {
        return {
            connection: null,
            service: {},
            panel: null,
            tableData: [],
            headers: [
                {
                    text: 'Code',
                    align: 'start',
                    sortable: false,
                    value: 'Code',
                },
                { text: 'Name', value: 'Name' },
                { text: 'Prefix', value: 'Prefix' },
                { text: 'Start Number', value: 'NumberStarts' },
                { text: 'End Number', value: 'NumberEnds' },
                { text: 'Is Hidden?', value: 'Hide' },
                // { text: '?', value: 'isSuspended' },
                // { text: 'Branch', value: 'branch' },
                { text: 'Actions', value: 'actions', sortable: false },
            ],
            items: ["India", "America", "Europe", "Japan"]
        };
    },
    components: { DataTables, EditTable },
    methods: {
        addservice() {
            // console.log("services:::::", this.service);
            this.service.number_starts = parseInt(this.service.number_starts)
            this.service.number_ends = parseInt(this.service.number_ends)
            this.service.default_time = parseInt(this.service.default_time)
            this.service.company_code = this.$store.getters.getUser.company_code
            this.service.company_name = this.$store.getters.getUser.company
            var date1 = new Date("01-01-2017 " + this.service.start_time + ":00");
            var date1 = new Date("01-01-2017 " + this.service.end_time + ":00");
            // var date1 = new Date(`01/01/1970 ${this.service.start_time}`);
            this.service.start_time = moment(date1).format("YYYY-MM-DD HH:mm:ss")
            this.service.end_time = moment(date1).format("YYYY-MM-DD HH:mm:ss")
            this.service.updated_by = this.$store.state.Auth.user.id
            // console.log("date", this.service.start_time);
            axios.post("/service/addservice", this.service).then(res => {
                this.$toast.success("service added successfully.")
                this.getServices()
            }).catch(err => {
                console.log(err.response);
                this.service.start_time = ""
                this.service.end_time = ""
                this.$toast.error("error occured while adding service!!!")
            })
        },
        getServices() {
            console.log("store>>>>>>>>>>>>>>>>>>>",);
            axios.get(`/service/getAllservice/${this.$store.getters.getUser.company_code}`).then(res => {
                console.log(">>>>>>>>>>>>>>>>>", res);

                res.data.message.forEach(element => {
                    let data = {}
                    data.Code = element.Code
                    data.Name = element.Name
                    data.Prefix = element.Prefix
                    data.NumberStarts = element.NumberStarts
                    data.NumberEnds = element.NumberEnds
                    data.Hide = element.Hide
                    data.id = element.ID
                    this.tableData.push(data)
                });

                this.$toast.success("services fetched successfully.")
                this.connection.send("Hello World !!!")
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting service!!!")
            })
        },
        clear() {

        },
    },
    mounted() {
        this.getServices()
    },
    created: async function () {
        // console.log("Starting connection to WebSocket Server")
        // this.connection = await new WebSocket("ws://localhost:8090/ws")
        // this.connection.onmessage = function (event) {
        //     console.log(event);
        // }
        
        // this.connection.onopen = function (event) {
        //     console.log(event)
        //     console.log("Successfully connected to the echo websocket server...")
        // }
        

    }
}
</script>

<style lang="scss" scoped></style>