<template>
    <v-app>
        <Header v-if="allconf" :conf="allconf" />
        <v-main>
            <div class="row fill-height">
                <audio ref="dingSound"
                    :src="`http://${this.$route.query.ip}:8090/uploaded/${allconf.audio_conf.bell_name}`"></audio>
                <div class="col-md-8 primary d-flex flex-column align-items-start justify-content-center h-100">
                    <!-- <h5 class="text-white ml-5 h3">NOW SERVING</h5> -->
                    <div class="d-flex align-center justify-content-around">
                        <v-card :class="allconf.ds_conf.show_url ? 'col-md-10' : 'col-md-12'" shaped>
                            <div class="row d-flex align-centerjustify-center">
                                <video v-if="allconf.video_conf && currentVideo && !allconf.ds_conf.show_url"
                                    ref="videoPlayer" :src="`http://${this.$route.query.ip}:8090/uploaded/` + currentVideo"
                                    autoplay controls @ended="playNextVideo" width="800" height="450"></video>
                                <iframe v-if="allconf.ds_conf && allconf.ds_conf.show_url" :src="allconf.ds_conf.url"
                                    height="600" style="width: 1000px;" class="p-2"></iframe>
                                <!-- <div class="col-md-7">
                                    <div class="d-flex flex-column">
                                        <h2 class="h1">COUNTER </h2>
                                        <h1 style="font-size: 10rem;">1 </h1>
                                    </div>
                                </div>
                                <div class="col-md-6">
                                    <div class="d-flex flex-column">
                                        <h2 class="h1">TICKET </h2>
                                        <h1 style="font-size: 10rem;">1000</h1>
                                    </div>
                                </div>
                                <div class="col-md-12">
                                    <div class="d-flex flex-column">
                                        <h2 class="h1" style="font-size: 5rem;">PAYMENT</h2>
                                    </div>
                                </div> -->
                            </div>
                        </v-card>

                    </div>
                </div>
                <div class="col-md-4 black text-white align-content-center">
                    <!-- <h3 class="mt-3">NOW SERVING</h3> -->
                    <!-- <div class="white w-100" style="height: 1px;"></div> -->
                    <!-- {{ activeClients }} -->
                    <!-- {{ activeClients }} -->
                    <div class="col-md-12 d-flex flex-wrap">
                        <div class="col-md-6">
                            <h4>Counter</h4>
                        </div>
                        <div class="col-md-6">
                            <h4>Ticket</h4>
                        </div>
                    </div>
                    <div v-for="(n, i) in activeClients" class="d-flex flex-column align-items-center">
                        <div class="col-md-12 d-flex flex-wrap">
                            <div class="col-md-6 text-left py-0 my-0">
                                <h2 style="font-size: 3rem;">{{ n.CounterNumber }}</h2>
                                <!-- {{ n.ID }} -->
                            </div>
                            <div class="col-md-6 text-center py-0 my-0">
                                <h2 style="font-size: 3rem;"> <span v-if="$store.state.Auth.activeTickets && $store.state.Auth.activeTickets.findIndex(
                                    j => j.ticket_payload.CounterID === n.ID) !== -1">
                                        {{ $store.state.Auth.activeTickets[$store.state.Auth.activeTickets.findIndex(
                                            j => j.ticket_payload.CounterID === n.ID)].ticket_payload.TicketName }}
                                    </span></h2>
                            </div>
                            <!-- 
                                ==>
                               
                            </h4> -->

                            <!-- {{ $store.state.Auth.activeTickets }} -->
                            <!-- <h2
                                v-if="activeTickets.length !== 0 && activeTickets[i].ticket_payload.CounterID === n.counter_id">
                                {{ activeTickets[i].ticket_payload.Service }}</h2> -->
                            <!-- <h2>{{ activeTickets }}</h2> -->
                            <div class="white w-100" style="height: 1px;"></div>
                        </div>
                    </div>
                    <v-card>

                    </v-card>
                </div>
            </div>
        </v-main>
        <v-footer elevation="4" height="64px" fixed bottom app>
            <v-row>
                <v-col>
                    <ScrollingAd v-if="allconf.announcement_conf && allconf.ds_conf.show_scroll"
                        :ad-text="allconf.announcement_conf.announce_text"
                        :scroll-speed="allconf.announcement_conf.speed" />
                </v-col>
            </v-row>
        </v-footer>
    </v-app>
