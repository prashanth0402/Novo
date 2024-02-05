<template>
    <div>

        <v-tabs v-model="activeTab" left>
            <v-tab>
                <!-- <v-icon left color="blue darken-4" medium> mdi-android</v-icon> -->
                IPO
            </v-tab>
            <v-tab>
                <!-- <v-icon left color="blue darken-4" medium> mdi-apple </v-icon> -->
                SGB
            </v-tab>
            <v-tab>
                <!-- <v-icon left color="blue darken-4" medium> mdi-apple </v-icon> -->
                NCB
            </v-tab>
        </v-tabs>
        <v-window v-model="activeTab" class="elevation-0 mt-5">
            <v-window-item v-for="(tab, index) in tabs" :key="index">
                <v-toolbar flat>
                            <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details
                                class="rounded px-2"></v-text-field>
                            <!-- <v-divider class="mx-4" inset vertical></v-divider> -->
                            <!-- <v-btn color="primary" text small class="mb-2 text-capitalize" @click="openNewItemDialog()">
                                + Add
                            </v-btn> -->
                        </v-toolbar>
                <v-data-table :headers="currentHeaders" :items="currentItems" :search="search" :items-per-page="10"
                    :footer-props="{ 'items-per-page-options': [10], pageText: '', }" :loading="loading"
                    loading-text="Loading... Please wait" no-data-text="No Records available"
                    :item-class="itemRowBackground">
                    <!-- Need to Check if we need else or not -->
                    <template v-slot:item.symbol="{ item }">
                        <v-row wrap class="d-flex text">
                            <v-col cols="6" class="d-flex flex-column text-left">
                                <span class="text--disabled">Symbol</span>
                                <span>{{ item.symbol }}</span>
                            </v-col>
                            <v-col class="d-flex flex-column text-right" cols="6">
                                <span class="text--disabled">Name</span>
                                <span> {{ item.name }}</span>
                            </v-col>
                        </v-row>
                        <v-row wrap class="d-flex text">
                            <v-col cols="6" class="d-flex flex-column text-left">
                                <span class="text--disabled">ISIN</span>
                                <span>{{ item.isin }}</span>
                            </v-col>
                            <v-col class="d-flex flex-column text-right" cols="6">
                                <span class="text--disabled">Exchange</span>
                                <span> {{ item.exchange }}</span>
                            </v-col>
                        </v-row>
                        <v-row wrap class="text">
                            <v-col class="text-center" cols="12">
                                <v-hover v-slot="{ hover }">
                                    <v-btn small width="80" class="text-capitalize white--text elevation-0"
                                        :class="hover ? 'secondary' : item.softDelete == 'N'?'red lighten-2':'green lighten-2'">
                                        <span v-if="item.softDelete == 'N'" @click="editItem(item, tab)">Delete</span>
                                        <span v-else @click="editItem(item, tab)">Restore</span>
                                    </v-btn>
                                </v-hover>
                            </v-col>
                        </v-row>
                    </template>
                </v-data-table>
            </v-window-item>
        </v-window>
        <MasterControlDialog :ControlDialog="ControlDialog" :CurrentTittle="CurrentTittle" @closeonly="closeonly"
            @closeDialog="closeDialog" :item="CurItem"></MasterControlDialog>
    </div>
</template>
<script>
import MasterControlDialog from './MasterControlDialog.vue';



export default {
    data() {
        return {
            CurrentTittle: "",
            CurItem: {},
            ControlDialog: false,
            loading: false,
            search: "",
            activeTab: 0,
            tabs: ["IPO", "SGB", "NCB"],
            IpoMobileHeader: [{ text: "", sortable: false, align: "center", value: "symbol" },],
            SgbMobileHeader: [{ text: "", sortable: false, align: "center", value: "symbol" },],
            NcbMobileHeader: [{ text: "", sortable: false, align: "center", value: "symbol" },],
        };
    },
    props: {
        IpoMasterData: [],
        SgbMasterData: [],
        NcbMasterData: []
    },
    computed: {
        currentHeaders() {
            return this.activeTab == 0 ? this.IpoMobileHeader : this.activeTab == 1 ? this.SgbMobileHeader : this.NcbMobileHeader;
        },
        currentItems() {
            return this.activeTab == 0 ? this.IpoMasterData : this.activeTab == 1 ? this.SgbMasterData : this.NcbMasterData;
        }
    },
    methods: {
        editItem(item, tab) {
            this.CurItem = { ...item };
            this.ControlDialog = true;
            this.CurrentTittle = tab;
        },
        closeDialog() {
            this.$emit("GetAllMaster");
            this.ControlDialog = false;
        },
        closeonly() {
            this.ControlDialog = false;
        },
        itemRowBackground: function (item) {
            if (item.softDelete == "Y") {
                return " grey lighten-3 black-text"
            } else {
                return "white black-text"
            }
        }
    },
    components: { MasterControlDialog }
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



