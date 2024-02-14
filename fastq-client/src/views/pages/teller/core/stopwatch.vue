<template>
    <div :class="time <= 60 ? 'text-success' : 'text-danger'">{{ time | secondsInMinutes }}</div>
</template>
  
<script>
import moment from "moment";

export default {
    name: "stopwatch",
    props: {
        running: {
            type: Boolean,
            default: false
        },
        resetWhenStart: {
            type: Boolean,
            default: false
        },
        restWhenStop: {
            type: Boolean,
            default: false
        },
        starter: {
            type: Number,
            default: 600
        },
    },
    watch: {
        running: function (newVal, oldVal) {
            if (newVal) this.startT();
            else this.stopT();
        }
    },
    mounted() {
        if (this.running) this.startT();
    },
    data() {
        return { time: 0, timer: null };
    },
    methods: {
        stopT: function () {
            clearInterval(this.timer);
            if (this.restWhenStop) this.resetT();
        },
        startT: function () {
            if (this.resetWhenStart) this.resetT();
            if (this.starter) this.time=this.starter
            this.timer = setInterval(() => {
                this.time++;
            }, 1000);
        },
        resetT: function () {
            this.time = 0;
        }
    },
    filters: {
        secondsInMinutes: function (seconds) {
            return moment("2015-01-01")
                .startOf("day")
                .seconds(seconds)
                .format("HH:mm:ss");
        }
    }
};
</script>
  
<style scoped></style>