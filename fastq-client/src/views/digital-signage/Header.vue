<template>
    <v-app-bar elevation="4" app height="80px">
        <div class="d-flex align-items-center justify-content-center p-3">
            <h2 class="m-0">
                {{ $store.state.Auth.user.company }}
            </h2>
            <v-chip class="ma-2" color="primary" label text-color="white">
                <v-icon left>
                    mdi-label
                </v-icon>
                <strong>HEAD OFFICE</strong>
            </v-chip>
        </div>
        <v-spacer></v-spacer>
        <div class="d-md-flex d-sm-none d-xs-none align-items-center justify-content-center flex-wrap p-3">
            <v-icon>
                mdi-calendar
            </v-icon>
            <v-chip class="ma-2" label text-color="grey darken-2">
                <h4 class="font-weight-regular">{{ m.format('ddd DD-MMM-YYYY') }}</h4>
            </v-chip>
            <v-icon>
                mdi-clock
            </v-icon>
            <v-chip class="ma-2" label text-color="grey darken-2">
                <h4 class="font-weight-regular">{{ m.format('HH:mm A') }}</h4>
            </v-chip>
        </div>
        <div class="d-flex align-items-center justify-content-center p-3">
            <v-btn fab small color="primary" class="mx-2">
                <v-icon color="white">mdi-cloud</v-icon>
            </v-btn>
        </div>
    </v-app-bar>
</template>

<script>
import moment from 'moment';

export default {
    data() {
        return {
            m: moment()
        }
    },

}
</script>

<style lang="scss" scoped></style>