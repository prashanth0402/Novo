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
          {{ dynamicText }}
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
    <v-card elevation="0" v-if="dynamicItem.length != 0">
      <v-card-title class="pa-1" v-if="Flag == 'O'">
        <!-- <v-progress-circular class="circular" indeterminate color="primary"></v-progress-circular> -->
        <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details></v-text-field>
      </v-card-title>
      <v-card-text>
        <v-slide-x-transition mode="out-in" appear>
          <v-data-table :key="tableKey" :headers="dynamicHeader" :items="dynamicItem" :search="search"
            :single-expand="singleExpand" :expanded.sync="expanded" :show-expand="Flag == 'O'" item-key="id"
            :loading="Load" loading-text="Loading... Please wait" class="table-text" no-results-text="Record not found"
            :footer-props="Flag == 'O' ? { 'items-per-page-options': [10] } : {
              'items-per-page-options': [10]
            }" @click:row="handleRowClick">
            <!-- :footer-props="Flag == 'O' ? { 'items-per-page-options': [5, 10, 15, -1] } : {
              'items-per-page-options': [10]
            }" -->
            <template v-slot:item.symbol="{ item }" v-if="Item != [] && Flag == 'I'">
              <v-row>
                <v-col>
                  <v-layout>
                    <v-flex v-if="item.blogLink == ''" class="text-uppercase">{{ item.symbol }}
                      <v-chip v-show="item.sme == true" x-small color="blue lighten-5" label
                        class="smechip--text mb-1">SME</v-chip>
                      &nbsp;
                      <v-menu top offset-x v-if="item.exchange == 'NSE'">
                        <template v-slot:activator="{ on, attrs }">
                          <v-icon left x-small @click="displayDemand(item)" v-bind="attrs"
                            v-on="on">mdi-information-outline</v-icon>
                        </template>
                        <Marketdata :master="item" :showDemand="showDemand" @closeDemand="closeDemand" />
                      </v-menu>
                    </v-flex>
                    <v-flex v-else>
                      <a target="_blank" :href="item.blogLink" color="black--text">{{ item.symbol }} </a>
                      <v-chip v-show="item.sme == true" x-small color="blue lighten-5" label
                        class="smechip--text mb-1">SME</v-chip>
                      &nbsp;
                      <v-menu top offset-x v-if="item.exchange == 'NSE'">
                        <template v-slot:activator="{ on, attrs }">
                          <v-icon left x-small @click="displayDemand(item)" v-bind="attrs"
                            v-on="on">mdi-information-outline</v-icon>
                        </template>
                        <Marketdata :master="item" :showDemand="showDemand" @closeDemand="closeDemand" />
                      </v-menu>
                    </v-flex>

                  </v-layout>
                  <v-flex class="text-small text-capitalize">{{ item.name }}</v-flex>

                </v-col>
              </v-row>
            </template>

            <template v-slot:item.bidDate="{ item }" v-if="Item != [] && Flag == 'I'">
              <v-layout class="d-flex flex-column">
                <v-flex>
                  <span class="text-wrapper">{{ item.bidDate }}</span>
                </v-flex>
              </v-layout>
            </template>

            <template v-slot:item.priceRange="{ item }" v-if="Item != [] && Flag == 'I'">
              <span class="text-wrapper">{{ item.priceRange }}</span>
            </template>
            <template v-slot:item.actions="{ item }" v-if="Item != [] && Flag == 'I'">
              <v-layout class="d-flex flex-column">
                <v-flex>

                  <!--
                // This below button is the combination of all the scenarios
                <v-btn small depressed :class="item.flag == 'Y' || item.pending == 'P' ? 'blue lighten-4 primary--text' :
                    item.preApply != 'pre' ? 'primary px-6' : 'primary'" v-bind="attrs" v-on="on"
                    :disabled="Modbutton && item.flag == 'Y' || item.pending == 'P' || item.upcoming == 'U' ? true : false"
                    :style="text == 'OFFLINE' ?
                      'width:70px;font-size:10px;' : item.upcoming == 'U' ||
                        item.preApply == 'pre' ? 'width:76px;font-size:10px;' : undefined">
                    {{ item.flag == 'Y' || item.pending == 'P' ? "Modify" : item.upcoming == 'U' &&
                      item.preApply == 'NA' ? "Upcoming" : item.preApply == "pre" ? "Pre-apply" : text }}
                  </v-btn>
                
                -->
                  <v-tooltip top color="rgba(17,17,17,.9)" transition="slide-y-reverse-transition"
                    v-if="item.category != ''">
                    <template v-slot:activator="{ on, attrs }">

                      <v-btn small depressed
                        @click=" item.applicationNo != '' ? signal('modify', item.id, item) : signal('bid', item.id, item)"
                        class="text-capitalize" :class="item.flag == 'Y' || item.pending == 'P' ? 'blue lighten-4 primary--text' :
                          item.preApply != 'pre' ? 'primary px-6' : 'primary'"
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
                    <span class="white--text tooltip-text" v-if="item.category != ''">Apply as
                      {{ item.category
                      }}</span>
                  </v-tooltip>

                  <v-menu top offset-y-reverse offset-x-reverse v-else>
                    <template v-slot:activator="{ on, attrs }">
                      <v-btn small depressed @click="getCategoryList(item.id)" :class="item.flag == 'Y' || item.pending == 'P' ? 'blue lighten-4 primary--text' :
                        item.preApply != 'pre' ? 'primary px-6' : 'primary'" v-bind="attrs" v-on="on"
                        :disabled="item.upcoming == 'U' && item.preApply != 'pre' ? true : false" :style="text == 'Offline' ?
                          'width:70px;font-size:10px;' : item.upcoming == 'U' ||
                            item.preApply == 'pre' ? 'width:76px;font-size:10px;' : undefined">
                        <span class="text-capitalize">
                          {{ item.flag == 'Y' || item.pending == 'P' ? "Modify" :
                            item.upcoming == 'U' &&
                              item.preApply == 'NA' ? "Upcoming" : item.preApply == "pre" ? "Pre-apply" : text }}
                        </span>
                      </v-btn>
                      <!-- :disabled="Modbutton && item.flag == 'Y' || item.pending == 'P' || item.upcoming == 'U' ? true : false" -->
                    </template>
                    <v-list outlined>
                      <v-list-item v-for="(category, index) in categoryArr" :key="index">
                        <v-list-item-title
                          @click=" category.flag != 'N' ? signal('modify', id, item, category,) : signal('bid', id, item, category)">
                          {{ category.text }}
                          <v-icon x-small left class="success--text mb-1"
                            v-show="category.flag == 'Y'">mdi-check-circle-outline</v-icon>
                        </v-list-item-title>
                      </v-list-item>
                    </v-list>
                  </v-menu>
                </v-flex>
              </v-layout>
            </template>

            <!-- ------- (History stepper Peace) ---------- -->
            <template v-slot:item.status="{ item }">
              <v-layout>
                <v-flex v-if="item.status == 'success' && item.cancelFlag == 'N'" class="success--text text-capitalize">
                  {{ item.status }}
                </v-flex>
                <v-flex v-else-if="item.status == 'success' && item.cancelFlag == 'Y'"
                  class="error--text text-capitalize">
                  cancelled
                </v-flex>
                <v-flex v-else-if="item.status == 'failed'" class="error--text text-capitalize">
                  {{ item.status }}
                </v-flex>
                <v-flex v-else class="warning--text lighten-5 text-capitalize">
                  {{ item.status }}
                </v-flex>
              </v-layout>
            </template>
            <template v-slot:item.applicationStatus="{ item }">
              <v-layout>
                <v-flex v-if="item.cancelFlag == 'Y'">-</v-flex>
                <v-flex v-else class="text-capitalize">{{ item.applicationStatus }}</v-flex>
              </v-layout>
            </template>
            <template v-slot:item.upiStatus="{ item }">
              <v-layout>
                <v-flex v-if="item.cancelFlag == 'Y'">-</v-flex>
                <v-flex v-else class="text-capitalize">{{ item.upiStatus }}</v-flex>
              </v-layout>
            </template>
            <template v-slot:expanded-item="{ headers, item }" lg12>
              <td style="background-color: white;" :colspan="headers.length">
                <!-- <span v-if="item.cancelFlag == 'Y'"> More info about {{ item.upiStatus }}</span> -->
                <Steppers :StepStruct="item" />
              </td>
            </template>
          </v-data-table>
        </v-slide-x-transition>
      </v-card-text>
      <HistoryRecDialog :HistoryRec="HistoryRec" :ShowHistoryRec="ShowHistoryRec" @closeHistoryRec="closeHistoryRec" />
    </v-card>
  </div>
