<template>
    <div>
        <!-- <p class="text text--secondary">Your order will be placed on the exchange (NSE/BSE) at the end of the subscription
            period.
            Ensure to keep
            sufficient balance in your Trading account on the last day of the issue. Credit from
            <br v-if="this.$vuetify.breakpoint.name == 'xl'">
            stocks sold on the closing day of the issue will not be considered towards the purchase of the SGB.
        </p> -->
        <v-flex xl8>
            <p class="text text--secondary" >{{ disclaimer }}</p>
        </v-flex>
    </div>
</template>

<script>
export default {
    props: {
        disclaimer: String
    }
}
</script>

<style scoped>
.text {
    font-size: .8em;
}
</style>