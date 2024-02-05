<template>
  <v-container class="mt-10">
    <v-overlay v-if="overlay">
      <v-progress-circular :size="50" color="white" indeterminate></v-progress-circular>
    </v-overlay>
    <v-layout class="mb-2">
      <v-flex :class="this.$vuetify.breakpoint.width < 800
        ? 'd-flex flex-column'
        : 'd-flex align-end'
        " lg9>
        <v-slide-y-transition mode="out-in" appear>
          <h2 :class="this.$vuetify.breakpoint.width < 800
            ? 'font-weight-medium d-flex align-center justify-start'
            : 'font-weight-medium d-flex align-center'
            " lg8>
            <v-img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/ipologo.webp"
              lazy-src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/ipologo.webp" height="30" width="30" contain
              class="mr-2 d-flex"></v-img>
            <span class="mt-2">IPO</span>
            <!-- removed the order name. -->
            <!-- <span v-else class="mt-2">IPO Orders <span v-if="historyArr.length == 0">(0)</span></span> -->
          </h2>
        </v-slide-y-transition>
        <TabBtn @show="show" :InvestCount="InvestCount" :OrderCount="OrderCount" lg4
          :class="this.$vuetify.breakpoint.width > 800 ? 'ml-5' : 'mb-2'" />
      </v-flex>
    </v-layout>

    <!-- <center class="d-flex justify-center">
      <v-flex>
        <v-progress-circular indeterminate :size="30" width="2" color="primary" v-if="circle"></v-progress-circular>
      </v-flex>
    </center> -->

    <!-- Desktop version Table -->
    <IpoTableDesk :dynamicText="dynamicText" :Item="activeIpo" :Load="loading" @Active="Signal" @RecallIpo="fetchMaster"
      v-if="hidetab" :history="historyArr" :loading="loading" @makeLoad="makeLoad" :Flag="Flag" :tableKey="tableKey"
      :categoryArr="categoryArr" @getCategoryList="getCategoryList" />
    <!-- Mobile version Table -->
    <IpoTableMobile :dynamicText="dynamicText" :Item="activeIpo" :Load="loading" @Active="Signal" @RecallIpo="fetchMaster"
      v-else :history="historyArr" :loading="loading" @makeLoad="makeLoad" :Flag="Flag" :tableKey="tableKey"
      :categoryArr="categoryArr" @getCategoryList="getCategoryList" />

    <BidDialog :ShowBid="bid" @CloseBid="closeBid" :DetailStruct="details" @CloseBidDg="closeBidonly"
      @RecallApi="fetchMaster" @RecallIpo="fetchMaster" :categoryArr="categoryArr" :selectedCategory="selectedCategory" />
    <ModifyDialog :ShowModify="modify" @CloseModify="closeModify" :DetailStruct="details" :ModifyDetail="ModifyData"
      :showAddBtn="showAddBtn" @showUpdate="showUpdate" @hideUpdate="hideUpdate" :modifyBtn="modifyBtn"
      @closeModifyBtn="closeModifyBtn" @addNewBid="addNew" :ipoAppTotal="ipoAppTotal" @modified="modifyBid"
      @closeSlot="closeSlot" @cancelBid="cancelBid" @Recall="fetchMaster" :copiedModifyData="copiedModifyData"
      :categoryArr="categoryArr" :amountPayable="amountPayable" />
  </v-container>
</template>

