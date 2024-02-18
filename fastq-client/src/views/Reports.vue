<template>
    <div class="container">
        <!-- <h1>Reports Page</h1> -->
        <!-- <TiimePickers /> -->
        <v-row class="d-flex align-center justify-center">
            <v-col cols="12" md="12" lg="12">
                <h2 class="grey--text"> Ticket Count by status</h2>
            </v-col>
            <v-col v-for="(value, key) in queriedData" :key="key" cols="12" md="4" lg="3">
                <v-card :color="colors[Math.floor((Math.random() * 3))]" elevation="3">
                    <v-card-title class="white--text text-h2">{{ value }}</v-card-title>
                    <v-card-subtitle class="white--text text-h6">{{ key }}</v-card-subtitle>
                    <v-card-text>
                        <v-row>
                            <!-- <v-col v-for="(value, key) in data" :key="key" cols="12">
                                <v-card-number>{{ value }}</v-card-number>
                            </v-col> -->
                        </v-row>
                    </v-card-text>
                </v-card>
            </v-col>
            <v-col cols="12" md="12" lg="12">
                <h2 class="grey--text"> Ticket Count by service</h2>
            </v-col>
            <v-col v-for="(value, key) in ticketbyservice" :key="key" cols="12" md="4" lg="3">
                <v-card :color="colors[Math.floor((Math.random() * 3))]" elevation="3">
                    <v-card-title class="white--text text-h2">{{ value }}</v-card-title>
                    <v-card-subtitle class="white--text text-h6">{{ key }}</v-card-subtitle>
                    <v-card-text>
                        <v-row>
                            <!-- <v-col v-for="(value, key) in data" :key="key" cols="12">
                                <v-card-number>{{ value }}</v-card-number>
                            </v-col> -->
                        </v-row>
                    </v-card-text>
                </v-card>
            </v-col>
            <v-col cols="12" md="12" lg="12" class="d-flex justify-content-between">
                <h2 class="grey--text"> Tickets Served By User</h2>
                <v-btn large text color="success" class="mt-1 text-md-button" @click="exportToExcel(userData, 'tickets_served'
                )"> <v-icon left>mdi-export</v-icon>
                    Export Excel
                </v-btn>
            </v-col>
            <v-col cols="12" md="12" lg="12">
                <DataTables :headers="tableheaders" :items="userData" />
            </v-col>
            <v-col cols="12" md="12" lg="12" class="d-flex justify-content-between">
                <h2 class="grey--text"> Active Time Of User Per Day</h2>
                <v-btn large text color="success" class="mt-1 text-md-button" @click="exportToExcel(activeData, 'active_time_user_per_day'
                )"> <v-icon left>mdi-export</v-icon>
                    Export Excel
                </v-btn>
            </v-col>
            <v-col cols="12" md="12" lg="12">
                <DataTables :headers="activeheaders" :items="activeData" />
            </v-col>
        </v-row>
    </div>
</template>
    
