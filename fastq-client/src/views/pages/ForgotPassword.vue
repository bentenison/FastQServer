<template>
    <v-app>
        <v-container fill-height>
            <v-layout row align-center justify-center>
                <v-card flat light max-width="400px" align-center>
                    <v-toolbar elevation="6" rounded color="primary" class="text-white" py-4>
                        <div class="d-flex row align-items-center justify-content-center w-100">
                            <h4 class="h4 pt-5">
                                FASTQ
                            </h4>
                            <p class="caption">Join queue anytime,anywhere</p>
                        </div>
                        <v-img :alt="platformName" contain height="48px" width="48px" position="center right"
                            src="../../assets/rsz_al_2.svg"></v-img>
                    </v-toolbar>
                    <v-divider></v-divider>
                    <v-card-text>
                        <p>Enter your email in order to change your password:</p>
                        <v-form>
                            <v-text-field outline label="Email" :rules="nameRules" type="text"
                                v-model="username"></v-text-field>
                        </v-form>
                    </v-card-text>
                    <v-divider></v-divider>
                    <v-card-actions :class="{ 'pa-3': $vuetify.breakpoint.smAndUp }">
                        <v-btn color="primary" class="" text :large="$vuetify.breakpoint.smAndUp">
                            <v-icon left color="primary">mdi-send</v-icon>
                            Send Link
                        </v-btn>
                    </v-card-actions>
                </v-card>
            </v-layout>
        </v-container>
        <v-footer padless fixed bottom elevation="4">
            <v-col class="text-center" cols="12">
                <v-row :class="{ 'd-none': $vuetify.breakpoint.mdAndDown, 'align-center': true }">
                    <v-col cols="12" class="d-flex align-items-center">
                        <v-img class="ml-2" src="../../assets/rsz_fastq_2.png" max-height="28" max-width="100"
                            contain></v-img>
                        <v-col class="d-flex flex-column p-0">
                            <p class="caption text-left p-0 mb-n1">Copyright © 2020 FASTQ Solutions limited. All Rights
                                Reserved
                            </p>
                            <p class="caption text-left p-0 mb-n1">Build version : 1.0.409 - 2023-03-28 01:50PM</p>
                        </v-col>
                    </v-col>
                </v-row>
                <!-- <v-spacer></v-spacer> -->

            </v-col>
        </v-footer>
    </v-app>
</template>

<script>
export default {
    data() {
        return {
            username: '',
            password: '',
            nameRules: [
                v => !!v || 'Name is required',
                v => v.length <= 10 || 'Name must be less than 10 characters',
            ],
            passwordRules: [
                v => !!v || 'Password is required',
                v => v.length <= 10 || 'Name must be less than 10 characters',
            ],
            emailRules: [
                v => !!v || 'E-mail is required',
                v => /.+@.+/.test(v) || 'E-mail must be valid',
            ],
        }
    },
    // authenticate() {
    //     this.$store.dispatch('AUTH_REQUEST', { loginId: this.username, password: this.password }).then(res => {
    //         // Redirect to next page after suucessfull login
    //         alert('Login : ' + res.isValid('MQLLogin'))
    //     }).catch(err => {
    //         alert(err)
    //         // Vue.$log.error(err)
    //     })
    // },
}
</script>

<style lang="scss" scoped></style>