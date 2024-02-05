<template>
    <div>
        <v-img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Coming Soon Desk copy W.webp"
            lazy-src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Coming Soon Desk copy W.webp" contain
            v-if="this.$vuetify.breakpoint.width > 800"></v-img>

        <v-img v-else src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Coming Soon Mob copy W.webp"
            lazy-src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Coming Soon Mob copy W.webp" contain></v-img>
    </div>
</template>

<script>
export default {

}
</script>

<style  scoped></style>
