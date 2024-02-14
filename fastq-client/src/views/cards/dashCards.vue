<template>
    <v-card class="mt-4 mx-auto" min-width="350">
        <v-sheet class="v-sheet--offset mx-auto" :color="color" elevation="12" max-width="calc(100% - 32px)">
            <!-- {{ sparkObj.labels }} ===== {{ sparkObj.countdata }} -->
            <v-sparkline :labels="sparkObj.labels" :value="sparkObj.countdata" :key="sparkObj.name" color="white" line-width="2"
                padding="16"></v-sparkline>
        </v-sheet>

        <v-card-text class="pt-0">
            <div class="text-h6 font-weight-light mb-2">
                {{ sparkObj.name }} - {{ sparkObj.countdata.reduce((partialSum, a) => partialSum + a, 0) }}
            </div>
            <div class="subheading font-weight-light grey--text">
                Last Campaign Performance
            </div>
            <v-divider class="my-2"></v-divider>
            <v-icon class="mr-2" small>
                mdi-clock
            </v-icon>
            <span class="text-caption grey--text font-weight-light">{{ sparkObj.lastFinished }}</span>
        </v-card-text>
    </v-card>
</template>

<script>

export default {
    props: ["color", "sparkObj"],
    data() {
        return {
            labels: [
                '12am',
                '3am',
                '6am',
                '9am',
                '12pm',
                '3pm',
                '6pm',
                '9pm',
            ],
            value: [
                200,
                675,
                410,
                390,
                310,
                460,
                250,
                240,
            ],
        }
    },

}

</script>

<style lang="scss" scoped>
.v-sheet--offset {
    top: -24px;
    position: relative;
}
</style>