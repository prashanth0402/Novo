<template>
    <div class="container my-10 d-flex justify-center">
        <div class="mobile-mock-bg rounded-xl">
            <v-row class="pa-5">
                <v-col class="d-flex justify-center align-center flex-column text-center" cols="12" xl="6" lg="6" md="6"
                    sm="6" xs="6" :class="`mt-${($vuetify.breakpoint.width < 591) ? 0 : 8}`">
                    <span class="d-none d-sm-flex text-h4 font-weight-bold mb-2 ">
                        Download and Start your Investment Journey
                    </span>
                    <span class="d-flex d-sm-none text-h5 font-weight-bold mb-2 "> Download and Start your
                        Investment Journey</span>
                    <v-row :class="`mt-${($vuetify.breakpoint.width < 950) ? 5 : 0}`">
                        <v-col class="d-flex ma-auto  justify-center" v-for="link, index in links" :key="index">
                            <a :href="link.path" v-if="index < 2" target="_blank">
                                <img :src="link.image" :lazy-src="link.image" class="subscribe_mobile" :alt="link.alt" />
                            </a>
                            <a :href="link.path" v-else target="_blank">
                                <img :src="link.image" :lazy-src="link.image" class="subscribe_mobile mt-1" :alt="link.alt" height="75%"/>
                            </a>
                        </v-col>
                    </v-row>
                </v-col>
                <v-col class="d-flex justify-center align-center" cols="12" xl="6" lg="6" md="6" sm="6" xs="6">
                    <!-- <v-img lazy-src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Mobile App Image NOVO.webp"
                        src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Mobile App Image NOVO.webp" contain
                        class="d-flex align-center justify-end mock_desk">
                    </v-img> -->
                    <v-img lazy-src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Mobile Novo.webp"
                        src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Mobile Novo.webp" contain
                        class="d-flex align-center justify-end mock_desk">
                    </v-img>
                </v-col>
            </v-row>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            links: [
                { image: "https://flattrade.s3.ap-south-1.amazonaws.com/promo/Novo App Store.webp", alt: "appstore", path: 'https://apps.apple.com/in/app/novo-ipo-sgb-mf-g-sec/id6473023986' },
                { image: "https://flattrade.s3.ap-south-1.amazonaws.com/promo/Novo Play store.webp", alt: "playstore", path: 'https://play.google.com/store/apps/details?id=com.novo.novoapp' },
                // { image: "https://flattrade.s3.ap-south-1.amazonaws.com/promo/web.webp", alt: "web", path: 'https://web.flattrade.in/#/' }
            ]
        }
    }
}
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@200;400;500;600;700;800;900&display=swap");

.v-application .text-h3,
.text-h4,
.text-h5 {
    font-family: 'Poppins', sans-serif !important;
}

.mobile-mock-bg {
    max-width: 1220px;
    background: #f2ecfd;
}

.mock_desk {
    max-width: 300px;
}

.mock_mobile {
    max-width: 50px;
}

.custom_container {
    max-width: 800px;
}

.col {
    width: 100%;
    padding: 5px;
}

.subscribe_desktop {
    max-width: 150px;
}

.subscribe_mobile {
    max-width: 150px;
}
</style>