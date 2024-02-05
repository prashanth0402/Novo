<template>
    <v-dialog v-model="dialog" persistent max-width="600px" :fullscreen="this.$vuetify.breakpoint.name == 'xs'"
        :transition="this.$vuetify.breakpoint.name == 'xs' ? 'dialog-bottom-transition' : undefined">

        <v-card class="pa-1" style="max-height: 80vh; overflow-x: hidden;">
            <!-- Common For Both Top Slide -->
            <v-card-title class="d-flex align-center" :class="$vuetify.breakpoint.xs ? 'px-3' : ''">

                <div style="display: flex;justify-content: space-between;width: 100%;align-items: center;">
                    <div class="d-flex align-center">
                        <v-img :src="iconVal" height="30" max-width="30" contain xs2></v-img>

                        <span class="font-weight-normal  ml-4 mr-2 " xs8
                            :class="$vuetify.breakpoint.xs ? 'body-2' : 'button'">{{ detail.name }}</span>
                        <span class="font-weight-normal" v-if="!slide1"
                            :class="$vuetify.breakpoint.xs ? 'ml-3 body-2' : 'button'">
                            -{{
                                Action == 'N' ? 'Order confirmation' : Action == 'M' ? ' Modify Order' : 'CancelOrder'
                            }}
                        </span>

                    </div>


                    <v-btn v-if="this.$vuetify.breakpoint.width > 500" class="red lighten-1 btn body-2  ml-2" dark
                        height="23" elevation="0" v-show="Action == 'M'" @click="ChangeActionFlag('D')" :disabled="disable">
                        <v-icon x-small class="pr-2">mdi-trash-can-outline</v-icon>
                        <span class=" text-capitalize"> cancel Bond </span>
                    </v-btn>
                    <v-btn icon xs2 @click="closePop">
                        <v-icon>mdi-close</v-icon>
                    </v-btn>
                </div>

                <v-row class="d-flex justify-end my-1 mr-2" v-if="this.$vuetify.breakpoint.width < 500">
                    <v-btn class="red lighten-1 btn body-2   ml-2" dark height="23" elevation="0" v-show="Action == 'M'"
                        @click="ChangeActionFlag('D')" :disabled="disable">
                        <v-icon x-small class="pr-2">mdi-trash-can-outline</v-icon>
                        <span class=" text-capitalize"> cancel Bond </span>
                    </v-btn>

                </v-row>

            </v-card-title>

            <v-divider></v-divider>

            <!-- Form -->
            <v-form ref="form" lazy-validation>

                <v-card-text>
                    <v-layout v-if="progress" class="mb-10">
                        <v-flex class="d-flex justify-center">
                            <v-progress-circular indeterminate color="primary" size="50"></v-progress-circular>
                        </v-flex>
                    </v-layout>

                    <!-- Inside the Section -->
                    <section v-else>

                        <!-- Slide1 -->
                        <v-row class="px-5" v-if="Action != 'D' && !slide2">
                            <v-col cols="6" sm="6">
                                <v-row class="mt-3 font-weight-medium subtitle-1 justify-start">
                                    Indicative Yield
                                </v-row>
                                <v-row class="button font-weight-light  justify-start">
                                    {{ detail.symbol }}
                                </v-row>
                            </v-col>


                            <v-col cols="6" sm="6">
                                <v-row class="mt-3 ml-4 font-weight-medium justify-end subtitle-1">
                                    Bid Close
                                </v-row>
                                <v-row class="button font-weight-light justify-end">
                                    {{ detail.dateTime }}
                                </v-row>
                            </v-col>
                        </v-row>

                        <!-- r2 -->
                        <v-row class="px-5" v-if="Action != 'D' && !slide2"
                            :class="this.$vuetify.breakpoint.name != 'xs' ? 'd-flex pt-5 ' : 'd-flex flex-column '">
                            <v-col class="d-flex justify-start" cols="12" xl="4" sm="12" xs="6" md="12" lg="12">
                                <v-row v-if="Action != 'D'" :class="this.$vuetify.breakpoint.width < 600 ? 'mt-1' : ''">
                                    <v-col class="ml-n4" cols="6" sm="6" lg="6">
                                        <v-text-field v-if="Action != 'D'" v-model.number="detail.minLot" label="No of Lot"
                                            outlined :rules="[customValidation]" @keypress="onlyForNumber"
                                            prepend-inner-icon="mdi-minus" append-icon="mdi-plus"
                                            @click:append="incrementValue1" @click:prepend-inner="decrementValue1"
                                            :error-messages="errText" :error="errVal" dense>
                                        </v-text-field>

                                    </v-col>
                                    <v-col cols="6" sm="6" lg="6" class=" ml-4 mt-2 ">
                                        <v-row class="font-weight-medium subtitle-1   d-flex justify-end">
                                            LotSize
                                        </v-row>
                                        <v-row class="button font-weight-light   d-flex justify-end">
                                            {{ detail.modifiedLotSize }}
                                        </v-row>
                                    </v-col>
                                </v-row>

                            </v-col>
                        </v-row>
                        <!-- r2 -->

                        <!-- r3 -->

                        <v-row class="px-5" v-if="Action != 'D' && !slide2"
                            :class="this.$vuetify.breakpoint.name != 'xs' ? 'd-flex ' : 'd-flex flex-column '">
                            <v-col>
                                <v-row :class="this.$vuetify.breakpoint.name != 'xs' ? 'd-flex  ' : 'd-flex flex-row '"
                                    v-if="Action != 'D' && !slide2">
                                    <v-col cols="4">
                                        <v-row class="font-weight-medium subtitle-1 justify-start ">
                                            Units
                                        </v-row>
                                        <v-row class=" button font-weight-light justify-start">
                                            {{ ncb.unit }}
                                        </v-row>
                                    </v-col>
                                    <v-col cols="4">
                                        <v-row class="font-weight-medium subtitle-1 justify-center">
                                            Price
                                        </v-row>
                                        <v-row class=" button font-weight-light justify-center">
                                            ₹ {{ detail.cutoffPrice }} / unit
                                        </v-row>
                                    </v-col>

                                    <v-col cols="4" sm="4">
                                        <v-row class="font-weight-medium subtitle-1  justify-end">
                                            Amount
                                        </v-row>
                                        <v-row class="button font-weight-light  justify-end">
                                            <!-- ₹  <span class=" ml-1 green--text  font-weight-bold "> {{ amount }}</span>  -->
                                            ₹ <span class="ml-1 button font-weight-light">{{ amount }}</span>
                                        </v-row>
                                    </v-col>
                                </v-row>
                            </v-col>
                        </v-row>


                        <!-- N -->
                        <v-row v-if="!slide1 || Action == 'M'">
                            <v-col cols="12" v-if="Action != 'M' && !slide1 && Action != 'D'">
                                <v-row class="mt-5 d-flex justify-center font-weight-medium subtitle-1">
                                    Investment Units
                                </v-row>
                                <v-row v-if="Action != 'D' && !slide1 && Action != 'M'"
                                    class="mx-16 px-14 d-flex justify-center button font-weight-light">
                                    {{ ncb.unit }}
                                </v-row>
                            </v-col>


                            <v-col v-if="Action == 'M' && slide2">

                                <v-row
                                    :class="this.$vuetify.breakpoint.name != 'xs' ? ' mt-6 d-flex px-5 justify-center' : 'mt-5 px-5 d-flex flex-row'"
                                    style="padding-right: -9px !important;padding-left: -12px !important;">

                                    <v-col class=" ml-n4 d-flex justify-end
                                     ">
                                        <v-text-field v-model.number="modify.lotSize" label="No of Lot" outlined
                                            :rules="[customValidation]" @keypress="onlyForNumber"
                                            prepend-inner-icon="mdi-minus" append-icon="mdi-plus"
                                            @click:append="incrementValue2" @click:prepend-inner="decrementValue2"
                                            :error-messages="errText" :error="errVal" dense>
                                        </v-text-field>
                                    </v-col>
                                    <v-col cols="6" sm="6" lg="6" class="mt-2">
                                        <v-row class=" font-weight-medium subtitle-1 d-flex justify-end ">
                                            LotSize
                                        </v-row>
                                        <v-row class="  button font-weight-light d-flex justify-end">
                                            {{ detail.modifiedLotSize }}
                                        </v-row>
                                    </v-col>
                                </v-row>


                                <v-row class="px-5">
                                    <v-col class="" cols="4" sm="4" lg="4">
                                        <v-row class=" font-weight-medium subtitle-1 d-flex justify-start ">
                                            Units
                                        </v-row>
                                        <v-row class="  button font-weight-light d-flex justify-start">
                                            {{ ncb.unit }}
                                        </v-row>
                                    </v-col>

                                    <v-col cols="4" sm="4" lg="4">
                                        <v-row class="font-weight-medium subtitle-1  d-flex justify-center">
                                            Price
                                        </v-row>
                                        <v-row class="button font-weight-light  d-flex justify-center">
                                            ₹ {{ detail.cutoffPrice }} / unit
                                        </v-row>
                                    </v-col>

                                    <v-col class="" cols="4" sm="4" lg="4">
                                        <v-row class="font-weight-medium subtitle-1  d-flex justify-end">
                                            Amount
                                        </v-row>
                                        <v-row class="button font-weight-light d-flex justify-end">
                                            <!-- ₹   <span class=" ml-1 green--text  font-weight-bold">  {{ amount }}</span> -->
                                            ₹ <span class=" ml-1 button font-weight-light"> {{ amount }}</span>
                                        </v-row>
                                    </v-col>
                                </v-row>

                            </v-col>

                            <v-row class="mt-4"
                                :class="this.$vuetify.breakpoint.name != 'xs' ? 'd-flex ml-12' : 'd-flex flex-row ml-7'"
                                v-if="Action != 'D' && !slide2">
                                <v-col cols="6">
                                    <v-row class="ml-5 font-weight-medium subtitle-1">
                                        Units
                                    </v-row>
                                    <v-row class="ml-5  button font-weight-light">
                                        {{ ncb.unit }}
                                    </v-row>
                                </v-col>

                                <v-col cols="6" sm="4">
                                    <v-row class="font-weight-medium subtitle-1">
                                        Amount
                                    </v-row>
                                    <v-row class="button font-weight-light">
                                        ₹ {{ amount }}
                                    </v-row>
                                </v-col>
                            </v-row>

                        </v-row>

                        <v-col cols="12" v-if="Action != 'M' && !slide1 && Action == 'D'">
                            <v-row class="mt-5 d-flex justify-center font-weight-medium subtitle-1">
                                Investment Units
                            </v-row>
                            <v-row v-if="Action == 'D' && !slide1 && Action != 'M'"
                                class="mx-16 px-14 d-flex justify-center button font-weight-light">
                                {{ modify.unit }}
                            </v-row>
                        </v-col>

                        <!-- Common For Both Bottom Slide -->

                        <!-- v-if="Action != 'D'" -->
                        <v-row no-gutters class="mt-5 mx-2">
                            <!-- <v-col :cols="this.$vuetify.breakpoint.width > 500 ? 2 : 3" > -->
                            <v-col
                                :cols="slide2 || Action == 'M' || Action == 'D' ? this.$vuetify.breakpoint.width > 500 ? 2 : 3 : 1">

                                <v-sheet class="pa-1 ma-1">
                                    <v-checkbox class="  d-flex justify-center" dense v-model="disclaimerCheckBox"
                                        v-if="slide2 || Action == 'M' || Action == 'D'" :rules="refRule"></v-checkbox>
                                    <v-icon color="primary" small v-else>mdi-information-variant-circle-outline</v-icon>
                                </v-sheet>
                            </v-col>
                            <v-col class="caption  d-block text-wrap" v-if="Action != 'D'">
                                <v-sheet class="pa-1 ma-1">
                                    Your order will be placed after debiting your trading account on the bid closing
                                    date of each
                                    security.
                                    Ensure to maintain sufficient balance. Bonds will be credited to your demat
                                    account
                                    directly by RBI
                                </v-sheet>
                            </v-col>
                            <v-col class="mt-1 d-block text-wrap" v-if="Action == 'D'">

                                <v-sheet class="pa-1 ma-1 mt-2">
                                    Refund will be issued to your Trading account, after you confirm the order
                                    cancellation.
                                </v-sheet>
                            </v-col>
                        </v-row>

                        <v-divider class="my-3"></v-divider>
                        <v-card-actions class="mt-4">
                            <v-row class="d-flex justify-end mr-2" width="100">
                                <v-btn icon class="secondary white--text mr-5" @click="back"
                                    v-if="slide2 != false && Action != 'M'">
                                    <v-icon>mdi-arrow-left</v-icon>
                                </v-btn>
                                <v-hover v-slot="{ hover }">
                                    <v-btn class="text-capitalize primary darken-1 elevation-0" @click="placeNcb" rounded
                                        :loading="btnLoad" :disabled="confirmdisabled"
                                        :class="hover ? 'secondary' : 'primary white--text'">Confirm</v-btn>
                                </v-hover>
                                <!-- :disabled="isFilled || (slide2 && !disclaimerCheckBox) || validatefield" -->
                            </v-row>
                        </v-card-actions>
                    </section>
                </v-card-text>
            </v-form>
        </v-card>
    </v-dialog>
