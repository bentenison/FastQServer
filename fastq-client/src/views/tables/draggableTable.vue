<template>
    <v-card flat>
        <!-- <v-data-table :key="reRender" :headers="headers" :items="items" :items-per-page="5" mobile :search="search">
           
            <template v-slot:body="{ items }">
                <tbody :class="Uclass">
                    <tr v-for="item in items" :key="item.email" class="text-center">
                        <td class="text-left">{{ item.file_name }}</td>
                        <td>{{ item.type }}</td>
                        <td>{{ item.size }}</td>
                        <td>{{ item.created_at }}</td>
                        <td>{{ item.company_name }}</td>

                    </tr>
                </tbody>
            </template>
        </v-data-table> -->
        <v-data-table :headers="headers" :items="items" item-key="id" :show-select="false" :disable-pagination="true"
            :hide-default-footer="true" class="page__table">
            <template v-slot:body="props">
                <draggable :list="props.items" tag="tbody">
                    <tr v-for="(item, index) in props.items" :key="index">
                        <td>
                            <v-icon small class="page__grab-icon">
                                mdi-arrow-all
                            </v-icon>
                        </td>
                        <td> {{ index + 1 }} </td>
                        <td class="text-left">{{ item.file_name }}</td>
                        <td>{{ item.type }}</td>
                        <td>{{ item.size }}</td>
                        <td>{{ item.created_at }}</td>
                        <td>{{ item.company_name }}</td>
                        <td>
                            <v-icon small>
                                mdi-pencil
                            </v-icon>
                        </td>
                    </tr>
                </draggable>
            </template>
        </v-data-table>
    </v-card>
</template>

<script>
import Draggable from 'vuedraggable';
export default {
    props: ["Uclass", "headers", "items"],
    components: {
        Draggable
    },
    data() {
        return {
            reRender: 0,
            dragNdrop: [],
            search: '',
            // headers: [
            //     {
            //         text: 'FileName',
            //         align: 'center',
            //         // filterable: false,
            //         value: 'name',
            //     },
            //     { text: 'Type', value: 'type', align: 'center', },
            //     { text: 'Size', value: 'size', align: 'center', },
            //     { text: 'Modified At', value: 'modified_at', align: 'center', },
            //     { text: 'Modified By', value: 'modified_by', align: 'center', },
            //     // { text: 'Action', value: 'action', align: 'center', sortable: false },
            // ],
            // desserts: [
            //     {
            //         name: '1.mp4',
            //         type: 'MP4',
            //         size: 120023,
            //         modified_at: '12/05/2023',
            //         modified_by: 'Tanvir',
            //         action: 1,
            //     },
            //     {
            //         name: '2.mp4',
            //         type: 'MP4',
            //         size: 120023,
            //         modified_at: '12/05/2023',
            //         modified_by: 'Admin',
            //         // action: 1,
            //     },
            // ],
        }
    },
    methods: {
        // initSortable() {
        //     let table = document.getElementsByClassName(this.Uclass)[0];
        //     console.log("elements:::::", table);
        //     // this way we avoid data binding
        //     this.dragNdrop = JSON.parse(JSON.stringify(this.items));
        //     const _self = this;
        //     console.log(this.dragNdrop);

        //     Sortable.create(table, {
        //         onEnd({
        //             newIndex,
        //             oldIndex
        //         }) {
        //             _self.dragNdrop.splice(
        //                 newIndex,
        //                 0,
        //                 ..._self.dragNdrop.splice(oldIndex, 1)
        //             );
        //         }
        //     });
        // },

    },
    mounted() {
        console.log("mounted");
        // this.initSortable();
    }
}
</script>

<style lang="scss" scoped>
.page--table {
  .page {
    &__table {
      margin-top: 20px;
    }
    &__grab-icon {
      cursor: move;
    }
  }
}
</style>