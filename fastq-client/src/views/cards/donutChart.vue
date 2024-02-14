<template>
    <v-card flat fill-height >
        <div class="example">
            <div id="chart" class="p-4">
                <apexchart type="donut" width="450" :options="chartOptions" :series="series"></apexchart>
            </div>
            <!-- <div class="actions">
                <b-btn class="btn primary-gradient" @click="appendData">
                    + ADD
                </b-btn>
    
                <b-btn class="btn danger-gradient" @click="removeData">
                    - REMOVE
                </b-btn>
    
                <b-btn class="btn secondary-gradient " @click="randomize">
                    RANDOMIZE
                </b-btn>
    
                <b-btn class="btn warning-gradient" @click="reset">
                    RESET
                </b-btn>
            </div> -->
        </div>
    </v-card>
</template>

<script>
export default {
    data() {
        return {
            series: this.data,
            chartOptions: {
                labels: this.legend,
                chart: {
                    fontFamily: 'Ubuntu, Arial, sans-serif',
                    width: 420,
                    type: 'donut',
                },
                plotOptions: {
                    pie: {
                        startAngle: -90,
                        endAngle: 270
                    }
                },
                dataLabels: {
                    enabled: false
                },
                fill: {
                    type: 'gradient',
                },
                colors: ["#0097A7", "#FF4081", "#00C853"],
                legend: {
                    show: false,
                },
                legend: {
                    formatter: function (val, opts) {
                        return val + " - " + opts.w.globals.series[opts.seriesIndex]
                    }
                },
                title: {
                    align: 'center',
                    text: 'Service Wise Ticket Count',
                    margin: 10,
                    offsetY: -10,
                    align: 'center',
                    style: {
                        padding: "30px",
                        color: '#444'
                    }
                },
                responsive: [{
                    breakpoint: 480,
                    options: {
                        chart: {
                            width: 200
                        },
                        legend: {
                            position: 'bottom'
                        }
                    }
                }]
            },
        }
    },
    props: {
        data: {
            type: Array,
            default: () => [],
        },
        legend: {
            type: Array,
            default: () => [],
        },
    },
    watch: {
        data: function () {
            this.chartSeries = this.data;
        },
        legend: function () {
            this.chartOptions.labels = this.legend;
        },
    },
    methods: {
        // appendData: function () {
        //     var arr = this.series.slice()
        //     arr.push(Math.floor(Math.random() * (100 - 1 + 1)) + 1)
        //     this.series = arr
        // },

        // removeData: function () {
        //     if (this.series.length === 1) return
        //     var arr = this.series.slice()
        //     arr.pop()
        //     this.series = arr
        // },

        // randomize: function () {
        //     this.series = this.series.map(function () {
        //         return Math.floor(Math.random() * (100 - 1 + 1)) + 1
        //     })
        // },

        // reset: function () {
        //     this.series = [44, 55, 13, 33]
        // }
    },
}
</script>

<style lang="scss" scoped>
.danger-gradient {
    background: linear-gradient(to right, #ffbf96, #fe7096);
    border: none;
    outline: none;
    color: whitesmoke;
    font-size: 0.75rem;
    font-weight: 500;
}

.primary-gradient {
    background: linear-gradient(to right, #90caf9, #047edf 99%);
    border: none;
    outline: none;
    color: whitesmoke;
    font-size: 0.75rem;
    font-weight: 500;
}

.secondary-gradient {
    background: linear-gradient(to right, #84d9d2, #07cdae);
    color: whitesmoke;
    border: none;
    outline: none;
    font-size: 0.75rem;
    font-weight: 500;
}

.warning-gradient {
    background: linear-gradient(to right, #f6e384, #ffd500);
    border: none;
    outline: none;
    font-size: 0.75rem;
    font-weight: 500;
}
</style>