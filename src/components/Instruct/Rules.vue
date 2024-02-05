<template>
    <v-container>
        <v-card class="pa-2 text elevation-0 light-blue lighten-5" outlined>
            <ul class="text-xs-subtitle-1 mt-2">
                <li>UPID is compulsory for applying online IPO.</li>
                <li>You can subscribe to IPO by placing live bids through our portal between 10 AM and 5 PM on trading
                    days
                    and until the IPO subscription is open.</li>
                <li>You can pre-apply to IPO by placing offline Bids placed between 5 PM to 10 AM and a mandate will be
                    generated next day.</li>
                <li>Once you place the IPO bids, you will receive a mandate request to your corresponding UPI app.</li>
                <li>The bidding will happen under the retail category when the IPO application amount is up to Rs 2
                    lakhs.
                </li>
                <li>The application amount is between Rs 2 lakhs to Rs 5 lakh, the bidding will happen under the HNI
                    category. HNI bidders are requested to submit their application before 3 PM on the last day of the
                    IPO.
                </li>
                <li>Please note the amount will be blocked but not deducted from your bank account.</li>
                <li>If there is no allotment of shares in your demat account, then the amount will be unblocked.</li>
            </ul>
        </v-card>
    </v-container>
</template>

<script>
export default {
    data() {
        return {
            dialog: false,
            data: false
        }
    }
}

</script>

<style scoped>
a {
    text-decoration: none;
    line-height: 1;
}

a:hover {
    text-decoration: underline;
}

.v-dialog__content {
    align-items: start;
}

.text {
    font-size: 12px;
}
</style>