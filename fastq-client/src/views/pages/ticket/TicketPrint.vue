<template>
    <v-app>
        <v-app-bar flat rounded app>
            <h2 class="" v-if="this.$store.state.Auth.user">{{ this.$store.state.Auth.user.company }}</h2>
            <v-spacer></v-spacer>
            <div class="d-flex">
                <v-chip class="ma-2" close color="cyan" label text-color="white">
                    <v-icon left>
                        mdi-account
                    </v-icon>
                    Admin
                </v-chip>
                <v-btn fab small color="#6A6A69" class="mx-2">
                    <v-icon color="white">mdi-cloud</v-icon>
                </v-btn>
            </div>
        </v-app-bar>
        <v-main>
            <v-row justify="center" v-if="this.$store.state.Auth.user">
                <v-dialog v-model="dialog" v-if="allconf.ds_conf.show_form" persistent max-width="600px">
                    <!-- <template v-slot:activator="{ on, attrs }">
                    <v-btn
                      color="primary"
                      dark
                      v-bind="attrs"
                      v-on="on"
                    >
                      Open Dialog
                    </v-btn>
                  </template> -->
                    <v-card>
                        <v-card-title>
                            <span class="text-h5">User Profile</span>
                        </v-card-title>
                        <v-card-text>
                            <v-container>
                                <v-row>
                                    <v-col cols="12" sm="6" md="4">
                                        <v-text-field label="Legal first name*" required></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="6" md="4">
                                        <v-text-field label="Legal last name*" hint="example of persistent helper text"
                                            persistent-hint required></v-text-field>
                                    </v-col>
                                    <v-col cols="12">
                                        <v-text-field label="Email*" required></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="6">
                                        <v-select :items="['0-17', '18-29', '30-54', '54+']" label="Age*"
                                            required></v-select>
                                    </v-col>
                                    <v-col cols="12" sm="6">
                                        <v-autocomplete
                                            :items="['Skiing', 'Ice hockey', 'Soccer', 'Basketball', 'Hockey', 'Reading', 'Writing', 'Coding', 'Basejump']"
                                            label="Interests" multiple></v-autocomplete>
                                    </v-col>
                                </v-row>
                            </v-container>
                            <small>*indicates required field</small>
                        </v-card-text>
                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn color="blue darken-1" text @click="dialog = false">
                                Close
                            </v-btn>
                            <v-btn color="blue darken-1" text @click="dialog = false">
                                Save
                            </v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
            </v-row>
            <div v-if="!this.$store.state.Auth.user" class="d-flex align-items-center justify-content-center">
                <h1>Company Not Found!! Contact Service Provider.</h1>
            </div>
            <div class="row fill-height align-items-center justify-content-center" v-if="this.$store.state.Auth.user">
                <div class="d-flex col-md-12 flex-column align-items-center">
                    <p class="text-h6">Please select from the Services below :</p>

                    <v-card v-for="n in services" :key="n.id" class="mb-3 col-md-4 text-center" @click="addTicket(n)" shaped
                        hover color="primary">
                        <h1 class="text-center white--text p-1">
                            {{ n.name }}
                        </h1>

                    </v-card>
                </div>
            </div>

        </v-main>

        <v-footer padless fixed bottom elevation="4">
            <v-col class="text-center" cols="12">
                <v-row :class="{ 'd-none': $vuetify.breakpoint.mdAndDown, 'align-center': true }">
                    <v-col cols="12" class="d-flex align-items-center">
                        <v-img class="ml-2" src="../../../assets/rsz_al_2.png" max-height="28" max-width="100"
                            contain></v-img>
                        <v-col class="d-flex flex-column p-0">
                            <p class="caption text-left p-0 mb-n1">Copyright © {{ new Date().getFullYear() }} FASTQ
                                Solutions limited. All Rights
                                Reserved
                            </p>
                            <p class="caption text-left p-0 mb-n1">
                                Build version : 1.0.409 - {{ m.format("DD MMM YY HH:MM A") }}</p>
                        </v-col>
                    </v-col>
                </v-row>
                <!-- <v-spacer></v-spacer> -->

            </v-col>
        </v-footer>

    </v-app>
</template>

