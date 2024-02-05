<template>
    <div>
        <!-- header -->
        <v-layout row wrap class="mb-5 mt-3" style="padding-left: 12px;">
            <div class="d-flex align-center">
                <v-img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/tresur.png" height="30" width="25" contain
                    class="d-flex"></v-img>
                <span class="font-weight-bold" style="margin-left: 13px;">T-Bills</span>
            </div>
        </v-layout>

        <!-- values -->
        <v-layout class="my-10" v-if="billdynamicItem.length == 0">
            <v-slide-x-transition mode="out-in" appear>
                <v-flex class="ml-5 text text--disabled">
                    <!-- <p v-if="billsresult.length == 0 && Flag == 'I'">No T-Bills are open for sale currently.</p> -->
                    <!-- <p v-if="billshistory.length == 0 && Flag == 'O'">You haven't invested in any T-Bills.</p> -->
                    <p>{{ billdynamicText }}</p>
                </v-flex>
            </v-slide-x-transition>
        </v-layout>


        <v-card class="elevation-0" v-if="billdynamicItem.length != 0">
            <v-card-title>
                <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line
                    hide-details></v-text-field>
            </v-card-title>

            <v-slide-x-transition mode="out-in" appear>


                <v-data-table :headers="billdynamicHeader" :items="billdynamicItem" :search="search" :items-per-page="10"
                    :key="tableKey"
                    :footer-props="Flag == 'O' ? { 'items-per-page-options': [5] } : { 'items-per-page-options': [5] }"
                    :loading="loading" loading-text="Loading... Please wait" no-data-text="No Records available">
                    <template v-slot:item.symbol="{ item }" v-if="billsresult != [] && Flag == 'I'">


                        <v-row wrap class="d-flex text font-weight-bold">
                            <v-col cols="6" class="d-flex flex-column text-left">
                                <span>{{ item.symbol }}</span>
                            </v-col>
                        </v-row>
                        <v-row wrap class="d-flex text mt-3">
                            <v-col cols="6" class="d-flex flex-column text-left">
                                <span class="text--disabled">Indicative yield</span>
                                <span>{{ item.name }}</span>
                            </v-col>

                            <v-col class="d-flex flex-column text-right" cols="6">
                                <span class="text--disabled">Bid Close date</span>
                                <span>{{ item.endDateWithTime }}</span>
                            </v-col>
                        </v-row>
                        <v-row wrap class="text">
                            <v-col class="d-flex text-left" cols="6">
                                <v-layout>
                                    <v-flex class="d-flex flex-column">
                                        <span class="text--disabled">Unit Limits</span>
                                        <span>{{ item.priceRange }}</span>

                                    </v-flex>

                                </v-layout>
                            </v-col>
                            <v-col class="d-flex flex-column text-right" cols="6">
                                <span class="text--disabled">Amount</span>
                                <v-col>
                                    <span v-if="item.amount !== 0">₹ </span>
                                    <span> {{ item.amount }}</span>
                                </v-col>
                            </v-col>
                        </v-row>
                        <v-row wrap class="text">
                            <v-col class="text-center" cols="12">
                                <v-hover v-slot="{ hover }">
                                    <v-btn small :disabled="item.diableActionBtn == undefined ? true : item.diableActionBtn"
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
                            </v-col>
                        </v-row>
                    </template>

                    <template v-slot:item.symbol="{ item }" v-else>
                        <v-row wrap class="d-flex text">
                            <v-col cols="6" class="d-flex flex-column text-left">
                                <span class="text--disabled">Security Name</span>
                                <span>{{ item.symbol }}</span>
                            </v-col>
                            <v-col class="d-flex flex-column text-right text" cols="6">
                                <span class="text--disabled">Ordered Date</span>
                                <span>{{ item.orderDate }}</span>
                            </v-col>
                        </v-row>
                        <v-row wrap class="text">
                            <v-col class="d-flex text-left" cols="6">
                                <v-layout>
                                    <v-flex class="d-flex flex-column">
                                        <span class="text--disabled">Int.RefNo</span>
                                        <span>{{ item.respOrderNo }}</span>
                                        <span class="text--disabled">OrderNo</span>
                                        <span>{{ item.orderNo }}</span>
                                        <span class="text--disabled">Unit price</span>
                                        <span>{{ item.requestedUnitPrice}}<span>
                                                <v-menu top offset-x>
                                                    <template v-slot:activator="{ on, attrs }">
                                                        <v-icon right small @click="displayDetail" v-bind="attrs" v-on="on"
                                                            color="primary">mdi-information-outline</v-icon>
                                                    </template>
                                                    <detailCard :master="item" :showDetail="showDetail"
                                                        @closeDetail="displayDetail" />
                                                </v-menu>
                                            </span>
                                        </span>
                                    </v-flex>

                                </v-layout>
                            </v-col>
                            <v-col class="d-flex justify-space-between flex-column text-right" cols="6">
                                <span class="text--disabled">Status</span>
                                <span
                                    :class="item.statusColor == 'G' ? 'text-capitalize success--text' : 'text-capitalize error--text darken-5'">
                                    {{ item.orderStatus }}
                                </span>

                                <v-layout class="d-flex justify-space-between flex-column">
                                    <v-flex class="d-flex flex-column">
                                        <span class="text--disabled">Units</span>
                                        <span>{{ item.requestedUnit }}</span>
                                    </v-flex>
                                    <v-flex class="d-flex flex-column">
                                        <span class="text--disabled">Total</span>
                                        <span>₹{{ formatedPrice(item.requestedAmount) }}</span>
                                    </v-flex>
                                </v-layout>
                            </v-col>
                        </v-row>
                    </template>
                </v-data-table>
            </v-slide-x-transition>
        </v-card>

    </div>
