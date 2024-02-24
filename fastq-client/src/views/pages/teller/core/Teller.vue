<template>
    <v-app>
        <div class="container-fluid h-100">
            <div class="row h-100" v-if="!floating">
                <div class="col-md-3 h-80" style="margin-top: 70px;">
                    <v-card class="h-100" flat>
                        <v-card-title primary-title
                            v-if="$store.state.Auth.counterUser.Firstname && $store.state.Auth.activeCounter.CounterName && $store.state.Auth.counterUser.Lastname">
                            {{ $store.state.Auth.activeCounter.CounterName.toUpperCase() }}<br />
                            {{ $store.state.Auth.counterUser.Firstname.toUpperCase() }} {{
                                $store.state.Auth.counterUser.Lastname.toUpperCase() }}
                        </v-card-title>
                        <hr />
                        <div class="d-flex p-2">
                            <v-card class="w-100" elevation="6">
                                <h4 class="p-1">CURRENT:</h4>
                                <h1 class="font-weight-bold px-2" v-if="$store.state.Auth.activeTicket">{{
                                    $store.state.Auth.activeTicket.TicketName }}</h1>
                            </v-card>
                        </div>
                        <hr>
                        <div class="d-flex p-2">
                            <v-card class="w-100" flat>
                                <div class="row">
                                    <div class="col-md-12 d-flex justify-content-between">
                                        <p class="px-2">Processing Time :</p>
                                        <div class="px-2">
                                            <Stopwatch :resetWhenStart="true" :running="startTimer"
                                                v-if="$store.state.Auth.activeTicket"
                                                :starter="calculateTimestampDifference($store.state.Auth.activeTicket.StartedServingAt, convertDateToCustomFormat(convertDateToIndianTimestampFormat(new Date())))" />
                                        </div>
                                    </div>
                                    <div class="col-md-12 d-flex justify-content-between">
                                        <p class="px-2">Waiting Time :</p>
                                        <div class="px-2">
                                            <Stopwatch :resetWhenStart="true" :running="startTimer"
                                                v-if="$store.state.Auth.activeTicket"
                                                :starter="calculateTimestampDifference($store.state.Auth.activeTicket.CreatedAt, convertDateToCustomFormat(convertDateToIndianTimestampFormat(new Date())))" />
                                        </div>
                                    </div>
                                    <div class="col-md-12 d-flex justify-content-between">
                                        <p class="px-2">Transaction :</p>
                                        <div class="px-2">
                                            <Stopwatch :resetWhenStart="true" :running="startTimer"
                                                v-if="$store.state.Auth.activeTicket"
                                                :starter="calculateTimestampDifference($store.state.Auth.activeTicket.StartedServingAt, convertDateToCustomFormat(convertDateToIndianTimestampFormat(new Date()))) + calculateTimestampDifference($store.state.Auth.activeTicket.CreatedAt, convertDateToCustomFormat(convertDateToIndianTimestampFormat(new Date())))" />
                                        </div>
                                    </div>

                                </div>
                            </v-card>
                        </div>
                        <hr>
                        <div class="d-flex p-2">
                            <v-card class="w-100" elevation="6">
                                <h4 class="p-2">LAST CALLED :</h4>
                                <div class="d-flex justify-content-around" v-if="lastcalled !== 'NO_RESULT'">
                                    <div class="d-flex flex-column">
                                        <h1 class="font-weight-bold">{{ lastcalled.TicketName }}</h1>
                                        <h6 class="font-weight-bold">{{ lastcalled.Service }}</h6>
                                    </div>
                                    <div class="d-flex flex-column">
                                        <!-- <h1 class="font-weight-bold">1</h1> -->
                                        <h6 class="font-weight-bold">{{ $store.state.Auth.activeCounter.CounterName }}</h6>
                                    </div>
                                </div>
                            </v-card>
                        </div>
                        <hr>

                    </v-card>
                </div>
                <div class="col-md-9 h-100 mt-5">
                    <v-card class="h-100" flat>
                        <div class="d-flex row p-2" v-if="false">
                            <div class="col-md-4">

                                <v-select v-model="user" :items="users" item-text="name" item-value="" return-object
                                    label="Transfer To" solo dense></v-select>
                            </div>
                        </div>
                        <div class="container" style="margin-top: 60px;">
                            <div class="row text-white justify-content-between py-2">
                                <v-card
                                    class="col-md-2 pink accent-2 d-flex flex-column align-items-center text-white main-card"
                                    hover ripple @click.prevent="getnextticket" :class="{ 'disabled': disableNext }">
                                    <v-icon size="48px" color="white">
                                        mdi-arrow-right-bold-circle
                                    </v-icon>
                                    <h6>Call Next</h6>
                                </v-card>
                                <v-card class="col-md-2 primary d-flex flex-column align-items-center  text-white" hover
                                    ripple :class="{ 'disabled': true }">
                                    <v-icon size="48px" color="white">
                                        mdi-volume-high
                                    </v-icon>
                                    <h6>Call Again</h6>
                                </v-card>
                                <v-card class="col-md-2 warning accent-2 d-flex flex-column align-items-center text-white"
                                    hover ripple :class="{ 'disabled': true }">
                                    <v-icon size="48px" color="white">
                                        mdi-lock
                                    </v-icon>
                                    <h6>Standby</h6>
                                </v-card>
                                <v-card class="col-md-2 red accent-2 d-flex flex-column align-items-center  text-white"
                                    hover ripple @click.prevent="updateLastTicketStatus('NO_SHOW')">
                                    <v-icon size="48px" color="white">
                                        mdi-close-octagon
                                    </v-icon>
                                    <h6>No Show</h6>
                                </v-card>
                                <v-card class="col-md-2 green accent-4 d-flex flex-column align-items-center  text-white"
                                    hover ripple @click.prevent="updateLastTicketStatus('FINISHED')">
                                    <v-icon size="48px" color="white">
                                        mdi-check
                                    </v-icon>
                                    <h6>Finish</h6>
                                </v-card>
                            </div>
                        </div>
                        <div class="container mt-3 w-100">
                            <div class="row">
                                <div class="col-md-3">
                                    <div class="d-flex flex-column">
                                        <h5>Done</h5>
                                        <h1>{{ finished.length }}</h1>
                                    </div>
                                </div>
                                <!-- <div class="col-md-9 d-flex flex-wrap gap-2 justify-center align-center" v-if="false">
                                    <v-col cols="12" sm="12" class="m-0 p-0">
                                        <p class=" grey--text text-center">Customer Information</p>

                                    </v-col>
                                    <v-col cols="12" sm="4" class="m-0 p-0">
                                        
                                        <v-text-field label="customer name" class="m-0 p-0" hide-details="auto"
                                            placeholder="" solo dense></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="4" class="m-0 p-0">
                                        
                                        <v-text-field label="customer name" class="m-0 p-0" placeholder=""
                                            hide-details="auto" solo dense></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="4" class="m-0 p-0">
                                        
                                        <v-text-field label="customer name" class="m-0 p-0" hide-details="auto"
                                            placeholder="" solo dense></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="4" class="m-0 p-0">
                                        
                                        <v-text-field label="mo. number" class="m-0 p-0" hide-details="auto" placeholder=""
                                            solo dense></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="12" class="m-0 p-0 text-center">
                                        
                                        <v-btn color="primary text-white green accent-4" style="margin-top:2px">
                                            <v-icon left>
                                                mdi-content-save
                                            </v-icon>
                                            save
                                        </v-btn>
                                    </v-col>
                                </div> -->


                            </div>
                        </div>
                        <div class="col-md-12 d-flex justify-content-between h-50 mt-3">
                            <v-card class="col-md-4 h-100" flat>
                                <div class="row">
                                    <v-toolbar color="pink accent-2 darken-2" class="mb-2">
                                        <v-card-title primary-title class="text-white ">
                                            WAITING: {{ waiting.length }}
                                        </v-card-title>
                                    </v-toolbar>
                                    <div class="p-3">

                                        <!-- <v-card class="d-flex p-2 align-items-center mb-3 justify-content-between"
                                            elevation="5" outlined shaped>
                                            <div>1001</div>
                                            <Stopwatch :running="true" />
                                        </v-card> -->
                                        <v-card class="d-flex p-2 align-items-center justify-content-between mb-2"
                                            elevation="5" outlined shaped v-for="ticket in  waiting " :key="ticket.ID">
                                            <div>{{ ticket.TicketName }}</div>
                                            <!-- {{ ticket.CreatedAt }} {{ convertDateToIndianTimestampFormat(new Date()) }}  -->
                                            <!-- {{ convertDateToCustomFormat(convertDateToIndianTimestampFormat(new Date())) }} -->
                                            <!-- {{ calculateTimestampDifference(ticket.CreatedAt,
                                                convertDateToIndianTimestampFormat(new Date())) }} -->
                                            <!-- {{ calculateTimestampDifference(ticket.CreatedAt, ) }} -->
                                            <Stopwatch :resetWhenStart="true" :running="true"
                                                :starter="calculateTimestampDifference(ticket.CreatedAt, convertDateToCustomFormat(convertDateToIndianTimestampFormat(new Date())))" />
                                            <v-menu open-on-hover offset-y :close-on-content-click="false" top
                                                content-class="bg-white p-3">
                                                <template v-slot:activator="{ on, attrs }">
                                                    <v-icon v-bind="attrs" v-on="on">
                                                        mdi-information
                                                    </v-icon>
                                                </template>
                                                <table width="100%" cellpadding="3" class="text-caption bg-white">
                                                    <tbody class="">
                                                        <tr>
                                                            <th id="titleHeader" colspan="3">Queue Details</th>
                                                        </tr>
                                                        <tr class="p-3">
                                                            <td id="label1" width="30%">Information</td>
                                                            <td width="5%">:</td>
                                                            <td id="informationList" width="65%" align="right">john</td>
                                                        </tr>
                                                        <tr>
                                                            <td id="label1" width="30%">Mobile</td>
                                                            <td width="5%">:</td>
                                                            <td id="mobileNumber" width="65%" align="right">123456789</td>
                                                        </tr>
                                                        <tr>
                                                            <td id="label1" width="30%">Email</td>
                                                            <td width="5%">:</td>
                                                            <td id="email" width="65%" align="right">h@mail.com</td>
                                                        </tr>
                                                        <tr>
                                                            <td id="label1" width="30%">date</td>
                                                            <td width="5%">:</td>
                                                            <td id="dateStr" width="65%" align="right">June 25, 2023</td>
                                                        </tr>
                                                        <tr>
                                                            <td id="label1" width="30%">Transaction</td>
                                                            <td width="5%">:</td>
                                                            <td id="transactionName" width="65%" align="right">Inquiry</td>
                                                        </tr>

                                                        <tr>
                                                            <td id="label2" width="30%">Transaction Types</td>
                                                            <td width="5%">:</td>
                                                            <td width="65%" id="transactionTypesList" align="right"></td>
                                                        </tr>

                                                    </tbody>
                                                </table>
                                            </v-menu>
                                        </v-card>
                                    </div>

                                </div>

                            </v-card>
                            <v-card class="col-md-4 h-100" flat>
                                <div class="row">
                                    <v-toolbar color="red accent-2">
                                        <v-card-title primary-title class="text-white">
                                            NO SHOW : {{ noShow.length }}
                                        </v-card-title>
                                    </v-toolbar>
                                    <div class="p-3">

                                        <v-card class="mt-3 p-2 d-flex align-items-center justify-content-between"
                                            elevation="5" outlined shaped v-for=" noshow  in  noShow " :key="noshow.ID">
                                            <div>{{ noshow.TicketName }}</div>
                                            <Stopwatch :resetWhenStart="true" :running="false" />
                                        </v-card>
                                    </div>
                                </div>

                            </v-card>
                            <v-card class="col-md-4 h-100" flat>
                                <div class="row">
                                    <v-toolbar color="green accent-4">
                                        <v-card-title class="text-white ">
                                            TRANSACTIONS: {{ finished.length }}
                                        </v-card-title>
                                    </v-toolbar>
                                    <div class="p-3">
                                        <v-card class="mt-3 p-2 d-flex align-items-center justify-content-between"
                                            elevation="5" outlined shaped v-for=" finish  in  finished " :key="finish.ID">
                                            <div>{{ finish.TicketName }}</div>
                                            <div class="text-success">
                                                {{ calculateTimestampDifference(finish.CreatedAt,
                                                    finish.UpdatedAt) | secondsInMinutes }}
                                            </div>
                                        </v-card>
                                        <hr>
                                    </div>
                                </div>

                            </v-card>


                        </div>
                    </v-card>
                </div>

                <!-- <div class="row"> -->
                <!-- </div> -->
            </div>
            <div class="floating-window w-100" style="margin-top:70px;" v-if="floating">
                <div class="d-flex flex-wrap align-items-center justify-content-between row w-100">
                    <div class="col-6 d-flex flex-wrap">
                        <h5 class="p-1">Current:</h5>
                        <h2 class="font-weight-bold px-2" v-if="$store.state.Auth.activeTicket">{{
                            $store.state.Auth.activeTicket.TicketName || 'NO TICKET' }}
                        </h2>
                    </div>
                    <div class="col-6 d-flex flex-wrap">
                        <div class="">
                            <p class="px-2 my-0">Processing Time :</p>
                            <div class="px-2 my-0">
                                <Stopwatch :resetWhenStart="true" :running="startTimer"
                                    v-if="$store.state.Auth.activeTicket"
                                    :starter="calculateTimestampDifference($store.state.Auth.activeTicket.StartedServingAt, convertDateToCustomFormat(convertDateToIndianTimestampFormat(new Date())))" />
                            </div>

                        </div>
                        <div class="">
                            <p class="px-2 my-0">Waiting Time :</p>
                            <div class="px-2 my-0">
                                <Stopwatch :resetWhenStart="true" :running="startTimer"
                                    v-if="$store.state.Auth.activeTicket"
                                    :starter="calculateTimestampDifference($store.state.Auth.activeTicket.CreatedAt, convertDateToCustomFormat(convertDateToIndianTimestampFormat(new Date())))" />
                            </div>

                        </div>
                        <div class="">
                            <p class="px-2 my-0">Transaction :</p>
                            <div class="px-2 my-0">
                                <Stopwatch :resetWhenStart="true" :running="startTimer"
                                    v-if="$store.state.Auth.activeTicket"
                                    :starter="calculateTimestampDifference($store.state.Auth.activeTicket.StartedServingAt, convertDateToCustomFormat(convertDateToIndianTimestampFormat(new Date()))) + calculateTimestampDifference($store.state.Auth.activeTicket.CreatedAt, convertDateToCustomFormat(convertDateToIndianTimestampFormat(new Date())))" />
                            </div>

                        </div>
                    </div>
                    <div class="col-6 d-flex flex-wrap">
                        <h5 class="p-1">Waiting:</h5>
                        <h2 class="font-weight-bold px-2">{{ waiting.length }}
                        </h2>
                    </div>
                    <div class="col-6 d-flex">
                        <v-card-title primary-title
                            v-if="$store.state.Auth.counterUser && $store.state.Auth.counterUser.Firstname && $store.state.Auth.activeCounter.CounterName && $store.state.Auth.counterUser.Lastname">
                            {{ $store.state.Auth.activeCounter.CounterName.toUpperCase() }}<br />
                            {{ $store.state.Auth.counterUser.Firstname.toUpperCase() }} {{
                                $store.state.Auth.counterUser.Lastname.toUpperCase() }}
                        </v-card-title>
                    </div>
                    <div class="col-12 w-100 d-flex justify-content-center">
                        <v-card class="w-90 d-flex px-2" elevation="6">
                            <h5 class="p-2">LAST CALLED :</h5>
                            <div class="d-flex justify-content-around" v-if="lastcalled !== 'NO_RESULT'">
                                <div class="d-flex flex-column">
                                    <h1 class="font-weight-bold">{{ lastcalled.TicketName }}</h1>
                                    <h6 class="font-weight-bold">{{ lastcalled.Service }}</h6>
                                </div>
                            </div>
                        </v-card>
                    </div>
                    <!-- <v-card class="w-100" elevation="6">
                    </v-card> -->
                </div>
            </div>
        </div>
    </v-app>
