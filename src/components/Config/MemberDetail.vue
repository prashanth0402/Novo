<template>
    <div class="container">
        <v-layout row wrap class="mb-5" style="padding-left: 12px;">
            <v-flex xs12 lg9 class="mt-6 d-flex justify-left">
                <v-icon left color="blue darken-4" medium>
                    mdi-wallet-membership
                </v-icon>
                <span class="text-subtitle-1">Member credentials</span>
            </v-flex>
        </v-layout>

        <v-card>
            <v-card-title>
                <!-- Order Placement Radio Group -->
                <v-col class="d-flex justify-end">
                    <v-btn text class="primary--text text-caption font-weight-medium"
                        @click="!editable ? editedItem() : editable = false">
                        <div v-if="!editable" class="d-flex align-center">
                            <v-icon small left>mdi-pencil</v-icon>
                            <span>Edit</span>
                        </div>
                        <div v-else class="d-flex align-center">
                            <v-icon small left>mdi-arrow-left</v-icon>
                            <span>Back</span>
                        </div>
                    </v-btn>
                </v-col>
            </v-card-title>
            <v-slide-x-transition mode="out-in" appear>
                <v-card-text>
                    <v-form ref="form" v-model="valid" lazy-validation>
                        <v-row>

                            <!-- <v-col cols="12" sm="12" md="12" class="d-flex"
                                :class="this.$vuetify.breakpoint.name == 'xs' ? 'd-flex flex-column align-start' : 'd-flex align-center'">
                                <span class="text-subtitle-1 font-weight-medium mr-10">Order Preference*</span>
                                <v-radio-group v-model="membercredentials.selectedOrder" mandatory row
                                    class="orderPlacement"
                                    :rules="[v => !!membercredentials.selectedOrder || 'Select an order placement option']">
                                    <v-radio label="NSE" value="NSE" :disabled="!editable"></v-radio>
                                    <v-radio label="BSE" value="BSE" :disabled="!editable"></v-radio>
                                </v-radio-group>
                            </v-col> -->

                            <!-- NSE Member credentials -->
                            <v-col cols="12" sm="12" md="12" xs="6" v-if="detail[0] != undefined">
                                <span class="text-subtitle-1 font-weight-medium">NSE<sup>*</sup></span>
                                <div :class="this.$vuetify.breakpoint.name == 'xs' ? 'd-flex flex-column' : 'd-flex mt-1'">
                                    <v-text-field v-model="detail[0].memberId" label="Member Id*" required outlined
                                        :disabled="!editable" dense
                                        :class="this.$vuetify.breakpoint.name == 'xs' ? 'mr-0' : 'mr-5'"></v-text-field>
                                    <v-text-field v-model="detail[0].login" label="Login*" required outlined
                                        :disabled="!editable" dense
                                        :class="this.$vuetify.breakpoint.name == 'xs' ? 'mr-0' : 'mr-5'"></v-text-field>
                                    <v-text-field v-model="detail[0].password" label="Password*" required outlined
                                        :disabled="!editable" dense></v-text-field>
                                </div>
                            </v-col>

                            <!-- BSE Member credentials -->
                            <v-col cols="12" sm="12" md="12" v-if="detail[1] != undefined">
                                <span class="text-subtitle-1 font-weight-medium">BSE<sup>*</sup></span>
                                <div :class="this.$vuetify.breakpoint.name == 'xs' ? 'd-flex flex-column' : 'd-flex mt-1'">
                                    <v-text-field v-model="detail[1].memberId" label="Member Id*" required outlined
                                        :disabled="!editable" dense
                                        :class="this.$vuetify.breakpoint.name == 'xs' ? 'mr-0' : 'mr-5'"></v-text-field>
                                    <v-text-field v-model="detail[1].login" label="Login*" required outlined
                                        :disabled="!editable" dense
                                        :class="this.$vuetify.breakpoint.name == 'xs' ? 'mr-0' : 'mr-5'"></v-text-field>
                                    <v-text-field v-model="detail[1].password" label="Password*" required outlined
                                        :disabled="!editable" dense
                                        :class="this.$vuetify.breakpoint.name == 'xs' ? 'mr-0' : 'mr-5'"></v-text-field>
                                    <v-text-field v-model="detail[1].ibbsid" label="Ibbsid*" required outlined
                                        :disabled="!editable" dense></v-text-field>
                                </div>
                            </v-col>
                            <!-- Checkbox Group -->
                            <v-col cols="12" sm="12" md="12"
                                :class="this.$vuetify.breakpoint.name == 'xs' ? 'd-flex flex-column' : 'd-flex'">
                                <span class="text-subtitle-1 font-weight-medium mr-10">Allowed Modules<sup>*</sup></span>
                                <div class="d-flex" v-for="module, idx in modules" :key="idx">

                                    <v-checkbox :rules="[v => !!selectedModules.length || 'Select at least one module']"
                                        v-model="selectedModules" :label="module.toUpperCase()" :value="module"
                                        class="mr-5 mt-0" :disabled="!editable" @click="PushSegment"></v-checkbox>
                                </div>
                            </v-col>
                            <v-col>
                                <div v-for="module, idx in selectedModules" :key="idx">
                                    <!-- <div v-for="Segments, idx in updatedArray2" :key="idx"> -->
                                    <div v-for="Segments, idx in SegmentsArr" :key="idx">
                                        <!-- <div v-if="module != Segments.segments ? SegmentsArr.push({
                                            segments: module, nse: false,
                                            nseShare: 0, bse: true, bseShare: 0
                                        })
                                            : undefined"></div> -->
                                        <v-row v-if="module == Segments.segments">

                                            <v-col xl="1" lg="1" sm="1" md="1" xs="1" class="d-flex align-center">
                                                <span class="text-subtitle-1 font-weight-medium">{{
                                                    Segments.segments.toUpperCase()
                                                }}<sup>*</sup></span>
                                            </v-col>
                                            <v-col xl="1" lg="1" sm="2" md="2" xs="2">
                                                <v-checkbox v-model="Segments.nse" label="NSE"
                                                    :disabled="!editable || Segments.bseShare == 100 || nseFilled"></v-checkbox>
                                                <v-slide-y-transition>
                                                    <v-text-field v-model.number="Segments.nseShare" v-if="Segments.nse"
                                                        label="NSE" required outlined
                                                        :disabled="!editable || Segments.bseShare == 100" dense
                                                        @keypress="onlyForNumber" suffix="%"></v-text-field>
                                                </v-slide-y-transition>
                                            </v-col>
                                            <v-col xl="1" lg="1" sm="2" md="2" xs="2">
                                                <v-checkbox v-model="Segments.bse" label="BSE"
                                                    :disabled="!editable || Segments.nseShare == 100 || bseFilled"></v-checkbox>
                                                <v-slide-y-transition>

                                                    <v-text-field v-model.number="Segments.bseShare" v-if="Segments.bse"
                                                        label="BSE" required outlined
                                                        :disabled="!editable || Segments.nseShare == 100" dense
                                                        @keypress="onlyForNumber" suffix="%"></v-text-field>
                                                </v-slide-y-transition>
                                            </v-col>
                                        </v-row>
                                    </div>
                                </div>
                            </v-col>
                        </v-row>
                    </v-form>
                </v-card-text>
            </v-slide-x-transition>
            <v-card-actions class="px-4 pb-4">
                <v-spacer></v-spacer>
                <v-btn depressed :disabled="!valid || !isFormFilled() || !editable || issave" color="blue darken-1"
                    class="white--text px-5" @click="saveMembercredentials()">
                    Save
                </v-btn>
            </v-card-actions>
        </v-card>
        <v-dialog v-model="dialog" persistent max-width="600">
            <v-card>
                <v-card-text class="text-center d-flex align-center justify-space-between pa-5">
                    <span class="primary--text subtitle-1">
                        Kindly Logout and Login again ,Thankyou!
                    </span>
                    <v-btn to="/" text>Logout</v-btn>
                </v-card-text>
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
import EventServices from '../../services/EventServices';
export default {
    data() {
        return {
            dialog: false,
            editable: true,
            valid: true,
            selectedModulescopy:[],
            selectedModules: [], // holding allowes models
            detail: [],
            membercredentials: {
                flag: '',
                selectedModules: '',
                credentials: [],
                segmentDetails: []
            },
            Detailcopy:[],
            isDetailsChanged:true,
            allow: false,
            segmentArrcopy:[],
            SegmentsArr: [], // Holding segments
            ErrorArr: [],
            segmentShares: [],
            //changed  Gsec as Ncb in Array and in Allow Modules Table also chaged Sgb/Gsec/Ipo to Sgb/Ncb/Ipo
            modules: ["Ipo", "Sgb", "Ncb"],
            nseFilled: false,
            bseFilled: false
        };
    },
    methods: {
        PushSegment() {
            // var store = []
            for (let index = 0; index < this.selectedModules.length; index++) {
                // this.seg.segments = this.selectedModules[index]

                // Check if the value property is not present in the secondArray
                const isValueNotPresent = !this.SegmentsArr.some(item => item.segments === this.selectedModules[index]);
                // If the value is not present, push the object into firstArray
                if (isValueNotPresent) {
                    this.SegmentsArr.push({ segments: this.selectedModules[index], nse: false, nseShare: 0, bse: false, bseShare: 0 });
                    console.log( this.SegmentsArr," this.SegmentsArr");
                }
            }
            // this.SegmentsArr = store
        },
        onlyForNumber($event) {
            let keyCode = $event.keyCode ? $event.keyCode : $event.which;

            if (
                (keyCode < 48 || keyCode > 57) &&
                (keyCode !== 46 || keyCode == 46) &&
                (keyCode !== -46 || keyCode == -46)
            ) {
                $event.preventDefault();
            }
        },
        isFormFilled() {
            return (
                this.selectedModules.length > 0
            );
        },
        editedItem() {
            this.editable = true;
            this.$refs.form.resetValidation()
        },
        RemoveSegments() {
            const filteredData = this.SegmentsArr.filter(item => item.nseShare != 0 || item.bseShare !== 0);
            for (let index = 0; index < filteredData.length; index++) {
                for (let id = 0; id < this.selectedModules.length; id++) {
                    if (this.selectedModules[id] == this.SegmentsArr[index].segments) {
                        this.segmentShares.push(this.SegmentsArr[index])
                        // console.log( this.segmentShares)
                    }
                }
            }
        },
        saveMembercredentials() {
            this.editable = false;
            this.RemoveSegments()
            this.membercredentials.segmentDetails = this.segmentShares
            this.membercredentials.selectedModules =
                this.selectedModules.join('/');
            // TO remove the not valid field index
            this.removeIndex(this.detail)
            if (this.allow == true) {
                EventServices.SetMemberDetail(this.membercredentials)
                    .then((response) => {
                        if (response.data.status == "S") {
                            this.MessageBar("S", "MemberDetails saved successfully")
                            // window.location.reload();
                            this.MemberDetail()
                        } else {
                            this.MessageBar("E", response.data.errMsg)
                            // this.MemberDetail()
                        }
                        if (this.membercredentials.flag == "N") {
                            this.dialog = true;
                        } else {
                            this.dialog = false;
                        }

                    })
                    .catch((error) => {
                        this.MessageBar("E", error)
                    })
                this.$refs.form.resetValidation()
            }
        },
        removeIndex(arr) {
            const filteredArr = arr.filter(item => {
                if (item.exchange === "NSE") {
                    return (item.memberId != "" && item.login != "" && item.password != "");
                } else {
                    return (item.memberId != "" && item.login != "" && item.password != "" && item.ibbsid != "");
                }
            });
            if ((arr[0].memberId == "" || arr[0].login == "" || arr[0].password == "") &&
                (arr[1].memberId == "" || arr[1].login == "" || arr[1].password == ""
                    || arr[1].ibbsid == "")) {
                this.MessageBar('E', "Atleast Fill either One NSE or BSE properly")
                this.allow = false;
            } else {
                this.allow = true;
            }
            this.membercredentials.credentials = filteredArr
        },
        MemberDetail() {
            this.editable = false;
            EventServices.GetMemberDetail()
                .then((response) => {
                    // console.log(response.data);
                    if (response.data.status == "S") {
                        if (response.data.segmentDetails != null) {

                            this.SegmentsArr = response.data.segmentDetails
                            this.segmentArrcopy = JSON.stringify(response.data.segmentDetails)
                        } else {
                            this.SegmentsArr = []
                        }
                        // console.log(this.SegmentsArr.length,"response");
                        // this.setDynamicArray()
                        const TempArr = response.data.memberDetail.credentials
                        if (TempArr != null) {
                            for (let i = 0; i < TempArr.length; i++) {
                                if (TempArr.length == 1) {
                                    if (TempArr[i].exchange == "NSE") {
                                        TempArr.push({
                                            memberId: '',
                                            login: '',
                                            password: '',
                                            ibbsid: '',
                                            exchange: 'BSE',
                                        })
                                    } else {
                                        TempArr.push({
                                            memberId: '',
                                            login: '',
                                            password: '',
                                            ibbsid: '',
                                            exchange: 'NSE',
                                        })
                                    }
                                    // TODO: sort by NSE
                                    TempArr.sort((a) => {
                                        if (a.exchange === "NSE") {
                                            return -1; // "NSE" comes before "BSE"
                                        } else {
                                            return 1; // "BSE" comes after "NSE"
                                        }
                                    });
                                    this.detail = TempArr
                                }

                            }
                            this.detail = response.data.memberDetail.credentials
                            this.Detailcopy = JSON.stringify( response.data.memberDetail.credentials)
                            this.membercredentials = response.data.memberDetail
                            // this.membercredentialsCopy = JSON.stringify(response.data.memberDetail)
                            this.selectedModules = response.data.memberDetail.selectedModules.split('/');
                            this.selectedModulescopy = JSON.stringify(response.data.memberDetail.selectedModules.split('/'))
                            if (this.SegmentsArr == null) {
                                this.PushSegment()
                            }
                        } else {
                            this.membercredentials.flag = "N"
                            this.editable = true
                            this.detail.push(
                                {
                                    memberId: '',
                                    login: '',
                                    password: '',
                                    ibbsid: '',
                                    exchange: 'NSE',
                                },
                                {
                                    memberId: '',
                                    login: '',
                                    password: '',
                                    ibbsid: '',
                                    exchange: 'BSE',
                                }
                            )
                        }
                    } else {
                        this.MessageBar("E", response.data.errMsg);
                    }
                })
                .catch((error) => {
                    this.MessageBar("E", error)
                })
        },
        pushValues() {
            for (const value of this.selectedModules) {
                if (!this.SegmentsArr.includes(value)) {
                    this.SegmentsArr.push(value);
                }
            }
        },
    },
    computed:{
        issave(){
            if (JSON.stringify(this.detail) != this.Detailcopy || (JSON.stringify(this.selectedModules) != this.selectedModulescopy) || (this.segmentArrcopy != JSON.stringify(this.SegmentsArr))){
                    return false
                }else if (JSON.stringify(this.detail) == this.Detailcopy || (JSON.stringify(this.selectedModules) == this.selectedModulescopy)|| (this.segmentArrcopy == JSON.stringify(this.SegmentsArr))){
                     return true
                }else{
                    return true
                }
        }
    },
    watch: {
        SegmentsArr: {
            handler(arr) {

                for (var i = 0; i < arr.length; i++) {
                    if (arr[i].bse != false && arr[i].nse != false) {

                        if (arr[i].nseShare + arr[i].bseShare > 101) {
                            this.MessageBar("E", "Segment percentage exceeded 100")
                        }
                        // else if (arr[i].nseShare + arr[i].bseShare <= 100) {
                        if (arr[i].bseShare == 100) {
                            arr[i].nseShare = 0
                        } else if (arr[i].nseShare == 100) {
                            arr[i].bseShare = 0
                        }
                        if (arr[i].nseShare + arr[i].bseShare != 100) {
                            this.MessageBar("E", arr[i].segments + " segment count below 100")
                        }
                        //  else {
                            //     if (arr[i].bseShare != 0 || arr[i].bseShare != "") {
                        //         arr[i].nseShare = 100 - arr[i].bseShare
                        //     } else if (arr[i].nseShare != 0 || arr[i].nseShare != "") {
                        //         arr[i].bseShare = 100 - arr[i].nseShare
                        //     }
                        // }
                    } else {
                        if (arr[i].bseShare == 100) {
                            arr[i].nseShare = 0
                        } else if (arr[i].nseShare == 100) {
                            // console.log("arr[i].nseShare == 100");
                            arr[i].bseShare = 0
                        } else {
                            if (arr[i].bseShare != 0 || arr[i].bseShare != "") {
                                arr[i].nseShare = 100 - arr[i].bseShare
                            } else if (arr[i].nseShare != 0 || arr[i].nseShare != "") {
                                arr[i].bseShare = 100 - arr[i].nseShare
                            }
                        }
                    }
                }

            }, deep: true
        },
        detail: {
            handler() {
                for (var i = 0; i < this.SegmentsArr.length; i++) {
                    if (this.detail[0].memberId != "" && this.detail[0].login != "", this.detail[0].password != "" && this.detail[0].exchange != "") {
                        this.nseFilled = false;
                    } else {
                        this.SegmentsArr[i].nse = false
                        this.nseFilled = true;
                    }

                    if (this.detail[1].memberId != "" && this.detail[1].login != "", this.detail[1].password != "" && this.detail[1].exchange != "" &&
                        this.detail[1].ibbsid != "") {
                        this.bseFilled = false;
                    } else {
                        this.SegmentsArr[i].bse = false
                        this.bseFilled = true;
                    }
                }
            }, deep: true
        }
    },
    mounted() {
        this.MemberDetail()
    }
};
</script>

<style scoped>
@media only screen and (max-width: 600px) {
    .orderPlacement .v-input--radio-group__input {
        flex-direction: column;
    }
}

::v-deep.v-card__subtitle,
.v-card__title {
    padding: 0px !important;
}
</style>