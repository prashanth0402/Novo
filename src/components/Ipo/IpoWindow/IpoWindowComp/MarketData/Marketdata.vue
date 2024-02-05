<template>
    <div>
        <v-card v-click-outside="newShow" :width="demandArr.length > 0 || catwiseArr.length > 0 ? 500 : undefined" outlined
            class="caption">

            <v-card-title>
                <v-layout class="d-flex flex-column">
                    <v-flex class="d-flex flex-column justify-start">
                        <span class="font-weight-medium">{{ master.symbol }}</span>
                        <p class="caption">Subscription details</p>
                    </v-flex>
                    <!-- <v-icon @click="closeDemand" class="error--text mb-1">mdi-close</v-icon> -->
                    <v-flex>
                        <span class="caption text--disabled" v-if="text != ''">{{ text }}.</span>
                    </v-flex>
                </v-layout>

            </v-card-title>
            <v-card-text>
                <v-row>
                    <v-col v-if="demandArr.length > 0">
                        <h4>By Price</h4>
                        <v-layout class="text--disabled">
                            <v-flex>Price</v-flex>
                            <v-flex class="text-right">quantity</v-flex>
                        </v-layout>
                        <v-divider></v-divider>
                        <v-layout v-for=" demand, index  in  demandArr " :key="index">
                            <v-flex>
                                <span v-if="demand.cutoff">Cutoff({{ demand.price }})</span>
                                <span v-else>{{ demand.price }}</span>
                            </v-flex>
                            <v-flex class="text-right">{{ formatedPrice(demand.quantity) }}</v-flex>
                        </v-layout>
                        <v-divider class="my-2"></v-divider>
                        <v-layout class="d-flex ma-2 justify-space-between">
                            <v-flex class="d-flex flex-column">
                                <span>Total subscription</span>
                                <span>Offered Qty</span>
                                <span>Subscription %</span>
                            </v-flex>
                            <v-flex class="d-flex text-right flex-column">
                                <span>{{ formatedPrice(totalQty) }}</span>
                                <span>{{ formatedPrice(master.issueSize) }}</span>
                                <span :class="(totalQty / master.issueSize) < 0 ? 'error--text' : 'success--text'">{{
                                    (totalQty / parseInt(master.issueSize)).toFixed(2) }} Times </span> 
                            </v-flex>
                        </v-layout>
                    </v-col>
                    <v-divider vertical v-if="$vuetify.breakpoint.width >= 600"></v-divider>
                    <v-col v-if="catwiseArr.length > 0">
                        <h4>By Category</h4>
                        <v-layout class="text--disabled">
                            <v-flex>Category</v-flex>
                            <v-flex class="text-right">quantity</v-flex>
                        </v-layout>
                        <v-divider></v-divider>
                        <v-layout v-for=" catwise, index  in  catwiseArr " :key="index">

                            <v-flex>{{ catwise.category }}</v-flex>
                            <v-flex class="text-right">{{ formatedPrice(catwise.quantity) }}</v-flex>
                        </v-layout>
                    </v-col>
                </v-row>
            </v-card-text>
        </v-card>
    </div>
</template>

<script>
import EventServices from '../../../../../services/EventServices'
export default {
    data() {
        return {
            Id:0,
            demandArr: [],
            catwiseArr: [],
            totalQty: 0,
            Subscription: 0,
            show: "",
            text: ""
        }
    },
    methods: {
        getIpoMktdata() {
            if (this.showDemand) {
                this.$globalData.overlay = true;
                 this.Id = this.historyDialog ? this.master.masterId : this.master.id 
                EventServices.GetIpoMktData(this.Id)
                    .then((response) => {
                        this.$globalData.overlay = false;
                        if (response.data.status == "S") {
                            this.demandArr = response.data.ipoMktDemandArr == null ? [] : response.data.ipoMktDemandArr
                            this.catwiseArr = response.data.ipoMktCatwiseArr == null ? [] : response.data.ipoMktCatwiseArr
                            this.text = response.data.noDataText
                        } else {
                            this.MessageBar("E", response.data.errMsg)
                        }
                        this.CalcTotalQuantity()
                    })
                    .catch((error) => {
                        this.$globalData.overlay = false;
                        this.MessageBar("E", error)
                    })
            }
        },
        formatedPrice(item) {
            if (item != undefined) {

                //  To change the issuesize in item from integer to string value
                if (item >= 10000000) {
                    item =
                        (item / 10000000).toFixed(2) + " Crores";
                } else if (item >= 100000) {
                    item =
                        (item / 100000).toFixed(2) + " Lakhs";
                } else {
                    item = item.toLocaleString('en-IN')
                }
                return item
            }
        },
        CalcTotalQuantity() {
            for (let i = 0; i < this.demandArr.length; i++) {
                this.totalQty += this.demandArr[i].quantity
            }
        },
        closeDemand() {
            this.demandArr = [];
            this.catwiseArr = [];
            this.totalQty = 0;
            this.Subscription = 0;
            this.$emit('closeDemand')
        },
    },
    props: {
        master: Object,
        showDemand: Boolean,
        historyDialog:Boolean
    },
    computed: {
        newShow: {
            get() {
                if (this.show != '') {
                    return false
                } else {
                    this.getIpoMktdata()
                    return this.showDemand
                }
            },
            set(value) {
                this.closeDemand()
                this.show = value
            }
        },
    },
}
</script>

<style  scoped></style>