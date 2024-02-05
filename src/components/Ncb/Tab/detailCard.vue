<template>
    <div>
        <v-card v-click-outside="newShow" width="350" outlined>
            <v-card-title style="background-color:#BBDEFB ;">
                <v-layout class="d-flex align-center">
                    <!-- <img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/sgblogo.webp" alt="" height="22px"
                        width="22px"> -->
                    <v-flex class="d-flex justify-space-between">
                        <span class="caption">&nbsp;{{ master.name }}</span>
                        <!-- <v-icon @click="closeDetail" class="error--text mb-1">mdi-close</v-icon> -->
                    </v-flex>
                </v-layout>

            </v-card-title>
            <v-card-text>
                <v-row>
                    <v-col>
                        <v-layout class="text--disabled">
                            <v-flex class="d-flex flex-column">
                                <span>Allotment Status</span>
                                <span>Form Verification Status</span>
                            </v-flex>
                            <v-flex class="d-flex flex-column text-right">
                                <span>{{ master.rbiStatus == '' ? '-' : master.rbiStatus }}</span>
                                <span>{{ master.dpStatus == '' ? '-' : master.dpStatus }}</span>
                            </v-flex>
                        </v-layout>
                        <v-divider class="my-2"></v-divider>
                        <v-row>
                            <v-col class="text-left d-flex flex-column">
                                <span>Requested Unit</span>
                                <span>Requested Unit Price</span>
                                <span>Amount Payable</span>
                            </v-col>
                            <v-col class="text-right d-flex flex-column">
                                <span>{{ master.requestedUnit }}</span>
                                <span>{{ master.requestedUnitPrice }}</span>
                                <span>&#x20b9; {{ master.requestedAmount }}</span>
                            </v-col>
                        </v-row>
                        <v-divider class="my-2"></v-divider>
                        <v-row>
                            <v-col class="text-left d-flex flex-column">
                                <span>Applied Unit</span>
                                <span>Applied Unit Price</span>
                                <span> Amount Payable</span>
                            </v-col>
                            <v-col class="text-right d-flex flex-column">
                                <span>{{ master.appliedUnit }}</span>
                                <span>{{ master.appliedUnitPrice }}</span>
                                <span>&#x20b9; {{ master.appliedAmount }}</span>
                            </v-col>
                        </v-row>
                        <v-divider class="my-2"></v-divider>
                        <v-row>
                            <v-col class="text-left d-flex flex-column">
                                <span>Alloted Unit</span>
                                <span>Alloted Unit Price</span>
                                <span>Amount Payable</span>
                            </v-col>
                            <v-col class="text-right d-flex flex-column">
                                <span>{{ master.allotedUnit }}</span>
                                <span>{{ master.allotedUnitPrice }}</span>
                                <span>&#x20b9; {{ master.allotedAmount }}</span>
                            </v-col>
                        </v-row>
                    </v-col>
                </v-row>
            </v-card-text>
        </v-card>
    </div>
</template>

<script>
export default {
    data() {
        return {
            show: "",
        }
    },
    methods: {
        formatedPrice(item) {
            if (item != undefined) {
                return item.toLocaleString('en-IN')
            }
        },
        closeDetail() {
            this.$emit('closeDetail')
        },
    },
    props: {
        master: Object,
        showDetail: Boolean,
    },
    computed: {
        newShow: {
            get() {
                if (this.show != '') {
                    return false
                } else {
                    return this.showDetail
                }
            },
            set(value) {
                this.closeDetail()
                this.show = value
            }
        },
    },
}
</script>

<style  scoped>
.v-responsive {
    flex: none;
}
</style>