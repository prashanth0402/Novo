<template>
    <div>
        <!-- <v-layout class="mb-2 d-none d-sm-flex">
            <v-flex class="d-flex justify-start">
                <v-icon :disabled="closeIcon" @click="close()">mdi-close</v-icon>
            </v-flex>
        </v-layout> -->
    
        <v-layout class="mb-4">
            <v-flex lg8>
                <v-layout class="d-flex flex-column align-start">
                    <v-flex class="subtitle-1 font-weight-bold" v-if="IpoDetail.blogLink == ''">{{ IpoDetail.symbol
                    }}</v-flex>
                    <v-flex class="subtitle-1 font-weight-bold" v-else><a target="_blank" :href="IpoDetail.blogLink">{{
                        IpoDetail.symbol }} <v-icon size="15" color="#1976d2">mdi-open-in-new</v-icon></a></v-flex>
                    <v-flex class="text ">{{ IpoDetail.name }}</v-flex>
                    <!-- <v-flex>
                <v-menu top offset-x>
                    <template v-slot:activator="{ on, attrs }">
                        <span class="primary--text">Market Demand</span>
                        <v-icon right x-small @click="displayDemand()" v-bind="attrs" class="primary--text"
                            v-on="on">mdi-information-outline</v-icon>
                    </template>
                    <Marketdata :master="IpoDetail" :showDemand="showDemand" @closeDemand="closeDemand" />
                    
                </v-menu>
            </v-flex> -->
                </v-layout>
            </v-flex>
            <v-flex lg4 class="d-flex justify-end">
                <!-- ipodetail value changed to boolean value by pavithra -->
                <v-chip v-show="IpoDetail.sme == true" x-small color="blue lighten-5" label
                    class="smechip--text">SME</v-chip>
            </v-flex>
            
        </v-layout>
        <v-row>
            <v-col cols="12">
                <v-divider></v-divider>

                <v-layout class="ma-1">
                    <v-flex>Application No</v-flex>
                    <v-flex class="d-flex justify-end font-weight-black">{{ modifyData.appNo }}</v-flex>
                </v-layout>

                <v-divider></v-divider>

                <v-layout class="ma-1">
                    <v-flex>Issue Date</v-flex>
                    <v-flex class="d-flex justify-end font-weight-black">{{ IpoDetail.bidDate }}</v-flex>
                </v-layout>

                <v-divider></v-divider>

                <v-layout class="ma-1">
                    <!-- Issue Size word removed by pavithra -->
                    <!-- <v-flex>Issue size <span class="body-2">(No. of shares)</span></v-flex> -->
                    <v-flex>No. of shares</v-flex>
                    <v-flex class="d-flex justify-end font-weight-black">{{ formatIssuesize(IpoDetail.issueSize) }}</v-flex>
                </v-layout>

                <v-divider></v-divider>

                <v-layout class="ma-1">
                    <v-flex>Issue price </v-flex>
                    <v-flex class="d-flex justify-end font-weight-black">{{ IpoDetail.priceRange }}</v-flex>
                </v-layout>

                <v-divider></v-divider>

                <v-layout class="ma-1">
                    <v-flex>Lot size</v-flex>
                    <v-flex class="d-flex justify-end font-weight-black">{{ IpoDetail.lotSize }}</v-flex>
                </v-layout>
                <v-divider></v-divider>

                <v-layout class="ma-1">
                    <v-flex>Discount</v-flex>
                    <v-flex class="d-flex justify-end font-weight-black">{{ discountStruct.discountType == "A" &&
                        discountStruct.discountPrice == 0 ? 'N/A' : discountStruct.discountType == "P" ?
                        discountStruct.discountPrice + "%" : "₹" + discountStruct.discountPrice }}</v-flex>
                </v-layout>
            </v-col>
        </v-row>

        <!-- {{ discountStruct }} -->

        <v-layout class="ma-2" v-if="IpoDetail.drhpLink != ''">
            <v-flex class="d-flex align-center primary--text"><a target="_blank" :href="IpoDetail.drhpLink"
                    class="caption">DRHP
                    <v-icon size="10" color="#1976d2">mdi-open-in-new</v-icon></a></v-flex>
        </v-layout>
        <v-layout class="ma-1">
            <v-flex justify-start>
                <v-card class="elevation-0 card-color pa-2" v-if="IpoDetail.sme == true">
                    <v-card-text class="caption">
                        This stock belongs to the SME (Small & Medium enterprises) segment which usually has low liquidity and is hence also riskier. It will be traded in a lot size of {{ IpoDetail.lotSize }} shares after listing. Selling in secondary market post listing is subjective to the RMS policy of FLATTRADE. Please refer RMS policy for more details.
                    </v-card-text>
                </v-card>
            </v-flex>
        </v-layout>
    </div>
</template>
  
<script>
// import Marketdata from '../../../MarketData/Marketdata.vue';

export default {
    name: "issueDetails",
    data(){
        return{showDemand:false}
    },
    props: {
        IpoDetail: {},
        modifyData: {},
        closeIcon: Boolean,
        discountStruct: {}
    
    },
    methods: {
        close() {
            this.$emit("closeModify");
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
    // components: { Marketdata }
};
</script>
<style scoped>
.v-application a {
    text-decoration: none;
}

/* a{
    text-decoration: none;
} */
.text {
    font-size: 12px;
    color: black;
}

.v-card__subtitle,
.v-card__text {
    padding: 5px;
}

.card-color {
    background-color: #F0F0F0;
}
</style>
  