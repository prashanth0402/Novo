<template>
    <div>
        <v-dialog v-model="ShowModify" max-width="1000px" persistent :fullscreen="isFullscreen"
            :transition="isFullscreen ? 'dialog-bottom-transition' : undefined" overlay-color="#fff">
            <v-overlay v-if="overlay">
                <v-progress-circular :size="50" color="white" indeterminate></v-progress-circular>
            </v-overlay>
            <v-card>
                <v-card-text :style="this.$vuetify.breakpoint.width <= 650 ? 'padding:  2px 10px;' : undefined">
                    <v-row class="pt-5">
                        <v-col cols="12" lg="5" md="6" sm="5" xs="12">
                            <IpoDetail :IpoDetail="DetailStruct" :modifyData="ModifyDetail" @closeModify="closeModify"
                                :closeIcon="closeIcon" :discountStruct="discountStruct" />
                        </v-col>
                        <v-col cols="12" lg="7" md="6" sm="7" xs="12">
                            <ModifyInput :modifyData="ModifyDetail" :issueDetails="DetailStruct" :showAddBtn="showAddBtn"
                                :modifyBtn="modifyBtn" @closeModifyBtn="closeModifyBtn" @showUpdate="showUpdate"
                                @hideUpdate="hideUpdate" @addNew="addNew" @modified="modified" :ipoAppTotal="ipoAppTotal"
                                @closeSlot="closeSlot" @cancelBid="cancelBid" @Recall="Recall" @closeModify="closeModify"
                                :copiedModifyData="copiedModifyData" @DgOverlay="DgOverlay" @DgOverlay1="DgOverlay1"
                                :categoryArr="categoryArr" @discountStruct="discountMethod"
                                :amountPayable="amountPayable" />
                        </v-col>
                    </v-row>
                </v-card-text>
            </v-card>
        </v-dialog>
    </div>
</template>
<script>
import IpoDetail from './ModifyDialogComp/ModifyLeft/IpoDetail.vue';
import ModifyInput from './ModifyDialogComp/ModifyRight/ModifyInput.vue';
export default {
    components: {
        IpoDetail,
        ModifyInput,
    },
    methods: {
        closeModify() {
            this.$emit("CloseModify", this.closeMod)
            this.closeIcon = false;
        },
        addNew() {
            this.$emit('addNewBid')
        },
        modified(idx) {
            this.$emit('modified', idx)
        },
        closeSlot(idx, signal) {
            this.$emit('closeSlot', idx, signal)
        },
        cancelBid() {
            this.$emit('cancelBid')
        },
        Recall() {
            this.$emit('Recall')
        },
        // Pending() {
        //     this.$emit('Pending')
        // },
        showUpdate() {
            this.$emit('showUpdate')
        },
        closeModifyBtn() {
            this.$emit('closeModifyBtn')
        },
        hideUpdate(bool) {
            this.$emit('hideUpdate', bool)
        },
        DgOverlay() {
            this.closeIcon = true
        },
        DgOverlay1() {
            this.closeIcon = false
        },
        discountMethod(struct) {
            this.discountStruct = struct
        }

    },
    data() {
        return {
            closeMod: false,
            overlay: false,
            closeIcon: false,
            // category Arr not displaying on next component so created an variable and send  it has prop
            Category: this.categoryArr,
            discountStruct: {}
        }
    },
    props: {
        categoryArr: Array,
        ShowModify: Boolean,
        showAddBtn: Boolean,
        modifyBtn: Boolean,
        DetailStruct: {},
        ModifyDetail: {},
        ipoAppTotal: Number,
        copiedModifyData: {},
        amountPayable: Number,
    },
    computed: {
        isFullscreen() {
            return this.$vuetify.breakpoint.width < 700;
        },
    },

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