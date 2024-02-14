<template>
    <v-app>
        <div>
            <v-toolbar prominent lazy-src transition src="https://picsum.photos/1920/1080" class="pt-2">
                <v-img class="mx-2" src="../../../assets/rsz_al_2.png" max-height="68" max-width="100" contain></v-img>

                <v-toolbar-title class="text-white">{{ $store.state.Auth.user.company }} </v-toolbar-title>

                <v-spacer></v-spacer>
                <!-- <v-toolbar-title> -->

                <!-- <v-toolbar-items class="d-flex"> -->
                <v-menu offset-y origin="center center" class="p-0 m-0" transition="scale-transition">
                    <template class="" v-slot:activator="{ on, attrs }">
                        <p v-bind="attrs" v-on="on" class="text-white">
                            Company Administrator - {{ $store.state.Auth.user.firstname }}{{ $store.state.Auth.user.lastname }}
                            <v-icon color="white">mdi-menu-down</v-icon>
                        </p>
                    </template>
                    <v-list>
                        <v-list-item>
                            <v-list-item-icon>
                                <v-icon>mdi-cog</v-icon>
                            </v-list-item-icon>
                            <v-list-item-title class="px-2">LogOut</v-list-item-title>
                        </v-list-item>
                    </v-list>
                </v-menu>
                <!-- </v-toolbar-items> -->
            </v-toolbar>
        </div>
        <div class="row fill-height">
            <!-- <div class="menu-content d-flex  justify-content-around">
                <div class="row h-100 p-4 w-100 mt-5">
                    <div class="d-flex align-items-center justify-content-around flex-wrap">

                    </div>


                </div>
            </div> -->
            <div class="col-md-12 d-flex flex-wrap align-items-center justify-content-around mt-n16">
                <v-card class="col-md-3 d-flex flex-column align-items-center justify-content-center" flat>
                    <v-icon size="100px" color="primary">mdi-monitor-screenshot</v-icon>
                    <v-card-actions>
                        <!-- <v-chip close-icon="mdi-delete" class="text-button" color="#00838F" link outlined>Get Ticket Screen</v-chip> -->
                        <v-btn text
                            :class="{ 'caption': $vuetify.breakpoint.mdAndDown, 'button': $vuetify.breakpoint.mdAndUp }"
                            color="primary" to="/ticket">Get Ticket Screen</v-btn>
                    </v-card-actions>
                </v-card>
                <v-card class="col-md-3 d-flex flex-column align-items-center justify-content-center" flat>
                    <v-icon size="100px" color="primary">mdi-monitor-dashboard</v-icon>
                    <v-card-actions>
                        <!-- <v-chip close-icon="mdi-delete" class="text-button" color="#00838F" link outlined></v-chip> -->
                        <v-btn text
                            :class="{ 'caption': $vuetify.breakpoint.mdAndDown, 'button': $vuetify.breakpoint.mdAndUp }"
                            color="primary" to="/ds">Customer Display Screen</v-btn>
                    </v-card-actions>
                </v-card>
                <v-card class="col-md-3 d-flex flex-column align-items-center justify-content-center" flat>
                    <v-icon size="100px" color="primary">mdi-ticket-account</v-icon>
                    <v-card-actions>
                        <v-btn text
                            :class="{ 'caption': $vuetify.breakpoint.mdAndDown, 'button': $vuetify.breakpoint.mdAndUp }"
                            color="primary" to="/teller">Teller call Next</v-btn>
                    </v-card-actions>
                </v-card>
                <v-card class="col-md-2 d-flex flex-column align-items-center justify-content-center" flat>
                    <v-icon size="100px" color="primary">mdi-badge-account</v-icon>
                    <v-card-actions>
                        <v-btn text
                            :class="{ 'caption': $vuetify.breakpoint.mdAndDown, 'button': $vuetify.breakpoint.mdAndUp }"
                            color="primary" to="/dashboard">
                            Admin Dashboard
                        </v-btn>
                    </v-card-actions>
                </v-card>
            </div>

        </div>

        <v-footer padless fixed app>
            <v-card class="flex" elevation-10 tile color="#E7E6E4">
                <!-- <v-card-title>
                    <strong class="subheading">Get connected with us on social networks!</strong>

                    <v-spacer></v-spacer>

                    <v-btn v-for="icon in icons" :key="icon" class="mx-4" dark icon>
                        <v-icon size="24px" color="#6A6A69">
                            {{ icon }}
                        </v-icon>
                    </v-btn>
                </v-card-title> -->

                <v-card-text class="py-2 white--text cyan darken-3 text-center">
                    {{ new Date().getFullYear() }} — <strong>FASTQ Solutions limited</strong>
                </v-card-text>
            </v-card>
        </v-footer>
        <!-- </div> -->
    </v-app>
</template>

<script>
export default {
    data: () => ({
        icons: [
            'mdi-facebook',
            'mdi-twitter',
            'mdi-linkedin',
            'mdi-instagram',
        ],
        admins: [
            ['Management', 'mdi-account-multiple-outline'],
            ['Settings', 'mdi-cog-outline'],
        ],
    }),
    mounted() {
        console.log("store", this.$store.state.Auth);
    },
}
</script>

<style lang="scss" scoped>
.menu-content {
    background-color: #E7E6E4;
    min-height: 75vh;
}
</style>