<template>
    <v-container>
        <v-slide-y-transition mode="out-in" appear>
            <v-breadcrumbs :items="items">
                <template v-slot:divider>
                    <v-icon>mdi-chevron-right</v-icon>
                </template>
            </v-breadcrumbs>
        </v-slide-y-transition>

        <ReportCard @IpoReport="IpoReport" @SgbReport="SgbReport" @GsecReport="GsecReport" @TbillReport="TbillReport"
            @SdlReport="SdlReport" />
        <v-slide-x-transition mode="out-in" appear>
            <v-layout class="mb-2 d-flex align-center justify-center">

                <v-flex :class="this.$vuetify.breakpoint.width <= 800 ? 'd-flex flex-column d-flex' : 'd-flex align-center'"
                    lg9>
                    <v-flex :class="this.$vuetify.breakpoint.width <= 800 ? 'order-last font-weight-normal d-flex align-center justify-space-between' :
                        'd-flex align-center justify-space-between'">
                        <h2 lg8 class="font-weight-medium">
                            <v-icon>mdi-file-document-box</v-icon>
                            <span>Report</span>
                        </h2>
                        <div class="text-right d-flex d-sm-none">
                            <v-autocomplete label="Category" outlined dense
                                v-if="this.$vuetify.breakpoint.width <= 800 && activeTab == '0'"
                                class="small-autocomplete    text-right d-flex d-sm-none " v-model="category"
                                :items="categoryArr" item-value="value" item-text="text"
                                background-color="white"></v-autocomplete>
                            <v-autocomplete label="Status" outlined dense v-if="this.$vuetify.breakpoint.width <= 800"
                                class="small-autocomplete    text-right d-flex d-sm-none " v-model="status"
                                :items="statusArr" item-value="value" item-text="text"
                                background-color="white"></v-autocomplete>

                            <v-btn @click="downloadCSV" small class="primary--text text-right d-flex d-sm-none" text
                                :disabled="currentItems.length == 0" v-if="this.$vuetify.breakpoint.width <= 800">
                                <span class="caption text-capitalize">Download
                                    <v-icon class="ml-2" x-small>mdi-cloud-download</v-icon>
                                </span>
                            </v-btn>
                        </div>
                    </v-flex>
                    <v-tabs v-model="activeTab" :class="$vuetify.breakpoint.width > 800 ? 'ml-5' : 'mb-2'"
                        active-class="black--text">
                        <v-tab v-for="tab, idx in tabs" :key="idx">
                            <span>{{ tab.name }}</span>
                        </v-tab>
                    </v-tabs>
                </v-flex>
                <v-flex class="justify-end align-center d-none d-sm-flex">
                    <v-autocomplete label="Category" outlined dense
                        v-if="this.$vuetify.breakpoint.width > 800 && activeTab == '0'" class="small-autocomplete "
                        :items="categoryArr" v-model="category" item-value="value" item-text="text"
                        background-color="white"></v-autocomplete>
                    <v-autocomplete label="Status" outlined dense v-if="this.$vuetify.breakpoint.width > 800"
                        class="small-autocomplete text-right ml-2" v-model="status" :items="statusArr" item-value="value"
                        item-text="text" background-color="white"></v-autocomplete>
                    <v-divider v-if="activeTab == 0" vertical class="pa-2"></v-divider>
                    <v-btn @click="downloadCSV" small class="primary--text text-right" text
                        :disabled="currentItems.length == 0" v-if="this.$vuetify.breakpoint.width > 800">
                        <span class="caption text-capitalize">Download
                            <v-icon class="ml-2" x-small>mdi-cloud-download</v-icon>
                        </span>
                    </v-btn>
                </v-flex>
            </v-layout>
        </v-slide-x-transition>
        <ReportDesk v-if="hidetab" :loading="loading" :headers="headers" :records="FilteredItem" :Choice="Choice"
            @dialog="openPop" @Ncbdialog="openNcbPop" />
        <ReportMobile v-else :loading="loading" :records="FilteredItem" :Choice="Choice" @dialog="openPop"
            @Ncbdialog="openNcbPop" />

        <PlaceSgb :dialog="dialog" @closeSgbPop="closePop" :Action="actionFlag" :detail="detail" />

        <placeNcb :dialog="Ncbdialog" @closeNcbPop="closeNcbPop" :Action="actionFlag" :detail="Ncbdetail" :iconVal="iconVal" />
    </v-container>
</template>