</template>




<script>
import Eventservice from "@/services/EventServices.js"
export default {
    name: "PlaceOrderDialog",
    data() {
        return {
            erromsgcheckbox: "",
            validatefield: false,
            slide1: false,
            slide2: false,
            disable: false,
            progress: false,
            amount: 0,
            counter: 0,
            errCheckBoxVal: "",
            errCheckBox: false,
            Item: ["Online", "Offline"],
            refRule: [(v) => !!v || "required"],
            customValidation: (v) => !!v || 'Unit must be greater than 0',
            errText: "",
            mtempunit: 0,
            errVal: false,
            ncb: {
                masterId: 0,
                symbol: "",
                series: "",
                name: "",
                closeDate: "",
                actionCode: "",
                unit: 0,
                applicationNo: "",
                lotSize: 0.0,
                orderNo: 0,
                minlot: 0,
                amount: 0.0,
                price: 0.0
            },
            temp: {},
            btnLoad: false,
            disclaimerCheckBox: false,
            CopyAction: "", // To store the Previous ActionFlag
        }
    },
    computed: {
        isFilled() {
            return this.detail.modifiedLotSize !== 0 && this.disclaimerCheckBox != this.disclaimerCheckBox;
        },
        confirmdisabled() {
            if (this.Action == "M") {
                if ((this.modify.lotSize >= this.minLot && this.modify.lotSize <= (this.detail.maxQuantity / this.detail.minBidQuantity)) && (JSON.stringify(this.modify) !== JSON.stringify(this.copymodify))) {
                    return false
                } else {
                    return true
                }


            } else if (this.Action == "N") {
                if (this.detail.minLot >= this.minLot && this.detail.minLot <= (this.detail.maxQuantity / this.detail.minBidQuantity)) {
                    return false
                } else {
                    return true
                }
            } else {
                return false
            }

        }
    },
    props: {
        detail: {},
        copymodify: {},
        modify: {},
        minLot: Number,
        Action: String,
        dialog: Boolean,
        tableType: String,
        iconVal: String,
    },
    methods: {

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
        incrementValue1() {
            if (this.detail.minLot < (this.detail.maxQuantity / this.detail.minBidQuantity)) {
                this.detail.minLot++
            }
        },
        decrementValue1() {
            if (this.detail.minLot > this.minLot) {
                this.detail.minLot--
            }
        },
        incrementValue2() {
            if (this.modify.lotSize < (this.detail.maxQuantity / this.detail.minBidQuantity)) {
                this.modify.lotSize++
            }
        },

        decrementValue2() {
            if (this.modify.lotSize > this.minLot) {
                this.modify.lotSize--
            }
        },
        closePop() {
            this.slide1 = false;
            this.slide2 = false;
            this.disclaimerCheckBox = false;
            this.ncb.id = 0;
            this.unit = 100;
            this.ncb.lotSize = 1;
            this.ncb.amount = 0;
            this.counter = 0;
            this.$emit('closeNcbPop');
            if (this.Action == "D") {
                // this.$emit("ChangeActionFlag", this.CopyAction)
                this.disclaimerCheckBox = false
                this.slide1 = false;
                this.slide2 = true;
            }
        },

        back() {
            this.$refs.form.resetValidation();
            this.slide1 = true;
            this.slide2 = false;
            this.counter = 0;
            this.disclaimerCheckBox = false
            if (this.Action == "D") {
                this.$emit("ChangeActionFlag", this.CopyAction)
                this.disclaimerCheckBox = false
                this.slide1 = false;
                this.slide2 = true;
            }
        },
        ChangeActionFlag(action) {

            this.$emit("ChangeActionFlag", action)
            this.disclaimerCheckBox = false
            // this.$emit('closeNcbPop');
        },
        placeNcb() {

            if (this.Action !== 'D') {
                if (this.$refs.form.validate()) {
                    if (this.Action == "N") {
                        this.counter++;
                        if (this.counter == 1) {
                            if (this.detail.minLot >= this.minLot && this.detail.minLot <= (this.detail.maxQuantity / this.detail.minBidQuantity)) {
                                // this.counter++;
                                // if (this.detail.modifiedLotSize > 0 && this.errVal != true) 

                                this.slide1 = false;
                                this.slide2 = true
                                this.disclaimerCheckBox = false
                                // this.counter++;

                            }
                            this.temp = this.ncb
                        } else if (this.counter == 2) {
                            if (this.slide2 == true && this.counter == 2) {
                                if (this.disclaimerCheckBox == true) {
                                    this.construct()
                                    this.counter = 0;
                                    // this.disclaimerCheckBox = true
                                }
                            }
                        }
                    } else if (this.Action == "M") {
                        this.counter++;
                        if (this.modify.lotSize >= this.minLot && this.modify.lotSize <= (this.detail.maxQuantity / this.detail.minBidQuantity)) {

                            if (this.counter == 1) {
                                if (this.errVal != true) {
                                    this.slide1 = false;
                                    this.slide2 = true;
                                    this.construct()
                                    this.counter = 0;
                                }
                            } else if (this.counter == 2) {
                                if (this.slide2 == true && this.counter == 2) {
                                    this.construct()
                                    this.counter = 0;
                                }
                            }
                        }
                    }
                }
            } else if (this.Action == 'D') {
                if (this.$refs.form.validate()) {
                    this.ncb.masterId = this.detail.id;
                    this.ncb.actionCode = this.Action;
                    this.ncb.symbol = this.detail.symbol;
                    this.ncb.series = this.detail.series
                    this.ncb.name = this.detail.name;
                    this.ncb.orderNo = this.modify.orderNo;
                    this.ncb.closeDate = this.detail.closeDate
                    this.ncb.lotSize = this.detail.modifiedLotSize;
                    this.ncb.amount = this.amount;
                    this.ncb.unit = this.modify.unit;
                    this.ncb.price = this.detail.cutoffPrice;
                    this.ncb.applicationNo = this.modify.applicationNo;
                    this.CallncbPlace();
                }
            }
        },
        construct() {
            this.ncb = this.temp
            this.ncb.masterId = this.detail.id;
            this.ncb.actionCode = this.Action;
            this.ncb.symbol = this.detail.symbol;
            this.ncb.series = this.detail.series
            this.ncb.name = this.detail.name;
            this.ncb.closeDate = this.detail.closeDate
            this.ncb.orderNo = this.modify.orderNo;
            this.ncb.lotSize = this.detail.modifiedLotSize
            this.ncb.amount = this.amount;
            this.ncb.price = this.detail.cutoffPrice;
            this.ncb.applicationNo = this.modify.applicationNo;
            this.CallncbPlace();
        },
        CallncbPlace() {
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
                        this.slide1 = false;
                        this.slide2 = false;
                        this.$emit('closeNcbPop');
                        this.$emit("RecallNcb")
                    } else if (response.data.status == "E") {
                        if (response.data.orderStatus == "") {
                            this.MessageBar("E", response.data.errMsg)
                        } else {
                            this.MessageBar("E", response.data.errMsg)
                        }
                    }
                })
                .catch((error) => {
                    this.disable = false;
                    this.btnLoad = false;
                    this.MessageBar("E", error)
                });
            this.slide1 = false;
            this.slide2 = false;
            this.disclaimerCheckBox = false;
            this.counter = 0;
            this.$emit('closeNcbPop');
        },
        calculateAmount() {
            if (this.Action == 'D') {
                this.Action = "D"
                this.errText = 'Invalid value for ncb.unit or detail.lotSize'
                this.errVal = true;
            }
            if (this.Action == 'N') {
                this.ncb.unit = this.detail.modifiedLotSize * this.detail.minLot
                this.amount = this.ncb.unit * this.detail.cutoffPrice;
                if (this.detail.minLot >= this.minLot && this.detail.minLot <= (this.detail.maxQuantity / this.detail.minBidQuantity)) {
                    this.errVal = false;
                    this.errText = "";
                    this.validatefield = false
                } else {
                    this.errText = "No.of.Lot must between " + this.minLot + " to " + (this.detail.maxQuantity / this.detail.minBidQuantity);
                    this.errVal = true;
                    this.validatefield = true
                }
            }
            if (this.Action == 'M') {
                this.ncb.unit = Number(this.detail.modifiedLotSize * this.modify.lotSize)
                this.amount = this.ncb.unit * this.detail.cutoffPrice;
                if (this.modify.lotSize >= this.minLot && this.modify.lotSize <= (this.detail.maxQuantity / this.detail.minBidQuantity)) {
                    this.errVal = false;
                    this.errText = "";
                    this.validatefield = false

                } else {
                    this.errText = "No.of.Lot must between " + this.minLot + " to " + (this.detail.maxQuantity / this.detail.minBidQuantity);
                    this.errVal = true;
                    this.validatefield = true

                }
            }
            if (this.Action == 'N') {
                this.slide1 = true
                this.slide2 = false
                if (this.detail.modifiedLotSize >= this.detail.lotValue) {
                    this.color = 'success--text'
                } else {
                    this.color = 'error--text'
                }
            } else if (this.Action == 'M') {
                this.slide1 = false
                this.slide2 = true
            }
        }
    },
    watch: {
        detail: {
            handler: function () {
                this.calculateAmount();
            },
            deep: true,
        },
        modify: {
            handler: function () {
                this.calculateAmount();
            },
            deep: true,
        },
        dialog: function (bool) {
            if (bool == true) {
                this.slide1 = true;
                this.progress = true;
                setTimeout(() => {
                    // if ('id' in this.modify && this.Action == "M") {
                    //     this.modify.price = this.amount;
                    //     this.ncb.oldUnit = this.detail.unit;
                    // } else {
                    //     this.ncb.amount = this.amount;
                    //     this.unit = this.detail.minBidQuantity;
                    // }
                    this.progress = false;
                }, 500);
            } else {
                // this.$emit('EmptyModify');
            }
        },
        Action: {
            immediate: true,
            handler(value, oldVal) {
                if (value == "D") {
                    this.CopyAction = oldVal;
                    this.slide1 = false;
                } else if (value == "M") {
                    this.ncb.amount = this.modify.amount
                    this.CopyAction = oldVal;
                }
            }
        }

    },

};


