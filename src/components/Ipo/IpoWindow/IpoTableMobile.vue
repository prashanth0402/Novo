<template>
  <div>
    <v-layout class="my-10" v-if="dynamicItem.length == 0">
      <v-slide-x-transition mode="out-in" appear>
        <v-flex class="ml-5 text text--disabled">
          <!-- <p v-if="dynamicItem.length == 0 && Flag == 'I'">No IPOs are open for sale currently.</p>
          <p v-if="dynamicItem.length == 0 && Flag == 'O'">You haven't invested in any IPOs.</p> -->
          <!-- <p v-if="dynamicItem.length == 0">
            {{ Flag == 'I' ? "No IPOs are open for sale currently." : "You haven't invested in any IPOs." }}
          </p> -->
          <p>{{ dynamicText }}</p>
        </v-flex>
      </v-slide-x-transition>
    </v-layout>

    <v-layout class="d-flex mb-2" v-if="Flag == 'I' && ShowMsg && dynamicItem.length != 0">
      <v-flex class="d-flex justify-end">
        <div class="d-flex">
          <span class="info--text d-flex align-center subtitle-2">
            <v-icon size="20" class="info--text mr-2">mdi-information</v-icon>
            Offline applications will be acceptable
          </span>
        </div>
      </v-flex>
    </v-layout>

    <v-card class="elevation-0" v-if="dynamicItem.length != 0">
      <v-card-title class="pa-1">
        <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details></v-text-field>
      </v-card-title>

      <v-slide-x-transition mode="out-in" appear>
        <v-data-table ref="myDatatable" :key="tableKey" :headers="dynamicHeader" :items="dynamicItem" :search="search"
          :items-per-page="15" item-key="id" :loading="Load" loading-text="Loading... Please wait" class="table-text"
          no-results-text="Record not found" :footer-props="{ 'items-per-page-options': [5] }"
          @click:row="handleRowClick">
          <template v-slot:item.symbol="{ item }" v-if="Item != [] && Flag == 'I'">
            <v-row wrap class="d-flex text">
              <v-col cols="6" class="d-flex flex-column text-left">
                <span class="text--disabled">Symbol
                  &nbsp;
                  <v-menu top offset-x v-if="item.exchange == 'NSE'">
                    <template v-slot:activator="{ on, attrs }">
                      <v-icon left x-small @click="displayDemand(item)" v-bind="attrs"
                        v-on="on">mdi-information-outline</v-icon>
                    </template>
                  </v-menu>
                </span>
                <span v-if="item.blogLink == ''" class="text-uppercase">{{ item.symbol }}
                  <!-- <span v-show="item.sme == true" class="text">( SME )</span> -->
                  <v-chip v-show="item.sme == true" x-small color="blue lighten-5" label
                    class="smechip--text mb-1">SME</v-chip>
                </span>
                <span v-else>
                  <a target="_blank" :href="item.blogLink">{{ item.symbol }} </a>
                  <!-- <span v-show="item.sme == true" class="text">( SME ) &nbsp;</span> -->
                  <v-chip v-show="item.sme == true" x-small color="blue lighten-5" label
                    class="smechip--text mb-1">SME</v-chip>
                </span>
                <span class="text--disabled">Name</span>
                <span class="text-capitalize">{{ item.name }}</span>
              </v-col>
              <v-col class="d-flex flex-column text-right" cols="6">
                <span class="text--disabled">Bidding date</span>
                <span>{{ item.bidDate }}</span>
                <span class="text--disabled">Lot size</span>
                <span>{{ item.lotSize }}</span>
              </v-col>
            </v-row>
            <v-row wrap class="text">
              <v-col class="d-flex text-left" cols="6">
                <v-layout>
                  <v-flex class="d-flex flex-column">
                    <span class="text--disabled">Min-Bid (Qty)</span>
                    <span>{{ item.minBidQuantity }}</span>
                  </v-flex>
                </v-layout>
              </v-col>
              <v-col class="d-flex flex-column text-right" cols="6">
                <span class="text--disabled">Pricing</span>
                <span>{{ item.priceRange }}</span>
                <!-- <span class="text--disabled">Amount</span>
                    <span>₹ {{ item.minPrice }}</span> -->
              </v-col>
            </v-row>
            <v-row wrap class="text">
              <v-col class="text-center" cols="12">
                <!-- <v-hover v-slot="{ hover }">
                  <v-btn small
                  :disabled="item.flag == 'Y' ? Modbutton : false || item.pending == 'P' || item.upcoming == 'U' && item.preApply != 'pre' ? true : false"
                  :class="hover ? 'secondary' : item.flag == 'N' ? 'primary white--text' : 'blue lighten-4 primary--text'"
                  @click="item.flag == 'N' ? signal('bid', item.id, item) : signal('modify', item.id, item)"
                  elevation="0" width="100">
                  <span class="text-capitalize" v-if="item.flag == 'N'">
                    {{ item.preApply == 'pre' ? 'Pre-Apply' : item.upcoming == 'U' ? 'Upcoming' : text }}
                  </span>
                  <span class="text-capitalize" v-else>Modify</span>
                </v-btn>
              </v-hover> -->
                <v-tooltip top color="rgba(17,17,17,.9)" transition="slide-y-reverse-transition"
                  v-if="item.category != ''">
                  <template v-slot:activator="{ on, attrs }">
                    <v-btn small depressed
                      @click="item.applicationNo != '' ? signal('modify', item.id, item) : signal('bid', item.id, item)"
                      :class="item.flag == 'Y' || item.pending == 'P' ? 'blue lighten-4 primary--text' :
                        item.preApply != 'pre' ? 'primary px-6' : 'primary'" class="text-capitalize"
                      v-bind="item.category != '' && item.flag != 'Y' && item.applicationNo == '' ? attrs : undefined"
                      v-on="item.category != '' && item.flag != 'Y' && item.applicationNo == '' ? on : undefined"
                      :disabled="item.upcoming == 'U' && item.preApply != 'pre' ? true : false" :style="text == 'Offline' ?
                        'width:70px;font-size:10px;' : item.upcoming == 'U' ||
                          item.preApply == 'pre' ? 'width:76px;font-size:10px;' : undefined">
                      {{ item.flag == 'Y' || item.pending == 'P' ? "Modify" : item.upcoming == 'U' &&
                        item.preApply == 'NA' ? "Upcoming" : item.preApply == "pre" ? "Pre-apply" : text }}
                    </v-btn>
                    <!-- :disabled="Modbutton && item.flag == 'Y' || item.pending == 'P' || item.upcoming == 'U' ? true : false" -->
                  </template>
                  <span class="white--text tooltip-text" v-if="item.category != ''">Apply as {{ item.category
                  }}</span>
                </v-tooltip>
                <!-- <v-menu top offset-y v-else>
                  <template v-slot:activator="{ on, attrs }">
                  </template>
                  <v-list>
                    <v-list-item v-for="(category, index) in categoryArr" :key="index">
                      <v-list-item-title
                      @click=" category.flag != 'N' ? signal('modify', item.id, item, category) : signal('bid', item.id, item, category)">
                      {{ category.text }}</v-list-item-title>
                    </v-list-item>
                  </v-list> 
                </v-menu>
              -->
                <v-bottom-sheet v-model="sheet" v-else>
                  <template v-slot:activator="{ on, attrs }">
                    <v-btn small depressed @click="getCategoryList(item.id)" class="text-capitalize"
                      :class="item.flag == 'Y' || item.pending == 'P' ? 'blue lighten-4 primary--text' : item.preApply != 'pre' ? 'primary px-6' : 'primary'"
                      v-bind="attrs" v-on="on" :disabled="item.upcoming == 'U' && item.preApply != 'pre' ? true : false"
                      :style="text == 'Offline' ?
                        'width:70px;font-size:10px;' : item.upcoming == 'U' || item.preApply == 'pre' ? 'width:76px;font-size:10px;' : undefined">
                      {{ item.flag == 'Y' || item.pending == 'P' ? "Modify" : item.upcoming == 'U' &&
                        item.preApply == 'NA' ? "Upcoming" : item.preApply == "pre" ? "Pre-apply" : text }}
                    </v-btn>
                    <!-- :disabled="Modbutton && item.flag == 'Y' || item.pending == 'P' || item.upcoming == 'U' ? true : false" -->
                  </template>
                  <v-list>
                    <v-list-item v-for="(category, index) in categoryArr" :key="index">
                      <v-list-item-avatar>
                        <v-avatar size="32px" tile>
                          <!-- <img :src="`https://cdn.vuetifyjs.com/images/bottom-sheets/${tile.img}`" :alt="tile.title"> -->
                          <v-icon> {{ category.code == 'IND' ? 'mdi-account-outline' : category.code == 'EMP' ?
                            'mdi-account-tie-outline' : 'mdi-account-star-outline' }}</v-icon>
                        </v-avatar>
                      </v-list-item-avatar>
                      <v-list-item-title
                        @click=" category.flag != 'N' ? signal('modify', id, rowRec, category) : signal('bid', id, rowRec, category)">{{
                          category.text }}
                        <v-chip x-small label class="success lighten-5 success--text" v-show="category.flag == 'Y'">
                          <span>Applied</span>
                          <v-icon x-small right color="success">mdi-check-circle-outline</v-icon>
                        </v-chip>
                      </v-list-item-title>
                    </v-list-item>
                  </v-list>
                </v-bottom-sheet>

              </v-col>
            </v-row>
          </template>

          <!-- Order Mobile -->
          <template v-slot:item.symbol="{ item }" v-else>
            <v-row wrap class="d-flex text">
              <v-col cols="6" class="d-flex flex-column text-left">
                <span class="text--disabled text-uppercase">Symbol</span>
                <span>{{ item.symbol }}</span>
                <!-- <span class="text--disabled">Unit</span>
                  <span>{{ item.unit }}</span> -->
              </v-col>
              <v-col class="d-flex flex-column text-right text" cols="6">
                <span class="text--disabled">Ordered Date</span>
                <span>{{ item.date }}</span>
              </v-col>
            </v-row>
            <v-row wrap class="text">
              <v-col class="d-flex text-left" cols="6">
                <v-layout>
                  <v-flex class="d-flex flex-column">
                    <span class="text--disabled">Order no.</span>
                    <span>{{ item.applicationNo }}</span>
                  </v-flex>
                </v-layout>
              </v-col>
              <v-col class="d-flex justify-space-between flex-column text-right" cols="6">
                <span class="text--disabled">Status</span>
                <span v-if="item.cancelFlag == 'N'"
                  :class="item.status == 'success' ? 'text-capitalize success--text' : 'text-capitalize error--text darken-5'">
                  {{ item.status }}
                </span>
                <span v-else class="error--text"> Bid cancelled</span>
              </v-col>
            </v-row>
          </template>
        </v-data-table>
      </v-slide-x-transition>
    </v-card>
    <MarketDataMobile :master="rowRec" :showDemand="showDemand" @closeDemand="closeDemand"></MarketDataMobile>
    <HistoryRecDialog :HistoryRec="HistoryRec" :ShowHistoryRec="ShowHistoryRec" @closeHistoryRec="closeHistoryRec"
      :Item="rowRec" :Flag="Flag" />
  </div>
