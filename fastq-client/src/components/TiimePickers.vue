<template>
  <v-row class="">
    <v-col cols="6" class="">
      <v-menu
        ref="menu"
        v-model="frommenu"
        :close-on-content-click="false"
        transition="scale-transition"
        offset-y
        min-width="auto"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="from"
            label="from date"
            hide-details="auto"
            prepend-icon="mdi-calendar"
            readonly
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker
          v-model="from"
          :active-picker.sync="activePicker"
          hide-details="auto"
          :max="
            new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
              .toISOString()
              .substring(0, 10)
          "
          min="1950-01-01"
          @change="save"
        ></v-date-picker>
      </v-menu>
    </v-col>
    <v-col cols="6" class="">
      <v-menu
        ref="menu"
        v-model="tomenu"
        :close-on-content-click="false"
        transition="scale-transition"
        offset-y
        min-width="auto"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="to"
            label="to date"
            prepend-icon="mdi-calendar"
            hide-details="auto"
            readonly
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker
          v-model="to"
          :active-picker.sync="toactivePicker"
          hide-details="auto"
          :max="
            new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
              .toISOString()
              .substring(0, 10)
          "
          min="1950-01-01"
          @change="save"
        ></v-date-picker>
      </v-menu>
    </v-col>
    <!-- <div class="mb-6">Active picker: <code>{{ activePicker || 'null' }}</code></div> -->
  </v-row>
</template>

<script>
export default {
  data: () => ({
    activePicker: null,
    toactivePicker: null,
    date: null,
    from: null,
    to: null,
    tomenu: false,
    frommenu: false,
  }),
  watch: {
    menu(val) {
      val && setTimeout(() => (this.activePicker = "YEAR"));
    },
  },
  methods: {
    save(date) {
      let time = {
        to: this.to,
        from: this.from,
      };
      this.$eventBus.$emit("time-selected", time);
      this.$refs.menu.save(date);
    },
  },
};
</script>