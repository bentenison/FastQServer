<template>
    <v-card flat>
        <v-card-title>
            <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details></v-text-field>
        </v-card-title>
        <v-data-table :headers="headers" :items="items" :items-per-page="5" mobile :search="search">
            <template v-slot:item.action="{ item }">
                <!-- <v-icon small class="mr-2">
                    mdi-pencil
                </v-icon> -->
                <v-icon color="error" @click="">
                    mdi-delete
                </v-icon>
            </template>
        </v-data-table>
    </v-card>
</template>

<script>
export default {
    props: ["headers", "items"],
    data() {
        return {
            search: '',

        }
    },
}
</script>

<style lang="scss" scoped></style>