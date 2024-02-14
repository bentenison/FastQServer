<template>
    <div class="container h-100">
        <div class="row h-100">
            <div class="row col-md-12 h-100">
                <v-card class="col-md-12 h-auto" flat>
                    <v-divider></v-divider>
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
                                    <v-text-field label="" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="4" class="">
                                    <p class="h6 grey--text">Counter Name:</p>
                                    <v-text-field label="" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
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
                                <v-btn class="mr-3" color="primary"><v-icon left>mdi-content-save</v-icon>Save</v-btn>
                                <v-btn color="warning"><v-icon left>mdi-restore</v-icon>Reset</v-btn>
                            </v-col>

                        </div>
                    </div>
                    <!-- <v-btn color="primary" class=""><v-icon left>mdi-content-save-all</v-icon>Save config</v-btn> -->
                </v-card>
            </div>
        </div>
    </div>
</template>
    
<script setup lang='ts'>

</script>
    
<style></style>