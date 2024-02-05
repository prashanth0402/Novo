<template>
    <div>
        <v-dialog v-model="dialog" max-width="850" persistent :fullscreen="this.$vuetify.breakpoint.name == 'xs'"
            :transition="this.$vuetify.breakpoint.name == 'xs' ? 'dialog-bottom-transition' : undefined">
            <v-card class="pa-1">
                <v-card-title>
                    <v-row>
                        <v-col cols="12">
                            <v-layout class="d-flex justify-space-between">
                                <v-flex class="d-flex align-center">
                                    <v-img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/sgblogo.webp"
                                        height="35" width="35" contain class="mr-2" xs2></v-img>
                                    <span class="font-weight-normal caption" xs8>{{
                                        detail.name }}</span>
                                    &nbsp;
                                    <span v-if="Action == 'N' && !slide1" class="font-weight-normal caption">- Order
                                        confirmation</span>
                                    <span v-if="Action == 'M' && !slide1" class="font-weight-normal caption">- Modify
                                        Order</span>
                                    <span v-if="Action == 'D' && !slide1" class="font-weight-normal caption">- Cancel
                                        Order</span>
                                </v-flex>
                                <v-flex class="d-flex justify-end">
                                    <v-btn icon xs2>
                                        <v-icon @click="closePop">mdi-close</v-icon>
                                    </v-btn>
                                </v-flex>
                            </v-layout>
                        </v-col>
                        <v-col cols="12" v-if="!slide2 && !progress">
                            <v-row :class="$vuetify.breakpoint.name == 'xs' ? 'd-flex flex-column' : undefined">
                                <v-col class="text-left" cols="12" xl="9" lg="9" md="9" sm="12" xs="12">
                                    <v-alert border="left" color="blue lighten-5" dense v-show="detail.infoText != ''">
                                        <span
                                            :class="this.$vuetify.breakpoint.name == 'xs' ? 'text black--text' : 'caption'">
                                            <v-icon small left>mdi-information-outline</v-icon>
                                            {{ detail.infoText }}
                                        </span>
                                    </v-alert>
                                </v-col>
                                <v-col class="text-right" cols="12" xl="3" lg="3" md="3" sm="12" xs="12">
                                    <v-btn class="red lighten-1 btn body-2 my-3 ml-2" dark height="30" elevation="0"
                                        v-show="detail.cancelAllowed" @click="ChangeActionFlag('D')" :disabled="disable">
                                        <v-icon x-small left>mdi-trash-can-outline</v-icon>
                                        <span class="caption text-capitalize"> cancel</span>
                                    </v-btn>
                                </v-col>
                            </v-row>
                        </v-col>
                    </v-row>
                </v-card-title>
                <v-form ref="form" lazy-validation>
                    <v-card-text>
                        <!-- This progress circle used for UI presentaion and help to avoid showing value jumping -->
                        <v-layout v-if="progress" class="mb-10">
                            <v-flex class="d-flex justify-center">
                                <v-progress-circular indeterminate color="primary" size="50"></v-progress-circular>
                            </v-flex>
                        </v-layout>
                        <!-- This section contains all the details of the SGB -->
                        <section v-else>

                            <v-row v-if="slide1 && Action != 'D'">
                                <v-col class="d-flex flex-column justify-start" v-if="Action != 'R'">
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
                                <v-col class="d-flex flex-column ">
                                    <span class="text--disabled">ISIN</span>
                                    <span>{{ detail.isin }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column  text-center" v-if="detail.orderNo != ''">
                                    <span class="text--disabled">Order No</span>
                                    <span>{{ detail.orderNo }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column  text-end">
                                    <span class="text--disabled">Bidding Period</span>
                                    <span>{{ detail.dateRangeWithTime }}</span>
                                </v-col>
                            </v-row>

                            <!-- Report -->

                            <v-row v-if="!detail.modifyAllowed">
                                <v-col class="d-flex flex-column  align-start">
                                    <span class="text--disabled">UnitPrice</span>
                                    <span>{{ formatedPrice(detail.unitPrice) }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column text-center">
                                    <span class="text--disabled">{{ detail.discountText }}</span>
                                    <span>{{ detail.discountAmt }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column text-end">
                                    <span class="text--disabled">Amount Payable</span>
                                    <span>₹ {{ amount.toLocaleString('en-IN') }}</span>
                                </v-col>
                            </v-row>

                            <v-row class="d-flex " v-if="!detail.modifyAllowed">
                                <v-col class="d-flex flex-column justify-start">
                                    <span class="text--disabled">AppliedUnit</span>
                                    <span>{{ detail.appliedUnit }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column  align-start">
                                    <span class="text--disabled">Requested Unit</span>
                                    <span>{{ detail.requestedUnit }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column  text-center">
                                    <span class="text--disabled">Allocated Unit</span>
                                    <span>{{ detail.allotedUnit }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column  text-end">
                                    <span class="text--disabled">Requested Amount</span>
                                    <span>₹ {{ formatedPrice(detail.requestedAmount) }}</span>
                                </v-col>
                            </v-row>
                            <v-row
                                :class="this.$vuetify.breakpoint.name != 'xs' ? 'd-flex mt-5' : 'd-flex flex-column mt-5'">
                                <v-col cols="12" xl="8" sm="10" xs="6" md="6" v-if="slide1 && Action != 'D'"
                                    class="align-center">

                                    <v-text-field v-if="detail.modifyAllowed" v-model.number="sgb.unit" label="Unit to buy"
                                        outlined @keypress="onlyForNumber" :rules="[customValidation]"
                                        prepend-inner-icon="mdi-minus" append-icon="mdi-plus" @click:append="incrementValue"
                                        @click:prepend-inner="decrementValue" :error-messages="errText" :error="errVal"
                                        :disabled="!detail.modifyAllowed">
                                    </v-text-field>
                                    <!-- <v-col v-if="Action == 'R'"> -->
                                    <v-col v-if="detail.modifyAllowed">
                                        <v-row class="d-flex ">
                                            <!-- <v-col class="d-flex flex-column justify-start" v-if="!detail.modifyAllowed">
                                                <span class="text--disabled">AppliedUnit</span>
                                                <span class="text-center">{{ detail.appliedUnit }}</span>
                                            </v-col> -->
                                            <v-col class="d-flex flex-column  align-start">
                                                <span class="text--disabled">UnitPrice</span>
                                                <span>{{ formatedPrice(detail.unitPrice) }}</span>
                                            </v-col>
                                            <v-col class="d-flex flex-column text-start">
                                                <span class="text--disabled">{{ detail.discountText }}</span>
                                                <span class="text-left">{{ detail.discountAmt }}</span>
                                            </v-col>
                                            <v-col class="d-flex align-start flex-column text-end">
                                                <span class="text--disabled">Amount Payable</span>
                                                <span>₹ {{ amount.toLocaleString('en-IN') }}</span>
                                            </v-col>
                                        </v-row>
                                    </v-col>
                                    <!-- </v-col> -->
                                </v-col>
                                <!-- <v-col cols="12" xl="4" sm="6" xs="4" md="4"
                                    v-if="slide1 && Action != 'D' && Action != 'R' && detail.modifyAllowed"
                                    :class="this.$vuetify.breakpoint.name < 'sm' ? 'd-flex flex-column text-end' :
                                        this.$vuetify.breakpoint.name == 'xl' ? 'd-flex flex-column text-end' : 'd-flex flex-column text-start'" class="justify-center">
                                    <span class="text--disabled">Amount</span>
                                    <span>₹ {{ amount.toLocaleString('en-IN') }}</span>
                                </v-col> -->
                                <v-col cols="12" v-if="slide2 || Action == 'D'">
                                    <v-layout class="d-flex flex-column mb-2">
                                        <v-flex class="d-flex align-center flex-column mb-2">
                                            <span class="text--disabled">Investment
                                                quantity</span>
                                            <span class="font-weight-bold">{{ sgb.unit }}
                                                units</span>
                                        </v-flex>
                                        <v-flex class="d-flex align-center" v-if="Action != 'D'">
                                            <v-checkbox v-model="disclaimerCheckBox" :rules="refRule" dense xl2 lg2 sm2 md2
                                                xs5></v-checkbox>
                                            <span class=" mt-1 d-block text-wrap pb-4" xl10 lg10 sm10 md10 xs7>
                                                {{ detail.SItext }}
                                            </span>
                                        </v-flex>
                                        <v-flex v-else class="text-center">
                                            <span>{{ detail.SIrefundText }}</span>
                                        </v-flex>
                                    </v-layout>
                                    <v-divider></v-divider>
                                </v-col>
                                <v-col class="d-flex justify-end align-center pb-4" v-if="Action != 'R'">
                                    <v-btn icon class="secondary white--text mr-5" @click="back" v-if="slide2 != false">
                                        <v-icon>mdi-arrow-left</v-icon>
                                    </v-btn>
                                    <v-btn class="text-capitalize primary darken-1 elevation-0" @click="placeSgb" rounded
                                        :loading="btnLoad" :disabled="isFilled && Action != 'D'">Confirm</v-btn>
                                </v-col>
                            </v-row>
                        </section>
                    </v-card-text>
                </v-form>
                <!-- {{ detail }} -->
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
            return this.sgb.unit == this.detail.appliedUnit
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
                    this.counter++;
                    if (this.Action == "N") {
                        // console.log("Counter in Action N", this.counter);
                        // if (this.accountbal >= this.amount) {
                        if (this.counter == 1) {
                            if (this.sgb.unit >= this.detail.minBidQty && this.errVal != true) {
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
                this.sgb.masterId = this.detail.id;
                this.sgb.actionCode = this.Action;
                this.sgb.bidId = this.detail.bidId;
                this.sgb.orderNo = this.detail.orderNo;
                this.sgb.price = this.detail.unitPrice - this.detail.discountAmt;
                this.sgb.amount = this.detail.unitPrice * this.detail.appliedUnit;
                this.sgb.unit = this.detail.appliedUnit;
                this.sgb.oldUnit = this.detail.unit;
                this.sgb.SIvalue = this.detail.SIvalue
                this.sgb.SItext = this.detail.SItext
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
                this.sgb.SIvalue = this.detail.SIvalue;
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
                        if (response.data.orderStatus == "") {
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
        incrementValue() {
            this.sgb.unit++;
        },
        decrementValue() {
            if (this.sgb.unit > this.detail.minBidQty) {
                this.sgb.unit--;
            }
        },
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
        choose(item) {
            if (item == "Online") {
                this.sgb.price = this.detail.minPrice;
            } else {
                this.sgb.price = this.detail.maxPrice;
            }
        },
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
                            this.accountbal = response.data.accountBalance
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

                    if (val.unit <= this.detail.maxBidQty && val.unit >= this.detail.minBidQty) {
                        this.errVal = false;
                        this.errText = "";
                    } else {
                        this.errText = "Unit must be within the range of " + this.detail.minBidQty + " - " + this.detail.maxBidQty;
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
                this.disclaimerCheckBox = this.detail.SIvalue;
                this.FetchClientFund();
                setTimeout(() => {
                    if ('id' in this.detail && this.Action == "M" || this.Action == "R") {
                        // if (this.Action == "M") {
                        this.sgb.unit = this.detail.appliedUnit;
                        this.sgb.price = this.detail.unitPrice - this.detail.discountAmt;

                    } else {
                        this.sgb.unit = this.detail.minBidQty;
                        this.sgb.price = this.detail.unitPrice - this.detail.discountAmt;
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
                    this.sgb.price = this.detail.unitPrice
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