<template>
    <div class="container h-100">
        <div class="row h-100">
            <div class="row col-md-12 h-100">
                <v-card class="col-md-12 h-auto" flat>

                    <div class="row d-flex" v-if="false">
                        <v-col cols="12" sm="3">
                            <v-select value="" :items="branches" item-text="Name" item-value="" return-object label="Branch"
                                solo></v-select>
                        </v-col>
                        <!-- <v-col cols="12" sm="4">
                            <v-select value="" label="User Type" solo></v-select>
                        </v-col> -->
                        <v-col cols="12" sm="4">
                            <v-text-field value="" label="User Name" solo></v-text-field>
                        </v-col>

                        <v-col cols="12" sm="4">
                            <v-text-field value="" label="First Name" solo></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="4">
                            <v-text-field value="" label="Last Name" solo></v-text-field>
                        </v-col>
                        <v-col class="pl-0 col-md-2">
                            <v-btn large text color="primary" class="mt-1 text-md-button"> <v-icon left>mdi-magnify</v-icon>
                                Search
                            </v-btn>
                        </v-col>
                    </div>
                    <v-divider></v-divider>
                    <EditTable :data="tableData" :headers="headers" :name="'user'" />
                    <v-col class="pl-0 col-md-2">
                        <v-btn large text color="success" class="mt-1 text-md-button"> <v-icon left>mdi-export</v-icon>
                            Export Excel
                        </v-btn>
                    </v-col>
                    <div class="row">
                        <div class="col-md-12 mt-5">
                            <h5 class="h4 grey--text mt-3">Add a New Staff</h5>
                            <v-col cols="6" sm="6" class="">
                                <p class="h6 grey--text">Code</p>
                                <v-file-input label="Select a file to upload" filled
                                    prepend-icon="mdi-image-edit"></v-file-input>
                                <p class="caption">Maximum file size is 1 MB , only jpg/png files allowed Maximum dimension
                                    for width is
                                    200</p>
                                <v-btn color="primary" elevation="2" outlined :loading="false"><v-icon left>
                                        mdi-upload
                                    </v-icon>Upload</v-btn>
                            </v-col>
                            <div class="col-md-12 d-flex flex-wrap">
                                <!-- <v-col cols="12" sm="5" class="m-0 mr-md-3 p-0">
                                    <p class="h6 grey--text">User Type</p>
                                    <v-text-field label="user type" class="m-0 p-0" placeholder="" outlined
                                        dense></v-text-field>
                                </v-col> -->
                                <v-col cols="12" sm="10" class="ml-n3">
                                    <p class="h6 grey--text">Title</p>
                                    <v-text-field v-model="user.title" label="title" class="m-0 col-md-4 p-0" placeholder=""
                                        solo></v-text-field>
                                </v-col>
                                <!-- <v-col cols="12" sm="5" class="mr-3 p-0">
                                    <p class="h6 grey--text">Branch</p>
                                    <v-select v-model="branch" :items="branches" item-text="Name" item-value=""
                                        return-object label="Branch" solo></v-select>
                                </v-col> -->
                                <v-col cols="12" sm="5" class="m-0 mr-md-3 p-0">
                                    <p class="h6 grey--text">FirstName</p>
                                    <v-text-field v-model="user.firstname" label="firstName" class="m-0 p-0" placeholder=""
                                        solo></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 mr-3 p-0">
                                    <p class="h6 grey--text">LastName</p>
                                    <v-text-field v-model="user.lastname" label="lastName" class="m-0 p-0" placeholder=""
                                        solo></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 mr-md-3 p-0">
                                    <p class="h6 grey--text">Email</p>
                                    <v-text-field v-model="user.email" label="email" class="m-0 p-0" placeholder=""
                                        solo></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 ml-md-3 p-0">
                                    <p class="h6 grey--text">Password</p>
                                    <v-text-field v-model="user.password" label="password" class="m-0 p-0" placeholder=""
                                        solo></v-text-field>
                                </v-col>
                                <!-- <v-col cols="12" sm="12" class="">
                                    <p class="h6 grey--text">Suspended</p>
                                    <v-checkbox v-model="ex4" label="Suspended" color="primary" class="m-0 p-0"
                                        value="suspended" hide-details></v-checkbox>
                                </v-col> -->
                                <v-col cols="12" sm="12" class="">
                                    <!-- <p class="h6 grey--text">Services</p> -->
                                    <v-btn class="mr-3" color="primary" @click="addUser"><v-icon
                                            left>mdi-content-save</v-icon>Save</v-btn>
                                    <v-btn color="warning" @click.prevent="reset"><v-icon
                                            left>mdi-restore</v-icon>Reset</v-btn>
                                </v-col>
                            </div>

                        </div>
                    </div>
                    <!-- <v-btn color="primary" class=""><v-icon left>mdi-content-save-all</v-icon>Save config</v-btn> -->
                </v-card>

                <div class="row col-md-12">
                    <v-expansion-panels v-model="panel" :readonly="readonly" multiple flat style="z-index:10">
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
            branches: [],
            users: [],
            panel: null,
            readonly: false,
            branch: null,
            user: {},
            tableData: [],
            headers: [
                {
                    text: 'Email',
                    align: 'start',
                    sortable: false,
                    value: 'email',
                },
                { text: 'FirstName', value: 'firstname' },
                { text: 'LastName', value: 'lastname' },
                { text: 'Date Created', value: 'date' },
                // { text: 'User Type', value: 'uType' },
                // { text: 'Suspended?', value: 'isSuspended' },
                { text: 'Branch', value: 'branch' },
                { text: 'Actions', value: 'actions', sortable: false },
            ],
            items: ["India", "America", "Europe", "Japan"],


        };
    },
    components: { DataTables, EditTable },
    methods: {
        getBranches() {
            axios.get(`/branch/getAllBranch/${this.$store.state.Auth.user.company_code}`).then(res => {
                console.log("res:::::", res);
                // this.$toast.success("branch added successfully.")
                this.branches = res.data
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting branches!!!")
            })

        },
        addUser() {
            // this.user.branch_name = this.branch.Name
            // this.user.branch_code = this.branch.Code
            this.user.company_code = this.$store.state.Auth.user.company_code
            this.user.created_by = this.$store.state.Auth.user.id
            this.user.updated_by = this.$store.state.Auth.user.id
            this.user.company_name = this.$store.state.Auth.user.company
            // console.log("usr", this.user);
            if (!this.user.email || this.user.email === "") {
                this.$toast.error("email is a required parameter")
                return
            }
            axios.post(`/user/adduser`, this.user).then(res => {
                console.log("res:::::", res);
                this.$toast.success("user added successfully.")
                this.branches = res.data
                this.reset()
                this.getAllUser()
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while adding user!!!")
            })

        },
        getAllUser() {
            let payload = {
                company_code: this.$store.state.Auth.user.company_code,
                branch_code: ""
            }
            this.tableData = []
            axios.post(`/user/getAlluser`, payload).then(res => {
                console.log("res:::::", res);
                // this.$toast.success("branch added successfully.")
                res.data.message.forEach(element => {
                    let data = {}
                    data.email = element.Email
                    data.firstname = element.Firstname
                    data.lastname = element.Lastname
                    data.date = element.CreatedAt
                    data.branch = element.CompanyCode
                    data.id = element.ID
                    // data.Hide = element.Hide
                    // data.id = element.ID
                    this.tableData.push(data)
                });
                this.users = res.data
                // this.reset()
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting all user!!!")
            })
        },
        getUser() {
            let payload = {
                company_code: "",
                branch_code: ""
            }
            axios.post(`/user/getuser`, this.user).then(res => {
                console.log("res:::::", res);
                // this.$toast.success("branch added successfully.")
                this.branches = res.data
                this.reset()
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while adding user!!!")
            })
        },
        reset() {
            this.branch = {}
            this.user = {}
        },
    },
    mounted() {
        this.getBranches()
        this.getAllUser()
    }
}
</script>

<style lang="scss" scoped></style>