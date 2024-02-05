<template>
    <div class="d-flex justify-center">
        <v-container max-width="1200">
            <v-row justify="center" class="my-10 pa-5 ma-auto custom_container card">
                <v-col class="d-flex justify-center align-center flex-column text-center" cols="12" xl="11" lg="11" md="12"
                    sm="12" xs="12">

                    <span class="d-none d-sm-flex text-h5  font-weight-bold mb-2 contentHead--text">
                        Achieve your financial goals by investing in diversified financial instruments
                    </span>
                    <span class="d-flex d-sm-none text-h6  font-weight-bold mb-2 contentHead--text">
                        Achieve your financial goals by investing in diversified financial instruments
                    </span>
                    <span :class="this.$vuetify.breakpoint.name == 'xs' ? 'text-h6 text--secondary' :
                        'text-h5 text--secondary'">Multiple invest choices</span>
                </v-col>
            </v-row>
            <v-row justify="center" class="custom_container ma-auto my-16 pa-5">
                <v-col>
                    <v-row v-for="cont in  content " :key="cont.text" justify="center" class="pa-5">
                        <v-col justify="center" align="start" cols="12">
                            <v-row :class="$vuetify.breakpoint.width < 960 ? 'd-felx flex-column align-center' : undefined">
                                <v-col cols="3" class="d-flex align-center justify-center">
                                    <v-img :lazy-src="cont.image" :src="cont.image" contain max-width="100"></v-img>
                                </v-col>
                                <v-col :class="$vuetify.breakpoint.width <= 959 ? 'd-flex flex-column text-center' :
                                    'd-flex flex-column text-left'" cols="9">
                                    <span class="text-h6 text-capitalize font-weight-bold mb-2 contentHead--text">{{
                                        cont.text
                                    }}</span>
                                    <p class="body-1 content--text button">{{ cont.sentence }}</p>
                                </v-col>
                            </v-row>
                        </v-col>
                    </v-row>
                </v-col>
            </v-row>
        </v-container>
    </div>
</template>

<script>
export default {
    data() {
        return {
            content: [
                {
                    image: "https://flattrade.s3.ap-south-1.amazonaws.com/promo/Ipo2.webp", text: "IPO", sentence: "Seize the opportunity to invest in quality companies and grow your wealth "
                },
                // https://flattrade.s3.ap-south-1.amazonaws.com/promo/mutualfunds2.webp
                { image: "https://flattrade.s3.ap-south-1.amazonaws.com/promo/mutalfundsdash.webp", text: "Direct Mutual Funds", sentence: "Take control over your long-term investments by investing in commission-free MF schemes" },
                { image: "https://flattrade.s3.ap-south-1.amazonaws.com/promo/sgb.webp", text: "Sovereign Gold Bonds", sentence: "Terrific choice to invest in gold and earn a fixed interest of 2.5% as well as enjoy the appreciation in price of the yellow metal" },
                { image: "https://flattrade.s3.ap-south-1.amazonaws.com/promo/gsec2.webp", text: "G-sec", sentence: "Secure your investments by investing in government-backed risk-free securities which are also highly liquid " },
                { image: "https://flattrade.s3.ap-south-1.amazonaws.com/promo/corporatebond2.webp", text: "Corporate Bonds", sentence: "Highly rated corporate bonds that offer better returns than traditional bank deposits" },

            ]
        }
    }
}
</script>

<style scoped>
.card {
    /* border: 1px solid ; */
    border-radius: 10px;
    border-color: transparent;
    /* background-color:#78c2f349    */
    background: linear-gradient(90deg, #879EF5 -91.93%, rgba(173, 187, 241, 0.170062) 10.46%, rgba(135, 158, 245, 0) 62.98%),
        linear-gradient(270deg, #879EF5 -91.93%, rgba(173, 187, 241, 0.170062) 10.46%, rgba(135, 158, 245, 0) 62.98%);

}

@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@200;400;500;600;700;800;900&display=swap");

h1,
h4,
p {
    font-family: 'Poppins', sans-serif !important;
}

h4 {
    font-size: 1.25rem;
    font-weight: 600;
}

.v-application .text-h6 {
    font-size: 1.1rem !important;
}

.float_card {
    background: #FFFFFF !important;
    box-shadow: 0px 4px 16px #EAECF5 !important;
    border-radius: 16px !important;
    border: 0;
}

.img-fluid {
    max-width: 100%;
    height: 90px;
}

@media only screen and (min-width:1024px) {
    @keyframes slide-forward {
        0% {
            transform: translateY(0px);
            opacity: 0.8;
        }

        100% {
            transform: translateY(-10%);
            position: relative;
            opacity: 1;
        }

    }
}

@media only screen and (min-width:768px) {
    .slide-in {
        transform: translateY(10%);
        position: relative;
        opacity: 1;
    }
}

.slide-in {
    animation: slide-forward .5s ease forwards;
}

.elevation {
    box-shadow: 1px 0px 54px 16px rgb(42, 57, 78, 0.1);
}

.col {
    width: 100%;
    padding: 5px;
}

.custom_container {
    max-width: 800px;
}
</style>