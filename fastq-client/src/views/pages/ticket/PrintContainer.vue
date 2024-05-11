<template>
    <div class="">
      <component
        :is="'TicketTemplateOne'"
        :ip="ipaddress"
        :company="companycode"
        :key="1"
      ></component>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  import TicketTemplateOne from "./TicketTemplateOne.vue";
  import * as types from "../../../store/types";
  // import TemplateOne from "./SignageTemplate1.vue";
  export default {
    components: {
        TicketTemplateOne,
    },
    data() {
      return {
        companycode: null,
        ipaddress: null,
      };
    },
    methods: {
      getcustomizations() {
        return new Promise((resolve, reject) => {
          // this.id = this.$route.query.company;
          // if (this.id) {
          this.$store.commit(types.MUTATE_LOADER_ON);
          axios
            .get(`/customize/getcustomizations`)
            .then((res) => {
              this.customization = res.data;
              console.log("object,cust", this.customization);
              this.$store.commit(types.MUTATE_LOADER_OFF);
            })
            .catch((err) => {
              console.log(err.response);
              reject(err.response);
              this.$store.commit(types.MUTATE_LOADER_OFF);
              this.$toast.error("error occured while getting company!!!");
            });
          // }
        });
      },
    },
    created() {
      this.getcustomizations();
      this.companycode = this.$route.query.company;
      this.ipaddress = this.$route.query.ip;
    },
  };
  </script>
  
  <style scoped>
  </style>