<script>
import axios from 'axios';
import moment from 'moment';
import * as types from "../../../store/types"
export default {
    data() {
        return {
            services: [],
            lastNumber: null,
            dialog: false,
            allconf: null,
            m: moment()
        }
    },
    methods: {
        getAllServices() {
            let payload = {
                company_code: this.$store.state.Auth.user.company_code,
                branch_code: ""
            }
            axios.get(`/service/getAllservice/${this.$store.state.Auth.user.company_code}`).then(res => {
                console.log("res:::::", res);
                // this.$toast.success("branch added successfully.")
                res.data.message.forEach(element => {
                    let data = {}
                    data.name = element.Name
                    data.prefix = element.Prefix
                    // data.lastname = element.Lastname
                    // data.date = element.CreatedAt
                    // data.branch = element.CompanyCode
                    data.id = element.ID
                    data.numberStarts = element.NumberStarts
                    data.numberEnds = element.NumberEnds
                    data.code = element.Code
                    data.startTime = element.StartTime
                    data.endTime = element.EndTime
                    // data.Hide = element.Hide
                    // data.id = element.ID
                    this.services.push(data)
                });
                // this.services = res.data
                // this.reset()
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting all user!!!")
            })
        },
        addTicket(ticket) {
            this.dialog = true
            this.lastNumber++
            let ticketname = ticket.prefix + "-" + this.lastNumber
            let date1 = new Date(ticket.startTime)
            let date2 = new Date(ticket.endTime)
            // console.log("::::::::::::::::::::", this.lastNumber++);
            let ticketPayload = {
                service: ticket.name,
                ticket_status: 'CREATED',
                ticket_number: this.lastNumber.toString(),
                started_serving_at: moment(date1).format("YYYY-MM-DD HH:mm:ss"),
                end_serving_at: moment(date2).format("YYYY-MM-DD HH:mm:ss"),
                ticket_name: ticketname,
                company_code: this.$store.state.Auth.user.company_code,
                company_name: this.$store.state.Auth.user.company,
                updated_by: this.$store.state.Auth.user.id,
            }
            axios.post("/ticket/addticket", ticketPayload).then(res => {
                console.log("res:::::", res);
                this.$toast.success("ticket added successfully.")
                this.printTicket(ticketPayload)
                this.branch = {}
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while adding branch!!!")
            })

        },
        printTicket(ticketPayload) {
            ticketPayload.timeDate = moment(new Date()).format("DD MMM YY HH:MM A")
            this.waitingTickets().then(res => {
                ticketPayload.Waiting = res
                console.log("object", ticketPayload.Waiting);
                axios.post("http://localhost:4000/api/print", ticketPayload).then(res => {
                    console.log("res:::::", res);
                    this.$toast.success("ticket printed successfully.")
                    this.branch = {}
                }).catch(err => {
                    console.log(err.response);
                    this.$toast.error("error occured while printing ticket!!!")
                })
            })
        },
        waitingTickets() {
            return new Promise((resolve, reject) => {
                axios.get(`/report/allwaitingTickets/${this.$store.state.Auth.user.company_code}`).then(res => {
                    console.log("res:::::", res);
                    resolve(res.data)
                }).catch(err => {
                    console.log(err.response);
                    reject(err)
                    // this.$toast.error("error occured while printing ticket!!!")
                })

            })
        },
        lastticketNumber() {
            axios.get("/ticket/getLastNumber").then(res => {
                console.log("res:::::", res);
                this.lastNumber = res.data.number
                // this.$toast.success("branch added successfully.")
                // this.branch = {}
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting last ticket!!!")
            })

        },
        getAllConfigs() {
            axios.get(`/config/getallconfig/${this.$store.state.Auth.user.company_code}`).then(res => {
                console.log("res :::::", res);
                this.allconf = res.data
                // this.allconf.schedular_conf.forEach(element => {
                //     console.log("single schedular", element);
                //     this.dayWiseVideo[element.day].push(element.video_id)
                // });
                // console.log("schedular", this.dayWiseVideo);
                // this.$store.commit("SET_USER", res.data)
                // this.$store.commit(types.MUTATE_LOADER_OFF)
            }).catch(err => {
                console.log(err.response);
                // reject(err.response)
                // this.$store.commit(types.MUTATE_LOADER_OFF)
                this.$toast.error("error occured while getting connected clients!!!")
            })

        },
        getCompany() {
            // console.log("object", this.$route);
            return new Promise((resolve, reject) => {
                if (Object.keys(this.$route.query).includes("company")) {
                    this.id = this.$route.query.company
                    this.$store.commit(types.MUTATE_LOADER_ON);
                    axios.get(`/app/getCompany/${this.id}`).then(res => {
                        console.log("res :::::", res);
                        this.$store.commit("SET_USER", res.data)
                        this.$store.commit(types.MUTATE_LOADER_OFF)
                        resolve()
                    }).catch(err => {
                        console.log(err.response);
                        reject(err.response)
                        this.$store.commit(types.MUTATE_LOADER_OFF)
                        this.$toast.error("error occured while getting company!!!")
                    })
                }
            })
        },

    },
    mounted() {
        console.log("object", this.$route.query);


    },
    created() {
        this.getCompany().then(res => {
            this.getAllServices()
            this.lastticketNumber()
            this.getAllConfigs()
        })
    }
}
</script>

<style lang="scss" scoped>
#app {
    background-color: #eee;
}
</style>