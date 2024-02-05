<template>
    <div class="">
        <v-bottom-sheet v-model="newShow" persistent :no-click-animation="true">
            <v-sheet class="text-center">
                <v-btn class="" text dense @click="closeDemand">
                    <v-icon>mdi-chevron-down</v-icon>
                </v-btn>
                <!-- <div class="py-3">
            This is a bottom sheet using the persistent prop
          </div> -->
                <v-container>
                    <v-layout class="d-flex flex-column">
                        <v-flex class="d-flex flex-column " >
                            <span class="font-weight-medium text-left">{{ master.symbol }}</span>
                            <p class="caption text-left">Subscription details</p>
                        </v-flex>
                        <!-- <v-icon @click="closeDemand" class="error--text mb-1">mdi-close</v-icon> -->
                        <v-flex>
                            <span class="caption text--disabled" v-if="text != ''">{{ text }}.</span>
                        </v-flex>
                    </v-layout>
                    <v-layout>
                        <v-row>
                            <v-col cols="12" v-if="demandArr.length > 0">
                                <v-row class="text-left">
                                    <v-col cols="6">
                                        <span>Total subscription</span>
                                    </v-col>
                                    <v-col cols="6" class="text-right">
                                        <span>{{ formatedPrice(totalQty) }}</span>
                                    </v-col>
                                    <v-col cols="6">
                                        <span>Offered Qty</span>
                                    </v-col>
                                    <v-col cols="6" class="text-right">
                                        <span>{{ formatedPrice(master.issueSize) }}</span>
                                    </v-col>
                                    <v-col cols="6">
                                        <span>Subscription %</span>
                                    </v-col>
                                    <v-col cols="6" class="text-right">
                                        <span
                                            :class="(totalQty / master.issueSize) < 0 ? 'error--text' : 'success--text'">{{
                                                (totalQty / parseInt(master.issueSize)).toFixed(2) }} Times </span>
                                    </v-col>
                                    <v-col cols="12" class="text-right mr-5">
                                        <p class="primary--text" v-if="Showdata == false" @click="Showdata = true">more
                                        </p>
                                        <p class="primary--text" v-else @click="Showdata = false, closeSheet()">less</p>
                                    </v-col>
                                </v-row>

                            </v-col>
                            <v-col cols="12">
                                <span class="d-flex justify-space-between" v-if="Showdata && demandArr.length > 0"
                                    @click="ShowPrice = !ShowPrice, ShowCategory = false">By Price <v-icon
                                        v-if="ShowPrice">mdi-chevron-up</v-icon><v-icon
                                        v-else>mdi-chevron-down</v-icon></span>
                                <!-- <v-virtual-scroll :items="demandArr" :item-height="50" > -->

                                <v-simple-table v-if="ShowPrice" height="200" :fixed-header="true">
                                    <template v-slot:default>
                                        <thead>
                                            <tr>
                                                <th class="text-left">
                                                    Price
                                                </th>
                                                <th class="text-right">
                                                    quantity
                                                </th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            <tr v-for=" demand, index  in  demandArr " :key="index">
                                                <td class="text-left" v-if="demand.cutoff">Cutoff({{ demand.price }})
                                                </td>
                                                <td class="text-left" v-else>{{ demand.price }}</td>
                                                <td class="text-right">{{ formatedPrice(demand.quantity) }}</td>
                                            </tr>
                                        </tbody>
                                    </template>
                                </v-simple-table>
                                <!-- </v-virtual-scroll> -->
                            </v-col>
                            <v-col cols="12">
                                <span class="d-flex justify-space-between" v-if="Showdata && catwiseArr.length > 0"
                                    @click="ShowCategory = !ShowCategory, ShowPrice = false">By Category <v-icon
                                        v-if="ShowCategory">mdi-chevron-up</v-icon><v-icon
                                        v-else>mdi-chevron-down</v-icon></span>

                                <v-simple-table v-if="ShowCategory" height="200" :fixed-header="true">
                                    <template v-slot:default>
                                        <thead>
                                            <tr>
                                                <th class="text-left">
                                                    Category
                                                </th>
                                                <th class="text-right">
                                                    quantity
                                                </th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            <tr v-for=" catwise, index  in  catwiseArr " :key="index">
                                                <td class="text-left">{{ catwise.category }}</td>
                                                <td class="text-right">{{ formatedPrice(catwise.quantity) }}</td>
                                            </tr>
                                        </tbody>
                                    </template>
                                </v-simple-table>
                            </v-col>
                        </v-row>
                    </v-layout>
                </v-container>
            </v-sheet>
        </v-bottom-sheet>
    </div>
</template>
<script>
import EventServices from '@/services/EventServices';
export default {
    data() {
        return {
            ShowPrice: false,
            ShowCategory: false,
            Showdata: false,
            Id: 0,
            demandArr: [],
            catwiseArr: [],
            totalQty: 0,
            Subscription: 0,
            show: "",
            text: ""
        }
    },
    props: {
        showDemand: Boolean,
        master: Object,
        historyDialog: Boolean
    },
    methods: {
        closeSheet() {
            this.ShowPrice = false
            this.ShowCategory = false
            this.Showdata = false
        },
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
            this.Id = 0
            this.$emit('closeDemand')
            this.closeSheet()
        },
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