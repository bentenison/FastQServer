<template>
    <div class="container h-100">
        <div class="row h-100">
            <div class="row col-md-12 h-100">
                <v-card class="col-md-12 h-auto" flat>
                    <h6 class="grey--text">Select Customer Display Screen</h6>
                    <ul class="col-md-12">
                        <li class="col-md-6"><input type="radio" name="test" value="SignageTemplate1"
                                v-model="dsconf.template" id="cb1" />

                            <label for="cb1"><v-img max-height="100" max-width="150"
                                    src="../../assets/templates-insp/template1.png" /></label>
                        </li>
                        <li class="col-md-6"><input type="radio" value="SignageTemplate2" v-model="dsconf.template"
                                name="test" id="cb2" />
                            <label for="cb2"><v-img max-height="100" max-width="150"
                                    src="../../assets/templates-insp/template5.png" /></label>
                        </li>
                        <li class="col-md-6"><input type="radio" value="SignageTemplate3" v-model="dsconf.template"
                                name="test" id="cb3" />
                            <label for="cb3"><v-img max-height="100" max-width="150"
                                    src="../../assets/templates-insp/template3.png" /></label>
                        </li>
                        <li class="col-md-6"><input type="radio" value="SignageTemplate4" v-model="dsconf.template"
                                name="test" id="cb4" />
                            <label for="cb4"><v-img max-height="100" max-width="150"
                                    src="../../assets/templates-insp/template4.png" /></label>
                        </li>
                    </ul>
                    <v-checkbox v-model="dsconf.confirmation_delay"
                        label="Show Queue Ticket Confirmation Message Pop Up Delay (seconds)" color="primary"
                        hide-details></v-checkbox>
                    <v-checkbox v-model="dsconf.retain_q" label="Retain Called Queues on Display" color="primary"
                        hide-details></v-checkbox>
                    <v-checkbox v-model="dsconf.show_form" label="Show Customer Data Entry Form on Kiosk" color="primary"
                        hide-details></v-checkbox>
                    <div class="col-md-4 ml-3" v-if="false">
                        <v-checkbox v-model="name" label="Name" color="primary" hide-details></v-checkbox>
                        <v-checkbox v-model="phone" label="Phone" color="primary" hide-details></v-checkbox>
                        <v-checkbox v-model="email" label="Email" color="primary" hide-details></v-checkbox>
                        <v-checkbox v-model="date" label="Date Field" color="primary" hide-details></v-checkbox>
                        <div class="col-md-6 pl-3 mt-3" v-if="date">
                            <v-text-field class="pl-5" label="date label" dense></v-text-field>
                        </div>
                        <v-checkbox v-model="ex6" label="Data Field #1" color="primary" hide-details></v-checkbox>
                        <div class="col-md-6 pl-3 mt-3" v-if="ex6">
                            <v-text-field class="pl-5" label="date label" dense></v-text-field>
                        </div>
                        <v-checkbox v-model="ex7" label="Data Field #2" color="primary" hide-details></v-checkbox>
                        <div class="col-md-6 pl-3 mt-3" v-if="ex7">
                            <v-text-field class="pl-5" label="date label" dense></v-text-field>
                        </div>
                        <v-checkbox v-model="ex4" label="Data Field #3" color="primary" hide-details></v-checkbox>
                        <v-checkbox v-model="ex4" label="Data Field #4" color="primary" hide-details></v-checkbox>
                    </div>
                    <v-checkbox v-model="dsconf.show_scroll" label="Show the scroll message" color="primary"
                        hide-details></v-checkbox>
                    <v-checkbox v-model="dsconf.show_dt" label="Show the current date and time" color="primary"
                        hide-details></v-checkbox>
                    <v-checkbox v-model="dsconf.show_url" label="Show URL on DS" color="primary"
                        hide-details></v-checkbox>
                        <v-text-field label="enter url to display" class="my-4 p-0 col-md-6" v-if="dsconf.show_url" placeholder="" v-model="dsconf.url"
                        outlined dense></v-text-field>
                    <v-btn color="primary" class="mt-2" v-if="createOne" @click="adddsConf"><v-icon
                            left>mdi-content-save-all</v-icon>Save
                        config</v-btn>
                    <v-btn color="primary" class="mt-2" v-if="!createOne" @click="updateDSConf"><v-icon left>mdi-update</v-icon>Update
                        config</v-btn>
                </v-card>
                <div class="row col-md-12">
                    <v-expansion-panels v-model="panel" :readonly="readonly" multiple flat>
                        <v-expansion-panel>
                            <v-expansion-panel-header class="grey--text">
                                <p class="text-center h5">AUDIT TRAIL</p>
                            </v-expansion-panel-header>
                            <v-expansion-panel-content>
                                <DataTables />
                            </v-expansion-panel-content>
                        </v-expansion-panel>
                    </v-expansion-panels>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import DataTables from '../tables/dataTables.vue';
