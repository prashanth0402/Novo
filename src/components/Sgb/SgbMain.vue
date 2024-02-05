<template>
    <v-container>
        <v-layout class="mb-2">
            <v-flex :class="this.$vuetify.breakpoint.width < 800 ? 'd-flex flex-column' : 'd-flex'" lg9>
                <v-slide-y-transition mode="out-in" appear>

                    <h2 :class="this.$vuetify.breakpoint.width < 800 ? 'font-weight-medium d-flex align-center justify-start' : 'font-weight-medium d-flex align-center'"
                        lg8>
                        <v-img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/sgblogo.webp"
                            lazy-src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/sgblogo.webp" height="40"
                            width="40" contain class="mr-2"></v-img>
                        <!-- <v-img src="../../assets/Sgblogo.png" height="40" width="40" contain class="mr-2"></v-img> -->
                        <span>Sovereign Gold Bond</span>

                        <!-- removed the order name. -->
                        <!-- <span v-else>SGB Orders <span v-if="historyArr.length == 0">(0)</span></span> -->
                    </h2>
                </v-slide-y-transition>
                <TabBtn @show="show" :InvestCount="InvestCount" :OrderCount="OrderCount" lg4
                    :class="this.$vuetify.breakpoint.width > 800 ? 'ml-5' : 'mb-2'" />
            </v-flex>
        </v-layout>
        <SgbMaster v-if="hidetab" :dynamicText="dynamicText" :result="result" :history="historyArr" :loading="loading"
            @showCircle="showCircle" :Flag="Flag" @passVal="passVal" :tableKey="tableKey" :masterFound="masterFound"
            :historyFound="historyFound" />

        <SgbMobile v-else :dynamicText="dynamicText" :result="result" :history="historyArr" :loading="loading"
            @showCircle="showCircle" :Flag="Flag" @passVal="passVal" :tableKey="tableKey" :masterFound="masterFound"
            :historyFound="historyFound" />

        <PlaceSgb :detail="detail" :dialog="dialog" @closeSgbPop="closePop" :Action="actionFlag" :modify="modify"
            @EmptyModify="emptyModify" @RecallSgb="reCall()" @ChangeActionFlag="ChangeActionFlag" />

        <BottomText class="mt-5" :disclaimer="disclaimer" />

    </v-container>
</template>

