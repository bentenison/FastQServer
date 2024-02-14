<template>
    <div class="container h-100">
        <div class="row h-100">
            <div class="row col-md-12 h-100">
                <v-card class="col-md-12 h-auto" flat>

                    <div class="row">
                        <div class="col-md-12 mt-5">
                            <h5 class="h4 grey--text">Edit Counter</h5>
                            <v-col cols="12" sm="7" class="m-0 p-0">
                                <p class="h6 grey--text">Branch</p>
                                <v-text-field label="Code" class="m-0 p-0" v-model="counter.BranchName" placeholder=""
                                    outlined dense></v-text-field>
                            </v-col>
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
                                <!-- <v-col cols="12" sm="4" class="">
                                    <p class="h6 grey--text">User:</p>
                                    <v-text-field label="" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                                </v-col> -->
                                <v-col cols="12" sm="4" class="">
                                    <p class="h6 grey--text">Counter Number:</p>
                                    <v-text-field label="" v-model="counter.CounterNumber" class="m-0 p-0" placeholder=""
                                        outlined dense></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="4" class="">
                                    <p class="h6 grey--text">Counter Name:</p>
                                    <v-text-field label="" v-model="counter.CounterName" class="m-0 p-0" placeholder=""
                                        outlined dense></v-text-field>
                                </v-col>
                            </div>
                            <!-- <div class="row d-flex flex-wrap">

                                <v-col cols="12" sm="4" class="">
                                    <v-checkbox v-model="ex4" label="All Queue numbers for this counter are clickable."
                                        color="primary" class="" value="Inquiry"></v-checkbox>
                                    <v-checkbox v-model="ex4" label="Allow to edit workflow" class=" " color="primary"
                                        value="Payment"></v-checkbox>
                                    <v-checkbox v-model="ex4" label="Merge Section." class=" " color="primary"
                                        value="Payment"></v-checkbox>
                                    <v-checkbox v-model="ex4" label="Can End Queue" class="" color="primary"
                                        value="Payment"></v-checkbox>
                                </v-col>
                                <v-col cols="12" sm="4" class="">
                                    <v-checkbox v-model="ex4" label="Can transfer Queue numbers. " color="primary" class=""
                                        value="Inquiry"></v-checkbox>
                                    <v-checkbox v-model="ex4" label="Can add transaction on transfer " class="m-0 "
                                        color="primary" value="Payment"></v-checkbox>
                                    <v-checkbox v-model="ex4" label="Can Assign User To Transfer Queue." class="m-0 "
                                        color="primary" value="Payment"></v-checkbox>
                                    <v-checkbox v-model="ex4" label="Can prioritize transfer Queue numbers." class="m-0 "
                                        color="primary" value="Payment"></v-checkbox>
                                    <v-checkbox v-model="ex4" label="Finish Transferred Queue Numbers." class="m-0 "
                                        color="primary" value="Payment"></v-checkbox>
                                </v-col>
                                <v-col cols="12" sm="4" class="">
                                    <v-checkbox v-model="ex4" label="Arrange Sub Section queues based on time."
                                        color="primary" class="" value="Inquiry"></v-checkbox>
                                    <v-checkbox v-model="ex4" label="Can Assign User." class="m-0 " color="primary"
                                        value="Payment"></v-checkbox>
                                    <v-checkbox v-model="ex4" label="Can cancel Queue Ticket anytime" class="m-0 "
                                        color="primary" value="Payment"></v-checkbox>
                                </v-col>
                            </div> -->
                            <v-col cols="12" sm="12" class="">
                                <!-- <p class="h6 grey--text">Services</p> -->
                                <v-btn class="mr-3" color="primary" @click="updateCounter"><v-icon
                                        left>mdi-content-save</v-icon>Save</v-btn>
                                <v-btn color="warning"><v-icon left>mdi-restore</v-icon>Reset</v-btn>
                            </v-col>

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
import EditTable from '../tables/editTable.vue';

export default {
    data() {
        return {
            panel: "",
            readonly: true,
            counter: null,
            id: null,
            items: ["India", "America", "Europe", "Japan"]
        };
    },
    components: { DataTables, EditTable },
    methods: {
        deleteCounter() {
            axios.get(`/counter/deletecounter/${this.id}`).then(res => {
                console.log("res:::::", res);
                this.$toast.success("counter deleted successfully.")
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while deleting user!!!")
            })
        },
        getcounter() {
            let payload = {
                id: this.id,
                company_code: this.$store.state.Auth.user.company_code,
                branch_code: this.$store.state.Auth.user.company_code,
            }
            // console.log("object:::::::", payload);
            axios.post(`/counter/getcounter`, payload).then(res => {
                console.log("res:::::", res);
                this.counter = res.data.message
                // this.$toast.success("branch added successfully.")
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting counter!!!")
            })
        },
        updateCounter() {
            let payload = {
                id: this.id,
                counter_name: this.counter.CounterName,
                counter_number: this.counter.CounterNumber,
                branch_code: this.$store.getters.getUser.company_code,
                branch_name: this.$store.getters.getUser.company,
                company_code: this.$store.getters.getUser.company_code,
                company_name: this.$store.getters.getUser.company,
                created_by: this.$store.state.Auth.user.id,
                updated_by: this.$store.state.Auth.user.id,
            }
            axios.post(`/counter/updatecounter`, payload).then(res => {
                console.log("res:::::", res);
                this.$toast.success("counter updated successfully.")

            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while updating counter!!!")
            })
        },
    },
    mounted() {
        this.getcounter()
    },
    created() {
        this.id = this.$route.params.id
    }
}
</script>

<style lang="scss" scoped></style>