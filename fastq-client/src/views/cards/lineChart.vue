<template>
    <v-card class="mt-4 mx-auto w-100" flat>
        <div id="chart">
            <apexchart type="line" height="300" :options="chartOptions" :series="series" >
            </apexchart>
        </div>
    </v-card>
</template>

<script>
export default {
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
    data() {
        return {
            // series: [{
            //     name: "0",
            //     data: this.seriesdata[0].data
            // },
            // {
            //     name: "1",
            //     data: this.seriesdata[1].data
            // }],
            series: this.seriesdata,
            chartOptions: {
                chart: {
                    fontFamily: 'Ubuntu, Arial, sans-serif',
                    height: 350,
                    type: 'line',
                    dropShadow: {
                        enabled: true,
                        color: '#000',
                        top: 18,
                        left: 7,
                        blur: 10,
                        opacity: 0.2
                    },
                    toolbar: {
                        show: false
                    }
                },
                colors: ["#FF4081", '#545454'],
                dataLabels: {
                    enabled: true,
                },
                stroke: {
                    curve: 'smooth'
                },
                title: {
                    text: 'Created and Finished tickets per day',
                    align: 'center'
                },
                grid: {
                    borderColor: '#e7e7e7',
                    row: {
                        colors: ['#f3f3f3', 'transparent'], // takes an array which will be repeated on columns
                        opacity: 0.5
                    },
                },
                markers: {
                    size: 1
                },
                xaxis: {
                    categories: this.legend,
                    title: {
                        text: 'Date'
                    }
                },
                yaxis: {
                    title: {
                        text: 'Counts'
                    },
                    min: 5,
                    max: 40
                },
                legend: {
                    position: 'top',
                    horizontalAlign: 'right',
                    floating: true,
                    offsetY: -25,
                    offsetX: -5
                }
            },

        }
    },
    watch: {

        // legend: function () {
        //     this.chartOptions.labels = this.legend;
        // },
    },
    created() {
        console.log("seriesData",this.seriesdata);
    }
}
</script>

<style lang="scss" scoped></style>