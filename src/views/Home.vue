<template>
  <div class="home">
    <!-- <v-progress-linear indeterminate color="primary" v-if="loading"></v-progress-linear> -->
    <MainComponent v-if="allowed" class="mt-10" />
    <Rules v-if="allowed" class="mt-10" />
  </div>
</template>

<script>
import MainComponent from '../components/MainComponent.vue';
import EventServices from '../services/EventServices';
import Rules from '../components/Instruct/Rules.vue';
export default {
  name: "Home",
  components: {
    MainComponent,
    Rules,
  },
  data: () => ({
    person: {
      Id: "",
    },
    allowed: false,
    loading: true,
  }),
  methods: {
    Token() {
      this.$globalData.overlay = true;
      this.loading = true;
      EventServices.tokenValidation()
        .then((response) => {
          if (response.data.status != "S") {
            this.$globalData.overlay = false;
            this.loading = false;
            this.allowed = false;
            window.location = this.LoginUrl;
            // console.log("Valid Failed", response.data);
          } else {
            // console.log("Valid");
            this.person.Id = response.data.ClientId;
            EventServices.RouterValidation(this.$route.path)
              .then((response) => {
                if (response.data.status != "S") {
                  this.$globalData.overlay = false;
                  this.loading = false;
                  //   window.location = this.LoginUrl;
                  if (this.$globalData.links.length > 0) {
                    this.$router.replace(this.$globalData.links[0])
                  } else {
                    this.$router.replace('/')
                  }
                } else {
                  //allow the access to this page
                  this.$globalData.overlay = false;
                  this.loading = false;
                  this.allowed = true;
                }
              })
              .catch((error) => {
                this.$globalData.overlay = false;
                this.loading = false;
                this.MessageBar("E", error);
              });
          }
        })
        .catch(() => {
          this.$globalData.overlay = false;
          this.loading = false;
          this.allowed = false;
          window.location = this.LoginUrl;
        });
    }
  },
  created() {
    if (this.$globalData.logged == true) {
      if (this.$globalData.logged == true) {
        this.Token();
      } else {
        this.$router.replace("/")
      };
    } else {
      this.$router.replace("/")
    }
  },
};
</script>
  