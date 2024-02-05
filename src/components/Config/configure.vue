<template>
    <v-container class="d-flex justify-center align-center mt-16 flex-column">
        <v-card class="text-right light-blue lighten-5 mb-4" :width="width"
            :elevation="this.$vuetify.breakpoint.name == 'xs' ? 0 : 1">
            <v-card-title class="text">
                {{ title }}
            </v-card-title>
            <v-card-text>
                <div v-if="alert">
                    <AddUser @Recall="recalluser" />
                </div>
                <div class="d-flex justify-center flex-column text-wrap">
                    <v-flex class="d-flex justify-center" v-if="alert2">
                        <v-radio-group v-model="Capture.type" row>
                            <v-radio label="NSE" value="NSE"></v-radio>
                            <v-radio label="BSE" value="BSE"></v-radio>
                        </v-radio-group>
                    </v-flex>
                    <v-flex>
                        <v-form ref="form" v-if="field">
                            <v-layout :class="flex">
                                <v-flex :class="flex">
                                    <v-text-field label="MemberId" :rules="model" v-model="Capture.member" dense
                                        outlined></v-text-field>
                                    <v-text-field label="LoginId" :rules="model" v-model="Capture.loginId" dense outlined
                                        class="mx-md-4"></v-text-field>
                                    <v-text-field :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
                                        :rules="[rules.req, rules.min, rules.max, rules.rule]"
                                        :type="show ? 'text' : 'password'" label="Password" v-model="Capture.password"
                                        class="input-group--focused" @click:append="show = !show" dense
                                        outlined></v-text-field>
                                    <v-btn color="info" @click="submit()" class="mx-3" :disabled="disableBtn">submit</v-btn>
                                </v-flex>
                            </v-layout>
                        </v-form>
                    </v-flex>
                </div>
                <div v-if="alert3" class="d-flex justify-center">
                    <v-radio-group v-model="swap" row>
                        <v-radio label="NSE" value="radio-1"></v-radio>
                        <v-radio label="BSE" value="radio-2"></v-radio>
                        <v-btn class="info elevation-0" small :disabled="swapBtn">confirm</v-btn>
                    </v-radio-group>
                </div>
            </v-card-text>
            <v-card-actions class="justify-end justify-sm-center mx-sm-auto">
                <v-btn @click="showalert(1)" class="info--text mr-2 text-wrap" text>
                    Add User
                </v-btn>
                <v-btn @click="showalert(2)" class="info--text  text-wrap" text>
                    Change password
                </v-btn>
                <v-btn @click="showalert(3)" class="info--text mr-2 text-wrap" text>
                    Change directory
                </v-btn>
            </v-card-actions>
        </v-card>
        <!-- <v-card :elevation="this.$vuetify.breakpoint.name == 'xs' ? 0 : 1" :width="tableWidth">
            <v-data-table no-data-text="No Records available" :headers="header" :items="details" :items-per-page="10"
                :loading="loading" :hide-default-header="this.$vuetify.breakpoint.name == 'xs' ? true : false">
                <template v-slot:item.name="{ item }" style="width: 100%;"
                    v-if="this.$vuetify.breakpoint.name == 'xs' || this.$vuetify.breakpoint.name == 'sm'">
                    <table class="text">
                        <tr>
                            <td xs4>
                                <span>AdminId: {{ item.name }}</span>
                            </td>
                            <td xs4>
                                <span>Role: {{ item.role }}</span>
                            </td>
                            <td xs4>
                                <span>Status: {{ item.status }}</span>
                            </td>
                        </tr>
                    </table>
                </template>
            </v-data-table>
        </v-card> -->
        <NewConfig />
    </v-container>
</template>


