<template>
    <div class="container h-100">
        <div class="row h-100">
            <div class="row col-md-12 h-100">
                <v-card class="col-md-7 h-auto" flat>
                    <div class="d-flex flex-column">
                        <p class="h6 grey--text">Company Name</p>
                        <p class="h4 grey--text">{{ $store.state.Auth.user.company }}</p>
                    </div>
                    <div class="d-flex flex-column">
                        <p class="h6 grey--text">Contact</p>
                        <p class="h4 grey--text">{{ $store.state.Auth.user.email }}</p>
                    </div>
                    <div class="d-flex flex-column">
                        <p class="h6 grey--text">Company Code</p>
                        <p class="h4 grey--text">{{ $store.state.Auth.user.company_code }}</p>
                    </div>
                    <div class="d-flex flex-column">
                        <p class="h6 grey--text">Locale Code</p>
                        <v-text-field name="localCode" label="locale" id="id" class="w-80"></v-text-field>
                    </div>
                    <div class="d-flex flex-column">
                        <p class="h6 grey--text">Country:</p>
                        <!-- <v-select :items="items" v-model="$store.state.Auth.user.country" label="country"></v-select> -->
                        <p class="h4 grey--text">{{ $store.state.Auth.user.country }}</p>
                    </div>
                    <div class="d-flex flex-column">
                        <p class="h6 grey--text">Time Zone:</p>
                        <!-- <v-select :items="items" v-model="$store.state.Auth.user.timezone" label="timezone"></v-select> -->
                        <p class="h4 grey--text">{{ $store.state.Auth.user.timezone }}</p>
                    </div>

                    <v-btn color="primary" class=""><v-icon left>mdi-content-save-all</v-icon>Save config</v-btn>
                </v-card>
                <div class="col-md-5 p-3">
                    <h5 class="h6 grey--text">Company Logo</h5>
                    <v-file-input label="Select a file to upload" filled prepend-icon="mdi-image-edit"></v-file-input>
                    <p class="caption">Maximum file size is 1 MB , only jpg/png files allowed</p>
                    <v-btn color="primary" elevation="2" outlined :loading="false"><v-icon left>
                            mdi-upload
                        </v-icon>Upload</v-btn>
                </div>
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

export default {
    data() {
        return {
            panel: "",
            readonly: false,
            items: ["India", "America", "Europe", "Japan"]
        };
    },
    components: { DataTables }
}
</script>

<style lang="scss" scoped></style>