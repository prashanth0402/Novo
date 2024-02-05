<template>
    <div>
        <v-layout class="my-10" v-if="dynamicItem.length == 0">
            <v-slide-x-transition mode="out-in" appear>
                <v-flex class="ml-5 text text--disabled">
                    <!-- <p v-if="result.length == 0 && Flag1 == 'I'">No SGBs are open for sale currently.</p>
                    <p v-if="history.length == 0 && Flag1 == 'O'">You haven't invested in any SGBs.</p> -->
                    <!-- <p v-if="dynamicItem.length == 0">
                        {{ Flag == 'I' ? "No SGBs are open for sale currently." : "You haven't invested in any SGBs." }}
                    </p> -->
                    <p>{{ dynamicText }}</p>

                </v-flex>
            </v-slide-x-transition>
        </v-layout>
        <v-card elevation="0" v-if="dynamicItem.length != 0">
            <v-card-title v-if="Flag == 'O'">
                <v-text-field v-model="search" dense append-icon="mdi-magnify" label="Search" single-line></v-text-field>
            </v-card-title>
            <v-slide-x-transition mode="out-in" appear>
                <v-card-text>
                    <v-data-table :headers="dynamicHeader" :items="dynamicItem" :search="search" :key="tableKey"
                        :footer-props="Flag == 'O' ? { 'items-per-page-options': [5, 10, 15, -1] } : { 'items-per-page-options': [10] }"
                        :loading="loading" loading-text="Loading... Please wait" no-data-text="No Records available">

                        <!-- ===================Order History Changes -->
                        <template v-if="Flag == 'O'" v-slot:item.name="{ item }">
                            <v-row>
                                <v-col class="d-flex flex-column text-left">
                                    <span class="mb-1">{{ item.name }}</span>
                                    <span class="caption">Ordered Date: {{ item.orderDate }}</span>
                                </v-col>
                            </v-row>
                        </template>
                        <template v-if="Flag === 'O'" v-slot:item.unit="{ item }">
                            <v-menu top offset-x>
                                <template v-slot:activator="{ on, attrs }">
                                    <span>{{ item.requestedUnit }}</span>
                                    <v-icon right small @click="displayDetail" v-bind="attrs" v-on="on"
                                        color="primary">mdi-information-outline</v-icon>
                                </template>
                                <detailCard :master="item" :showDetail="showDetail" @closeDetail="displayDetail" />
                            </v-menu>
                        </template>
                        <!-- =========================================================== -->

                        <template v-slot:item.price="{ item }">
                            <span> {{ formatedPrice(item.price) }}</span>
                        </template>
                        <template v-slot:item.minPrice="{ item }">
                            <span>{{ formatedPrice(item.minPrice) }}</span>
                        </template>
                        <template v-slot:item.status="{ item }">
                            <span :class="item.statusColor == 'G' ? 'green--text text-capitalize' : 'red--text'">
                                {{ item.orderStatus }}
                            </span>

                        </template>
                        <template v-slot:item.actions="{ item }" class="elevation-0">
                            <v-layout>
                                <v-flex class="d-flex align-center">
                                    <v-hover v-slot="{ hover }">
                                        <v-btn small
                                            :disabled="item.disableActionBtn == undefined ? true : item.disableActionBtn"
                                            width="100"
                                            v-if="!(item.actionFlag == '') && !(item.actionFlag == undefined) && !(item.buttonText == '') && !(item.buttonText == undefined)"
                                            :class="hover ? 'secondary' : item.actionFlag == 'M' || item.actionFlag == 'A' || item.actionFlag == 'U' || item.actionFlag == 'C' ? 'blue lighten-4 primary--text' : 'primary white--text'"
                                            @click="item.actionFlag == 'B' || item.actionFlag == 'P' ? sendTo(item, 'N') : sendTo(item, 'M')"
                                            elevation="0">
                                            <span class="text-capitalize">
                                                {{ item.buttonText }}
                                            </span>
                                        </v-btn>
                                    </v-hover>
                                </v-flex>
                            </v-layout>
                        </template>
                        <template v-slot:item.total="{ item }">
                            <span>₹ {{ formatedPrice(item.requestedAmount) }}</span>
                        </template>
                    </v-data-table>
                </v-card-text>
            </v-slide-x-transition>

        </v-card>
    </div>
</template>

