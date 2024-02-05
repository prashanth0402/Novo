<template>
    <div>
        <v-dialog v-model="dialog" max-width="650" persistent :fullscreen="this.$vuetify.breakpoint.name == 'xs'"
            :transition="this.$vuetify.breakpoint.name == 'xs' ? 'dialog-bottom-transition' : undefined"
            overlay-color="#fff">
            <v-card outlined>

                <v-form ref="form" lazy-validation>
                    <v-card-title>
                        <v-row>
                            <v-col cols="12" xl="10" lg="10" md="10" sm="12" xs="12">
                                <v-layout class="d-flex flex-column justify-space-between">
                                    <v-flex class="d-flex align-center">
                                        <v-img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/sgblogo.webp"
                                            height="35" width="35" contain class="mr-2" xs2></v-img>
                                        <span class="font-weight-normal caption" xs8>{{
                                            detail.name }}</span>
                                        <!-- &nbsp; <span v-if="Action == 'N' && !slide1" class="font-weight-normal caption">- Order
                                            confirmation</span>
                                        <span v-if="Action == 'M' && !slide1" class="font-weight-normal caption">- Modify
                                            Order</span>
                                        <span v-if="Action == 'D' && !slide1" class="font-weight-normal caption">- Cancel
                                            Order</span> -->
                                    </v-flex>
                                    <span class="text primary--text"
                                        v-if="detail.infoText != '' && detail.infoText != undefined && Action != 'R'">
                                        <v-icon x-small left color="primary">mdi-information-outline</v-icon>
                                        {{ detail.infoText }}
                                    </span>
                                </v-layout>
                            </v-col>
                            <v-col cols="12" xl="2" lg="2" md="2" sm="12" xs="12" v-if="slide1 && !progress">
                                <v-flex class="d-flex justify-end">
                                    <v-btn class="red lighten-1 btn caption text-capitalize" dark height="25" elevation="0"
                                        v-show="detail.cancelAllowed == '' || detail.cancelAllowed == undefined ? false : detail.cancelAllowed"
                                        @click="ChangeActionFlag('D')" :disabled="disable">
                                        <!-- <v-icon small>mdi-trash-can-outline</v-icon> -->
                                        cancel Bid
                                    </v-btn>
                                </v-flex>
                            </v-col>
                        </v-row>
                    </v-card-title>

                    <v-divider></v-divider>

                    <!-- This progress circle used for UI presentaion and help to avoid showing value jumping -->
                    <v-layout v-if="progress" class="mb-10">
                        <v-flex class="d-flex justify-center">
                            <v-progress-circular indeterminate color="primary" size="50"></v-progress-circular>
                        </v-flex>
                    </v-layout>
                    <!-- This section contains all the details of the SGB -->
                    <section v-else>
                        <div v-if="slide1">
                            <v-card-text>
                                <v-row v-if="slide1 && Action != 'D'">
                                    <v-col class="d-flex flex-column justify-start" v-if="Action != 'R'" cols="12" xl="4"
                                        lg="4" md="4" sm="4" xs="12">
                                        <span><b>Current Ledger Balance</b></span>
                                        <span>₹
                                            <span>
                                                <img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/load.gif"
                                                    v-show="image">
                                            </span>
                                            <span :class="color" v-show="!image">
                                                {{ accountbal == 0 ? 0 : accountbal }}
                                            </span>
                                        </span>
                                    </v-col>
                                    <v-col class="d-flex flex-column justify-start" cols="12" xl="4" lg="4" md="4" sm="4"
                                        xs="6">
                                        <span class="text--disabled">Unit Price</span>
                                        <span>{{ detail.unitPrice }}</span>
                                    </v-col>
                                    <v-col class="d-flex flex-column justify-start" cols="12" xl="4" lg="4" md="4" sm="4"
                                        xs="6">
                                        <span class="text--disabled">{{ detail.discountText }}</span>
                                        <span>{{ detail.discountAmt }}</span>
                                    </v-col>
                                </v-row>
                            </v-card-text>
                            <v-divider></v-divider>
                            <v-card-text style="background-color:#fff9a0 ;">
                                <!-- <v-card-text> -->

                                <v-row>
                                    <v-col class="d-flex flex-column justify-start" cols="12" xl="4" lg="4" md="4" sm="5"
                                        xs="12">
                                        <span class="text--disabled">Bidding Starts</span>
                                        <span>{{ detail.startDateWithTime }}</span>
                                    </v-col>
                                    <v-col cols="12" xl="3" lg="3" md="3" sm="2"
                                        v-if="$vuetify.breakpoint.name != 'xs'"></v-col>
                                    <v-col class="d-flex flex-column align-start" cols="12" xl="4" lg="4" md="4" sm="5"
                                        xs="12">
                                        <div v-if="detail.modifyAllowed == '' && detail.modifyAllowed != undefined ? true : !detail.modifyAllowed"
                                            class="d-flex flex-column justify-start align-start">
                                            <span class="text--disabled">Applied unit</span>
                                            <span>{{ detail.appliedUnit }}</span>
                                        </div>
                                        <div v-else>
                                            <span>Units to buy</span>
                                            <v-text-field background-color="#fff" v-model.number="sgb.unit" outlined dense
                                                type="number" :min="parseInt(detail.minBidQty)"
                                                :max="parseInt(detail.maxBidQty)" @keypress="onlyForNumber"
                                                :rules="[customValidation]" width="100%" :error-messages="errText"
                                                :error="errVal"
                                                :disabled="detail.modifyAllowed == '' && detail.modifyAllowed != undefined ? true : !detail.modifyAllowed">
                                            </v-text-field>
                                        </div>
                                    </v-col>
                                    <v-col class="d-flex flex-column justify-start" cols="12" xl="4" lg="4" md="4" sm="4"
                                        xs="12">
                                        <span class="text--disabled">Bidding Ends</span>
                                        <span>{{ detail.endDateWithTime }}</span>
                                    </v-col>
                                    <v-col class="d-flex flex-column justify-start" cols="12" xl="4" lg="4" md="4" sm="4"
                                        xs="12">
                                        <span class="text--disabled">ISIN</span>
                                        <span>{{ detail.isin }}</span>
                                    </v-col>
                                    <v-col class="d-flex flex-column justify-start" cols="12" xl="4" lg="4" md="4" sm="4"
                                        xs="12" v-if="detail.orderNo != ''">
                                        <span class="text--disabled">OrderNo</span>
                                        <span>{{ detail.orderNo }}</span>
                                    </v-col>
                                </v-row>
                            </v-card-text>
                        </div>
                        <div v-else-if="slide2">
                            <v-card-text>
                                <v-layout class="d-flex flex-column mb-2">
                                    <v-flex class="d-flex align-center flex-column mb-2">
                                        <span class="text--disabled">Investment
                                            quantity</span>
                                        <span class="font-weight-bold">{{ sgb.unit }}
                                            units</span>
                                    </v-flex>
                                    <v-flex class="d-flex align-center"
                                        v-if="Action != 'D' && (detail.showSI == '' ? false : detail.showSI == undefined ? false : detail.showSI)">
                                        <v-checkbox v-model="disclaimerCheckBox" :rules="refRule" dense xl2 lg2 sm2 md2
                                            xs5></v-checkbox>
                                        <span class=" mt-1 d-block text-wrap pb-4" xl10 lg10 sm10 md10 xs7>
                                            {{ detail.SItext }}
                                        </span>
                                    </v-flex>
                                    <v-flex v-else-if="Action == 'D'" class="text-center">
                                        <span>{{ detail.SIrefundText }}</span>
                                    </v-flex>
                                </v-layout>
                            </v-card-text>
                        </div>
                        <v-divider></v-divider>
                        <v-card-text>
                            <v-row>
                                <v-col cols=12 xl="7" lg="7" md="6" sm="12" xs="12">
                                    <v-row d-flex justify-space-evenly v-if="slide1">
                                        <v-col class="d-flex flex-column align-end" cols="12" xl="4" lg="4" md="4" sm="4"
                                            xs="6">
                                            <span class="text--disabled">Total.</span>
                                            <span v-if="Action != 'R'">₹ {{ (sgb.unit * (sgb.price +
                                                parseInt(detail.discountAmt))).toLocaleString('en-IN')
                                            }}</span>
                                            <span v-else>{{ (parseInt(detail.appliedUnit) * (parseInt(detail.unitPrice) +
                                                parseInt(detail.discountAmt))).toLocaleString('en-IN') }}</span>
                                        </v-col>
                                        <v-col class="d-flex flex-column align-end" cols="12" xl="4" lg="4" md="4" sm="4"
                                            xs="6">
                                            <span class="text--disabled">Discount.</span>
                                            <span v-if="Action != 'R'">₹ {{ ((sgb.unit * (sgb.price +
                                                parseInt(detail.discountAmt))) - amount).toLocaleString('en-IN')
                                            }}</span>
                                            <span v-else>{{
                                                (detail.appliedUnit * (detail.unitPrice + detail.discountAmt) -
                                                    (detail.appliedUnit
                                                        *
                                                        detail.unitPrice)) }}</span>
                                        </v-col>
                                        <v-col class="d-flex flex-column align-end" cols="12" xl="4" lg="4" md="4" sm="4"
                                            xs="6">
                                            <span class="text--disabled">Amt payable.</span>
                                            <span class="success--text font-weight-bold">₹
                                                {{ Action != 'R' ?
                                                    amount.toLocaleString('en-IN') : (detail.appliedUnit *
                                                        detail.unitPrice).toLocaleString('en-IN')
                                                }}</span>
                                        </v-col>
                                    </v-row>
                                </v-col>

                                <v-col class="d-flex align-center justify-end" cols=12 xl="5" lg="5" md="6" sm="12" xs="12">
                                    <v-btn icon small class="secondary white--text mr-2" @click="back"
                                        v-if="slide2 != false">
                                        <v-icon>mdi-arrow-left</v-icon>
                                    </v-btn>
                                    <v-btn class="text-capitalize white--text elevation-0" @click="placeSgb"
                                        v-if="Action != 'R'" :loading="btnLoad" :disabled="isFilled && Action != 'D'"
                                        color="#4184f3" small>Confirm</v-btn>
                                    <v-btn class="text-capitalize elevation-0 ml-2" outlined text small
                                        @click="closePop">Close</v-btn>
                                </v-col>
                            </v-row>
                        </v-card-text>
                    </section>
                    <!-- </v-card-text> -->
                </v-form>
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
import Eventservice from "@/services/EventServices.js"
export default {
    data() {
        return {
            color: '',
            accountbal: 0,
            amount: 0,
            sgb: {
                masterId: 0,
                bidId: "",
                unit: 0,
                price: 0,
                actionCode: "",
                orderNo: "",
                amount: 0,
                oldUnit: 0,
                SIvalue: false,
                SItext: ""
            },
            disable: false,
            inputValue: "",
            customValidation: (v) => !!v || 'Unit must be greater than 0',
            refRule: [(v) => !!v || "required"],
            disclaimerCheckBox: false,
            slide1: false,
            slide2: false,
            Item: ["Online", "Offline"],
            errText: "",
            errVal: false,
            btnLoad: false,
            progress: false,
            image: false,
            counter: 0,
            temp: {},
            CopyAction: "", // To store the Previous ActionFlag
            // CalcAmt: 0,
        }
    },
    props: {
        detail: {},
        // modify: {},
        Action: String,
        dialog: Boolean
    },
    computed: {
        isFilled() {
            return this.sgb.unit == parseInt(this.detail.appliedUnit)
        },

    },
    methods: {
        // This method is to check the enter values are number
        onlyForNumber($event) {
            let keyCode = $event.keyCode ? $event.keyCode : $event.which;

            if (
                (keyCode < 48 || keyCode > 57) &&
                (keyCode !== 46 || keyCode == 46) &&
                (keyCode !== -46 || keyCode == -46)
            ) {
                $event.preventDefault();
            }
        },
        formatedPrice(item) {
            if (item != undefined) {
                return item.toLocaleString('en-IN')
            }
        },
        placeSgb() {
            // if (this.accountbal >= this.CalcAmt) {
            //     console.log("more than", this.Action);

            // } else {
            //     console.log("less than");
            // }
            if (this.Action != "D") {
                if (this.$refs.form.validate()) {
                    if (this.errVal == false) {
                        this.counter++;
                    }
                    if (this.Action == "N") {
                        // console.log("Counter in Action N", this.counter);
                        // if (this.accountbal >= this.amount) {
                        if (this.counter == 1) {
                            if (this.sgb.unit >= parseInt(this.detail.minBidQty) && this.errVal != true) {
                                this.slide1 = false;
                                this.slide2 = true
                            }
                            this.temp = this.sgb
                        } else if (this.counter == 2) {
                            if (this.slide2 == true && this.counter == 2) {
                                this.construct()
                                this.counter = 0;
                            }
                        }
                        // }
                        // else {
                        //     this.counter = 0;
                        //     this.MessageBar("E", "You don't have sufficient fund balance in your trading account. Please make sufficient fund transfer to intiate SGB request")
                        // }
                    } else if (this.Action == "M") {
                        // console.log("Counter in Action M", this.counter);
                        // if (this.accountbal >= this.CalcAmt) {
                        if (this.counter == 1) {
                            if (this.errVal != true) {
                                this.slide1 = false;
                                this.slide2 = true;
                            }
                            this.temp = this.sgb
                        } else if (this.counter == 2) {
                            if (this.slide2 == true && this.counter == 2) {
                                this.construct()
                                this.counter = 0;
                            }
                        }
                        // } else {
                        //     this.counter = 0;
                        //     this.MessageBar("E", "You don't have sufficient fund balance in your trading account. Please make sufficient fund transfer to intiate SGB request")
                        // }
                    }
                }
            } else {
                this.sgb.masterId = parseInt(this.detail.id);
                this.sgb.actionCode = this.Action;
                this.sgb.bidId = this.detail.bidId;
                this.sgb.orderNo = this.detail.orderNo;
                this.sgb.price = parseInt(this.detail.unitPrice) - parseInt(this.detail.discountAmt);
                this.sgb.amount = parseInt(this.detail.unitPrice) * parseInt(this.detail.appliedUnit);
                this.sgb.unit = parseInt(this.detail.appliedUnit);
                // this.sgb.oldUnit = parseInt(this.detail.unit);
                this.sgb.oldUnit = parseInt(this.detail.appliedUnit);
                this.sgb.SIvalue = this.detail.SIvalue == '' ? false : this.detail.SIvalue;
                this.sgb.SItext = this.detail.SItext;
                this.CallSgbPlace();
            }
        },
        construct() {
            this.sgb = this.temp // assign the copy struct
            this.sgb.masterId = this.detail.id;
            this.sgb.actionCode = this.Action;
            this.sgb.bidId = this.detail.bidId;
            this.sgb.orderNo = this.detail.orderNo;
            this.sgb.amount = this.detail.total;
            this.sgb.oldUnit = this.detail.appliedUnit;
            this.sgb.SItext = this.detail.SItext;
            if (this.Action == 'N') {
                this.sgb.SIvalue = this.disclaimerCheckBox;
            } else {
                this.sgb.SIvalue = this.detail.SIvalue == '' ? false : this.detail.SIvalue;
            }
            this.CallSgbPlace();
        },
        CallSgbPlace() {
            this.disable = true;
            this.$globalData.overlay = true;
            this.btnLoad = true;
            Eventservice.SGBPlaceOrder(this.sgb)
                .then((response) => {
                    this.$globalData.overlay = false;
                    this.btnLoad = false;
                    this.disable = false;
                    if (response.data.status == "S") {
                        this.MessageBar("S", response.data.orderStatus)
                        this.disclaimerCheckBox = false;
                        this.sgb.unit = 0;
                        this.slide1 = false;
                        this.slide2 = false;
                        this.$emit('closeSgbPop');
                        this.$emit("RecallSgb")
                        this.$refs.form.resetValidation();
                        // location.reload();
                    } else if (response.data.status == "E") {
                        if (response.data.orderStatus == "" || response.data.orderStatus == undefined) {
                            this.MessageBar("E", response.data.errMsg)
                        } else {
                            this.MessageBar("E", response.data.orderStatus)
                        }
                    }
                })
                .catch((error) => {
                    this.disable = false;
                    this.btnLoad = false;
                    this.MessageBar("E", error)
                });
            this.disclaimerCheckBox = false;
            this.sgb.unit = 0;
            this.slide1 = false;
            this.slide2 = false;
            this.$emit('closeSgbPop');
            this.$refs.form.resetValidation();
        },
        closePop() {
            this.slide1 = false;
            this.slide2 = false;
            this.sgb.id = 0;
            this.sgb.unit = 0;
            this.sgb.price = 0;
            this.counter = 0;
            this.inputValue = "";
            this.accountbal = 0;
            this.$emit('closeSgbPop');
            this.$emit('EmptyModify');
            this.$refs.form.resetValidation();

        },
        // incrementValue() {
        //     this.sgb.unit++;
        // },
        // decrementValue() {
        //     if (this.sgb.unit > this.detail.minBidQty) {
        //         this.sgb.unit--;
        //     }
        // },
        back() {
            this.slide1 = true;
            this.slide2 = false;
            this.disclaimerCheckBox = false;
            this.counter = 0;
            if (this.Action == "D") {
                this.$emit("ChangeActionFlag", this.CopyAction)
            }
            // this.$refs.form.resetValidation();
        },
        // choose(item) {
        //     if (item == "Online") {
        //         this.sgb.price = parseInt(this.detail.minPrice);
        //     } else {
        //         this.sgb.price = parseInt(this.detail.maxPrice);
        //     }
        // },
        ChangeActionFlag(action) {
            this.slide1 = !this.slide1;
            this.slide2 = true;
            this.$emit("ChangeActionFlag", action)
        },
        FetchClientFund() {
            this.image = true;
            if (this.Action != "D" && this.Action != "R") {
                Eventservice.FetchFund()
                    .then((response) => {
                        this.image = false
                        if (response.data.status == "S") {
                            this.accountbal = parseInt(response.data.accountBalance)
                        }
                    })
                    .catch((error) => {
                        this.MessageBar("E", error)
                    });
            }
        }
    },
    watch: {
        sgb: {
            handler(val) {
                if (this.Action != "D" && this.Action != "R") {
                    this.amount = 0;
                    // this.amount = this.amount + (val.price - this.detail.discountAmt) * val.unit;

                    //  The price below excludes after the discount price
                    this.amount = val.price * val.unit;
                    // this.CalcAmt = (this.amount - this.detail.total)

                    if (val.unit <= parseInt(this.detail.maxBidQty) && val.unit >= parseInt(this.detail.minBidQty)) {
                        this.errVal = false;
                        this.errText = "";
                    } else if (val.unit < parseInt(this.detail.maxBidQty)) {
                        this.errText = "Min. bidQty " + this.detail.minBidQty;
                        this.errVal = true;
                    } else if (val.unit > parseInt(this.detail.minBidQty)) {
                        this.errText = "Max. bidQty " + parseInt(this.detail.maxBidQty);
                        this.errVal = true;
                    }
                    if (this.Action == 'N' || this.Action == 'M') {
                        if (this.accountbal >= this.amount) {
                            this.color = 'primary--text'
                        } else {
                            this.color = 'error--text'
                        }
                    }
                    // else if (this.Action == 'M') {
                    // if (this.accountbal >= this.CalcAmt) {
                    // this.accountbal = this.accountbal - this.CalcAmt
                    // this.color = 'success--text'
                    // } else {
                    // this.accountbal = this.accountbal + this.CalcAmt
                    // this.color = 'error--text'
                    // }
                    // }
                }
            },
            deep: true,
            // immediate: true,
        },
        dialog: function (bool) {
            if (bool == true) {
                this.slide1 = true;
                this.progress = true;
                this.disclaimerCheckBox = this.detail.SIvalue == '' ? false : this.detail.SIvalue;
                this.FetchClientFund();
                setTimeout(() => {
                    if ('id' in this.detail && this.Action == "M" || this.Action == "A" || this.Action == "R") {
                        // if (this.Action == "M") {
                        this.sgb.unit = parseInt(this.detail.appliedUnit);
                        this.sgb.price = parseInt(this.detail.unitPrice) - parseInt(this.detail.discountAmt);

                    } else {
                        this.sgb.unit = parseInt(this.detail.minBidQty);
                        this.sgb.price = parseInt(this.detail.unitPrice) - parseInt(this.detail.discountAmt);
                    }
                    this.progress = false;
                }, 500); // Change the time delay to your desired value in milliseconds
            } else {
                this.$emit('EmptyModify');
            }
        },
        Action: {
            immediate: true,
            handler(value, oldVal) {
                if (value == "D") {
                    this.CopyAction = oldVal;
                    this.slide1 = false;
                } else if (value == "M") {
                    this.sgb.price = parseInt(this.detail.unitPrice)
                    this.CopyAction = oldVal;
                }
            }
        }
    },

}
</script>

<style scoped>
.v-btn:not(.v-btn--round).v-size--default {
    height: 36px;
    min-width: 64px;
    padding: 0px 25px;
    margin-right: 10px;
}

.v-dialog>.v-card>.v-card__title {
    padding: 16px 10px 10px;
}

::v-deep .v-messages__message {
    width: 100%;
    font-size: 10px;
}

::v-deep .v-text-field input {
    flex: 1 1 auto;
    line-height: 20px;
    padding: 8px 20 px 8px;
}

.v-responsive {
    flex: none;
}

.text {
    font-size: 9px;
    color: grey;
    line-height: 8px;
}

::v-deep .v-sheet.v-card:not(.v-sheet--outlined) {
    box-shadow: none !important;
}

::v-deep .v-dialog {
    box-shadow: 5px 5px 12px #d7d7d7,
        -5px -5px 12px #e9e9e9 !important;
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
    padding: 10px !important;
}
</style>