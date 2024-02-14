<template>
    <div class="container h-100">
        <div class="row h-100">
            <div class="row col-md-12 h-100">
                <v-card class="col-md-12 h-auto" flat>

                    <div class="row d-flex">
                        <v-col cols="12" sm="4">
                            <v-text-field value="" label="Branch Code" solo></v-text-field>
                        </v-col>

                        <v-col cols="12" sm="4">
                            <v-text-field value="" label="Branch Name" solo></v-text-field>
                        </v-col>
                        <v-col class="pl-0 col-md-4">
                            <v-btn large text color="primary" class="mt-1 text-md-button"> <v-icon left>mdi-magnify</v-icon>
                                Search
                            </v-btn>
                        </v-col>
                    </div>
                    <EditTable :data="tableData" :headers="headers" :name="'branch'" />

                    <h5 class="h4 grey--text mt-5">Add a New Branch</h5>
                    <div class="row">
                        <div class="col-md-5 mt-5">
                            <v-col cols="12" sm="12" class="m-0 p-0">
                                <p class="h6 grey--text">Code</p>
                                <v-text-field label="Code" class="m-0 p-0" placeholder="" v-model="branch.code" outlined
                                    dense></v-text-field>
                            </v-col>
                            <v-col cols="12" sm="12" class="m-0 p-0">
                                <p class="h6 grey--text">Name</p>
                                <v-text-field label="Name" class="m-0 p-0" placeholder="" outlined dense
                                    v-model="branch.name"></v-text-field>
                            </v-col>
                            <v-col cols="12" sm="12" class="m-0 p-0">
                                <p class="h6 grey--text">Address</p>
                                <v-text-field label="Address" class="m-0 p-0" placeholder="" outlined dense
                                    v-model="branch.address"></v-text-field>
                            </v-col>
                            <v-col cols="12" sm="12" class="m-0 p-0">
                                <p class="h6 grey--text">Phone</p>
                                <v-text-field label="Phone" class="m-0 p-0" placeholder="" outlined dense
                                    v-model="branch.phone"></v-text-field>
                            </v-col>
                            <!-- <v-col cols="12" sm="12" class="">
                                <p class="h6 grey--text">Services</p>
                                <v-checkbox v-model="services" label="Inquiry" color="primary" class="m-0 p-0"
                                    value="Inquiry" hide-details></v-checkbox>
                                <v-checkbox v-model="services" label="Payment" class="m-0 " color="primary" value="Payment"
                                    hide-details></v-checkbox>
                            </v-col> -->
                            <v-col cols="12" sm="12" class="">
                                <!-- <p class="h6 grey--text">Services</p> -->
                                <v-btn class="mr-3" color="primary" @click.prevent="addBranch"><v-icon
                                        left>mdi-content-save</v-icon>Save</v-btn>
                                <v-btn color="warning"><v-icon left>mdi-restore</v-icon>Reset</v-btn>
                            </v-col>

                        </div>
                        <div class="col-md-5 mt-5">
                            <!-- <h5 class="h4 grey--text">Add a New Branch</h5> -->
                            <v-col cols="12" sm="12" class="m-0 p-0">
                                <p class="h6 grey--text">Email</p>
                                <v-text-field label="Email" class="m-0 p-0" placeholder="" v-model="branch.email" outlined
                                    dense></v-text-field>
                            </v-col>
                            <v-col cols="12" sm="12" class="m-0 p-0">
                                <p class="h6 grey--text">Password</p>
                                <v-text-field label="Password" class="m-0 p-0" placeholder="" outlined dense
                                    v-model="branch.password"></v-text-field>
                            </v-col>

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
import DataTables from '../tables/dataTables.vue';
import EditTable from '../tables/editTable.vue';
import axios from 'axios';
export default {
    data() {
        return {
            panel: null,
            readonly: false,
            branch: {},
            branches: [],
            tableData: [],
            headers: [
                {
                    text: 'Code',
                    align: 'start',
                    sortable: false,
                    value: 'code',
                },
                { text: 'Branch Name', value: 'name' },
                { text: 'License Details', value: 'license' },
                { text: 'Services', value: 'services' },
                // { text: 'Protein (g)', value: 'protein' },
                { text: 'Actions', value: 'actions', sortable: false },
            ],
            items: ["India", "America", "Europe", "Japan"]
        };
    },
    components: { DataTables, EditTable },
    methods: {
        addBranch() {
            // this.branch.services = this.services
            console.log("services:::::", this.branch);
            this.branch.companyCode = this.$store.state.Auth.user.company_code
            axios.post("/branch/addbranch", this.branch).then(res => {
                console.log("res:::::", res);
                this.$toast.success("branch added successfully.")
                this.branch = {}
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while adding branch!!!")
            })

        },
        getBranches() {
            axios.get(`/branch/getAllBranch/${this.$store.state.Auth.user.company_code}`).then(res => {
                console.log("res:::::", res);
                // this.$toast.success("branch added successfully.")
                res.data.forEach(element => {
                    let data = {}
                    data.code = element.Code
                    data.name = element.Name
                    data.license = element.License
                    data.services = element.Services
                    data.company_code = element.CompanyCode
                    data.id = element.ID
                    // data.Hide = element.Hide
                    // data.id = element.ID
                    this.tableData.push(data)
                });
                // this.branches = res.data
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting branches!!!")
            })

        },

        reset() {
            this.branch = {}
            this.service = {}
        }
    },
    mounted() {
        this.getBranches()
    }
}
</script>

<style lang="scss" scoped></style>