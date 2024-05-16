<template>
  <v-card class="p-2">
    <v-text-field
      v-model="search"
      append-icon="mdi-magnify"
      label="Search"
      single-line
      hide-details
    ></v-text-field>
    <v-data-table :headers="headers" :items="data" mobile :search="search">
      <template v-slot:item.license="{ item }">
        <div v-html="item.license"></div>
      </template>
      <template v-slot:item.actions="{ item }">
        <v-btn
          icon
          id="no-background-hover"
          :to="`/edit/${name}/${item.id}`"
          v-if="name !== 'counterservice'"
        >
          <v-icon small class="mr-2" color="info" @click="editItem(item)">
            mdi-pencil
          </v-icon>
        </v-btn>
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
              <v-btn color="red darken-1" text @click="deleteItem(item)"
                >OK</v-btn
              >
              <v-spacer></v-spacer>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </template>
      <template v-slot:no-data>
        <v-btn color="primary" @click="initialize"> Reset </v-btn>
      </template>
    </v-data-table>
  </v-card>
</template>

<script lang="js">
import axios from 'axios'


export default {
    props: ['data', 'headers', 'name'],
    data: () => ({
        dialog: false,
        dialogDelete: false,
        search: '',
        editedIndex: -1,
        editedItem: {
            name: '',
            calories: 0,
            fat: 0,
            carbs: 0,
            protein: 0,
        },
        defaultItem: {
            name: '',
            calories: 0,
            fat: 0,
            carbs: 0,
            protein: 0,
        },
        show:false
    }),

    computed: {
        formTitle() {
            return this.editedIndex === -1 ? 'New Item' : 'Edit Item'
        },
    },

    watch: {
        dialog(val) {
            val || this.close()
        },
        dialogDelete(val) {
            val || this.closeDelete()
        },
    },

    created() {
        // this.initialize()
    },

    methods: {
        initialize() {
            this.desserts = [
                {
                    code: 'default',
                    name: "Head Office",
                    // license: 6.0,
                    services: 'Payment,Inquiry',
                    license: `FREE<br>No of Users: 3<br>No of Counters: 3<br>No of Services: 3<br>No of Queue Tickets Per Day: 10`,
                },
            ]
        },

        editItem(item) {
            this.editedItem = Object.assign({}, item)
            this.dialog = true
        },
        deleteItem(item) {
            // this.editedIndex = this.desserts.indexOf(item)
            // this.editedItem = Object.assign({}, item)
            console.log("item", item);
            // return
            switch (this.name) {
                case "counterservice":
                axios.get(`/service/deleteassignedServices/${item.id}`).then(res => {
                        // console.log("user", res.data);
                        window.location.reload()
                        this.$toast.success('assigned services deleted successfully!!')
                    }).catch(err => {
                        this.$toast.error('error deleting service!!',err)
                    })
                    break;
                case "user":
                    axios.get(`/user/deleteuser/${item.id}`).then(res => {
                        // console.log("user", res.data);
                        window.location.reload()
                        this.$toast.success('user deleted successfully!!')
                    }).catch(err => {
                        this.$toast.error('error deleting user!!')
                    })
                    break;
                case "service":
                    axios.get(`/service/deleteservice/${item.id}`).then(res => {
                        // console.log("service", res.data);
                        window.location.reload()
                        this.$toast.success('service deleted successfully!!')
                    }).catch(err => {
                        this.$toast.error('error deleting service!!')
                    })
                    break;
                case "counter":
                    axios.get(`/counter/deletecounter/${item.id}`).then(res => {
                        // console.log("counter", res.data);
                        window.location.reload()
                        this.$toast.success('counter deleted successfully!!')
                    }).catch(err => {
                        this.$toast.error('error deleting counter!!')
                    })
                    // this.show = true
                    break;

                default:
                    break;
            }
            this.show = false
            // this.dialogDelete = true
        },

        deleteItemConfirm() {
            this.desserts.splice(this.editedIndex, 1)
            this.closeDelete()
        },

        close() {
            this.dialog = false
            this.$nextTick(() => {
                this.editedItem = Object.assign({}, this.defaultItem)
                this.editedIndex = -1
            })
        },

        closeDelete() {
            this.dialogDelete = false
            this.$nextTick(() => {
                this.editedItem = Object.assign({}, this.defaultItem)
                this.editedIndex = -1
            })
        },

        save() {
            if (this.editedIndex > -1) {
                Object.assign(this.desserts[this.editedIndex], this.editedItem)
            } else {
                this.desserts.push(this.editedItem)
            }
            this.close()
        },
    },
}
</script>
