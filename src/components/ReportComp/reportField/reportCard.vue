<template>
    <div class="d-flex justify-center">
        <v-card class="mb-5 elevation-1" :loading="loading">
            <template slot="progress">
                <v-progress-linear color="primary" indeterminate></v-progress-linear>
            </template>
            <v-card-title class="mb-1" dense solo>
                <h3 class="primary--text">Search</h3>
            </v-card-title>
            <v-card-text>
                <v-slide-y-transition mode="out-in" appear>
                    <v-container>
                        <v-form ref="form" lazy-validation>
                            <v-row>
                                <v-col cols="12" sm="6" xs="6" md="4">
                                    <v-autocomplete v-model="input.module" background-color="white" dense outlined
                                        label="Segment" :items="Item" :rules="Rules" item-text="text"
                                        item-value="value"></v-autocomplete>
                                </v-col>
                                <v-col cols="12" sm="6" xs="6" md="4">
                                    <v-text-field label="Symbol" outlined v-model="input.symbol" dense
                                        @keydown.enter="proccess" background-color="white"></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="6" xs="6" md="4">
                                    <v-text-field label="Client Id" outlined v-model.number="input.clientId" dense
                                        @keydown.enter="proccess" background-color="white"></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="6" xs="6" md="4">
                                    <v-menu v-model="menu1" :close-on-content-click="false" min-width="auto">
                                        <template v-slot:activator="{ on, attrs }" style="left:40px;top: 343px;">
                                            <v-text-field v-model="input.fromDate" label="From" readonly v-bind="attrs"
                                                v-on="on" dense outlined background-color="white"></v-text-field>
                                        </template>
                                        <v-date-picker v-model="input.fromDate" @input="menu1 = false" :max="maxDate"
                                            color="primary" no-title></v-date-picker>
                                    </v-menu>
                                </v-col>
                                <v-col cols="12" sm="6" xs="6" md="4" xl="4">
                                    <v-menu v-model="menu2" :close-on-content-click="false" :nudge-right="40"
                                        min-width="auto">
                                        <template v-slot:activator="{ on, attrs }" style="left:40px;top: 343px;">
                                            <v-text-field v-model="input.toDate" label="To" readonly v-bind="attrs"
                                                v-on="on" dense outlined background-color="white"></v-text-field>
                                        </template>
                                        <v-date-picker v-model="input.toDate" @input="menu2 = false" :max="maxDate"
                                            color="primary" no-title></v-date-picker>
                                    </v-menu>
                                </v-col>


                            </v-row>
                        </v-form>

                    </v-container>
                </v-slide-y-transition>
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn @click="proccess" class="info lighten-1 elevation-0" :disabled="!isButtonDisabled" width="150">
                    <v-icon class="mr-2" size="20">mdi-magnify</v-icon>search</v-btn>
            </v-card-actions>
        </v-card>
    </div>
</template>

<script>
import EventServices from '@/services/EventServices';
export default {
    data() {
        return {
            valid: true,
            Item: [
                { text: "IPO", value: "Ipo" },
                { text: "SGB", value: "Sgb" },
                { text: "Gsec", value: "G-sec" },
                { text: "TBill", value: "TBill" },
                { text: "SDL", value: "SDL" }
            ],
            isMobile: false,
            Rules: [
                v => !!v || 'required',
            ],
            input: {
                module: "",
                symbol: "",
                clientId: "",
                fromDate: "",
                toDate: "",
                category: "",
            },
            ipoArr: [],
            sgbArr: [],
            gsecArr: [],
            tbillArr: [],
            sdlArr: [],
            ShowReportRec: false,
            ReportRec: {},
            screenWidth: window.innerWidth,
            menu1: false,
            menu2: false,
            loading: false,
        }
    },

    methods: {
        // to resize the screen
        handleResize() {
            this.screenWidth = window.innerWidth;
        },
        proccess() {
            if (this.$refs.form.validate()) {
                this.loading = true;
                this.$globalData.overlay = true
                EventServices.GetReport(this.input)
                    .then((response) => {
                        this.loading = false;

                        if (response.data.status == "S") {
                            this.MessageBar('S', "Record Fetch successfully")

                            this.ipoArr = response.data.ipoArr == null ? [] : response.data.ipoArr
                            this.sgbArr = response.data.sgbArr == null ? [] : response.data.sgbArr

                            this.gsecArr = response.data.gsecArr == null ? [] : response.data.gsecArr
                            // console.log("this.gsecArr",  this.gsecArr);

                            this.tbillArr = response.data.tbillArr == null ? [] : response.data.tbillArr
                            // console.log("this.tbillArr",  this.tbillArr);

                            this.sdlArr = response.data.sdlArr == null ? [] : response.data.sdlArr
                            // console.log("this.sdlArr",  this.sdlArr);
                            // if (this.ipoArr == null) {
                            //     this.ipoArr = []
                            // }
                            // if (this.sgbArr == null) {
                            //     this.sgbArr = []
                            // }

                            if (this.input.module == "Ipo") {
                                this.$emit("IpoReport", this.ipoArr, this.input.module)
                            } else if (this.input.module == "Sgb") {
                                this.$emit("SgbReport", this.sgbArr, this.input.module)
                            }else if (this.input.module == "G-sec") {
                                this.$emit("GsecReport", this.gsecArr, this.input.module)
                            }else if (this.input.module == "TBill") {
                                this.$emit("TbillReport", this.tbillArr, this.input.module)
                            }else if (this.input.module == "SDL") {
                                this.$emit("SdlReport", this.sdlArr, this.input.module)
                            }
                            this.$globalData.overlay = false

                            this.emptyObject();
                            this.$refs.form.resetValidation()

                        } else {
                            this.$globalData.overlay = false
                            this.MessageBar('E', response.data.errMsg)
                        }
                    })
                    .catch((error) => {
                        this.loading = false;
                        this.MessageBar('E', error)
                        this.$globalData.overlay = false

                    });
            }
        },
        emptyObject() {
            this.input.clientId = "";
            this.input.module = "";
            this.input.symbol = "";
            this.input.fromDate = "";
            this.input.toDate = "";
        },
        // async manualFetch() {
        //     this.$globalData.overlay = true;
        //     await EventServices.FetchManual()
        //         .then((response) => {
        //             this.$globalData.overlay = false;
        //             if (response.data.status == 'S') {
        //                 this.MessageBar('S', "Data Fetched successfully")
        //             } else {
        //                 this.MessageBar('E', response.data.errMsg)
        //             }
        //         })
        //         .catch((error) => {
        //             this.$globalData.overlay = false;
        //             this.MessageBar('E', error)
        //         })
        // },
    },
    computed: {
        maxDate() {
            const today = new Date().toISOString().split("T")[0];
            return today;
        },
        isButtonDisabled() {
            return this.input.module != "" && (this.input.fromDate != "" && this.input.toDate != "") || this.input.symbol != "" || this.input.clientId != "";
        }
    },
}

</script>

<style scoped>
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
    padding: 0px 5px !important;
}
</style>