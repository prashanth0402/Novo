<template>
    <div>
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
        <v-window v-model="activeTab" class="elevation-0 mt-9">
            <v-window-item v-for="(tab, index) in tabs" :key="index">
                <v-card v-if="activeTab === index" elevation="0">
                    <v-data-table :headers="currentHeaders" :items="currentItems" :search="search" :items-per-page="10"
                        :footer-props="{ 'items-per-page-options': [10], pageText: '', }" :loading="loading" 
                        loading-text="Loading... Please wait" no-data-text="No Records available">
                        <template v-slot:top>
                            <v-toolbar flat>
                                <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details
                                class="rounded"></v-text-field>
                                <v-divider class="mx-2" inset vertical></v-divider>
                                <v-btn color="primary" text small class="mt-2 text-capitalize" @click="openNewItemDialog">
                                    + Add
                                </v-btn>
                            </v-toolbar>
                        </template>
                        <!-- Need to Check if we need else or not -->
                        <template v-slot:item.version="{ item }">
                            <v-row wrap class="d-flex text">
                                <v-col cols="6" class="d-flex flex-column text-left">
                                    <span class="text--disabled">Device Name</span>
                                    <span>{{ item.os }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column text-right" cols="6">
                                    <span class="text--disabled">Version</span>
                                    <span> {{ item.version }}</span>
                                </v-col>
                            </v-row>
                            <v-row wrap class="d-flex text">
                                <v-col cols="6" class="d-flex flex-column text-left">
                                    <span class="text--disabled">Force Update</span>
                                    <span>{{ item.forceUpdate == 'Y' ? 'Yes' : 'No' }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column text-right" cols="6">
                                    <span class="text--disabled">Status</span>
                                    <span> {{ item.appStatus == 'Y' ? 'Active' : 'InActive' }}</span>
                                </v-col>
                            </v-row>
                            <v-row wrap class="text">
                                <v-col class="d-flex text-left" cols="6">
                                    <v-layout>
                                        <v-flex class="d-flex flex-column">
                                            <span class="text--disabled">CreatedBy</span>
                                            <span>{{ item.createdBy }}</span>
                                        </v-flex>
                                    </v-layout>
                                </v-col>
                                <v-col class="d-flex flex-column text-right" cols="6">
                                    <span class="text--disabled">UpdatedBy</span>
                                    <span>{{ item.updatedBy }}</span>
                                </v-col>
                            </v-row>
                            <v-row wrap class="text">
                                <v-col class="d-flex text-left" cols="6">
                                    <v-layout>
                                        <v-flex class="d-flex flex-column">

                                            <span class="text--disabled">CreatedDate</span>
                                            <span> {{ item.createdDate }}</span>
                                        </v-flex>
                                    </v-layout>
                                </v-col>
                                <v-col class="d-flex flex-column text-right" cols="6">
                                    <span class="text--disabled">UpdatedDate</span>
                                    <span>{{ item.updatedDate }}</span>
                                </v-col>
                            </v-row>
                            <v-row wrap class="text">
                                <v-col class="text-center" cols="12">
                                    <v-hover v-slot="{ hover }">
                                        <v-btn small width="80"
                                            :class="hover ? 'secondary' : 'text-capitalize blue lighten-4 primary--text elevation-0'">
                                            <span @click="editItem(item)">Edit</span>
                                        </v-btn>
                                    </v-hover>
                                </v-col>
                            </v-row>
                        </template>
                    </v-data-table>
                </v-card>
            </v-window-item>
        </v-window>
        <versionDialog :VersionDialog="VersionDialog" :VersionManger="VersionManger" @closeDialog="closeDialog" @closeOnly="closeOnly" :VersionMangerCopy="VersionMangerCopy">
        </versionDialog>
    </div>
</template>
<script>
import EventServices from '../../../services/EventServices';
import versionDialog from './versionDialog.vue';
export default {
    data() {
        return {
            loading: false,
            DefaultVersionManger: {
                id: 0,
                os: "",
                forceUpdate:null,
                version: "",
                appStatus: null,
            },
            VersionManger: {
                id: 0,
                os: "",
                forceUpdate:null,
                version: "",
                appStatus: null,
            },
            VersionMangerCopy:{},
            VersionDialog: false,
            search: "",
            activeTab: 0,
            tabs: ['ANDROID', 'IOS'],
            titleIcon: ['mdi-android', 'mdi-apple'],
            Android: [],
            Ios: [],
            AndroidMobile: [{ text: "", sortable: false, align: "center", value: "version"},],
            IosMobile: [{ text: "", sortable: false, align: "center", value: "version"},],
        }
    },
    components: {
        versionDialog
    },
    computed: {

    currentHeaders() {
            return this.activeTab === 0 ? this.AndroidMobile : this.IosMobile;
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
                        this.Android = response.data.androidVersionList == null ? [] :response.data.androidVersionList
                        this.Ios = response.data.iosVersionList  == null ? [] :response.data.iosVersionList
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
}
</script>

<style scoped>
::v-deep .v-data-table__mobile-row__header {
    padding-right: 0px !important;
}

.text {
    font-size: .9em;
}

.v-card__title {
    padding: 4px;
}

::v-deep .v-data-table__mobile-row__cell {
    width: 100% !important;
}

::v-deep .v-data-table__mobile-row__cell table tr td:first-child {
    text-align: left;
}

::v-deep .v-data-table__mobile-row__cell table tr td:last-child {
    text-align: right;
}

::v-deep .v-data-footer {
    justify-content: end;
    /* padding: 0;
  padding-left: 200px; */
}

::v-deep .v-data-table__mobile-row__cell table tr {
    padding-bottom: 3px;
}

.row {
    margin-top: 2px;
    margin-bottom: 2px;
}

::v-deep .v-data-table>.v-data-table__wrapper .v-data-table__mobile-row {
    height: initial;
    min-height: 2px;
}

::v-deep .v-data-table>.v-data-table__wrapper>table>tbody>tr>td {
    padding: 16px !important;
}

.col-xl,
.col-xl-auto,
.col-xl-12,
.col-xl-11,
.col-xl-10,
.col-xl-9,
.col-xl-8,
.col-xl-7,
.col-xl-6,
.col-xl-5,
.col-xl-4,
.col-xl-3,
.col-xl-2,
.col-xl-1,
.col-lg,
.col-lg-auto,
.col-lg-12,
.col-lg-11,
.col-lg-10,
.col-lg-9,
.col-lg-8,
.col-lg-7,
.col-lg-6,
.col-lg-5,
.col-lg-4,
.col-lg-3,
.col-lg-2,
.col-lg-1,
.col-md,
.col-md-auto,
.col-md-12,
.col-md-11,
.col-md-10,
.col-md-9,
.col-md-8,
.col-md-7,
.col-md-6,
.col-md-5,
.col-md-4,
.col-md-3,
.col-md-2,
.col-md-1,
.col-sm,
.col-sm-auto,
.col-sm-12,
.col-sm-11,
.col-sm-10,
.col-sm-9,
.col-sm-8,
.col-sm-7,
.col-sm-6,
.col-sm-5,
.col-sm-4,
.col-sm-3,
.col-sm-2,
.col-sm-1,
.col,
.col-auto,
.col-12,
.col-11,
.col-10,
.col-9,
.col-8,
.col-7,
.col-6,
.col-5,
.col-4,
.col-3,
.col-2,
.col-1 {
    width: 100%;
    padding: 0px !important;
}
</style>