<template>
    <div>
        <v-container class="mt-16">
            <v-flex>
                <v-icon left color="blue darken-4" medium class="mb-2">mdi-domain-plus</v-icon>
                <span class="black--text subtitle-1 text--primary mb-3">
                    DOMAIN MASTER </span>
            </v-flex>
            <v-slide-x-transition mode="out-in" appear>
                <v-card class="elevation-0" outlined>
                    <v-form ref="form" lazyform-validation>
                        <v-data-table :headers="headers" :items="AddDomain" :items-per-page="10" item-key="id"
                            :search="search" height="400" fixed-header loading-text="Loading please wait..."
                            no-data-text="No Records available" no-results-text="Record not found">
                            <template v-slot:top>
                                <v-toolbar flat>
                                    <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line
                                        hide-details></v-text-field>
                                    <v-divider class="mx-4" inset vertical></v-divider>
                                    <v-btn color="primary" text small class="mb-2 text-capitalize"
                                        @click="openDialog('Domain', 0)">
                                        + Add
                                    </v-btn>
                                </v-toolbar>
                            </template>
                            <template v-slot:item.domainName="{ item }">
                                <span>{{ item.domainName }} </span>
                            </template>
                            <template v-slot:item.brokerName="{ item }">
                                <span> {{ item.brokerName }}</span>
                            </template>
                            <template v-slot:item.adminCount="{ item }">
                                <v-btn small icon class="primary--text info lighten-5"
                                    @click="openDialog('', item.id)"><span>{{
                                        item.adminCount
                                    }}</span>
                                </v-btn>
                            </template>
                            <template v-slot:item.status="{ item }">
                                <span>{{ item.status == 'Y' ? 'Active' : 'Inctive' }}</span>
                            </template>

                            <template v-slot:item.actions="{ item }">
                                <v-layout>
                                    <v-hover v-slot="{ hover }">
                                        <v-btn small icon
                                            :class="!item.editable ? hover ? 'secondary' : 'blue lighten-4' : 'error'">
                                            <v-icon small @click="openDialog('Domain', item.id, 'Edit', item)"
                                                :class="!item.editable ? hover ? 'white--text' : 'primary--text' : 'white--text'">
                                                {{ item.editable ? 'mdi-close' : 'mdi-pencil' }}
                                            </v-icon>
                                        </v-btn>
                                    </v-hover>
                                    <v-btn small icon class="success ml-2" v-if="item.editable">
                                        <v-icon text small @click="applyChanges(item)" color="white">
                                            mdi-check
                                        </v-icon>
                                    </v-btn>
                                </v-layout>
                            </template>
                        </v-data-table>
                    </v-form>
                </v-card>
            </v-slide-x-transition>

            <OpenDomainDialog :dialog="dialog" :AddDialogFor="AddDialogFor" :ActionFlag="ActionFlag"
                :domainValue="domainValue" :brokerId="brokerId" :RoleArr="RoleArr" :AddBroker="AddBroker"
                :BrokerArr="BrokerArr" @closeDialog="openDialog" @RecallApi="RecallApi" @SetDomain="SetDomain" />

        </v-container>
    </div>
</template>

<script>
import OpenDomainDialog from "./OnboardingDialog.vue"
import EventServices from '../../../services/EventServices';
export default {
    data: () => ({
        AddDialogFor: "",
        ActionFlag: "",
        StatusBar: [{ status: 'Y', desc: 'Activate' }, { status: 'N', desc: 'InActivte' }],
        DefaultDomain: {
            Flag: "A",
            Id: 0,
            DomainName: "",
            BrokerName: "",
            Status: "",
        },
        Domaindata: {
            Flag: "",
            Id: 0,
            DomainName: "",
            BrokerName: "",
            Status: "",
        },
        Domaindatacopy: {},
        domainValue: {},
        dialog: false,
        DomainRules: [
            (v) => !!v || "Domain Name is required",
            (v) =>
                /^https:\/\/[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(v) ||
                "Domain Name must be valid (e.g., https://example.com)",
        ],
        nameRules: [
            (v) => !!v || "Name is required",
            (v) => (v && v.length <= 20) || "Name must be less than 20 characters",
        ],
        search: '',
        headers: [
            { text: "BrokerName", value: "brokerName" },
            { text: "DomainName", value: "domainName" },
            { text: "Raw Domain", value: "rawDomain" },
            { text: "App Name", value: "appName" },
            { text: "AuthURL", value: "authURL" },
            { text: "Type", value: "type" },
            { text: "No.of Admins", value: "adminCount" },
            { text: "Status", value: "status" },
            { text: "CreatedBy", value: "createdBy" },
            { text: "CreatedDate", value: "createdDate" },
            { text: "", value: "actions", editable: false, sortable: false },
        ],
        loading: false,
        brokerId: 0
    }),
    props: {
        AddDomain: Array,
        RoleArr: Array,
        AddBroker: Array,
        BrokerArr: Array
    },
    components: {
        OpenDomainDialog
    },
    methods: {
        openDialog(whatFor, id, action, item) {
            this.AddDialogFor = whatFor
            this.ActionFlag = action
            this.dialog = !this.dialog
            this.brokerId = id
            this.domainValue = item
        },
        close() {
            this.$nextTick(() => {
                this.Domaindata = Object.assign({}, this.DefaultDomain);
                this.editedIndex = -1;
            });
            this.ActionFlag = ""
        },
        applyChanges(item) {
            item.editable = false;
            this.Domaindata.Id = item.id
            this.Domaindata.DomainName = item.domainName;
            this.Domaindata.BrokerName = item.brokerName;
            this.Domaindata.Status = item.status;
            this.Domaindata.Flag = "E"
            if (this.$refs.form.validate()) {
                this.SetDomain()
                this.close()
            } else {
                this.MessageBar("E", "Please Enter an valid Details To save")
            }
        },
        SetDomain(domainData) {
            this.$globalData.overlay = true;
            EventServices.SetDomain(domainData)
                .then((response) => {
                    this.$globalData.overlay = false;
                    if (response.data.status == "S") {
                        this.MessageBar("S", response.data.errMsg);
                        this.$emit("RecallApi")
                        this.dialog = false;
                        this.ActionFlag = ""
                    } else {
                        this.MessageBar("E", response.data.errMsg);
                    }
                })
                .catch((error) => {
                    this.$globalData.overlay = false;
                    this.MessageBar("E", error);
                });

        },
        RecallApi() {
            this.$emit("RecallApi")
        }
    }
}
</script>

<style scoped></style>