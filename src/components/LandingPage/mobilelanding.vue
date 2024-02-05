<template>
  <v-container>
    <v-layout row wrap class="mb-8">
      <v-flex class="d-flex justify-center">
        <table>
          <tr>
            <td align="right" class="pb-2">
              <img
                src="https://flattrade.s3.ap-south-1.amazonaws.com/logo/flattradegrey.png"
                height="11px"
              />&nbsp;
            </td>
          </tr>
          <tr>
            <td>
              <img class="" src="https://flattrade.s3.ap-south-1.amazonaws.com/logo/novologo.png" height="35" />
            </td>
          </tr>
        </table>
      </v-flex>
    </v-layout>

    <v-layout row wrap class="subtitle-1">
      <v-flex class="d-flex justify-center pa-5">
        <v-container>
          <span>
            <b class="d-flex justify-center">Apply for IPOs online</b>
            <p class="text-wrapper d-flex justify-center ">
              Apply online and invest in companies listing on the Indian
              exchanges <br v-if="this.$vuetify.breakpoint.name == 'sm'">
              with an IPO ( Initial Public Offering) using mobile UPI with
              your Flattrade account.
            </p>
          </span>
        </v-container>
      </v-flex>
    </v-layout>
    <!-- <br v-if="this.$vuetify.breakpoint.name == 'xs'"> 
    <br v-if="this.$vuetify.breakpoint.name == 'xs'">-->

    <v-layout row wrap class="mb-3">
      <v-flex class="d-flex justify-center">
        <v-btn @click="navigateTo" class="blue darken-3 rounded-0" elevation="0" dark height="30">Login</v-btn>
      </v-flex>
    </v-layout>
    <v-layout row wrap>
      <v-flex class="d-flex justify-center">
        <span class="grey--text caption">Don't have an account?

          <a target="_blank" :href="url">Sign up</a> now!
        </span>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
//import UsageTerms from "./UsageTerms.vue";

export default {
  name: "MobileView",

  components: {
    // UsageTerms,
  },

  data() {
    return {
      url: "https://ekyc.flattrade.in/openaccount/?utm_source=novo&utm_content=novo",
      // dialog: false,
    };
  },
  methods: {
    navigateTo() {
      this.$emit("GetRedirectUrl")
      // window.location.href = this.redirectUrl;
      // console.error = function () { };
      // console.warn = function () { };

    },
  },
};
</script>

<style scoped>
.text-wrapper {
  display: flex;
  flex-wrap: wrap;
  font-size: 14px;
  line-height: 24px;
}
</style>