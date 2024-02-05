<template>
    <div>
            <VersionController  v-if="allowed" class="mt-10"></VersionController>
    </div>
</template>
<script>
import EventServices from "@/services/EventServices";
import VersionController from "../components/versioncontrol/versionController.vue";

export default {
    components: {
    VersionController
},
    data() {
        return {
            allowed: false,
            loading: true,
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
                                    //   window.location = this.LoginUrl;
                                    this.$router.replace(this.$globalData.links[0].path)
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
                    window.location = this.LoginUrl;
                });
        },
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
    }

}
</script>