</template>

<script>
import HistoryRecDialog from '../ApplicationHistory/HistoryRecDialog.vue';
import EventServices from '@/services/EventServices';
// import Marketdata from './IpoWindowComp/MarketData/Marketdata.vue';
import MarketDataMobile from './IpoWindowComp/MarketData/MarketDataMobile.vue';
export default {
  components: {
    HistoryRecDialog,
    // Marketdata,
    MarketDataMobile
  },
  data() {
    return {
      // currentTime: "",
      text: "Bid",
      Day: 0,
      loading: false,
      Modbutton: false,
      Bidbutton: false,
      search: "",
      InvestMobile: [{ text: "", sortable: false, align: "center", value: "symbol" },],
      OrderMobile: [{ text: "", sortable: false, align: "center", value: "symbol" },],
      id: 0,
      ShowMsg: false,
      sheet: false,
      // newlyadded
      header: [],
      // store: [],
      HistoryRec: {},
      ShowHistoryRec: false,
      singleExpand: true,
      expanded: [],
      rowRec: {},
      showDemand: false,
      offeredQty: 0
    };
  },
  props: {
    Item: [],
    history: [],
    categoryArr: [],
    Load: Boolean,
    Flag: String,
    tableKey: Number,
    dynamicText: String
  },
  methods: {
    getCategoryList(id) {
      this.id = id
      this.rowRec = this.Item.find(item => item.id === id);
      this.$emit("getCategoryList", id)
    },
    signal(flag, id, master, category) {
      let appNo = category === undefined ? master.applicationNo : category.applicationNo;
      let categoryText = category === undefined ? master.category : category.text;
      let code = category === undefined ? master.code : category.code;
      if (flag == "modify") {
        if (this.Modbutton) {
          this.showRec(id, appNo)
        } else {
          if (master.preApply == "pre") {
            this.showRec(id, appNo)
          } else {
            this.$emit("Active", flag, id, master, categoryText, code);
          }
        }
      } else {
        this.$emit("Active", flag, id, master, categoryText, code);
      }
      this.sheet = false
    },
    // to find current time
    // getCurrentTime() {
    //   const currentTime = new Date();
    //   let hours = currentTime.getHours();
    //   let minutes = currentTime.getMinutes();
    //   let seconds = currentTime.getSeconds();
    //   hours = (hours < 10 ? "0" : "") + hours;
    //   minutes = (minutes < 10 ? "0" : "") + minutes;
    //   seconds = (seconds < 10 ? "0" : "") + seconds;
    //   //this.currentTime = `08:00:00`;

    //   this.currentTime = `${hours}:${minutes}:${seconds}`;
    // },
    //This method is used to send the row value to history dialog when the user click 
    handleRowClick(item) {
      if (this.Flag != "I") {
        if (window.getSelection().toString().length === 0) {
          this.showRec(item.masterId, item.applicationNo);
          this.$emit("ReAssign", item)
          this.rowRec = item
        }
      }
    },
    // To get the application detail when the user click the modify button
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
                (this.HistoryRec.issueSize / 10000000).toFixed(2) + "Crores";
            } else if (this.HistoryRec.issueSize >= 100000) {
              this.HistoryRec.issueSize =
                (this.HistoryRec.issueSize / 100000).toFixed(2) + "Lakhs";
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
    // This method is used to close the history dialog
    closeHistoryRec() {
      this.ShowHistoryRec = false;
    },
    displayDemand(item) {
      this.rowRec = item
      this.showDemand = true;
    },
    closeDemand() {
      this.showDemand = false;
    }
  },
  mounted() {
    // To get Day
    this.Day = new Date().getDay();
    // To get current Time
    // this.getCurrentTime();
    // setInterval(this.getCurrentTime, 1000);
  },
  computed: {
    dynamicHeader() {
      return this.Flag == "I" ? this.InvestMobile : this.OrderMobile
    },
    dynamicItem() {
      return this.Flag == "I" ? this.Item : this.history
    }
  },
  watch: {
    // Watching current time and do some functionality
    '$globalData.currentTime': function (time) {
      if (time >= "10:00:00" && time <= "17:00:00") {
        this.text = 'Bid'
        this.ShowMsg = false;
        this.Modbutton = false;
        if (this.Day === 6 || this.Day === 0) {
          this.text = 'Offline';
          this.ShowMsg = true;
          this.Modbutton = true;
        }
      } else {
        this.ShowMsg = true;
        this.$emit('BidButton',)
        this.text = 'Offline';
        this.Modbutton = true;
      }
    },
    // Watching Item struct to activate/deactive loading in the table
    Item: function (arr) {
      if (arr == null) {
        this.loading = true;
      } else {
        this.loading = false;
      }
    },
  },
}

</script>

<style scoped>
a {
  text-decoration: none;
}

/* .v-data-footer{
    padding-left: 10px;
 } */
.text {
  font-size: 10px;
}

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