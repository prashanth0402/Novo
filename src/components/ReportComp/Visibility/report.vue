<template>
    <div class="d-flex  justify-center flex-column">
        <v-card elevation="0">
            <v-slide-x-transition mode="out-in" appear>
                <v-data-table :headers="headers" :items="records" :search="search" :items-per-page="10" :loaging="loading"
                    item-key="Symbol" loading-text="Loading... Please wait" no-data-text="No Records available"
                    @click:row="handleRowClick">
                    <template v-slot:item.status="{ item }">
                        <span :class="item.status == 'success' ? 'green--text text-capitalize' : 'red--text'"
                            v-if="item.flag != 'Y'">
                            {{ item.status }}
                        </span>
                        <span v-else class="error--text">Bond cancelled</span>
                    </template>

                    <template v-slot:item.orderStatus="{ item }">
                        <span
                            :class="item.statusColor == 'G' ? 'text-capitalize success--text' : 'text-capitalize error--text darken-5'">
                            {{ item.orderStatus }}
                        </span>
                    </template>

                    <!-- <template v-slot:item.category="{ item }">
          <th v-if="item.category!=''">{{ item.category }}</th>
        </template> -->

                </v-data-table>
            </v-slide-x-transition>
        </v-card>
        <v-layout>
            <HistoryDetails :ReportRec="ReportRec" :ShowReportRec="ShowReportRec" @closeHistoryRec="closeHistoryRec" />
        </v-layout>
    </div>
</template>

<script>
import EventServices from '@/services/EventServices';
import HistoryDetails from '../../Ipo/ApplicationHistory/HistoryRecDialog.vue';
export default {
    components: {
        // ReportDesk,
        HistoryDetails,
    },
    props: {
        loading: Boolean,
        records: Array,
        Choice: String,
        headers: Array,

    },
    data() {
        return {
            search: "",
            isMobile: false,
            // header: [
            //     { text: 'Symbol', align: 'left', value: 'symbol', sortable: false },
            //     { text: 'Exchange', value: 'exchange', align: "center", sortable: true },
            //     { text: 'Application no.', value: 'applicationNo', align: "center", sortable: false },
            //     { text: 'Apply Date', value: 'applyDate', align: "center", sortable: true },
            //     { text: 'Applied Time', value: 'appliedTime', align: "center", sortable: false },
            //     { text: 'ClientId', value: 'clientId', align: "center", sortable: false },
            //     { text: 'Category', value: 'category', align: "center", sortable: false },

            //     { text: 'Status', value: 'status', align: "center", sortable: false },

            // ],
            ShowReportRec: false,
            ReportRec: {},
        }
    },
    methods: {
        // to resize the screen
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
                        this.ReportRec = response.data;
                        for (let ix = 0; ix < this.ReportRec.modifyDetails.length; ix++) {
                            this.ReportRec.modifyDetails[ix].signal = 'O'
                        }
                        this.ShowReportRec = true;
                        this.$globalData.overlay = false
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
    computed: {
        isSmallScreen() {
            return this.screenWidth <= 700;
        },

    },
}
</script>

<style scoped>
.text-large {
    font-size: 15px;
}
</style>

<style scoped>
.text {
    font-size: 10px;
}

::v-deep .v-data-table__mobile-row__cell {
    width: 100% !important;
}

::v-deep .v-data-table__mobile-row__cell table,
::v-deep .v-data-table__mobile-row__cell table tr,
::v-deep .v-data-table__mobile-row__cell table td {
    width: 100%
}

::v-deep .v-data-table__mobile-row__cell table tr td:first-child {
    text-align: left;
}

::v-deep .v-data-table__mobile-row__cell table tr td:last-child {
    text-align: right;
}

::v-deep .v-data-footer {
    justify-content: end;
}

::v-deep .v-data-table__mobile-row__cell table tr {
    padding-bottom: 5px;
}

.row {
    margin-top: 10px;
    margin-bottom: 10px;
}

::v-deep .v-data-table>.v-data-table__wrapper .v-data-table__mobile-row {
    height: initial;
    min-height: 2px;
}

/* ::v-deep .v-data-footer__pagination {
    display: none;
} */
</style>