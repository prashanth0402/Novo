<template>
    <div>
        <!-- <v-progress-linear indeterminate color="primary" v-if="loading"></v-progress-linear> -->

        <div class="container mt-10" v-if="allowed">
            <v-slide-x-transition mode="out-in" appear>
                <v-row>
                    <v-col cols="12">
                        <v-layout class="mt-10 d-flex justify-center align-center">
                            <v-flex>
                                <p>Press the button to fetch Updated or New IPO, SGB & NCB Details ---></p>
                            </v-flex>
                            <v-flex class="d-flex justify-end align-center">
                                <v-btn @click="manualFetch" small width="150" class="caption text-capitalize elevation-0"
                                    color="blue lighten-4">
                                    <v-icon size="15" class="primary--text mr-1">mdi-cached</v-icon>manualFetch</v-btn>
                            </v-flex>
                        </v-layout>
                    </v-col>
                    <v-col cols="12">
                        <v-divider></v-divider>
                    </v-col>
                    <v-col cols="12">
                        <v-layout class="mt-10 d-flex justify-center align-center">
                            <v-flex>
                                <p>Press the button to Run Offline Application Scheduler ---></p>
                            </v-flex>
                            <v-flex class="d-flex justify-end align-center">
                                <v-btn @click="ManualOfflineSch" small width="150"
                                    class="caption text-capitalize elevation-0" color="blue lighten-4">
                                    <v-icon size="15" class="primary--text mr-1">mdi-cached</v-icon>ManualOfflineSch</v-btn>
                            </v-flex>
                        </v-layout>
                    </v-col>
                    <v-col cols="12">
                        <v-divider></v-divider>
                    </v-col>
                    <v-col cols="12">
                        <v-layout class="mt-10 d-flex justify-center align-center">
                            <v-flex>

                                    <p> Press the button to Run RBI reading G-mail Scheduler</p>
                                    <v-col v-if="this.$vuetify.breakpoint.width >= 800" cols="4">
                                        <v-menu v-model="menu2" :close-on-content-click="false" :nudge-right="40"
                                            transition="scale-transition" offset-y>
                                            <template v-slot:activator="{ on, attrs }">
                                                <v-text-field v-model="date" label="Gmail Date" prepend-inner-icon="mdi-calendar"
                                                    readonly v-bind="attrs" v-on="on" outlined></v-text-field>
                                            </template>
                                            <v-date-picker v-model="date" @input="menu2 = false"
                                                :max="maxDate"></v-date-picker>
                                        </v-menu>
                                    </v-col>
                                    <v-col v-else cols="6">
                                        <v-menu v-model="menu2" :close-on-content-click="false" :nudge-right="40"
                                            transition="scale-transition" offset-y>
                                            <template v-slot:activator="{ on, attrs }">
                                                <v-text-field v-model="date" label="Gmail Date" prepend-inner-icon="mdi-calendar"
                                                    readonly v-bind="attrs" v-on="on" outlined></v-text-field>
                                            </template>
                                            <v-date-picker v-model="date" @input="menu2 = false"
                                                :max="maxDate"></v-date-picker>
                                        </v-menu>
                                    </v-col>
                            </v-flex>
                            <v-flex class="d-flex justify-end align-start">
                                <v-btn @click="GmailReader" small width="150" class="caption text-capitalize elevation-0"
                                    color="blue lighten-4" >
                                    <v-icon size="15"
                                        class="primary--text mr-1">mdi-cached</v-icon>FetchGmailScheduler</v-btn>
                                        <!-- :disabled="date == ''" -->
                            </v-flex>
                        </v-layout>
                    </v-col>
                    <v-col cols="12">
                        <v-divider></v-divider>
                    </v-col>
                </v-row>
            </v-slide-x-transition>
        </div>
    </div>
</template>

<script>
import EventServices from '../../../services/EventServices';
export default {
    data() {
        return {
            allowed: false,
            date: "",
            menu2: false
            // loading: true,
        };
    },
    computed: {
        maxDate() {
            const today = new Date().toISOString().split("T")[0];
            return today;
        },
    },
    methods: {
        async manualFetch() {
            this.$globalData.overlay = true;
            await EventServices.FetchManual()
                .then((response) => {
                    this.$globalData.overlay = false;
                    if (response.data.status == 'S') {
                        this.MessageBar('S', "Data Fetched successfully")
                    } else {
                        this.MessageBar('E', response.data.errMsg)
                    }
                })
                .catch((error) => {
                    this.$globalData.overlay = false;
                    this.MessageBar('E', error)
                })
        },
        ManualOfflineSch() {
            this.$globalData.overlay = true;
            EventServices.ManualOfflineSch()
                .then((response) => {
                    this.$globalData.overlay = false;
                    if (response.data.status == 'S') {
                        this.MessageBar('S', "Schedular run successfully")
                    } else {
                        this.MessageBar('E', response.data.errMsg)
                    }
                })
                .catch((error) => {
                    this.$globalData.overlay = false;
                    this.MessageBar('E', error)
                })
        },

        GmailReader() {
            // console.log(this.date, "Date");
            this.$globalData.overlay = true;
            EventServices.GmailReader(this.date)
                .then((response) => {
                    this.$globalData.overlay = false;
                    if (response.data.status == 'S') {
                        this.MessageBar('S', "Mail Readed successfully")
                    } else {
                        this.MessageBar('E', response.data.errMsg)
                    }
                })
                .catch((error) => {
                    this.$globalData.overlay = false;
                    this.MessageBar('E', error)
                })
        },
        Token() {
            this.$globalData.overlay = true;
            // this.loading = true;
            EventServices.tokenValidation()
                .then((response) => {
                    if (response.data.status != "S") {
                        this.$globalData.overlay = false;
                        // this.loading = false;
                        window.location = this.LoginUrl;
                    } else {
                        EventServices.RouterValidation(this.$route.path)
                            .then((response) => {
                                if (response.data.status != "S") {
                                    this.$globalData.overlay = false;
                                    // this.loading = false;
                                    //   window.location = this.LoginUrl;
                                    this.$router.replace('/ipo')
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
}
</script>
