<template>
    <div>
        <v-dialog v-model="ShowBid" max-width="1000" persistent :fullscreen="isFullscreen"
            :transition="isFullscreen ? 'dialog-bottom-transition' : undefined" height="100%" overlay-color="#fff">
            <v-card>
                <v-card-text :style="this.$vuetify.breakpoint.width <= 650 ? 'padding:  2px 10px;' : undefined">
                    <v-row class="pt-5">
                        <v-col cols="12" lg="5" md="6" sm="5" xs="12">
                            <IpoDetail :IpoDetail="DetailStruct" :hideLot="hideLot" :discountStruct="discountStruct" />
                        </v-col>
                        <v-col cols="12" lg="7" md="6" sm="7" xs="12">
                            <BidInput :ApplyDetails="DetailStruct" @closeDialog="closeBid" @closeDg="closeOnly"
                                @Recall="recallMethod" :categoryArr="categoryArr" :selectedCategory="selectedCategory"
                                @discountStruct="discountMethod" />
                        </v-col>
                    </v-row>
                </v-card-text>
            </v-card>
        </v-dialog>
    </div>
</template>
<script>
import IpoDetail from './BidDialogComp/IpoDetail.vue';
import BidInput from './BidDialogComp/BidInput.vue';
export default {
    components: {
        IpoDetail,
        BidInput,
    },
    data() {
        return {
            CloseBid: false,
            button: false,
            hideLot: true,
            discountStruct: {}
        }
    },
    props: {
        ShowBid: Boolean,
        DetailStruct: {},
        categoryArr: Array,
        selectedCategory: String
    },
    methods: {
        closeBid() {
            this.$emit("CloseBid", this.CloseBid);
        },
        closeOnly() {
            this.$emit("CloseBidDg", this.CloseBid);
        },
        recallMethod() {
            this.$emit("RecallApi");
        },
        discountMethod(struct) {
            this.discountStruct = struct
        }
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