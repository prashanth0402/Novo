<template>
    <div>
        <v-tabs v-model="activeTab" left>
            <v-tab>Header</v-tab>
            <v-tab>Detail</v-tab>
        </v-tabs>
        <v-window v-model="activeTab" class="elevation-0 mt-9">
            <v-window-item v-for="(tab, index) in tabs" :key="index">
                <v-card v-if="activeTab === index" elevation="0">
                    <!-- <v-toolbar-title class="black--text subtitle-1 text--primary ml-5"></v-toolbar-title> -->
                    <v-layout row wrap class="mb-5" style="padding-left: 12px;">
                        <v-flex xs12 lg9 class="mt-6 d-flex justify-left">
                            <v-icon left color="blue darken-4" medium>
                                {{ currentIcon }}
                            </v-icon>
                            <span class="text-subtitle-1 text-capitalize">
                                {{ currentTitle }}
                            </span>
                        </v-flex>
                    </v-layout>
                    <v-data-table :headers="currentHeaders" :items="currentItems" :search="search" :items-per-page="10"
                        :footer-props="{ 'items-per-page-options': [10], }" :loading="loading"
                        loading-text="Loading... Please wait" no-data-text="No Records available">
                        <template v-slot:top>
                            <v-toolbar flat>
                                <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line
                                    hide-details></v-text-field>
                                <v-divider class="mx-4" inset vertical></v-divider>
                                <v-btn color="primary" text small class="mb-2 text-capitalize" @click="openNewItemDialog()"
                                    :disabled="currentTitle == 'Detail' && details.length <= 0 && DetailItem.headerId == '' ? true : false">
                                    + Add
                                </v-btn>
                            </v-toolbar>
                        </template>
                        <template v-slot:item.code="{ item }" v-if="activeTab == 0">
                            <v-row wrap class="d-flex text">
                                <v-col cols="6" class="d-flex flex-column text-left">
                                    <span class="text--disabled">Code</span>
                                    <span>{{ item.code }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column text-right" cols="6">
                                    <span class="text--disabled">Description</span>
                                    <span> {{ item.description }}</span>
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
                                    <span class="text--disabled">CreatedDate</span>
                                    <span> {{ item.createdDate }}</span>

                                </v-col>
                            </v-row>
                            <v-row wrap class="text">
                                <v-col class="d-flex text-left" cols="6">
                                    <v-layout>
                                        <v-flex class="d-flex flex-column">
                                            <span class="text--disabled">UpdatedBy</span>
                                            <span>{{ item.updatedBy }}</span>
                                        </v-flex>
                                    </v-layout>
                                </v-col>
                                <v-col class="d-flex flex-column text-right" cols="6">
                                    <span class="text--disabled">UpdatedDate</span>
                                    <span>{{ item.updatedDate }}</span>
                                </v-col>
                            </v-row>
                            <!-- <v-row wrap class="text">
                                <v-col class="d-flex text-left" cols="6">
                                    <v-layout>
                                        <v-flex class="d-flex flex-column">
                                            <span class="text--disabled">CreatedBy</span>
                                            <span>{{ item.createdBy }}</span>
                                        </v-flex>
                                    </v-layout>
                                </v-col>
                            </v-row> -->
                            <v-row wrap class="text text-center d-flex justify-center">
                                <v-col >
                                    <v-hover v-slot="{ hover }">
                                        <v-btn small width="80"
                                            :class="hover ? 'secondary' : 'text-capitalize blue lighten-4 primary--text elevation-0'">
                                            <span @click="editItem(item)">Edit</span>
                                        </v-btn>
                                    </v-hover>
                                </v-col>
                                <v-col class="d-flex justify-center " >
                                    <v-hover v-slot="{ hover }">
                                        <v-btn small width="100" @click="FetchLookUp(item.id)"
                                            :class="hover ? 'secondary' : 'text-capitalize blue lighten-4 primary--text elevation-0'">
                                            <span>Details <v-icon x-small>mdi-chevron-right</v-icon>
                                            </span>
                                        </v-btn>
                                    </v-hover>
                                </v-col>
                            </v-row>
                        </template>
                        <template v-slot:item.code="{ item }" v-else>
                            <v-row wrap class="d-flex text">
                                <v-col cols="6" class="d-flex flex-column text-left">
                                    <span class="text--disabled">Code</span>
                                    <span>{{ item.code }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column text-right" cols="6">
                                    <span class="text--disabled">Description</span>
                                    <span>{{ item.description }}</span>
                                </v-col>
                            </v-row>
                            <!-- <v-row wrap class="d-flex text">
                                <v-col cols="6" class="d-flex flex-column text-left">
                                    <span class="text--disabled">Parent Menu</span>
                                    <span> {{ item.parentId == 0 ? 'NA' : item.parentId }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column text-right" cols="6">
                                    <span class="text--disabled">Status</span>
                                    <span> {{ item.type == 'Y' ? 'Active' : 'InActive' }}</span>
                                </v-col>
                            </v-row> -->
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
        <LookUpDialog :NewDialog="NewDialog" :HeaderDialog="HeaderDialog" :HeaderItem="HeaderItem" :DetailItem="DetailItem"
            :DetailDialog="DetailDialog" @closeDialog="closeDialog" @closeOnly="closeOnly" @DetailsDialog="DetailsDialog"
            :DetailItemcopy="DetailItemcopy" :HeaderItemcopy="HeaderItemcopy">
        </LookUpDialog>
    </div>
</template>
<script>
import EventServices from '../../../../services/EventServices';
import LookUpDialog from './LookUpDialog.vue';
export default {
    props: {
        Header: Array
    },
    components: {
        LookUpDialog
    },
    data() {
        return {
            loading: false,
            defaulHead: {
                id: "",
                user: "",
                code: "",
                description: "",
            },
            HeaderItem: {
                id: "",
                user: "",
                code: "",
                description: "",
            },
            defaultDetail: {
                id: "",
                headerId: "",
                code: "",
                description: "",
                attribute: "",
            },
            DetailItem: {
                id: "",
                headerId: "",
                code: "",
                description: "",
                attribute: "",
            },
            HeaderItemcopy:{},
            DetailItemcopy:{},
            NewDialog: false,
            DetailDialog: false,
            HeaderDialog: false,
            search: "",
            activeTab: 0,
            isAddingNewItem: true,
            tabs: ['Header', 'Detail'],
            titleIcon: ['mdi-account-edit', 'mdi-text-box-edit-outline'],
            HeadersMobile: [{ text: "", sortable: false, align: "center", value: "code" },],
            DetailsMobile: [{ text: "", sortable: false, align: "center", value: "code" },],

            details: []
        }
    },
    computed: {
        currentHeaders() {
            return this.activeTab === 0 ? this.HeadersMobile : this.DetailsMobile;
        },
        currentItems() {
            return this.activeTab === 0 ? this.Header : this.details;
        },
        currentTitle() {
            return this.tabs[this.activeTab];
        },
        currentIcon() {
            return this.titleIcon[this.activeTab];
        }
    },

    methods: {
        DetailsDialog(HeaderId) {
            this.DetailDialog = false
            this.HeaderDialog = false
            this.NewDialog = false
            this.FetchLookUp(HeaderId)
            this.close()

        },
        closeDialog() {
            this.close()
            this.DetailDialog = false
            this.HeaderDialog = false
            this.NewDialog = false
            this.GetHeaders()
        },
        closeOnly() {
            this.close()
            this.DetailDialog = false
            this.HeaderDialog = false
            this.NewDialog = false
        },
        GetHeaders() {
            this.$emit('GetHeaders')
        },

        openNewItemDialog() {
            if (this.currentTitle == 'Header') {
                this.HeaderDialog = true
                this.NewDialog = true;
            } else if (this.currentTitle == 'Detail') {
                this.DetailDialog = true
                this.NewDialog = true;
            }
        },
        close() {
            this.$nextTick(() => {
                this.HeaderItem = Object.assign({}, this.defaulHead)
                this.DetailItem = Object.assign({}, this.defaultDetail)
                this.HeaderItemcopy = Object.assign({}, this.defaulHead)
                this.DetailItemcopy = Object.assign({}, this.defaultDetail)
                // this.editedIndex = -1
            })
        },
        editItem(item) {

            if (this.currentTitle == 'Header') {
                this.HeaderItem.id = item.id
                this.HeaderItem.code = item.code
                this.HeaderItem.description = item.description
                this.HeaderItem.user = item.user
                this.HeaderItemcopy = {...this.HeaderItem}
                this.HeaderDialog = true
                this.NewDialog = true;
            } else if (this.currentTitle == 'Detail') {
                this.DetailItem.id = item.id
                this.DetailItem.headerId = item.headerId
                this.DetailItem.code = item.code
                this.DetailItem.description = item.description
                this.DetailItem.attribute = item.attribute
                this.DetailItemcopy = {...this.DetailItem}
                this.DetailDialog = true
                this.NewDialog = true;
            }
        },
        FetchDetail(event, row) {
            const item = row.item;
            if (this.activeTab == 0) {
                this.FetchLookUp(item.id)
            }

        },
        FetchLookUp(HeaderId) {
            this.$globalData.overlay = true;
            EventServices.FetchLookUpDetails(HeaderId)
                .then((response) => {
                    if (response.data.status == "S") {
                        if (response.data.details != [] && response.data.details != null) {
                            this.details = response.data.details;
                            this.DetailItem.headerId = HeaderId
                            this.activeTab = 1;
                        } else {
                            this.details = []
                            this.activeTab = 1
                            this.DetailItem.headerId = HeaderId
                        }
                        this.$globalData.overlay = false;
                    } else {
                        this.MessageBar("E", response.data.errMsg)
                        this.$globalData.overlay = false;

                    }
                })
                .catch((error) => {
                    this.MessageBar("E", error)
                    this.$globalData.overlay = false;

                });
        },
    },

};

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