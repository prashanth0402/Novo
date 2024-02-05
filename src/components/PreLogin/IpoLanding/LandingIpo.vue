<template>
    <div>
        <div class="banner_gradiant">
            <v-container>
                <v-row class="d-flex align-center justify-center custom_container">
                    <v-col cols="12">
                        <v-row class="page-Bg">
                            <v-col cols="12" xl="7" lg="7" md="6" sm="12" xs="12" justify="center" align="center"
                                v-if="!($vuetify.breakpoint.width < 950)">
                                <v-img lazy-src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/IPO dash resized.webp"
                                    src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/IPO dash resized.webp" contain
                                    max-width="80%"></v-img>
                            </v-col>
                            <v-col class="d-flex justify-center align-center flex-column" cols="12" xl="5" lg="5" md="6"
                                sm="12" xs="12" :class="`mt-${($vuetify.breakpoint.width < 950) ? 10 : 0}`">
                                <v-col class="d-flex flex-column text-center justify-center align-center">
                                    <span class="text-capitalize font-weight-bold mb-2 contentHead--text"
                                        :class="$vuetify.breakpoint.width > 720 ? 'text-h3' : 'text-h5'">
                                        Initial Public Offering</span>
                                    <p class="button content--text">
                                        Invest early in quality Indian companies and grow your
                                        wealth steadily</p>
                                    <v-hover v-slot="{ hover }">
                                        <v-btn :color="hover ? 'secondary mt-5' : 'btnColor white--text mt-5'" width="150"
                                            depressed style="border-radius: 10px;" @click="reDirect">
                                            <span>Login</span>
                                        </v-btn>
                                    </v-hover>
                                </v-col>
                            </v-col>
                        </v-row>
                    </v-col>
                </v-row>
            </v-container>
        </div>
        <InvestmentChoice />
        <MobileMock />
        <SignupMock />
    </div>
</template>
<script>

import InvestmentChoice from "./InvestmentChoice/InvestmentChoice.vue"
import MobileMock from "../DownloadMockUp/MobileMock.vue";
import EventService from "../../../services/EventServices";
import SignupMock from "../SignUp/signupMock.vue";
export default {
    components: {
        InvestmentChoice,
        MobileMock,
        SignupMock
    },
    methods: {
        reDirect() {
            EventService.GetRedirectURL()
                .then((response) => {
                    if (response.data.status == "S") {

                        this.$globalData.host = response.data.redirectUrl.host
                        this.$globalData.appName = response.data.redirectUrl.appName
                        this.$globalData.url = response.data.redirectUrl.url

                        this.AssignUrl();

                        // this.redirectUrl = this.redirectUrl + this.$globalData.url
                        // console.log("Redirect: " + this.redirectUrl, "Url: " + this.$globalData.url);

                        window.location.href = this.redirectUrl;

                    }
                })
                .catch((error) => {
                    this.MessageBar("E", error);
                });

        },
    }
};

</script>

<style scoped>
.v-application .text-h1,
.text-h2,
.text-h3,
.text-h4,
.text-h5,
.text-h6,
.v-application p {
    font-family: 'Poppins', sans-serif !important;
}

.v-btn {
    line-height: 3.290rem !important;
    font-family: 'Poppins', sans-serif !important;
}

.banner_gradiant {
    /* background: linear-gradient(115.96deg,
            #a1b4fa -25.22%,
            #cdd7fd 5.05%,
            rgba(229, 235, 254, 0.724631) 36.45%,
            rgba(249, 251, 255, 0.5) 88.4%);*/
    background: linear-gradient(180deg, #879EF5 -91.93%, rgba(173, 187, 241, 0.170062) 10.46%, rgba(135, 158, 245, 0) 62.98%);

}

.custom_container {
    max-width: 100%;
    margin: 0 auto;
}

@media (min-width: 850px) {
    .page-Bg {
        /* background: url("https://flattrade.s3.ap-south-1.amazonaws.com/promo/BG for Landing 1185.webp"); */
        /* background: url("https://flattrade.s3.ap-south-1.amazonaws.com/promo/BG New.png"); */
        background-size: cover;
        background-repeat: no-repeat;
    }
}

/*
.page-Bg {
    background: none;
    background-size: cover;
    background-repeat: no-repeat;
} 

.page-Bg {
    background: url("https://flattrade.s3.ap-south-1.amazonaws.com/promo/Bg New Merge.png");
    background-size: cover;
    aspect-ratio: 16/9;
    background-repeat: no-repeat;
}*/

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
