<template>
  <v-app class="body-d">
    <audio
      ref="dingSound"
      :src="`http://${this.ip}:8090/uploaded/${allconf.audio_conf.bell_name}`"
    ></audio>
    <div class="container-fluid bg-transparent">
      <div class="row d-flex justify-content-between">
        <div class="wrapper col-7 p-0 m-0 d-flex">
          <div class="col-4 col-md-2 col-lg-2 text-center">
            <a class="navbar-brand m-0" href="#">
              <img
                src="https://www.moi.gov.kw/main/images/assets/common/logo-moi.svg"
                style="height: 120px"
              />
            </a>
          </div>
          <div class="col-1 align-self-center">
            <div class="row">
              <div class="col text-center">
                <img
                  src="https://www.moi.gov.kw/main/images/assets/common/en/state-of-kuwait.svg"
                  class="text-center main-header-title-en"
                />
              </div>
            </div>
            <div class="row">
              <div class="col text-center">
                <img
                  src="https://www.moi.gov.kw/main/images/assets/common/en/ministry-of-interior.svg"
                  class="text-center main-header-title-en"
                />
              </div>
            </div>
          </div>
        </div>
        <div class="col-5 align-self-center">
          <div
            class="row d-flex flex-column justify-content-center align-items-center"
          >
            <div class="col-12 text-center temp_primary m-0 p-0">
              <h1>Time is : {{ time }}</h1>
            </div>
            <div class="col-12 text-center temp_primary m-0 p-0">
              <h1>Date is : {{ date }}</h1>
            </div>
          </div>
        </div>
      </div>
    </div>
    <v-main>
      <div class="row mt-5">
        <div class="col-md-12 d-flex align-items-center justify-content-center">
          <div class="card px-5 bg_primary">
            <h1>Ticket #</h1>
          </div>
        </div>
        <div
          class="col-6 d-flex align-items-center justify-content-center"
          v-for="(n, i) in activeClients"
          :key="i"
        >
          <div
            class="card px-5 d-flex align-items-center justify-content-center bg_primary mr-4"
            style="height: 7rem; min-width: 25rem"
          >
          <!-- :class="count > 0 ? 'anim' : ''" -->
            <h1
            class="text"
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
            </h1>
          </div>
          <div
            class="card second d-flex align-items-center justify-content-center bg_primary text-transparent"
            style="width: 7rem; height: 7rem; border-radius: 50%"
          >
          <!-- :class="count > 0 ? 'anim' : ''" -->
            <h1
              class="font-weight-bold py-2 text"
            >
              {{ n.CounterNumber }}
            </h1>
          </div>
        </div>
      </div>
    </v-main>
    <!-- <v-footer fixed height="64px" tag="footer"> </v-footer> -->
  </v-app>
