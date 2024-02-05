<template>
    <div>
        <v-dialog v-model="dialog" max-width="800px" persistent :fullscreen="this.$vuetify.breakpoint.width <= 600"
            :transition="this.$vuetify.breakpoint.width <= 600 ? 'dialog-bottom-transition' : undefined">
            <v-card>
                <v-card-title>
                    <v-layout>
                        <v-flex>
                            <span class="text-wrap text-body-1">{{ formTitle }}</span>
                        </v-flex>
                        <v-flex class="d-flex justify-end">
                            <v-btn icon small color="error" @click="closeDialog">
                                <v-icon>mdi-close</v-icon>
                            </v-btn>
                        </v-flex>
                    </v-layout>
                </v-card-title>
                <v-card-text>
                    <v-container>
                        <v-row class="d-flex flex-column">
                            <v-col v-if="!AddPop && AddDialogFor == ''">
                                <v-data-table :headers="headers" :items="filteredArray" :items-per-page="10" item-key="id"
                                    :search="search" height="400" fixed-header loading-text="Loading please wait..."
                                    no-data-text="No Records available" no-results-text="Record not found">
                                    <template v-slot:top>
                                        <v-toolbar flat>
                                            <v-text-field v-model="search" append-icon="mdi-magnify" label="Search"
                                                single-line hide-details></v-text-field>
                                            <v-divider class="mx-4" inset vertical></v-divider>
                                            <v-btn color="primary" text small class="mb-2 text-capitalize" @click="hidePop">
                                                + Add
                                            </v-btn>
                                        </v-toolbar>
                                    </template>
                                    <template v-slot:item.status="{ item }">
                                        <span>{{ item.status == "Y" ? 'Active' : 'InActive' }}</span>
                                    </template>
                                    <template v-slot:item.actions="{ item }">
                                        <v-hover v-slot="{ hover }">
                                            <v-btn small icon :class="hover ? 'secondary' : 'blue lighten-4'"
                                                @click="hidePop(item, 'E')">
                                                <v-icon small
                                                    :class="hover ? 'white--text' : 'primary--text'">mdi-pencil</v-icon>
                                            </v-btn>
                                        </v-hover>
                                    </template>
                                </v-data-table>
                            </v-col>
                            <v-col v-else>
                                <v-form ref="form" lazyform-validation>
                                    <!-- Broker admin adding field -->
                                    <v-row v-if="dialog && AddDialogFor == ''">
                                        <v-col cols="12" sm="3" xs="6" md="3" lg="3">
                                            <v-autocomplete v-model="brokerData.brokerId" :items="filterBroker"
                                                item-text="brokerName" :rules="required" item-value="brokerId"
                                                label="Broker Name" outlined dense required></v-autocomplete>

                                        </v-col>
                                        <v-col cols="12" sm="3" xs="6" md="3" lg="3">
                                            <v-text-field label="Client Id" v-model="brokerData.clientId" :rules="IdRules"
                                                outlined dense requried></v-text-field>
                                        </v-col>
                                        <v-col cols="12" sm="3" xs="6" md="3" lg="3">
                                            <v-autocomplete v-model="brokerData.roleId" :items="RoleArr"
                                                item-text="roleName" item-value="roleId" dense required label="Role"
                                                outlined :rules="RoleRules"></v-autocomplete></v-col>
                                        <v-col cols="12" sm="3" xs="6" md="3" lg="3">
                                            <v-autocomplete v-model="brokerData.status" :items="StatusBar"
                                                item-value="value" item-text="text" dense required outlined label="Status"
                                                :rules="StatusRules"></v-autocomplete>
                                        </v-col>
                                    </v-row>
                                    <v-row v-if="dialog && AddDialogFor != ''">
                                        <v-col cols="12" sm="3" xs="6" md="3" lg="3">
                                            <v-text-field v-model="domainData.brokerName" label="Broker Name"
                                                :rules="required" outlined dense required>
                                            </v-text-field>
                                        </v-col>
                                        <v-col cols="12" sm="3" xs="6" md="3" lg="3">
                                            <v-text-field label="Domain Addr..." v-model="domainData.domainName"
                                                :rules="DomainRules" outlined dense requried></v-text-field>
                                        </v-col>
                                        <v-col cols="12" sm="3" xs="6" md="3" lg="3">
                                            <v-text-field label="Raw Domain" v-model="domainData.rawDomain"
                                                :rules="required" outlined dense requried></v-text-field>
                                        </v-col>
                                        <v-col cols="12" sm="3" xs="6" md="3" lg="3">
                                            <v-text-field label="AppName" v-model="domainData.appName" :rules="required"
                                                outlined dense requried></v-text-field>
                                        </v-col>
                                        <v-col cols="12" sm="3" xs="6" md="3" lg="3">
                                            <v-text-field label="AuthURL" v-model="domainData.authURL" :rules="required"
                                                outlined dense requried></v-text-field>
                                        </v-col>
                                        <v-col cols="12" sm="3" xs="6" md="3" lg="3">
                                            <v-text-field label="Type" v-model="domainData.type" :rules="required" outlined
                                                dense requried></v-text-field>
                                        </v-col>
                                        <v-col cols="12" sm="3" xs="6" md="3" lg="3">
                                            <v-autocomplete v-model="domainData.status" :items="StatusBar"
                                                item-value="value" item-text="text" dense required outlined label="Status"
                                                :rules="StatusRules"></v-autocomplete>
                                        </v-col>
                                    </v-row>
                                </v-form>
                                <v-card-actions>
                                    <v-spacer></v-spacer>
                                    <v-btn color="error" text @click="hidePop"> Cancel </v-btn>
                                    <v-btn color="blue darken-1" text :disabled="isSaveBroker" @click="save"> Save
                                    </v-btn>
                                </v-card-actions>
                            </v-col>
                        </v-row>
                    </v-container>
                </v-card-text>
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
import EventServices from '../../../services/EventServices';
export default {
    data: () => ({
        brokerData: {
            id: 0,
            brokerId: 0,
            clientId: "",
            roleId: 0,
            status: ""
        },
        defaultData: {
            id: 0,
            brokerId: 0,
            clientId: "",
            roleId: 0,
            status: ""
        },
        domainData: {
            Id: 0,
            brokerName: "",
            domainName: "",
            rawDomain: "",
            appName: "",
            authURL: "",
            type: "",
            status: "",
            brokerId: 0,
        },
        brokerDataCopy: {},
        domainDataCopy: {},
        AddPop: false,
        search: "",
        flag: "",
        IdRules: [
            (v) => !!v || "clientid is required",
        ],
        required: [
            (v) => !!v || "required",
        ],
        RoleRules: [
            (v) => !!v || "Role is required",
        ],
        StatusRules: [
            (v) => !!v || "Status required",
        ],
        DomainRules: [
            (v) => !!v || "Address is required",
            (v) =>
                /^https:\/\/[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(v) ||
                "Domain Name must be valid (e.g., https://example.com)",
        ],
        headers: [
            { text: "BrokerName", value: "brokerName" },
            { text: "ClientId", value: "clientId" },
            { text: "Role", value: "roleName" },
            { text: "Status", value: "status" },
            { text: "Actions", value: "actions", sortable: false, editable: false },
        ],
        StatusBar: [{ text: 'Activate', value: 'Y' }, { text: 'InActivate', value: 'N' }],
    }),
    props: {
        dialog: Boolean,
        brokerId: Number,
        RoleArr: Array,
        AddBroker: Array,
        BrokerArr: Array,
        AddDialogFor: String,
        ActionFlag: String,
        domainValue: Object
    },
    computed: {
        filteredArray() {
            return this.AddBroker.filter(item => item.brokerId === this.brokerId);
        },
        filterBroker() {
            return this.BrokerArr.filter(item => item.brokerId === this.brokerId);
        },
        formTitle() {
            return this.flag == "E" ? 'Ediding Admin' : this.AddPop ? 'Adding Admin' : this.AddDialogFor == '' ? 'Broker Admin' : this.ActionFlag == 'Edit' && this.AddDialogFor != "" ? 'Editing Domain' : 'Adding Domain'
        },
        // issave() {
        //     return JSON.stringify(this.brokerData) == JSON.stringify(this.brokerDataCopy)
        // },
        isSaveBroker() {
            if (this.dialog == true) {
                if (this.AddDialogFor == "") {
                    // Check if any of the properties are empty or null
                    if (
                        this.brokerData.brokerId == "" ||
                        this.brokerData.clientId == "" ||
                        this.brokerData.status == null ||
                        this.brokerData.roleId == null || JSON.stringify(this.brokerData) == JSON.stringify(this.brokerDataCopy)
                    ) {
                        // If any of the conditions are true, return true
                        return true;
                    } else {
                        // All values are filled, return false
                        return false;
                    }
                } else {
                    if (
                        this.domainData.brokerName == "" ||
                        this.domainData.domainName == "" ||
                        this.domainData.rawDomain == "" ||
                        this.domainData.appName == "" ||
                        this.domainData.authURL == "" ||
                        this.domainData.type == "" ||
                        this.domainData.status == "" || JSON.stringify(this.domainData) == JSON.stringify(this.domainDataCopy)
                    ) {
                        // If any of the conditions are true, return true
                        return true;
                    } else {
                        // All values are filled, return false
                        return false;
                    }
                }
            }
            // If this.dialog is true nor false,It returns true
            return true;
        }
    },
    methods: {
        hidePop(item, flag) {
            this.AddPop = !this.AddPop;
            this.flag = flag
            if (this.AddPop != true) {
                this.$nextTick(() => {
                    this.brokerData = Object.assign({}, this.defaultData);
                    this.editedIndex = -1;
                });
            } else {
                if (this.AddDialogFor == 'Domain') {
                    this.closeDialog()
                } else {
                    this.flag = flag
                    if (this.flag == "E") {
                        this.brokerData = item;
                        this.brokerDataCopy = { ...this.brokerData }
                    } else {
                        this.brokerData = this.defaultData
                    }
                }
            }
        },
        closeDialog() {
            this.AddPop = false;
            this.$emit("closeDialog")
        },
        save() {
            if (this.AddDialogFor == "Domain") {
                this.domainData.brokerId = this.brokerId
                this.$emit("SetDomain", this.domainData)
            } else {
                this.brokerData.brokerId = this.brokerId
                EventServices.Adduser(this.brokerData)
                    .then((response) => {
                        if (response.data.status == "S") {
                            this.MessageBar("S", response.data.errMsg);
                            this.AddPop = false;
                            this.$emit("RecallApi")
                        } else {
                            this.MessageBar("E", response.data.errMsg);
                        }
                    })
                    .catch((error) => {
                        this.MessageBar("E", error);
                    });
            }
        },
    },
    watch: {
        ActionFlag: function (action) {
            if (this.AddDialogFor != '' && action == "Edit") {
                this.domainData = this.domainValue;
                this.domainDataCopy = { ...this.domainValue }
            } else {
                this.domainData = {}
                this.domainDataCopy = {}
            }
        }
    }
}
</script>

<style scoped></style>