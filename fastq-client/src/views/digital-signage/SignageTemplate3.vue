<template>
    <v-app>
        <Header />
        <v-main>
            <div class="row fill-height">
                <div class="col-md-12 primary h-60 d-flex align-items-center justify-content-center">
                    <div class="row col-md-10 h-100 mt-5">
                        <h2 class="text-white h3">NOW SERVING</h2>
                        <v-card width="100%" class="d-flex flex-column mt-n8 justify-content-center" height="70%" :style="{'border-radius':'24px'}" color="white">
                          <div class="row ">
                            <div class="col-md-4">
                                <h4>SERVICE</h4>
                            </div>
                            <div class="col-md-4">
                                <h4>COUNTER</h4>
                            </div>
                            <div class="col-md-4">
                                <h4>TICKET</h4>
                            </div>
                          </div>
                          <div class="row ">
                            <div class="col-md-4">
                                <h3 style="font-size: 4rem;">PAYMENT</h3>
                            </div>
                            <div class="col-md-4">
                                <h3 style="font-size: 4rem;">1</h3>
                            </div>
                            <div class="col-md-4">
                                <h3 style="font-size: 4rem;">2002</h3>
                            </div>
                          </div>
                        </v-card>
                    </div>
                </div>
                <div class="col-md-12 black h-40 d-flex align-items-center justify-content-center">
                    <div class="row col-md-10 h-100 text-white mt-5">
                        <h2 class="text-white h3">PREVIOUSLY CALLED</h2>
                        <v-card width="100%" class="row d-flex align-items-start justify-content-around" height="80%" color="black">
                            <v-card class="col-md-2 p-3" shaped>
                             <v-card-title primary-title>
                                <h2>1-2003</h2>
                             </v-card-title>
                           </v-card>
                           <v-card class="col-md-2 p-3" shaped>
                             <v-card-title primary-title>
                                <h2>1-2003</h2>
                             </v-card-title>
                           </v-card>
                           <v-card class="col-md-2 p-3" shaped>
                             <v-card-title primary-title>
                                <h2>1-2003</h2>
                             </v-card-title>
                           </v-card>
                           <v-card class="col-md-2 p-3" shaped>
                             <v-card-title primary-title>
                                <h2>1-2003</h2>
                             </v-card-title>
                           </v-card>
                           <v-card class="col-md-2 p-3 text-center" shaped>
                             <v-card-title primary-title>
                                <h2>1-2003</h2>
                             </v-card-title>
                           </v-card>
                        </v-card>
                    </div>

                </div>
            </div>
        </v-main>
        <v-footer fixed height="64px" tag="footer">

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