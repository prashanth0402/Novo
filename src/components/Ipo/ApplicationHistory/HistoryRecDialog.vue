<template>
    <div>
        <v-dialog v-model="Dialog" max-width="1000px" persistent :fullscreen="this.$vuetify.breakpoint.width <= 650"
            :transition="this.$vuetify.breakpoint.width <= 800 ? 'dialog-bottom-transition' : undefined"  overlay-color="#fff">
            <v-card>
                <v-card-text :style="this.$vuetify.breakpoint.width <= 650 ? 'padding:  2px 10px;' : undefined">
                    <v-row class="pt-5">
                        <v-col cols="12" lg="5" md="6" sm="12" xs="12" class="mb-lg-10 mb-md-10 ">
                            <IpoDetail :HistoryRec="Struc" @closeHistoryRec="closeHistoryRec" />
                        </v-col>
                        <v-col cols="12" lg="7" md="6" sm="12" xs="12">
                            <ModifyInput :HistoryRec="Struc" @closeHistoryRec="closeHistoryRec" />
                        </v-col>
                        <Steppers :Item="Item" v-if="check" v-show="Flag == 'O'" />
                    </v-row>
                </v-card-text>
            </v-card>
        </v-dialog>
    </div>
</template>
<script>
import IpoDetail from './HistoryRecDialogComp/HistoryRecLeft/IpoDetail.vue';
import ModifyInput from './HistoryRecDialogComp/HistoryRecRight/HistoryRecInput.vue';
import Steppers from "../IpoWindow/IpoWindowComp/Steppers.vue"
export default {
    components: {
        IpoDetail,
        ModifyInput,
        Steppers,

    },
    props: {
        HistoryRec: {},
        ReportRec: {},
        ShowHistoryRec: Boolean,
        ShowReportRec: Boolean,
        Item: {},
        Flag: String

    },
    methods: {
        closeHistoryRec() {
            this.$emit('closeHistoryRec');
        }
    },
    computed: {
        Dialog: {
            get() {
                if (this.ShowHistoryRec == false) {
                    return this.ShowReportRec
                } else {
                    return this.ShowHistoryRec
                }
            }
        },
        Struc: {
            get() {
                if (this.HistoryRec != null) {
                    return this.HistoryRec
                } else {
                    return this.ReportRec
                }
            }
        },
        check: {
            get() {
                if (this.Item != undefined) {
                    return true
                } else {
                    return false
                }
            }
        }

    }
}
</script>
<style scoped>
::v-deep .v-sheet.v-card:not(.v-sheet--outlined) {
    box-shadow: none !important;
}

::v-deep .v-dialog {
    box-shadow: 5px 5px 12px #d7d7d7,
        -5px -5px 12px #e9e9e9 !important;
}
</style>