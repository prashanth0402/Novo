<!-- <template>
    <div class="d-flex align-center justify-center flex-column ">
        <v-card class="elevation-0 light-blue lighten-5" xl12 lg12>
            <v-form ref="form" lazy-validation class="d-flex mt-4">
                <v-layout :class="flex">
                    <v-flex class="mr-4" xl3 md3 sm2>
                        <v-text-field v-model="add.clientId" label="UserId" dense required :rules="[v => !!v || 'required']"
                            :error="Err" @input="validation"></v-text-field>
                    </v-flex>
                    <v-flex class="mr-4" xl3 md2 sm2>
                        <v-autocomplete :items="items" v-model="add.role" item-text="role" item-value="roleId" dense label="Role" required
                            :rules="[v => !!v || 'required']" :error="Err">
                        </v-autocomplete>
                    </v-flex>
                    <v-flex xl4 md6 sm7>
                        <v-radio-group v-model="add.flag" row>
                            <v-radio label="Active" color="primary" value="Active" required class="mr-5">
                            </v-radio>
                            <v-radio label="Deactive" color="error" value="DeActive" required>
                            </v-radio>
                        </v-radio-group>
                    </v-flex>
                    <v-flex xl2 md2 sm1 class="mr-sm-10">
                        <v-btn block class="primary lighten-2 elevation-0" @click="addUser" :loading="btnLoad"
                            :disabled="btnLoad">ADD</v-btn>
                    </v-flex>
                </v-layout>
            </v-form>
        </v-card>
    </div>
</template>

<script>
import EventServices from "../../../services/EventServices"
export default {
    data() {
        return {
            valid: false,
            add: {
                clientId: "",
                role: "",
                flag: ""
            },
            items: [],
            btnLoad: false,
            Err: false,
        }
    },
    methods: {
        async addUser() {
            this.btnLoad = true;
            await this.validation()
            if (this.Err == false) {
                await EventServices.Adduser(this.add)
                    .then((response) => {
                        if (response.data.status == "S") {
                            this.btnLoad = false;
                            this.MessageBar('S', "User added successfully")
                            this.$refs.form.reset();
                            this.$emit('Recall')
                        } else {
                            this.btnLoad = false;
                            this.MessageBar('E', "Unable to add user right now")
                        }
                    })
                    .catch((error) => {
                        this.btnLoad = false;
                        this.MessageBar('E', error)
                    })
            } else {
                this.btnLoad = false;
                this.Err = true;
                this.SnackBar('E', "Fields cannot be empty")
            }
        },
        validation() {
            if (this.add.clientId != "" && this.add.role != "" && this.add.flag != "") {
                this.Err = false;
            } else {
                this.Err = true;
            }
        },
        GetUserDetails() {
            EventServices.Getuser()
                .then((response) => {
                    this.loading = false
                    if (response.data.status == "S") {
                        this.items = response.data.adminListArr
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
    computed: {
        flex: {
            get() {
                if (this.$vuetify.breakpoint.name == 'xs') { return 'd-flex flex-column justify-center' }
                else if (this.$vuetify.breakpoint.name == 'sm') { return 'd-flex justify-end align-center' }
                else if (this.$vuetify.breakpoint.name == 'md') { return 'd-flex justify-end align-center' }
                else if (this.$vuetify.breakpoint.name == 'lg') { return 'd-flex justify-end align-center' }
                else if (this.$vuetify.breakpoint.name == 'xl') { return 'd-flex justify-end align-center' }
                else { return 0 }
            }
        },
    },
    mounted() {
        this.GetUserDetails()
    }
}
</script> -->