</script>


<!-- <style scoped>
/* ::v-deep.v-input .v-label {
    padding: 15px 16px;
    height: 20px;
    line-height: 5px;
    letter-spacing: normal;
    border-radius: 50px;
    background-color: #ebebebd3;

}

::v-deep.v-text-field--outlined.v-input--dense .v-label--active {
    transform: translateY(-25px) scale(0.75) !important;
}

::v-deep .label {
    left: -14px;
    right: auto;
    position: absolute;
    padding: 5px;
    top: 20px;
    border: 1px solid #8c8c8c !important;
    border-radius: 40px;
}

.v-messages {
    font-size: 50px !important;
} */
</style> -->
<!-- // incQty() {
    //     if (this.ncb.series == "GS") {
    //         if (this.ncb.unit > this.detail.minBidQuantity || this.ncb.unit <= this.detail.gsecmaxQuantity) {
    //             this.ncb.unit = this.ncb.unit + parseInt(this.detail.multiples);
    //             this.icon1Clicked = true;
    //             this.icon2Clicked = false;
    //         }
    //     } else if (this.ncb.series == "TB") {
    //         if (this.ncb.unit > this.detail.minBidQuantity || this.ncb.unit <= this.detail.tbillmaxQuantity) {
    //             this.ncb.unit = this.ncb.unit + parseInt(this.detail.multiples);
    //             this.icon1Clicked = true;
    //             this.icon2Clicked = false;
    //         }
    //     } else {
    //         if (this.ncb.unit > this.detail.minBidQuantity || this.ncb.unit <= this.detail.sdlmaxQuantity) {
    //             this.ncb.unit = this.ncb.unit + parseInt(this.detail.multiples);
    //             this.icon1Clicked = true;
    //             this.icon2Clicked = false;
    //         }
    //     }

    // }, -->
