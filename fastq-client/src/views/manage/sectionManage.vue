<template>
    <div class="container h-100">
        <div class="row h-100">
            <div class="row col-md-12 h-100">
                <v-card class="col-md-12 h-auto" flat>
                    <div class="row d-flex">
                        <v-col cols="12" sm="4">
                            <v-text-field value="" label="Branch Code" solo></v-text-field>
                        </v-col>
                        <v-col class="pl-0 col-md-4">
                            <v-btn large text color="primary" class="mt-1 text-md-button"> <v-icon left>mdi-magnify</v-icon>
                                Search
                            </v-btn>
                        </v-col>
                    </div>
                    <v-divider></v-divider>
                    <EditTable :data="tableData" :headers="headers" :name="'section'" />
                    <v-col class="pl-0 col-md-2">
                        <v-btn large text color="success" class="mt-1 text-md-button"> <v-icon left>mdi-export</v-icon>
                            Export Excel
                        </v-btn>
                    </v-col>
                    <div class="row">
                        <div class="col-md-12 mt-5">
                            <h5 class="h4 grey--text mt-3">Add a New Section</h5>
                            <div class="col-md-12 d-flex flex-wrap">
                                <div class="col-md-12 d-flex flex-wrap justify-content-between">
                                    <v-col cols="12" sm="5" class="m-0 p-0">
                                        <p class="h6 grey--text">Name</p>
                                        <v-text-field label="" class="m-0 p-0" placeholder="" outlined
                                            dense></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="5" class="m-0  p-0">
                                        <p class="h6 grey--text">Benchmark Waiting Time (mins)</p>
                                        <v-text-field label="" class="m-0 p-0" placeholder="" outlined
                                            dense></v-text-field>
                                            <p class="caption p-0 m-0 mt-n5">Note: This is the normal expected waiting time for this Section.</p>
                                    </v-col>
                                    <v-col cols="12" sm="5" class="m-0  p-0">
                                        <p class="h6 grey--text">Benchmark Processing Time(mins)</p>
                                        <v-text-field label="" class="m-0 p-0" placeholder="" outlined
                                            dense></v-text-field>
                                            <p class="caption p-0 m-0 mt-n5">Note: This is the normal expected processing time for this Section.</p>
                                    </v-col>
                                    <v-col cols="12" sm="5" class="m-0  p-0">
                                        <p class="h6 grey--text">Description</p>
                                        <v-text-field label="" class="m-0 p-0" placeholder="" outlined
                                            dense></v-text-field>
                                    </v-col>
                                </div>

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
import DataTables from '../tables/dataTables.vue';
import EditTable from '../tables/editTable.vue';

export default {
    data() {
        return {
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
    components: { DataTables, EditTable }
}
</script>

<style lang="scss" scoped></style>