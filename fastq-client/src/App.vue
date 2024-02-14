<template>
  <div class="">
    <loader v-if="this.$store.state.Auth.loader" object="#0097A7" color1="#ffffff" color2="#AD1457" size="5" speed="2"
      bg="#343a40" objectbg="#999793" opacity="80" name="circular"></loader>
    <router-view />

  </div>
</template>

<script>
export default {
  name: "App",

  data: () => ({
    //
  }),
};
</script>

<style lang="scss">
@import url('https://fonts.googleapis.com/css2?family=Ubuntu:wght@300;400;500;700&display=swap');
* {
  font-family: "Ubuntu";
  text-align: left;
}

.container-manual {
  min-height: calc(100vh - 160px);
}
</style>