<!-- 
    // else if (this.detail.series == "TB" && this.ncb.unit > minBidQuantity) {
        //     this.ncb.unit = Math.max(this.ncb.unit - parseInt(this.detail.multiples), minBidQuantity);
        //     this.icon2Clicked = true;
        //     this.icon1Clicked = false;
        // } else if (this.detail.series !== "GS" && this.detail.series !== "TB" && this.ncb.unit > minBidQuantity) {
        //     this.ncb.unit = Math.max(this.ncb.unit - parseInt(this.detail.multiples), minBidQuantity);
        //     this.icon2Clicked = true;
        //     this.icon1Clicked = false;
        // } -->

<!-- 
        // decQty() {
            //     if (this.detail.series == "GS") {
            //         if (this.ncb.unit != 100 && this.ncb.unit != "" && this.ncb.unit <= this.detail.gsecmaxQuantity) {
            //             this.ncb.unit = this.ncb.unit - parseInt(this.detail.multiples);
            //             this.icon2Clicked = true;
            //             this.icon1Clicked = false;
            //         }
            //     } else if (this.detail.series == "TB") {
            //         if (this.ncb.unit != 100 && this.ncb.unit != "" && this.ncb.unit <= this.detail.tbillmaxQuantity) {
            //             this.ncb.unit = this.ncb.unit - parseInt(this.detail.multiples);
            //             this.icon2Clicked = true;
            //             this.icon1Clicked = false;
            //         }
            //     } else {
            //         if (this.ncb.unit != 100 && this.ncb.unit != "" && this.ncb.unit <= this.detail.sdlmaxQuantity) {
            //             this.ncb.unit = this.ncb.unit - parseInt(this.detail.multiples);
            //             this.icon2Clicked = true;
            //             this.icon1Clicked = false;
            //         }
            //     }
            // }, -->
