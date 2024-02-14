<template>
    <v-card flat class="mt-5" fill-height >
        <div class="example">
            <div id="chart" class="w-100">
                <apexchart type="bar" height="300" class="w-100" :options="chartOptions" :series="series"
                    v-if="seriesdata.length > 0"></apexchart>
            </div>
        </div>
    </v-card>
</template>
  
<script>
export default {
    name: 'BarExample',
    data() {
        return {
            series: [{
                name: 'Ticket Counts',
                data: this.seriesdata
            }],
            chartOptions: {
                labels: this.legend,
                colors: ["#00C853"],
                chart: {
                    fontFamily: 'Ubuntu, Arial, sans-serif',
                    height: 'auto',
                    type: 'bar',
                },
                plotOptions: {
                    bar: {
                        borderRadius: 10,
                        dataLabels: {
                            position: 'top', // top, center, bottom
                        },
                    }
                },
                dataLabels: {
                    enabled: true,
                    formatter: function (val) {
                        return val;
                    },
                    offsetY: -20,
                    style: {
                        fontSize: '12px',
                        colors: ["#304758"]
                    }
                },

                xaxis: {
                    categories: this.legend,
                    position: 'top',
                    axisBorder: {
                        show: false
                    },
                    axisTicks: {
                        show: false
                    },
                    crosshairs: {
                        fill: {
                            type: 'gradient',
                            gradient: {
                                colorFrom: '#0097A7',
                                colorTo: '#0097A7',
                                stops: [0, 100],
                                opacityFrom: 0.4,
                                opacityTo: 0.5,
                            }
                        }
                    },
                    tooltip: {
                        enabled: true,
                    }
                },
                yaxis: {
                    axisBorder: {
                        show: false
                    },
                    axisTicks: {
                        show: false,
                    },
                    labels: {
                        show: false,
                        formatter: function (val) {
                            return val + "%";
                        }
                    }

                },
                title: {
                    text: 'Total ticket counts by status',
                    // floating: true,
                    margin: 10,
                    // offsetY: 330
                    align: 'center',
                    style: {
                        padding: "10px",
                        color: '#444'
                    }
                }
            },
        }
    },
    props: {
        seriesdata: {
            type: Array,
            default: () => [],
        },
        legend: {
            type: Array,
            default: () => [],
        },
    },
    watch: {

        // legend: function () {
        //     this.chartOptions.labels = this.legend;
        // },
    },
    methods: {

    },
    created() {
        console.log("object>>>>>>>>>", this.seriesdata);
    }
}
</script>
  
<style lang="scss" scoped>
.btn-block {
    display: block;
    min-width: 100%;
}

.primary-gradient {
    background: linear-gradient(to right, #da8cff, #9a55ff);
    border: none;
    outline: none;
}
</style>