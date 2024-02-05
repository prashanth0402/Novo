<template>
  <v-container>

    <v-form ref="form" @input="validation()" lazy-validation>

      <v-row class="d-flex flex-column">
        <v-col cols="12">
          <!-- <v-layout class="d-none d-sm-flex flex-column">
            <v-flex class="d-flex justify-end">
              <v-btn icon :disabled="Disable">
                <v-icon @click="closeBid">mdi-close</v-icon>
              </v-btn>
            </v-flex>
          </v-layout> -->
          <v-alert border="left" color="blue lighten-5" dense class="mt-2">
            <span :class="this.$vuetify.breakpoint.name == 'xs' ? 'text black--text' : 'caption'">
              <v-icon small>mdi-information-outline</v-icon> IPO window will remain
              open from 10 AM till 5 PM,<br> However your may apply after 5PM as offline.
            </span>
          </v-alert>
          <v-layout class="d-none d-sm-flex ">
            <v-flex lg6 sm6>

              <v-text-field v-model="upiId" label="UPI ID" outlined dense class="rounded-0" :rules="atSymbolRule"
                height="5" @input="validateUPI"></v-text-field>
            </v-flex>
            <v-flex lg6 sm6>

              <v-autocomplete v-model="Code" :items="CategoryList" item-text="text" item-value="value" label="Category"
                :rules="refRule" outlined dense class="rounded-0" @change="updateMaxval()" readonly></v-autocomplete>
            </v-flex>
          </v-layout>
          <div class="text-end mb-3 d-none d-sm-flex justify-end">
            <span v-show="!hideAdd" class="mr-2" sm6>
              <v-btn text small class="primary--text" @click="addBid">+ add</v-btn>
            </span>
            <span v-show="!hideCancel" v-if="orderInput.bids.length != 1 && isError != true" sm6>
              <v-btn text small class="error--text" @click="removeslot">cancel</v-btn>
            </span>
          </div>

          <!-- Mobile View -->
          <v-layout class="d-flex d-sm-none flex-column">
            <v-layout>
              <v-flex xs12>
                <v-text-field v-model="upiId" label="UPI ID" outlined dense class="rounded-0" :rules="atSymbolRule"
                  height="10"></v-text-field>
              </v-flex>
            </v-layout>

            <v-layout>
              <v-flex xs12>
                <v-autocomplete v-model="orderInput.category" :items="CategoryList" item-text="text" item-value="value"
                  label="Category" :rules="refRule" outlined dense class="rounded-0" @change="updateMaxval"
                  readonly></v-autocomplete>
              </v-flex>
            </v-layout>
          </v-layout>
          <v-layout v-for="(n, idx) in orderInput.bids" :key="idx" class="d-flex ">
            <!--Desktop bid inputs -->
            <v-flex class="pt-3" lg1 md1 sm1 xs1><b>{{ idx + 1 }}.</b></v-flex>
            <v-flex lg4 md4 sm4 :xs="ApplyDetails.cutOffFlag == 'Y' ? 4 : 5">
              <v-text-field v-model.number="orderInput.bids[idx].quantity" label="No of Lot" background-color="white"
                type="number" min="1" outlined dense @keypress="onlyForNumber" @input="checkPrice(idx)"
                :rules="orderInput.bids[idx].price == 0 ? [] : refRule"
                :error="orderInput.bids[idx].ErrField"></v-text-field>
            </v-flex>
            <v-flex class="ml-2" v-if="ApplyDetails.cutOffFlag == 'Y'">
              <!-- <v-checkbox v-model="orderInput.bids[idx].cutOff" @click="calAmtByLotChange(idx)"
                        @change="disablePrice(orderInput.bids[idx].cutOff, idx)" dense
                        :disabled="ApplyDetails.cutOffFlag == 'N' ? true : false || disableCheckbox(n)"></v-checkbox> -->
              <v-checkbox v-model="orderInput.bids[idx].cutOff" @change="disablePrice(orderInput.bids[idx].cutOff, idx)"
                dense :disable="ApplyDetails.cutOffFlag == 'N' ? true : false" @click="checkPrice(idx)"></v-checkbox>
            </v-flex>
            <v-flex xs4 class="mt-2 d-inline" style="font-size: 12px;" v-if="ApplyDetails.cutOffFlag == 'Y'">
              <b>Cutoff-price</b>
            </v-flex>
            <v-flex :xs="ApplyDetails.cutOffFlag == 'Y' ? 4 : 5">
              <v-text-field v-model.number="orderInput.bids[idx].price" label="Price" background-color="white" outlined
                type="number" :min="ApplyDetails.minPrice" dense :max="ApplyDetails.cutOffPrice"
                :rules="orderInput.bids[idx].quantity == 0 ? [] : refRule" :error="orderInput.bids[idx].ErrField"
                @input="checkPrice(idx)" :error-messages="orderInput.bids[idx].ErrText"
                :disabled="priceArr[idx].priceValid" @keypress="onlyForNumber" class="ml-3"></v-text-field>
            </v-flex>
          </v-layout>
        </v-col>
        <v-col class="d-flex d-sm-none flex-column">
          <v-layout class="text-center mb-2">
            <v-flex v-show="!hideAdd">
              <v-btn text small class="primary--text text-capitalize" @click="addBid">+ add</v-btn>
            </v-flex>
            <v-flex v-show="!hideCancel" v-if="orderInput.bids.length != 1 && isError != true">
              <v-btn text small class="error--text text-capitalize" @click="removeslot">cancel</v-btn>
            </v-flex>
          </v-layout>
        </v-col>
        <v-col>
          <v-row>
            <v-col cols="2" xs="3">
              <v-checkbox v-model="disclaimerCheckBox" :rules="refRule" dense></v-checkbox>
            </v-col>
            <v-col cols="10" xs="9">
              <span class="caption mt-1 d-block text-wrap">
                I hereby undertake that I have read the Red Herring Prospectus and I
                am an eligible UPI bidder as per the application provisions of the
                SEBI (Issue of Capital and Disclosure Requirement) Regulation, 2009.
              </span>
            </v-col>
          </v-row>
          <v-divider class="my-2"></v-divider>
          <v-layout class="mt-3 d-flex align-center" xs12>
            <v-flex class="d-flex justify-start flex-column" style="font-size: 12px;" xs8>
              <b v-show="discountPrice != 0">Amount: &nbsp;<span :class="design">₹{{ amount }}.0</span></b>
              <b v-show="discountPrice != 0">Discount: &nbsp;<span :class="design">₹{{ amount - Math.ceil(amountpayable)
              }}.0</span></b>
              <b>Amount payable: &nbsp;<span :class="design">₹{{ Math.ceil(amountpayable) }}.0</span></b>
            </v-flex>
            <v-flex class="d-flex justify-end">
              <v-flex class="d-flex d-sm-none">
                <!-- <v-btn :disabled="Disable" height="25" @click="closeBid"
                  class="elevation-0 red lighten-1 white--text mr-2 pa-4">
                  close
                </v-btn> -->
              </v-flex>
              <v-btn class="primary body-2" small elevation="0" :disabled="isError" :loading="btnLoading"
                @click="PlaceOrder(orderInput)">
                Submit
              </v-btn>
              <v-btn class="text-capitalize elevation-0 ml-2" outlined text small :disabled="Disable"
                @click="closeBid">Close</v-btn>

            </v-flex>
          </v-layout>
        </v-col>
      </v-row>
    </v-form>
  </v-container>
