<template>
    <div id="app">
        <!-- <v-app-bar app color="#fff" elevation="0" :class="{ 'scrolled': isScrolled }"> -->
        <v-app-bar app color="#fff" :elevation="isScrolled ? 1 : 0">
            <div class="d-flex align-center justify-center container">
                <v-toolbar-title>
                    <img :src="!($vuetify.breakpoint.width < 950) ? 'https://flattrade.s3.ap-south-1.amazonaws.com/promo/Novo Transp.webp'
                        : 'https://flattrade.s3.ap-south-1.amazonaws.com/promo/Novo Fav 1.webp'" alt="logo" height="25"
                        @click="$globalData.logged == true ? changeRoute('Dashboard', '/dashboard', 'N') : changeRoute('Home', '/', '')"
                        style="cursor: pointer;">
                </v-toolbar-title>
                <v-spacer></v-spacer>

                <v-spacer v-if="$vuetify.breakpoint.width < 950"></v-spacer>

                <!-- This Menu change dynamically according to the user login or not -->
                <v-row v-else class="d-flex justify-end align-center">
                    <v-col :cols="$globalData.logged ? 11 : 12" class="d-flex justify-end align-center">
                        <v-menu v-for="btn in routers" :key="btn.router" offset-y v-show="clientId != ''">
                            <template v-slot:activator="{ attrs, on }">

                                <v-hover v-slot="{ hover }" v-if="btn.router != 'Login'">
                                    <span :class="{ 'active-tab': btn.path == $route.path }">
                                        <a v-bind="attrs" v-on="on"
                                            :class="hover ? 'primary--text mx-5 font-weight-medium' : 'black--text mx-5 font-weight-medium'"
                                            @click="changeRoute(btn.router, btn.path, 'N')" small text>
                                            {{ btn.router }}
                                            <v-icon v-if="btn.path == ''">mdi-menu-down</v-icon>
                                        </a>
                                    </span>
                                </v-hover>
                                <v-hover v-slot="{ hover }" v-else>
                                    <a v-bind="attrs" v-on="on"
                                        :class="hover ? 'gradiant white--text mx-5 font-weight-bold' : 'btn black--text mx-5 font-weight-bold'"
                                        @click="changeRoute(btn.router, btn.path, '')" small text>
                                        {{ btn.router }}
                                    </a>
                                </v-hover>
                            </template>
                            <SubMenu :parentMenuId="btn.routerId" v-if="btn.path == ''" />
                        </v-menu>
                    </v-col>
                    <v-col cols="1" class="d-flex justify-center" v-if="$globalData.logged">
                        <v-menu bottom offset-x-reverse offset-y rounded v-if="this.$vuetify.breakpoint.width
                            >= 900 && hideId" :close-on-click="true">
                            <template v-slot:activator="{ on, attrs }">
                                <v-avatar v-bind="attrs" v-on="on" color="blue lighten-4" size="45"
                                    style="cursor: pointer;">
                                    <span class="primary--text button font-weight-medium">{{ ClientName }}</span>
                                </v-avatar>
                            </template>

                            <v-list class="profile-menu mt-3" outlined>
                                <v-list-item>
                                    <v-list-item-title class="d-flex align-center">
                                        <v-img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Account.webp"
                                            contain width="25"></v-img>
                                        &nbsp; <span class="black--text font-weight-medium">{{ clientId }}</span>
                                    </v-list-item-title>
                                </v-list-item>
                                <v-list-item style="cursor: pointer;" @click="changeRoute('Logout', '/', 'Y')">
                                    <v-list-item-title class="d-flex align-center">
                                        <v-img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Power.webp" contain
                                            width="25"></v-img> &nbsp;
                                        <span> LogOut</span>
                                    </v-list-item-title>
                                </v-list-item>
                            </v-list>
                        </v-menu>
                    </v-col>
                </v-row>

                <v-app-bar-nav-icon @click.stop="drawer = !drawer"
                    v-if="this.$vuetify.breakpoint.width < 950"></v-app-bar-nav-icon>
                <!-- Switch Theme -->
                <!-- <v-icon v-if="!this.$vuetify.theme.dark" color="yellow" @click="toggleTheme"
                    size="30">mdi-white-balance-sunny</v-icon>
                <v-icon v-else color="white" @click="toggleTheme" size="30">mdi-weather-night</v-icon> -->
            </div>
        </v-app-bar>
        <v-navigation-drawer v-model="drawer" app v-if="this.$vuetify.breakpoint.width < 950" right>
            <v-row class="text-center pa-2 mt-2" v-if="hideId">
                <v-col class="d-flex flex-column align-center">
                    <v-progress-circular indeterminate color="primary" v-if="load"></v-progress-circular>
                    <v-avatar color="blue lighten-4" size="60" style="cursor: pointer;">
                        <span class="primary--text button font-weight-medium">{{ ClientName }}</span>
                    </v-avatar>
                    <span class="content--text mt-2">{{ clientId }}</span>
                </v-col>
            </v-row>
            <div :class="`mt-${!hideId ? 5 : 0}`">
                <v-list v-for="link in routers" :key="link.routerId">
                    <v-list-item @click="changeRoute(link.router, link.path)" v-if="link.path != ''"
                        :class="link.path == $route.path ? 'blue lighten-5' : undefined">
                        <span class="content--text">{{ link.router }}</span>
                    </v-list-item>

                    <v-list-group v-if="link.path == ''">
                        <template v-slot:activator>
                            <v-list-item-title>
                                <span>{{ link.router }}</span>
                            </v-list-item-title>
                        </template>

                        <SubMenu :parentMenuId="link.routerId" />
                    </v-list-group>
                </v-list>
            </div>
            <v-list v-if="hideId">
                <v-list-item-icon v-if="!hideId">
                    <v-img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/icons8-power-94.png" contain
                        width="30"></v-img>
                </v-list-item-icon>
                <v-list-item @click="changeRoute('Logout', '/', 'Y')"> LogOut</v-list-item>
            </v-list>
        </v-navigation-drawer>
    </div>
