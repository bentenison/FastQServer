<template>
    <div class="container h-100">
        <div class="row h-100">
            <div class="row col-md-12 h-100">
                <v-card class="col-md-12 h-auto" flat>

                    <!-- <div class="row d-flex">
                        <v-col cols="12" sm="3">
                            <v-text-field value="" label="Counter Number" solo></v-text-field>
                        </v-col> -->

                        <!-- <v-col cols="12" sm="3">
                            <v-text-field value="" label="Branch" solo></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="3">
                            <v-text-field value="" label="Section" solo></v-text-field>
                        </v-col> -->
                        <!-- <v-col class="pl-0 col-md-2">
                            <v-btn large text color="primary" class="mt-1 text-md-button"> <v-icon left>mdi-magnify</v-icon>
                                Search
                            </v-btn>
                        </v-col>
                    </div>
                    <v-divider></v-divider> -->
                    <EditTable :data="tableData" :headers="headers" :name="'counter'" />
                    <div class="row">
                        <div class="col-md-12 mt-5">
                            <h5 class="h4 grey--text">Add a New Counter</h5>
                            <!-- <v-col cols="12" sm="7" class="m-0 p-0">
                                <p class="h6 grey--text">Branch</p>
                                <v-select :items="items" label="Section" solo></v-select>
                            </v-col> -->
                            <!-- <div class="row d-flex flex-wrap">
                                <v-col cols="12" sm="4" class="pr-1 m-0 col-md-4 col-sm-6">
                                    <v-select :items="items" label="Section" solo></v-select>
                                </v-col>
                                <v-col cols="12" sm="4" class="pr-1 m-0 col-md-4 col-sm-6">
                                    <v-select :items="items" label="Sub Section" solo></v-select>
                                </v-col>
                                <v-col class="pl-0 col-md-4">
                                    <v-btn large text color="primary" class="mt-1 text-md-button"> Add Another
                                    </v-btn>
                                </v-col>
                            </div> -->
                            <div class="row d-flex flex-wrap">
                                <v-col cols="12" sm="4" class="">
                                    <p class="h6 grey--text">Counter Number:</p>
                                    <v-text-field label="" v-model="counter.number" class="m-0 p-0" placeholder="" outlined
                                        dense></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="4" class="">
                                    <p class="h6 grey--text">Counter Name:</p>
                                    <v-text-field label="" v-model="counter.name" class="m-0 p-0" placeholder="" outlined
                                        dense></v-text-field>
                                </v-col>
                            </div>
                            <!-- <v-col cols="12" sm="12" class="">
                                <p class="h6 grey--text">Services</p>
                                <v-checkbox v-model="ex4" label="Inquiry" color="primary" class="m-0 p-0" value="Inquiry"
                                    hide-details></v-checkbox>
                                <v-checkbox v-model="ex4" label="Payment" class="m-0 " color="primary" value="Payment"
                                    hide-details></v-checkbox>
                            </v-col> -->
                            <v-col cols="12" sm="12" class="">
                                <!-- <p class="h6 grey--text">Services</p> -->
                                <v-btn class="mr-3" color="primary" @click="addcounter"><v-icon
                                        left>mdi-content-save</v-icon>Save</v-btn>
                                <v-btn color="warning"><v-icon left>mdi-restore</v-icon>Reset</v-btn>
                            </v-col>

                        </div>
                    </div>
                    <!-- <v-btn color="primary" class=""><v-icon left>mdi-content-save-all</v-icon>Save config</v-btn> -->
                </v-card>

                <div class="row col-md-12">
                    <v-expansion-panels v-model="panel" :readonly="readonly" multiple flat style="z-index:10" v-if="false">
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

export default {
    data() {
        return {
            tableData: [],
            counter: {},
            panel: "",
            readonly: true,
            headers: [
                {
                    text: 'CounterNumber',
                    align: 'start',
                    sortable: false,
                    value: 'counterNumber',
                },
                { text: 'CounterName', value: 'counterName' },
                { text: 'BranchCode', value: 'branchCode' },
                { text: 'BranchName', value: 'branchName' },
                { text: 'CreatedAt', value: 'date' },

                { text: 'Actions', value: 'actions', sortable: false },
            ],
            items: ["India", "America", "Europe", "Japan"]
        };
    },
    components: { DataTables, EditTable },
    methods: {
        addcounter() {
            let payload = {
                counter_name: this.counter.name,
                counter_number: this.counter.number,
                branch_code: this.$store.getters.getUser.company_code,
                branch_name: this.$store.getters.getUser.company,
                company_code: this.$store.getters.getUser.company_code,
                company_name: this.$store.getters.getUser.company,
                created_by: this.$store.state.Auth.user.id,
                updated_by: this.$store.state.Auth.user.id,
            }

            // this.service.updated_by = this.$store.state.Auth.user.id
            // console.log("date", this.service.start_time);
            axios.post("/counter/addcounter", payload).then(res => {
                this.$toast.success("counter added successfully.")
                this.counter = {}
                this.getAllCounters()
            }).catch(err => {
                console.log(err.response);
                this.service.start_time = ""
                this.service.end_time = ""
                this.$toast.error("error occured while adding counter!!!")
            })
        },
        getAllCounters() {
            console.log("store>>>>>>>>>>>>>>>>>>>",);
            this.tableData = []
            axios.get(`/counter/getAllcounter/${this.$store.getters.getUser.company_code}`).then(res => {
                console.log(">>>>>>>>>>>>>>>>>", res);
                res.data.message.forEach(element => {
                    let data = {}
                    data.counterNumber = element.CounterNumber
                    data.counterName = element.CounterName
                    data.branchCode = element.BranchCode
                    data.branchName = element.BranchName
                    data.date = element.CreatedAt
                    data.id = element.ID
                    this.tableData.push(data)
                });

                this.$toast.success("counters fetched successfully.")
                // this.connection.send("Hello World !!!")
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting counters!!!")
            })
        },
    },
    mounted() {
        // this.addcounter()
        this.getAllCounters()
    }
}
</script>

<style lang="scss" scoped></style>