</template>

<script>
import axios from "axios";
import Header from "./Header.vue"
import VueNativeSock from 'vue-native-websocket';
import * as types from "../../store/types"
import ScrollingAd from "./comp/scrolling-ad.vue";
import moment from "moment";
export default {
    components: {
        Header,
        ScrollingAd
    },
    data() {
        return {
            id: null,
            socket: null,
            interval: null,
            clients: null,
            activeClients: [],
            activeTickets: [],
            allconf: null,
            days: ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'],
            videoList: [

                // Add more video URLs as needed
            ],
            m: moment(),
            currentVideoIndex: 0,
            dayWiseVideo: {},
            start: false,
        }
    },
    computed: {
        currentVideo() {
            return this.dayWiseVideo[this.m.format('dddd')][this.currentVideoIndex];
        },
        // activeTicketsMap() {
        //     // Any logic you need based on the data properties
        //     return this.activeTickets;
        // },
    },
    // watch: {
    //     activeTickets: {
    //         handler(newVal, oldVal) {
    //             // Handle changes in activeCounters
    //             this.$forceUpdate()
    //         },
    //         deep: true, // Watch nested properties
    //     },
    // },

    methods: {
        playNextVideo() {
            // Increment index to play the next video
            this.currentVideoIndex = (this.currentVideoIndex + 1) % this.dayWiseVideo[this.m.format('dddd')].length;
            // Restart the video
            if (this.$refs.videoPlayer) {
                this.$refs.videoPlayer.load();
                var playPromise = this.$refs.videoPlayer.play();

                if (playPromise !== undefined) {
                    playPromise.then(_ => {
                        this.start = true
                        // Automatic playback started!
                        // Show playing UI.
                    }).catch(error => {
                        this.start = true
                        // Auto-play was prevented
                        // Show paused UI.
                    });
                }

            }
            // this.$refs.videoPlayer.play();
        },
        callTicket(ticketNumber, counterNumber) {
            // Send a request to the server for text-to-speech
            // Handle the TTS response and call the TTS function
            this.performTextToSpeech(`Ticket number ${ticketNumber} is now ready on counter ${counterNumber}.`);
        },
        add(client) {
            const { length } = this.activeClients;
            const found = this.activeClients.some(el => el.counter_id === client.counter_id);
            if (!found) this.activeClients.push(client);
        },
        performTextToSpeech(text) {
            // Use Web Speech API or an external library for TTS
            const synth = window.speechSynthesis;
            const utterance = new SpeechSynthesisUtterance(text);

            // Optional: Adjust voice and other settings if needed
            // utterance.voice = ...

            // Speak the text
            synth.speak(utterance);
        },
        // Other methods for your Vue component
        connectToServer() {
            // this.connected = !this.connected;

            // Handle incoming messages
            this.socket.onmessage = (event) => {
                const response = JSON.parse(event.data)
                console.log('Received message from server:', response);
                switch (response.action) {
                    case "next":
                        // Vue.set(this.activeTickets, response.ticket_payload.CounterID, response);
                        // this.activeTickets.set(response.ticket_payload.CounterID, response)
                        console.log("active tickets", this.$store.state.Auth.activeTickets);
                        if (this.$store.state.Auth.activeTickets.length > 0) {
                            let ticketIndex = this.$store.state.Auth.activeTickets.findIndex(
                                i => i.ticket_payload.CounterID === response.ticket_payload.CounterID);
                            if (ticketIndex != -1) {
                                // REMOVE_ACTIVE
                                this.$store.commit("REMOVE_ACTIVE", ticketIndex)
                                // this.activeTickets.splice(index, 1)
                            }
                            // this.$forceUpdate();
                        }
                        // this.activeTickets.push(response)
                        this.$store.commit("ADD_ACTIVE", response)
                        // this.$forceUpdate();
                        this.playSound()
                        break;
                    case "finish":
                        // this.activeTickets = this.activeTickets.splice(this.activeTickets.findIndex(a => (a.ticket_payload.CounterID === response.ticket_payload.CounterID && a.ticket_payload.TicketName === response.ticket_payload.TicketName)), 1)
                        this.playSound()
                        if (this.$store.state.Auth.activeTickets.length > 0) {
                            let index = this.$store.state.Auth.activeTickets.findIndex(
                                i => i.ticket_payload.ID === response.ticket_payload.ID &&
                                    i.ticket_payload.CounterID === response.ticket_payload.CounterID);
                            // console.log("Fount on index", index);
                            // finally, remove matched item
                            // this.activeTickets.splice(index, 1)
                            if (index !== -1) {
                                this.$store.commit("REMOVE_ACTIVE", index)
                            }
                            console.log("Active ticket finished>>>>>>>>>>>", this.$store.state.Auth.activeTickets);
                            // this.$forceUpdate();
                            break;
                        }

                    default:
                        break;
                }
                // Process the response as needed
            };

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
        sendMessage() {
            // Send a message to the server
            if (this.socket.readyState === WebSocket.OPEN) {
                const data = 'Hello, Server!'; // Your message data
                this.socket.send(data);
                console.log('Sent message to server:', data);
            }

        },
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
        getServerIP() {

        },
        getActiveCounters(clientip, name) {
            // this.getConnected().then(connected => {
            // })
            // if (this.activeClients.legth > 0) {
            //     // for (let j = 0; j < this.clients.length; j++) {
            //     for (let index = 0; index < this.activeClients.length; index++) {
            //         const element = this.activeClients[index];
            //         if (element.client_ip === clientip) {
            //             continue
            //         } else {
            //             this.activeClients = this.activeClients.splice(index, 1)
            //         }
            //     }

            //     // }

            // } else {
            //     this.add(name)
            // }
            // for (const key in connected) {
            //     if (key === clientip) {
            //         // console.log("key fouind", key);

            //     } else {
            //         for (let index = 0; index < this.activeClients.length; index++) {
            //             const element = this.activeClients[index];
            //             console.log("object", element.client_ip === key);
            //             if (element.client_ip === key) {
            //                 // console.log("continuimng>>>>>>>>>>>>>.");
            //                 this.activeClients = this.activeClients.splice(index, 1)
            //                 // break
            //                 // continue
            //             }
            //         }

            //     }
            // };
            // axios.get(`http://${clientip}:3000/api/check`).then(res => {
            //     console.log("res client:::::", res);
            //     this.add(name)
            //     // this.activeClients.push(name)
            //     // this.$store.commit("SET_USER", res.data)
            //     // this.$store.commit(types.MUTATE_LOADER_OFF)
            // }).catch(err => {
            //     console.log("eroro active counters", err.response);
            //     this.activeClients.splice(this.activeClients.findIndex(a => a.counter_id === name.counter_id), 1)
            //     // this.activeClients.delete(name)
            //     // reject(err.response)
            //     // this.$store.commit(types.MUTATE_LOADER_OFF)
            //     // this.$toast.error("error occured while getting connected clients!!!")
            // })
        },
        getConnected() {
            return new Promise((resolve, reject) => {
                axios.get(`/getconnectedclients`).then(res => {
                    // console.log("res ::::: connected clients", res);
                    resolve(res.data)
                    // this.$store.commit("SET_USER", res.data)
                    // this.$store.commit(types.MUTATE_LOADER_OFF)
                }).catch(err => {
                    console.log(err.response);
                    reject(err)
                    // reject(err.response)
                    // this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while getting connected clients!!!")
                })
            })

        },
        getAllConfigs() {
            axios.get(`/config/getallconfig/${this.$store.state.Auth.user.company_code}`).then(res => {
                console.log("res :::::", res);
                this.allconf = res.data
                this.allconf.schedular_conf.forEach(element => {
                    // console.log("single schedular", element);
                    this.dayWiseVideo[element.day].push(element.video_id)
                });
                // console.log("schedular", this.dayWiseVideo);
                // this.$store.commit("SET_USER", res.data)
                // this.$store.commit(types.MUTATE_LOADER_OFF)
            }).catch(err => {
                console.log(err.response);
                // reject(err.response)
                // this.$store.commit(types.MUTATE_LOADER_OFF)
                // this.$toast.error("error occured while getting connected clients!!!")
            })

        },
        GetAllClients() {
            return new Promise((resolve, reject) => {
                axios.get(`/app/getAllClients`).then(res => {
                    this.clients = res.data
                    console.log("res >>>>>>>>>>:::::", res);
                    resolve(res.data)
                    // this.$store.commit("SET_USER", res.data)
                    // this.$store.commit(types.MUTATE_LOADER_OFF)
                }).catch(err => {
                    console.log(err.response);
                    // reject(err.response)
                    reject(err)
                    // this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while getting connected clients!!!")
                })

            })

        },
        customSetInterval(callback, interval) {
            function loop() {
                callback();
                setTimeout(loop, interval);
            }

            // Initial call to start the loop
            loop();
        },
        playSound() {
            // console.log(this.$refs.dingSound);
            if (this.$refs.dingSound) {
                let playPromise = this.$refs.dingSound.play();
                if (playPromise !== undefined) {
                    playPromise.then(_ => {
                        // this.start = true
                        // Automatic playback started!
                        // Show playing UI.
                    }).catch(error => {
                        // Auto-play was prevented
                        // Show paused UI.
                    });
                }
            }
        },
        initVideos() {
            this.days.forEach(element => {
                this.dayWiseVideo[element] = []
            });
        },
        addcounter(client) {
            const { length } = this.activeClients;
            const found = this.activeClients.some(el => el.counter_id === client.counter_id);
            if (!found) this.activeClients.push(client);
        },
        getAllcounters(id) {
            return new Promise((resolve, reject) => {
                let payload = {
                    id: id,
                    company_code: this.$store.state.Auth.user.company_code,
                    branch_code: this.$store.state.Auth.user.company_code,
                }
                // console.log("object:::::::", payload);
                axios.post(`/counter/getcounter`, payload).then(res => {
                    console.log("res:::::", res);
                    // this.counter = res.data.message
                    resolve(res.data.message)
                    // this.$toast.success("branch added successfully.")
                }).catch(err => {
                    console.log(err.response);
                    this.$toast.error("error occured while getting counter!!!")
                    reject(err)
                })
            })
        }
    },
    beforeMount() {
        // Disconnect from the server when the component is destroyed
        if (this.socket) {
            // this.socket.send("disconnect")
            this.socket.close();
            // console.log('Disconnected from WebSocket server');
            clearInterval(this.interval)
        }
    },
    mounted() {
        this.initVideos()
        this.getCompany()


        this.socket = new WebSocket(`ws://${this.$route.query.ip}:8090/ws`);
        this.connectToServer()
        // this.getConnected()
        // console.log("object", this.clients);
        this.interval = setInterval(() => {
            this.GetAllClients().then(res => {
                let temp = []
                res.forEach(element => {
                    console.log("element is", element);
                    this.getAllcounters(element.counter_id).then(counter => {
                        console.log("Counters are,", counter);
                        temp.push(counter)
                    })
                });
                // this.activeClients.push(counter)
                this.activeClients = temp
                // for (let i = 0; i < this.clients.length; i++) {
                //     // console.log("single object", this.clients[i]);
                //     if (this.clients[i].client_ip) {

                //     }
                //     // console.log("object map", this.activeClients);
                // }

            }).catch(err => {
                console.error(err);
                this.$toast.error("error occured while getting systems!!!")
            })
            //     this.getActiveCounters()
        }, 3000)
        this.getAllConfigs()
        setInterval(() => {
            this.getAllConfigs()
        }, 10000);
        setTimeout(() => {
            this.playSound()
        }, 3000);
    },
    created() {

    }
}
</script>

<style lang="scss" scoped></style>