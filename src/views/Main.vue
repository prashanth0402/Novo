<template>
  <v-app>

    <v-overlay align-center v-if="$globalData.overlay" color="#FFFFFF">
      <v-flex d-flex justify-center>
        <Loader />
      </v-flex>
    </v-overlay>

    <!-- Snack bar appears top right -->
    <v-snackbar class="mt-15" :color="$globalData.backgroundColor" rounded="0" elevation="1" width="200px"
      v-model="$globalData.snackbar" top right :timeout="$globalData.timeout">
      <v-layout>
        <!-- <v-flex lg1> <v-icon medium color="orange lighten-1">mdi-alert</v-icon></v-flex> -->
        <v-flex lg1 sm1 xs1 md1 xl1>
          <v-icon medium :color="$globalData.IconColor">{{
            $globalData.Icon
          }}</v-icon></v-flex>

        <v-flex lg11 sm11 xs11 md11 xl11>
          <span class="black--text body-2 pl-2">{{
            $globalData.alerttitle
          }}</span>
          <div class="caption pl-2" style="color: #424242">
            {{ $globalData.msg }}
          </div>
        </v-flex>
      </v-layout>
    </v-snackbar>

    <!-- hearder comes here -->
    <Header :clientId="person.Id" :hideId="allowApp" :ClientName="clientname" />


    <v-main>
      <router-view />
    </v-main>


    <OfflineBanner />
    <!-- footer comes here -->
    <Footer class="mt-16" v-if="isNotHome && person.Id == null || this.$route.path == '/'" />

    <v-footer padless v-if="isNotHome && person.Id != null && this.$route.path != '/'" class=" footer" style="width:">
      <v-col style="padding: 5px;">
        <span class="font-weight-light text-size d-flex justify-center pa-2">FLATTRADE © {{ new
          Date().getFullYear()
        }}. All
          rights reserved. SEBI Registration No. INZ000201438 |
          MemberCode for NSE: 14572 | BSE:6524 | MCX: 16765 | ICEX: 2010</span>
      </v-col>
    </v-footer>
  </v-app>
</template>
<script>
// import NavMenu from "../components/Menu/NavMenu.vue";
// import VerticalMenu from "../components/Menu/VerticalMenu.vue";
import EventService from "@/services/EventServices.js";
import Header from "../components/Menu/Header.vue"
import Footer from "../components/Footer/MainFooter.vue"
import OfflineBanner from "../components/OfflineBanner/OfflineBanner.vue";
import Loader from "../components/customLoader/loader.vue";
export default {
  name: "Main",

  data: () => ({
    person: {
      Id: "",
    },
    clientname: "",
    drawer: false,
    userSubMenus: [],
    CurrentPath: "",
    allowApp: true,
  }),
  components: {
    Header,
    // NavMenu,
    // VerticalMenu,
    Footer,
    OfflineBanner,
    Loader
  },
  beforeMount() {
    EventService.tokenValidation()
      .then((response) => {
        if (response.data.status != "S") {
          window.location = this.LoginUrl;
        } else {
          this.person.Id = response.data.clientId;
          // console.log("calling verifyAccess");
          this.verify();
          this.getClientName();
        }
      })
      .catch(() => {
        window.location = this.LoginUrl;
      });
    // console.error = function () { };
    // console.warn = function () { };

  },
  beforeDestroy() {
    EventService.DeleteCookie()
      .then((response) => {
        if (response.data.status == "S") {
          window.location = this.LoginUrl;
        }
      })
      .catch((error) => {
        this.MessageBar('E', error)
      });
  },
  computed: {
    isNotHome() {
      return this.$route.path != "/home" && !this.$route.meta.hideAppBar;

    },
    hidetab: {
      get() {
        if (this.$vuetify.breakpoint.name == 'xs') {
          return false
        } else {
          return true
        }
      }
    }
  },
  methods: {
    onClickOutsideStandard(event) {
      if (this.drawer && !event.target.closest(".v-navigation-drawer")) {
        this.closeDrawer()
      }
    },
    CloseMenu() {
      this.closeDrawer()
    },
    verify() {
      EventService.verifyClient()
        .then((response) => {
          if (response.data.status == "S") {
            this.$globalData.links = response.data.routerArr

            if (this.$globalData.links[0].path == "" || this.$globalData.links == null) {

              EventService.GetSubMenu(this.$globalData.links[0].routerId)
                .then((response) => {
                  if (response.data.subMenuArr == null) {
                    this.userSubMenus = []
                  } else {
                    this.userSubMenus = response.data.subMenuArr;
                  }
                  if (this.$route.path !== this.userSubMenus[0].path) {
                    this.$router.replace(this.userSubMenus[0].path)
                  }

                })
                .catch((error) => {
                  this.MessageBar('E', error.response)
                });
            }
          } else {
            this.$router.replace("/underConstruction")
          }
        })
        .catch((error) => {
          this.MessageBar('E', error)
        });
    },

    closeDrawer() {
      this.drawer = false;
    },
    getDirectory() {
      EventService.GetDirectory()
        .then((response) => {
          if (response.data.status == "S") {
            this.$store.commit('setStream', response.data.stream)
          }
        })
        .catch((error) => {
          this.MessageBar('E', error)
        });
    },
    getClientName() {
      EventService.GetClientName()
        .then((response) => {
          if (response.data.status == "S") {
            this.clientname = response.data.clientName;
          } else {
            this.MessageBar('E', response.data.errMsg)
          }
        })
        .catch((error) => {
          this.MessageBar('E', error)
        });
    },
   
  },
}

</script>

<style scoped>
::v-deep.v-main {
  padding: 46px 0px 0px !important;
}

.footer {
  background: #f8f9fc;
  /* padding: 60px 0; */
  position: relative;
  color: #666;
}

.v-app-bar {
  background: rgba(255, 255, 255, 0.3);
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
  backdrop-filter: blur(4.5px);
  -webkit-backdrop-filter: blur(4.5px);
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.18);
}

.text-size {
  font-size: 11px;
}

::v-deep .v-list v-sheet theme--light {
  max-height: 63px;
}

::v-deep .v-menu__content theme--light v-menu__content--fixed {
  max-height: 63px;
}

@media (max-width:600px) {
  .text-size {
    font-size: 8px;
  }
}
</style>
  