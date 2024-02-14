<template>
  <v-app>
    <div class="container-fluid d-flex align-items-center flex-wrap h-100">
      <div class="col-md-6">
        <h1 class="grey--text">Sign up to FastQ</h1>
        <v-img src="../../assets/signup_art.svg" aspect-ratio="1.7" contain></v-img>

      </div>
      <v-card class="m-auto" min-width="500" flat>
        <v-card-title class="text-h6 font-weight-regular justify-space-between">
          <v-img alt="" contain height="48px" position="center left" src="../../assets/sec-icon.svg"></v-img>
          <span>{{ currentTitle }} </span>

          <v-avatar color="#AD1457" class="subheading white--text m-2" size="24" v-text="step"></v-avatar>
        </v-card-title>

        <v-window v-model="step">
          <v-window-item :value="1">
            <v-card-text>
              <v-form ref="signupForm">


                <v-text-field label="Email" :rules="emailRules" value="" placeholder="john@fastqsolutions.com"
                  v-model="email"></v-text-field>
                <span class="text-caption grey--text text--darken-1">
                  This is the email you will use to login to your FastQ account
                </span>
                <div class="row d-flex flex-wrap">
                  <v-text-field label="FirstName" :rules="nameRules" value="" placeholder="John" v-model="firstname"
                    class="mt-2 col-md-6"></v-text-field>
                  <v-text-field label="LastName" :rules="nameRules" value="" placeholder="Doe" v-model="lastname"
                    class="mt-2 col-md-6"></v-text-field>
                </div>
                <div class="row d-flex flex-wrap">
                  <v-text-field label="CompanyName" :rules="companyRules" value="" placeholder="fastq solution ltd"
                    v-model="company" class="mt-2 col-md-6"></v-text-field>
                  <v-text-field label="Timezone" :rules="nameRules" value="" placeholder="" v-model="timezone"
                    class="mt-2  col-md-6"></v-text-field>
                </div>
                <v-select label="Country" :rules="nameRules" :items="countries" value="" placeholder="India"
                  v-model="country" class="mt-2"></v-select>
              </v-form>
            </v-card-text>
          </v-window-item>

          <v-window-item :value="2">
            <v-card-text>
              <v-text-field label="Password" v-model="password" :rules="passwordRules" type="password"></v-text-field>
              <v-text-field :rules="[confirmPasswordRules, (password === confirmPassword) || 'Password must match']"
                v-model="confirmPassword" label="Confirm Password" type="password"></v-text-field>
              <span class="text-caption grey--text text--darken-1">
                Please enter a password for your account
              </span>
            </v-card-text>
          </v-window-item>

          <v-window-item :value="3">
            <div class="pa-4 text-center">
              <v-img class="mb-4" contain height="128" src="../../assets/sec-icon.svg"></v-img>
              <h3 class="text-h6 font-weight-medium mb-2">
                Welcome to FASTQ
              </h3>
              <span class="text-body-1 grey--text">Thanks for signing up!</span>
            </div>
          </v-window-item>
        </v-window>

        <v-divider></v-divider>

        <v-card-actions class="row d-flex flex-column">
          <div class="d-flex col-md-12 justify-content-between">
            <v-btn :disabled="step === 1" text @click="step--">
              Back
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn color="cyan darken-3" text @click.prevent="createAccount">
              {{ step === 3 ? "SignUp" : "Next" }}
            </v-btn>
          </div>

          <p class="grey--text text-center">Already have an account?<router-link to="/login"
              class="text-button">&nbsp;&nbsp;SIGNIN</router-link></p>
        </v-card-actions>
      </v-card>
    </div>
    <v-footer padless fixed elevation="4" app>
      <v-col class="text-center" cols="12">
        <v-row :class="{ 'd-none': $vuetify.breakpoint.mdAndDown, 'align-center': true }">
          <v-col cols="12" class="d-flex align-items-center">
            <v-img class="ml-2" src="../../assets/sec-icon.svg" max-height="28" max-width="100" contain></v-img>
            <v-col class="d-flex p-0 flex-column">
              <p class="caption text-left p-0 m-0">Copyright © 2020 FASTQ Solutions limited. All Rights Reserved
              </p>
              <p class="caption text-left p-0 m-0">Build version : 1.0.409 - 2023-03-28 01:50PM</p>
            </v-col>
          </v-col>
        </v-row>
        <!-- <v-spacer></v-spacer> -->
        <!-- {{ new Date().getFullYear() }} — <strong>FASTQ solutions limited</strong> -->
      </v-col>
    </v-footer>
  </v-app>
</template>

<script>
import axios from 'axios'

export default {
  data: () => ({
    step: 1,
    email: '',
    firstname: '',
    lastname: '',
    company: '',
    country: '',
    timezone: '',
    confirmPassword: '',
    password: '',
    countries: [],
    success:false,
    nameRules: [
      v => !!v || 'Name is required',
      v => v.length <= 10 || 'Name must be less than 10 characters',
    ],
    companyRules: [
      v => !!v || 'Company name is required',
      v => v.length <= 20 || 'Name must be less than 20 characters',
    ],
    passwordRules: [
      v => !!v || 'Password is required',
      v => v.length <= 10 || 'Password must be less than 10 characters',
      v => v.length >= 6 || 'Password must be more than 6 characters',
    ],
    confirmPasswordRules: [
      v => !!v || 'Confirm password is required',
      v => v.length <= 10 || 'Confirm password must be less than 10 characters',
      v => v.length >= 6 || 'Confirm password must be more than 6 characters',
    ],
    emailRules: [
      v => !!v || 'E-mail is required',
      v => /.+@.+/.test(v) || 'E-mail must be valid',
    ],
  }),

  computed: {
    currentTitle() {
      switch (this.step) {
        case 1: return 'Sign-up'
        case 2: return 'Create a password'
        default: return 'Account created'
      }
    },
  },

  methods: {
    createAccount() {
      if (this.step === 3) {
        if (this.password !== this.confirmPassword) {
          this.$toast.error("Password didn't match.")
          return
        }
        if (this.$refs.signupForm.validate()) {
          let payload = {}
          payload.email = this.email
          payload.firstname = this.firstname
          payload.lastname = this.lastname
          payload.company = this.company
          payload.country = this.country
          payload.timezone = this.timezone
          payload.password = this.password
          payload.confirmPassword = this.confirmPassword

          axios.post("/signup", payload).then(res => {
            console.log("res:::", res);
            this.$toast.success(res.data.message)
            this.email = ""
            this.firstname = ""
            this.lastname = ""
            this.company = ""
            this.country = ""
            this.timezone = ""
            this.password = ""
            this.confirmPassword = ""

           return
            // this.$refs.signupForm.reset()
          }).catch(err => {
            this.$toast.error(err.type)
          })
        } else {
          this.$toast.error("All fields are required.")
          return
        }
      }
      this.step++
    },
    getCountries() {
      axios.get("/get-countries").then(res => {

        res.data.data.countries.forEach(element => {
          this.countries.push(element.country_name)
        });
        this.$toast.success("countries fetched!!")
      }).catch(err => {
        console.log(err.response);
        this.$toast.error(err.type)
      })
    }
  },
  mounted() {
    this.getCountries()
  },
}
</script>