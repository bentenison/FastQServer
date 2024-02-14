<template>
    <v-app>
        <Header />
        <v-main>
            <div class="row fill-height">
                <div class="col-md-8 black d-flex flex-column align-items-center h-100">
                    <div class="d-flex flex-column h-100 w-100 align-items-center mt-4">
                        <v-card width="99%" height="70%" color="black">
                          
                        </v-card>]
                        <div class="d-flex align-items-center justify-content-around row w-100">
                            <v-card class="col-md-2 p-3 text-white" color="grey darken-4">
                                <h2 style="font-size: 1rem;">PREVIOUSLY
                                    CALLED</h2>
                            </v-card>
                            <v-card class="col-md-2 p-3 text-white" color="grey darken-4">
                                <h2 style="font-size: 2rem;">2-1001</h2>
                            </v-card>
                            <v-card class="col-md-2 p-3 text-white" color="grey darken-4">
                                <h2 style="font-size: 2rem;">2-1001</h2>
                            </v-card>
                            <v-card class="col-md-2 p-3 text-white" color="grey darken-4">
                                <h2 style="font-size: 2rem;">2-1001</h2>
                            </v-card>
                            <v-card class="col-md-2 p-3 text-white" color="grey darken-4">
                                <h2 style="font-size: 2rem;">2-1001</h2>
                            </v-card>
                        </div>
                    </div>

                </div>
                <div class="col-md-4 primary text-white align-content-center">
                    <h4 class="text-center">NOW SERVING</h4>
                    <div class="col-md-12 d-flex flex-column align-items-center justify-content-between">
                        <div class="row col-md-12 d-flex justify-content-center align-items-center">
                          <div class="col-md-6"></div>
                          <div class="col-md-3"><h5>COUNTER</h5></div>
                          <div class="col-md-3"><h5>TICKET</h5></div>
                        </div>
                        <v-card class="row col-md-12 d-flex justify-content-center align-items-center white text-black my-3" shaped>
                          <div class="col-md-6"><h1>PAYMENT</h1></div>
                          <div class="col-md-3"><h1>1</h1></div>
                          <div class="col-md-3"><h1>1000</h1></div>
                        </v-card>
                        <v-card class="row col-md-12 d-flex justify-content-center align-items-center white text-black my-3" shaped>
                          <div class="col-md-6"><h1>PAYMENT</h1></div>
                          <div class="col-md-3"><h1>1</h1></div>
                          <div class="col-md-3"><h1>1000</h1></div>
                        </v-card>
                        <v-card class="row col-md-12 d-flex justify-content-center align-items-center white text-black my-3" shaped>
                          <div class="col-md-6"><h1>PAYMENT</h1></div>
                          <div class="col-md-3"><h1>1</h1></div>
                          <div class="col-md-3"><h1>1000</h1></div>
                        </v-card>
                        <v-card class="row col-md-12 d-flex justify-content-center align-items-center white text-black my-3" shaped>
                          <div class="col-md-6"><h1>PAYMENT</h1></div>
                          <div class="col-md-3"><h1>1</h1></div>
                          <div class="col-md-3"><h1>1000</h1></div>
                        </v-card>
                    </div>

                </div>
            </div>
        </v-main>


        <v-footer fixed bottoms height="64px" elevation="4" padless>

        </v-footer>

    </v-app>
</template>

<script>
import Header from "./Header.vue"
export default {
    components: {
        Header
    }
}
</script>

<style lang="scss" scoped></style>