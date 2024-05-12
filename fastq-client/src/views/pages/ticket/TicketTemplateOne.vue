<template>
  <v-app>
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
      <v-row justify="center" v-if="this.$store.state.Auth.user">
        <v-dialog
          v-model="dialog"
          v-if="allconf.ds_conf.show_form"
          persistent
          max-width="600px"
        >
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
                    <v-text-field
                      label="Legal first name*"
                      required
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="6" md="4">
                    <v-text-field
                      label="Legal last name*"
                      hint="example of persistent helper text"
                      persistent-hint
                      required
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12">
                    <v-text-field label="Email*" required></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="6">
                    <v-select
                      :items="['0-17', '18-29', '30-54', '54+']"
                      label="Age*"
                      required
                    ></v-select>
                  </v-col>
                  <v-col cols="12" sm="6">
                    <v-autocomplete
                      :items="[
                        'Skiing',
                        'Ice hockey',
                        'Soccer',
                        'Basketball',
                        'Hockey',
                        'Reading',
                        'Writing',
                        'Coding',
                        'Basejump',
                      ]"
                      label="Interests"
                      multiple
                    ></v-autocomplete>
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
      <div
        v-if="!this.$store.state.Auth.user"
        class="d-flex align-items-center justify-content-center"
      >
        <h1>Company Not Found!! Contact Service Provider.</h1>
      </div>
      <div
        class="row fill-height align-items-center justify-content-center"
        v-if="this.$store.state.Auth.user"
      >
        <!-- <div class="d-flex col-md-12 flex-column align-items-center">
            <v-card
              v-for="n in services"
              :key="n.id"
              class="mb-3 col-md-4 text-center"
              @click="addTicket(n)"
              shaped
              hover
              color="primary"
            >
              <h1 class="text-center white--text p-1">
                {{ n.name }}
              </h1>
            </v-card>
          </div> -->
        <div class="row col-md-12 align-content-center justify-content-center">
          <!-- <p class="text-h6">Please select from the services below :</p> -->
          <v-card
            v-for="n in services"
            :key="n.id"
            class="mb-3 col-md-4 mr-5 text-center bg_primary text"
            @click="addTicket(n)"
            shaped
            hover
            :class="{ disabled: wait }"
          >
            <h1
              class="text-center white--text p-1"
              v-if="customization.ticket_language === 'en-US'"
            >
              {{ n.name }}
            </h1>
            <h1 class="text-center white--text p-1" v-else>
              {{ n.arabic_name }}
            </h1>
          </v-card>
        </div>
      </div>
    </v-main>

    <v-footer padless fixed bottom elevation="4">
      <v-col class="text-center" cols="12">
        <!-- :class="{ 'd-none': $vuetify.breakpoint.mdAndDown, 'align-center': true }" -->
        <v-row>
          <v-col cols="12" class="d-flex align-items-center">
            <v-img
              class="ml-2"
              src="../../../assets/rsz_al_2.png"
              max-height="28"
              max-width="100"
              contain
            ></v-img>
            <v-col class="d-flex flex-column p-0">
              <p class="caption text-left p-0 mb-n1">
                Copyright © {{ new Date().getFullYear() }} FASTQ Solutions
                limited. All Rights Reserved
              </p>
              <p class="caption text-left p-0 mb-n1">
                Build version : 1.0.0 - {{ m.format("YYYY-MM-DD HH:mm:ss A") }}
              </p>
            </v-col>
          </v-col>
        </v-row>
        <!-- <v-spacer></v-spacer> -->
      </v-col>
    </v-footer>
  </v-app>
</template>
  
  <script>