</template>

<script>
import HistoryRecDialog from '../../ApplicationHistory/HistoryRecDialog.vue';
import Steppers from '../../IpoWindow/IpoWindowComp/Steppers.vue';
import EventServices from '@/services/EventServices';
import Marketdata from './MarketData/Marketdata.vue';
export default {
  components: {
    HistoryRecDialog,
    Steppers,
    Marketdata
  },
  data() {
    return {
      // currentTime: "",
      text: "Bid",
      Day: 0,
      id: 0,
      loading: false,
      Modbutton: false,
      Bidbutton: false,
      search: "",
      Invest: [
        { text: "Instrument", sortable: false, align: "left", value: "symbol" },
        { text: "Bidding Date", sortable: false, align: "center", value: "bidDate" },
        { text: "Price range", sortable: false, align: "center", value: "priceRange" },
        { text: "Minimum qty", sortable: false, align: "center", value: "minBidQuantity" },
        { text: "", sortable: false, align: "center", value: "actions" },
        { text: "", sortable: false, align: "center", value: "data-table-expand" },
      ],
      Order: [
        { text: "Instruments", align: "left", value: "symbol" },
        { text: "Bid Date", align: "center", value: "date" },
        { text: "ApplicationNo", align: "center", value: "applicationNo" },
        { text: "Status", align: "center", value: "status" },
        { text: "Depository status", align: "center", value: "applicationStatus" },
        { text: "Upi status", align: "center", value: "upiStatus" },
        { text: "", sortable: false, align: "center", value: "data-table-expand" },
      ],
      ShowMsg: false,
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
    Items: [],
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
      this.id = id;
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
    },
    // // to find current time
    // getCurrentTime() {
    //   console.log("calling getCurrentTime");
    //   const currentTime = new Date();
    //   let hours = currentTime.getHours();
    //   let minutes = currentTime.getMinutes();
    //   let seconds = currentTime.getSeconds();
    //   hours = (hours < 10 ? "0" : "") + hours;
    //   minutes = (minutes < 10 ? "0" : "") + minutes;
    //   seconds = (seconds < 10 ? "0" : "") + seconds;
    //   //this.currentTime = `08:00:00`;

    //   // this.currentTime = `${hours}:${minutes}:${seconds}`;
    //   // this.$globalData.currentTime = this.currentTime;
    // },

    // ------->>>>>>>>>>>>>> (History Peace) <<<<<<<<<<<<<----------
    handleRowClick(item) {
      if (this.Flag != "I") {
        if (window.getSelection().toString().length === 0) {
          this.showRec(item.masterId, item.applicationNo);
          this.$emit("ReAssign", item)
          // this.rowRec = item
        }
      }
    },
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

    // ------->>>>>>>>>>>>>> (Table Swaping Peace) <<<<<<<<<<<<<----------

    // show(flag) {
    //   this.$emit("makeLoad", true);
    //   if (flag == "I") {
    //     this.header = this.Invest
    //     this.store = this.Item;
    //   } else {
    //     this.header = this.Order
    //     this.store = this.history;
    //   }
    //   this.$emit("refresh");
    //   if (this.store.length > 0) {
    //     this.$emit("makeLoad", false);
    //   } else {
    //     setTimeout(() => {
    //       this.$emit("makeLoad", false);
    //     }, 500)
    //   }

    // },
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
      return this.Flag == "I" ? this.Invest : this.Order
    },
    dynamicItem() {
      return this.Flag == "I" ? this.Item : this.history
    }
  },
  watch: {
    '$globalData.currentTime': function (time) {
      if (time >= "10:00:00" && time <= "17:00:00") {
        this.text = 'Bid'
        // this.Modbutton = true;
        this.Modbutton = false;
        this.ShowMsg = false;
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
    Item: function (arr) {
      if (arr == null) {
        this.loading = true;
      } else {
        this.loading = false;
      }
    },
    // Flag: {
    //   handler(flag) {
    //     this.show(flag)
    //   }, immediate: true
    // },
    // tableKey: function () {
    //   this.show(this.Flag)
    // }
  },
}

</script>

<style scoped>
/* .v-tooltip {
  position: relative;
}

.v-tooltip::before:hover {
  position: absolute;
  content: '|';
  height: 20px;
  width: 20px;
  transform: rotate(45deg);
  background: #000;
  border-top: 1px solid #dddadabd;
  border-left: 1px solid #dddadabd;
  top: -10px;
  right: 20px !important;
} */

.tooltip-text {
  font-size: 12px;
  font-weight: 600;
}

.v-application a {
  text-decoration: none;
}

a:hover {
  text-decoration: underline;
}

.text {
  font-size: 1em;
  font-weight: 400;
}

.text-small {
  font-size: 10px;
}

.table-text {
  font-size: 13px;
}

::v-deep .v-data-footer {
  padding-left: 200px;
  justify-content: end;
}

::v-deep .v-data-table>.v-data-table__wrapper .v-data-table__mobile-row {
  height: initial;
  min-height: 5px;
}

@media screen and (min-width: 800px) {
  .text-wrapper {
    font-size: 12px;
  }
}

.v-list-item {
  min-height: 32px;
  padding: 0px 10px;
}

.v-list-item__title {
  font-size: 12px;
}

.v-list-item__title:hover {
  color: #1976d2;
  transition: .3s all ease-in-out;
}

.v-list-item:hover {
  cursor: pointer;
}

.v-menu__content {
  box-shadow: none !important;
}
</style>