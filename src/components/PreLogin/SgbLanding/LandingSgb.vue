<template>
  <div>
    <div class="banner_gradiant">
      <v-container>
        <v-row class="d-flex align-center justify-center custom_container">
          <v-col class="page-Bg">
            <v-row>
              <v-col cols="12" xl="7" lg="7" md="6" sm="12" xs="12" justify="center" align="center"
                v-if="!($vuetify.breakpoint.width < 950)">
                <v-img lazy-src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Sgb Resized.webp"
                  src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Sgb Resized.webp" contain
                  max-width="80%"></v-img>
              </v-col>
              <v-col class="d-flex justify-center align-center flex-column" cols="12" xl="5" lg="5" md="6" sm="12" xs="12"
                :class="`mt-${($vuetify.breakpoint.width < 950) ? 10 : 0}`">
                <v-col class="d-flex flex-column text-center justify-center align-center">
                  <span class="text-capitalize font-weight-bold mb-2 content--text"
                    :class="$vuetify.breakpoint.width > 720 ? 'text-h3' : 'text-h5'">Sovereign Gold Bonds</span>
                  <p class="content--text button">Invest in the yellow metal and also earn 2.5% interest per annum.
                    There is no capital gains tax.</p>
                  <v-hover v-slot="{ hover }">
                    <v-btn :color="hover ? 'secondary mt-5' : 'btnColor white--text mt-5'" width="150" depressed
                      style="border-radius: 10px;" @click="GetRedirectUrl">Login</v-btn>
                  </v-hover>
                </v-col>
              </v-col>
            </v-row>
          </v-col>
        </v-row>
      </v-container>
    </div>
    <InvestSgb />
    <MobileMock />
    <SignupMock />
  </div>
</template>

<script>
import InvestSgb from "../SgbLanding/InvestmentChoice/InvestSgb.vue"
import MobileMock from "../DownloadMockUp/MobileMock.vue";
import EventService from "../../../services/EventServices";
import SignupMock from "../SignUp/signupMock.vue";
export default {
  components: {
    InvestSgb,
    MobileMock,
    SignupMock
  },
  methods: {
    async GetRedirectUrl() {
      await EventService.GetRedirectURL()
        .then((response) => {
          if (response.data.status == "S") {

            this.$globalData.host = response.data.redirectUrl.host
            this.$globalData.appName = response.data.redirectUrl.appName
            this.$globalData.url = response.data.redirectUrl.url

            this.AssignUrl();
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

.banner_gradiant {
  background: linear-gradient(360deg, #ffffff -91.93%, #fffbec00 52.46%, #f7e70052 85%);

}

.custom_container {
  max-width: 100%;
  margin: 0 auto;
}

@media (min-width: 850px) {
  .page-Bg {
    background-size: cover;
    background-repeat: no-repeat;
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
