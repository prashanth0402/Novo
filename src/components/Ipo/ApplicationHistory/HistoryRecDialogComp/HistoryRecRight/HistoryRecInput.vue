<template>
  <div>
    <v-alert border="left" color="red lighten-5" dense class="">
      <span :class="this.$vuetify.breakpoint.name == 'xs' ? 'text black--text' : 'caption'">
        <v-icon small>mdi-information-outline</v-icon> You can Modify the bid between 10 AM and 5 PM
      </span>
    </v-alert>
    <v-card :class="this.$vuetify.breakpoint >= 'sm' ? 'pa-1' : 'pa-5 mb-3 mt-md-5 mt-lg-5'" elevation="0"
      color="blue lighten-5" rounded="0">
      <v-layout>
        <v-flex :class="!hidetab ? 'text d-flex flex-column justify-start' : 'button'">
          <v-layout>
            <v-flex class="d-flex justify-center font-weight-black">
              UPI ID
            </v-flex>
          </v-layout>
          <v-layout>
            <v-flex class="d-flex justify-center">
              {{ HistoryRec.upi }}
            </v-flex>
          </v-layout>
        </v-flex>
        <v-flex :class="!hidetab ? 'd-flex flex-column justify-center text' : 'button'" xs2>
          <v-layout>
            <v-flex class="d-flex justify-center font-weight-black"><b> Category</b></v-flex>
          </v-layout>
          <v-layout>
            <v-flex class="d-flex justify-center">
              {{ HistoryRec.category }}</v-flex>
          </v-layout>
        </v-flex>
        <v-flex :class="!hidetab ? 'text d-flex flex-column justify-end' : 'button'" xs5>
          <v-layout>
            <v-flex class="d-flex justify-center"><b> Amount payable </b>
            </v-flex>
          </v-layout>

          <v-layout>
            <v-flex class="d-flex justify-center">
              {{ HistoryRec.total }}
              <span v-show="HistoryRec.discount != 'N/A'" class="primary--text">*</span>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-layout>
    </v-card>
    <v-card class="my-2 elevation-0 pa-2 text orange lighten-5 d-flex align-center rounded-0">
      <v-layout :class="!hidetab ? 'text d-flex flex-column text-center' : 'd-flex flex-column text-center caption'">
        <v-flex><b>Depository Status</b></v-flex>
        <v-flex>{{ HistoryRec.dpStatus }}</v-flex>
      </v-layout>
      <v-layout :class="!hidetab ? 'text d-flex flex-column text-center' : 'd-flex flex-column text-center caption'">
        <v-flex><b>Upi Status</b></v-flex>
        <v-flex>{{ HistoryRec.upiStatus }}</v-flex>
      </v-layout>
    </v-card>
    <v-card class="text blue lighten-5 rounded-0 elevation-0  flex-column hidden-md-and-up">
      <v-layout class="ma-2 pt-2">
        <v-flex>Issue Date</v-flex>
        <v-flex class="d-flex justify-end">{{ HistoryRec.issueDate }}</v-flex>
      </v-layout>


      <v-layout class="ma-2 ">
        <v-flex>Issue size <span>(No. of shares)</span></v-flex>
        <v-flex class="d-flex justify-end">{{ formatIssuesize(HistoryRec.issueSize) }}</v-flex>
      </v-layout>


      <v-layout class="ma-2 ">
        <v-flex>Issue price </v-flex>
        <v-flex class="d-flex justify-end">{{ HistoryRec.issuePrice }}</v-flex>
      </v-layout>


      <v-layout class="ma-2 pb-2">
        <v-flex>Lot size</v-flex>
        <v-flex class="d-flex justify-end">{{ HistoryRec.lotSize }}</v-flex>
      </v-layout>
    </v-card>

    <div v-for="bid, idx  in HistoryRec.modifyDetails" :key="idx" class="mt-2">
      <!-- {{ modifyData.modifyDetails[idx].signal }} -->
      <ReadOnly :bid="bid" />
    </div>
    <span v-show="HistoryRec.discount != 'N/A'">* price after discount if any</span>

    <!-- <v-layout class="mb-2 d-flex d-sm-none"> -->
    <v-layout class="mb-2">
      <v-layout class="ma-2 mt-5">
        <v-flex class="d-flex justify-end">
          <v-flex class="d-flex d-sm-none" v-if="Allotmenturl && HistoryRec.errReason == ''">
            <a :href="HistoryRec.registrarLink" target="_blank" style="text-decoration: none;"> Check your Allotment
              <v-icon small color="blue lighten-1">mdi-information-outline</v-icon>
            </a>
          </v-flex>
          <v-layout class="ma-2 d-flex d-sm-none" v-if="HistoryRec.errReason != ''">
            <v-flex>
              <h5>
                <span class="error--text text"> {{ HistoryRec.errReason }}</span>
              </h5>
            </v-flex>
          </v-layout>
         
          <!-- <v-btn @click="close()" class="red lighten-1 white--text elevation-0 px-5" small>close</v-btn> -->
          <v-flex >
            <v-menu top offset-x>
              <template v-slot:activator="{ on, attrs }">
                <span class="primary--text">Market Demand</span>
                <v-icon right x-small @click="displayDemand()" v-bind="attrs" class="primary--text"
                  v-on="on">mdi-information-outline</v-icon>
              </template>
              <Marketdata v-if="this.$vuetify.breakpoint.width >= 600" :master="HistoryRec" :historyDialog="historyDialog" :showDemand="showDemand" @closeDemand="closeDemand" />
              <MarketDataMobile v-else :master="HistoryRec" :historyDialog="historyDialog" :showDemand="showDemand" @closeDemand="closeDemand"></MarketDataMobile>
            </v-menu>
          </v-flex>
          
        </v-flex>
        <v-btn class="text-capitalize elevation-0 ml-2 red lighten-2 white--text" outlined text small
          @click="close">Close</v-btn>
      </v-layout>
    </v-layout>

  </div>
</template>
<script>
import MarketDataMobile from "../../../IpoWindow/IpoWindowComp/MarketData/MarketDataMobile.vue";
import Marketdata from "../../../IpoWindow/IpoWindowComp/MarketData/Marketdata.vue";
import ReadOnly from "./ReadOnly.vue";
export default {
  name: "modifyDetail",
  data(){
    return{
      showDemand:false,historyDialog:true
    }
  },

  components: {
    ReadOnly,
    Marketdata,
    MarketDataMobile
},
  props: {
    HistoryRec: {}
  },
  methods: {
    close() {
      this.$emit("closeHistoryRec");
    },
    formatIssuesize(issueSize) {
      //  To change the issuesize in item from integer to string value
      if (issueSize >= 10000000) {
        issueSize = (issueSize / 10000000).toFixed(2) + " Crores";
      } else if (issueSize >= 100000) {
        issueSize = (issueSize / 100000).toFixed(2) + " Lakhs";
      } else {
        issueSize = issueSize.toLocaleString('en-IN');
      }
      return issueSize
    },
    displayDemand() {
      this.showDemand = true;
    },
    closeDemand() {
      this.showDemand = false;
    }
  },
  computed: {
    hidetab: {
      get() {
        if (this.$vuetify.breakpoint.name == 'xs') {
          return false
        } else {
          return true
        }
      }
    },
    Allotmenturl() {
      return this.HistoryRec.registrarLink.startsWith('http', 0) || this.HistoryRec.registrarLink.startsWith('https', 0);
    },
  },

}
</script>
<style scoped>
.text {
  font-size: 10px;
}
</style>