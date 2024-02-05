<template>
    <v-card elevation="0">
        <v-tabs v-model="activeTab" left>
            <v-tab>
                <v-icon left color="blue darken-4" medium> mdi-android</v-icon>
                 Android
            </v-tab>
            <v-tab>
                <v-icon left color="blue darken-4" medium> mdi-apple </v-icon>
                IOS
            </v-tab>
        </v-tabs>
        <v-window v-model="activeTab" class="elevation-0 mt-5">
            <v-window-item v-for="(tab, index) in tabs" :key="index">
                <v-card v-if="activeTab === index" elevation="0">
                    <!-- <v-toolbar-title class="black--text subtitle-1 text--primary ml-5">{{ currentTitle }}</v-toolbar-title> -->
                    <v-card elevation="0">
                        <v-toolbar flat>
                            <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details
                                class="rounded px-2"></v-text-field>
                            <v-divider class="mx-4" inset vertical></v-divider>
                            <v-btn color="primary" text small class="mb-2 text-capitalize" @click="openNewItemDialog()">
                                + Add
                            </v-btn>
                        </v-toolbar>
                        <v-data-table :search="search" :headers="currentHeaders" :items="currentItems" :items-per-page="10"
                            fixed-header :loading="loading">
                            <template v-slot:item.id="{ item, index }">
                                <span>{{ index + 1 }}</span>
                            </template>
                            <template v-slot:item.forceUpdate="{ item }">
                                {{ item.forceUpdate == 'Y' ? 'Yes' : 'No' }}
                            </template>
                            <template v-slot:item.appStatus="{ item }">
                                {{ item.appStatus == 'Y' ? 'Active' : 'InActive' }}
                            </template>
                            <template v-slot:item.actions="{ item }">
                                <v-hover v-slot="{ hover }">
                                    <v-btn small icon
                                        :class="hover ? 'secondary white--text' : 'blue lighten-4 primary--text'">
                                        <v-icon small @click="editItem(item)">
                                            mdi-pencil
                                        </v-icon>
                                    </v-btn>
                                </v-hover>
                            </template>
                        </v-data-table>
                    </v-card>
                </v-card>
            </v-window-item>
        </v-window>
        <versionDialog :VersionDialog="VersionDialog" :VersionManger="VersionManger" @closeDialog="closeDialog"
            @closeOnly="closeOnly" :VersionMangerCopy="VersionMangerCopy">
        </versionDialog>
    </v-card>
</template>
<script>
import EventServices from '@/services/EventServices';
import versionDialog from './versionDialog.vue';
export default {
    components: {
        versionDialog
    },
    data() {
        return {
            loading: false,
            DefaultVersionManger: {
                id: 0,
                os: "",
                forceUpdate: null,
                version: "",
                appStatus: null,
            },
            VersionManger: {
                id: 0,
                os: "",
                forceUpdate: null,
                version: "",
                appStatus: null,
            },
            VersionMangerCopy: {},
            VersionDialog: false,
            search: "",
            activeTab: 0,
            tabs: ['ANDROID', 'IOS'],
            titleIcon: ['mdi-android', 'mdi-apple'],
            Android: [],
            Ios: [],
            AndroidHeaders: [
                { text: 'S.No', align: 'start', sortable: false, value: 'id' },
                { text: 'DeviceName', value: 'os' },
                { text: 'Version', value: 'version' },
                { text: 'ForceUpdate', value: 'forceUpdate' },
                { text: 'Status', value: 'appStatus' },
                { text: 'CreatedBy', value: 'createdBy' },
                { text: 'CreatedDate', value: 'createdDate' },
                { text: 'UpdatedBy', value: 'updatedBy' },
                { text: 'UpdatedDate', value: 'updatedDate' },
                { text: "Actions", value: "actions", sortable: false },
            ],

            IosHeaders: [
                { text: 'S.No', align: 'start', sortable: false, value: 'id' },
                { text: 'DeviceName', value: 'os' },
                { text: 'Version', value: 'version' },
                { text: 'ForceUpdate', value: 'forceUpdate' },
                { text: 'Status', value: 'appStatus' },
                { text: 'CreatedBy', value: 'createdBy' },
                { text: 'CreatedDate', value: 'createdDate' },
                { text: 'UpdatedBy', value: 'updatedBy' },
                { text: 'UpdatedDate', value: 'updatedDate' },
                { text: "Actions", value: "actions", sortable: false },
            ],
        }
    },
    computed: {
        currentHeaders() {
            return this.activeTab === 0 ? this.AndroidHeaders : this.IosHeaders;
        },
        currentItems() {
            return this.activeTab === 0 ? this.Android : this.Ios;
        },
        currentTitle() {
            return this.tabs[this.activeTab];
        },
        currentIcon() {
            return this.titleIcon[this.activeTab];
        }
    },

    methods: {
        closeDialog() {
            this.close()
            this.VersionDialog = false
            this.GetVersion()
        },
        closeOnly() {
            this.close()
            this.VersionDialog = false
        },
        GetVersion() {
            this.loading = true
            EventServices.GetVersion()
                .then((response) => {
                    if (response.data.status == 'S') {
                        this.Android = response.data.androidVersionList == null ? [] : response.data.androidVersionList
                        this.Ios = response.data.iosVersionList == null ? [] : response.data.iosVersionList
                        this.loading = false
                    } else {
                        this.MessageBar('E', response.data.errMsg)
                    }
                })
                .catch((error) => {
                    console.log("Error :", error)

                })
        },
        openNewItemDialog() {
            if (this.currentTitle == 'ANDROID') {
                this.VersionManger.os = "ANDROID"
                this.VersionDialog = true
            } else if (this.currentTitle == 'IOS') {
                this.VersionManger.os = "IOS"
                this.VersionDialog = true
            }
        },
        close() {
            this.$nextTick(() => {
                this.VersionManger = Object.assign({}, this.DefaultVersionManger)
                this.VersionMangerCopy = Object.assign({}, this.DefaultVersionManger)
            })
        },
        editItem(item) {
            if (this.currentTitle == 'ANDROID') {
                this.VersionManger.id = item.id
                this.VersionManger.os = item.os
                this.VersionManger.version = item.version
                this.VersionManger.forceUpdate = item.forceUpdate
                this.VersionManger.appStatus = item.appStatus
                this.VersionMangerCopy = { ...this.VersionManger };
                this.VersionDialog = true
            } else if (this.currentTitle == 'IOS') {
                this.VersionManger.id = item.id
                this.VersionManger.os = item.os
                this.VersionManger.version = item.version
                this.VersionManger.forceUpdate = item.forceUpdate
                this.VersionManger.appStatus = item.appStatus
                this.VersionMangerCopy = { ...this.VersionManger };
                this.VersionDialog = true
            }
        },
    },
    mounted() {
        this.GetVersion()
    }

};

</script>

<style scoped>

::v-deep .theme--light.v-data-table > .v-data-table__wrapper > table > thead > tr:last-child > th {
    /* border-bottom: thin solid rgba(0, 0, 0, 0.12); */
    white-space: nowrap;
    /* vertical-align: sub; */
}

</style>