</template>

<script>
import axios from 'axios';
import Stopwatch from './stopwatch.vue';
import moment from 'moment';

export default {
    components: { Stopwatch },
    filters: {
        secondsInMinutes: function (seconds) {
            return moment("2015-01-01")
                .startOf("day")
                .seconds(seconds)
                .format("HH:mm:ss");
        }
    },
    data() {
        return {
            floating: false,
            noShow: [],
            finished: [],
            waiting: [],
            lastcalled: null,
            present: false,
            users: [],
            user: null,
            counters: [],
            disableNext: false,
            socket: null,
            serverIp: null,
            startTimer: false,
            intervalId: null,
        }
    },
    methods: {
        convertDateToCustomFormat(date) {
            // Create a moment object from the JavaScript Date
            const formattedDate = moment(date).format('YYYY-MM-DDTHH:mm:ss[Z]');

            return formattedDate;
        },
        convertDateToIndianTimestampFormat(date) {
            // Adjust the date and time for IST (GMT+5:30)
            date.setUTCHours(date.getUTCHours() + 5);
            date.setUTCMinutes(date.getUTCMinutes() + 30);

            // Convert to the specified timestamp format
            const timestamp = date.toISOString().replace('Z', '+05:30');
            return timestamp;
        },
        calculateTimestampDifference(timestamp1, timestamp2) {
            const date1 = new Date(timestamp1);
            const date2 = new Date(timestamp2);
            // console.log(">>>>>>>>>>>>>>>>>>>>", date2);
            // console.log(">>>>>>>>>>>>>>>>>>>> created AT", date1);
            // Calculate the difference in milliseconds
            const differenceInMilliseconds = date2 - date1;

            // Convert milliseconds to seconds
            const differenceInSeconds = differenceInMilliseconds / 1000;

            return differenceInSeconds;
        },

        getnextticket() {
            this.getlastongoingticket().then(res => {
                console.log("last ongoing", res);
                if (res === "NO_RESULT") {

                } else {
                    this.updateJustTicketStatus('FINISHED', res.ID)
                }
                //get last ongoing ticket directly first if not available go for this API

            })
            let payload = {
                id: "",
                company_code: this.$store.state.Auth.user.company_code,
                branch_code: this.$store.state.Auth.user.company_code,
            }
            axios.post("/ticket/gettickettoprocess", payload).then(res => {
                // if (res.status)
                // console.log("res::::: next", res);
                if (res.data && res.data !== "NO_RESULT") {
                    res.data.CounterID = this.$store.state.Auth.activeCounter.ID
                    res.data.StartedServingAt = this.convertDateToCustomFormat(this.convertDateToIndianTimestampFormat(new Date()))
                    // this.$store.dispatch('SET_CUTRRENT', res.data).then(ok => {
                    this.$store.commit("SET_ACTIVE_TICKET", res.data)
                    this.startTimer = true
                    this.updatestartTime()
                    this.disableNext = true
                    this.updateLastTicketStatus("STARTED")
                    this.sendMessage(res.data, "next")
                    this.reUpdate()
                    // })
                }

                // this.noShow = res.data
                // this.$toast.success("branch added successfully.")
                // this.branch = {}
            }).catch(err => {
                console.log(err);
                this.$toast.error("error occured while getting ticket!!!")
            })

        },
        getlastongoingticket() {
            return new Promise((resolve, reject) => {
                axios.get(`/ticket/getlaststarted/${this.$store.state.Auth.counterUser.ID}`).then(res => {
                    console.log("res::::: last ongoing", res);
                    if (res.data && res.data !== "NO_RESULT") {
                        this.present = true
                        this.startTimer = true
                        // this.$store.commit("SET_ACTIVE_TICKET", res.data)
                        resolve(res.data)
                    }
                    resolve(res.data)
                    // this.$store.commit("SET_ACTIVE_TICKET", res.data)
                    // this.updatestartTime()
                    // this.updateLastTicketStatus("STARTED")
                    // this.noShow = res.data
                    // this.$toast.success("branch added successfully.")
                    // this.branch = {}
                }).catch(err => {
                    this.present = false
                    console.log("error getting error", err.response.data);
                    this.$toast.error("error occured while getting ticket!!!")
                    reject(err)
                })

            })
        },
        updateLastTicketStatus(status) {
            let payload = {
                ticket_status: status,
                // updated_at: "",
                updated_by: this.$store.state.Auth.counterUser.ID,
                id: this.$store.state.Auth.activeTicket.ID

            }
            axios.post("/ticket/updateticketstatus", payload).then(res => {
                // console.log("res:::::", res);
                if (status === "NO_SHOW" || status === "FINISHED") {
                    this.disableNext = false
                } else {
                    this.disableNext = true
                }
                console.log("Status", status);
                if ((status === "NO_SHOW" || status === "FINISHED") && this.$store.state.Auth.activeTicket && this.$store.state.Auth.activeTicket !== "NO_RESULT") {
                    this.sendMessage(this.$store.state.Auth.activeTicket, "finish")
                    this.startTimer = false
                    this.reUpdate()
                    this.updateEndTime()
                    this.$store.commit("SET_ACTIVE_TICKET", null)
                }
                // this.noShow = res.data
                // this.$toast.success("branch added successfully.")
                // this.branch = {}
            }).catch(err => {
                console.log("error iin  ticket status", err.response);
                this.$toast.error("error occured while updating last ticket!!!")
            })

        },
        updateJustTicketStatus(status, id) {
            if (this.$store.state.Auth.activeTicket) {
                let payload = {
                    ticket_status: status,
                    // updated_at: "",
                    updated_by: this.$store.state.Auth.counterUser.ID,
                    id: id,

                }
                axios.post("/ticket/updateticketstatus", payload).then(res => {
                    // this.$store.commit("SET_ACTIVE_TICKET", null)
                    this.reUpdate()
                    this.updateEndTime()
                }).catch(err => {
                    console.log("error iin  ticket status", err.response);
                    this.$toast.error("error occured while updating last ticket!!!")
                })

            }

        },
        updatestartTime() {
            let payload = {
                // started_serving_at: "",
                // updated_at: "",
                updated_by: this.$store.state.Auth.counterUser.ID,
                id: this.$store.state.Auth.activeTicket.ID,
                counterId: this.$store.state.Auth.activeCounter.ID
            }
            axios.post("/ticket/updatetickettime", payload).then(res => {
                // console.log("res:::::", res);
                // this.noShow = res.data
                // this.$toast.success("branch added successfully.")
                // this.branch = {}
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting noshow ticket!!!")
            })

        },
        updateEndTime() {
            let payload = {
                // end_serving_at: "",
                // updated_at: "",
                updated_by: this.$store.state.Auth.counterUser.ID,
                id: this.$store.state.Auth.activeTicket.ID,
            }
            axios.post("/ticket/updateendtime", payload).then(res => {
                // console.log("res:::::", res);
                // this.noShow = res.data
                // this.$toast.success("branch added successfully.")
                // this.branch = {}
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting noshow ticket!!!")
            })

        },
        getnoshow() {
            axios.get(`/ticket/getnoshow/${this.$store.state.Auth.activeCounter.ID}`).then(res => {
                // console.log("res:::::", res);
                if (!res.data) {
                    this.noShow = []
                    return
                }
                this.noShow = res.data
                // this.$toast.success("branch added successfully.")
                // this.branch = {}
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting noshow ticket!!!")
            })
        },
        getwaiting() {
            axios.get("/ticket/getwaiting").then(res => {
                // console.log("res:::::", res);
                if (!res.data) {
                    this.waiting = []
                    return
                }
                this.waiting = res.data
                // this.$toast.success("branch added successfully.")
                // this.branch = {}
            }).catch(err => {
                console.log("error response in this", err.response);
                this.$toast.error("error occured while getting waiting ticket!!!")
            })
        },
        getcompleted() {
            axios.get(`/ticket/getfinished/${this.$store.state.Auth.activeCounter.ID}`).then(res => {
                // console.log("res:::::", res);
                if (!res.data) {
                    this.finished = []
                    return
                }
                this.finished = res.data
                // this.$toast.success("branch added successfully.")
                // this.branch = {}
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting completed ticket!!!")
            })
        },
        getlastcalled() {
            axios.get("/ticket/getlastcalled").then(res => {
                // console.log("res:::::", res);
                this.lastcalled = res.data
                // this.$toast.success("branch added successfully.")
                // this.branch = {}
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting last ticket!!!")
            })
        },
        getServerIP() {
            return new Promise((resolve, reject) => {
                if (window.ipc) {
                    window.ipc.send("SERVER_IP", {})
                    window.ipc.on("SERVER_IP_REPLY", res => {
                        console.log("from electron process", res);
                        // this.$store.commit(types.MUTATE_LOADER_OFF)
                        this.serverIp = res.IP
                        resolve(res)
                    })
                }
                // axios.get("http://localhost:3000/api/getServerIP").then(res => {
                //     console.log("res serverip:::::", res);
                //     this.serverIp = res.data.IP
                //     resolve(res.data)
                //     // this.$toast.success("branch added successfully.")
                //     // this.branch = {}
                // }).catch(err => {
                //     reject(err)
                //     console.log(err.response);
                //     this.$toast.error("error occured while getting last ticket!!!")
                // })

            })
        },
        getAllUser() {
            let payload = {
                company_code: this.$store.state.Auth.user.company_code,
                branch_code: ""
            }
            axios.post(`/user/getAlluser`, payload).then(res => {
                // console.log("res:::::", res);
                // this.$toast.success("branch added successfully.")
                res.data.message.forEach(element => {
                    let data = {}
                    data.email = element.Email
                    data.name = element.Firstname + " " + element.Lastname
                    data.branch = element.CompanyCode
                    data.id = element.ID
                    // data.Hide = element.Hide
                    // data.id = element.ID
                    this.users.push(data)
                });
                // this.users = res.data
                // this.reset()
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting all user!!!")
            })
        },
        getAllCounters() {
            // get those counters which are running here (active counters)
            axios.get(`/counter/getAllcounter/${this.$store.getters.getUser.company_code}`).then(res => {
                // console.log(">>>>>>>>>>>>>>>>>", res);
                res.data.message.forEach(element => {
                    let data = {}
                    data.counterNumber = element.CounterNumber
                    data.counterName = element.CounterName
                    data.branchCode = element.BranchCode
                    data.branchName = element.BranchName
                    data.date = element.CreatedAt
                    data.id = element.ID
                    this.counters.push(data)
                });

                this.$toast.success("counters fetched successfully.")
                // this.connection.send("Hello World !!!")
            }).catch(err => {
                console.log(err.response);
                this.$toast.error("error occured while getting counters!!!")
            })
        },
        connectToServer() {
            // this.connected = !this.connected;

            // // Handle incoming messages
            // this.socket.onmessage = (event) => {
            //     const response = event.data;
            //     console.log('Received message from server:', response);
            //     // Process the response as needed
            // };

            // Handle connection close
            this.socket.onclose = (event) => {
                console.log('WebSocket connection closed:', event);
                // You may want to handle reconnection logic here if needed
            };
            this.socket.onopen = (event) => {
                // this.sendMessage()
                console.log('WebSocket connection opened:', event);
                // You may want to handle reconnection logic here if needed
            };
            this.socket.onerror = (event) => {
                console.log('WebSocket connection error:', event);
                // You may want to handle reconnection logic here if needed
            };
        },
        sendMessage(d, action) {
            // Send a message to the server
            if (this.socket.readyState === WebSocket.OPEN) {
                // const data = 'Hello, Server!'; // Your message data
                let payload = {
                    ip: "",
                    ticket_payload: d,
                    action: action
                }
                this.socket.send(JSON.stringify(payload));
                console.log('Sent message to server:', d);
            } else {
                this.socket = new WebSocket(`ws://${res.IP}:8090/ws`);
                this.connectToServer()
                this.sendMessage(d, action)
            }

        },
        reUpdate() {
            this.getwaiting()
            this.getcompleted()
            this.getlastcalled()
            this.getnoshow()
        },
        activatefloat() {
            console.log("float mode activated");
            if (window.ipc) {
                window.ipc.send("FLOAT_MODE", {})
                this.floating = true
                // window.ipc.on("SERVER_IP_REPLY", res => {
                //     console.log("from electron process", res);
                //     // this.$store.commit(types.MUTATE_LOADER_OFF)
                //     this.serverIp = res.IP
                //     resolve(res)
                // })
            }
        },
        deactivatefloat() {
            console.log("float mode deactivated");
            if (window.ipc) {
                window.ipc.send("FLOAT_MODE_OFF", {})
                this.floating = false
                // window.ipc.on("SERVER_IP_REPLY", res => {
                //     console.log("from electron process", res);
                //     // this.$store.commit(types.MUTATE_LOADER_OFF)
                //     this.serverIp = res.IP
                //     resolve(res)
                // })
            }
        }

    },
    beforeDestroy() {
        // Disconnect from the server when the component is destroyed
        if (this.socket) {
            this.socket.send("disconnect")
            this.socket.close();
            // console.log('Disconnected from WebSocket server');
        }
        if (this.intervalId) {
            clearInterval(this.intervalId)
        }

        this.$eventBus.$off('float-mode')
        this.$eventBus.$off('float-mode-off')

    },
    mounted() {
        this.getlastongoingticket().then(res => {
            this.reUpdate()
            this.getAllUser()
            this.getAllCounters()
        })
    },
    created() {
        this.$eventBus.$on('float-mode', this.activatefloat)
        this.$eventBus.$on('float-mode-off', this.deactivatefloat)
        this.getServerIP().then(res => {
            // console.log("server ip address", res);
            this.socket = new WebSocket(`ws://${res.IP}:8090/ws`);
            this.connectToServer()
        })
        this.intervalId = setInterval(this.getwaiting, 2000);
        if (window.ipc) {
            // window.ipc.send("SERVER_IP", {})
            window.ipc.on("next-ticket", res => {
                console.log("from electron process", res);
                // alert("next")
                this.getnextticket()
                // this.$store.commit(types.MUTATE_LOADER_OFF)
                // this.serverIp = res.IP
                // resolve(res)
            })
            window.ipc.on("finish-ticket", res => {
                console.log("from electron process", res);
                this.updateLastTicketStatus('FINISHED')
                // this.$store.commit(types.MUTATE_LOADER_OFF)
                // this.serverIp = res.IP
                // resolve(res)
            })
            window.ipc.on("no-show-ticket", res => {
                console.log("from electron process", res);
                this.updateLastTicketStatus('NO_SHOW')
                // this.$store.commit(types.MUTATE_LOADER_OFF)
                // this.serverIp = res.IP
                // resolve(res)
            })
        }
        if (!this.floating) {
            window.ipc.send("FLOAT_MODE_OFF", {})
            this.$eventBus.$emit('float-mode-off')
            // window.reload()
            this.floating = false
        }
    }
}
</script>

<style lang="scss" scoped>
.sm-card-heading {
    background: #00E676;
    border-radius: 6px 6px 0px 0px;
    // border: 1px solid red;
}

.v-card__title {
    font-size: 1rem;
    font-family: "Ubuntu" !important;
    // border: 1px solid rgb(17, 16, 16);
}


.timer-wait {
    background-color: red;
}

.disabled {
    pointer-events: none;
    /* Disable click events */
    opacity: 0.5;
    /* Optionally, reduce opacity to visually indicate that it's disabled */
}
</style>