<script>
import IpoTableDesk from "./IpoWindowComp/IpoTableDesk.vue";
import IpoTableMobile from "./IpoTableMobile.vue";
import BidDialog from "./IpoWindowComp/Bid/BidDialog.vue";
import ModifyDialog from "./IpoWindowComp/Modify/ModifyDialog.vue";
import EventServices from "@/services/EventServices";
import TabBtn from "../../Sgb/Tab/tabBtn.vue";
export default {
  data() {
    return {
      activeIpo: [],
      id: 0,
      bid: false,
      modify: false,
      ipoAppTotal: 0,
      details: {},
      ModifyData: {
        symbol: "",
        masterId: 0,
        appNo: "",
        upi: "",
        category: "",
        total: 0,
        errMsg: "",
        status: "",
        changed: false,
        modifyDetails: [],
      },
      signal: "",
      lotSize: 0,
      showAddBtn: false,
      modifyBtn: false,
      randomNumber: 0,
      copiedModifyData: {},
      loading: false,
      overlay: false,
      // tempAmt1: 0,
      // tempAmt2: 0,
      // tempAmt3: 0,
      // Newly added
      historyArr: [],
      Flag: "I", // It helps to make condition on @click:row
      InvestCount: 0,
      OrderCount: 0,
      detail: {},
      // modify: {},
      dialog: false,
      actionFlag: "",
      tableKey: 0,
      circle: false,
      loadingtext: "",
      categoryArr: [],
      selectedCategory: "",
      amountPayable: 0,
      limitStruct: {}
    };
  },
  computed: {
    // It helps to hide the Desktop table when screen size changes below 600px
    hidetab: {
      get() {
        if (this.$vuetify.breakpoint.width <= 600) {
          return false;
        } else {
          return true;
        }
      },
    },
    dynamicText: {
      get() {
        if (this.loading == true) {
          return this.loadingtext
        } else {
          if (this.Flag == "I" && this.activeIpo.length == 0) {
            return "No IPOs are open for sale currently."
          } else if (this.Flag == "O" && this.historyArr.length == 0) {
            return "You haven't invested in any IPOs."
          } else {
            return 'No IPOs are open for sale currently.'
          }
        }
      },
      set(text) {
        this.loadingtext = text
      }
    }
  },
  components: {
    IpoTableDesk,
    IpoTableMobile,
    BidDialog,
    ModifyDialog,
    TabBtn,
  },
  methods: {
    // This method used to pass the information about the bids from the ipo table to ipo bid / modify dialog
    async Signal(indicator, id, master, category, code) {
      // console.log(indicator, id, master.symbol, category, code);
      this.selectedCategory = category
      if (indicator == "bid") {
        this.id = id;
        this.bid = true;
        // //  To change the issuesize in item from integer to string value
        // if (master.issueSize >= 10000000) {
        //   master.issueSize = (master.issueSize / 10000000).toFixed(2) + " Crores";
        // } else if (master.issueSize >= 100000) {
        //   master.issueSize = (master.issueSize / 100000).toFixed(2) + " Lakhs";
        // } else {
        //   master.issueSize = master.issueSize.toString();
        // }
        // finally assign item to details
        this.details = master;
        this.getCategory(id)
      } else {
        this.overlay = true;
        this.id = id;
        await this.GetModifyIpo(id, code);
        this.modify = true;

        // finally assign item to details
        this.details = master;
        this.getCategory(id)
      }
      // this.getCategory(id)
    },
    getCategoryList(id) {
      this.getCategroyPurFlag(id)
    },
    // To close the bid dialog when bid was placed by the user and show the message card with certain information
    closeBid(bid) {
      this.bid = bid;
      this.$emit("showalert");
    },
    // To close the bid dialog only
    closeBidonly(bid) {
      this.bid = bid;
    },
    closeModify(mod) {
      this.modify = mod;
    },
    // This method is used to make the update btn in modify dailog normal
    showUpdate() {
      this.modifyBtn = false;
    },
    // This method is used to make the update btn in modify dailog disable
    hideUpdate(bool) {
      if (bool == true) {
        this.modifyBtn = false;
      } else {
        this.modifyBtn = true;
      }
    },
    closeModifyBtn() {
      this.modifyBtn = true;
    },
    // When user presses the Modify button this method is fetch the already placed bid details
    async GetModifyIpo(id, code) {
      this.modifyBtn = true;
      await EventServices.GetModify(id, code)
        .then((response) => {
          this.overlay = false;
          if (response.data.status == "S") {
            this.copiedModifyData = JSON.parse(JSON.stringify(response.data)); // To have the permenant copy of the struct
            this.ModifyData = response.data;
            for (let ix = 0; ix < this.ModifyData.modifyDetails.length; ix++) {
              this.ModifyData.modifyDetails[ix].signal = "O";
            }
            if (this.ModifyData.modifyDetails.length < 3) {
              this.showAddBtn = true;
            }
            this.loading = false;
          } else if (response.data.status == "I") {
            this.$router.replace("/login");
          } else {
            this.MessageBar("E", "Please wait for sometime...");
          }
        })
        .catch((error) => {
          this.overlay = false;
          this.MessageBar("E", error.response);
        });
    },
    // This method is used to fetch the Ipomaster details from the API
    fetchMaster() {
      this.activeIpo = [];
      this.loading = true;
      this.loadingtext = "Loading please wait...";
      this.fetchHistory(); // To get the history detail
      EventServices.GetActiveIpo()
        .then((response) => {
          this.loading = false;
          if (response.data.status == "S") {
            if (response.data.ipoDetail != null) {
              this.activeIpo = response.data.ipoDetail;
              this.loading = false;
              // console.log("IPOMASTER", this.activeIpo);
            } else {
              this.loading = false;
            }
            this.Flag = "I";
          } else {
            this.MessageBar("E", response.data.errMsg);
            this.loading = false;
          }
        })
        .catch((error) => {
          this.MessageBar("E", error);
          this.loading = false;
        });
    },
    getCategory(id) {
      this.$globalData.overlay = true;
      // this.categoryArr = []
      EventServices.GetCategory(id, this.$route.path)
        .then((response) => {
          this.$globalData.overlay = false;
          if (response.data.status == "S") {
            this.categoryArr = []
            this.categoryArr = response.data.categoryArr
          } else {
            this.MessageBar("E", response.data.errMsg)
          }
        })
        .catch((error) => {
          this.$globalData.overlay = false;
          this.MessageBar("E", error)
        });
    },
    getCategroyPurFlag(id) {
      this.$globalData.overlay = true;
      // this.categoryArr = []
      EventServices.GetCategroyPurFlag(id)
        .then((response) => {
          this.$globalData.overlay = false;
          if (response.data.status == "S") {
            this.categoryArr = []
            this.categoryArr = response.data.orderedCategory
          } else {
            this.MessageBar("E", response.data.errMsg)
          }
        })
        .catch((error) => {
          this.$globalData.overlay = false;
          this.MessageBar("E", error)
        });
    },
    // This method is used to add the bids struct inside the bids array in modifyInput.vue
    addNew() {
      this.generateRandomNumber();
      for (let i = 0; i < this.activeIpo.length; i++) {
        if (this.ModifyData.masterId == this.activeIpo[i].id) {
          this.lotSize = this.activeIpo[i].lotSize;
          this.signal = "N";
        }
      }
      // * push the new bid struct in ModifiyData
      this.ModifyData.modifyDetails.push({
        lotSize: 0,
        bidReferenceNo: this.randomNumber,
        cutOff: false,
        price: 0,
        quantity: 0,
        activityType: "new",
        color: "green lighten-1",
        signal: this.signal,
        id: 0,
      });
    },
    // This method is used to change the bids array index value to previous value in modifyInput.vue
    closeSlot(idx) {
      for (let ix = 0; ix < this.ModifyData.modifyDetails.length; ix++) {
        if (idx == ix) {
          if (
            this.ModifyData.modifyDetails[ix].signal == "N" &&
            this.ModifyData.modifyDetails[ix].activityType == "new"
          ) {
            this.ModifyData.modifyDetails.splice(ix, 1);
          } else {
            if (this.ModifyData.modifyDetails[ix].signal == "N") {
              this.ModifyData.modifyDetails[ix].signal = "O";
              this.ModifyData.modifyDetails[ix].activityType =
                this.copiedModifyData.modifyDetails[ix].activityType;
              this.ModifyData.modifyDetails[ix].quantity =
                this.copiedModifyData.modifyDetails[ix].quantity;
              this.ModifyData.modifyDetails[ix].price =
                this.copiedModifyData.modifyDetails[ix].price;
              this.ModifyData.modifyDetails[ix].cutOff =
                this.copiedModifyData.modifyDetails[ix].cutOff;
            }
          }
        }
      }
      let length = this.ModifyData.modifyDetails.length;
      if (length < 3) {
        this.showAddBtn = true;
      }
      // To trigger Watcher abnormally
      this.ModifyData.modifyDetails.push({});
      this.ModifyData.modifyDetails.pop();
    },
    modifyBid(index) {
      for (let ix = 0; ix < this.ModifyData.modifyDetails.length; ix++) {
        if (ix == index) {
          this.ModifyData.modifyDetails[ix].activityType = "modify";
          this.ModifyData.modifyDetails[ix].color = "orange lighten-1";
          this.ModifyData.modifyDetails[ix].signal = "N";
        }
      }
      // To trigger Watcher abnormally
      this.ModifyData.modifyDetails.push({});
      this.ModifyData.modifyDetails.pop();
    },
    // This method is used to change the activity type of the bids when user choose delete application
    cancelBid() {
      for (let idx = 0; idx < this.ModifyData.modifyDetails.length; idx++) {
        this.ModifyData.modifyDetails[idx].activityType = "cancel";
        this.ModifyData.modifyDetails[idx].signal = "N";
      }
      // To trigger Watcher abnormally
      this.ModifyData.modifyDetails.push({});
      this.ModifyData.modifyDetails.pop();
    },
    // To generate temporary bidReference number
    generateRandomNumber() {
      const randomNumber = Math.floor(Math.random() * 1000);
      this.randomNumber = randomNumber.toString().padStart(6, "0");
    },
    // This method is used to fetch the IpoHistory details from the API
    fetchHistory() {
      this.historyArr = [];
      EventServices.GetHistory()
        .then((response) => {
          if (response.data.status == "S") {
            if (response.data.history != null) {
              this.historyArr = response.data.history;
            }
          } else {
            this.MessageBar("E", response.data.errMsg)
          }
        })
        .catch((error) => {
          this.MessageBar("E", error)

        });
      this.tableKey++;
      setTimeout(() => {
        this.CalcInvestAndOrder();
      }, 200);
    },
    // This method is used to fetch the particular records for the application when user click on the order table row
    showRec(id, no) {
      this.$globalData.overlay = true;
      EventServices.GetHistoryRecors(id, no)
        .then((response) => {
          if (response.data.status == "S") {
            this.HistoryRec = response.data;
            for (let ix = 0; ix < this.HistoryRec.modifyDetails.length; ix++) {
              this.HistoryRec.modifyDetails[ix].signal = "O";
            }
            this.ShowHistoryRec = true;
            this.$globalData.overlay = false;

            //  To change the issuesize in item from integer to string value
            if (this.HistoryRec.issueSize >= 10000000) {
              this.HistoryRec.issueSize =
                (this.HistoryRec.issueSize / 10000000).toFixed(2) + " Crores";
            } else if (this.HistoryRec.issueSize >= 100000) {
              this.HistoryRec.issueSize =
                (this.HistoryRec.issueSize / 100000).toFixed(2) + " Lakhs";
            } else {
              this.HistoryRec.issueSize = this.HistoryRec.issueSize.toString();
            }
          } else if (response.data.status == "I") {
            this.$router.replace("/login");
          } else {
            this.MessageBar("E", "Please wait for sometime...");
            this.$globalData.overlay = false;
          }
        })
        .catch((error) => {
          this.MessageBar("E", error);
          this.$globalData.overlay = false;
        });
    },
    // This method is used to close the history popup
    closeHistoryRec() {
      this.ShowHistoryRec = false;
    },
    // TO make the progress circle appers when the table values are switched for a moment
    makeLoad(circle) {
      this.circle = circle;
    },
    // It used to calculate the no of application where invested and No of application has been successfully placed for orders
    CalcInvestAndOrder() {
      this.InvestCount = 0;
      this.OrderCount = 0;
      setTimeout(() => {
        for (var i = 0; i < this.activeIpo.length; i++) {
          if (this.activeIpo[i].flag == "Y") {
            this.InvestCount++;
          }
        }
        for (var j = 0; j < this.historyArr.length; j++) {
          if (
            this.historyArr[j].cancelFlag == "N" &&
            this.historyArr[j].status == "success"
          ) {
            this.OrderCount++;
          }
        }
      }, 200);
    },
    // It's used to refresh the table when the GetIpoMaster || GetIpohistory was called
    async refreshTable() {
      await this.fetchMaster();
      // Update the tableKey to trigger re-render
      this.tableKey += 1;
    },
    // This method is used to store the flag from the tab component and send it to IpoTable for conditional purpose
    show(flag) {
      this.Flag = flag;
    },
    CalcAmountpayable(value1, value2) {
      let amount = 0
      if (this.limitStruct.discountType == "A") {
        amount =
          value1.lotSize *
          parseInt(value2.quantity) *
          (parseInt(value2.price) - parseInt(this.limitStruct.discountPrice));
      } else {
        amount =
          value1.lotSize *
          parseInt(value2.quantity) *
          (parseInt(value2.price) -
            // (parseInt(this.limitStruct.discountPrice) / parseInt(value2.price)) * 100
            (parseInt(value2.price) * (parseInt(this.limitStruct.discountPrice)/100))

          );
      }
      return amount
    }
  },
  // created() {
  //   this.fetchMaster();
  //   this.Flag = "I";
  //   this.show(this.Flag)
  // },
  mounted: async function () {
    await this.fetchMaster();
    this.show(this.Flag);
    if (this.$globalData.currentTime == "") {
      setInterval(this.GetCurrentTime, 1000)
    }
  },

  watch: {
    categoryArr: function () {
      this.ModifyData.modifyDetails.push({})
      this.ModifyData.modifyDetails.pop()
    },
    ModifyData: {
      // immediate: true,
      handler(newValue) {
        this.limitStruct = this.categoryArr.find(item => item.code == this.ModifyData.category || item.value == this.ModifyData.category)
        let tempAmt1 = 0, tempAmt2 = 0, tempAmt3 = 0, discount1 = 0, discount2 = 0, discount3 = 0;
        this.ipoAppTotal = 0;
        this.amountPayable = 0;

        // console.log("limitStruct", this.limitStruct, this.ModifyData.category);
        // this.tempAmt1 = 0;
        // this.tempAmt2 = 0;
        // this.tempAmt3 = 0;
        for (let i = 0; i < newValue.modifyDetails.length; i++) {
          for (let j = 0; j < this.activeIpo.length; j++) {
            if (
              parseInt(newValue.modifyDetails[i].price) > 0 &&
              parseInt(newValue.modifyDetails[i].quantity) > 0
            ) {
              if (this.activeIpo[j].id == newValue.masterId) {
                if (i == 0) {
                  tempAmt1 =
                    parseInt(newValue.modifyDetails[0].price) *
                    parseInt(newValue.modifyDetails[0].quantity) *
                    this.activeIpo[j].lotSize;

                  discount1 = this.CalcAmountpayable(this.activeIpo[j], newValue.modifyDetails[i])
                } else if (i == 1) {
                  tempAmt2 =
                    parseInt(newValue.modifyDetails[1].price) *
                    parseInt(newValue.modifyDetails[1].quantity) *
                    this.activeIpo[j].lotSize;
                  discount2 = this.CalcAmountpayable(this.activeIpo[j], newValue.modifyDetails[i])

                } else {
                  tempAmt3 =
                    parseInt(newValue.modifyDetails[2].price) *
                    parseInt(newValue.modifyDetails[2].quantity) *
                    this.activeIpo[j].lotSize;
                  discount3 = this.CalcAmountpayable(this.activeIpo[j], newValue.modifyDetails[i])
                }
                this.ipoAppTotal = Math.max(
                  tempAmt1,
                  tempAmt2,
                  tempAmt3
                );
                this.amountPayable = Math.max(
                  discount1,
                  discount2,
                  discount3
                );
              }
              // this.ipoAppTotal = this.ipoAppTotal + (parseInt(newValue.modifyDetails[i].price) * parseInt(newValue.modifyDetails[i].quantity) *
              // parseInt(this.details.lotSize));
              if (this.ipoAppTotal > this.limitStruct.maxvalue) {
                this.modifyBtn = true;
                this.showAddBtn = false;

              } else {
                this.showAddBtn = true;
              }
            }

            if (newValue.modifyDetails.length == 3) {
              this.showAddBtn = false;
            }

            if (newValue.modifyDetails[i].signal == "O") {
              this.modifyBtn = true;
            } else {
              if (this.ipoAppTotal <= this.limitStruct.maxvalue) {
                if (
                  (newValue.modifyDetails[i].price != 0 &&
                    newValue.modifyDetails[i].quantity != 0)
                ) {
                  this.modifyBtn = false;

                } else {
                  this.modifyBtn = true;
                }
              } else {
                this.modifyBtn = true
              }
            }

          }
        }
      },
      deep: true,
    },
  },
};
</script>

<style scoped>
.v-responsive {
  flex: none;
}
</style>
