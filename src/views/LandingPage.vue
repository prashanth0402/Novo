<template>
  <v-app>
    <Header :hideId="hide" />
    <v-main class="relative">
      <router-view />
    </v-main>
    <!-- QrCode comes heare -->
    <QrCode :open="openQr" @close="closeQr" class="absolute" />
    <OfflineBanner />
    <Footer class="mt-16" />
  </v-app>
</template>

<script>
import EventService from "../services/EventServices";
import Header from "../components/Menu/Header.vue";
import Footer from "../components/Footer/MainFooter.vue";
import OfflineBanner from "../components/OfflineBanner/OfflineBanner.vue";
import QrCode from "../components/QrScanner/qrCode.vue";
export default {
  data() {
    return {
      screenWidth: window.innerWidth,
      hide: true,
      openQr: true,

    };
  },
  methods: {
    closeQr() {
      this.openQr = false;
    },
    handleResize() {
      this.screenWidth = window.innerWidth;
    },
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
  },
  computed: {
    isSmallScreen() {
      return this.screenWidth >= 600;
    },
  },
  mounted() {
    this.hide = false;
    window.addEventListener("resize", this.handleResize);
    // console.error = function () { };
    // console.warn = function () { };

  },
  beforeDestroy() {
    this.hide = true;
    window.removeEventListener("resize", this.handleResize);
  },
  components: {
    Header,
    Footer,
    OfflineBanner,
    QrCode
  },
};
</script>

<style>
.relative {
  position: relative;
  padding: 45px 0px !important;
}

.absolute {
  position: absolute;
  left: 0;
}
</style>
