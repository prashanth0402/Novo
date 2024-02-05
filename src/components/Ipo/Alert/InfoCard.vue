<template>
    <v-container>
        <v-alert color="blue lighten-4" dark border="left" >
            <span class="black--text text font-weight-light">
                IPO window will remain open from <b> 10 AM till 5 PM </b> on trading days. <br>
                Please use your Virtual Payment Address (UPI ID) for placing bids. <br>
                The UPI collect request from the bank might be delayed.Accept it to apply whenever you receive the collect
                request on your UPI app (might take up to end of the day).
            </span>
        </v-alert>
    </v-container>
</template>

<script>
export default {
    name: "InfoCard",
}
</script>
<style scoped>
.text {
    font-size: 12px;
    display: inline-block;
}
</style>