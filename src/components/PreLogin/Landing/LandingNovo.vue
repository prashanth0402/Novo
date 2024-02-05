<template>
    <div>
        <div class="banner_gradiant ma-auto">
            <v-container>
                <v-row class="d-flex align-center justify-center custom_container">
                    <v-col cols="12">
                        <v-row class="page-Bg">
                            <v-col cols="12" xl="7" lg="6" md="6" sm="12" xs="12" justify="center" align="center"
                                v-if="!($vuetify.breakpoint.width < 950)">
                                <v-img lazy-src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Novo Dash Resized.webp"
                                    :src="isWebpSupported ? 'https://flattrade.s3.ap-south-1.amazonaws.com/promo/Novo Dash Resized.webp' : ''"
                                    contain min-width="70%" />
                            </v-col>
                            <v-col class="d-flex justify-center align-center flex-column" cols="12" xl="5" lg="6" md="6"
                                sm="12" xs="12" :class="`mt-${($vuetify.breakpoint.width < 950) ? 10 : 0}`">
                                <v-flex class="d-flex flex-column text-center justify-center align-center" xl12 lg12 md12
                                    sm12 xs12>
                                    <span class="font-weight-bold text-capitalize font-weight-normal mb-2 contentHead--text"
                                        :class="$vuetify.breakpoint.width > 720 ? 'text-h3' : 'text-h5'">primary
                                        market</span>
                                    <span class="secondary--text text-capitalize font-weight-normal mb-2"
                                        :class="$vuetify.breakpoint.width > 720 ? ' text-h5' : 'subtitle-1 font-weight-medium'">
                                        Start your investment journey
                                    </span>
                                    <p class="button content--text" style="font-size: 14px;">
                                        Invest in IPO, zero commission direct mutual funds, sovereign gold bonds, government
                                        bills, and corporate bonds.
                                    </p>
                                    <v-hover v-slot="{ hover }">
                                        <v-btn :color="hover ? 'secondary mt-5' : 'btnColor white--text mt-5'
                                            " width="150" depressed style="border-radius: 10px"
                                            @click="reDirect">Login</v-btn>
                                    </v-hover>
                                </v-flex>
                            </v-col>
                        </v-row>
                    </v-col>
                </v-row>
            </v-container>
        </div>
        <InvestmentChoice />
        <MobileMock />
        <signupMock />
    </div>
</template>

<script>
import InvestmentChoice from "./InvestmentChoice/InvestmentChoice.vue";
import MobileMock from "../DownloadMockUp/MobileMock.vue";
import signupMock from "../SignUp/signupMock.vue";
export default {
    components: {
        InvestmentChoice,
        MobileMock,
        signupMock
    },
    methods: {
        reDirect() {
            this.$emit("GetRedirectUrl");
        },
    },
    computed: {
        isWebpSupported() {
            const elem = document.createElement('canvas');
            return elem.toDataURL('image/webp').indexOf('data:image/webp') === 0;
        }
    }

};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@200;400;500;600;700;800;900&display=swap");

.v-application .text-h1,
.text-h2,
.text-h3,
.text-h4,
.text-h5,
.text-h6,
.v-application p {
    font-family: 'Poppins', sans-serif !important;
}

.banner_gradiant {
    background: linear-gradient(180deg, #879EF5 -91.93%, rgba(173, 187, 241, 0.170062) 10.46%, rgba(135, 158, 245, 0) 62.98%);
}

.custom_container {
    max-width: 80%;
    margin: 0 auto;
}

/* @import "@/path-to-your-styles.css"; */
@media (min-width: 850px) {
    .page-Bg {
        background-size: cover;
        background-repeat: no-repeat;
    }
}

@media only screen and (max-width:950px) {
    .custom_container {
        max-width: 100%;
    }
}

.col-xl,
.col-xl-auto,
.col-xl-12,
.col-xl-11,
.col-xl-10,
.col-xl-9,
.col-xl-8,
.col-xl-7,
.col-xl-6,
.col-xl-5,
.col-xl-4,
.col-xl-3,
.col-xl-2,
.col-xl-1,
.col-lg,
.col-lg-auto,
.col-lg-12,
.col-lg-11,
.col-lg-10,
.col-lg-9,
.col-lg-8,
.col-lg-7,
.col-lg-6,
.col-lg-5,
.col-lg-4,
.col-lg-3,
.col-lg-2,
.col-lg-1,
.col-md,
.col-md-auto,
.col-md-12,
.col-md-11,
.col-md-10,
.col-md-9,
.col-md-8,
.col-md-7,
.col-md-6,
.col-md-5,
.col-md-4,
.col-md-3,
.col-md-2,
.col-md-1,
.col-sm,
.col-sm-auto,
.col-sm-12,
.col-sm-11,
.col-sm-10,
.col-sm-9,
.col-sm-8,
.col-sm-7,
.col-sm-6,
.col-sm-5,
.col-sm-4,
.col-sm-3,
.col-sm-2,
.col-sm-1,
.col,
.col-auto,
.col-12,
.col-11,
.col-10,
.col-9,
.col-8,
.col-7,
.col-6,
.col-5,
.col-4,
.col-3,
.col-2,
.col-1 {
    width: 100%;
    padding: 10px 20px !important;
}
</style>
