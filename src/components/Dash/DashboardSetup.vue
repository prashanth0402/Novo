<template>
    <div class="mt-10">
        <v-container>
            <v-breadcrumbs :items="crums">
                <template v-slot:divider>
                    <v-icon>mdi-chevron-right</v-icon>
                </template>
            </v-breadcrumbs>
            <v-layout row wrap class="mb-5" style="padding-left: 12px;">
                <v-flex xs12 lg9 class="mt-6 d-flex justify-left">
                    <v-icon left color="blue darken-4" medium>
                        mdi-view-dashboard-edit
                    </v-icon>
                    <span class="text-subtitle-1">Dashboard Setup</span>
                </v-flex>
            </v-layout>
            <v-layout class="d-flex flex-column">
                <v-flex>
                    <v-slide-x-transition appear mode="out-in">
                        <v-data-table :headers="headers" :items="items" sort-by="calories" class="elevation-0"
                            :loading="loading" :search="search" :items-per-page="10">
                            <template v-slot:top>
                                <v-toolbar flat>
                                    <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line
                                        hide-details></v-text-field>
                                    <v-divider class="mx-4" inset vertical></v-divider>
                                    <v-btn color="primary" text small class="mb-2 text-capitalize" @click="add">
                                        + Add
                                    </v-btn>
                                </v-toolbar>
                            </template>

                            <template v-slot:item.actions="{ item }">
                                <v-hover v-slot="{ hover }">
                                    <v-btn small icon :class="hover ? 'secondary' : 'blue lighten-4'">
                                        <v-icon small @click="editItem(item)"
                                            :class="hover ? 'white--text' : 'primary--text'"> mdi-pencil
                                        </v-icon>
                                    </v-btn>
                                </v-hover>
                            </template>
                            <template v-slot:item.preview="{ item }">
                                <v-btn icon small @click="previewCard(item)">
                                    <v-icon>mdi-eye-outline</v-icon>
                                </v-btn>
                            </template>
                        </v-data-table>
                    </v-slide-x-transition>
                </v-flex>
            </v-layout>
        </v-container>

        <v-dialog v-model="dialog" max-width="800px" persistent>
            <v-card>
                <v-card-title>
                    <v-layout>
                        <v-flex>
                            <span class="text-h5">{{ formTitle }}</span>
                        </v-flex>
                        <v-flex class="d-flex justify-end" v-if="preview">
                            <v-btn icon @click="previewCard">
                                <v-icon color="error">mdi-close</v-icon>
                            </v-btn>
                        </v-flex>
                    </v-layout>
                </v-card-title>
                <v-card-text v-if="!preview">
                    <v-slide-x-transition mode="out-in" appear>
                        <v-container>
                            <v-form ref="form" lazy-validation>
                                <v-row>
                                    <v-col cols="12" sm="6" xs="6" md="4">
                                        <v-text-field v-model="editedItem.fullname" label="FullName"
                                            :rules="rules"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="6" xs="6" md="4">
                                        <v-text-field v-model="editedItem.name" label="Name" :rules="rules"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="6" xs="6" md="4">
                                        <v-autocomplete label="path" v-model="editedItem.taskId" :items="routerArr"
                                            item-text="routerName" item-value="taskId"></v-autocomplete>
                                    </v-col>
                                    <v-col cols="12" sm="6" xs="6" md="4">
                                        <v-text-field v-model="editedItem.image" label="Image"
                                            :rules="[urlValidationRule]"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="6" xs="6" md="4">
                                        <v-text-field v-model="editedItem.color" label="BgColor"
                                            :rules="rules"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="6" xs="6" md="4">
                                        <v-autocomplete v-model="editedItem.status" label="Status" :rules="rules"
                                            :items="status" item-value="value" item-text="text"></v-autocomplete>
                                    </v-col>
                                </v-row>
                            </v-form>
                        </v-container>
                    </v-slide-x-transition>
                </v-card-text>

                <v-card-actions v-if="!preview">
                    <v-spacer></v-spacer>
                    <v-btn color="error" text @click="close"> Cancel </v-btn>
                    <v-btn color="blue darken-1" text @click="save" :disabled="issave"> Save </v-btn>
                </v-card-actions>
                <DashCard v-else :segment="editedItem" :class="preview ? 'pa-10' : undefined" />
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
import EventServices from "@/services/EventServices";
import DashCard from "./DashCard.vue";
export default {
    data: () => ({
        crums: [],
        search: "",
        dialog: false,
        dialogDelete: false,
        checkBox: false,
        preview: false,
        rules: [(value) => !!value || "Required."],
        status: [
            { text: "Active", value: "Y" },
            { text: "InActive", value: "N" },
        ],
        headers: [
            {
                text: "FullName",
                align: "left",
                sortable: false,
                value: "fullname",
            },
            {
                text: "Name",
                align: "start",
                sortable: false,
                value: "name",
            },
            { text: "Image", value: "image", sortable: false },
            { text: "CardColor", value: "color", sortable: false },
            { text: "Status", align: "center", value: "status", sortable: false },
            { text: "", align: "center", value: "actions", sortable: false },
            { text: "", align: "center", value: "preview", sortable: false }
        ],
        headers2: [
            {
                text: "ISIN",
                align: "start",
                sortable: false,
                value: "symbol",
            },
        ],
        items: [],
        routerArr: [],
        editedIndex: -1,
        editedItem: {
            id: 0,
            fullname: "",
            name: "",
            image: "",
            taskId: 0,
            color: "",
            status: ""
        },
        defaultItem: {
            id: 0,
            fullname: "",
            name: "",
            image: "",
            taskId: 0,
            color: "",
            status: ""
        },
        copy: [],
        loading: false,
        menu1: false,
        menu2: false,
        menu3: false,
        menu4: false,
    }),
    computed: {
        formTitle() {
            return this.preview ? 'Preview Component' : this.editedIndex === -1 ? "Add Link" : "Edit Link";
        },
        issave() {
            return this.editedItem.fullname == "" || this.editedItem.name == "" || this.editedItem.image == "" || this.editedItem.color == "" || this.editedItem.status == "";
        },
        urlValidationRule() {
            return (value) => {
                // Check if the value starts with 'https://'
                if (value && !value.startsWith('https://')) {
                    return 'URL must start with "https://"';
                }
                return true; // Validation passed
            };
        },
    },
    created() {
        this.initialize();
    },
    methods: {
        previewCard(item) {
            this.dialog = !this.dialog;
            this.preview = !this.preview;

            if (this.preview == true) {
                this.editedItem = item
            } else {
                this.editedItem = this.defaultItem
            }
        },
        initialize() {
            this.loading = true;
            EventServices.GetDashboardDetail(this.$route.path)
                .then((response) => {
                    this.loading = false;
                    if (response.data.status == "S") {
                        this.items = response.data.segmentArr;
                        this.routerArr = response.data.routerArr;
                    }
                })
                .catch((error) => {
                    this.loading = false;
                    this.MessageBar("E", error);
                });
        },
        editItem(item) {
            this.editedIndex = this.items.indexOf(item);
            this.editedItem = Object.assign({}, item);
            this.dialog = true;
        },
        add() {
            this.dialog = true;
        },
        close() {
            // console.log("Closed")
            this.dialog = false;
            this.$nextTick(() => {
                this.editedItem = Object.assign({}, this.defaultItem);
                this.editedIndex = -1;
            });
            this.$refs.form.resetValidation();
        },
        save() {
            this.$refs.form.validate();
            if (this.editedItem.isin != "" &&
                this.editedItem.drhpLink != "" &&
                this.editedItem.allotmentFinal != "" &&
                this.editedItem.refundInitiate != "" &&
                this.editedItem.dematTransfer != "" &&
                this.editedItem.listingDate != "") {
                this.$globalData.overlay = true;
                EventServices.SetDashbord(this.editedItem)
                    .then((response) => {
                        this.$globalData.overlay = false;
                        if (response.data.status == "S") {
                            this.initialize();
                            this.MessageBar("S", "");
                        }
                        else {
                            this.MessageBar("E", response.data.errMsg);
                        }
                    })
                    .catch((error) => {
                        this.$globalData.overlay = false;
                        this.MessageBar("E", error);
                    });
                this.close();
            }
            else {
                this.MessageBar("E", "Fill all the details");
            }
        },
    },
    components: { DashCard }
};

</script>

<style  scoped></style>