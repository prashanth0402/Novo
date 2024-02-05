<template>
    <div>
        <LandingSgb />
    </div>
</template>

<script>
import LandingSgb from "../components/PreLogin/SgbLanding/LandingSgb.vue"
import EventService from "../services/EventServices";
export default {
    components: {
        LandingSgb
    },
    methods: {
        async GetRedirectUrl() {
            await EventService.GetRedirectURL()
                .then((response) => {
                    if (response.data.status == "S") {

                        this.$globalData.host = response.data.redirectUrl.host
                        this.$globalData.appName = response.data.redirectUrl.appName
                        this.$globalData.url = response.data.redirectUrl.url

                        this.AssignUrl();

                        // this.redirectUrl = this.redirectUrl + this.$globalData.url
                        // console.log("Redirect: " + this.redirectUrl, "Url: " + this.$globalData.url);

                        window.location.href = this.redirectUrl;

                    }
                })
                .catch((error) => {
                    this.MessageBar("E", error);
                });
        },

    }
}
</script>

<style scoped></style>