<template>
    <div class="container">
        <v-tabs v-model="activeTab">
            <v-tab> <v-icon left color="blue darken-4" medium> mdi-link-variant </v-icon> <span
                    class="text-subtitle-1 text-capitalize">
                    Ipo Link Detail </span></v-tab>
            <v-tab> <v-icon left color="blue darken-4" medium> mdi-registered-trademark </v-icon><span
                    class="text-subtitle-1 text-capitalize">
                    RegistrarList </span></v-tab>
        </v-tabs>

        <div v-if="activeTab == 0">
            <DesktopView v-if="this.$vuetify.breakpoint.width >= 823" />
            <MobileView v-else />
        </div>
        <div v-else>
            <IpoRegistrarDetails v-if="this.$vuetify.breakpoint.width >= 823"></IpoRegistrarDetails>
            <IpoRegistrarMobileView v-else></IpoRegistrarMobileView>
        </div>
    </div>
</template>

<script>
import MobileView from './View/MobileView.vue';
import DesktopView from './View/DesktopView.vue';
import IpoRegistrarDetails from './IpoRegistrarDetails/IpoRegistrarDetails.vue';
import IpoRegistrarMobileView from './IpoRegistrarDetails/IpoRegistrarMobileView.vue';
export default {
    components: {
        MobileView, DesktopView,
        IpoRegistrarDetails,
        IpoRegistrarMobileView
    },
    data() {
        return {
            activeTab: 0,
            tab: ["IPOLINK", "IPOREGISTRAR"]
        }
    }
}
</script>