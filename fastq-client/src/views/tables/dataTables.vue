<template>
  <v-card flat>
    <v-card-title>
      <v-text-field
        v-model="search"
        append-icon="mdi-magnify"
        label="Search"
        single-line
        hide-details
      ></v-text-field>
    </v-card-title>
    <v-data-table
      :headers="headers"
      :items="items"
      :items-per-page="5"
      mobile
      :search="search"
    >
      <template v-slot:item.actions="{ item }">
        <!-- <v-btn
          icon
          id="no-background-hover"
          :to="`/edit/${name}/${item.id}`"
          v-if="name !== 'counterservice'"
        >
          <v-icon small class="mr-2" color="info" @click="editItem(item)">
            mdi-pencil
          </v-icon>
        </v-btn> -->
        <v-icon small color="error" @click="show = true"> mdi-delete </v-icon>
        <v-dialog v-model="show" max-width="500px">
          <v-card>
            <v-card-title class="text-h5"
              >Are you sure you want to delete this item?</v-card-title
            >
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="green darken-1" text @click="show = false"
                >Cancel</v-btn
              >
              <v-btn color="red darken-1" text @click="deleteImage(item)"
                >OK</v-btn
              >
              <v-spacer></v-spacer>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
import axios from 'axios';
export default {
  props: ["headers", "items"],
  data() {
    return {
      search: "",
      show: false,
    };
  },
  methods: {
    deleteImage(item) {
      axios
        .get(`/config/deleteuploaded/${item.id}`)
        .then((res) => {
          console.log("user", res.data);
          this.show = false;
            window.location.reload();
          this.$toast.success("item deleted successfully!!");
        })
        .catch((err) => {
          this.$toast.error("error deleting service!!", err);
        });
    },
  },
};
</script>

<style lang="scss" scoped></style>