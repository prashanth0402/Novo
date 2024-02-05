<template>
    <v-container>
        <v-card class="elevation-0">
            <v-card-title>
                <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line
                    hide-details></v-text-field>
            </v-card-title>
            <v-slide-x-transition mode="out-in" appear>
                <v-data-table :headers="header" :items="records" :search="search"
                    :footer-props="{ 'items-per-page-options': [5] }" :loading="loading" @click:row="handleRowClick"
                    loading-text="Loading... Please wait" no-data-text="No Records available">
                    <template v-slot:item.symbol="{ item }">
                        <v-row wrap class="d-flex text">
                            <v-col cols="6" class="d-flex flex-column text-left">
                                <span class="text--disabled">Symbol</span>
                                <span>{{ item.symbol }} <v-chip x-small class="blue lighten-5 primary--text">{{
                                    item.exchange }}</v-chip></span>
                                <!-- <span class="text--disabled">Unit</span>
                            <span>{{ item.unit }}</span> -->
                            </v-col>
                            <v-col class="d-flex flex-column text-right text" cols="6">
                                <span class="text--disabled">Ordered Date</span>
                                <span>{{ item.applyDate == undefined ? item.orderDate : item.applyDate }}</span>
                            </v-col>
                        </v-row>
                        <v-row wrap class="text">
                            <v-col class="d-flex text-left" cols="6">
                                <v-layout>
                                    <v-flex class="d-flex flex-column">
                                        <span class="text--disabled">Int.RefNo</span>
                                        <span>{{ item.applicationNo == undefined ? item.orderNo : item.applicationNo
                                        }}</span>
                                        <span class="text--disabled" v-if="item.exchOrderNo != undefined">ExchRefno</span>
                                        <span>{{ item.exchOrderNo }}</span>
                                        <span class="text--disabled"
                                            v-if="item.category != '' && item.category != undefined">Category</span>
                                        <span>{{ item.category }}</span>
                                        <!-- <span class="text--disabled">Order Time</span>
                                        <span>{{ item.appliedTime }}</span> -->
                                    </v-flex>

                                </v-layout>
                            </v-col>
                            <v-col class="d-flex justify-space-between flex-column text-right" cols="6">
                                <span class="text--disabled">Status</span>
                                <span class="text-capitalize "
                                    :class="item.status == 'success' ? 'success--text' : 'error--text darken-5'">
                                    {{ item.status }}
                                </span>
                                <span v-if="item.status == 'cancelled'" class="error--text"> Bond cancelled</span>

                                <span
                                    :class="item.statusColor == 'G' ? 'text-capitalize success--text' : 'text-capitalize error--text darken-5'">
                                    {{ item.orderStatus }}
                                </span>

                                <v-flex class="d-flex justify-space-between flex-column">
                                    <span class="text--disabled">ClientId &nbsp;</span>
                                    <span>{{ item.clientId }}</span>
                                </v-flex>
                            </v-col>
                        </v-row>
                    </template>
                </v-data-table>
            </v-slide-x-transition>
        </v-card>

        <ReportMobile :ReportRec="ReportRec" :ShowReportRec="ShowReportRec" @closeHistoryRec="closeHistoryRec"
            v-if="this.$vuetify.breakpoint.width <= 800" />
    </v-container>
</template>
<script>
import ReportMobile from '../../Ipo/ApplicationHistory/HistoryRecDialog.vue';
import EventServices from '@/services/EventServices';
export default {
    components: {
        ReportMobile,
    },
    props: {
        loading: Boolean,
        records: Array,
        Choice: String,

    },
    data() {
        return {
            search: "",
            header: [
                {
                    text: "",
                    sortable: false,
                    align: "center",
                    value: "symbol",
                },
            ],
            ShowReportRec: false,
            ReportRec: {},
        }
    },
    methods: {
        handleRowClick(item) {
            if (this.Choice == "Ipo" && this.Choice != "Sgb") {
                if (window.getSelection().toString().length === 0) {
                    this.showRec(item.masterId, item.applicationNo)
                }
            } else if (this.Choice == "Sgb") {
                this.$emit("dialog", item)
            } if (this.Choice == "G-sec" && this.Choice != "Sgb") {
                this.$emit("Ncbdialog", item)
            } if (this.Choice == "TBill" && this.Choice != "Sgb") {
                this.$emit("Ncbdialog", item)
            } else if (this.Choice == "SDL" && this.Choice != "Sgb") {
                this.$emit("Ncbdialog", item)
            }
        },
        showRec(id, no) {
            this.$globalData.overlay = true
            EventServices.GetHistoryRecors(id, no)
                .then((response) => {
                    if (response.data.status == "S") {
                        this.$globalData.overlay = false
                        this.ReportRec = response.data;
                        for (let ix = 0; ix < this.ReportRec.modifyDetails.length; ix++) {
                            this.ReportRec.modifyDetails[ix].signal = 'O'
                        }
                        this.ShowReportRec = true;
                        //  To change the issuesize in item from integer to string value
                        if (this.ReportRec.issueSize >= 10000000) {
                            this.ReportRec.issueSize = (this.ReportRec.issueSize / 10000000).toFixed(2) + "Cr";
                        } else if (this.ReportRec.issueSize >= 100000) {
                            this.ReportRec.issueSize = (this.ReportRec.issueSize / 100000).toFixed(2) + "L";
                        } else {
                            this.ReportRec.issueSize = this.ReportRec.issueSize.toString();
                        }
                    } else if (response.data.status == "I") {
                        this.$router.replace("/login");
                    } else {
                        this.MessageBar("E", 'Please wait for sometime...');
                        this.$globalData.overlay = false
                    }
                })
                .catch((error) => {
                    this.MessageBar("E", error.response)
                    this.$globalData.overlay = false
                });
        },
        closeHistoryRec() {
            this.ShowReportRec = false
        },
    },
    // updated() {
    //     if (this.records.length != 0) {
    //         this.downloadBtn = false
    //     } else {
    //         this.downloadBtn = true
    //     }
    // }
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
    padding: 20px !important;
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