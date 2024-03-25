<template>
  <v-app>
    <Header v-if="allconf" :conf="allconf" />
    <audio
      ref="dingSound"
      :src="`http://${this.$route.query.ip}:8090/uploaded/${allconf.audio_conf.bell_name}`"
    ></audio>
    <v-main>
      <div class="row fill-height">
        <div class="col-md-8 black d-flex flex-column align-items-center h-100">
          <div class="d-flex flex-column h-100 w-100 align-items-center mt-4">
            <!-- :class="allconf.ds_conf.retain_q ? 'justify-content-center' : ''" -->
            <v-card width="99%" height="70%" color="black">
              <!-- <v-card :class="allconf.ds_conf.show_url ? 'col-md-12' : 'col-md-12'" shaped> -->
              <div class="row d-flex align-items-center justify-content-center">
                <video
                  v-if="
                    allconf.video_conf &&
                    currentVideo &&
                    !allconf.ds_conf.show_url &&
                    !allconf.ds_conf.retain_q
                  "
                  ref="videoPlayer"
                  :src="
                    `http://${this.$route.query.ip}:8090/uploaded/` +
                    currentVideo
                  "
                  autoplay
                  controls
                  @ended="playNextVideo"
                  width="800"
                  height="500"
                  style="
                    width: 100%;
                    height: calc(100vh - 130px);
                    border: 5px solid #ad1457;
                  "
                ></video>
                <v-carousel
                  v-if="allconf.ds_conf && allconf.ds_conf.retain_q"
                  class="d-flex flex-column"
                  style="
                    width: 100%;
                    border: 5px solid #ad1457;
                    height: calc(100vh - 145px);
                  "
                  continuous
                  cycle
                  hide-delimiters
                >
                  <v-carousel-item
                    v-for="(item, i) in imageData"
                    reverse-transition="fade-transition"
                    transition="fade-transition"
                    :key="i"
                    class="mt-5"
                  >
                    <!-- class="d-flex align-items-center justify-content-center" -->
                    <!-- :aspect-ratio="2" -->
                    <v-img
                      :src="`http://${$route.query.ip}:8090/uploaded/${item}`"
                      :alt="item"
                      class="mt-5"
                      style="width: 100%; min-height: auto"
                      contain
                    />
                  </v-carousel-item>
                </v-carousel>
                <embed
                  v-if="
                    allconf.ds_conf &&
                    allconf.ds_conf.show_url &&
                    !allconf.ds_conf.retain_q
                  "
                  :src="allconf.ds_conf.url"
                  height="500"
                  style="
                    width: 100%;
                    border: 5px solid #ad1457;
                    height: calc(100vh - 150px);
                  "
                  class="p-2"
                />
              </div>
              <!-- </v-card> -->
            </v-card>
            <div
              class="d-flex align-items-center justify-content-around row w-100"
              v-if="false"
            >
              <v-card class="col-md-2 text-white p-2" color="grey darken-4">
                <h2 style="font-size: 1rem">PREVIOUSLY CALLED</h2>
              </v-card>
              <v-card class="col-md-2 text-white p-2" color="grey darken-4">
                <h3 style="font-size: 2rem">2-1001</h3>
              </v-card>
              <v-card class="col-md-2 text-white p-2" color="grey darken-4">
                <h2 style="font-size: 2rem">2-1001</h2>
              </v-card>
              <v-card class="col-md-2 text-white p-2" color="grey darken-4">
                <h2 style="font-size: 2rem">2-1001</h2>
              </v-card>
              <v-card class="col-md-2 text-white p-2" color="grey darken-4">
                <h2 style="font-size: 2rem">2-1001</h2>
              </v-card>
            </div>
          </div>
        </div>
        <div class="col-md-4 black text-white align-content-center">
          <!-- <h4 class="text-center">NOW SERVING</h4> -->
          <div
            class="col-md-12 d-flex flex-column align-items-center justify-content-between"
          >
            <div
              class="row col-md-12 d-flex justify-content-center align-items-center"
            >
              <!-- <div class="col-md-6"></div> -->
              <div class="col-md-6">
                <h3 class="text-center">COUNTER</h3>
              </div>
              <div class="col-md-6">
                <h3 class="text-center">TICKET</h3>
              </div>
            </div>
            <v-card
              class="row col-md-12 text-center d-flex justify-content-center gradient-card text-white align-items-center mb-4 p-0 m-0 mr-1 p-1"
              shaped
              v-for="(n, i) in activeClients"
              :key="i"
            >
              <!-- <div class="col-md-6"><h3>PAYMENT</h3></div> -->
              <div class="col-md-6 p-0 m-0">
                <h1 class="text-center" style="font-size: 4vw">
                  {{ n.CounterNumber }}
                </h1>
              </div>
              <div class="col-md-6 p-0 m-0">
                <h1 class="text-center" style="font-size: 4vw">
                  <span
                    v-if="
                      $store.state.Auth.activeTickets &&
                      $store.state.Auth.activeTickets.findIndex(
                        (j) => j.ticket_payload.CounterID === n.ID
                      ) !== -1
                    "
                  >
                    {{
                      $store.state.Auth.activeTickets[
                        $store.state.Auth.activeTickets.findIndex(
                          (j) => j.ticket_payload.CounterID === n.ID
                        )
                      ].ticket_payload.TicketName
                    }}
                  </span>
                </h1>
              </div>
            </v-card>
          </div>
        </div>
      </div>
    </v-main>

    <v-footer fixed bottoms height="64px" elevation="4" padless>
      <v-row>
        <v-col>
          <ScrollingAd
            v-if="allconf.announcement_conf && allconf.ds_conf.show_scroll"
            :ad-text="allconf.announcement_conf.announce_text"
            :scroll-speed="allconf.announcement_conf.speed"
          />
        </v-col>
      </v-row>
    </v-footer>
  </v-app>
