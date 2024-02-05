<template>
  <div>
    <v-layout>
      <v-flex class="d-flex justify-end">
        <v-btn class="red lighten-1 btn body-2 mt-2 ml-2" dark height="23" elevation="0" @click="dialog = !dialog"
          :disabled="disable">
          <v-icon x-small class="pr-2">mdi-trash-can-outline</v-icon>
          <span class="caption text-capitalize"> Delete Application </span>
        </v-btn>
      </v-flex>
    </v-layout>

    <v-alert border="left" color="blue lighten-5" dense class="mt-3">
      <span :class="this.$vuetify.breakpoint.name == 'xs' ? 'text black--text' : 'caption'">
        <v-icon small>mdi-information-outline</v-icon> IPO window will remain open
        from 10 AM till 5 PM</span>
    </v-alert>

    <v-card class="pa-5 mb-3" elevation="0" color="blue lighten-5" rounded="0">
      <v-layout :class="this.$vuetify.breakpoint.name != 'xs' ? 'body-2' : 'caption'">
        <v-flex>
          <v-layout>
            <v-flex class="d-flex justify-center font-weight-black">
              UPI ID
            </v-flex>
          </v-layout>
          <v-layout>
            <v-flex class="d-flex justify-center">
              {{ modifyData.upi }}
            </v-flex>
          </v-layout>
        </v-flex>
        <v-flex>
          <v-layout>
            <v-flex class="d-flex justify-center font-weight-black"><b> Category</b></v-flex>
          </v-layout>
          <v-layout>
            <v-flex class="d-flex justify-center">
              {{ modifyData.category }}</v-flex>
          </v-layout>
        </v-flex>
        <v-flex>
          <v-layout>
            <v-flex class="d-flex justify-center"><b> Amount payable </b>
            </v-flex>
          </v-layout>

          <v-layout>
            <v-flex class="d-flex justify-center">
              {{ amountPayable }}
            </v-flex>
          </v-layout>
        </v-flex>
      </v-layout>
    </v-card>

    <v-layout>
      <v-flex class="d-flex justify-end">
        <v-btn class="body-2 mb-2 ml-2" dark text color="primary" elevation="0" @click="AddNewBid()" v-show="showAddBtn">
          <v-icon x-small class="pr-1">mdi-plus</v-icon>
          <span class="caption text-capitalize"> Add </span>
        </v-btn>
      </v-flex>
    </v-layout>
    <div v-for="bid, idx  in modifyData.modifyDetails" :key="idx">
      <!-- {{ modifyData.modifyDetails[idx].signal }} -->
      <ReadOnly :bid="bid" :Disable="disable" :Idx="idx" @modify="ModifyBid" :detail="issueDetails"
        v-if="modifyData.modifyDetails[idx].signal === 'O' || modifyData.modifyDetails[idx].activityType === 'cancel'" />

      <ModifyBid ref="Element" :bid="bid" :Disable="disable" :Idx="idx" :issueDetails="issueDetails"
        :totalAmt="modifyData.total" @closeSlot="closeSlot(idx, modifyData.modifyDetails[idx].signal)"
        @UpdateTotal="updateTotal" @hideUpdate="hideUpdate" v-else-if="modifyData.modifyDetails[idx].signal === 'N'" />
    </div>
    <v-layout class="mt-3 body-2 d-flex align-center mb-2">
      <v-flex class="justify-start">
        <span class="error--text" v-show="msg">The total amount should not exceed {{ formatNumberWithCommas(this.MaxValue)
        }}</span>
      </v-flex>
      <v-flex class="d-flex justify-end">
        <v-layout class="mt-3">
          <v-flex class="d-flex justify-end align-center ">
            <!-- <v-btn :disabled="closeBtn" @click="closeModify"
              class="red lighten-1 white--text mr-3 elevation-0 d-flex d-sm-none pa-4" small>
              close
            </v-btn> -->
            <v-btn class="body-2 white--text elevation-0" small :disabled="modifyBtn || isError" :loading="btnLoading"
              @click="UpdateBidToExchange(modifyData, 'N')" color="#4184f3">
              <span> Update </span>
            </v-btn>
            <v-btn class="text-capitalize elevation-0 ml-2 red lighten-2 white--text" outlined text small @click="closeModify"
              :disabled="closeBtn">Close</v-btn>
          </v-flex>
        </v-layout>
        <!-- <v-btn class="blue darken-3 body-2 pa-4 white--text" height="25" elevation="0" :disabled="modifyBtn || isError"
          :loading="btnLoading" @click="UpdateBidToExchange(modifyData, 'N')">
          <span> Update </span>
        </v-btn> -->
      </v-flex>
    </v-layout>
    <!-- final popup -->
    <v-dialog v-model="dialog" width="600" persistent>
      <v-card>
        <v-card-text class="d-flex justify-space-between align-center pa-10">
          <v-row>
            <v-col cols="12" xl="6" lg="6" md="6" sm="12" xs="12" class="d-flex justify-center align-center text-center">
              <span class="text-subtitle-1 black--text">Do you want to delete application?</span>
            </v-col>
            <v-col cols="12" xl="3" lg="3" md="3" sm="12" xs="12" class="d-flex justify-center align-center">
              <v-btn @click="cancelBid('N')" class="error elevation-0" small width="100">No</v-btn>
            </v-col>
            <v-col cols="12" xl="3" lg="3" md="3" sm="12" xs="12" class="d-flex justify-center align-center">
              <v-btn @click="UpdateBidToExchange(modifyData, 'Y')" class="primary elevation-0" small
                width="100">Yes</v-btn>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-dialog>
  </div>
