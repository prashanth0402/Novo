<template>
  <div>

    <!-- Header -->


    <v-layout row wrap class="mb-5 mt-3" style="padding-left: 12px;">
      <div class="d-flex align-center">
        <v-img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/tresur.png" height="25" width="25" contain
          class="d-flex"></v-img>
        <span class="font-weight-bold" style="margin-left: 13px;">T-Bills</span>
      </div>
    </v-layout>

    <!-- values -->
    <v-layout class="my-10" v-if="billdynamicItem.length == 0">
      <v-slide-x-transition mode="out-in" appear>
        <v-flex class="ml-5 text text--disabled">
          <!-- <p v-if="billsresult.length == 0 && Flag == 'I'">No T-Bills are open for sale currently.</p>
          <p v-if="billshistory.length == 0 && Flag == 'O'">You haven't invested in any T-Bills.</p> -->
          <p>
            {{ billdynamicText }}
          </p>
        </v-flex>
      </v-slide-x-transition>
    </v-layout>

    <!-- card -->
    <v-card elevation="0" v-if="billdynamicItem.length != 0">

      <!-- search bar -->
      <v-card-title v-if="Flag == 'O'">
        <v-text-field v-model="search" dense append-icon="mdi-magnify" label="Search" single-line></v-text-field>
      </v-card-title>
      <v-slide-x-transition mode="out-in" appear>
        <v-card-text>

          <v-data-table :headers="billdynamicHeader" :items="billdynamicItem" :search="search" :key="tableKey"
            :footer-props="Flag == 'O' ? { 'items-per-page-options': [5, 10, 15, -1] } : { 'items-per-page-options': [10] }"
            :loading="loading" loading-text="Loading... Please wait" no-data-text="No Records available">

            <template v-if="Flag === 'O'" v-slot:item.requestedUnit="{ item }">
              <v-menu top offset-x>
                <template v-slot:activator="{ on, attrs }">
                  <span>{{ item.requestedUnit }}</span>
                  <v-icon right small @click="displayDetail" v-bind="attrs" v-on="on"
                    color="primary">mdi-information-outline</v-icon>
                </template>
                <detailCard :master="item" :showDetail="showDetail" @closeDetail="displayDetail" />
              </v-menu>
            </template>

            <template v-slot:item.status="{ item }">
              <span :class="item.statusColor == 'G' ? 'green--text text-capitalize' : 'red--text'">
                {{ item.orderStatus }}
              </span>

            </template>

            <template v-slot:item.amount="{ item }">
              <span v-if="item.amount !== 0">₹</span>
              <span>{{ formatedPrice(item.amount) }}</span>
            </template>


            <template v-slot:item.requestedAmount="{ item }">
              <span v-if="item.requestedAmount !== 0">₹</span>
              <span>{{ formatedPrice(item.requestedAmount) }}</span>
            </template>

            <!-- Button -->
            <template v-slot:item.actions="{ item }" class="elevation-0">
              <v-layout>
                <v-flex class="d-flex align-center">
                  <v-hover v-slot="{ hover }">
                    <v-btn small :disabled="item.diableActionBtn == undefined ? true : item.diableActionBtn" width="100"
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



          </v-data-table>

        </v-card-text>
      </v-slide-x-transition>

    </v-card>

  </div>
</template>
  
<script>
import detailCard from "../../Tab/detailCard.vue"
export default {
  name: "TBillTable",
  data() {
    return {
      search: '',
      tabs: ['I', 'O'],
      // store: [],
      header: [],
      Invest: [
        { text: 'Security Name', align: 'start', sortable: false, value: "symbol", },
        { text: 'Indicative yield', sortable: false, align: "start", value: 'name' },
        { text: 'Bid Close date', sortable: false, align: "start", value: 'endDateWithTime' },
        { text: 'Unit Limits', sortable: false, align: "start", value: 'priceRange' },
        { text: 'Amount', sortable: false, align: "start", value: 'amount' },
        { text: '', sortable: false, align: "start", value: "actions" },
      ],

      Order: [
        { text: 'Security Name', align: 'start', sortable: false, value: "symbol", },
        { text: 'Int.Order No', sortable: false, align: "start", value: 'respOrderNo' },
        { text: 'Order No', sortable: false, align: "start", value: 'orderNo' },
        { text: 'Bid Order date', sortable: true, align: "start", value: 'orderDate' },
        { text: 'Unit Price', sortable: false, align: "start", value: 'requestedUnitPrice' },
        { text: 'Units', sortable: true, align: "center", value: 'requestedUnit' },
        { text: 'Amount', sortable: true, align: "start", value: 'requestedAmount' },
        { text: 'Status', sortable: false, align: "start", value: 'status' },
      ],

      showDetail: false,
    }
  },
  methods: {

    sendTo(item, indicator) {
      this.$emit("passVal", item, indicator)
    },
    formatedPrice(item) {
      if (item != undefined) {
        return item.toLocaleString('en-IN');
      }
    },
    // buttonText(item) {
    //   if (item != undefined) {
    //     return this.$globalData.currentTime >= item.startTime && this.$globalData.currentTime <= item.endTime ? "PlaceOrder" : "Offline"
    //   } else {
    //     return "PlaceOrder"
    //   }
    // },
    // Modbtn(item) {
    //   if (item != undefined) {
    //     return (this.$globalData.currentTime <= item.startTime && this.$globalData.currentTime >= item.endTime)
    //   }
    // },
    displayDetail() {
      this.showDetails = !this.showDetails
    },

    // show(flag) {
    //   // this.$emit("showCircle")
    //   // setTimeout(() => {
    //     if (flag == "I") {
    //       this.Flag = flag;
    //       this.header = this.Invest;
    //       if (this.billsresult != null) {
    //         this.store = this.billsresult;
    //       } else {
    //         this.store = []
    //       }

    //     } else {
    //       this.Flag = flag;
    //       this.header = this.Order;
    //       this.store = this.billshistory;
    //     }
    //     // this.$emit("showCircle")
    //   // }, 250);
    // },
  },
  computed: {
    billdynamicHeader() {
      return this.Flag == "I" ? this.Invest : this.Order
    },
    billdynamicItem: {

      get() {
        if (this.Flag == "I" && this. tbillmasterFound == "Y") {
          return this.billsresult
        } else if (this.Flag == "O" && this.tbillhistoryFound == "Y") {
          return this.billshistory
        } else {
          return []
        }
      }
    }
    // return this.Flag == "I" ? this.billsresult : this.billshistory

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
  }
  // created() {
  //   this.show(this.Flag);
  // },
  // watch: {
  //   Flag: {
  //     handler(newFlag) {
  //       this.show(newFlag);
  //     },
  //     immediate: true,
  //   },
  //   tableKey: {
  //     handler() {
  //       this.show(this.Flag);
  //     },
  //     immediate: true,
  //   }
  // },
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


::v-deep .v-data-footer {
  justify-content: end;

}
</style>