<script>
import detailCard from "../Tab/detailCard.vue"
export default {
    name: "SgbMaster",
    data() {
        return {
            search: '',
            header: [],
            Invest: [
                {
                    text: 'Tranche',
                    align: 'start',
                    sortable: false,
                    value: "name",
                },
                { text: 'Subscription period', sortable: false, align: "start", value: 'dateRange' },
                { text: 'Unit Price', sortable: false, align: "start", value: 'unitPrice' },
                { text: 'Ordered Unit', sortable: false, align: "start", value: 'appliedUnit' },
                { text: "", sortable: false, align: "start", value: "actions" },
            ],
            Order: [
                {
                    text: 'Tranche',
                    sortable: true,
                    align: 'start',
                    value: 'name'
                },
                { text: 'Int.RefNo', sortable: true, align: "center", value: 'orderNo' },
                { text: 'Exch OrderNo.', sortable: false, align: "center", value: 'exchOrderNo' },
                { text: 'Unit Price', sortable: false, align: "start", value: 'requestedUnitPrice' },
                { text: 'Units', sortable: true, align: "center", value: 'unit' },
                { text: 'Amount Payable', sortable: true, align: "start", value: 'total' },
                { text: 'Order status', sortable: false, align: "start", value: 'status' },
            ],
            tabs: ['I', 'O'],
            showDetail: false,
        }
    },
    methods: {
        // This method is used to show the popup and pass the row details to the popup
        sendTo(item, indicator) {
            this.$emit("passVal", item, indicator)
        },
        formatedPrice(item) {
            if (item != undefined) {
                return item.toLocaleString('en-IN');
            }
        },
        // buttonText(item) {
        //     if (item != undefined) {
        //         return this.$globalData.currentTime >= item.startTime && this.$globalData.currentTime <= item.endTime ? "PlaceOrder" : "Offline"
        //     } else {
        //         return "PlaceOrder"
        //     }
        // },
        // Modbtn(item) {
        //     if (item != undefined) {
        //         return (this.$globalData.currentTime <= item.startTime && this.$globalData.currentTime >= item.endTime)
        //     }
        // },
        displayDetail() {
            this.showDetails = !this.showDetails
        },

        //! Use this method incase you need the row value when click on the row
        // handleRowClick(item) {
        //     if (this.Flag == 'O') {
        //         if (window.getSelection().toString().length === 0) {
        //             this.sendTo(item, "R")
        //         }
        //     }
        // }, 
    },
    props: {
        loading: Boolean,
        result: Array,
        history: Array,
        Flag: String,
        tableKey: Number,
        // btnText: String,
        // Modbtn: Boolean,
        dynamicText: String,
        masterFound: String,
        historyFound: String,
    },
    computed: {
        dynamicHeader() {
            return this.Flag == "I" ? this.Invest : this.Order
        },
        dynamicItem: {
            get() {
                if (this.Flag == "I" && this.masterFound == "Y") {
                    return this.result
                } else if (this.Flag == "O" && this.historyFound == "Y") {
                    return this.history
                } else {
                    return []
                }
            }
        },
    },
    components: {
        detailCard
    }
}
</script>

<style scoped>
.text {
    font-size: 1em;
    font-weight: 400;
}


::v-deep .v-data-table>.v-data-table__wrapper .v-data-table__mobile-row {
    height: initial;
    min-height: 10px;
}

::v-deep .v-data-table>.v-data-table__wrapper>table>tbody>tr>td {
    height: 55px;
}

/* ::v-deep .v-data-table > .v-data-table__wrapper > table > tbody > tr > td:last-child {
   width: 1%;
} 
::v-deep .v-data-table__mobile-row__cell {
    width: 100% !important;
} */
/* 
::v-deep .v-data-table__mobile-row__cell table tr td:first-child {
    text-align: left;
}

::v-deep .v-data-table__mobile-row__cell table tr td:last-child {
    text-align: right;
} */

::v-deep .v-data-footer {
    justify-content: end;
    /* padding: 0;
  padding-left: 200px; */
}

/* ::v-deep .v-data-table__mobile-row__cell table tr {
    padding-bottom: 3px;
}

.row {
    margin-top: 2px;
    margin-bottom: 2px;
}

::v-deep .v-data-table>.v-data-table__wrapper .v-data-table__mobile-row {
    height: initial;
    min-height: 2px;
} */
</style>