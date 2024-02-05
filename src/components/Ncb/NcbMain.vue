<template>
    <v-container>

        <v-layout class="mb-2">
            <v-flex :class="this.$vuetify.breakpoint.width < 800 ? 'd-flex flex-column' : 'd-flex'" lg9>
                <v-slide-y-transition mode="out-in" appear>
                    <h2 :class="this.$vuetify.breakpoint.width < 800 ? 'font-weight-medium d-flex align-center justify-start' : 'font-weight-medium d-flex align-center'"
                        lg8>

                        <span>
                            NCB
                        </span>
                    </h2>
                </v-slide-y-transition>
                <TabBtn @show="show" :InvestCount="InvestCount" :OrderCount="OrderCount" lg4
                    :class="this.$vuetify.breakpoint.width > 800 ? 'ml-5' : 'mb-2'" />
            </v-flex>

        </v-layout>

        <!-- Table -->

        <GoiDatedBonds v-if="hidetab" :goiresult="goiresult" :goihistory="goihistoryArr" :loading="loading"
            :goidynamicText="goidynamicText" @showCircle="showCircle" :Flag="Flag" @passVal="passVal" :tableKey="tableKey"
            :goimasterFound="goimasterFound" :goihistoryFound="goihistoryFound" />


        <GoiMobile v-else :goiresult="goiresult" :goihistory="goihistoryArr" :loading="loading"
            :goidynamicText="goidynamicText" @showCircle="showCircle" :Flag="Flag" @passVal="passVal" :tableKey="tableKey"
            :goimasterFound="goimasterFound" :goihistoryFound="goihistoryFound" />

        <Tbilltable v-if="hidetab" :billsresult="billsresult" :billshistory="billshistoryArr"
            :billdynamicText="billdynamicText" :loading="loading" @showCircle="showCircle" :Flag="Flag" @passVal="passVal"
            :tableKey="tableKey" :tbillmasterFound="tbillmasterFound" :tbillhistoryFound="tbillhistoryFound" />

        <Tbillmobile v-else :billsresult="billsresult" :billshistory="billshistoryArr" :loading="loading"
            :billdynamicText="billdynamicText" @showCircle="showCircle" :Flag="Flag" @passVal="passVal" :tableKey="tableKey"
            :tbillmasterFound="tbillmasterFound" :tbillhistoryFound="tbillhistoryFound" />

        <SdlTable v-if="hidetab" :loansresult="loansresult" :loanshistory="loanshistoryArr" :loading="loading"
            :sdldynamicText="sdldynamicText" @showCircle="showCircle" :Flag="Flag" @passVal="passVal" :tableKey="tableKey"
            :sdlmasterFound="sdlmasterFound" :sdlhistoryFound="sdlhistoryFound" />


        <SdlMobile v-else :loansresult="loansresult" :loanshistory="loanshistoryArr" :loading="loading"
            :sdldynamicText="sdldynamicText" @showCircle="showCircle" :Flag="Flag" @passVal="passVal"
            :tableKey="tableKey" :sdlmasterFound="sdlmasterFound" :sdlhistoryFound="sdlhistoryFound"/>


        <Placeorder :detail="detail" :copymodify="copymodify" :minLot="minLot" :dialog="dialog" @closeNcbPop="closePop"
            @RecallNcb="reCall()" :modify="modify" :Action="actionFlag" @EmptyModify="emptyModify"
            @ChangeActionFlag="ChangeActionFlag" :iconVal="iconVal" />

        <BottomText class="mt-5" :disclaimer="disclaimer" />

    </v-container>
</template>


<script>
import TabBtn from './Tab/tabBtn.vue';
import GoiDatedBonds from './NcbInvert/GoiMaster/GoiTable.vue';
import GoiMobile from './NcbInvert/GoiMaster/GoiMobile.vue';
import Tbilltable from './NcbInvert/TBillMaster/TbillTable.vue';
import Tbillmobile from './NcbInvert/TBillMaster/TbillMobile.vue';
import SdlTable from './NcbInvert/SDLMaster/SdlTable.vue';
import SdlMobile from './NcbInvert/SDLMaster/SdlMobile.vue';
import BottomText from './BottomText.vue';
import Placeorder from './NcbPlaceOrder/placeNcb.vue';
import EventService from '@/services/EventServices.js';

