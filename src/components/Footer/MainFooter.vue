<template>
    <div>
        <Bankdetaildialog :dialog="dialog" :indicator="indicator" @hideBankDetails="hideBankDetails" />

        <Footer1 />
        <v-divider></v-divider>
        <Footer2 @showBankDetails="showBankDetails" />
        <v-divider></v-divider>
        <Footer3 />
    </div>
</template>

<script>
import Footer1 from './FooterComp/Footer1.vue';
import Footer2 from './FooterComp/Footer2.vue';
import Footer3 from './FooterComp/Footer3.vue';
import Bankdetaildialog from './FooterComp/bankdetaildialog.vue';
export default {
    data: () => ({
        dialog: false,
        indicator: ''
    }),
    components: {
        Footer1,
        Footer2,
        Footer3,
        Bankdetaildialog
    },
    methods: {
        showBankDetails(val) {
            this.dialog = true
            this.indicator = val
        },
        hideBankDetails() {
            this.dialog = false

        }
    }
}
</script>

<style  scoped></style>