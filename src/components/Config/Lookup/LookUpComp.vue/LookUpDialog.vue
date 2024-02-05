<template>
    <div>
        <v-dialog v-model="NewDialog" :max-width="600" persistent :fullscreen="this.$vuetify.breakpoint.width <= 600"
            :transition="this.$vuetify.breakpoint.width <= 600 ? 'dialog-bottom-transition' : undefined">
            <v-card elevation="0">
                <v-card-title class="">
                    <span class="text-h5">{{ FormTitle }}</span>
                </v-card-title>
                <v-card-text class="mt-6">
                    <v-slide-x-transition mode="out-in" appear>
                        <v-container>
                            <v-layout>
                                <v-form>
                                    <!-- Role Master Dialog -->
                                    <v-row v-if="HeaderDialog">
                                        <v-col cols="12" lg="6" sm="6" md="6">
                                            <v-text-field v-model="HeaderItem.code" label="Code" outlined dense
                                                :rules="nameRules"></v-text-field>
                                        </v-col>
                                        <v-col lg="6" cols="12" sm="6" md="6">
                                            <v-text-field v-model="HeaderItem.description" outlined label="Description"
                                                dense :rules="nameRules"> </v-text-field>
                                        </v-col>
                                    </v-row>
                                    <!-- Task Master Dialog -->
                                    <v-row v-if="DetailDialog">
                                        <v-col xl="4" lg="4" cols="12" sm="4" md="4">
                                            <v-text-field v-model="DetailItem.code" :rules="nameRules"
                                                label="Code" outlined dense></v-text-field>
                                        </v-col>
                                        <v-col xl="4" lg="4" cols="12" sm="4" md="4">
                                            <v-text-field v-model="DetailItem.description" :rules="nameRules" outlined
                                                label="Description" dense></v-text-field>
                                        </v-col>
                                        <v-col xl="4" lg="4" cols="12" sm="4" md="4">
                                            <v-text-field v-model="DetailItem.attribute" :rules="nameRules" outlined dense label="Attribute"></v-text-field>                                        </v-col>
                                    </v-row>
                                </v-form>
                            </v-layout>
                        </v-container>
                    </v-slide-x-transition>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="closeDialog">
                        Cancel
                    </v-btn>
                    <v-btn color="blue darken-1" text @click="save" :disabled="issave">
                        Save
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>
<script>
import EventServices from '@/services/EventServices';
export default {
    data() {
        return {
            nameRules: [v => !!v || 'required',],
        };
    },
    props: {
        HeaderItemcopy: {},
        DetailItemcopy: {},
        HeaderItem: {},
        NewDialog: Boolean,
        DetailItem: {},
        HeaderDialog: Boolean,
        DetailDialog: Boolean
    },
    computed: {
        FormTitle() {
            if (this.DetailDialog == true) {

                return this.DetailItem.id == "" ? 'Add Details' : 'Edit Details'
            } else {

                return this.HeaderItem.id == "" ? 'Add Header' : 'Edit Header'
            }
        },
        issave() {
            if (this.DetailDialog) {

                return this.DetailItem.code == "" || this.DetailItem.description == "" || JSON.stringify(this.DetailItem) == JSON.stringify(this.DetailItemcopy)
            } else {
                return this.HeaderItem.code == "" || this.HeaderItem.description == "" || JSON.stringify(this.HeaderItemcopy) == JSON.stringify(this.HeaderItem)
            }

        }
    },
    methods: {

        closeDialog() {
            this.$emit("closeOnly")

            // Close the dialog and reset editedItem
        },
        SetHeader() {
            if (this.HeaderItem.id == "") {
                EventServices.AddLookUpHeader(this.HeaderItem)
                    .then((response) => {

                        if (response.data.status == 'S') {
                            this.MessageBar('S', response.data.errMsg)
                            this.$emit("closeDialog")
                        } else {
                            this.MessageBar('E', response.data.errMsg)
                        }
                    })
                    .catch((error) => {
                        this.MessageBar('E', error)
                    })
            } else {
                
                EventServices.UpdateLookUpHeader(this.HeaderItem)
                    .then((response) => {
                        if (response.data.status == 'S') {
                            this.MessageBar('S', response.data.errMsg)
                            this.$emit("closeDialog")

                        } else {
                            this.MessageBar('E', response.data.errMsg)
                        }
                    })
                    .catch((error) => {
                        this.MessageBar('E', error)
                    })
            }
        },
        SetDetails() {
            if (this.DetailItem.id == "") {
                EventServices.AddLookUpDetails(this.DetailItem)
                    .then((response) => {
                        if (response.data.status == 'S') {
                            this.MessageBar('S', response.data.errMsg)
                            this.$emit("DetailsDialog", this.DetailItem.headerId)
                        } else {
                            this.MessageBar('E', response.data.errMsg)
                        }
                    })
                    .catch((error) => {
                        this.MessageBar('E', error)
                    })
            } else {
                EventServices.UpdateLookUpDetails(this.DetailItem)
                    .then((response) => {
                        if (response.data.status == 'S') {
                            this.MessageBar('S', response.data.errMsg)
                            this.$emit("DetailsDialog", this.DetailItem.headerId)
                        } else {
                            this.MessageBar('E', response.data.errMsg)
                        }
                    })
                    .catch((error) => {
                        this.MessageBar('E', error)
                    })
            }
        },
        save() {
            this.HeaderDialog == true ? this.SetHeader() : this.SetDetails()
        },
    },
};
</script>
    