export default {
    name: "NCB",
    components: {
        TabBtn,
        GoiDatedBonds,
        GoiMobile,
        Tbilltable,
        Tbillmobile,
        SdlTable,
        SdlMobile,
        Placeorder,
        BottomText
    },
    data() {
        return {
            copymodify: {},
            goiresult: [],
            goihistoryArr: [],
            billsresult: [],
            billshistoryArr: [],
            loansresult: [],
            loanshistoryArr: [],
            loading: false,
            Flag: "",
            tableKey: 0,
            InvestCount: 0,
            OrderCount: 0,
            detail: {},
            copydetail: {},
            minLot: 0,
            modify: {},
            dialog: false,
            actionFlag: "",
            currentTime: "",
            circle: false,
            iconVal: "",
            loadingtext: "",

            //masterDatatext
            goinoDataText: "",
            tbillnoDataText: "",
            sdlnoDataText: "",

            //Historydta
            goihistorynoDataText: "",
            tbillhistorynoDataText: "",
            sdlhistorynoDataText: "",

            //masterfound
            goimasterFound: "",
            tbillmasterFound: "",
            sdlmasterFound: "",

            //historyfound
            goihistoryFound: "",
            tbillhistoryFound: "",
            sdlhistoryFound: "",


            disclaimer: "",
            Day: 0,
        }
    },
    methods: {


        openDialog() {
            this.dialog = true;
        },

        closePop() {
            this.dialog = false;
            this.detail = {}
        },
        // loading
        showCircle() {
            this.circle = !this.circle
        },
        // Show
        show(flag) {
            this.Flag = flag;
        },
        // flag
        reCall() {
            this.getNcb();
        },
        // table
        getNcb() {
            this.loading = true;
            this.loadingtext = "loading please wait..."
            EventService.GetNcbMaster()
                .then((response) => {
                    if (response.data.status == "S") {

                        this.disclaimer = response.data.disclaimer == undefined ? '' : response.data.disclaimer

                        this.goimasterFound = response.data.goimasterFound
                        this.goinoDataText = response.data.goinoDataText == undefined ? '' : response.data.goinoDataText

                        this.tbillmasterFound = response.data.tbillmasterFound
                        this.tbillnoDataText = response.data.tbillnoDataText == undefined ? '' : response.data.tbillnoDataText

                        this.sdlmasterFound = response.data.sdlmasterFound
                        this.sdlnoDataText = response.data.sdlnoDataText == undefined ? '' : response.data.sdlnoDataText

                        this.InvestCount = response.data.investCount == undefined ? 0 : parseInt(response.data.investCount)

                        //G-sec
                        if (response.data.gSecDetail != null) {

                            if (this.goimasterFound == "Y") {
                                this.goiresult = response.data.gSecDetail;
                            } else {
                                this.goiresult = [];
                            }

                            this.goiresult.forEach(item => {
                                if (item.name) {
                                    item.name = item.name.replace(/[^a-zA-Z0-9.%\s]/g, '');
                                    item.name = item.name.replace('MISSING', '');
                                }
                            });

                        } else {
                            this.goinoDataText = response.data.goinoDataText
                            this.goiresult = [];

                        }

                        //T-bill
                        if (response.data.tBillDetail != null) {


                            if (this.tbillmasterFound == "Y") {
                                this.billsresult = response.data.tBillDetail;
                            } else {
                                this.billsresult = [];
                            }

                            this.billsresult.forEach(item => {
                                if (item.name) {
                                    item.name = item.name.replace(/[^a-zA-Z0-9.%\s]/g, '');
                                    item.name = item.name.replace('MISSING', '');
                                }
                            });

                        } else {
                            this.tbillnoDataText = response.data.tbillnoDataText
                            this.billsresult = [];

                        }

                        //SDL
                        if (response.data.sdlDetail != null) {

                            if (this.sdlmasterFound == "Y") {
                                this.loansresult = response.data.sdlDetail;
                            } else {
                                this.loansresult = [];
                            }

                            this.loansresult.forEach(item => {
                                if (item.name) {
                                    item.name = item.name.replace(/[^a-zA-Z0-9.%\s]/g, '');
                                    item.name = item.name.replace('MISSING', '');
                                }
                            });
                        } else {
                            this.sdlnoDataText = response.data.sdlnoDataText
                            this.loansresult = [];

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
            this.getNcbOrderHistory()
        },
        getNcbOrderHistory() {
            EventService.getNcbOrderHistory()
                .then((response) => {
                    this.loading = false;
                    this.OrderCount = response.data.orderCount

                    //G-sec
                    if (response.data.status == "S") {

                        this.goihistoryFound = response.data.goihistoryFound
                        if (this.goihistoryFound == "Y") {
                            this.goihistoryArr = response.data.gSecOrderHistoryArr != null ? response.data.gSecOrderHistoryArr : [];

                            this.goihistoryArr.forEach(item => {
                                if (item.name) {
                                    item.name = item.name.replace(/[^a-zA-Z0-9.%\s]/g, '');
                                    item.name = item.name.replace('MISSING', '');
                                }
                            });

                        } else {
                            this.goihistoryArr = []
                        }

                        this.tbillhistoryFound = response.data.tbillhistoryFound
                        if (this.tbillhistoryFound == "Y") {
                            this.billshistoryArr = response.data.tBillOrderHistoryArr != null ? response.data.tBillOrderHistoryArr : [];

                            this.billshistoryArr.forEach(item => {
                                if (item.name) {
                                    item.name = item.name.replace(/[^a-zA-Z0-9.%\s]/g, '');
                                    item.name = item.name.replace('MISSING', '');
                                }
                            });

                        } else {
                            this.billshistoryArr = []
                        }

                        this.sdlhistoryFound = response.data.sdlhistoryFound
                        if (this.sdlhistoryFound == "Y") {
                            this.loanshistoryArr = response.data.sdlOrderHistoryArr != null ? response.data.sdlOrderHistoryArr : [];

                            this.loanshistoryArr.forEach(item => {
                                if (item.name) {
                                    item.name = item.name.replace(/[^a-zA-Z0-9.%\s]/g, '');
                                    item.name = item.name.replace('MISSING', '');
                                }
                            });

                        } else {
                            this.loanshistoryArr = []
                        }
                        

                        this.goihistorynoDataText = response.data.goihistorynoDataText == undefined ? '' : response.data.goihistorynoDataText
                        this.OrderCount = response.data.orderCount == undefined ? 0 : parseInt(response.data.orderCount)

                        this.tbillhistorynoDataText = response.data.tbillhistorynoDataText == undefined ? '' : response.data.tbillhistorynoDataText
                        this.OrderCount = response.data.orderCount == undefined ? 0 : parseInt(response.data.orderCount)

                        this.sdlhistorynoDataText = response.data.sdlhistorynoDataText == undefined ? '' : response.data.sdlhistorynoDataText
                        this.OrderCount = response.data.orderCount == undefined ? 0 : parseInt(response.data.orderCount)

                        this.loading = false;
                    }
                    
                })
                .catch((error) => {
                    this.loading = false;
                    this.MessageBar("E", error)
                });

            // this.refreshTable()
        },
                        //     if (response.data.gSecOrderHistoryArr != null) {
                        //         this.goihistoryArr = response.data.gSecOrderHistoryArr;
                        //         this.goihistorynoDataText = response.data.goihistorynoDataText
                        //         this.loading = false;
                        //         this.goihistoryArr.forEach(item => {
                        //             if (item.name) {
                        //                 item.name = item.name.replace(/[^a-zA-Z0-9.%\s]/g, '');
                        //                 item.name = item.name.replace('MISSING', '');
                        //             }
                        //         });
                        //     } else {
                        //         this.goihistoryArr = [];
                        //         this.goihistorynoDataText = response.data.goihistorynoDataText
                        //     }

                        //     //T-Bill
                        //     if (response.data.tBillOrderHistoryArr != null) {
                        //         this.billshistoryArr = response.data.tBillOrderHistoryArr;
                        //         this.tbillhistorynoDataText = response.data.tbillhistorynoDataText
                        //         this.loading = false;
                        //         this.billshistoryArr.forEach(item => {
                        //             if (item.name) {
                        //                 item.name = item.name.replace(/[^a-zA-Z0-9.%\s]/g, '');
                        //                 item.name = item.name.replace('MISSING', '');
                        //             }
                        //         });
                        //     } else {
                        //         this.billshistoryArr = [];
                        //         this.tbillhistorynoDataText = response.data.tbillhistorynoDataText
                        //     }

                        //     //SDL
                        //     if (response.data.sdlOrderHistoryArr != null) {
                        //         this.loanshistoryArr = response.data.sdlOrderHistoryArr;
                        //         this.sdlhistorynoDataText = response.data.sdlhistorynoDataText
                        //         this.loading = false;
                        //         this.loanshistoryArr.forEach(item => {
                        //             if (item.name) {
                        //                 item.name = item.name.replace(/[^a-zA-Z0-9.%\s]/g, '');
                        //                 item.name = item.name.replace('MISSING', '');
                        //             }
                        //         });
                        //     } else {
                        //         this.loanshistoryArr = [];
                        //         this.sdlhistorynoDataText = response.data.sdlhistorynoDataText
                        //     }
                        // } else {
                        //     this.MessageBar('E', response.data.errMsg)
                    // }
        // refreshTable() {
        //     this.tableKey += 1;
        //     setTimeout(() => {
        //         this.CalcInvestAndOrder();
        //     }, 200)

        // },



        passVal(item, indicator) {

            this.detail = { ...item };
            // this.copydetail = { ...item }
            // this.minLot = item.minLot
            this.actionFlag = indicator;
            this.dialog = true;
            if (this.actionFlag != "N") {
                this.getModify(item)
            }

            if (this.detail.series == "GS") {
                this.iconVal = "https://flattrade.s3.ap-south-1.amazonaws.com/promo/gseclogo.png"
            } else if (this.detail.series == "TB") {
                this.iconVal = "https://flattrade.s3.ap-south-1.amazonaws.com/promo/tresur.png"
            } else {
                this.iconVal = "https://flattrade.s3.ap-south-1.amazonaws.com/promo/SdlLogo.png"
            }
        },

        getModify(item) {
            this.modify = item
        },
        ChangeActionFlag(action) {
            this.actionFlag = action;
        },
        getCurrentTime() {
            const currentTime = new Date();
            let hours = currentTime.getHours();
            let minutes = currentTime.getMinutes();
            let seconds = currentTime.getSeconds();
            hours = (hours < 10 ? "0" : "") + hours;
            minutes = (minutes < 10 ? "0" : "") + minutes;
            seconds = (seconds < 10 ? "0" : "") + seconds;
            this.currentTime = `${hours}:${minutes}:${seconds}`;
        },
        emptyModify() {
            // console.log(this.modify.lotSize)
            this.modify = {};
        },
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
        goidynamicText: {
            get() {
                if (this.loading == true) {
                    return this.loadingtext
                } else {
                    if (this.Flag == "I" && this.goiresult.length == 0) {
                        return this.goimasterFound == "N" ? this.goinoDataText : "Record not found"
                    }
                    else if (this.Flag == "O" && this.goihistoryArr.length == 0) {
                        return this.goihistoryFound == "N" ? this.goihistorynoDataText : "Record not found"
                    }
                    else {
                        return "Records not found"
                    }
                }
            },
            set(text) {
                this.loadingtext = text
            }
        },
        billdynamicText: {
            get() {
                if (this.loading == true) {
                    return this.loadingtext
                } else {
                    if (this.Flag == "I" && this.billsresult.length == 0) {
                        return this.tbillmasterFound == "N" ? this.tbillnoDataText : "Record not found"
                    }
                    else if (this.Flag == "O" && this.billshistoryArr.length == 0) {
                        return this.tbillhistoryFound == "N" ? this.tbillhistorynoDataText : "Record not found"
                    }
                    else {
                        return "Records not found"
                    }
                }
            },
            set(text) {
                this.loadingtext = text
            }
        },
        sdldynamicText: {
            get() {
                if (this.loading == true) {
                    return this.loadingtext
                } else {
                    if (this.Flag == "I" && this.loansresult.length == 0) {
                        return this.sdlmasterFound == "N" ? this.sdlnoDataText : "Record not found"
                    }
                    else if (this.Flag == "O" && this.loanshistoryArr.length == 0) {
                        return this.sdlhistoryFound == "N" ? this.sdlhistorynoDataText : "Record not found"
                    }
                    else {
                        return "Records not found"
                    }
                }
            },
            set(text) {
                this.loadingtext = text
            }
        },
    },
    // created: async function () {
    //     if (this.$globalData.currentTime == "") {
    //         await setInterval(this.GetCurrentTime, 900);
    //     }
    //     await this.getNcb();
    //     this.Flag = "I"
    //     // To get Day
    //     this.Day = new Date().getDay();
    // },
    created: async function () {
        await this.getNcb();
        this.Flag = "I"
        this.Day = new Date().getDay();
        this.getCurrentTime();
        setInterval(this.getCurrentTime, 500);

    },
}
</script>

<style scoped>
.v-responsive {
    flex: none;
}
</style>