</template>
<script>
import EventService from "@/services/EventServices";
export default {
  name: "ApplyDetails",

  props: {
    ApplyDetails: {},
    categoryArr: Array,
    selectedCategory: String
  },

  data() {
    return {
      upiId: "",
      isValidUpi: false,
      Disable: false,
      refRule: [(v) => !!v || "required"],
      atSymbolRule: [
        (v) => {
          return /@/.test(v) || "Invalid UPI ID.";
        },
      ],
      MaxValue: 0,
      design: "",
      amount: 0,
      tempAmt1: 0,
      tempAmt2: 0,
      tempAmt3: 0,
      // to hide the add button above the text fiels
      hideAdd: false,
      hideCancel: false,
      // to disable the price text feild
      priceArr: [
        { priceValid: false },
        { priceValid: false },
        { priceValid: false },
      ],
      // To disable the check box
      check: [{ checkbox: false }, { checkbox: false }, { checkbox: false }],
      // To store the bidrefNo
      randomNumber: 0,
      // input struct to placed order
      orderInput: {
        upiId: "",
        upiEndPoint: "",
        category: "",
        symbol: "",
        masterId: 0,
        preApply: '',
        bids: [
          {
            quantity: 0,
            cutOff: false,
            price: 0,
            activityType: "new",
            bidReferenceNo: "",
            lotSize: 0,
            ErrField: false,
            ErrText: "",
          },
        ],
      },
      disclaimerCheckBox: false, // disclaimer chexkBox
      btnLoading: false, // loading for submit button
      // to display text feild as error
      // items: [{ text: "Individual investor", value: "IND" }], // Category DropDown
      upiItems: [], // upiEndPoint dropdown
      isError: false,
      discountPrice: 0.0, // discount
      discountType: "", // discount
      amountpayable: 0,
    };
  },
  methods: {
    // This method prevent the user to select all three check boxes
    // disableCheckbox(item) {
    //   return this.orderInput.bids.filter((dataItem) => dataItem.cutOff).length >= 1 && !item.cutOff;
    // },

    // This method is to disable the price textfeild when checkBox is enabled and vice versa
    disablePrice(checkBox, indicator) {
      if (indicator == 0) {
        if (checkBox == true) {
          this.orderInput.bids[indicator].cutOff = true;
          this.orderInput.bids[indicator].price = parseInt(
            this.ApplyDetails.cutOffPrice
          );
          this.priceArr[indicator].priceValid = true;
        } else if (checkBox == false) {
          this.orderInput.bids[indicator].cutOff = false;
          this.orderInput.bids[indicator].price = this.ApplyDetails.minPrice;
          this.priceArr[indicator].priceValid = false;
        }
        this.orderInput.bids[indicator].ErrText = ""
      } else if (indicator == 1) {
        if (checkBox == true) {
          this.orderInput.bids[indicator].cutOff = true;
          this.orderInput.bids[indicator].price = parseInt(
            this.ApplyDetails.cutOffPrice
          );
          this.priceArr[indicator].priceValid = true;
        } else if (checkBox == false) {
          this.orderInput.bids[indicator].cutOff = false;
          this.orderInput.bids[indicator].price = this.ApplyDetails.minPrice;
          this.priceArr[indicator].priceValid = false;
        }
        this.orderInput.bids[indicator].ErrText = ""
      } else if (indicator == 2) {
        if (checkBox == true) {
          this.orderInput.bids[indicator].cutOff = true;
          this.orderInput.bids[indicator].price = parseInt(
            this.ApplyDetails.cutOffPrice
          );
          this.priceArr[indicator].priceValid = true;
        } else if (checkBox == false) {
          this.orderInput.bids[indicator].cutOff = false;
          this.orderInput.bids[indicator].price = this.ApplyDetails.minPrice;
          this.priceArr[indicator].priceValid = false;
        }
        this.orderInput.bids[indicator].ErrText = ""
      }
    },
    addBid() {
      if (this.isError != true) {
        if (this.orderInput.bids.length == 1) {
          if (
            this.orderInput.bids[0].price != 0 &&
            this.orderInput.bids[0].quantity != 0 &&
            this.orderInput.bids[0].ErrField != true
          ) {
            this.generateRandomNumber();
            this.orderInput.bids.push({
              quantity: 0,
              cutOff: false,
              price: 0,
              activityType: "new",
              bidReferenceNo: this.randomNumber,
              lotSize: 0,
              ErrField: false,
              ErrText: "",
            });
          } else {
            if (this.orderInput.bids[0].price != 0) {
              this.orderInput.bids[0].ErrField = false;
            } else {
              this.orderInput.bids[0].ErrField = true;
              this.orderInput.bids[0].ErrText = this.ApplyDetails.priceRange;
            }
          }
        } else if (this.orderInput.bids.length == 2) {
          if (
            this.orderInput.bids[1].price != 0 &&
            this.orderInput.bids[1].quantity != 0
          ) {
            this.generateRandomNumber();
            this.orderInput.bids.push({
              quantity: 0,
              cutOff: false,
              price: 0,
              activityType: "new",
              bidReferenceNo: this.randomNumber,
              lotSize: 0,
              ErrField: false,
              ErrText: "",
            });
          } else {
            if (this.orderInput.bids[1].price != 0) {
              this.orderInput.bids[1].ErrField = false;
            } else {
              this.orderInput.bids[1].ErrField = true;
              this.orderInput.bids[1].ErrText = this.ApplyDetails.priceRange;
            }
          }
        }
      }
    },
    removeslot() {
      this.orderInput.bids.pop();
      const len = this.orderInput.bids.length;
      if (len == 2) {
        this.priceArr[2].priceValid = false;
      } else if (len == 1) {
        this.priceArr[1].priceValid = false;
      }
    },

    // this method is to check the enter values are number
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
    // this method is to check wheather the value entered in the price feild lies between the range
    checkPrice(indicator) {
      if (indicator == 0) {
        if (
          (parseInt(this.orderInput.bids[0].price) <
            parseInt(this.ApplyDetails.minPrice) ||
            parseInt(this.orderInput.bids[0].price) >
            parseInt(this.ApplyDetails.cutOffPrice)
          )
        ) {
          this.isError = true;
          this.orderInput.bids[0].ErrField = true;
          this.orderInput.bids[0].ErrText = this.ApplyDetails.priceRange;
        } else {
          this.isError = false;
          this.orderInput.bids[0].ErrField = false;
          this.orderInput.bids[0].ErrText = "";
        }
      } else if (indicator == 1) {
        if (
          (parseInt(this.orderInput.bids[1].price) <
            parseInt(this.ApplyDetails.minPrice) ||
            parseInt(this.orderInput.bids[1].price) >
            parseInt(this.ApplyDetails.cutOffPrice)) &&
          this.orderInput.bids[1].price != 0
        ) {
          this.isError = true;
          this.orderInput.bids[1].ErrField = true;
          this.orderInput.bids[1].ErrText = this.ApplyDetails.priceRange;
        } else {
          this.isError = false;
          this.orderInput.bids[1].ErrField = false;
          this.orderInput.bids[1].ErrText = "";
        }
      } else if (indicator == 2) {
        if (
          (parseInt(this.orderInput.bids[2].price) <
            parseInt(this.ApplyDetails.minPrice) ||
            parseInt(this.orderInput.bids[2].price) >
            parseInt(this.ApplyDetails.cutOffPrice)) &&
          this.orderInput.bids[2].price != 0
        ) {
          this.isError = true;
          this.orderInput.bids[2].ErrField = true;
          this.orderInput.bids[2].ErrText = this.ApplyDetails.priceRange;
        } else {
          this.isError = false;
          this.orderInput.bids[2].ErrField = false;
          this.orderInput.bids[2].ErrText = "";
        }
      }
    },

    // UpiEndNames() {
    //   EventService.GetUpi()
    //     .then((response) => {
    //       if (response.data.status == "S") {
    //         this.upiItems = response.data.upiArr;
    //       } else if (response.data.status == "E") {
    //         this.displayMsg.errorbar = true;
    //         this.displayMsg.err = response.data.errMsg;
    //       }
    //     })
    //     .catch((error) => {
    //       this.displayMsg.errorbar = true;
    //       this.displayMsg.err = error.response;
    //     });
    // },
    // To generate temporary bidReference number
    generateRandomNumber() {
      const randomNumber = Math.floor(Math.random() * 10000);
      this.randomNumber = randomNumber.toString().padStart(6, "0");
    },
    // To place a Order
    PlaceOrder(value) {
      this.$emit("BidOverlay")
      if (this.$refs.form.validate()) {
        if (
          this.orderInput.bids[0].quantity != 0 &&
          this.orderInput.bids[0].price != 0
        ) {
          this.Disable = true;
          this.btnLoading = true;
          value.symbol = this.ApplyDetails.symbol;
          value.masterId = this.ApplyDetails.id;
          value.preApply = this.ApplyDetails.preApply
          value.category = this.Code

          this.checkBids(this.orderInput.bids); // to remove the null value structure
          this.removeValue(); // to remove unwanted variable from the structure
          this.generateRandomNumber();
          this.splitString(this.upiId); // To split the upi endpoints
          for (let i = 0; i < this.orderInput.bids.length; i++) {
            this.orderInput.bids[i].lotSize = this.ApplyDetails.lotSize;
            this.orderInput.bids[0].bidReferenceNo = this.randomNumber;
          }

          EventService.PlaceOrder(value)
            .then((response) => {
              this.Disable = false;
              this.$emit("BidOverlay1")
              if (response.data.status == "S") {
                //this method is use to change the default falue of varaiable
                this.emptyInput();

                this.btnLoading = false;
                this.$refs.form.resetValidation();
                this.$emit("Recall");
                if (response.data.status == []) {
                  this.MessageBar(
                    "E",
                    "Unable to process your request try again!"
                  );
                  this.btnLoading = false;
                } else if (response.data.appStatus == "success") {
                  this.$emit("closeDialog");
                  this.orderInput.bids.length = 1;
                  this.btnLoading = false;
                } else if (response.data.appStatus == "Pending") {
                  this.$emit("closeDg");
                  this.btnLoading = false;
                } else {
                  this.btnLoading = false;
                  // this.$emit("closeDialog");
                  this.MessageBar("E", response.data.appReason);
                }
              } else if (response.data.status == "") {
                this.MessageBar(
                  "E",
                  "Unable to process your request try again!"
                );
                this.btnLoading = false;
              } else if (response.data.status == "I") {
                this.btnLoading = false;
                this.$router.replace("/login");
              } else if (response.data.status == "E") {
                this.btnLoading = false;
                this.Disable = false;
                // this.$emit("closeDg");
                this.MessageBar("E", response.data.errMsg);
              }
            })
            .catch((error) => {
              this.Disable = false;
              this.btnLoading = false;
              this.MessageBar("E", error);
            });
        } else {
          this.MessageBar("E", "Field cannot be empty");
        }
      } else {
        this.btnLoading = false;
      }
    },

    // To close a bid Dialog
    closeBid() {
      this.emptyInput();
      this.$refs.form.resetValidation();
      // To replace the length when user opens the bid
      this.orderInput.bids.length = 1;
      for (var pc = 0; pc < this.priceArr.length; pc++) {
        this.priceArr[pc].priceValid = false;
        this.isError = false;
      }
      this.MaxValue = 0
      this.$emit("closeDg");
    },
    //
    emptyInput() {
      this.upiId = "";
      for (var i = 0; i < this.orderInput.bids.length; i++) {
        this.orderInput.bids[i].quantity = 0;
        this.orderInput.bids[i].price = 0;
        this.orderInput.bids[i].cutOff = false;
        this.orderInput.bids[i].ErrField = false;
        this.orderInput.bids[i].ErrText = ''

      }
      this.orderInput.category = "IND"
      this.amount = 0;
      this.tempAmt1 = 0;
      this.tempAmt2 = 0;
      this.tempAmt3 = 0;
      this.disclaimerCheckBox = false;
      this.Disable = false;
    },

    // this method is to disable the submit button when price feild is in error
    validation() {
      if (
        (parseInt(this.orderInput.bids[0].price) <
          parseInt(this.ApplyDetails.minPrice) ||
          parseInt(this.orderInput.bids[0].price) >
          parseInt(this.ApplyDetails.cutOffPrice)) &&
        this.orderInput.bids[0].price != 0
      ) {
        return true;
      } else {
        return false;
      }
    },
    validateUPI() {
      if (this.upiId != null) {
        this.isValidUpi = this.upiId.includes('@');
      } else {
        this.isValidUpi = false
      }
    },
    checkBids(arr) {
      // Use a loop to filter out elements whos value is null
      for (let i = 0; i < arr.length; i++) {
        if (arr[i].price == 0 || arr[i].quantity == 0) {
          arr.splice(i, 1);
          i--;
        }
      }
    },
    // To split upiId and EndPoint seprately
    splitString(text) {
      const splitStrings = text.split("@");
      if (splitStrings.length >= 2) {
        this.orderInput.upiId = splitStrings[0];
        this.orderInput.upiEndPoint = "@" + splitStrings[1];
      }
    },
    // To remove unwanted vaules inside the array
    removeValue() {
      var testArr = [];
      for (let rv = 0; rv < this.orderInput.bids.length; rv++) {
        testArr.push({
          activityType: this.orderInput.bids[rv].activityType,
          bidReferenceNo: this.orderInput.bids[rv].bidReferenceNo,
          cutOff: this.orderInput.bids[rv].cutOff,
          lotSize: this.orderInput.bids[rv].lotSize,
          price: this.orderInput.bids[rv].price,
          quantity: this.orderInput.bids[rv].quantity,
        });
      }
      this.orderInput.bids = testArr;
    },

    updateMaxval() {
      for (let index = 0; index < this.categoryArr.length; index++) {
        if (this.Code === this.categoryArr[index].value) {
          this.MaxValue = this.categoryArr[index].maxvalue
          this.discountPrice = this.categoryArr[index].discountPrice
          this.discountType = this.categoryArr[index].discountType
          this.$emit("discountStruct", this.categoryArr[index])
        }

      }
    },
    formatNumberWithCommas(number) {
      return number.toLocaleString('en-IN');
    },
    CalcAmountpayable(value) {
      let amount
      if (this.discountType == "A") {
        amount =
          this.ApplyDetails.lotSize *
          parseInt(value.quantity) *
          (parseInt(value.price) - this.discountPrice);
      } else {
        amount =
          this.ApplyDetails.lotSize *
          parseInt(value.quantity) *
          (parseInt(value.price) -
            // (this.discountPrice / parseInt(value.price)) * 100
            (parseInt(value.price) * (this.discountPrice / 100))
          );
      }
      return amount
    }
  },

  updated() {
    this.orderInput.category = this.Code
    this.updateMaxval()
    // popup error msg when the amount exceeds 2 lahks
    if (this.amountpayable > this.MaxValue) {
      this.MessageBar("E", "The total amount should not exceed " + this.formatNumberWithCommas(this.MaxValue));
      this.isError = true;
    }

  },
  mounted() {
    // Your code that should run after the component is mounted
    // To get Upi endpoint from Database
    // this.UpiEndNames();
    // To generate bidReferenceNo
    this.generateRandomNumber();
  },
  //TODO Newly added
  computed: {
    hidetab: {
      get() {
        if (this.$vuetify.breakpoint.name < 650) {
          return false
        } else {
          return true
        }
      }
    },
    CategoryList() {
      const FilteredCategoryArr = this.categoryArr.filter(item => item.text == this.selectedCategory)
      return FilteredCategoryArr
    },
    Code() {
      const category = this.categoryArr.filter(item => item.text == this.selectedCategory)
      let code = ""
      for (let i = 0; i < category.length; i++) {
        code = category[i].value
      }
      return code
    }
  },

  watch: {
    // Commented by prashanth
    //  max value is not mapped on  on mounted show an undefiene Reading Error
    //  watcher is used to overcome only For First time Getting max of value
    // categoryArr: {
    //   immediate: true,
    //   handler() {
    //     this.updateMaxval()
    //   }
    // },
    orderInput: {
      handler(newval) {
        this.amount = 0;
        this.amountpayable = 0
        // this.tempAmt1 = 0;
        // this.tempAmt2 = 0;
        // this.tempAmt3 = 0;
        let tempAmt1 = 0, tempAmt2 = 0, tempAmt3 = 0, discount1 = 0, discount2 = 0, discount3 = 0;

        for (var i = 0; i < newval.bids.length; i++) {
          if (newval.bids[i].price != 0 && newval.bids[i].quantity != 0) {
            newval.bids[i].ErrField = false;
          } else {
            newval.bids[i].price = this.ApplyDetails.minPrice
            newval.bids[i].quantity = 1
          }
          if (
            parseInt(newval.bids[i].price) > 0 &&
            parseInt(newval.bids[i].quantity) > 0
          ) {
            if (i == 0) {
              tempAmt1 =
                newval.bids[0].price *
                newval.bids[0].quantity *
                this.ApplyDetails.lotSize;
              discount1 = this.CalcAmountpayable(newval.bids[i])

            } else if (i == 1) {
              tempAmt2 =
                newval.bids[1].price *
                newval.bids[1].quantity *
                this.ApplyDetails.lotSize;
              discount2 = this.CalcAmountpayable(newval.bids[i])
            } else {
              tempAmt3 =
                newval.bids[2].price *
                newval.bids[2].quantity *
                this.ApplyDetails.lotSize;
              discount3 = this.CalcAmountpayable(newval.bids[i])
            }
            this.amount = Math.max(tempAmt1, tempAmt2, tempAmt3);
            this.amountpayable = Math.max(discount1, discount2, discount3)
            // this.amount = this.amount + (parseInt(newval.bids[i].price) * parseInt(newval.bids[i].quantity) * parseInt(this.ApplyDetails.lotSize));
          }
        }
        if (this.amountpayable > this.MaxValue) {
          this.design = "red--text";
        } else if (this.amountpayable < this.MaxValue || this.amountpayable > 0) {
          this.design = "green--text";
          // this.isError = false;
        } else {
          this.design = "grey--text";
        }

        if (newval.bids.length == 3) {
          this.hideAdd = true;
        } else {
          this.hideAdd = false;
          if (newval.bids.length == 1) {
            this.hideCancel = true;
          } else if (newval.bids.length == 2) {
            // if (newval.bids[1].ErrField == true) {
            //   this.hideCancel = true;
            //   this.$refs.form.resetValidation();
            // } else {
            this.hideCancel = false;
            // }
          } else {
            this.hideCancel = false;
          }
        }
        this.orderInput.upiId = newval.upiId.toLowerCase();
      },
      deep: true,
    },
    ApplyDetails: {
      handler() {
        if (this.ApplyDetails != {}) {
          this.emptyInput();
          this.generateRandomNumber();
          this.orderInput.bids[0].bidReferenceNo = this.randomNumber;
          for (let pa = 0; pa < this.priceArr.length; pa++) {
            this.priceArr[pa].priceValid = false;
          }
        }
      },
      deep: true,
    },
  },
};

</script>

<style scoped>
.text {
  font-size: 10px;
  color: grey;
}

.text-wrap {
  font-size: 8px;
}

::v-deep .v-messages__message {
  font-size: 8px;
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
  padding: 5px !important;
}
</style>