<script>
import TabBtn from './Tab/tabBtn.vue';
// import PlaceSgb from './PlaceSgb/placeSgb.vue'
import PlaceSgb from './PlaceSgb/placeSgbNew.vue'
import SgbMaster from './SgbMaster/SgbMaster.vue';
import SgbMobile from './SgbMaster/SgbMobile.vue';
import BottomText from './BottomText.vue';
import EventService from '@/services/EventServices.js'
export default {
    components: {
        TabBtn,
        SgbMaster,
        SgbMobile,
        PlaceSgb,
        BottomText
    },
    data() {
        return {
            result: [],
            historyArr: [],
            loading: false,
            Flag: "", // It helps to make condition on @click:row
            InvestCount: 0,
            OrderCount: 0,
            detail: {},
            modify: {},
            dialog: false,
            actionFlag: "",
            tableKey: 0,
            currentTime: "",
            // btnText: "Place order",
            Day: 0,
            // Modbutton: false,
            circle: false,
            loadingtext: "",
            masterNoDataText: "",
            historyNoDataText: "",
            disclaimer: "",
            masterFound: "",
            historyFound: "",
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
        dynamicText: {
            get() {
                if (this.loading == true) {
                    return this.loadingtext
                } else {
                    if (this.Flag == "I") {
                        return this.masterFound == "N" ? this.masterNoDataText : "Record not found"
                    }
                    else if (this.Flag == "O") {
                        return this.historyFound == "N" ? this.historyNoDataText : "Record not found"
                    }
                    else {
                        return "Records not found"
                    }

                }
            },
            set(text) {
                this.loadingtext = text
            }
        }
    },
    methods: {
        showCircle() {
            this.circle = !this.circle
        },
        show(flag) {
            this.Flag = flag;
        },
        reCall() {
            this.getSgb();
            this.Flag = "I";
        },
        getSgb() {
            this.loading = true;
            this.loadingtext = "loading please wait..."
            EventService.GetSgbMaster()
                .then((response) => {
                    if (response.data.status == "S") {
                        this.disclaimer = response.data.disclaimer == undefined ? '' : response.data.disclaimer

                        this.masterFound = response.data.masterFound
                        this.masterNoDataText = response.data.noDataText == undefined ? '' : response.data.noDataText

                        if (response.data.sgbDetail != null) {
                            if (this.masterFound == "Y") {
                                this.result = response.data.sgbDetail;
                            } else {
                                this.result = [];
                            }
                            this.InvestCount = response.data.investCount == undefined ? 0 : parseInt(response.data.investCount)
                            // this.Flag = "I";
                        } else {
                            this.result = [];
                        }
                    } else {
                        this.MessageBar("E", response.data.errMsg)
                    }
                    this.loading = false;
                })
                .catch((error) => {
                    this.loading = false;
                    this.MessageBar("E", error)
                });
            this.getSgbHistory() // To call and get the history 
        },
        getSgbHistory() {
            EventService.GetSgbHistory()
                .then((response) => {
                    this.loading = false;
                    // console.log(response.data,"this si data")
                    if (response.data.status == "S") {
                        this.historyFound = response.data.historyFound
                        if (this.historyFound == "Y") {
                            this.historyArr = response.data.sgbOrderHistoryArr != null ? response.data.sgbOrderHistoryArr : [];
                        } else {
                            this.historyArr = []
                        }
                        this.historyNoDataText = response.data.historynoDataText == undefined ? '' : response.data.historynoDataText
                        this.OrderCount = response.data.orderCount == undefined ? 0 : parseInt(response.data.orderCount)

                        this.loading = false;
                    }
                })
                .catch((error) => {
                    this.loading = false;
                    this.MessageBar("E", error)
                });
        },

        passVal(item, indicator) {
            this.detail = item;
            this.actionFlag = indicator;
            this.dialog = true;
            if (this.actionFlag != "N") {
                this.getModify(item)
            }
        },
        getModify(item) {
            //     EventService.GetSgbModifyDetail(Id, OrderNo)
            //         .then((response) => {
            //             if (response.data.status == "S") {
            this.modify = item
            //             } else {
            //                 this.MessageBar("E", response.data.errMsg)
            //             }
            //         })
            //         .catch((error) => {
            //             this.MessageBar("E", error)
            //         });
        },
        closePop() {
            this.dialog = false;
            this.detail = {}
        },
        emptyModify() {
            this.modify = {};
        },
        ChangeActionFlag(action) {
            this.actionFlag = action;
        },
        // getCurrentTime() {
        //     const currentTime = new Date();
        //     let hours = currentTime.getHours();
        //     let minutes = currentTime.getMinutes();
        //     let seconds = currentTime.getSeconds();
        //     hours = (hours < 10 ? "0" : "") + hours;
        //     minutes = (minutes < 10 ? "0" : "") + minutes;
        //     seconds = (seconds < 10 ? "0" : "") + seconds;
        //     //this.currentTime = `08:00:00`;

        //     this.currentTime = `${hours}:${minutes}:${seconds}`;
        // },

    },
    created: async function () {
        if (this.$globalData.currentTime == "") {
            await setInterval(this.GetCurrentTime, 1000);
        }
        await this.getSgb();
        this.Flag = "I"
        // To get Day
        this.Day = new Date().getDay();
        // To get current Time
        // this.getCurrentTime();
        // setInterval(this.getCurrentTime, 1000);

    },
    // watch: {
    //     currentTime: function () {
    //         if (this.currentTime >= "10:00:00" && this.currentTime <= "17:00:00") {
    //             this.btnText = 'Place Order'
    //             this.Modbutton = false;
    //             // if (this.Day === 6 || this.Day === 0) {
    //             //     this.btnText = 'OFFLINE';
    //             //     this.ShowMsg = true;
    //             //     this.Modbutton = true;
    //             // }
    //         } else {
    //             this.ShowMsg = true;
    //             this.$emit('BidButton',)
    //             this.btnText = 'OFFLINE';
    //             this.Modbutton = true;
    //         }
    //     },

    // }
}
</script>

<style scoped>
.v-responsive {
    flex: none;
}
</style>