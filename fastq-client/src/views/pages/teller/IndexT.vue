<template>
    <v-app>
     <AppBarT></AppBarT>
    <MainView></MainView>
    </v-app>
</template>

<script>
import AppBarT from "./core/AppBarT.vue"
import Drawer from "../../core/Drawer.vue"
import MainView from "../../core/MainView.vue"
    export default {
        components:{
            AppBarT,
            Drawer,
            MainView
        },
        data() {
            return {
                
            }
        },  
    }
</script>

<style lang="sass" scoped>

</style>