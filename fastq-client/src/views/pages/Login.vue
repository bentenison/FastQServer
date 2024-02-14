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


                        <!-- <v-flex> -->
                        <v-img alt="platformName" contain height="48px" width="48px" position="center right"
                            src="../../assets/rsz_al_2.svg"></v-img>
                        <!-- </v-flex> -->
                        <!-- </v-layout> -->
                    </v-toolbar>
                    <v-divider></v-divider>
                    <v-card-text>
                        <p>Sign in with your username and password:</p>
                        <v-form ref="form" v-model="valid" lazy-validation>
                            <v-text-field outline label="Email" :rules="emailRules" type="text"
                                v-model="email"></v-text-field>
                            <v-text-field outline label="Password" :rules="passwordRules" type="password"
                                v-model="password"></v-text-field>
                        </v-form>
                    </v-card-text>
                    <v-divider></v-divider>
                    <v-card-actions :class="{ 'pa-3': $vuetify.breakpoint.smAndUp }">
                        <v-btn color="primary" text to="/forgot-password">
                            Forgot password?
                        </v-btn>
                        <v-spacer></v-spacer>
                        <v-btn color="primary" class="" text :large="$vuetify.breakpoint.smAndUp" @click="authenticate">
                            <v-icon left color="primary">mdi-lock</v-icon>
                            Login
                        </v-btn>
                    </v-card-actions>
                    <p class="grey--text text-center">Dont have an account?<router-link to="/signup"
                            class="text-button">&nbsp;&nbsp;SIGNUP</router-link></p>
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
            valid: true,
            email: '',
            password: '',
            nameRules: [
                v => !!v || 'Name is required',
                v => v.length <= 10 || 'Name must be less than 10 characters',
            ],
            passwordRules: [
                v => !!v || 'Password is required',
                // v => v.length <= 10 || 'Name must be less than 10 characters',
            ],
            emailRules: [
                v => !!v || 'E-mail is required',
                v => /.+@.+/.test(v) || 'E-mail must be valid',
            ],
        }
    },
    methods: {
        authenticate() {
            if (this.$refs.form.validate()) {
                this.$store.dispatch('AUTH_LOGIN', { email: this.email, password: this.password, role: 1 }).then(res => {
                    console.log("Results:::", res);
                    this.$toast.success('Logged In!!')
                    this.$router.push("/dashboard")
                    this.reset()
                }).catch(err => {
                    console.log(err);
                    this.$toast.error(err.response.data.error.message)
                })
            }
        },
        reset() {
            this.$refs.form.reset()
        },
    },
}
</script>

<style lang="scss" scoped></style>