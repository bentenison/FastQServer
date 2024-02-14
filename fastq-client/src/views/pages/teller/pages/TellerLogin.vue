<template>
    <v-app id="lock">
        <div class="container h-100 mt-5" :key="activated">
            <div class="row d-flex flex-row align-items-center justify-content-center mt-5" v-if="activated">
                <div class="col-md-5 py-5">
                    <v-card elevation="10">
                        <v-card-title class="primary text-white" v-if="this.activeCounter.CounterName">
                            Lock Screen For {{ this.activeCounter.CounterName.toUpperCase() }}
                        </v-card-title>
                        <v-card-text class="row p-4">
                            <v-text-field class="col-md-6" label="Email" type="text" v-model="counterUser.email"
                                :rules="rules"></v-text-field>
                            <v-text-field class="col-md-6" label="Password" type="password" v-model="counterUser.password"
                                :rules="rules"></v-text-field>
                        </v-card-text>
                        <v-card-actions>
                            <v-btn text color="primary text-white" @click="counterlogin">
                                <v-icon left>
                                    mdi-lock-open-variant
                                </v-icon>
                                Unlock
                            </v-btn>
                        </v-card-actions>
                    </v-card>
                </div>
            </div>
            <div class="row d-flex flex-row align-items-center justify-content-center mt-5" v-if="!activated">
                <div class="row col-md-8 d-flex align-items-center justify-content-center">
                    <v-card class="col-md-12 h-100" flat>
                        <v-divider></v-divider>
                        <div class="row">
                            <div class="col-md-12 mt-5">
                                <h5 class="h4 grey--text text-center">Add a New Counter</h5>
                                <div class="row d-flex flex-wrap align-items-center justify-content-center">
                                    <v-col cols="12" sm="5" class="">
                                        <p class="h6 grey--text">Counter Number:</p>
                                        <v-select v-model="counter" :items="tableData" item-text="counterName" item-value=""
                                            return-object label="Select Counter" solo></v-select>
                                    </v-col>
                                    <!-- <v-col cols="12" sm="5" class="">
                                        <p class="h6 grey--text">License Key:</p>
                                        <v-text-field label="" class="m-0 p-0" placeholder="" outlined dense></v-text-field>
                                    </v-col> -->
                                </div>
                                <v-col cols="12" sm="12" class="d-flex flex-wrap align-items-center justify-content-center">
                                    <!-- <p class="h6 grey--text">Services</p> -->
                                    <v-btn class="mr-3" color="primary" @click="activateCounter"><v-icon
                                            left>mdi-content-save</v-icon>Save</v-btn>
                                    <v-btn color="warning"><v-icon left>mdi-restore</v-icon>Reset</v-btn>
                                </v-col>

                            </div>
                        </div>
                        <!-- <v-btn color="primary" class=""><v-icon left>mdi-content-save-all</v-icon>Save config</v-btn> -->
                    </v-card>
                </div>
            </div>
        </div>
    </v-app>
</template>