</template>
<script>
import ReadOnly from "./ReadOnly.vue";
import ModifyBid from "./ModifyBid.vue";
import EventService from "@/services/EventServices";

// import ModifyInput from "../../../../../../Script/Modify/ModifyInput";
// export default {
//   mixins: [ModifyInput],
// };
export default {
  name: "modifyDetail",
  components: {
    ReadOnly,
    ModifyBid,
  },
  props: {
    categoryArr: Array,
    // To display the details of the Ipo
    issueDetails: {},
    // TO get the Placed bid details along with upi details
    modifyData: {},
    //  TO show the add button based on the lenght of the array
    showAddBtn: Boolean,
    // It's used to calculate the amount inside modifyData and used to show the amount payable
    ipoAppTotal: Number,
    modifyBtn: Boolean,
    copiedModifyData: {},
    amountPayable: Number
  },

  data() {
    return {
      MaxValue: 0,
      btnLoading: false, // loading for modify button
      dialog: false,
      cancelBtn: false,
      orderInput: {
        applicationNo: "",
        upiId: "",
        upiEndPoint: "",
        category: "",
        symbol: "",
        masterId: 0,
        bids: [],
      },
      msg: false,
      copyStruct: {},
      disable: false,
      isError: false,
      closeBtn: false,
    };
  },

  methods: {
    formatNumberWithCommas(number) {
      return number.toLocaleString('en-IN');
    },

    updateMaxval() {
      for (let index = 0; index < this.categoryArr.length; index++) {
        if (this.modifyData.category == this.categoryArr[index].value) {
          this.MaxValue = this.categoryArr[index].maxvalue
          // console.log("discountStruct in modifyInput",this.categoryArr[index]);
          this.$emit("discountStruct", this.categoryArr[index])
          break
        }

      }
    },
    closeModify() {
      this.$emit("closeModify");
    },
    AddNewBid() {
      this.$emit("showModifyBtn");
      if (this.modifyData.modifyDetails.length < 3) {
        this.$emit("addNew");
      }
    },
    cancelBid(flag) {
      if (flag == "Y") {
        this.$emit("cancelBid");
        this.dialog = false;
      } else {
        this.dialog = false;
      }
    },
    ModifyBid(index) {
      this.$emit("modified", index);
    },
    closeSlot(index, signal) {
      this.$emit("closeSlot", index, signal);
    },
    updateTotal(total) {
      this.modifyData.total = total;
    },
    UpdateBidToExchange(value, indicator) {
      if (
        (value.modifyDetails[0].quantity != 0 &&
          value.modifyDetails[0].price != 0) ||
        (value.modifyDetails[1].quantity != 0 &&
          value.modifyDetails[1].price != 0) ||
        (value.modifyDetails[2].quantity != 0 &&
          value.modifyDetails[2].price != 0)
      ) {
        if (indicator == "Y") {
          this.copyStruct = value;
        } else {
          this.copyStruct = JSON.parse(JSON.stringify(value)); // To have the permenant copy of the struct
        }
        this.disable = true;
        this.cancelBid(indicator); // To check if the user cancel the bid or not
        this.dialog = false;
        this.btnLoading = true;
        this.closeBtn = true;
        this.$emit("DgOverlay")
        // const localArr = value.modifyDetails // this variable is use to prevent the values change in original
        this.splitString(this.copyStruct.upi); // to split the upi and upi endname

        for (let MD = 0; MD < this.copyStruct.modifyDetails.length; MD++) {
          if (this.copyStruct.masterId == this.issueDetails.id) {
            this.orderInput.applicationNo = this.copyStruct.appNo;
            this.orderInput.symbol = this.issueDetails.symbol;
            this.orderInput.masterId = this.copyStruct.masterId;
            this.orderInput.category = this.modifyData.category

            this.copyStruct.modifyDetails[MD].lotSize =
              this.issueDetails.lotSize;
          }
        }
        // console.log("this.copyStruct",this.copyStruct);
        this.checkBids(this.copyStruct.modifyDetails); // to remove the non changed structure from the array
        this.removeValue(this.copyStruct.modifyDetails); // to remove the unNeccessaray values for Api
        // let count = 0;
        // console.log("AFTER Remove value",this.copyStruct.modifyDetails);

        // if (
        //   this.orderInput.bids.length >
        //   this.copiedModifyData.modifyDetails.length
        // ) {
        //   count++;
        // }
        if (indicator == "N") {
          let TempArr = [];
          // console.log();
          for (var MD = 0; MD < this.orderInput.bids.length; MD++) {
            // console.log(count, "outerloop", this.orderInput.bids);

            for (
              var CS = 0;
              CS < this.copiedModifyData.modifyDetails.length;
              CS++
            ) {
              // if (
              //   this.orderInput.bids.length ==
              //   this.copiedModifyData.modifyDetails.length
              // ) {
              if (
                this.orderInput.bids[MD].id ==
                this.copiedModifyData.modifyDetails[CS].id
              ) {
                if (
                  this.orderInput.bids[MD].price !=
                  this.copiedModifyData.modifyDetails[CS].price ||
                  this.orderInput.bids[MD].quantity !=
                  this.copiedModifyData.modifyDetails[CS].quantity
                ) {
                  TempArr.push(this.orderInput.bids[MD]);
                }
              }
            }
            if (this.orderInput.bids[MD].id == 0) {
              TempArr.push(this.orderInput.bids[MD]);
            }
          }
          if (TempArr.length == 0) {
            this.$emit("DgOverlay1")
            this.MessageBar("E", "No Changes were made");
            this.closeBtn = false;
            this.btnLoading = false;
            this.disable = false;
            // this.orderInput.bids = [];
          } else {
            this.orderInput.bids = TempArr;
            // console.log(" if Call Api", this.orderInput);
            this.Order(this.orderInput);
          }
        } else {
          // console.log(" else Call Api", this.orderInput);
          this.Order(this.orderInput);
        }
      } else {
        this.$emit("DgOverlay1")
        this.MessageBar("E", "Field Cannot be empty!");
        this.btnLoading = false;
      }
    },
    //placingOrder
    Order(value) {
      // console.log("order",value);
      EventService.PlaceOrder(value)
        .then((response) => {
          this.$emit("DgOverlay1")
          this.closeBtn = false;
          this.disable = false;
          // console.log("Order response",response)
          if (response.data.status == "S") {
            this.$emit("closeModify");
            this.btnLoading = false;
            this.$emit("Recall");
            if (response.data.appStatus == "") {
              this.MessageBar(
                "E",
                "Sorry Unable to process your Request right now !"
              );
            } else if (response.data.appStatus == "success") {
              this.$emit("closeDialog");
              // this.$emit("Recall");
              this.MessageBar("S", response.data.appStatus);
            } else if (response.data.appStatus == "Pending") {
              this.$emit("closeDialog");
              // this.$emit("Recall");
              this.MessageBar("S", response.data.appReason);
            }
          } else if (response.data.status == "I") {
            this.$router.replace("/login");
          } else {
            this.btnLoading = false;
            this.MessageBar("E", response.data.errMsg);
          }
        })
        .catch((error) => {
          this.closeBtn = false;
          this.$emit("DgOverlay1")
          this.disable = false;
          this.btnLoading = false;
          this.MessageBar("E", error);
        });
    },
    // Use a loop to filter out elements whos value is null
    checkBids(arr) {
      for (let i = 0; i < arr.length; i++) {
        if (arr[i].signal == "O") {
          arr.splice(i, 1);
          i--;
        } else {
          if (arr[i].price == 0 && arr[i].quantity == 0) {
            arr.splice(i, 1);
            i--;
          }
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
    removeValue(arr) {
      var testArr = [];
      for (let rv = 0; rv < arr.length; rv++) {
        testArr.push({
          id: arr[rv].id,
          activityType: arr[rv].activityType,
          bidReferenceNo: arr[rv].bidReferenceNo,
          cutOff: arr[rv].cutOff,
          lotSize: arr[rv].lotSize,
          price: arr[rv].price,
          quantity: arr[rv].quantity,
        });
      }
      this.orderInput.bids = testArr;
    },
    // To hide the Update button
    hideUpdate(bool) {
      this.isError = bool
    },
  },
  // updated() {
  //   this.updateMaxval()
  // },
  watch: {
    categoryArr: {
      //  deep: true,
      immediate: true,
      handler() {
        if (this.modifyData.category != "") {
          this.updateMaxval()
        }
      }
    },
    modifyData: {
      handler(value) {
        for (let i = 0; i < value.modifyDetails.length; i++) {
          if (value.modifyDetails[i].signal == "N") {
            if (
              value.modifyDetails[i].price != 0 &&
              value.modifyDetails[i].quantity != 0
            ) {
              if (this.amountPayable <= this.MaxValue) {
                this.$emit("showUpdate");

              }
            }
          }
        }
        // this.updateMaxval()
      },
      deep: true,
    },
    amountPayable: {
      handler(value) {
        if (value > this.MaxValue) {
          this.msg = true;
          this.$emit("closeModifyBtn");
        } else {
          this.msg = false;
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
}
</style>