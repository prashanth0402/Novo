<template>
  <v-app id="app">
    <Main v-if="logged && !loading" />
    <Login v-if="!logged && !loading" />
    <v-container v-if="loading">
      <v-app>
        <v-layout align-center v-if="$globalData.overlay">
          <v-flex d-flex justify-center>
            <!-- <img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Loader.webp" /> -->
            <Loader />
          </v-flex>
        </v-layout>
      </v-app>
    </v-container>
  </v-app>
</template>


<script>

import Main from "@/views/Main.vue";
import Login from "@/views/LandingPage.vue";
import EventService from "@/services/EventServices.js";
import Loader from "./components/customLoader/loader.vue";
export default {
  name: "App",
  components: {
    Main,
    Login,
    Loader
  },
  data: () => ({
    visible: false,
    logged: false,
    loading: false,
    abhiinput: {
      clientID: "",
      token: "",
    },
  }),
  methods: {

  },
  beforeMount() {
    this.loading = true;
    const urlParams = new URLSearchParams(window.location.search);
    // console.log(this.$route.path);
    if (
      this.$route.query.token != "" &&
      this.$route.query.LoginId != "" &&
      this.$route.query.token &&
      this.$route.query.LoginId
    ) {
      //console.log("with param");
      this.abhiinput.clientID = this.$route.query.LoginId;
      this.abhiinput.token = this.$route.query.token;

      //getting token and set the token in cookies
      EventService.token(this.abhiinput)
        .then((response) => {
          // console.log(response);
          if (response.data.status == "S") {
            if (urlParams.get("rd") && urlParams.get("rd") != "") {
              window.location = this.MainDomain + urlParams.get("rd");
            } else {
              this.logged = true;
              this.$globalData.logged = true;
              this.loading = false;
              // this.$router.replace("/ipo"); // Before dashboard create
              this.$router.replace("/dashboard"); // After dashboard Created
            }
          } else if (response.data.status == "I") {
            this.logged = false;
            this.$globalData.logged = false;
            this.loading = false;
          } else if (response.data.status == "E") {
            this.logged = false;
            this.$globalData.logged = false;
            this.loading = false;
          }
        })
        .catch(() => {
          this.logged = false;
          this.loading = false;
        });
    } else {
      // console.log("without param");

      EventService.tokenValidation()
        .then((response) => {
          if (response.data.status != "S") {
            this.logged = false;
            this.$globalData.logged = false;
            this.loading = false;
          } else {
            this.logged = true;
            this.$globalData.logged = true;
            this.loading = false;
            if (this.$route.path == "/") {
              // this.$router.replace("/ipo"); // Before dashboard create
              this.$router.replace("/dashboard"); // After dashboard Created
            }
          }
        })
        .catch(() => {
          this.logged = false;
          this.$globalData.logged = false;
          this.loading = false;
        });
    }
    // console.error = function () { };
    // console.warn = function () { };
  },
};
</script>


<style>
#app {
  font-family: 'Inter', sans-serif !important;
  font-size: -10px;
}
</style>