</template>
<script>
import axios from "axios";
import Header from "./Header.vue";
import ScrollingAd from "./comp/scrolling-ad.vue";
import moment from "moment";
import * as types from "../../store/types";
export default {
  props: ["ip", "company"],
  components: {
    Header,
    ScrollingAd,
  },
  data() {
    return {
      m: moment(),
      ipAddress: null,
      week: ["SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"],
      timerID: null,
      time: "",
      date: "",
      customization: null,
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
      start: false,
      reconnectAttempts: 0,
      maxReconnectAttempts: 100,
      reconnectInterval: 1000,
    };
  },
  computed: {},

  methods: {
    callTicket(ticketNumber, counterNumber) {
      // Send a request to the server for text-to-speech
      // Handle the TTS response and call the TTS function
      if (this.customization.tts_language === "en-US") {
        console.log("in if english");
        this.performTextToSpeech(
          `Ticket number ${ticketNumber} is now ready on counter ${counterNumber}.`
        );
      } else {
        this.performTextToSpeech(
          `التذكرة رقم ${ticketNumber} جاهزة الآن في المنضدة ${counterNumber}.`
        );
      }
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
        // console.log("voices", voices);
        voices.forEach((v) => {
          if (v.lang.includes(this.customization.tts_languag)) {
            utterance.voice = v;
          }
        });
        if (this.customization.tts_language === "en-US") {
          utterance.lang = "en-us";
        } else {
          utterance.lang = "ar-kw";
        }
        // utterance.text = text;
        utterance.pitch = 2;
        utterance.rate = 0.4;
        utterance.volume = 5;
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
                this.count = 0;
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
          case "recall":
            let index = this.activeClients.findIndex(
              (j) => response.ticket_payload.CounterID === j.ID
            );

            if (index !== -1) {
              this.count = 0;
              // if (this.count !== 1) {
              this.callTicket(
                response.ticket_payload.TicketName,
                this.activeClients[index].CounterNumber
              );
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
        this.id = this.company;
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
      this.socket = new WebSocket(`ws://${this.ip}:8090/ws`);
    },
    updateTime() {
      var cd = new Date();
      this.time =
        this.zeroPadding(cd.getHours(), 2) +
        ":" +
        this.zeroPadding(cd.getMinutes(), 2) +
        ":" +
        this.zeroPadding(cd.getSeconds(), 2);
      this.date =
        this.zeroPadding(cd.getFullYear(), 4) +
        "-" +
        this.zeroPadding(cd.getMonth() + 1, 2) +
        "-" +
        this.zeroPadding(cd.getDate(), 2) +
        " " +
        this.week[cd.getDay()];
    },
    zeroPadding(num, digit) {
      var zero = "";
      for (var i = 0; i < digit; i++) {
        zero += "0";
      }
      return (zero + num).slice(-digit);
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
    getcustomizations() {
      return new Promise((resolve, reject) => {
        // this.id = this.$route.query.company;
        // if (this.id) {
        this.$store.commit(types.MUTATE_LOADER_ON);
        axios
          .get(`/customize/getcustomizations`)
          .then((res) => {
            this.customization = res.data;
            // console.log("object,cust", this.customization);
            this.$store.commit(types.MUTATE_LOADER_OFF);
          })
          .catch((err) => {
            console.log(err.response);
            reject(err.response);
            this.$store.commit(types.MUTATE_LOADER_OFF);
            this.$toast.error("error occured while getting company!!!");
          });
        // }
      });
    },
  },

  beforeDestroy() {
    // Disconnect from the server when the component is destroyed
    if (this.socket) {
      this.socket.send("disconnect");
      this.socket.close();
      // console.log('Disconnected from WebSocket server');
      clearInterval(this.interval);
      clearInterval(this.timerID);
    }
  },
  mounted() {
    console.log("object,", this.ip);
    console.log("object,", this.company);
    this.getcustomizations();
    this.getCompany();
    this.connect();
    // this.callTicket("A002","2")
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
          // this.$toast.error("error occured while getting systems!!!");
        });
      //     this.getActiveCounters()
    }, 3000);
    this.getAllConfigs();
    setInterval(() => {
      this.getAllConfigs();
    }, 10000);
    setTimeout(() => {
      this.playSound();
    }, 3000);
  },
  created() {
    this.timerID = setInterval(this.updateTime, 1000);
    this.updateTime();
  },
};
</script>

<style lang="scss">
.main-header-title-en {
  width: 12.4em;
}
.temp_primary {
  color: #040873;
}
.temp_red {
  color: red;
}
.bg_primary {
  background-color: #040873;
  color: #eceae4;
}
.text {
  font-size: clamp(4rem, 7vw, 8rem);
}
.anim{
  animation: color-change 2s infinite;
}
@keyframes color-change {
  0% {
    color: red;
  }
  50% {
    color: #eceae4;
  }
  100% {
    color: red;
  }
}
.body-d {
  font-size: 22px;
  font-family: "Droid Arabic Kufi Regular", "Skia Regular", Arial, Tahoma,
    sans-serif;
  background-color: #eceae4 !important;
  background-image: url(https://www.moi.gov.kw/main/images/assets/common/bg-pattern.png) !important;
  height: 100%;
  /* margin-bottom: 280px; */
}
</style>