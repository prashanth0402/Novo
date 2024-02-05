<template>
  <v-container>
    <v-layout align-center>
      <v-flex lg2> </v-flex>
      <v-flex d-flex justify-end class="ma-5" lg2>
        <table>
          <tr>
            <td align="right" class="pb-1">
              <img
                src="https://flattrade.s3.ap-south-1.amazonaws.com/logo/flattradegrey.png"
                height="11px"
              />&nbsp;
            </td>
          </tr>
          <tr>
            <td>
              <img
                src="https://flattrade.s3.ap-south-1.amazonaws.com/logo/novologo.png"
                :height="height"
              />
            </td>
          </tr>
        </table>
      </v-flex>
      <!-- https://flattrade.s3.ap-south-1.amazonaws.com/logo/novo.svg
https://flattrade.s3.ap-south-1.amazonaws.com/logo/novo.svg -->
      <v-divider vertical></v-divider>
      <v-flex lg6 class="pa-6">
        <v-container>
          <v-layout row wrap class="mb-4 subtitle-1">
            <v-flex>
              <span>
                <b>Apply for IPOs online</b>
                <p>
                  Apply online and invest in companies listing on the Indian
                  exchanges with an IPO ( Initial Public Offering) using mobile
                  UPI with your Flattrade account.
                </p>
              </span>
            </v-flex>
          </v-layout>

          <v-layout row wrap class="mb-3">
            <v-flex>
              <v-btn
                @click="navigateTo"
                class="blue darken-3 rounded-0"
                dark
                height="30"
                >Login</v-btn
              >
            </v-flex>
          </v-layout>

          <v-layout row wrap>
            <v-flex>
              <span class="grey--text caption"
                >Don't have an account?
                <a target="_blank" :href="url">Sign up</a> now!
              </span>
            </v-flex>
          </v-layout>
        </v-container>
      </v-flex>
      <v-flex lg2> </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
//import UsageTerms from "./UsageTerms.vue";

export default {
  name: "DesktopView",

  components: {
    // UsageTerms,
  },
  data() {
    return {
      url: "https://ekyc.flattrade.in/openaccount/?utm_source=novo&utm_content=novo",
      //dialog: false,
    };
  },
  methods: {
    navigateTo() {
      this.$emit("GetRedirectUrl")
      // window.location.href = this.redirectUrl;
      // console.error = function () {};
      // console.warn = function () {};
    },
  },
  computed: {
    height: {
      get() {
        if (this.$vuetify.breakpoint.name == "xs") {
          return 40;
        } else if (this.$vuetify.breakpoint.name == "sm") {
          return 60;
        } else if (this.$vuetify.breakpoint.name == "md") {
          return 35;
        } else if (this.$vuetify.breakpoint.name == "lg") {
          return 50;
        } else if (this.$vuetify.breakpoint.name == "xl") {
          return 50;
        } else {
          return 0;
        }
      },
    },
  },
};
</script>

<style>
.v-application--wrap {
  justify-content: center;
}
</style>
