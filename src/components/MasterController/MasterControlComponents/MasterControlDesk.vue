<template>
    <div>

        <v-tabs v-model="activeTab" left>
            <v-tab>
                <!-- <v-icon left color="blue darken-4" medium> mdi-android</v-icon> -->
                IPO
            </v-tab>
            <v-tab>
                <!-- <v-icon left color="blue darken-4" medium> mdi-apple </v-icon> -->
                SGB
            </v-tab>
            <v-tab>
                <!-- <v-icon left color="blue darken-4" medium> mdi-apple </v-icon> -->
                NCB
            </v-tab>
        </v-tabs>
        <v-window v-model="activeTab" class="elevation-0 mt-5">
            <v-window-item v-for="(tab, index) in tabs" :key="index">
                <v-toolbar flat>
                            <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details
                                class="rounded px-2"></v-text-field>
                            <!-- <v-divider class="mx-4" inset vertical></v-divider> -->
                            <!-- <v-btn color="primary" text small class="mb-2 text-capitalize" @click="openNewItemDialog()">
                                + Add
                            </v-btn> -->
                        </v-toolbar>
                <v-data-table :search="search" :headers="currentHeaders" :items="currentItems" :items-per-page="10"
                            fixed-header :loading="loading" :item-class="itemRowBackground">
                            <template v-slot:item.id="{ item, index }" >
                                    <span>{{ index + 1 }}</span>
                            </template>
                            <template v-slot:item.actions="{ item }">
                                <!-- <v-hover v-slot="{ hover }"> -->
                                    <!-- <v-btn small 
                                        :class="hover ? 'secondary white--text' : 'blue lighten-4 primary--text'"> -->
                                        <v-icon v-if="item.softDelete=='N'" small @click="editItem(item,tab)" class="red--text">
                                            mdi-trash-can-outline
                                        </v-icon>
                                        <v-icon v-else small @click="editItem(item,tab)" class="green--text">
                                            mdi-restart
                                        </v-icon>
                                    <!-- </v-btn> -->
                                <!-- </v-hover> -->
                            </template>
                        </v-data-table>
            </v-window-item>
        </v-window>
        <MasterControlDialog :ControlDialog="ControlDialog" :CurrentTittle="CurrentTittle"  @closeonly="closeonly" @closeDialog="closeDialog"
         :item="CurItem"></MasterControlDialog>
    </div>
    </template>
