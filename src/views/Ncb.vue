<template>
    <div>
        <v-progress-linear indeterminate color="primary" v-if="loading"></v-progress-linear>
        <NcbMain v-if="allowed" class="mt-10" />
    </div>
</template>

<script>
import NcbMain from "../components/Ncb/NcbMain.vue"
import EventServices from "../services/EventServices";
export default {
    name: "NCB",
    components: {
        NcbMain,
    },
    data() {
        return {
            allowed: true,
            loading: false,
        };
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
                                   
                                    this.$router.replace(this.$globalData.links[0].path)
                                } else {
                              
                                    this.$globalData.overlay = false;
                                    this.loading = false;
                                    this.allowed = true;
                                }
                            })
                            .catch((errMsg) => {
                                this.$globalData.overlay = false;
                                this.loading = false;
                                this.MessageBar("E",  errMsg);
                            });
                    }
                })
                .catch(() => {
                    this.$globalData.overlay = false;
                    this.loading = false;
                    window.location = this.LoginUrl;
                });
        },
    },
}
</script>

<style scoped></style>