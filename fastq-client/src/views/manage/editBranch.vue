<template>
    <div class="container h-100">
        <div class="row h-100">
            <div class="row col-md-12 h-100">
                <v-card class="col-md-12 h-auto" flat>
                    <div class="row justify-content-between">
                        <div class="col-md-5 mt-5">
                            <h5 class="h4 grey--text">Add a New Branch</h5>
                            <v-divider inset height="3"></v-divider>
                            <v-col cols="12" sm="12" class="m-0 p-0">
                                <p class="h6 grey--text">Code</p>
                                <v-text-field label="Code" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                            </v-col>
                            <v-col cols="12" sm="12" class="m-0 p-0">
                                <p class="h6 grey--text">Name</p>
                                <v-text-field label="Name" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                            </v-col>
                            <v-col cols="12" sm="12" class="m-0 p-0">
                                <p class="h6 grey--text">Address</p>
                                <v-text-field label="Address" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                            </v-col>
                            <v-col cols="12" sm="12" class="m-0 p-0">
                                <p class="h6 grey--text">Phone</p>
                                <v-text-field label="Phone" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                            </v-col>
                            <v-col cols="12" sm="12" class="m-0 p-0">
                                <p class="h6 grey--text">License</p>
                                <v-textarea filled dense rows="2" name="input-7-4" label="key"
                                    value=""></v-textarea>
                                <p class="caption" v-html="html"></p>
                            </v-col>
                            <v-col cols="12" sm="12" class="">
                                <p class="h6 grey--text">Services</p>
                                <v-checkbox v-model="ex4" label="Inquiry" color="primary" class="m-0 p-0" value="Inquiry"
                                    hide-details></v-checkbox>
                                <v-checkbox v-model="ex4" label="Payment" class="m-0 " color="primary" value="Payment"
                                    hide-details></v-checkbox>
                            </v-col>
                            <v-col cols="12" sm="12" class="">
                                <p class="h6 grey--text pb-2 my-0">Customer Check in Ticket Page for this Branch:</p>
                                <router-link to="/edit/branch" class="grey--text pb-2">http://localhost:8080/#/edit/branch</router-link>
                                <p class="h6 grey--text pb-2 my-0 mt-2">Generate Ticket Page without login for this Branch:</p>
                                <router-link to="/edit/branch" class="grey--text">http://localhost:8080/#/edit/branch</router-link>
                                <p class="h6 grey--text pb-2 my-0 mt-2">TV Display without login for this Branch:</p>
                                <router-link to="/edit/branch" class="grey--text">http://localhost:8080/#/edit/branch</router-link>
                            </v-col>
                            <v-col cols="12" sm="12" class="">
                                    <!-- <p class="h6 grey--text">Services</p> -->
                                    <v-btn class="mr-3" color="primary"><v-icon left>mdi-content-save</v-icon>Save</v-btn>
                                    <v-btn color="warning"><v-icon left>mdi-restore</v-icon>Reset</v-btn>
                                </v-col>
                            <v-btn color="info" text to="/manage/counters"><v-icon left>mdi-view-dashboard-edit</v-icon> Manage Branch Counters</v-btn>

                        </div>
                        <div class="col-md-6 mt-5">
                            <p class="h6 grey--text">Queue Ticket Printer Details</p>
                            <p class="body-2 grey--text">Printer Code:</p>
                            <p class="body-2 grey--text">Printer Name:</p>
                            <p class="body-2 grey--text">Status:</p>
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
import DataTables from '../tables/dataTables.vue';
import EditTable from '../tables/editTable.vue';

export default {
    data() {
        return {
            html:`FREE<br>No of Users: 3<br>No of Counters: 3<br>No of Services: 3<br>No of Queue Tickets Per Day: 10`,
            items: ["India", "America", "Europe", "Japan"]
        };
    },
    components: { DataTables, EditTable }
}
</script>

<style lang="scss" scoped></style>