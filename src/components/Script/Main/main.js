import EventService from "@/services/EventServices.js";
export default {
    name: "Main",

    data: () => ({
        person: {
            Id: "",
        },
        // links: [],
        drawer: false,
        userSubMenus: [],
    }),

    beforeMount() {
        EventService.tokenValidation()
            .then((response) => {
                if (response.data.status != "S") {
                    window.location = this.LoginUrl;
                } else {
                    this.person.Id = response.data.clientId;
                    this.verify();
                    // this.getDirectory();
                }
            })
            .catch(() => {
                window.location = this.LoginUrl;
            });
        // console.error = function () { };
        // console.warn = function () { };

    },
    computed: {
        isNotHome() {
            // Replace '/home' with the actual path of your home route
            return this.$route.path !== '/';
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
            // console.log("close")
            this.closeDrawer()
        },
        changeRoute(route) {
            if (this.$route.path !== route) {
                // this.$router.push(route);
                this.$router.replace(route)
                if (route == '/') {
                    //console.log(route);
                    EventService.DeleteCookie()
                        .then((response) => {
                            if (response.data.status == "S") {
                                // this.$router.replace('/')
                                window.location = this.LoginUrl;
                            }
                        })
                        .catch((error) => {
                            this.MessageBar('E', error)
                        });
                }
            }
        },
        verify() {
            EventService.verifyClient()
                .then((response) => {
                    if (response.data.status == "S") {
                        this.$globalData.links = response.data.routerArr

                        if (this.$globalData.links[0].path == "" || this.$globalData.links == null) {

                            EventService.GetSubMenu(this.$globalData.links[0].routerId)
                                .then((response) => {
                                    //console.log(response.data);
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
                                    this.MessageBar('E', error)// Logs out the error
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
    },
}
