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
                                        <v-img :src="iconVal" height="30" max-width="30" contain xs2></v-img>
                                        <span class="font-weight-normal caption ml-2" xs8>{{
                                            detail.symbol }}</span>
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
                                        cancel Bid
                                    </v-btn>
                                </v-flex>
                            </v-col>
                        </v-row>
                    </v-card-title>

                    <v-divider></v-divider>

                    <v-layout v-if="progress" class="mb-10">
                        <v-flex class="d-flex justify-center">
                            <v-progress-circular indeterminate color="primary" size="50"></v-progress-circular>
                        </v-flex>
                    </v-layout>

                    <section v-else>
                        <div v-if="slide1">
                            <v-card-text>
                                <v-row v-if="slide1 && Action != 'D'">
                                    <v-col class="d-flex flex-column justify-start" cols="12" xl="4" lg="4" md="4" sm="4"
                                        xs="6">
                                        <span class="text--disabled">Indicatie yield</span>
                                        <span>{{ detail.name }}</span>
                                    </v-col>
                                    <v-col class="d-flex flex-column justify-start" cols="12" xl="4" lg="4" md="4" sm="4"
                                        xs="6">
                                        <span class="text--disabled">UnitPrice</span>
                                        <span>{{ detail.unitPrice }}</span>
                                    </v-col>
                                    <v-col class="d-flex flex-column justify-start" cols="12" xl="4" lg="4" md="4" sm="4"
                                        xs="6">
                                        <span class="text--disabled">Maturity</span>
                                        <span :class="detail.maturityDate == '-' ? 'ml-5':undefined">{{ detail.maturityDate }}</span>
                                    </v-col>
                                    <v-col class="d-flex flex-column justify-start" cols="12" xl="4" lg="4" md="4" sm="4"
                                        xs="6" v-if="detail.discountText == ''">
                                        <span class="text--disabled">{{ detail.discountText }}</span>
                                        <span>{{ detail.discountAmt }}</span>
                                    </v-col>
                                </v-row>
                            </v-card-text>
                            <v-divider></v-divider>

                            <v-card-text style="background-color:#BBDEFB ;">
                                <!-- <v-card-text> -->

                                <v-row>
                                    <v-col class="d-flex flex-column justify-start" cols="12" xl="4" lg="4" md="4" sm="5"
                                        xs="12">
                                        <span class="text--disabled">Last bid</span>
                                        <span>{{ detail.endDateWithTime }}</span>
                                    </v-col>
                                    <v-col cols="12" xl="3" lg="3" md="3" sm="2"
                                        v-if="$vuetify.breakpoint.name != 'xs'"></v-col>
                                    <v-col class="d-flex flex-column align-start" cols="12" xl="4" lg="4" md="4" sm="5"
                                        xs="12">
                                        <div v-if="!detail.modifyAllowed"
                                            class="d-flex flex-column justify-start align-end">
                                            <span class="text--disabled">Applied unit</span>
                                            <span>{{ detail.appliedUnit }}</span>
                                        </div>
                                        <div v-else>
                                            <span>Units to buy</span>
                                            <v-text-field v-model.number="ncb.unit" class="v-text1" background-color="#fff"
                                                style="height: 22px; position: relative" dense outlined
                                                @keypress="onlyForNumber" :rules="[customValidation]"
                                                :min="parseInt(detail.minBidQuantity)"
                                                :max="parseInt(detail.maxBidQuantity)" @keydown="handleQtyInput"
                                                :error-messages="errText" :error="errVal" @input="handleUnitInput"
                                                :disabled="detail.modifyAllowed == '' && detail.modifyAllowed != undefined ? true : !detail.modifyAllowed">
                                                <template v-slot:prepend-inner>
                                                    <span class="caption text--grey d-flex flex-column pt-2 icon"
                                                        style="position: absolute; right: 15px; top: 0; z-index: 5"><v-icon
                                                            @click="incQty" :class="{
                                                                iconcol1: icon1Clicked,
                                                                'other-class': !icon1Clicked,
                                                            }" size="13" @mouseleave="changeColor"
                                                            color="black">mdi-chevron-up</v-icon>

                                                        <v-icon size="13" @click="decQty" :class="{
                                                            iconcol2: icon2Clicked,
                                                            'other-class': !icon2Clicked,
                                                        }" @mouseleave="changeColor"
                                                            color="black">mdi-chevron-down</v-icon>
                                                    </span>
                                                </template>
                                            </v-text-field>
                                        </div>
                                    </v-col>
                                    <v-col class="d-flex flex-column justify-start" cols="12" xl="4" lg="4" md="4" sm="4"
                                        xs="12">
                                        <span class="text--disabled">Settlement</span>
                                        <span  :class="detail.settlementDate == '-' ? 'ml-8':undefined">{{ detail.settlementDate }}</span>
                                    </v-col>
                                    <v-col class="d-flex flex-column justify-start" cols="12" xl="4" lg="4" md="4" sm="4"
                                        xs="12" v-if="detail.isin == ''">
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
                                        <span class="font-weight-bold" v-if="Action != 'D'">{{ ncb.unit }}
                                            units</span>
                                        <span class="font-weight-bold" v-if="Action == 'D'">{{ detail.appliedUnit }}
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
                                    <v-flex v-else class="text-center">
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

                                            <span v-if="Action != 'R'">₹ {{ (ncb.unit * (ncb.price +
                                                parseInt(detail.discountAmt))).toLocaleString('en-IN')
                                            }}</span>
                                            <span v-else>{{ (parseInt(detail.appliedUnit) * (parseInt(detail.unitPrice) +
                                                parseInt(detail.discountAmt))).toLocaleString('en-IN') }}</span>

                                            <!-- <span>₹ {{ (ncb.unit * (ncb.price + detail.discountAmt)).toLocaleString('en-IN')
                                            }}</span> -->
                                        </v-col>
                                        <v-col class="d-flex flex-column align-end" cols="12" xl="4" lg="4" md="4" sm="4"
                                            xs="6" v-if="detail.discountText == ''">
                                            <span class="text--disabled">Discount.</span>
                                            <span v-if="Action != 'R'">₹ {{ ((ncb.unit * (ncb.price +
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
                                            <!-- <span class="success--text font-weight-bold">₹ {{ (Action
                                                !== 'R' ? amount :
                                                ncb.amount).toLocaleString('en-IN') }}
                                            </span> -->

                                            <span class="success--text font-weight-bold">₹ {{
                                                (amount).toLocaleString('en-IN') }}
                                            </span>
                                        </v-col>
                                    </v-row>
                                </v-col>

                                <v-col class="d-flex align-center justify-end" cols=12 xl="5" lg="5" md="6" sm="12" xs="12">

                                    <v-btn icon small class="secondary white--text mr-2" @click="back"
                                        v-if="slide2 != false">
                                        <v-icon>mdi-arrow-left</v-icon>
                                    </v-btn>
                                    <v-btn class="text-capitalize white--text elevation-0" @click="placeNcb"
                                        :loading="btnLoad" :disabled="isFilled && Action != 'D'" color="#4184f3"
                                        small v-if="Action != 'R'">Confirm</v-btn>
                                    <v-btn class="text-capitalize elevation-0 ml-2" outlined text small
                                        @click="closePop">Close</v-btn>
                                </v-col>
                            </v-row>
                        </v-card-text>
                    </section>
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
            amount: 0,
            ncb: {
                masterId: 0,
                unit: 0,
                oldUnit: 0,
                price: 0,
                amount: 0,
                actionType: "",
                orderNo: 0,
                series: "",
                SIvalue: false,
                SItext: ""
            },
            disable: false,
            inputValue: "",
            // customValidation: (v) => !!v || `Unit must be greater than ${this.detail.minBidQuantity}`,
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
            icon1Clicked: false,
            icon2Clicked: false,
            maxQty: 0
        }
    },
    props: {
        detail: {},
        iconVal: String,
        Action: String,
        dialog: Boolean
    },
    computed: {
        isFilled() {
            if (this.errVal) {
                return false;
            }

            return this.ncb.unit == parseInt(this.detail.appliedUnit);
        },

    },
    methods: {

        customValidation(value) {
            const enteredValue = parseInt(value);

            if (isNaN(enteredValue)) {
                return `Unit must be a number`;
            }

            if (enteredValue < parseInt(this.detail.minBidQuantity)) {
                return `Unit must be greater than or equal to ${this.detail.minBidQuantity}`;
            }

            if (enteredValue > parseInt(this.detail.maxBidQuantity)) {
                return `Unit must be less than or equal to ${this.detail.maxBidQuantity}`;
            }

            // Check if the entered value is a multiple of 100
            if (enteredValue % 100 !== 0) {
                return `Unit must be a multiple of 100`;
            }

            return true; // Validation passed
        },

        handleUnitInput(newValue) {
            const enteredValue = parseInt(newValue);

            if (isNaN(enteredValue)) {
                this.ncb.unit = this.detail.minBidQuantity;
            } else {
                this.ncb.unit = Math.min(
                    Math.max(enteredValue, this.detail.minBidQuantity),
                    this.detail.maxBidQuantity
                );
            }
        },
        changeColor() {
            this.icon1Clicked = false;
            this.icon2Clicked = false;
        },
        handleQtyInput(event) {
            if (event.key === "ArrowUp" || event.keyCode === 38) {
                if (this.ncb.unit < this.detail.minBidQuantity) {
                    this.ncb.unit = this.detail.minBidQuantity;
                } else if (this.ncb.unit < this.detail.maxBidQuantity) {

                    this.ncb.unit = parseInt(this.ncb.unit) + parseInt(this.detail.multiples);
                }
            } else if (event.key === "ArrowDown" || event.keyCode === 40) {
                if (
                    this.ncb.unit != 100 &&
                    this.ncb.unit != "" &&
                    this.ncb.unit >= this.detail.minBidQuantity
                ) {
                    this.ncb.unit = parseInt(this.ncb.unit) - parseInt(this.detail.multiples);
                }
            }
            // console.log('Calculated amount:', this.amount);
        },



        incQty() {
            if (this.ncb.unit < this.detail.minBidQuantity) {
                this.ncb.unit = this.detail.minBidQuantity;
            } else if (this.ncb.unit <= this.detail.maxBidQuantity) {
                this.ncb.unit = parseInt(this.ncb.unit) + parseInt(this.detail.multiples);
                this.icon1Clicked = true;
                this.icon2Clicked = false;
            }
        },

        decQty() {
            const minBidQuantity = this.detail.minBidQuantity;

            if (this.ncb.unit >= minBidQuantity) {
                this.ncb.unit = Math.max(this.ncb.unit - parseInt(this.detail.multiples), minBidQuantity);
                this.icon2Clicked = true;
                this.icon1Clicked = false;
            }

        },

        onlyForNumber($event) {
            let keyCode = $event.keyCode ? $event.keyCode : $event.which;

            if (
                (keyCode < 48 || keyCode > 57) &&
                ![8, 9, 37, 39, 46].includes(keyCode)
            ) {
                $event.preventDefault();
            }

            // Combine the current value and the entered key to check the range
            const enteredValue = parseInt(this.ncb.unit + $event.key);

            // Check if the entered value is within the desired range
            if (
                enteredValue < this.detail.minBidQuantity ||
                enteredValue > this.detail.maxBidQuantity
            ) {
                $event.preventDefault();
            }
        },
        formatedPrice(item) {
            if (item != undefined) {
                return item.toLocaleString('en-IN')
            }
        },
        placeNcb() {
            if (this.Action != "D") {
                if (this.$refs.form.validate()) {
                    if (this.Action == "N") {
                        this.counter++;
                        if (this.counter == 1) {
                            if (this.ncb.unit >= parseInt(this.detail.minBidQuantity) && this.errVal != true) {
                                this.slide1 = false;
                                this.slide2 = true;
                            }
                            this.temp = this.ncb
                        } else if (this.counter == 2) {
                            if (this.slide2 == true && this.counter == 2) {
                                this.construct()
                                this.counter = 0;
                            }
                        }
                    } else if (this.Action == "M") {
                        this.counter++;
                        if (this.counter == 1) {
                            if (this.errVal != true) {
                                this.slide1 = false;
                                this.slide2 = true;
                            }
                            this.temp = this.ncb
                        } else if (this.counter == 2) {
                            if (this.slide2 == true && this.counter == 2) {
                                this.construct()
                                this.counter = 0;
                            }
                        }

                    }

                } 
                // else {
                    // console.log('Form is not valid. Check for validation errors.');
                // }
            } else {
                this.ncb.masterId = parseInt(this.detail.id);
                this.ncb.actionType = this.Action;
                this.ncb.orderNo = this.detail.orderNo;
                this.ncb.price = parseInt(this.detail.unitPrice);
                this.ncb.unit = parseInt(this.detail.appliedUnit);
                this.ncb.oldUnit = parseInt(this.detail.appliedUnit);
                this.ncb.series = this.detail.series;
                this.ncb.SIvalue = this.detail.SIvalue == '' ? false : this.detail.SIvalue;
                this.ncb.SItext = this.detail.SItext
                this.CallNcbPlace();
            }
        },
        construct() {
            this.ncb = this.temp // assign the copy struct
            this.ncb.masterId = parseInt(this.detail.id);
            this.ncb.actionType = this.Action;
            this.ncb.orderNo = this.detail.orderNo;
            this.ncb.price = parseInt(this.detail.unitPrice);
            this.ncb.amount = parseInt(this.detail.unitPrice) * parseInt(this.ncb.unit);
            this.ncb.oldUnit = this.detail.appliedUnit;
            this.ncb.series = this.detail.series;
            this.ncb.SItext = this.detail.SItext;
            if (this.Action == 'N') {
                this.ncb.SIvalue = this.disclaimerCheckBox;
            } else {
                this.ncb.SIvalue = this.detail.SIvalue == '' ? false : this.detail.SIvalue;
            }
            this.CallNcbPlace();
        },
        CallNcbPlace() {
            this.disable = true;
            this.$globalData.overlay = true;
            this.btnLoad = true;
            Eventservice.NcbPlaceOrder(this.ncb)
                .then((response) => {
                    this.$globalData.overlay = false;
                    this.btnLoad = false;
                    this.disable = false;
                    if (response.data.status == "S") {
                        this.MessageBar("S", response.data.orderStatus)
                        this.disclaimerCheckBox = false;
                        this.ncb.unit = 100;
                        this.slide1 = false;
                        this.slide2 = false;
                        this.$emit('closeNcbPop');
                        this.$emit("RecallNcb")
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
            this.ncb.unit = 100;
            this.slide1 = false;
            this.slide2 = false;
            this.$emit('closeNcbPop');
            this.$refs.form.resetValidation();
        },
        closePop() {
            this.slide1 = false;
            this.slide2 = false;
            this.ncb.id = 0;
            this.ncb.unit = 100;
            this.inputValue = "";
            this.ncb.price = 0;
            this.counter = 0;
            this.$emit('closeNcbPop');
            this.$refs.form.resetValidation();

        },
        back() {
            this.slide1 = true;
            this.slide2 = false;
            this.disclaimerCheckBox = false;
            this.counter = 0;
            if (this.Action == "D") {
                this.$emit("ChangeActionFlag", this.CopyAction)
            }
        },
        ChangeActionFlag(action) {
            this.slide1 = !this.slide1;
            this.slide2 = true;
            this.$emit("ChangeActionFlag", action)
        },
    },
    watch: {
        ncb: {
            handler(val) {
                if (this.Action != "D" && this.Action != "R") {
                    this.amount = 0;
                    this.amount = val.price * val.unit;


                    if (val.unit <= parseInt(this.detail.maxBidQuantity) && val.unit >= parseInt(this.detail.minBidQuantity)) {
                        this.errVal = false;
                        this.errText = "";

                    } else if (val.unit < parseInt(this.detail.maxBidQuantity)) {
                        this.errText = "Min. bidQty " + this.detail.minBidQuantity;
                        this.errVal = true;

                    } else if (val.unit > parseInt(this.detail.minBidQuantity)) {
                        this.errText = "Max. bidQty " + parseInt(this.detail.maxBidQuantity);
                        this.errVal = true;
                    }
                } else {
                    this.amount = val.price * val.unit;
                }
            },
            deep: true,
        },

        dialog: function (bool) {
            if (bool == true) {
                this.slide1 = true;
                this.progress = true;
                this.disclaimerCheckBox = this.detail.SIvalue;
                // this.FetchClientFund();
                setTimeout(() => {
                    if ('id' in this.detail && this.Action == "M" || this.Action == "R") {
                        // if (this.Action == "M") {

                        this.ncb.unit = parseInt(this.detail.appliedUnit);
                        this.ncb.price = parseInt(this.detail.unitPrice) - parseInt(this.detail.discountAmt);


                    } else {

                        this.ncb.unit = parseInt(this.detail.minBidQuantity);
                        this.ncb.price = parseInt(this.detail.unitPrice) - parseInt(this.detail.discountAmt);
                    }
                    this.progress = false;
                }, 500); // Change the time delay to your desired value in milliseconds
            }
        },
        Action: {
            immediate: true,
            handler(value, oldVal) {
                if (value == "D") {
                    this.CopyAction = oldVal;
                    this.slide1 = false;
                } else if (value == "M") {
                    this.ncb.price = parseInt(this.detail.unitPrice);
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

::v-deep .v-text1 input {
    max-width: 73%;
}

.v-text1:hover .icon {
    display: flex !important;
}

.iconcol1 {
    background-color: rgb(121, 121, 121);
}

.iconcol2 {
    background-color: rgb(121, 121, 121);
}

.other-class {
    background-color: rgb(236, 236, 236);
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