<template>
  <div>
    <!-- <v-progress-linear indeterminate color="primary" v-if="loading"></v-progress-linear> -->
    <configure v-if="allowed" />

    <template>

    </template>
  </div>
</template >

<script>
import configure from '../components/Config/configure.vue';
import EventServices from "../services/EventServices";
EventServices;
export default {
  components: {
    configure,
  },
  data() {
    return {
      allowed: false,
      // loading: true,
    };
  },
  methods: {
    Token() {
      this.$globalData.overlay = true;
      this.loading = true;
      EventServices.tokenValidation()
        .then((response) => {
          if (response.data.status != "S") {
            this.$globalData.overlay = false;
            this.loading = false;
            window.location = this.LoginUrl;
          } else {
            EventServices.RouterValidation(this.$route.path)
              .then((response) => {
                if (response.data.status != "S") {
                  this.$globalData.overlay = false;
                  this.loading = false;
                  // window.location = this.LoginUrl;
                  this.$router.replace(this.$globalData.links[0].path)
                } else {
                  //allow the access to this page
                  this.$globalData.overlay = false;
                  // this.loading = false;
                  this.allowed = true;
                }
              })
              .catch((error) => {
                this.$globalData.overlay = false;
                // this.loading = false;
                this.MessageBar("E", error);
              });
          }
        })
        .catch(() => {
          this.$globalData.overlay = false;
          // this.loading = false;    
          window.location = this.LoginUrl;
        });
    },
  },
  created() {
     if (this.$globalData.logged == true) {
      this.Token();
    } else {
      this.$router.replace("/")
    };
  }
};
</script>

<style scoped>
::v-deep .v-stepper__content {
  margin: -18px -14px -23px 34px !important;
  border-left: 3px solid #1E88E5 !important;
  height: 50px !important;
}

::v-deep .v-stepper__wrapper {
  display: none !important;
}

::v-deep .v-sheet.v-stepper:not(.v-sheet--outlined) {
  box-shadow: none;
}
</style>