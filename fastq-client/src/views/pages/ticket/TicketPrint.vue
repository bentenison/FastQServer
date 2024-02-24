<template>
    <v-app>
        <v-app-bar flat rounded app>
            <!-- `http://${ipAddress}:8080/server/uploaded/logo.png -->
            <v-img class="mx-2"
                :src="ipAddress ? `https://${ipAddress}:443/uploaded/logo.png` : `http://localhost:8080/server/uploaded/logo.png`"
                max-height="45" max-width="100" contain></v-img>
            <h2 class="" v-if="this.$store.state.Auth.user">{{ this.$store.state.Auth.user.company.toUpperCase() }}</h2>
            <v-spacer></v-spacer>
            <div class="d-flex">
                <!-- <v-chip class="ma-2" close color="cyan" label text-color="white">
                    <v-icon left>
                        mdi-account
                    </v-icon>
                    Admin
                </v-chip> -->
                <v-btn fab small color="#6A6A69" class="mx-2">
                    <v-icon color="white">mdi-cloud</v-icon>
                </v-btn>
                <v-btn fab small color="#6A6A69" class="mx-2" @click.prevent="connectBluetoothPrinter">
                    <v-icon color="white">mdi-bluetooth</v-icon>
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
                    <!-- <p class="text-h6">Please select from the services below :</p> -->
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
import EscPosEncoder from '../../../../node_modules/esc-pos-encoder/dist/esc-pos-encoder.esm.js';
export default {
    data() {
        return {
            printCharacteristic: null,
            device: null,
            services: [],
            lastNumber: null,
            dialog: false,
            allconf: null,
            ipAddress: null,
            m: moment()
        }
    },
    methods: {
        printwithConnected(data) {
            const self = this
            self.device.gatt
                .connect()
                .then(server =>
                    server.getPrimaryService('000018f0-0000-1000-8000-00805f9b34fb')
                )
                .then(service =>
                    service.getCharacteristic('00002af1-0000-1000-8000-00805f9b34fb')
                )
                .then(characteristic => {
                    console.log('characteristic', characteristic)
                    self.printCharacteristic = characteristic
                    let encoder = new EscPosEncoder({
                        imageMode: 'raster'
                    });

                    // let result = encoder
                    //     .initialize()
                    //     .text('The quick brown fox jumps over the lazy dog')
                    //     .newline()
                    //     .qrcode('https://nielsleenheer.com')
                    //     .encode();

                    console.log("router", window.location);
                    const url = new URL(window.location.origin);

                    // Get the IP address from the URL
                    self.ipAddress = url.hostname;
                    let img = new Image();
                    if (self.ipAddress === "localhost") {
                        return
                    }
                    img.src = `https://${self.ipAddress}:443/uploaded/logo.png`;
                    // img.src = `http://localhost:8080/server/uploaded/logo.png`;

                    img.onload = async function () {
                        let result = encoder.align('center')
                            .image(img, 160, 160, 'atkinson').newline()
                            .encode()
                        console.log("hello", result);
                        const dataSize = result.length;
                        let startIndex = 0;
                        const maxChunkSize = characteristic.properties.writeWithoutResponse ? 512 : 20;
                        // Send chunks of data until the entire image is sent

                        // this.loop(0, result, device)
                        // characteristic.writeValue(result)
                        // 
                        self.printlogo(startIndex, result, maxChunkSize).then(async () => {

                            // let data = {
                            //     service: 'Parts Order',
                            //     ticket_status: 'CREATED',
                            //     ticket_number: '15',
                            //     started_serving_at: '2017-01-01 10:31:00',
                            //     end_serving_at: '2017-01-01 10:31:00',
                            //     ticket_name: 'A-15',
                            //     company_code: 'nDmmyD0',
                            //     company_name: 'fastq solutions ltd',
                            //     updated_by: '73ac3366-c04d-11ee-a5d5-ce47405e851e',
                            //     timeDate: '18 Feb 24 23:02 PM',
                            //     Waiting: '15'
                            // }
                            let text = encoder.align('center')
                                .width(3)
                                .height(3)
                                .line(data.ticket_name).bold().newline()
                                .width(2)
                                .height(2)
                                .line(data.service).bold().newline()
                                .width(2)
                                .height(2)
                                .line(`Waiting Count : ${data.Waiting}`)
                                .width(1)
                                .height(1)
                                .line(`Created At : ${data.timeDate}`)
                                .line("").newline()
                                .line("").newline()
                                .encode()
                            self.printCharacteristic.writeValue(text);
                            // self.device.gatt
                            //     .disconnect()
                        })
                    }
                    // self.sendTextData(device)
                })
                .catch(error => {
                    console.log("error", error);
                    // this.handleError(error, this.device)
                })
        },
        printlogo(startIndex, result, maxChunkSize) {
            return new Promise(async (resolve, reject) => {
                let dataSize = result.length
                while (startIndex < dataSize) {
                    const endIndex = Math.min(startIndex + maxChunkSize, dataSize);
                    const chunk = result.slice(startIndex, endIndex);
                    // Send the current chunk to the printer
                    await this.printCharacteristic.writeValue(chunk);
                    startIndex = endIndex;
                }
                resolve()
            })
        },
        connectBluetoothPrinter() {
            navigator.bluetooth
                .requestDevice({
                    filters: [
                        {
                            name: 'MPT-III',
                            services: ['000018f0-0000-1000-8000-00805f9b34fb']
                        }
                    ]
                },
                    {
                        optionalServices: ['00002af1-0000-1000-8000-00805f9b34fb']
                    })
                .then(device => {
                    console.log('device', device)
                    // window.alert(device)
                    this.device = device
                })
                .catch(err => {
                    console.log(err);
                })
        },
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
                // axios.post("http://localhost:4000/api/print", ticketPayload).then(res => {
                //     console.log("res:::::", res);
                //     this.$toast.success("ticket printed successfully.")
                //     this.branch = {}
                // }).catch(err => {
                //     console.log(err.response);
                //     this.$toast.error("error occured while printing ticket!!!")
                // })
                this.printwithConnected(ticketPayload)
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