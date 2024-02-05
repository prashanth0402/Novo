<template>
    <div class="mt-16">
        <v-container class="d-flex justify-center">
            <div class="custom_container ">
                <v-row class="d-flex justify-space-between align-center">
                    <v-col cols="12" xl="5" lg="6" md="6" sm="6" xs="8"
                        :class="$vuetify.breakpoint.width <= 600 ? 'text-center d-flex flex-column align-center justify-start' : 'text-left d-flex align-center justify-start'"
                        width="50px">
                        <span>
                            <span
                                :class="$vuetify.breakpoint.width < 920 ? 'text-h4 mb-2 font-weight-bold' : 'text-h3 font-weight-bold'">Sign
                                up now to start
                                investing</span>
                            <!-- <v-icon color="blue lighten-2" class="mt-2">mdi-arrow-right-circle-outline</v-icon> -->
                        </span>
                        <v-hover v-slot="{ hover }" v-if="$vuetify.breakpoint.width <= 600">
                            <v-btn class="text-h6 text-capitalize white--text elevation-0" large
                                :color="hover ? 'secondary' : 'btnColor'">
                                Sign up now
                            </v-btn>
                        </v-hover>
                    </v-col>
                    <v-col cols="12" xl="5" lg="6" md="6" sm="6" xs="8" class="text-left d-flex align-center justify-end"
                        v-if="$vuetify.breakpoint.width > 600">
                        <v-hover v-slot="{ hover }">
                            <v-btn class="text-h6 text-capitalize elevation-0 " large
                                :color="hover ? 'secondary' : 'btnColor'">
                                <a :href="url" target="_blank" class="white--text "> Sign up now</a>
                            </v-btn>
                        </v-hover>
                    </v-col>
                </v-row>
            </div>
        </v-container>
    </div>
</template>

<script>
export default {
    data() {
        return {
            url: "https://ekyc.flattrade.in/openaccount/?utm_source=novo&utm_content=novo",
            //dialog: false,
        };
    },
}
</script>

<style  scoped>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@200;400;500;600;700;800;900&display=swap");

a {
    text-decoration: none;
    color: #fff;
}

.v-application .text-h3,
.text-h4,
.v-btn {
    line-height: 3.290rem !important;
    font-family: 'Poppins', sans-serif !important;
}

.custom_container{
    max-width: 80%;
}
</style>