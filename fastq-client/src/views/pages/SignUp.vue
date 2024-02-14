<template>
    <v-app>
        <v-container fill-height>
            <v-layout row align-center justify-center>
                <v-card elevation="4" light width="450px" align-center>
                    <v-toolbar elevation="4" rounded color="primary" py-4>
                        <v-layout align-center justify-space-between>
                            <v-flex>
                                <!-- <v-flex> -->

                                <h3 pa-2>
                                    FASTQ
                                </h3>
                                <p>Join queue anytime,anywhere</p>
                                <!-- </v-flex> -->
                            </v-flex>
                            <v-flex>
                                <v-img :alt="platformName" class="ml-3" contain height="48px" position="center right"
                                    src="../../assets/logo.png"></v-img>
                            </v-flex>
                        </v-layout>
                    </v-toolbar>
                    <v-divider></v-divider>
                    <v-card-text>
                        <p>Sign up with us to start queue anytime, anywhere :</p>
                        <v-form>
                            <v-text-field outline label="Username" type="text" v-model="username"></v-text-field>
                            <v-text-field outline label="Fullname" type="text" v-model="fullname"></v-text-field>
                            <v-text-field outline hide-details label="Password" type="password"
                                v-model="password"></v-text-field>
                        </v-form>
                    </v-card-text>
                    <v-divider></v-divider>
                    <v-card-actions :class="{ 'pa-3': $vuetify.breakpoint.smAndUp }">
                        <v-btn color="info" text>
                           already have an account
                        </v-btn>
                        <v-spacer></v-spacer>
                        <v-btn color="info" :large="$vuetify.breakpoint.smAndUp">
                            <v-icon left>mdi-arrow-right-bold-circle-outline</v-icon>
                            SignUp
                        </v-btn>
                    </v-card-actions>
                </v-card>
            </v-layout>
        </v-container>
    </v-app>
</template>

<script>
export default {

}
</script>

<style lang="sass" scoped>

</style>