<template>
    <div>
        <v-row justify="center" class="custom_container ma-auto my-10 pa-5">
            <v-col>
                <v-row v-for="cont in  content " :key="cont.text" justify="center" class="pa-5">
                    <v-col justify="center" align="start" cols="12">
                        <v-row :class="$vuetify.breakpoint.width < 960 ? 'd-felx flex-column align-center' : undefined">
                            <v-col cols="3" class="d-flex align-center justify-center">
                                <v-img :lazy-src="cont.image" :src="cont.image" contain max-width="100"></v-img>
                            </v-col>
                            <v-col :class="$vuetify.breakpoint.width <= 850 ? 'd-flex flex-column text-center' :
                                'd-flex flex-column text-left'" cols="9">
                                <span class="text-h6 text-capitalize font-weight-bold mb-2 contentHead--text custom-font">{{
                                    cont.text
                                }}</span>
                                <p class="body-1 content--text button custom-font">{{ cont.sentence }}</p>
                            </v-col>
                        </v-row>
                    </v-col>
                </v-row>
            </v-col>
        </v-row>
    </div>
</template>

<script>
export default {
    data() {
        return {
            content: [
                {
                    image: "https://flattrade.s3.ap-south-1.amazonaws.com/promo/earlyinvest2.webp", text: "Early investor", sentence: "Make the first move and become an early shareholder in a company"
                },
                { image: "https://flattrade.s3.ap-south-1.amazonaws.com/promo/listinggains2.webp", text: "Listing gains", sentence: "You may achieve significant gains if the shares get listed at a premium" },
                { image: "https://flattrade.s3.ap-south-1.amazonaws.com/promo/pricetrans2.webp", text: "Price Transparency", sentence: "Access all information related to price and determine investor demand for a stock" },
            ]
        }
    }
}
</script>

<style scoped>
.col {
    width: 100%;
    padding: 5px;
}

.custom_container {
    max-width: 800px;
}

.v-application .text-h6 {
    font-family: "Nunito", sans-serif !important;
}
</style>