<script>
import TiimePickers from '@/components/TiimePickers.vue';
import axios from 'axios';
import DataTables from './tables/dataTables.vue';
import * as XLSX from 'xlsx';
export default {
    components: { TiimePickers, DataTables },
    data: () => ({
        headers: null,
        queriedData: null,
        ticketbyservice: null,
        colors: ["#0097A7", "#FF4081", "#00C853"],
        avgWaitingTime: null,
        avgWaitingTimeByService: null,
        tableheaders: [
            {
                text: 'CounterName',
                align: 'start',
                sortable: true,
                value: 'counter_name',
            },
            { text: 'BranchName', value: 'company_name' },
            { text: 'Service', value: 'service' },
            { text: 'TicketName', value: 'ticket_name' },
            { text: 'TicketStatus', value: 'ticket_status' },
            { text: 'UserName', value: 'user_name' },
            { text: 'TimeTaken (seconds)', align: 'center', value: 'time_taken' },
            // { text: 'CreatedAt', value: 'date' },

            { text: 'Actions', value: 'actions', sortable: false },
        ],
        activeheaders: [
            {
                text: 'UserName',
                align: 'start',
                sortable: false,
                value: 'UserName',
            },
            {
                text: 'ActiveTimeHours',
                value: 'ActiveTimeHours',
            },
            { text: 'ActiveTimeMinutes', value: 'ActiveTimeMinutes' },
            { text: 'ActiveTimeSeconds', value: 'ActiveTimeSeconds' },
            { text: 'FirstTicketTime', value: 'FirstTicketTime' },
            { text: 'LastTicketTime', value: 'LastTicketTime' },
            { text: 'ServingDate', value: 'ServingDate' },
            // { text: 'CreatedAt', value: 'date' },

            { text: 'Actions', value: 'actions', sortable: false },
        ],
        counterTableData: [],
        userData: [],
        activeData: [],
    }),
    watch: {

    },
    methods: {
        exportToExcel(data, name) {
            if (data.length <= 0) {
                this.$toast.error("no data to export!!")
                return
            }
            // Your JSON data
            // const jsonData = [
            //     { Name: 'John Doe', Age: 30, City: 'New York' },
            //     { Name: 'Jane Smith', Age: 25, City: 'Los Angeles' },
            //     // Add more rows as needed
            // ];

            // Create a worksheet
            const ws = XLSX.utils.json_to_sheet(data);

            // Create a workbook
            const wb = XLSX.utils.book_new();
            XLSX.utils.book_append_sheet(wb, ws, 'Sheet 1');

            // Save the workbook as an Excel file
            XLSX.writeFile(wb, `${name}.xlsx`);
        },
        async getTotalTicketsByService() {
            try {
                const response = await axios.get(`/report/tickets-by-service/${this.$store.state.Auth.user.company_code}`);
                // console.log("object", response.data);
                this.ticketbyservice = response.data
                // this.handleQueryResponse(response);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
        async getTicketsByStatus() {
            try {
                const response = await axios.get(`/report/tickets-by-status/${this.$store.state.Auth.user.company_code}`);
                this.handleQueryResponse(response);
                console.log("object", response.data);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },

        handleQueryResponse(response) {
            // Customize this method based on your data structure
            if (response.data) {
                // console.log(">>>>>>>>>>>>>>>", response.data);
                // this.headers = Object.keys(response.data[0]);
                this.queriedData = response.data;
                this.showData = true;
            } else {
                console.error('Invalid response data structure');
            }
        },
        async getTicketgetAverageServiceTimeByService() {
            try {
                const response = await axios.get(`/report/avg-service-time-by-service/${this.$store.state.Auth.user.company_code}`);
                // this.handleQueryResponse(response);
                console.log("object", response.data);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
        async getTicketsAverageServiceTimeByService() {
            try {
                const response = await axios.get(`/report/tickets-with-avg-service-time-by-counter/${this.$store.state.Auth.user.company_code}`);
                // this.handleQueryResponse(response);
                console.log("object with avg service time", response.data);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
        async getTicketgetAverageWatingTimeByService() {
            try {
                const response = await axios.get(`/report/avg-waiting-time-by-service/${this.$store.state.Auth.user.company_code}`);
                // this.handleQueryResponse(response);
                console.log("object", response.data);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
        async getTicketsByComanyBranch() {
            try {
                const response = await axios.get(`/report/tickets-by-company-branch/${this.$store.state.Auth.user.company_code}`);
                // this.handleQueryResponse(response);
                console.log("object company", response.data);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
        async getTicketServedByCounter() {
            try {
                const response = await axios.get(`/report/tickets-served-by-counter/${this.$store.state.Auth.user.company_code}`);
                // this.handleQueryResponse(response);
                console.log("object served by counter", response.data);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
        async getTicketswithUserInfo() {
            try {
                const response = await axios.get(`/report/tickets-with-user-info/${this.$store.state.Auth.user.company_code}`);
                // this.handleQueryResponse(response);
                this.userData = response.data
                console.log("object with user info", response.data);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
        async getTicketWithCounterInfo() {
            try {
                const response = await axios.get(`/report/tickets-with-counter-info/${this.$store.state.Auth.user.company_code}`);
                // this.handleQueryResponse(response);
                console.log("object with counter info", response.data);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
        async getActiveTimePerDay() {
            try {
                const response = await axios.get(`/report/active-time-of-user-per-day/${this.$store.state.Auth.user.company_code}`);
                // this.handleQueryResponse(response);
                this.activeData = response.data
                console.log("object active time of user per day", response.data);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
        async getTicketCounterUserInfo() {
            try {
                const response = await axios.get(`/report/tickets-by-company-branch-counter-info/${this.$store.state.Auth.user.company_code}`);
                // this.handleQueryResponse(response);
                console.log("object", response.data);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },

    },
    created() {
        this.getTicketsByStatus()
        this.getTotalTicketsByService()
        this.getTicketgetAverageServiceTimeByService()
        this.getTicketgetAverageWatingTimeByService()
        this.getTicketServedByCounter()
        this.getTicketsByComanyBranch()
        this.getTicketServedByCounter()
        this.getTicketswithUserInfo()
        this.getTicketWithCounterInfo()
        this.getActiveTimePerDay()
        this.getTicketsAverageServiceTimeByService()
    }
}

</script>
    
<style></style>