<script>
import MasterControlDialog from "./MasterControlDialog.vue"
export default{
    components:{MasterControlDialog},
    data() {
        return {

            ControlDialog:false,
            loading:false,
            search:"",
            activeTab: 0,
            tabs: ["IPO", "SGB","NCB"],
            IpoMasterHeader: [
                { text: 'S.No', align: 'start', sortable: false, value: 'id' },
                { text: 'Symbol', value: 'symbol' },
                { text: 'Name', value: 'name' },
                // { text: 'BiddingStartDate', value: 'biddingStartDate' },
                // { text: 'BiddingEndDate', value: 'biddingEndDate' },
                // { text: 'DailyStartTime', value: 'dailyStartTime' },
                // { text: 'DailyEndTime', value: 'dailyEndTime' },
                // { text: 'MaxPrice', value: 'maxPrice' },
                // { text: 'MinPrice', value: 'minPrice' },
                // { text: 'MinBidQuantity', value: 'minBidQuantity' },
                // { text: 'LotSize', value: 'lotSize' },
                // { text: 'Registrar', value: 'registrar' },
                // { text: 'T1ModStartDate', value: 't1ModStartDate' },
                // { text: 'T1ModEndDate', value: 't1ModEndDate' },
                // { text: 'T1ModStartTime', value: 't1ModStartTime' },
                // { text: 'T1ModEndTime', value: 't1ModEndTime' },
                // { text: 'TickSize', value: 'tickSize' },
                // { text: 'FaceValue', value: 'faceValue' },
                // { text: 'IssueSize', value: 'issueSize' },
                // { text: 'CutOffPrice', value: 'cutOffPrice' },
                { text: 'Isin', value: 'isin' },
                { text: 'IssueType', value: 'issueType' },
                // { text: 'SubType', value: 'subType' },
                { text: 'Exchange', value: 'exchange' },
                { text: 'SoftDelete', value: 'softDelete' },
                { text: "Actions", value: "actions", sortable: false },
            ],
            SgbMasterHeader: [
                  { text: 'S.No', align: 'start', sortable: false, value: 'id' },
                  { text: 'Symbol', value: 'symbol' },
                  { text: 'Series', value: 'series' },
                  { text: 'Name', value: 'name' },
                //   { text: 'IssueType', value: 'issueType' },
                //   { text: 'LotSize', value: 'lotSize' },
                //   { text: 'FaceValue', value: 'faceValue' },
                //   { text: 'MinBidQuantity', value: 'minBidQuantity' },
                //   { text: 'MinPrice', value: 'minPrice' },
                //   { text: 'MaxPrice', value: 'maxPrice' },
                //   { text: 'TickSize', value: 'tickSize' },
                //   { text: 'BiddingStartDate', value: 'biddingStartDate' },
                //   { text: 'BiddingEndDate', value: 'biddingEndDate' },
                //   { text: 'DailyStartTime', value: 'dailyStartTime' },
                //   { text: 'DailyEndTime', value: 'dailyEndTime' },
                //   { text: 'T1ModStartDate', value: 't1ModStartDate' },
                //   { text: 'T1ModEndDate', value: 't1ModEndDate' },
                //   { text: 'T1ModStartTime', value: 't1ModStartTime' },
                //   { text: 'T1ModEndTime', value: 't1ModEndTime' },
                  { text: 'Isin', value: 'isin' },
                //   { text: 'IssueSize', value: 'issueSize' },
                //   { text: 'IssueValueSize', value: 'issueValueSize' },
                //   { text: 'MaxQuantity', value: 'maxQuantity' },
                //   { text: 'AllotmentDate', value: 'allotmentDate' },
                //   { text: 'IncompleteModEndDate', value: 'incompleteModEndDate' },
                  { text: 'Exchange', value: 'exchange' },
                  { text: 'Redemption', value: 'redemption' },
                { text: 'SoftDelete', value: 'softDelete' },

                  { text: "Actions", value: "actions", sortable: false },
              ],
              NcbMasterHeader: [
                  { text: 'S.No', align: 'start', sortable: false, value: 'id' },
                  { text: 'Symbol', value: 'symbol' },
                //   { text: 'Series', value: 'series' },
                  { text: 'Name', value: 'name' },
                //   { text: 'LotSize', value: 'lotSize' },
                //   { text: 'FaceValue', value: 'faceValue' },
                //   { text: 'MinBidQuantity', value: 'minBidQuantity' },
                //   { text: 'MinPrice', value: 'minPrice' },
                //   { text: 'MaxPrice', value: 'maxPrice' },
                //   { text: 'TickSize', value: 'tickSize' },
                //   { text: 'Cutoffprice', value: 'cutOffPrice' },
                //   { text: 'BiddingStartDate', value: 'biddingStartDate' },
                //   { text: 'BiddingEndDate', value: 'biddingEndDate' },
                //   { text: 'DailyStartTime', value: 'dailyStartTime' },
                //   { text: 'DailyEndTime', value: 'dailyEndTime' },
                //   { text: 'T1ModStartDate', value: 't1ModStartDate' },
                //   { text: 'T1ModEndDate', value: 't1ModEndDate' },
                //   { text: 'T1ModStartTime', value: 't1ModStartTime' },
                //   { text: 'T1ModEndTime', value: 't1ModEndTime' },
                  { text: 'Isin', value: 'isin' },
                //   { text: 'IssueSize', value: 'issueSize' },
                //   { text: 'IssueValueSize', value: 'issueValueSize' },
                //   { text: 'MaxQuantity', value: 'maxQuantity' },
                //   { text: 'AllotmentDate', value: 'allotmentDate' },
                //   { text: 'LastDayBiddingEndTime', value: 'lastDayBiddingEndTime' },
                  { text: 'IssueValueSize', value: 'rbiName' },
                  { text: 'Exchange', value: 'exchange' },
                { text: 'SoftDelete', value: 'softDelete' },
                  { text: "Actions", value: "actions", sortable: false },
              ],
              CurItem:{},
              CurrentTittle:""
     
        };
    },
    props:{
        IpoMasterData:[],
        SgbMasterData:[],
        NcbMasterData:[]
    },
    computed:{
        currentHeaders(){
  return this.activeTab == 0 ? this.IpoMasterHeader : this.activeTab== 1 ? this.SgbMasterHeader : this.NcbMasterHeader

        },
        currentItems(){
  return this.activeTab == 0 ? this.IpoMasterData :  this.activeTab== 1 ? this.SgbMasterData : this.NcbMasterData


        }
    },
    methods:{
        editItem(item,tab){
            
            this.CurItem = {...item}
            this.ControlDialog = true
this.CurrentTittle=tab
        },
        closeDialog(){
            this.$emit("GetAllMaster")
            this.ControlDialog= false
        },
        closeonly(){
            this.ControlDialog = false
        },
        itemRowBackground: function (item) {  
      if (item.softDelete == "Y") {
        return " grey lighten-2 black-text"
      } else {
        return "white black-text"
      }
      //  return item.statusResp == "S" ?  : 
    }

    }
}
</script>