</template>

<script>
import axios from "axios";
import Header from "./Header.vue";
import ScrollingAd from "./comp/scrolling-ad.vue";
import moment from "moment";
import * as types from "../../store/types";
export default {
  components: {
    Header,
    ScrollingAd,
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
      count: 0,
      days: [
        "Sunday",
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday",
        "Saturday",
      ],
      videoList: [
        // Add more video URLs as needed
      ],
      imageData: [],
      items: [
        {
          src: "https://cdn.vuetifyjs.com/images/carousel/squirrel.jpg",
        },
        {
          src: "https://cdn.vuetifyjs.com/images/carousel/sky.jpg",
        },
        {
          src: "https://cdn.vuetifyjs.com/images/carousel/bird.jpg",
        },
        {
          src: "https://cdn.vuetifyjs.com/images/carousel/planet.jpg",
        },
      ],
      m: moment(),
      currentVideoIndex: 0,
      dayWiseVideo: {},
      start: false,
      reconnectAttempts: 0,
      maxReconnectAttempts: 100,
      reconnectInterval: 1000,
    };
  },
  computed: {
    currentVideo() {
      return this.dayWiseVideo[this.m.format("dddd")][this.currentVideoIndex];
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
      this.currentVideoIndex =
        (this.currentVideoIndex + 1) %
        this.dayWiseVideo[this.m.format("dddd")].length;
      // Restart the video
      if (this.$refs.videoPlayer) {
        this.$refs.videoPlayer.load();
        var playPromise = this.$refs.videoPlayer.play();

        if (playPromise !== undefined) {
          playPromise
            .then((_) => {
              this.start = true;
              // Automatic playback started!
              // Show playing UI.
            })
            .catch((error) => {
              this.start = true;
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
      this.performTextToSpeech(
        `Ticket number ${ticketNumber} is now ready on counter ${counterNumber}.`
      );
    },
    add(client) {
      const { length } = this.activeClients;
      const found = this.activeClients.some(
        (el) => el.counter_id === client.counter_id
      );
      if (!found) this.activeClients.push(client);
    },
    performTextToSpeech(text) {
      // Use Web Speech API or an external library for TTS
      let s = this.setSpeech();
      s.then((voices) => {
        var synthesis = window.speechSynthesis;
        var utterance = new SpeechSynthesisUtterance(text);

        // This overrides the text "Hello World" and is uttered instead
        utterance.voice = voices[1];
        // utterance.text = text;
        utterance.pitch = 1.8;
        utterance.rate = 0.9;
        utterance.volume = 0.8;
        console.log("Calling Once");
        if (this.count == 0) {
          synthesis.speak(utterance);
          this.count++;
        }
      });
      // const synth = window.speechSynthesis;
      // const utterance = new SpeechSynthesisUtterance(text);
      // var voices = window.speechSynthesis.getVoices();
      // // var utterance = new SpeechSynthesisUtterance("Hello World");
      // utterance.rate = 0.7;
      // utterance.voice = voices[1];

      // // Speak the text
      // synth.speak(utterance);
    },
    // Other methods for your Vue component
    connectToServer() {
      // this.connected = !this.connected;

      // Handle incoming messages
      this.socket.onmessage = (event) => {
        this.reconnectAttempts = 0;
        const response = JSON.parse(event.data);
        console.log("Received message from server:", response);
        switch (response.action) {
          case "next":
            // Vue.set(this.activeTickets, response.ticket_payload.CounterID, response);
            // this.activeTickets.set(response.ticket_payload.CounterID, response)
            // console.log("active tickets", this.$store.state.Auth.activeTickets);
            if (this.$store.state.Auth.activeTickets.length > 0) {
              let ticketIndex = this.$store.state.Auth.activeTickets.findIndex(
                (i) =>
                  i.ticket_payload.CounterID ===
                  response.ticket_payload.CounterID
              );
              if (ticketIndex != -1) {
                // REMOVE_ACTIVE
                this.$store.commit("REMOVE_ACTIVE", ticketIndex);
                // this.activeTickets.splice(index, 1)
              }
              // this.$forceUpdate();
            }
            // this.activeTickets.push(response)
            this.$store.commit("ADD_ACTIVE", response);
            // this.$forceUpdate();
            let allowVoice = this.allconf.ds_conf.show_dt;
            if (allowVoice) {
              //   this.activeClients.forEach((client) => {
              let index = this.activeClients.findIndex(
                (j) => response.ticket_payload.CounterID === j.ID
              );

              if (index !== -1) {
                // if (this.count !== 1) {
                this.callTicket(
                  response.ticket_payload.TicketName,
                  this.activeClients[index].CounterNumber
                );
                // }
              }
              //   });
            } else {
              this.playSound();
            }
            // this.playSound();
            break;
          case "finish":
            this.count = 0;
            // this.activeTickets = this.activeTickets.splice(this.activeTickets.findIndex(a => (a.ticket_payload.CounterID === response.ticket_payload.CounterID && a.ticket_payload.TicketName === response.ticket_payload.TicketName)), 1)
            this.playSound();
            if (this.$store.state.Auth.activeTickets.length > 0) {
              let index = this.$store.state.Auth.activeTickets.findIndex(
                (i) =>
                  i.ticket_payload.ID === response.ticket_payload.ID &&
                  i.ticket_payload.CounterID ===
                    response.ticket_payload.CounterID
              );
              // console.log("Fount on index", index);
              // finally, remove matched item
              // this.activeTickets.splice(index, 1)
              if (index !== -1) {
                this.$store.commit("REMOVE_ACTIVE", index);
              }
              // console.log("Active ticket finished>>>>>>>>>>>", this.$store.state.Auth.activeTickets);
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
        console.log("WebSocket connection closed:", event);
        this.reconnect();
        // You may want to handle reconnection logic here if needed
      };
      this.socket.onopen = (event) => {
        // this.sendMessage()
        console.log("WebSocket connection opened:", event);
        // You may want to handle reconnection logic here if needed
      };
      this.socket.onerror = (event) => {
        console.log("WebSocket connection error:", event);
        // You may want to handle reconnection logic here if needed
      };
    },
    sendMessage() {
      // Send a message to the server
      if (this.socket.readyState === WebSocket.OPEN) {
        const data = "Hello, Server!"; // Your message data
        this.socket.send(data);
        // console.log('Sent message to server:', data);
      }
    },
    getCompany() {
      return new Promise((resolve, reject) => {
        this.id = this.$route.query.company;
        if (this.id) {
          this.$store.commit(types.MUTATE_LOADER_ON);
          axios
            .get(`/app/getCompany/${this.id}`)
            .then((res) => {
              // console.log("res :::::", res);
              this.$store.commit("SET_USER", res.data);
              this.$store.commit(types.MUTATE_LOADER_OFF);
            })
            .catch((err) => {
              console.log(err.response);
              reject(err.response);
              this.$store.commit(types.MUTATE_LOADER_OFF);
              this.$toast.error("error occured while getting company!!!");
            });
        }
      });
    },
    getServerIP() {},
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
        axios
          .get(`/getconnectedclients`)
          .then((res) => {
            // console.log("res ::::: connected clients", res);
            resolve(res.data);
            // this.$store.commit("SET_USER", res.data)
            // this.$store.commit(types.MUTATE_LOADER_OFF)
          })
          .catch((err) => {
            console.log(err.response);
            reject(err);
            // reject(err.response)
            // this.$store.commit(types.MUTATE_LOADER_OFF)
            // this.$toast.error("error occured while getting connected clients!!!")
          });
      });
    },
    getAllConfigs() {
      axios
        .get(`/config/getallconfig/${this.$store.state.Auth.user.company_code}`)
        .then((res) => {
          // console.log("res :::::", res);
          this.allconf = res.data;
          this.allconf.schedular_conf.forEach((element) => {
            // console.log("single schedular", element);
            this.dayWiseVideo[element.day].push(element.video_id);
          });
          // console.log("schedular", this.dayWiseVideo);
          // this.$store.commit("SET_USER", res.data)
          // this.$store.commit(types.MUTATE_LOADER_OFF)
        })
        .catch((err) => {
          console.log(err.response);
          // reject(err.response)
          // this.$store.commit(types.MUTATE_LOADER_OFF)
          // this.$toast.error("error occured while getting connected clients!!!")
        });
    },
    GetAllClients() {
      return new Promise((resolve, reject) => {
        axios
          .get(`/app/getAllClients`)
          .then((res) => {
            this.clients = res.data;
            // console.log("res >>>>>>>>>>:::::", res);
            resolve(res.data);
            // this.$store.commit("SET_USER", res.data)
            // this.$store.commit(types.MUTATE_LOADER_OFF)
          })
          .catch((err) => {
            console.log(err.response);
            // reject(err.response)
            reject(err);
            // this.$store.commit(types.MUTATE_LOADER_OFF)
            // this.$toast.error("error occured while getting connected clients!!!")
          });
      });
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
          playPromise
            .then((_) => {
              // this.start = true
              // Automatic playback started!
              // Show playing UI.
            })
            .catch((error) => {
              // Auto-play was prevented
              // Show paused UI.
            });
        }
      }
    },
    initVideos() {
      this.days.forEach((element) => {
        this.dayWiseVideo[element] = [];
      });
    },
    addcounter(client) {
      const { length } = this.activeClients;
      const found = this.activeClients.some(
        (el) => el.counter_id === client.counter_id
      );
      if (!found) this.activeClients.push(client);
    },
    getAllcounters(id) {
      return new Promise((resolve, reject) => {
        let payload = {
          id: id,
          company_code: this.$store.state.Auth.user.company_code,
          branch_code: this.$store.state.Auth.user.company_code,
        };
        // console.log("object:::::::", payload);
        axios
          .post(`/counter/getcounter`, payload)
          .then((res) => {
            // console.log("res:::::", res);
            // this.counter = res.data.message
            resolve(res.data.message);
            // this.$toast.success("branch added successfully.")
          })
          .catch((err) => {
            console.log(err.response);
            this.$toast.error("error occured while getting counter!!!");
            reject(err);
          });
      });
    },
    isImage(filename) {
      // Define common image extensions
      const imageExtensions = [
        "jpg",
        "jpeg",
        "png",
        "gif",
        "bmp",
        "svg",
        "webp",
      ];

      // Get the file extension
      const ext = filename
        .slice(((filename.lastIndexOf(".") - 1) >>> 0) + 2)
        .toLowerCase();
      //   console.log("extension", ext);
      // Check if the file extension is in the list of image extensions
      return imageExtensions.includes(ext, 0);
    },
    GetAllImages() {
      return new Promise((resolve, reject) => {
        axios
          .get(
            `/config/getallvideo/${this.$store.state.Auth.user.company_code}`
          )
          .then((res) => {
            // this.allvideos = res.data;
            // console.log("res >>>>>>>>>>:::::", res);
            this.imageData = [];
            if (res.data) {
              res.data.forEach((element) => {
                // let data = {};
                // console.log(this.isImage(element.file_name));
                if (this.isImage(element.file_name)) {
                  // data.Hide = element.Hide
                  // data.id = element.ID
                  this.imageData.push(element.file_name);
                }
              });
              resolve(res.data);
            }
            // this.$store.commit("SET_USER", res.data)
            // this.$store.commit(types.MUTATE_LOADER_OFF)
          })
          .catch((err) => {
            console.log(err.response);
            // reject(err.response)
            reject(err);
            // this.$store.commit(types.MUTATE_LOADER_OFF)
            this.$toast.error(
              "error occured while getting connected clients!!!"
            );
          });
      });
    },
    reconnect() {
      if (this.reconnectAttempts < this.maxReconnectAttempts) {
        setTimeout(() => {
          this.reconnectAttempts++;
          this.connect();
        }, this.reconnectInterval);
      } else {
        console.log("Max reconnect attempts reached");
      }
    },
    connect() {
      this.socket = new WebSocket(`ws://${this.$route.query.ip}:8090/ws`);
    },
    setSpeech() {
      return new Promise(function (resolve, reject) {
        let synth = window.speechSynthesis;
        let id;

        id = setInterval(() => {
          if (synth.getVoices().length !== 0) {
            resolve(synth.getVoices());
            clearInterval(id);
          }
        }, 10);
      });
    },
  },
  beforeMount() {
    // Disconnect from the server when the component is destroyed
    if (this.socket) {
      // this.socket.send("disconnect")
      this.socket.close();
      // console.log('Disconnected from WebSocket server');
      clearInterval(this.interval);
    }
  },
  mounted() {
    this.initVideos();
    this.getCompany();
    this.connect();

    this.connectToServer();
    // this.getConnected()
    // console.log("object", this.clients);
    this.interval = setInterval(() => {
      this.GetAllClients()
        .then((res) => {
          let temp = [];
          res.forEach((element) => {
            // console.log("element is", element);
            this.getAllcounters(element.counter_id).then((counter) => {
              // console.log("Counters are,", counter);
              temp.push(counter);
            });
          });
          // this.activeClients.push(counter)
          this.activeClients = temp;
          // for (let i = 0; i < this.clients.length; i++) {
          //     // console.log("single object", this.clients[i]);
          //     if (this.clients[i].client_ip) {

          //     }
          //     // console.log("object map", this.activeClients);
          // }
        })
        .catch((err) => {
          console.error(err);
          this.$toast.error("error occured while getting systems!!!");
        });
      //     this.getActiveCounters()
    }, 3000);
    this.getAllConfigs();
    setInterval(() => {
      this.getAllConfigs();
      this.GetAllImages();
    }, 10000);
    setTimeout(() => {
      this.playSound();
    }, 3000);
  },
  created() {
    // this.callTicket("A-12", "3");
  },
};
</script>

<style lang="scss" scoped>
.gradient-card {
  background-image: linear-gradient(150deg, #ad1457 50%, #0097a7 35%);
}
.v-carousel .v-window-item {
  height: 100% !important;
}
</style>