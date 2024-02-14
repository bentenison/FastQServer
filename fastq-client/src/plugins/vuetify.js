import Vue from "vue";
import Vuetify from "vuetify/lib/framework";

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
      themes: {
        light: {
          primary: '#0097A7',
          error: '#FF5252',
          success:"#00c853"
        },
      },
    },
  });