import * as types from "../../store/types"
import axios from 'axios';
export default {
    data() {
        return {
            items: ["India", "America", "Europe", "Japan"],
            date: null,
            name: null,
            phone: null,
            email: null,
            ex6: null,
            ex7: null,
            panel: "",
            readonly: false,
            dsconf: {
                // id: "",
                // branch_code: "",
                // branch_name: "",
                // company_code: "",
                // company_name: "",
                // template: "",
                // retain_q: false,
                // confirmation_delay: "",
                // show_form: "",
                // show_scroll: "",
                // show_dt: "",
                // volume: "",
                // form_conf: "",
                // online_form: ""
            },
            createOne: false
        };
    },
    components: { DataTables },
    methods: {
        adddsConf() {
            console.log("object>>>>>>>>>>>", this.dsconf);
            return new Promise((resolve, reject) => {
                // this.id = this.$route.query.company
                // if (this.id) {
                // this.announcement.speed = parseInt(this.announcement.speed)
                this.dsconf.company_code = this.$store.state.Auth.user.company_code
                this.dsconf.company_name = this.$store.state.Auth.user.company
                this.dsconf.updated_by = this.$store.state.Auth.user.id
                this.$store.commit(types.MUTATE_LOADER_ON);
                axios.post(`/config/setdsconfig`, this.dsconf).then(res => {
                    console.log("res :::::", res);
                    this.$toast.success("ds config added!!!")
                    // this.$store.commit("SET_USER", res.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    resolve()
                }).catch(err => {
                    console.log(err.response);
                    reject(err.response)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while adding ds config!!!")
                })
                // }
            })
        },
        updateAnnouncement() {
            return new Promise((resolve, reject) => {
                // this.id = this.$route.query.company
                // if (this.id) {
                this.$store.commit(types.MUTATE_LOADER_ON);
                axios.post(`/config/updateannounce`).then(res => {
                    console.log("res :::::", res);
                    // this.$store.commit("SET_USER", res.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                }).catch(err => {
                    console.log(err.response);
                    reject(err.response)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while adding announcement!!!")
                })
                // }
            })
        },
        updateDSConf() {
            return new Promise((resolve, reject) => {
                // this.id = this.$route.query.company
                // if (this.id) {
                this.$store.commit(types.MUTATE_LOADER_ON);
                axios.post(`/config/updateds`,this.dsconf).then(res => {
                    console.log("res :::::", res);
                    this.$toast.success("DS Config Updated!!!")
                    // this.$store.commit("SET_USER", res.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                }).catch(err => {
                    console.log(err.response);
                    reject(err.response)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while updating DS Config!!!")
                })
                // }
            })
        },
        getDsConf() {
            return new Promise((resolve, reject) => {
                // this.id = this.$route.query.company
                // if (this.id) {
                this.$store.commit(types.MUTATE_LOADER_ON);
                axios.get(`/config/getdsconfig/${this.$store.state.Auth.user.company_code}`).then(res => {
                    console.log("res :::::", res);
                    this.dsconf = res.data
                    // this.$store.commit("SET_USER", res.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                }).catch(err => {
                    console.log(err.response.data.error.includes("sql:"));
                    if (err.response.data.error.includes("sql:")) {
                        this.createOne = true
                    }
                    // reject(err.response.data)
                    this.$store.commit(types.MUTATE_LOADER_OFF)
                    this.$toast.error("error occured while getting ds config!!!")
                })
                // }
            })
        },
    },
    created() {
        this.getDsConf()
    }
}
</script>

<style lang="scss" scoped>
ul {
    list-style-type: none;
}

li {
    display: inline-block;
}

input[type="radio"][id^="cb"] {
    display: none;
}

label {
    border: 1px solid #fff;
    padding: 10px;
    display: block;
    position: relative;
    margin: 10px;
    cursor: pointer;
}

label:before {
    background-color: white;
    color: white;
    content: " ";
    display: block;
    border-radius: 50%;
    border: 1px solid white;
    position: absolute;
    top: -5px;
    left: -5px;
    width: 25px;
    height: 25px;
    text-align: center;
    line-height: 28px;
    transition-duration: 0.4s;
    transform: scale(0);
}

label img {
    height: 100px;
    width: 100px;
    transition-duration: 0.2s;
    transform-origin: 50% 50%;
}

:checked+label {
    border-color: #ddd;
}

:checked+label:before {
    content: "✓";
    background-color: #0097A7;
    transform: scale(1);
}

:checked+label img {
    transform: scale(0.9);
    box-shadow: 0 0 5px #333;
    z-index: -1;
}
</style>