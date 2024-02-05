<template>
    <v-container>
        <v-slide-y-transition mode="out-in" appear>
            <v-breadcrumbs :items="items">
                <template v-slot:divider>
                    <v-icon>mdi-chevron-right</v-icon>
                </template>
            </v-breadcrumbs>
        </v-slide-y-transition>
        <v-slide-x-transition mode="out-in" appear>
            <LookUpMain  v-if="allowed" :Header="Header" @GetHeaders="GetHeaders" />
        </v-slide-x-transition>
    </v-container>
</template>

<script>
import LookUpMain from '../components/Config/Lookup/LookUpMain.vue';
import EventServices from '@/services/EventServices';
export default {
    components: {
        LookUpMain
    },
    data: () => ({
        items: [
            {
                text: 'Config',
                disabled: false,
            },
            {
                text: 'LookUp',
                disabled: true,
            },
        ],
        Header: [],
        allowed:false,
    }),
    methods: {
        GetHeaders() {
            this.GetLookUpHeader()
        },
        GetLookUpHeader() {
            this.$globalData.overlay = true;
            EventServices.GetLookUpHeader()
                .then((response) => {
                    if (response.data.status == "S") {
                        this.allowed = true;
                        this.$globalData.overlay = false;
                        this.Header = response.data.header;
                    } else {
                        this.allowed = true;
                        this.$globalData.overlay = false;
                        this.MessageBar("E", response.data.errMsg)
                    }
                })
                .catch((error) => {
                    this.allowed = true;
                    this.$globalData.overlay = false;
                    this.MessageBar("E", error)
                });

        }

    },
    mounted() {
        this.GetLookUpHeader()
    }
}
</script>