<script>
import ReportCard from "./reportField/reportCard.vue"
import ReportDesk from './Visibility/report.vue';
import ReportMobile from './Visibility/reportMobile.vue';
import EventServices from "@/services/EventServices";
// import PlaceSgb from "../Sgb/PlaceSgb/placeSgb.vue"
import PlaceSgb from "../Sgb/PlaceSgb/placeSgbNew.vue"
import placeNcb from "../Ncb/NcbPlaceOrder/placeNcb.vue";
import Papa from "papaparse";
export default {

    data() {
        return {
            category: "All",
            status: "All",
            categoryArr: [],
            statusArr: [
                { text: "All", value: 'All' }, { text: "Success", value: 'success' }, { text: "Failed", value: 'failed' },
                { text: "Pending", value: 'pending' },
            ],
            // id:"",
            Choice: "",
            downloadBtn: false,
            tabs: [
                { idx: 0, name: 'Ipo' },
                { idx: 1, name: 'Sgb' },
                { idx: 2, name: 'G-sec' },
                { idx: 3, name: 'TBill' },
                { idx: 4, name: 'SDL' },
            ],
            moreTabs: [
            ],
            // TodayIpoRecords: [],
            // TodaySgbRecords: [],
            records: [],
            ipoArr: [],
            sgbArr: [],
            gsecArr: [],
            tbillArr: [],
            sdlArr: [],
            iconVal: "",
            screenWidth: window.innerWidth,
            loading: false,
            activeTab: "",
            counter: {
                ipo: 0,
                sgb: 0,
                gsec: 0,
                tbill: 0,
                sdl: 0
            },
            dialog: false,
            Ncbdialog: false,
            actionFlag: "",
            modify: {},
            detail: {},
            Ncbdetail: {},
            items: [
                {
                    text: 'Activities',
                    disabled: true,
                    href: 'breadcrumbs_dashboard',
                },
                {
                    text: 'Report',
                    disabled: false,
                    href: '/report',
                },
            ],

            headeripo: [
                { text: 'Symbol', align: 'left', value: 'symbol', sortable: false },
                { text: 'Exchange', value: 'exchange', align: "center", sortable: true },
                { text: 'Application no.', value: 'applicationNo', align: "center", sortable: false },
                { text: 'Apply Date', value: 'applyDate', align: "center", sortable: true },
                { text: 'Applied Time', value: 'appliedTime', align: "center", sortable: false },
                { text: 'ClientId', value: 'clientId', align: "center", sortable: false },
                { text: 'Category', value: 'category', align: "center", sortable: false },
                { text: 'Status', value: 'status', align: "center", sortable: false },

            ],
            headersgb: [
                { text: 'Symbol', align: 'left', value: 'symbol', sortable: false },
                { text: 'Exchange', value: 'exchange', align: "center", sortable: true },
                { text: 'Int.RefNo', value: 'orderNo', align: "center", sortable: false },
                { text: 'ExchRefno.', value: 'exchOrderNo', align: "center", sortable: false },
                { text: 'Apply Date', value: 'orderDate', align: "center", sortable: true },
                { text: 'ClientId', value: 'clientId', align: "center", sortable: false },
                { text: 'Status', value: 'orderStatus', align: "center", sortable: false },

            ],
            headergsec: [
                { text: 'Symbol', align: 'left', value: 'symbol', sortable: false },
                { text: 'Exchange', value: 'exchange', align: "center", sortable: true },
                { text: 'Int.RefNo', value: 'orderNo', align: "center", sortable: false },
                { text: 'ExchRefno.', value: 'respOrderNo', align: "center", sortable: false },
                { text: 'Apply Date', value: 'orderDate', align: "center", sortable: true },
                { text: 'ClientId', value: 'clientId', align: "center", sortable: false },
                { text: 'Status', value: 'orderStatus', align: "center", sortable: false },

            ],
            headertbill: [
                { text: 'Symbol', align: 'left', value: 'symbol', sortable: false },
                { text: 'Exchange', value: 'exchange', align: "center", sortable: true },
                { text: 'Int.RefNo', value: 'orderNo', align: "center", sortable: false },
                { text: 'ExchRefno.', value: 'respOrderNo', align: "center", sortable: false },
                { text: 'Apply Date', value: 'orderDate', align: "center", sortable: true },
                { text: 'ClientId', value: 'clientId', align: "center", sortable: false },
                { text: 'Status', value: 'orderStatus', align: "center", sortable: false },

            ],
            headersdl: [
                { text: 'Symbol', align: 'left', value: 'symbol', sortable: false },
                { text: 'Exchange', value: 'exchange', align: "center", sortable: true },
                { text: 'Int.RefNo', value: 'orderNo', align: "center", sortable: false },
                { text: 'ExchRefno.', value: 'respOrderNo', align: "center", sortable: false },
                { text: 'Apply Date', value: 'orderDate', align: "center", sortable: true },
                { text: 'ClientId', value: 'clientId', align: "center", sortable: false },
                { text: 'Status', value: 'orderStatus', align: "center", sortable: false },

            ],
        }
    },
    components: {
        ReportCard,
        ReportDesk,
        ReportMobile,
        PlaceSgb,
        placeNcb
    },
    computed: {
        hidetab: {
            get() {
                if (this.$vuetify.breakpoint.name == 'xs') {
                    return false
                } else {
                    return true
                }
            },
        },
        isSmallScreen() {
            return this.screenWidth <= 800;
        },

        currentItems() {
            if (this.activeTab == 0) {
                if (this.category == '' || this.category == undefined || this.category == "All") {
                    return this.ipoArr == null ? [] : this.ipoArr;
                }
                else {
                    return this.ipoArr.filter(ipo => ipo.category === this.category)
                }
            } else if (this.activeTab == 1) {
                return this.sgbArr == null ? [] : this.sgbArr


            } else if (this.activeTab == 2) {
                return this.gsecArr == null ? [] : this.gsecArr

            } else if (this.activeTab == 3) {
                return this.tbillArr == null ? [] : this.tbillArr

            } else if (this.activeTab == 4) {
                return this.sdlArr == null ? [] : this.sdlArr

            } else {

                return []
            }

            // return this.activeTab == 0 ? this.ipoArr == null ? [] : this.ipoArr : this.sgbArr == null ? [] : this.sgbArr;

        },
        FilteredItem() {
            if (this.status == "All") {
                return this.currentItems
            } else {
                return this.currentItems.filter(item => this.activeTab == "0" ? item.status === this.status : item.orderStatus.toLowerCase() === this.status)
            }
        },
        headers() {
            switch (this.activeTab) {
                case 0: return this.headeripo
                case 1: return this.headersgb
                case 2: return this.headergsec
                case 3: return this.headertbill
                default: return this.headersdl
            }

            // if (this.activeTab == 0) {
            //     return this.headeripo
            // } else if (this.activeTab == 1) {
            //     return this.headersgb
            // } else if (this.activeTab == 2) {
            //     return this.headergsec
            // } else if (this.activeTab == 3) {
            //     return this.headertbill
            // } else {
            //     return this.headersdl
            // }
        },
    },
    methods: {
        IpoReport(value, ipo) {
            this.category = "All"
            this.setActiveTabFromApiResponse(ipo, value)
            this.ipoArr = value;
        },
        SgbReport(value, sgb) {
            this.setActiveTabFromApiResponse(sgb, value)
            this.sgbArr = value
        },
        GsecReport(value, gsec) {
            this.setActiveTabFromApiResponse(gsec, value)
            this.gsecArr = value

        },
        TbillReport(value, tbill) {
            this.setActiveTabFromApiResponse(tbill, value)
            this.tbillArr = value

        },
        SdlReport(value, sdl) {
            this.setActiveTabFromApiResponse(sdl, value)
            this.sdlArr = value

        },
        Default() {
            this.loading = true;
            EventServices.DefaultReport()
                .then((response) => {
                    if (response.data.status == "S") {
                        this.loading = false;
                        this.ipoArr = response.data.ipoArr;
                        this.sgbArr = response.data.sgbArr;
                        this.gsecArr = response.data.gsecArr;
                        this.tbillArr = response.data.tbillArr;
                        this.sdlArr = response.data.sdlArr

                        if (this.ipoArr == null) {
                            this.ipoArr = []
                        } else if (this.sgbArr == null) {
                            this.sgbArr = []
                        } else if (this.gsecArr == null) {
                            this.gsecArr = []
                        } else if (this.tbillArr == null) {
                            this.tbillArr = []
                        } else if (this.sdlArr == null) {
                            this.sdlArr = []
                        }




                        // this.setActiveTabFromApiResponse("Ipo", this.TodayIpoRecords);
                        this.Count(this.ipoArr, 'Ipo');
                        this.Count(this.sgbArr, 'Sgb');
                        this.Count(this.gsecArr, 'G-sec');
                        this.Count(this.tbillArr, 'TBill');
                        this.Count(this.sdlArr, 'SDL');
                        this.activeTab = 0
                    }
                })
                .catch(() => {
                    this.loading = false;
                    // this.MessageBar('E', error)
                });
        },
        Count(arr, tab) {
            this.counter.ipo = 0;
            this.counter.sgb = 0;
            this.counter.gsec = 0;
            this.counter.tbill = 0;
            this.counter.sdl = 0;

            if (arr.length > 0) {
                if (tab == 'Ipo') {
                    this.counter.ipo = arr.length
                } else if (tab == 'Sgb') {
                    this.counter.sgb = arr.length
                } else if (tab == 'G-sec') {
                    this.counter.gsec = arr.length
                } else if (tab == 'TBill') {
                    this.counter.tbill = arr.length
                } else if (tab == 'SDL') {
                    this.counter.sdl = arr.length
                }
            }
        },
        downloadCSV() {
            if (this.currentItems.length != 0) {
                const csv = Papa.unparse(this.currentItems);
                const blob = new Blob([csv], { type: "text/csv;charset=utf-8;" });
                const url = URL.createObjectURL(blob);
                const link = document.createElement("a");
                link.setAttribute("href", url);
                link.setAttribute("download", "data.csv");
                link.style.visibility = "hidden";
                document.body.appendChild(link);
                link.click();
                document.body.removeChild(link);
            } else {
                this.MessageBar('E', 'Records unavailabe')
            }
        },
        setActiveTabFromApiResponse(response, array) {
            this.records = []
            // Find the index of the tab based on the API response
            const tabIndex = this.tabs.findIndex(tab => tab.name.toLowerCase() === response.toLowerCase());
            this.Choice = response; // Used to deside the Dialog visibility
            if (tabIndex !== -1) {
                this.activeTab = tabIndex;
            }
            if (array.length > 0) {
                this.Count(array, response); // to count the length of the array
                this.records = array; // To assign the respective array base on the response
            }
        },
        openPop(item) {
            this.actionFlag = "R";
            this.dialog = true;
            this.detail = {
                name: item.name,
                unitPrice: item.requestedUnitPrice,
                appliedUnit: item.requestedUnit,
                startDateWithTime: item.startDateWithTime,
                endDateWithTime: item.endDateWithTime,
                isin: item.isin,
                discountText: item.discountText,
                discountAmt: item.discountAmt,
                orderNo: item.orderNo
            }
        },
        openNcbPop(item) {
            this.actionFlag = "R";
            this.Ncbdialog = true;
            this.Ncbdetail = {
                symbol: item.symbol,
                series: item.series,
                name: item.name,
                unitPrice: item.requestedUnitPrice,
                appliedUnit: item.requestedUnit,
                amount: item.requestedAmount,
                startDateWithTime: item.startDateWithTime,
                endDateWithTime: item.endDateWithTime,
                isin: item.isin,
                discountText: item.discountText,
                discountAmt: item.discountAmt,
                orderNo: item.orderNo
            }
            // this.detail = item
            if (this.Ncbdetail.series == "GS") {
                this.iconVal = "https://flattrade.s3.ap-south-1.amazonaws.com/promo/gseclogo.png"
            } else if (this.Ncbdetail.series == "TB") {
                this.iconVal = "https://flattrade.s3.ap-south-1.amazonaws.com/promo/tresur.png"
            } else {
                this.iconVal = "https://flattrade.s3.ap-south-1.amazonaws.com/promo/SdlLogo.png"
            }

        },
        closePop() {
            this.dialog = false;
        },
        closeNcbPop(){
            this.Ncbdialog = false;
        },
        getCategory(id, path) {
            this.$globalData.overlay = true;
            this.categoryArr = []
            EventServices.GetCategory(id, path)
                .then((response) => {
                    this.$globalData.overlay = false;
                    if (response.data.status == "S") {
                        this.categoryArr = response.data.categoryArr
                        this.categoryArr.unshift({ text: "All", value: "All" })

                    } else {
                        this.MessageBar("E", response.data.errMsg)
                    }
                })
                .catch((error) => {
                    this.$globalData.overlay = false;
                    this.MessageBar("E", error)
                });
        },
    },
    watch: {
        activeTab: function (active) {
            if (this.statusArr.length > 4) {
                this.statusArr.pop()
            }
            if (active == 0) {
                this.Choice = "Ipo";
                this.statusArr.push({ text: "Cancelled", value: 'user cancelled' })
            } else if (active == 1) {
                this.Choice = "Sgb";
                this.statusArr.push({ text: "Cancelled", value: 'bond cancelled' })
            } else if (active == 2) {
                this.Choice = "G-sec";
                this.statusArr.push({ text: "Cancelled", value: 'bond cancelled' })
            } else if (active == 3) {
                this.Choice = "TBill";
                this.statusArr.push({ text: "Cancelled", value: 'bond cancelled' })
            } else if (active == 4) {
                this.Choice = "SDL";
                this.statusArr.push({ text: "Cancelled", value: 'bond cancelled' })
            }
        }
    },

    mounted() {
        this.Default();
        this.getCategory(0, this.$route.path)
    }
}
</script>

<style scoped>
.small-autocomplete {
    max-width: 150px;
    font-size: 12px;
    height: 35px;
}

.small-autocomplete /deep/ label {
    font-size: 12px;

}
</style>