<template>
    <div>
        <!-- <v-layout class="mt-2 d-none d-sm-flex">
            <v-flex class="d-flex justify-start">
                <v-icon @click="close()">mdi-close</v-icon>
            </v-flex>
        </v-layout> -->

        <v-layout class="mb-4 d-flex justify-space-between mt-2">
            <v-flex lg8>
                <v-layout class="d-flex flex-column align-start">
                    <v-flex class="subtitle-1 font-weight-bold">{{ HistoryRec.symbol }}</v-flex>
                    <v-flex class="text ">{{ HistoryRec.name }}</v-flex>
                </v-layout>
            </v-flex>
            <v-flex lg2 class="d-flex justify-end">
                <v-chip v-show="HistoryRec.sme == true" x-small color="blue lighten-5" label
                    class="smechip--text">SME</v-chip>
            </v-flex>
        </v-layout>
        <v-row>
            <v-col class=" hidden-sm-and-down">

                <v-divider></v-divider>
                <v-layout class="ma-2 ">
                    <v-flex>Application No</v-flex>
                    <v-flex class="d-flex justify-end font-weight-black">{{ HistoryRec.appNo }}</v-flex>
                </v-layout>

                <v-divider></v-divider>

                <v-layout class="ma-2 ">
                    <v-flex>Issue Date</v-flex>
                    <v-flex class="d-flex justify-end font-weight-black">{{ HistoryRec.issueDate }}</v-flex>
                </v-layout>

                <v-divider></v-divider>

                <v-layout class="ma-2 ">
                    <v-flex><span>No. of shares</span></v-flex>
                    <v-flex class="d-flex justify-end font-weight-black">{{ formatIssuesize(HistoryRec.issueSize)
                    }}</v-flex>
                </v-layout>

                <v-divider></v-divider>

                <v-layout class="ma-2 ">
                    <v-flex>Issue price </v-flex>
                    <v-flex class="d-flex justify-end font-weight-black">{{ HistoryRec.issuePrice }}</v-flex>
                </v-layout>

                <v-divider></v-divider>

                <v-layout class="ma-2 ">
                    <v-flex>Lot size</v-flex>
                    <v-flex class="d-flex justify-end font-weight-black">{{ HistoryRec.lotSize }}</v-flex>
                </v-layout>
                <v-divider></v-divider>

                <v-layout class="ma-2 ">
                    <v-flex>Discount</v-flex>
                    <v-flex class="d-flex justify-end font-weight-black">{{ HistoryRec.discount }}</v-flex>
                </v-layout>
                <v-layout class="ma-2 mt-5">
                    <v-flex><a :href="HistoryRec.registrarLink" target="_blank" v-if="Allotmenturl"
                            style="text-decoration: none;">Check your Allotment
                            <v-icon small color="blue lighten-1">mdi-information-outline</v-icon></a></v-flex>
                </v-layout>

                <!-- <v-divider></v-divider> -->
                <v-layout class="ma-2" v-if="HistoryRec.errReason != ''">
                    <!-- <v-flex>
                        <h4>Failed Reason:</h4>
                    </v-flex> -->
                    <v-flex class="d-flex justify-start font-weight-black">
                        <span class="error--text">
                            {{ HistoryRec.errReason }}
                        </span>
                    </v-flex>
                </v-layout>
            </v-col>
        </v-row>
    </div>
</template>
  
<script>
export default {
    name: "issueDetails",

    props: {
        HistoryRec: {}
    },

    methods: {
        close() {
            this.$emit("closeHistoryRec");
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
        }
    },
    computed: {
        Allotmenturl() {
            return this.HistoryRec.registrarLink.startsWith('http', 0) || this.HistoryRec.registrarLink.startsWith('https', 0);
        }
    }
};
</script>

<style scoped>
.v-card__subtitle,
.v-card__text {
    padding: 5px !important;
}

.text {
    font-size: 10px;
}

.col {
    width: 100%;
    padding: 0px 16px !important;
}
</style>
  