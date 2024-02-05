<template>
    <div fluid name="Landing">
        <PreLogin @GetRedirectUrl="GetRedirectUrl" />
    </div>
</template>

<script>
import PreLogin from "../components/PreLogin/Landing/LandingNovo.vue";
import EventService from "../services/EventServices";
export default {
    components: {
        PreLogin
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
                        // console.log("redirectUrl: " + this.redirectUrl);
                        window.location.href = this.redirectUrl;

                    }
                })
                .catch((error) => {
                    this.MessageBar("E", error);
                });
        },

    },

}
</script>

<style scoped></style>