import axios from "axios";
import moment from "moment";
import * as types from "../../../store/types.js";
import EscPosEncoder from "../../../../node_modules/esc-pos-encoder/dist/esc-pos-encoder.esm.js";
export default {
  props: ["ip", "company"],
  data() {
    return {
      wait: false,
      printCharacteristic: null,
      device: null,
      services: [],
      lastNumber: null,
      dialog: false,
      allconf: null,
      ipAddress: null,
      m: moment(),
      bluetooth: true,
      week: ["SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"],
      timerID: null,
      time: "",
      date: "",
      customization: null,
    };
  },
  methods: {
    printwithConnected(data) {
      const self = this;
      self.device.gatt
        .connect()
        .then((server) =>
          server.getPrimaryService("000018f0-0000-1000-8000-00805f9b34fb")
        )
        .then((service) =>
          service.getCharacteristic("00002af1-0000-1000-8000-00805f9b34fb")
        )
        .then((characteristic) => {
          // console.log("characteristic", characteristic);
          self.printCharacteristic = characteristic;
          let encoder = new EscPosEncoder({
            imageMode: "raster",
          });

          // let result = encoder
          //     .initialize()
          //     .text('The quick brown fox jumps over the lazy dog')
          //     .newline()
          //     .qrcode('https://nielsleenheer.com')
          //     .encode();

          // console.log("router", window.location);
          const url = new URL(window.location.origin);

          // Get the IP address from the URL
          self.ipAddress = url.hostname;
          let img = new Image();
          if (self.ipAddress === "localhost") {
            // return
            img.src = `http://${self.ipAddress}:8080/server/uploaded/logo.png`;
          } else {
            img.src = `https://${self.ipAddress}:443/uploaded/logo.png`;
          }
          // img.src = `http://localhost:8080/server/uploaded/logo.png`;

          img.onload = async function () {
            let result = encoder
              .align("center")
              .image(img, 160, 160, "atkinson")
              .newline()
              .encode();
            // console.log("hello", result);
            const dataSize = result.length;
            let startIndex = 0;
            const maxChunkSize = characteristic.properties.writeWithoutResponse
              ? 512
              : 20;
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
              let text = encoder
                .align("center")
                .width(3)
                .height(3)
                .line(data.ticket_name)
                .bold()
                .newline()
                .width(2)
                .height(2)
                .line(data.service)
                .bold()
                .newline()
                .width(2)
                .height(2)
                .line(`Waiting Count : ${data.Waiting}`)
                .width(1)
                .height(1)
                .line(`Created At : ${data.timeDate}`)
                .line(`Estimated Waiting Time : ${data.estTime}`)
                .newline()
                .line("")
                .newline()
                .encode();
              self.printCharacteristic.writeValue(text);
              // self.device.gatt
              //     .disconnect()
            });
          };
          // self.sendTextData(device)
        })
        .catch((error) => {
          console.log("error", error);
          // this.handleError(error, this.device)
        });
    },
    printlogo(startIndex, result, maxChunkSize) {
      return new Promise(async (resolve, reject) => {
        let dataSize = result.length;
        while (startIndex < dataSize) {
          const endIndex = Math.min(startIndex + maxChunkSize, dataSize);
          const chunk = result.slice(startIndex, endIndex);
          // Send the current chunk to the printer
          await this.printCharacteristic.writeValue(chunk);
          startIndex = endIndex;
        }
        resolve();
      });
    },
    sendPairingRequest() {
      window.electronAPI.bluetoothPairingRequest((event, details) => {
        const response = {};

        switch (details.pairingKind) {
          case "confirm": {
            response.confirmed = window.confirm(
              `Do you want to connect to device ${details.deviceId}?`
            );
            break;
          }
          case "confirmPin": {
            response.confirmed = window.confirm(
              `Does the pin ${details.pin} match the pin displayed on device ${details.deviceId}?`
            );
            break;
          }
          case "providePin": {
            const pin = window.prompt(
              `Please provide a pin for ${details.deviceId}.`
            );
            if (pin) {
              response.pin = pin;
              response.confirmed = true;
            } else {
              response.confirmed = false;
            }
          }
        }

        window.electronAPI.bluetoothPairingResponse(response);
      });
    },
    connectBluetoothPrinter() {
      // if (window.electronAPI) {
      //   window.ipc.send("select-bluetooth-device");
      //   console.log("electron send ", window.ipc);
      // }
      navigator.bluetooth
        .requestDevice(
          {
            filters: [
              {
                name: "MPT-III",
                services: ["000018f0-0000-1000-8000-00805f9b34fb"],
              },
            ],
          },
          {
            optionalServices: ["00002af1-0000-1000-8000-00805f9b34fb"],
          }
        )
        .then((device) => {
          console.log("device", device);

          // if (window.electronAPI) {
          //   this.sendPairingRequest();
          // }
          // window.alert(device)
          this.device = device;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    getAllServices() {
      let payload = {
        company_code: this.$store.state.Auth.user.company_code,
        branch_code: "",
      };
      axios
        .get(
          `/service/getAllservice/${this.$store.state.Auth.user.company_code}`
        )
        .then((res) => {
          console.log("res:::::", res);
          // this.$toast.success("branch added successfully.")
          res.data.message.forEach((element) => {
            let data = {};
            data.name = element.Name;
            data.prefix = element.Prefix;
            // data.lastname = element.Lastname
            // data.date = element.CreatedAt
            // data.branch = element.CompanyCode
            data.id = element.ID;
            data.numberStarts = element.NumberStarts;
            data.numberEnds = element.NumberEnds;
            data.code = element.Code;
            data.startTime = element.StartTime;
            data.endTime = element.EndTime;
            data.arabic_name = element.Workflow;
            // data.Hide = element.Hide
            // data.id = element.ID
            this.services.push(data);
          });
          console.log("services", this.services);
          // this.services = res.data
          // this.reset()
        })
        .catch((err) => {
          console.log(err.response);
          this.$toast.error("error occured while getting services!!!");
        });
    },
    async addTicket(ticket) {
      // if (window.electronAPI) {
      //   if (!this.device) {
      //     await this.connectBluetoothPrinter();
      //     this.$toast.info("Please Try Again !!!");
      //     return;
      //   }
      // }
      console.log("Ticket is",ticket);
      this.wait = true;
      let serviceName = null
      if (this.customization.ticket_language === "en-US") {
        serviceName = ticket.name;
        } else {
            serviceName = ticket.arabic_name;
        }
      this.lastticketNumber(serviceName).then((ticketNumber) => {
        console.log("ticket number", ticketNumber);
        let lastNumber = parseInt(ticketNumber);
        this.services.forEach((element) => {
          if (element.name === ticket.name) {
            if (lastNumber === 0) {
              // console.log("element is", element);
              lastNumber = element.numberStarts;
              // console.log("name matched");
            } else if (lastNumber > element.numberEnds) {
              lastNumber = element.numberStarts;
            } else {
              lastNumber += 1;
            }
          }
        });
        console.log("this.lastnumber", lastNumber);
        this.dialog = true;
        let ticketname = ticket.prefix + "-" + lastNumber;
        let date1 = new Date();
        let date2 = new Date();
        let ticketPayload = {
          service: ticket.name,
          ticket_status: "CREATED",
          ticket_number: lastNumber.toString(),
          started_serving_at: moment(date1).format("YYYY-MM-DD HH:mm:ss"),
          end_serving_at: moment(date2).format("YYYY-MM-DD HH:mm:ss"),
          ticket_name: ticketname,
          company_code: this.$store.state.Auth.user.company_code,
          company_name: this.$store.state.Auth.user.company,
          updated_by: this.$store.state.Auth.user.id,
        };
        if (this.customization.ticket_language === "en-US") {
          ticketPayload.service = ticket.name;
        } else {
          ticketPayload.service = ticket.arabic_name;
        }
        // console.log("::::::::::::::::::::", this.lastNumber++);
        axios
          .post("/ticket/addticket", ticketPayload)
          .then((res) => {
            console.log("res:::::", res);
            this.wait = false;
            // this.$toast.success("ticket added successfully.");
            if (window.electronAPI) {
              this.printElectronTicket(ticketPayload);
            } else {
              this.printTicket(ticketPayload);
            }
            this.branch = {};
          })
          .catch((err) => {
            console.log(err.response);
            this.$toast.error("error occured while adding branch!!!");
          });
      });
    },
    printTicket(ticketPayload) {
      ticketPayload.timeDate = moment(new Date()).format("DD MMM YY HH:MM A");
      this.waitingTickets().then((res) => {
        ticketPayload.Waiting = res;
        this.estimatedTime().then((avgTime) => {
          ticketPayload.estTime = avgTime;
        });
        console.log("object", ticketPayload.Waiting);
        // axios.post("http://localhost:4000/api/print", ticketPayload).then(res => {
        //     console.log("res:::::", res);
        //     this.$toast.success("ticket printed successfully.")
        //     this.branch = {}
        // }).catch(err => {
        //     console.log(err.response);
        //     this.$toast.error("error occured while printing ticket!!!")
        // })
        this.printwithConnected(ticketPayload);
      });
    },
    printElectronTicket(ticketPayload) {
      ticketPayload.timeDate = moment(new Date()).format(
        "YYYY-MM-DD HH:mm:ss A"
      );
      let printerconf = JSON.parse(this.allconf.ticket_conf.header_text);
      console.log("printer conf", printerconf);
      ticketPayload.printer = printerconf.printerName;
      this.waitingTickets().then((res) => {
        ticketPayload.Waiting = res;
        this.estimatedTime().then((avgTime) => {
          ticketPayload.estTime = avgTime;
        });
        console.log("object", ticketPayload.Waiting);
        axios
          .post("http://localhost:4000/api/print", ticketPayload)
          .then((res) => {
            console.log("res:::::", res);
            // this.$toast.success("ticket printed successfully.")
            this.branch = {};
          })
          .catch((err) => {
            console.log(err.response);
            // this.$toast.error("error occured while printing ticket!!!")
          });
      });
    },
    waitingTickets() {
      return new Promise((resolve, reject) => {
        axios
          .get(
            `/report/allwaitingTickets/${this.$store.state.Auth.user.company_code}`
          )
          .then((res) => {
            console.log("res:::::", res);
            resolve(res.data);
          })
          .catch((err) => {
            console.log(err.response);
            reject(err);
            // this.$toast.error("error occured while printing ticket!!!")
          });
      });
    },
    estimatedTime() {
      return new Promise((resolve, reject) => {
        axios
          .get(
            `/ticket/getestimated/${this.$store.state.Auth.user.company_code}`
          )
          .then((res) => {
            console.log("res:::::", res);
            resolve(res.data);
          })
          .catch((err) => {
            console.log(err.response);
            reject(err);
            // this.$toast.error("error occured while printing ticket!!!")
          });
      });
    },
    lastticketNumber(name) {
      return new Promise((resolve, reject) => {
        axios
          .get(`/ticket/getLastNumber/${name}`)
          .then((res) => {
            console.log("res:::::", res);
            // this.lastNumber = res.data.number;
            resolve(res.data.number);
            // this.$toast.success("branch added successfully.")
            // this.branch = {}
          })
          .catch((err) => {
            console.log(err.response);
            this.$toast.error("error occured while getting last ticket!!!");
          });
      });
    },
    getAllConfigs() {
      axios
        .get(`/config/getallconfig/${this.$store.state.Auth.user.company_code}`)
        .then((res) => {
          // console.log("res :::::", res);
          this.allconf = res.data;
          // this.allconf.schedular_conf.forEach(element => {
          //     console.log("single schedular", element);
          //     this.dayWiseVideo[element.day].push(element.video_id)
          // });
          // console.log("schedular", this.dayWiseVideo);
          // this.$store.commit("SET_USER", res.data)
          // this.$store.commit(types.MUTATE_LOADER_OFF)
        })
        .catch((err) => {
          console.log(err.response);
          // reject(err.response)
          // this.$store.commit(types.MUTATE_LOADER_OFF)
        //   this.$toast.error("error occured while getting connected clients!!!");
        });
    },
    getCompany() {
      // console.log("object", this.$route);
      return new Promise((resolve, reject) => {
        if (this.company) {
          this.id = this.company;
          this.$store.commit(types.MUTATE_LOADER_ON);
          axios
            .get(`/app/getCompany/${this.id}`)
            .then((res) => {
              // console.log("res :::::", res);
              this.$store.commit("SET_USER", res.data);
              this.$store.commit(types.MUTATE_LOADER_OFF);
              resolve();
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
    getcustomizations() {
      return new Promise((resolve, reject) => {
        // this.id = this.$route.query.company;
        // if (this.id) {
        this.$store.commit(types.MUTATE_LOADER_ON);
        axios
          .get(`/customize/getcustomizations`)
          .then((res) => {
            this.customization = res.data;
            console.log("object,cust", this.customization);
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
  },
  mounted() {
    // console.log("object", this.$route.query);
    // console.log("object", window.electronAPI);
    if (window.electronAPI) {
      this.bluetooth = false;
    }
  },
  created() {
    this.getcustomizations();
    this.getCompany().then((res) => {
      this.getAllServices();
      this.getAllConfigs();
      // this.lastticketNumber()
    });
    this.timerID = setInterval(this.updateTime, 1000);
    setInterval(this.getAllConfigs, 5000);
    this.updateTime();
    const url = new URL(window.location.origin);

    // Get the IP address from the URL
    this.ipAddress = url.hostname;
  },
};
</script>
  
  <style lang="scss" scoped>
#app {
  background-color: #eee;
}
.time {
  text-shadow: 2px 4px 2px #0096a786;
}
.disabled {
  pointer-events: none;
  /* Disable click events */
  opacity: 0.5;
  /* Optionally, reduce opacity to visually indicate that it's disabled */
}
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
.anim {
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