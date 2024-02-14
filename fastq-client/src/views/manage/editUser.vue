<template>
    <div class="container h-100">
        <div class="row h-100">
            <div class="row col-md-12 h-100">
                <v-card class="col-md-12 h-auto" flat>
                    <div class="row">
                        <div class="col-md-12 mt-5">
                            <h5 class="h4 grey--text mt-3">Edit User</h5>
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
                                <v-col cols="12" sm="5" class="m-0 mr-md-3 p-0">
                                    <p class="h6 grey--text">Title</p>
                                    <v-text-field label="title" v-model="user.Title" class="m-0 col-md-4 p-0" placeholder=""
                                        outlined dense></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 ml-md-3 p-0">
                                    <p class="h6 grey--text">Branch</p>
                                    <v-text-field label="branch" v-model="user.BranchName" class="m-0 p-0" placeholder=""
                                        outlined dense></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 mr-md-3 p-0">
                                    <p class="h6 grey--text">FirstName</p>
                                    <v-text-field label="firstName" v-model="user.Firstname" class="m-0 p-0" placeholder=""
                                        outlined dense></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 ml-md-3 p-0">
                                    <p class="h6 grey--text">LastName</p>
                                    <v-text-field label="lastName" v-model="user.Lastname" class="m-0 p-0" placeholder=""
                                        outlined dense></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 mr-md-3 p-0">
                                    <p class="h6 grey--text">Email</p>
                                    <v-text-field label="email" v-model="user.Email" class="m-0 p-0" placeholder="" outlined
                                        dense></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="5" class="m-0 p-0">
                                    <p class="h6 grey--text">Password</p>
                                    <v-text-field label="password" v-model="user.Password" class="m-0 p-0" placeholder=""
                                        outlined dense></v-text-field>
                                </v-col>
                                <!-- <v-col cols="12" sm="12" class="">
                                    <p class="h6 grey--text">Suspended</p>
                                    <v-checkbox v-model="ex4" label="Suspended" color="primary" class="m-0 p-0"
                                        value="suspended" hide-details></v-checkbox>
                                </v-col> -->
                                <v-col cols="12" sm="12" class="">
                                    <!-- <p class="h6 grey--text">Services</p> -->
                                    <v-btn class="mr-3" color="primary" @click.prevent="updateUser"><v-icon
                                            left>mdi-content-save</v-icon>Save</v-btn>
                                    <!-- <v-btn color="warning"><v-icon left>mdi-restore</v-icon>Reset</v-btn> -->
                                </v-col>

                            </div>

                        </div>
                        <div class="col-md-12" v-if="false">
                            <div class="col-md-6">
                                <h5 class="h4 grey--text">Change Password</h5>
                                <v-col cols="12" sm="12" class="m-0 p-0">
                                    <p class="h6 grey--text">Email</p>
                                    <v-text-field label="email" class="m-0 p-0" placeholder="" outlined
                                        dense></v-text-field>
                                </v-col>

                                <v-col cols="12" sm="12" class="">
                                    <!-- <p class="h6 grey--text">Services</p> -->
                                    <v-btn class="mr-3" color="primary"><v-icon left>mdi-content-save</v-icon>Save</v-btn>
                                    <v-btn color="error"><v-icon left>mdi-cancel</v-icon>Cancel</v-btn>
                                </v-col>
                            </div>
                        </div>
                    </div>
                    <!-- <v-btn color="primary" class=""><v-icon left>mdi-content-save-all</v-icon>Save config</v-btn> -->
                </v-card>

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
// import EditTable from '../tables/editTable.vue';

export default {
    data() {
        return {
            tableData: [],
            id: null,
            user: {},
            panel: null,
            readonly: false,
        };
    },
    components: { DataTables },
    methods: {
        deleteUser() {
            axios.get(`/user/deleteuser/${this.id}`).then(res => {
                console.log("res:::::", res);
                this.$toast.success("user deleted successfully.")
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while deleting user!!!")
            })
        },
        getUser() {
            let payload = {
                id: this.id,
                company_code: this.$store.state.Auth.user.company_code,
                branch_code: ""
            }
            // console.log("object:::::::", payload);
            axios.post(`/user/getuser`, payload).then(res => {
                console.log("res:::::", res);
                this.user = res.data.message
                // this.$toast.success("branch added successfully.")
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting user!!!")
            })
        },
        updateUser() {
            let payload = {
                email: this.user.Email,
                password: this.user.Password,
                firstname: this.user.Firstname,
                lastname: this.user.Lastname,
                // suspended: this.user.,
                // image_url: "",
                title: this.user.Title,
                id: this.id,
                // user_type: "",
                updated_by: this.$store.state.Auth.user.id,
            }
            axios.post(`/user/updateuser`, payload).then(res => {
                console.log("res:::::", res);
                this.$toast.success("user updated successfully.")

            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while updating user!!!")
            })
        },
    },
    mounted() {
        this.getUser()
    },
    created() {
        this.id = this.$route.params.id
    }
}
</script>

<style lang="scss" scoped></style>