<!-- 

        // incQty() {
            //     if (this.detail.series == "GS") {
            //         if (this.ncb.unit < this.detail.minBidQuantity || this.ncb.unit >= this.detail.gsecmaxQuantity) {
            //             this.ncb.unit = this.detail.minBidQuantity;
            //         } else {
            //             this.ncb.unit = this.ncb.unit + parseInt(this.detail.multiples);
            //             this.icon1Clicked = true;
            //             this.icon2Clicked = false;
            //         }
            //     } else if (this.detail.series == "TB") {
            //         if (this.ncb.unit < this.detail.minBidQuantity || this.ncb.unit >= this.detail.tbillmaxQuantity) {
            //             this.ncb.unit = this.detail.minBidQuantity;
            //         } else {
            //             this.ncb.unit = this.ncb.unit + parseInt(this.detail.multiples);
            //             this.icon1Clicked = true;
            //             this.icon2Clicked = false;
            //         }
            //     } else {
            //         if (this.ncb.unit < this.detail.minBidQuantity || this.ncb.unit >= this.detail.sdlmaxQuantity) {
            //             this.ncb.unit = this.detail.minBidQuantity;
            //         } else {
            //             this.ncb.unit = this.ncb.unit + parseInt(this.detail.multiples);
            //             this.icon1Clicked = true;
            //             this.icon2Clicked = false;
            //         }
            //     }
            // }, -->
<!--     
            // rules: {
                //     required: value => !!value || `Unit must be greater than ${this.detail.minBidQuantity}`,
                //     min: v => v >= this.detail.minBidQuantity || `Unit must be greater than  ${this.detail.minBidQuantity}`,
                //     max: v => v <= this.detail.sdlmaxQuantity || `Unit must be less than and equal to ${this.detail.sdlmaxQuantity}`
                // }, -->