<script>
import axios from 'axios';
import * as types from "../../../../store/types"
import { localinstance } from "../../../../instance"
// const { ipcRenderer } = window.require('electron');
export default {
    data() {
        return {
            username: null,
            password: null,
            clientDetails: null,
            activated: false,
            id: null,
            counter: null,
            tableData: [],
            activeCounter: null,
            counterUser: {},
            user: null,
            rules: [
                value => !!value || 'Required.',
            ]
        }
    },
    methods: {
        getCompany() {
            return new Promise((resolve, reject) => {
                this.id = this.$route.query.company
                if (this.id) {
                    this.$store.commit(types.MUTATE_LOADER_ON);
                    axios.get(`/app/getCompany/${this.id}`).then(res => {
                        console.log("res :::::", res);
                        this.$store.commit("SET_USER", res.data)
                        this.$store.commit(types.MUTATE_LOADER_OFF)
                    }).catch(err => {
                        console.log(err.response);
                        reject(err.response)
                        this.$store.commit(types.MUTATE_LOADER_OFF)
                        this.$toast.error("error occured while getting company!!!")
                    })
                }
            })
        },
        checklicense() {
            return new Promise((resolve, reject) => {
                this.$store.commit(types.MUTATE_LOADER_ON)
                axios.get(`/app/checkLicense`).then(res => {
                    console.log("res:::::", res);
                    if (!res.data) {
                        this.$toast.error("license is not present, activate system!!!")
                        this.$store.commit(types.MUTATE_LOADER_OFF)
                        return
                    }
                    resolve(res.data)
                }).catch(err => {
                    console.log(err.response);
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while checking license!!!")
                    reject(err.response)
                })
            })
        },
        checkActivation() {
            return new Promise((resolve, reject) => {
                this.$store.commit(types.MUTATE_LOADER_ON)
                let payload = {
                    branchId: this.id,
                    cpuId: this.clientDetails.client_cpu,
                    hardDiskId: this.clientDetails.client_disk_id,
                }
                axios.post(`/app/checkActivation`, payload).then(res => {
                    console.log("res:::::", res);
                    this.activated = res.data
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    resolve(res.data)
                }).catch(err => {
                    console.log(err.response);
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while getting company!!!")
                    reject(err.response)
                })

            })
        },
        invokeMainProcess() {
            return new Promise((resolve, reject) => {
                this.$store.commit(types.MUTATE_LOADER_ON)
                if (window.ipc) {
                    window.ipc.send("GET_CLIENT", {})
                    window.ipc.on("GET_CLIENT_REPLY", res => {
                        console.log("from electron process", res);
                        this.$store.commit(types.MUTATE_LOADER_OFF)
                        this.clientDetails = res
                        resolve(res)
                    })
                }
            })
            // Send an IPC message to the main process
        },
        activateCounter() {
            return new Promise((resolve, reject) => {
                this.$store.commit(types.MUTATE_LOADER_ON)
                console.log(":::::::::::::", this.counter);
                // this.invokeMainProcess().then(res => {
                // this.clientDetails = res
                // console.log("CLIENTDETAILSSSSS", this.clientDetails);
                if (!this.clientDetails || !this.clientDetails.client_cpu === "" || !this.clientDetails.client_disk_id === "") {
                    this.$toast.error("error getting client details for activation.")
                }
                let payload = {
                    branchId: this.counter.branchCode,
                    cpuId: this.clientDetails.client_cpu,
                    hardDiskId: this.clientDetails.client_disk_id,
                    counterId: this.counter.id,
                    counterName: this.counter.counterName
                }
                axios.post(`/license/counteractivate`, payload).then(res => {
                    console.log("res:::::", res);
                    this.checkActivation().then(activated => {
                        this.$toast.success("Counter activated on the system!!!")
                        this.$store.commit(types.MUTATE_LOADER_OFF)
                        // this.$forceUpdate()
                        if (window){
                            resolve(res.data)
                            window.reload()
                        }
                    })
                }).catch(err => {
                    console.log(err.response);
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while activating counter!!!")
                    reject(err.response)
                })
                // })
            })

        },
        updateSystemIP() {
            if (window.ipc) {
                window.ipc.send("UPDATE_IP", {})
                window.ipc.on("UPDATE_IP_REPLY", res => {
                    console.log("from electron process", res);
                    // resolve(res)
                })
            }
            // return new Promise((resolve, reject) => {
            //     localinstance.get(`/api/updateIP`).then(res => {
            //         this.$toast.success("Counter activated on the system!!!")
            //         resolve()
            //     }).catch(err => {
            //         this.$toast.error("error occured while activating counter!!!")
            //         reject()
            //     })

            // })
        },
        checkCounter() {
            return new Promise((resolve, reject) => {
                this.$store.commit(types.MUTATE_LOADER_ON)
                console.log(":::::::::::::", this.counter);
                let payload = {
                    cpuId: this.clientDetails.client_cpu,
                    hardDiskId: this.clientDetails.client_disk_id,
                }
                axios.post(`/app/checkandgetcounter`, payload).then(res => {
                    console.log("res:::::", res);
                    this.activeCounter = res.data
                    this.$store.commit('SET_ACTIVE_COUNTER', res.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    resolve(res.data)
                }).catch(err => {
                    console.log(err.response);
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while getting active counter!!!")
                    reject(err.response)
                })
            })

        },
        GetIp() {
            return new Promise((resolve, reject) => {
                this.$store.commit(types.MUTATE_LOADER_ON)
                if (window.ipc) {
                    window.ipc.send("SERVER_IP", {})
                    window.ipc.on("SERVER_IP_REPLY", res => {
                        console.log("from electron process", res);
                        this.$store.commit(types.MUTATE_LOADER_OFF)
                        resolve(res)
                    })
                }
                // localinstance.get(`/api/getServerIP`).then(res => {
                //     console.log("res:::::", res);
                //     this.$store.commit(types.MUTATE_LOADER_OFF)
                //     resolve(res.data)
                // }).catch(err => {
                //     console.log(err.response);
                //     this.$store.commit(types.MUTATE_LOADER_OFF)
                //     this.$toast.error("error occured while getting server IP!!!")
                //     reject(err.response)
                // })
            })
            // Send an IPC message to the main process
        },
        getAllCounters() {
            // console.log("store>>>>>>>>>>>>>>>>>>>",);
            this.tableData = []
            this.$store.commit(types.MUTATE_LOADER_ON)
            axios.get(`/counter/getinactivecounter/${this.$store.state.Auth.user.company_code}`).then(res => {
                console.log(">>>>>>>>>>>>>>>>>", res);
                res.data.message.forEach(element => {
                    let data = {}
                    data.counterNumber = element.CounterNumber
                    data.counterName = element.CounterName
                    data.branchCode = element.BranchCode
                    data.branchName = element.BranchName
                    data.date = element.CreatedAt
                    data.id = element.ID
                    this.tableData.push(data)
                });

                this.$toast.success("counters fetched successfully.")
                // this.connection.send("Hello World !!!")
            }).catch(err => {
                this.$store.commit(types.MUTATE_LOADER_OFF)
                console.log(err.response);
                this.$toast.error("error occured while getting counters!!!")
            })
        },
        counterlogin() {
            return new Promise((resolve, reject) => {
                this.$store.commit(types.MUTATE_LOADER_ON)
                this.counterUser.counter_id = this.$store.state.Auth.activeCounter.ID
                axios.post(`/counter/auth`, this.counterUser).then(res => {
                    console.log("res:::::", res);
                    this.user = res.data
                    this.$store.commit("SET_ACTIVE_USER", this.user)
                    this.$toast.success("Logged In !!!")
                    this.$router.push("/teller/teller-screen")
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.updateSystemIP()
                    resolve(res.data)
                }).catch(err => {
                    // console.log(err.response);
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while authenticating counter!!!")
                    reject(err.response)
                })
            })

        },
    },
    created() {
        this.$store.commit(types.MUTATE_LOADER_ON)
        this.getCompany()
        this.invokeMainProcess().then(res => {
            this.clientDetails = res
            // console.log("CLIENTDETAILSSSSS", this.clientDetails);
            this.checklicense().then(res => {
                // this.GetIp().then(res => {
                this.checkActivation().then(res => {
                    // this.activated = res
                    this.checkCounter().then(res => {
                        // this.activeCounter = res
                        // this.$store.commit('SET_ACTIVE_COUNTER', this.activeCounter)
                        // console.log("object:::::::::::::::", this.activeCounter);
                    })
                }).catch(err => {
                    // console.log(err.response);
                    this.$toast.error("error occured while checking activation!!!")
                })
                // }).catch(err => {
                //     // console.log(err.response);
                //     this.$toast.error("error occured while getting server IP!!!")
                // })
            }).catch(err => {
                // console.log(err.response);
                this.$toast.error("error occured while checking license!!!")
            })

        }).catch(err => {
            // console.log(err.response);
            this.$toast.error("error occured while getting client details!!")
        })
        this.getAllCounters()
        // this.invokeMainProcess()
        console.log("window", window.ipc);
    }
}
</script>

<style lang="scss" scoped>
.illustrator-img {
    position: absolute;
    inset-block-end: 0;
    inset-inline-end: 5%;
}

#lock {
    //   background-image: url("https://picsum.photos/200/300");
    background-image: url("https://source.unsplash.com/user/wsanter");
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center center;
}
</style>