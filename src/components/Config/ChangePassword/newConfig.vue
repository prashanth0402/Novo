<template>
    <v-data-table :headers="headers" :items="details" sort-by="calories" class="elevation-1">
        <template v-slot:top>
            <v-toolbar flat>
                <v-toolbar-title>User Details</v-toolbar-title>
                <v-divider class="mx-4" inset vertical></v-divider>
                <v-spacer></v-spacer>
                <v-dialog v-model="dialog" max-width="500px" persistent>
                    <template v-slot:activator="{ on, attrs }">
                        <v-btn small text class="primary--text mb-2 elevation-0 caption text-capitalize" v-bind="attrs"
                            v-on="on">
                             + Add user
                        </v-btn>
                    </template>
                    <v-card>
                        <v-card-title>
                            <span class="text-h5">{{ formTitle }}</span>
                        </v-card-title>

                        <v-card-text>
                            <v-container>
                                <v-layout class="this.$vuetify.breakpoint.name < 'sm' ? 'd-flex flex-column' : 'd-flex'">
                                    <v-flex xs12 sm4 md4>
                                        <v-text-field v-model="editedItem.name" label="Client Id"  dense
                                            :rules="[v => !!v || 'required']" :disabled="disable"></v-text-field>
                                    </v-flex>
                                    <v-flex xs12 sm4 md4>
                                        <v-autocomplete :items="items" v-model="editedItem.role" item-text="role"
                                            item-value="role" dense label="Role" required :rules="[v => !!v || 'required']">
                                        </v-autocomplete>
                                    </v-flex>
                                    <v-flex xs12 sm4 md4>
                                        <v-autocomplete :items="type" v-model="editedItem.type" item-text="role"
                                            item-value="status" dense label="Type" required :rules="[v => !!v || 'required']">
                                        </v-autocomplete>
                                    </v-flex>
                                </v-layout>
                            </v-container>
                        </v-card-text>

                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn color="blue darken-1" text @click="close">
                                Cancel
                            </v-btn>
                            <v-btn color="blue darken-1" text @click="save">
                                Save
                            </v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
            </v-toolbar>
        </template>
        <template v-slot:item.actions="{ item }">
            <v-btn small icon>
                <v-icon small @click="editItem(item)">
                    mdi-pencil
                </v-icon>
            </v-btn>
        </template>

    </v-data-table>
</template>

<script>
import EventServices from '../../../services/EventServices'
export default {
    data: () => ({
        dialog: false,
        dialogDelete: false,
        headers: [
            { text: 'AdminId', align: 'center', value: 'name', sortable: false },
            { text: 'Role', align: 'center', value: 'role', sortable: false },
            { text: 'Type', align: 'center', value: 'status', sortable: false },
            { text: 'Actions', align: 'center', value: 'actions', sortable: false },
        ],
        details: [],
        items: [],
        type: ['Active', 'DeActive'],
        editedIndex: -1,
        editedItem: {
            name: '',
            role: '',
            type: '',
        },
        defaultItem: {
            name: '',
            role: '',
            type: '',
        },
        disable: false,
    }),

    computed: {
        formTitle() {
            return this.editedIndex === -1 ? 'New Item' : 'Edit User'
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

    methods: {

        editItem(item) {
            this.editedIndex = this.details.indexOf(item)
            this.editedItem = Object.assign({}, item)
            this.dialog = true
            this.disable = true
        },

        close() {
            this.dialog = false
            this.disable = false
            this.$nextTick(() => {
                this.editedItem = Object.assign({}, this.defaultItem)
                this.editedIndex = -1
            })
        },
        save() {
            if (this.editedIndex > -1) {
                Object.assign(this.details[this.editedIndex], this.editedItem)
            } else {
                this.details.push(this.editedItem)
            }
            this.close()
        },
        GetUserDetails() {
            this.loading = true
            EventServices.Getuser()
                .then((response) => {
                    this.loading = false
                    if (response.data.status == "S") {
                        this.details = response.data.adminListArr
                        this.items = response.data.adminListArr
                    } else {
                        this.MessageBar('E', response.data.errMsg)
                    }
                })
                .catch((error) => {
                    this.loading = false
                    this.MessageBar('E', error)
                })
        }
    },
    mounted() {
        this.GetUserDetails()
    }
}
</script>