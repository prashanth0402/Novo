<template>
    <div>
        <v-container>
            <v-breadcrumbs :items="items">
                <template v-slot:divider>
                    <v-icon>mdi-chevron-right</v-icon>
                </template>
            </v-breadcrumbs>

            <OnboardingDesk v-if="this.$vuetify.breakpoint.width >= 600" :RoleArr="RoleArr" :AddDomain="AddDomain"
                :AddBroker="AddBroker" :BrokerArr="BrokerArr" @RecallApi="RecallApi" />
            <OnboardingMob v-else :RoleArr="RoleArr" :AddDomain="AddDomain" :AddBroker="AddBroker" :BrokerArr="BrokerArr" />
        </v-container>
    </div>
</template>

<script>
import EventServices from '../../services/EventServices';
import OnboardingDesk from './OnBoarding/OnboardingDesk.vue';
import OnboardingMob from './OnBoarding/OnboardingMob.vue';
export default {
    components: {
        OnboardingDesk,
        OnboardingMob
    },
    data: () => ({
        items: [
            {
                text: 'Setup',
                disabled: false,
            },
            {
                text: 'Domain Setup',
                disabled: true,
            },
        ],
        loading: false,
        AddDomain: [],
        AddBroker: [],
        RoleArr: [],
        BrokerArr: []
    }),
    methods: {
        RecallApi() {
            this.GetDomainData()
            this.GetAdminData()
        },
        GetDomainData() {
            this.loading = true
            this.$globalData.overlay = true
            EventServices.GetDomainList()
                .then((response) => {
                    this.loading = false
                    this.$globalData.overlay = false
                    if (response.data.status == "S") {
                        // this.MessageBar("S", response.data.errMsg)
                        this.AddDomain = response.data.brokerListArr;
                    } else {
                        this.loading = false
                        this.MessageBar("E", response.data.errMsg);
                    }
                })
                .catch((error) => {
                    this.loading = false
                    this.$globalData.overlay = false
                    this.MessageBar("E", error);
                });
        },
        GetAdminData() {
            EventServices.GetBrokerList()
                .then((response) => {
                    if (response.data.status == "S") {
                        this.loading = false
                        this.$globalData.overlay = false
                        response.data.adminListArr != null ? this.AddBroker = response.data.adminListArr : this.AddBroker = []
                        response.data.roleListArr != null ? this.RoleArr = response.data.roleListArr : this.RoleArr = []
                        response.data.brokerNameArr != null ? this.BrokerArr = response.data.brokerNameArr : this.BrokerArr = []

                        // this.AddBroker = response.data.adminListArr != null ? response.data.adminListArr : []
                    } else {
                        this.loading = false
                        this.$globalData.overlay = false
                        this.MessageBar("E", response.data.errMsg);
                    }
                })
                .catch((error) => {
                    this.loading = false
                    this.MessageBar("E", error);
                });

        }
    },
    created() {
        this.GetAdminData()
        this.GetDomainData()
    }
}
</script>