</template>
<script>
import detailCard from "../../Tab/detailCard.vue"
export default {
    name: "TBillMobile",
    data() {
        return {
            search: '',
            header: [],
            InvestMobile: [{ text: "", sortable: false, align: "center", value: "symbol" },],
            OrderMobile: [{ text: "", sortable: false, align: "center", value: "symbol" },],
            showDetail: false,
        }
    },
    methods: {
        formatedPrice(item) {
            if (item != undefined) {
                return item.toLocaleString('en-IN')
            }
        },
        sendTo(item, indicator) {
            this.$emit("passVal", item, indicator)
        },
        buttonText(item) {
            if (item != undefined) {
                return this.$globalData.currentTime >= item.startTime && this.$globalData.currentTime <= item.endTime ? "PlaceOrder" : "Offline"
            } else {
                return "PlaceOrder"
            }
        },
        Modbtn(item) {
            if (item != undefined) {
                return (this.$globalData.currentTime <= item.startTime && this.$globalData.currentTime >= item.endTime)
            }
        },
        displayDetail() {
            this.showDetails = !this.showDetails
        },
        // show(flag) {
        //     // this.$emit("showCircle");
        //     // setTimeout(() => {
        //         if (flag == "I") {
        //             this.Flag = flag;
        //             if (this.$vuetify.breakpoint.width < 600) {
        //                 this.header = this.InvestMobile;
        //             }
        //             this.store = this.billsresult;

        //         } else {
        //             this.Flag = flag;
        //             if (this.$vuetify.breakpoint.width < 600) {
        //                 this.header = this.OrderMobile;
        //             }
        //             this.store = this.billshistory;
        //         }
        //         // this.$emit("showCircle");
        //     // }, 250);
        // },
    },
    computed: {
        billdynamicHeader() {
            return this.Flag == "I" ? this.InvestMobile : this.OrderMobile
        },
        billdynamicItem: {
            get() {
                if (this.Flag == "I" && this.tbillmasterFound == "Y") {
                    return this.billsresult
                } else if (this.Flag == "O" && this.tbillhistoryFound == "Y") {
                    return this.billshistory
                } else {
                    return []
                }
            }
            // return this.Flag == "I" ? this.billsresult : this.billshistory
        }
    },
    props: {
        loading: Boolean,
        billsresult: Array,
        billshistory: Array,
        Flag: String,
        tableKey: Number,
        // btnText: String,
        // Modbtn: Boolean,
        billdynamicText: String,
        tbillmasterFound: String,
        tbillhistoryFound: String,
    },
    components: {
        detailCard
    },
    // mounted() {
    //     this.show(this.Flag);
    // },
    // watch: {
    //     Flag: function (flag) {
    //         this.show(flag);
    //     },
    //     tableKey: {
    //         handler() {
    //             this.show(this.Flag);
    //         },
    //         immediate: true,
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