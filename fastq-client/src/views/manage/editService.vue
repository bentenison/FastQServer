<template>
    <div class="container h-100">
        <div class="row h-100">
            <div class="row col-md-12 h-100">
                <v-card class="col-md-12 h-auto" flat>
                    <div class="row">
                        <div class="col-md-12 mt-5">
                            <h5 class="h4 grey--text mt-3">Edit Service</h5>
                            <div class="col-md-12 d-flex flex-wrap">
                                <v-col cols="12" sm="5" class="m-0 p-0">
                                    <p class="h6 grey--text">Name</p>
                                    <v-text-field label="" v-model="service.Name" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="3" class="">
                                    <v-checkbox v-model="service.Hide" label="Hide This Transaction" color="primary" class="ml-md-5"
                                        :value="true"  hide-details></v-checkbox>
                                </v-col>
                                <v-col cols="12" sm="3" class="">
                                    <v-checkbox v-model="service.ShowDisplay" label="Show in Ticket Display" color="primary" class=""
                                        :value="true" hide-details></v-checkbox>
                                </v-col>
                            </div>
                            <div class="col-md-12 d-flex flex-wrap">
                                <div class="col-md-12 d-flex flex-wrap justify-content-between">
                                    <v-col cols="12" sm="2" class="m-0 p-0">
                                        <p class="h6 grey--text">Code</p>
                                        <v-text-field label="" v-model="service.Code" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="2" class="m-0  p-0">
                                        <p class="h6 grey--text">Prefix</p>
                                        <v-text-field label="" v-model="service.Prefix" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="3" class="m-0  p-0">
                                        <p class="h6 grey--text">Number Starts At</p>
                                        <v-text-field label="" v-model="service.NumberStarts" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="3" class="m-0  p-0">
                                        <p class="h6 grey--text">Number Ends At</p>
                                        <v-text-field label="" v-model="service.NumberEnds" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                                    </v-col>
                                </div>
                                <v-col cols="12" sm="5" class="m-0 mr-md-3 p-0">
                                    <p class="h6 grey--text">Description</p>
                                    <v-textarea label="" v-model="service.Description" rows="3" class="m-0 p-0" placeholder="" outlined
                                        dense></v-textarea>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 ml-md-3 p-0">
                                    <p class="h6 grey--text">Show Message</p>
                                    <v-textarea label=""  rows="3" class="m-0 p-0" placeholder="" outlined
                                        dense></v-textarea>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 mr-md-3 p-0">
                                    <p class="h6 grey--text">Start Time</p>
                                    <v-text-field label="" v-model="service.StartTime" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 ml-md-3 p-0">
                                    <p class="h6 grey--text">End Time</p>
                                    <v-text-field label="" v-model="service.EndTime" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 p-0">
                                    <p class="h6 grey--text">Default Processing Time</p>
                                    <v-text-field label="" v-model="service.DefaultTime" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="12" class="">
                                    <!-- <p class="h6 grey--text">Services</p> -->
                                    <v-btn class="mr-3" color="primary"><v-icon left>mdi-content-save</v-icon>Save</v-btn>
                                    <v-btn color="warning"><v-icon left>mdi-restore</v-icon>Reset</v-btn>
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

export default {
    data() {
        return {
            id: "",
            service: {},
            panel:"",
            tableData: [{
                email: 'john@gmail.com',
                firstName: "John",
                lastName: "Doe",
                // license: 6.0,
                date: '1-06-2023',
                uType: `Admin`,
                isSuspended: `No`,
                branch: `Head Office`,
            }],
            headers: [
                {
                    text: 'Email',
                    align: 'start',
                    sortable: false,
                    value: 'email',
                },
                { text: 'FirstName', value: 'firstName' },
                { text: 'LastName', value: 'lastName' },
                { text: 'Date Created', value: 'date' },
                { text: 'User Type', value: 'uType' },
                { text: 'Suspended?', value: 'isSuspended' },
                { text: 'Branch', value: 'branch' },
                { text: 'Actions', value: 'actions', sortable: false },
            ],
            items: ["India", "America", "Europe", "Japan"]
        };
    },
    methods: {
        getService() {
            let payload = {}
            payload.id = this.id
            payload.company_code = this.$store.getters.getUser.company_code
            axios.post(`/service/getservice`, payload).then(res => {
                console.log(">>>>>>>>>>>>>>>>>", res);
                this.service = res.data.message
                this.$toast.success("service fetched successfully.")
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting service!!!")
            })
        }
    },
    components: { DataTables, EditTable },
    mounted() {
        // console.log(">>>>>>>>>>>>>>>>>>>>>>route",this.$route);
        this.id = this.$route.params.id
        this.getService()
    },
}
</script>

<style lang="scss" scoped></style>