<template>
    <div class="container">
        <!-- <v-toolbar flat color="" app class="d-flex align-items-center mt-n5">
            <v-toolbar-title>
                <h2 class="grey--text">{{ $route.name }}</h2>
            </v-toolbar-title>
        </v-toolbar> -->

        <!-- <div class="col-md-12">
                <v-select v-model="select" :items="items" :error-messages="errors" label="Select" data-vv-name="select"
                    required pa-3></v-select>
            </div> -->
        <!-- </div> -->
        <!-- <TiimePickers /> -->
        <v-card width="100%" flat class="d-flex flex-wrap align-items-start justify-content-around py-3" fill-height>
            <!-- <div class="col-md-12 px-0 d-flex align-items-center justify-content-between">
                <div class="col-md-6 d-flex align-items-center justify-content-between ">
                    <h5 class="mr-2 grey--text">Show </h5>
                    <v-select v-model="select" :items="items" :error-messages="errors" label="Select" data-vv-name="select"
                        required pa-3></v-select>
                    <v-chip color="cyan" text-color="white" label link
                        :class="{ 'd-none': $vuetify.breakpoint.mdAndDown, 'ma-2': true }">
                        <v-icon left color="white">
                            mdi-monitor
                        </v-icon>
                        Activity Monitoring Screen
                    </v-chip>
                </div>
                <div class="col-md-4 d-flex align-items-center">
                    <p class=" h5 grey--text">12-05-2023 to 19-05-2023 </p>
                  
                </div>
            </div> -->
            <div class="card-wrapper" v-for="(spark, index) in sparkObj">
                <DashCards :color="colors[index]" :sparkObj="spark" />
            </div>
            <div class="row col-md-12 align-items-center ">
                <div class="col-md-6">
                    <ChartCards :seriesdata="barData" :legend="barlabel" />
                </div>
                <div class="col-md-5">
                    <donutChart :data="donutData" :legend="labels" />
                </div>
            </div>
            <div class="row col-md-12 align-items-center ">
                <div class="col-md-12">
                    <LineChart :seriesdata="chartSeries" v-if="chartSeries[0].data.length > 0" :legend="lineLable" />
                </div>
                <!-- <div class="col-md-6">
                    <AreaChart />
                </div> -->
            </div>
        </v-card>
    </div>
</template>

<script>
import DashCards from './cards/dashCards.vue';
import ChartCards from './cards/chartCards.vue';
import DonutChart from './cards/donutChart.vue';
import LineChart from './cards/lineChart.vue';
import AreaChart from './cards/areaChart.vue';
import TiimePickers from '@/components/TiimePickers.vue';
import axios from 'axios';
import moment from 'moment';

export default {
    components: {
        DashCards,
        ChartCards,
        DonutChart,
        LineChart,
        AreaChart,
        TiimePickers
    },
    data() {
        return {
            colors: ["#0097A7", "#FF4081", "#00C853"],
            selectedFile: null,
            labels: [],
            donutData: [],
            barData: [],
            barlabel: [],
            lineLable: [],
            createdhours: [],
            createdData: [],
            finishedHours: [],
            finishedData: [],
            noshowHours: [],
            noshowData: [],
            sparkObj: [],
            chartSeries: [
                {
                    name: 'Finished Tickets',
                    data: [],
                },
                {
                    name: 'Total Tickets',
                    data: [],
                },
            ],
        };
    },
    methods: {
        async getTotalTicketsByService() {
            try {
                const response = await axios.get(`/report/tickets-by-service/${this.$store.state.Auth.user.company_code}`);
                // console.log("object", response.data);
                for (const [key, value] of Object.entries(response.data)) {

                    console.log(key, value);
                    this.donutData.push(value)
                    this.labels.push(key)
                }
                // this.labels = Object.keys(response.data)

                // this.handleQueryResponse(response);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
        async getTicketsByStatus() {
            try {
                const response = await axios.get(`/report/tickets-by-status/${this.$store.state.Auth.user.company_code}`);
                for (const [key, value] of Object.entries(response.data)) {
                    // console.log(key, value);
                    this.barData.push(value)
                    this.barlabel.push(key)
                }
                // console.log("object", response.data);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
        async getHourlyCreated() {
            try {
                const response = await axios.get(`/report/ticket-today-created/${this.$store.state.Auth.user.company_code}`);
                for (let i = 0; i < response.data.length; i++) {
                    this.createdhours.push(response.data[i].hour.toString())
                    this.createdData.push(response.data[i].createdCount)
                }
                if (this.createdData.length === 1) {
                    this.createdData.push(0)
                }
                let spark = {
                    name: "Tickets Created on Hourly Basis",
                    lastFinished: `last ticket created on ${this.createdhours[this.createdhours.length - 1] - 12}`,
                    countdata: this.createdData,
                    labels: this.createdhours

                }
                this.sparkObj.push(spark)
                console.log("object hourly created", this.createdData, this.createdhours);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
        async getHourlyNoShow() {
            try {
                const response = await axios.get(`/report/ticket-today-noshow/${this.$store.state.Auth.user.company_code}`);
                for (let i = 0; i < response.data.length; i++) {
                    this.noshowHours.push(response.data[i].hour.toString())
                    this.noshowData.push(response.data[i].noshowCount)
                }
                if (this.noshowData.length === 1) {
                    this.noshowData.push(0)
                }
                let spark = {
                    name: "Tickets NoShow on Hourly Basis",
                    lastFinished: `last ticket finished on ${this.noshowHours[this.noshowHours.length - 1] - 12}`,
                    countdata: this.noshowData,
                    labels: this.noshowHours

                }
                this.sparkObj.push(spark)
                console.log("object hourly noshow", this.noshowData, this.noshowHours);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
        async getHourlyfinished() {
            try {
                const response = await axios.get(`/report/ticket-today-finished/${this.$store.state.Auth.user.company_code}`);
                for (let i = 0; i < response.data.length; i++) {
                    this.finishedHours.push(response.data[i].hour.toString())
                    this.finishedData.push(response.data[i].finishedCount)
                }
                if (this.finishedData.length === 1) {
                    this.finishedData.push(0)
                }
                let spark = {
                    name: "Tickets Finished on Hourly Basis",
                    lastFinished: `last ticket finished on ${this.finishedHours[this.finishedHours.length - 1] - 12}`,
                    countdata: this.finishedData,
                    labels: this.finishedHours

                }
                this.sparkObj.push(spark)
                console.log("object hourly finished", this.finishedData, this.finishedHours);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
        async fetchTicketStats() {
            try {
                const response = await axios.get(`/report/ticket-stats/${this.$store.state.Auth.user.company_code}`);
                const stats = response.data;

                // Extract data for chart
                const categories = stats.map((stat) => moment(stat.date).format('DD MMMM YYYY'));
                const finishedTickets = stats.map((stat) => stat.finishedTickets);
                const totalTickets = stats.map((stat) => stat.totalTickets);

                // Update chart data
                this.lineLable = categories;
                this.chartSeries[0].data = finishedTickets;
                this.chartSeries[1].data = totalTickets;
                console.log(" dataes", this.lineLable);
                console.log(" dataes", this.chartSeries);

            } catch (error) {
                console.error('Error fetching ticket stats:', error);
            }
        },
    },
    created() {
        this.getTotalTicketsByService()
        this.getTicketsByStatus()
        this.fetchTicketStats()
        this.getHourlyCreated()
        this.getHourlyfinished()
        this.getHourlyNoShow()
    }
}
</script>

<style lang="sass" scoped>

</style>