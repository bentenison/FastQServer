<template>
    <v-app>
     <AppBar></AppBar>
    <Drawer></Drawer>
    <MainView></MainView>
    </v-app>
</template>

<script>
import AppBar from "./core/AppBar.vue"
import Drawer from "./core/Drawer.vue"
import MainView from "./core/MainView.vue"
    export default {
        components:{
            AppBar,
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