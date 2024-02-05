import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import vuetify from "./plugins/vuetify";
import VueLottie from 'vue-lottie';

Vue.use(VueLottie);

Vue.config.productionTip = false;

Vue.prototype.$globalData = Vue.observable({
  logged: false,
  msg: "",
  alertColor: "",
  alert: "",
  backgroundColor: "",
  IconColor: "",
  snackbar: false,
  links: [],
  subMenu: [],
  overlay: false,
  timeout: -1,
  host: "",
  appName: "",
  url: "",
  currentTime:""
});
Vue.mixin({
  data: function () {
    return {
      LoginUrl: "",
      redirectUrl: "",
      MainDomain: "",

      // LoginUrl: "http://localhost:8080/home",
      // redirectUrl: "https://auth.flattrade.in/?app=novodev",
      // MainDomain: "http://localhost:8080"

      // LoginUrl: "https://novo.flattrade.in/home",
      // redirectUrl: "https://auth.flattrade.in/?app=novo",
      // MainDomain:"https://novo.flattrade.in"
    };
  },
  methods: {
    AssignUrl() {
      this.LoginUrl = "http://" + this.$globalData.host + "/home"
      this.redirectUrl = "https://" + this.$globalData.url
      this.MainDomain = "http://" + this.$globalData.host
    },
    MessageBar: function (indicator, Msg) {
      this.$globalData.msg = Msg;
      this.$globalData.alert = true;
      this.$globalData.snackbar = true;
      if (indicator == "S") {
        this.$globalData.backgroundColor = "green lighten-5";
        this.$globalData.IconColor = "green darken-2";
        this.$globalData.Icon = "mdi-check-circle-outline";
        this.$globalData.alerttitle = "Success";
        this.$globalData.timeout = 3000;
      } else if (indicator == "E") {
        this.$globalData.backgroundColor = "error lighten-5";
        this.$globalData.IconColor = "error darken-2";
        this.$globalData.Icon = "mdi-alert";
        this.$globalData.alerttitle = "Error";
        this.$globalData.timeout = 3000;
      }
    },
    closeAlert() {
      setTimeout(() => {
        this.$globalData.alert = false;
        this.$globalData.snackbar = false;
      }, 3000);
    },
     // to find current time
     GetCurrentTime() {
      // console.log("Calling time function");
      const currentTime = new Date();
      let hours = currentTime.getHours();
      let minutes = currentTime.getMinutes();
      let seconds = currentTime.getSeconds();
      hours = (hours < 10 ? "0" : "") + hours;
      minutes = (minutes < 10 ? "0" : "") + minutes;
      seconds = (seconds < 10 ? "0" : "") + seconds;
      //this.currentTime = `08:00:00`;

      this.$globalData.currentTime = `${hours}:${minutes}:${seconds}`;
    },
  },
});

new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App),
}).$mount("#app");