</template>

<script>
import EventService from "@/services/EventServices"
import SubMenu from './SubMenu.vue';
export default {
    components: {
        SubMenu
    },
    data: () => ({
        buttons: [
            { icon: '', router: "IPO", path: "/pre/ipo" },
            // { icon: '', router: "Mutual funds", path: "/pre/mutalFunds" },
            // { icon: '', router: "G-Secs", path: "/pre/gsec" },
            { icon: '', router: "SGB", path: "/pre/sgb" },
            // { icon: '', router: "Corporate bonds", path: "/pre/corporateBond" },
            { icon: '', router: "Login", path: "/" },
        ],
        appBarColor: 'rgba(248, 249, 252,0.8)', // Initial background color with some opacity
        drawer: false,
        load: false,
        isScrolled: false
    }),
    methods: {
        changeRoute(name, route, signal) {
            if (name == "Login") {
                this.GetRedirectUrl()
            } else {
                if (this.$route.path !== route && route != "") {
                    if (route == '/' && signal == 'Y') {
                        EventService.DeleteCookie()
                            .then((response) => {
                                if (response.data.status == "S") {
                                    window.location = this.LoginUrl;
                                }
                            })
                            .catch((error) => {
                                this.MessageBar('E', error)
                            });
                        signal = ""
                    }

                    this.$router.push(route)
                }
            }
        },
        // toggleTheme() {
        //     this.$vuetify.theme.dark = !this.$vuetify.theme.dark;
        //     this.handleScroll()
        // },
        //TO MAKE THE APPBAR LOOKS GLASS
        // handleScroll() {
        //     if (!this.$vuetify.theme.dark) {
        //         // Adjust the opacity based on the scroll position or any other conditions you need
        //         if (window.scrollY > 100) {
        //             this.appBarColor = 'rgba(255, 255, 255, 0.5)'; // Background becomes more transparent
        //         } else {
        //             this.appBarColor = 'rgba(255, 255, 255, 0.9)'; // Reset to initial opacity
        //         }
        //     } else {
        //         // Adjust the opacity based on the scroll position or any other conditions you need
        //         if (window.scrollY > 100) {
        //             this.appBarColor = 'rgba(39,39,39, 0.7)'; // Background becomes more transparent
        //         } else {
        //             this.appBarColor = 'rgba(39,39,39, 0.9)'; // Reset to initial opacity
        //         }
        //     }
        // },
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
        handleScroll() {
            const scrollPosition = window.scrollY;
            if (scrollPosition > 100) {
                this.isScrolled = true;
            } else {
                this.isScrolled = false;
            }
        }

    },
    computed: {
        routers: {
            get() {
                if (this.clientId != null) {
                    if (this.$route.path == "/dashboard") {

                        return this.$globalData.links.filter(link => link.router != "IPO" && link.router != "SGB" && link.router != "NCB" && link.router != "MutalFunds" && link.router != "CorporateBonds")
                    } else {
                        const order = ["IPO", "SGB", "NCB"];
                        return this.$globalData.links.slice().sort((a, b) => {
                            const indexA = order.indexOf(a.router);
                            const indexB = order.indexOf(b.router);

                            if (indexA === -1) {
                                return 1; // Move items not in the "order" array to the end
                            }
                            if (indexB === -1) {
                                return -1; // Move items not in the "order" array to the front
                            }
                            return indexA - indexB; // Sort based on the order index
                        });
                    }
                } else {
                    return this.buttons
                }
            }

        },
    },
    updated() {
        window.addEventListener('scroll', this.handleScroll);
    },
    beforeDestroy() {
        window.removeEventListener('scroll', this.handleScroll);
    },
    props: {
        clientId: String,
        hideId: Boolean,
        ClientName: String,
    },
    watch: {
        ClientName: function (name) {
            if (name == "") {
                this.load = true;
            } else {
                this.load = false;
            }
        }
    }
}
</script>

<style scoped>
.active-tab {
    padding-top: 5px;
    border-bottom: 2px solid #1976D2;
}

.scrolled {
    border-bottom: 1px solid #727272;
    /* Add your desired border style here */
}

span .link:hover {
    color: #1976D2;
    cursor: pointer;
}

.gradiant {
    background: #0965da;
    padding: 4px 20px;
    border: 1px solid #1976D2;
    border-radius: 25px;
}

.btn {
    color: #000;
    border: 1px solid #2A394E;
    padding: 4px 20px;
    border-radius: 25px;
}

/* Make the appbar like glass */
/* 
.v-app-bar {
    background: rgba(255, 255, 255, 0.3);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    border: 1px solid rgba(148, 148, 148, 0.596);
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    z-index: 1000;
} */

.v-list:nth-child(1) {
    padding: 8px 0;
}

.v-list:last-child {
    padding: 0;
}

.v-list {
    padding: 0;
}

.v-menu__content {
    box-shadow: none !important;
}


.profile-menu {
    position: relative;
    padding: 0 !important;
}

/* .profile-menu::before {
    position: absolute;
    content: '';
    height: 20px;
    width: 20px;
    transform: rotate(45deg);
    background: #fff;
    border-top: 1px solid #dddadabd;
    border-left: 1px solid #dddadabd;
    top: -10px;
    right: 20px !important;

} */


.v-col {
    padding: 0px !important;
}

.v-responsive {
    flex: none;
}
</style>