<script>
import AddUser from './AddUser/AddUser.vue'
import NewConfig from './ChangePassword/newConfig.vue'
import EventServices from '../../services/EventServices';
export default {
    components: {
        AddUser,
        NewConfig
    },
    data() {
        return {
            title: '',
            swap: '',
            alert: false,
            alert2: false,
            alert3: false,
            field: false,
            show: false,
            password: '',
            rules: {},
            model: [value => !!value || 'Required',],
            Capture: {
                type: '',
                password: ''
            },
            disableBtn: false,
            swapBtn: false,
            details: [],
            header: [
                { text: 'AdminId', align: 'center', value: 'name', class: 'primary lighten-3 white--text' },
                { text: 'Role', align: 'center', value: 'role', class: 'primary lighten-3 white--text' },
                { text: 'Status', align: 'center', value: 'status', class: 'primary lighten-3 white--text' },
            ],
            loading: false
        }
    },
    methods: {
        recalluser() {
            this.GetUserDetails()
        },
        showalert(indicator) {
            if (indicator == 1) {
                this.alert = !this.alert;
                this.alert2 = false;
                this.alert3 = false;
                this.field = false
                this.title = 'AddUser'
            } else if (indicator == 2) {
                this.Capture = {
                    type: '',
                    password: ''
                }
                this.field = false
                this.alert2 = !this.alert2;
                this.alert = false;
                this.alert3 = false;
                this.title = 'Change Password'
            } else {
                this.alert = false;
                this.alert2 = false;
                this.alert3 = !this.alert3;
                this.field = false
                this.title = 'Change Stream'
            }
        },
        submit() {
            this.disableBtn = true
            EventServices.ChangePassword(this.Capture)
                .then((response) => {
                    if (response.data.status == "S") {
                        this.MessageBar('S', response.data.errMsg)
                    } else {
                        this.MessageBar('E', response.data.errMsg)
                    }
                })
                .catch((error) => {
                    this.MessageBar('E', error)
                })
            this.disableBtn = false
            this.field = false
            this.alert2 = false
            this.Capture = {
                type: '',
                member: '',
                loginId: '',
                password: ''
            };
        },
        GetUserDetails() {
            this.loading = true
            EventServices.Getuser()
                .then((response) => {
                    this.loading = false
                    if (response.data.status == "S") {
                        this.details = response.data.adminListArr
                    } else {
                        this.MessageBar('E', response.data.errMsg)
                    }
                })
                .catch((error) => {
                    this.loading = false
                    this.MessageBar('E', error)
                })
        }
    },
    watch: {
        Capture: {
            immediate: true,
            handler(value) {
                if (value.type != "") {
                    if (value.type == 'NSE') {
                        this.rules = {
                            req: value => !!value || 'Required',
                            min: v => v.length >= 12 || 'Min 12 characters',
                            rule: value =>
                                /^(?=.*\d)(?=.*[a-zA-Z])(?=.*[^a-zA-Z0-9]).{12,15}$/.test(value) || 'Password must be like Example@123'
                        }
                        this.field = true;
                    } else {
                        this.rules = {
                            req: value => !!value || 'Required',
                            min: v => v.length >= 8 || 'Min 8 characters',
                            max: v => v.length <= 8 || 'Password does not exist 8 characters',
                            rule: value =>
                                /^(?=.*\d)(?=.*[a-zA-Z])(?=.*[^a-zA-Z0-9]).{8,8}$/.test(value) || 'Password must be like Example@123'
                        }
                        this.field = true;
                    }
                }
                if (value.member != '' && value.loginId != '' && value.password != '') {
                    this.disableBtn = false
                } else {
                    this.disableBtn = true
                }
            }, deep: true
        },
        title: function () {
            if (this.title == '') {
                this.title = 'Welcome'
            }
        },
        swap: {
            immediate: true,
            handler() {
                if (this.swap == '') {
                    this.swapBtn = true
                } else {
                    this.swapBtn = false
                }
            }
        }
    },
    computed: {
        width: {
            get() {
                if (this.$vuetify.breakpoint.name == 'xs') { return 500 }
                else if (this.$vuetify.breakpoint.name == 'sm') { return 600 }
                else if (this.$vuetify.breakpoint.name == 'md') { return 650 }
                else if (this.$vuetify.breakpoint.name == 'lg') { return 700 }
                else if (this.$vuetify.breakpoint.name == 'xl') { return 800 }
                else { return 0 }
            }
        },
        flex: {
            get() {
                if (this.$vuetify.breakpoint.name == 'xs') { return 'd-flex flex-column justify-center' }
                else if (this.$vuetify.breakpoint.name == 'sm') { return 'd-flex justify-end' }
                else if (this.$vuetify.breakpoint.name == 'md') { return 'd-flex justify-end' }
                else if (this.$vuetify.breakpoint.name == 'lg') { return 'd-flex justify-end' }
                else if (this.$vuetify.breakpoint.name == 'xl') { return 'd-flex justify-center' }
                else { return 0 }
            }
        },
        tableWidth: {
            get() {
                if (this.$vuetify.breakpoint.name == 'xs') { return 400 }
                else if (this.$vuetify.breakpoint.name == 'sm') { return 600 }
                else if (this.$vuetify.breakpoint.name == 'md') { return 900 }
                else if (this.$vuetify.breakpoint.name == 'lg') { return 1000 }
                else if (this.$vuetify.breakpoint.name == 'xl') { return 1300 }
                else { return 0 }
            }
        },
    },
    mounted() {
        this.GetUserDetails()
    }
}
</script>

<style scoped>
.text {
    font-size: 10px;
}

@media(max-width:600px) {
    .text-wrap {
        display: flex;
        flex-direction: column;
        font-size: 8px;
    }
}
</style> 